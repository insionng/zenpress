package handlers

import (
	"../utils"
)

type LogoutHandler struct {
	utils.BaseHandler
}

func (this *LogoutHandler) Get() {
	//退出，销毁session
	sess := this.StartSession()
	sess.Delete("userid")
	sess.Delete("username")

	this.Ctx.Redirect(302, "/")

}
