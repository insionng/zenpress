package main

import (
	"./handlers"
	"./handlers/root"
	"./models"
	"github.com/astaxie/beego"
	//"github.com/insionng/torgo"
	//"./torgo"
)

func main() {
	models.CreateDb()
	beego.SetStaticPath("/static", "./static")
	beego.SetStaticPath("/archives", "./archives")

	beego.Router("/", &handlers.MainHandler{})
	beego.Router("/category/:cid([0-9]+)", &handlers.MainHandler{})
	beego.Router("/search", &handlers.SearchHandler{})

	beego.Router("/node/:nid([0-9]+)", &handlers.NodeHandler{})
	beego.Router("/view/:tid([0-9]+)", &handlers.ViewHandler{})

	beego.Router("/register", &handlers.RegHandler{})
	beego.Router("/login", &handlers.LoginHandler{})
	beego.Router("/logout", &handlers.LogoutHandler{})

	beego.Router("/like/topic/:tid([0-9]+)", &handlers.LikeTopicHandler{})
	beego.Router("/hate/topic/:tid([0-9]+)", &handlers.HateTopicHandler{})

	beego.Router("/like/node/:nid([0-9]+)", &handlers.LikeNodeHandler{})
	beego.Router("/hate/node/:nid([0-9]+)", &handlers.HateNodeHandler{})

	beego.Router("/new/category", &handlers.NewCategoryHandler{})
	beego.Router("/new/node", &handlers.NewNodeHandler{})
	beego.Router("/new/topic", &handlers.NewTopicHandler{})
	beego.Router("/new/reply/:tid([0-9]+)", &handlers.NewReplyHandler{})

	beego.Router("/modify/category", &handlers.ModifyCategoryHandler{})
	beego.Router("/modify/node", &handlers.ModifyNodeHandler{})

	beego.Router("/topic/delete/:tid([0-9]+)", &handlers.TopicDeleteHandler{})
	beego.Router("/topic/edit/:tid([0-9]+)", &handlers.TopicEditHandler{})

	beego.Router("/node/delete/:nid([0-9]+)", &handlers.NodeDeleteHandler{})
	beego.Router("/node/edit/:nid([0-9]+)", &handlers.NodeEditHandler{})

	beego.Router("/delete/reply/:rid([0-9]+)", &handlers.DeleteReplyHandler{})

	//root routes
	beego.Router("/root", &root.RMainHandler{})
	beego.Router("/root-login", &root.RLoginHandler{})
	beego.Router("/root/account", &root.RAccountHandler{})

	beego.SessionOn = true
	beego.Run()
}
