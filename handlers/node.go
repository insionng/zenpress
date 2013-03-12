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

	///:cid(category-[0-9]+)/:nid(node-[0-9]+)+"/"
	//"/category/:cid([0-9]+)/node/:nid([0-9]+)"

	//inputs := this.Input()
	//id, _ := strconv.Atoi(inputs.Get("id"))
	nodeid, _ := strconv.Atoi(this.Ctx.Params[":nid"])

	this.Data["topics"] = models.GetAllTopicByNode(nodeid)
	nid_handler := models.GetNode(nodeid)
	nid_path := strconv.Itoa(int(nid_handler.Pid)) + "/" + strconv.Itoa(int(nid_handler.Id)) + "/"
	nid_name := "index.html"
	this.TplNames = "node.html"
	this.Layout = "layout.html"

	rs, _, _ := this.RenderString()

	utils.Writefile("./archives/"+nid_path, nid_name, rs)

	this.Redirect("/archives/"+nid_path+nid_name, 302)

	/*this.Render()*/
}
