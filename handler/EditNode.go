package handler

import (
	"fmt"
	"github.com/insionng/vodka"
	"github.com/insionng/zenpress/models"
	"github.com/vodka-contrib/sessions"
	"net/http"
)

func NodeEditGetHandler(self *vodka.Context) error {
	data := make(map[string]interface{})
	nid, _ := self.ParamInt64(":nid")
	nid_handler := models.GetNode(nid)
	data["inode"] = nid_handler
	data["icategory"] = models.GetCategory(nid_handler.Pid)

	return self.Render(http.StatusOK, "node_edit.html", data)
}

func NodeEditPostHandler(self *vodka.Context) error {
	nid, _ := self.ParamInt64(":nid")
	cid, _ := self.ParamInt64("categoryid")

	session := sessions.Default(self)
	var user models.User
	val := session.Get("user")
	if val != nil {
		user = val.(models.User)
	}
	uid := user.Id

	nid_title := self.FormEscape("title")
	nid_content := self.FormEscape("content")
	if nid_title != "" && nid_content != "" {
		models.EditNode(nid, cid, uid, nid_title, nid_content)
		return self.Redirect(302, fmt.Sprintf("/node/%v/", nid))
	} else {
		return self.Redirect(302, "/")
	}
}
