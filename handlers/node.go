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

func (this *NodeHandler) Get() {
	nodeid, _ := strconv.Atoi(this.Ctx.Params[":nid"])

	this.Data["topics"] = models.GetAllTopicByNode(nodeid, "hotness")
	nid_handler := models.GetNode(nodeid)
	nid_path := strconv.Itoa(int(nid_handler.Pid)) + "/" + strconv.Itoa(int(nid_handler.Id)) + "/"
	nid_name := "index.html"
	this.TplNames = "node.html"
	this.Layout = "layout.html"

	rs, _ := this.RenderString()

	utils.Writefile("./archives/"+nid_path, nid_name, rs)

	this.Redirect("/archives/"+nid_path+nid_name, 302)

	/*this.Render()*/
}
