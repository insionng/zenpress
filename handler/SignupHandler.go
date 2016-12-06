package handler

import (
	"errors"
	"fmt"
	"github.com/insionng/macross"
	"github.com/insionng/zenpress/helper"
	"github.com/insionng/zenpress/models"
)

func GetSignupHandler(self *macross.Context) error {
	return self.Render("signup")
}

func PostSignupHandler(self *macross.Context) error {

	data := make(map[string]interface{})

	username := self.Args("username").String()
	password := self.Args("password").String()
	usererr := helper.CheckUsername(username)

	if usererr == false {
		e := errors.New("Username error, Please to again")
		data["UsernameErr"] = e.Error()
		return e
	}

	passerr := helper.CheckPassword(password)
	if passerr == false {
		e := errors.New("Password error, Please to again")
		data["PasswordErr"] = e.Error()
		return e
	}

	pwd := helper.EncryptHash(password, nil)

	//now := torgo.Date(time.Now(), "Y-m-d H:i:s")

	userInfo := models.CheckUserByNickname(username)

	//fmt.Println(userInfo.Nickname)

	//  检查该用户是否已经被注册
	if userInfo.Nickname == "" {

		//注册用户
		regErr := models.AddUser(username+"@your.com", username, "", pwd, 1)
		fmt.Println("reg:s")
		fmt.Println(regErr)
		fmt.Println("reg:e ")
		//注册成功设置session
		//	self.SetSession("userid", userInfo.Id)
		//	self.SetSession("username", userInfo.Nickname)
		//	self.SetSession("userrole", userInfo.Role)
		//	self.SetSession("useremail", userInfo.Email)
		return self.Redirect("/signin/", macross.StatusFound)

	} else {
		e := errors.New("User already exists")
		data["UsernameErr"] = e.Error()
		return e
	}
	self.SetStore(data)
	return self.Render("signup")
}
