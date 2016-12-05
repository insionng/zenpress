package handler

import (
	"fmt"
	"github.com/insionng/macross"
	"github.com/insionng/macross/jwt"
	"github.com/insionng/zenpress/models"
)

func NewNodeGetHandler(self *macross.Context) error {
	data := make(map[string]interface{})
	data["categorys"] = models.GetAllCategory()
	self.SetStore(data)
	return self.Render("new_node")
}

func NewNodePostHandler(self *macross.Context) error {

	claims := jwt.GetMapClaims(self)
	var uid int64
	if jwtUserId, okay := claims["UserId"].(float64); okay {
		uid = int64(jwtUserId)
	}

	cid := self.Param("category").MustInt64()

	nid_title := self.Args("title").String()
	nid_content := self.Args("content").String()

	if nid_title != "" && nid_content != "" && cid != 0 {
		models.AddNode(nid_title, nid_content, cid, uid)
		return self.Redirect(fmt.Sprintf("/category/", cid), 302)
	} else {
		return self.Redirect("/", 302)
	}
}
