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
	self.AddActionHook("SigninHandler", SigninHandle)
	self.SetStore(map[string]var{
			"title": "Single!",
			"oh":    "SigninHandler in default",
	})
	self.DoActionHook("SigninHandler")
	return self.Next()
}