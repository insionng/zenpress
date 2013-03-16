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

func (self *ViewHandler) Get() {
	//  "/:cid(category-[0-9]+)/:nid(node-[0-9]+)/:tid(topic-[0-9]+)"
	//"/category/:cid([0-9]+)/node/:nid([0-9]+)/topic/:tid([0-9]+)"

	//cid, _ := strconv.Atoi(self.Ctx.Params[":cid"])
	//nid, _ := strconv.Atoi(self.Ctx.Params[":pid"])
	tid, _ := strconv.Atoi(self.Ctx.Params[":tid"])
	tid_handler := models.GetTopic(tid)
	self.Data["article"] = tid_handler
	self.Data["replys"] = models.GetReplyByPid(tid, 0, 0, "id")
	tid_path := strconv.Itoa(int(tid_handler.Cid)) + "/" + strconv.Itoa(int(tid_handler.Nid)) + "/"
	tid_name := strconv.Itoa(int(tid_handler.Id)) + ".html"
	self.TplNames = "view.html"
	self.Layout = "layout.html"

	rs, _ := self.RenderString()

	utils.Writefile("./archives/"+tid_path, tid_name, rs)

	self.Redirect("/archives/"+tid_path+tid_name, 302)

	/*self.Render()*/
}

func (self *ViewHandler) Post() {
	inputs := self.Input()
	tid, _ := strconv.Atoi(inputs.Get("comment_parent"))
	uid, _ := strconv.Atoi(inputs.Get("comment_userid"))

	author := inputs.Get("author")
	email := inputs.Get("email")
	website := inputs.Get("website")

	rc := inputs.Get("comment")
	models.AddReply(tid, uid, rc, author, email, website)
	self.Ctx.Redirect(302, "/")
}
