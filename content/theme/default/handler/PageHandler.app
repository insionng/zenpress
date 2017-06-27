PageHandler = fn(self) {
	self.SetStore(map[string]var{
			"title": "#页面# in default",
			"oh":    "Page in default",
	})
	self.AddFilterHook("index_template", fixTpl)
	return self.String("Page")
}

PageHandle = fn() {
	str = "<PageHandle are Action!!!!!!>"
	println(str)
	return str
}

fixTpl = fn(b) {
	return []byte(fmt.Sprintf("@#page## %s ##page#@", b))
}