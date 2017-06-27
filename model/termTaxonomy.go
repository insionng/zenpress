package model

// TermTaxonomy 分类信息表，区分wp_terms信息的分类类型，有category、link_category和tag三种分类类型。
type TermTaxonomy struct {
	TermTaxonomyID uint64 `gorm:"primary_key"`
	TermID         uint64 `gorm:"not null default 0 unique(term_id_taxonomy) BIGINT(20)"`
	Taxonomy       string `gorm:"not null default '' unique(term_id_taxonomy) index VARCHAR(32)"`
	Description    string `gorm:"not null LONGTEXT"`
	Parent         int64  `gorm:"not null default 0 BIGINT(20)"`
	Count          int64  `gorm:"not null default 0 BIGINT(20)"`
}
