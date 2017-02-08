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

// GetTermmetaCountViaMetaId Get Termmeta via MetaId
func GetTermmetaCountViaMetaId(iMetaId int64) int64 {
	n, _ := Engine.Where("meta_id = ?", iMetaId).Count(&Termmeta{MetaId: iMetaId})
	return n
}

// GetTermmetaCountViaTermId Get Termmeta via TermId
func GetTermmetaCountViaTermId(iTermId int64) int64 {
	n, _ := Engine.Where("term_id = ?", iTermId).Count(&Termmeta{TermId: iTermId})
	return n
}

// GetTermmetaCountViaMetaKey Get Termmeta via MetaKey
func GetTermmetaCountViaMetaKey(iMetaKey string) int64 {
	n, _ := Engine.Where("meta_key = ?", iMetaKey).Count(&Termmeta{MetaKey: iMetaKey})
	return n
}

// GetTermmetaCountViaMetaValue Get Termmeta via MetaValue
func GetTermmetaCountViaMetaValue(iMetaValue string) int64 {
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

// GetTermmetasViaMetaId Get Termmetas via MetaId
func GetTermmetasViaMetaId(offset int, limit int, MetaId_ int64, field string) (*[]*Termmeta, error) {
	var _Termmeta = new([]*Termmeta)
	err := Engine.Table("termmeta").Where("meta_id = ?", MetaId_).Limit(limit, offset).Desc(field).Find(_Termmeta)
	return _Termmeta, err
}

// GetTermmetasViaTermId Get Termmetas via TermId
func GetTermmetasViaTermId(offset int, limit int, TermId_ int64, field string) (*[]*Termmeta, error) {
	var _Termmeta = new([]*Termmeta)
	err := Engine.Table("termmeta").Where("term_id = ?", TermId_).Limit(limit, offset).Desc(field).Find(_Termmeta)
	return _Termmeta, err
}

// GetTermmetasViaMetaKey Get Termmetas via MetaKey
func GetTermmetasViaMetaKey(offset int, limit int, MetaKey_ string, field string) (*[]*Termmeta, error) {
	var _Termmeta = new([]*Termmeta)
	err := Engine.Table("termmeta").Where("meta_key = ?", MetaKey_).Limit(limit, offset).Desc(field).Find(_Termmeta)
	return _Termmeta, err
}

// GetTermmetasViaMetaValue Get Termmetas via MetaValue
func GetTermmetasViaMetaValue(offset int, limit int, MetaValue_ string, field string) (*[]*Termmeta, error) {
	var _Termmeta = new([]*Termmeta)
	err := Engine.Table("termmeta").Where("meta_value = ?", MetaValue_).Limit(limit, offset).Desc(field).Find(_Termmeta)
	return _Termmeta, err
}

// HasTermmetaViaMetaId Has Termmeta via MetaId
func HasTermmetaViaMetaId(iMetaId int64) bool {
	if has, err := Engine.Where("meta_id = ?", iMetaId).Get(new(Termmeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTermmetaViaTermId Has Termmeta via TermId
func HasTermmetaViaTermId(iTermId int64) bool {
	if has, err := Engine.Where("term_id = ?", iTermId).Get(new(Termmeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTermmetaViaMetaKey Has Termmeta via MetaKey
func HasTermmetaViaMetaKey(iMetaKey string) bool {
	if has, err := Engine.Where("meta_key = ?", iMetaKey).Get(new(Termmeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasTermmetaViaMetaValue Has Termmeta via MetaValue
func HasTermmetaViaMetaValue(iMetaValue string) bool {
	if has, err := Engine.Where("meta_value = ?", iMetaValue).Get(new(Termmeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetTermmetaViaMetaId Get Termmeta via MetaId
func GetTermmetaViaMetaId(iMetaId int64) (*Termmeta, error) {
	var _Termmeta = &Termmeta{MetaId: iMetaId}
	has, err := Engine.Get(_Termmeta)
	if has {
		return _Termmeta, err
	} else {
		return nil, err
	}
}

// GetTermmetaViaTermId Get Termmeta via TermId
func GetTermmetaViaTermId(iTermId int64) (*Termmeta, error) {
	var _Termmeta = &Termmeta{TermId: iTermId}
	has, err := Engine.Get(_Termmeta)
	if has {
		return _Termmeta, err
	} else {
		return nil, err
	}
}

// GetTermmetaViaMetaKey Get Termmeta via MetaKey
func GetTermmetaViaMetaKey(iMetaKey string) (*Termmeta, error) {
	var _Termmeta = &Termmeta{MetaKey: iMetaKey}
	has, err := Engine.Get(_Termmeta)
	if has {
		return _Termmeta, err
	} else {
		return nil, err
	}
}

// GetTermmetaViaMetaValue Get Termmeta via MetaValue
func GetTermmetaViaMetaValue(iMetaValue string) (*Termmeta, error) {
	var _Termmeta = &Termmeta{MetaValue: iMetaValue}
	has, err := Engine.Get(_Termmeta)
	if has {
		return _Termmeta, err
	} else {
		return nil, err
	}
}

// SetTermmetaViaMetaId Set Termmeta via MetaId
func SetTermmetaViaMetaId(iMetaId int64, termmeta *Termmeta) (int64, error) {
	termmeta.MetaId = iMetaId
	return Engine.Insert(termmeta)
}

// SetTermmetaViaTermId Set Termmeta via TermId
func SetTermmetaViaTermId(iTermId int64, termmeta *Termmeta) (int64, error) {
	termmeta.TermId = iTermId
	return Engine.Insert(termmeta)
}

// SetTermmetaViaMetaKey Set Termmeta via MetaKey
func SetTermmetaViaMetaKey(iMetaKey string, termmeta *Termmeta) (int64, error) {
	termmeta.MetaKey = iMetaKey
	return Engine.Insert(termmeta)
}

// SetTermmetaViaMetaValue Set Termmeta via MetaValue
func SetTermmetaViaMetaValue(iMetaValue string, termmeta *Termmeta) (int64, error) {
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

// PutTermmetaViaMetaId Put Termmeta via MetaId
func PutTermmetaViaMetaId(MetaId_ int64, iTermmeta *Termmeta) (int64, error) {
	row, err := Engine.Update(iTermmeta, &Termmeta{MetaId: MetaId_})
	return row, err
}

// PutTermmetaViaTermId Put Termmeta via TermId
func PutTermmetaViaTermId(TermId_ int64, iTermmeta *Termmeta) (int64, error) {
	row, err := Engine.Update(iTermmeta, &Termmeta{TermId: TermId_})
	return row, err
}

// PutTermmetaViaMetaKey Put Termmeta via MetaKey
func PutTermmetaViaMetaKey(MetaKey_ string, iTermmeta *Termmeta) (int64, error) {
	row, err := Engine.Update(iTermmeta, &Termmeta{MetaKey: MetaKey_})
	return row, err
}

// PutTermmetaViaMetaValue Put Termmeta via MetaValue
func PutTermmetaViaMetaValue(MetaValue_ string, iTermmeta *Termmeta) (int64, error) {
	row, err := Engine.Update(iTermmeta, &Termmeta{MetaValue: MetaValue_})
	return row, err
}

// UpdateTermmetaViaMetaId via map[string]interface{}{}
func UpdateTermmetaViaMetaId(iMetaId int64, iTermmetaMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Termmeta)).Where("meta_id = ?", iMetaId).Update(iTermmetaMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTermmetaViaTermId via map[string]interface{}{}
func UpdateTermmetaViaTermId(iTermId int64, iTermmetaMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Termmeta)).Where("term_id = ?", iTermId).Update(iTermmetaMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTermmetaViaMetaKey via map[string]interface{}{}
func UpdateTermmetaViaMetaKey(iMetaKey string, iTermmetaMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Termmeta)).Where("meta_key = ?", iMetaKey).Update(iTermmetaMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateTermmetaViaMetaValue via map[string]interface{}{}
func UpdateTermmetaViaMetaValue(iMetaValue string, iTermmetaMap *map[string]interface{}) error {
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

// DeleteTermmetaViaMetaId Delete Termmeta via MetaId
func DeleteTermmetaViaMetaId(iMetaId int64) (err error) {
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

// DeleteTermmetaViaTermId Delete Termmeta via TermId
func DeleteTermmetaViaTermId(iTermId int64) (err error) {
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

// DeleteTermmetaViaMetaKey Delete Termmeta via MetaKey
func DeleteTermmetaViaMetaKey(iMetaKey string) (err error) {
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

// DeleteTermmetaViaMetaValue Delete Termmeta via MetaValue
func DeleteTermmetaViaMetaValue(iMetaValue string) (err error) {
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
