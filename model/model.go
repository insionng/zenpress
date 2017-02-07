package model

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"

	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

var (
	app       = "app"
	Engine    *xorm.Engine
	HasEngine bool

	DataType      = "mysql"
	DBConnect     = "root:rootpass@/wp?charset=utf8"
	DBTablePrefix = "wp_"
)

func ConDb() (*xorm.Engine, error) {
	switch {
	case DataType == "memory":
		return xorm.NewEngine("tidb", "memory://tidb/tidb")

	case DataType == "goleveldb":
		if DBConnect != "" {
			return xorm.NewEngine("tidb", DBConnect)
		}
		return xorm.NewEngine("tidb", "goleveldb://"+app+"/data/tidb/tidb")

	case DataType == "boltdb":
		if DBConnect != "" {
			return xorm.NewEngine("tidb", DBConnect)
		}
		return xorm.NewEngine("tidb", "boltdb://"+app+"/data/tidb/tidb")
	case DataType == "sqlite":
		if DBConnect != "" {
			return xorm.NewEngine("sqlite3", DBConnect)
		}
		return xorm.NewEngine("sqlite3", app+"/data/sqlite.db")

	case DataType == "mysql":
		return xorm.NewEngine("mysql", DBConnect)
		//return xorm.NewEngine("mysql", "root:YouPass@/db?charset=utf8")

	case DataType == "postgres":
		return xorm.NewEngine("postgres", DBConnect)
		//return xorm.NewEngine("postgres", "user=postgres password=jn!@#$%^&* dbname=pgsql sslmode=disable")

		// "user=postgres password=jn!@#$%^&* dbname=yougam sslmode=disable maxcons=10 persist=true"
		//return xorm.NewEngine("postgres", "host=192.168.1.113 user=postgres password=jn!@#$%^&* dbname=yougam sslmode=disable")
		//return xorm.NewEngine("postgres", "host=127.0.0.1 port=6432 user=postgres password=jn!@#$%^&* dbname=yougam sslmode=disable")
	}
	return nil, errors.New("Unknown database type..")
}

func SetEngine() (*xorm.Engine, error) {
	var _error error
	if Engine, _error = ConDb(); _error != nil {
		return nil, fmt.Errorf("Fail to connect to database: %s", _error.Error())
	} else {
		Engine.SetTableMapper(core.NewPrefixMapper(core.GonicMapper{}, DBTablePrefix))
		Engine.SetColumnMapper(core.GonicMapper{})

		cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 10240)
		Engine.SetDefaultCacher(cacher)

		logDir := "logs"
		if _, e := os.Open(logDir); e != nil {
			os.MkdirAll(logDir, os.ModePerm)
		}
		
		logPath := path.Join(logDir, "xorm.log")
		f, err := os.Create(logPath)
		if err != nil {
			return Engine, fmt.Errorf("Fail to create xorm.log: %s", err.Error())
		}

		Engine.SetLogger(xorm.NewSimpleLogger(f))
		Engine.ShowSQL(false)

		if location, err := time.LoadLocation("Asia/Shanghai"); err == nil {
			Engine.TZLocation = location
		}

		return Engine, err
	}
}

func NewEngine() error {
	var _error error
	Engine, _error = SetEngine()
	return _error
}

func init() {

	var _error error
	if Engine, _error = SetEngine(); _error != nil {
		log.Fatal("app.models.init() errors:", _error.Error())
	}

	if _error = createTables(Engine); _error != nil {
		log.Fatal("Fail to creatTables errors:", _error.Error())
	}

	inits()

}

func Ping() error {
	return Engine.Ping()
}

func createTables(Engine *xorm.Engine) error {
	return Engine.Sync2(&BlogVersions{},&Blogs{},&Commentmeta{},&Comments{},&Links{},&Options{},&Postmeta{},&Posts{},&RegistrationLog{},&Signups{},&Site{},&Sitemeta{},&TermRelationships{},&TermTaxonomy{},&Termmeta{},&Terms{},&Usermeta{},&Users{})
}

func inits() {
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("The app system has started!")
}
