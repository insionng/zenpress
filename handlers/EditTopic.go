package handlers

import (
	"../libs"
	"../models"
	"strconv"
	"time"
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

	tid_title := inputs.Get("title")
	tid_content := inputs.Get("content")
	if tid_title != "" && tid_content != "" {
		var tp models.Topic
		tp.Id = int64(tid)
		tp.Cid = cid
		tp.Nid = int64(nid)
		tp.Title = tid_title
		tp.Content = tid_content
		tp.Created = time.Now()
		models.SaveTopic(tp)
		self.Ctx.Redirect(302, "/view/"+self.Ctx.Params[":tid"])
	}
	self.Ctx.Redirect(302, "/")
}
