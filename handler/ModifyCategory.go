package handler

import (
	"fmt"
	"time"

	"github.com/insionng/macross"
	"github.com/insionng/zenpress/models"
)

func ModifyCatGetHandler(self *macross.Context) error {
	return self.Render("modify_category")
}

func ModifyCatPostHandler(self *macross.Context) error {

	cid := self.Param("categoryid").MustInt64()

	cat_title := self.Args("title").String()
	cat_content := self.Args("content").String()
	if cid != 0 && cat_title != "" && cat_content != "" {
		var cat models.Category
		cat.Id = int64(cid)
		cat.Title = cat_title
		cat.Content = cat_content
		cat.Created = time.Now().Unix()
		models.UpdateCategory(cat.Id, cat)
		return self.Redirect(fmt.Sprintf("/category/%v/", cid), 302)
	} else {
		return self.Redirect("/", 302)
	}
}
