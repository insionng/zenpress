package model_test

import (
	"testing"

	"github.com/insionng/zenpress/model"

	"github.com/stretchr/testify/assert"
)

func TestNewPost(t *testing.T) {
	assert := assert.New(t)
	var post = new(model.Post)
	post.PostTitle = "测试文章标题"
	post.PostContent = "测试文章内容"
	db := model.NewPost(post)

	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

	db = model.DeletePost(post.ID)
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

}
func TestAddPost(t *testing.T) {
	assert := assert.New(t)
	db := model.AddPost("testkey", "http://www.at3.net", "http://at3.net/image.png")

	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

}
func TestGetPost(t *testing.T) {
	assert := assert.New(t)
	var post = new(model.Post)
	post.PostTitle = "测试文章标题"
	post.PostContent = "http://www.at3.net"
	db := model.NewPost(post)

	db, opt := model.GetPost(post.ID)
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
		assert.Equal("http://www.at3.net", opt.PostContent, "they should be equal")
	}

}

func TestUpdatePost(t *testing.T) {
	assert := assert.New(t)
	var post = new(model.Post)
	post.PostTitle = "测试文章标题"
	post.PostContent = "http://www.at3.net"
	db := model.NewPost(post)

	db, opt := model.UpdatePost(post.ID, "http://at3.net")
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
		assert.Equal("http://at3.net", opt.PostContent, "they should be equal")
	}

}

func TestDeletePost(t *testing.T) {
	assert := assert.New(t)
	var post = new(model.Post)
	post.PostTitle = "测试文章标题"
	post.PostContent = "http://www.at3.net"
	db := model.NewPost(post)

	db = model.DeletePost(post.ID)
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

}
