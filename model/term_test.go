package model_test

import (
	"testing"

	"github.com/insionng/zenpress/model"

	"github.com/stretchr/testify/assert"
)

func TestNewTerm(t *testing.T) {
	assert := assert.New(t)
	var term = new(model.Term)
	term.Name = "分类名称"
	term.Slug = "slug-name"
	db := model.NewTerm(term)

	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

	db = model.DeleteTerm(term.ID)
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

}
func TestAddTerm(t *testing.T) {
	assert := assert.New(t)
	db := model.AddTerm("testkey", "testslug", 1)

	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

}
func TestGetTerm(t *testing.T) {
	assert := assert.New(t)
	var term = new(model.Term)
	term.Name = "分类名称"
	term.Slug = "slug-name"
	db := model.NewTerm(term)

	db, opt := model.GetTerm(term.ID)
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
		assert.Equal("slug-name", opt.Slug, "they should be equal")
	}

}

func TestUpdateTerm(t *testing.T) {
	assert := assert.New(t)
	var term = new(model.Term)
	term.Name = "分类名称"
	term.Slug = "slug-name"
	db := model.NewTerm(term)

	db, opt := model.UpdateTerm(term.ID, "名称")
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
		assert.Equal("名称", opt.Name, "they should be equal")
	}

}

func TestDeleteTerm(t *testing.T) {
	assert := assert.New(t)
	var term = new(model.Term)
	term.Name = "分类名称"
	term.Slug = "slug-name"
	db := model.NewTerm(term)

	db = model.DeleteTerm(term.ID)
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

}
