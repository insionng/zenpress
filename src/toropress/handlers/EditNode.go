package handlers

import (
	"toropress/libs"
	"toropress/models"
)

type NodeEditHandler struct {
	libs.RootAuthHandler
}

func (self *NodeEditHandler) Get() {
	nid, _ := self.GetInt(":nid")
	nid_handler := models.GetNode(nid)
	self.Data["inode"] = nid_handler
	self.Data["icategory"] = models.GetCategory(nid_handler.Pid)

	self.Layout = "layout.html"
	self.TplNames = "node_edit.html"
	self.Render()
}

func (self *NodeEditHandler) Post() {
	nid, _ := self.GetInt(":nid")
	cid, _ := self.GetInt("categoryid")
	uid, _ := self.GetSession("userid").(int64)
	nid_title := self.GetString("title")
	nid_content := self.GetString("content")
	if nid_title != "" && nid_content != "" {
		models.EditNode(nid, cid, uid, nid_title, nid_content)
		self.Redirect("/node/"+self.GetString(":nid"), 302)
	} else {
		self.Redirect("/", 302)
	}
}
