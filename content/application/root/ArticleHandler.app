RootArticleHandler = fn(self) {
	hook.AddAction("ArticleHandler", ArticleHandle)
	self.SetStore(map[string]var{
			"title": "#首页# in Application",
			"oh":    "ArticleHandler in Application",
	})
	hook.DoAction("ArticleHandler")
	return self.Render("index")
}

ArticleHandle = fn() {
	str = "<ArticleHandle are Action!!!!!!>"
	println(str)
	return str
}