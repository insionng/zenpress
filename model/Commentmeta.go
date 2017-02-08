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

// GetCommentmetaCountViaMetaId Get Commentmeta via MetaId
func GetCommentmetaCountViaMetaId(iMetaId int64) int64 {
	n, _ := Engine.Where("meta_id = ?", iMetaId).Count(&Commentmeta{MetaId: iMetaId})
	return n
}

// GetCommentmetaCountViaCommentId Get Commentmeta via CommentId
func GetCommentmetaCountViaCommentId(iCommentId int64) int64 {
	n, _ := Engine.Where("comment_id = ?", iCommentId).Count(&Commentmeta{CommentId: iCommentId})
	return n
}

// GetCommentmetaCountViaMetaKey Get Commentmeta via MetaKey
func GetCommentmetaCountViaMetaKey(iMetaKey string) int64 {
	n, _ := Engine.Where("meta_key = ?", iMetaKey).Count(&Commentmeta{MetaKey: iMetaKey})
	return n
}

// GetCommentmetaCountViaMetaValue Get Commentmeta via MetaValue
func GetCommentmetaCountViaMetaValue(iMetaValue string) int64 {
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

// GetCommentmetasViaMetaId Get Commentmetas via MetaId
func GetCommentmetasViaMetaId(offset int, limit int, MetaId_ int64, field string) (*[]*Commentmeta, error) {
	var _Commentmeta = new([]*Commentmeta)
	err := Engine.Table("commentmeta").Where("meta_id = ?", MetaId_).Limit(limit, offset).Desc(field).Find(_Commentmeta)
	return _Commentmeta, err
}

// GetCommentmetasViaCommentId Get Commentmetas via CommentId
func GetCommentmetasViaCommentId(offset int, limit int, CommentId_ int64, field string) (*[]*Commentmeta, error) {
	var _Commentmeta = new([]*Commentmeta)
	err := Engine.Table("commentmeta").Where("comment_id = ?", CommentId_).Limit(limit, offset).Desc(field).Find(_Commentmeta)
	return _Commentmeta, err
}

// GetCommentmetasViaMetaKey Get Commentmetas via MetaKey
func GetCommentmetasViaMetaKey(offset int, limit int, MetaKey_ string, field string) (*[]*Commentmeta, error) {
	var _Commentmeta = new([]*Commentmeta)
	err := Engine.Table("commentmeta").Where("meta_key = ?", MetaKey_).Limit(limit, offset).Desc(field).Find(_Commentmeta)
	return _Commentmeta, err
}

// GetCommentmetasViaMetaValue Get Commentmetas via MetaValue
func GetCommentmetasViaMetaValue(offset int, limit int, MetaValue_ string, field string) (*[]*Commentmeta, error) {
	var _Commentmeta = new([]*Commentmeta)
	err := Engine.Table("commentmeta").Where("meta_value = ?", MetaValue_).Limit(limit, offset).Desc(field).Find(_Commentmeta)
	return _Commentmeta, err
}

// HasCommentmetaViaMetaId Has Commentmeta via MetaId
func HasCommentmetaViaMetaId(iMetaId int64) bool {
	if has, err := Engine.Where("meta_id = ?", iMetaId).Get(new(Commentmeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentmetaViaCommentId Has Commentmeta via CommentId
func HasCommentmetaViaCommentId(iCommentId int64) bool {
	if has, err := Engine.Where("comment_id = ?", iCommentId).Get(new(Commentmeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentmetaViaMetaKey Has Commentmeta via MetaKey
func HasCommentmetaViaMetaKey(iMetaKey string) bool {
	if has, err := Engine.Where("meta_key = ?", iMetaKey).Get(new(Commentmeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentmetaViaMetaValue Has Commentmeta via MetaValue
func HasCommentmetaViaMetaValue(iMetaValue string) bool {
	if has, err := Engine.Where("meta_value = ?", iMetaValue).Get(new(Commentmeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetCommentmetaViaMetaId Get Commentmeta via MetaId
func GetCommentmetaViaMetaId(iMetaId int64) (*Commentmeta, error) {
	var _Commentmeta = &Commentmeta{MetaId: iMetaId}
	has, err := Engine.Get(_Commentmeta)
	if has {
		return _Commentmeta, err
	} else {
		return nil, err
	}
}

// GetCommentmetaViaCommentId Get Commentmeta via CommentId
func GetCommentmetaViaCommentId(iCommentId int64) (*Commentmeta, error) {
	var _Commentmeta = &Commentmeta{CommentId: iCommentId}
	has, err := Engine.Get(_Commentmeta)
	if has {
		return _Commentmeta, err
	} else {
		return nil, err
	}
}

// GetCommentmetaViaMetaKey Get Commentmeta via MetaKey
func GetCommentmetaViaMetaKey(iMetaKey string) (*Commentmeta, error) {
	var _Commentmeta = &Commentmeta{MetaKey: iMetaKey}
	has, err := Engine.Get(_Commentmeta)
	if has {
		return _Commentmeta, err
	} else {
		return nil, err
	}
}

// GetCommentmetaViaMetaValue Get Commentmeta via MetaValue
func GetCommentmetaViaMetaValue(iMetaValue string) (*Commentmeta, error) {
	var _Commentmeta = &Commentmeta{MetaValue: iMetaValue}
	has, err := Engine.Get(_Commentmeta)
	if has {
		return _Commentmeta, err
	} else {
		return nil, err
	}
}

// SetCommentmetaViaMetaId Set Commentmeta via MetaId
func SetCommentmetaViaMetaId(iMetaId int64, commentmeta *Commentmeta) (int64, error) {
	commentmeta.MetaId = iMetaId
	return Engine.Insert(commentmeta)
}

// SetCommentmetaViaCommentId Set Commentmeta via CommentId
func SetCommentmetaViaCommentId(iCommentId int64, commentmeta *Commentmeta) (int64, error) {
	commentmeta.CommentId = iCommentId
	return Engine.Insert(commentmeta)
}

// SetCommentmetaViaMetaKey Set Commentmeta via MetaKey
func SetCommentmetaViaMetaKey(iMetaKey string, commentmeta *Commentmeta) (int64, error) {
	commentmeta.MetaKey = iMetaKey
	return Engine.Insert(commentmeta)
}

// SetCommentmetaViaMetaValue Set Commentmeta via MetaValue
func SetCommentmetaViaMetaValue(iMetaValue string, commentmeta *Commentmeta) (int64, error) {
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

// PutCommentmetaViaMetaId Put Commentmeta via MetaId
func PutCommentmetaViaMetaId(MetaId_ int64, iCommentmeta *Commentmeta) (int64, error) {
	row, err := Engine.Update(iCommentmeta, &Commentmeta{MetaId: MetaId_})
	return row, err
}

// PutCommentmetaViaCommentId Put Commentmeta via CommentId
func PutCommentmetaViaCommentId(CommentId_ int64, iCommentmeta *Commentmeta) (int64, error) {
	row, err := Engine.Update(iCommentmeta, &Commentmeta{CommentId: CommentId_})
	return row, err
}

// PutCommentmetaViaMetaKey Put Commentmeta via MetaKey
func PutCommentmetaViaMetaKey(MetaKey_ string, iCommentmeta *Commentmeta) (int64, error) {
	row, err := Engine.Update(iCommentmeta, &Commentmeta{MetaKey: MetaKey_})
	return row, err
}

// PutCommentmetaViaMetaValue Put Commentmeta via MetaValue
func PutCommentmetaViaMetaValue(MetaValue_ string, iCommentmeta *Commentmeta) (int64, error) {
	row, err := Engine.Update(iCommentmeta, &Commentmeta{MetaValue: MetaValue_})
	return row, err
}

// UpdateCommentmetaViaMetaId via map[string]interface{}{}
func UpdateCommentmetaViaMetaId(iMetaId int64, iCommentmetaMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Commentmeta)).Where("meta_id = ?", iMetaId).Update(iCommentmetaMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentmetaViaCommentId via map[string]interface{}{}
func UpdateCommentmetaViaCommentId(iCommentId int64, iCommentmetaMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Commentmeta)).Where("comment_id = ?", iCommentId).Update(iCommentmetaMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentmetaViaMetaKey via map[string]interface{}{}
func UpdateCommentmetaViaMetaKey(iMetaKey string, iCommentmetaMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Commentmeta)).Where("meta_key = ?", iMetaKey).Update(iCommentmetaMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentmetaViaMetaValue via map[string]interface{}{}
func UpdateCommentmetaViaMetaValue(iMetaValue string, iCommentmetaMap *map[string]interface{}) error {
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

// DeleteCommentmetaViaMetaId Delete Commentmeta via MetaId
func DeleteCommentmetaViaMetaId(iMetaId int64) (err error) {
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

// DeleteCommentmetaViaCommentId Delete Commentmeta via CommentId
func DeleteCommentmetaViaCommentId(iCommentId int64) (err error) {
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

// DeleteCommentmetaViaMetaKey Delete Commentmeta via MetaKey
func DeleteCommentmetaViaMetaKey(iMetaKey string) (err error) {
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

// DeleteCommentmetaViaMetaValue Delete Commentmeta via MetaValue
func DeleteCommentmetaViaMetaValue(iMetaValue string) (err error) {
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
