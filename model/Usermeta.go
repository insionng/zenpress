package model

type Usermeta struct {
	UmetaId   int64  `xorm:"not null pk autoincr BIGINT(20)"`
	UserId    int64  `xorm:"not null default 0 index BIGINT(20)"`
	MetaKey   string `xorm:"index VARCHAR(255)"`
	MetaValue string `xorm:"LONGTEXT"`
}

// GetUsermetasCount Usermetas' Count
func GetUsermetasCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&Usermeta{})
	return total, err
}

// GetUsermetaCountByUmetaId Get Usermeta via UmetaId
func GetUsermetaCountByUmetaId(iUmetaId int64) int64 {
	n, _ := Engine.Where("umeta_id = ?", iUmetaId).Count(&Usermeta{UmetaId: iUmetaId})
	return n
}

// GetUsermetaCountByUserId Get Usermeta via UserId
func GetUsermetaCountByUserId(iUserId int64) int64 {
	n, _ := Engine.Where("user_id = ?", iUserId).Count(&Usermeta{UserId: iUserId})
	return n
}

// GetUsermetaCountByMetaKey Get Usermeta via MetaKey
func GetUsermetaCountByMetaKey(iMetaKey string) int64 {
	n, _ := Engine.Where("meta_key = ?", iMetaKey).Count(&Usermeta{MetaKey: iMetaKey})
	return n
}

// GetUsermetaCountByMetaValue Get Usermeta via MetaValue
func GetUsermetaCountByMetaValue(iMetaValue string) int64 {
	n, _ := Engine.Where("meta_value = ?", iMetaValue).Count(&Usermeta{MetaValue: iMetaValue})
	return n
}

// GetUsermetasByUmetaIdAndUserIdAndMetaKey Get Usermetas via UmetaIdAndUserIdAndMetaKey
func GetUsermetasByUmetaIdAndUserIdAndMetaKey(offset int, limit int, UmetaId_ int64, UserId_ int64, MetaKey_ string) (*[]*Usermeta, error) {
	var _Usermeta = new([]*Usermeta)
	err := Engine.Table("usermeta").Where("umeta_id = ? and user_id = ? and meta_key = ?", UmetaId_, UserId_, MetaKey_).Limit(limit, offset).Find(_Usermeta)
	return _Usermeta, err
}

// GetUsermetasByUmetaIdAndUserIdAndMetaValue Get Usermetas via UmetaIdAndUserIdAndMetaValue
func GetUsermetasByUmetaIdAndUserIdAndMetaValue(offset int, limit int, UmetaId_ int64, UserId_ int64, MetaValue_ string) (*[]*Usermeta, error) {
	var _Usermeta = new([]*Usermeta)
	err := Engine.Table("usermeta").Where("umeta_id = ? and user_id = ? and meta_value = ?", UmetaId_, UserId_, MetaValue_).Limit(limit, offset).Find(_Usermeta)
	return _Usermeta, err
}

// GetUsermetasByUmetaIdAndMetaKeyAndMetaValue Get Usermetas via UmetaIdAndMetaKeyAndMetaValue
func GetUsermetasByUmetaIdAndMetaKeyAndMetaValue(offset int, limit int, UmetaId_ int64, MetaKey_ string, MetaValue_ string) (*[]*Usermeta, error) {
	var _Usermeta = new([]*Usermeta)
	err := Engine.Table("usermeta").Where("umeta_id = ? and meta_key = ? and meta_value = ?", UmetaId_, MetaKey_, MetaValue_).Limit(limit, offset).Find(_Usermeta)
	return _Usermeta, err
}

// GetUsermetasByUserIdAndMetaKeyAndMetaValue Get Usermetas via UserIdAndMetaKeyAndMetaValue
func GetUsermetasByUserIdAndMetaKeyAndMetaValue(offset int, limit int, UserId_ int64, MetaKey_ string, MetaValue_ string) (*[]*Usermeta, error) {
	var _Usermeta = new([]*Usermeta)
	err := Engine.Table("usermeta").Where("user_id = ? and meta_key = ? and meta_value = ?", UserId_, MetaKey_, MetaValue_).Limit(limit, offset).Find(_Usermeta)
	return _Usermeta, err
}

