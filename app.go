package main

import (
	"./handlers"
	//"github.com/insionng/torgo"
	"./libs/torgo"
	//"lihuashu.com/insionng/torgo"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	torgo.TorApp.SetStaticPath("/static", "./static")
	torgo.TorApp.SetStaticPath("/archives", "./archives")

	torgo.RegisterHandler("/", &handlers.MainHandler{})
	torgo.RegisterHandler("/category/:cid([0-9]+)", &handlers.MainHandler{})
	torgo.RegisterHandler("/search", &handlers.SearchHandler{})

	torgo.RegisterHandler("/node/:nid([0-9]+)", &handlers.NodeHandler{})
	torgo.RegisterHandler("/view/:tid([0-9]+)", &handlers.ViewHandler{})

	torgo.RegisterHandler("/register", &handlers.RegHandler{})
	torgo.RegisterHandler("/login", &handlers.LoginHandler{})
	torgo.RegisterHandler("/logout", &handlers.LogoutHandler{})

	torgo.RegisterHandler("/like/topic/:tid([0-9]+)", &handlers.LikeTopicHandler{})
	torgo.RegisterHandler("/hate/topic/:tid([0-9]+)", &handlers.HateTopicHandler{})

	torgo.RegisterHandler("/like/node/:nid([0-9]+)", &handlers.LikeNodeHandler{})
	torgo.RegisterHandler("/hate/node/:nid([0-9]+)", &handlers.HateNodeHandler{})

	torgo.RegisterHandler("/new/category", &handlers.NewCategoryHandler{})
	torgo.RegisterHandler("/new/node", &handlers.NewNodeHandler{})
	torgo.RegisterHandler("/new/topic", &handlers.NewTopicHandler{})

	torgo.RegisterHandler("/modify/category", &handlers.ModifyCategoryHandler{})
	torgo.RegisterHandler("/modify/node", &handlers.ModifyNodeHandler{})

	torgo.RegisterHandler("/topic/delete/:tid([0-9]+)", &handlers.TopicDeleteHandler{})
	torgo.RegisterHandler("/topic/edit/:tid([0-9]+)", &handlers.TopicEditHandler{})

	torgo.RegisterHandler("/node/delete/:nid([0-9]+)", &handlers.NodeDeleteHandler{})
	torgo.RegisterHandler("/node/edit/:nid([0-9]+)", &handlers.NodeEditHandler{})

	torgo.RegisterHandler("/delete/reply/:rid([0-9]+)", &handlers.DeleteReplyHandler{})

	torgo.SessionOn = true
	torgo.Run()
}
