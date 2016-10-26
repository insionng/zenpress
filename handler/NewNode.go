package handler

import (
	"fmt"
	"net/http"

	"github.com/Unknwon/com"
	"github.com/insionng/vodka"
	"github.com/insionng/zenpress/models"
	"github.com/vodka-contrib/session"
)

func NewNodeGetHandler(self vodka.Context) error {
	data := make(map[string]interface{})
	data["categorys"] = models.GetAllCategory()
	self.SetStore(data)
	return self.Render(http.StatusOK, "new_node.html")
}

func NewNodePostHandler(self vodka.Context) error {

	sess := session.GetStore(self)
	var user models.User
	val := sess.Get("user")
	if val != nil {
		user = val.(models.User)
	}
	uid := user.Id
	cid := com.StrTo(self.Param("category")).MustInt64()

	nid_title := self.FormValue("title")
	nid_content := self.FormValue("content")

	if nid_title != "" && nid_content != "" && cid != 0 {
		models.AddNode(nid_title, nid_content, cid, uid)
		return self.Redirect(302, fmt.Sprintf("/category/", cid))
	} else {
		return self.Redirect(302, "/")
	}
}
