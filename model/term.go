package model

// Term 文章分类、链接分类、标签的信息表。
type Term struct {
	TermID    uint64 `gorm:"primary_key"`
	Name      string `gorm:"not null default '' index VARCHAR(200)"`
	Slug      string `gorm:"not null default '' index VARCHAR(200)"`
	TermGroup int64  `gorm:"not null default 0 BIGINT(10)"`
}
