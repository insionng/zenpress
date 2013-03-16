package handlers

import (
	"../libs"
	"../models"
	"strconv"
)

type NewTopicHandler struct {
	libs.AuthHandler
}

func (self *NewTopicHandler) Get() {
	self.TplNames = "topic_new.html"
	self.Layout = "layout.html"
	self.Data["nodes"] = models.GetAllNode()
	self.Render()
}

func (self *NewTopicHandler) Post() {
	inputs := self.Input()
	nodeid, _ := strconv.Atoi(inputs.Get("nodeid"))
	cid := int(models.GetNode(nodeid).Pid)
	models.AddTopic(inputs.Get("title"), inputs.Get("content"), cid, nodeid)

	self.Ctx.Redirect(302, "/")
}
