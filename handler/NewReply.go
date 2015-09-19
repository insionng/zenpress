package handler

import (
	"fmt"
	"github.com/insionng/vodka"
	"github.com/insionng/zenpress/models"
	"github.com/vodka-contrib/sessions"
	"net/url"
)

func NewReplyPostHandler(self *vodka.Context) error {
	tid := self.FormInt64("comment_parent")

	session := sessions.Default(self)
	var user models.User
	val := session.Get("user")
	if val != nil {
		user = val.(models.User)
	}

	sess_userid := user.Id

	gmt, _ := self.Request().Cookie("gmt")

	if gmt != nil {

		if _, err := url.QueryUnescape(gmt.Value); err == nil {

			author := self.FormEscape("author")
			email := self.FormEscape("email")
			website := self.FormEscape("website")
			rc := self.FormEscape("comment")

			if author != "" && email != "" && tid != 0 && rc != "" {
				if err := models.AddReply(tid, sess_userid, rc, author, email, website); err != nil {
					return err
				}
				return self.Redirect(302, fmt.Sprintf("/view/", tid))
			} else {
				return self.Redirect(302, "/")
			}
		} else {
			return self.Redirect(302, "/")
		}

	} else {
		return self.Redirect(302, "/")
	}

}
