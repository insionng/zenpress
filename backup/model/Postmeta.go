package model

// Postmeta 存储文章的一些相关信息，如文章缩略图地址、缩略图长宽高和alt信息、
// 文章所在分类的URL、文章自定义的Description和Keywords、文章访问次数等；
type Postmeta struct {
	MetaId    int64  `xorm:"not null pk autoincr BIGINT(20)"`     //自增唯一ID
	PostId    int64  `xorm:"not null default 0 index BIGINT(20)"` //对应文章ID
	MetaKey   string `xorm:"index VARCHAR(255)"`                  //键名
	MetaValue string `xorm:"LONGTEXT"`                            //键值
}
