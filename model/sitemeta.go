package model

type Sitemeta struct {
	MetaId    int64  `xorm:"not null pk autoincr BIGINT(20)"`
	SiteId    int64  `xorm:"not null default 0 index BIGINT(20)"`
	MetaKey   string `xorm:"index VARCHAR(255)"`
	MetaValue string `xorm:"LONGTEXT"`
}

// GetSitemetasCount Sitemetas' Count
func GetSitemetasCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&Sitemeta{})
	return total, err
}

// GetSitemetaCountByMetaId Get Sitemeta via MetaId
func GetSitemetaCountByMetaId(iMetaId int64) int64 {
	n, _ := Engine.Where("meta_id = ?", iMetaId).Count(&Sitemeta{MetaId: iMetaId})
	return n
}

// GetSitemetaCountBySiteId Get Sitemeta via SiteId
func GetSitemetaCountBySiteId(iSiteId int64) int64 {
	n, _ := Engine.Where("site_id = ?", iSiteId).Count(&Sitemeta{SiteId: iSiteId})
	return n
}

// GetSitemetaCountByMetaKey Get Sitemeta via MetaKey
func GetSitemetaCountByMetaKey(iMetaKey string) int64 {
	n, _ := Engine.Where("meta_key = ?", iMetaKey).Count(&Sitemeta{MetaKey: iMetaKey})
	return n
}

// GetSitemetaCountByMetaValue Get Sitemeta via MetaValue
func GetSitemetaCountByMetaValue(iMetaValue string) int64 {
	n, _ := Engine.Where("meta_value = ?", iMetaValue).Count(&Sitemeta{MetaValue: iMetaValue})
	return n
}

// GetSitemetasByMetaIdAndSiteIdAndMetaKey Get Sitemetas via MetaIdAndSiteIdAndMetaKey
func GetSitemetasByMetaIdAndSiteIdAndMetaKey(offset int, limit int, MetaId_ int64, SiteId_ int64, MetaKey_ string) (*[]*Sitemeta, error) {
	var _Sitemeta = new([]*Sitemeta)
	err := Engine.Table("sitemeta").Where("meta_id = ? and site_id = ? and meta_key = ?", MetaId_, SiteId_, MetaKey_).Limit(limit, offset).Find(_Sitemeta)
	return _Sitemeta, err
}

// GetSitemetasByMetaIdAndSiteIdAndMetaValue Get Sitemetas via MetaIdAndSiteIdAndMetaValue
func GetSitemetasByMetaIdAndSiteIdAndMetaValue(offset int, limit int, MetaId_ int64, SiteId_ int64, MetaValue_ string) (*[]*Sitemeta, error) {
	var _Sitemeta = new([]*Sitemeta)
	err := Engine.Table("sitemeta").Where("meta_id = ? and site_id = ? and meta_value = ?", MetaId_, SiteId_, MetaValue_).Limit(limit, offset).Find(_Sitemeta)
	return _Sitemeta, err
}

// GetSitemetasByMetaIdAndMetaKeyAndMetaValue Get Sitemetas via MetaIdAndMetaKeyAndMetaValue
func GetSitemetasByMetaIdAndMetaKeyAndMetaValue(offset int, limit int, MetaId_ int64, MetaKey_ string, MetaValue_ string) (*[]*Sitemeta, error) {
	var _Sitemeta = new([]*Sitemeta)
	err := Engine.Table("sitemeta").Where("meta_id = ? and meta_key = ? and meta_value = ?", MetaId_, MetaKey_, MetaValue_).Limit(limit, offset).Find(_Sitemeta)
	return _Sitemeta, err
}

// GetSitemetasBySiteIdAndMetaKeyAndMetaValue Get Sitemetas via SiteIdAndMetaKeyAndMetaValue
func GetSitemetasBySiteIdAndMetaKeyAndMetaValue(offset int, limit int, SiteId_ int64, MetaKey_ string, MetaValue_ string) (*[]*Sitemeta, error) {
	var _Sitemeta = new([]*Sitemeta)
	err := Engine.Table("sitemeta").Where("site_id = ? and meta_key = ? and meta_value = ?", SiteId_, MetaKey_, MetaValue_).Limit(limit, offset).Find(_Sitemeta)
	return _Sitemeta, err
}

// GetSitemetasByMetaIdAndSiteId Get Sitemetas via MetaIdAndSiteId
func GetSitemetasByMetaIdAndSiteId(offset int, limit int, MetaId_ int64, SiteId_ int64) (*[]*Sitemeta, error) {
	var _Sitemeta = new([]*Sitemeta)
	err := Engine.Table("sitemeta").Where("meta_id = ? and site_id = ?", MetaId_, SiteId_).Limit(limit, offset).Find(_Sitemeta)
	return _Sitemeta, err
}

