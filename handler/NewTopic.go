package handler

import (
	"fmt"
	"net/http"

	"github.com/Unknwon/com"
	"github.com/insionng/vodka"
	"github.com/insionng/zenpress/models"
	"github.com/vodka-contrib/session"
)

func NewTopicGetHandler(self vodka.Context) error {
	self.Set("nodes", models.GetAllNode())
	return self.Render(http.StatusOK, "topic_new.html")
}

func NewTopicPostHandler(self vodka.Context) error {
	nid := com.StrTo(self.Param("nodeid")).MustInt64()
	cid := models.GetNode(nid).Pid
	sess := session.GetStore(self)
	var user models.User
	val := sess.Get("user")
	if val != nil {
		user = val.(models.User)
	}
	uid := user.Id
	author := user.Nickname

	tid_title := self.FormValue("title")
	tid_content := self.FormValue("content")

	if tid_title != "" && tid_content != "" {
		models.AddTopic(tid_title, tid_content, cid, nid, uid, author)
		return self.Redirect(302, fmt.Sprintf("/node/", nid))
	} else {
		return self.Redirect(302, "/")
	}
}
