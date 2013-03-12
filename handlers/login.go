package handlers

import (
	"../libs"
	"../models"
	"../utils"
)

type LoginHandler struct {
	libs.BaseHandler
}

func (this *LoginHandler) Get() {

	sess := this.StartSession()
	sess_username, _ := sess.Get("username").(string)
	//如果未登录
	if sess_username == "" {
		this.TplNames = "login.html"
		this.Render()
	} else { //如果已登录
		this.Ctx.Redirect(302, "/")
	}

}

func (this *LoginHandler) Post() {
	this.TplNames = "login.html"
	this.Ctx.Request.ParseForm()
	username := this.Ctx.Request.Form.Get("username")
	password := this.Ctx.Request.Form.Get("password")

	userInfo := models.GetUserByNickname(username)

	if utils.Validate_password(userInfo.Password, password) {

		//登录成功设置session
		sess := this.StartSession()
		sess.Set("userid", userInfo.Id)
		sess.Set("username", userInfo.Nickname)
		sess.Set("userrole", userInfo.Role)
		sess.Set("useremail", userInfo.Email)

		this.Ctx.Redirect(302, "/")
	} else {

		this.Ctx.Redirect(302, "/login")
	}
}
