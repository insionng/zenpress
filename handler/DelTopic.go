package handler

import (
	"github.com/insionng/vodka"
	"github.com/insionng/zenpress/models"
)

func TopicDeleteHandler(self *vodka.Context) error {
	tid, _ := self.ParamInt64(":tid")
	models.DelTopic(tid)
	return self.Redirect(302, "/")
}
