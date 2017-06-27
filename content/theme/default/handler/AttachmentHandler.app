AttachmentHandler = fn(self) {
	hook.AddAction("AttachmentHandler", AttachmentHandle)
	self.SetStore(map[string]var{
			"title": "#附件# in default",
			"oh":    "AttachmentHandler in default",
	})
	hook.DoAction("AttachmentHandler")
	return self.Render("index")
}

AttachmentHandle = fn() {
	str = "<AttachmentHandle are Action!!!!!!>"
	println(str)
	return str
}