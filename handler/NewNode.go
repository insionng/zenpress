package handler

import (
	"fmt"
	"github.com/insionng/vodka"
	"github.com/insionng/zenpress/models"
	"github.com/vodka-contrib/sessions"
	"net/http"
)

func NewNodeGetHandler(self *vodka.Context) error {
	data := make(map[string]interface{})
	data["categorys"] = models.GetAllCategory()
	return self.Render(http.StatusOK, "new_node.html", data)
}

func NewNodePostHandler(self *vodka.Context) error {

	session := sessions.Default(self)
	var user models.User
	val := session.Get("user")
	if val != nil {
		user = val.(models.User)
	}
	uid := user.Id
	cid, _ := self.ParamInt64("category")

	nid_title := self.FormEscape("title")
	nid_content := self.FormEscape("content")

	if nid_title != "" && nid_content != "" && cid != 0 {
		models.AddNode(nid_title, nid_content, cid, uid)
		return self.Redirect(302, fmt.Sprintf("/category/", cid))
	} else {
		return self.Redirect(302, "/")
	}
}
