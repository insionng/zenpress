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
	self.Data["icategory"] = models.GetCategory(int(models.GetNode(nid).Pid))

	self.Layout = "layout.html"
	self.TplNames = "node_edit.html"
	self.Render()
}

func (self *NodeEditHandler) Post() {
	inputs := self.Input()
	nid, _ := strconv.Atoi(self.Ctx.Params[":nid"])
	cid, _ := strconv.Atoi(inputs.Get("categoryid"))

	var nd models.Node
	nd.Id = int64(nid)
	nd.Pid = int64(cid)
	nd.Title = inputs.Get("title")
	nd.Content = inputs.Get("content")
	nd.Created = time.Now()
	models.SaveNode(nd)
	self.Ctx.Redirect(302, "/")
}
