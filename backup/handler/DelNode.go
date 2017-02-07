package handler

import (
	"github.com/insionng/macross"
	"github.com/insionng/zenpress/models"
)

func NodeDeleteHandler(self *macross.Context) error {
	nid := self.Param("nid").MustInt64()
	models.DelNode(nid)
	return self.Redirect("/", 302)
}
