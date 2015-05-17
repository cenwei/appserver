package main

import (
	"github.com/hprose/hprose-go/hprose"
	"github.com/sharelog/appserver/authtoken"
	"github.com/sharelog/appserver/log"
	"github.com/sharelog/appserver/symmetric"
)

var _ hprose.Filter = symmetricEncryption{}

// symmetric encryption implements hprose.Filter
// workflow:
// get token from hprose context by headerAccessToken
// get private key from session by token & sessionKey
// encrypt and decrypt data using private key which is unique for clients
type symmetricEncryption struct {
	symmetric         symmetric.Symmetric
	headerAccessToken string // the token header name
	tokenStore        authtoken.TokenStore
}

func getTokenFromContext(headerAccessToken string, context hprose.Context) string {
	ctx, ok := context.(*hprose.HttpContext)
	if ok {
		return ctx.Request.Header.Get(headerAccessToken)
	}
	return ""
}

func (f symmetricEncryption) filterFunc(data []byte, context hprose.Context, isInput bool) []byte {
	accessToken := getTokenFromContext(f.headerAccessToken, context)
	if accessToken != "" {
		token, err := f.tokenStore.Get(accessToken)
		switch {
		case authtoken.IsMongoErrorNotFound(err):
			panic("Invalid access token, please login first")
		case err != nil:
			log.Errorf("mongodb error: %v", err)
			panic("Server error")
		}
		if isInput {
			return f.symmetric.Decrypt(data, token.PrivateKey)
		}
		return f.symmetric.Encrypt(data, token.PrivateKey)
	}
	return data
}

func (f symmetricEncryption) InputFilter(data []byte, context hprose.Context) []byte {
	return f.filterFunc(data, context, true)
}

func (f symmetricEncryption) OutputFilter(data []byte, context hprose.Context) []byte {
	return f.filterFunc(data, context, false)
}
