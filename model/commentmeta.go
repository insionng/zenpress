package model

import "github.com/jinzhu/gorm"

// Commentmeta 文章评论额外信息表
type Commentmeta struct {
	ID        uint64 `gorm:"primary_key"`
	CommentID uint64 `xorm:"not null default 0 index BIGINT(20)"`
	MetaKey   string `xorm:"index VARCHAR(255)"`
	MetaValue string `xorm:"LONGTEXT"`
}

// NewCommentmeta 创建评论数据
func NewCommentmeta(commentmeta *Commentmeta) (db *gorm.DB) {
	db = Database.Create(commentmeta)
	return
}

// AddCommentmeta 新增评论数据
func AddCommentmeta(commentID uint64, metaKey, metaValue string) (db *gorm.DB) {
	db = Database.Create(&Commentmeta{CommentID: commentID, MetaKey: metaKey, MetaValue: metaValue})
	return
}

// GetCommentmeta 获得评论数据
func GetCommentmeta(id uint64) (db *gorm.DB, Commentmeta Commentmeta) {
	db = Database.First(&Commentmeta, "id = ?", id)
	return
}

// UpdateCommentmeta 更新评论数据
func UpdateCommentmeta(id, commentID uint64, metaKey, metaValue string) (db *gorm.DB, commentmeta Commentmeta) {
	commentmeta = Commentmeta{CommentID: commentID, MetaKey: metaKey, MetaValue: metaValue}
	db = Database.Model(&commentmeta).Update("id", id)
	return
}

// DeleteCommentmeta 删除评论数据
func DeleteCommentmeta(id uint64) (db *gorm.DB) {
	db = Database.Delete(Commentmeta{}, "id = ?", id)
	return
}
