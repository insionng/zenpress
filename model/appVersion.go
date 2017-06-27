package model

import (
	"time"
)

type AppVersion struct {
	AppID       uint64    `gorm:"primary_key"`
	DbVersion   string    `gorm:"not null default '' index VARCHAR(20)"`
	LastUpdated time.Time `gorm:"not null default '0000-00-00 00:00:00' DATETIME"`
}
