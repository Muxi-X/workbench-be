package auth

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAuthRequest(t *testing.T) {
	InitVar()
	var (
		code    = "XEYDGF0ZPH-SRY2P7O5VRW"
		token   string
		refresh string

		clientID     = "4b194ad8-7d97-4dca-b078-6c3c65b31c75"
		clientSecret = "8c066b19-e507-4887-88f3-7e7edd99bfd8"
	)

	Convey("Test register", t, func() {
		err := RegisterRequest("muxi101", "muxi@123.com", "muxi")
		So(err, ShouldBeNil)
	})

	Convey("Test getting token", t, func() {
		t, err := GetTokenRequest(code, clientID, clientSecret)
		So(err, ShouldBeNil)

		token = t.AccessToken
		refresh = t.RefreshToken
	})

	Convey("Test refreshing token", t, func() {
		t, err := RefreshTokenRequest(refresh, clientID, clientSecret)
		So(err, ShouldBeNil)
		So(t.AccessToken, ShouldNotEqual, token)

		token = t.AccessToken
	})

	Convey("Test getting user info", t, func() {
		u, err := GetInfoRequest(token)
		So(err, ShouldBeNil)
		So(u, ShouldNotBeNil)
	})
}
