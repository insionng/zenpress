package main

import (

	//"github.com/insionng/zenpress/handlers"
	//"github.com/insionng/zenpress/handlers/root"
	//"github.com/insionng/zenpress/models"

	"github.com/insionng/vodka"
	m "github.com/insionng/vodka/middleware"
	"github.com/insionng/zenpress/handler"
	"github.com/vodka-contrib/jwt"
	"github.com/vodka-contrib/pongor"
	"net/http"
)

// Handler
func hello(self *vodka.Context) error {
	data := map[string]interface{}{}
	data["title"] = "你好，世界"
	self.Render(http.StatusOK, "index.html", data)
	return nil
}

var (
	key = "AppSkey"
)

func TokenHandler(self *vodka.Context) error {

	var claims = map[string]interface{}{}
	claims["username"] = "Insion"
	token, err := jwt.NewToken(key, claims)
	if err != nil {
		return err
	}
	// show the token use for test
	return self.String(http.StatusOK, "%s", token)
}

func main() {

	v := vodka.New()
	v.Use(m.Logger())
	v.Use(m.Recover())
	v.Use(m.Gzip())

	v.Static("/static/", "static")
	v.SetRenderer(pongor.Renderor())

	g := v.Group("")
	g.Get("/", handler.MainHandler)

	// Restricted group
	r := v.Group("")
	jwt.JWTContextKey = key
	jwt.Bearer = "zenpress"
	r.Use(jwt.JWTAuther(jwt.Options{
		KeyFunc: func(ctx *vodka.Context) (string, error) {
			return jwt.JWTContextKey, nil
		},
	}))
	r.Any("/jwt/", hello)

	/*

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

	*/
	v.Run(9000)
}
