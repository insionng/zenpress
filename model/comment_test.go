package model_test

import (
	"testing"

	"github.com/insionng/zenpress/model"

	"github.com/stretchr/testify/assert"
)

func TestNewComment(t *testing.T) {
	assert := assert.New(t)
	var comment = new(model.Comment)
	comment.CommentAuthor = "测试评论作者"
	comment.CommentContent = "测试评论内容"
	db := model.NewComment(comment)

	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

	db = model.DeleteComment(comment.ID)
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

}
func TestAddComment(t *testing.T) {
	assert := assert.New(t)
	db := model.AddComment("testkey", "http://www.at3.net", "http://at3.net/image.png")

	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

}
func TestGetComment(t *testing.T) {
	assert := assert.New(t)
	var comment = new(model.Comment)
	comment.CommentAuthor = "测试评论作者"
	comment.CommentContent = "http://www.at3.net"
	db := model.NewComment(comment)

	db, opt := model.GetComment(comment.ID)
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
		assert.Equal("http://www.at3.net", opt.CommentContent, "they should be equal")
	}

}

func TestUpdateComment(t *testing.T) {
	assert := assert.New(t)
	var comment = new(model.Comment)
	comment.CommentAuthor = "测试评论作者"
	comment.CommentContent = "http://www.at3.net"
	db := model.NewComment(comment)

	db, opt := model.UpdateComment(comment.ID, "http://at3.net")
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
		assert.Equal("http://at3.net", opt.CommentContent, "they should be equal")
	}

}

func TestDeleteComment(t *testing.T) {
	assert := assert.New(t)
	var comment = new(model.Comment)
	comment.CommentAuthor = "测试评论作者"
	comment.CommentContent = "http://www.at3.net"
	db := model.NewComment(comment)

	db = model.DeleteComment(comment.ID)
	if assert.NotNil(db) {
		assert.Equal(nil, db.Error, "they should be equal")
	}

}
