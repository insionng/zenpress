package model

import (
	"time"
)

type RegistrationLog struct {
	Id             int64     `xorm:"BIGINT(20)"`
	Email          string    `xorm:"not null default '' VARCHAR(255)"`
	Ip             string    `xorm:"not null default '' index VARCHAR(30)"`
	BlogId         int64     `xorm:"not null default 0 BIGINT(20)"`
	DateRegistered time.Time `xorm:"not null default '0000-00-00 00:00:00' DATETIME"`
}

// GetRegistrationLogsCount RegistrationLogs' Count
func GetRegistrationLogsCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&RegistrationLog{})
	return total, err
}

// GetRegistrationLogCountById Get RegistrationLog via Id
func GetRegistrationLogCountById(iId int64) int64 {
	n, _ := Engine.Where("ID = ?", iId).Count(&RegistrationLog{Id: iId})
	return n
}

// GetRegistrationLogCountByEmail Get RegistrationLog via Email
func GetRegistrationLogCountByEmail(iEmail string) int64 {
	n, _ := Engine.Where("email = ?", iEmail).Count(&RegistrationLog{Email: iEmail})
	return n
}

// GetRegistrationLogCountByIp Get RegistrationLog via Ip
func GetRegistrationLogCountByIp(iIp string) int64 {
	n, _ := Engine.Where("IP = ?", iIp).Count(&RegistrationLog{Ip: iIp})
	return n
}

// GetRegistrationLogCountByBlogId Get RegistrationLog via BlogId
func GetRegistrationLogCountByBlogId(iBlogId int64) int64 {
	n, _ := Engine.Where("blog_id = ?", iBlogId).Count(&RegistrationLog{BlogId: iBlogId})
	return n
}

// GetRegistrationLogCountByDateRegistered Get RegistrationLog via DateRegistered
func GetRegistrationLogCountByDateRegistered(iDateRegistered time.Time) int64 {
	n, _ := Engine.Where("date_registered = ?", iDateRegistered).Count(&RegistrationLog{DateRegistered: iDateRegistered})
	return n
}

// GetRegistrationLogsByIdAndEmailAndIp Get RegistrationLogs via IdAndEmailAndIp
func GetRegistrationLogsByIdAndEmailAndIp(offset int, limit int, Id_ int64, Email_ string, Ip_ string) (*[]*RegistrationLog, error) {
	var _RegistrationLog = new([]*RegistrationLog)
	err := Engine.Table("registration_log").Where("id = ? and email = ? and ip = ?", Id_, Email_, Ip_).Limit(limit, offset).Find(_RegistrationLog)
	return _RegistrationLog, err
}

// GetRegistrationLogsByIdAndEmailAndBlogId Get RegistrationLogs via IdAndEmailAndBlogId
func GetRegistrationLogsByIdAndEmailAndBlogId(offset int, limit int, Id_ int64, Email_ string, BlogId_ int64) (*[]*RegistrationLog, error) {
	var _RegistrationLog = new([]*RegistrationLog)
	err := Engine.Table("registration_log").Where("id = ? and email = ? and blog_id = ?", Id_, Email_, BlogId_).Limit(limit, offset).Find(_RegistrationLog)
	return _RegistrationLog, err
}

// GetRegistrationLogsByIdAndEmailAndDateRegistered Get RegistrationLogs via IdAndEmailAndDateRegistered
func GetRegistrationLogsByIdAndEmailAndDateRegistered(offset int, limit int, Id_ int64, Email_ string, DateRegistered_ time.Time) (*[]*RegistrationLog, error) {
	var _RegistrationLog = new([]*RegistrationLog)
	err := Engine.Table("registration_log").Where("id = ? and email = ? and date_registered = ?", Id_, Email_, DateRegistered_).Limit(limit, offset).Find(_RegistrationLog)
	return _RegistrationLog, err
}

// GetRegistrationLogsByIdAndIpAndBlogId Get RegistrationLogs via IdAndIpAndBlogId
func GetRegistrationLogsByIdAndIpAndBlogId(offset int, limit int, Id_ int64, Ip_ string, BlogId_ int64) (*[]*RegistrationLog, error) {
	var _RegistrationLog = new([]*RegistrationLog)
	err := Engine.Table("registration_log").Where("id = ? and ip = ? and blog_id = ?", Id_, Ip_, BlogId_).Limit(limit, offset).Find(_RegistrationLog)
	return _RegistrationLog, err
}

