package handlers

import (
	"../libs"
	"../models"
	"../utils"
)

type LoginHandler struct {
	libs.BaseHandler
}

func (self *LoginHandler) Get() {

	sess_username, _ := self.GetSession("username").(string)
	//如果未登录
	if sess_username == "" {
		self.TplNames = "login.html"
		self.Render()
	} else { //如果已登录
		self.Ctx.Redirect(302, "/")
	}

}

func (self *LoginHandler) Post() {
	self.TplNames = "login.html"
	self.Ctx.Request.ParseForm()
	username := self.Ctx.Request.Form.Get("username")
	password := self.Ctx.Request.Form.Get("password")

	userInfo := models.GetUserByNickname(username)

	if utils.Validate_password(userInfo.Password, password) {

		//登录成功设置session
		self.SetSession("userid", userInfo.Id)
		self.SetSession("username", userInfo.Nickname)
		self.SetSession("userrole", userInfo.Role)
		self.SetSession("useremail", userInfo.Email)

		self.Ctx.Redirect(302, "/")
	} else {

		self.Ctx.Redirect(302, "/login")
	}
}
