package model_test

import (
	"testing"

	"github.com/insionng/zenpress/model"

	"github.com/stretchr/testify/assert"
)

func TestNewLink(t *testing.T) {
	assert := assert.New(t)
	var link = new(model.Link)
	link.LinkName = "测试链接"
	link.LinkURL = "http://www.at3.net"
	db := model.NewLink(link)

	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

	db = model.DeleteLink("测试链接")
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

}
func TestAddLink(t *testing.T) {
	assert := assert.New(t)
	db := model.AddLink("testkey", "http://www.at3.net", "http://at3.net/image.png")

	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

}
func TestGetLink(t *testing.T) {
	assert := assert.New(t)

	db, opt := model.GetLink("testkey")
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
		assert.Equal("http://www.at3.net", opt.LinkURL, "they should be equal")
	}

}

func TestUpdateLink(t *testing.T) {
	assert := assert.New(t)

	db, opt := model.UpdateLink("testkey", "http://at3.net")
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
		assert.Equal("http://at3.net", opt.LinkURL, "they should be equal")
	}

}

func TestDeleteLink(t *testing.T) {
	assert := assert.New(t)

	db := model.DeleteLink("testkey")
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

}
