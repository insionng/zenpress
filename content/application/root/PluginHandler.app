RootPluginHandler = fn(self) {
	hook.AddAction("PluginHandler", PluginHandle)
	self.SetStore(map[string]var{
			"title": "#首页# in Application",
			"oh":    "PluginHandler in Application",
	})
	hook.DoAction("PluginHandler")
	return self.Render("index")
}

PluginHandle = fn() {
	str = "<PluginHandle are Action!!!!!!>"
	println(str)
	return str
}