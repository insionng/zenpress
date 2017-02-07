package model

type TermRelationships struct {
	ObjectId       int64 `xorm:"not null pk default 0 BIGINT(20)"`
	TermTaxonomyId int64 `xorm:"not null pk default 0 index BIGINT(20)"`
	TermOrder      int   `xorm:"not null default 0 INT(11)"`
}

// GetTermRelationshipsesCount TermRelationshipss' Count
func GetTermRelationshipsesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&TermRelationships{})
	return total, err
}

// GetTermRelationshipsCountByObjectId Get TermRelationships via ObjectId
func GetTermRelationshipsCountByObjectId(iObjectId int64) int64 {
	n, _ := Engine.Where("object_id = ?", iObjectId).Count(&TermRelationships{ObjectId: iObjectId})
	return n
}

// GetTermRelationshipsCountByTermTaxonomyId Get TermRelationships via TermTaxonomyId
func GetTermRelationshipsCountByTermTaxonomyId(iTermTaxonomyId int64) int64 {
	n, _ := Engine.Where("term_taxonomy_id = ?", iTermTaxonomyId).Count(&TermRelationships{TermTaxonomyId: iTermTaxonomyId})
	return n
}

// GetTermRelationshipsCountByTermOrder Get TermRelationships via TermOrder
func GetTermRelationshipsCountByTermOrder(iTermOrder int) int64 {
	n, _ := Engine.Where("term_order = ?", iTermOrder).Count(&TermRelationships{TermOrder: iTermOrder})
	return n
}

// GetTermRelationshipsesByObjectIdAndTermTaxonomyIdAndTermOrder Get TermRelationshipses via ObjectIdAndTermTaxonomyIdAndTermOrder
func GetTermRelationshipsesByObjectIdAndTermTaxonomyIdAndTermOrder(offset int, limit int, ObjectId_ int64, TermTaxonomyId_ int64, TermOrder_ int) (*[]*TermRelationships, error) {
	var _TermRelationships = new([]*TermRelationships)
	err := Engine.Table("term_relationships").Where("object_id = ? and term_taxonomy_id = ? and term_order = ?", ObjectId_, TermTaxonomyId_, TermOrder_).Limit(limit, offset).Find(_TermRelationships)
	return _TermRelationships, err
}

// GetTermRelationshipsesByObjectIdAndTermTaxonomyId Get TermRelationshipses via ObjectIdAndTermTaxonomyId
func GetTermRelationshipsesByObjectIdAndTermTaxonomyId(offset int, limit int, ObjectId_ int64, TermTaxonomyId_ int64) (*[]*TermRelationships, error) {
	var _TermRelationships = new([]*TermRelationships)
	err := Engine.Table("term_relationships").Where("object_id = ? and term_taxonomy_id = ?", ObjectId_, TermTaxonomyId_).Limit(limit, offset).Find(_TermRelationships)
	return _TermRelationships, err
}

// GetTermRelationshipsesByObjectIdAndTermOrder Get TermRelationshipses via ObjectIdAndTermOrder
func GetTermRelationshipsesByObjectIdAndTermOrder(offset int, limit int, ObjectId_ int64, TermOrder_ int) (*[]*TermRelationships, error) {
	var _TermRelationships = new([]*TermRelationships)
	err := Engine.Table("term_relationships").Where("object_id = ? and term_order = ?", ObjectId_, TermOrder_).Limit(limit, offset).Find(_TermRelationships)
	return _TermRelationships, err
}

// GetTermRelationshipsesByTermTaxonomyIdAndTermOrder Get TermRelationshipses via TermTaxonomyIdAndTermOrder
func GetTermRelationshipsesByTermTaxonomyIdAndTermOrder(offset int, limit int, TermTaxonomyId_ int64, TermOrder_ int) (*[]*TermRelationships, error) {
	var _TermRelationships = new([]*TermRelationships)
	err := Engine.Table("term_relationships").Where("term_taxonomy_id = ? and term_order = ?", TermTaxonomyId_, TermOrder_).Limit(limit, offset).Find(_TermRelationships)
	return _TermRelationships, err
}

// GetTermRelationshipses Get TermRelationshipses via field
func GetTermRelationshipses(offset int, limit int, field string) (*[]*TermRelationships, error) {
	var _TermRelationships = new([]*TermRelationships)
	err := Engine.Table("term_relationships").Limit(limit, offset).Desc(field).Find(_TermRelationships)
	return _TermRelationships, err
}

