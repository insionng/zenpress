package handlers

import (
	"../libs"
	"../models"
)

type NewCategoryHandler struct {
	libs.RootAuthHandler
}

func (self *NewCategoryHandler) Get() {
	self.TplNames = "new_category.html"
	self.Layout = "layout.html"

	self.Render()
}

func (self *NewCategoryHandler) Post() {
	inputs := self.Input()
	models.AddCategory(inputs.Get("title"), inputs.Get("content"))

	self.Ctx.Redirect(302, "/")
}
