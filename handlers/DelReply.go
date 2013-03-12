package handlers

import (
	"../libs"
	"../models"
	"strconv"
)

type DeleteReplyHandler struct {
	libs.RootAuthHandler
}

func (this *DeleteReplyHandler) Get() {
	rid, _ := strconv.Atoi(this.Ctx.Params[":rid"])
	models.DelReply(rid)
	this.Ctx.Redirect(302, "/")
}
