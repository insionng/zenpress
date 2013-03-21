package handlers

import (
	"../libs"
	"../models"
	"../utils"
	"strconv"
)

type NodeHandler struct {
	libs.BaseHandler
}

func (self *NodeHandler) Get() {
	inputs := self.Input()
	page, _ := strconv.Atoi(inputs.Get("page"))
	nodeid, _ := strconv.Atoi(self.Ctx.Params[":nid"])

	limit := 25
	rcs := len(models.GetAllTopicByNodeid(nodeid, 0, 0, "hotness"))
	pages, pageout, beginnum, endnum, offset := utils.Pages(rcs, page, limit)
	self.Data["pagesbar"] = utils.Pagesbar("", rcs, pages, pageout, beginnum, endnum, 1)
	self.Data["nodeid"] = nodeid
	self.Data["topics"] = models.GetAllTopicByNodeid(nodeid, offset, limit, "hotness")
	nid_handler := models.GetNode(nodeid)
	nid_path := strconv.Itoa(int(nid_handler.Pid)) + "/" + strconv.Itoa(int(nid_handler.Id)) + "/"
	nid_name := "index.html"
	self.TplNames = "node.html"
	self.Layout = "layout.html"

	if sess_userrole, _ := self.GetSession("userrole").(int64); sess_userrole == -1000 {
		self.Render()
	} else {
		rs, _ := self.RenderString()
		utils.Writefile("./archives/"+nid_path, nid_name, rs)
		self.Redirect("/archives/"+nid_path+nid_name, 301)
	}
}
