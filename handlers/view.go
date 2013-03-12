package handlers

import (
	"../libs"
	"../models"
	"../utils"
	"strconv"
)

type ViewHandler struct {
	libs.BaseHandler
}

func (this *ViewHandler) Get() {
	//  "/:cid(category-[0-9]+)/:nid(node-[0-9]+)/:tid(topic-[0-9]+)"
	//"/category/:cid([0-9]+)/node/:nid([0-9]+)/topic/:tid([0-9]+)"

	//cid, _ := strconv.Atoi(this.Ctx.Params[":cid"])
	//nid, _ := strconv.Atoi(this.Ctx.Params[":pid"])
	tid, _ := strconv.Atoi(this.Ctx.Params[":tid"])
	tid_handler := models.GetTopic(tid)
	this.Data["article"] = tid_handler
	this.Data["replys"] = models.GetReplyByPid(tid, 0, 0, "id")
	tid_path := strconv.Itoa(int(tid_handler.Cid)) + "/" + strconv.Itoa(int(tid_handler.Nid)) + "/"
	tid_name := strconv.Itoa(int(tid_handler.Id)) + ".html"
	this.TplNames = "view.html"
	this.Layout = "layout.html"

	rs, _, _ := this.RenderString()

	utils.Writefile("./archives/"+tid_path, tid_name, rs)

	this.Redirect("/archives/"+tid_path+tid_name, 302)

	/*this.Render()*/
}

func (this *ViewHandler) Post() {
	inputs := this.Input()
	tid, _ := strconv.Atoi(inputs.Get("comment_parent"))
	uid, _ := strconv.Atoi(inputs.Get("comment_userid"))

	author := inputs.Get("author")
	email := inputs.Get("email")
	website := inputs.Get("website")

	rc := inputs.Get("comment")
	models.AddReply(tid, uid, rc, author, email, website)
	this.Ctx.Redirect(302, "/")
}
