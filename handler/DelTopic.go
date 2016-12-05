package handler

import (
	"github.com/insionng/macross"
	"github.com/insionng/zenpress/models"
)

func TopicDeleteHandler(self *macross.Context) error {
	tid := self.Param("tid").MustInt64()
	models.DelTopic(tid)
	return self.Redirect("/", 302)
}
