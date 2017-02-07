package model

type Commentmeta struct {
	MetaId    int64  `xorm:"not null pk autoincr BIGINT(20)"`
	CommentId int64  `xorm:"not null default 0 index BIGINT(20)"`
	MetaKey   string `xorm:"index VARCHAR(255)"`
	MetaValue string `xorm:"LONGTEXT"`
}

// GetCommentmetasCount Commentmetas' Count
func GetCommentmetasCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&Commentmeta{})
	return total, err
}

// GetCommentmetaCountByMetaId Get Commentmeta via MetaId
func GetCommentmetaCountByMetaId(iMetaId int64) int64 {
	n, _ := Engine.Where("meta_id = ?", iMetaId).Count(&Commentmeta{MetaId: iMetaId})
	return n
}

// GetCommentmetaCountByCommentId Get Commentmeta via CommentId
func GetCommentmetaCountByCommentId(iCommentId int64) int64 {
	n, _ := Engine.Where("comment_id = ?", iCommentId).Count(&Commentmeta{CommentId: iCommentId})
	return n
}

// GetCommentmetaCountByMetaKey Get Commentmeta via MetaKey
func GetCommentmetaCountByMetaKey(iMetaKey string) int64 {
	n, _ := Engine.Where("meta_key = ?", iMetaKey).Count(&Commentmeta{MetaKey: iMetaKey})
	return n
}

// GetCommentmetaCountByMetaValue Get Commentmeta via MetaValue
func GetCommentmetaCountByMetaValue(iMetaValue string) int64 {
	n, _ := Engine.Where("meta_value = ?", iMetaValue).Count(&Commentmeta{MetaValue: iMetaValue})
	return n
}

// GetCommentmetasByMetaIdAndCommentIdAndMetaKey Get Commentmetas via MetaIdAndCommentIdAndMetaKey
func GetCommentmetasByMetaIdAndCommentIdAndMetaKey(offset int, limit int, MetaId_ int64, CommentId_ int64, MetaKey_ string) (*[]*Commentmeta, error) {
	var _Commentmeta = new([]*Commentmeta)
	err := Engine.Table("commentmeta").Where("meta_id = ? and comment_id = ? and meta_key = ?", MetaId_, CommentId_, MetaKey_).Limit(limit, offset).Find(_Commentmeta)
	return _Commentmeta, err
}

// GetCommentmetasByMetaIdAndCommentIdAndMetaValue Get Commentmetas via MetaIdAndCommentIdAndMetaValue
func GetCommentmetasByMetaIdAndCommentIdAndMetaValue(offset int, limit int, MetaId_ int64, CommentId_ int64, MetaValue_ string) (*[]*Commentmeta, error) {
	var _Commentmeta = new([]*Commentmeta)
	err := Engine.Table("commentmeta").Where("meta_id = ? and comment_id = ? and meta_value = ?", MetaId_, CommentId_, MetaValue_).Limit(limit, offset).Find(_Commentmeta)
	return _Commentmeta, err
}

// GetCommentmetasByMetaIdAndMetaKeyAndMetaValue Get Commentmetas via MetaIdAndMetaKeyAndMetaValue
func GetCommentmetasByMetaIdAndMetaKeyAndMetaValue(offset int, limit int, MetaId_ int64, MetaKey_ string, MetaValue_ string) (*[]*Commentmeta, error) {
	var _Commentmeta = new([]*Commentmeta)
	err := Engine.Table("commentmeta").Where("meta_id = ? and meta_key = ? and meta_value = ?", MetaId_, MetaKey_, MetaValue_).Limit(limit, offset).Find(_Commentmeta)
	return _Commentmeta, err
}

// GetCommentmetasByCommentIdAndMetaKeyAndMetaValue Get Commentmetas via CommentIdAndMetaKeyAndMetaValue
func GetCommentmetasByCommentIdAndMetaKeyAndMetaValue(offset int, limit int, CommentId_ int64, MetaKey_ string, MetaValue_ string) (*[]*Commentmeta, error) {
	var _Commentmeta = new([]*Commentmeta)
	err := Engine.Table("commentmeta").Where("comment_id = ? and meta_key = ? and meta_value = ?", CommentId_, MetaKey_, MetaValue_).Limit(limit, offset).Find(_Commentmeta)
	return _Commentmeta, err
}

// GetCommentmetasByMetaIdAndCommentId Get Commentmetas via MetaIdAndCommentId
func GetCommentmetasByMetaIdAndCommentId(offset int, limit int, MetaId_ int64, CommentId_ int64) (*[]*Commentmeta, error) {
	var _Commentmeta = new([]*Commentmeta)
	err := Engine.Table("commentmeta").Where("meta_id = ? and comment_id = ?", MetaId_, CommentId_).Limit(limit, offset).Find(_Commentmeta)
	return _Commentmeta, err
}

