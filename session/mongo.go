package session

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const indexModified = "modified"

// Mongo implements session.Session interface
type Mongo struct {
	mongoSession   *mgo.Session
	databaseName   string
	collectionName string
	expireDuration time.Duration
	logFunc        func(format string, args ...interface{}) // inject for logFuncging
}

// NewMongo returns session.Session interface implemented using mongo
func NewMongo(
	session *mgo.Session,
	databaseName string,
	collectionName string,
	expireDuration time.Duration,
	logFunc func(format string, args ...interface{}),
) Session {
	conn := session.Clone()
	defer conn.Close()
	c := conn.DB(databaseName).C(collectionName)
	c.EnsureIndex(mgo.Index{
		Key:         []string{indexModified},
		Background:  true,
		Sparse:      true,
		ExpireAfter: expireDuration,
	})
	if logFunc == nil {
		logFunc = log.Printf
	}
	return Mongo{
		mongoSession:   session,
		databaseName:   databaseName,
		collectionName: collectionName,
		expireDuration: expireDuration,
		logFunc:        logFunc,
	}
}

// for functions below, Mongo.logFunc will logFunc all errors

// Get returns the value according to token & key
func (m Mongo) Get(token string, key string) interface{} {
	conn := m.mongoSession.Clone()
	defer conn.Close()
	c := conn.DB(m.databaseName).C(m.collectionName)
	data, err := m.findID(c, token)
	if err != nil {
		return nil
	}
	return data[key]
}

// Set set value in mongodb with token & key
func (m Mongo) Set(token string, key string, val interface{}) error {
	conn := m.mongoSession.Clone()
	defer conn.Close()
	c := conn.DB(m.databaseName).C(m.collectionName)
	data, err := m.findID(c, token)
	if err != nil {
		return err
	}

	data[key] = val
	data[indexModified] = time.Now()

	// update if exist, else insert
	if _, ok := data["_id"]; ok {
		if err := c.UpdateId(token, data); err != nil {
			m.logFunc("mongo cant update document with token %s: %v", token, err)
			return err
		}
	} else {
		data["_id"] = token
		if err := c.Insert(data); err != nil {
			m.logFunc("mongo cant insert document with token %s: %v", token, err)
			return err
		}
	}

	return nil
}

// Delete delete a key from mongo according to token
func (m Mongo) Delete(token string, key string) {
	conn := m.mongoSession.Clone()
	defer conn.Close()
	c := conn.DB(m.databaseName).C(m.collectionName)
	data, err := m.findID(c, token)
	if err != nil {
		return
	}

	// if not exist, do nothing
	if _, ok := data["_id"]; !ok {
		return
	}

	// update document
	delete(data, key)
	if err := c.UpdateId(token, data); err != nil {
		m.logFunc("mongo cant update document with token %s: %v", token, err)
	}
}

// Expire makes a token expired
func (m Mongo) Expire(token string) {
	conn := m.mongoSession.Clone()
	defer conn.Close()
	c := conn.DB(m.databaseName).C(m.collectionName)
	if err := c.RemoveId(token); err != nil {
		m.logFunc("mongo cant remove document with token %s: %v", token, err)
	}
}

// help functions
func (m Mongo) findID(c *mgo.Collection, token string) (bson.M, error) {
	data := bson.M{}
	err := c.FindId(token).One(&data)
	switch err {
	case nil, mgo.ErrNotFound:
		err = nil
	default:
		m.logFunc("mongo cant query document with token %s: %v", token, err)
	}
	return data, err
}
