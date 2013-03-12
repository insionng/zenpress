package handlers

import (
	"../libs"
)

type LogoutHandler struct {
	libs.BaseHandler
}

func (this *LogoutHandler) Get() {
	//退出，销毁session
	sess := this.StartSession()
	sess.Delete("userid")
	sess.Delete("username")

	this.Ctx.Redirect(302, "/")

}
