IndexHandler = fn(self) {
	
	self.AddActionHook("IndexHandler", IndexHandle)
	self.SetStore(map[string]var{
			"title": "#首页# in Application",
			"oh":    "IndexHandler in Application",
	})
	self.DoActionHook("IndexHandler")

	self.AddFilterHook("index_template", fn(b) {
		return []byte(fmt.Sprintf("@#### %s ###@", b))
	})
	
	return self.Render("index")
}

IndexHandle = fn() {
	str = "<IndexHandle are Action!!!!!!>"
	println(str)
	return str
}

fixText = fn(b) {
	return []byte(fmt.Sprintf("@#### %s ###@", b))
}