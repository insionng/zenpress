package model

type Termmeta struct {
	MetaId    int64  `xorm:"not null pk autoincr BIGINT(20)"`     //自增唯一ID
	TermId    int64  `xorm:"not null default 0 index BIGINT(20)"` //分类ID
	MetaKey   string `xorm:"index VARCHAR(255)"`                  //键名
	MetaValue string `xorm:"LONGTEXT"`                            //键值
}
