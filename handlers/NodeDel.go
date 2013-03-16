package handlers

import (
	"../libs"
	"../models"
	"strconv"
)

type NodeDeleteHandler struct {
	libs.RootAuthHandler
}

func (self *NodeDeleteHandler) Get() {
	nid, _ := strconv.Atoi(self.Ctx.Params[":nid"])
	models.DelNode(nid)
	self.Ctx.Redirect(302, "/")
}
