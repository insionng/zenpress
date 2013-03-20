package root

import (
	"../../libs"
	"../../models"
	//"../../utils"

)

type RMainHandler struct {
	libs.RootHandler
}

func (self *RMainHandler) Get() {
	tpcc_today, tpcc_week, tpcc_month := models.TopicCount()
	cs_categorys, cs_nodes, cs_topics, cs_menbers := models.Counts()

	self.Data["tpcc_today"] = tpcc_today
	self.Data["tpcc_week"] = tpcc_week
	self.Data["tpcc_month"] = tpcc_month
	self.Data["cs_categorys"] = cs_categorys
	self.Data["cs_nodes"] = cs_nodes
	self.Data["cs_topics"] = cs_topics
	self.Data["cs_menbers"] = cs_menbers

	self.TplNames = "root/index.html"
	self.Render()
}
