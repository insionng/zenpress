app.Get("/", IndexHandler)
app.Get("/post", SingleHandler)
app.Get("/page", PageHandler)
app.Get("/category", CategoryHandler)
app.Get("/tag", TagHandler)
app.Get("/taxonomy", TaxonomyHandler)
app.Get("/author", AuthorHandler)
app.Get("/attachment", AttachmentHandler)
app.Get("/date", DateHandler)
app.Get("/archive", ArchiveHandler)
app.Get("/search", SearchHandler)

//root routers
rg = app.Group("/root")

//Dashboard：控制面板
rg.Get("/", RootDashboardHandler)

//article：文章
rg.Get("/article", RootArticleHandler)

//media：媒体
rg.Get("/media", RootMediaHandler)

//media：链接
rg.Get("/link", RootLinkHandler)

//page：页面
rg.Get("/page", RootPageHandler)

//comments：评论
rg.Get("/comment", RootCommentHandler)

//theme：主题
rg.Get("/theme", RootThemeHandler)

//plugin：插件
rg.Get("/plugin", RootPluginHandler)

//users：用户
rg.Get("/user", RootUserHandler)

//tools：工具
rg.Get("/tool", RootToolHandler)

//options：设置
rg.Get("/option", RootOptionHandler)
