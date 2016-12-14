package model

import (
	"time"
)

// Comment 存储评论信息，如评论内容、评论所属文章、评论人昵称、邮箱、URL等；
type Comment struct {
	CommentId          int64     `xorm:"not null pk autoincr BIGINT(20)"`                                                        //自增唯一ID
	CommentPostId      int64     `xorm:"not null default 0 index BIGINT(20)"`                                                    //对应文章ID
	CommentAuthor      string    `xorm:"not null TINYTEXT"`                                                                      //评论者
	CommentAuthorEmail string    `xorm:"not null default '' index VARCHAR(100)"`                                                 //评论者邮箱
	CommentAuthorUrl   string    `xorm:"not null default '' VARCHAR(200)"`                                                       //评论者网址
	CommentAuthorIp    string    `xorm:"not null default '' VARCHAR(100)"`                                                       //评论者IP
	CommentDate        time.Time `xorm:"not null default '0000-00-00 00:00:00' DATETIME"`                                        //评论时间
	CommentDateGmt     time.Time `xorm:"not null default '0000-00-00 00:00:00' index(comment_approved_date_gmt) index DATETIME"` //评论时间（GMT+0时间）
	CommentContent     string    `xorm:"not null TEXT"`                                                                          //评论正文
	CommentKarma       int       `xorm:"not null default 0 INT(11)"`                                                             //未知
	CommentApproved    string    `xorm:"not null default '1' index(comment_approved_date_gmt) VARCHAR(20)"`                      //评论是否被批准
	CommentAgent       string    `xorm:"not null default '' VARCHAR(255)"`                                                       //评论者的USER AGENT
	CommentType        string    `xorm:"not null default '' VARCHAR(20)"`                                                        //评论类型(pingback/普通)
	CommentParent      int64     `xorm:"not null default 0 index BIGINT(20)"`                                                    //父评论ID
	UserId             int64     `xorm:"not null default 0 BIGINT(20)"`                                                          //评论者用户ID（不一定存在）
}
