RootDashboardHandler = fn(self) {
	self.AddActionHook("DashboardHandler", DashboardHandle)
	self.SetStore(map[string]var{
			"title": "#后端首页# in default",
			"oh":    "DashboardHandler in default",
	})
	self.DoActionHook("DashboardHandler")
	return self.Render("index")
}

DashboardHandle = fn() {
	str = "<DashboardHandle are Action!!!!!!>"
	println(str)
	return str
}