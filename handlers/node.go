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
	nid, _ := strconv.Atoi(self.Ctx.Params[":nid"])

	nid_handler := models.GetNode(nid)
	nid_handler.Views = nid_handler.Views + 1
	models.UpdateNode(nid_handler)

	limit := 25
	rcs := len(models.GetAllTopicByNid(nid, 0, 0, "hotness"))
	pages, pageout, beginnum, endnum, offset := utils.Pages(rcs, page, limit)
	self.Data["pagesbar"] = utils.Pagesbar("", rcs, pages, pageout, beginnum, endnum, 1)
	self.Data["nodeid"] = nid
	self.Data["topics"] = models.GetAllTopicByNid(nid, offset, limit, "hotness")

	self.TplNames = "node.html"
	self.Layout = "layout.html"

	if nid != 0 {
		if sess_userrole, _ := self.GetSession("userrole").(int64); sess_userrole == -1000 {
			self.Render()
		} else {
			nid_path := strconv.Itoa(int(nid_handler.Pid)) + "/" + strconv.Itoa(int(nid_handler.Id)) + "/"
			nid_name := "index.html"
			rs, _ := self.RenderString()
			utils.Writefile("./archives/"+nid_path, nid_name, rs)
			self.Redirect("/archives/"+nid_path+nid_name, 301)
		}
	} else {
		self.Redirect("/", 302)
	}

}
