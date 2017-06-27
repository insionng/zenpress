IndexHandler = fn(self) {
	if self.Request.Method == makross.GET {
		fmt.Println("Request Method Is [GET]!")
	} else {
		fmt.Println("Request Method Is Not [GET]!")
	}
	self.AddActionHook("IndexHandler", IndexHandle)
	self.SetStore(map[string]var{
			"title": "#默认主题首页#",
			"oh":    "IndexHandler in default",
	})
	self.DoActionHook("IndexHandler")
	//self.AddFilterHook("index_template", indexFixTpl)
	return self.Render("index")
}

IndexHandle = fn() {
	str = "<IndexHandle are Action in IndexHandler!!!!!>"
	println(str)
	return str
}

indexFixTpl = fn(b) {
	return []byte(fmt.Sprintf("<h1>过滤器</h1>%s", b))
}