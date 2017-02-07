package model

type Termmeta struct {
	MetaId    int64  `xorm:"not null pk autoincr BIGINT(20)"`
	TermId    int64  `xorm:"not null default 0 index BIGINT(20)"`
	MetaKey   string `xorm:"index VARCHAR(255)"`
	MetaValue string `xorm:"LONGTEXT"`
}

// GetTermmetasCount Termmetas' Count
func GetTermmetasCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&Termmeta{})
	return total, err
}

// GetTermmetaCountByMetaId Get Termmeta via MetaId
func GetTermmetaCountByMetaId(iMetaId int64) int64 {
	n, _ := Engine.Where("meta_id = ?", iMetaId).Count(&Termmeta{MetaId: iMetaId})
	return n
}

// GetTermmetaCountByTermId Get Termmeta via TermId
func GetTermmetaCountByTermId(iTermId int64) int64 {
	n, _ := Engine.Where("term_id = ?", iTermId).Count(&Termmeta{TermId: iTermId})
	return n
}

// GetTermmetaCountByMetaKey Get Termmeta via MetaKey
func GetTermmetaCountByMetaKey(iMetaKey string) int64 {
	n, _ := Engine.Where("meta_key = ?", iMetaKey).Count(&Termmeta{MetaKey: iMetaKey})
	return n
}

// GetTermmetaCountByMetaValue Get Termmeta via MetaValue
func GetTermmetaCountByMetaValue(iMetaValue string) int64 {
	n, _ := Engine.Where("meta_value = ?", iMetaValue).Count(&Termmeta{MetaValue: iMetaValue})
	return n
}

// GetTermmetasByMetaIdAndTermIdAndMetaKey Get Termmetas via MetaIdAndTermIdAndMetaKey
func GetTermmetasByMetaIdAndTermIdAndMetaKey(offset int, limit int, MetaId_ int64, TermId_ int64, MetaKey_ string) (*[]*Termmeta, error) {
	var _Termmeta = new([]*Termmeta)
	err := Engine.Table("termmeta").Where("meta_id = ? and term_id = ? and meta_key = ?", MetaId_, TermId_, MetaKey_).Limit(limit, offset).Find(_Termmeta)
	return _Termmeta, err
}

// GetTermmetasByMetaIdAndTermIdAndMetaValue Get Termmetas via MetaIdAndTermIdAndMetaValue
func GetTermmetasByMetaIdAndTermIdAndMetaValue(offset int, limit int, MetaId_ int64, TermId_ int64, MetaValue_ string) (*[]*Termmeta, error) {
	var _Termmeta = new([]*Termmeta)
	err := Engine.Table("termmeta").Where("meta_id = ? and term_id = ? and meta_value = ?", MetaId_, TermId_, MetaValue_).Limit(limit, offset).Find(_Termmeta)
	return _Termmeta, err
}

// GetTermmetasByMetaIdAndMetaKeyAndMetaValue Get Termmetas via MetaIdAndMetaKeyAndMetaValue
func GetTermmetasByMetaIdAndMetaKeyAndMetaValue(offset int, limit int, MetaId_ int64, MetaKey_ string, MetaValue_ string) (*[]*Termmeta, error) {
	var _Termmeta = new([]*Termmeta)
	err := Engine.Table("termmeta").Where("meta_id = ? and meta_key = ? and meta_value = ?", MetaId_, MetaKey_, MetaValue_).Limit(limit, offset).Find(_Termmeta)
	return _Termmeta, err
}

// GetTermmetasByTermIdAndMetaKeyAndMetaValue Get Termmetas via TermIdAndMetaKeyAndMetaValue
func GetTermmetasByTermIdAndMetaKeyAndMetaValue(offset int, limit int, TermId_ int64, MetaKey_ string, MetaValue_ string) (*[]*Termmeta, error) {
	var _Termmeta = new([]*Termmeta)
	err := Engine.Table("termmeta").Where("term_id = ? and meta_key = ? and meta_value = ?", TermId_, MetaKey_, MetaValue_).Limit(limit, offset).Find(_Termmeta)
	return _Termmeta, err
}

