package model

import (
	"fmt"
)

type Permission struct {
	ID               uint64 `gorm:"primary_key"`
	Pid              uint64
	URL              string
	Name             string
	Description      string
	Roles            []*Role
	ChildPermissions []*Permission `gorm:"-"`
}

func FindPermissionById(id int) Permission {
	var permission Permission
	Database.Where("id = ?", id).First(&permission)
	return permission
}

func FindPermissions() []*Permission {

	var permissions []*Permission
	Database.Find(&permissions)
	return permissions
}

func FindPermissionsByPid(pid int) []*Permission {
	var permissions []*Permission
	Database.Where("pid = ?", pid).Find(&permissions)
	return permissions
}

func SavePermission(permission *Permission) uint64 {
	Database.Create(permission)
	return permission.ID
}

func UpdatePermission(permission *Permission) uint64 {
	Database.Update(permission)
	return permission.ID
}

func DeletePermission(permission *Permission) {
	Enforcer.DeletePermission(fmt.Sprintf("%v", permission.ID))
	Database.Delete(permission)
}

func DeleteRolePermissionByPermissionId(permission_id int) {
	Enforcer.DeletePermission(fmt.Sprintf("%v", permission_id))
	Database.Exec("delete from role_permissions where permission_id = ?", permission_id)
}
