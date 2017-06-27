RootDashboardHandler = fn(self) {
	hook.AddAction("DashboardHandler", DashboardHandle)
	self.SetStore(map[string]var{
			"title": "#全局22首页# in Application",
			"oh":    "DashboardHandler in Application",
	})
	hook.DoAction("DashboardHandler")
	return self.Render("index")
}

DashboardHandle = fn() {
	str = "<DashboardHandle are Action!!!!!!>"
	println(str)
	return str
}