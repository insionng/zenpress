package model

type Site struct {
	ID     uint64 `gorm:"primary_key"`
	Domain string `gorm:"not null default '' index(domain) VARCHAR(200)"`
	Path   string `gorm:"not null default '' index(domain) VARCHAR(100)"`
}
