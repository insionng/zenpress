package model

import (
	"time"
)

type App struct {
	AppID       uint64    `gorm:"primary_key"`
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
