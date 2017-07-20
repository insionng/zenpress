//root routers
root = app.Group("/root")
root.Use(switchr.SwitchrWithConfig(&switchr.SwitchrConfig{Theme: theme, Filter: filter, Reload: reload}))

//Dashboard：控制面板
root.Any("/", RootDashboardHandler)

//article：文章
root.Any("/article", RootArticleHandler)

//media：媒体
root.Any("/media", RootMediaHandler)

//media：链接
root.Any("/link", RootLinkHandler)

//page：页面
root.Any("/page", RootPageHandler)

//comments：评论
root.Any("/comment", RootCommentHandler)

//theme：主题
root.Any("/theme", RootThemeHandler)

//plugin：插件
root.Any("/plugin", RootPluginHandler)

//users：用户
root.Any("/user", RootUserHandler)

//tools：工具
root.Any("/tool", RootToolHandler)

//options：设置
root.Any("/option", RootOptionHandler)
