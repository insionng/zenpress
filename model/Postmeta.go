package model

type Postmeta struct {
	MetaId    int64  `xorm:"not null pk autoincr BIGINT(20)"`
	PostId    int64  `xorm:"not null default 0 index BIGINT(20)"`
	MetaKey   string `xorm:"index VARCHAR(255)"`
	MetaValue string `xorm:"LONGTEXT"`
}

// GetPostmetasCount Postmetas' Count
func GetPostmetasCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&Postmeta{})
	return total, err
}

// GetPostmetaCountByMetaId Get Postmeta via MetaId
func GetPostmetaCountByMetaId(iMetaId int64) int64 {
	n, _ := Engine.Where("meta_id = ?", iMetaId).Count(&Postmeta{MetaId: iMetaId})
	return n
}

// GetPostmetaCountByPostId Get Postmeta via PostId
func GetPostmetaCountByPostId(iPostId int64) int64 {
	n, _ := Engine.Where("post_id = ?", iPostId).Count(&Postmeta{PostId: iPostId})
	return n
}

// GetPostmetaCountByMetaKey Get Postmeta via MetaKey
func GetPostmetaCountByMetaKey(iMetaKey string) int64 {
	n, _ := Engine.Where("meta_key = ?", iMetaKey).Count(&Postmeta{MetaKey: iMetaKey})
	return n
}

// GetPostmetaCountByMetaValue Get Postmeta via MetaValue
func GetPostmetaCountByMetaValue(iMetaValue string) int64 {
	n, _ := Engine.Where("meta_value = ?", iMetaValue).Count(&Postmeta{MetaValue: iMetaValue})
	return n
}

// GetPostmetasByMetaIdAndPostIdAndMetaKey Get Postmetas via MetaIdAndPostIdAndMetaKey
func GetPostmetasByMetaIdAndPostIdAndMetaKey(offset int, limit int, MetaId_ int64, PostId_ int64, MetaKey_ string) (*[]*Postmeta, error) {
	var _Postmeta = new([]*Postmeta)
	err := Engine.Table("postmeta").Where("meta_id = ? and post_id = ? and meta_key = ?", MetaId_, PostId_, MetaKey_).Limit(limit, offset).Find(_Postmeta)
	return _Postmeta, err
}

// GetPostmetasByMetaIdAndPostIdAndMetaValue Get Postmetas via MetaIdAndPostIdAndMetaValue
func GetPostmetasByMetaIdAndPostIdAndMetaValue(offset int, limit int, MetaId_ int64, PostId_ int64, MetaValue_ string) (*[]*Postmeta, error) {
	var _Postmeta = new([]*Postmeta)
	err := Engine.Table("postmeta").Where("meta_id = ? and post_id = ? and meta_value = ?", MetaId_, PostId_, MetaValue_).Limit(limit, offset).Find(_Postmeta)
	return _Postmeta, err
}

// GetPostmetasByMetaIdAndMetaKeyAndMetaValue Get Postmetas via MetaIdAndMetaKeyAndMetaValue
func GetPostmetasByMetaIdAndMetaKeyAndMetaValue(offset int, limit int, MetaId_ int64, MetaKey_ string, MetaValue_ string) (*[]*Postmeta, error) {
	var _Postmeta = new([]*Postmeta)
	err := Engine.Table("postmeta").Where("meta_id = ? and meta_key = ? and meta_value = ?", MetaId_, MetaKey_, MetaValue_).Limit(limit, offset).Find(_Postmeta)
	return _Postmeta, err
}

// GetPostmetasByPostIdAndMetaKeyAndMetaValue Get Postmetas via PostIdAndMetaKeyAndMetaValue
func GetPostmetasByPostIdAndMetaKeyAndMetaValue(offset int, limit int, PostId_ int64, MetaKey_ string, MetaValue_ string) (*[]*Postmeta, error) {
	var _Postmeta = new([]*Postmeta)
	err := Engine.Table("postmeta").Where("post_id = ? and meta_key = ? and meta_value = ?", PostId_, MetaKey_, MetaValue_).Limit(limit, offset).Find(_Postmeta)
	return _Postmeta, err
}

