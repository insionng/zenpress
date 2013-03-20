package root

import (
	"../../libs"
	"../../models"
	"../../utils"
)

type RLoginHandler struct {
	libs.BaseHandler
}

func (self *RLoginHandler) Get() {

	sess_userrole, _ := self.GetSession("userrole").(int64)
	//如果未登录root
	if sess_userrole != -1000 {
		self.TplNames = "root/login.html"
		self.Render()
	} else { //如果已登录root
		self.Ctx.Redirect(302, "/root")
	}

}

func (self *RLoginHandler) Post() {
	self.TplNames = "root/login.html"
	self.Ctx.Request.ParseForm()
	username := self.Ctx.Request.Form.Get("username")
	password := self.Ctx.Request.Form.Get("password")

	userInfo := models.GetUserByNickname(username)

	if utils.Validate_password(userInfo.Password, password) && userInfo.Role == -1000 {

		//登录成功设置session
		self.SetSession("userid", userInfo.Id)
		self.SetSession("username", userInfo.Nickname)
		self.SetSession("userrole", userInfo.Role)
		self.SetSession("useremail", userInfo.Email)

		self.Ctx.Redirect(302, "/root")
	} else {

		self.Ctx.Redirect(302, "/root/login")
	}
}
