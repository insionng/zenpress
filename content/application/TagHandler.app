TagHandler = fn(self) {
	hook.AddAction("TagHandler", TagHandle)
	self.SetStore(map[string]var{
			"title": "#标签# in Application",
			"oh":    "TagHandler in Application",
	})
	hook.DoAction("TagHandler")
	return self.Render("index")
}

TagHandle = fn() {
	str = "<TagHandle are Action!!!!!!>"
	println(str)
	return str
}