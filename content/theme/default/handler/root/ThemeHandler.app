RootThemeHandler = fn(self) {
	hook.AddAction("ThemeHandler", ThemeHandle)
	self.SetStore(map[string]var{
			"title": "#首页# in Application",
			"oh":    "ThemeHandler in Application",
	})
	hook.DoAction("ThemeHandler")
	return self.Render("index")
}

ThemeHandle = fn() {
	str = "<ThemeHandle are Action!!!!!!>"
	println(str)
	return str
}