// GetPostmetasByMetaIdAndPostId Get Postmetas via MetaIdAndPostId
func GetPostmetasByMetaIdAndPostId(offset int, limit int, MetaId_ int64, PostId_ int64) (*[]*Postmeta, error) {
	var _Postmeta = new([]*Postmeta)
	err := Engine.Table("postmeta").Where("meta_id = ? and post_id = ?", MetaId_, PostId_).Limit(limit, offset).Find(_Postmeta)
	return _Postmeta, err
}

// GetPostmetasByMetaIdAndMetaKey Get Postmetas via MetaIdAndMetaKey
func GetPostmetasByMetaIdAndMetaKey(offset int, limit int, MetaId_ int64, MetaKey_ string) (*[]*Postmeta, error) {
	var _Postmeta = new([]*Postmeta)
	err := Engine.Table("postmeta").Where("meta_id = ? and meta_key = ?", MetaId_, MetaKey_).Limit(limit, offset).Find(_Postmeta)
	return _Postmeta, err
}

// GetPostmetasByMetaIdAndMetaValue Get Postmetas via MetaIdAndMetaValue
func GetPostmetasByMetaIdAndMetaValue(offset int, limit int, MetaId_ int64, MetaValue_ string) (*[]*Postmeta, error) {
	var _Postmeta = new([]*Postmeta)
	err := Engine.Table("postmeta").Where("meta_id = ? and meta_value = ?", MetaId_, MetaValue_).Limit(limit, offset).Find(_Postmeta)
	return _Postmeta, err
}

// GetPostmetasByPostIdAndMetaKey Get Postmetas via PostIdAndMetaKey
func GetPostmetasByPostIdAndMetaKey(offset int, limit int, PostId_ int64, MetaKey_ string) (*[]*Postmeta, error) {
	var _Postmeta = new([]*Postmeta)
	err := Engine.Table("postmeta").Where("post_id = ? and meta_key = ?", PostId_, MetaKey_).Limit(limit, offset).Find(_Postmeta)
	return _Postmeta, err
}

// GetPostmetasByPostIdAndMetaValue Get Postmetas via PostIdAndMetaValue
func GetPostmetasByPostIdAndMetaValue(offset int, limit int, PostId_ int64, MetaValue_ string) (*[]*Postmeta, error) {
	var _Postmeta = new([]*Postmeta)
	err := Engine.Table("postmeta").Where("post_id = ? and meta_value = ?", PostId_, MetaValue_).Limit(limit, offset).Find(_Postmeta)
	return _Postmeta, err
}

// GetPostmetasByMetaKeyAndMetaValue Get Postmetas via MetaKeyAndMetaValue
func GetPostmetasByMetaKeyAndMetaValue(offset int, limit int, MetaKey_ string, MetaValue_ string) (*[]*Postmeta, error) {
	var _Postmeta = new([]*Postmeta)
	err := Engine.Table("postmeta").Where("meta_key = ? and meta_value = ?", MetaKey_, MetaValue_).Limit(limit, offset).Find(_Postmeta)
	return _Postmeta, err
}

// GetPostmetas Get Postmetas via field
func GetPostmetas(offset int, limit int, field string) (*[]*Postmeta, error) {
	var _Postmeta = new([]*Postmeta)
	err := Engine.Table("postmeta").Limit(limit, offset).Desc(field).Find(_Postmeta)
	return _Postmeta, err
}

// GetPostmetasByMetaId Get Postmetas via MetaId
func GetPostmetasByMetaId(offset int, limit int, MetaId_ int64, field string) (*[]*Postmeta, error) {
	var _Postmeta = new([]*Postmeta)
	err := Engine.Table("postmeta").Where("meta_id = ?", MetaId_).Limit(limit, offset).Desc(field).Find(_Postmeta)
	return _Postmeta, err
}

// GetPostmetasByPostId Get Postmetas via PostId
func GetPostmetasByPostId(offset int, limit int, PostId_ int64, field string) (*[]*Postmeta, error) {
	var _Postmeta = new([]*Postmeta)
	err := Engine.Table("postmeta").Where("post_id = ?", PostId_).Limit(limit, offset).Desc(field).Find(_Postmeta)
	return _Postmeta, err
}

// GetPostmetasByMetaKey Get Postmetas via MetaKey
func GetPostmetasByMetaKey(offset int, limit int, MetaKey_ string, field string) (*[]*Postmeta, error) {
	var _Postmeta = new([]*Postmeta)
	err := Engine.Table("postmeta").Where("meta_key = ?", MetaKey_).Limit(limit, offset).Desc(field).Find(_Postmeta)
	return _Postmeta, err
}

