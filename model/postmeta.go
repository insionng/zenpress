package model

// Postmeta 文章额外数据表，例如文章浏览次数，文章的自定义字段等都存储在这里。
type Postmeta struct {
	MetaID    uint64 `gorm:"primary_key"`
	PostID    int64  `gorm:"not null default 0 index BIGINT(20)"`
	MetaKey   string `gorm:"index VARCHAR(255)"`
	MetaValue string `gorm:"LONGTEXT"`
}
