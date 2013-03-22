package handlers

import (
	"../libs"
	"../models"
	"strconv"
)

type NewNodeHandler struct {
	libs.AuthHandler
}

func (self *NewNodeHandler) Get() {
	self.TplNames = "new_node.html"
	self.Layout = "layout.html"

	self.Data["categorys"] = models.GetAllCategory()
	self.Render()
}

func (self *NewNodeHandler) Post() {
	inputs := self.Input()
	cid, _ := strconv.Atoi(inputs.Get("category"))
	uid, _ := self.GetSession("userid").(int64)
	nid_title := inputs.Get("title")
	nid_content := inputs.Get("content")

	if nid_title != "" && nid_content != "" && cid != 0 {
		models.AddNode(nid_title, nid_content, cid, int(uid))
		self.Ctx.Redirect(302, "/category/"+inputs.Get("category"))
	} else {
		self.Ctx.Redirect(302, "/")
	}
}
