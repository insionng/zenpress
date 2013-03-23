package main

import (
	"./handlers"
	"./handlers/root"
	"./models"
	"github.com/insionng/torgo"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	models.CreateDb()
	torgo.TorApp.SetStaticPath("/static", "./static")
	torgo.TorApp.SetStaticPath("/archives", "./archives")

	torgo.Route("/", &handlers.MainHandler{})
	torgo.Route("/category/:cid([0-9]+)", &handlers.MainHandler{})
	torgo.Route("/search", &handlers.SearchHandler{})

	torgo.Route("/node/:nid([0-9]+)", &handlers.NodeHandler{})
	torgo.Route("/view/:tid([0-9]+)", &handlers.ViewHandler{})

	torgo.Route("/register", &handlers.RegHandler{})
	torgo.Route("/login", &handlers.LoginHandler{})
	torgo.Route("/logout", &handlers.LogoutHandler{})

	torgo.Route("/like/topic/:tid([0-9]+)", &handlers.LikeTopicHandler{})
	torgo.Route("/hate/topic/:tid([0-9]+)", &handlers.HateTopicHandler{})

	torgo.Route("/like/node/:nid([0-9]+)", &handlers.LikeNodeHandler{})
	torgo.Route("/hate/node/:nid([0-9]+)", &handlers.HateNodeHandler{})

	torgo.Route("/new/category", &handlers.NewCategoryHandler{})
	torgo.Route("/new/node", &handlers.NewNodeHandler{})
	torgo.Route("/new/topic", &handlers.NewTopicHandler{})
	torgo.Route("/new/reply/:tid([0-9]+)", &handlers.NewReplyHandler{})

	torgo.Route("/modify/category", &handlers.ModifyCategoryHandler{})
	torgo.Route("/modify/node", &handlers.ModifyNodeHandler{})

	torgo.Route("/topic/delete/:tid([0-9]+)", &handlers.TopicDeleteHandler{})
	torgo.Route("/topic/edit/:tid([0-9]+)", &handlers.TopicEditHandler{})

	torgo.Route("/node/delete/:nid([0-9]+)", &handlers.NodeDeleteHandler{})
	torgo.Route("/node/edit/:nid([0-9]+)", &handlers.NodeEditHandler{})

	torgo.Route("/delete/reply/:rid([0-9]+)", &handlers.DeleteReplyHandler{})

	//root routes
	torgo.Route("/root", &root.RMainHandler{})
	torgo.Route("/root/login", &root.RLoginHandler{})
	torgo.Route("/root/account", &root.RAccountHandler{})
	torgo.Route("/root/change_password", &root.RChangePasswordHandler{})
	torgo.Route("/root/category_list", &root.RCategoryListHandler{})

	torgo.SessionOn = true
	torgo.Run()
}
