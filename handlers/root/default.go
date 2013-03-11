package root

import (
	//"../../models"
	"../../utils"
)

type RootHandler struct {
	utils.AuthHandler
}

func (this *RootHandler) Get() {

	this.TplNames = "root/index.html"

	this.Render()

}
