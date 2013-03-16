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
	nodeid, _ := strconv.Atoi(self.Ctx.Params[":nid"])

	self.Data["topics"] = models.GetAllTopicByNode(nodeid, "hotness")
	nid_handler := models.GetNode(nodeid)
	nid_path := strconv.Itoa(int(nid_handler.Pid)) + "/" + strconv.Itoa(int(nid_handler.Id)) + "/"
	nid_name := "index.html"
	self.TplNames = "node.html"
	self.Layout = "layout.html"

	rs, _ := self.RenderString()

	utils.Writefile("./archives/"+nid_path, nid_name, rs)

	self.Redirect("/archives/"+nid_path+nid_name, 302)

	/*self.Render()*/
}
