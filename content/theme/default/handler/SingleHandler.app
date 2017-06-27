SingleHandler = fn(self) {
	self.AddActionHook("SingleHandler", SingleHandle)
	self.SetStore(map[string]var{
			"title": "Single!",
			"oh":    "SingleHandler in default",
	})
	self.DoActionHook("SingleHandler")
	return self.Render("single")
}

SingleHandle = fn() {
	str = "<SingleHandle are Action!!!!!!>"
	println(str)
	return str
}