// GetPostmetasByMetaValue Get Postmetas via MetaValue
func GetPostmetasByMetaValue(offset int, limit int, MetaValue_ string, field string) (*[]*Postmeta, error) {
	var _Postmeta = new([]*Postmeta)
	err := Engine.Table("postmeta").Where("meta_value = ?", MetaValue_).Limit(limit, offset).Desc(field).Find(_Postmeta)
	return _Postmeta, err
}

// HasPostmetaByMetaId Has Postmeta via MetaId
func HasPostmetaByMetaId(iMetaId int64) bool {
	if has, err := Engine.Where("meta_id = ?", iMetaId).Get(new(Postmeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasPostmetaByPostId Has Postmeta via PostId
func HasPostmetaByPostId(iPostId int64) bool {
	if has, err := Engine.Where("post_id = ?", iPostId).Get(new(Postmeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasPostmetaByMetaKey Has Postmeta via MetaKey
func HasPostmetaByMetaKey(iMetaKey string) bool {
	if has, err := Engine.Where("meta_key = ?", iMetaKey).Get(new(Postmeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasPostmetaByMetaValue Has Postmeta via MetaValue
func HasPostmetaByMetaValue(iMetaValue string) bool {
	if has, err := Engine.Where("meta_value = ?", iMetaValue).Get(new(Postmeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetPostmetaByMetaId Get Postmeta via MetaId
func GetPostmetaByMetaId(iMetaId int64) (*Postmeta, error) {
	var _Postmeta = &Postmeta{MetaId: iMetaId}
	has, err := Engine.Get(_Postmeta)
	if has {
		return _Postmeta, err
	} else {
		return nil, err
	}
}

// GetPostmetaByPostId Get Postmeta via PostId
func GetPostmetaByPostId(iPostId int64) (*Postmeta, error) {
	var _Postmeta = &Postmeta{PostId: iPostId}
	has, err := Engine.Get(_Postmeta)
	if has {
		return _Postmeta, err
	} else {
		return nil, err
	}
}

// GetPostmetaByMetaKey Get Postmeta via MetaKey
func GetPostmetaByMetaKey(iMetaKey string) (*Postmeta, error) {
	var _Postmeta = &Postmeta{MetaKey: iMetaKey}
	has, err := Engine.Get(_Postmeta)
	if has {
		return _Postmeta, err
	} else {
		return nil, err
	}
}

// GetPostmetaByMetaValue Get Postmeta via MetaValue
func GetPostmetaByMetaValue(iMetaValue string) (*Postmeta, error) {
	var _Postmeta = &Postmeta{MetaValue: iMetaValue}
	has, err := Engine.Get(_Postmeta)
	if has {
		return _Postmeta, err
	} else {
		return nil, err
	}
}

// SetPostmetaByMetaId Set Postmeta via MetaId
func SetPostmetaByMetaId(iMetaId int64, postmeta *Postmeta) (int64, error) {
	postmeta.MetaId = iMetaId
	return Engine.Insert(postmeta)
}

// SetPostmetaByPostId Set Postmeta via PostId
func SetPostmetaByPostId(iPostId int64, postmeta *Postmeta) (int64, error) {
	postmeta.PostId = iPostId
	return Engine.Insert(postmeta)
}

// SetPostmetaByMetaKey Set Postmeta via MetaKey
func SetPostmetaByMetaKey(iMetaKey string, postmeta *Postmeta) (int64, error) {
	postmeta.MetaKey = iMetaKey
	return Engine.Insert(postmeta)
}

// SetPostmetaByMetaValue Set Postmeta via MetaValue
func SetPostmetaByMetaValue(iMetaValue string, postmeta *Postmeta) (int64, error) {
	postmeta.MetaValue = iMetaValue
	return Engine.Insert(postmeta)
}

// AddPostmeta Add Postmeta via all columns
func AddPostmeta(iMetaId int64, iPostId int64, iMetaKey string, iMetaValue string) error {
	_Postmeta := &Postmeta{MetaId: iMetaId, PostId: iPostId, MetaKey: iMetaKey, MetaValue: iMetaValue}
	if _, err := Engine.Insert(_Postmeta); err != nil {
		return err
	} else {
		return nil
	}
}

// PostPostmeta Post Postmeta via iPostmeta
func PostPostmeta(iPostmeta *Postmeta) (int64, error) {
	_, err := Engine.Insert(iPostmeta)
	return iPostmeta.MetaId, err
}

// PutPostmeta Put Postmeta
func PutPostmeta(iPostmeta *Postmeta) (int64, error) {
	_, err := Engine.Id(iPostmeta.MetaId).Update(iPostmeta)
	return iPostmeta.MetaId, err
}

// PutPostmetaByMetaId Put Postmeta via MetaId
func PutPostmetaByMetaId(MetaId_ int64, iPostmeta *Postmeta) (int64, error) {
	row, err := Engine.Update(iPostmeta, &Postmeta{MetaId: MetaId_})
	return row, err
}

// PutPostmetaByPostId Put Postmeta via PostId
func PutPostmetaByPostId(PostId_ int64, iPostmeta *Postmeta) (int64, error) {
	row, err := Engine.Update(iPostmeta, &Postmeta{PostId: PostId_})
	return row, err
}

// PutPostmetaByMetaKey Put Postmeta via MetaKey
func PutPostmetaByMetaKey(MetaKey_ string, iPostmeta *Postmeta) (int64, error) {
	row, err := Engine.Update(iPostmeta, &Postmeta{MetaKey: MetaKey_})
	return row, err
}

// PutPostmetaByMetaValue Put Postmeta via MetaValue
func PutPostmetaByMetaValue(MetaValue_ string, iPostmeta *Postmeta) (int64, error) {
	row, err := Engine.Update(iPostmeta, &Postmeta{MetaValue: MetaValue_})
	return row, err
}

// UpdatePostmetaByMetaId via map[string]interface{}{}
func UpdatePostmetaByMetaId(iMetaId int64, iPostmetaMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Postmeta)).Where("meta_id = ?", iMetaId).Update(iPostmetaMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdatePostmetaByPostId via map[string]interface{}{}
func UpdatePostmetaByPostId(iPostId int64, iPostmetaMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Postmeta)).Where("post_id = ?", iPostId).Update(iPostmetaMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdatePostmetaByMetaKey via map[string]interface{}{}
func UpdatePostmetaByMetaKey(iMetaKey string, iPostmetaMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Postmeta)).Where("meta_key = ?", iMetaKey).Update(iPostmetaMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdatePostmetaByMetaValue via map[string]interface{}{}
func UpdatePostmetaByMetaValue(iMetaValue string, iPostmetaMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Postmeta)).Where("meta_value = ?", iMetaValue).Update(iPostmetaMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeletePostmeta Delete Postmeta
func DeletePostmeta(iMetaId int64) (int64, error) {
	row, err := Engine.Id(iMetaId).Delete(new(Postmeta))
	return row, err
}

// DeletePostmetaByMetaId Delete Postmeta via MetaId
func DeletePostmetaByMetaId(iMetaId int64) (err error) {
	var has bool
	var _Postmeta = &Postmeta{MetaId: iMetaId}
	if has, err = Engine.Get(_Postmeta); (has == true) && (err == nil) {
		if row, err := Engine.Where("meta_id = ?", iMetaId).Delete(new(Postmeta)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeletePostmetaByPostId Delete Postmeta via PostId
func DeletePostmetaByPostId(iPostId int64) (err error) {
	var has bool
	var _Postmeta = &Postmeta{PostId: iPostId}
	if has, err = Engine.Get(_Postmeta); (has == true) && (err == nil) {
		if row, err := Engine.Where("post_id = ?", iPostId).Delete(new(Postmeta)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeletePostmetaByMetaKey Delete Postmeta via MetaKey
func DeletePostmetaByMetaKey(iMetaKey string) (err error) {
	var has bool
	var _Postmeta = &Postmeta{MetaKey: iMetaKey}
	if has, err = Engine.Get(_Postmeta); (has == true) && (err == nil) {
		if row, err := Engine.Where("meta_key = ?", iMetaKey).Delete(new(Postmeta)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeletePostmetaByMetaValue Delete Postmeta via MetaValue
func DeletePostmetaByMetaValue(iMetaValue string) (err error) {
	var has bool
	var _Postmeta = &Postmeta{MetaValue: iMetaValue}
	if has, err = Engine.Get(_Postmeta); (has == true) && (err == nil) {
		if row, err := Engine.Where("meta_value = ?", iMetaValue).Delete(new(Postmeta)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
