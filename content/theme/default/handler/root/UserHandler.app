RootUserHandler = fn(self) {
	hook.AddAction("UserHandler", UserHandle)
	self.SetStore(map[string]var{
			"title": "#首页# in Application",
			"oh":    "UserHandler in Application",
	})
	hook.DoAction("UserHandler")
	return self.Render("index")
}

UserHandle = fn() {
	str = "<UserHandle are Action!!!!!!>"
	println(str)
	return str
}