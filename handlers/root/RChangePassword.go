package root

import (
	"../../libs"
	//"../../models"
	//"../../utils"

)

type RChangePasswordHandler struct {
	libs.RootHandler
}

func (self *RChangePasswordHandler) Get() {
	self.TplNames = "root/change_password.html"
	self.Render()
}
