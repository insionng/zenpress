package model

type TermTaxonomy struct {
	TermTaxonomyId int64  `xorm:"not null pk autoincr BIGINT(20)"`
	TermId         int64  `xorm:"not null default 0 unique(term_id_taxonomy) BIGINT(20)"`
	Taxonomy       string `xorm:"not null default '' unique(term_id_taxonomy) index VARCHAR(32)"`
	Description    string `xorm:"not null LONGTEXT"`
	Parent         int64  `xorm:"not null default 0 BIGINT(20)"`
	Count          int64  `xorm:"not null default 0 BIGINT(20)"`
}

// GetTermTaxonomiesCount TermTaxonomys' Count
func GetTermTaxonomiesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&TermTaxonomy{})
	return total, err
}

// GetTermTaxonomyCountByTermTaxonomyId Get TermTaxonomy via TermTaxonomyId
func GetTermTaxonomyCountByTermTaxonomyId(iTermTaxonomyId int64) int64 {
	n, _ := Engine.Where("term_taxonomy_id = ?", iTermTaxonomyId).Count(&TermTaxonomy{TermTaxonomyId: iTermTaxonomyId})
	return n
}

// GetTermTaxonomyCountByTermId Get TermTaxonomy via TermId
func GetTermTaxonomyCountByTermId(iTermId int64) int64 {
	n, _ := Engine.Where("term_id = ?", iTermId).Count(&TermTaxonomy{TermId: iTermId})
	return n
}

// GetTermTaxonomyCountByTaxonomy Get TermTaxonomy via Taxonomy
func GetTermTaxonomyCountByTaxonomy(iTaxonomy string) int64 {
	n, _ := Engine.Where("taxonomy = ?", iTaxonomy).Count(&TermTaxonomy{Taxonomy: iTaxonomy})
	return n
}

// GetTermTaxonomyCountByDescription Get TermTaxonomy via Description
func GetTermTaxonomyCountByDescription(iDescription string) int64 {
	n, _ := Engine.Where("description = ?", iDescription).Count(&TermTaxonomy{Description: iDescription})
	return n
}

// GetTermTaxonomyCountByParent Get TermTaxonomy via Parent
func GetTermTaxonomyCountByParent(iParent int64) int64 {
	n, _ := Engine.Where("parent = ?", iParent).Count(&TermTaxonomy{Parent: iParent})
	return n
}

// GetTermTaxonomyCountByCount Get TermTaxonomy via Count
func GetTermTaxonomyCountByCount(iCount int64) int64 {
	n, _ := Engine.Where("count = ?", iCount).Count(&TermTaxonomy{Count: iCount})
	return n
}

