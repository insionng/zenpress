package handlers

import (
	"fmt"
	"toropress/helper"
	"toropress/libs"
	"toropress/models"
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
	username := self.GetString("username")
	password := self.GetString("password")

	if username != "" && password != "" {

		if userInfo := models.GetUserByNickname(username); userInfo.Password != "" {
//			fmt.Println("uname: " + userInfo.Nickname)
//			fmt.Println("pwd: " + userInfo.Password)
//			err := helper.Validate_password(userInfo.Password, password)
//
//			fmt.Println(err)

			if helper.Validate_password(userInfo.Password, password) {

				//登录成功设置session
				self.SetSession("userid", userInfo.Id)
				self.SetSession("username", userInfo.Nickname)
				self.SetSession("userrole", userInfo.Role)
				self.SetSession("useremail", userInfo.Email)

				fmt.Println("============userid===============")
				fmt.Println(self.GetSession("userid"))
				fmt.Println("============username===============")
				fmt.Println(self.GetSession("username"))




				self.Ctx.Redirect(302, "/")
			} else {

				self.Ctx.Redirect(302, "/login")
			}
		} else {

			self.Ctx.Redirect(302, "/login")
		}
	} else {

		self.Ctx.Redirect(302, "/login")
	}
}
