package model

type Terms struct {
	TermId    int64  `xorm:"not null pk autoincr BIGINT(20)"`
	Name      string `xorm:"not null default '' index VARCHAR(200)"`
	Slug      string `xorm:"not null default '' index VARCHAR(200)"`
	TermGroup int64  `xorm:"not null default 0 BIGINT(10)"`
}

// GetTermsesCount Termss' Count
func GetTermsesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&Terms{})
	return total, err
}

// GetTermsCountByTermId Get Terms via TermId
func GetTermsCountByTermId(iTermId int64) int64 {
	n, _ := Engine.Where("term_id = ?", iTermId).Count(&Terms{TermId: iTermId})
	return n
}

// GetTermsCountByName Get Terms via Name
func GetTermsCountByName(iName string) int64 {
	n, _ := Engine.Where("name = ?", iName).Count(&Terms{Name: iName})
	return n
}

// GetTermsCountBySlug Get Terms via Slug
func GetTermsCountBySlug(iSlug string) int64 {
	n, _ := Engine.Where("slug = ?", iSlug).Count(&Terms{Slug: iSlug})
	return n
}

// GetTermsCountByTermGroup Get Terms via TermGroup
func GetTermsCountByTermGroup(iTermGroup int64) int64 {
	n, _ := Engine.Where("term_group = ?", iTermGroup).Count(&Terms{TermGroup: iTermGroup})
	return n
}

// GetTermsesByTermIdAndNameAndSlug Get Termses via TermIdAndNameAndSlug
func GetTermsesByTermIdAndNameAndSlug(offset int, limit int, TermId_ int64, Name_ string, Slug_ string) (*[]*Terms, error) {
	var _Terms = new([]*Terms)
	err := Engine.Table("terms").Where("term_id = ? and name = ? and slug = ?", TermId_, Name_, Slug_).Limit(limit, offset).Find(_Terms)
	return _Terms, err
}

// GetTermsesByTermIdAndNameAndTermGroup Get Termses via TermIdAndNameAndTermGroup
func GetTermsesByTermIdAndNameAndTermGroup(offset int, limit int, TermId_ int64, Name_ string, TermGroup_ int64) (*[]*Terms, error) {
	var _Terms = new([]*Terms)
	err := Engine.Table("terms").Where("term_id = ? and name = ? and term_group = ?", TermId_, Name_, TermGroup_).Limit(limit, offset).Find(_Terms)
	return _Terms, err
}

// GetTermsesByTermIdAndSlugAndTermGroup Get Termses via TermIdAndSlugAndTermGroup
func GetTermsesByTermIdAndSlugAndTermGroup(offset int, limit int, TermId_ int64, Slug_ string, TermGroup_ int64) (*[]*Terms, error) {
	var _Terms = new([]*Terms)
	err := Engine.Table("terms").Where("term_id = ? and slug = ? and term_group = ?", TermId_, Slug_, TermGroup_).Limit(limit, offset).Find(_Terms)
	return _Terms, err
}

// GetTermsesByNameAndSlugAndTermGroup Get Termses via NameAndSlugAndTermGroup
func GetTermsesByNameAndSlugAndTermGroup(offset int, limit int, Name_ string, Slug_ string, TermGroup_ int64) (*[]*Terms, error) {
	var _Terms = new([]*Terms)
	err := Engine.Table("terms").Where("name = ? and slug = ? and term_group = ?", Name_, Slug_, TermGroup_).Limit(limit, offset).Find(_Terms)
	return _Terms, err
}

// GetTermsesByTermIdAndName Get Termses via TermIdAndName
func GetTermsesByTermIdAndName(offset int, limit int, TermId_ int64, Name_ string) (*[]*Terms, error) {
	var _Terms = new([]*Terms)
	err := Engine.Table("terms").Where("term_id = ? and name = ?", TermId_, Name_).Limit(limit, offset).Find(_Terms)
	return _Terms, err
}

// GetTermsesByTermIdAndSlug Get Termses via TermIdAndSlug
func GetTermsesByTermIdAndSlug(offset int, limit int, TermId_ int64, Slug_ string) (*[]*Terms, error) {
	var _Terms = new([]*Terms)
	err := Engine.Table("terms").Where("term_id = ? and slug = ?", TermId_, Slug_).Limit(limit, offset).Find(_Terms)
	return _Terms, err
}

// GetTermsesByTermIdAndTermGroup Get Termses via TermIdAndTermGroup
func GetTermsesByTermIdAndTermGroup(offset int, limit int, TermId_ int64, TermGroup_ int64) (*[]*Terms, error) {
	var _Terms = new([]*Terms)
	err := Engine.Table("terms").Where("term_id = ? and term_group = ?", TermId_, TermGroup_).Limit(limit, offset).Find(_Terms)
	return _Terms, err
}

