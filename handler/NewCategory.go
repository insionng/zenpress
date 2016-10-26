package handler

import (
	"net/http"

	"github.com/insionng/vodka"
	"github.com/insionng/zenpress/models"
)

func NewCategoryGetHandler(self vodka.Context) error {
	return self.Render(http.StatusOK, "new_category.html")
}

func NewCategoryPostHandler(self vodka.Context) error {

	t := self.FormValue("title")
	c := self.FormValue("content")
	if t != "" && c != "" {
		models.AddCategory(t, c)
	}
	return self.Redirect(302, "/")
}
