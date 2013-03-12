package main

import (
	"./handlers"
	"github.com/insionng/torgo"
	//"lihuashu.com/insionng/torgo"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	torgo.TorApp.SetStaticPath("/static", "./static")
	torgo.TorApp.SetStaticPath("/archives", "./archives")

	torgo.RegisterController("/", &handlers.MainHandler{})
	torgo.RegisterController("/category/:cid([0-9]+)", &handlers.MainHandler{})

	torgo.RegisterController("/node/:nid([0-9]+)", &handlers.NodeHandler{})
	torgo.RegisterController("/view/:tid([0-9]+)", &handlers.ViewHandler{})

	torgo.RegisterController("/register", &handlers.RegHandler{})
	torgo.RegisterController("/login", &handlers.LoginHandler{})
	torgo.RegisterController("/logout", &handlers.LogoutHandler{})

	torgo.RegisterController("/like/topic/:tid([0-9]+)", &handlers.LikeTopicHandler{})
	torgo.RegisterController("/hate/topic/:tid([0-9]+)", &handlers.HateTopicHandler{})

	torgo.RegisterController("/like/node/:nid([0-9]+)", &handlers.LikeNodeHandler{})
	torgo.RegisterController("/hate/node/:nid([0-9]+)", &handlers.HateNodeHandler{})

	torgo.RegisterController("/new/category", &handlers.NewCategoryHandler{})
	torgo.RegisterController("/new/node", &handlers.NewNodeHandler{})
	torgo.RegisterController("/new/topic", &handlers.NewTopicHandler{})

	torgo.RegisterController("/modify/category", &handlers.ModifyCategoryHandler{})
	torgo.RegisterController("/modify/node", &handlers.ModifyNodeHandler{})
	torgo.RegisterController("/modify/topic", &handlers.ModifyTopicHandler{})

	torgo.RegisterController("/topic/delete/:tid([0-9]+)", &handlers.TopicDeleteHandler{})
	torgo.RegisterController("/topic/edit/:tid([0-9]+)", &handlers.TopicEditHandler{})

	torgo.RegisterController("/node/delete/:nid([0-9]+)", &handlers.NodeDeleteHandler{})
	torgo.RegisterController("/node/edit/:nid([0-9]+)", &handlers.NodeEditHandler{})

	torgo.RegisterController("/delete/reply/:rid([0-9]+)", &handlers.DeleteReplyHandler{})

	torgo.SessionOn = true
	torgo.Run()
}
