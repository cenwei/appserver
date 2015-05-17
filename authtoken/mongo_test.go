package authtoken

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/mgo.v2"
)

func TestMongo(t *testing.T) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		t.Logf("no mongodb service, stop testing: %v", err)
		return
	}

	Convey("Given a fresh mongo collection", t, func() {
		mongo := NewMongoStore(session, "mgo", "token_store", 5*time.Second)
		mongo.Expire("token")
		mongo.Expire("other-token")

		Convey("When get data", func() {
			_, err := mongo.Get("token")

			Convey("The error should be mgo.ErrNotFound", func() {
				So(IsMongoErrorNotFound(err), ShouldBeTrue)
			})
		})
	})

	Convey("Given a mongo collection with data", t, func() {
		mongo := NewMongoStore(session, "mgo", "token_store", 5*time.Second)
		t := time.Now()
		token := Token{
			AccessToken: "token",
			PrivateKey:  []byte("key"),
			UserID:      1,
			Modified:    t,
		}
		mongo.Set(token)

		Convey("When get token", func() {
			tokenTmp, _ := mongo.Get("token")
			// should set time otherwise test wouldn't pass because mondb store time in us level
			tokenTmp.Modified = t

			Convey("The tokenTmp should be resemble to token", func() {
				So(tokenTmp, ShouldResemble, token)
			})
		})

		Convey("When get token after document is deleted", func() {
			mongo.Expire("token")
			_, err := mongo.Get("token")

			Convey("The error should be mgo.ErrNotFound", func() {
				So(IsMongoErrorNotFound(err), ShouldBeTrue)
			})
		})
	})
}