// GetTermsesByNameAndSlug Get Termses via NameAndSlug
func GetTermsesByNameAndSlug(offset int, limit int, Name_ string, Slug_ string) (*[]*Terms, error) {
	var _Terms = new([]*Terms)
	err := Engine.Table("terms").Where("name = ? and slug = ?", Name_, Slug_).Limit(limit, offset).Find(_Terms)
	return _Terms, err
}

// GetTermsesByNameAndTermGroup Get Termses via NameAndTermGroup
func GetTermsesByNameAndTermGroup(offset int, limit int, Name_ string, TermGroup_ int64) (*[]*Terms, error) {
	var _Terms = new([]*Terms)
	err := Engine.Table("terms").Where("name = ? and term_group = ?", Name_, TermGroup_).Limit(limit, offset).Find(_Terms)
	return _Terms, err
}

// GetTermsesBySlugAndTermGroup Get Termses via SlugAndTermGroup
func GetTermsesBySlugAndTermGroup(offset int, limit int, Slug_ string, TermGroup_ int64) (*[]*Terms, error) {
	var _Terms = new([]*Terms)
	err := Engine.Table("terms").Where("slug = ? and term_group = ?", Slug_, TermGroup_).Limit(limit, offset).Find(_Terms)
	return _Terms, err
}

// GetTermses Get Termses via field
func GetTermses(offset int, limit int, field string) (*[]*Terms, error) {
	var _Terms = new([]*Terms)
	err := Engine.Table("terms").Limit(limit, offset).Desc(field).Find(_Terms)
	return _Terms, err
}

// GetTermsesByTermId Get Termss via TermId
func GetTermsesByTermId(offset int, limit int, TermId_ int64, field string) (*[]*Terms, error) {
	var _Terms = new([]*Terms)
	err := Engine.Table("terms").Where("term_id = ?", TermId_).Limit(limit, offset).Desc(field).Find(_Terms)
	return _Terms, err
}

// GetTermsesByName Get Termss via Name
func GetTermsesByName(offset int, limit int, Name_ string, field string) (*[]*Terms, error) {
	var _Terms = new([]*Terms)
	err := Engine.Table("terms").Where("name = ?", Name_).Limit(limit, offset).Desc(field).Find(_Terms)
	return _Terms, err
}

// GetTermsesBySlug Get Termss via Slug
func GetTermsesBySlug(offset int, limit int, Slug_ string, field string) (*[]*Terms, error) {
	var _Terms = new([]*Terms)
	err := Engine.Table("terms").Where("slug = ?", Slug_).Limit(limit, offset).Desc(field).Find(_Terms)
	return _Terms, err
}

// GetTermsesByTermGroup Get Termss via TermGroup
func GetTermsesByTermGroup(offset int, limit int, TermGroup_ int64, field string) (*[]*Terms, error) {
	var _Terms = new([]*Terms)
	err := Engine.Table("terms").Where("term_group = ?", TermGroup_).Limit(limit, offset).Desc(field).Find(_Terms)
	return _Terms, err
}