// GetTermRelationshipsesByObjectId Get TermRelationshipss via ObjectId
func GetTermRelationshipsesByObjectId(offset int, limit int, ObjectId_ int64, field string) (*[]*TermRelationships, error) {
	var _TermRelationships = new([]*TermRelationships)
	err := Engine.Table("term_relationships").Where("object_id = ?", ObjectId_).Limit(limit, offset).Desc(field).Find(_TermRelationships)
	return _TermRelationships, err
}

// GetTermRelationshipsesByTermTaxonomyId Get TermRelationshipss via TermTaxonomyId
func GetTermRelationshipsesByTermTaxonomyId(offset int, limit int, TermTaxonomyId_ int64, field string) (*[]*TermRelationships, error) {
	var _TermRelationships = new([]*TermRelationships)
	err := Engine.Table("term_relationships").Where("term_taxonomy_id = ?", TermTaxonomyId_).Limit(limit, offset).Desc(field).Find(_TermRelationships)
	return _TermRelationships, err
}

// GetTermRelationshipsesByTermOrder Get TermRelationshipss via TermOrder
func GetTermRelationshipsesByTermOrder(offset int, limit int, TermOrder_ int, field string) (*[]*TermRelationships, error) {
	var _TermRelationships = new([]*TermRelationships)
	err := Engine.Table("term_relationships").Where("term_order = ?", TermOrder_).Limit(limit, offset).Desc(field).Find(_TermRelationships)
	return _TermRelationships, err
}

