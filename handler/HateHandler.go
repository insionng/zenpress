package handler

import (
	"fmt"
	"time"

	"github.com/insionng/macross"
	"github.com/insionng/zenpress/helper"
	"github.com/insionng/zenpress/models"
)

func HateHandler(self *macross.Context) error {

	if helper.IsSpider(string(self.Request.Header.UserAgent())) != true {
		name := self.Args("name").String()
		id := self.Param("id").MustInt64()

		if name == "topic" {

			tp := models.GetTopic(id)
			tp.Hotdown = tp.Hotdown + 1
			tp.Hotscore = helper.Hotness_Score(tp.Hotup, tp.Hotdown)
			tp.Hotness = helper.Hotness(tp.Hotup, tp.Hotdown, time.Now())

			models.PutTopic(id, tp)
			return self.String(fmt.Sprintf("%v", tp.Hotdown))

		} else if name == "node" {

			nd := models.GetNode(id)
			nd.Hotdown = nd.Hotdown + 1
			nd.Hotscore = helper.Hotness_Score(nd.Hotup, nd.Hotdown)
			nd.Hotness = helper.Hotness(nd.Hotup, nd.Hotdown, time.Now())

			models.PutNode(id, nd)
			return self.String("StatusOK")
		}

	}
	return self.String("StatusUnauthorized", macross.StatusUnauthorized)
}
