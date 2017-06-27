RootCommentHandler = fn(self) {
	hook.AddAction("CommentHandler", CommentHandle)
	self.SetStore(map[string]var{
			"title": "#首页# in Application",
			"oh":    "CommentHandler in Application",
	})
	hook.DoAction("CommentHandler")
	return self.Render("index")
}

CommentHandle = fn() {
	str = "<CommentHandle are Action!!!!!!>"
	println(str)
	return str
}