package handler

import (
	"fmt"
	"net/http"

	"github.com/Unknwon/com"
	"github.com/insionng/vodka"
	"github.com/insionng/zenpress/models"
	"github.com/vodka-contrib/session"
)

func TopicEditGetHandler(self vodka.Context) error {
	data := make(map[string]interface{})
	tid := com.StrTo(self.Param("tid")).MustInt64()
	tid_handler := models.GetTopic(tid)
	data["topic"] = tid_handler
	data["inode"] = models.GetNode(tid_handler.Nid)

	return self.Render(http.StatusOK, "topic_edit.html", data)
}

func TopicEditPostHandler(self vodka.Context) error {
	tid := com.StrTo(self.Param("tid")).MustInt64()
	nid := com.StrTo(self.Param("nodeid")).MustInt64()
	cid := models.GetNode(nid).Pid

	sess := session.GetStore(self)
	var user models.User
	val := sess.Get("user")
	if val != nil {
		user = val.(models.User)
	}
	uid := user.Id
	tid_title := self.FormValue("title")
	tid_content := self.FormValue("content")

	if tid_title != "" && tid_content != "" {
		models.EditTopic(tid, nid, cid, uid, tid_title, tid_content)
		return self.Redirect(302, fmt.Sprintf("/view/%v/", tid))
	} else {
		return self.Redirect(302, "/")
	}
}
