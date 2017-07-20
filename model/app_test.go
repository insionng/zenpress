package model_test

import (
	"testing"

	"github.com/insionng/zenpress/model"

	"github.com/stretchr/testify/assert"
)

func TestNewApp(t *testing.T) {
	assert := assert.New(t)
	var app = new(model.App)
	app.Domain = "http://at3.net"
	app.Path = "/zen"
	db := model.NewApp(app)

	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

	db = model.DeleteApp(app.ID)
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

}
func TestAddApp(t *testing.T) {
	assert := assert.New(t)
	db := model.AddApp("http://at3.net", "/zen")

	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

}
func TestGetApp(t *testing.T) {
	assert := assert.New(t)
	var app = new(model.App)
	app.Domain = "http://at3.net"
	app.Path = "/zen"
	db := model.NewApp(app)

	db, opt := model.GetApp(app.ID)
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
		assert.Equal("http://at3.net", opt.Domain, "they should be equal")
	}

}

func TestUpdateApp(t *testing.T) {
	assert := assert.New(t)
	var app = new(model.App)
	app.Domain = "http://at3.net"
	app.Path = "/zen"
	db := model.NewApp(app)

	db, opt := model.UpdateApp(app.ID, "http://www.at3.net", "/zen")
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
		assert.Equal("http://www.at3.net", opt.Domain, "they should be equal")
	}

}

func TestDeleteApp(t *testing.T) {
	assert := assert.New(t)
	var app = new(model.App)
	app.Domain = "http://at3.net"
	app.Path = "/zen"
	db := model.NewApp(app)

	db = model.DeleteApp(app.ID)
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

}
