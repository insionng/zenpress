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
	nid, _ := strconv.Atoi(inputs.Get("nodeid"))
	cid := int(models.GetNode(nid).Pid)
	uid, _ := self.GetSession("userid").(int64)
	tid_title := inputs.Get("title")
	tid_content := inputs.Get("content")
	if tid_title != "" && tid_content != "" {
		models.AddTopic(inputs.Get("title"), inputs.Get("content"), cid, nid, int(uid))
		self.Ctx.Redirect(302, "/node/"+inputs.Get("nodeid"))
	}
	self.Ctx.Redirect(302, "/")
}
