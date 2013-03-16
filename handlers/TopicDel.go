package handlers

import (
	"../libs"
	"../models"
	"strconv"
)

type TopicDeleteHandler struct {
	libs.RootAuthHandler
}

func (self *TopicDeleteHandler) Get() {
	tid, _ := strconv.Atoi(self.Ctx.Params[":tid"])
	models.DelTopic(tid)
	self.Ctx.Redirect(302, "/")
}
