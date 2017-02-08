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

// GetTermsCountViaTermId Get Terms via TermId
func GetTermsCountViaTermId(iTermId int64) int64 {
	n, _ := Engine.Where("term_id = ?", iTermId).Count(&Terms{TermId: iTermId})
	return n
}

// GetTermsCountViaName Get Terms via Name
func GetTermsCountViaName(iName string) int64 {
	n, _ := Engine.Where("name = ?", iName).Count(&Terms{Name: iName})
	return n
}

// GetTermsCountViaSlug Get Terms via Slug
func GetTermsCountViaSlug(iSlug string) int64 {
	n, _ := Engine.Where("slug = ?", iSlug).Count(&Terms{Slug: iSlug})
	return n
}

// GetTermsCountViaTermGroup Get Terms via TermGroup
func GetTermsCountViaTermGroup(iTermGroup int64) int64 {
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

// GetTermsesViaTermId Get Termss via TermId
func GetTermsesViaTermId(offset int, limit int, TermId_ int64, field string) (*[]*Terms, error) {
	var _Terms = new([]*Terms)
	err := Engine.Table("terms").Where("term_id = ?", TermId_).Limit(limit, offset).Desc(field).Find(_Terms)
	return _Terms, err
}

// GetTermsesViaName Get Termss via Name
func GetTermsesViaName(offset int, limit int, Name_ string, field string) (*[]*Terms, error) {
	var _Terms = new([]*Terms)
	err := Engine.Table("terms").Where("name = ?", Name_).Limit(limit, offset).Desc(field).Find(_Terms)
	return _Terms, err
}

// GetTermsesViaSlug Get Termss via Slug
func GetTermsesViaSlug(offset int, limit int, Slug_ string, field string) (*[]*Terms, error) {
	var _Terms = new([]*Terms)
	err := Engine.Table("terms").Where("slug = ?", Slug_).Limit(limit, offset).Desc(field).Find(_Terms)
	return _Terms, err
}

// GetTermsesViaTermGroup Get Termss via TermGroup
func GetTermsesViaTermGroup(offset int, limit int, TermGroup_ int64, field string) (*[]*Terms, error) {
	var _Terms = new([]*Terms)
	err := Engine.Table("terms").Where("term_group = ?", TermGroup_).Limit(limit, offset).Desc(field).Find(_Terms)
	return _Terms, err
}

// HasTermsViaTermId Has Terms via TermId
func HasTermsViaTermId(iTermId int64) bool {
	if has, err := Engine.Where("term_id = ?", iTermId).Get(new(Terms)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTermsViaName Has Terms via Name
func HasTermsViaName(iName string) bool {
	if has, err := Engine.Where("name = ?", iName).Get(new(Terms)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTermsViaSlug Has Terms via Slug
func HasTermsViaSlug(iSlug string) bool {
	if has, err := Engine.Where("slug = ?", iSlug).Get(new(Terms)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTermsViaTermGroup Has Terms via TermGroup
func HasTermsViaTermGroup(iTermGroup int64) bool {
	if has, err := Engine.Where("term_group = ?", iTermGroup).Get(new(Terms)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetTermsViaTermId Get Terms via TermId
func GetTermsViaTermId(iTermId int64) (*Terms, error) {
	var _Terms = &Terms{TermId: iTermId}
	has, err := Engine.Get(_Terms)
	if has {
		return _Terms, err
	} else {
		return nil, err
	}
}

// GetTermsViaName Get Terms via Name
func GetTermsViaName(iName string) (*Terms, error) {
	var _Terms = &Terms{Name: iName}
	has, err := Engine.Get(_Terms)
	if has {
		return _Terms, err
	} else {
		return nil, err
	}
}

// GetTermsViaSlug Get Terms via Slug
func GetTermsViaSlug(iSlug string) (*Terms, error) {
	var _Terms = &Terms{Slug: iSlug}
	has, err := Engine.Get(_Terms)
	if has {
		return _Terms, err
	} else {
		return nil, err
	}
}

// GetTermsViaTermGroup Get Terms via TermGroup
func GetTermsViaTermGroup(iTermGroup int64) (*Terms, error) {
	var _Terms = &Terms{TermGroup: iTermGroup}
	has, err := Engine.Get(_Terms)
	if has {
		return _Terms, err
	} else {
		return nil, err
	}
}

// SetTermsViaTermId Set Terms via TermId
func SetTermsViaTermId(iTermId int64, terms *Terms) (int64, error) {
	terms.TermId = iTermId
	return Engine.Insert(terms)
}

// SetTermsViaName Set Terms via Name
func SetTermsViaName(iName string, terms *Terms) (int64, error) {
	terms.Name = iName
	return Engine.Insert(terms)
}

// SetTermsViaSlug Set Terms via Slug
func SetTermsViaSlug(iSlug string, terms *Terms) (int64, error) {
	terms.Slug = iSlug
	return Engine.Insert(terms)
}

// SetTermsViaTermGroup Set Terms via TermGroup
func SetTermsViaTermGroup(iTermGroup int64, terms *Terms) (int64, error) {
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

// PutTermsViaTermId Put Terms via TermId
func PutTermsViaTermId(TermId_ int64, iTerms *Terms) (int64, error) {
	row, err := Engine.Update(iTerms, &Terms{TermId: TermId_})
	return row, err
}

// PutTermsViaName Put Terms via Name
func PutTermsViaName(Name_ string, iTerms *Terms) (int64, error) {
	row, err := Engine.Update(iTerms, &Terms{Name: Name_})
	return row, err
}

// PutTermsViaSlug Put Terms via Slug
func PutTermsViaSlug(Slug_ string, iTerms *Terms) (int64, error) {
	row, err := Engine.Update(iTerms, &Terms{Slug: Slug_})
	return row, err
}

// PutTermsViaTermGroup Put Terms via TermGroup
func PutTermsViaTermGroup(TermGroup_ int64, iTerms *Terms) (int64, error) {
	row, err := Engine.Update(iTerms, &Terms{TermGroup: TermGroup_})
	return row, err
}

// UpdateTermsViaTermId via map[string]interface{}{}
func UpdateTermsViaTermId(iTermId int64, iTermsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Terms)).Where("term_id = ?", iTermId).Update(iTermsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTermsViaName via map[string]interface{}{}
func UpdateTermsViaName(iName string, iTermsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Terms)).Where("name = ?", iName).Update(iTermsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTermsViaSlug via map[string]interface{}{}
func UpdateTermsViaSlug(iSlug string, iTermsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Terms)).Where("slug = ?", iSlug).Update(iTermsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTermsViaTermGroup via map[string]interface{}{}
func UpdateTermsViaTermGroup(iTermGroup int64, iTermsMap *map[string]interface{}) error {
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

// DeleteTermsViaTermId Delete Terms via TermId
func DeleteTermsViaTermId(iTermId int64) (err error) {
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

// DeleteTermsViaName Delete Terms via Name
func DeleteTermsViaName(iName string) (err error) {
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

// DeleteTermsViaSlug Delete Terms via Slug
func DeleteTermsViaSlug(iSlug string) (err error) {
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

// DeleteTermsViaTermGroup Delete Terms via TermGroup
func DeleteTermsViaTermGroup(iTermGroup int64) (err error) {
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
