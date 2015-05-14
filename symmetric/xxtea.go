package symmetric

import "github.com/xxtea/xxtea-go/xxtea"

var _ Symmetric = XXTEA{} // compile time interface checker

// XXTEA implements symmetric-algorithm interface using xxtea
type XXTEA struct{}

// Encrypt implements Symmetric.Encrypt
func (XXTEA) Encrypt(data, key []byte) []byte { return xxtea.Encrypt(data, key) }

// Decrypt implements Symmetric.Decrypt
func (XXTEA) Decrypt(data, key []byte) []byte { return xxtea.Decrypt(data, key) }
