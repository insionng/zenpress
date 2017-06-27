package model

import (
	"time"
)

// Comment 文章评论信息表
type Comment struct {
	CommentID          uint64    `gorm:"primary_key"`
	CommentPostID      int64     `gorm:"not null default 0 index BIGINT(20)"`
	CommentAuthor      string    `gorm:"not null TINYTEXT"`
	CommentAuthorEmail string    `gorm:"not null default '' index VARCHAR(100)"`
	CommentAuthorURL   string    `gorm:"not null default '' VARCHAR(200)"`
	CommentAuthorIP    string    `gorm:"not null default '' VARCHAR(100)"`
	CommentDate        time.Time `gorm:"not null default '0000-00-00 00:00:00' DATETIME"`
	CommentDateGmt     time.Time `gorm:"not null default '0000-00-00 00:00:00' index index(comment_approved_date_gmt) DATETIME"`
	CommentContent     string    `gorm:"not null TEXT"`
	CommentKarma       int       `gorm:"not null default 0 INT(11)"`
	CommentApproved    string    `gorm:"not null default '1' index(comment_approved_date_gmt) VARCHAR(20)"`
	CommentAgent       string    `gorm:"not null default '' VARCHAR(255)"`
	CommentType        string    `gorm:"not null default '' VARCHAR(20)"`
	CommentParent      int64     `gorm:"not null default 0 index BIGINT(20)"`
	UserID             int64     `gorm:"not null default 0 BIGINT(20)"`
}