// GetTermmetasByMetaIdAndTermId Get Termmetas via MetaIdAndTermId
func GetTermmetasByMetaIdAndTermId(offset int, limit int, MetaId_ int64, TermId_ int64) (*[]*Termmeta, error) {
	var _Termmeta = new([]*Termmeta)
	err := Engine.Table("termmeta").Where("meta_id = ? and term_id = ?", MetaId_, TermId_).Limit(limit, offset).Find(_Termmeta)
	return _Termmeta, err
}

// GetTermmetasByMetaIdAndMetaKey Get Termmetas via MetaIdAndMetaKey
func GetTermmetasByMetaIdAndMetaKey(offset int, limit int, MetaId_ int64, MetaKey_ string) (*[]*Termmeta, error) {
	var _Termmeta = new([]*Termmeta)
	err := Engine.Table("termmeta").Where("meta_id = ? and meta_key = ?", MetaId_, MetaKey_).Limit(limit, offset).Find(_Termmeta)
	return _Termmeta, err
}

// GetTermmetasByMetaIdAndMetaValue Get Termmetas via MetaIdAndMetaValue
func GetTermmetasByMetaIdAndMetaValue(offset int, limit int, MetaId_ int64, MetaValue_ string) (*[]*Termmeta, error) {
	var _Termmeta = new([]*Termmeta)
	err := Engine.Table("termmeta").Where("meta_id = ? and meta_value = ?", MetaId_, MetaValue_).Limit(limit, offset).Find(_Termmeta)
	return _Termmeta, err
}

// GetTermmetasByTermIdAndMetaKey Get Termmetas via TermIdAndMetaKey
func GetTermmetasByTermIdAndMetaKey(offset int, limit int, TermId_ int64, MetaKey_ string) (*[]*Termmeta, error) {
	var _Termmeta = new([]*Termmeta)
	err := Engine.Table("termmeta").Where("term_id = ? and meta_key = ?", TermId_, MetaKey_).Limit(limit, offset).Find(_Termmeta)
	return _Termmeta, err
}

// GetTermmetasByTermIdAndMetaValue Get Termmetas via TermIdAndMetaValue
func GetTermmetasByTermIdAndMetaValue(offset int, limit int, TermId_ int64, MetaValue_ string) (*[]*Termmeta, error) {
	var _Termmeta = new([]*Termmeta)
	err := Engine.Table("termmeta").Where("term_id = ? and meta_value = ?", TermId_, MetaValue_).Limit(limit, offset).Find(_Termmeta)
	return _Termmeta, err
}

// GetTermmetasByMetaKeyAndMetaValue Get Termmetas via MetaKeyAndMetaValue
func GetTermmetasByMetaKeyAndMetaValue(offset int, limit int, MetaKey_ string, MetaValue_ string) (*[]*Termmeta, error) {
	var _Termmeta = new([]*Termmeta)
	err := Engine.Table("termmeta").Where("meta_key = ? and meta_value = ?", MetaKey_, MetaValue_).Limit(limit, offset).Find(_Termmeta)
	return _Termmeta, err
}

// GetTermmetas Get Termmetas via field
func GetTermmetas(offset int, limit int, field string) (*[]*Termmeta, error) {
	var _Termmeta = new([]*Termmeta)
	err := Engine.Table("termmeta").Limit(limit, offset).Desc(field).Find(_Termmeta)
	return _Termmeta, err
}

// GetTermmetasByMetaId Get Termmetas via MetaId
func GetTermmetasByMetaId(offset int, limit int, MetaId_ int64, field string) (*[]*Termmeta, error) {
	var _Termmeta = new([]*Termmeta)
	err := Engine.Table("termmeta").Where("meta_id = ?", MetaId_).Limit(limit, offset).Desc(field).Find(_Termmeta)
	return _Termmeta, err
}

// GetTermmetasByTermId Get Termmetas via TermId
func GetTermmetasByTermId(offset int, limit int, TermId_ int64, field string) (*[]*Termmeta, error) {
	var _Termmeta = new([]*Termmeta)
	err := Engine.Table("termmeta").Where("term_id = ?", TermId_).Limit(limit, offset).Desc(field).Find(_Termmeta)
	return _Termmeta, err
}

// GetTermmetasByMetaKey Get Termmetas via MetaKey
func GetTermmetasByMetaKey(offset int, limit int, MetaKey_ string, field string) (*[]*Termmeta, error) {
	var _Termmeta = new([]*Termmeta)
	err := Engine.Table("termmeta").Where("meta_key = ?", MetaKey_).Limit(limit, offset).Desc(field).Find(_Termmeta)
	return _Termmeta, err
}

