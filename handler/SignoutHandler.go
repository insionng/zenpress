package handler

import (
	"github.com/insionng/vodka"
	"github.com/vodka-contrib/session"
)

func SignoutHandler(self vodka.Context) error {
	sess := session.GetStore(self)
	sess.Delete("user")
	//return self.JSON( 200, "okay")
	return self.Redirect(302, "/")
}
