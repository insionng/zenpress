package libs

import (
	"../models"
	"github.com/insionng/torgo"
	//"../../torgo"
	//"lihuashu.com/insionng/torgo"
	//"github.com/astaxie/beego"
)

type BaseHandler struct {
	torgo.Handler
}

type AuthHandler struct {
	torgo.Handler
}

type RootAuthHandler struct {
	torgo.Handler
}

func (self *BaseHandler) Prepare() {
	sess_username, _ := self.GetSession("username").(string)
	sess_uid, _ := self.GetSession("userid").(int64)
	sess_role, _ := self.GetSession("userrole").(int64)
	sess_email, _ := self.GetSession("useremail").(string)
	if sess_role <= 0 {
		self.Data["Userid"] = 0
		self.Data["Username"] = ""
		self.Data["Userrole"] = 0
		self.Data["Useremail"] = ""
	} else {
		self.Data["Userid"] = sess_uid
		self.Data["Username"] = sess_username
		self.Data["Userrole"] = sess_role
		self.Data["Useremail"] = sess_email
	}

	self.Data["categorys"] = models.GetAllCategory()
	self.Data["nodes"] = models.GetAllNode()
	self.Data["nodes_10s"] = models.GetAllNodeByCategoryId(0, 0, 10, "id")
	self.Data["replys_5s"] = models.GetReplyByPid(0, 0, 5, "id")
	self.Data["replys_10s"] = models.GetReplyByPid(0, 0, 10, "id")
}

func (self *AuthHandler) Prepare() {
	sess_role, _ := self.GetSession("userrole").(int64)
	if sess_role <= 0 {
		self.Ctx.Redirect(302, "/login")
	}

	sess_username, _ := self.GetSession("username").(string)
	sess_uid, _ := self.GetSession("userid").(int64)
	sess_email, _ := self.GetSession("useremail").(string)
	if sess_role <= 0 {
		self.Data["Userid"] = 0
		self.Data["Username"] = ""
		self.Data["Userrole"] = 0
		self.Data["Useremail"] = ""
	} else {
		self.Data["Userid"] = sess_uid
		self.Data["Username"] = sess_username
		self.Data["Userrole"] = sess_role
		self.Data["Useremail"] = sess_email
	}

	self.Data["categorys"] = models.GetAllCategory()
	self.Data["nodes"] = models.GetAllNode()
	self.Data["nodes_10s"] = models.GetAllNodeByCategoryId(0, 0, 10, "id")
	self.Data["replys_5s"] = models.GetReplyByPid(0, 0, 5, "id")
	self.Data["replys_10s"] = models.GetReplyByPid(0, 0, 10, "id")
}

func (self *RootAuthHandler) Prepare() {
	sess_role, _ := self.GetSession("userrole").(int64)

	if sess_role != 100 {
		self.Ctx.Redirect(302, "/login")
	}

	sess_username, _ := self.GetSession("username").(string)
	sess_uid, _ := self.GetSession("userid").(int64)
	sess_email, _ := self.GetSession("useremail").(string)
	if sess_role <= 0 {
		self.Data["Userid"] = 0
		self.Data["Username"] = ""
		self.Data["Userrole"] = 0
		self.Data["Useremail"] = ""
	} else {
		self.Data["Userid"] = sess_uid
		self.Data["Username"] = sess_username
		self.Data["Userrole"] = sess_role
		self.Data["Useremail"] = sess_email
	}

	self.Data["categorys"] = models.GetAllCategory()
	self.Data["nodes"] = models.GetAllNode()
	self.Data["nodes_10s"] = models.GetAllNodeByCategoryId(0, 0, 10, "id")
	self.Data["replys_5s"] = models.GetReplyByPid(0, 0, 5, "id")
	self.Data["replys_10s"] = models.GetReplyByPid(0, 0, 10, "id")
}
