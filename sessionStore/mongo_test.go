package sessionStore

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
		mongo := NewMongoSessionStore(session, "mgo", "token_store", 5*time.Second, t.Logf)
		mongo.Expire("token")
		mongo.Expire("other-token")

		Convey("When get data", func() {
			data := mongo.Get("token", "key")

			Convey("The data should be nil", func() {
				So(data, ShouldBeNil)
			})
		})
	})

	Convey("Given a mongo collection with data", t, func() {
		mongo := NewMongoSessionStore(session, "mgo", "token_store", 5*time.Second, t.Logf)
		mongo.Set("token", "key", "random")
		mongo.Set("other-token", "key", "random")

		Convey("When get data", func() {
			data := mongo.Get("token", "key")

			Convey("The data should be equal to \"random\"", func() {
				So(data, ShouldEqual, "random")
			})
		})

		Convey("When get data after correponding key deleted", func() {
			mongo.Delete("token", "key")
			data := mongo.Get("token", "key")

			Convey("The data should be nil", func() {
				So(data, ShouldBeNil)
			})
		})

		Convey("When get data after document is deleted", func() {
			mongo.Expire("other-token")
			data := mongo.Get("other-token", "key")

			Convey("The data should be nil", func() {
				So(data, ShouldBeNil)
			})
		})
	})
}
