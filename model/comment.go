package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Comment 评论评论信息表
type Comment struct {
	ID                 uint64    `gorm:"primary_key"`
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

// NewComment 创建评论
func NewComment(comment *Comment) (db *gorm.DB) {
	db = Database.Create(comment)
	return
}

// AddComment 新增评论
func AddComment(commentAuthor, commentAuthorURL, commentContent string) (db *gorm.DB) {
	db = Database.Create(&Comment{CommentAuthor: commentAuthor, CommentContent: commentContent, CommentAuthorURL: commentAuthorURL})
	return
}

// GetComment 获得评论
func GetComment(id uint64) (db *gorm.DB, Comment Comment) {
	db = Database.First(&Comment, "id = ?", id)
	return
}

// UpdateComment 更新评论
func UpdateComment(id uint64, commentContent string) (db *gorm.DB, comment Comment) {
	comment = Comment{CommentContent: commentContent}
	db = Database.Model(&comment).Update("id", id)
	return
}

// DeleteComment 删除评论
func DeleteComment(id uint64) (db *gorm.DB) {
	db = Database.Delete(Comment{}, "id = ?", id)
	return
}
