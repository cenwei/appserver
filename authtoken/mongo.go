package authtoken

import (
	"time"

	"gopkg.in/mgo.v2"
)

// MongoStore implements Store interface
type MongoStore struct {
	mongoSession   *mgo.Session
	databaseName   string
	collectionName string
}

// NewMongoStore returns Store interface implemented by mongo
func NewMongoStore(
	mongoSession *mgo.Session,
	databaseName string,
	collectionName string,
	expireDuration time.Duration,
) TokenStore {
	conn := mongoSession.Clone()
	defer conn.Close()
	c := conn.DB(databaseName).C(collectionName)
	c.EnsureIndex(mgo.Index{
		Key:         []string{"modified"},
		Background:  true,
		Sparse:      true,
		ExpireAfter: expireDuration,
	})
	return MongoStore{
		mongoSession:   mongoSession,
		databaseName:   databaseName,
		collectionName: collectionName,
	}
}

// Get reads the specified access token from mongodb
func (m MongoStore) Get(accessToken string) (Token, error) {
	conn := m.mongoSession.Clone()
	defer conn.Close()
	c := conn.DB(m.databaseName).C(m.collectionName)

	data := Token{}
	err := c.FindId(accessToken).One(&data)
	return data, err
}

// Set writes the specified key-value into mongodb, or update if there exist any
func (m MongoStore) Set(token Token) error {
	conn := m.mongoSession.Clone()
	defer conn.Close()
	c := conn.DB(m.databaseName).C(m.collectionName)

	_, err := c.UpsertId(token.AccessToken, token)
	return err
}

// Expire removes a document from mongodb if it exists
func (m MongoStore) Expire(accessToken string) error {
	conn := m.mongoSession.Clone()
	defer conn.Close()
	c := conn.DB(m.databaseName).C(m.collectionName)

	return c.RemoveId(accessToken)
}

// IsMongoErrorNotFound checks if error == mgo.ErrNotFound
func IsMongoErrorNotFound(err error) bool {
	return mgo.ErrNotFound == err
}
