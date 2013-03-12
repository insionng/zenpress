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

func (this *ModifyNodeHandler) Get() {
	this.Layout = "layout.html"
	this.TplNames = "modify_node.html"
	this.Render()
}

func (this *ModifyNodeHandler) Post() {
	inputs := this.Input()
	cid, _ := strconv.Atoi(inputs.Get("categoryid"))
	nid, _ := strconv.Atoi(inputs.Get("nodeid"))

	var nd models.Node
	nd.Id = int64(nid)
	nd.Pid = int64(cid)
	nd.Title = inputs.Get("title")
	nd.Content = inputs.Get("content")
	nd.Created = time.Now()
	models.SaveNode(nd)
	this.Ctx.Redirect(302, "/")
}
