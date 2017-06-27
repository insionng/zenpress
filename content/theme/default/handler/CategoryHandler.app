CategoryHandler = fn(self) {
	self.AddActionHook("CategoryHandler", CategoryHandle)
	self.SetStore(map[string]var{
			"title": "Category!",
			"oh":    "CategoryHandler in default",
	})
	self.DoActionHook("CategoryHandler")
	return self.Render("index")
}

CategoryHandle = fn() {
	str = "<CategoryHandle are Action!!!!!!>"
	println(str)
	return str
}