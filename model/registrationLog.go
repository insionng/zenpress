package model

import (
	"time"
)

type RegistrationLog struct {
	ID             uint64    `gorm:"primary_key"`
	Email          string    `gorm:"not null default '' VARCHAR(255)"`
	IP             string    `gorm:"not null default '' index VARCHAR(30)"`
	AppID          uint64    `gorm:"not null default 0 BIGINT(20)"`
	DateRegistered time.Time `gorm:"not null default '0000-00-00 00:00:00' DATETIME"`
}
