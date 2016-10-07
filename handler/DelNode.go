package handler

import (
	"github.com/Unknwon/com"
	"github.com/insionng/vodka"
	"github.com/insionng/zenpress/models"
)

func NodeDeleteHandler(self vodka.Context) error {
	nid := com.StrTo(self.Param("nid")).MustInt64()
	models.DelNode(nid)
	return self.Redirect(302, "/")
}
