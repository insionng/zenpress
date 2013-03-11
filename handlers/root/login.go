package root

import (
	"../../models"
	"../../utils"
	"time"
)

type LoginHandler struct {
	utils.BaseHandler
}

func (this *LoginHandler) Get() {
	this.TplNames = "login.html"
	this.Render()
}

func (this *LoginHandler) Post() {
	this.TplNames = "root/login.html"
	this.Ctx.Request.ParseForm()
	username := this.Ctx.Request.Form.Get("username")
	password := this.Ctx.Request.Form.Get("password")

	userInfo := models.GetUserByNickname(username)

	if utils.Validate_password(userInfo.Password, password) {
		var users models.User
		users.Last_login_time = time.Now()
		models.SaveUser(users)

		//登录成功设置session
		sess := this.StartSession()
		sess.Set("userid", userInfo.Id)
		sess.Set("username", userInfo.Nickname)

		this.Ctx.Redirect(302, "/")
	} else {

		this.Ctx.Redirect(302, "/login")
	}
}
