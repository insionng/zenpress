SearchHandler = fn(self) {
	hook.AddAction("SearchHandler", SearchHandle)
	self.SetStore(map[string]var{
			"title": "#搜索# in default",
			"oh":    "SearchHandler in default",
	})
	hook.DoAction("SearchHandler")
	return self.Render("index")
}

SearchHandle = fn() {
	str = "<SearchHandle are Action!!!!!!>"
	println(str)
	return str
}