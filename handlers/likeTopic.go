package handlers

import (
	"../libs"
	"../models"
	"../utils"
	"strconv"
)

type LikeTopicHandler struct {
	libs.BaseHandler
}

func (this *LikeTopicHandler) Get() {
	//inputs := this.Input()
	//id, _ := strconv.Atoi(inputs.Get("id"))
	if utils.IsSpider(this.Ctx.Request.UserAgent()) != true {

		id, _ := strconv.Atoi(this.Ctx.Params[":tid"])

		tp := models.GetTopic(id)
		tp.Hotup = tp.Hotup + 1
		tp.Hotness = utils.Hotness(tp.Hotup, tp.Hotdown, tp.Created)

		models.SaveTopic(tp)

		this.Ctx.WriteString("success")

	} else {
		this.Ctx.WriteString("R u spider?")
	}

}
