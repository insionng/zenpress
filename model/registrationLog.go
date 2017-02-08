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

// GetRegistrationLogCountViaId Get RegistrationLog via Id
func GetRegistrationLogCountViaId(iId int64) int64 {
	n, _ := Engine.Where("ID = ?", iId).Count(&RegistrationLog{Id: iId})
	return n
}

// GetRegistrationLogCountViaEmail Get RegistrationLog via Email
func GetRegistrationLogCountViaEmail(iEmail string) int64 {
	n, _ := Engine.Where("email = ?", iEmail).Count(&RegistrationLog{Email: iEmail})
	return n
}

// GetRegistrationLogCountViaIp Get RegistrationLog via Ip
func GetRegistrationLogCountViaIp(iIp string) int64 {
	n, _ := Engine.Where("IP = ?", iIp).Count(&RegistrationLog{Ip: iIp})
	return n
}

// GetRegistrationLogCountViaBlogId Get RegistrationLog via BlogId
func GetRegistrationLogCountViaBlogId(iBlogId int64) int64 {
	n, _ := Engine.Where("blog_id = ?", iBlogId).Count(&RegistrationLog{BlogId: iBlogId})
	return n
}

// GetRegistrationLogCountViaDateRegistered Get RegistrationLog via DateRegistered
func GetRegistrationLogCountViaDateRegistered(iDateRegistered time.Time) int64 {
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

// GetRegistrationLogsViaId Get RegistrationLogs via Id
func GetRegistrationLogsViaId(offset int, limit int, Id_ int64, field string) (*[]*RegistrationLog, error) {
	var _RegistrationLog = new([]*RegistrationLog)
	err := Engine.Table("registration_log").Where("ID = ?", Id_).Limit(limit, offset).Desc(field).Find(_RegistrationLog)
	return _RegistrationLog, err
}

// GetRegistrationLogsViaEmail Get RegistrationLogs via Email
func GetRegistrationLogsViaEmail(offset int, limit int, Email_ string, field string) (*[]*RegistrationLog, error) {
	var _RegistrationLog = new([]*RegistrationLog)
	err := Engine.Table("registration_log").Where("email = ?", Email_).Limit(limit, offset).Desc(field).Find(_RegistrationLog)
	return _RegistrationLog, err
}

// GetRegistrationLogsViaIp Get RegistrationLogs via Ip
func GetRegistrationLogsViaIp(offset int, limit int, Ip_ string, field string) (*[]*RegistrationLog, error) {
	var _RegistrationLog = new([]*RegistrationLog)
	err := Engine.Table("registration_log").Where("IP = ?", Ip_).Limit(limit, offset).Desc(field).Find(_RegistrationLog)
	return _RegistrationLog, err
}

// GetRegistrationLogsViaBlogId Get RegistrationLogs via BlogId
func GetRegistrationLogsViaBlogId(offset int, limit int, BlogId_ int64, field string) (*[]*RegistrationLog, error) {
	var _RegistrationLog = new([]*RegistrationLog)
	err := Engine.Table("registration_log").Where("blog_id = ?", BlogId_).Limit(limit, offset).Desc(field).Find(_RegistrationLog)
	return _RegistrationLog, err
}

// GetRegistrationLogsViaDateRegistered Get RegistrationLogs via DateRegistered
func GetRegistrationLogsViaDateRegistered(offset int, limit int, DateRegistered_ time.Time, field string) (*[]*RegistrationLog, error) {
	var _RegistrationLog = new([]*RegistrationLog)
	err := Engine.Table("registration_log").Where("date_registered = ?", DateRegistered_).Limit(limit, offset).Desc(field).Find(_RegistrationLog)
	return _RegistrationLog, err
}

// HasRegistrationLogViaId Has RegistrationLog via Id
func HasRegistrationLogViaId(iId int64) bool {
	if has, err := Engine.Where("ID = ?", iId).Get(new(RegistrationLog)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasRegistrationLogViaEmail Has RegistrationLog via Email
func HasRegistrationLogViaEmail(iEmail string) bool {
	if has, err := Engine.Where("email = ?", iEmail).Get(new(RegistrationLog)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasRegistrationLogViaIp Has RegistrationLog via Ip
func HasRegistrationLogViaIp(iIp string) bool {
	if has, err := Engine.Where("IP = ?", iIp).Get(new(RegistrationLog)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasRegistrationLogViaBlogId Has RegistrationLog via BlogId
func HasRegistrationLogViaBlogId(iBlogId int64) bool {
	if has, err := Engine.Where("blog_id = ?", iBlogId).Get(new(RegistrationLog)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasRegistrationLogViaDateRegistered Has RegistrationLog via DateRegistered
func HasRegistrationLogViaDateRegistered(iDateRegistered time.Time) bool {
	if has, err := Engine.Where("date_registered = ?", iDateRegistered).Get(new(RegistrationLog)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetRegistrationLogViaId Get RegistrationLog via Id
func GetRegistrationLogViaId(iId int64) (*RegistrationLog, error) {
	var _RegistrationLog = &RegistrationLog{Id: iId}
	has, err := Engine.Get(_RegistrationLog)
	if has {
		return _RegistrationLog, err
	} else {
		return nil, err
	}
}

// GetRegistrationLogViaEmail Get RegistrationLog via Email
func GetRegistrationLogViaEmail(iEmail string) (*RegistrationLog, error) {
	var _RegistrationLog = &RegistrationLog{Email: iEmail}
	has, err := Engine.Get(_RegistrationLog)
	if has {
		return _RegistrationLog, err
	} else {
		return nil, err
	}
}

// GetRegistrationLogViaIp Get RegistrationLog via Ip
func GetRegistrationLogViaIp(iIp string) (*RegistrationLog, error) {
	var _RegistrationLog = &RegistrationLog{Ip: iIp}
	has, err := Engine.Get(_RegistrationLog)
	if has {
		return _RegistrationLog, err
	} else {
		return nil, err
	}
}

// GetRegistrationLogViaBlogId Get RegistrationLog via BlogId
func GetRegistrationLogViaBlogId(iBlogId int64) (*RegistrationLog, error) {
	var _RegistrationLog = &RegistrationLog{BlogId: iBlogId}
	has, err := Engine.Get(_RegistrationLog)
	if has {
		return _RegistrationLog, err
	} else {
		return nil, err
	}
}

// GetRegistrationLogViaDateRegistered Get RegistrationLog via DateRegistered
func GetRegistrationLogViaDateRegistered(iDateRegistered time.Time) (*RegistrationLog, error) {
	var _RegistrationLog = &RegistrationLog{DateRegistered: iDateRegistered}
	has, err := Engine.Get(_RegistrationLog)
	if has {
		return _RegistrationLog, err
	} else {
		return nil, err
	}
}

// SetRegistrationLogViaId Set RegistrationLog via Id
func SetRegistrationLogViaId(iId int64, registration_log *RegistrationLog) (int64, error) {
	registration_log.Id = iId
	return Engine.Insert(registration_log)
}

// SetRegistrationLogViaEmail Set RegistrationLog via Email
func SetRegistrationLogViaEmail(iEmail string, registration_log *RegistrationLog) (int64, error) {
	registration_log.Email = iEmail
	return Engine.Insert(registration_log)
}

// SetRegistrationLogViaIp Set RegistrationLog via Ip
func SetRegistrationLogViaIp(iIp string, registration_log *RegistrationLog) (int64, error) {
	registration_log.Ip = iIp
	return Engine.Insert(registration_log)
}

// SetRegistrationLogViaBlogId Set RegistrationLog via BlogId
func SetRegistrationLogViaBlogId(iBlogId int64, registration_log *RegistrationLog) (int64, error) {
	registration_log.BlogId = iBlogId
	return Engine.Insert(registration_log)
}

// SetRegistrationLogViaDateRegistered Set RegistrationLog via DateRegistered
func SetRegistrationLogViaDateRegistered(iDateRegistered time.Time, registration_log *RegistrationLog) (int64, error) {
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

// PutRegistrationLogViaId Put RegistrationLog via Id
func PutRegistrationLogViaId(Id_ int64, iRegistrationLog *RegistrationLog) (int64, error) {
	row, err := Engine.Update(iRegistrationLog, &RegistrationLog{Id: Id_})
	return row, err
}

// PutRegistrationLogViaEmail Put RegistrationLog via Email
func PutRegistrationLogViaEmail(Email_ string, iRegistrationLog *RegistrationLog) (int64, error) {
	row, err := Engine.Update(iRegistrationLog, &RegistrationLog{Email: Email_})
	return row, err
}

// PutRegistrationLogViaIp Put RegistrationLog via Ip
func PutRegistrationLogViaIp(Ip_ string, iRegistrationLog *RegistrationLog) (int64, error) {
	row, err := Engine.Update(iRegistrationLog, &RegistrationLog{Ip: Ip_})
	return row, err
}

// PutRegistrationLogViaBlogId Put RegistrationLog via BlogId
func PutRegistrationLogViaBlogId(BlogId_ int64, iRegistrationLog *RegistrationLog) (int64, error) {
	row, err := Engine.Update(iRegistrationLog, &RegistrationLog{BlogId: BlogId_})
	return row, err
}

// PutRegistrationLogViaDateRegistered Put RegistrationLog via DateRegistered
func PutRegistrationLogViaDateRegistered(DateRegistered_ time.Time, iRegistrationLog *RegistrationLog) (int64, error) {
	row, err := Engine.Update(iRegistrationLog, &RegistrationLog{DateRegistered: DateRegistered_})
	return row, err
}

// UpdateRegistrationLogViaId via map[string]interface{}{}
func UpdateRegistrationLogViaId(iId int64, iRegistrationLogMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(RegistrationLog)).Where("ID = ?", iId).Update(iRegistrationLogMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateRegistrationLogViaEmail via map[string]interface{}{}
func UpdateRegistrationLogViaEmail(iEmail string, iRegistrationLogMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(RegistrationLog)).Where("email = ?", iEmail).Update(iRegistrationLogMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateRegistrationLogViaIp via map[string]interface{}{}
func UpdateRegistrationLogViaIp(iIp string, iRegistrationLogMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(RegistrationLog)).Where("IP = ?", iIp).Update(iRegistrationLogMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateRegistrationLogViaBlogId via map[string]interface{}{}
func UpdateRegistrationLogViaBlogId(iBlogId int64, iRegistrationLogMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(RegistrationLog)).Where("blog_id = ?", iBlogId).Update(iRegistrationLogMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateRegistrationLogViaDateRegistered via map[string]interface{}{}
func UpdateRegistrationLogViaDateRegistered(iDateRegistered time.Time, iRegistrationLogMap *map[string]interface{}) error {
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

// DeleteRegistrationLogViaId Delete RegistrationLog via Id
func DeleteRegistrationLogViaId(iId int64) (err error) {
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

// DeleteRegistrationLogViaEmail Delete RegistrationLog via Email
func DeleteRegistrationLogViaEmail(iEmail string) (err error) {
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

// DeleteRegistrationLogViaIp Delete RegistrationLog via Ip
func DeleteRegistrationLogViaIp(iIp string) (err error) {
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

// DeleteRegistrationLogViaBlogId Delete RegistrationLog via BlogId
func DeleteRegistrationLogViaBlogId(iBlogId int64) (err error) {
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

// DeleteRegistrationLogViaDateRegistered Delete RegistrationLog via DateRegistered
func DeleteRegistrationLogViaDateRegistered(iDateRegistered time.Time) (err error) {
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
