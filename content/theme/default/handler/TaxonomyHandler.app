TaxonomyHandler = fn(self) {
	hook.AddAction("TaxonomyHandler", TaxonomyHandle)
	self.SetStore(map[string]var{
			"title": "#类别# in default",
			"oh":    "TaxonomyHandler in default",
	})
	hook.DoAction("TaxonomyHandler")
	return self.Render("index")
}

TaxonomyHandle = fn() {
	str = "<TaxonomyHandle are Action!!!!!!>"
	println(str)
	return str
}