package handlers

import (
	"../models"
	"../utils"
	"strconv"
)

type DeleteHandler struct {
	utils.RootAuthHandler
}

func (this *DeleteHandler) Get() {
	tpid, _ := strconv.Atoi(this.Ctx.Params[":id"])
	models.DelTopic(tpid)
	this.Ctx.Redirect(302, "/")
}
