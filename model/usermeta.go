package model

// Usermeta 用户额外信息表
type Usermeta struct {
	UmetaID   uint64 `gorm:"primary_key"`
	UserID    uint64 `gorm:"not null default 0 index BIGINT(20)"`
	MetaKey   string `gorm:"index VARCHAR(255)"`
	MetaValue string `gorm:"LONGTEXT"`
}
