package model_test

import (
	"testing"

	"github.com/insionng/zenpress/model"

	"github.com/stretchr/testify/assert"
)

func TestNewSite(t *testing.T) {
	assert := assert.New(t)
	var site = new(model.Site)
	site.Domain = "http://at3.net"
	site.Path = "/zen"
	db := model.NewSite(site)

	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

	db = model.DeleteSite(site.ID)
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

}
func TestAddSite(t *testing.T) {
	assert := assert.New(t)
	db := model.AddSite("http://at3.net", "/zen")

	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

}
func TestGetSite(t *testing.T) {
	assert := assert.New(t)
	var site = new(model.Site)
	site.Domain = "http://at3.net"
	site.Path = "/zen"
	db := model.NewSite(site)

	db, opt := model.GetSite(site.ID)
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
		assert.Equal("http://at3.net", opt.Domain, "they should be equal")
	}

}

func TestUpdateSite(t *testing.T) {
	assert := assert.New(t)
	var site = new(model.Site)
	site.Domain = "http://at3.net"
	site.Path = "/zen"
	db := model.NewSite(site)

	db, opt := model.UpdateSite(site.ID, "http://www.at3.net", "/zen")
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
		assert.Equal("http://www.at3.net", opt.Domain, "they should be equal")
	}

}

func TestDeleteSite(t *testing.T) {
	assert := assert.New(t)
	var site = new(model.Site)
	site.Domain = "http://at3.net"
	site.Path = "/zen"
	db := model.NewSite(site)

	db = model.DeleteSite(site.ID)
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

}
