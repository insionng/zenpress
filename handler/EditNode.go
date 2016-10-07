package handler

import (
	"fmt"
	"net/http"

	"github.com/Unknwon/com"
	"github.com/insionng/vodka"
	"github.com/insionng/zenpress/models"
	"github.com/vodka-contrib/session"
)

func NodeEditGetHandler(self vodka.Context) error {
	data := make(map[string]interface{})
	nid := com.StrTo(self.Param("nid")).MustInt64()
	nid_handler := models.GetNode(nid)
	data["inode"] = nid_handler
	data["icategory"] = models.GetCategory(nid_handler.Pid)

	return self.Render(http.StatusOK, "node_edit.html", data)
}

func NodeEditPostHandler(self vodka.Context) error {
	nid := com.StrTo(self.Param("nid")).MustInt64()
	cid := com.StrTo(self.Param("categoryid")).MustInt64()

	sess := session.GetStore(self)
	var user models.User
	val := sess.Get("user")
	if val != nil {
		user = val.(models.User)
	}
	uid := user.Id

	nid_title := self.FormValue("title")
	nid_content := self.FormValue("content")
	if nid_title != "" && nid_content != "" {
		models.EditNode(nid, cid, uid, nid_title, nid_content)
		return self.Redirect(302, fmt.Sprintf("/node/%v/", nid))
	} else {
		return self.Redirect(302, "/")
	}
}
