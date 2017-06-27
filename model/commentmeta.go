package model

// Commentmeta 文章评论额外信息表
type Commentmeta struct {
	MetaID    uint64 `gorm:"primary_key"`
	CommentID uint64 `xorm:"not null default 0 index BIGINT(20)"`
	MetaKey   string `xorm:"index VARCHAR(255)"`
	MetaValue string `xorm:"LONGTEXT"`
}
