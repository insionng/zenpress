package model_test

import (
	"testing"

	"github.com/insionng/zenpress/model"

	"github.com/stretchr/testify/assert"
)

func TestNewUsermeta(t *testing.T) {
	assert := assert.New(t)
	var usermeta = new(model.Usermeta)
	usermeta.UserID = 1
	usermeta.MetaKey = "testMetakey"
	usermeta.MetaValue = "testMetaValue"
	db := model.NewUsermeta(usermeta)

	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

	db = model.DeleteUsermeta(usermeta.ID)
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

}
func TestAddUsermeta(t *testing.T) {
	assert := assert.New(t)
	db := model.AddUsermeta(1, "testMetaKey", "testMetaValue")
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

}
func TestGetUsermeta(t *testing.T) {
	assert := assert.New(t)
	var usermeta = new(model.Usermeta)
	usermeta.UserID = 1
	usermeta.MetaKey = "testMetakey"
	usermeta.MetaValue = "testMetaValue"
	db := model.NewUsermeta(usermeta)

	db, opt := model.GetUsermeta(usermeta.ID)
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
		assert.Equal("testMetaValue", opt.MetaValue, "they should be equal")
	}

}

func TestUpdateUsermeta(t *testing.T) {
	assert := assert.New(t)
	var usermeta = new(model.Usermeta)
	usermeta.UserID = 1
	usermeta.MetaKey = "testMetakey"
	usermeta.MetaValue = "testMetaValue"
	db := model.NewUsermeta(usermeta)

	db, opt := model.UpdateUsermeta(usermeta.ID, usermeta.ID, "testMetakeyUpdated", "testMetaValueUpdated")
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
		assert.Equal("testMetaValueUpdated", opt.MetaValue, "they should be equal")
	}

}

func TestDeleteUsermeta(t *testing.T) {
	assert := assert.New(t)
	var usermeta = new(model.Usermeta)
	usermeta.UserID = 1
	usermeta.MetaKey = "testMetakey"
	usermeta.MetaValue = "testMetaValue"
	db := model.NewUsermeta(usermeta)

	db = model.DeleteUsermeta(usermeta.ID)
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

}
