package libs

import (
	"../models"
	"github.com/insionng/torgo"
	//"../torgo"
)

var (
	sess_username string
	sess_uid      int64
	sess_role     int64
	sess_email    string
)

type BaseHandler struct {
	torgo.Handler
}

type AuthHandler struct {
	BaseHandler
}

type RootAuthHandler struct {
	BaseHandler
}

type RootHandler struct {
	BaseHandler
}

//用户等级划分：正数是普通用户，负数是管理员各种等级划分，为0则尚未注册
func (self *BaseHandler) Prepare() {
	sess_username, _ = self.GetSession("username").(string)
	sess_uid, _ = self.GetSession("userid").(int64)
	sess_role, _ = self.GetSession("userrole").(int64)
	sess_email, _ = self.GetSession("useremail").(string)
	if sess_role == 0 {
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
	self.Data["topics_5s"] = models.GetAllTopic(0, 5, "id")
	self.Data["topics_10s"] = models.GetAllTopic(0, 10, "id")
	self.Data["nodes_10s"] = models.GetAllNodeByCid(0, 0, 10, "id")
	self.Data["replys_5s"] = models.GetReplyByPid(0, 0, 5, "id")
	self.Data["replys_10s"] = models.GetReplyByPid(0, 0, 10, "id")
}

//会员或管理员前台权限认证
func (self *AuthHandler) Prepare() {
	self.BaseHandler.Prepare()

	if sess_role == 0 {
		self.Ctx.Redirect(302, "/login")
	}
}

//管理员前台权限认证
func (self *RootAuthHandler) Prepare() {
	self.BaseHandler.Prepare()

	if sess_role != -1000 {
		self.Ctx.Redirect(302, "/login")
	}
}

//管理员后台后台认证
func (self *RootHandler) Prepare() {
	self.BaseHandler.Prepare()

	if sess_role != -1000 {
		self.Ctx.Redirect(302, "/root/login")
	}
}
