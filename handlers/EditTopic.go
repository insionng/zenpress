package handlers

import (
	"toropress/libs"
	"toropress/models"
)

type TopicEditHandler struct {
	libs.RootAuthHandler
}

func (self *TopicEditHandler) Get() {
	tid, _ := self.GetInt(":tid")
	tid_handler := models.GetTopic(tid)
	self.Data["topic"] = tid_handler
	self.Data["inode"] = models.GetNode(tid_handler.Nid)

	self.Layout = "layout.html"
	self.TplNames = "topic_edit.html"
	self.Render()
}

func (self *TopicEditHandler) Post() {
	tid, _ := self.GetInt(":tid")
	nid, _ := self.GetInt("nodeid")
	cid := models.GetNode(nid).Pid
	uid, _ := self.GetSession("userid").(int64)
	tid_title := self.GetString("title")
	tid_content := self.GetString("content")

	if tid_title != "" && tid_content != "" {
		models.EditTopic(tid, nid, cid, uid, tid_title, tid_content)
		self.Redirect("/view/"+self.GetString(":tid"), 302)
	} else {
		self.Redirect("/", 302)
	}
}
