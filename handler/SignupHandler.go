package handler

import (
	"errors"
	"github.com/insionng/macross"
	"github.com/insionng/zenpress/helper"
	"github.com/insionng/zenpress/models"
	"github.com/macross-contrib/session"
	"log"
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

	userInfo := models.CheckUserByNickname(username)

	//fmt.Println(userInfo.Nickname)

	//  检查该用户是否已经被注册
	if userInfo.Nickname == "" {

		//注册用户
		err := models.AddUser(username+"@your.com", username, "", pwd, 1)
		if err != nil {
			errors.New("AddUser Errors")
		}

		//注册成功设置Session
		sess := session.GetStore(self)
		err = sess.Set("username", userInfo.Nickname)
		if err != nil {
			log.Printf("sess.set %v \n", err)
		}

		return self.Redirect("/signin/", macross.StatusFound)

	} else {
		e := errors.New("User already exists")
		data["UsernameErr"] = e.Error()
		return e
	}
	self.SetStore(data)
	return self.Render("signup")
}
