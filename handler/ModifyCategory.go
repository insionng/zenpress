package handler

import (
	"fmt"
	"github.com/insionng/vodka"
	"github.com/insionng/zenpress/models"
	"net/http"
	"time"
)

func ModifyCatGetHandler(self *vodka.Context) error {
	return self.Render(http.StatusOK, "modify_category.html", nil)
}

func ModifyCatPostHandler(self *vodka.Context) error {

	cid, _ := self.ParamInt64("categoryid")

	cat_title := self.ParamEscape("title")
	cat_content := self.ParamEscape("content")
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
