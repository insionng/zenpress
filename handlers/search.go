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
	page, _ := strconv.Atoi(inputs.Get("page"))
	keyword := inputs.Get("keyword")
	limit := 25
	rcs := len(models.SearchAllTopicByContent(keyword, 0, 0, "id"))
	pages, pageout, beginnum, endnum, offset := utils.Pages(rcs, page, limit)
	self.Data["pagesbar"] = utils.Pagesbar(keyword, rcs, pages, pageout, beginnum, endnum)
	self.Data["search_hotness"] = models.SearchAllTopicByContent(keyword, offset, limit, "hotness")
	self.TplNames = "search.html"
	self.Layout = "layout.html"

	self.Render()
}
