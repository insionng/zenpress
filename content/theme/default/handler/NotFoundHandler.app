NotFoundHandler = fn(self) {
	hook.AddAction("NotFoundHandler", NotFoundHandle)
	self.SetStore(map[string]var{
			"title": "#缺失页面# in default",
			"oh":    "NotFoundHandler in default",
	})
	hook.DoAction("NotFoundHandler")
	return self.Render("index")
}

NotFoundHandle = fn() {
	str = "<NotFoundHandle are Action!!!!!!>"
	println(str)
	return str
}