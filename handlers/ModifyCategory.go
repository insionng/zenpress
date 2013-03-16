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

func (self *ModifyCategoryHandler) Get() {
	self.TplNames = "modify_category.html"
	self.Layout = "layout.html"

	self.Render()
}

func (self *ModifyCategoryHandler) Post() {
	inputs := self.Input()
	cid, _ := strconv.Atoi(inputs.Get("categoryid"))

	var cat models.Category
	cat.Id = int64(cid)
	cat.Title = inputs.Get("title")
	cat.Content = inputs.Get("content")
	cat.Created = time.Now()
	models.SaveCategory(cat)
	self.Ctx.Redirect(302, "/")
}
