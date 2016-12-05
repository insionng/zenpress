package handler

import (
	"fmt"
	"time"

	"github.com/insionng/macross"
	"github.com/insionng/zenpress/helper"
	"github.com/insionng/zenpress/models"
)

func LikeHandler(self *macross.Context) error {

	if helper.IsSpider(string(self.Request.Header.UserAgent())) != true {
		name := self.Args("name").String()
		id := self.Param("id").MustInt64()

		if name == "topic" {

			tp := models.GetTopic(id)
			tp.Hotup = tp.Hotup + 1
			tp.Hotscore = helper.Hotness_Score(tp.Hotup, tp.Hotdown)
			tp.Hotness = helper.Hotness(tp.Hotup, tp.Hotdown, time.Now())

			models.PutTopic(id, tp)
			return self.String(fmt.Sprintf("%v", tp.Hotup))
		} else if name == "node" {

			nd := models.GetNode(id)
			nd.Hotup = nd.Hotup + 1
			nd.Hotscore = helper.Hotness_Score(nd.Hotup, nd.Hotdown)
			nd.Hotness = helper.Hotness(nd.Hotup, nd.Hotdown, time.Now())

			models.PutNode(id, nd)

			return self.String("StatusOK")
		}

	}
	return self.String("StatusUnauthorized", macross.StatusUnauthorized)
}
