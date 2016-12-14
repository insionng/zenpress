package model

import (
	"time"
)

// Link 存储友情链接信息，如友链名称、URL、打开方式、描述、是否可见等；
type Link struct {
	LinkId          int64     `xorm:"not null pk autoincr BIGINT(20)"`                 //自增唯一ID
	LinkUrl         string    `xorm:"not null default '' VARCHAR(255)"`                //链接URL
	LinkName        string    `xorm:"not null default '' VARCHAR(255)"`                //链接标题
	LinkImage       string    `xorm:"not null default '' VARCHAR(255)"`                //链接图片
	LinkTarget      string    `xorm:"not null default '' VARCHAR(25)"`                 //链接打开方式
	LinkDescription string    `xorm:"not null default '' VARCHAR(255)"`                //链接描述
	LinkVisible     string    `xorm:"not null default 'Y' index VARCHAR(20)"`          //是否可见（Y/N）
	LinkOwner       int64     `xorm:"not null default 1 BIGINT(20)"`                   //添加者用户ID
	LinkRating      int       `xorm:"not null default 0 INT(11)"`                      //评分等级
	LinkUpdated     time.Time `xorm:"not null default '0000-00-00 00:00:00' DATETIME"` //链接更新时间
	LinkRel         string    `xorm:"not null default '' VARCHAR(255)"`                //XFN关系
	LinkNotes       string    `xorm:"not null MEDIUMTEXT"`                             //XFN注释
	LinkRss         string    `xorm:"not null default '' VARCHAR(255)"`                //链接RSS地址
}
