package handlers

import (
	"strconv"
	"time"
	"toropress/helper"
	"toropress/libs"
	"toropress/models"
)

type HateHandler struct {
	libs.BaseHandler
}

func (self *HateHandler) Get() {

	if helper.IsSpider(self.Ctx.Request.UserAgent()) != true {
		name := self.GetString(":name")
		id, _ := self.GetInt(":id")

		if name == "topic" {

			tp := models.GetTopic(id)
			tp.Hotdown = tp.Hotdown + 1
			tp.Hotscore = helper.Hotness_Score(tp.Hotup, tp.Hotdown)
			tp.Hotness = helper.Hotness(tp.Hotup, tp.Hotdown, time.Now())

			models.PutTopic(id, tp)
			//&spades; 没用 ({{.article.Hotdown}})
			self.Ctx.WriteString(strconv.Itoa(int(tp.Hotdown)))

		} else if name == "node" {

			nd := models.GetNode(id)
			nd.Hotdown = nd.Hotdown + 1
			nd.Hotscore = helper.Hotness_Score(nd.Hotup, nd.Hotdown)
			nd.Hotness = helper.Hotness(nd.Hotup, nd.Hotdown, time.Now())

			models.PutNode(id, nd)

			self.Ctx.WriteString("node hated")
		} else {
			self.Abort("401")
		}

	} else {
		self.Abort("401")
	}

}
