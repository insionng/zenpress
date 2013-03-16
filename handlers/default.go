package handlers

import (
	"../libs"
	"../models"
	"strconv"
)

type MainHandler struct {
	libs.BaseHandler
}

func (self *MainHandler) Get() {
	models.Ct()

	id, _ := strconv.Atoi(self.Ctx.Params[":cid"])
	offset := 0
	limit := 25

	self.Data["nodes_hotness"] = models.GetAllNodeByCategoryId(id, offset, limit, "hotness")
	self.Layout = "layout.html"
	self.TplNames = "index.html"
	self.Render()

}
