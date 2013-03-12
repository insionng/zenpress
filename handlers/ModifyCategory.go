package handlers

import (
	"../libs"
	"../models"
	"strconv"
	"time"
)

type ModifyCategoryHandler struct {
	libs.RootAuthHandler
}

func (this *ModifyCategoryHandler) Get() {
	this.TplNames = "modify_category.html"
	this.Layout = "layout.html"

	this.Render()
}

func (this *ModifyCategoryHandler) Post() {
	inputs := this.Input()
	cid, _ := strconv.Atoi(inputs.Get("categoryid"))

	var cat models.Category
	cat.Id = int64(cid)
	cat.Title = inputs.Get("title")
	cat.Content = inputs.Get("content")
	cat.Created = time.Now()
	models.SaveCategory(cat)
	this.Ctx.Redirect(302, "/")
}
