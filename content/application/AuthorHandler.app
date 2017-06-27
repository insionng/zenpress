AuthorHandler = fn(self) {
	hook.AddAction("AuthorHandler", AuthorHandle)
	self.SetStore(map[string]var{
			"title": "#作者# in Application",
			"oh":    "AuthorHandler in Application",
	})
	hook.DoAction("AuthorHandler")
	return self.Render("index")
}

AuthorHandle = fn() {
	str = "<AuthorHandle are Action!!!!!!>"
	println(str)
	return str
}