// GetCommentmetasByMetaIdAndMetaKey Get Commentmetas via MetaIdAndMetaKey
func GetCommentmetasByMetaIdAndMetaKey(offset int, limit int, MetaId_ int64, MetaKey_ string) (*[]*Commentmeta, error) {
	var _Commentmeta = new([]*Commentmeta)
	err := Engine.Table("commentmeta").Where("meta_id = ? and meta_key = ?", MetaId_, MetaKey_).Limit(limit, offset).Find(_Commentmeta)
	return _Commentmeta, err
}

// GetCommentmetasByMetaIdAndMetaValue Get Commentmetas via MetaIdAndMetaValue
func GetCommentmetasByMetaIdAndMetaValue(offset int, limit int, MetaId_ int64, MetaValue_ string) (*[]*Commentmeta, error) {
	var _Commentmeta = new([]*Commentmeta)
	err := Engine.Table("commentmeta").Where("meta_id = ? and meta_value = ?", MetaId_, MetaValue_).Limit(limit, offset).Find(_Commentmeta)
	return _Commentmeta, err
}

// GetCommentmetasByCommentIdAndMetaKey Get Commentmetas via CommentIdAndMetaKey
func GetCommentmetasByCommentIdAndMetaKey(offset int, limit int, CommentId_ int64, MetaKey_ string) (*[]*Commentmeta, error) {
	var _Commentmeta = new([]*Commentmeta)
	err := Engine.Table("commentmeta").Where("comment_id = ? and meta_key = ?", CommentId_, MetaKey_).Limit(limit, offset).Find(_Commentmeta)
	return _Commentmeta, err
}

// GetCommentmetasByCommentIdAndMetaValue Get Commentmetas via CommentIdAndMetaValue
func GetCommentmetasByCommentIdAndMetaValue(offset int, limit int, CommentId_ int64, MetaValue_ string) (*[]*Commentmeta, error) {
	var _Commentmeta = new([]*Commentmeta)
	err := Engine.Table("commentmeta").Where("comment_id = ? and meta_value = ?", CommentId_, MetaValue_).Limit(limit, offset).Find(_Commentmeta)
	return _Commentmeta, err
}

// GetCommentmetasByMetaKeyAndMetaValue Get Commentmetas via MetaKeyAndMetaValue
func GetCommentmetasByMetaKeyAndMetaValue(offset int, limit int, MetaKey_ string, MetaValue_ string) (*[]*Commentmeta, error) {
	var _Commentmeta = new([]*Commentmeta)
	err := Engine.Table("commentmeta").Where("meta_key = ? and meta_value = ?", MetaKey_, MetaValue_).Limit(limit, offset).Find(_Commentmeta)
	return _Commentmeta, err
}

// GetCommentmetas Get Commentmetas via field
func GetCommentmetas(offset int, limit int, field string) (*[]*Commentmeta, error) {
	var _Commentmeta = new([]*Commentmeta)
	err := Engine.Table("commentmeta").Limit(limit, offset).Desc(field).Find(_Commentmeta)
	return _Commentmeta, err
}

// GetCommentmetasByMetaId Get Commentmetas via MetaId
func GetCommentmetasByMetaId(offset int, limit int, MetaId_ int64, field string) (*[]*Commentmeta, error) {
	var _Commentmeta = new([]*Commentmeta)
	err := Engine.Table("commentmeta").Where("meta_id = ?", MetaId_).Limit(limit, offset).Desc(field).Find(_Commentmeta)
	return _Commentmeta, err
}

// GetCommentmetasByCommentId Get Commentmetas via CommentId
func GetCommentmetasByCommentId(offset int, limit int, CommentId_ int64, field string) (*[]*Commentmeta, error) {
	var _Commentmeta = new([]*Commentmeta)
	err := Engine.Table("commentmeta").Where("comment_id = ?", CommentId_).Limit(limit, offset).Desc(field).Find(_Commentmeta)
	return _Commentmeta, err
}

// GetCommentmetasByMetaKey Get Commentmetas via MetaKey
func GetCommentmetasByMetaKey(offset int, limit int, MetaKey_ string, field string) (*[]*Commentmeta, error) {
	var _Commentmeta = new([]*Commentmeta)
	err := Engine.Table("commentmeta").Where("meta_key = ?", MetaKey_).Limit(limit, offset).Desc(field).Find(_Commentmeta)
	return _Commentmeta, err
}

