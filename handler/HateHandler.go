package handler

import (
	"github.com/insionng/vodka"
	"github.com/insionng/zenpress/helper"
	"github.com/insionng/zenpress/models"
	"net/http"
	"time"
)

func HateHandler(self *vodka.Context) error {

	if helper.IsSpider(self.Request().UserAgent()) != true {
		name := self.FormEscape(":name")
		id, _ := self.ParamInt64(":id")

		if name == "topic" {

			tp := models.GetTopic(id)
			tp.Hotdown = tp.Hotdown + 1
			tp.Hotscore = helper.Hotness_Score(tp.Hotup, tp.Hotdown)
			tp.Hotness = helper.Hotness(tp.Hotup, tp.Hotdown, time.Now())

			models.PutTopic(id, tp)
			return self.String(http.StatusOK, "%v", tp.Hotdown)

		} else if name == "node" {

			nd := models.GetNode(id)
			nd.Hotdown = nd.Hotdown + 1
			nd.Hotscore = helper.Hotness_Score(nd.Hotup, nd.Hotdown)
			nd.Hotness = helper.Hotness(nd.Hotup, nd.Hotdown, time.Now())

			models.PutNode(id, nd)

			return self.Status(200)
		}

	}
	return self.Status(401)
}
