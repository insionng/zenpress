package handlers

import (
	"../libs"
	"../models"
	"strconv"
	"time"
)

type TopicEditHandler struct {
	libs.RootAuthHandler
}

func (this *TopicEditHandler) Get() {
	tid, _ := strconv.Atoi(this.Ctx.Params[":tid"])
	this.Data["Topic"] = models.GetTopic(tid)
	this.Layout = "layout.html"
	this.TplNames = "edit.html"
	this.Render()
}

func (this *TopicEditHandler) Post() {
	tid, _ := strconv.Atoi(this.Ctx.Params[":tid"])

	inputs := this.Input()
	var tp models.Topic
	tp.Id = int64(tid)

	tp.Title = inputs.Get("title")
	tp.Content = inputs.Get("content")
	tp.Created = time.Now()
	models.SaveTopic(tp)
	this.Ctx.Redirect(302, "/")
}
