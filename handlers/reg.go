package handlers

import (
	"../libs"
	"../models"
	"../utils"
	"fmt"
	"regexp"
)

type RegHandler struct {
	libs.BaseHandler
}

func (this *RegHandler) Get() {
	this.TplNames = "reg.html"
	this.Render()
}

func (this *RegHandler) Post() {
	this.TplNames = "reg.html"
	this.Ctx.Request.ParseForm()
	username := this.Ctx.Request.Form.Get("username")
	password := this.Ctx.Request.Form.Get("password")
	usererr := checkUsername(username)

	fmt.Println(usererr)
	if usererr == false {
		this.Data["UsernameErr"] = "Username error, Please to again"
		return
	}

	passerr := checkPassword(password)
	if passerr == false {
		this.Data["PasswordErr"] = "Password error, Please to again"
		return
	}

	pwd := utils.Encrypt_password(password, nil)

	//now := torgo.Date(time.Now(), "Y-m-d H:i:s")

	userInfo := models.GetUserByNickname(username)

	if userInfo.Nickname == "" {
		models.AddUser(username+"@insion.co", username, pwd, 0)

		//登录成功设置session
		sess := this.StartSession()
		sess.Set("userid", userInfo.Id)
		sess.Set("username", userInfo.Nickname)
		sess.Set("userrole", userInfo.Role)
		sess.Set("useremail", userInfo.Email)
		this.Ctx.Redirect(302, "/login")
	} else {
		this.Data["UsernameErr"] = "User already exists"
	}
	this.Render()
}

func checkPassword(password string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", password); !ok {
		return false
	}
	return true
}

func checkUsername(username string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", username); !ok {
		return false
	}
	return true
}
