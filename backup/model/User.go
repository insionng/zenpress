package model

import (
	"time"
)

// User 存储用户名、密码、昵称、邮箱、注册时间等信息；
type User struct {
	Id                int64     `xorm:"BIGINT(20)"`                                      //自增唯一ID
	UserLogin         string    `xorm:"not null default '' index VARCHAR(60)"`           //登录名
	UserPass          string    `xorm:"not null default '' VARCHAR(255)"`                //密码
	UserNicename      string    `xorm:"not null default '' index VARCHAR(50)"`           //昵称
	UserEmail         string    `xorm:"not null default '' index VARCHAR(100)"`          //邮箱
	UserUrl           string    `xorm:"not null default '' VARCHAR(100)"`                //网址
	UserRegistered    time.Time `xorm:"not null default '0000-00-00 00:00:00' DATETIME"` //注册时间
	UserActivationKey string    `xorm:"not null default '' VARCHAR(255)"`                //激活码
	UserStatus        int       `xorm:"not null default 0 INT(11)"`                      //用户状态
	DisplayName       string    `xorm:"not null default '' VARCHAR(250)"`                //显示名称
}
