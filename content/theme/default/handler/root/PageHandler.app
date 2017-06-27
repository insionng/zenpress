RootPageHandler = fn(self) {
	hook.AddAction("PageHandler", PageHandle)
	self.SetStore(map[string]var{
			"title": "#首页# in Application",
			"oh":    "PageHandler in Application",
	})
	hook.DoAction("PageHandler")
	return self.Render("index")
}

PageHandle = fn() {
	str = "<PageHandle are Action!!!!!!>"
	println(str)
	return str
}