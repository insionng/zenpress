package main

import (
	"github.com/insionng/vodka"
	//"github.com/insionng/zenpress/handlers"

	"github.com/insionng/zenpress/handler"
	"net/http"

	m "github.com/insionng/vodka/middleware"
	"github.com/vodka-contrib/jwt"
	"github.com/vodka-contrib/pongor"
	"github.com/vodka-contrib/sessions"
	"github.com/vodka-contrib/vodkapprof"
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

	store := sessions.NewCookieStore()
	v.Use(sessions.Sessions("vodka", store))

	v.Use(m.Logger())
	v.Use(m.Recover())
	v.Use(m.Gzip())

	v.Static("/static/", "static")
	v.Favicon("static/ico/favicon.ico")
	v.SetRenderer(pongor.Renderor())

	g := v.Group("")
	g.Get("/", handler.MainHandler)
	g.Get("/signup/", handler.SignupHandler)
	g.Get("/signin/", handler.SigninHandler)
	g.Get("/signout/", handler.SignoutHandler)

	g.Any("/search/", handler.SearchHandler)
	g.Get("/node/:nid/", handler.NodeHandler)
	g.Get("/view/:tid/", handler.ViewHandler)
	g.Get("/category/:cid/", handler.MainHandler)

	// Restricted group
	r := v.Group("")
	jwt.JWTContextKey = key
	jwt.Bearer = "Zenpress"
	r.Use(jwt.JWTAuther(jwt.Options{
		KeyFunc: func(ctx *vodka.Context) (string, error) {
			return jwt.JWTContextKey, nil
		},
	}))

	r.Get("/new/category/", handler.NewCategoryGetHandler)
	r.Post("/new/category/", handler.NewCategoryPostHandler)

	r.Get("/new/node/", handler.NewNodeGetHandler)
	r.Post("/new/node/", handler.NewNodePostHandler)

	r.Get("/new/topic/", handler.NewTopicGetHandler)
	r.Post("/new/topic/", handler.NewTopicPostHandler)

	r.Post("/new/reply/:tid/", handler.NewReplyPostHandler)

	r.Get("/modify/category/", handler.ModifyCatGetHandler)
	r.Post("/modify/category/", handler.ModifyCatPostHandler)

	r.Get("/modify/node/", handler.ModifyNodeGetHandler)
	r.Post("/modify/node/", handler.ModifyNodePostHandler)

	/*


		//hotness
		beego.Router("/like/:name:string/:id:int", &handlers.LikeHandler{})
		beego.Router("/hate/:name:string/:id:int", &handlers.HateHandler{})


		beego.Router("/topic/delete/:tid:int", &handlers.TopicDeleteHandler{})
		beego.Router("/topic/edit/:tid:int", &handlers.TopicEditHandler{})

		beego.Router("/node/delete/:nid:int", &handlers.NodeDeleteHandler{})
		beego.Router("/node/edit/:nid:int", &handlers.NodeEditHandler{})

		beego.Router("/delete/reply/:rid:int", &handlers.DeleteReplyHandler{})

	*/

	// e.g. /debug/pprof, /debug/pprof/heap, etc.
	vodkapprof.Wrapper(v)
	v.Run(9000)
}