// GetRegistrationLogsByIdAndIpAndDateRegistered Get RegistrationLogs via IdAndIpAndDateRegistered
func GetRegistrationLogsByIdAndIpAndDateRegistered(offset int, limit int, Id_ int64, Ip_ string, DateRegistered_ time.Time) (*[]*RegistrationLog, error) {
	var _RegistrationLog = new([]*RegistrationLog)
	err := Engine.Table("registration_log").Where("id = ? and ip = ? and date_registered = ?", Id_, Ip_, DateRegistered_).Limit(limit, offset).Find(_RegistrationLog)
	return _RegistrationLog, err
}

// GetRegistrationLogsByIdAndBlogIdAndDateRegistered Get RegistrationLogs via IdAndBlogIdAndDateRegistered
func GetRegistrationLogsByIdAndBlogIdAndDateRegistered(offset int, limit int, Id_ int64, BlogId_ int64, DateRegistered_ time.Time) (*[]*RegistrationLog, error) {
	var _RegistrationLog = new([]*RegistrationLog)
	err := Engine.Table("registration_log").Where("id = ? and blog_id = ? and date_registered = ?", Id_, BlogId_, DateRegistered_).Limit(limit, offset).Find(_RegistrationLog)
	return _RegistrationLog, err
}

// GetRegistrationLogsByEmailAndIpAndBlogId Get RegistrationLogs via EmailAndIpAndBlogId
func GetRegistrationLogsByEmailAndIpAndBlogId(offset int, limit int, Email_ string, Ip_ string, BlogId_ int64) (*[]*RegistrationLog, error) {
	var _RegistrationLog = new([]*RegistrationLog)
	err := Engine.Table("registration_log").Where("email = ? and ip = ? and blog_id = ?", Email_, Ip_, BlogId_).Limit(limit, offset).Find(_RegistrationLog)
	return _RegistrationLog, err
}

// GetRegistrationLogsByEmailAndIpAndDateRegistered Get RegistrationLogs via EmailAndIpAndDateRegistered
func GetRegistrationLogsByEmailAndIpAndDateRegistered(offset int, limit int, Email_ string, Ip_ string, DateRegistered_ time.Time) (*[]*RegistrationLog, error) {
	var _RegistrationLog = new([]*RegistrationLog)
	err := Engine.Table("registration_log").Where("email = ? and ip = ? and date_registered = ?", Email_, Ip_, DateRegistered_).Limit(limit, offset).Find(_RegistrationLog)
	return _RegistrationLog, err
}

// GetRegistrationLogsByEmailAndBlogIdAndDateRegistered Get RegistrationLogs via EmailAndBlogIdAndDateRegistered
func GetRegistrationLogsByEmailAndBlogIdAndDateRegistered(offset int, limit int, Email_ string, BlogId_ int64, DateRegistered_ time.Time) (*[]*RegistrationLog, error) {
	var _RegistrationLog = new([]*RegistrationLog)
	err := Engine.Table("registration_log").Where("email = ? and blog_id = ? and date_registered = ?", Email_, BlogId_, DateRegistered_).Limit(limit, offset).Find(_RegistrationLog)
	return _RegistrationLog, err
}

// GetRegistrationLogsByIpAndBlogIdAndDateRegistered Get RegistrationLogs via IpAndBlogIdAndDateRegistered
func GetRegistrationLogsByIpAndBlogIdAndDateRegistered(offset int, limit int, Ip_ string, BlogId_ int64, DateRegistered_ time.Time) (*[]*RegistrationLog, error) {
	var _RegistrationLog = new([]*RegistrationLog)
	err := Engine.Table("registration_log").Where("ip = ? and blog_id = ? and date_registered = ?", Ip_, BlogId_, DateRegistered_).Limit(limit, offset).Find(_RegistrationLog)
	return _RegistrationLog, err
}

// GetRegistrationLogsByIdAndEmail Get RegistrationLogs via IdAndEmail
func GetRegistrationLogsByIdAndEmail(offset int, limit int, Id_ int64, Email_ string) (*[]*RegistrationLog, error) {
	var _RegistrationLog = new([]*RegistrationLog)
	err := Engine.Table("registration_log").Where("id = ? and email = ?", Id_, Email_).Limit(limit, offset).Find(_RegistrationLog)
	return _RegistrationLog, err
}