// GetTermmetasByMetaValue Get Termmetas via MetaValue
func GetTermmetasByMetaValue(offset int, limit int, MetaValue_ string, field string) (*[]*Termmeta, error) {
	var _Termmeta = new([]*Termmeta)
	err := Engine.Table("termmeta").Where("meta_value = ?", MetaValue_).Limit(limit, offset).Desc(field).Find(_Termmeta)
	return _Termmeta, err
}

// HasTermmetaByMetaId Has Termmeta via MetaId
func HasTermmetaByMetaId(iMetaId int64) bool {
	if has, err := Engine.Where("meta_id = ?", iMetaId).Get(new(Termmeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTermmetaByTermId Has Termmeta via TermId
func HasTermmetaByTermId(iTermId int64) bool {
	if has, err := Engine.Where("term_id = ?", iTermId).Get(new(Termmeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTermmetaByMetaKey Has Termmeta via MetaKey
func HasTermmetaByMetaKey(iMetaKey string) bool {
	if has, err := Engine.Where("meta_key = ?", iMetaKey).Get(new(Termmeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTermmetaByMetaValue Has Termmeta via MetaValue
func HasTermmetaByMetaValue(iMetaValue string) bool {
	if has, err := Engine.Where("meta_value = ?", iMetaValue).Get(new(Termmeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetTermmetaByMetaId Get Termmeta via MetaId
func GetTermmetaByMetaId(iMetaId int64) (*Termmeta, error) {
	var _Termmeta = &Termmeta{MetaId: iMetaId}
	has, err := Engine.Get(_Termmeta)
	if has {
		return _Termmeta, err
	} else {
		return nil, err
	}
}

// GetTermmetaByTermId Get Termmeta via TermId
func GetTermmetaByTermId(iTermId int64) (*Termmeta, error) {
	var _Termmeta = &Termmeta{TermId: iTermId}
	has, err := Engine.Get(_Termmeta)
	if has {
		return _Termmeta, err
	} else {
		return nil, err
	}
}

// GetTermmetaByMetaKey Get Termmeta via MetaKey
func GetTermmetaByMetaKey(iMetaKey string) (*Termmeta, error) {
	var _Termmeta = &Termmeta{MetaKey: iMetaKey}
	has, err := Engine.Get(_Termmeta)
	if has {
		return _Termmeta, err
	} else {
		return nil, err
	}
}

// GetTermmetaByMetaValue Get Termmeta via MetaValue
func GetTermmetaByMetaValue(iMetaValue string) (*Termmeta, error) {
	var _Termmeta = &Termmeta{MetaValue: iMetaValue}
	has, err := Engine.Get(_Termmeta)
	if has {
		return _Termmeta, err
	} else {
		return nil, err
	}
}

// SetTermmetaByMetaId Set Termmeta via MetaId
func SetTermmetaByMetaId(iMetaId int64, termmeta *Termmeta) (int64, error) {
	termmeta.MetaId = iMetaId
	return Engine.Insert(termmeta)
}

// SetTermmetaByTermId Set Termmeta via TermId
func SetTermmetaByTermId(iTermId int64, termmeta *Termmeta) (int64, error) {
	termmeta.TermId = iTermId
	return Engine.Insert(termmeta)
}

// SetTermmetaByMetaKey Set Termmeta via MetaKey
func SetTermmetaByMetaKey(iMetaKey string, termmeta *Termmeta) (int64, error) {
	termmeta.MetaKey = iMetaKey
	return Engine.Insert(termmeta)
}

// SetTermmetaByMetaValue Set Termmeta via MetaValue
func SetTermmetaByMetaValue(iMetaValue string, termmeta *Termmeta) (int64, error) {
	termmeta.MetaValue = iMetaValue
	return Engine.Insert(termmeta)
}

// AddTermmeta Add Termmeta via all columns
func AddTermmeta(iMetaId int64, iTermId int64, iMetaKey string, iMetaValue string) error {
	_Termmeta := &Termmeta{MetaId: iMetaId, TermId: iTermId, MetaKey: iMetaKey, MetaValue: iMetaValue}
	if _, err := Engine.Insert(_Termmeta); err != nil {
		return err
	} else {
		return nil
	}
}

// PostTermmeta Post Termmeta via iTermmeta
func PostTermmeta(iTermmeta *Termmeta) (int64, error) {
	_, err := Engine.Insert(iTermmeta)
	return iTermmeta.MetaId, err
}

// PutTermmeta Put Termmeta
func PutTermmeta(iTermmeta *Termmeta) (int64, error) {
	_, err := Engine.Id(iTermmeta.MetaId).Update(iTermmeta)
	return iTermmeta.MetaId, err
}

// PutTermmetaByMetaId Put Termmeta via MetaId
func PutTermmetaByMetaId(MetaId_ int64, iTermmeta *Termmeta) (int64, error) {
	row, err := Engine.Update(iTermmeta, &Termmeta{MetaId: MetaId_})
	return row, err
}

// PutTermmetaByTermId Put Termmeta via TermId
func PutTermmetaByTermId(TermId_ int64, iTermmeta *Termmeta) (int64, error) {
	row, err := Engine.Update(iTermmeta, &Termmeta{TermId: TermId_})
	return row, err
}

// PutTermmetaByMetaKey Put Termmeta via MetaKey
func PutTermmetaByMetaKey(MetaKey_ string, iTermmeta *Termmeta) (int64, error) {
	row, err := Engine.Update(iTermmeta, &Termmeta{MetaKey: MetaKey_})
	return row, err
}

// PutTermmetaByMetaValue Put Termmeta via MetaValue
func PutTermmetaByMetaValue(MetaValue_ string, iTermmeta *Termmeta) (int64, error) {
	row, err := Engine.Update(iTermmeta, &Termmeta{MetaValue: MetaValue_})
	return row, err
}

// UpdateTermmetaByMetaId via map[string]interface{}{}
func UpdateTermmetaByMetaId(iMetaId int64, iTermmetaMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Termmeta)).Where("meta_id = ?", iMetaId).Update(iTermmetaMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTermmetaByTermId via map[string]interface{}{}
func UpdateTermmetaByTermId(iTermId int64, iTermmetaMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Termmeta)).Where("term_id = ?", iTermId).Update(iTermmetaMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTermmetaByMetaKey via map[string]interface{}{}
func UpdateTermmetaByMetaKey(iMetaKey string, iTermmetaMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Termmeta)).Where("meta_key = ?", iMetaKey).Update(iTermmetaMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTermmetaByMetaValue via map[string]interface{}{}
func UpdateTermmetaByMetaValue(iMetaValue string, iTermmetaMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Termmeta)).Where("meta_value = ?", iMetaValue).Update(iTermmetaMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteTermmeta Delete Termmeta
func DeleteTermmeta(iMetaId int64) (int64, error) {
	row, err := Engine.Id(iMetaId).Delete(new(Termmeta))
	return row, err
}

// DeleteTermmetaByMetaId Delete Termmeta via MetaId
func DeleteTermmetaByMetaId(iMetaId int64) (err error) {
	var has bool
	var _Termmeta = &Termmeta{MetaId: iMetaId}
	if has, err = Engine.Get(_Termmeta); (has == true) && (err == nil) {
		if row, err := Engine.Where("meta_id = ?", iMetaId).Delete(new(Termmeta)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteTermmetaByTermId Delete Termmeta via TermId
func DeleteTermmetaByTermId(iTermId int64) (err error) {
	var has bool
	var _Termmeta = &Termmeta{TermId: iTermId}
	if has, err = Engine.Get(_Termmeta); (has == true) && (err == nil) {
		if row, err := Engine.Where("term_id = ?", iTermId).Delete(new(Termmeta)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteTermmetaByMetaKey Delete Termmeta via MetaKey
func DeleteTermmetaByMetaKey(iMetaKey string) (err error) {
	var has bool
	var _Termmeta = &Termmeta{MetaKey: iMetaKey}
	if has, err = Engine.Get(_Termmeta); (has == true) && (err == nil) {
		if row, err := Engine.Where("meta_key = ?", iMetaKey).Delete(new(Termmeta)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteTermmetaByMetaValue Delete Termmeta via MetaValue
func DeleteTermmetaByMetaValue(iMetaValue string) (err error) {
	var has bool
	var _Termmeta = &Termmeta{MetaValue: iMetaValue}
	if has, err = Engine.Get(_Termmeta); (has == true) && (err == nil) {
		if row, err := Engine.Where("meta_value = ?", iMetaValue).Delete(new(Termmeta)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
