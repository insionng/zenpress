package handler

import (
	"fmt"
	"github.com/insionng/macross"
	"github.com/insionng/macross/jwt"
	"github.com/insionng/zenpress/models"
)

func NodeEditGetHandler(self *macross.Context) error {
	data := make(map[string]interface{})
	nid := self.Param("nid").MustInt64()
	nid_handler := models.GetNode(nid)
	data["inode"] = nid_handler
	data["icategory"] = models.GetCategory(nid_handler.Pid)
	self.SetStore(data)
	return self.Render("node_edit")
}

func NodeEditPostHandler(self *macross.Context) error {
	nid := self.Param("nid").MustInt64()
	cid := self.Param("categoryid").MustInt64()

	claims := jwt.GetMapClaims(self)
	var uid int64
	if jwtUserId, okay := claims["UserId"].(float64); okay {
		uid = int64(jwtUserId)
	}

	nid_title := self.Args("title").String()
	nid_content := self.Args("content").String()
	if nid_title != "" && nid_content != "" {
		models.EditNode(nid, cid, uid, nid_title, nid_content)
		return self.Redirect(fmt.Sprintf("/node/%v/", nid), 302)
	} else {
		return self.Redirect("/", 302)
	}
}
