package handlers

import (
	"../libs"
	"../models"
	"../utils"
	"strconv"
)

type HateNodeHandler struct {
	libs.BaseHandler
}

func (this *HateNodeHandler) Get() {
	//inputs := this.Input()
	//id, _ := strconv.Atoi(inputs.Get("id"))
	if utils.IsSpider(this.Ctx.Request.UserAgent()) != true {

		id, _ := strconv.Atoi(this.Ctx.Params[":nid"])

		nd := models.GetNode(id)
		nd.Hotdown = nd.Hotdown + 1
		nd.Hotness = utils.Hotness(nd.Hotup, nd.Hotdown, nd.Created)

		models.SaveNode(nd)

		this.Ctx.WriteString("success")
		//this.Ctx.Redirect(302, "/")

	} else {
		this.Ctx.WriteString("R u spider?")
	}

}
