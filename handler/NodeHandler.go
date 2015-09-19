package handler

import (
	"github.com/insionng/vodka"
	"github.com/insionng/zenpress/helper"
	"github.com/insionng/zenpress/models"
	"net/http"
)

func NodeHandler(self *vodka.Context) error {
	data := make(map[string]interface{})
	page, _ := self.ParamInt64("page")
	nid, _ := self.ParamInt64(":nid")

	nid_handler := models.GetNode(nid)
	nid_handler.Views = nid_handler.Views + 1
	models.UpdateNode(nid, nid_handler)

	limit := 25
	rcs := len(models.GetAllTopicByNid(nid, 0, 0, 0, "hotness"))
	pages, pageout, beginnum, endnum, offset := helper.Pages(rcs, int(page), limit)
	data["pagesbar"] = helper.Pagesbar("", rcs, pages, pageout, beginnum, endnum, 1)
	data["nodeid"] = nid
	data["topics_hotness"] = models.GetAllTopicByNid(nid, offset, limit, 0, "hotness")
	data["topics_latest"] = models.GetAllTopicByNid(nid, offset, limit, 0, "id")

	if nid != 0 {
		return self.Render(http.StatusOK, "node.html", data)
	} else {
		return self.Redirect(302, "/")
	}

}
