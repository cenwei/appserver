package symmetric

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func testSymmetric(sa Symmetric, t *testing.T) {
	Convey("Given a symmetric algorithm", t, func() {

		Convey("When encrypt & decrypt with same key", func() {
			key := []byte("something-very-safe")
			originalText := []byte("original text")
			cipher := sa.Encrypt(originalText, key)
			decryptedText := sa.Decrypt(cipher, key)

			Convey("decrypted data should be equal to original data", func() {
				So(decryptedText, ShouldResemble, originalText)
			})
		})

		Convey("When encrypt & decrypt with different key", func() {
			key := []byte("something-very-safe")
			otherKey := []byte("other-key")
			originalText := []byte("original text")
			cipher := sa.Encrypt(originalText, key)
			decryptedText := sa.Decrypt(cipher, otherKey)

			Convey("decrypted data should not be equal to original data", func() {
				So(decryptedText, ShouldNotResemble, originalText)
			})
		})
	})
}
