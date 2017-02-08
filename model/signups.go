package model

import (
	"time"
)

type Signups struct {
	SignupId      int64     `xorm:"not null pk autoincr BIGINT(20)"`
	Domain        string    `xorm:"not null default '' index(domain_path) VARCHAR(200)"`
	Path          string    `xorm:"not null default '' index(domain_path) VARCHAR(100)"`
	Title         string    `xorm:"not null LONGTEXT"`
	UserLogin     string    `xorm:"not null default '' index(user_login_email) VARCHAR(60)"`
	UserEmail     string    `xorm:"not null default '' index index(user_login_email) VARCHAR(100)"`
	Registered    time.Time `xorm:"not null default '0000-00-00 00:00:00' DATETIME"`
	Activated     time.Time `xorm:"not null default '0000-00-00 00:00:00' DATETIME"`
	Active        int       `xorm:"not null default 0 TINYINT(1)"`
	ActivationKey string    `xorm:"not null default '' index VARCHAR(50)"`
	Meta          string    `xorm:"LONGTEXT"`
}

// GetSignupsesCount Signupss' Count
func GetSignupsesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&Signups{})
	return total, err
}

// GetSignupsCountViaSignupId Get Signups via SignupId
func GetSignupsCountViaSignupId(iSignupId int64) int64 {
	n, _ := Engine.Where("signup_id = ?", iSignupId).Count(&Signups{SignupId: iSignupId})
	return n
}

// GetSignupsCountViaDomain Get Signups via Domain
func GetSignupsCountViaDomain(iDomain string) int64 {
	n, _ := Engine.Where("domain = ?", iDomain).Count(&Signups{Domain: iDomain})
	return n
}

// GetSignupsCountViaPath Get Signups via Path
func GetSignupsCountViaPath(iPath string) int64 {
	n, _ := Engine.Where("path = ?", iPath).Count(&Signups{Path: iPath})
	return n
}

// GetSignupsCountViaTitle Get Signups via Title
func GetSignupsCountViaTitle(iTitle string) int64 {
	n, _ := Engine.Where("title = ?", iTitle).Count(&Signups{Title: iTitle})
	return n
}

// GetSignupsCountViaUserLogin Get Signups via UserLogin
func GetSignupsCountViaUserLogin(iUserLogin string) int64 {
	n, _ := Engine.Where("user_login = ?", iUserLogin).Count(&Signups{UserLogin: iUserLogin})
	return n
}

// GetSignupsCountViaUserEmail Get Signups via UserEmail
func GetSignupsCountViaUserEmail(iUserEmail string) int64 {
	n, _ := Engine.Where("user_email = ?", iUserEmail).Count(&Signups{UserEmail: iUserEmail})
	return n
}

// GetSignupsCountViaRegistered Get Signups via Registered
func GetSignupsCountViaRegistered(iRegistered time.Time) int64 {
	n, _ := Engine.Where("registered = ?", iRegistered).Count(&Signups{Registered: iRegistered})
	return n
}

// GetSignupsCountViaActivated Get Signups via Activated
func GetSignupsCountViaActivated(iActivated time.Time) int64 {
	n, _ := Engine.Where("activated = ?", iActivated).Count(&Signups{Activated: iActivated})
	return n
}

// GetSignupsCountViaActive Get Signups via Active
func GetSignupsCountViaActive(iActive int) int64 {
	n, _ := Engine.Where("active = ?", iActive).Count(&Signups{Active: iActive})
	return n
}

// GetSignupsCountViaActivationKey Get Signups via ActivationKey
func GetSignupsCountViaActivationKey(iActivationKey string) int64 {
	n, _ := Engine.Where("activation_key = ?", iActivationKey).Count(&Signups{ActivationKey: iActivationKey})
	return n
}

// GetSignupsCountViaMeta Get Signups via Meta
func GetSignupsCountViaMeta(iMeta string) int64 {
	n, _ := Engine.Where("meta = ?", iMeta).Count(&Signups{Meta: iMeta})
	return n
}