// GetCommentmetasByMetaValue Get Commentmetas via MetaValue
func GetCommentmetasByMetaValue(offset int, limit int, MetaValue_ string, field string) (*[]*Commentmeta, error) {
	var _Commentmeta = new([]*Commentmeta)
	err := Engine.Table("commentmeta").Where("meta_value = ?", MetaValue_).Limit(limit, offset).Desc(field).Find(_Commentmeta)
	return _Commentmeta, err
}

// HasCommentmetaByMetaId Has Commentmeta via MetaId
func HasCommentmetaByMetaId(iMetaId int64) bool {
	if has, err := Engine.Where("meta_id = ?", iMetaId).Get(new(Commentmeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentmetaByCommentId Has Commentmeta via CommentId
func HasCommentmetaByCommentId(iCommentId int64) bool {
	if has, err := Engine.Where("comment_id = ?", iCommentId).Get(new(Commentmeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentmetaByMetaKey Has Commentmeta via MetaKey
func HasCommentmetaByMetaKey(iMetaKey string) bool {
	if has, err := Engine.Where("meta_key = ?", iMetaKey).Get(new(Commentmeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentmetaByMetaValue Has Commentmeta via MetaValue
func HasCommentmetaByMetaValue(iMetaValue string) bool {
	if has, err := Engine.Where("meta_value = ?", iMetaValue).Get(new(Commentmeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetCommentmetaByMetaId Get Commentmeta via MetaId
func GetCommentmetaByMetaId(iMetaId int64) (*Commentmeta, error) {
	var _Commentmeta = &Commentmeta{MetaId: iMetaId}
	has, err := Engine.Get(_Commentmeta)
	if has {
		return _Commentmeta, err
	} else {
		return nil, err
	}
}

// GetCommentmetaByCommentId Get Commentmeta via CommentId
func GetCommentmetaByCommentId(iCommentId int64) (*Commentmeta, error) {
	var _Commentmeta = &Commentmeta{CommentId: iCommentId}
	has, err := Engine.Get(_Commentmeta)
	if has {
		return _Commentmeta, err
	} else {
		return nil, err
	}
}

// GetCommentmetaByMetaKey Get Commentmeta via MetaKey
func GetCommentmetaByMetaKey(iMetaKey string) (*Commentmeta, error) {
	var _Commentmeta = &Commentmeta{MetaKey: iMetaKey}
	has, err := Engine.Get(_Commentmeta)
	if has {
		return _Commentmeta, err
	} else {
		return nil, err
	}
}

// GetCommentmetaByMetaValue Get Commentmeta via MetaValue
func GetCommentmetaByMetaValue(iMetaValue string) (*Commentmeta, error) {
	var _Commentmeta = &Commentmeta{MetaValue: iMetaValue}
	has, err := Engine.Get(_Commentmeta)
	if has {
		return _Commentmeta, err
	} else {
		return nil, err
	}
}

// SetCommentmetaByMetaId Set Commentmeta via MetaId
func SetCommentmetaByMetaId(iMetaId int64, commentmeta *Commentmeta) (int64, error) {
	commentmeta.MetaId = iMetaId
	return Engine.Insert(commentmeta)
}

// SetCommentmetaByCommentId Set Commentmeta via CommentId
func SetCommentmetaByCommentId(iCommentId int64, commentmeta *Commentmeta) (int64, error) {
	commentmeta.CommentId = iCommentId
	return Engine.Insert(commentmeta)
}

// SetCommentmetaByMetaKey Set Commentmeta via MetaKey
func SetCommentmetaByMetaKey(iMetaKey string, commentmeta *Commentmeta) (int64, error) {
	commentmeta.MetaKey = iMetaKey
	return Engine.Insert(commentmeta)
}

// SetCommentmetaByMetaValue Set Commentmeta via MetaValue
func SetCommentmetaByMetaValue(iMetaValue string, commentmeta *Commentmeta) (int64, error) {
	commentmeta.MetaValue = iMetaValue
	return Engine.Insert(commentmeta)
}

// AddCommentmeta Add Commentmeta via all columns
func AddCommentmeta(iMetaId int64, iCommentId int64, iMetaKey string, iMetaValue string) error {
	_Commentmeta := &Commentmeta{MetaId: iMetaId, CommentId: iCommentId, MetaKey: iMetaKey, MetaValue: iMetaValue}
	if _, err := Engine.Insert(_Commentmeta); err != nil {
		return err
	} else {
		return nil
	}
}

// PostCommentmeta Post Commentmeta via iCommentmeta
func PostCommentmeta(iCommentmeta *Commentmeta) (int64, error) {
	_, err := Engine.Insert(iCommentmeta)
	return iCommentmeta.MetaId, err
}

// PutCommentmeta Put Commentmeta
func PutCommentmeta(iCommentmeta *Commentmeta) (int64, error) {
	_, err := Engine.Id(iCommentmeta.MetaId).Update(iCommentmeta)
	return iCommentmeta.MetaId, err
}

// PutCommentmetaByMetaId Put Commentmeta via MetaId
func PutCommentmetaByMetaId(MetaId_ int64, iCommentmeta *Commentmeta) (int64, error) {
	row, err := Engine.Update(iCommentmeta, &Commentmeta{MetaId: MetaId_})
	return row, err
}

// PutCommentmetaByCommentId Put Commentmeta via CommentId
func PutCommentmetaByCommentId(CommentId_ int64, iCommentmeta *Commentmeta) (int64, error) {
	row, err := Engine.Update(iCommentmeta, &Commentmeta{CommentId: CommentId_})
	return row, err
}

// PutCommentmetaByMetaKey Put Commentmeta via MetaKey
func PutCommentmetaByMetaKey(MetaKey_ string, iCommentmeta *Commentmeta) (int64, error) {
	row, err := Engine.Update(iCommentmeta, &Commentmeta{MetaKey: MetaKey_})
	return row, err
}

// PutCommentmetaByMetaValue Put Commentmeta via MetaValue
func PutCommentmetaByMetaValue(MetaValue_ string, iCommentmeta *Commentmeta) (int64, error) {
	row, err := Engine.Update(iCommentmeta, &Commentmeta{MetaValue: MetaValue_})
	return row, err
}

// UpdateCommentmetaByMetaId via map[string]interface{}{}
func UpdateCommentmetaByMetaId(iMetaId int64, iCommentmetaMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Commentmeta)).Where("meta_id = ?", iMetaId).Update(iCommentmetaMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentmetaByCommentId via map[string]interface{}{}
func UpdateCommentmetaByCommentId(iCommentId int64, iCommentmetaMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Commentmeta)).Where("comment_id = ?", iCommentId).Update(iCommentmetaMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentmetaByMetaKey via map[string]interface{}{}
func UpdateCommentmetaByMetaKey(iMetaKey string, iCommentmetaMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Commentmeta)).Where("meta_key = ?", iMetaKey).Update(iCommentmetaMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentmetaByMetaValue via map[string]interface{}{}
func UpdateCommentmetaByMetaValue(iMetaValue string, iCommentmetaMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Commentmeta)).Where("meta_value = ?", iMetaValue).Update(iCommentmetaMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteCommentmeta Delete Commentmeta
func DeleteCommentmeta(iMetaId int64) (int64, error) {
	row, err := Engine.Id(iMetaId).Delete(new(Commentmeta))
	return row, err
}

// DeleteCommentmetaByMetaId Delete Commentmeta via MetaId
func DeleteCommentmetaByMetaId(iMetaId int64) (err error) {
	var has bool
	var _Commentmeta = &Commentmeta{MetaId: iMetaId}
	if has, err = Engine.Get(_Commentmeta); (has == true) && (err == nil) {
		if row, err := Engine.Where("meta_id = ?", iMetaId).Delete(new(Commentmeta)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentmetaByCommentId Delete Commentmeta via CommentId
func DeleteCommentmetaByCommentId(iCommentId int64) (err error) {
	var has bool
	var _Commentmeta = &Commentmeta{CommentId: iCommentId}
	if has, err = Engine.Get(_Commentmeta); (has == true) && (err == nil) {
		if row, err := Engine.Where("comment_id = ?", iCommentId).Delete(new(Commentmeta)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentmetaByMetaKey Delete Commentmeta via MetaKey
func DeleteCommentmetaByMetaKey(iMetaKey string) (err error) {
	var has bool
	var _Commentmeta = &Commentmeta{MetaKey: iMetaKey}
	if has, err = Engine.Get(_Commentmeta); (has == true) && (err == nil) {
		if row, err := Engine.Where("meta_key = ?", iMetaKey).Delete(new(Commentmeta)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentmetaByMetaValue Delete Commentmeta via MetaValue
func DeleteCommentmetaByMetaValue(iMetaValue string) (err error) {
	var has bool
	var _Commentmeta = &Commentmeta{MetaValue: iMetaValue}
	if has, err = Engine.Get(_Commentmeta); (has == true) && (err == nil) {
		if row, err := Engine.Where("meta_value = ?", iMetaValue).Delete(new(Commentmeta)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
