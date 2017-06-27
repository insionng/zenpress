package model

import (
	"fmt"
	"strconv"
)

type Role struct {
	ID          uint64 `gorm:"primary_key"`
	Name        string // `orm:"unique"`
	Users       []*User
	Permissions []*Permission `gorm:"many2many:role_permissions;"`
}

func FindRoleById(id int) Role {
	var role Role
	Database.Where("id = ?", id).First(&role)
	return role
}

func FindRoles() []*Role {
	var roles []*Role
	Database.Find(&roles)
	return roles
}

func SaveRole(role *Role) uint64 {
	Database.Create(role)
	return role.ID
}

func UpdateRole(role *Role) {
	Database.Update(role, "name")
}

func DeleteRole(role *Role) {
	Enforcer.DeleteRole(fmt.Sprintf("%v", role.ID))
	Database.Delete(role)
}

func DeleteRolePermissionByRoleId(role_id int) {
	Enforcer.DeletePermissionsForUser(strconv.Itoa(role_id))
	Database.Exec("delete from role_permissions where role_id = ?", role_id)
}

func SaveRolePermission(role_id int, permission_id int) {
	Enforcer.AddPermissionForUser(fmt.Sprintf("%v", role_id), fmt.Sprintf("%v", permission_id))
	Database.Exec("insert into role_permissions (role_id, permission_id) values (?, ?)", role_id, permission_id)
}

type Result struct {
	ID           uint64
	RoleID       uint64
	PermissionID uint64
}

func FindRolePermissionByRoleId(role_id int) []Result {
	var res []Result
	Database.Raw("select id, role_id, permission_id from role_permissions where role_id = ?", role_id).Scan(&res)
	return res
}
