package handler

import (
	//"github.com/insionng/toropress/helper"
	"net/http"
	//"github.com/insionng/toropress/libs"
	//"github.com/insionng/toropress/models"
	"github.com/insionng/vodka"
)

func MainHandler(self *vodka.Context) error {
	return self.Render(http.StatusOK, "index.html", nil)
}

/*
func MainGetHandler(self *vodka.Context) {
	data := make(map[string]interface{}, 0)
	page, _ := self.ParamInt64("page")
	curtab, _ := self.ParamInt("tab")
	cid, _ := self.ParamInt(":cid")
	limit := 25
	home := "false"
	if cid == 0 {
		home = "true"
	}

	data["home"] = home
	data["curcate"] = cid
	data["curtab"] = curtab

	topics_rcs := len(models.GetAllTopicByCid(cid, 0, 0, 0, "hotness"))
	topics_pages, topics_pageout, topics_beginnum, topics_endnum, offset := helper.Pages(topics_rcs, int(page), limit)

	data["topics_latest"] = models.GetAllTopicByCid(cid, offset, limit, 0, "id")
	data["topics_hotness"] = models.GetAllTopicByCid(cid, offset, limit, 0, "hotness")
	data["topics_views"] = models.GetAllTopicByCid(cid, offset, limit, 0, "views")
	data["topics_reply_count"] = models.GetAllTopicByCid(cid, offset, limit, 0, "reply_count")

	data["topics_pagesbar_tab1"] = helper.Pagesbar("tab=1&", topics_rcs, topics_pages, topics_pageout, topics_beginnum, topics_endnum, 1)
	data["topics_pagesbar_tab2"] = helper.Pagesbar("tab=2&", topics_rcs, topics_pages, topics_pageout, topics_beginnum, topics_endnum, 1)
	data["topics_pagesbar_tab3"] = helper.Pagesbar("tab=3&", topics_rcs, topics_pages, topics_pageout, topics_beginnum, topics_endnum, 1)
	data["topics_pagesbar_tab4"] = helper.Pagesbar("tab=4&", topics_rcs, topics_pages, topics_pageout, topics_beginnum, topics_endnum, 1)

	nodes_rcs := len(models.GetAllNodeByCid(cid, 0, 0, 0, "hotness"))
	nodes_pages, nodes_pageout, nodes_beginnum, nodes_endnum, offset := helper.Pages(nodes_rcs, int(page), limit)

	data["nodes_latest"] = models.GetAllNodeByCid(cid, offset, limit, 0, "id")
	data["nodes_hotness"] = models.GetAllNodeByCid(cid, offset, limit, 0, "hotness")

	data["nodes_pagesbar_tab5"] = helper.Pagesbar("tab=5&", nodes_rcs, nodes_pages, nodes_pageout, nodes_beginnum, nodes_endnum, 1)
	data["nodes_pagesbar_tab6"] = helper.Pagesbar("tab=6&", nodes_rcs, nodes_pages, nodes_pageout, nodes_beginnum, nodes_endnum, 1)

	self.Render(http.StatusOK, "index.html", data)

}
*/
