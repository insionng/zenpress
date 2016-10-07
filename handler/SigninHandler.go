package handler

import (
	"net/http"

	"github.com/insionng/vodka"
	"github.com/insionng/zenpress/helper"
	"github.com/insionng/zenpress/models"
	"github.com/vodka-contrib/session"
)

func SigninGetHandler(self vodka.Context) error {

	sess := session.GetStore(self)

	val := sess.Get("user")
	if val != nil {
		return self.Redirect(302, "/")
	} else {
		return self.Render(http.StatusOK, "signin.html", nil)
	}

}

func SigninPostHandler(self vodka.Context) error {
	username := self.FormValue("username")
	password := self.FormValue("password")

	if username != "" && password != "" {

		if userInfo := models.GetUserByNickname(username); userInfo.Password != "" {

			if helper.ValidateHash(userInfo.Password, password) {
				sess := session.GetStore(self)

				err := sess.Set("user", userInfo)
				if err == nil {
					return self.Redirect(302, "/")
				}

			}
		}
	}
	return self.Redirect(302, "/signin/")
}