// GetSignupsesBySignupIdAndDomainAndPath Get Signupses via SignupIdAndDomainAndPath
func GetSignupsesBySignupIdAndDomainAndPath(offset int, limit int, SignupId_ int64, Domain_ string, Path_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and domain = ? and path = ?", SignupId_, Domain_, Path_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndDomainAndTitle Get Signupses via SignupIdAndDomainAndTitle
func GetSignupsesBySignupIdAndDomainAndTitle(offset int, limit int, SignupId_ int64, Domain_ string, Title_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and domain = ? and title = ?", SignupId_, Domain_, Title_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndDomainAndUserLogin Get Signupses via SignupIdAndDomainAndUserLogin
func GetSignupsesBySignupIdAndDomainAndUserLogin(offset int, limit int, SignupId_ int64, Domain_ string, UserLogin_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and domain = ? and user_login = ?", SignupId_, Domain_, UserLogin_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndDomainAndUserEmail Get Signupses via SignupIdAndDomainAndUserEmail
func GetSignupsesBySignupIdAndDomainAndUserEmail(offset int, limit int, SignupId_ int64, Domain_ string, UserEmail_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and domain = ? and user_email = ?", SignupId_, Domain_, UserEmail_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndDomainAndRegistered Get Signupses via SignupIdAndDomainAndRegistered
func GetSignupsesBySignupIdAndDomainAndRegistered(offset int, limit int, SignupId_ int64, Domain_ string, Registered_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and domain = ? and registered = ?", SignupId_, Domain_, Registered_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndDomainAndActivated Get Signupses via SignupIdAndDomainAndActivated
func GetSignupsesBySignupIdAndDomainAndActivated(offset int, limit int, SignupId_ int64, Domain_ string, Activated_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and domain = ? and activated = ?", SignupId_, Domain_, Activated_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndDomainAndActive Get Signupses via SignupIdAndDomainAndActive
func GetSignupsesBySignupIdAndDomainAndActive(offset int, limit int, SignupId_ int64, Domain_ string, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and domain = ? and active = ?", SignupId_, Domain_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndDomainAndActivationKey Get Signupses via SignupIdAndDomainAndActivationKey
func GetSignupsesBySignupIdAndDomainAndActivationKey(offset int, limit int, SignupId_ int64, Domain_ string, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and domain = ? and activation_key = ?", SignupId_, Domain_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndDomainAndMeta Get Signupses via SignupIdAndDomainAndMeta
func GetSignupsesBySignupIdAndDomainAndMeta(offset int, limit int, SignupId_ int64, Domain_ string, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and domain = ? and meta = ?", SignupId_, Domain_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndPathAndTitle Get Signupses via SignupIdAndPathAndTitle
func GetSignupsesBySignupIdAndPathAndTitle(offset int, limit int, SignupId_ int64, Path_ string, Title_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and path = ? and title = ?", SignupId_, Path_, Title_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndPathAndUserLogin Get Signupses via SignupIdAndPathAndUserLogin
func GetSignupsesBySignupIdAndPathAndUserLogin(offset int, limit int, SignupId_ int64, Path_ string, UserLogin_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and path = ? and user_login = ?", SignupId_, Path_, UserLogin_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndPathAndUserEmail Get Signupses via SignupIdAndPathAndUserEmail
func GetSignupsesBySignupIdAndPathAndUserEmail(offset int, limit int, SignupId_ int64, Path_ string, UserEmail_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and path = ? and user_email = ?", SignupId_, Path_, UserEmail_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndPathAndRegistered Get Signupses via SignupIdAndPathAndRegistered
func GetSignupsesBySignupIdAndPathAndRegistered(offset int, limit int, SignupId_ int64, Path_ string, Registered_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and path = ? and registered = ?", SignupId_, Path_, Registered_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndPathAndActivated Get Signupses via SignupIdAndPathAndActivated
func GetSignupsesBySignupIdAndPathAndActivated(offset int, limit int, SignupId_ int64, Path_ string, Activated_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and path = ? and activated = ?", SignupId_, Path_, Activated_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndPathAndActive Get Signupses via SignupIdAndPathAndActive
func GetSignupsesBySignupIdAndPathAndActive(offset int, limit int, SignupId_ int64, Path_ string, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and path = ? and active = ?", SignupId_, Path_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndPathAndActivationKey Get Signupses via SignupIdAndPathAndActivationKey
func GetSignupsesBySignupIdAndPathAndActivationKey(offset int, limit int, SignupId_ int64, Path_ string, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and path = ? and activation_key = ?", SignupId_, Path_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndPathAndMeta Get Signupses via SignupIdAndPathAndMeta
func GetSignupsesBySignupIdAndPathAndMeta(offset int, limit int, SignupId_ int64, Path_ string, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and path = ? and meta = ?", SignupId_, Path_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndTitleAndUserLogin Get Signupses via SignupIdAndTitleAndUserLogin
func GetSignupsesBySignupIdAndTitleAndUserLogin(offset int, limit int, SignupId_ int64, Title_ string, UserLogin_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and title = ? and user_login = ?", SignupId_, Title_, UserLogin_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndTitleAndUserEmail Get Signupses via SignupIdAndTitleAndUserEmail
func GetSignupsesBySignupIdAndTitleAndUserEmail(offset int, limit int, SignupId_ int64, Title_ string, UserEmail_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and title = ? and user_email = ?", SignupId_, Title_, UserEmail_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndTitleAndRegistered Get Signupses via SignupIdAndTitleAndRegistered
func GetSignupsesBySignupIdAndTitleAndRegistered(offset int, limit int, SignupId_ int64, Title_ string, Registered_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and title = ? and registered = ?", SignupId_, Title_, Registered_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndTitleAndActivated Get Signupses via SignupIdAndTitleAndActivated
func GetSignupsesBySignupIdAndTitleAndActivated(offset int, limit int, SignupId_ int64, Title_ string, Activated_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and title = ? and activated = ?", SignupId_, Title_, Activated_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndTitleAndActive Get Signupses via SignupIdAndTitleAndActive
func GetSignupsesBySignupIdAndTitleAndActive(offset int, limit int, SignupId_ int64, Title_ string, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and title = ? and active = ?", SignupId_, Title_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndTitleAndActivationKey Get Signupses via SignupIdAndTitleAndActivationKey
func GetSignupsesBySignupIdAndTitleAndActivationKey(offset int, limit int, SignupId_ int64, Title_ string, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and title = ? and activation_key = ?", SignupId_, Title_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndTitleAndMeta Get Signupses via SignupIdAndTitleAndMeta
func GetSignupsesBySignupIdAndTitleAndMeta(offset int, limit int, SignupId_ int64, Title_ string, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and title = ? and meta = ?", SignupId_, Title_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndUserLoginAndUserEmail Get Signupses via SignupIdAndUserLoginAndUserEmail
func GetSignupsesBySignupIdAndUserLoginAndUserEmail(offset int, limit int, SignupId_ int64, UserLogin_ string, UserEmail_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and user_login = ? and user_email = ?", SignupId_, UserLogin_, UserEmail_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndUserLoginAndRegistered Get Signupses via SignupIdAndUserLoginAndRegistered
func GetSignupsesBySignupIdAndUserLoginAndRegistered(offset int, limit int, SignupId_ int64, UserLogin_ string, Registered_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and user_login = ? and registered = ?", SignupId_, UserLogin_, Registered_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndUserLoginAndActivated Get Signupses via SignupIdAndUserLoginAndActivated
func GetSignupsesBySignupIdAndUserLoginAndActivated(offset int, limit int, SignupId_ int64, UserLogin_ string, Activated_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and user_login = ? and activated = ?", SignupId_, UserLogin_, Activated_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndUserLoginAndActive Get Signupses via SignupIdAndUserLoginAndActive
func GetSignupsesBySignupIdAndUserLoginAndActive(offset int, limit int, SignupId_ int64, UserLogin_ string, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and user_login = ? and active = ?", SignupId_, UserLogin_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndUserLoginAndActivationKey Get Signupses via SignupIdAndUserLoginAndActivationKey
func GetSignupsesBySignupIdAndUserLoginAndActivationKey(offset int, limit int, SignupId_ int64, UserLogin_ string, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and user_login = ? and activation_key = ?", SignupId_, UserLogin_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndUserLoginAndMeta Get Signupses via SignupIdAndUserLoginAndMeta
func GetSignupsesBySignupIdAndUserLoginAndMeta(offset int, limit int, SignupId_ int64, UserLogin_ string, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and user_login = ? and meta = ?", SignupId_, UserLogin_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndUserEmailAndRegistered Get Signupses via SignupIdAndUserEmailAndRegistered
func GetSignupsesBySignupIdAndUserEmailAndRegistered(offset int, limit int, SignupId_ int64, UserEmail_ string, Registered_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and user_email = ? and registered = ?", SignupId_, UserEmail_, Registered_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndUserEmailAndActivated Get Signupses via SignupIdAndUserEmailAndActivated
func GetSignupsesBySignupIdAndUserEmailAndActivated(offset int, limit int, SignupId_ int64, UserEmail_ string, Activated_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and user_email = ? and activated = ?", SignupId_, UserEmail_, Activated_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndUserEmailAndActive Get Signupses via SignupIdAndUserEmailAndActive
func GetSignupsesBySignupIdAndUserEmailAndActive(offset int, limit int, SignupId_ int64, UserEmail_ string, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and user_email = ? and active = ?", SignupId_, UserEmail_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndUserEmailAndActivationKey Get Signupses via SignupIdAndUserEmailAndActivationKey
func GetSignupsesBySignupIdAndUserEmailAndActivationKey(offset int, limit int, SignupId_ int64, UserEmail_ string, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and user_email = ? and activation_key = ?", SignupId_, UserEmail_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndUserEmailAndMeta Get Signupses via SignupIdAndUserEmailAndMeta
func GetSignupsesBySignupIdAndUserEmailAndMeta(offset int, limit int, SignupId_ int64, UserEmail_ string, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and user_email = ? and meta = ?", SignupId_, UserEmail_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndRegisteredAndActivated Get Signupses via SignupIdAndRegisteredAndActivated
func GetSignupsesBySignupIdAndRegisteredAndActivated(offset int, limit int, SignupId_ int64, Registered_ time.Time, Activated_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and registered = ? and activated = ?", SignupId_, Registered_, Activated_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndRegisteredAndActive Get Signupses via SignupIdAndRegisteredAndActive
func GetSignupsesBySignupIdAndRegisteredAndActive(offset int, limit int, SignupId_ int64, Registered_ time.Time, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and registered = ? and active = ?", SignupId_, Registered_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndRegisteredAndActivationKey Get Signupses via SignupIdAndRegisteredAndActivationKey
func GetSignupsesBySignupIdAndRegisteredAndActivationKey(offset int, limit int, SignupId_ int64, Registered_ time.Time, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and registered = ? and activation_key = ?", SignupId_, Registered_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndRegisteredAndMeta Get Signupses via SignupIdAndRegisteredAndMeta
func GetSignupsesBySignupIdAndRegisteredAndMeta(offset int, limit int, SignupId_ int64, Registered_ time.Time, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and registered = ? and meta = ?", SignupId_, Registered_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndActivatedAndActive Get Signupses via SignupIdAndActivatedAndActive
func GetSignupsesBySignupIdAndActivatedAndActive(offset int, limit int, SignupId_ int64, Activated_ time.Time, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and activated = ? and active = ?", SignupId_, Activated_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndActivatedAndActivationKey Get Signupses via SignupIdAndActivatedAndActivationKey
func GetSignupsesBySignupIdAndActivatedAndActivationKey(offset int, limit int, SignupId_ int64, Activated_ time.Time, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and activated = ? and activation_key = ?", SignupId_, Activated_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndActivatedAndMeta Get Signupses via SignupIdAndActivatedAndMeta
func GetSignupsesBySignupIdAndActivatedAndMeta(offset int, limit int, SignupId_ int64, Activated_ time.Time, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and activated = ? and meta = ?", SignupId_, Activated_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndActiveAndActivationKey Get Signupses via SignupIdAndActiveAndActivationKey
func GetSignupsesBySignupIdAndActiveAndActivationKey(offset int, limit int, SignupId_ int64, Active_ int, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and active = ? and activation_key = ?", SignupId_, Active_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndActiveAndMeta Get Signupses via SignupIdAndActiveAndMeta
func GetSignupsesBySignupIdAndActiveAndMeta(offset int, limit int, SignupId_ int64, Active_ int, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and active = ? and meta = ?", SignupId_, Active_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndActivationKeyAndMeta Get Signupses via SignupIdAndActivationKeyAndMeta
func GetSignupsesBySignupIdAndActivationKeyAndMeta(offset int, limit int, SignupId_ int64, ActivationKey_ string, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and activation_key = ? and meta = ?", SignupId_, ActivationKey_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndPathAndTitle Get Signupses via DomainAndPathAndTitle
func GetSignupsesByDomainAndPathAndTitle(offset int, limit int, Domain_ string, Path_ string, Title_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and path = ? and title = ?", Domain_, Path_, Title_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndPathAndUserLogin Get Signupses via DomainAndPathAndUserLogin
func GetSignupsesByDomainAndPathAndUserLogin(offset int, limit int, Domain_ string, Path_ string, UserLogin_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and path = ? and user_login = ?", Domain_, Path_, UserLogin_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndPathAndUserEmail Get Signupses via DomainAndPathAndUserEmail
func GetSignupsesByDomainAndPathAndUserEmail(offset int, limit int, Domain_ string, Path_ string, UserEmail_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and path = ? and user_email = ?", Domain_, Path_, UserEmail_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndPathAndRegistered Get Signupses via DomainAndPathAndRegistered
func GetSignupsesByDomainAndPathAndRegistered(offset int, limit int, Domain_ string, Path_ string, Registered_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and path = ? and registered = ?", Domain_, Path_, Registered_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndPathAndActivated Get Signupses via DomainAndPathAndActivated
func GetSignupsesByDomainAndPathAndActivated(offset int, limit int, Domain_ string, Path_ string, Activated_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and path = ? and activated = ?", Domain_, Path_, Activated_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndPathAndActive Get Signupses via DomainAndPathAndActive
func GetSignupsesByDomainAndPathAndActive(offset int, limit int, Domain_ string, Path_ string, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and path = ? and active = ?", Domain_, Path_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndPathAndActivationKey Get Signupses via DomainAndPathAndActivationKey
func GetSignupsesByDomainAndPathAndActivationKey(offset int, limit int, Domain_ string, Path_ string, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and path = ? and activation_key = ?", Domain_, Path_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndPathAndMeta Get Signupses via DomainAndPathAndMeta
func GetSignupsesByDomainAndPathAndMeta(offset int, limit int, Domain_ string, Path_ string, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and path = ? and meta = ?", Domain_, Path_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndTitleAndUserLogin Get Signupses via DomainAndTitleAndUserLogin
func GetSignupsesByDomainAndTitleAndUserLogin(offset int, limit int, Domain_ string, Title_ string, UserLogin_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and title = ? and user_login = ?", Domain_, Title_, UserLogin_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndTitleAndUserEmail Get Signupses via DomainAndTitleAndUserEmail
func GetSignupsesByDomainAndTitleAndUserEmail(offset int, limit int, Domain_ string, Title_ string, UserEmail_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and title = ? and user_email = ?", Domain_, Title_, UserEmail_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndTitleAndRegistered Get Signupses via DomainAndTitleAndRegistered
func GetSignupsesByDomainAndTitleAndRegistered(offset int, limit int, Domain_ string, Title_ string, Registered_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and title = ? and registered = ?", Domain_, Title_, Registered_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndTitleAndActivated Get Signupses via DomainAndTitleAndActivated
func GetSignupsesByDomainAndTitleAndActivated(offset int, limit int, Domain_ string, Title_ string, Activated_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and title = ? and activated = ?", Domain_, Title_, Activated_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndTitleAndActive Get Signupses via DomainAndTitleAndActive
func GetSignupsesByDomainAndTitleAndActive(offset int, limit int, Domain_ string, Title_ string, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and title = ? and active = ?", Domain_, Title_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndTitleAndActivationKey Get Signupses via DomainAndTitleAndActivationKey
func GetSignupsesByDomainAndTitleAndActivationKey(offset int, limit int, Domain_ string, Title_ string, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and title = ? and activation_key = ?", Domain_, Title_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndTitleAndMeta Get Signupses via DomainAndTitleAndMeta
func GetSignupsesByDomainAndTitleAndMeta(offset int, limit int, Domain_ string, Title_ string, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and title = ? and meta = ?", Domain_, Title_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndUserLoginAndUserEmail Get Signupses via DomainAndUserLoginAndUserEmail
func GetSignupsesByDomainAndUserLoginAndUserEmail(offset int, limit int, Domain_ string, UserLogin_ string, UserEmail_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and user_login = ? and user_email = ?", Domain_, UserLogin_, UserEmail_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndUserLoginAndRegistered Get Signupses via DomainAndUserLoginAndRegistered
func GetSignupsesByDomainAndUserLoginAndRegistered(offset int, limit int, Domain_ string, UserLogin_ string, Registered_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and user_login = ? and registered = ?", Domain_, UserLogin_, Registered_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndUserLoginAndActivated Get Signupses via DomainAndUserLoginAndActivated
func GetSignupsesByDomainAndUserLoginAndActivated(offset int, limit int, Domain_ string, UserLogin_ string, Activated_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and user_login = ? and activated = ?", Domain_, UserLogin_, Activated_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndUserLoginAndActive Get Signupses via DomainAndUserLoginAndActive
func GetSignupsesByDomainAndUserLoginAndActive(offset int, limit int, Domain_ string, UserLogin_ string, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and user_login = ? and active = ?", Domain_, UserLogin_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndUserLoginAndActivationKey Get Signupses via DomainAndUserLoginAndActivationKey
func GetSignupsesByDomainAndUserLoginAndActivationKey(offset int, limit int, Domain_ string, UserLogin_ string, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and user_login = ? and activation_key = ?", Domain_, UserLogin_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndUserLoginAndMeta Get Signupses via DomainAndUserLoginAndMeta
func GetSignupsesByDomainAndUserLoginAndMeta(offset int, limit int, Domain_ string, UserLogin_ string, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and user_login = ? and meta = ?", Domain_, UserLogin_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndUserEmailAndRegistered Get Signupses via DomainAndUserEmailAndRegistered
func GetSignupsesByDomainAndUserEmailAndRegistered(offset int, limit int, Domain_ string, UserEmail_ string, Registered_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and user_email = ? and registered = ?", Domain_, UserEmail_, Registered_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndUserEmailAndActivated Get Signupses via DomainAndUserEmailAndActivated
func GetSignupsesByDomainAndUserEmailAndActivated(offset int, limit int, Domain_ string, UserEmail_ string, Activated_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and user_email = ? and activated = ?", Domain_, UserEmail_, Activated_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndUserEmailAndActive Get Signupses via DomainAndUserEmailAndActive
func GetSignupsesByDomainAndUserEmailAndActive(offset int, limit int, Domain_ string, UserEmail_ string, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and user_email = ? and active = ?", Domain_, UserEmail_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndUserEmailAndActivationKey Get Signupses via DomainAndUserEmailAndActivationKey
func GetSignupsesByDomainAndUserEmailAndActivationKey(offset int, limit int, Domain_ string, UserEmail_ string, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and user_email = ? and activation_key = ?", Domain_, UserEmail_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndUserEmailAndMeta Get Signupses via DomainAndUserEmailAndMeta
func GetSignupsesByDomainAndUserEmailAndMeta(offset int, limit int, Domain_ string, UserEmail_ string, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and user_email = ? and meta = ?", Domain_, UserEmail_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndRegisteredAndActivated Get Signupses via DomainAndRegisteredAndActivated
func GetSignupsesByDomainAndRegisteredAndActivated(offset int, limit int, Domain_ string, Registered_ time.Time, Activated_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and registered = ? and activated = ?", Domain_, Registered_, Activated_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndRegisteredAndActive Get Signupses via DomainAndRegisteredAndActive
func GetSignupsesByDomainAndRegisteredAndActive(offset int, limit int, Domain_ string, Registered_ time.Time, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and registered = ? and active = ?", Domain_, Registered_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndRegisteredAndActivationKey Get Signupses via DomainAndRegisteredAndActivationKey
func GetSignupsesByDomainAndRegisteredAndActivationKey(offset int, limit int, Domain_ string, Registered_ time.Time, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and registered = ? and activation_key = ?", Domain_, Registered_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndRegisteredAndMeta Get Signupses via DomainAndRegisteredAndMeta
func GetSignupsesByDomainAndRegisteredAndMeta(offset int, limit int, Domain_ string, Registered_ time.Time, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and registered = ? and meta = ?", Domain_, Registered_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndActivatedAndActive Get Signupses via DomainAndActivatedAndActive
func GetSignupsesByDomainAndActivatedAndActive(offset int, limit int, Domain_ string, Activated_ time.Time, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and activated = ? and active = ?", Domain_, Activated_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndActivatedAndActivationKey Get Signupses via DomainAndActivatedAndActivationKey
func GetSignupsesByDomainAndActivatedAndActivationKey(offset int, limit int, Domain_ string, Activated_ time.Time, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and activated = ? and activation_key = ?", Domain_, Activated_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndActivatedAndMeta Get Signupses via DomainAndActivatedAndMeta
func GetSignupsesByDomainAndActivatedAndMeta(offset int, limit int, Domain_ string, Activated_ time.Time, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and activated = ? and meta = ?", Domain_, Activated_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndActiveAndActivationKey Get Signupses via DomainAndActiveAndActivationKey
func GetSignupsesByDomainAndActiveAndActivationKey(offset int, limit int, Domain_ string, Active_ int, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and active = ? and activation_key = ?", Domain_, Active_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndActiveAndMeta Get Signupses via DomainAndActiveAndMeta
func GetSignupsesByDomainAndActiveAndMeta(offset int, limit int, Domain_ string, Active_ int, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and active = ? and meta = ?", Domain_, Active_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndActivationKeyAndMeta Get Signupses via DomainAndActivationKeyAndMeta
func GetSignupsesByDomainAndActivationKeyAndMeta(offset int, limit int, Domain_ string, ActivationKey_ string, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and activation_key = ? and meta = ?", Domain_, ActivationKey_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndTitleAndUserLogin Get Signupses via PathAndTitleAndUserLogin
func GetSignupsesByPathAndTitleAndUserLogin(offset int, limit int, Path_ string, Title_ string, UserLogin_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and title = ? and user_login = ?", Path_, Title_, UserLogin_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndTitleAndUserEmail Get Signupses via PathAndTitleAndUserEmail
func GetSignupsesByPathAndTitleAndUserEmail(offset int, limit int, Path_ string, Title_ string, UserEmail_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and title = ? and user_email = ?", Path_, Title_, UserEmail_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndTitleAndRegistered Get Signupses via PathAndTitleAndRegistered
func GetSignupsesByPathAndTitleAndRegistered(offset int, limit int, Path_ string, Title_ string, Registered_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and title = ? and registered = ?", Path_, Title_, Registered_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndTitleAndActivated Get Signupses via PathAndTitleAndActivated
func GetSignupsesByPathAndTitleAndActivated(offset int, limit int, Path_ string, Title_ string, Activated_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and title = ? and activated = ?", Path_, Title_, Activated_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndTitleAndActive Get Signupses via PathAndTitleAndActive
func GetSignupsesByPathAndTitleAndActive(offset int, limit int, Path_ string, Title_ string, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and title = ? and active = ?", Path_, Title_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndTitleAndActivationKey Get Signupses via PathAndTitleAndActivationKey
func GetSignupsesByPathAndTitleAndActivationKey(offset int, limit int, Path_ string, Title_ string, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and title = ? and activation_key = ?", Path_, Title_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndTitleAndMeta Get Signupses via PathAndTitleAndMeta
func GetSignupsesByPathAndTitleAndMeta(offset int, limit int, Path_ string, Title_ string, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and title = ? and meta = ?", Path_, Title_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndUserLoginAndUserEmail Get Signupses via PathAndUserLoginAndUserEmail
func GetSignupsesByPathAndUserLoginAndUserEmail(offset int, limit int, Path_ string, UserLogin_ string, UserEmail_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and user_login = ? and user_email = ?", Path_, UserLogin_, UserEmail_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndUserLoginAndRegistered Get Signupses via PathAndUserLoginAndRegistered
func GetSignupsesByPathAndUserLoginAndRegistered(offset int, limit int, Path_ string, UserLogin_ string, Registered_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and user_login = ? and registered = ?", Path_, UserLogin_, Registered_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndUserLoginAndActivated Get Signupses via PathAndUserLoginAndActivated
func GetSignupsesByPathAndUserLoginAndActivated(offset int, limit int, Path_ string, UserLogin_ string, Activated_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and user_login = ? and activated = ?", Path_, UserLogin_, Activated_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndUserLoginAndActive Get Signupses via PathAndUserLoginAndActive
func GetSignupsesByPathAndUserLoginAndActive(offset int, limit int, Path_ string, UserLogin_ string, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and user_login = ? and active = ?", Path_, UserLogin_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndUserLoginAndActivationKey Get Signupses via PathAndUserLoginAndActivationKey
func GetSignupsesByPathAndUserLoginAndActivationKey(offset int, limit int, Path_ string, UserLogin_ string, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and user_login = ? and activation_key = ?", Path_, UserLogin_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndUserLoginAndMeta Get Signupses via PathAndUserLoginAndMeta
func GetSignupsesByPathAndUserLoginAndMeta(offset int, limit int, Path_ string, UserLogin_ string, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and user_login = ? and meta = ?", Path_, UserLogin_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndUserEmailAndRegistered Get Signupses via PathAndUserEmailAndRegistered
func GetSignupsesByPathAndUserEmailAndRegistered(offset int, limit int, Path_ string, UserEmail_ string, Registered_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and user_email = ? and registered = ?", Path_, UserEmail_, Registered_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndUserEmailAndActivated Get Signupses via PathAndUserEmailAndActivated
func GetSignupsesByPathAndUserEmailAndActivated(offset int, limit int, Path_ string, UserEmail_ string, Activated_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and user_email = ? and activated = ?", Path_, UserEmail_, Activated_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndUserEmailAndActive Get Signupses via PathAndUserEmailAndActive
func GetSignupsesByPathAndUserEmailAndActive(offset int, limit int, Path_ string, UserEmail_ string, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and user_email = ? and active = ?", Path_, UserEmail_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndUserEmailAndActivationKey Get Signupses via PathAndUserEmailAndActivationKey
func GetSignupsesByPathAndUserEmailAndActivationKey(offset int, limit int, Path_ string, UserEmail_ string, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and user_email = ? and activation_key = ?", Path_, UserEmail_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndUserEmailAndMeta Get Signupses via PathAndUserEmailAndMeta
func GetSignupsesByPathAndUserEmailAndMeta(offset int, limit int, Path_ string, UserEmail_ string, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and user_email = ? and meta = ?", Path_, UserEmail_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndRegisteredAndActivated Get Signupses via PathAndRegisteredAndActivated
func GetSignupsesByPathAndRegisteredAndActivated(offset int, limit int, Path_ string, Registered_ time.Time, Activated_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and registered = ? and activated = ?", Path_, Registered_, Activated_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndRegisteredAndActive Get Signupses via PathAndRegisteredAndActive
func GetSignupsesByPathAndRegisteredAndActive(offset int, limit int, Path_ string, Registered_ time.Time, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and registered = ? and active = ?", Path_, Registered_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndRegisteredAndActivationKey Get Signupses via PathAndRegisteredAndActivationKey
func GetSignupsesByPathAndRegisteredAndActivationKey(offset int, limit int, Path_ string, Registered_ time.Time, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and registered = ? and activation_key = ?", Path_, Registered_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndRegisteredAndMeta Get Signupses via PathAndRegisteredAndMeta
func GetSignupsesByPathAndRegisteredAndMeta(offset int, limit int, Path_ string, Registered_ time.Time, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and registered = ? and meta = ?", Path_, Registered_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndActivatedAndActive Get Signupses via PathAndActivatedAndActive
func GetSignupsesByPathAndActivatedAndActive(offset int, limit int, Path_ string, Activated_ time.Time, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and activated = ? and active = ?", Path_, Activated_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndActivatedAndActivationKey Get Signupses via PathAndActivatedAndActivationKey
func GetSignupsesByPathAndActivatedAndActivationKey(offset int, limit int, Path_ string, Activated_ time.Time, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and activated = ? and activation_key = ?", Path_, Activated_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndActivatedAndMeta Get Signupses via PathAndActivatedAndMeta
func GetSignupsesByPathAndActivatedAndMeta(offset int, limit int, Path_ string, Activated_ time.Time, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and activated = ? and meta = ?", Path_, Activated_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndActiveAndActivationKey Get Signupses via PathAndActiveAndActivationKey
func GetSignupsesByPathAndActiveAndActivationKey(offset int, limit int, Path_ string, Active_ int, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and active = ? and activation_key = ?", Path_, Active_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndActiveAndMeta Get Signupses via PathAndActiveAndMeta
func GetSignupsesByPathAndActiveAndMeta(offset int, limit int, Path_ string, Active_ int, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and active = ? and meta = ?", Path_, Active_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndActivationKeyAndMeta Get Signupses via PathAndActivationKeyAndMeta
func GetSignupsesByPathAndActivationKeyAndMeta(offset int, limit int, Path_ string, ActivationKey_ string, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and activation_key = ? and meta = ?", Path_, ActivationKey_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByTitleAndUserLoginAndUserEmail Get Signupses via TitleAndUserLoginAndUserEmail
func GetSignupsesByTitleAndUserLoginAndUserEmail(offset int, limit int, Title_ string, UserLogin_ string, UserEmail_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("title = ? and user_login = ? and user_email = ?", Title_, UserLogin_, UserEmail_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByTitleAndUserLoginAndRegistered Get Signupses via TitleAndUserLoginAndRegistered
func GetSignupsesByTitleAndUserLoginAndRegistered(offset int, limit int, Title_ string, UserLogin_ string, Registered_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("title = ? and user_login = ? and registered = ?", Title_, UserLogin_, Registered_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByTitleAndUserLoginAndActivated Get Signupses via TitleAndUserLoginAndActivated
func GetSignupsesByTitleAndUserLoginAndActivated(offset int, limit int, Title_ string, UserLogin_ string, Activated_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("title = ? and user_login = ? and activated = ?", Title_, UserLogin_, Activated_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByTitleAndUserLoginAndActive Get Signupses via TitleAndUserLoginAndActive
func GetSignupsesByTitleAndUserLoginAndActive(offset int, limit int, Title_ string, UserLogin_ string, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("title = ? and user_login = ? and active = ?", Title_, UserLogin_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByTitleAndUserLoginAndActivationKey Get Signupses via TitleAndUserLoginAndActivationKey
func GetSignupsesByTitleAndUserLoginAndActivationKey(offset int, limit int, Title_ string, UserLogin_ string, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("title = ? and user_login = ? and activation_key = ?", Title_, UserLogin_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByTitleAndUserLoginAndMeta Get Signupses via TitleAndUserLoginAndMeta
func GetSignupsesByTitleAndUserLoginAndMeta(offset int, limit int, Title_ string, UserLogin_ string, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("title = ? and user_login = ? and meta = ?", Title_, UserLogin_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByTitleAndUserEmailAndRegistered Get Signupses via TitleAndUserEmailAndRegistered
func GetSignupsesByTitleAndUserEmailAndRegistered(offset int, limit int, Title_ string, UserEmail_ string, Registered_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("title = ? and user_email = ? and registered = ?", Title_, UserEmail_, Registered_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByTitleAndUserEmailAndActivated Get Signupses via TitleAndUserEmailAndActivated
func GetSignupsesByTitleAndUserEmailAndActivated(offset int, limit int, Title_ string, UserEmail_ string, Activated_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("title = ? and user_email = ? and activated = ?", Title_, UserEmail_, Activated_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByTitleAndUserEmailAndActive Get Signupses via TitleAndUserEmailAndActive
func GetSignupsesByTitleAndUserEmailAndActive(offset int, limit int, Title_ string, UserEmail_ string, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("title = ? and user_email = ? and active = ?", Title_, UserEmail_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByTitleAndUserEmailAndActivationKey Get Signupses via TitleAndUserEmailAndActivationKey
func GetSignupsesByTitleAndUserEmailAndActivationKey(offset int, limit int, Title_ string, UserEmail_ string, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("title = ? and user_email = ? and activation_key = ?", Title_, UserEmail_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByTitleAndUserEmailAndMeta Get Signupses via TitleAndUserEmailAndMeta
func GetSignupsesByTitleAndUserEmailAndMeta(offset int, limit int, Title_ string, UserEmail_ string, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("title = ? and user_email = ? and meta = ?", Title_, UserEmail_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByTitleAndRegisteredAndActivated Get Signupses via TitleAndRegisteredAndActivated
func GetSignupsesByTitleAndRegisteredAndActivated(offset int, limit int, Title_ string, Registered_ time.Time, Activated_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("title = ? and registered = ? and activated = ?", Title_, Registered_, Activated_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByTitleAndRegisteredAndActive Get Signupses via TitleAndRegisteredAndActive
func GetSignupsesByTitleAndRegisteredAndActive(offset int, limit int, Title_ string, Registered_ time.Time, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("title = ? and registered = ? and active = ?", Title_, Registered_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByTitleAndRegisteredAndActivationKey Get Signupses via TitleAndRegisteredAndActivationKey
func GetSignupsesByTitleAndRegisteredAndActivationKey(offset int, limit int, Title_ string, Registered_ time.Time, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("title = ? and registered = ? and activation_key = ?", Title_, Registered_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByTitleAndRegisteredAndMeta Get Signupses via TitleAndRegisteredAndMeta
func GetSignupsesByTitleAndRegisteredAndMeta(offset int, limit int, Title_ string, Registered_ time.Time, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("title = ? and registered = ? and meta = ?", Title_, Registered_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByTitleAndActivatedAndActive Get Signupses via TitleAndActivatedAndActive
func GetSignupsesByTitleAndActivatedAndActive(offset int, limit int, Title_ string, Activated_ time.Time, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("title = ? and activated = ? and active = ?", Title_, Activated_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByTitleAndActivatedAndActivationKey Get Signupses via TitleAndActivatedAndActivationKey
func GetSignupsesByTitleAndActivatedAndActivationKey(offset int, limit int, Title_ string, Activated_ time.Time, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("title = ? and activated = ? and activation_key = ?", Title_, Activated_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByTitleAndActivatedAndMeta Get Signupses via TitleAndActivatedAndMeta
func GetSignupsesByTitleAndActivatedAndMeta(offset int, limit int, Title_ string, Activated_ time.Time, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("title = ? and activated = ? and meta = ?", Title_, Activated_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByTitleAndActiveAndActivationKey Get Signupses via TitleAndActiveAndActivationKey
func GetSignupsesByTitleAndActiveAndActivationKey(offset int, limit int, Title_ string, Active_ int, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("title = ? and active = ? and activation_key = ?", Title_, Active_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByTitleAndActiveAndMeta Get Signupses via TitleAndActiveAndMeta
func GetSignupsesByTitleAndActiveAndMeta(offset int, limit int, Title_ string, Active_ int, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("title = ? and active = ? and meta = ?", Title_, Active_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByTitleAndActivationKeyAndMeta Get Signupses via TitleAndActivationKeyAndMeta
func GetSignupsesByTitleAndActivationKeyAndMeta(offset int, limit int, Title_ string, ActivationKey_ string, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("title = ? and activation_key = ? and meta = ?", Title_, ActivationKey_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserLoginAndUserEmailAndRegistered Get Signupses via UserLoginAndUserEmailAndRegistered
func GetSignupsesByUserLoginAndUserEmailAndRegistered(offset int, limit int, UserLogin_ string, UserEmail_ string, Registered_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_login = ? and user_email = ? and registered = ?", UserLogin_, UserEmail_, Registered_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserLoginAndUserEmailAndActivated Get Signupses via UserLoginAndUserEmailAndActivated
func GetSignupsesByUserLoginAndUserEmailAndActivated(offset int, limit int, UserLogin_ string, UserEmail_ string, Activated_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_login = ? and user_email = ? and activated = ?", UserLogin_, UserEmail_, Activated_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserLoginAndUserEmailAndActive Get Signupses via UserLoginAndUserEmailAndActive
func GetSignupsesByUserLoginAndUserEmailAndActive(offset int, limit int, UserLogin_ string, UserEmail_ string, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_login = ? and user_email = ? and active = ?", UserLogin_, UserEmail_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserLoginAndUserEmailAndActivationKey Get Signupses via UserLoginAndUserEmailAndActivationKey
func GetSignupsesByUserLoginAndUserEmailAndActivationKey(offset int, limit int, UserLogin_ string, UserEmail_ string, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_login = ? and user_email = ? and activation_key = ?", UserLogin_, UserEmail_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserLoginAndUserEmailAndMeta Get Signupses via UserLoginAndUserEmailAndMeta
func GetSignupsesByUserLoginAndUserEmailAndMeta(offset int, limit int, UserLogin_ string, UserEmail_ string, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_login = ? and user_email = ? and meta = ?", UserLogin_, UserEmail_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserLoginAndRegisteredAndActivated Get Signupses via UserLoginAndRegisteredAndActivated
func GetSignupsesByUserLoginAndRegisteredAndActivated(offset int, limit int, UserLogin_ string, Registered_ time.Time, Activated_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_login = ? and registered = ? and activated = ?", UserLogin_, Registered_, Activated_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserLoginAndRegisteredAndActive Get Signupses via UserLoginAndRegisteredAndActive
func GetSignupsesByUserLoginAndRegisteredAndActive(offset int, limit int, UserLogin_ string, Registered_ time.Time, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_login = ? and registered = ? and active = ?", UserLogin_, Registered_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserLoginAndRegisteredAndActivationKey Get Signupses via UserLoginAndRegisteredAndActivationKey
func GetSignupsesByUserLoginAndRegisteredAndActivationKey(offset int, limit int, UserLogin_ string, Registered_ time.Time, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_login = ? and registered = ? and activation_key = ?", UserLogin_, Registered_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserLoginAndRegisteredAndMeta Get Signupses via UserLoginAndRegisteredAndMeta
func GetSignupsesByUserLoginAndRegisteredAndMeta(offset int, limit int, UserLogin_ string, Registered_ time.Time, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_login = ? and registered = ? and meta = ?", UserLogin_, Registered_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserLoginAndActivatedAndActive Get Signupses via UserLoginAndActivatedAndActive
func GetSignupsesByUserLoginAndActivatedAndActive(offset int, limit int, UserLogin_ string, Activated_ time.Time, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_login = ? and activated = ? and active = ?", UserLogin_, Activated_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserLoginAndActivatedAndActivationKey Get Signupses via UserLoginAndActivatedAndActivationKey
func GetSignupsesByUserLoginAndActivatedAndActivationKey(offset int, limit int, UserLogin_ string, Activated_ time.Time, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_login = ? and activated = ? and activation_key = ?", UserLogin_, Activated_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserLoginAndActivatedAndMeta Get Signupses via UserLoginAndActivatedAndMeta
func GetSignupsesByUserLoginAndActivatedAndMeta(offset int, limit int, UserLogin_ string, Activated_ time.Time, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_login = ? and activated = ? and meta = ?", UserLogin_, Activated_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserLoginAndActiveAndActivationKey Get Signupses via UserLoginAndActiveAndActivationKey
func GetSignupsesByUserLoginAndActiveAndActivationKey(offset int, limit int, UserLogin_ string, Active_ int, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_login = ? and active = ? and activation_key = ?", UserLogin_, Active_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserLoginAndActiveAndMeta Get Signupses via UserLoginAndActiveAndMeta
func GetSignupsesByUserLoginAndActiveAndMeta(offset int, limit int, UserLogin_ string, Active_ int, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_login = ? and active = ? and meta = ?", UserLogin_, Active_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserLoginAndActivationKeyAndMeta Get Signupses via UserLoginAndActivationKeyAndMeta
func GetSignupsesByUserLoginAndActivationKeyAndMeta(offset int, limit int, UserLogin_ string, ActivationKey_ string, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_login = ? and activation_key = ? and meta = ?", UserLogin_, ActivationKey_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserEmailAndRegisteredAndActivated Get Signupses via UserEmailAndRegisteredAndActivated
func GetSignupsesByUserEmailAndRegisteredAndActivated(offset int, limit int, UserEmail_ string, Registered_ time.Time, Activated_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_email = ? and registered = ? and activated = ?", UserEmail_, Registered_, Activated_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserEmailAndRegisteredAndActive Get Signupses via UserEmailAndRegisteredAndActive
func GetSignupsesByUserEmailAndRegisteredAndActive(offset int, limit int, UserEmail_ string, Registered_ time.Time, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_email = ? and registered = ? and active = ?", UserEmail_, Registered_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserEmailAndRegisteredAndActivationKey Get Signupses via UserEmailAndRegisteredAndActivationKey
func GetSignupsesByUserEmailAndRegisteredAndActivationKey(offset int, limit int, UserEmail_ string, Registered_ time.Time, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_email = ? and registered = ? and activation_key = ?", UserEmail_, Registered_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserEmailAndRegisteredAndMeta Get Signupses via UserEmailAndRegisteredAndMeta
func GetSignupsesByUserEmailAndRegisteredAndMeta(offset int, limit int, UserEmail_ string, Registered_ time.Time, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_email = ? and registered = ? and meta = ?", UserEmail_, Registered_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserEmailAndActivatedAndActive Get Signupses via UserEmailAndActivatedAndActive
func GetSignupsesByUserEmailAndActivatedAndActive(offset int, limit int, UserEmail_ string, Activated_ time.Time, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_email = ? and activated = ? and active = ?", UserEmail_, Activated_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserEmailAndActivatedAndActivationKey Get Signupses via UserEmailAndActivatedAndActivationKey
func GetSignupsesByUserEmailAndActivatedAndActivationKey(offset int, limit int, UserEmail_ string, Activated_ time.Time, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_email = ? and activated = ? and activation_key = ?", UserEmail_, Activated_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserEmailAndActivatedAndMeta Get Signupses via UserEmailAndActivatedAndMeta
func GetSignupsesByUserEmailAndActivatedAndMeta(offset int, limit int, UserEmail_ string, Activated_ time.Time, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_email = ? and activated = ? and meta = ?", UserEmail_, Activated_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserEmailAndActiveAndActivationKey Get Signupses via UserEmailAndActiveAndActivationKey
func GetSignupsesByUserEmailAndActiveAndActivationKey(offset int, limit int, UserEmail_ string, Active_ int, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_email = ? and active = ? and activation_key = ?", UserEmail_, Active_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserEmailAndActiveAndMeta Get Signupses via UserEmailAndActiveAndMeta
func GetSignupsesByUserEmailAndActiveAndMeta(offset int, limit int, UserEmail_ string, Active_ int, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_email = ? and active = ? and meta = ?", UserEmail_, Active_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserEmailAndActivationKeyAndMeta Get Signupses via UserEmailAndActivationKeyAndMeta
func GetSignupsesByUserEmailAndActivationKeyAndMeta(offset int, limit int, UserEmail_ string, ActivationKey_ string, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_email = ? and activation_key = ? and meta = ?", UserEmail_, ActivationKey_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByRegisteredAndActivatedAndActive Get Signupses via RegisteredAndActivatedAndActive
func GetSignupsesByRegisteredAndActivatedAndActive(offset int, limit int, Registered_ time.Time, Activated_ time.Time, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("registered = ? and activated = ? and active = ?", Registered_, Activated_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByRegisteredAndActivatedAndActivationKey Get Signupses via RegisteredAndActivatedAndActivationKey
func GetSignupsesByRegisteredAndActivatedAndActivationKey(offset int, limit int, Registered_ time.Time, Activated_ time.Time, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("registered = ? and activated = ? and activation_key = ?", Registered_, Activated_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByRegisteredAndActivatedAndMeta Get Signupses via RegisteredAndActivatedAndMeta
func GetSignupsesByRegisteredAndActivatedAndMeta(offset int, limit int, Registered_ time.Time, Activated_ time.Time, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("registered = ? and activated = ? and meta = ?", Registered_, Activated_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByRegisteredAndActiveAndActivationKey Get Signupses via RegisteredAndActiveAndActivationKey
func GetSignupsesByRegisteredAndActiveAndActivationKey(offset int, limit int, Registered_ time.Time, Active_ int, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("registered = ? and active = ? and activation_key = ?", Registered_, Active_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByRegisteredAndActiveAndMeta Get Signupses via RegisteredAndActiveAndMeta
func GetSignupsesByRegisteredAndActiveAndMeta(offset int, limit int, Registered_ time.Time, Active_ int, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("registered = ? and active = ? and meta = ?", Registered_, Active_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByRegisteredAndActivationKeyAndMeta Get Signupses via RegisteredAndActivationKeyAndMeta
func GetSignupsesByRegisteredAndActivationKeyAndMeta(offset int, limit int, Registered_ time.Time, ActivationKey_ string, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("registered = ? and activation_key = ? and meta = ?", Registered_, ActivationKey_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByActivatedAndActiveAndActivationKey Get Signupses via ActivatedAndActiveAndActivationKey
func GetSignupsesByActivatedAndActiveAndActivationKey(offset int, limit int, Activated_ time.Time, Active_ int, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("activated = ? and active = ? and activation_key = ?", Activated_, Active_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByActivatedAndActiveAndMeta Get Signupses via ActivatedAndActiveAndMeta
func GetSignupsesByActivatedAndActiveAndMeta(offset int, limit int, Activated_ time.Time, Active_ int, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("activated = ? and active = ? and meta = ?", Activated_, Active_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByActivatedAndActivationKeyAndMeta Get Signupses via ActivatedAndActivationKeyAndMeta
func GetSignupsesByActivatedAndActivationKeyAndMeta(offset int, limit int, Activated_ time.Time, ActivationKey_ string, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("activated = ? and activation_key = ? and meta = ?", Activated_, ActivationKey_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByActiveAndActivationKeyAndMeta Get Signupses via ActiveAndActivationKeyAndMeta
func GetSignupsesByActiveAndActivationKeyAndMeta(offset int, limit int, Active_ int, ActivationKey_ string, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("active = ? and activation_key = ? and meta = ?", Active_, ActivationKey_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndDomain Get Signupses via SignupIdAndDomain
func GetSignupsesBySignupIdAndDomain(offset int, limit int, SignupId_ int64, Domain_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and domain = ?", SignupId_, Domain_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndPath Get Signupses via SignupIdAndPath
func GetSignupsesBySignupIdAndPath(offset int, limit int, SignupId_ int64, Path_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and path = ?", SignupId_, Path_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndTitle Get Signupses via SignupIdAndTitle
func GetSignupsesBySignupIdAndTitle(offset int, limit int, SignupId_ int64, Title_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and title = ?", SignupId_, Title_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndUserLogin Get Signupses via SignupIdAndUserLogin
func GetSignupsesBySignupIdAndUserLogin(offset int, limit int, SignupId_ int64, UserLogin_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and user_login = ?", SignupId_, UserLogin_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndUserEmail Get Signupses via SignupIdAndUserEmail
func GetSignupsesBySignupIdAndUserEmail(offset int, limit int, SignupId_ int64, UserEmail_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and user_email = ?", SignupId_, UserEmail_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndRegistered Get Signupses via SignupIdAndRegistered
func GetSignupsesBySignupIdAndRegistered(offset int, limit int, SignupId_ int64, Registered_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and registered = ?", SignupId_, Registered_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndActivated Get Signupses via SignupIdAndActivated
func GetSignupsesBySignupIdAndActivated(offset int, limit int, SignupId_ int64, Activated_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and activated = ?", SignupId_, Activated_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndActive Get Signupses via SignupIdAndActive
func GetSignupsesBySignupIdAndActive(offset int, limit int, SignupId_ int64, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and active = ?", SignupId_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndActivationKey Get Signupses via SignupIdAndActivationKey
func GetSignupsesBySignupIdAndActivationKey(offset int, limit int, SignupId_ int64, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and activation_key = ?", SignupId_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesBySignupIdAndMeta Get Signupses via SignupIdAndMeta
func GetSignupsesBySignupIdAndMeta(offset int, limit int, SignupId_ int64, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ? and meta = ?", SignupId_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndPath Get Signupses via DomainAndPath
func GetSignupsesByDomainAndPath(offset int, limit int, Domain_ string, Path_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and path = ?", Domain_, Path_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndTitle Get Signupses via DomainAndTitle
func GetSignupsesByDomainAndTitle(offset int, limit int, Domain_ string, Title_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and title = ?", Domain_, Title_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndUserLogin Get Signupses via DomainAndUserLogin
func GetSignupsesByDomainAndUserLogin(offset int, limit int, Domain_ string, UserLogin_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and user_login = ?", Domain_, UserLogin_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndUserEmail Get Signupses via DomainAndUserEmail
func GetSignupsesByDomainAndUserEmail(offset int, limit int, Domain_ string, UserEmail_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and user_email = ?", Domain_, UserEmail_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndRegistered Get Signupses via DomainAndRegistered
func GetSignupsesByDomainAndRegistered(offset int, limit int, Domain_ string, Registered_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and registered = ?", Domain_, Registered_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndActivated Get Signupses via DomainAndActivated
func GetSignupsesByDomainAndActivated(offset int, limit int, Domain_ string, Activated_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and activated = ?", Domain_, Activated_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndActive Get Signupses via DomainAndActive
func GetSignupsesByDomainAndActive(offset int, limit int, Domain_ string, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and active = ?", Domain_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndActivationKey Get Signupses via DomainAndActivationKey
func GetSignupsesByDomainAndActivationKey(offset int, limit int, Domain_ string, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and activation_key = ?", Domain_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByDomainAndMeta Get Signupses via DomainAndMeta
func GetSignupsesByDomainAndMeta(offset int, limit int, Domain_ string, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ? and meta = ?", Domain_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndTitle Get Signupses via PathAndTitle
func GetSignupsesByPathAndTitle(offset int, limit int, Path_ string, Title_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and title = ?", Path_, Title_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndUserLogin Get Signupses via PathAndUserLogin
func GetSignupsesByPathAndUserLogin(offset int, limit int, Path_ string, UserLogin_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and user_login = ?", Path_, UserLogin_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndUserEmail Get Signupses via PathAndUserEmail
func GetSignupsesByPathAndUserEmail(offset int, limit int, Path_ string, UserEmail_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and user_email = ?", Path_, UserEmail_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndRegistered Get Signupses via PathAndRegistered
func GetSignupsesByPathAndRegistered(offset int, limit int, Path_ string, Registered_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and registered = ?", Path_, Registered_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndActivated Get Signupses via PathAndActivated
func GetSignupsesByPathAndActivated(offset int, limit int, Path_ string, Activated_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and activated = ?", Path_, Activated_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndActive Get Signupses via PathAndActive
func GetSignupsesByPathAndActive(offset int, limit int, Path_ string, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and active = ?", Path_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndActivationKey Get Signupses via PathAndActivationKey
func GetSignupsesByPathAndActivationKey(offset int, limit int, Path_ string, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and activation_key = ?", Path_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByPathAndMeta Get Signupses via PathAndMeta
func GetSignupsesByPathAndMeta(offset int, limit int, Path_ string, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ? and meta = ?", Path_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByTitleAndUserLogin Get Signupses via TitleAndUserLogin
func GetSignupsesByTitleAndUserLogin(offset int, limit int, Title_ string, UserLogin_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("title = ? and user_login = ?", Title_, UserLogin_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByTitleAndUserEmail Get Signupses via TitleAndUserEmail
func GetSignupsesByTitleAndUserEmail(offset int, limit int, Title_ string, UserEmail_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("title = ? and user_email = ?", Title_, UserEmail_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByTitleAndRegistered Get Signupses via TitleAndRegistered
func GetSignupsesByTitleAndRegistered(offset int, limit int, Title_ string, Registered_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("title = ? and registered = ?", Title_, Registered_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByTitleAndActivated Get Signupses via TitleAndActivated
func GetSignupsesByTitleAndActivated(offset int, limit int, Title_ string, Activated_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("title = ? and activated = ?", Title_, Activated_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByTitleAndActive Get Signupses via TitleAndActive
func GetSignupsesByTitleAndActive(offset int, limit int, Title_ string, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("title = ? and active = ?", Title_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByTitleAndActivationKey Get Signupses via TitleAndActivationKey
func GetSignupsesByTitleAndActivationKey(offset int, limit int, Title_ string, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("title = ? and activation_key = ?", Title_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByTitleAndMeta Get Signupses via TitleAndMeta
func GetSignupsesByTitleAndMeta(offset int, limit int, Title_ string, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("title = ? and meta = ?", Title_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserLoginAndUserEmail Get Signupses via UserLoginAndUserEmail
func GetSignupsesByUserLoginAndUserEmail(offset int, limit int, UserLogin_ string, UserEmail_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_login = ? and user_email = ?", UserLogin_, UserEmail_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserLoginAndRegistered Get Signupses via UserLoginAndRegistered
func GetSignupsesByUserLoginAndRegistered(offset int, limit int, UserLogin_ string, Registered_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_login = ? and registered = ?", UserLogin_, Registered_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserLoginAndActivated Get Signupses via UserLoginAndActivated
func GetSignupsesByUserLoginAndActivated(offset int, limit int, UserLogin_ string, Activated_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_login = ? and activated = ?", UserLogin_, Activated_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserLoginAndActive Get Signupses via UserLoginAndActive
func GetSignupsesByUserLoginAndActive(offset int, limit int, UserLogin_ string, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_login = ? and active = ?", UserLogin_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserLoginAndActivationKey Get Signupses via UserLoginAndActivationKey
func GetSignupsesByUserLoginAndActivationKey(offset int, limit int, UserLogin_ string, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_login = ? and activation_key = ?", UserLogin_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserLoginAndMeta Get Signupses via UserLoginAndMeta
func GetSignupsesByUserLoginAndMeta(offset int, limit int, UserLogin_ string, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_login = ? and meta = ?", UserLogin_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserEmailAndRegistered Get Signupses via UserEmailAndRegistered
func GetSignupsesByUserEmailAndRegistered(offset int, limit int, UserEmail_ string, Registered_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_email = ? and registered = ?", UserEmail_, Registered_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserEmailAndActivated Get Signupses via UserEmailAndActivated
func GetSignupsesByUserEmailAndActivated(offset int, limit int, UserEmail_ string, Activated_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_email = ? and activated = ?", UserEmail_, Activated_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserEmailAndActive Get Signupses via UserEmailAndActive
func GetSignupsesByUserEmailAndActive(offset int, limit int, UserEmail_ string, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_email = ? and active = ?", UserEmail_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserEmailAndActivationKey Get Signupses via UserEmailAndActivationKey
func GetSignupsesByUserEmailAndActivationKey(offset int, limit int, UserEmail_ string, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_email = ? and activation_key = ?", UserEmail_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByUserEmailAndMeta Get Signupses via UserEmailAndMeta
func GetSignupsesByUserEmailAndMeta(offset int, limit int, UserEmail_ string, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_email = ? and meta = ?", UserEmail_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByRegisteredAndActivated Get Signupses via RegisteredAndActivated
func GetSignupsesByRegisteredAndActivated(offset int, limit int, Registered_ time.Time, Activated_ time.Time) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("registered = ? and activated = ?", Registered_, Activated_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByRegisteredAndActive Get Signupses via RegisteredAndActive
func GetSignupsesByRegisteredAndActive(offset int, limit int, Registered_ time.Time, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("registered = ? and active = ?", Registered_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByRegisteredAndActivationKey Get Signupses via RegisteredAndActivationKey
func GetSignupsesByRegisteredAndActivationKey(offset int, limit int, Registered_ time.Time, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("registered = ? and activation_key = ?", Registered_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByRegisteredAndMeta Get Signupses via RegisteredAndMeta
func GetSignupsesByRegisteredAndMeta(offset int, limit int, Registered_ time.Time, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("registered = ? and meta = ?", Registered_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByActivatedAndActive Get Signupses via ActivatedAndActive
func GetSignupsesByActivatedAndActive(offset int, limit int, Activated_ time.Time, Active_ int) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("activated = ? and active = ?", Activated_, Active_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByActivatedAndActivationKey Get Signupses via ActivatedAndActivationKey
func GetSignupsesByActivatedAndActivationKey(offset int, limit int, Activated_ time.Time, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("activated = ? and activation_key = ?", Activated_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByActivatedAndMeta Get Signupses via ActivatedAndMeta
func GetSignupsesByActivatedAndMeta(offset int, limit int, Activated_ time.Time, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("activated = ? and meta = ?", Activated_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByActiveAndActivationKey Get Signupses via ActiveAndActivationKey
func GetSignupsesByActiveAndActivationKey(offset int, limit int, Active_ int, ActivationKey_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("active = ? and activation_key = ?", Active_, ActivationKey_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByActiveAndMeta Get Signupses via ActiveAndMeta
func GetSignupsesByActiveAndMeta(offset int, limit int, Active_ int, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("active = ? and meta = ?", Active_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupsesByActivationKeyAndMeta Get Signupses via ActivationKeyAndMeta
func GetSignupsesByActivationKeyAndMeta(offset int, limit int, ActivationKey_ string, Meta_ string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("activation_key = ? and meta = ?", ActivationKey_, Meta_).Limit(limit, offset).Find(_Signups)
	return _Signups, err
}

// GetSignupses Get Signupses via field
func GetSignupses(offset int, limit int, field string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Limit(limit, offset).Desc(field).Find(_Signups)
	return _Signups, err
}

// GetSignupsesViaSignupId Get Signupss via SignupId
func GetSignupsesViaSignupId(offset int, limit int, SignupId_ int64, field string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("signup_id = ?", SignupId_).Limit(limit, offset).Desc(field).Find(_Signups)
	return _Signups, err
}

// GetSignupsesViaDomain Get Signupss via Domain
func GetSignupsesViaDomain(offset int, limit int, Domain_ string, field string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("domain = ?", Domain_).Limit(limit, offset).Desc(field).Find(_Signups)
	return _Signups, err
}

// GetSignupsesViaPath Get Signupss via Path
func GetSignupsesViaPath(offset int, limit int, Path_ string, field string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("path = ?", Path_).Limit(limit, offset).Desc(field).Find(_Signups)
	return _Signups, err
}

// GetSignupsesViaTitle Get Signupss via Title
func GetSignupsesViaTitle(offset int, limit int, Title_ string, field string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("title = ?", Title_).Limit(limit, offset).Desc(field).Find(_Signups)
	return _Signups, err
}

// GetSignupsesViaUserLogin Get Signupss via UserLogin
func GetSignupsesViaUserLogin(offset int, limit int, UserLogin_ string, field string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_login = ?", UserLogin_).Limit(limit, offset).Desc(field).Find(_Signups)
	return _Signups, err
}

// GetSignupsesViaUserEmail Get Signupss via UserEmail
func GetSignupsesViaUserEmail(offset int, limit int, UserEmail_ string, field string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("user_email = ?", UserEmail_).Limit(limit, offset).Desc(field).Find(_Signups)
	return _Signups, err
}

// GetSignupsesViaRegistered Get Signupss via Registered
func GetSignupsesViaRegistered(offset int, limit int, Registered_ time.Time, field string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("registered = ?", Registered_).Limit(limit, offset).Desc(field).Find(_Signups)
	return _Signups, err
}

// GetSignupsesViaActivated Get Signupss via Activated
func GetSignupsesViaActivated(offset int, limit int, Activated_ time.Time, field string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("activated = ?", Activated_).Limit(limit, offset).Desc(field).Find(_Signups)
	return _Signups, err
}

// GetSignupsesViaActive Get Signupss via Active
func GetSignupsesViaActive(offset int, limit int, Active_ int, field string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("active = ?", Active_).Limit(limit, offset).Desc(field).Find(_Signups)
	return _Signups, err
}

// GetSignupsesViaActivationKey Get Signupss via ActivationKey
func GetSignupsesViaActivationKey(offset int, limit int, ActivationKey_ string, field string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("activation_key = ?", ActivationKey_).Limit(limit, offset).Desc(field).Find(_Signups)
	return _Signups, err
}

// GetSignupsesViaMeta Get Signupss via Meta
func GetSignupsesViaMeta(offset int, limit int, Meta_ string, field string) (*[]*Signups, error) {
	var _Signups = new([]*Signups)
	err := Engine.Table("signups").Where("meta = ?", Meta_).Limit(limit, offset).Desc(field).Find(_Signups)
	return _Signups, err
}

// HasSignupsViaSignupId Has Signups via SignupId
func HasSignupsViaSignupId(iSignupId int64) bool {
	if has, err := Engine.Where("signup_id = ?", iSignupId).Get(new(Signups)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasSignupsViaDomain Has Signups via Domain
func HasSignupsViaDomain(iDomain string) bool {
	if has, err := Engine.Where("domain = ?", iDomain).Get(new(Signups)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasSignupsViaPath Has Signups via Path
func HasSignupsViaPath(iPath string) bool {
	if has, err := Engine.Where("path = ?", iPath).Get(new(Signups)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasSignupsViaTitle Has Signups via Title
func HasSignupsViaTitle(iTitle string) bool {
	if has, err := Engine.Where("title = ?", iTitle).Get(new(Signups)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasSignupsViaUserLogin Has Signups via UserLogin
func HasSignupsViaUserLogin(iUserLogin string) bool {
	if has, err := Engine.Where("user_login = ?", iUserLogin).Get(new(Signups)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasSignupsViaUserEmail Has Signups via UserEmail
func HasSignupsViaUserEmail(iUserEmail string) bool {
	if has, err := Engine.Where("user_email = ?", iUserEmail).Get(new(Signups)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasSignupsViaRegistered Has Signups via Registered
func HasSignupsViaRegistered(iRegistered time.Time) bool {
	if has, err := Engine.Where("registered = ?", iRegistered).Get(new(Signups)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasSignupsViaActivated Has Signups via Activated
func HasSignupsViaActivated(iActivated time.Time) bool {
	if has, err := Engine.Where("activated = ?", iActivated).Get(new(Signups)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasSignupsViaActive Has Signups via Active
func HasSignupsViaActive(iActive int) bool {
	if has, err := Engine.Where("active = ?", iActive).Get(new(Signups)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasSignupsViaActivationKey Has Signups via ActivationKey
func HasSignupsViaActivationKey(iActivationKey string) bool {
	if has, err := Engine.Where("activation_key = ?", iActivationKey).Get(new(Signups)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasSignupsViaMeta Has Signups via Meta
func HasSignupsViaMeta(iMeta string) bool {
	if has, err := Engine.Where("meta = ?", iMeta).Get(new(Signups)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetSignupsViaSignupId Get Signups via SignupId
func GetSignupsViaSignupId(iSignupId int64) (*Signups, error) {
	var _Signups = &Signups{SignupId: iSignupId}
	has, err := Engine.Get(_Signups)
	if has {
		return _Signups, err
	} else {
		return nil, err
	}
}

// GetSignupsViaDomain Get Signups via Domain
func GetSignupsViaDomain(iDomain string) (*Signups, error) {
	var _Signups = &Signups{Domain: iDomain}
	has, err := Engine.Get(_Signups)
	if has {
		return _Signups, err
	} else {
		return nil, err
	}
}

// GetSignupsViaPath Get Signups via Path
func GetSignupsViaPath(iPath string) (*Signups, error) {
	var _Signups = &Signups{Path: iPath}
	has, err := Engine.Get(_Signups)
	if has {
		return _Signups, err
	} else {
		return nil, err
	}
}

// GetSignupsViaTitle Get Signups via Title
func GetSignupsViaTitle(iTitle string) (*Signups, error) {
	var _Signups = &Signups{Title: iTitle}
	has, err := Engine.Get(_Signups)
	if has {
		return _Signups, err
	} else {
		return nil, err
	}
}

// GetSignupsViaUserLogin Get Signups via UserLogin
func GetSignupsViaUserLogin(iUserLogin string) (*Signups, error) {
	var _Signups = &Signups{UserLogin: iUserLogin}
	has, err := Engine.Get(_Signups)
	if has {
		return _Signups, err
	} else {
		return nil, err
	}
}

// GetSignupsViaUserEmail Get Signups via UserEmail
func GetSignupsViaUserEmail(iUserEmail string) (*Signups, error) {
	var _Signups = &Signups{UserEmail: iUserEmail}
	has, err := Engine.Get(_Signups)
	if has {
		return _Signups, err
	} else {
		return nil, err
	}
}

// GetSignupsViaRegistered Get Signups via Registered
func GetSignupsViaRegistered(iRegistered time.Time) (*Signups, error) {
	var _Signups = &Signups{Registered: iRegistered}
	has, err := Engine.Get(_Signups)
	if has {
		return _Signups, err
	} else {
		return nil, err
	}
}

// GetSignupsViaActivated Get Signups via Activated
func GetSignupsViaActivated(iActivated time.Time) (*Signups, error) {
	var _Signups = &Signups{Activated: iActivated}
	has, err := Engine.Get(_Signups)
	if has {
		return _Signups, err
	} else {
		return nil, err
	}
}

// GetSignupsViaActive Get Signups via Active
func GetSignupsViaActive(iActive int) (*Signups, error) {
	var _Signups = &Signups{Active: iActive}
	has, err := Engine.Get(_Signups)
	if has {
		return _Signups, err
	} else {
		return nil, err
	}
}

// GetSignupsViaActivationKey Get Signups via ActivationKey
func GetSignupsViaActivationKey(iActivationKey string) (*Signups, error) {
	var _Signups = &Signups{ActivationKey: iActivationKey}
	has, err := Engine.Get(_Signups)
	if has {
		return _Signups, err
	} else {
		return nil, err
	}
}

// GetSignupsViaMeta Get Signups via Meta
func GetSignupsViaMeta(iMeta string) (*Signups, error) {
	var _Signups = &Signups{Meta: iMeta}
	has, err := Engine.Get(_Signups)
	if has {
		return _Signups, err
	} else {
		return nil, err
	}
}

// SetSignupsViaSignupId Set Signups via SignupId
func SetSignupsViaSignupId(iSignupId int64, signups *Signups) (int64, error) {
	signups.SignupId = iSignupId
	return Engine.Insert(signups)
}

// SetSignupsViaDomain Set Signups via Domain
func SetSignupsViaDomain(iDomain string, signups *Signups) (int64, error) {
	signups.Domain = iDomain
	return Engine.Insert(signups)
}

// SetSignupsViaPath Set Signups via Path
func SetSignupsViaPath(iPath string, signups *Signups) (int64, error) {
	signups.Path = iPath
	return Engine.Insert(signups)
}

// SetSignupsViaTitle Set Signups via Title
func SetSignupsViaTitle(iTitle string, signups *Signups) (int64, error) {
	signups.Title = iTitle
	return Engine.Insert(signups)
}

// SetSignupsViaUserLogin Set Signups via UserLogin
func SetSignupsViaUserLogin(iUserLogin string, signups *Signups) (int64, error) {
	signups.UserLogin = iUserLogin
	return Engine.Insert(signups)
}

// SetSignupsViaUserEmail Set Signups via UserEmail
func SetSignupsViaUserEmail(iUserEmail string, signups *Signups) (int64, error) {
	signups.UserEmail = iUserEmail
	return Engine.Insert(signups)
}

// SetSignupsViaRegistered Set Signups via Registered
func SetSignupsViaRegistered(iRegistered time.Time, signups *Signups) (int64, error) {
	signups.Registered = iRegistered
	return Engine.Insert(signups)
}

// SetSignupsViaActivated Set Signups via Activated
func SetSignupsViaActivated(iActivated time.Time, signups *Signups) (int64, error) {
	signups.Activated = iActivated
	return Engine.Insert(signups)
}

// SetSignupsViaActive Set Signups via Active
func SetSignupsViaActive(iActive int, signups *Signups) (int64, error) {
	signups.Active = iActive
	return Engine.Insert(signups)
}

// SetSignupsViaActivationKey Set Signups via ActivationKey
func SetSignupsViaActivationKey(iActivationKey string, signups *Signups) (int64, error) {
	signups.ActivationKey = iActivationKey
	return Engine.Insert(signups)
}

// SetSignupsViaMeta Set Signups via Meta
func SetSignupsViaMeta(iMeta string, signups *Signups) (int64, error) {
	signups.Meta = iMeta
	return Engine.Insert(signups)
}

// AddSignups Add Signups via all columns
func AddSignups(iSignupId int64, iDomain string, iPath string, iTitle string, iUserLogin string, iUserEmail string, iRegistered time.Time, iActivated time.Time, iActive int, iActivationKey string, iMeta string) error {
	_Signups := &Signups{SignupId: iSignupId, Domain: iDomain, Path: iPath, Title: iTitle, UserLogin: iUserLogin, UserEmail: iUserEmail, Registered: iRegistered, Activated: iActivated, Active: iActive, ActivationKey: iActivationKey, Meta: iMeta}
	if _, err := Engine.Insert(_Signups); err != nil {
		return err
	} else {
		return nil
	}
}

// PostSignups Post Signups via iSignups
func PostSignups(iSignups *Signups) (int64, error) {
	_, err := Engine.Insert(iSignups)
	return iSignups.SignupId, err
}

// PutSignups Put Signups
func PutSignups(iSignups *Signups) (int64, error) {
	_, err := Engine.Id(iSignups.SignupId).Update(iSignups)
	return iSignups.SignupId, err
}

// PutSignupsViaSignupId Put Signups via SignupId
func PutSignupsViaSignupId(SignupId_ int64, iSignups *Signups) (int64, error) {
	row, err := Engine.Update(iSignups, &Signups{SignupId: SignupId_})
	return row, err
}

// PutSignupsViaDomain Put Signups via Domain
func PutSignupsViaDomain(Domain_ string, iSignups *Signups) (int64, error) {
	row, err := Engine.Update(iSignups, &Signups{Domain: Domain_})
	return row, err
}

// PutSignupsViaPath Put Signups via Path
func PutSignupsViaPath(Path_ string, iSignups *Signups) (int64, error) {
	row, err := Engine.Update(iSignups, &Signups{Path: Path_})
	return row, err
}

// PutSignupsViaTitle Put Signups via Title
func PutSignupsViaTitle(Title_ string, iSignups *Signups) (int64, error) {
	row, err := Engine.Update(iSignups, &Signups{Title: Title_})
	return row, err
}

// PutSignupsViaUserLogin Put Signups via UserLogin
func PutSignupsViaUserLogin(UserLogin_ string, iSignups *Signups) (int64, error) {
	row, err := Engine.Update(iSignups, &Signups{UserLogin: UserLogin_})
	return row, err
}

// PutSignupsViaUserEmail Put Signups via UserEmail
func PutSignupsViaUserEmail(UserEmail_ string, iSignups *Signups) (int64, error) {
	row, err := Engine.Update(iSignups, &Signups{UserEmail: UserEmail_})
	return row, err
}

// PutSignupsViaRegistered Put Signups via Registered
func PutSignupsViaRegistered(Registered_ time.Time, iSignups *Signups) (int64, error) {
	row, err := Engine.Update(iSignups, &Signups{Registered: Registered_})
	return row, err
}

// PutSignupsViaActivated Put Signups via Activated
func PutSignupsViaActivated(Activated_ time.Time, iSignups *Signups) (int64, error) {
	row, err := Engine.Update(iSignups, &Signups{Activated: Activated_})
	return row, err
}

// PutSignupsViaActive Put Signups via Active
func PutSignupsViaActive(Active_ int, iSignups *Signups) (int64, error) {
	row, err := Engine.Update(iSignups, &Signups{Active: Active_})
	return row, err
}

// PutSignupsViaActivationKey Put Signups via ActivationKey
func PutSignupsViaActivationKey(ActivationKey_ string, iSignups *Signups) (int64, error) {
	row, err := Engine.Update(iSignups, &Signups{ActivationKey: ActivationKey_})
	return row, err
}

// PutSignupsViaMeta Put Signups via Meta
func PutSignupsViaMeta(Meta_ string, iSignups *Signups) (int64, error) {
	row, err := Engine.Update(iSignups, &Signups{Meta: Meta_})
	return row, err
}

// UpdateSignupsViaSignupId via map[string]interface{}{}
func UpdateSignupsViaSignupId(iSignupId int64, iSignupsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Signups)).Where("signup_id = ?", iSignupId).Update(iSignupsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateSignupsViaDomain via map[string]interface{}{}
func UpdateSignupsViaDomain(iDomain string, iSignupsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Signups)).Where("domain = ?", iDomain).Update(iSignupsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateSignupsViaPath via map[string]interface{}{}
func UpdateSignupsViaPath(iPath string, iSignupsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Signups)).Where("path = ?", iPath).Update(iSignupsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateSignupsViaTitle via map[string]interface{}{}
func UpdateSignupsViaTitle(iTitle string, iSignupsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Signups)).Where("title = ?", iTitle).Update(iSignupsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateSignupsViaUserLogin via map[string]interface{}{}
func UpdateSignupsViaUserLogin(iUserLogin string, iSignupsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Signups)).Where("user_login = ?", iUserLogin).Update(iSignupsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateSignupsViaUserEmail via map[string]interface{}{}
func UpdateSignupsViaUserEmail(iUserEmail string, iSignupsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Signups)).Where("user_email = ?", iUserEmail).Update(iSignupsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateSignupsViaRegistered via map[string]interface{}{}
func UpdateSignupsViaRegistered(iRegistered time.Time, iSignupsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Signups)).Where("registered = ?", iRegistered).Update(iSignupsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateSignupsViaActivated via map[string]interface{}{}
func UpdateSignupsViaActivated(iActivated time.Time, iSignupsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Signups)).Where("activated = ?", iActivated).Update(iSignupsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateSignupsViaActive via map[string]interface{}{}
func UpdateSignupsViaActive(iActive int, iSignupsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Signups)).Where("active = ?", iActive).Update(iSignupsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateSignupsViaActivationKey via map[string]interface{}{}
func UpdateSignupsViaActivationKey(iActivationKey string, iSignupsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Signups)).Where("activation_key = ?", iActivationKey).Update(iSignupsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateSignupsViaMeta via map[string]interface{}{}
func UpdateSignupsViaMeta(iMeta string, iSignupsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Signups)).Where("meta = ?", iMeta).Update(iSignupsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteSignups Delete Signups
func DeleteSignups(iSignupId int64) (int64, error) {
	row, err := Engine.Id(iSignupId).Delete(new(Signups))
	return row, err
}

// DeleteSignupsViaSignupId Delete Signups via SignupId
func DeleteSignupsViaSignupId(iSignupId int64) (err error) {
	var has bool
	var _Signups = &Signups{SignupId: iSignupId}
	if has, err = Engine.Get(_Signups); (has == true) && (err == nil) {
		if row, err := Engine.Where("signup_id = ?", iSignupId).Delete(new(Signups)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteSignupsViaDomain Delete Signups via Domain
func DeleteSignupsViaDomain(iDomain string) (err error) {
	var has bool
	var _Signups = &Signups{Domain: iDomain}
	if has, err = Engine.Get(_Signups); (has == true) && (err == nil) {
		if row, err := Engine.Where("domain = ?", iDomain).Delete(new(Signups)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteSignupsViaPath Delete Signups via Path
func DeleteSignupsViaPath(iPath string) (err error) {
	var has bool
	var _Signups = &Signups{Path: iPath}
	if has, err = Engine.Get(_Signups); (has == true) && (err == nil) {
		if row, err := Engine.Where("path = ?", iPath).Delete(new(Signups)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteSignupsViaTitle Delete Signups via Title
func DeleteSignupsViaTitle(iTitle string) (err error) {
	var has bool
	var _Signups = &Signups{Title: iTitle}
	if has, err = Engine.Get(_Signups); (has == true) && (err == nil) {
		if row, err := Engine.Where("title = ?", iTitle).Delete(new(Signups)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteSignupsViaUserLogin Delete Signups via UserLogin
func DeleteSignupsViaUserLogin(iUserLogin string) (err error) {
	var has bool
	var _Signups = &Signups{UserLogin: iUserLogin}
	if has, err = Engine.Get(_Signups); (has == true) && (err == nil) {
		if row, err := Engine.Where("user_login = ?", iUserLogin).Delete(new(Signups)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteSignupsViaUserEmail Delete Signups via UserEmail
func DeleteSignupsViaUserEmail(iUserEmail string) (err error) {
	var has bool
	var _Signups = &Signups{UserEmail: iUserEmail}
	if has, err = Engine.Get(_Signups); (has == true) && (err == nil) {
		if row, err := Engine.Where("user_email = ?", iUserEmail).Delete(new(Signups)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteSignupsViaRegistered Delete Signups via Registered
func DeleteSignupsViaRegistered(iRegistered time.Time) (err error) {
	var has bool
	var _Signups = &Signups{Registered: iRegistered}
	if has, err = Engine.Get(_Signups); (has == true) && (err == nil) {
		if row, err := Engine.Where("registered = ?", iRegistered).Delete(new(Signups)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteSignupsViaActivated Delete Signups via Activated
func DeleteSignupsViaActivated(iActivated time.Time) (err error) {
	var has bool
	var _Signups = &Signups{Activated: iActivated}
	if has, err = Engine.Get(_Signups); (has == true) && (err == nil) {
		if row, err := Engine.Where("activated = ?", iActivated).Delete(new(Signups)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteSignupsViaActive Delete Signups via Active
func DeleteSignupsViaActive(iActive int) (err error) {
	var has bool
	var _Signups = &Signups{Active: iActive}
	if has, err = Engine.Get(_Signups); (has == true) && (err == nil) {
		if row, err := Engine.Where("active = ?", iActive).Delete(new(Signups)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteSignupsViaActivationKey Delete Signups via ActivationKey
func DeleteSignupsViaActivationKey(iActivationKey string) (err error) {
	var has bool
	var _Signups = &Signups{ActivationKey: iActivationKey}
	if has, err = Engine.Get(_Signups); (has == true) && (err == nil) {
		if row, err := Engine.Where("activation_key = ?", iActivationKey).Delete(new(Signups)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteSignupsViaMeta Delete Signups via Meta
func DeleteSignupsViaMeta(iMeta string) (err error) {
	var has bool
	var _Signups = &Signups{Meta: iMeta}
	if has, err = Engine.Get(_Signups); (has == true) && (err == nil) {
		if row, err := Engine.Where("meta = ?", iMeta).Delete(new(Signups)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
