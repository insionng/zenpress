package handler

import (
	"fmt"
	"net/url"

	"github.com/Unknwon/com"
	"github.com/insionng/vodka"
	"github.com/insionng/zenpress/models"
	"github.com/vodka-contrib/session"
)

func NewReplyPostHandler(self vodka.Context) error {
	tid := com.StrTo(self.FormValue("comment_parent")).MustInt64()

	sess := session.GetStore(self)
	var user models.User
	val := sess.Get("user")
	if val != nil {
		user = val.(models.User)
	}

	sess_userid := user.Id

	gmt, _ := self.Request().Cookie("gmt")

	if gmt != nil {
		gv := gmt.Value()

		if _, err := url.QueryUnescape(gv); err == nil {

			author := self.FormValue("author")
			email := self.FormValue("email")
			website := self.FormValue("website")
			rc := self.FormValue("comment")

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
