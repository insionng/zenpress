package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type App struct {
	ID          uint64    `gorm:"primary_key"`
	SiteID      uint64    `gorm:"not null default 0 BIGINT(20)"`
	Domain      string    `gorm:"not null default '' index(domain) VARCHAR(200)"`
	Path        string    `gorm:"not null default '' index(domain) VARCHAR(100)"`
	Registered  time.Time `gorm:"not null default '0000-00-00 00:00:00' DATETIME"`
	LastUpdated time.Time `gorm:"not null default '0000-00-00 00:00:00' DATETIME"`
	Public      int       `gorm:"not null default 1 TINYINT(2)"`
	Archived    int       `gorm:"not null default 0 TINYINT(2)"`
	Mature      int       `gorm:"not null default 0 TINYINT(2)"`
	Spam        int       `gorm:"not null default 0 TINYINT(2)"`
	Deleted     int       `gorm:"not null default 0 TINYINT(2)"`
	LangID      int       `gorm:"not null default 0 index INT(11)"`
}

// NewApp 创建应用
func NewApp(app *App) (db *gorm.DB) {
	db = Database.Create(app)
	return
}

// AddApp 新增应用
func AddApp(domain, path string) (db *gorm.DB) {
	db = Database.Create(&App{Domain: domain, Path: path})
	return
}

// GetApp 获得应用
func GetApp(id uint64) (db *gorm.DB, App App) {
	db = Database.First(&App, "id = ?", id)
	return
}

// UpdateApp 更新应用
func UpdateApp(id uint64, domain, path string) (db *gorm.DB, app App) {
	app = App{Domain: domain, Path: path}
	db = Database.Model(&app).Update("id", id)
	return
}

// DeleteApp 删除应用
func DeleteApp(id uint64) (db *gorm.DB) {
	db = Database.Delete(App{}, "id = ?", id)
	return
}