// GetRegistrationLogsByIdAndIp Get RegistrationLogs via IdAndIp
func GetRegistrationLogsByIdAndIp(offset int, limit int, Id_ int64, Ip_ string) (*[]*RegistrationLog, error) {
	var _RegistrationLog = new([]*RegistrationLog)
	err := Engine.Table("registration_log").Where("id = ? and ip = ?", Id_, Ip_).Limit(limit, offset).Find(_RegistrationLog)
	return _RegistrationLog, err
}

// GetRegistrationLogsByIdAndBlogId Get RegistrationLogs via IdAndBlogId
func GetRegistrationLogsByIdAndBlogId(offset int, limit int, Id_ int64, BlogId_ int64) (*[]*RegistrationLog, error) {
	var _RegistrationLog = new([]*RegistrationLog)
	err := Engine.Table("registration_log").Where("id = ? and blog_id = ?", Id_, BlogId_).Limit(limit, offset).Find(_RegistrationLog)
	return _RegistrationLog, err
}

// GetRegistrationLogsByIdAndDateRegistered Get RegistrationLogs via IdAndDateRegistered
func GetRegistrationLogsByIdAndDateRegistered(offset int, limit int, Id_ int64, DateRegistered_ time.Time) (*[]*RegistrationLog, error) {
	var _RegistrationLog = new([]*RegistrationLog)
	err := Engine.Table("registration_log").Where("id = ? and date_registered = ?", Id_, DateRegistered_).Limit(limit, offset).Find(_RegistrationLog)
	return _RegistrationLog, err
}

// GetRegistrationLogsByEmailAndIp Get RegistrationLogs via EmailAndIp
func GetRegistrationLogsByEmailAndIp(offset int, limit int, Email_ string, Ip_ string) (*[]*RegistrationLog, error) {
	var _RegistrationLog = new([]*RegistrationLog)
	err := Engine.Table("registration_log").Where("email = ? and ip = ?", Email_, Ip_).Limit(limit, offset).Find(_RegistrationLog)
	return _RegistrationLog, err
}

// GetRegistrationLogsByEmailAndBlogId Get RegistrationLogs via EmailAndBlogId
func GetRegistrationLogsByEmailAndBlogId(offset int, limit int, Email_ string, BlogId_ int64) (*[]*RegistrationLog, error) {
	var _RegistrationLog = new([]*RegistrationLog)
	err := Engine.Table("registration_log").Where("email = ? and blog_id = ?", Email_, BlogId_).Limit(limit, offset).Find(_RegistrationLog)
	return _RegistrationLog, err
}

// GetRegistrationLogsByEmailAndDateRegistered Get RegistrationLogs via EmailAndDateRegistered
func GetRegistrationLogsByEmailAndDateRegistered(offset int, limit int, Email_ string, DateRegistered_ time.Time) (*[]*RegistrationLog, error) {
	var _RegistrationLog = new([]*RegistrationLog)
	err := Engine.Table("registration_log").Where("email = ? and date_registered = ?", Email_, DateRegistered_).Limit(limit, offset).Find(_RegistrationLog)
	return _RegistrationLog, err
}

// GetRegistrationLogsByIpAndBlogId Get RegistrationLogs via IpAndBlogId
func GetRegistrationLogsByIpAndBlogId(offset int, limit int, Ip_ string, BlogId_ int64) (*[]*RegistrationLog, error) {
	var _RegistrationLog = new([]*RegistrationLog)
	err := Engine.Table("registration_log").Where("ip = ? and blog_id = ?", Ip_, BlogId_).Limit(limit, offset).Find(_RegistrationLog)
	return _RegistrationLog, err
}

// GetRegistrationLogsByIpAndDateRegistered Get RegistrationLogs via IpAndDateRegistered
func GetRegistrationLogsByIpAndDateRegistered(offset int, limit int, Ip_ string, DateRegistered_ time.Time) (*[]*RegistrationLog, error) {
	var _RegistrationLog = new([]*RegistrationLog)
	err := Engine.Table("registration_log").Where("ip = ? and date_registered = ?", Ip_, DateRegistered_).Limit(limit, offset).Find(_RegistrationLog)
	return _RegistrationLog, err
}

// GetRegistrationLogsByBlogIdAndDateRegistered Get RegistrationLogs via BlogIdAndDateRegistered
func GetRegistrationLogsByBlogIdAndDateRegistered(offset int, limit int, BlogId_ int64, DateRegistered_ time.Time) (*[]*RegistrationLog, error) {
	var _RegistrationLog = new([]*RegistrationLog)
	err := Engine.Table("registration_log").Where("blog_id = ? and date_registered = ?", BlogId_, DateRegistered_).Limit(limit, offset).Find(_RegistrationLog)
	return _RegistrationLog, err
}

