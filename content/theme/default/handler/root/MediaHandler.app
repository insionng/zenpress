RootMediaHandler = fn(self) {
	hook.AddAction("MediaHandler", MediaHandle)
	self.SetStore(map[string]var{
			"title": "#首页# in Application",
			"oh":    "MediaHandler in Application",
	})
	hook.DoAction("MediaHandler")
	return self.Render("index")
}

MediaHandle = fn() {
	str = "<MediaHandle are Action!!!!!!>"
	println(str)
	return str
}