package handler

import (
	"fmt"
	"time"

	"github.com/insionng/macross"
	"github.com/insionng/zenpress/models"
)

func ModifyNodeGetHandler(self *macross.Context) error {
	return self.Render("modify_node")
}

func ModifyNodePostHandler(self *macross.Context) error {

	cid := self.Param("categoryid").MustInt64()
	nid := self.Param("nodeid").MustInt64()

	nd_title := self.FormValue("title")
	nd_content := self.FormValue("content")
	if cid != 0 && nid != 0 && nd_title != "" && nd_content != "" {
		nd := new(models.Node)
		nd.Id = int64(nid)
		nd.Pid = int64(cid)
		nd.Title = nd_title
		nd.Content = nd_content
		nd.Created = time.Now().Unix()
		models.UpdateNode(nd.Id, nd)
		return self.Redirect(fmt.Sprintf("/node/%v/", nid), macross.StatusFound)
	} else {
		return self.Redirect("/", macross.StatusFound)
	}
}
