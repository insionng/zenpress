package model

import (
	"time"
)

type Users struct {
	Id                int64     `xorm:"BIGINT(20)"`
	UserLogin         string    `xorm:"not null default '' index VARCHAR(60)"`
	UserPass          string    `xorm:"not null default '' VARCHAR(255)"`
	UserNicename      string    `xorm:"not null default '' index VARCHAR(50)"`
	UserEmail         string    `xorm:"not null default '' index VARCHAR(100)"`
	UserUrl           string    `xorm:"not null default '' VARCHAR(100)"`
	UserRegistered    time.Time `xorm:"not null default '0000-00-00 00:00:00' DATETIME"`
	UserActivationKey string    `xorm:"not null default '' VARCHAR(255)"`
	UserStatus        int       `xorm:"not null default 0 INT(11)"`
	DisplayName       string    `xorm:"not null default '' VARCHAR(250)"`
	Spam              int       `xorm:"not null default 0 TINYINT(2)"`
	Deleted           int       `xorm:"not null default 0 TINYINT(2)"`
}

// GetUsersesCount Userss' Count
func GetUsersesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&Users{})
	return total, err
}

// GetUsersCountViaId Get Users via Id
func GetUsersCountViaId(iId int64) int64 {
	n, _ := Engine.Where("ID = ?", iId).Count(&Users{Id: iId})
	return n
}

// GetUsersCountViaUserLogin Get Users via UserLogin
func GetUsersCountViaUserLogin(iUserLogin string) int64 {
	n, _ := Engine.Where("user_login = ?", iUserLogin).Count(&Users{UserLogin: iUserLogin})
	return n
}

// GetUsersCountViaUserPass Get Users via UserPass
func GetUsersCountViaUserPass(iUserPass string) int64 {
	n, _ := Engine.Where("user_pass = ?", iUserPass).Count(&Users{UserPass: iUserPass})
	return n
}

// GetUsersCountViaUserNicename Get Users via UserNicename
func GetUsersCountViaUserNicename(iUserNicename string) int64 {
	n, _ := Engine.Where("user_nicename = ?", iUserNicename).Count(&Users{UserNicename: iUserNicename})
	return n
}

// GetUsersCountViaUserEmail Get Users via UserEmail
func GetUsersCountViaUserEmail(iUserEmail string) int64 {
	n, _ := Engine.Where("user_email = ?", iUserEmail).Count(&Users{UserEmail: iUserEmail})
	return n
}

// GetUsersCountViaUserUrl Get Users via UserUrl
func GetUsersCountViaUserUrl(iUserUrl string) int64 {
	n, _ := Engine.Where("user_url = ?", iUserUrl).Count(&Users{UserUrl: iUserUrl})
	return n
}

// GetUsersCountViaUserRegistered Get Users via UserRegistered
func GetUsersCountViaUserRegistered(iUserRegistered time.Time) int64 {
	n, _ := Engine.Where("user_registered = ?", iUserRegistered).Count(&Users{UserRegistered: iUserRegistered})
	return n
}

// GetUsersCountViaUserActivationKey Get Users via UserActivationKey
func GetUsersCountViaUserActivationKey(iUserActivationKey string) int64 {
	n, _ := Engine.Where("user_activation_key = ?", iUserActivationKey).Count(&Users{UserActivationKey: iUserActivationKey})
	return n
}

// GetUsersCountViaUserStatus Get Users via UserStatus
func GetUsersCountViaUserStatus(iUserStatus int) int64 {
	n, _ := Engine.Where("user_status = ?", iUserStatus).Count(&Users{UserStatus: iUserStatus})
	return n
}

// GetUsersCountViaDisplayName Get Users via DisplayName
func GetUsersCountViaDisplayName(iDisplayName string) int64 {
	n, _ := Engine.Where("display_name = ?", iDisplayName).Count(&Users{DisplayName: iDisplayName})
	return n
}

// GetUsersCountViaSpam Get Users via Spam
func GetUsersCountViaSpam(iSpam int) int64 {
	n, _ := Engine.Where("spam = ?", iSpam).Count(&Users{Spam: iSpam})
	return n
}

// GetUsersCountViaDeleted Get Users via Deleted
func GetUsersCountViaDeleted(iDeleted int) int64 {
	n, _ := Engine.Where("deleted = ?", iDeleted).Count(&Users{Deleted: iDeleted})
	return n
}

