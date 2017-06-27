PageHandler = fn(self) {
	hook.AddAction("PageHandle", PageHandle)
	self.SetStore(map[string]var{
			"title": "#页面# in Application",
			"oh":    "Page in Application",
	})
	hook.DoAction("PageHandle")
	return self.Render("index")
}

PageHandle = fn() {
	str = "<PageHandle are Action!!!!!!>"
	println(str)
	return str
}