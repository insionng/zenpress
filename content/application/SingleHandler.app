SingleHandler = fn(self) {
	hook.AddAction("SingleHandler", SingleHandle)
	self.SetStore(map[string]var{
			"title": "#日志# in Application",
			"oh":    "SingleHandler in Application",
	})
	hook.DoAction("SingleHandler")
	return self.Render("index")
}

SingleHandle = fn() {
	str = "<SingleHandle are Action!!!!!!>"
	println(str)
	return str
}