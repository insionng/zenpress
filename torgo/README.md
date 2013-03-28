## Torgo version 0.7
=======

Torgo是一个与python的Tornado web框架相似的轻量级Golang Web框架，基于beego基础上增加各种功能


##licensed

Torgo is licensed under the BSD Licence


## Install
============
    go get github.com/insionng/torgo


## Base on Beego
============
增加了开发需要的一些辅助功能和特性，譬如:
输出html文档RanderString函数
控制器由Controller改为Handler，以符合从Python Tornado 转到Golang的用户使用习惯

调整优化了一些默认参数等等：
譬如默认的http服务端口不再是beego的8080 而直接是80,以方便本地调试
AutoRender自动渲染默认是false 以灵活控制渲染
增加一些模版函数

以下文档援引自beego，略作修改拿来展示我这个山寨版Beego和正版Beego的细节上的区别:)


## Quick Start
============
Here is the canonical "Hello, world" example app for torgo:
```go
package main

import (
	"github.com/insionng/torgo"
)

type MainHandler struct {
	torgo.Handler
}

func (self *MainHandler) Get() {
	self.Ctx.WriteString("hello world")
}

func main() {
	torgo.Route("/", &MainHandler{})
	//torgo.HttpPort = 80 // default
	torgo.Run()
}
```
	
	hello world
	
