package model

// Usermeta 存储分类和标签的描述信息、父子关系、所属包含的文章数等；
type Usermeta struct {
	UmetaId   int64  `xorm:"not null pk autoincr BIGINT(20)"`     //自增唯一ID
	UserId    int64  `xorm:"not null default 0 index BIGINT(20)"` //对应用户ID
	MetaKey   string `xorm:"index VARCHAR(255)"`                  //键名
	MetaValue string `xorm:"LONGTEXT"`                            //键值
}
