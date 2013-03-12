package handlers

import (
	"../libs"
	"../models"
	"strconv"
)

type TopicDeleteHandler struct {
	libs.RootAuthHandler
}

func (this *TopicDeleteHandler) Get() {
	tid, _ := strconv.Atoi(this.Ctx.Params[":tid"])
	models.DelTopic(tid)
	this.Ctx.Redirect(302, "/")
}
