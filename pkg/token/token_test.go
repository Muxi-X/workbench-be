package token

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestToken(t *testing.T) {
	var token string
	var id uint32 = 2020

	Convey("Test token", t, func() {
		Convey("Test token generation", func() {
			var err error
			token, err = GenerateToken(TokenPayload{
				ID:      id,
				Expired: time.Hour * 24,
			})
			So(err, ShouldBeNil)
		})

		Convey("Test token resolution", func() {
			curID, err := ResolveToken(token)
			So(err, ShouldBeNil)
			So(curID, ShouldEqual, id)
		})
	})
}
