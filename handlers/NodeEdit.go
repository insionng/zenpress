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

func (this *NodeEditHandler) Get() {
	nid, _ := strconv.Atoi(this.Ctx.Params[":nid"])
	this.Data["icategory"] = models.GetCategory(int(models.GetNode(nid).Pid))

	this.Layout = "layout.html"
	this.TplNames = "node_edit.html"
	this.Render()
}

func (this *NodeEditHandler) Post() {
	inputs := this.Input()
	nid, _ := strconv.Atoi(this.Ctx.Params[":nid"])
	cid, _ := strconv.Atoi(inputs.Get("categoryid"))

	var nd models.Node
	nd.Id = int64(nid)
	nd.Pid = int64(cid)
	nd.Title = inputs.Get("title")
	nd.Content = inputs.Get("content")
	nd.Created = time.Now()
	models.SaveNode(nd)
	this.Ctx.Redirect(302, "/")
}
