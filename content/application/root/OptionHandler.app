RootOptionHandler = fn(self) {
	hook.AddAction("OptionHandler", OptionHandle)
	self.SetStore(map[string]var{
			"title": "#首页# in Application",
			"oh":    "OptionHandler in Application",
	})
	hook.DoAction("OptionHandler")
	return self.Render("index")
}

OptionHandle = fn() {
	str = "<OptionHandle are Action!!!!!!>"
	println(str)
	return str
}