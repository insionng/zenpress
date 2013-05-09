package handlers

import (
	"../libs"
	"../models"
	"../utils"
)

type MainHandler struct {
	libs.BaseHandler
}

func (self *MainHandler) Get() {
	page, _ := self.GetInt("page")
	cid, _ := self.GetInt(":cid")
	home := "false"
	if cid == 0 {
		home = "true"
	}
	rcs := len(models.GetAllNodeByCid(cid, 0, 0, "hotness"))

	limit := 25
	pages, pageout, beginnum, endnum, offset := utils.Pages(rcs, int(page), limit)
	self.Data["home"] = home
	self.Data["curcate"] = cid
	self.Data["pagesbar"] = utils.Pagesbar("", rcs, pages, pageout, beginnum, endnum, 1)
	self.Data["nodes_latest"] = models.GetAllNodeByCid(cid, offset, limit, "id")
	self.Data["nodes_hotness"] = models.GetAllNodeByCid(cid, offset, limit, "hotness")
	self.Layout = "layout.html"
	self.TplNames = "index.html"
	//self.Render()

}
