package model_test

import (
	"testing"
	"github.com/insionng/zenpress/model"

	"github.com/stretchr/testify/assert"
)

func TestAddOption(t *testing.T) {
	assert := assert.New(t)
	db := model.AddOption("testkey", "testvalue", "yes")

	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

}
func TestGetOption(t *testing.T) {
	assert := assert.New(t)

	db, opt := model.GetOption("testkey")
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
		assert.Equal("testvalue", opt.OptionValue, "they should be equal")
	}

}

func TestUpdateOption(t *testing.T) {
	assert := assert.New(t)

	db, opt := model.UpdateOption("testkey", "testupdate")
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
		assert.Equal("testupdate", opt.OptionValue, "they should be equal")
	}

}

func TestDeleteOption(t *testing.T) {
	assert := assert.New(t)

	db := model.DeleteOption("testkey")
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

}
