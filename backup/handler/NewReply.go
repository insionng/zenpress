package handler

import (
	"fmt"
	"github.com/insionng/macross"
	"github.com/insionng/macross/jwt"
	"github.com/insionng/zenpress/models"
)

func NewReplyPostHandler(self *macross.Context) error {
	tid := self.Args("comment_parent").MustInt64()

	claims := jwt.GetMapClaims(self)
	var uid int64
	if jwtUserId, okay := claims["UserId"].(float64); okay {
		uid = int64(jwtUserId)
	}
	var author string
	if jwtUsername, okay := claims["Username"].(string); okay {
		author = jwtUsername
	}

	email := self.Args("email").String()
	website := self.Args("website").String()
	rc := self.Args("comment").String()

	if author != "" && email != "" && tid != 0 && rc != "" {
		if err := models.AddReply(tid, uid, rc, author, email, website); err != nil {
			return err
		}
		return self.Redirect(fmt.Sprintf("/view/", tid), macross.StatusFound)
	}
	return self.Redirect("/", macross.StatusFound)

}
