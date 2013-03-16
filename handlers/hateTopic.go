package handlers

import (
	"../libs"
	"../models"
	"../utils"
	"strconv"
)

type HateTopicHandler struct {
	libs.BaseHandler
}

func (self *HateTopicHandler) Get() {
	//inputs := self.Input()
	//id, _ := strconv.Atoi(inputs.Get("id"))
	if utils.IsSpider(self.Ctx.Request.UserAgent()) != true {

		id, _ := strconv.Atoi(self.Ctx.Params[":tid"])

		tp := models.GetTopic(id)
		tp.Hotdown = tp.Hotdown + 1
		tp.Hotness = utils.Hotness(tp.Hotup, tp.Hotdown, tp.Created)

		models.SaveTopic(tp)

		self.Ctx.WriteString("success")

	} else {
		self.Ctx.WriteString("R u spider?")
	}

}