// HasTermRelationshipsByObjectId Has TermRelationships via ObjectId
func HasTermRelationshipsByObjectId(iObjectId int64) bool {
	if has, err := Engine.Where("object_id = ?", iObjectId).Get(new(TermRelationships)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTermRelationshipsByTermTaxonomyId Has TermRelationships via TermTaxonomyId
func HasTermRelationshipsByTermTaxonomyId(iTermTaxonomyId int64) bool {
	if has, err := Engine.Where("term_taxonomy_id = ?", iTermTaxonomyId).Get(new(TermRelationships)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTermRelationshipsByTermOrder Has TermRelationships via TermOrder
func HasTermRelationshipsByTermOrder(iTermOrder int) bool {
	if has, err := Engine.Where("term_order = ?", iTermOrder).Get(new(TermRelationships)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetTermRelationshipsByObjectId Get TermRelationships via ObjectId
func GetTermRelationshipsByObjectId(iObjectId int64) (*TermRelationships, error) {
	var _TermRelationships = &TermRelationships{ObjectId: iObjectId}
	has, err := Engine.Get(_TermRelationships)
	if has {
		return _TermRelationships, err
	} else {
		return nil, err
	}
}

// GetTermRelationshipsByTermTaxonomyId Get TermRelationships via TermTaxonomyId
func GetTermRelationshipsByTermTaxonomyId(iTermTaxonomyId int64) (*TermRelationships, error) {
	var _TermRelationships = &TermRelationships{TermTaxonomyId: iTermTaxonomyId}
	has, err := Engine.Get(_TermRelationships)
	if has {
		return _TermRelationships, err
	} else {
		return nil, err
	}
}

// GetTermRelationshipsByTermOrder Get TermRelationships via TermOrder
func GetTermRelationshipsByTermOrder(iTermOrder int) (*TermRelationships, error) {
	var _TermRelationships = &TermRelationships{TermOrder: iTermOrder}
	has, err := Engine.Get(_TermRelationships)
	if has {
		return _TermRelationships, err
	} else {
		return nil, err
	}
}

// SetTermRelationshipsByObjectId Set TermRelationships via ObjectId
func SetTermRelationshipsByObjectId(iObjectId int64, term_relationships *TermRelationships) (int64, error) {
	term_relationships.ObjectId = iObjectId
	return Engine.Insert(term_relationships)
}

// SetTermRelationshipsByTermTaxonomyId Set TermRelationships via TermTaxonomyId
func SetTermRelationshipsByTermTaxonomyId(iTermTaxonomyId int64, term_relationships *TermRelationships) (int64, error) {
	term_relationships.TermTaxonomyId = iTermTaxonomyId
	return Engine.Insert(term_relationships)
}

// SetTermRelationshipsByTermOrder Set TermRelationships via TermOrder
func SetTermRelationshipsByTermOrder(iTermOrder int, term_relationships *TermRelationships) (int64, error) {
	term_relationships.TermOrder = iTermOrder
	return Engine.Insert(term_relationships)
}

// AddTermRelationships Add TermRelationships via all columns
func AddTermRelationships(iObjectId int64, iTermTaxonomyId int64, iTermOrder int) error {
	_TermRelationships := &TermRelationships{ObjectId: iObjectId, TermTaxonomyId: iTermTaxonomyId, TermOrder: iTermOrder}
	if _, err := Engine.Insert(_TermRelationships); err != nil {
		return err
	} else {
		return nil
	}
}

// PostTermRelationships Post TermRelationships via iTermRelationships
func PostTermRelationships(iTermRelationships *TermRelationships) (int64, error) {
	_, err := Engine.Insert(iTermRelationships)
	return iTermRelationships.ObjectId, err
}

// PutTermRelationships Put TermRelationships
func PutTermRelationships(iTermRelationships *TermRelationships) (int64, error) {
	_, err := Engine.Id(iTermRelationships.ObjectId).Update(iTermRelationships)
	return iTermRelationships.ObjectId, err
}

// PutTermRelationshipsByObjectId Put TermRelationships via ObjectId
func PutTermRelationshipsByObjectId(ObjectId_ int64, iTermRelationships *TermRelationships) (int64, error) {
	row, err := Engine.Update(iTermRelationships, &TermRelationships{ObjectId: ObjectId_})
	return row, err
}

// PutTermRelationshipsByTermTaxonomyId Put TermRelationships via TermTaxonomyId
func PutTermRelationshipsByTermTaxonomyId(TermTaxonomyId_ int64, iTermRelationships *TermRelationships) (int64, error) {
	row, err := Engine.Update(iTermRelationships, &TermRelationships{TermTaxonomyId: TermTaxonomyId_})
	return row, err
}

// PutTermRelationshipsByTermOrder Put TermRelationships via TermOrder
func PutTermRelationshipsByTermOrder(TermOrder_ int, iTermRelationships *TermRelationships) (int64, error) {
	row, err := Engine.Update(iTermRelationships, &TermRelationships{TermOrder: TermOrder_})
	return row, err
}

// UpdateTermRelationshipsByObjectId via map[string]interface{}{}
func UpdateTermRelationshipsByObjectId(iObjectId int64, iTermRelationshipsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(TermRelationships)).Where("object_id = ?", iObjectId).Update(iTermRelationshipsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTermRelationshipsByTermTaxonomyId via map[string]interface{}{}
func UpdateTermRelationshipsByTermTaxonomyId(iTermTaxonomyId int64, iTermRelationshipsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(TermRelationships)).Where("term_taxonomy_id = ?", iTermTaxonomyId).Update(iTermRelationshipsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTermRelationshipsByTermOrder via map[string]interface{}{}
func UpdateTermRelationshipsByTermOrder(iTermOrder int, iTermRelationshipsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(TermRelationships)).Where("term_order = ?", iTermOrder).Update(iTermRelationshipsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteTermRelationships Delete TermRelationships
func DeleteTermRelationships(iObjectId int64) (int64, error) {
	row, err := Engine.Id(iObjectId).Delete(new(TermRelationships))
	return row, err
}

// DeleteTermRelationshipsByObjectId Delete TermRelationships via ObjectId
func DeleteTermRelationshipsByObjectId(iObjectId int64) (err error) {
	var has bool
	var _TermRelationships = &TermRelationships{ObjectId: iObjectId}
	if has, err = Engine.Get(_TermRelationships); (has == true) && (err == nil) {
		if row, err := Engine.Where("object_id = ?", iObjectId).Delete(new(TermRelationships)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteTermRelationshipsByTermTaxonomyId Delete TermRelationships via TermTaxonomyId
func DeleteTermRelationshipsByTermTaxonomyId(iTermTaxonomyId int64) (err error) {
	var has bool
	var _TermRelationships = &TermRelationships{TermTaxonomyId: iTermTaxonomyId}
	if has, err = Engine.Get(_TermRelationships); (has == true) && (err == nil) {
		if row, err := Engine.Where("term_taxonomy_id = ?", iTermTaxonomyId).Delete(new(TermRelationships)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteTermRelationshipsByTermOrder Delete TermRelationships via TermOrder
func DeleteTermRelationshipsByTermOrder(iTermOrder int) (err error) {
	var has bool
	var _TermRelationships = &TermRelationships{TermOrder: iTermOrder}
	if has, err = Engine.Get(_TermRelationships); (has == true) && (err == nil) {
		if row, err := Engine.Where("term_order = ?", iTermOrder).Delete(new(TermRelationships)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
