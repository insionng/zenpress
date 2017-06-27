DateHandler = fn(self) {
	hook.AddAction("DateHandler", DateHandle)
	self.SetStore(map[string]var{
			"title": "#日期# in Application",
			"oh":    "DateHandler in Application",
	})
	hook.DoAction("DateHandler")
	return self.Render("index")
}

DateHandle = fn() {
	str = "<DateHandle are Action!!!!!!>"
	println(str)
	return str
}