package handlers

import (
	"../libs"
	"../models"
	"strconv"
	"time"
)

type NodeEditHandler struct {
	libs.RootAuthHandler
}

func (self *NodeEditHandler) Get() {
	nid, _ := strconv.Atoi(self.Ctx.Params[":nid"])
	nid_handler := models.GetNode(nid)
	self.Data["inode"] = nid_handler
	self.Data["icategory"] = models.GetCategory(int(nid_handler.Pid))

	self.Layout = "layout.html"
	self.TplNames = "node_edit.html"
	self.Render()
}

func (self *NodeEditHandler) Post() {
	inputs := self.Input()
	nid, _ := strconv.Atoi(self.Ctx.Params[":nid"])
	cid, _ := strconv.Atoi(inputs.Get("categoryid"))

	nid_title := inputs.Get("title")
	nid_content := inputs.Get("content")
	if nid_title != "" && nid_content != "" {
		var nd models.Node
		nd.Id = int64(nid)
		nd.Pid = int64(cid)
		nd.Title = nid_title
		nd.Content = nid_content
		nd.Created = time.Now()
		models.SaveNode(nd)
		self.Ctx.Redirect(302, "/node/"+self.Ctx.Params[":nid"])
	} else {
		self.Ctx.Redirect(302, "/")
	}
}
