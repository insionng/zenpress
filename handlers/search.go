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
	limit := 5
	rcs := len(models.GetAllTopic(0, 0, "id"))
	pages, pageout, beginnum, endnum, offset := utils.Pages(rcs, page, limit)
	self.Data["pages"] = pages
	self.Data["page"] = pageout
	self.Data["beginnum"] = beginnum
	self.Data["endnum"] = endnum

	self.Data["search_hotness"] = models.SearchAllTopicByTitle(keyword, offset, limit, "hotness")
	self.TplNames = "search.html"
	self.Layout = "layout.html"

	self.Render()
}
