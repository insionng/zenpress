package handlers

import (
	"../models"
	"../utils"
	"strconv"
	"time"
)

type EditHandler struct {
	utils.RootAuthHandler
}

func (this *EditHandler) Get() {
	id, _ := strconv.Atoi(this.Ctx.Params[":id"])
	this.Data["article"] = models.GetTopic(id)
	this.Layout = "layout.html"
	this.TplNames = "new.html"
	this.Render()
}

func (this *EditHandler) Post() {
	id, _ := strconv.Atoi(this.Ctx.Params[":id"])

	inputs := this.Input()
	var tp models.Topic
	tp.Id = int64(id)

	tp.Title = inputs.Get("title")
	tp.Content = inputs.Get("content")
	tp.Created = time.Now()
	models.SaveTopic(tp)
	this.Ctx.Redirect(302, "/")
}
