package libs

import (
	"../models"
	"github.com/insionng/torgo"
	//"lihuashu.com/insionng/torgo"
	//"github.com/astaxie/beego"
)

type BaseHandler struct {
	torgo.Controller
}

type AuthHandler struct {
	BaseHandler
}

type RootAuthHandler struct {
	BaseHandler
}

func (this *BaseHandler) Prepare() {
	sess := this.StartSession()
	sess_username, _ := sess.Get("username").(string)
	sess_uid, _ := sess.Get("userid").(int64)
	sess_role, _ := sess.Get("userrole").(int64)
	sess_email, _ := sess.Get("useremail").(string)

	if sess_uid == 0 {
		this.Data["Userid"] = 0
		this.Data["Username"] = ""
		this.Data["Userrole"] = 0
		this.Data["Useremail"] = ""
	} else {
		this.Data["Userid"] = sess_uid
		this.Data["Username"] = sess_username
		this.Data["Userrole"] = sess_role
		this.Data["Useremail"] = sess_email
	}

	this.Data["categorys"] = models.GetAllCategory()
	this.Data["nodes"] = models.GetAllNode()
	this.Data["nodes_10s"] = models.GetAllNodeByCategoryId(0, 0, 10, "id")
	this.Data["replys_5s"] = models.GetReplyByPid(0, 0, 5, "id")
	this.Data["replys_10s"] = models.GetReplyByPid(0, 0, 10, "id")
}

func (this *AuthHandler) Prepare() {
	sess := this.StartSession()
	sess_uid, _ := sess.Get("userid").(int64)
	if sess_uid == 0 {
		this.Ctx.Redirect(302, "/login")
	}

	sess_username, _ := sess.Get("username").(string)
	sess_role, _ := sess.Get("userrole").(int64)
	sess_email, _ := sess.Get("useremail").(string)

	if sess_uid == 0 {
		this.Data["Userid"] = 0
		this.Data["Username"] = ""
		this.Data["Userrole"] = 0
		this.Data["Useremail"] = ""
	} else {
		this.Data["Userid"] = sess_uid
		this.Data["Username"] = sess_username
		this.Data["Userrole"] = sess_role
		this.Data["Useremail"] = sess_email
	}
	this.Data["categorys"] = models.GetAllCategory()
	this.Data["nodes"] = models.GetAllNode()
	this.Data["nodes_10s"] = models.GetAllNodeByCategoryId(0, 0, 10, "id")
	this.Data["replys_5s"] = models.GetReplyByPid(0, 0, 5, "id")
	this.Data["replys_10s"] = models.GetReplyByPid(0, 0, 10, "id")
}

func (this *RootAuthHandler) Prepare() {
	sess := this.StartSession()
	sess_role, _ := sess.Get("userrole").(int64)

	if sess_role != 100 {
		this.Ctx.Redirect(302, "/login")
	}

	sess_username, _ := sess.Get("username").(string)
	sess_uid, _ := sess.Get("userid").(int64)
	sess_email, _ := sess.Get("useremail").(string)

	if sess_uid == 0 {
		this.Data["Userid"] = 0
		this.Data["Username"] = ""
		this.Data["Userrole"] = 0
		this.Data["Useremail"] = ""
	} else {
		this.Data["Userid"] = sess_uid
		this.Data["Username"] = sess_username
		this.Data["Userrole"] = sess_role
		this.Data["Useremail"] = sess_email
	}
	this.Data["categorys"] = models.GetAllCategory()
	this.Data["nodes"] = models.GetAllNode()
	this.Data["nodes_10s"] = models.GetAllNodeByCategoryId(0, 0, 10, "id")
	this.Data["replys_5s"] = models.GetReplyByPid(0, 0, 5, "id")
	this.Data["replys_10s"] = models.GetReplyByPid(0, 0, 10, "id")
}
