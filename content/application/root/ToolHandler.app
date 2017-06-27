RootToolHandler = fn(self) {
	hook.AddAction("ToolHandler", ToolHandle)
	self.SetStore(map[string]var{
			"title": "#首页# in Application",
			"oh":    "ToolHandler in Application",
	})
	hook.DoAction("ToolHandler")
	return self.Render("index")
}

ToolHandle = fn() {
	str = "<ToolHandle are Action!!!!!!>"
	println(str)
	return str
}