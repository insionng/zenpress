RootLinkHandler = fn(self) {
	hook.AddAction("LinkHandler", LinkHandle)
	self.SetStore(map[string]var{
			"title": "#首页# in Application",
			"oh":    "LinkHandler in Application",
	})
	hook.DoAction("LinkHandler")
	return self.Render("index")
}

LinkHandle = fn() {
	str = "<LinkHandle are Action!!!!!!>"
	println(str)
	return str
}