// GetUsersesByIdAndUserLoginAndUserPass Get Userses via IdAndUserLoginAndUserPass
func GetUsersesByIdAndUserLoginAndUserPass(offset int, limit int, Id_ int64, UserLogin_ string, UserPass_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_login = ? and user_pass = ?", Id_, UserLogin_, UserPass_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserLoginAndUserNicename Get Userses via IdAndUserLoginAndUserNicename
func GetUsersesByIdAndUserLoginAndUserNicename(offset int, limit int, Id_ int64, UserLogin_ string, UserNicename_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_login = ? and user_nicename = ?", Id_, UserLogin_, UserNicename_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserLoginAndUserEmail Get Userses via IdAndUserLoginAndUserEmail
func GetUsersesByIdAndUserLoginAndUserEmail(offset int, limit int, Id_ int64, UserLogin_ string, UserEmail_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_login = ? and user_email = ?", Id_, UserLogin_, UserEmail_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserLoginAndUserUrl Get Userses via IdAndUserLoginAndUserUrl
func GetUsersesByIdAndUserLoginAndUserUrl(offset int, limit int, Id_ int64, UserLogin_ string, UserUrl_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_login = ? and user_url = ?", Id_, UserLogin_, UserUrl_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserLoginAndUserRegistered Get Userses via IdAndUserLoginAndUserRegistered
func GetUsersesByIdAndUserLoginAndUserRegistered(offset int, limit int, Id_ int64, UserLogin_ string, UserRegistered_ time.Time) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_login = ? and user_registered = ?", Id_, UserLogin_, UserRegistered_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserLoginAndUserActivationKey Get Userses via IdAndUserLoginAndUserActivationKey
func GetUsersesByIdAndUserLoginAndUserActivationKey(offset int, limit int, Id_ int64, UserLogin_ string, UserActivationKey_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_login = ? and user_activation_key = ?", Id_, UserLogin_, UserActivationKey_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserLoginAndUserStatus Get Userses via IdAndUserLoginAndUserStatus
func GetUsersesByIdAndUserLoginAndUserStatus(offset int, limit int, Id_ int64, UserLogin_ string, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_login = ? and user_status = ?", Id_, UserLogin_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserLoginAndDisplayName Get Userses via IdAndUserLoginAndDisplayName
func GetUsersesByIdAndUserLoginAndDisplayName(offset int, limit int, Id_ int64, UserLogin_ string, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_login = ? and display_name = ?", Id_, UserLogin_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserLoginAndSpam Get Userses via IdAndUserLoginAndSpam
func GetUsersesByIdAndUserLoginAndSpam(offset int, limit int, Id_ int64, UserLogin_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_login = ? and spam = ?", Id_, UserLogin_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserLoginAndDeleted Get Userses via IdAndUserLoginAndDeleted
func GetUsersesByIdAndUserLoginAndDeleted(offset int, limit int, Id_ int64, UserLogin_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_login = ? and deleted = ?", Id_, UserLogin_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserPassAndUserNicename Get Userses via IdAndUserPassAndUserNicename
func GetUsersesByIdAndUserPassAndUserNicename(offset int, limit int, Id_ int64, UserPass_ string, UserNicename_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_pass = ? and user_nicename = ?", Id_, UserPass_, UserNicename_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserPassAndUserEmail Get Userses via IdAndUserPassAndUserEmail
func GetUsersesByIdAndUserPassAndUserEmail(offset int, limit int, Id_ int64, UserPass_ string, UserEmail_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_pass = ? and user_email = ?", Id_, UserPass_, UserEmail_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserPassAndUserUrl Get Userses via IdAndUserPassAndUserUrl
func GetUsersesByIdAndUserPassAndUserUrl(offset int, limit int, Id_ int64, UserPass_ string, UserUrl_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_pass = ? and user_url = ?", Id_, UserPass_, UserUrl_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserPassAndUserRegistered Get Userses via IdAndUserPassAndUserRegistered
func GetUsersesByIdAndUserPassAndUserRegistered(offset int, limit int, Id_ int64, UserPass_ string, UserRegistered_ time.Time) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_pass = ? and user_registered = ?", Id_, UserPass_, UserRegistered_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserPassAndUserActivationKey Get Userses via IdAndUserPassAndUserActivationKey
func GetUsersesByIdAndUserPassAndUserActivationKey(offset int, limit int, Id_ int64, UserPass_ string, UserActivationKey_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_pass = ? and user_activation_key = ?", Id_, UserPass_, UserActivationKey_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserPassAndUserStatus Get Userses via IdAndUserPassAndUserStatus
func GetUsersesByIdAndUserPassAndUserStatus(offset int, limit int, Id_ int64, UserPass_ string, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_pass = ? and user_status = ?", Id_, UserPass_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserPassAndDisplayName Get Userses via IdAndUserPassAndDisplayName
func GetUsersesByIdAndUserPassAndDisplayName(offset int, limit int, Id_ int64, UserPass_ string, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_pass = ? and display_name = ?", Id_, UserPass_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserPassAndSpam Get Userses via IdAndUserPassAndSpam
func GetUsersesByIdAndUserPassAndSpam(offset int, limit int, Id_ int64, UserPass_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_pass = ? and spam = ?", Id_, UserPass_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserPassAndDeleted Get Userses via IdAndUserPassAndDeleted
func GetUsersesByIdAndUserPassAndDeleted(offset int, limit int, Id_ int64, UserPass_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_pass = ? and deleted = ?", Id_, UserPass_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserNicenameAndUserEmail Get Userses via IdAndUserNicenameAndUserEmail
func GetUsersesByIdAndUserNicenameAndUserEmail(offset int, limit int, Id_ int64, UserNicename_ string, UserEmail_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_nicename = ? and user_email = ?", Id_, UserNicename_, UserEmail_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserNicenameAndUserUrl Get Userses via IdAndUserNicenameAndUserUrl
func GetUsersesByIdAndUserNicenameAndUserUrl(offset int, limit int, Id_ int64, UserNicename_ string, UserUrl_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_nicename = ? and user_url = ?", Id_, UserNicename_, UserUrl_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserNicenameAndUserRegistered Get Userses via IdAndUserNicenameAndUserRegistered
func GetUsersesByIdAndUserNicenameAndUserRegistered(offset int, limit int, Id_ int64, UserNicename_ string, UserRegistered_ time.Time) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_nicename = ? and user_registered = ?", Id_, UserNicename_, UserRegistered_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserNicenameAndUserActivationKey Get Userses via IdAndUserNicenameAndUserActivationKey
func GetUsersesByIdAndUserNicenameAndUserActivationKey(offset int, limit int, Id_ int64, UserNicename_ string, UserActivationKey_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_nicename = ? and user_activation_key = ?", Id_, UserNicename_, UserActivationKey_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserNicenameAndUserStatus Get Userses via IdAndUserNicenameAndUserStatus
func GetUsersesByIdAndUserNicenameAndUserStatus(offset int, limit int, Id_ int64, UserNicename_ string, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_nicename = ? and user_status = ?", Id_, UserNicename_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserNicenameAndDisplayName Get Userses via IdAndUserNicenameAndDisplayName
func GetUsersesByIdAndUserNicenameAndDisplayName(offset int, limit int, Id_ int64, UserNicename_ string, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_nicename = ? and display_name = ?", Id_, UserNicename_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserNicenameAndSpam Get Userses via IdAndUserNicenameAndSpam
func GetUsersesByIdAndUserNicenameAndSpam(offset int, limit int, Id_ int64, UserNicename_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_nicename = ? and spam = ?", Id_, UserNicename_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserNicenameAndDeleted Get Userses via IdAndUserNicenameAndDeleted
func GetUsersesByIdAndUserNicenameAndDeleted(offset int, limit int, Id_ int64, UserNicename_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_nicename = ? and deleted = ?", Id_, UserNicename_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserEmailAndUserUrl Get Userses via IdAndUserEmailAndUserUrl
func GetUsersesByIdAndUserEmailAndUserUrl(offset int, limit int, Id_ int64, UserEmail_ string, UserUrl_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_email = ? and user_url = ?", Id_, UserEmail_, UserUrl_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserEmailAndUserRegistered Get Userses via IdAndUserEmailAndUserRegistered
func GetUsersesByIdAndUserEmailAndUserRegistered(offset int, limit int, Id_ int64, UserEmail_ string, UserRegistered_ time.Time) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_email = ? and user_registered = ?", Id_, UserEmail_, UserRegistered_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserEmailAndUserActivationKey Get Userses via IdAndUserEmailAndUserActivationKey
func GetUsersesByIdAndUserEmailAndUserActivationKey(offset int, limit int, Id_ int64, UserEmail_ string, UserActivationKey_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_email = ? and user_activation_key = ?", Id_, UserEmail_, UserActivationKey_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserEmailAndUserStatus Get Userses via IdAndUserEmailAndUserStatus
func GetUsersesByIdAndUserEmailAndUserStatus(offset int, limit int, Id_ int64, UserEmail_ string, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_email = ? and user_status = ?", Id_, UserEmail_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserEmailAndDisplayName Get Userses via IdAndUserEmailAndDisplayName
func GetUsersesByIdAndUserEmailAndDisplayName(offset int, limit int, Id_ int64, UserEmail_ string, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_email = ? and display_name = ?", Id_, UserEmail_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserEmailAndSpam Get Userses via IdAndUserEmailAndSpam
func GetUsersesByIdAndUserEmailAndSpam(offset int, limit int, Id_ int64, UserEmail_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_email = ? and spam = ?", Id_, UserEmail_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserEmailAndDeleted Get Userses via IdAndUserEmailAndDeleted
func GetUsersesByIdAndUserEmailAndDeleted(offset int, limit int, Id_ int64, UserEmail_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_email = ? and deleted = ?", Id_, UserEmail_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserUrlAndUserRegistered Get Userses via IdAndUserUrlAndUserRegistered
func GetUsersesByIdAndUserUrlAndUserRegistered(offset int, limit int, Id_ int64, UserUrl_ string, UserRegistered_ time.Time) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_url = ? and user_registered = ?", Id_, UserUrl_, UserRegistered_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserUrlAndUserActivationKey Get Userses via IdAndUserUrlAndUserActivationKey
func GetUsersesByIdAndUserUrlAndUserActivationKey(offset int, limit int, Id_ int64, UserUrl_ string, UserActivationKey_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_url = ? and user_activation_key = ?", Id_, UserUrl_, UserActivationKey_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserUrlAndUserStatus Get Userses via IdAndUserUrlAndUserStatus
func GetUsersesByIdAndUserUrlAndUserStatus(offset int, limit int, Id_ int64, UserUrl_ string, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_url = ? and user_status = ?", Id_, UserUrl_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserUrlAndDisplayName Get Userses via IdAndUserUrlAndDisplayName
func GetUsersesByIdAndUserUrlAndDisplayName(offset int, limit int, Id_ int64, UserUrl_ string, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_url = ? and display_name = ?", Id_, UserUrl_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserUrlAndSpam Get Userses via IdAndUserUrlAndSpam
func GetUsersesByIdAndUserUrlAndSpam(offset int, limit int, Id_ int64, UserUrl_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_url = ? and spam = ?", Id_, UserUrl_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserUrlAndDeleted Get Userses via IdAndUserUrlAndDeleted
func GetUsersesByIdAndUserUrlAndDeleted(offset int, limit int, Id_ int64, UserUrl_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_url = ? and deleted = ?", Id_, UserUrl_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserRegisteredAndUserActivationKey Get Userses via IdAndUserRegisteredAndUserActivationKey
func GetUsersesByIdAndUserRegisteredAndUserActivationKey(offset int, limit int, Id_ int64, UserRegistered_ time.Time, UserActivationKey_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_registered = ? and user_activation_key = ?", Id_, UserRegistered_, UserActivationKey_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserRegisteredAndUserStatus Get Userses via IdAndUserRegisteredAndUserStatus
func GetUsersesByIdAndUserRegisteredAndUserStatus(offset int, limit int, Id_ int64, UserRegistered_ time.Time, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_registered = ? and user_status = ?", Id_, UserRegistered_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserRegisteredAndDisplayName Get Userses via IdAndUserRegisteredAndDisplayName
func GetUsersesByIdAndUserRegisteredAndDisplayName(offset int, limit int, Id_ int64, UserRegistered_ time.Time, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_registered = ? and display_name = ?", Id_, UserRegistered_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserRegisteredAndSpam Get Userses via IdAndUserRegisteredAndSpam
func GetUsersesByIdAndUserRegisteredAndSpam(offset int, limit int, Id_ int64, UserRegistered_ time.Time, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_registered = ? and spam = ?", Id_, UserRegistered_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserRegisteredAndDeleted Get Userses via IdAndUserRegisteredAndDeleted
func GetUsersesByIdAndUserRegisteredAndDeleted(offset int, limit int, Id_ int64, UserRegistered_ time.Time, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_registered = ? and deleted = ?", Id_, UserRegistered_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserActivationKeyAndUserStatus Get Userses via IdAndUserActivationKeyAndUserStatus
func GetUsersesByIdAndUserActivationKeyAndUserStatus(offset int, limit int, Id_ int64, UserActivationKey_ string, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_activation_key = ? and user_status = ?", Id_, UserActivationKey_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserActivationKeyAndDisplayName Get Userses via IdAndUserActivationKeyAndDisplayName
func GetUsersesByIdAndUserActivationKeyAndDisplayName(offset int, limit int, Id_ int64, UserActivationKey_ string, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_activation_key = ? and display_name = ?", Id_, UserActivationKey_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserActivationKeyAndSpam Get Userses via IdAndUserActivationKeyAndSpam
func GetUsersesByIdAndUserActivationKeyAndSpam(offset int, limit int, Id_ int64, UserActivationKey_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_activation_key = ? and spam = ?", Id_, UserActivationKey_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserActivationKeyAndDeleted Get Userses via IdAndUserActivationKeyAndDeleted
func GetUsersesByIdAndUserActivationKeyAndDeleted(offset int, limit int, Id_ int64, UserActivationKey_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_activation_key = ? and deleted = ?", Id_, UserActivationKey_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserStatusAndDisplayName Get Userses via IdAndUserStatusAndDisplayName
func GetUsersesByIdAndUserStatusAndDisplayName(offset int, limit int, Id_ int64, UserStatus_ int, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_status = ? and display_name = ?", Id_, UserStatus_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserStatusAndSpam Get Userses via IdAndUserStatusAndSpam
func GetUsersesByIdAndUserStatusAndSpam(offset int, limit int, Id_ int64, UserStatus_ int, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_status = ? and spam = ?", Id_, UserStatus_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserStatusAndDeleted Get Userses via IdAndUserStatusAndDeleted
func GetUsersesByIdAndUserStatusAndDeleted(offset int, limit int, Id_ int64, UserStatus_ int, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_status = ? and deleted = ?", Id_, UserStatus_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndDisplayNameAndSpam Get Userses via IdAndDisplayNameAndSpam
func GetUsersesByIdAndDisplayNameAndSpam(offset int, limit int, Id_ int64, DisplayName_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and display_name = ? and spam = ?", Id_, DisplayName_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndDisplayNameAndDeleted Get Userses via IdAndDisplayNameAndDeleted
func GetUsersesByIdAndDisplayNameAndDeleted(offset int, limit int, Id_ int64, DisplayName_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and display_name = ? and deleted = ?", Id_, DisplayName_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndSpamAndDeleted Get Userses via IdAndSpamAndDeleted
func GetUsersesByIdAndSpamAndDeleted(offset int, limit int, Id_ int64, Spam_ int, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and spam = ? and deleted = ?", Id_, Spam_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserPassAndUserNicename Get Userses via UserLoginAndUserPassAndUserNicename
func GetUsersesByUserLoginAndUserPassAndUserNicename(offset int, limit int, UserLogin_ string, UserPass_ string, UserNicename_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_pass = ? and user_nicename = ?", UserLogin_, UserPass_, UserNicename_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserPassAndUserEmail Get Userses via UserLoginAndUserPassAndUserEmail
func GetUsersesByUserLoginAndUserPassAndUserEmail(offset int, limit int, UserLogin_ string, UserPass_ string, UserEmail_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_pass = ? and user_email = ?", UserLogin_, UserPass_, UserEmail_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserPassAndUserUrl Get Userses via UserLoginAndUserPassAndUserUrl
func GetUsersesByUserLoginAndUserPassAndUserUrl(offset int, limit int, UserLogin_ string, UserPass_ string, UserUrl_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_pass = ? and user_url = ?", UserLogin_, UserPass_, UserUrl_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserPassAndUserRegistered Get Userses via UserLoginAndUserPassAndUserRegistered
func GetUsersesByUserLoginAndUserPassAndUserRegistered(offset int, limit int, UserLogin_ string, UserPass_ string, UserRegistered_ time.Time) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_pass = ? and user_registered = ?", UserLogin_, UserPass_, UserRegistered_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserPassAndUserActivationKey Get Userses via UserLoginAndUserPassAndUserActivationKey
func GetUsersesByUserLoginAndUserPassAndUserActivationKey(offset int, limit int, UserLogin_ string, UserPass_ string, UserActivationKey_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_pass = ? and user_activation_key = ?", UserLogin_, UserPass_, UserActivationKey_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserPassAndUserStatus Get Userses via UserLoginAndUserPassAndUserStatus
func GetUsersesByUserLoginAndUserPassAndUserStatus(offset int, limit int, UserLogin_ string, UserPass_ string, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_pass = ? and user_status = ?", UserLogin_, UserPass_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserPassAndDisplayName Get Userses via UserLoginAndUserPassAndDisplayName
func GetUsersesByUserLoginAndUserPassAndDisplayName(offset int, limit int, UserLogin_ string, UserPass_ string, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_pass = ? and display_name = ?", UserLogin_, UserPass_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserPassAndSpam Get Userses via UserLoginAndUserPassAndSpam
func GetUsersesByUserLoginAndUserPassAndSpam(offset int, limit int, UserLogin_ string, UserPass_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_pass = ? and spam = ?", UserLogin_, UserPass_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserPassAndDeleted Get Userses via UserLoginAndUserPassAndDeleted
func GetUsersesByUserLoginAndUserPassAndDeleted(offset int, limit int, UserLogin_ string, UserPass_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_pass = ? and deleted = ?", UserLogin_, UserPass_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserNicenameAndUserEmail Get Userses via UserLoginAndUserNicenameAndUserEmail
func GetUsersesByUserLoginAndUserNicenameAndUserEmail(offset int, limit int, UserLogin_ string, UserNicename_ string, UserEmail_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_nicename = ? and user_email = ?", UserLogin_, UserNicename_, UserEmail_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserNicenameAndUserUrl Get Userses via UserLoginAndUserNicenameAndUserUrl
func GetUsersesByUserLoginAndUserNicenameAndUserUrl(offset int, limit int, UserLogin_ string, UserNicename_ string, UserUrl_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_nicename = ? and user_url = ?", UserLogin_, UserNicename_, UserUrl_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserNicenameAndUserRegistered Get Userses via UserLoginAndUserNicenameAndUserRegistered
func GetUsersesByUserLoginAndUserNicenameAndUserRegistered(offset int, limit int, UserLogin_ string, UserNicename_ string, UserRegistered_ time.Time) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_nicename = ? and user_registered = ?", UserLogin_, UserNicename_, UserRegistered_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserNicenameAndUserActivationKey Get Userses via UserLoginAndUserNicenameAndUserActivationKey
func GetUsersesByUserLoginAndUserNicenameAndUserActivationKey(offset int, limit int, UserLogin_ string, UserNicename_ string, UserActivationKey_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_nicename = ? and user_activation_key = ?", UserLogin_, UserNicename_, UserActivationKey_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserNicenameAndUserStatus Get Userses via UserLoginAndUserNicenameAndUserStatus
func GetUsersesByUserLoginAndUserNicenameAndUserStatus(offset int, limit int, UserLogin_ string, UserNicename_ string, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_nicename = ? and user_status = ?", UserLogin_, UserNicename_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserNicenameAndDisplayName Get Userses via UserLoginAndUserNicenameAndDisplayName
func GetUsersesByUserLoginAndUserNicenameAndDisplayName(offset int, limit int, UserLogin_ string, UserNicename_ string, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_nicename = ? and display_name = ?", UserLogin_, UserNicename_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserNicenameAndSpam Get Userses via UserLoginAndUserNicenameAndSpam
func GetUsersesByUserLoginAndUserNicenameAndSpam(offset int, limit int, UserLogin_ string, UserNicename_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_nicename = ? and spam = ?", UserLogin_, UserNicename_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserNicenameAndDeleted Get Userses via UserLoginAndUserNicenameAndDeleted
func GetUsersesByUserLoginAndUserNicenameAndDeleted(offset int, limit int, UserLogin_ string, UserNicename_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_nicename = ? and deleted = ?", UserLogin_, UserNicename_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserEmailAndUserUrl Get Userses via UserLoginAndUserEmailAndUserUrl
func GetUsersesByUserLoginAndUserEmailAndUserUrl(offset int, limit int, UserLogin_ string, UserEmail_ string, UserUrl_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_email = ? and user_url = ?", UserLogin_, UserEmail_, UserUrl_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserEmailAndUserRegistered Get Userses via UserLoginAndUserEmailAndUserRegistered
func GetUsersesByUserLoginAndUserEmailAndUserRegistered(offset int, limit int, UserLogin_ string, UserEmail_ string, UserRegistered_ time.Time) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_email = ? and user_registered = ?", UserLogin_, UserEmail_, UserRegistered_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserEmailAndUserActivationKey Get Userses via UserLoginAndUserEmailAndUserActivationKey
func GetUsersesByUserLoginAndUserEmailAndUserActivationKey(offset int, limit int, UserLogin_ string, UserEmail_ string, UserActivationKey_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_email = ? and user_activation_key = ?", UserLogin_, UserEmail_, UserActivationKey_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserEmailAndUserStatus Get Userses via UserLoginAndUserEmailAndUserStatus
func GetUsersesByUserLoginAndUserEmailAndUserStatus(offset int, limit int, UserLogin_ string, UserEmail_ string, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_email = ? and user_status = ?", UserLogin_, UserEmail_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserEmailAndDisplayName Get Userses via UserLoginAndUserEmailAndDisplayName
func GetUsersesByUserLoginAndUserEmailAndDisplayName(offset int, limit int, UserLogin_ string, UserEmail_ string, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_email = ? and display_name = ?", UserLogin_, UserEmail_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserEmailAndSpam Get Userses via UserLoginAndUserEmailAndSpam
func GetUsersesByUserLoginAndUserEmailAndSpam(offset int, limit int, UserLogin_ string, UserEmail_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_email = ? and spam = ?", UserLogin_, UserEmail_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserEmailAndDeleted Get Userses via UserLoginAndUserEmailAndDeleted
func GetUsersesByUserLoginAndUserEmailAndDeleted(offset int, limit int, UserLogin_ string, UserEmail_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_email = ? and deleted = ?", UserLogin_, UserEmail_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserUrlAndUserRegistered Get Userses via UserLoginAndUserUrlAndUserRegistered
func GetUsersesByUserLoginAndUserUrlAndUserRegistered(offset int, limit int, UserLogin_ string, UserUrl_ string, UserRegistered_ time.Time) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_url = ? and user_registered = ?", UserLogin_, UserUrl_, UserRegistered_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserUrlAndUserActivationKey Get Userses via UserLoginAndUserUrlAndUserActivationKey
func GetUsersesByUserLoginAndUserUrlAndUserActivationKey(offset int, limit int, UserLogin_ string, UserUrl_ string, UserActivationKey_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_url = ? and user_activation_key = ?", UserLogin_, UserUrl_, UserActivationKey_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserUrlAndUserStatus Get Userses via UserLoginAndUserUrlAndUserStatus
func GetUsersesByUserLoginAndUserUrlAndUserStatus(offset int, limit int, UserLogin_ string, UserUrl_ string, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_url = ? and user_status = ?", UserLogin_, UserUrl_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserUrlAndDisplayName Get Userses via UserLoginAndUserUrlAndDisplayName
func GetUsersesByUserLoginAndUserUrlAndDisplayName(offset int, limit int, UserLogin_ string, UserUrl_ string, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_url = ? and display_name = ?", UserLogin_, UserUrl_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserUrlAndSpam Get Userses via UserLoginAndUserUrlAndSpam
func GetUsersesByUserLoginAndUserUrlAndSpam(offset int, limit int, UserLogin_ string, UserUrl_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_url = ? and spam = ?", UserLogin_, UserUrl_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserUrlAndDeleted Get Userses via UserLoginAndUserUrlAndDeleted
func GetUsersesByUserLoginAndUserUrlAndDeleted(offset int, limit int, UserLogin_ string, UserUrl_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_url = ? and deleted = ?", UserLogin_, UserUrl_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserRegisteredAndUserActivationKey Get Userses via UserLoginAndUserRegisteredAndUserActivationKey
func GetUsersesByUserLoginAndUserRegisteredAndUserActivationKey(offset int, limit int, UserLogin_ string, UserRegistered_ time.Time, UserActivationKey_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_registered = ? and user_activation_key = ?", UserLogin_, UserRegistered_, UserActivationKey_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserRegisteredAndUserStatus Get Userses via UserLoginAndUserRegisteredAndUserStatus
func GetUsersesByUserLoginAndUserRegisteredAndUserStatus(offset int, limit int, UserLogin_ string, UserRegistered_ time.Time, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_registered = ? and user_status = ?", UserLogin_, UserRegistered_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserRegisteredAndDisplayName Get Userses via UserLoginAndUserRegisteredAndDisplayName
func GetUsersesByUserLoginAndUserRegisteredAndDisplayName(offset int, limit int, UserLogin_ string, UserRegistered_ time.Time, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_registered = ? and display_name = ?", UserLogin_, UserRegistered_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserRegisteredAndSpam Get Userses via UserLoginAndUserRegisteredAndSpam
func GetUsersesByUserLoginAndUserRegisteredAndSpam(offset int, limit int, UserLogin_ string, UserRegistered_ time.Time, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_registered = ? and spam = ?", UserLogin_, UserRegistered_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserRegisteredAndDeleted Get Userses via UserLoginAndUserRegisteredAndDeleted
func GetUsersesByUserLoginAndUserRegisteredAndDeleted(offset int, limit int, UserLogin_ string, UserRegistered_ time.Time, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_registered = ? and deleted = ?", UserLogin_, UserRegistered_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserActivationKeyAndUserStatus Get Userses via UserLoginAndUserActivationKeyAndUserStatus
func GetUsersesByUserLoginAndUserActivationKeyAndUserStatus(offset int, limit int, UserLogin_ string, UserActivationKey_ string, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_activation_key = ? and user_status = ?", UserLogin_, UserActivationKey_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserActivationKeyAndDisplayName Get Userses via UserLoginAndUserActivationKeyAndDisplayName
func GetUsersesByUserLoginAndUserActivationKeyAndDisplayName(offset int, limit int, UserLogin_ string, UserActivationKey_ string, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_activation_key = ? and display_name = ?", UserLogin_, UserActivationKey_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserActivationKeyAndSpam Get Userses via UserLoginAndUserActivationKeyAndSpam
func GetUsersesByUserLoginAndUserActivationKeyAndSpam(offset int, limit int, UserLogin_ string, UserActivationKey_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_activation_key = ? and spam = ?", UserLogin_, UserActivationKey_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserActivationKeyAndDeleted Get Userses via UserLoginAndUserActivationKeyAndDeleted
func GetUsersesByUserLoginAndUserActivationKeyAndDeleted(offset int, limit int, UserLogin_ string, UserActivationKey_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_activation_key = ? and deleted = ?", UserLogin_, UserActivationKey_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserStatusAndDisplayName Get Userses via UserLoginAndUserStatusAndDisplayName
func GetUsersesByUserLoginAndUserStatusAndDisplayName(offset int, limit int, UserLogin_ string, UserStatus_ int, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_status = ? and display_name = ?", UserLogin_, UserStatus_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserStatusAndSpam Get Userses via UserLoginAndUserStatusAndSpam
func GetUsersesByUserLoginAndUserStatusAndSpam(offset int, limit int, UserLogin_ string, UserStatus_ int, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_status = ? and spam = ?", UserLogin_, UserStatus_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserStatusAndDeleted Get Userses via UserLoginAndUserStatusAndDeleted
func GetUsersesByUserLoginAndUserStatusAndDeleted(offset int, limit int, UserLogin_ string, UserStatus_ int, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_status = ? and deleted = ?", UserLogin_, UserStatus_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndDisplayNameAndSpam Get Userses via UserLoginAndDisplayNameAndSpam
func GetUsersesByUserLoginAndDisplayNameAndSpam(offset int, limit int, UserLogin_ string, DisplayName_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and display_name = ? and spam = ?", UserLogin_, DisplayName_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndDisplayNameAndDeleted Get Userses via UserLoginAndDisplayNameAndDeleted
func GetUsersesByUserLoginAndDisplayNameAndDeleted(offset int, limit int, UserLogin_ string, DisplayName_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and display_name = ? and deleted = ?", UserLogin_, DisplayName_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndSpamAndDeleted Get Userses via UserLoginAndSpamAndDeleted
func GetUsersesByUserLoginAndSpamAndDeleted(offset int, limit int, UserLogin_ string, Spam_ int, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and spam = ? and deleted = ?", UserLogin_, Spam_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserNicenameAndUserEmail Get Userses via UserPassAndUserNicenameAndUserEmail
func GetUsersesByUserPassAndUserNicenameAndUserEmail(offset int, limit int, UserPass_ string, UserNicename_ string, UserEmail_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_nicename = ? and user_email = ?", UserPass_, UserNicename_, UserEmail_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserNicenameAndUserUrl Get Userses via UserPassAndUserNicenameAndUserUrl
func GetUsersesByUserPassAndUserNicenameAndUserUrl(offset int, limit int, UserPass_ string, UserNicename_ string, UserUrl_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_nicename = ? and user_url = ?", UserPass_, UserNicename_, UserUrl_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserNicenameAndUserRegistered Get Userses via UserPassAndUserNicenameAndUserRegistered
func GetUsersesByUserPassAndUserNicenameAndUserRegistered(offset int, limit int, UserPass_ string, UserNicename_ string, UserRegistered_ time.Time) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_nicename = ? and user_registered = ?", UserPass_, UserNicename_, UserRegistered_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserNicenameAndUserActivationKey Get Userses via UserPassAndUserNicenameAndUserActivationKey
func GetUsersesByUserPassAndUserNicenameAndUserActivationKey(offset int, limit int, UserPass_ string, UserNicename_ string, UserActivationKey_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_nicename = ? and user_activation_key = ?", UserPass_, UserNicename_, UserActivationKey_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserNicenameAndUserStatus Get Userses via UserPassAndUserNicenameAndUserStatus
func GetUsersesByUserPassAndUserNicenameAndUserStatus(offset int, limit int, UserPass_ string, UserNicename_ string, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_nicename = ? and user_status = ?", UserPass_, UserNicename_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserNicenameAndDisplayName Get Userses via UserPassAndUserNicenameAndDisplayName
func GetUsersesByUserPassAndUserNicenameAndDisplayName(offset int, limit int, UserPass_ string, UserNicename_ string, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_nicename = ? and display_name = ?", UserPass_, UserNicename_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserNicenameAndSpam Get Userses via UserPassAndUserNicenameAndSpam
func GetUsersesByUserPassAndUserNicenameAndSpam(offset int, limit int, UserPass_ string, UserNicename_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_nicename = ? and spam = ?", UserPass_, UserNicename_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserNicenameAndDeleted Get Userses via UserPassAndUserNicenameAndDeleted
func GetUsersesByUserPassAndUserNicenameAndDeleted(offset int, limit int, UserPass_ string, UserNicename_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_nicename = ? and deleted = ?", UserPass_, UserNicename_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserEmailAndUserUrl Get Userses via UserPassAndUserEmailAndUserUrl
func GetUsersesByUserPassAndUserEmailAndUserUrl(offset int, limit int, UserPass_ string, UserEmail_ string, UserUrl_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_email = ? and user_url = ?", UserPass_, UserEmail_, UserUrl_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserEmailAndUserRegistered Get Userses via UserPassAndUserEmailAndUserRegistered
func GetUsersesByUserPassAndUserEmailAndUserRegistered(offset int, limit int, UserPass_ string, UserEmail_ string, UserRegistered_ time.Time) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_email = ? and user_registered = ?", UserPass_, UserEmail_, UserRegistered_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserEmailAndUserActivationKey Get Userses via UserPassAndUserEmailAndUserActivationKey
func GetUsersesByUserPassAndUserEmailAndUserActivationKey(offset int, limit int, UserPass_ string, UserEmail_ string, UserActivationKey_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_email = ? and user_activation_key = ?", UserPass_, UserEmail_, UserActivationKey_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserEmailAndUserStatus Get Userses via UserPassAndUserEmailAndUserStatus
func GetUsersesByUserPassAndUserEmailAndUserStatus(offset int, limit int, UserPass_ string, UserEmail_ string, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_email = ? and user_status = ?", UserPass_, UserEmail_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserEmailAndDisplayName Get Userses via UserPassAndUserEmailAndDisplayName
func GetUsersesByUserPassAndUserEmailAndDisplayName(offset int, limit int, UserPass_ string, UserEmail_ string, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_email = ? and display_name = ?", UserPass_, UserEmail_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserEmailAndSpam Get Userses via UserPassAndUserEmailAndSpam
func GetUsersesByUserPassAndUserEmailAndSpam(offset int, limit int, UserPass_ string, UserEmail_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_email = ? and spam = ?", UserPass_, UserEmail_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserEmailAndDeleted Get Userses via UserPassAndUserEmailAndDeleted
func GetUsersesByUserPassAndUserEmailAndDeleted(offset int, limit int, UserPass_ string, UserEmail_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_email = ? and deleted = ?", UserPass_, UserEmail_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserUrlAndUserRegistered Get Userses via UserPassAndUserUrlAndUserRegistered
func GetUsersesByUserPassAndUserUrlAndUserRegistered(offset int, limit int, UserPass_ string, UserUrl_ string, UserRegistered_ time.Time) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_url = ? and user_registered = ?", UserPass_, UserUrl_, UserRegistered_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserUrlAndUserActivationKey Get Userses via UserPassAndUserUrlAndUserActivationKey
func GetUsersesByUserPassAndUserUrlAndUserActivationKey(offset int, limit int, UserPass_ string, UserUrl_ string, UserActivationKey_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_url = ? and user_activation_key = ?", UserPass_, UserUrl_, UserActivationKey_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserUrlAndUserStatus Get Userses via UserPassAndUserUrlAndUserStatus
func GetUsersesByUserPassAndUserUrlAndUserStatus(offset int, limit int, UserPass_ string, UserUrl_ string, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_url = ? and user_status = ?", UserPass_, UserUrl_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserUrlAndDisplayName Get Userses via UserPassAndUserUrlAndDisplayName
func GetUsersesByUserPassAndUserUrlAndDisplayName(offset int, limit int, UserPass_ string, UserUrl_ string, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_url = ? and display_name = ?", UserPass_, UserUrl_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserUrlAndSpam Get Userses via UserPassAndUserUrlAndSpam
func GetUsersesByUserPassAndUserUrlAndSpam(offset int, limit int, UserPass_ string, UserUrl_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_url = ? and spam = ?", UserPass_, UserUrl_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserUrlAndDeleted Get Userses via UserPassAndUserUrlAndDeleted
func GetUsersesByUserPassAndUserUrlAndDeleted(offset int, limit int, UserPass_ string, UserUrl_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_url = ? and deleted = ?", UserPass_, UserUrl_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserRegisteredAndUserActivationKey Get Userses via UserPassAndUserRegisteredAndUserActivationKey
func GetUsersesByUserPassAndUserRegisteredAndUserActivationKey(offset int, limit int, UserPass_ string, UserRegistered_ time.Time, UserActivationKey_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_registered = ? and user_activation_key = ?", UserPass_, UserRegistered_, UserActivationKey_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserRegisteredAndUserStatus Get Userses via UserPassAndUserRegisteredAndUserStatus
func GetUsersesByUserPassAndUserRegisteredAndUserStatus(offset int, limit int, UserPass_ string, UserRegistered_ time.Time, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_registered = ? and user_status = ?", UserPass_, UserRegistered_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserRegisteredAndDisplayName Get Userses via UserPassAndUserRegisteredAndDisplayName
func GetUsersesByUserPassAndUserRegisteredAndDisplayName(offset int, limit int, UserPass_ string, UserRegistered_ time.Time, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_registered = ? and display_name = ?", UserPass_, UserRegistered_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserRegisteredAndSpam Get Userses via UserPassAndUserRegisteredAndSpam
func GetUsersesByUserPassAndUserRegisteredAndSpam(offset int, limit int, UserPass_ string, UserRegistered_ time.Time, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_registered = ? and spam = ?", UserPass_, UserRegistered_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserRegisteredAndDeleted Get Userses via UserPassAndUserRegisteredAndDeleted
func GetUsersesByUserPassAndUserRegisteredAndDeleted(offset int, limit int, UserPass_ string, UserRegistered_ time.Time, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_registered = ? and deleted = ?", UserPass_, UserRegistered_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserActivationKeyAndUserStatus Get Userses via UserPassAndUserActivationKeyAndUserStatus
func GetUsersesByUserPassAndUserActivationKeyAndUserStatus(offset int, limit int, UserPass_ string, UserActivationKey_ string, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_activation_key = ? and user_status = ?", UserPass_, UserActivationKey_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserActivationKeyAndDisplayName Get Userses via UserPassAndUserActivationKeyAndDisplayName
func GetUsersesByUserPassAndUserActivationKeyAndDisplayName(offset int, limit int, UserPass_ string, UserActivationKey_ string, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_activation_key = ? and display_name = ?", UserPass_, UserActivationKey_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserActivationKeyAndSpam Get Userses via UserPassAndUserActivationKeyAndSpam
func GetUsersesByUserPassAndUserActivationKeyAndSpam(offset int, limit int, UserPass_ string, UserActivationKey_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_activation_key = ? and spam = ?", UserPass_, UserActivationKey_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserActivationKeyAndDeleted Get Userses via UserPassAndUserActivationKeyAndDeleted
func GetUsersesByUserPassAndUserActivationKeyAndDeleted(offset int, limit int, UserPass_ string, UserActivationKey_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_activation_key = ? and deleted = ?", UserPass_, UserActivationKey_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserStatusAndDisplayName Get Userses via UserPassAndUserStatusAndDisplayName
func GetUsersesByUserPassAndUserStatusAndDisplayName(offset int, limit int, UserPass_ string, UserStatus_ int, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_status = ? and display_name = ?", UserPass_, UserStatus_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserStatusAndSpam Get Userses via UserPassAndUserStatusAndSpam
func GetUsersesByUserPassAndUserStatusAndSpam(offset int, limit int, UserPass_ string, UserStatus_ int, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_status = ? and spam = ?", UserPass_, UserStatus_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserStatusAndDeleted Get Userses via UserPassAndUserStatusAndDeleted
func GetUsersesByUserPassAndUserStatusAndDeleted(offset int, limit int, UserPass_ string, UserStatus_ int, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_status = ? and deleted = ?", UserPass_, UserStatus_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndDisplayNameAndSpam Get Userses via UserPassAndDisplayNameAndSpam
func GetUsersesByUserPassAndDisplayNameAndSpam(offset int, limit int, UserPass_ string, DisplayName_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and display_name = ? and spam = ?", UserPass_, DisplayName_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndDisplayNameAndDeleted Get Userses via UserPassAndDisplayNameAndDeleted
func GetUsersesByUserPassAndDisplayNameAndDeleted(offset int, limit int, UserPass_ string, DisplayName_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and display_name = ? and deleted = ?", UserPass_, DisplayName_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndSpamAndDeleted Get Userses via UserPassAndSpamAndDeleted
func GetUsersesByUserPassAndSpamAndDeleted(offset int, limit int, UserPass_ string, Spam_ int, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and spam = ? and deleted = ?", UserPass_, Spam_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndUserEmailAndUserUrl Get Userses via UserNicenameAndUserEmailAndUserUrl
func GetUsersesByUserNicenameAndUserEmailAndUserUrl(offset int, limit int, UserNicename_ string, UserEmail_ string, UserUrl_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and user_email = ? and user_url = ?", UserNicename_, UserEmail_, UserUrl_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndUserEmailAndUserRegistered Get Userses via UserNicenameAndUserEmailAndUserRegistered
func GetUsersesByUserNicenameAndUserEmailAndUserRegistered(offset int, limit int, UserNicename_ string, UserEmail_ string, UserRegistered_ time.Time) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and user_email = ? and user_registered = ?", UserNicename_, UserEmail_, UserRegistered_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndUserEmailAndUserActivationKey Get Userses via UserNicenameAndUserEmailAndUserActivationKey
func GetUsersesByUserNicenameAndUserEmailAndUserActivationKey(offset int, limit int, UserNicename_ string, UserEmail_ string, UserActivationKey_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and user_email = ? and user_activation_key = ?", UserNicename_, UserEmail_, UserActivationKey_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndUserEmailAndUserStatus Get Userses via UserNicenameAndUserEmailAndUserStatus
func GetUsersesByUserNicenameAndUserEmailAndUserStatus(offset int, limit int, UserNicename_ string, UserEmail_ string, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and user_email = ? and user_status = ?", UserNicename_, UserEmail_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndUserEmailAndDisplayName Get Userses via UserNicenameAndUserEmailAndDisplayName
func GetUsersesByUserNicenameAndUserEmailAndDisplayName(offset int, limit int, UserNicename_ string, UserEmail_ string, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and user_email = ? and display_name = ?", UserNicename_, UserEmail_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndUserEmailAndSpam Get Userses via UserNicenameAndUserEmailAndSpam
func GetUsersesByUserNicenameAndUserEmailAndSpam(offset int, limit int, UserNicename_ string, UserEmail_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and user_email = ? and spam = ?", UserNicename_, UserEmail_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndUserEmailAndDeleted Get Userses via UserNicenameAndUserEmailAndDeleted
func GetUsersesByUserNicenameAndUserEmailAndDeleted(offset int, limit int, UserNicename_ string, UserEmail_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and user_email = ? and deleted = ?", UserNicename_, UserEmail_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndUserUrlAndUserRegistered Get Userses via UserNicenameAndUserUrlAndUserRegistered
func GetUsersesByUserNicenameAndUserUrlAndUserRegistered(offset int, limit int, UserNicename_ string, UserUrl_ string, UserRegistered_ time.Time) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and user_url = ? and user_registered = ?", UserNicename_, UserUrl_, UserRegistered_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndUserUrlAndUserActivationKey Get Userses via UserNicenameAndUserUrlAndUserActivationKey
func GetUsersesByUserNicenameAndUserUrlAndUserActivationKey(offset int, limit int, UserNicename_ string, UserUrl_ string, UserActivationKey_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and user_url = ? and user_activation_key = ?", UserNicename_, UserUrl_, UserActivationKey_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndUserUrlAndUserStatus Get Userses via UserNicenameAndUserUrlAndUserStatus
func GetUsersesByUserNicenameAndUserUrlAndUserStatus(offset int, limit int, UserNicename_ string, UserUrl_ string, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and user_url = ? and user_status = ?", UserNicename_, UserUrl_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndUserUrlAndDisplayName Get Userses via UserNicenameAndUserUrlAndDisplayName
func GetUsersesByUserNicenameAndUserUrlAndDisplayName(offset int, limit int, UserNicename_ string, UserUrl_ string, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and user_url = ? and display_name = ?", UserNicename_, UserUrl_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndUserUrlAndSpam Get Userses via UserNicenameAndUserUrlAndSpam
func GetUsersesByUserNicenameAndUserUrlAndSpam(offset int, limit int, UserNicename_ string, UserUrl_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and user_url = ? and spam = ?", UserNicename_, UserUrl_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndUserUrlAndDeleted Get Userses via UserNicenameAndUserUrlAndDeleted
func GetUsersesByUserNicenameAndUserUrlAndDeleted(offset int, limit int, UserNicename_ string, UserUrl_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and user_url = ? and deleted = ?", UserNicename_, UserUrl_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndUserRegisteredAndUserActivationKey Get Userses via UserNicenameAndUserRegisteredAndUserActivationKey
func GetUsersesByUserNicenameAndUserRegisteredAndUserActivationKey(offset int, limit int, UserNicename_ string, UserRegistered_ time.Time, UserActivationKey_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and user_registered = ? and user_activation_key = ?", UserNicename_, UserRegistered_, UserActivationKey_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndUserRegisteredAndUserStatus Get Userses via UserNicenameAndUserRegisteredAndUserStatus
func GetUsersesByUserNicenameAndUserRegisteredAndUserStatus(offset int, limit int, UserNicename_ string, UserRegistered_ time.Time, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and user_registered = ? and user_status = ?", UserNicename_, UserRegistered_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndUserRegisteredAndDisplayName Get Userses via UserNicenameAndUserRegisteredAndDisplayName
func GetUsersesByUserNicenameAndUserRegisteredAndDisplayName(offset int, limit int, UserNicename_ string, UserRegistered_ time.Time, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and user_registered = ? and display_name = ?", UserNicename_, UserRegistered_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndUserRegisteredAndSpam Get Userses via UserNicenameAndUserRegisteredAndSpam
func GetUsersesByUserNicenameAndUserRegisteredAndSpam(offset int, limit int, UserNicename_ string, UserRegistered_ time.Time, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and user_registered = ? and spam = ?", UserNicename_, UserRegistered_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndUserRegisteredAndDeleted Get Userses via UserNicenameAndUserRegisteredAndDeleted
func GetUsersesByUserNicenameAndUserRegisteredAndDeleted(offset int, limit int, UserNicename_ string, UserRegistered_ time.Time, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and user_registered = ? and deleted = ?", UserNicename_, UserRegistered_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndUserActivationKeyAndUserStatus Get Userses via UserNicenameAndUserActivationKeyAndUserStatus
func GetUsersesByUserNicenameAndUserActivationKeyAndUserStatus(offset int, limit int, UserNicename_ string, UserActivationKey_ string, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and user_activation_key = ? and user_status = ?", UserNicename_, UserActivationKey_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndUserActivationKeyAndDisplayName Get Userses via UserNicenameAndUserActivationKeyAndDisplayName
func GetUsersesByUserNicenameAndUserActivationKeyAndDisplayName(offset int, limit int, UserNicename_ string, UserActivationKey_ string, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and user_activation_key = ? and display_name = ?", UserNicename_, UserActivationKey_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndUserActivationKeyAndSpam Get Userses via UserNicenameAndUserActivationKeyAndSpam
func GetUsersesByUserNicenameAndUserActivationKeyAndSpam(offset int, limit int, UserNicename_ string, UserActivationKey_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and user_activation_key = ? and spam = ?", UserNicename_, UserActivationKey_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndUserActivationKeyAndDeleted Get Userses via UserNicenameAndUserActivationKeyAndDeleted
func GetUsersesByUserNicenameAndUserActivationKeyAndDeleted(offset int, limit int, UserNicename_ string, UserActivationKey_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and user_activation_key = ? and deleted = ?", UserNicename_, UserActivationKey_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndUserStatusAndDisplayName Get Userses via UserNicenameAndUserStatusAndDisplayName
func GetUsersesByUserNicenameAndUserStatusAndDisplayName(offset int, limit int, UserNicename_ string, UserStatus_ int, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and user_status = ? and display_name = ?", UserNicename_, UserStatus_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndUserStatusAndSpam Get Userses via UserNicenameAndUserStatusAndSpam
func GetUsersesByUserNicenameAndUserStatusAndSpam(offset int, limit int, UserNicename_ string, UserStatus_ int, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and user_status = ? and spam = ?", UserNicename_, UserStatus_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndUserStatusAndDeleted Get Userses via UserNicenameAndUserStatusAndDeleted
func GetUsersesByUserNicenameAndUserStatusAndDeleted(offset int, limit int, UserNicename_ string, UserStatus_ int, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and user_status = ? and deleted = ?", UserNicename_, UserStatus_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndDisplayNameAndSpam Get Userses via UserNicenameAndDisplayNameAndSpam
func GetUsersesByUserNicenameAndDisplayNameAndSpam(offset int, limit int, UserNicename_ string, DisplayName_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and display_name = ? and spam = ?", UserNicename_, DisplayName_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndDisplayNameAndDeleted Get Userses via UserNicenameAndDisplayNameAndDeleted
func GetUsersesByUserNicenameAndDisplayNameAndDeleted(offset int, limit int, UserNicename_ string, DisplayName_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and display_name = ? and deleted = ?", UserNicename_, DisplayName_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndSpamAndDeleted Get Userses via UserNicenameAndSpamAndDeleted
func GetUsersesByUserNicenameAndSpamAndDeleted(offset int, limit int, UserNicename_ string, Spam_ int, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and spam = ? and deleted = ?", UserNicename_, Spam_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserEmailAndUserUrlAndUserRegistered Get Userses via UserEmailAndUserUrlAndUserRegistered
func GetUsersesByUserEmailAndUserUrlAndUserRegistered(offset int, limit int, UserEmail_ string, UserUrl_ string, UserRegistered_ time.Time) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_email = ? and user_url = ? and user_registered = ?", UserEmail_, UserUrl_, UserRegistered_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserEmailAndUserUrlAndUserActivationKey Get Userses via UserEmailAndUserUrlAndUserActivationKey
func GetUsersesByUserEmailAndUserUrlAndUserActivationKey(offset int, limit int, UserEmail_ string, UserUrl_ string, UserActivationKey_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_email = ? and user_url = ? and user_activation_key = ?", UserEmail_, UserUrl_, UserActivationKey_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserEmailAndUserUrlAndUserStatus Get Userses via UserEmailAndUserUrlAndUserStatus
func GetUsersesByUserEmailAndUserUrlAndUserStatus(offset int, limit int, UserEmail_ string, UserUrl_ string, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_email = ? and user_url = ? and user_status = ?", UserEmail_, UserUrl_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserEmailAndUserUrlAndDisplayName Get Userses via UserEmailAndUserUrlAndDisplayName
func GetUsersesByUserEmailAndUserUrlAndDisplayName(offset int, limit int, UserEmail_ string, UserUrl_ string, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_email = ? and user_url = ? and display_name = ?", UserEmail_, UserUrl_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserEmailAndUserUrlAndSpam Get Userses via UserEmailAndUserUrlAndSpam
func GetUsersesByUserEmailAndUserUrlAndSpam(offset int, limit int, UserEmail_ string, UserUrl_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_email = ? and user_url = ? and spam = ?", UserEmail_, UserUrl_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserEmailAndUserUrlAndDeleted Get Userses via UserEmailAndUserUrlAndDeleted
func GetUsersesByUserEmailAndUserUrlAndDeleted(offset int, limit int, UserEmail_ string, UserUrl_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_email = ? and user_url = ? and deleted = ?", UserEmail_, UserUrl_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserEmailAndUserRegisteredAndUserActivationKey Get Userses via UserEmailAndUserRegisteredAndUserActivationKey
func GetUsersesByUserEmailAndUserRegisteredAndUserActivationKey(offset int, limit int, UserEmail_ string, UserRegistered_ time.Time, UserActivationKey_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_email = ? and user_registered = ? and user_activation_key = ?", UserEmail_, UserRegistered_, UserActivationKey_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserEmailAndUserRegisteredAndUserStatus Get Userses via UserEmailAndUserRegisteredAndUserStatus
func GetUsersesByUserEmailAndUserRegisteredAndUserStatus(offset int, limit int, UserEmail_ string, UserRegistered_ time.Time, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_email = ? and user_registered = ? and user_status = ?", UserEmail_, UserRegistered_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserEmailAndUserRegisteredAndDisplayName Get Userses via UserEmailAndUserRegisteredAndDisplayName
func GetUsersesByUserEmailAndUserRegisteredAndDisplayName(offset int, limit int, UserEmail_ string, UserRegistered_ time.Time, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_email = ? and user_registered = ? and display_name = ?", UserEmail_, UserRegistered_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserEmailAndUserRegisteredAndSpam Get Userses via UserEmailAndUserRegisteredAndSpam
func GetUsersesByUserEmailAndUserRegisteredAndSpam(offset int, limit int, UserEmail_ string, UserRegistered_ time.Time, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_email = ? and user_registered = ? and spam = ?", UserEmail_, UserRegistered_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserEmailAndUserRegisteredAndDeleted Get Userses via UserEmailAndUserRegisteredAndDeleted
func GetUsersesByUserEmailAndUserRegisteredAndDeleted(offset int, limit int, UserEmail_ string, UserRegistered_ time.Time, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_email = ? and user_registered = ? and deleted = ?", UserEmail_, UserRegistered_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserEmailAndUserActivationKeyAndUserStatus Get Userses via UserEmailAndUserActivationKeyAndUserStatus
func GetUsersesByUserEmailAndUserActivationKeyAndUserStatus(offset int, limit int, UserEmail_ string, UserActivationKey_ string, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_email = ? and user_activation_key = ? and user_status = ?", UserEmail_, UserActivationKey_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserEmailAndUserActivationKeyAndDisplayName Get Userses via UserEmailAndUserActivationKeyAndDisplayName
func GetUsersesByUserEmailAndUserActivationKeyAndDisplayName(offset int, limit int, UserEmail_ string, UserActivationKey_ string, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_email = ? and user_activation_key = ? and display_name = ?", UserEmail_, UserActivationKey_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserEmailAndUserActivationKeyAndSpam Get Userses via UserEmailAndUserActivationKeyAndSpam
func GetUsersesByUserEmailAndUserActivationKeyAndSpam(offset int, limit int, UserEmail_ string, UserActivationKey_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_email = ? and user_activation_key = ? and spam = ?", UserEmail_, UserActivationKey_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserEmailAndUserActivationKeyAndDeleted Get Userses via UserEmailAndUserActivationKeyAndDeleted
func GetUsersesByUserEmailAndUserActivationKeyAndDeleted(offset int, limit int, UserEmail_ string, UserActivationKey_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_email = ? and user_activation_key = ? and deleted = ?", UserEmail_, UserActivationKey_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserEmailAndUserStatusAndDisplayName Get Userses via UserEmailAndUserStatusAndDisplayName
func GetUsersesByUserEmailAndUserStatusAndDisplayName(offset int, limit int, UserEmail_ string, UserStatus_ int, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_email = ? and user_status = ? and display_name = ?", UserEmail_, UserStatus_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserEmailAndUserStatusAndSpam Get Userses via UserEmailAndUserStatusAndSpam
func GetUsersesByUserEmailAndUserStatusAndSpam(offset int, limit int, UserEmail_ string, UserStatus_ int, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_email = ? and user_status = ? and spam = ?", UserEmail_, UserStatus_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserEmailAndUserStatusAndDeleted Get Userses via UserEmailAndUserStatusAndDeleted
func GetUsersesByUserEmailAndUserStatusAndDeleted(offset int, limit int, UserEmail_ string, UserStatus_ int, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_email = ? and user_status = ? and deleted = ?", UserEmail_, UserStatus_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserEmailAndDisplayNameAndSpam Get Userses via UserEmailAndDisplayNameAndSpam
func GetUsersesByUserEmailAndDisplayNameAndSpam(offset int, limit int, UserEmail_ string, DisplayName_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_email = ? and display_name = ? and spam = ?", UserEmail_, DisplayName_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserEmailAndDisplayNameAndDeleted Get Userses via UserEmailAndDisplayNameAndDeleted
func GetUsersesByUserEmailAndDisplayNameAndDeleted(offset int, limit int, UserEmail_ string, DisplayName_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_email = ? and display_name = ? and deleted = ?", UserEmail_, DisplayName_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserEmailAndSpamAndDeleted Get Userses via UserEmailAndSpamAndDeleted
func GetUsersesByUserEmailAndSpamAndDeleted(offset int, limit int, UserEmail_ string, Spam_ int, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_email = ? and spam = ? and deleted = ?", UserEmail_, Spam_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserUrlAndUserRegisteredAndUserActivationKey Get Userses via UserUrlAndUserRegisteredAndUserActivationKey
func GetUsersesByUserUrlAndUserRegisteredAndUserActivationKey(offset int, limit int, UserUrl_ string, UserRegistered_ time.Time, UserActivationKey_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_url = ? and user_registered = ? and user_activation_key = ?", UserUrl_, UserRegistered_, UserActivationKey_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserUrlAndUserRegisteredAndUserStatus Get Userses via UserUrlAndUserRegisteredAndUserStatus
func GetUsersesByUserUrlAndUserRegisteredAndUserStatus(offset int, limit int, UserUrl_ string, UserRegistered_ time.Time, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_url = ? and user_registered = ? and user_status = ?", UserUrl_, UserRegistered_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserUrlAndUserRegisteredAndDisplayName Get Userses via UserUrlAndUserRegisteredAndDisplayName
func GetUsersesByUserUrlAndUserRegisteredAndDisplayName(offset int, limit int, UserUrl_ string, UserRegistered_ time.Time, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_url = ? and user_registered = ? and display_name = ?", UserUrl_, UserRegistered_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserUrlAndUserRegisteredAndSpam Get Userses via UserUrlAndUserRegisteredAndSpam
func GetUsersesByUserUrlAndUserRegisteredAndSpam(offset int, limit int, UserUrl_ string, UserRegistered_ time.Time, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_url = ? and user_registered = ? and spam = ?", UserUrl_, UserRegistered_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserUrlAndUserRegisteredAndDeleted Get Userses via UserUrlAndUserRegisteredAndDeleted
func GetUsersesByUserUrlAndUserRegisteredAndDeleted(offset int, limit int, UserUrl_ string, UserRegistered_ time.Time, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_url = ? and user_registered = ? and deleted = ?", UserUrl_, UserRegistered_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserUrlAndUserActivationKeyAndUserStatus Get Userses via UserUrlAndUserActivationKeyAndUserStatus
func GetUsersesByUserUrlAndUserActivationKeyAndUserStatus(offset int, limit int, UserUrl_ string, UserActivationKey_ string, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_url = ? and user_activation_key = ? and user_status = ?", UserUrl_, UserActivationKey_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserUrlAndUserActivationKeyAndDisplayName Get Userses via UserUrlAndUserActivationKeyAndDisplayName
func GetUsersesByUserUrlAndUserActivationKeyAndDisplayName(offset int, limit int, UserUrl_ string, UserActivationKey_ string, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_url = ? and user_activation_key = ? and display_name = ?", UserUrl_, UserActivationKey_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserUrlAndUserActivationKeyAndSpam Get Userses via UserUrlAndUserActivationKeyAndSpam
func GetUsersesByUserUrlAndUserActivationKeyAndSpam(offset int, limit int, UserUrl_ string, UserActivationKey_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_url = ? and user_activation_key = ? and spam = ?", UserUrl_, UserActivationKey_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserUrlAndUserActivationKeyAndDeleted Get Userses via UserUrlAndUserActivationKeyAndDeleted
func GetUsersesByUserUrlAndUserActivationKeyAndDeleted(offset int, limit int, UserUrl_ string, UserActivationKey_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_url = ? and user_activation_key = ? and deleted = ?", UserUrl_, UserActivationKey_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserUrlAndUserStatusAndDisplayName Get Userses via UserUrlAndUserStatusAndDisplayName
func GetUsersesByUserUrlAndUserStatusAndDisplayName(offset int, limit int, UserUrl_ string, UserStatus_ int, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_url = ? and user_status = ? and display_name = ?", UserUrl_, UserStatus_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserUrlAndUserStatusAndSpam Get Userses via UserUrlAndUserStatusAndSpam
func GetUsersesByUserUrlAndUserStatusAndSpam(offset int, limit int, UserUrl_ string, UserStatus_ int, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_url = ? and user_status = ? and spam = ?", UserUrl_, UserStatus_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserUrlAndUserStatusAndDeleted Get Userses via UserUrlAndUserStatusAndDeleted
func GetUsersesByUserUrlAndUserStatusAndDeleted(offset int, limit int, UserUrl_ string, UserStatus_ int, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_url = ? and user_status = ? and deleted = ?", UserUrl_, UserStatus_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserUrlAndDisplayNameAndSpam Get Userses via UserUrlAndDisplayNameAndSpam
func GetUsersesByUserUrlAndDisplayNameAndSpam(offset int, limit int, UserUrl_ string, DisplayName_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_url = ? and display_name = ? and spam = ?", UserUrl_, DisplayName_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserUrlAndDisplayNameAndDeleted Get Userses via UserUrlAndDisplayNameAndDeleted
func GetUsersesByUserUrlAndDisplayNameAndDeleted(offset int, limit int, UserUrl_ string, DisplayName_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_url = ? and display_name = ? and deleted = ?", UserUrl_, DisplayName_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserUrlAndSpamAndDeleted Get Userses via UserUrlAndSpamAndDeleted
func GetUsersesByUserUrlAndSpamAndDeleted(offset int, limit int, UserUrl_ string, Spam_ int, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_url = ? and spam = ? and deleted = ?", UserUrl_, Spam_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserRegisteredAndUserActivationKeyAndUserStatus Get Userses via UserRegisteredAndUserActivationKeyAndUserStatus
func GetUsersesByUserRegisteredAndUserActivationKeyAndUserStatus(offset int, limit int, UserRegistered_ time.Time, UserActivationKey_ string, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_registered = ? and user_activation_key = ? and user_status = ?", UserRegistered_, UserActivationKey_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserRegisteredAndUserActivationKeyAndDisplayName Get Userses via UserRegisteredAndUserActivationKeyAndDisplayName
func GetUsersesByUserRegisteredAndUserActivationKeyAndDisplayName(offset int, limit int, UserRegistered_ time.Time, UserActivationKey_ string, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_registered = ? and user_activation_key = ? and display_name = ?", UserRegistered_, UserActivationKey_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserRegisteredAndUserActivationKeyAndSpam Get Userses via UserRegisteredAndUserActivationKeyAndSpam
func GetUsersesByUserRegisteredAndUserActivationKeyAndSpam(offset int, limit int, UserRegistered_ time.Time, UserActivationKey_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_registered = ? and user_activation_key = ? and spam = ?", UserRegistered_, UserActivationKey_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserRegisteredAndUserActivationKeyAndDeleted Get Userses via UserRegisteredAndUserActivationKeyAndDeleted
func GetUsersesByUserRegisteredAndUserActivationKeyAndDeleted(offset int, limit int, UserRegistered_ time.Time, UserActivationKey_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_registered = ? and user_activation_key = ? and deleted = ?", UserRegistered_, UserActivationKey_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserRegisteredAndUserStatusAndDisplayName Get Userses via UserRegisteredAndUserStatusAndDisplayName
func GetUsersesByUserRegisteredAndUserStatusAndDisplayName(offset int, limit int, UserRegistered_ time.Time, UserStatus_ int, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_registered = ? and user_status = ? and display_name = ?", UserRegistered_, UserStatus_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserRegisteredAndUserStatusAndSpam Get Userses via UserRegisteredAndUserStatusAndSpam
func GetUsersesByUserRegisteredAndUserStatusAndSpam(offset int, limit int, UserRegistered_ time.Time, UserStatus_ int, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_registered = ? and user_status = ? and spam = ?", UserRegistered_, UserStatus_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserRegisteredAndUserStatusAndDeleted Get Userses via UserRegisteredAndUserStatusAndDeleted
func GetUsersesByUserRegisteredAndUserStatusAndDeleted(offset int, limit int, UserRegistered_ time.Time, UserStatus_ int, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_registered = ? and user_status = ? and deleted = ?", UserRegistered_, UserStatus_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserRegisteredAndDisplayNameAndSpam Get Userses via UserRegisteredAndDisplayNameAndSpam
func GetUsersesByUserRegisteredAndDisplayNameAndSpam(offset int, limit int, UserRegistered_ time.Time, DisplayName_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_registered = ? and display_name = ? and spam = ?", UserRegistered_, DisplayName_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserRegisteredAndDisplayNameAndDeleted Get Userses via UserRegisteredAndDisplayNameAndDeleted
func GetUsersesByUserRegisteredAndDisplayNameAndDeleted(offset int, limit int, UserRegistered_ time.Time, DisplayName_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_registered = ? and display_name = ? and deleted = ?", UserRegistered_, DisplayName_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserRegisteredAndSpamAndDeleted Get Userses via UserRegisteredAndSpamAndDeleted
func GetUsersesByUserRegisteredAndSpamAndDeleted(offset int, limit int, UserRegistered_ time.Time, Spam_ int, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_registered = ? and spam = ? and deleted = ?", UserRegistered_, Spam_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserActivationKeyAndUserStatusAndDisplayName Get Userses via UserActivationKeyAndUserStatusAndDisplayName
func GetUsersesByUserActivationKeyAndUserStatusAndDisplayName(offset int, limit int, UserActivationKey_ string, UserStatus_ int, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_activation_key = ? and user_status = ? and display_name = ?", UserActivationKey_, UserStatus_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserActivationKeyAndUserStatusAndSpam Get Userses via UserActivationKeyAndUserStatusAndSpam
func GetUsersesByUserActivationKeyAndUserStatusAndSpam(offset int, limit int, UserActivationKey_ string, UserStatus_ int, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_activation_key = ? and user_status = ? and spam = ?", UserActivationKey_, UserStatus_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserActivationKeyAndUserStatusAndDeleted Get Userses via UserActivationKeyAndUserStatusAndDeleted
func GetUsersesByUserActivationKeyAndUserStatusAndDeleted(offset int, limit int, UserActivationKey_ string, UserStatus_ int, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_activation_key = ? and user_status = ? and deleted = ?", UserActivationKey_, UserStatus_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserActivationKeyAndDisplayNameAndSpam Get Userses via UserActivationKeyAndDisplayNameAndSpam
func GetUsersesByUserActivationKeyAndDisplayNameAndSpam(offset int, limit int, UserActivationKey_ string, DisplayName_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_activation_key = ? and display_name = ? and spam = ?", UserActivationKey_, DisplayName_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserActivationKeyAndDisplayNameAndDeleted Get Userses via UserActivationKeyAndDisplayNameAndDeleted
func GetUsersesByUserActivationKeyAndDisplayNameAndDeleted(offset int, limit int, UserActivationKey_ string, DisplayName_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_activation_key = ? and display_name = ? and deleted = ?", UserActivationKey_, DisplayName_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserActivationKeyAndSpamAndDeleted Get Userses via UserActivationKeyAndSpamAndDeleted
func GetUsersesByUserActivationKeyAndSpamAndDeleted(offset int, limit int, UserActivationKey_ string, Spam_ int, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_activation_key = ? and spam = ? and deleted = ?", UserActivationKey_, Spam_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserStatusAndDisplayNameAndSpam Get Userses via UserStatusAndDisplayNameAndSpam
func GetUsersesByUserStatusAndDisplayNameAndSpam(offset int, limit int, UserStatus_ int, DisplayName_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_status = ? and display_name = ? and spam = ?", UserStatus_, DisplayName_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserStatusAndDisplayNameAndDeleted Get Userses via UserStatusAndDisplayNameAndDeleted
func GetUsersesByUserStatusAndDisplayNameAndDeleted(offset int, limit int, UserStatus_ int, DisplayName_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_status = ? and display_name = ? and deleted = ?", UserStatus_, DisplayName_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserStatusAndSpamAndDeleted Get Userses via UserStatusAndSpamAndDeleted
func GetUsersesByUserStatusAndSpamAndDeleted(offset int, limit int, UserStatus_ int, Spam_ int, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_status = ? and spam = ? and deleted = ?", UserStatus_, Spam_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByDisplayNameAndSpamAndDeleted Get Userses via DisplayNameAndSpamAndDeleted
func GetUsersesByDisplayNameAndSpamAndDeleted(offset int, limit int, DisplayName_ string, Spam_ int, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("display_name = ? and spam = ? and deleted = ?", DisplayName_, Spam_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserLogin Get Userses via IdAndUserLogin
func GetUsersesByIdAndUserLogin(offset int, limit int, Id_ int64, UserLogin_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_login = ?", Id_, UserLogin_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserPass Get Userses via IdAndUserPass
func GetUsersesByIdAndUserPass(offset int, limit int, Id_ int64, UserPass_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_pass = ?", Id_, UserPass_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserNicename Get Userses via IdAndUserNicename
func GetUsersesByIdAndUserNicename(offset int, limit int, Id_ int64, UserNicename_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_nicename = ?", Id_, UserNicename_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserEmail Get Userses via IdAndUserEmail
func GetUsersesByIdAndUserEmail(offset int, limit int, Id_ int64, UserEmail_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_email = ?", Id_, UserEmail_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserUrl Get Userses via IdAndUserUrl
func GetUsersesByIdAndUserUrl(offset int, limit int, Id_ int64, UserUrl_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_url = ?", Id_, UserUrl_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserRegistered Get Userses via IdAndUserRegistered
func GetUsersesByIdAndUserRegistered(offset int, limit int, Id_ int64, UserRegistered_ time.Time) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_registered = ?", Id_, UserRegistered_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserActivationKey Get Userses via IdAndUserActivationKey
func GetUsersesByIdAndUserActivationKey(offset int, limit int, Id_ int64, UserActivationKey_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_activation_key = ?", Id_, UserActivationKey_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndUserStatus Get Userses via IdAndUserStatus
func GetUsersesByIdAndUserStatus(offset int, limit int, Id_ int64, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and user_status = ?", Id_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndDisplayName Get Userses via IdAndDisplayName
func GetUsersesByIdAndDisplayName(offset int, limit int, Id_ int64, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and display_name = ?", Id_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndSpam Get Userses via IdAndSpam
func GetUsersesByIdAndSpam(offset int, limit int, Id_ int64, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and spam = ?", Id_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByIdAndDeleted Get Userses via IdAndDeleted
func GetUsersesByIdAndDeleted(offset int, limit int, Id_ int64, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("id = ? and deleted = ?", Id_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserPass Get Userses via UserLoginAndUserPass
func GetUsersesByUserLoginAndUserPass(offset int, limit int, UserLogin_ string, UserPass_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_pass = ?", UserLogin_, UserPass_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserNicename Get Userses via UserLoginAndUserNicename
func GetUsersesByUserLoginAndUserNicename(offset int, limit int, UserLogin_ string, UserNicename_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_nicename = ?", UserLogin_, UserNicename_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserEmail Get Userses via UserLoginAndUserEmail
func GetUsersesByUserLoginAndUserEmail(offset int, limit int, UserLogin_ string, UserEmail_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_email = ?", UserLogin_, UserEmail_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserUrl Get Userses via UserLoginAndUserUrl
func GetUsersesByUserLoginAndUserUrl(offset int, limit int, UserLogin_ string, UserUrl_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_url = ?", UserLogin_, UserUrl_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserRegistered Get Userses via UserLoginAndUserRegistered
func GetUsersesByUserLoginAndUserRegistered(offset int, limit int, UserLogin_ string, UserRegistered_ time.Time) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_registered = ?", UserLogin_, UserRegistered_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserActivationKey Get Userses via UserLoginAndUserActivationKey
func GetUsersesByUserLoginAndUserActivationKey(offset int, limit int, UserLogin_ string, UserActivationKey_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_activation_key = ?", UserLogin_, UserActivationKey_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndUserStatus Get Userses via UserLoginAndUserStatus
func GetUsersesByUserLoginAndUserStatus(offset int, limit int, UserLogin_ string, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and user_status = ?", UserLogin_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndDisplayName Get Userses via UserLoginAndDisplayName
func GetUsersesByUserLoginAndDisplayName(offset int, limit int, UserLogin_ string, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and display_name = ?", UserLogin_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndSpam Get Userses via UserLoginAndSpam
func GetUsersesByUserLoginAndSpam(offset int, limit int, UserLogin_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and spam = ?", UserLogin_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserLoginAndDeleted Get Userses via UserLoginAndDeleted
func GetUsersesByUserLoginAndDeleted(offset int, limit int, UserLogin_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ? and deleted = ?", UserLogin_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserNicename Get Userses via UserPassAndUserNicename
func GetUsersesByUserPassAndUserNicename(offset int, limit int, UserPass_ string, UserNicename_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_nicename = ?", UserPass_, UserNicename_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserEmail Get Userses via UserPassAndUserEmail
func GetUsersesByUserPassAndUserEmail(offset int, limit int, UserPass_ string, UserEmail_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_email = ?", UserPass_, UserEmail_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserUrl Get Userses via UserPassAndUserUrl
func GetUsersesByUserPassAndUserUrl(offset int, limit int, UserPass_ string, UserUrl_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_url = ?", UserPass_, UserUrl_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserRegistered Get Userses via UserPassAndUserRegistered
func GetUsersesByUserPassAndUserRegistered(offset int, limit int, UserPass_ string, UserRegistered_ time.Time) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_registered = ?", UserPass_, UserRegistered_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserActivationKey Get Userses via UserPassAndUserActivationKey
func GetUsersesByUserPassAndUserActivationKey(offset int, limit int, UserPass_ string, UserActivationKey_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_activation_key = ?", UserPass_, UserActivationKey_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndUserStatus Get Userses via UserPassAndUserStatus
func GetUsersesByUserPassAndUserStatus(offset int, limit int, UserPass_ string, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and user_status = ?", UserPass_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndDisplayName Get Userses via UserPassAndDisplayName
func GetUsersesByUserPassAndDisplayName(offset int, limit int, UserPass_ string, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and display_name = ?", UserPass_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndSpam Get Userses via UserPassAndSpam
func GetUsersesByUserPassAndSpam(offset int, limit int, UserPass_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and spam = ?", UserPass_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserPassAndDeleted Get Userses via UserPassAndDeleted
func GetUsersesByUserPassAndDeleted(offset int, limit int, UserPass_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ? and deleted = ?", UserPass_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndUserEmail Get Userses via UserNicenameAndUserEmail
func GetUsersesByUserNicenameAndUserEmail(offset int, limit int, UserNicename_ string, UserEmail_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and user_email = ?", UserNicename_, UserEmail_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndUserUrl Get Userses via UserNicenameAndUserUrl
func GetUsersesByUserNicenameAndUserUrl(offset int, limit int, UserNicename_ string, UserUrl_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and user_url = ?", UserNicename_, UserUrl_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndUserRegistered Get Userses via UserNicenameAndUserRegistered
func GetUsersesByUserNicenameAndUserRegistered(offset int, limit int, UserNicename_ string, UserRegistered_ time.Time) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and user_registered = ?", UserNicename_, UserRegistered_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndUserActivationKey Get Userses via UserNicenameAndUserActivationKey
func GetUsersesByUserNicenameAndUserActivationKey(offset int, limit int, UserNicename_ string, UserActivationKey_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and user_activation_key = ?", UserNicename_, UserActivationKey_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndUserStatus Get Userses via UserNicenameAndUserStatus
func GetUsersesByUserNicenameAndUserStatus(offset int, limit int, UserNicename_ string, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and user_status = ?", UserNicename_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndDisplayName Get Userses via UserNicenameAndDisplayName
func GetUsersesByUserNicenameAndDisplayName(offset int, limit int, UserNicename_ string, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and display_name = ?", UserNicename_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndSpam Get Userses via UserNicenameAndSpam
func GetUsersesByUserNicenameAndSpam(offset int, limit int, UserNicename_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and spam = ?", UserNicename_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserNicenameAndDeleted Get Userses via UserNicenameAndDeleted
func GetUsersesByUserNicenameAndDeleted(offset int, limit int, UserNicename_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ? and deleted = ?", UserNicename_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserEmailAndUserUrl Get Userses via UserEmailAndUserUrl
func GetUsersesByUserEmailAndUserUrl(offset int, limit int, UserEmail_ string, UserUrl_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_email = ? and user_url = ?", UserEmail_, UserUrl_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserEmailAndUserRegistered Get Userses via UserEmailAndUserRegistered
func GetUsersesByUserEmailAndUserRegistered(offset int, limit int, UserEmail_ string, UserRegistered_ time.Time) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_email = ? and user_registered = ?", UserEmail_, UserRegistered_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserEmailAndUserActivationKey Get Userses via UserEmailAndUserActivationKey
func GetUsersesByUserEmailAndUserActivationKey(offset int, limit int, UserEmail_ string, UserActivationKey_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_email = ? and user_activation_key = ?", UserEmail_, UserActivationKey_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserEmailAndUserStatus Get Userses via UserEmailAndUserStatus
func GetUsersesByUserEmailAndUserStatus(offset int, limit int, UserEmail_ string, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_email = ? and user_status = ?", UserEmail_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserEmailAndDisplayName Get Userses via UserEmailAndDisplayName
func GetUsersesByUserEmailAndDisplayName(offset int, limit int, UserEmail_ string, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_email = ? and display_name = ?", UserEmail_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserEmailAndSpam Get Userses via UserEmailAndSpam
func GetUsersesByUserEmailAndSpam(offset int, limit int, UserEmail_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_email = ? and spam = ?", UserEmail_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserEmailAndDeleted Get Userses via UserEmailAndDeleted
func GetUsersesByUserEmailAndDeleted(offset int, limit int, UserEmail_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_email = ? and deleted = ?", UserEmail_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserUrlAndUserRegistered Get Userses via UserUrlAndUserRegistered
func GetUsersesByUserUrlAndUserRegistered(offset int, limit int, UserUrl_ string, UserRegistered_ time.Time) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_url = ? and user_registered = ?", UserUrl_, UserRegistered_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserUrlAndUserActivationKey Get Userses via UserUrlAndUserActivationKey
func GetUsersesByUserUrlAndUserActivationKey(offset int, limit int, UserUrl_ string, UserActivationKey_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_url = ? and user_activation_key = ?", UserUrl_, UserActivationKey_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserUrlAndUserStatus Get Userses via UserUrlAndUserStatus
func GetUsersesByUserUrlAndUserStatus(offset int, limit int, UserUrl_ string, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_url = ? and user_status = ?", UserUrl_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserUrlAndDisplayName Get Userses via UserUrlAndDisplayName
func GetUsersesByUserUrlAndDisplayName(offset int, limit int, UserUrl_ string, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_url = ? and display_name = ?", UserUrl_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserUrlAndSpam Get Userses via UserUrlAndSpam
func GetUsersesByUserUrlAndSpam(offset int, limit int, UserUrl_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_url = ? and spam = ?", UserUrl_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserUrlAndDeleted Get Userses via UserUrlAndDeleted
func GetUsersesByUserUrlAndDeleted(offset int, limit int, UserUrl_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_url = ? and deleted = ?", UserUrl_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserRegisteredAndUserActivationKey Get Userses via UserRegisteredAndUserActivationKey
func GetUsersesByUserRegisteredAndUserActivationKey(offset int, limit int, UserRegistered_ time.Time, UserActivationKey_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_registered = ? and user_activation_key = ?", UserRegistered_, UserActivationKey_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserRegisteredAndUserStatus Get Userses via UserRegisteredAndUserStatus
func GetUsersesByUserRegisteredAndUserStatus(offset int, limit int, UserRegistered_ time.Time, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_registered = ? and user_status = ?", UserRegistered_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserRegisteredAndDisplayName Get Userses via UserRegisteredAndDisplayName
func GetUsersesByUserRegisteredAndDisplayName(offset int, limit int, UserRegistered_ time.Time, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_registered = ? and display_name = ?", UserRegistered_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserRegisteredAndSpam Get Userses via UserRegisteredAndSpam
func GetUsersesByUserRegisteredAndSpam(offset int, limit int, UserRegistered_ time.Time, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_registered = ? and spam = ?", UserRegistered_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserRegisteredAndDeleted Get Userses via UserRegisteredAndDeleted
func GetUsersesByUserRegisteredAndDeleted(offset int, limit int, UserRegistered_ time.Time, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_registered = ? and deleted = ?", UserRegistered_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserActivationKeyAndUserStatus Get Userses via UserActivationKeyAndUserStatus
func GetUsersesByUserActivationKeyAndUserStatus(offset int, limit int, UserActivationKey_ string, UserStatus_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_activation_key = ? and user_status = ?", UserActivationKey_, UserStatus_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserActivationKeyAndDisplayName Get Userses via UserActivationKeyAndDisplayName
func GetUsersesByUserActivationKeyAndDisplayName(offset int, limit int, UserActivationKey_ string, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_activation_key = ? and display_name = ?", UserActivationKey_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserActivationKeyAndSpam Get Userses via UserActivationKeyAndSpam
func GetUsersesByUserActivationKeyAndSpam(offset int, limit int, UserActivationKey_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_activation_key = ? and spam = ?", UserActivationKey_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserActivationKeyAndDeleted Get Userses via UserActivationKeyAndDeleted
func GetUsersesByUserActivationKeyAndDeleted(offset int, limit int, UserActivationKey_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_activation_key = ? and deleted = ?", UserActivationKey_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserStatusAndDisplayName Get Userses via UserStatusAndDisplayName
func GetUsersesByUserStatusAndDisplayName(offset int, limit int, UserStatus_ int, DisplayName_ string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_status = ? and display_name = ?", UserStatus_, DisplayName_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserStatusAndSpam Get Userses via UserStatusAndSpam
func GetUsersesByUserStatusAndSpam(offset int, limit int, UserStatus_ int, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_status = ? and spam = ?", UserStatus_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByUserStatusAndDeleted Get Userses via UserStatusAndDeleted
func GetUsersesByUserStatusAndDeleted(offset int, limit int, UserStatus_ int, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_status = ? and deleted = ?", UserStatus_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByDisplayNameAndSpam Get Userses via DisplayNameAndSpam
func GetUsersesByDisplayNameAndSpam(offset int, limit int, DisplayName_ string, Spam_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("display_name = ? and spam = ?", DisplayName_, Spam_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesByDisplayNameAndDeleted Get Userses via DisplayNameAndDeleted
func GetUsersesByDisplayNameAndDeleted(offset int, limit int, DisplayName_ string, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("display_name = ? and deleted = ?", DisplayName_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUsersesBySpamAndDeleted Get Userses via SpamAndDeleted
func GetUsersesBySpamAndDeleted(offset int, limit int, Spam_ int, Deleted_ int) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("spam = ? and deleted = ?", Spam_, Deleted_).Limit(limit, offset).Find(_Users)
	return _Users, err
}

// GetUserses Get Userses via field
func GetUserses(offset int, limit int, field string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Limit(limit, offset).Desc(field).Find(_Users)
	return _Users, err
}

// GetUsersesViaId Get Userss via Id
func GetUsersesViaId(offset int, limit int, Id_ int64, field string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("ID = ?", Id_).Limit(limit, offset).Desc(field).Find(_Users)
	return _Users, err
}

// GetUsersesViaUserLogin Get Userss via UserLogin
func GetUsersesViaUserLogin(offset int, limit int, UserLogin_ string, field string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_login = ?", UserLogin_).Limit(limit, offset).Desc(field).Find(_Users)
	return _Users, err
}

// GetUsersesViaUserPass Get Userss via UserPass
func GetUsersesViaUserPass(offset int, limit int, UserPass_ string, field string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_pass = ?", UserPass_).Limit(limit, offset).Desc(field).Find(_Users)
	return _Users, err
}

// GetUsersesViaUserNicename Get Userss via UserNicename
func GetUsersesViaUserNicename(offset int, limit int, UserNicename_ string, field string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_nicename = ?", UserNicename_).Limit(limit, offset).Desc(field).Find(_Users)
	return _Users, err
}

// GetUsersesViaUserEmail Get Userss via UserEmail
func GetUsersesViaUserEmail(offset int, limit int, UserEmail_ string, field string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_email = ?", UserEmail_).Limit(limit, offset).Desc(field).Find(_Users)
	return _Users, err
}

// GetUsersesViaUserUrl Get Userss via UserUrl
func GetUsersesViaUserUrl(offset int, limit int, UserUrl_ string, field string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_url = ?", UserUrl_).Limit(limit, offset).Desc(field).Find(_Users)
	return _Users, err
}

// GetUsersesViaUserRegistered Get Userss via UserRegistered
func GetUsersesViaUserRegistered(offset int, limit int, UserRegistered_ time.Time, field string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_registered = ?", UserRegistered_).Limit(limit, offset).Desc(field).Find(_Users)
	return _Users, err
}

// GetUsersesViaUserActivationKey Get Userss via UserActivationKey
func GetUsersesViaUserActivationKey(offset int, limit int, UserActivationKey_ string, field string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_activation_key = ?", UserActivationKey_).Limit(limit, offset).Desc(field).Find(_Users)
	return _Users, err
}

// GetUsersesViaUserStatus Get Userss via UserStatus
func GetUsersesViaUserStatus(offset int, limit int, UserStatus_ int, field string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("user_status = ?", UserStatus_).Limit(limit, offset).Desc(field).Find(_Users)
	return _Users, err
}

// GetUsersesViaDisplayName Get Userss via DisplayName
func GetUsersesViaDisplayName(offset int, limit int, DisplayName_ string, field string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("display_name = ?", DisplayName_).Limit(limit, offset).Desc(field).Find(_Users)
	return _Users, err
}

// GetUsersesViaSpam Get Userss via Spam
func GetUsersesViaSpam(offset int, limit int, Spam_ int, field string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("spam = ?", Spam_).Limit(limit, offset).Desc(field).Find(_Users)
	return _Users, err
}

// GetUsersesViaDeleted Get Userss via Deleted
func GetUsersesViaDeleted(offset int, limit int, Deleted_ int, field string) (*[]*Users, error) {
	var _Users = new([]*Users)
	err := Engine.Table("users").Where("deleted = ?", Deleted_).Limit(limit, offset).Desc(field).Find(_Users)
	return _Users, err
}

// HasUsersViaId Has Users via Id
func HasUsersViaId(iId int64) bool {
	if has, err := Engine.Where("ID = ?", iId).Get(new(Users)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsersViaUserLogin Has Users via UserLogin
func HasUsersViaUserLogin(iUserLogin string) bool {
	if has, err := Engine.Where("user_login = ?", iUserLogin).Get(new(Users)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsersViaUserPass Has Users via UserPass
func HasUsersViaUserPass(iUserPass string) bool {
	if has, err := Engine.Where("user_pass = ?", iUserPass).Get(new(Users)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsersViaUserNicename Has Users via UserNicename
func HasUsersViaUserNicename(iUserNicename string) bool {
	if has, err := Engine.Where("user_nicename = ?", iUserNicename).Get(new(Users)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsersViaUserEmail Has Users via UserEmail
func HasUsersViaUserEmail(iUserEmail string) bool {
	if has, err := Engine.Where("user_email = ?", iUserEmail).Get(new(Users)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsersViaUserUrl Has Users via UserUrl
func HasUsersViaUserUrl(iUserUrl string) bool {
	if has, err := Engine.Where("user_url = ?", iUserUrl).Get(new(Users)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsersViaUserRegistered Has Users via UserRegistered
func HasUsersViaUserRegistered(iUserRegistered time.Time) bool {
	if has, err := Engine.Where("user_registered = ?", iUserRegistered).Get(new(Users)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsersViaUserActivationKey Has Users via UserActivationKey
func HasUsersViaUserActivationKey(iUserActivationKey string) bool {
	if has, err := Engine.Where("user_activation_key = ?", iUserActivationKey).Get(new(Users)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsersViaUserStatus Has Users via UserStatus
func HasUsersViaUserStatus(iUserStatus int) bool {
	if has, err := Engine.Where("user_status = ?", iUserStatus).Get(new(Users)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsersViaDisplayName Has Users via DisplayName
func HasUsersViaDisplayName(iDisplayName string) bool {
	if has, err := Engine.Where("display_name = ?", iDisplayName).Get(new(Users)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsersViaSpam Has Users via Spam
func HasUsersViaSpam(iSpam int) bool {
	if has, err := Engine.Where("spam = ?", iSpam).Get(new(Users)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasUsersViaDeleted Has Users via Deleted
func HasUsersViaDeleted(iDeleted int) bool {
	if has, err := Engine.Where("deleted = ?", iDeleted).Get(new(Users)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetUsersViaId Get Users via Id
func GetUsersViaId(iId int64) (*Users, error) {
	var _Users = &Users{Id: iId}
	has, err := Engine.Get(_Users)
	if has {
		return _Users, err
	} else {
		return nil, err
	}
}

// GetUsersViaUserLogin Get Users via UserLogin
func GetUsersViaUserLogin(iUserLogin string) (*Users, error) {
	var _Users = &Users{UserLogin: iUserLogin}
	has, err := Engine.Get(_Users)
	if has {
		return _Users, err
	} else {
		return nil, err
	}
}

// GetUsersViaUserPass Get Users via UserPass
func GetUsersViaUserPass(iUserPass string) (*Users, error) {
	var _Users = &Users{UserPass: iUserPass}
	has, err := Engine.Get(_Users)
	if has {
		return _Users, err
	} else {
		return nil, err
	}
}

// GetUsersViaUserNicename Get Users via UserNicename
func GetUsersViaUserNicename(iUserNicename string) (*Users, error) {
	var _Users = &Users{UserNicename: iUserNicename}
	has, err := Engine.Get(_Users)
	if has {
		return _Users, err
	} else {
		return nil, err
	}
}

// GetUsersViaUserEmail Get Users via UserEmail
func GetUsersViaUserEmail(iUserEmail string) (*Users, error) {
	var _Users = &Users{UserEmail: iUserEmail}
	has, err := Engine.Get(_Users)
	if has {
		return _Users, err
	} else {
		return nil, err
	}
}

// GetUsersViaUserUrl Get Users via UserUrl
func GetUsersViaUserUrl(iUserUrl string) (*Users, error) {
	var _Users = &Users{UserUrl: iUserUrl}
	has, err := Engine.Get(_Users)
	if has {
		return _Users, err
	} else {
		return nil, err
	}
}

// GetUsersViaUserRegistered Get Users via UserRegistered
func GetUsersViaUserRegistered(iUserRegistered time.Time) (*Users, error) {
	var _Users = &Users{UserRegistered: iUserRegistered}
	has, err := Engine.Get(_Users)
	if has {
		return _Users, err
	} else {
		return nil, err
	}
}

// GetUsersViaUserActivationKey Get Users via UserActivationKey
func GetUsersViaUserActivationKey(iUserActivationKey string) (*Users, error) {
	var _Users = &Users{UserActivationKey: iUserActivationKey}
	has, err := Engine.Get(_Users)
	if has {
		return _Users, err
	} else {
		return nil, err
	}
}

// GetUsersViaUserStatus Get Users via UserStatus
func GetUsersViaUserStatus(iUserStatus int) (*Users, error) {
	var _Users = &Users{UserStatus: iUserStatus}
	has, err := Engine.Get(_Users)
	if has {
		return _Users, err
	} else {
		return nil, err
	}
}

// GetUsersViaDisplayName Get Users via DisplayName
func GetUsersViaDisplayName(iDisplayName string) (*Users, error) {
	var _Users = &Users{DisplayName: iDisplayName}
	has, err := Engine.Get(_Users)
	if has {
		return _Users, err
	} else {
		return nil, err
	}
}

// GetUsersViaSpam Get Users via Spam
func GetUsersViaSpam(iSpam int) (*Users, error) {
	var _Users = &Users{Spam: iSpam}
	has, err := Engine.Get(_Users)
	if has {
		return _Users, err
	} else {
		return nil, err
	}
}

// GetUsersViaDeleted Get Users via Deleted
func GetUsersViaDeleted(iDeleted int) (*Users, error) {
	var _Users = &Users{Deleted: iDeleted}
	has, err := Engine.Get(_Users)
	if has {
		return _Users, err
	} else {
		return nil, err
	}
}

// SetUsersViaId Set Users via Id
func SetUsersViaId(iId int64, users *Users) (int64, error) {
	users.Id = iId
	return Engine.Insert(users)
}

// SetUsersViaUserLogin Set Users via UserLogin
func SetUsersViaUserLogin(iUserLogin string, users *Users) (int64, error) {
	users.UserLogin = iUserLogin
	return Engine.Insert(users)
}

// SetUsersViaUserPass Set Users via UserPass
func SetUsersViaUserPass(iUserPass string, users *Users) (int64, error) {
	users.UserPass = iUserPass
	return Engine.Insert(users)
}

// SetUsersViaUserNicename Set Users via UserNicename
func SetUsersViaUserNicename(iUserNicename string, users *Users) (int64, error) {
	users.UserNicename = iUserNicename
	return Engine.Insert(users)
}

// SetUsersViaUserEmail Set Users via UserEmail
func SetUsersViaUserEmail(iUserEmail string, users *Users) (int64, error) {
	users.UserEmail = iUserEmail
	return Engine.Insert(users)
}

// SetUsersViaUserUrl Set Users via UserUrl
func SetUsersViaUserUrl(iUserUrl string, users *Users) (int64, error) {
	users.UserUrl = iUserUrl
	return Engine.Insert(users)
}

// SetUsersViaUserRegistered Set Users via UserRegistered
func SetUsersViaUserRegistered(iUserRegistered time.Time, users *Users) (int64, error) {
	users.UserRegistered = iUserRegistered
	return Engine.Insert(users)
}

// SetUsersViaUserActivationKey Set Users via UserActivationKey
func SetUsersViaUserActivationKey(iUserActivationKey string, users *Users) (int64, error) {
	users.UserActivationKey = iUserActivationKey
	return Engine.Insert(users)
}

// SetUsersViaUserStatus Set Users via UserStatus
func SetUsersViaUserStatus(iUserStatus int, users *Users) (int64, error) {
	users.UserStatus = iUserStatus
	return Engine.Insert(users)
}

// SetUsersViaDisplayName Set Users via DisplayName
func SetUsersViaDisplayName(iDisplayName string, users *Users) (int64, error) {
	users.DisplayName = iDisplayName
	return Engine.Insert(users)
}

// SetUsersViaSpam Set Users via Spam
func SetUsersViaSpam(iSpam int, users *Users) (int64, error) {
	users.Spam = iSpam
	return Engine.Insert(users)
}

// SetUsersViaDeleted Set Users via Deleted
func SetUsersViaDeleted(iDeleted int, users *Users) (int64, error) {
	users.Deleted = iDeleted
	return Engine.Insert(users)
}

// AddUsers Add Users via all columns
func AddUsers(iId int64, iUserLogin string, iUserPass string, iUserNicename string, iUserEmail string, iUserUrl string, iUserRegistered time.Time, iUserActivationKey string, iUserStatus int, iDisplayName string, iSpam int, iDeleted int) error {
	_Users := &Users{Id: iId, UserLogin: iUserLogin, UserPass: iUserPass, UserNicename: iUserNicename, UserEmail: iUserEmail, UserUrl: iUserUrl, UserRegistered: iUserRegistered, UserActivationKey: iUserActivationKey, UserStatus: iUserStatus, DisplayName: iDisplayName, Spam: iSpam, Deleted: iDeleted}
	if _, err := Engine.Insert(_Users); err != nil {
		return err
	} else {
		return nil
	}
}

// PostUsers Post Users via iUsers
func PostUsers(iUsers *Users) (int64, error) {
	_, err := Engine.Insert(iUsers)
	return iUsers.Id, err
}

// PutUsers Put Users
func PutUsers(iUsers *Users) (int64, error) {
	_, err := Engine.Id(iUsers.Id).Update(iUsers)
	return iUsers.Id, err
}

// PutUsersViaId Put Users via Id
func PutUsersViaId(Id_ int64, iUsers *Users) (int64, error) {
	row, err := Engine.Update(iUsers, &Users{Id: Id_})
	return row, err
}

// PutUsersViaUserLogin Put Users via UserLogin
func PutUsersViaUserLogin(UserLogin_ string, iUsers *Users) (int64, error) {
	row, err := Engine.Update(iUsers, &Users{UserLogin: UserLogin_})
	return row, err
}

// PutUsersViaUserPass Put Users via UserPass
func PutUsersViaUserPass(UserPass_ string, iUsers *Users) (int64, error) {
	row, err := Engine.Update(iUsers, &Users{UserPass: UserPass_})
	return row, err
}

// PutUsersViaUserNicename Put Users via UserNicename
func PutUsersViaUserNicename(UserNicename_ string, iUsers *Users) (int64, error) {
	row, err := Engine.Update(iUsers, &Users{UserNicename: UserNicename_})
	return row, err
}

// PutUsersViaUserEmail Put Users via UserEmail
func PutUsersViaUserEmail(UserEmail_ string, iUsers *Users) (int64, error) {
	row, err := Engine.Update(iUsers, &Users{UserEmail: UserEmail_})
	return row, err
}

// PutUsersViaUserUrl Put Users via UserUrl
func PutUsersViaUserUrl(UserUrl_ string, iUsers *Users) (int64, error) {
	row, err := Engine.Update(iUsers, &Users{UserUrl: UserUrl_})
	return row, err
}

// PutUsersViaUserRegistered Put Users via UserRegistered
func PutUsersViaUserRegistered(UserRegistered_ time.Time, iUsers *Users) (int64, error) {
	row, err := Engine.Update(iUsers, &Users{UserRegistered: UserRegistered_})
	return row, err
}

// PutUsersViaUserActivationKey Put Users via UserActivationKey
func PutUsersViaUserActivationKey(UserActivationKey_ string, iUsers *Users) (int64, error) {
	row, err := Engine.Update(iUsers, &Users{UserActivationKey: UserActivationKey_})
	return row, err
}

// PutUsersViaUserStatus Put Users via UserStatus
func PutUsersViaUserStatus(UserStatus_ int, iUsers *Users) (int64, error) {
	row, err := Engine.Update(iUsers, &Users{UserStatus: UserStatus_})
	return row, err
}

// PutUsersViaDisplayName Put Users via DisplayName
func PutUsersViaDisplayName(DisplayName_ string, iUsers *Users) (int64, error) {
	row, err := Engine.Update(iUsers, &Users{DisplayName: DisplayName_})
	return row, err
}

// PutUsersViaSpam Put Users via Spam
func PutUsersViaSpam(Spam_ int, iUsers *Users) (int64, error) {
	row, err := Engine.Update(iUsers, &Users{Spam: Spam_})
	return row, err
}

// PutUsersViaDeleted Put Users via Deleted
func PutUsersViaDeleted(Deleted_ int, iUsers *Users) (int64, error) {
	row, err := Engine.Update(iUsers, &Users{Deleted: Deleted_})
	return row, err
}

// UpdateUsersViaId via map[string]interface{}{}
func UpdateUsersViaId(iId int64, iUsersMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Users)).Where("ID = ?", iId).Update(iUsersMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsersViaUserLogin via map[string]interface{}{}
func UpdateUsersViaUserLogin(iUserLogin string, iUsersMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Users)).Where("user_login = ?", iUserLogin).Update(iUsersMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsersViaUserPass via map[string]interface{}{}
func UpdateUsersViaUserPass(iUserPass string, iUsersMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Users)).Where("user_pass = ?", iUserPass).Update(iUsersMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsersViaUserNicename via map[string]interface{}{}
func UpdateUsersViaUserNicename(iUserNicename string, iUsersMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Users)).Where("user_nicename = ?", iUserNicename).Update(iUsersMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsersViaUserEmail via map[string]interface{}{}
func UpdateUsersViaUserEmail(iUserEmail string, iUsersMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Users)).Where("user_email = ?", iUserEmail).Update(iUsersMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsersViaUserUrl via map[string]interface{}{}
func UpdateUsersViaUserUrl(iUserUrl string, iUsersMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Users)).Where("user_url = ?", iUserUrl).Update(iUsersMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsersViaUserRegistered via map[string]interface{}{}
func UpdateUsersViaUserRegistered(iUserRegistered time.Time, iUsersMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Users)).Where("user_registered = ?", iUserRegistered).Update(iUsersMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsersViaUserActivationKey via map[string]interface{}{}
func UpdateUsersViaUserActivationKey(iUserActivationKey string, iUsersMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Users)).Where("user_activation_key = ?", iUserActivationKey).Update(iUsersMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsersViaUserStatus via map[string]interface{}{}
func UpdateUsersViaUserStatus(iUserStatus int, iUsersMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Users)).Where("user_status = ?", iUserStatus).Update(iUsersMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsersViaDisplayName via map[string]interface{}{}
func UpdateUsersViaDisplayName(iDisplayName string, iUsersMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Users)).Where("display_name = ?", iDisplayName).Update(iUsersMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsersViaSpam via map[string]interface{}{}
func UpdateUsersViaSpam(iSpam int, iUsersMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Users)).Where("spam = ?", iSpam).Update(iUsersMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateUsersViaDeleted via map[string]interface{}{}
func UpdateUsersViaDeleted(iDeleted int, iUsersMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Users)).Where("deleted = ?", iDeleted).Update(iUsersMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteUsers Delete Users
func DeleteUsers(iId int64) (int64, error) {
	row, err := Engine.Id(iId).Delete(new(Users))
	return row, err
}

// DeleteUsersViaId Delete Users via Id
func DeleteUsersViaId(iId int64) (err error) {
	var has bool
	var _Users = &Users{Id: iId}
	if has, err = Engine.Get(_Users); (has == true) && (err == nil) {
		if row, err := Engine.Where("ID = ?", iId).Delete(new(Users)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsersViaUserLogin Delete Users via UserLogin
func DeleteUsersViaUserLogin(iUserLogin string) (err error) {
	var has bool
	var _Users = &Users{UserLogin: iUserLogin}
	if has, err = Engine.Get(_Users); (has == true) && (err == nil) {
		if row, err := Engine.Where("user_login = ?", iUserLogin).Delete(new(Users)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsersViaUserPass Delete Users via UserPass
func DeleteUsersViaUserPass(iUserPass string) (err error) {
	var has bool
	var _Users = &Users{UserPass: iUserPass}
	if has, err = Engine.Get(_Users); (has == true) && (err == nil) {
		if row, err := Engine.Where("user_pass = ?", iUserPass).Delete(new(Users)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsersViaUserNicename Delete Users via UserNicename
func DeleteUsersViaUserNicename(iUserNicename string) (err error) {
	var has bool
	var _Users = &Users{UserNicename: iUserNicename}
	if has, err = Engine.Get(_Users); (has == true) && (err == nil) {
		if row, err := Engine.Where("user_nicename = ?", iUserNicename).Delete(new(Users)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsersViaUserEmail Delete Users via UserEmail
func DeleteUsersViaUserEmail(iUserEmail string) (err error) {
	var has bool
	var _Users = &Users{UserEmail: iUserEmail}
	if has, err = Engine.Get(_Users); (has == true) && (err == nil) {
		if row, err := Engine.Where("user_email = ?", iUserEmail).Delete(new(Users)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsersViaUserUrl Delete Users via UserUrl
func DeleteUsersViaUserUrl(iUserUrl string) (err error) {
	var has bool
	var _Users = &Users{UserUrl: iUserUrl}
	if has, err = Engine.Get(_Users); (has == true) && (err == nil) {
		if row, err := Engine.Where("user_url = ?", iUserUrl).Delete(new(Users)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsersViaUserRegistered Delete Users via UserRegistered
func DeleteUsersViaUserRegistered(iUserRegistered time.Time) (err error) {
	var has bool
	var _Users = &Users{UserRegistered: iUserRegistered}
	if has, err = Engine.Get(_Users); (has == true) && (err == nil) {
		if row, err := Engine.Where("user_registered = ?", iUserRegistered).Delete(new(Users)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsersViaUserActivationKey Delete Users via UserActivationKey
func DeleteUsersViaUserActivationKey(iUserActivationKey string) (err error) {
	var has bool
	var _Users = &Users{UserActivationKey: iUserActivationKey}
	if has, err = Engine.Get(_Users); (has == true) && (err == nil) {
		if row, err := Engine.Where("user_activation_key = ?", iUserActivationKey).Delete(new(Users)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsersViaUserStatus Delete Users via UserStatus
func DeleteUsersViaUserStatus(iUserStatus int) (err error) {
	var has bool
	var _Users = &Users{UserStatus: iUserStatus}
	if has, err = Engine.Get(_Users); (has == true) && (err == nil) {
		if row, err := Engine.Where("user_status = ?", iUserStatus).Delete(new(Users)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsersViaDisplayName Delete Users via DisplayName
func DeleteUsersViaDisplayName(iDisplayName string) (err error) {
	var has bool
	var _Users = &Users{DisplayName: iDisplayName}
	if has, err = Engine.Get(_Users); (has == true) && (err == nil) {
		if row, err := Engine.Where("display_name = ?", iDisplayName).Delete(new(Users)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsersViaSpam Delete Users via Spam
func DeleteUsersViaSpam(iSpam int) (err error) {
	var has bool
	var _Users = &Users{Spam: iSpam}
	if has, err = Engine.Get(_Users); (has == true) && (err == nil) {
		if row, err := Engine.Where("spam = ?", iSpam).Delete(new(Users)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteUsersViaDeleted Delete Users via Deleted
func DeleteUsersViaDeleted(iDeleted int) (err error) {
	var has bool
	var _Users = &Users{Deleted: iDeleted}
	if has, err = Engine.Get(_Users); (has == true) && (err == nil) {
		if row, err := Engine.Where("deleted = ?", iDeleted).Delete(new(Users)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
