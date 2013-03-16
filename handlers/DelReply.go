package handlers

import (
	"../libs"
	"../models"
	"strconv"
)

type DeleteReplyHandler struct {
	libs.RootAuthHandler
}

func (self *DeleteReplyHandler) Get() {
	rid, _ := strconv.Atoi(self.Ctx.Params[":rid"])
	models.DelReply(rid)
	self.Ctx.Redirect(302, "/")
}