// GetRegistrationLogs Get RegistrationLogs via field
func GetRegistrationLogs(offset int, limit int, field string) (*[]*RegistrationLog, error) {
	var _RegistrationLog = new([]*RegistrationLog)
	err := Engine.Table("registration_log").Limit(limit, offset).Desc(field).Find(_RegistrationLog)
	return _RegistrationLog, err
}

// GetRegistrationLogsById Get RegistrationLogs via Id
func GetRegistrationLogsById(offset int, limit int, Id_ int64, field string) (*[]*RegistrationLog, error) {
	var _RegistrationLog = new([]*RegistrationLog)
	err := Engine.Table("registration_log").Where("ID = ?", Id_).Limit(limit, offset).Desc(field).Find(_RegistrationLog)
	return _RegistrationLog, err
}

// GetRegistrationLogsByEmail Get RegistrationLogs via Email
func GetRegistrationLogsByEmail(offset int, limit int, Email_ string, field string) (*[]*RegistrationLog, error) {
	var _RegistrationLog = new([]*RegistrationLog)
	err := Engine.Table("registration_log").Where("email = ?", Email_).Limit(limit, offset).Desc(field).Find(_RegistrationLog)
	return _RegistrationLog, err
}

// GetRegistrationLogsByIp Get RegistrationLogs via Ip
func GetRegistrationLogsByIp(offset int, limit int, Ip_ string, field string) (*[]*RegistrationLog, error) {
	var _RegistrationLog = new([]*RegistrationLog)
	err := Engine.Table("registration_log").Where("IP = ?", Ip_).Limit(limit, offset).Desc(field).Find(_RegistrationLog)
	return _RegistrationLog, err
}

// GetRegistrationLogsByBlogId Get RegistrationLogs via BlogId
func GetRegistrationLogsByBlogId(offset int, limit int, BlogId_ int64, field string) (*[]*RegistrationLog, error) {
	var _RegistrationLog = new([]*RegistrationLog)
	err := Engine.Table("registration_log").Where("blog_id = ?", BlogId_).Limit(limit, offset).Desc(field).Find(_RegistrationLog)
	return _RegistrationLog, err
}

// GetRegistrationLogsByDateRegistered Get RegistrationLogs via DateRegistered
func GetRegistrationLogsByDateRegistered(offset int, limit int, DateRegistered_ time.Time, field string) (*[]*RegistrationLog, error) {
	var _RegistrationLog = new([]*RegistrationLog)
	err := Engine.Table("registration_log").Where("date_registered = ?", DateRegistered_).Limit(limit, offset).Desc(field).Find(_RegistrationLog)
	return _RegistrationLog, err
}

