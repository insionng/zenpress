package handler

import (
	"github.com/insionng/macross"
	"github.com/insionng/zenpress/models"
)

func GetNewCategoryHandler(self *macross.Context) error {
	return self.Render("new_category")
}

func PostNewCategoryHandler(self *macross.Context) error {

	t := self.Args("title").String()
	c := self.FormValue("content")
	if t != "" && c != "" {
		models.AddCategory(t, c)
	}
	return self.Redirect("/", 302)
}
