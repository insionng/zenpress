package handlers

import (
	"fmt"
	"toropress/libs"
	"toropress/models"
)

type NewTopicHandler struct {
	libs.AuthHandler
}

func (self *NewTopicHandler) Get() {
	self.TplNames = "topic_new.html"
	self.Layout = "layout.html"
	self.Data["nodes"] = models.GetAllNode()
	self.Render()
}

func (self *NewTopicHandler) Post() {
	nid, _ := self.GetInt("nodeid")
	cid := models.GetNode(nid).Pid
	uid, _ := self.GetSession("userid").(int64)
	author,_ := self.GetSession("username").(string)

	fmt.Println("=============================")
	fmt.Println(self.GetSession("userid"))
	fmt.Println(uid)

	tid_title := self.GetString("title")
	tid_content := self.GetString("content")

	if tid_title != "" && tid_content != "" {
		models.AddTopic(tid_title, tid_content, cid, nid, uid, author)
		self.Ctx.Redirect(302, "/node/"+self.GetString("nodeid"))
	} else {
		self.Ctx.Redirect(302, "/")
	}
}
