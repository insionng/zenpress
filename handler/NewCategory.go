package handler

import (
	"github.com/insionng/vodka"
	"github.com/insionng/zenpress/models"
	"net/http"
)

func NewCategoryGetHandler(self *vodka.Context) error {
	return self.Render(http.StatusOK, "new_category.html", nil)
}

func NewCategoryPostHandler(self *vodka.Context) error {

	t := self.FormEscape("title")
	c := self.FormEscape("content")
	if t != "" && c != "" {
		models.AddCategory(t, c)
	}
	return self.Redirect(302, "/")
}
