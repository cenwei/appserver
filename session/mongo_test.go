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

	Convey("given a clean mongo collection", t, func() {
		mongo := NewMongoSessionStore(session, "mgo", "token_store", 5*time.Second, t.Logf)

		Convey("get data from fresh mongo", func() {
			So(mongo.Get("token", "key"), ShouldBeNil)
		})

		Convey("given a mongo collection with data", func() {
			mongo.Set("token", "key", "random")

			Convey("should get the same data", func() {
				So(mongo.Get("token", "key"), ShouldEqual, "random")
			})
		})

		Convey("delete corresponding key", func() {
			mongo.Delete("token", "key")

			Convey("should not get the data", func() {
				So(mongo.Get("token", "key"), ShouldBeNil)
			})
		})

		Convey("insert corresponding key and delete the collection", func() {
			mongo.Set("token", "key", "random")
			mongo.Expire("token")

			Convey("should not get the data", func() {
				So(mongo.Get("token", "key"), ShouldBeNil)
			})
		})
	})
}