// GetSitemetasByMetaIdAndMetaKey Get Sitemetas via MetaIdAndMetaKey
func GetSitemetasByMetaIdAndMetaKey(offset int, limit int, MetaId_ int64, MetaKey_ string) (*[]*Sitemeta, error) {
	var _Sitemeta = new([]*Sitemeta)
	err := Engine.Table("sitemeta").Where("meta_id = ? and meta_key = ?", MetaId_, MetaKey_).Limit(limit, offset).Find(_Sitemeta)
	return _Sitemeta, err
}

// GetSitemetasByMetaIdAndMetaValue Get Sitemetas via MetaIdAndMetaValue
func GetSitemetasByMetaIdAndMetaValue(offset int, limit int, MetaId_ int64, MetaValue_ string) (*[]*Sitemeta, error) {
	var _Sitemeta = new([]*Sitemeta)
	err := Engine.Table("sitemeta").Where("meta_id = ? and meta_value = ?", MetaId_, MetaValue_).Limit(limit, offset).Find(_Sitemeta)
	return _Sitemeta, err
}

// GetSitemetasBySiteIdAndMetaKey Get Sitemetas via SiteIdAndMetaKey
func GetSitemetasBySiteIdAndMetaKey(offset int, limit int, SiteId_ int64, MetaKey_ string) (*[]*Sitemeta, error) {
	var _Sitemeta = new([]*Sitemeta)
	err := Engine.Table("sitemeta").Where("site_id = ? and meta_key = ?", SiteId_, MetaKey_).Limit(limit, offset).Find(_Sitemeta)
	return _Sitemeta, err
}

// GetSitemetasBySiteIdAndMetaValue Get Sitemetas via SiteIdAndMetaValue
func GetSitemetasBySiteIdAndMetaValue(offset int, limit int, SiteId_ int64, MetaValue_ string) (*[]*Sitemeta, error) {
	var _Sitemeta = new([]*Sitemeta)
	err := Engine.Table("sitemeta").Where("site_id = ? and meta_value = ?", SiteId_, MetaValue_).Limit(limit, offset).Find(_Sitemeta)
	return _Sitemeta, err
}

// GetSitemetasByMetaKeyAndMetaValue Get Sitemetas via MetaKeyAndMetaValue
func GetSitemetasByMetaKeyAndMetaValue(offset int, limit int, MetaKey_ string, MetaValue_ string) (*[]*Sitemeta, error) {
	var _Sitemeta = new([]*Sitemeta)
	err := Engine.Table("sitemeta").Where("meta_key = ? and meta_value = ?", MetaKey_, MetaValue_).Limit(limit, offset).Find(_Sitemeta)
	return _Sitemeta, err
}

// GetSitemetas Get Sitemetas via field
func GetSitemetas(offset int, limit int, field string) (*[]*Sitemeta, error) {
	var _Sitemeta = new([]*Sitemeta)
	err := Engine.Table("sitemeta").Limit(limit, offset).Desc(field).Find(_Sitemeta)
	return _Sitemeta, err
}

// GetSitemetasByMetaId Get Sitemetas via MetaId
func GetSitemetasByMetaId(offset int, limit int, MetaId_ int64, field string) (*[]*Sitemeta, error) {
	var _Sitemeta = new([]*Sitemeta)
	err := Engine.Table("sitemeta").Where("meta_id = ?", MetaId_).Limit(limit, offset).Desc(field).Find(_Sitemeta)
	return _Sitemeta, err
}

// GetSitemetasBySiteId Get Sitemetas via SiteId
func GetSitemetasBySiteId(offset int, limit int, SiteId_ int64, field string) (*[]*Sitemeta, error) {
	var _Sitemeta = new([]*Sitemeta)
	err := Engine.Table("sitemeta").Where("site_id = ?", SiteId_).Limit(limit, offset).Desc(field).Find(_Sitemeta)
	return _Sitemeta, err
}

// GetSitemetasByMetaKey Get Sitemetas via MetaKey
func GetSitemetasByMetaKey(offset int, limit int, MetaKey_ string, field string) (*[]*Sitemeta, error) {
	var _Sitemeta = new([]*Sitemeta)
	err := Engine.Table("sitemeta").Where("meta_key = ?", MetaKey_).Limit(limit, offset).Desc(field).Find(_Sitemeta)
	return _Sitemeta, err
}

