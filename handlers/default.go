package handlers

import (
	"../models"
	"../utils"
	"strconv"
)

type MainHandler struct {
	utils.BaseHandler
}

func (this *MainHandler) Get() {

	id, _ := strconv.Atoi(this.Ctx.Params[":cid"])
	offset := 0
	limit := 25

	this.Data["nodes"] = models.GetAllNodeByCategoryId(id, offset, limit, "hotness")
	this.Data["nodes_1"] = models.GetAllNodeByCategoryId(1, offset, 10, "id")
	this.Data["categorys"] = models.GetAllCategory()
	this.Data["allnode"] = models.GetAllNode()
	this.Data["replys_5s"] = models.GetReplyByPid(0, 0, 5, "id")
	this.Layout = "layout.html"
	this.TplNames = "index.html"

	this.Render()

}
