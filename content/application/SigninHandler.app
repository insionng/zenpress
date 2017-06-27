SigninHandler = fn(self) {
	switch self.Request.Method == makross.GET {
	case true:
		fmt.Println("Request Method Is [GET]!")
		self.AddActionHook("SigninHandler", Ftest)
	default:
		fmt.Println("Request Method Is Not [GET]!")
	}
	self.AddActionHook("SigninHandler", SigninHandle)
	self.SetStore(map[string]var{
			"title": "Single!",
			"oh":    "SigninHandler in default",
	})
	self.DoActionHook("SigninHandler")
	return self.Next()
}

Ftest = fn() {
	str = "Request Method Is [GET]!"
	println(str)
	return str
}

SigninHandle = fn() {
	str = "<SigninHandle are Action!!!!!!>"
	println(str)
	return str
}