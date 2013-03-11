package utils

import (
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
	sess_uid := sess.Get("userid")
	sess_username := sess.Get("username")
	sess_role := sess.Get("userrole")
	if sess_uid == nil {
		this.Data["Userid"] = nil
		this.Data["Username"] = nil
		this.Data["Userrole"] = nil
	} else {
		this.Data["Userid"] = sess_uid
		this.Data["Username"] = sess_username
		this.Data["Userrole"] = sess_role
	}
}

func (this *AuthHandler) Prepare() {
	sess := this.StartSession()
	sess_uid := sess.Get("userid")
	if sess_uid == nil {
		this.Ctx.Redirect(302, "/login")
	}
}

func (this *RootAuthHandler) Prepare() {
	sess := this.StartSession()
	sess_uid := sess.Get("userid")
	sess_role := sess.Get("userrole")
	if sess_uid == nil || sess_role != 100 {
		this.Ctx.Redirect(302, "/login")
	}
}
