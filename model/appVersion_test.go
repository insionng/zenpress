package model_test

import (
	"testing"

	"github.com/insionng/zenpress/model"

	"github.com/stretchr/testify/assert"
)

func TestNewAppVersion(t *testing.T) {
	assert := assert.New(t)
	var appVersion = new(model.AppVersion)
	appVersion.AppID = 1
	appVersion.DbVersion = "3.69"
	db := model.NewAppVersion(appVersion)

	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

	db = model.DeleteAppVersion(appVersion.AppID)
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

}
func TestAddAppVersion(t *testing.T) {
	assert := assert.New(t)
	db := model.AddAppVersion(1, "3.69")

	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

}
func TestGetAppVersion(t *testing.T) {
	assert := assert.New(t)
	var appVersion = new(model.AppVersion)
	appVersion.AppID = 1
	appVersion.DbVersion = "3.69"
	db := model.NewAppVersion(appVersion)

	db, opt := model.GetAppVersion(appVersion.AppID)
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
		assert.Equal("3.69", opt.DbVersion, "they should be equal")
	}

}

func TestUpdateAppVersion(t *testing.T) {
	assert := assert.New(t)
	var appVersion = new(model.AppVersion)
	appVersion.AppID = 1
	appVersion.DbVersion = "3.69"
	db := model.NewAppVersion(appVersion)

	db, opt := model.UpdateAppVersion(appVersion.AppID, "3.99")
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
		assert.Equal("3.99", opt.DbVersion, "they should be equal")
	}

}

func TestDeleteAppVersion(t *testing.T) {
	assert := assert.New(t)
	var appVersion = new(model.AppVersion)
	appVersion.AppID = 1
	appVersion.DbVersion = "3.69"
	db := model.NewAppVersion(appVersion)

	db = model.DeleteAppVersion(appVersion.AppID)
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

}
