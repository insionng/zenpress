package model

// TermRelationship 存储文章和分类、标签的相互对应关系；
type TermRelationship struct {
	ObjectId       int64 `xorm:"not null pk default 0 BIGINT(20)"`       //对应文章ID/链接ID
	TermTaxonomyId int64 `xorm:"not null pk default 0 index BIGINT(20)"` //对应分类方法ID
	TermOrder      int   `xorm:"not null default 0 INT(11)"`             //排序
}
