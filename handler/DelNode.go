package handler

import (
	"github.com/insionng/vodka"
	"github.com/insionng/zenpress/models"
)

func NodeDeleteHandler(self *vodka.Context) error {
	nid, _ := self.ParamInt64(":nid")
	models.DelNode(nid)
	return self.Redirect(302, "/")
}
