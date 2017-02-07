package models

import (
	_ "github.com/insionng/zenpress/libraries/go-sql-driver/mysql"
	_ "github.com/insionng/zenpress/libraries/lib/pq"

	_ "github.com/go-xorm/tidb"
	_ "github.com/pingcap/tidb"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"

	"errors"
	"fmt"
	"github.com/insionng/zenpress/helper"
	"log"
	"os"
	"path"
	"time"
)

var (
	Engine    *xorm.Engine
	HasEngine bool

	DataType       = helper.DataType
	DBConnect      = helper.DBConnect
	FileStorageDir = "../"
	DBTablePrefix  = "nilyes_"
)

func ConDb() (*xorm.Engine, error) {
	switch {
	case DataType == "memory":
		return xorm.NewEngine("tidb", "memory://tidb/tidb")

	case DataType == "goleveldb":
		if DBConnect != "" {
			return xorm.NewEngine("tidb", DBConnect)
		}
		return xorm.NewEngine("tidb", "goleveldb://"+FileStorageDir+"data/tidb/tidb")

	case DataType == "boltdb":
		if DBConnect != "" {
			return xorm.NewEngine("tidb", DBConnect)
		}
		return xorm.NewEngine("tidb", "boltdb://"+FileStorageDir+"data/tidb/tidb")
	case DataType == "sqlite":
		if DBConnect != "" {
			return xorm.NewEngine("sqlite3", DBConnect)
		}
		return xorm.NewEngine("sqlite3", FileStorageDir+"data/sqlite.db")

	case DataType == "mysql":
		return xorm.NewEngine("mysql", DBConnect)
		//return xorm.NewEngine("mysql", "root:YouPass@/db?charset=utf8")

	case DataType == "postgres":
		return xorm.NewEngine("postgres", DBConnect)
		//return xorm.NewEngine("postgres", "user=postgres password=jn!@#$%^&* dbname=pgsql sslmode=disable")

		// "user=postgres password=jn!@#$%^&* dbname=zenpress sslmode=disable maxcons=10 persist=true"
		//return xorm.NewEngine("postgres", "host=192.168.1.113 user=postgres password=jn!@#$%^&* dbname=zenpress sslmode=disable")
		//return xorm.NewEngine("postgres", "host=127.0.0.1 port=6432 user=postgres password=jn!@#$%^&* dbname=zenpress sslmode=disable")
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

		logPath := path.Join(FileStorageDir+"logs", "xorm.log")
		os.MkdirAll(path.Dir(logPath), os.ModePerm)
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
		log.Fatal("zenpress.models.init() errors:", _error.Error())
	}

	if _error = createTables(Engine); _error != nil {
		log.Fatal("Fail to creatTables errors:", _error.Error())
	}

	initData()

}

func Ping() error {
	return Engine.Ping()
}

func createTables(Engine *xorm.Engine) error {
	return Engine.Sync2(&User{}, &Balance{}, &Message{}, &HistoryMessage{}, &Friend{}, &Category{}, &Node{}, &Topic{}, &Page{}, &Reply{}, &Attachment{}, &Link{}, &Notification{}, &UserMark{}, &TopicMark{}, &ReplyMark{}, &NodeMark{}, &ReportMark{})
}

func initData() {
	//用户等级划分：正数是普通用户，负数是管理员各种等级划分，为0则尚未注册
	if usr, err := GetUserByRole(-1000); usr == nil && err == nil {
		if row, err := AddUser("root@zenpress.com", "root", "root", "root", helper.EncryptHash("rootpass", nil), "", "", "", 1, -1000); err == nil && row > 0 {
			log.Println("Default Email:root@zenpress.com ,Username:root ,Password:rootpass")

			if usr, err := GetUserByRole(-1000); usr != nil && err == nil {
				SetAmountByUid(usr.Id, 2, 2000, "注册收益2000金币")
			}

		} else {
			log.Println("create root got errors:", err)
		}

	}

	if cats, err := GetCategories(0, 0, "id"); cats == nil && err == nil {
		log.Println(AddCategory("默认分类", "默认分类内容简介", "", 0))
	}

	fmt.Println("-----------------------------------------------------------")
	fmt.Println("The zenpress system has started!")
}
