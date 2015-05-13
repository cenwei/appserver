package main

import (
	"github.com/hprose/hprose-go/hprose"
	"github.com/sharelog/appserver/session"
	"github.com/xxtea/xxtea-go/xxtea"
)

// xxtea filter
// get token from hprose context by headerAccessToken
// get xxtea private key from session by token & sessionXXTEA
// encrypt and decrypt data using xxtea private key which is unique for clients
type xxteaFilter struct {
	getter            session.Getter
	headerAccessToken string // the token header name
	sessionXXTEA      string // the xxtea key name in map
}

func getTokenFromContext(headerAccessToken string, context hprose.Context) string {
	ctx, ok := context.(*hprose.HttpContext)
	if ok {
		return ctx.Request.Header.Get(headerAccessToken)
	}
	return ""
}

func (f xxteaFilter) filterFunc(data []byte, context hprose.Context, isInput bool) []byte {
	token := getTokenFromContext(f.headerAccessToken, context)
	if token != "" {
		if key, ok := f.getter.Get(token, f.sessionXXTEA).([]byte); ok && key != nil {
			if isInput {
				return xxtea.Decrypt(data, key)
			}
			return xxtea.Encrypt(data, key)
		}
	}
	return data
}

func (f xxteaFilter) InputFilter(data []byte, context hprose.Context) []byte {
	return f.filterFunc(data, context, true)
}

func (f xxteaFilter) OutputFilter(data []byte, context hprose.Context) []byte {
	return f.filterFunc(data, context, false)
}
