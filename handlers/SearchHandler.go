package handlers

import (
	"toropress/helper"
	"toropress/libs"
	"toropress/models"
)

type SearchHandler struct {
	libs.BaseHandler
}

func (self *SearchHandler) Get() {
	if keyword := self.GetString("keyword"); keyword != "" {
		page, _ := self.GetInt("page")
		limit := 25

		rcs := len(*models.SearchTopic(keyword, 0, 0, "id"))
		pages, pageout, beginnum, endnum, offset := helper.Pages(rcs, int(page), limit)
		self.Data["search_hotness"] = models.SearchTopic(keyword, offset, limit, "hotness")

		keywordz := "keyword=" + keyword + "&"
		self.Data["pagesbar"] = helper.Pagesbar(keywordz, rcs, pages, pageout, beginnum, endnum, 1)

	}
	self.TplNames = "search.html"
	self.Layout = "layout.html"

	self.Render()
}