A more complete example use of torgo exists here:[toropress](https://github.com/insionng/toropress)


	
## Router
============
In torgo, a route is a struct paired with a URL-matching pattern. The struct has many method with the same name of http method to serve the http response. Each route is associated with a block.
```go
torgo.Route("/", &handlers.MainHandler{})
torgo.Route("/admin", &admin.UserHandler{})
torgo.Route("/admin/index", &admin.ArticleHandler{})
torgo.Route("/admin/addpkg", &admin.AddHandler{})
```
You can specify custom regular expressions for routes:
```go
torgo.Route("/admin/editpkg/:id([0-9]+)", &admin.EditHandler{})
torgo.Route("/admin/delpkg/:id([0-9]+)", &admin.DelHandler{})
torgo.Route("/:pkg(.*)", &handlers.MainHandler{})
```	
You can also create routes for static files:

	torgo.TorApp.SetStaticPath("/static","/public")
	
This will serve any files in /static, including files in subdirectories. For example request `/static/logo.gif` or `/static/style/main.css` will server with the file in the path `/pulic/logo.gif` or `/public/style/main.css`

## Filters / Middleware
============
You can apply filters to routes, which is useful for enforcing security, redirects, etc.

You can, for example, filter all request to enforce some type of security:
```go
var FilterUser = func(w http.ResponseWriter, r *http.Request) {
    if r.URL.User == nil || r.URL.User.Username() != "admin" {
        http.Error(w, "", http.StatusUnauthorized)
    }
}

torgo.Filter(FilterUser)
```	
You can also apply filters only when certain REST URL Parameters exist:
```go
torgo.Route("/:id([0-9]+)", &admin.EditHandler{})
torgo.FilterParam("id", func(rw http.ResponseWriter, r *http.Request) {
    ...
})
```	
Additionally, You can apply filters only when certain prefix URL path exist:
```go
torgo.FilterPrefixPath("/admin", func(rw http.ResponseWriter, r *http.Request) {
    … auth 
})
```		

## Handler / Struct
============ 	
To implement a torgo Handler, embed the `torgo.Handler` struct:
```go
type xxxHandler struct {
	torgo.Handler
}
```
`torgo.Handler` satisfieds the `torgo.HandlerInterface` interface, which defines the following methods:

- Init(ct *Context, cn string)

	this function is init the Context, ChildStruct' name and the Handler's variables.
	
- Prepare()

   this function is Run before the HTTP METHOD's Function,as follow defined. In the ChildStruct you can define this function to auth user or init database et.
   
- Get()

	When the HTTP' Method is GET, the torgo router will run this function.Default is HTTP-403. In the ChildStruct you must define the same functon to logical processing.
	
- Post()

	When the HTTP' Method is POST, the torgo router will run this function.Default is HTTP-403. In the ChildStruct you must define the same functon to logical processing.

- Delete()

	When the HTTP' Method is DELETE, the torgo router will run this function.Default is HTTP-403. In the ChildStruct you must define the same functon to logical processing.

- Put()

	When the HTTP' Method is PUT, the torgo router will run this function.Default is HTTP-403. In the ChildStruct you must define the same functon to logical processing.

- Head()

	When the HTTP' Method is HEAD, the torgo router will run this function.Default is HTTP-403. In the ChildStruct you must define the same functon to logical processing.

- Patch()

	When the HTTP' Method is PATCH, the torgo router will run this function.Default is HTTP-403. In the ChildStruct you must define the same functon to logical processing.

- Options()

	When the HTTP' Method is OPTIONS, the torgo router will run this function.Default is HTTP-403. In the ChildStruct you must define the same functon to logical processing.

- Finish()

	this function is run after the HTTP METHOD's Function,as previous defined. In the ChildStruct you can define this function to close database et.

- Render() error

	this function is to render the template as user defined. In the strcut you need to call.
	
- RenderString() (string, error)

- RenderBytes() ([]byte, error)


So you can define ChildStruct method to accomplish the interface's method, now let us see an example:
```go
type AddHandler struct {
	torgo.Handler
}

func (self *AddHandler) Prepare() {

}

func (self *AddHandler) Get() {
	self.Layout = "admin/layout.html"
	self.TplNames = "admin/add.tpl"
}

func (self *AddHandler) Post() {
	//data deal with
	self.Ctx.Request.ParseForm()
	pkgname := self.Ctx.Request.Form.Get("pkgname")
	content := self.Ctx.Request.Form.Get("content")
	torgo.Info(self.Ctx.Request.Form)
	pk := models.GetCruPkg(pkgname)
	if pk.Id == 0 {
		var pp models.PkgEntity
		pp.Pid = 0
		pp.Pathname = pkgname
		pp.Intro = pkgname
		models.InsertPkg(pp)
		pk = models.GetCruPkg(pkgname)
	}
	var at models.Article
	at.Pkgid = pk.Id
	at.Content = content
	models.InsertArticle(at)
	self.Ctx.Redirect(302, "/admin/index")
}
```
## View / Template
============ 		
### template view path

The default viewPath is `/views`, you can put the template file in the views. torgo will find the template from viewpath.

also you can modify the viewpaths like this:

	torgo.ViewsPath = "/myviewpath"
	
### template names
torgo will find the template from viewpath. the file is set by user like：

	self.TplNames = "admin/add.tpl"
	
then torgo will find the file in the path:`/views/admin/add.tpl`

if you don't set TplNames,torgo will find like this:

	c.TplNames = c.ChildName + "/" + c.Ctx.Request.Method + "." + c.TplExt

So if the ChildName="AddHandler",Request Method= "POST",default TplEXT="html"	
So torgo will file the file in the path:`/view/AddHandler/POST.tpl`

### autoRender
In the handler you need to call render function. torgo will not auto call this function after HTTP Method Call.

You can enable automatic invokation of autorender via the AutoRender Flag:
```go
torgo.AutoRender = false
```

### layout
torgo supports layouts for views. For example:
```go
self.Layout = "admin/layout.html"
self.TplNames = "admin/add.tpl"	
```

In layout.html you must define the variable like this to show sub template's content:

	{{.LayoutContent}}

torgo first parses the TplNames files, renders their content, and appends it to data["LayoutContent"].

### template function
torgo support users to define template function like this:
```go
func hello(in string)(out string){
	out = in + "world"
	return
}

torgo.AddFuncMap("hi",hello)
```

then in you template you can use it like this:

	{{.Content | hi}}
	
torgo has three default defined funtion:

- torgoTplFuncMap["markdown"] = MarkDown

	MarkDown parses a string in MarkDown format and returns HTML. Used by the template parser as "markdown"

- torgoTplFuncMap["dateformat"] = DateFormat

	DateFormat takes a time and a layout string and returns a string with the formatted date. Used by the template parser as "dateformat"

- torgoTplFuncMap["compare"] = Compare	

	Compare is a quick and dirty comparison function. It will convert whatever you give it to strings and see if the two values are equal.Whitespace is trimmed. Used by the template parser as "eq"
	
### JSON/XML output
You can use `torgo.Handler.ServeJson` or `torgo.Handler.ServeXml` for serializing to Json and Xml. I found myself constantly writing code to serialize, set content type, content length, etc. Feel free to use these functions to eliminate redundant code in your app.
		
Helper function for serving Json, sets content type to application/json:
```go
func (self *AddHandler) Get() {
    mystruct := { ... }
	self.Data["json"] = &mystruct
    self.ServeJson()
}
```
Helper function for serving Xml, sets content type to application/xml:
```go
func (self *AddHandler) Get() {
    mystruct := { ... }
	self.Data["xml"]=&mystruct
    self.ServeXml()
}
```

## Torgo Variables
============ 
torgo has many default variables, as follow is a list to show:

- TorApp       *App

	global app init by the torgo. You needn't to init it, just use it.
	
- AppName      string

	appname is what you project named, default is torgo

- AppPath      string

	this is the project path

- StaticDir    map[string]string

	staticdir store the map which request url to the static file path
	
	default is the request url has prefix `static`, then server the path in the app path
	
- HttpAddr     string

	http listen address, defult is ""

- HttpPort     int

	http listen port, default is 80

- RecoverPanic bool

	RecoverPanic mean when the program panic  whether the process auto recover,default is true

- AutoRender   bool

	whether run the Render function, default is false

- ViewsPath    string

	the template path, default is ./views

- RunMode      string //"dev" or "prod"

	the runmode ,default is prod

- AppConfig    *Config

    Appconfig is a result that parse file from conf/app.conf, if this file not exist then the variable is nil. if the file exist, then return the Config as follow.
	
- PprofOn bool

	default is false. turn on pprof, if set to true. you can visit like this:
	
		/debug/pprof
		/debug/pprof/cmdline
		/debug/pprof/profile
		/debug/pprof/symbol	
	this serves via its HTTP server runtime profiling data in the format expected by the pprof visualization tool. For more information about pprof, see	http://golang.org/pkg/net/http/pprof/

## Config
============ 

torgo support parse ini file, torgo will parse the default file in the path `conf/app.conf`

throw this conf file you can set many Torgo Variables to change default values.

app.conf

	appname = beepkg
	httpaddr = "127.0.0.1"
	httpport = 9090
	runmode ="dev"
	autorender = false
	autorecover = false
	viewspath = "myview"

this variables will replace the default torgo variable's values

you can also set you own variables such as database setting

	mysqluser = "root"
	mysqlpass = "rootpass"
	mysqlurls = "127.0.0.1"
	mysqldb   = "torgo"
	
In you app you can get the config like this:

	torgo.AppConfig.String("mysqluser")
	torgo.AppConfig.String("mysqlpass")
	torgo.AppConfig.String("mysqlurls")
	torgo.AppConfig.String("mysqldb")

## Logger
============ 	
torgo has a default log named BeeLogger which output to os.Stdout.

you can change it output with the standard log.Logger like this:

	fd,err := os.OpenFile("/opt/app/beepkg/beepkg.log", os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		torgo.Critical("openfile beepkg.log:", err)
		return
	}
	lg := log.New(fd, "", log.Ldate|log.Ltime)
	torgo.SetLogger(lg)
	

### Supported log levels
- Trace - For pervasive information on states of all elementary constructs. Use 'Trace' for in-depth debugging to find problem parts of a function, to check values of temporary variables, etc.
- Debug - For detailed system behavior reports and diagnostic messages to help to locate problems during development.
- Info - For general information on the application's work. Use 'Info' level in your code so that you could leave it 'enabled' even in production. So it is a 'production log level'.
- Warn - For indicating small errors, strange situations, failures that are automatically handled in a safe manner.
- Error - For severe failures that affects application's workflow, not fatal, however (without forcing app shutdown).
- Critical - For producing final messages before application’s death. Note: critical messages force immediate flush because in critical situation it is important to avoid log message losses if app crashes.
- Off - A special log level used to turn off logging

torgo has follow functions:

- Trace(v ...interface{})
- Debug(v ...interface{})
- Info(v ...interface{})
- Warn(v ...interface{})
- Error(v ...interface{})
- Critical(v ...interface{})

you can set log levels like this :

	torgo.SetLevel(torgo.LevelError)

after set the log levels, in the logs function which below the setlevels willn't output anything

after set levels to torgo.LevelError

Trace, Debug, Info, Warn will not output anything. So you can change it when in dev and prod mode.