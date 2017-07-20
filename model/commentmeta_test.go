package model_test

import (
	"testing"

	"github.com/insionng/zenpress/model"

	"github.com/stretchr/testify/assert"
)

func TestNewCommentmeta(t *testing.T) {
	assert := assert.New(t)
	var commentmeta = new(model.Commentmeta)
	commentmeta.CommentID = 1
	commentmeta.MetaKey = "testMetakey"
	commentmeta.MetaValue = "testMetaValue"
	db := model.NewCommentmeta(commentmeta)

	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

	db = model.DeleteCommentmeta(commentmeta.ID)
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

}
func TestAddCommentmeta(t *testing.T) {
	assert := assert.New(t)
	db := model.AddCommentmeta(1, "testMetaKey", "testMetaValue")
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

}
func TestGetCommentmeta(t *testing.T) {
	assert := assert.New(t)
	var commentmeta = new(model.Commentmeta)
	commentmeta.CommentID = 1
	commentmeta.MetaKey = "testMetakey"
	commentmeta.MetaValue = "testMetaValue"
	db := model.NewCommentmeta(commentmeta)

	db, opt := model.GetCommentmeta(commentmeta.ID)
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
		assert.Equal("testMetaValue", opt.MetaValue, "they should be equal")
	}

}

func TestUpdateCommentmeta(t *testing.T) {
	assert := assert.New(t)
	var commentmeta = new(model.Commentmeta)
	commentmeta.CommentID = 1
	commentmeta.MetaKey = "testMetakey"
	commentmeta.MetaValue = "testMetaValue"
	db := model.NewCommentmeta(commentmeta)

	db, opt := model.UpdateCommentmeta(commentmeta.ID, commentmeta.ID, "testMetakeyUpdated", "testMetaValueUpdated")
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
		assert.Equal("testMetaValueUpdated", opt.MetaValue, "they should be equal")
	}

}

func TestDeleteCommentmeta(t *testing.T) {
	assert := assert.New(t)
	var commentmeta = new(model.Commentmeta)
	commentmeta.CommentID = 1
	commentmeta.MetaKey = "testMetakey"
	commentmeta.MetaValue = "testMetaValue"
	db := model.NewCommentmeta(commentmeta)

	db = model.DeleteCommentmeta(commentmeta.ID)
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

}