// GetUsermetasByUmetaIdAndUserId Get Usermetas via UmetaIdAndUserId
func GetUsermetasByUmetaIdAndUserId(offset int, limit int, UmetaId_ int64, UserId_ int64) (*[]*Usermeta, error) {
	var _Usermeta = new([]*Usermeta)
	err := Engine.Table("usermeta").Where("umeta_id = ? and user_id = ?", UmetaId_, UserId_).Limit(limit, offset).Find(_Usermeta)
	return _Usermeta, err
}

// GetUsermetasByUmetaIdAndMetaKey Get Usermetas via UmetaIdAndMetaKey
func GetUsermetasByUmetaIdAndMetaKey(offset int, limit int, UmetaId_ int64, MetaKey_ string) (*[]*Usermeta, error) {
	var _Usermeta = new([]*Usermeta)
	err := Engine.Table("usermeta").Where("umeta_id = ? and meta_key = ?", UmetaId_, MetaKey_).Limit(limit, offset).Find(_Usermeta)
	return _Usermeta, err
}

// GetUsermetasByUmetaIdAndMetaValue Get Usermetas via UmetaIdAndMetaValue
func GetUsermetasByUmetaIdAndMetaValue(offset int, limit int, UmetaId_ int64, MetaValue_ string) (*[]*Usermeta, error) {
	var _Usermeta = new([]*Usermeta)
	err := Engine.Table("usermeta").Where("umeta_id = ? and meta_value = ?", UmetaId_, MetaValue_).Limit(limit, offset).Find(_Usermeta)
	return _Usermeta, err
}

// GetUsermetasByUserIdAndMetaKey Get Usermetas via UserIdAndMetaKey
func GetUsermetasByUserIdAndMetaKey(offset int, limit int, UserId_ int64, MetaKey_ string) (*[]*Usermeta, error) {
	var _Usermeta = new([]*Usermeta)
	err := Engine.Table("usermeta").Where("user_id = ? and meta_key = ?", UserId_, MetaKey_).Limit(limit, offset).Find(_Usermeta)
	return _Usermeta, err
}

// GetUsermetasByUserIdAndMetaValue Get Usermetas via UserIdAndMetaValue
func GetUsermetasByUserIdAndMetaValue(offset int, limit int, UserId_ int64, MetaValue_ string) (*[]*Usermeta, error) {
	var _Usermeta = new([]*Usermeta)
	err := Engine.Table("usermeta").Where("user_id = ? and meta_value = ?", UserId_, MetaValue_).Limit(limit, offset).Find(_Usermeta)
	return _Usermeta, err
}

// GetUsermetasByMetaKeyAndMetaValue Get Usermetas via MetaKeyAndMetaValue
func GetUsermetasByMetaKeyAndMetaValue(offset int, limit int, MetaKey_ string, MetaValue_ string) (*[]*Usermeta, error) {
	var _Usermeta = new([]*Usermeta)
	err := Engine.Table("usermeta").Where("meta_key = ? and meta_value = ?", MetaKey_, MetaValue_).Limit(limit, offset).Find(_Usermeta)
	return _Usermeta, err
}

// GetUsermetas Get Usermetas via field
func GetUsermetas(offset int, limit int, field string) (*[]*Usermeta, error) {
	var _Usermeta = new([]*Usermeta)
	err := Engine.Table("usermeta").Limit(limit, offset).Desc(field).Find(_Usermeta)
	return _Usermeta, err
}

// GetUsermetasByUmetaId Get Usermetas via UmetaId
func GetUsermetasByUmetaId(offset int, limit int, UmetaId_ int64, field string) (*[]*Usermeta, error) {
	var _Usermeta = new([]*Usermeta)
	err := Engine.Table("usermeta").Where("umeta_id = ?", UmetaId_).Limit(limit, offset).Desc(field).Find(_Usermeta)
	return _Usermeta, err
}

// GetUsermetasByUserId Get Usermetas via UserId
func GetUsermetasByUserId(offset int, limit int, UserId_ int64, field string) (*[]*Usermeta, error) {
	var _Usermeta = new([]*Usermeta)
	err := Engine.Table("usermeta").Where("user_id = ?", UserId_).Limit(limit, offset).Desc(field).Find(_Usermeta)
	return _Usermeta, err
}

// GetUsermetasByMetaKey Get Usermetas via MetaKey
func GetUsermetasByMetaKey(offset int, limit int, MetaKey_ string, field string) (*[]*Usermeta, error) {
	var _Usermeta = new([]*Usermeta)
	err := Engine.Table("usermeta").Where("meta_key = ?", MetaKey_).Limit(limit, offset).Desc(field).Find(_Usermeta)
	return _Usermeta, err
}

