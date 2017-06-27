package core

import (
	"fmt"
	"io/ioutil"
	"os"

	gomakross "github.com/insionng/makross"
	"github.com/insionng/makross/file"
	gologger "github.com/insionng/makross/logger"
	gopongor "github.com/insionng/makross/pongor"
	gosession "github.com/insionng/makross/session"
	gostatic "github.com/insionng/makross/static"
	goswitchr "github.com/insionng/zenpress/module/switchr"

	mfmt "github.com/insionng/zenpress/extension/fmt"
	"github.com/insionng/zenpress/extension/github.com/insionng/makross"
	"github.com/insionng/zenpress/extension/github.com/insionng/makross/logger"
	"github.com/insionng/zenpress/extension/github.com/insionng/makross/pongor"
	"github.com/insionng/zenpress/extension/github.com/insionng/makross/session"
	"github.com/insionng/zenpress/extension/github.com/insionng/makross/static"
	"github.com/insionng/zenpress/extension/github.com/insionng/zenpress/model"
	"github.com/insionng/zenpress/extension/github.com/insionng/zenpress/module/hook"
	"github.com/insionng/zenpress/extension/github.com/insionng/zenpress/module/switchr"

	"qlang.io/cl/qlang"
	"qlang.io/lib/qlang.all"
)

var (
	Q = qlang.New()
)

func VmByte(code []byte) error {
	return Q.SafeExec(code, "")
}

func VmString(code string) error {
	return VmByte([]byte(code))
}

func init() {
	qall.InitSafe(true)
	qlang.Import("makross", makross.Exports)
	qlang.Import("logger", logger.Exports)
	qlang.Import("switchr", switchr.Exports)
	qlang.Import("pongor", pongor.Exports)
	qlang.Import("static", static.Exports)
	qlang.Import("session", session.Exports)
	qlang.Import("model", model.Exports)

	qlang.Import("hook", hook.Exports)
	qlang.Import("fmt", mfmt.Exports)
}

func AddFunc(name string, function interface{}, pack ...string) {
	Exports := map[string]interface{}{
		name: function}
	var pname string
	if len(pack) > 0 {
		pname = pack[0]
	}
	qlang.Import(pname, Exports)
}

func dirs(dir string) (dirlist []string) {
	f, _ := os.Open(dir)
	defer f.Close()
	dirs, _ := f.Readdir(0)
	for _, fileInfo := range dirs {
		if fileInfo.IsDir() {
			dirlist = append(dirlist, fileInfo.Name())
		}
	}
	return dirlist
}

func Plugins() {
	pluginDir := "content/plugin"
	plugin := dirs(pluginDir)
	for _, plugin := range plugin {

		b, e := ioutil.ReadFile(fmt.Sprintf("%s/%s/%s.app", pluginDir, plugin, plugin))
		if e != nil {
			fmt.Printf("Plugin[%v] ReadFile Error:%v\n", plugin, e)
			continue
		}

		if len(b) > 0 {
			e = VmByte(b)
			if e != nil {
				fmt.Printf("Plugin[%v] VmByte Error:%v\n", plugin, e)
				//panic(e)
			}
		}
	}
}

func readfile(filename string) (b []byte, e error) {
	b, e = ioutil.ReadFile(filename)
	if e != nil {
		return nil, e
	}
	return
}

func Codes(theme string) []byte {
	var applicationDir = "content/application"
	var themeDir = "content/theme"
	var themeApps, rootApps string

	//读取前端逻辑代码
	files := []string{"IndexHandler", "SingleHandler", "PageHandler", "CategoryHandler", "TagHandler", "TaxonomyHandler", "AuthorHandler", "AttachmentHandler", "DateHandler", "ArchiveHandler", "SearchHandler", "NotFoundHandler"}
	for _, file := range files {
		var b []byte
		var e error
		if b, e = readfile(fmt.Sprintf("%s/%s/handler/%s", themeDir, theme, fmt.Sprintf("%s.app", file))); e != nil {
			if b, e = readfile(fmt.Sprintf("%s/%s", applicationDir, fmt.Sprintf("%s.app", file))); e != nil {
				panic(fmt.Sprintf("Not found %s.app in %s", file, applicationDir))
			}
		}
		themeApps = themeApps + fmt.Sprintf("%s\n", b)
	}

	//读取后端逻辑代码
	rootfiles := []string{"DashboardHandler", "ArticleHandler", "MediaHandler", "LinkHandler", "PageHandler", "CommentHandler", "ThemeHandler", "PluginHandler", "UserHandler", "ToolHandler", "OptionHandler", "NotFoundHandler"}
	for _, file := range rootfiles {
		var b []byte
		var e error
		if b, e = readfile(fmt.Sprintf("%s/%s/handler/root/%s", themeDir, theme, fmt.Sprintf("%s.app", file))); e != nil {
			if b, e = readfile(fmt.Sprintf("%s/root/%s", applicationDir, fmt.Sprintf("%s.app", file))); e != nil {
				panic(fmt.Sprintf("Not found %s.app in %s/root", file, applicationDir))
			}
		}
		if len(b) > 0 {
			rootApps = rootApps + fmt.Sprintf("%s\n", b)
		}
	}

	//读取主程序逻辑
	var application = fmt.Sprintf("%s/%s", applicationDir, "Application.app")
	appCode, e := readfile(application)
	if e != nil {
		panic(fmt.Sprintf("Not found %s in %s", application, applicationDir))
	}
	if len(appCode) > 0 {
		appCodes := []byte(fmt.Sprintf("%s\n%s\n%s", themeApps, rootApps, appCode))
		return appCodes
	}

	return nil
}