// HasTermsByTermId Has Terms via TermId
func HasTermsByTermId(iTermId int64) bool {
	if has, err := Engine.Where("term_id = ?", iTermId).Get(new(Terms)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTermsByName Has Terms via Name
func HasTermsByName(iName string) bool {
	if has, err := Engine.Where("name = ?", iName).Get(new(Terms)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTermsBySlug Has Terms via Slug
func HasTermsBySlug(iSlug string) bool {
	if has, err := Engine.Where("slug = ?", iSlug).Get(new(Terms)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTermsByTermGroup Has Terms via TermGroup
func HasTermsByTermGroup(iTermGroup int64) bool {
	if has, err := Engine.Where("term_group = ?", iTermGroup).Get(new(Terms)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetTermsByTermId Get Terms via TermId
func GetTermsByTermId(iTermId int64) (*Terms, error) {
	var _Terms = &Terms{TermId: iTermId}
	has, err := Engine.Get(_Terms)
	if has {
		return _Terms, err
	} else {
		return nil, err
	}
}

// GetTermsByName Get Terms via Name
func GetTermsByName(iName string) (*Terms, error) {
	var _Terms = &Terms{Name: iName}
	has, err := Engine.Get(_Terms)
	if has {
		return _Terms, err
	} else {
		return nil, err
	}
}

// GetTermsBySlug Get Terms via Slug
func GetTermsBySlug(iSlug string) (*Terms, error) {
	var _Terms = &Terms{Slug: iSlug}
	has, err := Engine.Get(_Terms)
	if has {
		return _Terms, err
	} else {
		return nil, err
	}
}

// GetTermsByTermGroup Get Terms via TermGroup
func GetTermsByTermGroup(iTermGroup int64) (*Terms, error) {
	var _Terms = &Terms{TermGroup: iTermGroup}
	has, err := Engine.Get(_Terms)
	if has {
		return _Terms, err
	} else {
		return nil, err
	}
}

// SetTermsByTermId Set Terms via TermId
func SetTermsByTermId(iTermId int64, terms *Terms) (int64, error) {
	terms.TermId = iTermId
	return Engine.Insert(terms)
}

// SetTermsByName Set Terms via Name
func SetTermsByName(iName string, terms *Terms) (int64, error) {
	terms.Name = iName
	return Engine.Insert(terms)
}

// SetTermsBySlug Set Terms via Slug
func SetTermsBySlug(iSlug string, terms *Terms) (int64, error) {
	terms.Slug = iSlug
	return Engine.Insert(terms)
}

// SetTermsByTermGroup Set Terms via TermGroup
func SetTermsByTermGroup(iTermGroup int64, terms *Terms) (int64, error) {
	terms.TermGroup = iTermGroup
	return Engine.Insert(terms)
}

// AddTerms Add Terms via all columns
func AddTerms(iTermId int64, iName string, iSlug string, iTermGroup int64) error {
	_Terms := &Terms{TermId: iTermId, Name: iName, Slug: iSlug, TermGroup: iTermGroup}
	if _, err := Engine.Insert(_Terms); err != nil {
		return err
	} else {
		return nil
	}
}

// PostTerms Post Terms via iTerms
func PostTerms(iTerms *Terms) (int64, error) {
	_, err := Engine.Insert(iTerms)
	return iTerms.TermId, err
}

// PutTerms Put Terms
func PutTerms(iTerms *Terms) (int64, error) {
	_, err := Engine.Id(iTerms.TermId).Update(iTerms)
	return iTerms.TermId, err
}

// PutTermsByTermId Put Terms via TermId
func PutTermsByTermId(TermId_ int64, iTerms *Terms) (int64, error) {
	row, err := Engine.Update(iTerms, &Terms{TermId: TermId_})
	return row, err
}

// PutTermsByName Put Terms via Name
func PutTermsByName(Name_ string, iTerms *Terms) (int64, error) {
	row, err := Engine.Update(iTerms, &Terms{Name: Name_})
	return row, err
}

// PutTermsBySlug Put Terms via Slug
func PutTermsBySlug(Slug_ string, iTerms *Terms) (int64, error) {
	row, err := Engine.Update(iTerms, &Terms{Slug: Slug_})
	return row, err
}

// PutTermsByTermGroup Put Terms via TermGroup
func PutTermsByTermGroup(TermGroup_ int64, iTerms *Terms) (int64, error) {
	row, err := Engine.Update(iTerms, &Terms{TermGroup: TermGroup_})
	return row, err
}

// UpdateTermsByTermId via map[string]interface{}{}
func UpdateTermsByTermId(iTermId int64, iTermsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Terms)).Where("term_id = ?", iTermId).Update(iTermsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTermsByName via map[string]interface{}{}
func UpdateTermsByName(iName string, iTermsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Terms)).Where("name = ?", iName).Update(iTermsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTermsBySlug via map[string]interface{}{}
func UpdateTermsBySlug(iSlug string, iTermsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Terms)).Where("slug = ?", iSlug).Update(iTermsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTermsByTermGroup via map[string]interface{}{}
func UpdateTermsByTermGroup(iTermGroup int64, iTermsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Terms)).Where("term_group = ?", iTermGroup).Update(iTermsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteTerms Delete Terms
func DeleteTerms(iTermId int64) (int64, error) {
	row, err := Engine.Id(iTermId).Delete(new(Terms))
	return row, err
}

// DeleteTermsByTermId Delete Terms via TermId
func DeleteTermsByTermId(iTermId int64) (err error) {
	var has bool
	var _Terms = &Terms{TermId: iTermId}
	if has, err = Engine.Get(_Terms); (has == true) && (err == nil) {
		if row, err := Engine.Where("term_id = ?", iTermId).Delete(new(Terms)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteTermsByName Delete Terms via Name
func DeleteTermsByName(iName string) (err error) {
	var has bool
	var _Terms = &Terms{Name: iName}
	if has, err = Engine.Get(_Terms); (has == true) && (err == nil) {
		if row, err := Engine.Where("name = ?", iName).Delete(new(Terms)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteTermsBySlug Delete Terms via Slug
func DeleteTermsBySlug(iSlug string) (err error) {
	var has bool
	var _Terms = &Terms{Slug: iSlug}
	if has, err = Engine.Get(_Terms); (has == true) && (err == nil) {
		if row, err := Engine.Where("slug = ?", iSlug).Delete(new(Terms)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteTermsByTermGroup Delete Terms via TermGroup
func DeleteTermsByTermGroup(iTermGroup int64) (err error) {
	var has bool
	var _Terms = &Terms{TermGroup: iTermGroup}
	if has, err = Engine.Get(_Terms); (has == true) && (err == nil) {
		if row, err := Engine.Where("term_group = ?", iTermGroup).Delete(new(Terms)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