// GetUsermetasByMetaValue Get Usermetas via MetaValue
func GetUsermetasByMetaValue(offset int, limit int, MetaValue_ string, field string) (*[]*Usermeta, error) {
	var _Usermeta = new([]*Usermeta)
	err := Engine.Table("usermeta").Where("meta_value = ?", MetaValue_).Limit(limit, offset).Desc(field).Find(_Usermeta)
	return _Usermeta, err
}

// HasUsermetaByUmetaId Has Usermeta via UmetaId
func HasUsermetaByUmetaId(iUmetaId int64) bool {
	if has, err := Engine.Where("umeta_id = ?", iUmetaId).Get(new(Usermeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsermetaByUserId Has Usermeta via UserId
func HasUsermetaByUserId(iUserId int64) bool {
	if has, err := Engine.Where("user_id = ?", iUserId).Get(new(Usermeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsermetaByMetaKey Has Usermeta via MetaKey
func HasUsermetaByMetaKey(iMetaKey string) bool {
	if has, err := Engine.Where("meta_key = ?", iMetaKey).Get(new(Usermeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsermetaByMetaValue Has Usermeta via MetaValue
func HasUsermetaByMetaValue(iMetaValue string) bool {
	if has, err := Engine.Where("meta_value = ?", iMetaValue).Get(new(Usermeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetUsermetaByUmetaId Get Usermeta via UmetaId
func GetUsermetaByUmetaId(iUmetaId int64) (*Usermeta, error) {
	var _Usermeta = &Usermeta{UmetaId: iUmetaId}
	has, err := Engine.Get(_Usermeta)
	if has {
		return _Usermeta, err
	} else {
		return nil, err
	}
}

// GetUsermetaByUserId Get Usermeta via UserId
func GetUsermetaByUserId(iUserId int64) (*Usermeta, error) {
	var _Usermeta = &Usermeta{UserId: iUserId}
	has, err := Engine.Get(_Usermeta)
	if has {
		return _Usermeta, err
	} else {
		return nil, err
	}
}

// GetUsermetaByMetaKey Get Usermeta via MetaKey
func GetUsermetaByMetaKey(iMetaKey string) (*Usermeta, error) {
	var _Usermeta = &Usermeta{MetaKey: iMetaKey}
	has, err := Engine.Get(_Usermeta)
	if has {
		return _Usermeta, err
	} else {
		return nil, err
	}
}

// GetUsermetaByMetaValue Get Usermeta via MetaValue
func GetUsermetaByMetaValue(iMetaValue string) (*Usermeta, error) {
	var _Usermeta = &Usermeta{MetaValue: iMetaValue}
	has, err := Engine.Get(_Usermeta)
	if has {
		return _Usermeta, err
	} else {
		return nil, err
	}
}

// SetUsermetaByUmetaId Set Usermeta via UmetaId
func SetUsermetaByUmetaId(iUmetaId int64, usermeta *Usermeta) (int64, error) {
	usermeta.UmetaId = iUmetaId
	return Engine.Insert(usermeta)
}

// SetUsermetaByUserId Set Usermeta via UserId
func SetUsermetaByUserId(iUserId int64, usermeta *Usermeta) (int64, error) {
	usermeta.UserId = iUserId
	return Engine.Insert(usermeta)
}

// SetUsermetaByMetaKey Set Usermeta via MetaKey
func SetUsermetaByMetaKey(iMetaKey string, usermeta *Usermeta) (int64, error) {
	usermeta.MetaKey = iMetaKey
	return Engine.Insert(usermeta)
}

// SetUsermetaByMetaValue Set Usermeta via MetaValue
func SetUsermetaByMetaValue(iMetaValue string, usermeta *Usermeta) (int64, error) {
	usermeta.MetaValue = iMetaValue
	return Engine.Insert(usermeta)
}

// AddUsermeta Add Usermeta via all columns
func AddUsermeta(iUmetaId int64, iUserId int64, iMetaKey string, iMetaValue string) error {
	_Usermeta := &Usermeta{UmetaId: iUmetaId, UserId: iUserId, MetaKey: iMetaKey, MetaValue: iMetaValue}
	if _, err := Engine.Insert(_Usermeta); err != nil {
		return err
	} else {
		return nil
	}
}

// PostUsermeta Post Usermeta via iUsermeta
func PostUsermeta(iUsermeta *Usermeta) (int64, error) {
	_, err := Engine.Insert(iUsermeta)
	return iUsermeta.UmetaId, err
}

// PutUsermeta Put Usermeta
func PutUsermeta(iUsermeta *Usermeta) (int64, error) {
	_, err := Engine.Id(iUsermeta.UmetaId).Update(iUsermeta)
	return iUsermeta.UmetaId, err
}

// PutUsermetaByUmetaId Put Usermeta via UmetaId
func PutUsermetaByUmetaId(UmetaId_ int64, iUsermeta *Usermeta) (int64, error) {
	row, err := Engine.Update(iUsermeta, &Usermeta{UmetaId: UmetaId_})
	return row, err
}

// PutUsermetaByUserId Put Usermeta via UserId
func PutUsermetaByUserId(UserId_ int64, iUsermeta *Usermeta) (int64, error) {
	row, err := Engine.Update(iUsermeta, &Usermeta{UserId: UserId_})
	return row, err
}

// PutUsermetaByMetaKey Put Usermeta via MetaKey
func PutUsermetaByMetaKey(MetaKey_ string, iUsermeta *Usermeta) (int64, error) {
	row, err := Engine.Update(iUsermeta, &Usermeta{MetaKey: MetaKey_})
	return row, err
}

// PutUsermetaByMetaValue Put Usermeta via MetaValue
func PutUsermetaByMetaValue(MetaValue_ string, iUsermeta *Usermeta) (int64, error) {
	row, err := Engine.Update(iUsermeta, &Usermeta{MetaValue: MetaValue_})
	return row, err
}

// UpdateUsermetaByUmetaId via map[string]interface{}{}
func UpdateUsermetaByUmetaId(iUmetaId int64, iUsermetaMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Usermeta)).Where("umeta_id = ?", iUmetaId).Update(iUsermetaMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsermetaByUserId via map[string]interface{}{}
func UpdateUsermetaByUserId(iUserId int64, iUsermetaMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Usermeta)).Where("user_id = ?", iUserId).Update(iUsermetaMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsermetaByMetaKey via map[string]interface{}{}
func UpdateUsermetaByMetaKey(iMetaKey string, iUsermetaMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Usermeta)).Where("meta_key = ?", iMetaKey).Update(iUsermetaMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsermetaByMetaValue via map[string]interface{}{}
func UpdateUsermetaByMetaValue(iMetaValue string, iUsermetaMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Usermeta)).Where("meta_value = ?", iMetaValue).Update(iUsermetaMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteUsermeta Delete Usermeta
func DeleteUsermeta(iUmetaId int64) (int64, error) {
	row, err := Engine.Id(iUmetaId).Delete(new(Usermeta))
	return row, err
}

// DeleteUsermetaByUmetaId Delete Usermeta via UmetaId
func DeleteUsermetaByUmetaId(iUmetaId int64) (err error) {
	var has bool
	var _Usermeta = &Usermeta{UmetaId: iUmetaId}
	if has, err = Engine.Get(_Usermeta); (has == true) && (err == nil) {
		if row, err := Engine.Where("umeta_id = ?", iUmetaId).Delete(new(Usermeta)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsermetaByUserId Delete Usermeta via UserId
func DeleteUsermetaByUserId(iUserId int64) (err error) {
	var has bool
	var _Usermeta = &Usermeta{UserId: iUserId}
	if has, err = Engine.Get(_Usermeta); (has == true) && (err == nil) {
		if row, err := Engine.Where("user_id = ?", iUserId).Delete(new(Usermeta)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsermetaByMetaKey Delete Usermeta via MetaKey
func DeleteUsermetaByMetaKey(iMetaKey string) (err error) {
	var has bool
	var _Usermeta = &Usermeta{MetaKey: iMetaKey}
	if has, err = Engine.Get(_Usermeta); (has == true) && (err == nil) {
		if row, err := Engine.Where("meta_key = ?", iMetaKey).Delete(new(Usermeta)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsermetaByMetaValue Delete Usermeta via MetaValue
func DeleteUsermetaByMetaValue(iMetaValue string) (err error) {
	var has bool
	var _Usermeta = &Usermeta{MetaValue: iMetaValue}
	if has, err = Engine.Get(_Usermeta); (has == true) && (err == nil) {
		if row, err := Engine.Where("meta_value = ?", iMetaValue).Delete(new(Usermeta)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
