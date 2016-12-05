package handler

import (
	"github.com/insionng/macross"
	"github.com/insionng/zenpress/models"
)

func NewCategoryGetHandler(self *macross.Context) error {
	return self.Render("new_category")
}

func NewCategoryPostHandler(self *macross.Context) error {

	t := self.Args("title").String()
	c := self.FormValue("content")
	if t != "" && c != "" {
		models.AddCategory(t, c)
	}
	return self.Redirect("/", 302)
}
