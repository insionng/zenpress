package handler

import (
	"fmt"
	"github.com/insionng/macross"
	"github.com/insionng/macross/jwt"
	"github.com/insionng/zenpress/models"
)

func TopicEditGetHandler(self *macross.Context) error {

	tid := self.Param("tid").MustInt64()
	tid_handler := models.GetTopic(tid)
	self.Set("topic", tid_handler)
	self.Set("inode", models.GetNode(tid_handler.Nid))

	return self.Render("topic_edit")
}

func TopicEditPostHandler(self *macross.Context) error {
	tid := self.Param("tid").MustInt64()
	nid := self.Param("nodeid").MustInt64()
	cid := models.GetNode(nid).Pid

	claims := jwt.GetMapClaims(self)
	var uid int64
	if jwtUserId, okay := claims["UserId"].(float64); okay {
		uid = int64(jwtUserId)
	}

	tid_title := self.Args("title").String()
	tid_content := self.Args("content").String()

	if tid_title != "" && tid_content != "" {
		models.EditTopic(tid, nid, cid, uid, tid_title, tid_content)
		return self.Redirect(fmt.Sprintf("/view/%v/", tid), 302)
	} else {
		return self.Redirect("/", 302)
	}
}
