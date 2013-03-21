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
	tid, _ := strconv.Atoi(self.Ctx.Params[":tid"])
	tid_handler := models.GetTopic(tid)
	self.Data["article"] = tid_handler
	self.Data["replys"] = models.GetReplyByPid(tid, 0, 0, "id")
	tid_path := strconv.Itoa(int(tid_handler.Cid)) + "/" + strconv.Itoa(int(tid_handler.Nid)) + "/"
	tid_name := strconv.Itoa(int(tid_handler.Id)) + ".html"
	self.TplNames = "view.html"
	self.Layout = "layout.html"

	if sess_userrole, _ := self.GetSession("userrole").(int64); sess_userrole == -1000 {
		self.Render()
	} else {
		rs, _ := self.RenderString()
		utils.Writefile("./archives/"+tid_path, tid_name, rs)
		self.Redirect("/archives/"+tid_path+tid_name, 301)
	}
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
