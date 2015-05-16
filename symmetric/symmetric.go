package symmetric

// Symmetric abstracts the symmetric-algorithm interface
type Symmetric interface {
	Encrypt(data []byte, key []byte) []byte
	Decrypt(data []byte, key []byte) []byte
}
