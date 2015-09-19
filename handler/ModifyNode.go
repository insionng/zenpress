package handler

import (
	"fmt"
	"github.com/insionng/vodka"
	"github.com/insionng/zenpress/models"
	"net/http"
	"time"
)

func ModifyNodeGetHandler(self *vodka.Context) error {
	return self.Render(http.StatusOK, "modify_node.html", nil)
}

func ModifyNodePostHandler(self *vodka.Context) error {

	cid, _ := self.ParamInt64("categoryid")
	nid, _ := self.ParamInt64("nodeid")

	nd_title := self.FormEscape("title")
	nd_content := self.FormEscape("content")
	if cid != 0 && nid != 0 && nd_title != "" && nd_content != "" {
		nd := new(models.Node)
		nd.Id = int64(nid)
		nd.Pid = int64(cid)
		nd.Title = nd_title
		nd.Content = nd_content
		nd.Created = time.Now()
		models.UpdateNode(nd.Id, nd)
		return self.Redirect(302, fmt.Sprintf("/node/%v/", nid))
	} else {
		return self.Redirect(302, "/")
	}
}
