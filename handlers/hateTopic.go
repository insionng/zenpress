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

func (this *HateTopicHandler) Get() {
	//inputs := this.Input()
	//id, _ := strconv.Atoi(inputs.Get("id"))
	if utils.IsSpider(this.Ctx.Request.UserAgent()) != true {

		id, _ := strconv.Atoi(this.Ctx.Params[":tid"])

		tp := models.GetTopic(id)
		tp.Hotdown = tp.Hotdown + 1
		tp.Hotness = utils.Hotness(tp.Hotup, tp.Hotdown, tp.Created)

		models.SaveTopic(tp)

		this.Ctx.WriteString("success")

	} else {
		this.Ctx.WriteString("R u spider?")
	}

}
