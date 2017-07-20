package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Post 文章信息表，包括了日志、附件、页面等等信息。是WordPress最重要的一个数据表。
type Post struct {
	ID                  uint64    `gorm:"primary_key"`
	PostAuthor          uint64    `gorm:"index"`
	PostDate            time.Time `gorm:"not null default '0000-00-00 00:00:00' index(TYPE_STATUS_DATE) DATETIME"`
	PostDateGmt         time.Time `gorm:"not null default '0000-00-00 00:00:00' DATETIME"`
	PostContent         string    `gorm:"not null LONGTEXT"`
	PostTitle           string    `gorm:"not null"`
	PostExcerpt         string    `gorm:"not null"`
	PostStatus          string    `gorm:"not null default 'publish' index(TYPE_STATUS_DATE) VARCHAR(20)"`
	CommentStatus       string    `gorm:"not null default 'open' VARCHAR(20)"`
	PingStatus          string    `gorm:"not null default 'open' VARCHAR(20)"`
	PostPassword        string    `gorm:"not null default '' VARCHAR(255)"`
	PostName            string    `gorm:"not null default '' index VARCHAR(200)"`
	ToPing              string    `gorm:"not null"`
	Pinged              string    `gorm:"not null"`
	PostModified        time.Time `gorm:"not null default '0000-00-00 00:00:00' DATETIME"`
	PostModifiedGmt     time.Time `gorm:"not null default '0000-00-00 00:00:00' DATETIME"`
	PostContentFiltered string    `gorm:"not null LONGTEXT"`
	PostParent          uint64    `gorm:"index"`
	GUID                string
	MenuOrder           int
	PostType            string `gorm:"not null default 'post' index(TYPE_STATUS_DATE) VARCHAR(20)"`
	PostMimeType        string
	CommentCount        uint64
}

// NewPost 创建文章
func NewPost(post *Post) (db *gorm.DB) {
	db = Database.Create(post)
	return
}

// AddPost 新增文章
func AddPost(postName, postTitle, postContent string) (db *gorm.DB) {
	db = Database.Create(&Post{PostName: postName, PostContent: postContent, PostTitle: postTitle})
	return
}

// GetPost 获得文章
func GetPost(id uint64) (db *gorm.DB, Post Post) {
	db = Database.First(&Post, "id = ?", id)
	return
}

// UpdatePost 更新文章
func UpdatePost(id uint64, postContent string) (db *gorm.DB, post Post) {
	post = Post{PostContent: postContent}
	db = Database.Model(&post).Update("id", id)
	return
}

// DeletePost 删除文章
func DeletePost(id uint64) (db *gorm.DB) {
	db = Database.Delete(Post{}, "id = ?", id)
	return
}
