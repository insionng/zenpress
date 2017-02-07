package model

// Commentmeta 仅存储Akismet或手工审核的评论是否为垃圾评论的判断结果；
type Commentmeta struct {
	MetaId    int64  `xorm:"not null pk autoincr BIGINT(20)"`     //自增唯一ID
	CommentId int64  `xorm:"not null default 0 index BIGINT(20)"` //对应评论ID
	MetaKey   string `xorm:"index VARCHAR(255)"`                  //键名
	MetaValue string `xorm:"LONGTEXT"`                            //键值
}
