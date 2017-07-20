package model

import "github.com/jinzhu/gorm"

// Usermeta 用户额外信息表
type Usermeta struct {
	ID        uint64 `gorm:"primary_key"`
	UserID    uint64 `gorm:"not null default 0 index BIGINT(20)"`
	MetaKey   string `gorm:"index VARCHAR(255)"`
	MetaValue string `gorm:"LONGTEXT"`
}

// NewUsermeta 创建用户数据
func NewUsermeta(usermeta *Usermeta) (db *gorm.DB) {
	db = Database.Create(usermeta)
	return
}

// AddUsermeta 新增用户数据
func AddUsermeta(userID uint64, metaKey, metaValue string) (db *gorm.DB) {
	db = Database.Create(&Usermeta{UserID: userID, MetaKey: metaKey, MetaValue: metaValue})
	return
}

// GetUsermeta 获得用户数据
func GetUsermeta(id uint64) (db *gorm.DB, Usermeta Usermeta) {
	db = Database.First(&Usermeta, "id = ?", id)
	return
}

// UpdateUsermeta 更新用户数据
func UpdateUsermeta(id, userID uint64, metaKey, metaValue string) (db *gorm.DB, usermeta Usermeta) {
	usermeta = Usermeta{UserID: userID, MetaKey: metaKey, MetaValue: metaValue}
	db = Database.Model(&usermeta).Update("id", id)
	return
}

// DeleteUsermeta 删除用户数据
func DeleteUsermeta(id uint64) (db *gorm.DB) {
	db = Database.Delete(Usermeta{}, "id = ?", id)
	return
}
