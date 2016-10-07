package handler

import (
	"github.com/Unknwon/com"
	"github.com/insionng/vodka"
	"github.com/insionng/zenpress/models"
)

func TopicDeleteHandler(self vodka.Context) error {
	tid := com.StrTo(self.Param("tid")).MustInt64()
	models.DelTopic(tid)
	return self.Redirect(302, "/")
}
