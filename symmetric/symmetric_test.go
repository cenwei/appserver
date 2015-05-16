package symmetric

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func testSymmetric(sa Symmetric, t *testing.T) {
	Convey("given a symmetric algorithm", t, func() {

		Convey("with same key, decrypted data should equals to original data", func() {
			key := []byte("something-very-safe")
			originalText := []byte("original text")
			cipher := sa.Encrypt(originalText, key)
			decryptedText := sa.Decrypt(cipher, key)

			So(decryptedText, ShouldResemble, originalText)
		})

		Convey("with two different key, decrypted data should not equal to original data", func() {
			key := []byte("something-very-safe")
			otherKey := []byte("other-key")
			originalText := []byte("original text")
			cipher := sa.Encrypt(originalText, key)
			decryptedText := sa.Decrypt(cipher, otherKey)

			So(decryptedText, ShouldNotResemble, originalText)
		})
	})
}
