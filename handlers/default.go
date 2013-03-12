package handlers

import (
	"../libs"
	"../models"
	"strconv"
)

type MainHandler struct {
	libs.BaseHandler
}

func (this *MainHandler) Get() {
	models.Ct()

	id, _ := strconv.Atoi(this.Ctx.Params[":cid"])
	offset := 0
	limit := 25

	this.Data["nodes_hotness"] = models.GetAllNodeByCategoryId(id, offset, limit, "hotness")
	this.Layout = "layout.html"
	this.TplNames = "index.html"
	this.Render()

}
