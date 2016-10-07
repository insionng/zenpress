package handler

import (
	"net/http"
	"runtime"
	"time"

	"github.com/Unknwon/com"
	"github.com/insionng/vodka"
	"github.com/insionng/zenpress/helper"
	"github.com/insionng/zenpress/models"
)

func MainHandler(self vodka.Context) error {
	data := make(map[string]interface{})
	///
	data["nodes"] = models.GetAllNode()
	data["nodes_hotness_topbar"] = models.GetAllNodeByCid(0, 0, 16, 0, "hotness")
	data["topics_5s"] = models.GetAllTopic(0, 5, "id")
	data["topics_10s"] = models.GetAllTopic(0, 10, "id")
	data["nodes_10s"] = models.GetAllNodeByCid(0, 0, 10, 0, "id")
	data["replys_5s"] = models.GetReplyByPid(0, 0, 5, "id")
	data["replys_10s"] = models.GetReplyByPid(0, 0, 10, "id")

	data["author"] = models.GetKV("author")
	data["title"] = models.GetKV("title")
	data["title_en"] = models.GetKV("title_en")
	data["keywords"] = models.GetKV("keywords")
	data["description"] = models.GetKV("description")

	data["company"] = models.GetKV("company")
	data["copyright"] = models.GetKV("copyright")
	data["site_email"] = models.GetKV("site_email")

	data["tweibo"] = models.GetKV("tweibo")
	data["sweibo"] = models.GetKV("sweibo")
	data["timenow"] = time.Now()
	data["statistics"] = models.GetKV("statistics")
	data["remoteproto"] = self.Request().Scheme()
	data["remotehost"] = self.Request().Host
	data["remoteos"] = runtime.GOOS
	data["remotearch"] = runtime.GOARCH
	data["remotecpus"] = runtime.NumCPU()
	data["golangver"] = runtime.Version()
	///
	page := com.StrTo(self.Param("page")).MustInt64()
	curtab := com.StrTo(self.Param("tab")).MustInt64()
	cid := com.StrTo(self.Param("cid")).MustInt64()

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

	return self.Render(http.StatusOK, "index.html", nil)

}
