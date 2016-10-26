package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/insionng/vodka"
	"github.com/insionng/zenpress/helper"
	"github.com/insionng/zenpress/models"
)

func SignupGetHandler(self vodka.Context) error {
	return self.Render(http.StatusOK, "signup.html")
}

func SignupPostHandler(self vodka.Context) error {

	data := make(map[string]interface{})

	username := self.FormValue("username")
	password := self.FormValue("password")
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
		regErr := models.AddUser(username+"@insion.co", username, "", pwd, 1)
		fmt.Println("reg:s")
		fmt.Println(regErr)
		fmt.Println("reg:e ")
		//注册成功设置session
		//	self.SetSession("userid", userInfo.Id)
		//	self.SetSession("username", userInfo.Nickname)
		//	self.SetSession("userrole", userInfo.Role)
		//	self.SetSession("useremail", userInfo.Email)
		return self.Redirect(302, "/signin/")

	} else {
		e := errors.New("User already exists")
		data["UsernameErr"] = e.Error()
		return e
	}
	self.SetStore(data)
	return self.Render(http.StatusOK, "signup.html")
}
