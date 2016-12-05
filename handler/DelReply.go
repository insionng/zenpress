package handler

import (
	"github.com/insionng/macross"
	"github.com/insionng/zenpress/models"
)

func DeleteReplyHandler(self *macross.Context) error {
	rid := self.Param("rid").MustInt64()
	models.DelReply(rid)
	return self.Redirect("/", 302)
}
