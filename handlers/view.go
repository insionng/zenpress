package handlers

import (
	"../models"
	"../utils"
	"strconv"
)

type ViewHandler struct {
	utils.BaseHandler
}

func (this *ViewHandler) Get() {
	//  "/:cid(category-[0-9]+)/:nid(node-[0-9]+)/:tid(topic-[0-9]+)"
	//"/category/:cid([0-9]+)/node/:nid([0-9]+)/topic/:tid([0-9]+)"

	//cid, _ := strconv.Atoi(this.Ctx.Params[":cid"])
	//nid, _ := strconv.Atoi(this.Ctx.Params[":pid"])
	tid, _ := strconv.Atoi(this.Ctx.Params[":tid"])

	this.Data["article"] = models.GetTopic(tid)

	this.TplNames = "view.html"
	this.Layout = "layout.html"

	rs, _, _ := this.RenderString()

	utils.Writefile("./archives"+this.Ctx.Request.RequestURI, this.Ctx.Params[":tid"], ".html", rs)

	this.Redirect("/archives"+this.Ctx.Request.RequestURI+".html", 302)

	/*this.Render()*/
}
