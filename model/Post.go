package model

import (
	"time"
)

// Post 存储文章信息，如文章标题、正文、摘要、作者、发布时间、
// 访问密码、评论数、修改时间、文章地址（非静态化之前的，带？和数字ID）等；
type Post struct {
	Id                  int64     `xorm:"index(type_status_date) BIGINT(20)"`                                      //自增唯一ID
	PostAuthor          int64     `xorm:"not null default 0 index BIGINT(20)"`                                     //对应作者ID
	PostDate            time.Time `xorm:"not null default '0000-00-00 00:00:00' index(type_status_date) DATETIME"` //发布时间
	PostDateGmt         time.Time `xorm:"not null default '0000-00-00 00:00:00' DATETIME"`                         //发布时间（GMT+0时间）
	PostContent         string    `xorm:"not null LONGTEXT"`                                                       //正文
	PostTitle           string    `xorm:"not null TEXT"`                                                           //标题
	PostExcerpt         string    `xorm:"not null TEXT"`                                                           //摘录
	PostStatus          string    `xorm:"not null default 'publish' index(type_status_date) VARCHAR(20)"`          //文章状态（publish/auto-draft/inherit等）
	CommentStatus       string    `xorm:"not null default 'open' VARCHAR(20)"`                                     //PING状态（open/closed）                                  //评论状态（open/closed）
	PingStatus          string    `xorm:"not null default 'open' VARCHAR(20)"`
	PostPassword        string    `xorm:"not null default '' VARCHAR(255)"`                            //文章密码
	PostName            string    `xorm:"not null default '' index VARCHAR(200)"`                      //文章缩略名
	ToPing              string    `xorm:"not null TEXT"`                                               //未知
	Pinged              string    `xorm:"not null TEXT"`                                               //已经PING过的链接
	PostModified        time.Time `xorm:"not null default '0000-00-00 00:00:00' DATETIME"`             //修改时间
	PostModifiedGmt     time.Time `xorm:"not null default '0000-00-00 00:00:00' DATETIME"`             //修改时间（GMT+0时间）
	PostContentFiltered string    `xorm:"not null LONGTEXT"`                                           //未知
	PostParent          int64     `xorm:"not null default 0 index BIGINT(20)"`                         //父文章，主要用于PAGE
	Guid                string    `xorm:"not null default '' VARCHAR(255)"`                            //未知
	MenuOrder           int       `xorm:"not null default 0 INT(11)"`                                  //排序ID
	PostType            string    `xorm:"not null default 'post' index(type_status_date) VARCHAR(20)"` //文章类型（post/page等）
	PostMimeType        string    `xorm:"not null default '' VARCHAR(100)"`                            //MIME类型
	CommentCount        int64     `xorm:"not null default 0 BIGINT(20)"`                               //评论总数
}