// GetSitemetasByMetaValue Get Sitemetas via MetaValue
func GetSitemetasByMetaValue(offset int, limit int, MetaValue_ string, field string) (*[]*Sitemeta, error) {
	var _Sitemeta = new([]*Sitemeta)
	err := Engine.Table("sitemeta").Where("meta_value = ?", MetaValue_).Limit(limit, offset).Desc(field).Find(_Sitemeta)
	return _Sitemeta, err
}

// HasSitemetaByMetaId Has Sitemeta via MetaId
func HasSitemetaByMetaId(iMetaId int64) bool {
	if has, err := Engine.Where("meta_id = ?", iMetaId).Get(new(Sitemeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasSitemetaBySiteId Has Sitemeta via SiteId
func HasSitemetaBySiteId(iSiteId int64) bool {
	if has, err := Engine.Where("site_id = ?", iSiteId).Get(new(Sitemeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasSitemetaByMetaKey Has Sitemeta via MetaKey
func HasSitemetaByMetaKey(iMetaKey string) bool {
	if has, err := Engine.Where("meta_key = ?", iMetaKey).Get(new(Sitemeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasSitemetaByMetaValue Has Sitemeta via MetaValue
func HasSitemetaByMetaValue(iMetaValue string) bool {
	if has, err := Engine.Where("meta_value = ?", iMetaValue).Get(new(Sitemeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetSitemetaByMetaId Get Sitemeta via MetaId
func GetSitemetaByMetaId(iMetaId int64) (*Sitemeta, error) {
	var _Sitemeta = &Sitemeta{MetaId: iMetaId}
	has, err := Engine.Get(_Sitemeta)
	if has {
		return _Sitemeta, err
	} else {
		return nil, err
	}
}

// GetSitemetaBySiteId Get Sitemeta via SiteId
func GetSitemetaBySiteId(iSiteId int64) (*Sitemeta, error) {
	var _Sitemeta = &Sitemeta{SiteId: iSiteId}
	has, err := Engine.Get(_Sitemeta)
	if has {
		return _Sitemeta, err
	} else {
		return nil, err
	}
}

// GetSitemetaByMetaKey Get Sitemeta via MetaKey
func GetSitemetaByMetaKey(iMetaKey string) (*Sitemeta, error) {
	var _Sitemeta = &Sitemeta{MetaKey: iMetaKey}
	has, err := Engine.Get(_Sitemeta)
	if has {
		return _Sitemeta, err
	} else {
		return nil, err
	}
}

// GetSitemetaByMetaValue Get Sitemeta via MetaValue
func GetSitemetaByMetaValue(iMetaValue string) (*Sitemeta, error) {
	var _Sitemeta = &Sitemeta{MetaValue: iMetaValue}
	has, err := Engine.Get(_Sitemeta)
	if has {
		return _Sitemeta, err
	} else {
		return nil, err
	}
}

// SetSitemetaByMetaId Set Sitemeta via MetaId
func SetSitemetaByMetaId(iMetaId int64, sitemeta *Sitemeta) (int64, error) {
	sitemeta.MetaId = iMetaId
	return Engine.Insert(sitemeta)
}

// SetSitemetaBySiteId Set Sitemeta via SiteId
func SetSitemetaBySiteId(iSiteId int64, sitemeta *Sitemeta) (int64, error) {
	sitemeta.SiteId = iSiteId
	return Engine.Insert(sitemeta)
}

// SetSitemetaByMetaKey Set Sitemeta via MetaKey
func SetSitemetaByMetaKey(iMetaKey string, sitemeta *Sitemeta) (int64, error) {
	sitemeta.MetaKey = iMetaKey
	return Engine.Insert(sitemeta)
}

// SetSitemetaByMetaValue Set Sitemeta via MetaValue
func SetSitemetaByMetaValue(iMetaValue string, sitemeta *Sitemeta) (int64, error) {
	sitemeta.MetaValue = iMetaValue
	return Engine.Insert(sitemeta)
}

// AddSitemeta Add Sitemeta via all columns
func AddSitemeta(iMetaId int64, iSiteId int64, iMetaKey string, iMetaValue string) error {
	_Sitemeta := &Sitemeta{MetaId: iMetaId, SiteId: iSiteId, MetaKey: iMetaKey, MetaValue: iMetaValue}
	if _, err := Engine.Insert(_Sitemeta); err != nil {
		return err
	} else {
		return nil
	}
}

// PostSitemeta Post Sitemeta via iSitemeta
func PostSitemeta(iSitemeta *Sitemeta) (int64, error) {
	_, err := Engine.Insert(iSitemeta)
	return iSitemeta.MetaId, err
}

// PutSitemeta Put Sitemeta
func PutSitemeta(iSitemeta *Sitemeta) (int64, error) {
	_, err := Engine.Id(iSitemeta.MetaId).Update(iSitemeta)
	return iSitemeta.MetaId, err
}

// PutSitemetaByMetaId Put Sitemeta via MetaId
func PutSitemetaByMetaId(MetaId_ int64, iSitemeta *Sitemeta) (int64, error) {
	row, err := Engine.Update(iSitemeta, &Sitemeta{MetaId: MetaId_})
	return row, err
}

// PutSitemetaBySiteId Put Sitemeta via SiteId
func PutSitemetaBySiteId(SiteId_ int64, iSitemeta *Sitemeta) (int64, error) {
	row, err := Engine.Update(iSitemeta, &Sitemeta{SiteId: SiteId_})
	return row, err
}

// PutSitemetaByMetaKey Put Sitemeta via MetaKey
func PutSitemetaByMetaKey(MetaKey_ string, iSitemeta *Sitemeta) (int64, error) {
	row, err := Engine.Update(iSitemeta, &Sitemeta{MetaKey: MetaKey_})
	return row, err
}

// PutSitemetaByMetaValue Put Sitemeta via MetaValue
func PutSitemetaByMetaValue(MetaValue_ string, iSitemeta *Sitemeta) (int64, error) {
	row, err := Engine.Update(iSitemeta, &Sitemeta{MetaValue: MetaValue_})
	return row, err
}

// UpdateSitemetaByMetaId via map[string]interface{}{}
func UpdateSitemetaByMetaId(iMetaId int64, iSitemetaMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Sitemeta)).Where("meta_id = ?", iMetaId).Update(iSitemetaMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateSitemetaBySiteId via map[string]interface{}{}
func UpdateSitemetaBySiteId(iSiteId int64, iSitemetaMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Sitemeta)).Where("site_id = ?", iSiteId).Update(iSitemetaMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateSitemetaByMetaKey via map[string]interface{}{}
func UpdateSitemetaByMetaKey(iMetaKey string, iSitemetaMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Sitemeta)).Where("meta_key = ?", iMetaKey).Update(iSitemetaMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateSitemetaByMetaValue via map[string]interface{}{}
func UpdateSitemetaByMetaValue(iMetaValue string, iSitemetaMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Sitemeta)).Where("meta_value = ?", iMetaValue).Update(iSitemetaMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteSitemeta Delete Sitemeta
func DeleteSitemeta(iMetaId int64) (int64, error) {
	row, err := Engine.Id(iMetaId).Delete(new(Sitemeta))
	return row, err
}

// DeleteSitemetaByMetaId Delete Sitemeta via MetaId
func DeleteSitemetaByMetaId(iMetaId int64) (err error) {
	var has bool
	var _Sitemeta = &Sitemeta{MetaId: iMetaId}
	if has, err = Engine.Get(_Sitemeta); (has == true) && (err == nil) {
		if row, err := Engine.Where("meta_id = ?", iMetaId).Delete(new(Sitemeta)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteSitemetaBySiteId Delete Sitemeta via SiteId
func DeleteSitemetaBySiteId(iSiteId int64) (err error) {
	var has bool
	var _Sitemeta = &Sitemeta{SiteId: iSiteId}
	if has, err = Engine.Get(_Sitemeta); (has == true) && (err == nil) {
		if row, err := Engine.Where("site_id = ?", iSiteId).Delete(new(Sitemeta)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteSitemetaByMetaKey Delete Sitemeta via MetaKey
func DeleteSitemetaByMetaKey(iMetaKey string) (err error) {
	var has bool
	var _Sitemeta = &Sitemeta{MetaKey: iMetaKey}
	if has, err = Engine.Get(_Sitemeta); (has == true) && (err == nil) {
		if row, err := Engine.Where("meta_key = ?", iMetaKey).Delete(new(Sitemeta)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteSitemetaByMetaValue Delete Sitemeta via MetaValue
func DeleteSitemetaByMetaValue(iMetaValue string) (err error) {
	var has bool
	var _Sitemeta = &Sitemeta{MetaValue: iMetaValue}
	if has, err = Engine.Get(_Sitemeta); (has == true) && (err == nil) {
		if row, err := Engine.Where("meta_value = ?", iMetaValue).Delete(new(Sitemeta)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
