package handlers

import (
	"strconv"
	"time"
	"toropress/helper"
	"toropress/libs"
	"toropress/models"
)

type LikeHandler struct {
	libs.BaseHandler
}

func (self *LikeHandler) Get() {

	if helper.IsSpider(self.Ctx.Request.UserAgent()) != true {
		name := self.GetString(":name")
		id, _ := self.GetInt(":id")

		if name == "topic" {

			tp := models.GetTopic(id)
			tp.Hotup = tp.Hotup + 1
			tp.Hotscore = helper.Hotness_Score(tp.Hotup, tp.Hotdown)
			tp.Hotness = helper.Hotness(tp.Hotup, tp.Hotdown, time.Now())

			models.PutTopic(id, tp)
			//&hearts; 有用 ({{.article.Hotup}})
			self.Ctx.WriteString(strconv.Itoa(int(tp.Hotup)))
		} else if name == "node" {

			nd := models.GetNode(id)
			nd.Hotup = nd.Hotup + 1
			nd.Hotscore = helper.Hotness_Score(nd.Hotup, nd.Hotdown)
			nd.Hotness = helper.Hotness(nd.Hotup, nd.Hotdown, time.Now())

			models.PutNode(id, nd)

			self.Ctx.WriteString("node liked")
		} else {
			self.Abort("401")
		}

	} else {
		self.Abort("401")
	}

}
