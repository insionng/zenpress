TaxonomyHandler = fn(self) {
	hook.AddAction("TaxonomyHandler", TaxonomyHandle)
	self.SetStore(map[string]var{
			"title": "#类别# in Application",
			"oh":    "TaxonomyHandler in Application",
	})
	hook.DoAction("TaxonomyHandler")
	return self.Render("index")
}

TaxonomyHandle = fn() {
	str = "<TaxonomyHandle are Action!!!!!!>"
	println(str)
	return str
}