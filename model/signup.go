package model

import (
	"time"
)

type Signup struct {
	SignupID      uint64    `gorm:"primary_key"`
	Domain        string    `gorm:"not null default '' index(domain_path) VARCHAR(200)"`
	Path          string    `gorm:"not null default '' index(domain_path) VARCHAR(100)"`
	Title         string    `gorm:"not null LONGTEXT"`
	UserLogin     string    `gorm:"not null default '' index(user_login_email) VARCHAR(60)"`
	UserEmail     string    `gorm:"not null default '' index index(user_login_email) VARCHAR(100)"`
	Registered    time.Time `gorm:"not null default '0000-00-00 00:00:00' DATETIME"`
	Activated     time.Time `gorm:"not null default '0000-00-00 00:00:00' DATETIME"`
	Active        int       `gorm:"not null default 0 TINYINT(1)"`
	ActivationKey string    `gorm:"not null default '' index VARCHAR(50)"`
	Meta          string    `gorm:"LONGTEXT"`
}
