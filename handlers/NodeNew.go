package handlers

import (
	"../libs"
	"../models"
	"strconv"
)

type NewNodeHandler struct {
	libs.AuthHandler
}

func (self *NewNodeHandler) Get() {
	self.TplNames = "new_node.html"
	self.Layout = "layout.html"

	self.Data["categorys"] = models.GetAllCategory()
	self.Render()
}

func (self *NewNodeHandler) Post() {
	inputs := self.Input()
	cid, _ := strconv.Atoi(inputs.Get("category"))

	models.AddNode(inputs.Get("title"), inputs.Get("content"), cid)

	self.Ctx.Redirect(302, "/")
}
