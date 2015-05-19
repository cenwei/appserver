package authtoken

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Token is an access token associated to userid & symmetric-algorithm key
type Token struct {
	AccessToken interface{} `bson:"_id"`
	UserID      interface{}
	PrivateKey  []byte
	Modified    time.Time `bson:"modified"`
}

// NewToken creates a new token ready for use with userID & privateKey
func NewToken(userID int64, privateKey []byte) Token {
	return Token{
		AccessToken: bson.NewObjectId(),
		UserID:      userID,
		PrivateKey:  privateKey,
		Modified:    time.Now(),
	}
}

// TokenStore represents a persistent storage for Token
type TokenStore interface {
	Get(accessToken string) (Token, error)
	Set(token Token) error
	Expire(accessToken string) error
}
