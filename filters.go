package main

import (
	"github.com/hprose/hprose-go/hprose"
	"github.com/sharelog/appserver/session"
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
	sessionKey        string // the key name in session
	getter            session.Getter
}

func getTokenFromContext(headerAccessToken string, context hprose.Context) string {
	ctx, ok := context.(*hprose.HttpContext)
	if ok {
		return ctx.Request.Header.Get(headerAccessToken)
	}
	return ""
}

func (f symmetricEncryption) filterFunc(data []byte, context hprose.Context, isInput bool) []byte {
	token := getTokenFromContext(f.headerAccessToken, context)
	if token != "" {
		if key, ok := f.getter.Get(token, f.sessionKey).([]byte); ok && key != nil {
			if isInput {
				return f.symmetric.Decrypt(data, key)
			}
			return f.symmetric.Encrypt(data, key)
		}
	}
	return data
}

func (f symmetricEncryption) InputFilter(data []byte, context hprose.Context) []byte {
	return f.filterFunc(data, context, true)
}

func (f symmetricEncryption) OutputFilter(data []byte, context hprose.Context) []byte {
	return f.filterFunc(data, context, false)
}
