package model_test

import (
	"testing"

	"github.com/insionng/zenpress/model"

	"github.com/stretchr/testify/assert"
)

func TestNewPostmeta(t *testing.T) {
	assert := assert.New(t)
	var postmeta = new(model.Postmeta)
	postmeta.PostID = 1
	postmeta.MetaKey = "testMetakey"
	postmeta.MetaValue = "testMetaValue"
	db := model.NewPostmeta(postmeta)

	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

	db = model.DeletePostmeta(postmeta.ID)
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

}
func TestAddPostmeta(t *testing.T) {
	assert := assert.New(t)
	db := model.AddPostmeta(1, "testMetaKey", "testMetaValue")
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

}
func TestGetPostmeta(t *testing.T) {
	assert := assert.New(t)
	var postmeta = new(model.Postmeta)
	postmeta.PostID = 1
	postmeta.MetaKey = "testMetakey"
	postmeta.MetaValue = "testMetaValue"
	db := model.NewPostmeta(postmeta)

	db, opt := model.GetPostmeta(postmeta.ID)
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
		assert.Equal("testMetaValue", opt.MetaValue, "they should be equal")
	}

}

func TestUpdatePostmeta(t *testing.T) {
	assert := assert.New(t)
	var postmeta = new(model.Postmeta)
	postmeta.PostID = 1
	postmeta.MetaKey = "testMetakey"
	postmeta.MetaValue = "testMetaValue"
	db := model.NewPostmeta(postmeta)

	db, opt := model.UpdatePostmeta(postmeta.ID, postmeta.ID, "testMetakeyUpdated", "testMetaValueUpdated")
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
		assert.Equal("testMetaValueUpdated", opt.MetaValue, "they should be equal")
	}

}

func TestDeletePostmeta(t *testing.T) {
	assert := assert.New(t)
	var postmeta = new(model.Postmeta)
	postmeta.PostID = 1
	postmeta.MetaKey = "testMetakey"
	postmeta.MetaValue = "testMetaValue"
	db := model.NewPostmeta(postmeta)

	db = model.DeletePostmeta(postmeta.ID)
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

}
