package model

import "github.com/jinzhu/gorm"

// Term 文章分类、链接分类、标签的信息表。
type Term struct {
	ID        uint64 `gorm:"primary_key"`
	Name      string `gorm:"not null default '' index VARCHAR(200)"`
	Slug      string `gorm:"not null default '' index VARCHAR(200)"`
	TermGroup uint64 `gorm:"not null default 0 BIGINT(10)"`
}

// NewTerm 创建分类
func NewTerm(term *Term) (db *gorm.DB) {
	db = Database.Create(term)
	return
}

// AddTerm 新增分类
func AddTerm(name, slug string, termGroup uint64) (db *gorm.DB) {
	db = Database.Create(&Term{Name: name, Slug: slug, TermGroup: termGroup})
	return
}

// GetTerm 获得分类
func GetTerm(id uint64) (db *gorm.DB, Term Term) {
	db = Database.First(&Term, "id = ?", id)
	return
}

// UpdateTerm 更新分类
func UpdateTerm(id uint64, name string) (db *gorm.DB, term Term) {
	term = Term{Name: name}
	db = Database.Model(&term).Update("id", id)
	return
}

// DeleteTerm 删除分类
func DeleteTerm(id uint64) (db *gorm.DB) {
	db = Database.Delete(Term{}, "id = ?", id)
	return
}
