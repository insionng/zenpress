package model

type TermTaxonomy struct {
	TermTaxonomyId int64  `xorm:"not null pk autoincr BIGINT(20)"`                                //分类方法ID
	TermId         int64  `xorm:"not null default 0 unique(term_id_taxonomy) BIGINT(20)"`         //分类ID
	Taxonomy       string `xorm:"not null default '' unique(term_id_taxonomy) index VARCHAR(32)"` //分类方法(category/post_tag)
	Description    string `xorm:"not null LONGTEXT"`                                              //未知
	Parent         int64  `xorm:"not null default 0 BIGINT(20)"`                                  //所属父分类方法ID
	Count          int64  `xorm:"not null default 0 BIGINT(20)"`                                  //文章数统计
}
