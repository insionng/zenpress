package handler

import (
	"github.com/insionng/vodka"
	"github.com/vodka-contrib/sessions"
)

func SignoutHandler(self *vodka.Context) error {
	session := sessions.Default(self)
	session.Delete("user")
	//return self.JSON( 200, "okay")
	return self.Redirect(302, "/")
}