// GetTermTaxonomiesByTermTaxonomyIdAndTermIdAndTaxonomy Get TermTaxonomies via TermTaxonomyIdAndTermIdAndTaxonomy
func GetTermTaxonomiesByTermTaxonomyIdAndTermIdAndTaxonomy(offset int, limit int, TermTaxonomyId_ int64, TermId_ int64, Taxonomy_ string) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("term_taxonomy_id = ? and term_id = ? and taxonomy = ?", TermTaxonomyId_, TermId_, Taxonomy_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTermTaxonomyIdAndTermIdAndDescription Get TermTaxonomies via TermTaxonomyIdAndTermIdAndDescription
func GetTermTaxonomiesByTermTaxonomyIdAndTermIdAndDescription(offset int, limit int, TermTaxonomyId_ int64, TermId_ int64, Description_ string) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("term_taxonomy_id = ? and term_id = ? and description = ?", TermTaxonomyId_, TermId_, Description_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTermTaxonomyIdAndTermIdAndParent Get TermTaxonomies via TermTaxonomyIdAndTermIdAndParent
func GetTermTaxonomiesByTermTaxonomyIdAndTermIdAndParent(offset int, limit int, TermTaxonomyId_ int64, TermId_ int64, Parent_ int64) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("term_taxonomy_id = ? and term_id = ? and parent = ?", TermTaxonomyId_, TermId_, Parent_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTermTaxonomyIdAndTermIdAndCount Get TermTaxonomies via TermTaxonomyIdAndTermIdAndCount
func GetTermTaxonomiesByTermTaxonomyIdAndTermIdAndCount(offset int, limit int, TermTaxonomyId_ int64, TermId_ int64, Count_ int64) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("term_taxonomy_id = ? and term_id = ? and count = ?", TermTaxonomyId_, TermId_, Count_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTermTaxonomyIdAndTaxonomyAndDescription Get TermTaxonomies via TermTaxonomyIdAndTaxonomyAndDescription
func GetTermTaxonomiesByTermTaxonomyIdAndTaxonomyAndDescription(offset int, limit int, TermTaxonomyId_ int64, Taxonomy_ string, Description_ string) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("term_taxonomy_id = ? and taxonomy = ? and description = ?", TermTaxonomyId_, Taxonomy_, Description_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTermTaxonomyIdAndTaxonomyAndParent Get TermTaxonomies via TermTaxonomyIdAndTaxonomyAndParent
func GetTermTaxonomiesByTermTaxonomyIdAndTaxonomyAndParent(offset int, limit int, TermTaxonomyId_ int64, Taxonomy_ string, Parent_ int64) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("term_taxonomy_id = ? and taxonomy = ? and parent = ?", TermTaxonomyId_, Taxonomy_, Parent_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTermTaxonomyIdAndTaxonomyAndCount Get TermTaxonomies via TermTaxonomyIdAndTaxonomyAndCount
func GetTermTaxonomiesByTermTaxonomyIdAndTaxonomyAndCount(offset int, limit int, TermTaxonomyId_ int64, Taxonomy_ string, Count_ int64) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("term_taxonomy_id = ? and taxonomy = ? and count = ?", TermTaxonomyId_, Taxonomy_, Count_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTermTaxonomyIdAndDescriptionAndParent Get TermTaxonomies via TermTaxonomyIdAndDescriptionAndParent
func GetTermTaxonomiesByTermTaxonomyIdAndDescriptionAndParent(offset int, limit int, TermTaxonomyId_ int64, Description_ string, Parent_ int64) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("term_taxonomy_id = ? and description = ? and parent = ?", TermTaxonomyId_, Description_, Parent_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTermTaxonomyIdAndDescriptionAndCount Get TermTaxonomies via TermTaxonomyIdAndDescriptionAndCount
func GetTermTaxonomiesByTermTaxonomyIdAndDescriptionAndCount(offset int, limit int, TermTaxonomyId_ int64, Description_ string, Count_ int64) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("term_taxonomy_id = ? and description = ? and count = ?", TermTaxonomyId_, Description_, Count_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTermTaxonomyIdAndParentAndCount Get TermTaxonomies via TermTaxonomyIdAndParentAndCount
func GetTermTaxonomiesByTermTaxonomyIdAndParentAndCount(offset int, limit int, TermTaxonomyId_ int64, Parent_ int64, Count_ int64) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("term_taxonomy_id = ? and parent = ? and count = ?", TermTaxonomyId_, Parent_, Count_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTermIdAndTaxonomyAndDescription Get TermTaxonomies via TermIdAndTaxonomyAndDescription
func GetTermTaxonomiesByTermIdAndTaxonomyAndDescription(offset int, limit int, TermId_ int64, Taxonomy_ string, Description_ string) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("term_id = ? and taxonomy = ? and description = ?", TermId_, Taxonomy_, Description_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTermIdAndTaxonomyAndParent Get TermTaxonomies via TermIdAndTaxonomyAndParent
func GetTermTaxonomiesByTermIdAndTaxonomyAndParent(offset int, limit int, TermId_ int64, Taxonomy_ string, Parent_ int64) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("term_id = ? and taxonomy = ? and parent = ?", TermId_, Taxonomy_, Parent_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTermIdAndTaxonomyAndCount Get TermTaxonomies via TermIdAndTaxonomyAndCount
func GetTermTaxonomiesByTermIdAndTaxonomyAndCount(offset int, limit int, TermId_ int64, Taxonomy_ string, Count_ int64) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("term_id = ? and taxonomy = ? and count = ?", TermId_, Taxonomy_, Count_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTermIdAndDescriptionAndParent Get TermTaxonomies via TermIdAndDescriptionAndParent
func GetTermTaxonomiesByTermIdAndDescriptionAndParent(offset int, limit int, TermId_ int64, Description_ string, Parent_ int64) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("term_id = ? and description = ? and parent = ?", TermId_, Description_, Parent_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTermIdAndDescriptionAndCount Get TermTaxonomies via TermIdAndDescriptionAndCount
func GetTermTaxonomiesByTermIdAndDescriptionAndCount(offset int, limit int, TermId_ int64, Description_ string, Count_ int64) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("term_id = ? and description = ? and count = ?", TermId_, Description_, Count_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTermIdAndParentAndCount Get TermTaxonomies via TermIdAndParentAndCount
func GetTermTaxonomiesByTermIdAndParentAndCount(offset int, limit int, TermId_ int64, Parent_ int64, Count_ int64) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("term_id = ? and parent = ? and count = ?", TermId_, Parent_, Count_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTaxonomyAndDescriptionAndParent Get TermTaxonomies via TaxonomyAndDescriptionAndParent
func GetTermTaxonomiesByTaxonomyAndDescriptionAndParent(offset int, limit int, Taxonomy_ string, Description_ string, Parent_ int64) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("taxonomy = ? and description = ? and parent = ?", Taxonomy_, Description_, Parent_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTaxonomyAndDescriptionAndCount Get TermTaxonomies via TaxonomyAndDescriptionAndCount
func GetTermTaxonomiesByTaxonomyAndDescriptionAndCount(offset int, limit int, Taxonomy_ string, Description_ string, Count_ int64) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("taxonomy = ? and description = ? and count = ?", Taxonomy_, Description_, Count_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTaxonomyAndParentAndCount Get TermTaxonomies via TaxonomyAndParentAndCount
func GetTermTaxonomiesByTaxonomyAndParentAndCount(offset int, limit int, Taxonomy_ string, Parent_ int64, Count_ int64) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("taxonomy = ? and parent = ? and count = ?", Taxonomy_, Parent_, Count_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByDescriptionAndParentAndCount Get TermTaxonomies via DescriptionAndParentAndCount
func GetTermTaxonomiesByDescriptionAndParentAndCount(offset int, limit int, Description_ string, Parent_ int64, Count_ int64) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("description = ? and parent = ? and count = ?", Description_, Parent_, Count_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTermTaxonomyIdAndTermId Get TermTaxonomies via TermTaxonomyIdAndTermId
func GetTermTaxonomiesByTermTaxonomyIdAndTermId(offset int, limit int, TermTaxonomyId_ int64, TermId_ int64) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("term_taxonomy_id = ? and term_id = ?", TermTaxonomyId_, TermId_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTermTaxonomyIdAndTaxonomy Get TermTaxonomies via TermTaxonomyIdAndTaxonomy
func GetTermTaxonomiesByTermTaxonomyIdAndTaxonomy(offset int, limit int, TermTaxonomyId_ int64, Taxonomy_ string) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("term_taxonomy_id = ? and taxonomy = ?", TermTaxonomyId_, Taxonomy_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTermTaxonomyIdAndDescription Get TermTaxonomies via TermTaxonomyIdAndDescription
func GetTermTaxonomiesByTermTaxonomyIdAndDescription(offset int, limit int, TermTaxonomyId_ int64, Description_ string) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("term_taxonomy_id = ? and description = ?", TermTaxonomyId_, Description_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTermTaxonomyIdAndParent Get TermTaxonomies via TermTaxonomyIdAndParent
func GetTermTaxonomiesByTermTaxonomyIdAndParent(offset int, limit int, TermTaxonomyId_ int64, Parent_ int64) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("term_taxonomy_id = ? and parent = ?", TermTaxonomyId_, Parent_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTermTaxonomyIdAndCount Get TermTaxonomies via TermTaxonomyIdAndCount
func GetTermTaxonomiesByTermTaxonomyIdAndCount(offset int, limit int, TermTaxonomyId_ int64, Count_ int64) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("term_taxonomy_id = ? and count = ?", TermTaxonomyId_, Count_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTermIdAndTaxonomy Get TermTaxonomies via TermIdAndTaxonomy
func GetTermTaxonomiesByTermIdAndTaxonomy(offset int, limit int, TermId_ int64, Taxonomy_ string) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("term_id = ? and taxonomy = ?", TermId_, Taxonomy_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTermIdAndDescription Get TermTaxonomies via TermIdAndDescription
func GetTermTaxonomiesByTermIdAndDescription(offset int, limit int, TermId_ int64, Description_ string) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("term_id = ? and description = ?", TermId_, Description_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTermIdAndParent Get TermTaxonomies via TermIdAndParent
func GetTermTaxonomiesByTermIdAndParent(offset int, limit int, TermId_ int64, Parent_ int64) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("term_id = ? and parent = ?", TermId_, Parent_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTermIdAndCount Get TermTaxonomies via TermIdAndCount
func GetTermTaxonomiesByTermIdAndCount(offset int, limit int, TermId_ int64, Count_ int64) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("term_id = ? and count = ?", TermId_, Count_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTaxonomyAndDescription Get TermTaxonomies via TaxonomyAndDescription
func GetTermTaxonomiesByTaxonomyAndDescription(offset int, limit int, Taxonomy_ string, Description_ string) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("taxonomy = ? and description = ?", Taxonomy_, Description_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTaxonomyAndParent Get TermTaxonomies via TaxonomyAndParent
func GetTermTaxonomiesByTaxonomyAndParent(offset int, limit int, Taxonomy_ string, Parent_ int64) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("taxonomy = ? and parent = ?", Taxonomy_, Parent_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTaxonomyAndCount Get TermTaxonomies via TaxonomyAndCount
func GetTermTaxonomiesByTaxonomyAndCount(offset int, limit int, Taxonomy_ string, Count_ int64) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("taxonomy = ? and count = ?", Taxonomy_, Count_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByDescriptionAndParent Get TermTaxonomies via DescriptionAndParent
func GetTermTaxonomiesByDescriptionAndParent(offset int, limit int, Description_ string, Parent_ int64) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("description = ? and parent = ?", Description_, Parent_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByDescriptionAndCount Get TermTaxonomies via DescriptionAndCount
func GetTermTaxonomiesByDescriptionAndCount(offset int, limit int, Description_ string, Count_ int64) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("description = ? and count = ?", Description_, Count_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByParentAndCount Get TermTaxonomies via ParentAndCount
func GetTermTaxonomiesByParentAndCount(offset int, limit int, Parent_ int64, Count_ int64) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("parent = ? and count = ?", Parent_, Count_).Limit(limit, offset).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomies Get TermTaxonomies via field
func GetTermTaxonomies(offset int, limit int, field string) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Limit(limit, offset).Desc(field).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTermTaxonomyId Get TermTaxonomys via TermTaxonomyId
func GetTermTaxonomiesByTermTaxonomyId(offset int, limit int, TermTaxonomyId_ int64, field string) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("term_taxonomy_id = ?", TermTaxonomyId_).Limit(limit, offset).Desc(field).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTermId Get TermTaxonomys via TermId
func GetTermTaxonomiesByTermId(offset int, limit int, TermId_ int64, field string) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("term_id = ?", TermId_).Limit(limit, offset).Desc(field).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByTaxonomy Get TermTaxonomys via Taxonomy
func GetTermTaxonomiesByTaxonomy(offset int, limit int, Taxonomy_ string, field string) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("taxonomy = ?", Taxonomy_).Limit(limit, offset).Desc(field).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByDescription Get TermTaxonomys via Description
func GetTermTaxonomiesByDescription(offset int, limit int, Description_ string, field string) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("description = ?", Description_).Limit(limit, offset).Desc(field).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByParent Get TermTaxonomys via Parent
func GetTermTaxonomiesByParent(offset int, limit int, Parent_ int64, field string) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("parent = ?", Parent_).Limit(limit, offset).Desc(field).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// GetTermTaxonomiesByCount Get TermTaxonomys via Count
func GetTermTaxonomiesByCount(offset int, limit int, Count_ int64, field string) (*[]*TermTaxonomy, error) {
	var _TermTaxonomy = new([]*TermTaxonomy)
	err := Engine.Table("term_taxonomy").Where("count = ?", Count_).Limit(limit, offset).Desc(field).Find(_TermTaxonomy)
	return _TermTaxonomy, err
}

// HasTermTaxonomyByTermTaxonomyId Has TermTaxonomy via TermTaxonomyId
func HasTermTaxonomyByTermTaxonomyId(iTermTaxonomyId int64) bool {
	if has, err := Engine.Where("term_taxonomy_id = ?", iTermTaxonomyId).Get(new(TermTaxonomy)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTermTaxonomyByTermId Has TermTaxonomy via TermId
func HasTermTaxonomyByTermId(iTermId int64) bool {
	if has, err := Engine.Where("term_id = ?", iTermId).Get(new(TermTaxonomy)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTermTaxonomyByTaxonomy Has TermTaxonomy via Taxonomy
func HasTermTaxonomyByTaxonomy(iTaxonomy string) bool {
	if has, err := Engine.Where("taxonomy = ?", iTaxonomy).Get(new(TermTaxonomy)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTermTaxonomyByDescription Has TermTaxonomy via Description
func HasTermTaxonomyByDescription(iDescription string) bool {
	if has, err := Engine.Where("description = ?", iDescription).Get(new(TermTaxonomy)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTermTaxonomyByParent Has TermTaxonomy via Parent
func HasTermTaxonomyByParent(iParent int64) bool {
	if has, err := Engine.Where("parent = ?", iParent).Get(new(TermTaxonomy)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTermTaxonomyByCount Has TermTaxonomy via Count
func HasTermTaxonomyByCount(iCount int64) bool {
	if has, err := Engine.Where("count = ?", iCount).Get(new(TermTaxonomy)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetTermTaxonomyByTermTaxonomyId Get TermTaxonomy via TermTaxonomyId
func GetTermTaxonomyByTermTaxonomyId(iTermTaxonomyId int64) (*TermTaxonomy, error) {
	var _TermTaxonomy = &TermTaxonomy{TermTaxonomyId: iTermTaxonomyId}
	has, err := Engine.Get(_TermTaxonomy)
	if has {
		return _TermTaxonomy, err
	} else {
		return nil, err
	}
}

// GetTermTaxonomyByTermId Get TermTaxonomy via TermId
func GetTermTaxonomyByTermId(iTermId int64) (*TermTaxonomy, error) {
	var _TermTaxonomy = &TermTaxonomy{TermId: iTermId}
	has, err := Engine.Get(_TermTaxonomy)
	if has {
		return _TermTaxonomy, err
	} else {
		return nil, err
	}
}

// GetTermTaxonomyByTaxonomy Get TermTaxonomy via Taxonomy
func GetTermTaxonomyByTaxonomy(iTaxonomy string) (*TermTaxonomy, error) {
	var _TermTaxonomy = &TermTaxonomy{Taxonomy: iTaxonomy}
	has, err := Engine.Get(_TermTaxonomy)
	if has {
		return _TermTaxonomy, err
	} else {
		return nil, err
	}
}

// GetTermTaxonomyByDescription Get TermTaxonomy via Description
func GetTermTaxonomyByDescription(iDescription string) (*TermTaxonomy, error) {
	var _TermTaxonomy = &TermTaxonomy{Description: iDescription}
	has, err := Engine.Get(_TermTaxonomy)
	if has {
		return _TermTaxonomy, err
	} else {
		return nil, err
	}
}

// GetTermTaxonomyByParent Get TermTaxonomy via Parent
func GetTermTaxonomyByParent(iParent int64) (*TermTaxonomy, error) {
	var _TermTaxonomy = &TermTaxonomy{Parent: iParent}
	has, err := Engine.Get(_TermTaxonomy)
	if has {
		return _TermTaxonomy, err
	} else {
		return nil, err
	}
}

// GetTermTaxonomyByCount Get TermTaxonomy via Count
func GetTermTaxonomyByCount(iCount int64) (*TermTaxonomy, error) {
	var _TermTaxonomy = &TermTaxonomy{Count: iCount}
	has, err := Engine.Get(_TermTaxonomy)
	if has {
		return _TermTaxonomy, err
	} else {
		return nil, err
	}
}

// SetTermTaxonomyByTermTaxonomyId Set TermTaxonomy via TermTaxonomyId
func SetTermTaxonomyByTermTaxonomyId(iTermTaxonomyId int64, term_taxonomy *TermTaxonomy) (int64, error) {
	term_taxonomy.TermTaxonomyId = iTermTaxonomyId
	return Engine.Insert(term_taxonomy)
}

// SetTermTaxonomyByTermId Set TermTaxonomy via TermId
func SetTermTaxonomyByTermId(iTermId int64, term_taxonomy *TermTaxonomy) (int64, error) {
	term_taxonomy.TermId = iTermId
	return Engine.Insert(term_taxonomy)
}

// SetTermTaxonomyByTaxonomy Set TermTaxonomy via Taxonomy
func SetTermTaxonomyByTaxonomy(iTaxonomy string, term_taxonomy *TermTaxonomy) (int64, error) {
	term_taxonomy.Taxonomy = iTaxonomy
	return Engine.Insert(term_taxonomy)
}

// SetTermTaxonomyByDescription Set TermTaxonomy via Description
func SetTermTaxonomyByDescription(iDescription string, term_taxonomy *TermTaxonomy) (int64, error) {
	term_taxonomy.Description = iDescription
	return Engine.Insert(term_taxonomy)
}

// SetTermTaxonomyByParent Set TermTaxonomy via Parent
func SetTermTaxonomyByParent(iParent int64, term_taxonomy *TermTaxonomy) (int64, error) {
	term_taxonomy.Parent = iParent
	return Engine.Insert(term_taxonomy)
}

// SetTermTaxonomyByCount Set TermTaxonomy via Count
func SetTermTaxonomyByCount(iCount int64, term_taxonomy *TermTaxonomy) (int64, error) {
	term_taxonomy.Count = iCount
	return Engine.Insert(term_taxonomy)
}

// AddTermTaxonomy Add TermTaxonomy via all columns
func AddTermTaxonomy(iTermTaxonomyId int64, iTermId int64, iTaxonomy string, iDescription string, iParent int64, iCount int64) error {
	_TermTaxonomy := &TermTaxonomy{TermTaxonomyId: iTermTaxonomyId, TermId: iTermId, Taxonomy: iTaxonomy, Description: iDescription, Parent: iParent, Count: iCount}
	if _, err := Engine.Insert(_TermTaxonomy); err != nil {
		return err
	} else {
		return nil
	}
}

// PostTermTaxonomy Post TermTaxonomy via iTermTaxonomy
func PostTermTaxonomy(iTermTaxonomy *TermTaxonomy) (int64, error) {
	_, err := Engine.Insert(iTermTaxonomy)
	return iTermTaxonomy.TermTaxonomyId, err
}

// PutTermTaxonomy Put TermTaxonomy
func PutTermTaxonomy(iTermTaxonomy *TermTaxonomy) (int64, error) {
	_, err := Engine.Id(iTermTaxonomy.TermTaxonomyId).Update(iTermTaxonomy)
	return iTermTaxonomy.TermTaxonomyId, err
}

// PutTermTaxonomyByTermTaxonomyId Put TermTaxonomy via TermTaxonomyId
func PutTermTaxonomyByTermTaxonomyId(TermTaxonomyId_ int64, iTermTaxonomy *TermTaxonomy) (int64, error) {
	row, err := Engine.Update(iTermTaxonomy, &TermTaxonomy{TermTaxonomyId: TermTaxonomyId_})
	return row, err
}

// PutTermTaxonomyByTermId Put TermTaxonomy via TermId
func PutTermTaxonomyByTermId(TermId_ int64, iTermTaxonomy *TermTaxonomy) (int64, error) {
	row, err := Engine.Update(iTermTaxonomy, &TermTaxonomy{TermId: TermId_})
	return row, err
}

// PutTermTaxonomyByTaxonomy Put TermTaxonomy via Taxonomy
func PutTermTaxonomyByTaxonomy(Taxonomy_ string, iTermTaxonomy *TermTaxonomy) (int64, error) {
	row, err := Engine.Update(iTermTaxonomy, &TermTaxonomy{Taxonomy: Taxonomy_})
	return row, err
}

// PutTermTaxonomyByDescription Put TermTaxonomy via Description
func PutTermTaxonomyByDescription(Description_ string, iTermTaxonomy *TermTaxonomy) (int64, error) {
	row, err := Engine.Update(iTermTaxonomy, &TermTaxonomy{Description: Description_})
	return row, err
}

// PutTermTaxonomyByParent Put TermTaxonomy via Parent
func PutTermTaxonomyByParent(Parent_ int64, iTermTaxonomy *TermTaxonomy) (int64, error) {
	row, err := Engine.Update(iTermTaxonomy, &TermTaxonomy{Parent: Parent_})
	return row, err
}

// PutTermTaxonomyByCount Put TermTaxonomy via Count
func PutTermTaxonomyByCount(Count_ int64, iTermTaxonomy *TermTaxonomy) (int64, error) {
	row, err := Engine.Update(iTermTaxonomy, &TermTaxonomy{Count: Count_})
	return row, err
}

// UpdateTermTaxonomyByTermTaxonomyId via map[string]interface{}{}
func UpdateTermTaxonomyByTermTaxonomyId(iTermTaxonomyId int64, iTermTaxonomyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(TermTaxonomy)).Where("term_taxonomy_id = ?", iTermTaxonomyId).Update(iTermTaxonomyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTermTaxonomyByTermId via map[string]interface{}{}
func UpdateTermTaxonomyByTermId(iTermId int64, iTermTaxonomyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(TermTaxonomy)).Where("term_id = ?", iTermId).Update(iTermTaxonomyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTermTaxonomyByTaxonomy via map[string]interface{}{}
func UpdateTermTaxonomyByTaxonomy(iTaxonomy string, iTermTaxonomyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(TermTaxonomy)).Where("taxonomy = ?", iTaxonomy).Update(iTermTaxonomyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTermTaxonomyByDescription via map[string]interface{}{}
func UpdateTermTaxonomyByDescription(iDescription string, iTermTaxonomyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(TermTaxonomy)).Where("description = ?", iDescription).Update(iTermTaxonomyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTermTaxonomyByParent via map[string]interface{}{}
func UpdateTermTaxonomyByParent(iParent int64, iTermTaxonomyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(TermTaxonomy)).Where("parent = ?", iParent).Update(iTermTaxonomyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTermTaxonomyByCount via map[string]interface{}{}
func UpdateTermTaxonomyByCount(iCount int64, iTermTaxonomyMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(TermTaxonomy)).Where("count = ?", iCount).Update(iTermTaxonomyMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteTermTaxonomy Delete TermTaxonomy
func DeleteTermTaxonomy(iTermTaxonomyId int64) (int64, error) {
	row, err := Engine.Id(iTermTaxonomyId).Delete(new(TermTaxonomy))
	return row, err
}

// DeleteTermTaxonomyByTermTaxonomyId Delete TermTaxonomy via TermTaxonomyId
func DeleteTermTaxonomyByTermTaxonomyId(iTermTaxonomyId int64) (err error) {
	var has bool
	var _TermTaxonomy = &TermTaxonomy{TermTaxonomyId: iTermTaxonomyId}
	if has, err = Engine.Get(_TermTaxonomy); (has == true) && (err == nil) {
		if row, err := Engine.Where("term_taxonomy_id = ?", iTermTaxonomyId).Delete(new(TermTaxonomy)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteTermTaxonomyByTermId Delete TermTaxonomy via TermId
func DeleteTermTaxonomyByTermId(iTermId int64) (err error) {
	var has bool
	var _TermTaxonomy = &TermTaxonomy{TermId: iTermId}
	if has, err = Engine.Get(_TermTaxonomy); (has == true) && (err == nil) {
		if row, err := Engine.Where("term_id = ?", iTermId).Delete(new(TermTaxonomy)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteTermTaxonomyByTaxonomy Delete TermTaxonomy via Taxonomy
func DeleteTermTaxonomyByTaxonomy(iTaxonomy string) (err error) {
	var has bool
	var _TermTaxonomy = &TermTaxonomy{Taxonomy: iTaxonomy}
	if has, err = Engine.Get(_TermTaxonomy); (has == true) && (err == nil) {
		if row, err := Engine.Where("taxonomy = ?", iTaxonomy).Delete(new(TermTaxonomy)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteTermTaxonomyByDescription Delete TermTaxonomy via Description
func DeleteTermTaxonomyByDescription(iDescription string) (err error) {
	var has bool
	var _TermTaxonomy = &TermTaxonomy{Description: iDescription}
	if has, err = Engine.Get(_TermTaxonomy); (has == true) && (err == nil) {
		if row, err := Engine.Where("description = ?", iDescription).Delete(new(TermTaxonomy)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteTermTaxonomyByParent Delete TermTaxonomy via Parent
func DeleteTermTaxonomyByParent(iParent int64) (err error) {
	var has bool
	var _TermTaxonomy = &TermTaxonomy{Parent: iParent}
	if has, err = Engine.Get(_TermTaxonomy); (has == true) && (err == nil) {
		if row, err := Engine.Where("parent = ?", iParent).Delete(new(TermTaxonomy)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteTermTaxonomyByCount Delete TermTaxonomy via Count
func DeleteTermTaxonomyByCount(iCount int64) (err error) {
	var has bool
	var _TermTaxonomy = &TermTaxonomy{Count: iCount}
	if has, err = Engine.Get(_TermTaxonomy); (has == true) && (err == nil) {
		if row, err := Engine.Where("count = ?", iCount).Delete(new(TermTaxonomy)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
