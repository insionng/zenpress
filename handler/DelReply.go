package handler

import (
	"github.com/Unknwon/com"
	"github.com/insionng/vodka"
	"github.com/insionng/zenpress/models"
)

func DeleteReplyHandler(self vodka.Context) error {
	rid := com.StrTo(self.Param("rid")).MustInt64()
	models.DelReply(rid)
	return self.Redirect(302, "/")
}
