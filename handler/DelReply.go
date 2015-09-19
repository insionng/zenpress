package handler

import (
	"github.com/insionng/vodka"
	"github.com/insionng/zenpress/models"
)

func DeleteReplyHandler(self *vodka.Context) error {
	rid, _ := self.ParamInt64(":rid")
	models.DelReply(rid)
	return self.Redirect(302, "/")
}
