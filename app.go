package main

import (
	"./handlers"
	"./handlers/root"
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

	torgo.RegisterController("/:cid([0-9]+)/:nid([0-9]+)", &handlers.NodeHandler{})
	torgo.RegisterController("/:cid([0-9]+)/:nid([0-9]+)/:tid([0-9]+)", &handlers.ViewHandler{})

	torgo.RegisterController("/register", &handlers.RegHandler{})
	torgo.RegisterController("/login", &handlers.LoginHandler{})
	torgo.RegisterController("/logout", &handlers.LogoutHandler{})

	torgo.RegisterController("/like/topic/:tid([0-9]+)", &handlers.LikeTopicHandler{})
	torgo.RegisterController("/hate/topic/:tid([0-9]+)", &handlers.HateTopicHandler{})

	torgo.RegisterController("/like/node/:nid([0-9]+)", &handlers.LikeNodeHandler{})
	torgo.RegisterController("/hate/node/:nid([0-9]+)", &handlers.HateNodeHandler{})

	torgo.RegisterController("/category/new", &handlers.NewCategoryHandler{})
	torgo.RegisterController("/node/new", &handlers.NewNodeHandler{})
	torgo.RegisterController("/topic/new", &handlers.NewTopicHandler{})

	torgo.RegisterController("/topic/delete/:id([0-9]+)", &handlers.DeleteHandler{})
	torgo.RegisterController("/topic/edit/:id([0-9]+)", &handlers.EditHandler{})

	torgo.RegisterController("/root", &root.RootHandler{})
	torgo.RegisterController("/root/login", &root.LoginHandler{})

	torgo.SessionOn = true
	torgo.Run()
}
