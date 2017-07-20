package model

import (
	"os"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"errors"
	"fmt"
	"log"
	"time"
)

const (
	app = "./"

	DataType            = "sqlite"
	DatabaseTablePrefix = "zen_"
)

var (
	Database    *gorm.DB
	HasDatabase bool

	//DatabaseConn        = "root:rootpass@/wp?charset=utf8"
	DatabaseConn = "content/storage/data/sqlite.db"
	//DatabaseConn = "../content/storage/data/sqlite_test.db"
)

// Model base model definition, including fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`, which could be embedded in your models
//    type User struct {
//      model.Model
//    }
type Model struct {
	ID        uint64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func ConnDatabase(databaseConn string) (*gorm.DB, error) {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return DatabaseTablePrefix + defaultTableName
	}

	switch {
	case DataType == "sqlite":
		return gorm.Open("sqlite3", databaseConn)

	case DataType == "mysql":
		return gorm.Open("mysql", databaseConn)

	case DataType == "postgres":
		return gorm.Open("postgres", databaseConn)
	}

	return nil, errors.New("Unknown Database Type")
}

func SetDatabase(databaseConn string, maxIdleConns, maxOpenConns int) (*gorm.DB, error) {
	var _error error
	if Database, _error = ConnDatabase(databaseConn); _error != nil {
		return nil, fmt.Errorf("Failed to connect database: %s,%s", _error.Error(), DatabaseConn)
	}

	//Database Connection Pool
	Database.DB().SetMaxIdleConns(maxIdleConns)
	Database.DB().SetMaxOpenConns(maxOpenConns)

	Database.Debug().LogMode(true)

	logWriter, err := os.OpenFile("content/storage/logs/gorm.log", os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(fmt.Errorf("Create gorm log file error:%v", err))
	}
	Database.SetLogger(log.New(logWriter, "\r\n", 0))
	return Database, nil
}

func NewDatabase(databaseConn string, maxIdleConns, maxOpenConns int) error {
	var _error error
	Database, _error = SetDatabase(databaseConn, maxIdleConns, maxOpenConns)
	return _error
}

func init() {
	var _error error
	if Database, _error = SetDatabase(DatabaseConn, 10, 100); _error != nil {
		log.Fatal("app.models.init() errors:", _error.Error())
	}

	CreateTables(Database)
	message()
}

func CreateTables(Database *gorm.DB) {
	if DataType == "mysql" {
		Database.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&AppVersion{}, &App{}, &Commentmeta{}, &Comment{}, &Link{}, &Option{}, &Postmeta{}, &Post{}, &RegistrationLog{}, &Signup{}, &Site{}, &Sitemeta{}, &TermRelationship{}, &TermTaxonomy{}, &Termmeta{}, &Term{}, &Usermeta{}, &User{}, &Role{}, &Permission{}, &RolePermission{})
	} else {
		Database.AutoMigrate(&AppVersion{}, &App{}, &Commentmeta{}, &Comment{}, &Link{}, &Option{}, &Postmeta{}, &Post{}, &RegistrationLog{}, &Signup{}, &Site{}, &Sitemeta{}, &TermRelationship{}, &TermTaxonomy{}, &Termmeta{}, &Term{}, &Usermeta{}, &User{}, &Role{}, &Permission{}, &RolePermission{})
	}
}

func Ping() error {
	return Database.DB().Ping()
}

func message() {
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("The Zenpress system has started!")
}
