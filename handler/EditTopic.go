package handler

import (
	"fmt"
	"github.com/insionng/vodka"
	"github.com/insionng/zenpress/models"
	"github.com/vodka-contrib/sessions"
	"net/http"
)

func TopicEditGetHandler(self *vodka.Context) error {
	data := make(map[string]interface{})
	tid, _ := self.ParamInt64(":tid")
	tid_handler := models.GetTopic(tid)
	data["topic"] = tid_handler
	data["inode"] = models.GetNode(tid_handler.Nid)

	return self.Render(http.StatusOK, "topic_edit.html", data)
}

func TopicEditPostHandler(self *vodka.Context) error {
	tid, _ := self.ParamInt64(":tid")
	nid, _ := self.ParamInt64("nodeid")
	cid := models.GetNode(nid).Pid

	session := sessions.Default(self)
	var user models.User
	val := session.Get("user")
	if val != nil {
		user = val.(models.User)
	}
	uid := user.Id
	tid_title := self.FormEscape("title")
	tid_content := self.FormEscape("content")

	if tid_title != "" && tid_content != "" {
		models.EditTopic(tid, nid, cid, uid, tid_title, tid_content)
		return self.Redirect(302, fmt.Sprintf("/view/%v/", tid))
	} else {
		return self.Redirect(302, "/")
	}
}
