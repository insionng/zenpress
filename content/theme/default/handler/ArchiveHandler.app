ArchiveHandler = fn(self) {
	hook.AddAction("ArchiveHandler", ArchiveHandle)
	self.SetStore(map[string]var{
			"title": "#归档# in default",
			"oh":    "ArchiveHandler in default",
	})
	hook.DoAction("ArchiveHandler")
	return self.Render("index")
}

ArchiveHandle = fn() {
	str = "<ArchiveHandle are Action!!!!!!>"
	println(str)
	return str
}