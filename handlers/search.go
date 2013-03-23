package handlers

import (
	"../libs"
	"../models"
	"../utils"
	"strconv"
)

type SearchHandler struct {
	libs.BaseHandler
}

func (self *SearchHandler) Get() {
	inputs := self.Input()
	if keyword := inputs.Get("keyword"); keyword != "" {
		page, _ := strconv.Atoi(inputs.Get("page"))
		limit := 25
		rcs := len(models.SearchTopic(keyword, 0, 0, "id"))
		pages, pageout, beginnum, endnum, offset := utils.Pages(rcs, page, limit)
		self.Data["pagesbar"] = utils.Pagesbar(keyword, rcs, pages, pageout, beginnum, endnum, 1)
		self.Data["search_hotness"] = models.SearchTopic(keyword, offset, limit, "hotness")
	}
	self.TplNames = "search.html"
	self.Layout = "layout.html"

	self.Render()
}
