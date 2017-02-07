package handler

import (
	"github.com/insionng/zenpress/helper"
	"github.com/insionng/zenpress/model"
	"github.com/insionng/macross"
)

func GetCommentsesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetCommentsesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["commentssCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetCommentsesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsCountByCommentIdHandler(self *macross.Context) error {
	CommentId_ := self.Args("comment_ID").MustInt64()
	_int64 := model.GetCommentsCountByCommentId(CommentId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["commentsCount"] = 0
	}
	m["commentsCount"] = _int64
	return self.JSON(m)
}

func GetCommentsCountByCommentPostIdHandler(self *macross.Context) error {
	CommentPostId_ := self.Args("comment_post_ID").MustInt64()
	_int64 := model.GetCommentsCountByCommentPostId(CommentPostId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["commentsCount"] = 0
	}
	m["commentsCount"] = _int64
	return self.JSON(m)
}

func GetCommentsCountByCommentAuthorHandler(self *macross.Context) error {
	CommentAuthor_ := self.Args("comment_author").String()
	_int64 := model.GetCommentsCountByCommentAuthor(CommentAuthor_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["commentsCount"] = 0
	}
	m["commentsCount"] = _int64
	return self.JSON(m)
}

func GetCommentsCountByCommentAuthorEmailHandler(self *macross.Context) error {
	CommentAuthorEmail_ := self.Args("comment_author_email").String()
	_int64 := model.GetCommentsCountByCommentAuthorEmail(CommentAuthorEmail_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["commentsCount"] = 0
	}
	m["commentsCount"] = _int64
	return self.JSON(m)
}

func GetCommentsCountByCommentAuthorUrlHandler(self *macross.Context) error {
	CommentAuthorUrl_ := self.Args("comment_author_url").String()
	_int64 := model.GetCommentsCountByCommentAuthorUrl(CommentAuthorUrl_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["commentsCount"] = 0
	}
	m["commentsCount"] = _int64
	return self.JSON(m)
}

func GetCommentsCountByCommentAuthorIpHandler(self *macross.Context) error {
	CommentAuthorIp_ := self.Args("comment_author_IP").String()
	_int64 := model.GetCommentsCountByCommentAuthorIp(CommentAuthorIp_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["commentsCount"] = 0
	}
	m["commentsCount"] = _int64
	return self.JSON(m)
}

func GetCommentsCountByCommentDateHandler(self *macross.Context) error {
	CommentDate_ := self.Args("comment_date").Time()
	_int64 := model.GetCommentsCountByCommentDate(CommentDate_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["commentsCount"] = 0
	}
	m["commentsCount"] = _int64
	return self.JSON(m)
}

func GetCommentsCountByCommentDateGmtHandler(self *macross.Context) error {
	CommentDateGmt_ := self.Args("comment_date_gmt").Time()
	_int64 := model.GetCommentsCountByCommentDateGmt(CommentDateGmt_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["commentsCount"] = 0
	}
	m["commentsCount"] = _int64
	return self.JSON(m)
}

func GetCommentsCountByCommentContentHandler(self *macross.Context) error {
	CommentContent_ := self.Args("comment_content").String()
	_int64 := model.GetCommentsCountByCommentContent(CommentContent_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["commentsCount"] = 0
	}
	m["commentsCount"] = _int64
	return self.JSON(m)
}

func GetCommentsCountByCommentKarmaHandler(self *macross.Context) error {
	CommentKarma_ := self.Args("comment_karma").MustInt()
	_int64 := model.GetCommentsCountByCommentKarma(CommentKarma_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["commentsCount"] = 0
	}
	m["commentsCount"] = _int64
	return self.JSON(m)
}

func GetCommentsCountByCommentApprovedHandler(self *macross.Context) error {
	CommentApproved_ := self.Args("comment_approved").String()
	_int64 := model.GetCommentsCountByCommentApproved(CommentApproved_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["commentsCount"] = 0
	}
	m["commentsCount"] = _int64
	return self.JSON(m)
}

func GetCommentsCountByCommentAgentHandler(self *macross.Context) error {
	CommentAgent_ := self.Args("comment_agent").String()
	_int64 := model.GetCommentsCountByCommentAgent(CommentAgent_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["commentsCount"] = 0
	}
	m["commentsCount"] = _int64
	return self.JSON(m)
}

func GetCommentsCountByCommentTypeHandler(self *macross.Context) error {
	CommentType_ := self.Args("comment_type").String()
	_int64 := model.GetCommentsCountByCommentType(CommentType_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["commentsCount"] = 0
	}
	m["commentsCount"] = _int64
	return self.JSON(m)
}

func GetCommentsCountByCommentParentHandler(self *macross.Context) error {
	CommentParent_ := self.Args("comment_parent").MustInt64()
	_int64 := model.GetCommentsCountByCommentParent(CommentParent_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["commentsCount"] = 0
	}
	m["commentsCount"] = _int64
	return self.JSON(m)
}

func GetCommentsCountByUserIdHandler(self *macross.Context) error {
	UserId_ := self.Args("user_id").MustInt64()
	_int64 := model.GetCommentsCountByUserId(UserId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["commentsCount"] = 0
	}
	m["commentsCount"] = _int64
	return self.JSON(m)
}

func GetCommentsesByCommentIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCommentId := self.Args("comment_ID").MustInt64()
	if (offset > 0) && helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentId(offset, limit, iCommentId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCommentPostId := self.Args("comment_post_ID").MustInt64()
	if (offset > 0) && helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostId(offset, limit, iCommentPostId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCommentAuthor := self.Args("comment_author").String()
	if (offset > 0) && helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthor(offset, limit, iCommentAuthor, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthor's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	if (offset > 0) && helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmail(offset, limit, iCommentAuthorEmail, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	if (offset > 0) && helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrl(offset, limit, iCommentAuthorUrl, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCommentAuthorIp := self.Args("comment_author_IP").String()
	if (offset > 0) && helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIp(offset, limit, iCommentAuthorIp, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCommentDate := self.Args("comment_date").Time()
	if (offset > 0) && helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDate(offset, limit, iCommentDate, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	if (offset > 0) && helper.IsHas(iCommentDateGmt) {
		_Comments, _error := model.GetCommentsesByCommentDateGmt(offset, limit, iCommentDateGmt, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateGmt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCommentContent := self.Args("comment_content").String()
	if (offset > 0) && helper.IsHas(iCommentContent) {
		_Comments, _error := model.GetCommentsesByCommentContent(offset, limit, iCommentContent, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCommentKarma := self.Args("comment_karma").MustInt()
	if (offset > 0) && helper.IsHas(iCommentKarma) {
		_Comments, _error := model.GetCommentsesByCommentKarma(offset, limit, iCommentKarma, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCommentApproved := self.Args("comment_approved").String()
	if (offset > 0) && helper.IsHas(iCommentApproved) {
		_Comments, _error := model.GetCommentsesByCommentApproved(offset, limit, iCommentApproved, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCommentAgent := self.Args("comment_agent").String()
	if (offset > 0) && helper.IsHas(iCommentAgent) {
		_Comments, _error := model.GetCommentsesByCommentAgent(offset, limit, iCommentAgent, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCommentType := self.Args("comment_type").String()
	if (offset > 0) && helper.IsHas(iCommentType) {
		_Comments, _error := model.GetCommentsesByCommentType(offset, limit, iCommentType, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCommentParent := self.Args("comment_parent").MustInt64()
	if (offset > 0) && helper.IsHas(iCommentParent) {
		_Comments, _error := model.GetCommentsesByCommentParent(offset, limit, iCommentParent, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUserId := self.Args("user_id").MustInt64()
	if (offset > 0) && helper.IsHas(iUserId) {
		_Comments, _error := model.GetCommentsesByUserId(offset, limit, iUserId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentPostIdAndCommentAuthorHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthor := self.Args("comment_author").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentPostIdAndCommentAuthor(offset, limit, iCommentId,iCommentPostId,iCommentAuthor)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentPostIdAndCommentAuthor's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentPostIdAndCommentAuthorEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorEmail := self.Args("comment_author_email").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentPostIdAndCommentAuthorEmail(offset, limit, iCommentId,iCommentPostId,iCommentAuthorEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentPostIdAndCommentAuthorEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentPostIdAndCommentAuthorUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorUrl := self.Args("comment_author_url").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentPostIdAndCommentAuthorUrl(offset, limit, iCommentId,iCommentPostId,iCommentAuthorUrl)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentPostIdAndCommentAuthorUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentPostIdAndCommentAuthorIpHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorIp := self.Args("comment_author_ip").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentPostIdAndCommentAuthorIp(offset, limit, iCommentId,iCommentPostId,iCommentAuthorIp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentPostIdAndCommentAuthorIp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentPostIdAndCommentDateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentDate := self.Args("comment_date").Time()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentPostIdAndCommentDate(offset, limit, iCommentId,iCommentPostId,iCommentDate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentPostIdAndCommentDate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentPostIdAndCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentPostIdAndCommentDateGmt(offset, limit, iCommentId,iCommentPostId,iCommentDateGmt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentPostIdAndCommentDateGmt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentPostIdAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentPostIdAndCommentContent(offset, limit, iCommentId,iCommentPostId,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentPostIdAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentPostIdAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentPostIdAndCommentKarma(offset, limit, iCommentId,iCommentPostId,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentPostIdAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentPostIdAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentPostIdAndCommentApproved(offset, limit, iCommentId,iCommentPostId,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentPostIdAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentPostIdAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentPostIdAndCommentAgent(offset, limit, iCommentId,iCommentPostId,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentPostIdAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentPostIdAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentPostIdAndCommentType(offset, limit, iCommentId,iCommentPostId,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentPostIdAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentPostIdAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentPostIdAndCommentParent(offset, limit, iCommentId,iCommentPostId,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentPostIdAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentPostIdAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentPostIdAndUserId(offset, limit, iCommentId,iCommentPostId,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentPostIdAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorAndCommentAuthorEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorEmail := self.Args("comment_author_email").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorAndCommentAuthorEmail(offset, limit, iCommentId,iCommentAuthor,iCommentAuthorEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorAndCommentAuthorEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorAndCommentAuthorUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorUrl := self.Args("comment_author_url").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorAndCommentAuthorUrl(offset, limit, iCommentId,iCommentAuthor,iCommentAuthorUrl)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorAndCommentAuthorUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorAndCommentAuthorIpHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorAndCommentAuthorIp(offset, limit, iCommentId,iCommentAuthor,iCommentAuthorIp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorAndCommentAuthorIp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorAndCommentDateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentDate := self.Args("comment_date").Time()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorAndCommentDate(offset, limit, iCommentId,iCommentAuthor,iCommentDate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorAndCommentDate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorAndCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorAndCommentDateGmt(offset, limit, iCommentId,iCommentAuthor,iCommentDateGmt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorAndCommentDateGmt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorAndCommentContent(offset, limit, iCommentId,iCommentAuthor,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorAndCommentKarma(offset, limit, iCommentId,iCommentAuthor,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorAndCommentApproved(offset, limit, iCommentId,iCommentAuthor,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorAndCommentAgent(offset, limit, iCommentId,iCommentAuthor,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorAndCommentType(offset, limit, iCommentId,iCommentAuthor,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorAndCommentParent(offset, limit, iCommentId,iCommentAuthor,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthor := self.Args("comment_author").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorAndUserId(offset, limit, iCommentId,iCommentAuthor,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentAuthorUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAuthorUrl := self.Args("comment_author_url").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentAuthorUrl(offset, limit, iCommentId,iCommentAuthorEmail,iCommentAuthorUrl)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentAuthorUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentAuthorIpHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentAuthorIp(offset, limit, iCommentId,iCommentAuthorEmail,iCommentAuthorIp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentAuthorIp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentDateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentDate := self.Args("comment_date").Time()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentDate(offset, limit, iCommentId,iCommentAuthorEmail,iCommentDate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentDate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentDateGmt(offset, limit, iCommentId,iCommentAuthorEmail,iCommentDateGmt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentDateGmt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentContent(offset, limit, iCommentId,iCommentAuthorEmail,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentKarma(offset, limit, iCommentId,iCommentAuthorEmail,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentApproved(offset, limit, iCommentId,iCommentAuthorEmail,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentAgent(offset, limit, iCommentId,iCommentAuthorEmail,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentType(offset, limit, iCommentId,iCommentAuthorEmail,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentParent(offset, limit, iCommentId,iCommentAuthorEmail,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorEmailAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorEmailAndUserId(offset, limit, iCommentId,iCommentAuthorEmail,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorEmailAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentAuthorIpHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentAuthorIp(offset, limit, iCommentId,iCommentAuthorUrl,iCommentAuthorIp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentAuthorIp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentDateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentDate := self.Args("comment_date").Time()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentDate(offset, limit, iCommentId,iCommentAuthorUrl,iCommentDate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentDate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentDateGmt(offset, limit, iCommentId,iCommentAuthorUrl,iCommentDateGmt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentDateGmt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentContent(offset, limit, iCommentId,iCommentAuthorUrl,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentKarma(offset, limit, iCommentId,iCommentAuthorUrl,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentApproved(offset, limit, iCommentId,iCommentAuthorUrl,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentAgent(offset, limit, iCommentId,iCommentAuthorUrl,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentType(offset, limit, iCommentId,iCommentAuthorUrl,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentParent(offset, limit, iCommentId,iCommentAuthorUrl,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorUrlAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorUrlAndUserId(offset, limit, iCommentId,iCommentAuthorUrl,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorUrlAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorIpAndCommentDateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentDate := self.Args("comment_date").Time()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorIpAndCommentDate(offset, limit, iCommentId,iCommentAuthorIp,iCommentDate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorIpAndCommentDate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorIpAndCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorIpAndCommentDateGmt(offset, limit, iCommentId,iCommentAuthorIp,iCommentDateGmt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorIpAndCommentDateGmt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorIpAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorIpAndCommentContent(offset, limit, iCommentId,iCommentAuthorIp,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorIpAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorIpAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorIpAndCommentKarma(offset, limit, iCommentId,iCommentAuthorIp,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorIpAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorIpAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorIpAndCommentApproved(offset, limit, iCommentId,iCommentAuthorIp,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorIpAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorIpAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorIpAndCommentAgent(offset, limit, iCommentId,iCommentAuthorIp,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorIpAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorIpAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorIpAndCommentType(offset, limit, iCommentId,iCommentAuthorIp,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorIpAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorIpAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorIpAndCommentParent(offset, limit, iCommentId,iCommentAuthorIp,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorIpAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorIpAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorIpAndUserId(offset, limit, iCommentId,iCommentAuthorIp,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorIpAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentDateAndCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentDate := self.Args("comment_date").Time()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentDateAndCommentDateGmt(offset, limit, iCommentId,iCommentDate,iCommentDateGmt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentDateAndCommentDateGmt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentDateAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentDate := self.Args("comment_date").Time()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentDateAndCommentContent(offset, limit, iCommentId,iCommentDate,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentDateAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentDateAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentDate := self.Args("comment_date").Time()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentDateAndCommentKarma(offset, limit, iCommentId,iCommentDate,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentDateAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentDateAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentDate := self.Args("comment_date").Time()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentDateAndCommentApproved(offset, limit, iCommentId,iCommentDate,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentDateAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentDateAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentDate := self.Args("comment_date").Time()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentDateAndCommentAgent(offset, limit, iCommentId,iCommentDate,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentDateAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentDateAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentDate := self.Args("comment_date").Time()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentDateAndCommentType(offset, limit, iCommentId,iCommentDate,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentDateAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentDateAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentDate := self.Args("comment_date").Time()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentDateAndCommentParent(offset, limit, iCommentId,iCommentDate,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentDateAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentDateAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentDate := self.Args("comment_date").Time()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentDateAndUserId(offset, limit, iCommentId,iCommentDate,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentDateAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentDateGmtAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentDateGmtAndCommentContent(offset, limit, iCommentId,iCommentDateGmt,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentDateGmtAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentDateGmtAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentDateGmtAndCommentKarma(offset, limit, iCommentId,iCommentDateGmt,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentDateGmtAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentDateGmtAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentDateGmtAndCommentApproved(offset, limit, iCommentId,iCommentDateGmt,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentDateGmtAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentDateGmtAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentDateGmtAndCommentAgent(offset, limit, iCommentId,iCommentDateGmt,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentDateGmtAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentDateGmtAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentDateGmtAndCommentType(offset, limit, iCommentId,iCommentDateGmt,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentDateGmtAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentDateGmtAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentDateGmtAndCommentParent(offset, limit, iCommentId,iCommentDateGmt,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentDateGmtAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentDateGmtAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentDateGmtAndUserId(offset, limit, iCommentId,iCommentDateGmt,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentDateGmtAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentContentAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentContent := self.Args("comment_content").String()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentContentAndCommentKarma(offset, limit, iCommentId,iCommentContent,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentContentAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentContentAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentContent := self.Args("comment_content").String()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentContentAndCommentApproved(offset, limit, iCommentId,iCommentContent,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentContentAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentContentAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentContent := self.Args("comment_content").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentContentAndCommentAgent(offset, limit, iCommentId,iCommentContent,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentContentAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentContentAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentContent := self.Args("comment_content").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentContentAndCommentType(offset, limit, iCommentId,iCommentContent,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentContentAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentContentAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentContent := self.Args("comment_content").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentContentAndCommentParent(offset, limit, iCommentId,iCommentContent,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentContentAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentContentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentContent := self.Args("comment_content").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentContentAndUserId(offset, limit, iCommentId,iCommentContent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentContentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentKarmaAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentKarmaAndCommentApproved(offset, limit, iCommentId,iCommentKarma,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentKarmaAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentKarmaAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentKarmaAndCommentAgent(offset, limit, iCommentId,iCommentKarma,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentKarmaAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentKarmaAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentKarmaAndCommentType(offset, limit, iCommentId,iCommentKarma,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentKarmaAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentKarmaAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentKarmaAndCommentParent(offset, limit, iCommentId,iCommentKarma,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentKarmaAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentKarmaAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentKarmaAndUserId(offset, limit, iCommentId,iCommentKarma,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentKarmaAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentApprovedAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentApprovedAndCommentAgent(offset, limit, iCommentId,iCommentApproved,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentApprovedAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentApprovedAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentApprovedAndCommentType(offset, limit, iCommentId,iCommentApproved,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentApprovedAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentApprovedAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentApprovedAndCommentParent(offset, limit, iCommentId,iCommentApproved,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentApprovedAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentApprovedAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentApproved := self.Args("comment_approved").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentApprovedAndUserId(offset, limit, iCommentId,iCommentApproved,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentApprovedAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAgentAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAgent := self.Args("comment_agent").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAgentAndCommentType(offset, limit, iCommentId,iCommentAgent,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAgentAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAgentAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAgent := self.Args("comment_agent").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAgentAndCommentParent(offset, limit, iCommentId,iCommentAgent,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAgentAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAgentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAgent := self.Args("comment_agent").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAgentAndUserId(offset, limit, iCommentId,iCommentAgent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAgentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentTypeAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentType := self.Args("comment_type").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentTypeAndCommentParent(offset, limit, iCommentId,iCommentType,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentTypeAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentTypeAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentType := self.Args("comment_type").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentTypeAndUserId(offset, limit, iCommentId,iCommentType,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentTypeAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentParentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentParent := self.Args("comment_parent").MustInt64()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentParentAndUserId(offset, limit, iCommentId,iCommentParent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentParentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorAndCommentAuthorEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorEmail := self.Args("comment_author_email").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorAndCommentAuthorEmail(offset, limit, iCommentPostId,iCommentAuthor,iCommentAuthorEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorAndCommentAuthorEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorAndCommentAuthorUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorUrl := self.Args("comment_author_url").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorAndCommentAuthorUrl(offset, limit, iCommentPostId,iCommentAuthor,iCommentAuthorUrl)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorAndCommentAuthorUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorAndCommentAuthorIpHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorAndCommentAuthorIp(offset, limit, iCommentPostId,iCommentAuthor,iCommentAuthorIp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorAndCommentAuthorIp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorAndCommentDateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentDate := self.Args("comment_date").Time()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorAndCommentDate(offset, limit, iCommentPostId,iCommentAuthor,iCommentDate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorAndCommentDate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorAndCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorAndCommentDateGmt(offset, limit, iCommentPostId,iCommentAuthor,iCommentDateGmt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorAndCommentDateGmt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorAndCommentContent(offset, limit, iCommentPostId,iCommentAuthor,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorAndCommentKarma(offset, limit, iCommentPostId,iCommentAuthor,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorAndCommentApproved(offset, limit, iCommentPostId,iCommentAuthor,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorAndCommentAgent(offset, limit, iCommentPostId,iCommentAuthor,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorAndCommentType(offset, limit, iCommentPostId,iCommentAuthor,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorAndCommentParent(offset, limit, iCommentPostId,iCommentAuthor,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthor := self.Args("comment_author").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorAndUserId(offset, limit, iCommentPostId,iCommentAuthor,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentAuthorUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAuthorUrl := self.Args("comment_author_url").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentAuthorUrl(offset, limit, iCommentPostId,iCommentAuthorEmail,iCommentAuthorUrl)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentAuthorUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentAuthorIpHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentAuthorIp(offset, limit, iCommentPostId,iCommentAuthorEmail,iCommentAuthorIp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentAuthorIp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentDateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentDate := self.Args("comment_date").Time()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentDate(offset, limit, iCommentPostId,iCommentAuthorEmail,iCommentDate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentDate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentDateGmt(offset, limit, iCommentPostId,iCommentAuthorEmail,iCommentDateGmt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentDateGmt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentContent(offset, limit, iCommentPostId,iCommentAuthorEmail,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentKarma(offset, limit, iCommentPostId,iCommentAuthorEmail,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentApproved(offset, limit, iCommentPostId,iCommentAuthorEmail,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentAgent(offset, limit, iCommentPostId,iCommentAuthorEmail,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentType(offset, limit, iCommentPostId,iCommentAuthorEmail,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentParent(offset, limit, iCommentPostId,iCommentAuthorEmail,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorEmailAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorEmailAndUserId(offset, limit, iCommentPostId,iCommentAuthorEmail,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorEmailAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentAuthorIpHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentAuthorIp(offset, limit, iCommentPostId,iCommentAuthorUrl,iCommentAuthorIp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentAuthorIp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentDateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentDate := self.Args("comment_date").Time()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentDate(offset, limit, iCommentPostId,iCommentAuthorUrl,iCommentDate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentDate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentDateGmt(offset, limit, iCommentPostId,iCommentAuthorUrl,iCommentDateGmt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentDateGmt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentContent(offset, limit, iCommentPostId,iCommentAuthorUrl,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentKarma(offset, limit, iCommentPostId,iCommentAuthorUrl,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentApproved(offset, limit, iCommentPostId,iCommentAuthorUrl,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentAgent(offset, limit, iCommentPostId,iCommentAuthorUrl,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentType(offset, limit, iCommentPostId,iCommentAuthorUrl,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentParent(offset, limit, iCommentPostId,iCommentAuthorUrl,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorUrlAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorUrlAndUserId(offset, limit, iCommentPostId,iCommentAuthorUrl,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorUrlAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentDateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentDate := self.Args("comment_date").Time()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentDate(offset, limit, iCommentPostId,iCommentAuthorIp,iCommentDate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentDate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentDateGmt(offset, limit, iCommentPostId,iCommentAuthorIp,iCommentDateGmt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentDateGmt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentContent(offset, limit, iCommentPostId,iCommentAuthorIp,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentKarma(offset, limit, iCommentPostId,iCommentAuthorIp,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentApproved(offset, limit, iCommentPostId,iCommentAuthorIp,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentAgent(offset, limit, iCommentPostId,iCommentAuthorIp,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentType(offset, limit, iCommentPostId,iCommentAuthorIp,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentParent(offset, limit, iCommentPostId,iCommentAuthorIp,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorIpAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorIpAndUserId(offset, limit, iCommentPostId,iCommentAuthorIp,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorIpAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentDateAndCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentDate := self.Args("comment_date").Time()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentDateAndCommentDateGmt(offset, limit, iCommentPostId,iCommentDate,iCommentDateGmt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentDateAndCommentDateGmt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentDateAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentDate := self.Args("comment_date").Time()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentDateAndCommentContent(offset, limit, iCommentPostId,iCommentDate,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentDateAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentDateAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentDate := self.Args("comment_date").Time()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentDateAndCommentKarma(offset, limit, iCommentPostId,iCommentDate,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentDateAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentDateAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentDate := self.Args("comment_date").Time()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentDateAndCommentApproved(offset, limit, iCommentPostId,iCommentDate,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentDateAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentDateAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentDate := self.Args("comment_date").Time()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentDateAndCommentAgent(offset, limit, iCommentPostId,iCommentDate,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentDateAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentDateAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentDate := self.Args("comment_date").Time()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentDateAndCommentType(offset, limit, iCommentPostId,iCommentDate,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentDateAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentDateAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentDate := self.Args("comment_date").Time()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentDateAndCommentParent(offset, limit, iCommentPostId,iCommentDate,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentDateAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentDateAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentDate := self.Args("comment_date").Time()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentDateAndUserId(offset, limit, iCommentPostId,iCommentDate,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentDateAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentDateGmtAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentDateGmtAndCommentContent(offset, limit, iCommentPostId,iCommentDateGmt,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentDateGmtAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentDateGmtAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentDateGmtAndCommentKarma(offset, limit, iCommentPostId,iCommentDateGmt,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentDateGmtAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentDateGmtAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentDateGmtAndCommentApproved(offset, limit, iCommentPostId,iCommentDateGmt,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentDateGmtAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentDateGmtAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentDateGmtAndCommentAgent(offset, limit, iCommentPostId,iCommentDateGmt,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentDateGmtAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentDateGmtAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentDateGmtAndCommentType(offset, limit, iCommentPostId,iCommentDateGmt,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentDateGmtAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentDateGmtAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentDateGmtAndCommentParent(offset, limit, iCommentPostId,iCommentDateGmt,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentDateGmtAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentDateGmtAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentDateGmtAndUserId(offset, limit, iCommentPostId,iCommentDateGmt,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentDateGmtAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentContentAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentContent := self.Args("comment_content").String()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentContentAndCommentKarma(offset, limit, iCommentPostId,iCommentContent,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentContentAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentContentAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentContent := self.Args("comment_content").String()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentContentAndCommentApproved(offset, limit, iCommentPostId,iCommentContent,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentContentAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentContentAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentContent := self.Args("comment_content").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentContentAndCommentAgent(offset, limit, iCommentPostId,iCommentContent,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentContentAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentContentAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentContent := self.Args("comment_content").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentContentAndCommentType(offset, limit, iCommentPostId,iCommentContent,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentContentAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentContentAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentContent := self.Args("comment_content").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentContentAndCommentParent(offset, limit, iCommentPostId,iCommentContent,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentContentAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentContentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentContent := self.Args("comment_content").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentContentAndUserId(offset, limit, iCommentPostId,iCommentContent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentContentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentKarmaAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentKarmaAndCommentApproved(offset, limit, iCommentPostId,iCommentKarma,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentKarmaAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentKarmaAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentKarmaAndCommentAgent(offset, limit, iCommentPostId,iCommentKarma,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentKarmaAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentKarmaAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentKarmaAndCommentType(offset, limit, iCommentPostId,iCommentKarma,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentKarmaAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentKarmaAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentKarmaAndCommentParent(offset, limit, iCommentPostId,iCommentKarma,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentKarmaAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentKarmaAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentKarmaAndUserId(offset, limit, iCommentPostId,iCommentKarma,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentKarmaAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentApprovedAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentApprovedAndCommentAgent(offset, limit, iCommentPostId,iCommentApproved,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentApprovedAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentApprovedAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentApprovedAndCommentType(offset, limit, iCommentPostId,iCommentApproved,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentApprovedAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentApprovedAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentApprovedAndCommentParent(offset, limit, iCommentPostId,iCommentApproved,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentApprovedAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentApprovedAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentApproved := self.Args("comment_approved").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentApprovedAndUserId(offset, limit, iCommentPostId,iCommentApproved,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentApprovedAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAgentAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAgent := self.Args("comment_agent").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAgentAndCommentType(offset, limit, iCommentPostId,iCommentAgent,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAgentAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAgentAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAgent := self.Args("comment_agent").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAgentAndCommentParent(offset, limit, iCommentPostId,iCommentAgent,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAgentAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAgentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAgent := self.Args("comment_agent").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAgentAndUserId(offset, limit, iCommentPostId,iCommentAgent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAgentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentTypeAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentType := self.Args("comment_type").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentTypeAndCommentParent(offset, limit, iCommentPostId,iCommentType,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentTypeAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentTypeAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentType := self.Args("comment_type").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentTypeAndUserId(offset, limit, iCommentPostId,iCommentType,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentTypeAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentParentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentParent := self.Args("comment_parent").MustInt64()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentParentAndUserId(offset, limit, iCommentPostId,iCommentParent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentParentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentAuthorUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAuthorUrl := self.Args("comment_author_url").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentAuthorUrl(offset, limit, iCommentAuthor,iCommentAuthorEmail,iCommentAuthorUrl)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentAuthorUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentAuthorIpHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentAuthorIp(offset, limit, iCommentAuthor,iCommentAuthorEmail,iCommentAuthorIp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentAuthorIp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentDateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentDate := self.Args("comment_date").Time()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentDate(offset, limit, iCommentAuthor,iCommentAuthorEmail,iCommentDate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentDate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentDateGmt(offset, limit, iCommentAuthor,iCommentAuthorEmail,iCommentDateGmt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentDateGmt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentContent(offset, limit, iCommentAuthor,iCommentAuthorEmail,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentKarma(offset, limit, iCommentAuthor,iCommentAuthorEmail,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentApproved(offset, limit, iCommentAuthor,iCommentAuthorEmail,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentAgent(offset, limit, iCommentAuthor,iCommentAuthorEmail,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentType(offset, limit, iCommentAuthor,iCommentAuthorEmail,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentParent(offset, limit, iCommentAuthor,iCommentAuthorEmail,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorEmailAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorEmailAndUserId(offset, limit, iCommentAuthor,iCommentAuthorEmail,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorEmailAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentAuthorIpHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentAuthorIp(offset, limit, iCommentAuthor,iCommentAuthorUrl,iCommentAuthorIp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentAuthorIp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentDateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentDate := self.Args("comment_date").Time()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentDate(offset, limit, iCommentAuthor,iCommentAuthorUrl,iCommentDate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentDate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentDateGmt(offset, limit, iCommentAuthor,iCommentAuthorUrl,iCommentDateGmt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentDateGmt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentContent(offset, limit, iCommentAuthor,iCommentAuthorUrl,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentKarma(offset, limit, iCommentAuthor,iCommentAuthorUrl,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentApproved(offset, limit, iCommentAuthor,iCommentAuthorUrl,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentAgent(offset, limit, iCommentAuthor,iCommentAuthorUrl,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentType(offset, limit, iCommentAuthor,iCommentAuthorUrl,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentParent(offset, limit, iCommentAuthor,iCommentAuthorUrl,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorUrlAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorUrlAndUserId(offset, limit, iCommentAuthor,iCommentAuthorUrl,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorUrlAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentDateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentDate := self.Args("comment_date").Time()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentDate(offset, limit, iCommentAuthor,iCommentAuthorIp,iCommentDate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentDate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentDateGmt(offset, limit, iCommentAuthor,iCommentAuthorIp,iCommentDateGmt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentDateGmt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentContent(offset, limit, iCommentAuthor,iCommentAuthorIp,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentKarma(offset, limit, iCommentAuthor,iCommentAuthorIp,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentApproved(offset, limit, iCommentAuthor,iCommentAuthorIp,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentAgent(offset, limit, iCommentAuthor,iCommentAuthorIp,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentType(offset, limit, iCommentAuthor,iCommentAuthorIp,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentParent(offset, limit, iCommentAuthor,iCommentAuthorIp,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorIpAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorIpAndUserId(offset, limit, iCommentAuthor,iCommentAuthorIp,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorIpAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentDateAndCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentDate := self.Args("comment_date").Time()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentDateAndCommentDateGmt(offset, limit, iCommentAuthor,iCommentDate,iCommentDateGmt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentDateAndCommentDateGmt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentDateAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentDate := self.Args("comment_date").Time()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentDateAndCommentContent(offset, limit, iCommentAuthor,iCommentDate,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentDateAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentDateAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentDate := self.Args("comment_date").Time()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentDateAndCommentKarma(offset, limit, iCommentAuthor,iCommentDate,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentDateAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentDateAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentDate := self.Args("comment_date").Time()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentDateAndCommentApproved(offset, limit, iCommentAuthor,iCommentDate,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentDateAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentDateAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentDate := self.Args("comment_date").Time()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentDateAndCommentAgent(offset, limit, iCommentAuthor,iCommentDate,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentDateAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentDateAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentDate := self.Args("comment_date").Time()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentDateAndCommentType(offset, limit, iCommentAuthor,iCommentDate,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentDateAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentDateAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentDate := self.Args("comment_date").Time()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentDateAndCommentParent(offset, limit, iCommentAuthor,iCommentDate,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentDateAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentDateAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentDate := self.Args("comment_date").Time()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentDateAndUserId(offset, limit, iCommentAuthor,iCommentDate,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentDateAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentDateGmtAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentDateGmtAndCommentContent(offset, limit, iCommentAuthor,iCommentDateGmt,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentDateGmtAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentDateGmtAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentDateGmtAndCommentKarma(offset, limit, iCommentAuthor,iCommentDateGmt,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentDateGmtAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentDateGmtAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentDateGmtAndCommentApproved(offset, limit, iCommentAuthor,iCommentDateGmt,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentDateGmtAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentDateGmtAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentDateGmtAndCommentAgent(offset, limit, iCommentAuthor,iCommentDateGmt,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentDateGmtAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentDateGmtAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentDateGmtAndCommentType(offset, limit, iCommentAuthor,iCommentDateGmt,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentDateGmtAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentDateGmtAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentDateGmtAndCommentParent(offset, limit, iCommentAuthor,iCommentDateGmt,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentDateGmtAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentDateGmtAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentDateGmtAndUserId(offset, limit, iCommentAuthor,iCommentDateGmt,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentDateGmtAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentContentAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentContent := self.Args("comment_content").String()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentContentAndCommentKarma(offset, limit, iCommentAuthor,iCommentContent,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentContentAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentContentAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentContent := self.Args("comment_content").String()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentContentAndCommentApproved(offset, limit, iCommentAuthor,iCommentContent,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentContentAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentContentAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentContent := self.Args("comment_content").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentContentAndCommentAgent(offset, limit, iCommentAuthor,iCommentContent,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentContentAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentContentAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentContent := self.Args("comment_content").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentContentAndCommentType(offset, limit, iCommentAuthor,iCommentContent,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentContentAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentContentAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentContent := self.Args("comment_content").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentContentAndCommentParent(offset, limit, iCommentAuthor,iCommentContent,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentContentAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentContentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentContent := self.Args("comment_content").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentContentAndUserId(offset, limit, iCommentAuthor,iCommentContent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentContentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentKarmaAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentKarmaAndCommentApproved(offset, limit, iCommentAuthor,iCommentKarma,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentKarmaAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentKarmaAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentKarmaAndCommentAgent(offset, limit, iCommentAuthor,iCommentKarma,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentKarmaAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentKarmaAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentKarmaAndCommentType(offset, limit, iCommentAuthor,iCommentKarma,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentKarmaAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentKarmaAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentKarmaAndCommentParent(offset, limit, iCommentAuthor,iCommentKarma,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentKarmaAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentKarmaAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentKarmaAndUserId(offset, limit, iCommentAuthor,iCommentKarma,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentKarmaAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentApprovedAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentApprovedAndCommentAgent(offset, limit, iCommentAuthor,iCommentApproved,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentApprovedAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentApprovedAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentApprovedAndCommentType(offset, limit, iCommentAuthor,iCommentApproved,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentApprovedAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentApprovedAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentApprovedAndCommentParent(offset, limit, iCommentAuthor,iCommentApproved,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentApprovedAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentApprovedAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentApproved := self.Args("comment_approved").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentApprovedAndUserId(offset, limit, iCommentAuthor,iCommentApproved,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentApprovedAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAgentAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAgent := self.Args("comment_agent").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAgentAndCommentType(offset, limit, iCommentAuthor,iCommentAgent,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAgentAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAgentAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAgent := self.Args("comment_agent").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAgentAndCommentParent(offset, limit, iCommentAuthor,iCommentAgent,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAgentAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAgentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAgent := self.Args("comment_agent").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAgentAndUserId(offset, limit, iCommentAuthor,iCommentAgent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAgentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentTypeAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentType := self.Args("comment_type").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentTypeAndCommentParent(offset, limit, iCommentAuthor,iCommentType,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentTypeAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentTypeAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentType := self.Args("comment_type").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentTypeAndUserId(offset, limit, iCommentAuthor,iCommentType,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentTypeAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentParentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentParent := self.Args("comment_parent").MustInt64()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentParentAndUserId(offset, limit, iCommentAuthor,iCommentParent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentParentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentAuthorIpHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentAuthorIp(offset, limit, iCommentAuthorEmail,iCommentAuthorUrl,iCommentAuthorIp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentAuthorIp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentDateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentDate := self.Args("comment_date").Time()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentDate(offset, limit, iCommentAuthorEmail,iCommentAuthorUrl,iCommentDate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentDate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentDateGmt(offset, limit, iCommentAuthorEmail,iCommentAuthorUrl,iCommentDateGmt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentDateGmt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentContent(offset, limit, iCommentAuthorEmail,iCommentAuthorUrl,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentKarma(offset, limit, iCommentAuthorEmail,iCommentAuthorUrl,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentApproved(offset, limit, iCommentAuthorEmail,iCommentAuthorUrl,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentAgent(offset, limit, iCommentAuthorEmail,iCommentAuthorUrl,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentType(offset, limit, iCommentAuthorEmail,iCommentAuthorUrl,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentParent(offset, limit, iCommentAuthorEmail,iCommentAuthorUrl,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndUserId(offset, limit, iCommentAuthorEmail,iCommentAuthorUrl,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentDateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentDate := self.Args("comment_date").Time()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentDate(offset, limit, iCommentAuthorEmail,iCommentAuthorIp,iCommentDate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentDate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentDateGmt(offset, limit, iCommentAuthorEmail,iCommentAuthorIp,iCommentDateGmt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentDateGmt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentContent(offset, limit, iCommentAuthorEmail,iCommentAuthorIp,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentKarma(offset, limit, iCommentAuthorEmail,iCommentAuthorIp,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentApproved(offset, limit, iCommentAuthorEmail,iCommentAuthorIp,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentAgent(offset, limit, iCommentAuthorEmail,iCommentAuthorIp,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentType(offset, limit, iCommentAuthorEmail,iCommentAuthorIp,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentParent(offset, limit, iCommentAuthorEmail,iCommentAuthorIp,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndUserId(offset, limit, iCommentAuthorEmail,iCommentAuthorIp,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentDate := self.Args("comment_date").Time()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentDateGmt(offset, limit, iCommentAuthorEmail,iCommentDate,iCommentDateGmt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentDateGmt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentDate := self.Args("comment_date").Time()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentContent(offset, limit, iCommentAuthorEmail,iCommentDate,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentDate := self.Args("comment_date").Time()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentKarma(offset, limit, iCommentAuthorEmail,iCommentDate,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentDate := self.Args("comment_date").Time()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentApproved(offset, limit, iCommentAuthorEmail,iCommentDate,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentDate := self.Args("comment_date").Time()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentAgent(offset, limit, iCommentAuthorEmail,iCommentDate,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentDate := self.Args("comment_date").Time()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentType(offset, limit, iCommentAuthorEmail,iCommentDate,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentDate := self.Args("comment_date").Time()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentParent(offset, limit, iCommentAuthorEmail,iCommentDate,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentDateAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentDate := self.Args("comment_date").Time()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentDateAndUserId(offset, limit, iCommentAuthorEmail,iCommentDate,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentDateAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndCommentContent(offset, limit, iCommentAuthorEmail,iCommentDateGmt,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndCommentKarma(offset, limit, iCommentAuthorEmail,iCommentDateGmt,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndCommentApproved(offset, limit, iCommentAuthorEmail,iCommentDateGmt,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndCommentAgent(offset, limit, iCommentAuthorEmail,iCommentDateGmt,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndCommentType(offset, limit, iCommentAuthorEmail,iCommentDateGmt,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndCommentParent(offset, limit, iCommentAuthorEmail,iCommentDateGmt,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndUserId(offset, limit, iCommentAuthorEmail,iCommentDateGmt,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentContentAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentContent := self.Args("comment_content").String()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentContentAndCommentKarma(offset, limit, iCommentAuthorEmail,iCommentContent,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentContentAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentContentAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentContent := self.Args("comment_content").String()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentContentAndCommentApproved(offset, limit, iCommentAuthorEmail,iCommentContent,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentContentAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentContentAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentContent := self.Args("comment_content").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentContentAndCommentAgent(offset, limit, iCommentAuthorEmail,iCommentContent,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentContentAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentContentAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentContent := self.Args("comment_content").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentContentAndCommentType(offset, limit, iCommentAuthorEmail,iCommentContent,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentContentAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentContentAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentContent := self.Args("comment_content").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentContentAndCommentParent(offset, limit, iCommentAuthorEmail,iCommentContent,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentContentAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentContentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentContent := self.Args("comment_content").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentContentAndUserId(offset, limit, iCommentAuthorEmail,iCommentContent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentContentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentKarmaAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentKarmaAndCommentApproved(offset, limit, iCommentAuthorEmail,iCommentKarma,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentKarmaAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentKarmaAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentKarmaAndCommentAgent(offset, limit, iCommentAuthorEmail,iCommentKarma,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentKarmaAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentKarmaAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentKarmaAndCommentType(offset, limit, iCommentAuthorEmail,iCommentKarma,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentKarmaAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentKarmaAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentKarmaAndCommentParent(offset, limit, iCommentAuthorEmail,iCommentKarma,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentKarmaAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentKarmaAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentKarmaAndUserId(offset, limit, iCommentAuthorEmail,iCommentKarma,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentKarmaAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentApprovedAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentApprovedAndCommentAgent(offset, limit, iCommentAuthorEmail,iCommentApproved,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentApprovedAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentApprovedAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentApprovedAndCommentType(offset, limit, iCommentAuthorEmail,iCommentApproved,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentApprovedAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentApprovedAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentApprovedAndCommentParent(offset, limit, iCommentAuthorEmail,iCommentApproved,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentApprovedAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentApprovedAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentApproved := self.Args("comment_approved").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentApprovedAndUserId(offset, limit, iCommentAuthorEmail,iCommentApproved,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentApprovedAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentAgentAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAgent := self.Args("comment_agent").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentAgentAndCommentType(offset, limit, iCommentAuthorEmail,iCommentAgent,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentAgentAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentAgentAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAgent := self.Args("comment_agent").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentAgentAndCommentParent(offset, limit, iCommentAuthorEmail,iCommentAgent,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentAgentAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentAgentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAgent := self.Args("comment_agent").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentAgentAndUserId(offset, limit, iCommentAuthorEmail,iCommentAgent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentAgentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentTypeAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentType := self.Args("comment_type").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentTypeAndCommentParent(offset, limit, iCommentAuthorEmail,iCommentType,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentTypeAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentTypeAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentType := self.Args("comment_type").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentTypeAndUserId(offset, limit, iCommentAuthorEmail,iCommentType,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentTypeAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentParentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentParent := self.Args("comment_parent").MustInt64()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentParentAndUserId(offset, limit, iCommentAuthorEmail,iCommentParent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentParentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentDateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentDate := self.Args("comment_date").Time()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentDate(offset, limit, iCommentAuthorUrl,iCommentAuthorIp,iCommentDate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentDate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentDateGmt(offset, limit, iCommentAuthorUrl,iCommentAuthorIp,iCommentDateGmt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentDateGmt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentContent(offset, limit, iCommentAuthorUrl,iCommentAuthorIp,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentKarma(offset, limit, iCommentAuthorUrl,iCommentAuthorIp,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentApproved(offset, limit, iCommentAuthorUrl,iCommentAuthorIp,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentAgent(offset, limit, iCommentAuthorUrl,iCommentAuthorIp,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentType(offset, limit, iCommentAuthorUrl,iCommentAuthorIp,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentParent(offset, limit, iCommentAuthorUrl,iCommentAuthorIp,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndUserId(offset, limit, iCommentAuthorUrl,iCommentAuthorIp,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentDate := self.Args("comment_date").Time()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentDateGmt(offset, limit, iCommentAuthorUrl,iCommentDate,iCommentDateGmt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentDateGmt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentDate := self.Args("comment_date").Time()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentContent(offset, limit, iCommentAuthorUrl,iCommentDate,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentDate := self.Args("comment_date").Time()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentKarma(offset, limit, iCommentAuthorUrl,iCommentDate,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentDate := self.Args("comment_date").Time()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentApproved(offset, limit, iCommentAuthorUrl,iCommentDate,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentDate := self.Args("comment_date").Time()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentAgent(offset, limit, iCommentAuthorUrl,iCommentDate,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentDate := self.Args("comment_date").Time()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentType(offset, limit, iCommentAuthorUrl,iCommentDate,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentDate := self.Args("comment_date").Time()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentParent(offset, limit, iCommentAuthorUrl,iCommentDate,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentDateAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentDate := self.Args("comment_date").Time()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentDateAndUserId(offset, limit, iCommentAuthorUrl,iCommentDate,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentDateAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndCommentContent(offset, limit, iCommentAuthorUrl,iCommentDateGmt,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndCommentKarma(offset, limit, iCommentAuthorUrl,iCommentDateGmt,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndCommentApproved(offset, limit, iCommentAuthorUrl,iCommentDateGmt,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndCommentAgent(offset, limit, iCommentAuthorUrl,iCommentDateGmt,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndCommentType(offset, limit, iCommentAuthorUrl,iCommentDateGmt,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndCommentParent(offset, limit, iCommentAuthorUrl,iCommentDateGmt,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndUserId(offset, limit, iCommentAuthorUrl,iCommentDateGmt,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentContentAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentContent := self.Args("comment_content").String()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentContentAndCommentKarma(offset, limit, iCommentAuthorUrl,iCommentContent,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentContentAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentContentAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentContent := self.Args("comment_content").String()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentContentAndCommentApproved(offset, limit, iCommentAuthorUrl,iCommentContent,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentContentAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentContentAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentContent := self.Args("comment_content").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentContentAndCommentAgent(offset, limit, iCommentAuthorUrl,iCommentContent,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentContentAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentContentAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentContent := self.Args("comment_content").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentContentAndCommentType(offset, limit, iCommentAuthorUrl,iCommentContent,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentContentAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentContentAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentContent := self.Args("comment_content").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentContentAndCommentParent(offset, limit, iCommentAuthorUrl,iCommentContent,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentContentAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentContentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentContent := self.Args("comment_content").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentContentAndUserId(offset, limit, iCommentAuthorUrl,iCommentContent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentContentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentKarmaAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentKarmaAndCommentApproved(offset, limit, iCommentAuthorUrl,iCommentKarma,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentKarmaAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentKarmaAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentKarmaAndCommentAgent(offset, limit, iCommentAuthorUrl,iCommentKarma,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentKarmaAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentKarmaAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentKarmaAndCommentType(offset, limit, iCommentAuthorUrl,iCommentKarma,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentKarmaAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentKarmaAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentKarmaAndCommentParent(offset, limit, iCommentAuthorUrl,iCommentKarma,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentKarmaAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentKarmaAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentKarmaAndUserId(offset, limit, iCommentAuthorUrl,iCommentKarma,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentKarmaAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentApprovedAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentApprovedAndCommentAgent(offset, limit, iCommentAuthorUrl,iCommentApproved,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentApprovedAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentApprovedAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentApprovedAndCommentType(offset, limit, iCommentAuthorUrl,iCommentApproved,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentApprovedAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentApprovedAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentApprovedAndCommentParent(offset, limit, iCommentAuthorUrl,iCommentApproved,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentApprovedAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentApprovedAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentApproved := self.Args("comment_approved").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentApprovedAndUserId(offset, limit, iCommentAuthorUrl,iCommentApproved,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentApprovedAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentAgentAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentAgent := self.Args("comment_agent").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentAgentAndCommentType(offset, limit, iCommentAuthorUrl,iCommentAgent,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentAgentAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentAgentAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentAgent := self.Args("comment_agent").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentAgentAndCommentParent(offset, limit, iCommentAuthorUrl,iCommentAgent,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentAgentAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentAgentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentAgent := self.Args("comment_agent").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentAgentAndUserId(offset, limit, iCommentAuthorUrl,iCommentAgent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentAgentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentTypeAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentType := self.Args("comment_type").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentTypeAndCommentParent(offset, limit, iCommentAuthorUrl,iCommentType,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentTypeAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentTypeAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentType := self.Args("comment_type").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentTypeAndUserId(offset, limit, iCommentAuthorUrl,iCommentType,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentTypeAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentParentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentParent := self.Args("comment_parent").MustInt64()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentParentAndUserId(offset, limit, iCommentAuthorUrl,iCommentParent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentParentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentDateAndCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentDate := self.Args("comment_date").Time()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentDateAndCommentDateGmt(offset, limit, iCommentAuthorIp,iCommentDate,iCommentDateGmt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentDateAndCommentDateGmt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentDateAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentDate := self.Args("comment_date").Time()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentDateAndCommentContent(offset, limit, iCommentAuthorIp,iCommentDate,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentDateAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentDateAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentDate := self.Args("comment_date").Time()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentDateAndCommentKarma(offset, limit, iCommentAuthorIp,iCommentDate,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentDateAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentDateAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentDate := self.Args("comment_date").Time()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentDateAndCommentApproved(offset, limit, iCommentAuthorIp,iCommentDate,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentDateAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentDateAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentDate := self.Args("comment_date").Time()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentDateAndCommentAgent(offset, limit, iCommentAuthorIp,iCommentDate,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentDateAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentDateAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentDate := self.Args("comment_date").Time()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentDateAndCommentType(offset, limit, iCommentAuthorIp,iCommentDate,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentDateAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentDateAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentDate := self.Args("comment_date").Time()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentDateAndCommentParent(offset, limit, iCommentAuthorIp,iCommentDate,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentDateAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentDateAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentDate := self.Args("comment_date").Time()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentDateAndUserId(offset, limit, iCommentAuthorIp,iCommentDate,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentDateAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentDateGmtAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentDateGmtAndCommentContent(offset, limit, iCommentAuthorIp,iCommentDateGmt,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentDateGmtAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentDateGmtAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentDateGmtAndCommentKarma(offset, limit, iCommentAuthorIp,iCommentDateGmt,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentDateGmtAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentDateGmtAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentDateGmtAndCommentApproved(offset, limit, iCommentAuthorIp,iCommentDateGmt,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentDateGmtAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentDateGmtAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentDateGmtAndCommentAgent(offset, limit, iCommentAuthorIp,iCommentDateGmt,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentDateGmtAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentDateGmtAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentDateGmtAndCommentType(offset, limit, iCommentAuthorIp,iCommentDateGmt,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentDateGmtAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentDateGmtAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentDateGmtAndCommentParent(offset, limit, iCommentAuthorIp,iCommentDateGmt,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentDateGmtAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentDateGmtAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentDateGmtAndUserId(offset, limit, iCommentAuthorIp,iCommentDateGmt,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentDateGmtAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentContentAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentContent := self.Args("comment_content").String()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentContentAndCommentKarma(offset, limit, iCommentAuthorIp,iCommentContent,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentContentAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentContentAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentContent := self.Args("comment_content").String()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentContentAndCommentApproved(offset, limit, iCommentAuthorIp,iCommentContent,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentContentAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentContentAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentContent := self.Args("comment_content").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentContentAndCommentAgent(offset, limit, iCommentAuthorIp,iCommentContent,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentContentAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentContentAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentContent := self.Args("comment_content").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentContentAndCommentType(offset, limit, iCommentAuthorIp,iCommentContent,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentContentAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentContentAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentContent := self.Args("comment_content").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentContentAndCommentParent(offset, limit, iCommentAuthorIp,iCommentContent,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentContentAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentContentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentContent := self.Args("comment_content").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentContentAndUserId(offset, limit, iCommentAuthorIp,iCommentContent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentContentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentKarmaAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentKarmaAndCommentApproved(offset, limit, iCommentAuthorIp,iCommentKarma,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentKarmaAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentKarmaAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentKarmaAndCommentAgent(offset, limit, iCommentAuthorIp,iCommentKarma,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentKarmaAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentKarmaAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentKarmaAndCommentType(offset, limit, iCommentAuthorIp,iCommentKarma,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentKarmaAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentKarmaAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentKarmaAndCommentParent(offset, limit, iCommentAuthorIp,iCommentKarma,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentKarmaAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentKarmaAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentKarmaAndUserId(offset, limit, iCommentAuthorIp,iCommentKarma,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentKarmaAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentApprovedAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentApprovedAndCommentAgent(offset, limit, iCommentAuthorIp,iCommentApproved,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentApprovedAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentApprovedAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentApprovedAndCommentType(offset, limit, iCommentAuthorIp,iCommentApproved,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentApprovedAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentApprovedAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentApprovedAndCommentParent(offset, limit, iCommentAuthorIp,iCommentApproved,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentApprovedAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentApprovedAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentApproved := self.Args("comment_approved").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentApprovedAndUserId(offset, limit, iCommentAuthorIp,iCommentApproved,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentApprovedAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentAgentAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentAgent := self.Args("comment_agent").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentAgentAndCommentType(offset, limit, iCommentAuthorIp,iCommentAgent,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentAgentAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentAgentAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentAgent := self.Args("comment_agent").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentAgentAndCommentParent(offset, limit, iCommentAuthorIp,iCommentAgent,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentAgentAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentAgentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentAgent := self.Args("comment_agent").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentAgentAndUserId(offset, limit, iCommentAuthorIp,iCommentAgent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentAgentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentTypeAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentType := self.Args("comment_type").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentTypeAndCommentParent(offset, limit, iCommentAuthorIp,iCommentType,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentTypeAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentTypeAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentType := self.Args("comment_type").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentTypeAndUserId(offset, limit, iCommentAuthorIp,iCommentType,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentTypeAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentParentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentParent := self.Args("comment_parent").MustInt64()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentParentAndUserId(offset, limit, iCommentAuthorIp,iCommentParent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentParentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentDateGmtAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentDateGmtAndCommentContent(offset, limit, iCommentDate,iCommentDateGmt,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentDateGmtAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentDateGmtAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentDateGmtAndCommentKarma(offset, limit, iCommentDate,iCommentDateGmt,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentDateGmtAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentDateGmtAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentDateGmtAndCommentApproved(offset, limit, iCommentDate,iCommentDateGmt,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentDateGmtAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentDateGmtAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentDateGmtAndCommentAgent(offset, limit, iCommentDate,iCommentDateGmt,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentDateGmtAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentDateGmtAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentDateGmtAndCommentType(offset, limit, iCommentDate,iCommentDateGmt,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentDateGmtAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentDateGmtAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentDateGmtAndCommentParent(offset, limit, iCommentDate,iCommentDateGmt,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentDateGmtAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentDateGmtAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentDateGmtAndUserId(offset, limit, iCommentDate,iCommentDateGmt,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentDateGmtAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentContentAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentContent := self.Args("comment_content").String()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentContentAndCommentKarma(offset, limit, iCommentDate,iCommentContent,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentContentAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentContentAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentContent := self.Args("comment_content").String()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentContentAndCommentApproved(offset, limit, iCommentDate,iCommentContent,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentContentAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentContentAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentContent := self.Args("comment_content").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentContentAndCommentAgent(offset, limit, iCommentDate,iCommentContent,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentContentAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentContentAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentContent := self.Args("comment_content").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentContentAndCommentType(offset, limit, iCommentDate,iCommentContent,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentContentAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentContentAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentContent := self.Args("comment_content").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentContentAndCommentParent(offset, limit, iCommentDate,iCommentContent,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentContentAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentContentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentContent := self.Args("comment_content").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentContentAndUserId(offset, limit, iCommentDate,iCommentContent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentContentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentKarmaAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentKarmaAndCommentApproved(offset, limit, iCommentDate,iCommentKarma,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentKarmaAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentKarmaAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentKarmaAndCommentAgent(offset, limit, iCommentDate,iCommentKarma,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentKarmaAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentKarmaAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentKarmaAndCommentType(offset, limit, iCommentDate,iCommentKarma,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentKarmaAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentKarmaAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentKarmaAndCommentParent(offset, limit, iCommentDate,iCommentKarma,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentKarmaAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentKarmaAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentKarmaAndUserId(offset, limit, iCommentDate,iCommentKarma,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentKarmaAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentApprovedAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentApprovedAndCommentAgent(offset, limit, iCommentDate,iCommentApproved,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentApprovedAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentApprovedAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentApprovedAndCommentType(offset, limit, iCommentDate,iCommentApproved,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentApprovedAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentApprovedAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentApprovedAndCommentParent(offset, limit, iCommentDate,iCommentApproved,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentApprovedAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentApprovedAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentApproved := self.Args("comment_approved").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentApprovedAndUserId(offset, limit, iCommentDate,iCommentApproved,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentApprovedAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentAgentAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentAgent := self.Args("comment_agent").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentAgentAndCommentType(offset, limit, iCommentDate,iCommentAgent,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentAgentAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentAgentAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentAgent := self.Args("comment_agent").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentAgentAndCommentParent(offset, limit, iCommentDate,iCommentAgent,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentAgentAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentAgentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentAgent := self.Args("comment_agent").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentAgentAndUserId(offset, limit, iCommentDate,iCommentAgent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentAgentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentTypeAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentType := self.Args("comment_type").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentTypeAndCommentParent(offset, limit, iCommentDate,iCommentType,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentTypeAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentTypeAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentType := self.Args("comment_type").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentTypeAndUserId(offset, limit, iCommentDate,iCommentType,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentTypeAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentParentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentParent := self.Args("comment_parent").MustInt64()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentParentAndUserId(offset, limit, iCommentDate,iCommentParent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentParentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateGmtAndCommentContentAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentContent := self.Args("comment_content").String()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentDateGmt) {
		_Comments, _error := model.GetCommentsesByCommentDateGmtAndCommentContentAndCommentKarma(offset, limit, iCommentDateGmt,iCommentContent,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateGmtAndCommentContentAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateGmtAndCommentContentAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentContent := self.Args("comment_content").String()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentDateGmt) {
		_Comments, _error := model.GetCommentsesByCommentDateGmtAndCommentContentAndCommentApproved(offset, limit, iCommentDateGmt,iCommentContent,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateGmtAndCommentContentAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateGmtAndCommentContentAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentContent := self.Args("comment_content").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentDateGmt) {
		_Comments, _error := model.GetCommentsesByCommentDateGmtAndCommentContentAndCommentAgent(offset, limit, iCommentDateGmt,iCommentContent,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateGmtAndCommentContentAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateGmtAndCommentContentAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentContent := self.Args("comment_content").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentDateGmt) {
		_Comments, _error := model.GetCommentsesByCommentDateGmtAndCommentContentAndCommentType(offset, limit, iCommentDateGmt,iCommentContent,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateGmtAndCommentContentAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateGmtAndCommentContentAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentContent := self.Args("comment_content").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentDateGmt) {
		_Comments, _error := model.GetCommentsesByCommentDateGmtAndCommentContentAndCommentParent(offset, limit, iCommentDateGmt,iCommentContent,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateGmtAndCommentContentAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateGmtAndCommentContentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentContent := self.Args("comment_content").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentDateGmt) {
		_Comments, _error := model.GetCommentsesByCommentDateGmtAndCommentContentAndUserId(offset, limit, iCommentDateGmt,iCommentContent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateGmtAndCommentContentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateGmtAndCommentKarmaAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentDateGmt) {
		_Comments, _error := model.GetCommentsesByCommentDateGmtAndCommentKarmaAndCommentApproved(offset, limit, iCommentDateGmt,iCommentKarma,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateGmtAndCommentKarmaAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateGmtAndCommentKarmaAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentDateGmt) {
		_Comments, _error := model.GetCommentsesByCommentDateGmtAndCommentKarmaAndCommentAgent(offset, limit, iCommentDateGmt,iCommentKarma,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateGmtAndCommentKarmaAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateGmtAndCommentKarmaAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentDateGmt) {
		_Comments, _error := model.GetCommentsesByCommentDateGmtAndCommentKarmaAndCommentType(offset, limit, iCommentDateGmt,iCommentKarma,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateGmtAndCommentKarmaAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateGmtAndCommentKarmaAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentDateGmt) {
		_Comments, _error := model.GetCommentsesByCommentDateGmtAndCommentKarmaAndCommentParent(offset, limit, iCommentDateGmt,iCommentKarma,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateGmtAndCommentKarmaAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateGmtAndCommentKarmaAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentDateGmt) {
		_Comments, _error := model.GetCommentsesByCommentDateGmtAndCommentKarmaAndUserId(offset, limit, iCommentDateGmt,iCommentKarma,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateGmtAndCommentKarmaAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateGmtAndCommentApprovedAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentDateGmt) {
		_Comments, _error := model.GetCommentsesByCommentDateGmtAndCommentApprovedAndCommentAgent(offset, limit, iCommentDateGmt,iCommentApproved,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateGmtAndCommentApprovedAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateGmtAndCommentApprovedAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentDateGmt) {
		_Comments, _error := model.GetCommentsesByCommentDateGmtAndCommentApprovedAndCommentType(offset, limit, iCommentDateGmt,iCommentApproved,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateGmtAndCommentApprovedAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateGmtAndCommentApprovedAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentDateGmt) {
		_Comments, _error := model.GetCommentsesByCommentDateGmtAndCommentApprovedAndCommentParent(offset, limit, iCommentDateGmt,iCommentApproved,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateGmtAndCommentApprovedAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateGmtAndCommentApprovedAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentApproved := self.Args("comment_approved").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentDateGmt) {
		_Comments, _error := model.GetCommentsesByCommentDateGmtAndCommentApprovedAndUserId(offset, limit, iCommentDateGmt,iCommentApproved,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateGmtAndCommentApprovedAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateGmtAndCommentAgentAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentAgent := self.Args("comment_agent").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentDateGmt) {
		_Comments, _error := model.GetCommentsesByCommentDateGmtAndCommentAgentAndCommentType(offset, limit, iCommentDateGmt,iCommentAgent,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateGmtAndCommentAgentAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateGmtAndCommentAgentAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentAgent := self.Args("comment_agent").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentDateGmt) {
		_Comments, _error := model.GetCommentsesByCommentDateGmtAndCommentAgentAndCommentParent(offset, limit, iCommentDateGmt,iCommentAgent,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateGmtAndCommentAgentAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateGmtAndCommentAgentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentAgent := self.Args("comment_agent").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentDateGmt) {
		_Comments, _error := model.GetCommentsesByCommentDateGmtAndCommentAgentAndUserId(offset, limit, iCommentDateGmt,iCommentAgent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateGmtAndCommentAgentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateGmtAndCommentTypeAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentType := self.Args("comment_type").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentDateGmt) {
		_Comments, _error := model.GetCommentsesByCommentDateGmtAndCommentTypeAndCommentParent(offset, limit, iCommentDateGmt,iCommentType,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateGmtAndCommentTypeAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateGmtAndCommentTypeAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentType := self.Args("comment_type").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentDateGmt) {
		_Comments, _error := model.GetCommentsesByCommentDateGmtAndCommentTypeAndUserId(offset, limit, iCommentDateGmt,iCommentType,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateGmtAndCommentTypeAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateGmtAndCommentParentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentParent := self.Args("comment_parent").MustInt64()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentDateGmt) {
		_Comments, _error := model.GetCommentsesByCommentDateGmtAndCommentParentAndUserId(offset, limit, iCommentDateGmt,iCommentParent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateGmtAndCommentParentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentContentAndCommentKarmaAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentContent := self.Args("comment_content").String()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentContent) {
		_Comments, _error := model.GetCommentsesByCommentContentAndCommentKarmaAndCommentApproved(offset, limit, iCommentContent,iCommentKarma,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentContentAndCommentKarmaAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentContentAndCommentKarmaAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentContent := self.Args("comment_content").String()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentContent) {
		_Comments, _error := model.GetCommentsesByCommentContentAndCommentKarmaAndCommentAgent(offset, limit, iCommentContent,iCommentKarma,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentContentAndCommentKarmaAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentContentAndCommentKarmaAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentContent := self.Args("comment_content").String()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentContent) {
		_Comments, _error := model.GetCommentsesByCommentContentAndCommentKarmaAndCommentType(offset, limit, iCommentContent,iCommentKarma,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentContentAndCommentKarmaAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentContentAndCommentKarmaAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentContent := self.Args("comment_content").String()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentContent) {
		_Comments, _error := model.GetCommentsesByCommentContentAndCommentKarmaAndCommentParent(offset, limit, iCommentContent,iCommentKarma,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentContentAndCommentKarmaAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentContentAndCommentKarmaAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentContent := self.Args("comment_content").String()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentContent) {
		_Comments, _error := model.GetCommentsesByCommentContentAndCommentKarmaAndUserId(offset, limit, iCommentContent,iCommentKarma,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentContentAndCommentKarmaAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentContentAndCommentApprovedAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentContent := self.Args("comment_content").String()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentContent) {
		_Comments, _error := model.GetCommentsesByCommentContentAndCommentApprovedAndCommentAgent(offset, limit, iCommentContent,iCommentApproved,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentContentAndCommentApprovedAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentContentAndCommentApprovedAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentContent := self.Args("comment_content").String()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentContent) {
		_Comments, _error := model.GetCommentsesByCommentContentAndCommentApprovedAndCommentType(offset, limit, iCommentContent,iCommentApproved,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentContentAndCommentApprovedAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentContentAndCommentApprovedAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentContent := self.Args("comment_content").String()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentContent) {
		_Comments, _error := model.GetCommentsesByCommentContentAndCommentApprovedAndCommentParent(offset, limit, iCommentContent,iCommentApproved,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentContentAndCommentApprovedAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentContentAndCommentApprovedAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentContent := self.Args("comment_content").String()
	iCommentApproved := self.Args("comment_approved").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentContent) {
		_Comments, _error := model.GetCommentsesByCommentContentAndCommentApprovedAndUserId(offset, limit, iCommentContent,iCommentApproved,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentContentAndCommentApprovedAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentContentAndCommentAgentAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentContent := self.Args("comment_content").String()
	iCommentAgent := self.Args("comment_agent").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentContent) {
		_Comments, _error := model.GetCommentsesByCommentContentAndCommentAgentAndCommentType(offset, limit, iCommentContent,iCommentAgent,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentContentAndCommentAgentAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentContentAndCommentAgentAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentContent := self.Args("comment_content").String()
	iCommentAgent := self.Args("comment_agent").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentContent) {
		_Comments, _error := model.GetCommentsesByCommentContentAndCommentAgentAndCommentParent(offset, limit, iCommentContent,iCommentAgent,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentContentAndCommentAgentAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentContentAndCommentAgentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentContent := self.Args("comment_content").String()
	iCommentAgent := self.Args("comment_agent").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentContent) {
		_Comments, _error := model.GetCommentsesByCommentContentAndCommentAgentAndUserId(offset, limit, iCommentContent,iCommentAgent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentContentAndCommentAgentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentContentAndCommentTypeAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentContent := self.Args("comment_content").String()
	iCommentType := self.Args("comment_type").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentContent) {
		_Comments, _error := model.GetCommentsesByCommentContentAndCommentTypeAndCommentParent(offset, limit, iCommentContent,iCommentType,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentContentAndCommentTypeAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentContentAndCommentTypeAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentContent := self.Args("comment_content").String()
	iCommentType := self.Args("comment_type").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentContent) {
		_Comments, _error := model.GetCommentsesByCommentContentAndCommentTypeAndUserId(offset, limit, iCommentContent,iCommentType,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentContentAndCommentTypeAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentContentAndCommentParentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentContent := self.Args("comment_content").String()
	iCommentParent := self.Args("comment_parent").MustInt64()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentContent) {
		_Comments, _error := model.GetCommentsesByCommentContentAndCommentParentAndUserId(offset, limit, iCommentContent,iCommentParent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentContentAndCommentParentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentKarmaAndCommentApprovedAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentKarma) {
		_Comments, _error := model.GetCommentsesByCommentKarmaAndCommentApprovedAndCommentAgent(offset, limit, iCommentKarma,iCommentApproved,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentKarmaAndCommentApprovedAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentKarmaAndCommentApprovedAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentKarma) {
		_Comments, _error := model.GetCommentsesByCommentKarmaAndCommentApprovedAndCommentType(offset, limit, iCommentKarma,iCommentApproved,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentKarmaAndCommentApprovedAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentKarmaAndCommentApprovedAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentKarma) {
		_Comments, _error := model.GetCommentsesByCommentKarmaAndCommentApprovedAndCommentParent(offset, limit, iCommentKarma,iCommentApproved,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentKarmaAndCommentApprovedAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentKarmaAndCommentApprovedAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentApproved := self.Args("comment_approved").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentKarma) {
		_Comments, _error := model.GetCommentsesByCommentKarmaAndCommentApprovedAndUserId(offset, limit, iCommentKarma,iCommentApproved,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentKarmaAndCommentApprovedAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentKarmaAndCommentAgentAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentAgent := self.Args("comment_agent").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentKarma) {
		_Comments, _error := model.GetCommentsesByCommentKarmaAndCommentAgentAndCommentType(offset, limit, iCommentKarma,iCommentAgent,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentKarmaAndCommentAgentAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentKarmaAndCommentAgentAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentAgent := self.Args("comment_agent").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentKarma) {
		_Comments, _error := model.GetCommentsesByCommentKarmaAndCommentAgentAndCommentParent(offset, limit, iCommentKarma,iCommentAgent,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentKarmaAndCommentAgentAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentKarmaAndCommentAgentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentAgent := self.Args("comment_agent").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentKarma) {
		_Comments, _error := model.GetCommentsesByCommentKarmaAndCommentAgentAndUserId(offset, limit, iCommentKarma,iCommentAgent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentKarmaAndCommentAgentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentKarmaAndCommentTypeAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentType := self.Args("comment_type").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentKarma) {
		_Comments, _error := model.GetCommentsesByCommentKarmaAndCommentTypeAndCommentParent(offset, limit, iCommentKarma,iCommentType,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentKarmaAndCommentTypeAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentKarmaAndCommentTypeAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentType := self.Args("comment_type").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentKarma) {
		_Comments, _error := model.GetCommentsesByCommentKarmaAndCommentTypeAndUserId(offset, limit, iCommentKarma,iCommentType,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentKarmaAndCommentTypeAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentKarmaAndCommentParentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentParent := self.Args("comment_parent").MustInt64()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentKarma) {
		_Comments, _error := model.GetCommentsesByCommentKarmaAndCommentParentAndUserId(offset, limit, iCommentKarma,iCommentParent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentKarmaAndCommentParentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentApprovedAndCommentAgentAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentAgent := self.Args("comment_agent").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentApproved) {
		_Comments, _error := model.GetCommentsesByCommentApprovedAndCommentAgentAndCommentType(offset, limit, iCommentApproved,iCommentAgent,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentApprovedAndCommentAgentAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentApprovedAndCommentAgentAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentAgent := self.Args("comment_agent").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentApproved) {
		_Comments, _error := model.GetCommentsesByCommentApprovedAndCommentAgentAndCommentParent(offset, limit, iCommentApproved,iCommentAgent,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentApprovedAndCommentAgentAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentApprovedAndCommentAgentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentAgent := self.Args("comment_agent").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentApproved) {
		_Comments, _error := model.GetCommentsesByCommentApprovedAndCommentAgentAndUserId(offset, limit, iCommentApproved,iCommentAgent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentApprovedAndCommentAgentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentApprovedAndCommentTypeAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentType := self.Args("comment_type").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentApproved) {
		_Comments, _error := model.GetCommentsesByCommentApprovedAndCommentTypeAndCommentParent(offset, limit, iCommentApproved,iCommentType,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentApprovedAndCommentTypeAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentApprovedAndCommentTypeAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentType := self.Args("comment_type").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentApproved) {
		_Comments, _error := model.GetCommentsesByCommentApprovedAndCommentTypeAndUserId(offset, limit, iCommentApproved,iCommentType,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentApprovedAndCommentTypeAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentApprovedAndCommentParentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentParent := self.Args("comment_parent").MustInt64()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentApproved) {
		_Comments, _error := model.GetCommentsesByCommentApprovedAndCommentParentAndUserId(offset, limit, iCommentApproved,iCommentParent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentApprovedAndCommentParentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAgentAndCommentTypeAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAgent := self.Args("comment_agent").String()
	iCommentType := self.Args("comment_type").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAgent) {
		_Comments, _error := model.GetCommentsesByCommentAgentAndCommentTypeAndCommentParent(offset, limit, iCommentAgent,iCommentType,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAgentAndCommentTypeAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAgentAndCommentTypeAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAgent := self.Args("comment_agent").String()
	iCommentType := self.Args("comment_type").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAgent) {
		_Comments, _error := model.GetCommentsesByCommentAgentAndCommentTypeAndUserId(offset, limit, iCommentAgent,iCommentType,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAgentAndCommentTypeAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAgentAndCommentParentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAgent := self.Args("comment_agent").String()
	iCommentParent := self.Args("comment_parent").MustInt64()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAgent) {
		_Comments, _error := model.GetCommentsesByCommentAgentAndCommentParentAndUserId(offset, limit, iCommentAgent,iCommentParent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAgentAndCommentParentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentTypeAndCommentParentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iCommentParent := self.Args("comment_parent").MustInt64()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentType) {
		_Comments, _error := model.GetCommentsesByCommentTypeAndCommentParentAndUserId(offset, limit, iCommentType,iCommentParent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentTypeAndCommentParentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentPostIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentPostId := self.Args("comment_post_id").MustInt64()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentPostId(offset, limit, iCommentId,iCommentPostId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentPostId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthor := self.Args("comment_author").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthor(offset, limit, iCommentId,iCommentAuthor)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthor's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorEmail := self.Args("comment_author_email").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorEmail(offset, limit, iCommentId,iCommentAuthorEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorUrl := self.Args("comment_author_url").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorUrl(offset, limit, iCommentId,iCommentAuthorUrl)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAuthorIpHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAuthorIp := self.Args("comment_author_ip").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAuthorIp(offset, limit, iCommentId,iCommentAuthorIp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAuthorIp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentDateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentDate := self.Args("comment_date").Time()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentDate(offset, limit, iCommentId,iCommentDate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentDate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentDateGmt(offset, limit, iCommentId,iCommentDateGmt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentDateGmt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentContent(offset, limit, iCommentId,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentKarma(offset, limit, iCommentId,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentApproved(offset, limit, iCommentId,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentAgent(offset, limit, iCommentId,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentType(offset, limit, iCommentId,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndCommentParent(offset, limit, iCommentId,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentIdAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsesByCommentIdAndUserId(offset, limit, iCommentId,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentIdAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthor := self.Args("comment_author").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthor(offset, limit, iCommentPostId,iCommentAuthor)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthor's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorEmail := self.Args("comment_author_email").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorEmail(offset, limit, iCommentPostId,iCommentAuthorEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorUrl := self.Args("comment_author_url").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorUrl(offset, limit, iCommentPostId,iCommentAuthorUrl)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAuthorIpHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAuthorIp := self.Args("comment_author_ip").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAuthorIp(offset, limit, iCommentPostId,iCommentAuthorIp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAuthorIp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentDateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentDate := self.Args("comment_date").Time()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentDate(offset, limit, iCommentPostId,iCommentDate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentDate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentDateGmt(offset, limit, iCommentPostId,iCommentDateGmt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentDateGmt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentContent(offset, limit, iCommentPostId,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentKarma(offset, limit, iCommentPostId,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentApproved(offset, limit, iCommentPostId,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentAgent(offset, limit, iCommentPostId,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentType(offset, limit, iCommentPostId,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndCommentParent(offset, limit, iCommentPostId,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentPostIdAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentPostId := self.Args("comment_post_id").MustInt64()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsesByCommentPostIdAndUserId(offset, limit, iCommentPostId,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentPostIdAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorEmail := self.Args("comment_author_email").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorEmail(offset, limit, iCommentAuthor,iCommentAuthorEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorUrl := self.Args("comment_author_url").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorUrl(offset, limit, iCommentAuthor,iCommentAuthorUrl)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAuthorIpHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAuthorIp(offset, limit, iCommentAuthor,iCommentAuthorIp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAuthorIp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentDateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentDate := self.Args("comment_date").Time()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentDate(offset, limit, iCommentAuthor,iCommentDate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentDate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentDateGmt(offset, limit, iCommentAuthor,iCommentDateGmt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentDateGmt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentContent(offset, limit, iCommentAuthor,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentKarma(offset, limit, iCommentAuthor,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentApproved(offset, limit, iCommentAuthor,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentAgent(offset, limit, iCommentAuthor,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentType(offset, limit, iCommentAuthor,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndCommentParent(offset, limit, iCommentAuthor,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthor := self.Args("comment_author").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsesByCommentAuthorAndUserId(offset, limit, iCommentAuthor,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAuthorUrl := self.Args("comment_author_url").String()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentAuthorUrl(offset, limit, iCommentAuthorEmail,iCommentAuthorUrl)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentAuthorUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentAuthorIpHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentAuthorIp(offset, limit, iCommentAuthorEmail,iCommentAuthorIp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentAuthorIp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentDateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentDate := self.Args("comment_date").Time()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentDate(offset, limit, iCommentAuthorEmail,iCommentDate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentDate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentDateGmt(offset, limit, iCommentAuthorEmail,iCommentDateGmt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentDateGmt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentContent(offset, limit, iCommentAuthorEmail,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentKarma(offset, limit, iCommentAuthorEmail,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentApproved(offset, limit, iCommentAuthorEmail,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentAgent(offset, limit, iCommentAuthorEmail,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentType(offset, limit, iCommentAuthorEmail,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndCommentParent(offset, limit, iCommentAuthorEmail,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorEmailAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsesByCommentAuthorEmailAndUserId(offset, limit, iCommentAuthorEmail,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorEmailAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentAuthorIpHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentAuthorIp := self.Args("comment_author_ip").String()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentAuthorIp(offset, limit, iCommentAuthorUrl,iCommentAuthorIp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentAuthorIp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentDateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentDate := self.Args("comment_date").Time()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentDate(offset, limit, iCommentAuthorUrl,iCommentDate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentDate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentDateGmt(offset, limit, iCommentAuthorUrl,iCommentDateGmt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentDateGmt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentContent(offset, limit, iCommentAuthorUrl,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentKarma(offset, limit, iCommentAuthorUrl,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentApproved(offset, limit, iCommentAuthorUrl,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentAgent(offset, limit, iCommentAuthorUrl,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentType(offset, limit, iCommentAuthorUrl,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndCommentParent(offset, limit, iCommentAuthorUrl,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorUrlAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsesByCommentAuthorUrlAndUserId(offset, limit, iCommentAuthorUrl,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorUrlAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentDateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentDate := self.Args("comment_date").Time()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentDate(offset, limit, iCommentAuthorIp,iCommentDate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentDate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentDateGmt(offset, limit, iCommentAuthorIp,iCommentDateGmt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentDateGmt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentContent(offset, limit, iCommentAuthorIp,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentKarma(offset, limit, iCommentAuthorIp,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentApproved(offset, limit, iCommentAuthorIp,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentAgent(offset, limit, iCommentAuthorIp,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentType(offset, limit, iCommentAuthorIp,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndCommentParent(offset, limit, iCommentAuthorIp,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAuthorIpAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAuthorIp := self.Args("comment_author_ip").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsesByCommentAuthorIpAndUserId(offset, limit, iCommentAuthorIp,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAuthorIpAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentDateGmt(offset, limit, iCommentDate,iCommentDateGmt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentDateGmt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentContent(offset, limit, iCommentDate,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentKarma(offset, limit, iCommentDate,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentApproved(offset, limit, iCommentDate,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentAgent(offset, limit, iCommentDate,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentType(offset, limit, iCommentDate,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndCommentParent(offset, limit, iCommentDate,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDate := self.Args("comment_date").Time()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsesByCommentDateAndUserId(offset, limit, iCommentDate,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateGmtAndCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentContent := self.Args("comment_content").String()

	if helper.IsHas(iCommentDateGmt) {
		_Comments, _error := model.GetCommentsesByCommentDateGmtAndCommentContent(offset, limit, iCommentDateGmt,iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateGmtAndCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateGmtAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentDateGmt) {
		_Comments, _error := model.GetCommentsesByCommentDateGmtAndCommentKarma(offset, limit, iCommentDateGmt,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateGmtAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateGmtAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentDateGmt) {
		_Comments, _error := model.GetCommentsesByCommentDateGmtAndCommentApproved(offset, limit, iCommentDateGmt,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateGmtAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateGmtAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentDateGmt) {
		_Comments, _error := model.GetCommentsesByCommentDateGmtAndCommentAgent(offset, limit, iCommentDateGmt,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateGmtAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateGmtAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentDateGmt) {
		_Comments, _error := model.GetCommentsesByCommentDateGmtAndCommentType(offset, limit, iCommentDateGmt,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateGmtAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateGmtAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentDateGmt) {
		_Comments, _error := model.GetCommentsesByCommentDateGmtAndCommentParent(offset, limit, iCommentDateGmt,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateGmtAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentDateGmtAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentDateGmt) {
		_Comments, _error := model.GetCommentsesByCommentDateGmtAndUserId(offset, limit, iCommentDateGmt,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentDateGmtAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentContentAndCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentContent := self.Args("comment_content").String()
	iCommentKarma := self.Args("comment_karma").MustInt()

	if helper.IsHas(iCommentContent) {
		_Comments, _error := model.GetCommentsesByCommentContentAndCommentKarma(offset, limit, iCommentContent,iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentContentAndCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentContentAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentContent := self.Args("comment_content").String()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentContent) {
		_Comments, _error := model.GetCommentsesByCommentContentAndCommentApproved(offset, limit, iCommentContent,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentContentAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentContentAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentContent := self.Args("comment_content").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentContent) {
		_Comments, _error := model.GetCommentsesByCommentContentAndCommentAgent(offset, limit, iCommentContent,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentContentAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentContentAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentContent := self.Args("comment_content").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentContent) {
		_Comments, _error := model.GetCommentsesByCommentContentAndCommentType(offset, limit, iCommentContent,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentContentAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentContentAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentContent := self.Args("comment_content").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentContent) {
		_Comments, _error := model.GetCommentsesByCommentContentAndCommentParent(offset, limit, iCommentContent,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentContentAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentContentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentContent := self.Args("comment_content").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentContent) {
		_Comments, _error := model.GetCommentsesByCommentContentAndUserId(offset, limit, iCommentContent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentContentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentKarmaAndCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentApproved := self.Args("comment_approved").String()

	if helper.IsHas(iCommentKarma) {
		_Comments, _error := model.GetCommentsesByCommentKarmaAndCommentApproved(offset, limit, iCommentKarma,iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentKarmaAndCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentKarmaAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentKarma) {
		_Comments, _error := model.GetCommentsesByCommentKarmaAndCommentAgent(offset, limit, iCommentKarma,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentKarmaAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentKarmaAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentKarma) {
		_Comments, _error := model.GetCommentsesByCommentKarmaAndCommentType(offset, limit, iCommentKarma,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentKarmaAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentKarmaAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentKarma) {
		_Comments, _error := model.GetCommentsesByCommentKarmaAndCommentParent(offset, limit, iCommentKarma,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentKarmaAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentKarmaAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentKarma := self.Args("comment_karma").MustInt()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentKarma) {
		_Comments, _error := model.GetCommentsesByCommentKarmaAndUserId(offset, limit, iCommentKarma,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentKarmaAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentApprovedAndCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentAgent := self.Args("comment_agent").String()

	if helper.IsHas(iCommentApproved) {
		_Comments, _error := model.GetCommentsesByCommentApprovedAndCommentAgent(offset, limit, iCommentApproved,iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentApprovedAndCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentApprovedAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentApproved) {
		_Comments, _error := model.GetCommentsesByCommentApprovedAndCommentType(offset, limit, iCommentApproved,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentApprovedAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentApprovedAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentApproved := self.Args("comment_approved").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentApproved) {
		_Comments, _error := model.GetCommentsesByCommentApprovedAndCommentParent(offset, limit, iCommentApproved,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentApprovedAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentApprovedAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentApproved := self.Args("comment_approved").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentApproved) {
		_Comments, _error := model.GetCommentsesByCommentApprovedAndUserId(offset, limit, iCommentApproved,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentApprovedAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAgentAndCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAgent := self.Args("comment_agent").String()
	iCommentType := self.Args("comment_type").String()

	if helper.IsHas(iCommentAgent) {
		_Comments, _error := model.GetCommentsesByCommentAgentAndCommentType(offset, limit, iCommentAgent,iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAgentAndCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAgentAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAgent := self.Args("comment_agent").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentAgent) {
		_Comments, _error := model.GetCommentsesByCommentAgentAndCommentParent(offset, limit, iCommentAgent,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAgentAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentAgentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentAgent := self.Args("comment_agent").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentAgent) {
		_Comments, _error := model.GetCommentsesByCommentAgentAndUserId(offset, limit, iCommentAgent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentAgentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentTypeAndCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iCommentParent := self.Args("comment_parent").MustInt64()

	if helper.IsHas(iCommentType) {
		_Comments, _error := model.GetCommentsesByCommentTypeAndCommentParent(offset, limit, iCommentType,iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentTypeAndCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentTypeAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentType := self.Args("comment_type").String()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentType) {
		_Comments, _error := model.GetCommentsesByCommentTypeAndUserId(offset, limit, iCommentType,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentTypeAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesByCommentParentAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentParent := self.Args("comment_parent").MustInt64()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iCommentParent) {
		_Comments, _error := model.GetCommentsesByCommentParentAndUserId(offset, limit, iCommentParent,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsesByCommentParentAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_Comments, _error := model.GetCommentses(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentses' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentsByCommentIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentId := self.Args("comment_ID").MustInt64()
	if helper.IsHas(iCommentId) {
		_Comments := model.HasCommentsByCommentId(iCommentId)
		var m = map[string]interface{}{}
		m["comments"] = _Comments
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentsByCommentId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentsByCommentPostIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentPostId := self.Args("comment_post_ID").MustInt64()
	if helper.IsHas(iCommentPostId) {
		_Comments := model.HasCommentsByCommentPostId(iCommentPostId)
		var m = map[string]interface{}{}
		m["comments"] = _Comments
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentsByCommentPostId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentsByCommentAuthorHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentAuthor := self.Args("comment_author").String()
	if helper.IsHas(iCommentAuthor) {
		_Comments := model.HasCommentsByCommentAuthor(iCommentAuthor)
		var m = map[string]interface{}{}
		m["comments"] = _Comments
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentsByCommentAuthor's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentsByCommentAuthorEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	if helper.IsHas(iCommentAuthorEmail) {
		_Comments := model.HasCommentsByCommentAuthorEmail(iCommentAuthorEmail)
		var m = map[string]interface{}{}
		m["comments"] = _Comments
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentsByCommentAuthorEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentsByCommentAuthorUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	if helper.IsHas(iCommentAuthorUrl) {
		_Comments := model.HasCommentsByCommentAuthorUrl(iCommentAuthorUrl)
		var m = map[string]interface{}{}
		m["comments"] = _Comments
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentsByCommentAuthorUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentsByCommentAuthorIpHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentAuthorIp := self.Args("comment_author_IP").String()
	if helper.IsHas(iCommentAuthorIp) {
		_Comments := model.HasCommentsByCommentAuthorIp(iCommentAuthorIp)
		var m = map[string]interface{}{}
		m["comments"] = _Comments
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentsByCommentAuthorIp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentsByCommentDateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentDate := self.Args("comment_date").Time()
	if helper.IsHas(iCommentDate) {
		_Comments := model.HasCommentsByCommentDate(iCommentDate)
		var m = map[string]interface{}{}
		m["comments"] = _Comments
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentsByCommentDate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentsByCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	if helper.IsHas(iCommentDateGmt) {
		_Comments := model.HasCommentsByCommentDateGmt(iCommentDateGmt)
		var m = map[string]interface{}{}
		m["comments"] = _Comments
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentsByCommentDateGmt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentsByCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentContent := self.Args("comment_content").String()
	if helper.IsHas(iCommentContent) {
		_Comments := model.HasCommentsByCommentContent(iCommentContent)
		var m = map[string]interface{}{}
		m["comments"] = _Comments
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentsByCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentsByCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentKarma := self.Args("comment_karma").MustInt()
	if helper.IsHas(iCommentKarma) {
		_Comments := model.HasCommentsByCommentKarma(iCommentKarma)
		var m = map[string]interface{}{}
		m["comments"] = _Comments
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentsByCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentsByCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentApproved := self.Args("comment_approved").String()
	if helper.IsHas(iCommentApproved) {
		_Comments := model.HasCommentsByCommentApproved(iCommentApproved)
		var m = map[string]interface{}{}
		m["comments"] = _Comments
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentsByCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentsByCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentAgent := self.Args("comment_agent").String()
	if helper.IsHas(iCommentAgent) {
		_Comments := model.HasCommentsByCommentAgent(iCommentAgent)
		var m = map[string]interface{}{}
		m["comments"] = _Comments
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentsByCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentsByCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentType := self.Args("comment_type").String()
	if helper.IsHas(iCommentType) {
		_Comments := model.HasCommentsByCommentType(iCommentType)
		var m = map[string]interface{}{}
		m["comments"] = _Comments
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentsByCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentsByCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentParent := self.Args("comment_parent").MustInt64()
	if helper.IsHas(iCommentParent) {
		_Comments := model.HasCommentsByCommentParent(iCommentParent)
		var m = map[string]interface{}{}
		m["comments"] = _Comments
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentsByCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentsByUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserId := self.Args("user_id").MustInt64()
	if helper.IsHas(iUserId) {
		_Comments := model.HasCommentsByUserId(iUserId)
		var m = map[string]interface{}{}
		m["comments"] = _Comments
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentsByUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsByCommentIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentId := self.Args("comment_ID").MustInt64()
	if helper.IsHas(iCommentId) {
		_Comments, _error := model.GetCommentsByCommentId(iCommentId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsByCommentId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsByCommentPostIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentPostId := self.Args("comment_post_ID").MustInt64()
	if helper.IsHas(iCommentPostId) {
		_Comments, _error := model.GetCommentsByCommentPostId(iCommentPostId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsByCommentPostId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsByCommentAuthorHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentAuthor := self.Args("comment_author").String()
	if helper.IsHas(iCommentAuthor) {
		_Comments, _error := model.GetCommentsByCommentAuthor(iCommentAuthor)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsByCommentAuthor's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsByCommentAuthorEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentAuthorEmail := self.Args("comment_author_email").String()
	if helper.IsHas(iCommentAuthorEmail) {
		_Comments, _error := model.GetCommentsByCommentAuthorEmail(iCommentAuthorEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsByCommentAuthorEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsByCommentAuthorUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentAuthorUrl := self.Args("comment_author_url").String()
	if helper.IsHas(iCommentAuthorUrl) {
		_Comments, _error := model.GetCommentsByCommentAuthorUrl(iCommentAuthorUrl)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsByCommentAuthorUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsByCommentAuthorIpHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentAuthorIp := self.Args("comment_author_IP").String()
	if helper.IsHas(iCommentAuthorIp) {
		_Comments, _error := model.GetCommentsByCommentAuthorIp(iCommentAuthorIp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsByCommentAuthorIp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsByCommentDateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentDate := self.Args("comment_date").Time()
	if helper.IsHas(iCommentDate) {
		_Comments, _error := model.GetCommentsByCommentDate(iCommentDate)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsByCommentDate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsByCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentDateGmt := self.Args("comment_date_gmt").Time()
	if helper.IsHas(iCommentDateGmt) {
		_Comments, _error := model.GetCommentsByCommentDateGmt(iCommentDateGmt)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsByCommentDateGmt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsByCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentContent := self.Args("comment_content").String()
	if helper.IsHas(iCommentContent) {
		_Comments, _error := model.GetCommentsByCommentContent(iCommentContent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsByCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsByCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentKarma := self.Args("comment_karma").MustInt()
	if helper.IsHas(iCommentKarma) {
		_Comments, _error := model.GetCommentsByCommentKarma(iCommentKarma)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsByCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsByCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentApproved := self.Args("comment_approved").String()
	if helper.IsHas(iCommentApproved) {
		_Comments, _error := model.GetCommentsByCommentApproved(iCommentApproved)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsByCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsByCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentAgent := self.Args("comment_agent").String()
	if helper.IsHas(iCommentAgent) {
		_Comments, _error := model.GetCommentsByCommentAgent(iCommentAgent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsByCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsByCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentType := self.Args("comment_type").String()
	if helper.IsHas(iCommentType) {
		_Comments, _error := model.GetCommentsByCommentType(iCommentType)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsByCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsByCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentParent := self.Args("comment_parent").MustInt64()
	if helper.IsHas(iCommentParent) {
		_Comments, _error := model.GetCommentsByCommentParent(iCommentParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsByCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentsByUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserId := self.Args("user_id").MustInt64()
	if helper.IsHas(iUserId) {
		_Comments, _error := model.GetCommentsByUserId(iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the GetCommentsByUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentsByCommentIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentId_ := self.Args("comment_ID").MustInt64()
	if helper.IsHas(CommentId_) {
		var iComments model.Comments
		self.Bind(&iComments)
		_Comments, _error := model.SetCommentsByCommentId(CommentId_, &iComments)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the SetCommentsByCommentId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentsByCommentPostIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentPostId_ := self.Args("comment_post_ID").MustInt64()
	if helper.IsHas(CommentPostId_) {
		var iComments model.Comments
		self.Bind(&iComments)
		_Comments, _error := model.SetCommentsByCommentPostId(CommentPostId_, &iComments)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the SetCommentsByCommentPostId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentsByCommentAuthorHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentAuthor_ := self.Args("comment_author").String()
	if helper.IsHas(CommentAuthor_) {
		var iComments model.Comments
		self.Bind(&iComments)
		_Comments, _error := model.SetCommentsByCommentAuthor(CommentAuthor_, &iComments)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the SetCommentsByCommentAuthor's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentsByCommentAuthorEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentAuthorEmail_ := self.Args("comment_author_email").String()
	if helper.IsHas(CommentAuthorEmail_) {
		var iComments model.Comments
		self.Bind(&iComments)
		_Comments, _error := model.SetCommentsByCommentAuthorEmail(CommentAuthorEmail_, &iComments)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the SetCommentsByCommentAuthorEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentsByCommentAuthorUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentAuthorUrl_ := self.Args("comment_author_url").String()
	if helper.IsHas(CommentAuthorUrl_) {
		var iComments model.Comments
		self.Bind(&iComments)
		_Comments, _error := model.SetCommentsByCommentAuthorUrl(CommentAuthorUrl_, &iComments)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the SetCommentsByCommentAuthorUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentsByCommentAuthorIpHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentAuthorIp_ := self.Args("comment_author_IP").String()
	if helper.IsHas(CommentAuthorIp_) {
		var iComments model.Comments
		self.Bind(&iComments)
		_Comments, _error := model.SetCommentsByCommentAuthorIp(CommentAuthorIp_, &iComments)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the SetCommentsByCommentAuthorIp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentsByCommentDateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentDate_ := self.Args("comment_date").Time()
	if helper.IsHas(CommentDate_) {
		var iComments model.Comments
		self.Bind(&iComments)
		_Comments, _error := model.SetCommentsByCommentDate(CommentDate_, &iComments)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the SetCommentsByCommentDate's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentsByCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentDateGmt_ := self.Args("comment_date_gmt").Time()
	if helper.IsHas(CommentDateGmt_) {
		var iComments model.Comments
		self.Bind(&iComments)
		_Comments, _error := model.SetCommentsByCommentDateGmt(CommentDateGmt_, &iComments)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the SetCommentsByCommentDateGmt's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentsByCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentContent_ := self.Args("comment_content").String()
	if helper.IsHas(CommentContent_) {
		var iComments model.Comments
		self.Bind(&iComments)
		_Comments, _error := model.SetCommentsByCommentContent(CommentContent_, &iComments)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the SetCommentsByCommentContent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentsByCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentKarma_ := self.Args("comment_karma").MustInt()
	if helper.IsHas(CommentKarma_) {
		var iComments model.Comments
		self.Bind(&iComments)
		_Comments, _error := model.SetCommentsByCommentKarma(CommentKarma_, &iComments)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the SetCommentsByCommentKarma's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentsByCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentApproved_ := self.Args("comment_approved").String()
	if helper.IsHas(CommentApproved_) {
		var iComments model.Comments
		self.Bind(&iComments)
		_Comments, _error := model.SetCommentsByCommentApproved(CommentApproved_, &iComments)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the SetCommentsByCommentApproved's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentsByCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentAgent_ := self.Args("comment_agent").String()
	if helper.IsHas(CommentAgent_) {
		var iComments model.Comments
		self.Bind(&iComments)
		_Comments, _error := model.SetCommentsByCommentAgent(CommentAgent_, &iComments)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the SetCommentsByCommentAgent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentsByCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentType_ := self.Args("comment_type").String()
	if helper.IsHas(CommentType_) {
		var iComments model.Comments
		self.Bind(&iComments)
		_Comments, _error := model.SetCommentsByCommentType(CommentType_, &iComments)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the SetCommentsByCommentType's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentsByCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentParent_ := self.Args("comment_parent").MustInt64()
	if helper.IsHas(CommentParent_) {
		var iComments model.Comments
		self.Bind(&iComments)
		_Comments, _error := model.SetCommentsByCommentParent(CommentParent_, &iComments)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the SetCommentsByCommentParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentsByUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserId_ := self.Args("user_id").MustInt64()
	if helper.IsHas(UserId_) {
		var iComments model.Comments
		self.Bind(&iComments)
		_Comments, _error := model.SetCommentsByUserId(UserId_, &iComments)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Comments)
	}
	herr.Message = "Can't get to the SetCommentsByUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddCommentsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentId_ := self.Args("comment_ID").MustInt64()
	CommentPostId_ := self.Args("comment_post_ID").MustInt64()
	CommentAuthor_ := self.Args("comment_author").String()
	CommentAuthorEmail_ := self.Args("comment_author_email").String()
	CommentAuthorUrl_ := self.Args("comment_author_url").String()
	CommentAuthorIp_ := self.Args("comment_author_IP").String()
	CommentDate_ := self.Args("comment_date").Time()
	CommentDateGmt_ := self.Args("comment_date_gmt").Time()
	CommentContent_ := self.Args("comment_content").String()
	CommentKarma_ := self.Args("comment_karma").MustInt()
	CommentApproved_ := self.Args("comment_approved").String()
	CommentAgent_ := self.Args("comment_agent").String()
	CommentType_ := self.Args("comment_type").String()
	CommentParent_ := self.Args("comment_parent").MustInt64()
	UserId_ := self.Args("user_id").MustInt64()

	if helper.IsHas(CommentId_) {
		_error := model.AddComments(CommentId_,CommentPostId_,CommentAuthor_,CommentAuthorEmail_,CommentAuthorUrl_,CommentAuthorIp_,CommentDate_,CommentDateGmt_,CommentContent_,CommentKarma_,CommentApproved_,CommentAgent_,CommentType_,CommentParent_,UserId_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddComments's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostCommentsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iComments model.Comments
	self.Bind(&iComments)
	_int64, _error := model.PostComments(&iComments)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	} else {
		var m = map[string]interface{}{}
		m["created"] = _int64
		return self.JSON(m, macross.StatusCreated)
	}
	return self.JSON(herr)
}

func PutCommentsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iComments model.Comments
	self.Bind(&iComments)
	_int64, _error := model.PutComments(&iComments)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	} else {
		var m = map[string]interface{}{}
		m["updated"] = _int64
		return self.JSON(m)
	}
	return self.JSON(herr)
}

func PutCommentsByCommentIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentId_ := self.Args("comment_ID").MustInt64()
	var iComments model.Comments
	self.Bind(&iComments)
	_int64, _error := model.PutCommentsByCommentId(CommentId_, &iComments)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentsByCommentPostIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentPostId_ := self.Args("comment_post_ID").MustInt64()
	var iComments model.Comments
	self.Bind(&iComments)
	_int64, _error := model.PutCommentsByCommentPostId(CommentPostId_, &iComments)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentsByCommentAuthorHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentAuthor_ := self.Args("comment_author").String()
	var iComments model.Comments
	self.Bind(&iComments)
	_int64, _error := model.PutCommentsByCommentAuthor(CommentAuthor_, &iComments)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentsByCommentAuthorEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentAuthorEmail_ := self.Args("comment_author_email").String()
	var iComments model.Comments
	self.Bind(&iComments)
	_int64, _error := model.PutCommentsByCommentAuthorEmail(CommentAuthorEmail_, &iComments)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentsByCommentAuthorUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentAuthorUrl_ := self.Args("comment_author_url").String()
	var iComments model.Comments
	self.Bind(&iComments)
	_int64, _error := model.PutCommentsByCommentAuthorUrl(CommentAuthorUrl_, &iComments)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentsByCommentAuthorIpHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentAuthorIp_ := self.Args("comment_author_IP").String()
	var iComments model.Comments
	self.Bind(&iComments)
	_int64, _error := model.PutCommentsByCommentAuthorIp(CommentAuthorIp_, &iComments)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentsByCommentDateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentDate_ := self.Args("comment_date").Time()
	var iComments model.Comments
	self.Bind(&iComments)
	_int64, _error := model.PutCommentsByCommentDate(CommentDate_, &iComments)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentsByCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentDateGmt_ := self.Args("comment_date_gmt").Time()
	var iComments model.Comments
	self.Bind(&iComments)
	_int64, _error := model.PutCommentsByCommentDateGmt(CommentDateGmt_, &iComments)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentsByCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentContent_ := self.Args("comment_content").String()
	var iComments model.Comments
	self.Bind(&iComments)
	_int64, _error := model.PutCommentsByCommentContent(CommentContent_, &iComments)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentsByCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentKarma_ := self.Args("comment_karma").MustInt()
	var iComments model.Comments
	self.Bind(&iComments)
	_int64, _error := model.PutCommentsByCommentKarma(CommentKarma_, &iComments)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentsByCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentApproved_ := self.Args("comment_approved").String()
	var iComments model.Comments
	self.Bind(&iComments)
	_int64, _error := model.PutCommentsByCommentApproved(CommentApproved_, &iComments)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentsByCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentAgent_ := self.Args("comment_agent").String()
	var iComments model.Comments
	self.Bind(&iComments)
	_int64, _error := model.PutCommentsByCommentAgent(CommentAgent_, &iComments)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentsByCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentType_ := self.Args("comment_type").String()
	var iComments model.Comments
	self.Bind(&iComments)
	_int64, _error := model.PutCommentsByCommentType(CommentType_, &iComments)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentsByCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentParent_ := self.Args("comment_parent").MustInt64()
	var iComments model.Comments
	self.Bind(&iComments)
	_int64, _error := model.PutCommentsByCommentParent(CommentParent_, &iComments)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentsByUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserId_ := self.Args("user_id").MustInt64()
	var iComments model.Comments
	self.Bind(&iComments)
	_int64, _error := model.PutCommentsByUserId(UserId_, &iComments)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentsByCommentIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentId_ := self.Args("comment_ID").MustInt64()
	var iComments model.Comments
	self.Bind(&iComments)
	var iMap = helper.StructToMap(iComments)
	_error := model.UpdateCommentsByCommentId(CommentId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentsByCommentPostIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentPostId_ := self.Args("comment_post_ID").MustInt64()
	var iComments model.Comments
	self.Bind(&iComments)
	var iMap = helper.StructToMap(iComments)
	_error := model.UpdateCommentsByCommentPostId(CommentPostId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentsByCommentAuthorHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentAuthor_ := self.Args("comment_author").String()
	var iComments model.Comments
	self.Bind(&iComments)
	var iMap = helper.StructToMap(iComments)
	_error := model.UpdateCommentsByCommentAuthor(CommentAuthor_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentsByCommentAuthorEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentAuthorEmail_ := self.Args("comment_author_email").String()
	var iComments model.Comments
	self.Bind(&iComments)
	var iMap = helper.StructToMap(iComments)
	_error := model.UpdateCommentsByCommentAuthorEmail(CommentAuthorEmail_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentsByCommentAuthorUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentAuthorUrl_ := self.Args("comment_author_url").String()
	var iComments model.Comments
	self.Bind(&iComments)
	var iMap = helper.StructToMap(iComments)
	_error := model.UpdateCommentsByCommentAuthorUrl(CommentAuthorUrl_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentsByCommentAuthorIpHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentAuthorIp_ := self.Args("comment_author_IP").String()
	var iComments model.Comments
	self.Bind(&iComments)
	var iMap = helper.StructToMap(iComments)
	_error := model.UpdateCommentsByCommentAuthorIp(CommentAuthorIp_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentsByCommentDateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentDate_ := self.Args("comment_date").Time()
	var iComments model.Comments
	self.Bind(&iComments)
	var iMap = helper.StructToMap(iComments)
	_error := model.UpdateCommentsByCommentDate(CommentDate_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentsByCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentDateGmt_ := self.Args("comment_date_gmt").Time()
	var iComments model.Comments
	self.Bind(&iComments)
	var iMap = helper.StructToMap(iComments)
	_error := model.UpdateCommentsByCommentDateGmt(CommentDateGmt_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentsByCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentContent_ := self.Args("comment_content").String()
	var iComments model.Comments
	self.Bind(&iComments)
	var iMap = helper.StructToMap(iComments)
	_error := model.UpdateCommentsByCommentContent(CommentContent_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentsByCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentKarma_ := self.Args("comment_karma").MustInt()
	var iComments model.Comments
	self.Bind(&iComments)
	var iMap = helper.StructToMap(iComments)
	_error := model.UpdateCommentsByCommentKarma(CommentKarma_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentsByCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentApproved_ := self.Args("comment_approved").String()
	var iComments model.Comments
	self.Bind(&iComments)
	var iMap = helper.StructToMap(iComments)
	_error := model.UpdateCommentsByCommentApproved(CommentApproved_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentsByCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentAgent_ := self.Args("comment_agent").String()
	var iComments model.Comments
	self.Bind(&iComments)
	var iMap = helper.StructToMap(iComments)
	_error := model.UpdateCommentsByCommentAgent(CommentAgent_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentsByCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentType_ := self.Args("comment_type").String()
	var iComments model.Comments
	self.Bind(&iComments)
	var iMap = helper.StructToMap(iComments)
	_error := model.UpdateCommentsByCommentType(CommentType_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentsByCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentParent_ := self.Args("comment_parent").MustInt64()
	var iComments model.Comments
	self.Bind(&iComments)
	var iMap = helper.StructToMap(iComments)
	_error := model.UpdateCommentsByCommentParent(CommentParent_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentsByUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserId_ := self.Args("user_id").MustInt64()
	var iComments model.Comments
	self.Bind(&iComments)
	var iMap = helper.StructToMap(iComments)
	_error := model.UpdateCommentsByUserId(UserId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentId_ := self.Args("comment_ID").MustInt64()
	_int64, _error := model.DeleteComments(CommentId_)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	} else {
		var m = map[string]interface{}{}
		m["deleted"] = _int64
		return self.JSON(m)
	}
	return self.JSON(herr)
}

func DeleteCommentsByCommentIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentId_ := self.Args("comment_ID").MustInt64()
	_error := model.DeleteCommentsByCommentId(CommentId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentsByCommentPostIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentPostId_ := self.Args("comment_post_ID").MustInt64()
	_error := model.DeleteCommentsByCommentPostId(CommentPostId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentsByCommentAuthorHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentAuthor_ := self.Args("comment_author").String()
	_error := model.DeleteCommentsByCommentAuthor(CommentAuthor_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentsByCommentAuthorEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentAuthorEmail_ := self.Args("comment_author_email").String()
	_error := model.DeleteCommentsByCommentAuthorEmail(CommentAuthorEmail_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentsByCommentAuthorUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentAuthorUrl_ := self.Args("comment_author_url").String()
	_error := model.DeleteCommentsByCommentAuthorUrl(CommentAuthorUrl_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentsByCommentAuthorIpHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentAuthorIp_ := self.Args("comment_author_IP").String()
	_error := model.DeleteCommentsByCommentAuthorIp(CommentAuthorIp_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentsByCommentDateHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentDate_ := self.Args("comment_date").Time()
	_error := model.DeleteCommentsByCommentDate(CommentDate_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentsByCommentDateGmtHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentDateGmt_ := self.Args("comment_date_gmt").Time()
	_error := model.DeleteCommentsByCommentDateGmt(CommentDateGmt_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentsByCommentContentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentContent_ := self.Args("comment_content").String()
	_error := model.DeleteCommentsByCommentContent(CommentContent_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentsByCommentKarmaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentKarma_ := self.Args("comment_karma").MustInt()
	_error := model.DeleteCommentsByCommentKarma(CommentKarma_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentsByCommentApprovedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentApproved_ := self.Args("comment_approved").String()
	_error := model.DeleteCommentsByCommentApproved(CommentApproved_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentsByCommentAgentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentAgent_ := self.Args("comment_agent").String()
	_error := model.DeleteCommentsByCommentAgent(CommentAgent_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentsByCommentTypeHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentType_ := self.Args("comment_type").String()
	_error := model.DeleteCommentsByCommentType(CommentType_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentsByCommentParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentParent_ := self.Args("comment_parent").MustInt64()
	_error := model.DeleteCommentsByCommentParent(CommentParent_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentsByUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserId_ := self.Args("user_id").MustInt64()
	_error := model.DeleteCommentsByUserId(UserId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
