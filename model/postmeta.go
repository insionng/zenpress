package model

import "github.com/jinzhu/gorm"

// Postmeta 文章额外数据表，例如文章浏览次数，文章的自定义字段等都存储在这里。
type Postmeta struct {
	ID        uint64 `gorm:"primary_key"`
	PostID    uint64 `gorm:"not null default 0 index BIGINT(20)"`
	MetaKey   string `gorm:"index VARCHAR(255)"`
	MetaValue string `gorm:"LONGTEXT"`
}

// NewPostmeta 创建文章数据
func NewPostmeta(postmeta *Postmeta) (db *gorm.DB) {
	db = Database.Create(postmeta)
	return
}

// AddPostmeta 新增文章数据
func AddPostmeta(postID uint64, metaKey, metaValue string) (db *gorm.DB) {
	db = Database.Create(&Postmeta{PostID: postID, MetaKey: metaKey, MetaValue: metaValue})
	return
}

// GetPostmeta 获得文章数据
func GetPostmeta(id uint64) (db *gorm.DB, Postmeta Postmeta) {
	db = Database.First(&Postmeta, "id = ?", id)
	return
}

// UpdatePostmeta 更新文章数据
func UpdatePostmeta(id, postID uint64, metaKey, metaValue string) (db *gorm.DB, postmeta Postmeta) {
	postmeta = Postmeta{PostID: postID, MetaKey: metaKey, MetaValue: metaValue}
	db = Database.Model(&postmeta).Update("id", id)
	return
}

// DeletePostmeta 删除文章数据
func DeletePostmeta(id uint64) (db *gorm.DB) {
	db = Database.Delete(Postmeta{}, "id = ?", id)
	return
}
