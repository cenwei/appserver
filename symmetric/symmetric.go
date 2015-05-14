package symmetric

import (
	"reflect"
	"testing"
)

// Symmetric abstracts the symmetric-algorithm interface
type Symmetric interface {
	Encrypt(data []byte, key []byte) []byte
	Decrypt(data []byte, key []byte) []byte
}

func testSymmetric(sa Symmetric, t *testing.T) {
	key := []byte("something-very-safe")
	original := []byte("original text")
	cipher := sa.Encrypt(original, key)

	if !reflect.DeepEqual(original, sa.Decrypt(cipher, key)) {
		t.Error("cant get original text")
	}
}
