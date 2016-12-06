package main

import (
	"github.com/insionng/macross"
	"github.com/insionng/macross/blimit"
	"github.com/insionng/macross/compress"
	"github.com/insionng/macross/cors"
	"github.com/insionng/macross/csrf"
	"github.com/insionng/macross/jwt"
	"github.com/insionng/macross/logger"
	"github.com/insionng/macross/pongor"
	"github.com/insionng/macross/recover"
	"github.com/insionng/macross/secure"
	"github.com/insionng/macross/slash"
	"github.com/insionng/macross/static"
	"github.com/insionng/zenpress/handler"
	"time"
)

func main() {

	v := macross.New()
	v.Use(slash.AddTrailingSlash())

	v.Use(logger.Logger())
	v.Use(recover.Recover())
	v.Use(compress.Gzip())
	v.Use(secure.Secure())
	v.Use(blimit.BodyLimit("2M"))
	v.Use(csrf.CSRFWithConfig(csrf.CSRFConfig{
		TokenLookup: "header:X-XSRF-TOKEN",
	}))
	v.Use(cors.CORS())
	v.Use(static.Static("static"))
	v.SetRenderer(pongor.Renderor())

	v.File("/favicon.ico", "static/ico/favicon.ico")

	v.Get("/", handler.MainHandler)

	v.Get("/signup/", handler.GetSignupHandler)
	v.Post("/signup/", handler.PostSignupHandler)

	v.Get("/signin/", handler.GetSigninHandler)
	v.Post("/signin/", handler.PostSigninHandler)

	v.Get("/signout/", handler.GetSignoutHandler)

	v.Any("/search/", handler.SearchHandler)
	v.Get("/node/<nid>/", handler.GetNodeHandler)
	v.Get("/view/<tid>/", handler.GetViewHandler)
	v.Get("/category/<cid>/", handler.MainHandler)

	r := v.Group("/apis/v1", jwt.JWTWithConfig(jwt.JWTConfig{
		SigningKey: []byte("ZeNpReSe"),
		Expires:    time.Minute * 60,
	}))

	v.Get("/new/category/", handler.GetNewCategoryHandler)
	r.Post("/new/category/", handler.PostNewCategoryHandler)

	v.Get("/new/node/", handler.NewNodeGetHandler)
	r.Post("/new/node/", handler.NewNodePostHandler)

	v.Get("/new/topic/", handler.NewTopicGetHandler)
	r.Post("/new/topic/", handler.NewTopicPostHandler)

	r.Post("/new/reply/<tid>/", handler.NewReplyPostHandler)

	v.Get("/modify/category/", handler.ModifyCatGetHandler)
	r.Post("/modify/category/", handler.ModifyCatPostHandler)

	v.Get("/modify/node/", handler.ModifyNodeGetHandler)
	r.Post("/modify/node/", handler.ModifyNodePostHandler)

	r.Any("/topic/delete/<tid>/", handler.TopicDeleteHandler)

	v.Get("/topic/edit/<tid>/", handler.TopicEditGetHandler)
	r.Post("/topic/edit/<tid>/", handler.TopicEditPostHandler)

	r.Any("/node/delete/<nid>/", handler.NodeDeleteHandler)

	v.Get("/node/edit/<nid>/", handler.NodeEditGetHandler)
	r.Post("/node/edit/<nid>/", handler.NodeEditPostHandler)

	r.Any("/delete/reply/<rid>/", handler.DeleteReplyHandler)

	//hotness
	r.Any("/like/<name>/<id>/", handler.LikeHandler)
	r.Any("/hate/<name>/<id>/", handler.HateHandler)

	v.Run(":9000")
}
