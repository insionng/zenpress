package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Unknwon/com"
	"github.com/insionng/vodka"
	"github.com/insionng/zenpress/models"
)

func ModifyCatGetHandler(self vodka.Context) error {
	return self.Render(http.StatusOK, "modify_category.html")
}

func ModifyCatPostHandler(self vodka.Context) error {

	cid := com.StrTo(self.Param("categoryid")).MustInt64()

	cat_title := self.FormValue("title")
	cat_content := self.FormValue("content")
	if cid != 0 && cat_title != "" && cat_content != "" {
		var cat models.Category
		cat.Id = int64(cid)
		cat.Title = cat_title
		cat.Content = cat_content
		cat.Created = time.Now().Unix()
		models.UpdateCategory(cat.Id, cat)
		return self.Redirect(302, fmt.Sprintf("/category/%v/", cid))
	} else {
		return self.Redirect(302, "/")
	}
}
