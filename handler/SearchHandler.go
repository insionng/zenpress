package handler

import (
	"github.com/insionng/macross"
	"github.com/insionng/zenpress/helper"
	"github.com/insionng/zenpress/models"
)

func SearchHandler(self *macross.Context) error {
	data := make(map[string]interface{})
	if keyword := self.Args("keyword").String(); len(keyword) != 0 {
		page := self.Param("page").MustInt()
		limit := 25

		rcs := len(*models.SearchTopic(keyword, 0, 0, "id"))
		pages, pageout, beginnum, endnum, offset := helper.Pages(rcs, int(page), limit)
		data["search_hotness"] = models.SearchTopic(keyword, offset, limit, "hotness")

		keywordz := "keyword=" + keyword + "&"
		data["pagesbar"] = helper.Pagesbar(keywordz, rcs, pages, pageout, beginnum, endnum, 1)

	}

	self.SetStore(data)
	return self.Render("search")
}
