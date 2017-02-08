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

// GetPostmetaCountViaMetaId Get Postmeta via MetaId
func GetPostmetaCountViaMetaId(iMetaId int64) int64 {
	n, _ := Engine.Where("meta_id = ?", iMetaId).Count(&Postmeta{MetaId: iMetaId})
	return n
}

// GetPostmetaCountViaPostId Get Postmeta via PostId
func GetPostmetaCountViaPostId(iPostId int64) int64 {
	n, _ := Engine.Where("post_id = ?", iPostId).Count(&Postmeta{PostId: iPostId})
	return n
}

// GetPostmetaCountViaMetaKey Get Postmeta via MetaKey
func GetPostmetaCountViaMetaKey(iMetaKey string) int64 {
	n, _ := Engine.Where("meta_key = ?", iMetaKey).Count(&Postmeta{MetaKey: iMetaKey})
	return n
}

// GetPostmetaCountViaMetaValue Get Postmeta via MetaValue
func GetPostmetaCountViaMetaValue(iMetaValue string) int64 {
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

// GetPostmetasViaMetaId Get Postmetas via MetaId
func GetPostmetasViaMetaId(offset int, limit int, MetaId_ int64, field string) (*[]*Postmeta, error) {
	var _Postmeta = new([]*Postmeta)
	err := Engine.Table("postmeta").Where("meta_id = ?", MetaId_).Limit(limit, offset).Desc(field).Find(_Postmeta)
	return _Postmeta, err
}

// GetPostmetasViaPostId Get Postmetas via PostId
func GetPostmetasViaPostId(offset int, limit int, PostId_ int64, field string) (*[]*Postmeta, error) {
	var _Postmeta = new([]*Postmeta)
	err := Engine.Table("postmeta").Where("post_id = ?", PostId_).Limit(limit, offset).Desc(field).Find(_Postmeta)
	return _Postmeta, err
}

// GetPostmetasViaMetaKey Get Postmetas via MetaKey
func GetPostmetasViaMetaKey(offset int, limit int, MetaKey_ string, field string) (*[]*Postmeta, error) {
	var _Postmeta = new([]*Postmeta)
	err := Engine.Table("postmeta").Where("meta_key = ?", MetaKey_).Limit(limit, offset).Desc(field).Find(_Postmeta)
	return _Postmeta, err
}

// GetPostmetasViaMetaValue Get Postmetas via MetaValue
func GetPostmetasViaMetaValue(offset int, limit int, MetaValue_ string, field string) (*[]*Postmeta, error) {
	var _Postmeta = new([]*Postmeta)
	err := Engine.Table("postmeta").Where("meta_value = ?", MetaValue_).Limit(limit, offset).Desc(field).Find(_Postmeta)
	return _Postmeta, err
}

// HasPostmetaViaMetaId Has Postmeta via MetaId
func HasPostmetaViaMetaId(iMetaId int64) bool {
	if has, err := Engine.Where("meta_id = ?", iMetaId).Get(new(Postmeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasPostmetaViaPostId Has Postmeta via PostId
func HasPostmetaViaPostId(iPostId int64) bool {
	if has, err := Engine.Where("post_id = ?", iPostId).Get(new(Postmeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasPostmetaViaMetaKey Has Postmeta via MetaKey
func HasPostmetaViaMetaKey(iMetaKey string) bool {
	if has, err := Engine.Where("meta_key = ?", iMetaKey).Get(new(Postmeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasPostmetaViaMetaValue Has Postmeta via MetaValue
func HasPostmetaViaMetaValue(iMetaValue string) bool {
	if has, err := Engine.Where("meta_value = ?", iMetaValue).Get(new(Postmeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetPostmetaViaMetaId Get Postmeta via MetaId
func GetPostmetaViaMetaId(iMetaId int64) (*Postmeta, error) {
	var _Postmeta = &Postmeta{MetaId: iMetaId}
	has, err := Engine.Get(_Postmeta)
	if has {
		return _Postmeta, err
	} else {
		return nil, err
	}
}

// GetPostmetaViaPostId Get Postmeta via PostId
func GetPostmetaViaPostId(iPostId int64) (*Postmeta, error) {
	var _Postmeta = &Postmeta{PostId: iPostId}
	has, err := Engine.Get(_Postmeta)
	if has {
		return _Postmeta, err
	} else {
		return nil, err
	}
}

// GetPostmetaViaMetaKey Get Postmeta via MetaKey
func GetPostmetaViaMetaKey(iMetaKey string) (*Postmeta, error) {
	var _Postmeta = &Postmeta{MetaKey: iMetaKey}
	has, err := Engine.Get(_Postmeta)
	if has {
		return _Postmeta, err
	} else {
		return nil, err
	}
}

// GetPostmetaViaMetaValue Get Postmeta via MetaValue
func GetPostmetaViaMetaValue(iMetaValue string) (*Postmeta, error) {
	var _Postmeta = &Postmeta{MetaValue: iMetaValue}
	has, err := Engine.Get(_Postmeta)
	if has {
		return _Postmeta, err
	} else {
		return nil, err
	}
}

// SetPostmetaViaMetaId Set Postmeta via MetaId
func SetPostmetaViaMetaId(iMetaId int64, postmeta *Postmeta) (int64, error) {
	postmeta.MetaId = iMetaId
	return Engine.Insert(postmeta)
}

// SetPostmetaViaPostId Set Postmeta via PostId
func SetPostmetaViaPostId(iPostId int64, postmeta *Postmeta) (int64, error) {
	postmeta.PostId = iPostId
	return Engine.Insert(postmeta)
}

// SetPostmetaViaMetaKey Set Postmeta via MetaKey
func SetPostmetaViaMetaKey(iMetaKey string, postmeta *Postmeta) (int64, error) {
	postmeta.MetaKey = iMetaKey
	return Engine.Insert(postmeta)
}

// SetPostmetaViaMetaValue Set Postmeta via MetaValue
func SetPostmetaViaMetaValue(iMetaValue string, postmeta *Postmeta) (int64, error) {
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

// PutPostmetaViaMetaId Put Postmeta via MetaId
func PutPostmetaViaMetaId(MetaId_ int64, iPostmeta *Postmeta) (int64, error) {
	row, err := Engine.Update(iPostmeta, &Postmeta{MetaId: MetaId_})
	return row, err
}

// PutPostmetaViaPostId Put Postmeta via PostId
func PutPostmetaViaPostId(PostId_ int64, iPostmeta *Postmeta) (int64, error) {
	row, err := Engine.Update(iPostmeta, &Postmeta{PostId: PostId_})
	return row, err
}

// PutPostmetaViaMetaKey Put Postmeta via MetaKey
func PutPostmetaViaMetaKey(MetaKey_ string, iPostmeta *Postmeta) (int64, error) {
	row, err := Engine.Update(iPostmeta, &Postmeta{MetaKey: MetaKey_})
	return row, err
}

// PutPostmetaViaMetaValue Put Postmeta via MetaValue
func PutPostmetaViaMetaValue(MetaValue_ string, iPostmeta *Postmeta) (int64, error) {
	row, err := Engine.Update(iPostmeta, &Postmeta{MetaValue: MetaValue_})
	return row, err
}

// UpdatePostmetaViaMetaId via map[string]interface{}{}
func UpdatePostmetaViaMetaId(iMetaId int64, iPostmetaMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Postmeta)).Where("meta_id = ?", iMetaId).Update(iPostmetaMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdatePostmetaViaPostId via map[string]interface{}{}
func UpdatePostmetaViaPostId(iPostId int64, iPostmetaMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Postmeta)).Where("post_id = ?", iPostId).Update(iPostmetaMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdatePostmetaViaMetaKey via map[string]interface{}{}
func UpdatePostmetaViaMetaKey(iMetaKey string, iPostmetaMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Postmeta)).Where("meta_key = ?", iMetaKey).Update(iPostmetaMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdatePostmetaViaMetaValue via map[string]interface{}{}
func UpdatePostmetaViaMetaValue(iMetaValue string, iPostmetaMap *map[string]interface{}) error {
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

// DeletePostmetaViaMetaId Delete Postmeta via MetaId
func DeletePostmetaViaMetaId(iMetaId int64) (err error) {
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

// DeletePostmetaViaPostId Delete Postmeta via PostId
func DeletePostmetaViaPostId(iPostId int64) (err error) {
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

// DeletePostmetaViaMetaKey Delete Postmeta via MetaKey
func DeletePostmetaViaMetaKey(iMetaKey string) (err error) {
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

// DeletePostmetaViaMetaValue Delete Postmeta via MetaValue
func DeletePostmetaViaMetaValue(iMetaValue string) (err error) {
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
