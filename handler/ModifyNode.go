package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Unknwon/com"
	"github.com/insionng/vodka"
	"github.com/insionng/zenpress/models"
)

func ModifyNodeGetHandler(self vodka.Context) error {
	return self.Render(http.StatusOK, "modify_node.html", nil)
}

func ModifyNodePostHandler(self vodka.Context) error {

	cid := com.StrTo(self.Param("categoryid")).MustInt64()
	nid := com.StrTo(self.Param("nodeid")).MustInt64()

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
		return self.Redirect(302, fmt.Sprintf("/node/%v/", nid))
	} else {
		return self.Redirect(302, "/")
	}
}
