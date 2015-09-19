package handler

import (
	"github.com/insionng/vodka"
	"github.com/insionng/zenpress/helper"
	"github.com/insionng/zenpress/models"
	"github.com/vodka-contrib/sessions"
	"net/http"
)

func SigninGetHandler(self *vodka.Context) error {

	session := sessions.Default(self)

	val := session.Get("user")
	if val != nil {
		return self.Redirect(302, "/")
	} else {
		return self.Render(http.StatusOK, "signin.html", nil)
	}

}

func SigninPostHandler(self *vodka.Context) error {
	username := self.FormEscape("username")
	password := self.FormEscape("password")

	if username != "" && password != "" {

		if userInfo := models.GetUserByNickname(username); userInfo.Password != "" {

			if helper.ValidateHash(userInfo.Password, password) {
				session := sessions.Default(self)
				session.Set("user", userInfo)

				if e := session.Save(); e == nil {
					return self.Redirect(302, "/")
				}

			}
		}
	}
	return self.Redirect(302, "/signin/")
}
