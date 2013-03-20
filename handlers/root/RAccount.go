package root

import (
	"../../libs"
	//"../../models"
	//"../../utils"

)

type RAccountHandler struct {
	libs.RootHandler
}

func (self *RAccountHandler) Get() {
	self.TplNames = "root/account.html"
	self.Render()
}
