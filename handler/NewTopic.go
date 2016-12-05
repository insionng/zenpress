package handler

import (
	"fmt"
	"github.com/insionng/macross"
	"github.com/insionng/macross/jwt"
	"github.com/insionng/zenpress/models"
)

func NewTopicGetHandler(self *macross.Context) error {
	self.Set("nodes", models.GetAllNode())
	return self.Render("topic_new")
}

func NewTopicPostHandler(self *macross.Context) error {
	nid := self.Param("nodeid").MustInt64()
	cid := models.GetNode(nid).Pid

	claims := jwt.GetMapClaims(self)
	var uid int64
	if jwtUserId, okay := claims["UserId"].(float64); okay {
		uid = int64(jwtUserId)
	}
	var author string
	if jwtUsername, okay := claims["Username"].(string); okay {
		author = jwtUsername
	}

	tid_title := self.Args("title").String()
	tid_content := self.Args("content").String()

	if tid_title != "" && tid_content != "" {
		models.AddTopic(tid_title, tid_content, cid, nid, uid, author)
		return self.Redirect(fmt.Sprintf("/node/", nid), macross.StatusFound)
	} else {
		return self.Redirect("/", macross.StatusFound)
	}
}
