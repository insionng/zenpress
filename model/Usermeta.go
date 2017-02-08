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

// GetUsermetaCountViaUmetaId Get Usermeta via UmetaId
func GetUsermetaCountViaUmetaId(iUmetaId int64) int64 {
	n, _ := Engine.Where("umeta_id = ?", iUmetaId).Count(&Usermeta{UmetaId: iUmetaId})
	return n
}

// GetUsermetaCountViaUserId Get Usermeta via UserId
func GetUsermetaCountViaUserId(iUserId int64) int64 {
	n, _ := Engine.Where("user_id = ?", iUserId).Count(&Usermeta{UserId: iUserId})
	return n
}

// GetUsermetaCountViaMetaKey Get Usermeta via MetaKey
func GetUsermetaCountViaMetaKey(iMetaKey string) int64 {
	n, _ := Engine.Where("meta_key = ?", iMetaKey).Count(&Usermeta{MetaKey: iMetaKey})
	return n
}

// GetUsermetaCountViaMetaValue Get Usermeta via MetaValue
func GetUsermetaCountViaMetaValue(iMetaValue string) int64 {
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

// GetUsermetasViaUmetaId Get Usermetas via UmetaId
func GetUsermetasViaUmetaId(offset int, limit int, UmetaId_ int64, field string) (*[]*Usermeta, error) {
	var _Usermeta = new([]*Usermeta)
	err := Engine.Table("usermeta").Where("umeta_id = ?", UmetaId_).Limit(limit, offset).Desc(field).Find(_Usermeta)
	return _Usermeta, err
}

// GetUsermetasViaUserId Get Usermetas via UserId
func GetUsermetasViaUserId(offset int, limit int, UserId_ int64, field string) (*[]*Usermeta, error) {
	var _Usermeta = new([]*Usermeta)
	err := Engine.Table("usermeta").Where("user_id = ?", UserId_).Limit(limit, offset).Desc(field).Find(_Usermeta)
	return _Usermeta, err
}

// GetUsermetasViaMetaKey Get Usermetas via MetaKey
func GetUsermetasViaMetaKey(offset int, limit int, MetaKey_ string, field string) (*[]*Usermeta, error) {
	var _Usermeta = new([]*Usermeta)
	err := Engine.Table("usermeta").Where("meta_key = ?", MetaKey_).Limit(limit, offset).Desc(field).Find(_Usermeta)
	return _Usermeta, err
}

// GetUsermetasViaMetaValue Get Usermetas via MetaValue
func GetUsermetasViaMetaValue(offset int, limit int, MetaValue_ string, field string) (*[]*Usermeta, error) {
	var _Usermeta = new([]*Usermeta)
	err := Engine.Table("usermeta").Where("meta_value = ?", MetaValue_).Limit(limit, offset).Desc(field).Find(_Usermeta)
	return _Usermeta, err
}

// HasUsermetaViaUmetaId Has Usermeta via UmetaId
func HasUsermetaViaUmetaId(iUmetaId int64) bool {
	if has, err := Engine.Where("umeta_id = ?", iUmetaId).Get(new(Usermeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsermetaViaUserId Has Usermeta via UserId
func HasUsermetaViaUserId(iUserId int64) bool {
	if has, err := Engine.Where("user_id = ?", iUserId).Get(new(Usermeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsermetaViaMetaKey Has Usermeta via MetaKey
func HasUsermetaViaMetaKey(iMetaKey string) bool {
	if has, err := Engine.Where("meta_key = ?", iMetaKey).Get(new(Usermeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsermetaViaMetaValue Has Usermeta via MetaValue
func HasUsermetaViaMetaValue(iMetaValue string) bool {
	if has, err := Engine.Where("meta_value = ?", iMetaValue).Get(new(Usermeta)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetUsermetaViaUmetaId Get Usermeta via UmetaId
func GetUsermetaViaUmetaId(iUmetaId int64) (*Usermeta, error) {
	var _Usermeta = &Usermeta{UmetaId: iUmetaId}
	has, err := Engine.Get(_Usermeta)
	if has {
		return _Usermeta, err
	} else {
		return nil, err
	}
}

// GetUsermetaViaUserId Get Usermeta via UserId
func GetUsermetaViaUserId(iUserId int64) (*Usermeta, error) {
	var _Usermeta = &Usermeta{UserId: iUserId}
	has, err := Engine.Get(_Usermeta)
	if has {
		return _Usermeta, err
	} else {
		return nil, err
	}
}

// GetUsermetaViaMetaKey Get Usermeta via MetaKey
func GetUsermetaViaMetaKey(iMetaKey string) (*Usermeta, error) {
	var _Usermeta = &Usermeta{MetaKey: iMetaKey}
	has, err := Engine.Get(_Usermeta)
	if has {
		return _Usermeta, err
	} else {
		return nil, err
	}
}

// GetUsermetaViaMetaValue Get Usermeta via MetaValue
func GetUsermetaViaMetaValue(iMetaValue string) (*Usermeta, error) {
	var _Usermeta = &Usermeta{MetaValue: iMetaValue}
	has, err := Engine.Get(_Usermeta)
	if has {
		return _Usermeta, err
	} else {
		return nil, err
	}
}

// SetUsermetaViaUmetaId Set Usermeta via UmetaId
func SetUsermetaViaUmetaId(iUmetaId int64, usermeta *Usermeta) (int64, error) {
	usermeta.UmetaId = iUmetaId
	return Engine.Insert(usermeta)
}

// SetUsermetaViaUserId Set Usermeta via UserId
func SetUsermetaViaUserId(iUserId int64, usermeta *Usermeta) (int64, error) {
	usermeta.UserId = iUserId
	return Engine.Insert(usermeta)
}

// SetUsermetaViaMetaKey Set Usermeta via MetaKey
func SetUsermetaViaMetaKey(iMetaKey string, usermeta *Usermeta) (int64, error) {
	usermeta.MetaKey = iMetaKey
	return Engine.Insert(usermeta)
}

// SetUsermetaViaMetaValue Set Usermeta via MetaValue
func SetUsermetaViaMetaValue(iMetaValue string, usermeta *Usermeta) (int64, error) {
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

// PutUsermetaViaUmetaId Put Usermeta via UmetaId
func PutUsermetaViaUmetaId(UmetaId_ int64, iUsermeta *Usermeta) (int64, error) {
	row, err := Engine.Update(iUsermeta, &Usermeta{UmetaId: UmetaId_})
	return row, err
}

// PutUsermetaViaUserId Put Usermeta via UserId
func PutUsermetaViaUserId(UserId_ int64, iUsermeta *Usermeta) (int64, error) {
	row, err := Engine.Update(iUsermeta, &Usermeta{UserId: UserId_})
	return row, err
}

// PutUsermetaViaMetaKey Put Usermeta via MetaKey
func PutUsermetaViaMetaKey(MetaKey_ string, iUsermeta *Usermeta) (int64, error) {
	row, err := Engine.Update(iUsermeta, &Usermeta{MetaKey: MetaKey_})
	return row, err
}

// PutUsermetaViaMetaValue Put Usermeta via MetaValue
func PutUsermetaViaMetaValue(MetaValue_ string, iUsermeta *Usermeta) (int64, error) {
	row, err := Engine.Update(iUsermeta, &Usermeta{MetaValue: MetaValue_})
	return row, err
}

// UpdateUsermetaViaUmetaId via map[string]interface{}{}
func UpdateUsermetaViaUmetaId(iUmetaId int64, iUsermetaMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Usermeta)).Where("umeta_id = ?", iUmetaId).Update(iUsermetaMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsermetaViaUserId via map[string]interface{}{}
func UpdateUsermetaViaUserId(iUserId int64, iUsermetaMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Usermeta)).Where("user_id = ?", iUserId).Update(iUsermetaMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsermetaViaMetaKey via map[string]interface{}{}
func UpdateUsermetaViaMetaKey(iMetaKey string, iUsermetaMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Usermeta)).Where("meta_key = ?", iMetaKey).Update(iUsermetaMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsermetaViaMetaValue via map[string]interface{}{}
func UpdateUsermetaViaMetaValue(iMetaValue string, iUsermetaMap *map[string]interface{}) error {
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

// DeleteUsermetaViaUmetaId Delete Usermeta via UmetaId
func DeleteUsermetaViaUmetaId(iUmetaId int64) (err error) {
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

// DeleteUsermetaViaUserId Delete Usermeta via UserId
func DeleteUsermetaViaUserId(iUserId int64) (err error) {
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

// DeleteUsermetaViaMetaKey Delete Usermeta via MetaKey
func DeleteUsermetaViaMetaKey(iMetaKey string) (err error) {
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

// DeleteUsermetaViaMetaValue Delete Usermeta via MetaValue
func DeleteUsermetaViaMetaValue(iMetaValue string) (err error) {
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
