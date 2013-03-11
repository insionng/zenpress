package handlers

import (
	"../models"
	"../utils"
	"strconv"
)

type NewTopicHandler struct {
	utils.RootAuthHandler
}

func (this *NewTopicHandler) Get() {
	this.TplNames = "new_topic.html"
	this.Layout = "layout.html"

	this.Render()
}

func (this *NewTopicHandler) Post() {
	inputs := this.Input()
	nodeid, _ := strconv.Atoi(inputs.Get("nodeid"))
	models.AddTopic(inputs.Get("title"), inputs.Get("content"), int64(nodeid))

	this.Ctx.Redirect(302, "/")
}
