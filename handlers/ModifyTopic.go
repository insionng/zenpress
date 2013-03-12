package handlers

import (
	"../libs"
	"../models"
	"strconv"
	"time"
)

type ModifyTopicHandler struct {
	libs.RootAuthHandler
}

func (this *ModifyTopicHandler) Get() {
	tid, _ := strconv.Atoi(this.Ctx.Params[":tid"])
	tid_handler := models.GetTopic(tid)
	this.Data["topic"] = tid_handler
	this.Data["inode"] = models.GetNode(int(tid_handler.Nid))

	this.Layout = "layout.html"
	this.TplNames = "topic_edit.html"
	this.Render()
}

func (this *ModifyTopicHandler) Post() {
	inputs := this.Input()
	tid, _ := strconv.Atoi(this.Ctx.Params[":tid"])
	nid, _ := strconv.Atoi(inputs.Get("nodeid"))
	cid := models.GetNode(nid).Pid

	var tp models.Topic
	tp.Id = int64(tid)
	tp.Cid = cid
	tp.Nid = int64(nid)
	tp.Title = inputs.Get("title")
	tp.Content = inputs.Get("content")
	tp.Created = time.Now()
	models.SaveTopic(tp)
	this.Ctx.Redirect(302, "/")
}
