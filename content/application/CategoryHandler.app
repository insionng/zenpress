CategoryHandler = fn(self) {
	hook.AddAction("CategoryHandler", CategoryHandle)
	self.SetStore(map[string]var{
			"title": "#分类# in Application",
			"oh":    "CategoryHandler in Application",
	})
	hook.DoAction("CategoryHandler")
	return self.Render("index")
}

CategoryHandle = fn() {
	str = "<CategoryHandle are Action!!!!!!>"
	println(str)
	return str
}