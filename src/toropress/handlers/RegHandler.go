package handlers

import (
	"fmt"
	"toropress/helper"
	"toropress/libs"
	"toropress/models"
)

type RegHandler struct {
	libs.BaseHandler
}

func (self *RegHandler) Get() {
	self.TplNames = "reg.html"
	self.Render()
}

func (self *RegHandler) Post() {
	self.TplNames = "reg.html"
	self.Ctx.Request.ParseForm()
	username := self.Ctx.Request.Form.Get("username")
	password := self.Ctx.Request.Form.Get("password")
	usererr := helper.CheckUsername(username)


	if usererr == false {
		self.Data["UsernameErr"] = "Username error, Please to again"
		return
	}

	passerr := helper.CheckPassword(password)
	if passerr == false {
		self.Data["PasswordErr"] = "Password error, Please to again"
		return
	}

	pwd := helper.Encrypt_password(password, nil)

	//now := torgo.Date(time.Now(), "Y-m-d H:i:s")

	userInfo := models.CheckUserByNickname(username)
	fmt.Println("=============================")
	fmt.Println(userInfo.Nickname)

	//  检查该用户是否已经被注册
		if userInfo.Nickname == "" {

			//注册用户
			regErr := models.AddUser(username+"@insion.co", username, "", pwd, 1)
			fmt.Println("reg:s")
			fmt.Println(regErr)
			fmt.Println("reg:e ")
			//注册成功设置session
		//	self.SetSession("userid", userInfo.Id)
		//	self.SetSession("username", userInfo.Nickname)
		//	self.SetSession("userrole", userInfo.Role)
		//	self.SetSession("useremail", userInfo.Email)
			self.Ctx.Redirect(302, "/login")




		} else {
			self.Data["UsernameErr"] = "User already exists"
		}
	self.Render()
}
