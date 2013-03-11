package handlers

import (
	"../models"
	"../utils"
	"strconv"
)

type NodeHandler struct {
	utils.BaseHandler
}

func (this *NodeHandler) Get() {

	///:cid(category-[0-9]+)/:nid(node-[0-9]+)+"/"
	//"/category/:cid([0-9]+)/node/:nid([0-9]+)"

	//inputs := this.Input()
	//id, _ := strconv.Atoi(inputs.Get("id"))
	nodeid, _ := strconv.Atoi(this.Ctx.Params[":nid"])
	this.Data["topics"] = models.GetAllTopicByNode(nodeid)
	this.TplNames = "node.html"
	this.Layout = "layout.html"

	rs, _, _ := this.RenderString()

	utils.Writefile("./archives"+this.Ctx.Request.RequestURI, this.Ctx.Params[":nid"], ".html", rs)

	this.Redirect("/archives"+this.Ctx.Request.RequestURI+".html", 302)

	/*this.Render()*/
}
