package handlers

import (
	"../libs"
	"../models"
	"strconv"
)

type NodeDeleteHandler struct {
	libs.RootAuthHandler
}

func (this *NodeDeleteHandler) Get() {
	nid, _ := strconv.Atoi(this.Ctx.Params[":nid"])
	models.DelNode(nid)
	this.Ctx.Redirect(302, "/")
}
