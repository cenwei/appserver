package main

import (
	"github.com/hprose/hprose-go/hprose"
	"github.com/xxtea/xxtea-go/xxtea"
)

type filterXXTEA struct{ key []byte }

func (f filterXXTEA) InputFilter(data []byte, _ hprose.Context) []byte {
	return xxtea.Decrypt(data, f.key)
}

func (f filterXXTEA) OutputFilter(data []byte, _ hprose.Context) []byte {
	return xxtea.Encrypt(data, f.key)
}
