package handler

import (
	"github.com/insionng/vodka"
	"github.com/insionng/zenpress/helper"
	"github.com/insionng/zenpress/models"
	"net/http"
)

func SearchHandler(self *vodka.Context) error {
	data := make(map[string]interface{})
	if keyword := self.FormEscape("keyword"); keyword != "" {
		page, _ := self.ParamInt("page")
		limit := 25

		rcs := len(*models.SearchTopic(keyword, 0, 0, "id"))
		pages, pageout, beginnum, endnum, offset := helper.Pages(rcs, int(page), limit)
		data["search_hotness"] = models.SearchTopic(keyword, offset, limit, "hotness")

		keywordz := "keyword=" + keyword + "&"
		data["pagesbar"] = helper.Pagesbar(keywordz, rcs, pages, pageout, beginnum, endnum, 1)

	}

	return self.Render(http.StatusOK, "search.html", data)
}