func FrontCodes(theme string) []byte {
	var applicationDir = "content/application"
	var themeDir = "content/theme"
	var themeApps, rootApps string

	//读取前端控制器逻辑代码
	files := []string{"IndexHandler", "SingleHandler", "PageHandler", "CategoryHandler", "TagHandler", "TaxonomyHandler", "AuthorHandler", "AttachmentHandler", "DateHandler", "ArchiveHandler", "SearchHandler", "SigninHandler", "NotFoundHandler"}
	for _, file := range files {
		var b []byte
		var e error
		if b, e = readfile(fmt.Sprintf("%s/%s/handler/%s", themeDir, theme, fmt.Sprintf("%s.app", file))); e != nil {
			if b, e = readfile(fmt.Sprintf("%s/%s", applicationDir, fmt.Sprintf("%s.app", file))); e != nil {
				panic(fmt.Sprintf("Not found %s.app in %s", file, applicationDir))
			}
		}
		themeApps = themeApps + fmt.Sprintf("%s\n", b)
	}

	//读取前端路由逻辑
	var application = fmt.Sprintf("%s/%s", applicationDir, "Frontkend.app")
	appCode, e := readfile(application)
	if e != nil {
		panic(fmt.Sprintf("Not found %s in %s", application, applicationDir))
	}
	if len(appCode) > 0 {
		appCodes := []byte(fmt.Sprintf("%s\n%s\n%s", themeApps, rootApps, appCode))
		return appCodes
	}

	return nil
}

func BackCodes(theme string) []byte {
	var applicationDir = "content/application"
	var themeDir = "content/theme"
	var themeApps, rootApps string

	//读取后端逻辑代码
	rootfiles := []string{"DashboardHandler", "ArticleHandler", "MediaHandler", "LinkHandler", "PageHandler", "CommentHandler", "ThemeHandler", "PluginHandler", "UserHandler", "ToolHandler", "OptionHandler", "NotFoundHandler"}
	for _, file := range rootfiles {
		var b []byte
		var e error
		if b, e = readfile(fmt.Sprintf("%s/%s/handler/root/%s", themeDir, theme, fmt.Sprintf("%s.app", file))); e != nil {
			if b, e = readfile(fmt.Sprintf("%s/root/%s", applicationDir, fmt.Sprintf("%s.app", file))); e != nil {
				panic(fmt.Sprintf("Not found %s.app in %s/root", file, applicationDir))
			}
		}
		if len(b) > 0 {
			rootApps = rootApps + fmt.Sprintf("%s\n", b)
		}
	}

	//读取后端路由逻辑
	var application = fmt.Sprintf("%s/%s", applicationDir, "Backend.app")
	appCode, e := readfile(application)
	if e != nil {
		panic(fmt.Sprintf("Not found %s in %s", application, applicationDir))
	}
	if len(appCode) > 0 {
		appCodes := []byte(fmt.Sprintf("%s\n%s\n%s", themeApps, rootApps, appCode))
		return appCodes
	}

	return nil
}

func GetAppByTheme(theme string, filter bool, reload bool) (*gomakross.Makross, bool) {
	logWriter, err := os.OpenFile("content/storage/logs/makross.log", os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(fmt.Errorf("Create makross log file error:%v", err))
	}
	app := gomakross.New()
	app.Use(gologger.LoggerWithConfig(gologger.LoggerConfig{Output: logWriter}))
	app.Use(gostatic.Static(fmt.Sprintf("content/theme/%s/public", theme)))
	app.Use(gosession.Sessioner(gosession.Options{"file", `{"cookieName":"makrossSessionId","gcLifetime":3600,"providerConfig":"./content/storage/session"}`}))
	app.Use(goswitchr.SwitchrWithConfig(goswitchr.SwitchrConfig{Theme: theme, Filter: filter, Reload: reload}))

	app.SetRenderer(gopongor.Renderor(gopongor.Option{Directory: fmt.Sprintf("content/theme/%s/template", theme), Reload: reload, Filter: filter}))

	Q.SetVar("app", app)

	appFrontCode := FrontCodes(theme)
	err = VmByte(appFrontCode)
	if err != nil {
		panic(fmt.Errorf("#1 执行前端代码出错：%v", err))
	}

	m, okay := Q.GetVar("app")
	if !okay {
		panic("cannot find app object in front code.")
	}

	app, okay = m.(*gomakross.Makross)
	if !okay {
		panic("cannot convert FrontCodes.app to (*makross.Makross)")
	}

	appBackCode := BackCodes(theme)
	Q.SetVar("theme", theme)
	Q.SetVar("filter", filter)
	Q.SetVar("reload", reload)

	err = VmByte(appBackCode)
	if err != nil {
		panic(fmt.Errorf("#1 执行后端代码出错：%v", err))
	}

	m, okay = Q.GetVar("app")
	if !okay {
		panic("cannot find app object in back code.")
	}

	app, okay = m.(*gomakross.Makross)
	if okay {
		app.Get("/*", file.Server(file.PathMap{
			//"/":      fmt.Sprintf("content/theme/%s/public", theme),
			"/root/": "/public/",
		}))
	}

	return app, okay

}
