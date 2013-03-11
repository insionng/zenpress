package handlers

import (
	"../models"
	"../utils"
)

type NewCategoryHandler struct {
	utils.RootAuthHandler
}

func (this *NewCategoryHandler) Get() {
	this.TplNames = "new_category.html"
	this.Layout = "layout.html"

	this.Render()
}

func (this *NewCategoryHandler) Post() {
	inputs := this.Input()
	models.AddCategory(inputs.Get("title"), inputs.Get("content"))

	this.Ctx.Redirect(302, "/")
}
