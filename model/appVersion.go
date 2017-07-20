package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// AppVersion 应用版本信息
type AppVersion struct {
	ID          uint64 `gorm:"primary_key"`
	AppID       uint64
	DbVersion   string    `gorm:"not null default '' index VARCHAR(20)"`
	LastUpdated time.Time `gorm:"not null default '0000-00-00 00:00:00' DATETIME"`
}

// NewAppVersion 创建应用版本
func NewAppVersion(appVersion *AppVersion) (db *gorm.DB) {
	db = Database.Create(appVersion)
	return
}

// AddAppVersion 新增应用版本
func AddAppVersion(appID uint64, dbVersion string) (db *gorm.DB) {
	db = Database.Create(&AppVersion{AppID: appID, DbVersion: dbVersion})
	return
}

// GetAppVersion 获得应用版本
func GetAppVersion(appID uint64) (db *gorm.DB, appVersion AppVersion) {
	db = Database.First(&appVersion, "app_id = ?", appID)
	return
}

// UpdateAppVersion 更新应用版本
func UpdateAppVersion(id uint64, dbVersion string) (db *gorm.DB, appVersion AppVersion) {
	appVersion = AppVersion{AppID: id, DbVersion: dbVersion}
	db = Database.Model(&appVersion).Update("id", id)
	return
}

// DeleteAppVersion 删除应用版本
func DeleteAppVersion(id uint64) (db *gorm.DB) {
	db = Database.Delete(AppVersion{}, "id = ?", id)
	return
}
