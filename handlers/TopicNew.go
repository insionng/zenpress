package handlers

import (
	"../libs"
	"../models"
	"strconv"
)

type NewTopicHandler struct {
	libs.AuthHandler
}

func (this *NewTopicHandler) Get() {
	this.TplNames = "topic_new.html"
	this.Layout = "layout.html"
	this.Data["nodes"] = models.GetAllNode()
	this.Render()
}

func (this *NewTopicHandler) Post() {
	inputs := this.Input()
	nodeid, _ := strconv.Atoi(inputs.Get("nodeid"))
	cid := int(models.GetNode(nodeid).Pid)
	models.AddTopic(inputs.Get("title"), inputs.Get("content"), cid, nodeid)

	this.Ctx.Redirect(302, "/")
}
