package handler

import (
	"fmt"
	"github.com/insionng/vodka"
	"github.com/insionng/zenpress/models"
	"github.com/vodka-contrib/sessions"
	"net/http"
)

func NewTopicGetHandler(self *vodka.Context) error {
	data := make(map[string]interface{})
	data["nodes"] = models.GetAllNode()
	return self.Render(http.StatusOK, "topic_new.html", data)
}

func NewTopicPostHandler(self *vodka.Context) error {
	nid, _ := self.ParamInt64("nodeid")
	cid := models.GetNode(nid).Pid
	session := sessions.Default(self)
	var user models.User
	val := session.Get("user")
	if val != nil {
		user = val.(models.User)
	}
	uid := user.Id
	author := user.Nickname

	tid_title := self.FormEscape("title")
	tid_content := self.FormEscape("content")

	if tid_title != "" && tid_content != "" {
		models.AddTopic(tid_title, tid_content, cid, nid, uid, author)
		return self.Redirect(302, fmt.Sprintf("/node/", nid))
	} else {
		return self.Redirect(302, "/")
	}
}
