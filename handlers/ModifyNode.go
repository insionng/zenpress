package handlers

import (
	"../libs"
	"../models"
	"strconv"
	"time"
)

type ModifyNodeHandler struct {
	libs.RootAuthHandler
}

func (self *ModifyNodeHandler) Get() {

	self.Layout = "layout.html"
	self.TplNames = "modify_node.html"
	self.Render()
}

func (self *ModifyNodeHandler) Post() {

	inputs := self.Input()
	cid, _ := strconv.Atoi(inputs.Get("categoryid"))
	nid, _ := strconv.Atoi(inputs.Get("nodeid"))

	var nd models.Node
	nd.Id = int64(nid)
	nd.Pid = int64(cid)
	nd.Title = inputs.Get("title")
	nd.Content = inputs.Get("content")
	nd.Created = time.Now()
	models.SaveNode(nd)
	self.Ctx.Redirect(302, "/")
}
