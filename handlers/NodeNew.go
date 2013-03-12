package handlers

import (
	"../libs"
	"../models"
	"strconv"
)

type NewNodeHandler struct {
	libs.AuthHandler
}

func (this *NewNodeHandler) Get() {
	this.TplNames = "new_node.html"
	this.Layout = "layout.html"

	this.Data["categorys"] = models.GetAllCategory()
	this.Render()
}

func (this *NewNodeHandler) Post() {
	inputs := this.Input()
	cid, _ := strconv.Atoi(inputs.Get("category"))

	models.AddNode(inputs.Get("title"), inputs.Get("content"), cid)

	this.Ctx.Redirect(302, "/")
}
