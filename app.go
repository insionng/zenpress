package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"toropress/handlers"
	"toropress/handlers/root"
	"toropress/models"
)

func main() {
	fmt.Println("Starting...")
	if !models.CreateDb() {
		fmt.Println("Database struct sync failed")
		return
	}
	fmt.Println("Database struct sync successfully")
	beego.SetStaticPath("/static", "static/")
	beego.SetStaticPath("/archives", "archives/")

	beego.Router("/", &handlers.MainHandler{})
	beego.Router("/category/:cid:int", &handlers.MainHandler{})
	beego.Router("/search", &handlers.SearchHandler{})

	beego.Router("/node/:nid:int", &handlers.NodeHandler{})
	beego.Router("/view/:tid:int", &handlers.ViewHandler{})

	beego.Router("/register", &handlers.RegHandler{})
	beego.Router("/login", &handlers.LoginHandler{})
	beego.Router("/logout", &handlers.LogoutHandler{})

	//hotness
	beego.Router("/like/:name:string/:id:int", &handlers.LikeHandler{})
	beego.Router("/hate/:name:string/:id:int", &handlers.HateHandler{})

	beego.Router("/new/category", &handlers.NewCategoryHandler{})
	beego.Router("/new/node", &handlers.NewNodeHandler{})
	beego.Router("/new/topic", &handlers.NewTopicHandler{})
	beego.Router("/new/reply/:tid:int", &handlers.NewReplyHandler{})

	beego.Router("/modify/category", &handlers.ModifyCategoryHandler{})
	beego.Router("/modify/node", &handlers.ModifyNodeHandler{})

	beego.Router("/topic/delete/:tid:int", &handlers.TopicDeleteHandler{})
	beego.Router("/topic/edit/:tid:int", &handlers.TopicEditHandler{})

	beego.Router("/node/delete/:nid:int", &handlers.NodeDeleteHandler{})
	beego.Router("/node/edit/:nid:int", &handlers.NodeEditHandler{})

	beego.Router("/delete/reply/:rid:int", &handlers.DeleteReplyHandler{})

	//root routes
	beego.Router("/root", &root.RMainHandler{})
	beego.Router("/root-login", &root.RLoginHandler{})
	beego.Router("/root/account", &root.RAccountHandler{})

	beego.SessionOn = true
	beego.Run()
}
