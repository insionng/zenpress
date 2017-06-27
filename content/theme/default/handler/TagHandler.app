TagHandler = fn(self) {
	hook.AddAction("TagHandler", TagHandle)
	self.SetStore(map[string]var{
			"title": "#标签# in default",
			"oh":    "TagHandler in default",
	})
	hook.DoAction("TagHandler")
	return self.Render("index")
}

TagHandle = fn() {
	str = "<TagHandle are Action!!!!!!>"
	println(str)
	return str
}