// HasRegistrationLogById Has RegistrationLog via Id
func HasRegistrationLogById(iId int64) bool {
	if has, err := Engine.Where("ID = ?", iId).Get(new(RegistrationLog)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasRegistrationLogByEmail Has RegistrationLog via Email
func HasRegistrationLogByEmail(iEmail string) bool {
	if has, err := Engine.Where("email = ?", iEmail).Get(new(RegistrationLog)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasRegistrationLogByIp Has RegistrationLog via Ip
func HasRegistrationLogByIp(iIp string) bool {
	if has, err := Engine.Where("IP = ?", iIp).Get(new(RegistrationLog)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasRegistrationLogByBlogId Has RegistrationLog via BlogId
func HasRegistrationLogByBlogId(iBlogId int64) bool {
	if has, err := Engine.Where("blog_id = ?", iBlogId).Get(new(RegistrationLog)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasRegistrationLogByDateRegistered Has RegistrationLog via DateRegistered
func HasRegistrationLogByDateRegistered(iDateRegistered time.Time) bool {
	if has, err := Engine.Where("date_registered = ?", iDateRegistered).Get(new(RegistrationLog)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetRegistrationLogById Get RegistrationLog via Id
func GetRegistrationLogById(iId int64) (*RegistrationLog, error) {
	var _RegistrationLog = &RegistrationLog{Id: iId}
	has, err := Engine.Get(_RegistrationLog)
	if has {
		return _RegistrationLog, err
	} else {
		return nil, err
	}
}

// GetRegistrationLogByEmail Get RegistrationLog via Email
func GetRegistrationLogByEmail(iEmail string) (*RegistrationLog, error) {
	var _RegistrationLog = &RegistrationLog{Email: iEmail}
	has, err := Engine.Get(_RegistrationLog)
	if has {
		return _RegistrationLog, err
	} else {
		return nil, err
	}
}

// GetRegistrationLogByIp Get RegistrationLog via Ip
func GetRegistrationLogByIp(iIp string) (*RegistrationLog, error) {
	var _RegistrationLog = &RegistrationLog{Ip: iIp}
	has, err := Engine.Get(_RegistrationLog)
	if has {
		return _RegistrationLog, err
	} else {
		return nil, err
	}
}

// GetRegistrationLogByBlogId Get RegistrationLog via BlogId
func GetRegistrationLogByBlogId(iBlogId int64) (*RegistrationLog, error) {
	var _RegistrationLog = &RegistrationLog{BlogId: iBlogId}
	has, err := Engine.Get(_RegistrationLog)
	if has {
		return _RegistrationLog, err
	} else {
		return nil, err
	}
}

// GetRegistrationLogByDateRegistered Get RegistrationLog via DateRegistered
func GetRegistrationLogByDateRegistered(iDateRegistered time.Time) (*RegistrationLog, error) {
	var _RegistrationLog = &RegistrationLog{DateRegistered: iDateRegistered}
	has, err := Engine.Get(_RegistrationLog)
	if has {
		return _RegistrationLog, err
	} else {
		return nil, err
	}
}

// SetRegistrationLogById Set RegistrationLog via Id
func SetRegistrationLogById(iId int64, registration_log *RegistrationLog) (int64, error) {
	registration_log.Id = iId
	return Engine.Insert(registration_log)
}

// SetRegistrationLogByEmail Set RegistrationLog via Email
func SetRegistrationLogByEmail(iEmail string, registration_log *RegistrationLog) (int64, error) {
	registration_log.Email = iEmail
	return Engine.Insert(registration_log)
}

// SetRegistrationLogByIp Set RegistrationLog via Ip
func SetRegistrationLogByIp(iIp string, registration_log *RegistrationLog) (int64, error) {
	registration_log.Ip = iIp
	return Engine.Insert(registration_log)
}

// SetRegistrationLogByBlogId Set RegistrationLog via BlogId
func SetRegistrationLogByBlogId(iBlogId int64, registration_log *RegistrationLog) (int64, error) {
	registration_log.BlogId = iBlogId
	return Engine.Insert(registration_log)
}

// SetRegistrationLogByDateRegistered Set RegistrationLog via DateRegistered
func SetRegistrationLogByDateRegistered(iDateRegistered time.Time, registration_log *RegistrationLog) (int64, error) {
	registration_log.DateRegistered = iDateRegistered
	return Engine.Insert(registration_log)
}

// AddRegistrationLog Add RegistrationLog via all columns
func AddRegistrationLog(iId int64, iEmail string, iIp string, iBlogId int64, iDateRegistered time.Time) error {
	_RegistrationLog := &RegistrationLog{Id: iId, Email: iEmail, Ip: iIp, BlogId: iBlogId, DateRegistered: iDateRegistered}
	if _, err := Engine.Insert(_RegistrationLog); err != nil {
		return err
	} else {
		return nil
	}
}

// PostRegistrationLog Post RegistrationLog via iRegistrationLog
func PostRegistrationLog(iRegistrationLog *RegistrationLog) (int64, error) {
	_, err := Engine.Insert(iRegistrationLog)
	return iRegistrationLog.Id, err
}

// PutRegistrationLog Put RegistrationLog
func PutRegistrationLog(iRegistrationLog *RegistrationLog) (int64, error) {
	_, err := Engine.Id(iRegistrationLog.Id).Update(iRegistrationLog)
	return iRegistrationLog.Id, err
}

// PutRegistrationLogById Put RegistrationLog via Id
func PutRegistrationLogById(Id_ int64, iRegistrationLog *RegistrationLog) (int64, error) {
	row, err := Engine.Update(iRegistrationLog, &RegistrationLog{Id: Id_})
	return row, err
}

// PutRegistrationLogByEmail Put RegistrationLog via Email
func PutRegistrationLogByEmail(Email_ string, iRegistrationLog *RegistrationLog) (int64, error) {
	row, err := Engine.Update(iRegistrationLog, &RegistrationLog{Email: Email_})
	return row, err
}

// PutRegistrationLogByIp Put RegistrationLog via Ip
func PutRegistrationLogByIp(Ip_ string, iRegistrationLog *RegistrationLog) (int64, error) {
	row, err := Engine.Update(iRegistrationLog, &RegistrationLog{Ip: Ip_})
	return row, err
}

// PutRegistrationLogByBlogId Put RegistrationLog via BlogId
func PutRegistrationLogByBlogId(BlogId_ int64, iRegistrationLog *RegistrationLog) (int64, error) {
	row, err := Engine.Update(iRegistrationLog, &RegistrationLog{BlogId: BlogId_})
	return row, err
}

// PutRegistrationLogByDateRegistered Put RegistrationLog via DateRegistered
func PutRegistrationLogByDateRegistered(DateRegistered_ time.Time, iRegistrationLog *RegistrationLog) (int64, error) {
	row, err := Engine.Update(iRegistrationLog, &RegistrationLog{DateRegistered: DateRegistered_})
	return row, err
}

// UpdateRegistrationLogById via map[string]interface{}{}
func UpdateRegistrationLogById(iId int64, iRegistrationLogMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(RegistrationLog)).Where("ID = ?", iId).Update(iRegistrationLogMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateRegistrationLogByEmail via map[string]interface{}{}
func UpdateRegistrationLogByEmail(iEmail string, iRegistrationLogMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(RegistrationLog)).Where("email = ?", iEmail).Update(iRegistrationLogMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateRegistrationLogByIp via map[string]interface{}{}
func UpdateRegistrationLogByIp(iIp string, iRegistrationLogMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(RegistrationLog)).Where("IP = ?", iIp).Update(iRegistrationLogMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateRegistrationLogByBlogId via map[string]interface{}{}
func UpdateRegistrationLogByBlogId(iBlogId int64, iRegistrationLogMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(RegistrationLog)).Where("blog_id = ?", iBlogId).Update(iRegistrationLogMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateRegistrationLogByDateRegistered via map[string]interface{}{}
func UpdateRegistrationLogByDateRegistered(iDateRegistered time.Time, iRegistrationLogMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(RegistrationLog)).Where("date_registered = ?", iDateRegistered).Update(iRegistrationLogMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteRegistrationLog Delete RegistrationLog
func DeleteRegistrationLog(iId int64) (int64, error) {
	row, err := Engine.Id(iId).Delete(new(RegistrationLog))
	return row, err
}

// DeleteRegistrationLogById Delete RegistrationLog via Id
func DeleteRegistrationLogById(iId int64) (err error) {
	var has bool
	var _RegistrationLog = &RegistrationLog{Id: iId}
	if has, err = Engine.Get(_RegistrationLog); (has == true) && (err == nil) {
		if row, err := Engine.Where("ID = ?", iId).Delete(new(RegistrationLog)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteRegistrationLogByEmail Delete RegistrationLog via Email
func DeleteRegistrationLogByEmail(iEmail string) (err error) {
	var has bool
	var _RegistrationLog = &RegistrationLog{Email: iEmail}
	if has, err = Engine.Get(_RegistrationLog); (has == true) && (err == nil) {
		if row, err := Engine.Where("email = ?", iEmail).Delete(new(RegistrationLog)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteRegistrationLogByIp Delete RegistrationLog via Ip
func DeleteRegistrationLogByIp(iIp string) (err error) {
	var has bool
	var _RegistrationLog = &RegistrationLog{Ip: iIp}
	if has, err = Engine.Get(_RegistrationLog); (has == true) && (err == nil) {
		if row, err := Engine.Where("IP = ?", iIp).Delete(new(RegistrationLog)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteRegistrationLogByBlogId Delete RegistrationLog via BlogId
func DeleteRegistrationLogByBlogId(iBlogId int64) (err error) {
	var has bool
	var _RegistrationLog = &RegistrationLog{BlogId: iBlogId}
	if has, err = Engine.Get(_RegistrationLog); (has == true) && (err == nil) {
		if row, err := Engine.Where("blog_id = ?", iBlogId).Delete(new(RegistrationLog)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteRegistrationLogByDateRegistered Delete RegistrationLog via DateRegistered
func DeleteRegistrationLogByDateRegistered(iDateRegistered time.Time) (err error) {
	var has bool
	var _RegistrationLog = &RegistrationLog{DateRegistered: iDateRegistered}
	if has, err = Engine.Get(_RegistrationLog); (has == true) && (err == nil) {
		if row, err := Engine.Where("date_registered = ?", iDateRegistered).Delete(new(RegistrationLog)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
