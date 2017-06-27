SigninHandler = fn(self) {
	switch self.Request.Method {
	case makross.POST:
		fmt.Println("SigninHandler Request Method Is [POST]!")
		PostSigninHandler(self)
	default:
		fmt.Println("SigninHandler Request Method Is [GET]!")
	}
	return self.Next()
}

SigninHandle = fn() {
	str = "<SigninHandle are Action in post!!!!!!>"
	println(str)
	return str
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