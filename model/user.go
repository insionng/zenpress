package model

import (
	"fmt"
	"time"

	"github.com/insionng/zenpress/helper"

	"github.com/casbin/casbin"
)

// User 用户基本信息表。存放系统所有用户基本信息。
type User struct {
	ID                uint64    `gorm:"primary_key"`
	Token             string    //`orm:"unique"`
	UserLogin         string    `gorm:"not null default '' index VARCHAR(60)"`
	UserPass          string    `gorm:"not null default '' VARCHAR(255)"`
	UserNicename      string    `gorm:"not null default '' index VARCHAR(50)"`
	UserEmail         string    `gorm:"not null default '' index VARCHAR(100)"`
	UserURL           string    `gorm:"not null default '' VARCHAR(100)"`
	UserRegistered    time.Time `gorm:"not null default '0000-00-00 00:00:00' DATETIME"`
	UserActivationKey string    `gorm:"not null default '' VARCHAR(255)"`
	UserStatus        int       `gorm:"not null default 0 INT(11)"`
	DisplayName       string    `gorm:"not null default '' VARCHAR(250)"`
	Spam              int       `gorm:"not null default 0 TINYINT(2)"`
	Deleted           int       `gorm:"not null default 0 TINYINT(2)"`
	Roles             []Role    `gorm:"many2many:user_roles;"`
}

var Enforcer *casbin.Enforcer = nil

func getAttr(name string, attr string) string {
	if attr != "url" {
		return ""
	}

	permissions := FindPermissions()
	for _, permission := range permissions {
		if name == fmt.Sprintf("%v", permission.ID) {
			return permission.URL
		}
	}
	return ""
}

func getAttrFunc(args ...interface{}) (interface{}, error) {
	name := args[0].(string)
	attr := args[1].(string)

	return (string)(getAttr(name, attr)), nil
}

func Init() {
	Enforcer = &casbin.Enforcer{}
	Enforcer.InitWithFile("content/config/rbac_model.conf", "")
	Enforcer.AddActionAttributeFunction(getAttrFunc)

	type UserRoleResult struct {
		UserID uint64
		RoleID uint64
	}
	var res []UserRoleResult
	Database.Raw("select user_id, role_id from user_roles").Scan(&res)
	for _, param := range res {
		Enforcer.AddRoleForUser(fmt.Sprintf("%v", param.UserID), fmt.Sprintf("%v", param.RoleID))
	}

	type RolePermissionResult struct {
		RoleID       uint64
		PermissionID uint64
	}
	var rez []RolePermissionResult
	Database.Raw("select role_id, permission_id from role_permissions").Scan(&rez)
	for _, param := range rez {
		Enforcer.AddPermissionForUser(fmt.Sprintf("%v", param.RoleID), fmt.Sprintf("%v", param.PermissionID))
	}
}

func FindUserById(id int) (bool, User) {

	var user User
	db := Database.Where("id = ?", id).First(&user)
	return db.Error == nil, user
}

func FindUserByToken(token string) (bool, User) {

	var user User
	db := Database.Where("token = ?", token).First(&user)

	return db.Error == nil, user
}

func Login(username string, password string) (bool, User) {
	var user User
	db := Database.Where("username = ? and password = ?", username, password).First(&user)
	return db.Error == nil, user
}

func FindUserByUserName(username string) (bool, User) {

	var user User
	db := Database.Where("username = ?", username).First(&user)
	return db.Error == nil, user
}

func SaveUser(user *User) uint64 {
	Database.Create(user)
	return user.ID
}

func UpdateUser(user *User) {
	Database.Update(user)
}

func PageUser(p int, size int) helper.Page {

	var user User
	var list []User
	qs := Database.Find(&user)
	var count int
	qs.Count(count)
	qs.Related(&user).Order("user_registered").Limit(size).Offset((p - 1) * size).Find(&list)
	//c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return helper.PageUtil(count, p, size, list)
}

func FindPermissionByUserIdAndPermissionName(userId int, name string) bool {
	permissions := FindPermissions()
	for _, permission := range permissions {
		if name == permission.Name {
			return Enforcer.Enforce(fmt.Sprintf("%v", userId), permission.URL)
		}
	}

	return false
}

func DeleteUser(user *User) {
	Enforcer.DeleteUser(fmt.Sprintf("%v", user.ID))
	Database.Delete(user)
}

func DeleteUserRolesByUserId(user_id int) {
	Enforcer.DeleteRolesForUser(fmt.Sprintf("%v", user_id))
	Database.Exec("delete from user_roles where user_id = ?", user_id)
}

func SaveUserRole(user_id int, role_id int) {
	Enforcer.AddRoleForUser(fmt.Sprintf("%v", user_id), fmt.Sprintf("%v", role_id))
	Database.Exec("insert into user_roles (user_id, role_id) values (?, ?)", user_id, role_id)
}

type UserRoleResult struct {
	ID     uint64
	UserID uint64
	RoleID uint64
}

func FindUserRolesByUserId(user_id int) []UserRoleResult {

	var res []UserRoleResult
	Database.Raw("select id, user_id, role_id from user_roles where user_id = ?", user_id).Scan(&res)
	return res
}
