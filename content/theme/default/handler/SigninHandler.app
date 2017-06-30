SigninHandler = fn(self) {
	switch self.Request.Method {
	case makross.POST:
		fmt.Println("SigninHandler Request Method Is [POST]!")
		PostSigninHandler(self)
	case makross.GET:
		fmt.Println("SigninHandler Request Method Is [GET]!")
		GetSigninHandler(self)
	}
	return self.Next()
}

SigninHandle = fn() {
	str = "<SigninHandle are Action in post!!!!!!>"
	println(str)
	return str
}

GetSigninHandler = fn(self) {

	IsSignin = false
	SignedUser = self.Session.Get("SignedUser")
	if SignedUser != nil {
		IsSignin = true
	}

	tplName = "signin"
	self.Set("catpage", "SigninHandler")

	remember, err = self.GetCookie("remember")

	if IsSignin { //如果已登录
		next = self.Args("next").String()
		if len(next) > 0 {
			return self.Redirect(next)
		}
		return self.Redirect("/")
	} else { //如果未登录
		if (remember != nil) && (err == nil) {			
			if remember.Value == "true" {
				self.Set("remember", "true")
			} else {
				self.Set("remember", nil)
			}
		}
	}

	return self.Render(tplName)
}

PostSigninHandler = fn(self) {
	IsCaptcha = true

	TplNames = "signin"
	cpt = captcha.NewCaptcha(&captcha.Options{})
	allow = false
	if IsCaptcha {
		cpt = captcha.Store(self)
		allow = cpt.VerifyReq(self)
	}
	if IsCaptcha && (!allow) {
		if len(self.Args(cpt.FieldCaptchaName).String()) > 0 {
			self.Flash.Error("验证码错误~")
		} else {
			self.Flash.Error("验证码为空~")
		}
		return self.Render(TplNames)
	}

	cc = cache.Store(self)

	//Secret = helper.MD5(self.Req.UserAgent() + helper.AesConstKey)
	self.Set("catpage", "SigninHandler")

	password = self.Args("password").String()
	self.Set("tmppassword", password)
	self.Set("tmpemail", self.Args("email").String())
	remember = self.Args("remember").String()

	if len(password) == 0 {
		self.Flash.Error("密码为空~")
		return self.Render(TplNames)
	}

	if helper.CheckPassword(password) == false {
		self.Flash.Error("密码含有非法字符或密码过短(至少4~30位密码)!")
		return self.Render(TplNames)
	}

	err = nil
	usr = new(model.User)
	email, username = ""
	isEmail = strings.Contains(self.Args("email").String(), "@")
	if isEmail {
		email = self.Args("email").String()
		if len(email) == 0 {
			self.Flash.Error("EMAIL为空~")
			return self.Render(TplNames)
		}

		if helper.CheckEmail(email) == false {
			self.Flash.Error("Email格式不合符规格~")
			return self.Render(TplNames)
		}

		usr, err = models.GetUserByEmail(email)
	} else {
		username = self.Args("email").String()
		if len(username) == 0 {
			self.Flash.Error("用户名称为空~")
			return self.Render(TplNames)
		}

		if helper.CheckUsername(username) == false {
			self.Flash.Error("用户名称格式不合符规格~")
			return self.Render(TplNames)
		}

		usr, err = models.GetUserByUsername(username)
	}

	if (usr != nil) && (err == nil) {

		if helper.ValidateHash(usr.Password, password) {

			//登录成功设置session
			self.Session.Set("SignedUserID", usr.Id)
			self.Session.Set("SignedUserName", usr.Username)
			self.Session.Set("SignedUser", usr)

			self.Set("IsSigned", true)
			self.Set("IsRoot", (usr.Role == -1000))
			self.Set("SignedUser", usr)
			self.Set("SignedUserID", usr.Id)
			self.Set("SignedUserName", usr.Username)
			cc.Set(fmt.Sprintf("SignedUser:%v", usr.Id), usr, 60*60*24)
			models.PutSignin2User(usr.Id, time.Now().Unix(), usr.SigninCount+1, self.RemoteAddress())

			//设置cookie
			cookie = new(makross.Cookie)
			cookie.SetName("remember")
			if remember == "true" {
				cookie.SetValue("true")
				cookie.SetExpire(time.Now().Add(time.Duration(31190400))) //361 days
				//使用flower作本地存储时的Email别名
				//self.SetSuperSecureCookie(Secret, "flower", usr.Email, 31190400)
			} else {
				cookie.SetValue("false") //取消记录
				cookie.SetExpire(time.Now().Add(time.Duration(-1)))
				//self.SetSuperSecureCookie(Secret, "flower", "", 3600) //删除数据
			}
			self.SetCookie(cookie)

			next = self.Args("next").String()
			if next != "" {
				return self.Redirect(next)
			}
			return self.Redirect("/")

		} else {
			self.Flash.Error("密码无法通过校验~")
			return self.Render(TplNames)
		}
	} else {
		self.Flash.Error("该账号不存在~")
		return self.Render(TplNames)
	}

	return self.Render(TplNames)
}
