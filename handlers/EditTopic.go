package handlers

import (
	"../libs"
	"../models"
	"strconv"
)

type TopicEditHandler struct {
	libs.RootAuthHandler
}

func (self *TopicEditHandler) Get() {
	tid, _ := strconv.Atoi(self.Ctx.Params[":tid"])
	tid_handler := models.GetTopic(tid)
	self.Data["topic"] = tid_handler
	self.Data["inode"] = models.GetNode(int(tid_handler.Nid))

	self.Layout = "layout.html"
	self.TplNames = "topic_edit.html"
	self.Render()
}

func (self *TopicEditHandler) Post() {
	inputs := self.Input()
	tid, _ := strconv.Atoi(self.Ctx.Params[":tid"])
	nid, _ := strconv.Atoi(inputs.Get("nodeid"))
	cid := models.GetNode(nid).Pid
	uid, _ := self.GetSession("userid").(int64)
	tid_title := inputs.Get("title")
	tid_content := inputs.Get("content")

	if tid_title != "" && tid_content != "" {
		models.EditTopic(tid, nid, int(cid), int(uid), tid_title, tid_content)
		self.Ctx.Redirect(302, "/view/"+self.Ctx.Params[":tid"])
	} else {
		self.Ctx.Redirect(302, "/")
	}
}
