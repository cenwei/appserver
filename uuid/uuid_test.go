package uuid

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"regexp"
)

func TestUUID(t *testing.T) {
	Convey("Given a new uuid v4", t, func() {
		uuid := New()

		Convey("It should be in format xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxx", func() {
			const uuid4Pattern = `[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}`
			matched, _ := regexp.MatchString(uuid4Pattern, uuid)
			So(matched, ShouldBeTrue)
		})
	})
}
