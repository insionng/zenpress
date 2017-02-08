package handler

import (
	"github.com/insionng/zenpress/helper"
	"github.com/insionng/zenpress/model"
	"github.com/insionng/macross"
)

func GetLinksesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetLinksesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["linkssCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetLinksesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksCountViaLinkIdHandler(self *macross.Context) error {
	LinkId_ := self.Args("link_id").MustInt64()
	_int64 := model.GetLinksCountViaLinkId(LinkId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["linksCount"] = 0
	}
	m["linksCount"] = _int64
	return self.JSON(m)
}

func GetLinksCountViaLinkUrlHandler(self *macross.Context) error {
	LinkUrl_ := self.Args("link_url").String()
	_int64 := model.GetLinksCountViaLinkUrl(LinkUrl_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["linksCount"] = 0
	}
	m["linksCount"] = _int64
	return self.JSON(m)
}

func GetLinksCountViaLinkNameHandler(self *macross.Context) error {
	LinkName_ := self.Args("link_name").String()
	_int64 := model.GetLinksCountViaLinkName(LinkName_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["linksCount"] = 0
	}
	m["linksCount"] = _int64
	return self.JSON(m)
}

func GetLinksCountViaLinkImageHandler(self *macross.Context) error {
	LinkImage_ := self.Args("link_image").String()
	_int64 := model.GetLinksCountViaLinkImage(LinkImage_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["linksCount"] = 0
	}
	m["linksCount"] = _int64
	return self.JSON(m)
}

func GetLinksCountViaLinkTargetHandler(self *macross.Context) error {
	LinkTarget_ := self.Args("link_target").String()
	_int64 := model.GetLinksCountViaLinkTarget(LinkTarget_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["linksCount"] = 0
	}
	m["linksCount"] = _int64
	return self.JSON(m)
}

func GetLinksCountViaLinkDescriptionHandler(self *macross.Context) error {
	LinkDescription_ := self.Args("link_description").String()
	_int64 := model.GetLinksCountViaLinkDescription(LinkDescription_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["linksCount"] = 0
	}
	m["linksCount"] = _int64
	return self.JSON(m)
}

func GetLinksCountViaLinkVisibleHandler(self *macross.Context) error {
	LinkVisible_ := self.Args("link_visible").String()
	_int64 := model.GetLinksCountViaLinkVisible(LinkVisible_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["linksCount"] = 0
	}
	m["linksCount"] = _int64
	return self.JSON(m)
}

func GetLinksCountViaLinkOwnerHandler(self *macross.Context) error {
	LinkOwner_ := self.Args("link_owner").MustInt64()
	_int64 := model.GetLinksCountViaLinkOwner(LinkOwner_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["linksCount"] = 0
	}
	m["linksCount"] = _int64
	return self.JSON(m)
}

func GetLinksCountViaLinkRatingHandler(self *macross.Context) error {
	LinkRating_ := self.Args("link_rating").MustInt()
	_int64 := model.GetLinksCountViaLinkRating(LinkRating_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["linksCount"] = 0
	}
	m["linksCount"] = _int64
	return self.JSON(m)
}

func GetLinksCountViaLinkUpdatedHandler(self *macross.Context) error {
	LinkUpdated_ := self.Args("link_updated").Time()
	_int64 := model.GetLinksCountViaLinkUpdated(LinkUpdated_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["linksCount"] = 0
	}
	m["linksCount"] = _int64
	return self.JSON(m)
}

func GetLinksCountViaLinkRelHandler(self *macross.Context) error {
	LinkRel_ := self.Args("link_rel").String()
	_int64 := model.GetLinksCountViaLinkRel(LinkRel_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["linksCount"] = 0
	}
	m["linksCount"] = _int64
	return self.JSON(m)
}

func GetLinksCountViaLinkNotesHandler(self *macross.Context) error {
	LinkNotes_ := self.Args("link_notes").String()
	_int64 := model.GetLinksCountViaLinkNotes(LinkNotes_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["linksCount"] = 0
	}
	m["linksCount"] = _int64
	return self.JSON(m)
}

func GetLinksCountViaLinkRssHandler(self *macross.Context) error {
	LinkRss_ := self.Args("link_rss").String()
	_int64 := model.GetLinksCountViaLinkRss(LinkRss_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["linksCount"] = 0
	}
	m["linksCount"] = _int64
	return self.JSON(m)
}

func GetLinksesViaLinkIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLinkId := self.Args("link_id").MustInt64()
	if (offset > 0) && helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesViaLinkId(offset, limit, iLinkId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesViaLinkId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesViaLinkUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLinkUrl := self.Args("link_url").String()
	if (offset > 0) && helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesViaLinkUrl(offset, limit, iLinkUrl, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesViaLinkUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesViaLinkNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLinkName := self.Args("link_name").String()
	if (offset > 0) && helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesViaLinkName(offset, limit, iLinkName, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesViaLinkName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesViaLinkImageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLinkImage := self.Args("link_image").String()
	if (offset > 0) && helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesViaLinkImage(offset, limit, iLinkImage, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesViaLinkImage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesViaLinkTargetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLinkTarget := self.Args("link_target").String()
	if (offset > 0) && helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesViaLinkTarget(offset, limit, iLinkTarget, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesViaLinkTarget's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesViaLinkDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLinkDescription := self.Args("link_description").String()
	if (offset > 0) && helper.IsHas(iLinkDescription) {
		_Links, _error := model.GetLinksesViaLinkDescription(offset, limit, iLinkDescription, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesViaLinkDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesViaLinkVisibleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLinkVisible := self.Args("link_visible").String()
	if (offset > 0) && helper.IsHas(iLinkVisible) {
		_Links, _error := model.GetLinksesViaLinkVisible(offset, limit, iLinkVisible, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesViaLinkVisible's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesViaLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLinkOwner := self.Args("link_owner").MustInt64()
	if (offset > 0) && helper.IsHas(iLinkOwner) {
		_Links, _error := model.GetLinksesViaLinkOwner(offset, limit, iLinkOwner, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesViaLinkOwner's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesViaLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLinkRating := self.Args("link_rating").MustInt()
	if (offset > 0) && helper.IsHas(iLinkRating) {
		_Links, _error := model.GetLinksesViaLinkRating(offset, limit, iLinkRating, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesViaLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesViaLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLinkUpdated := self.Args("link_updated").Time()
	if (offset > 0) && helper.IsHas(iLinkUpdated) {
		_Links, _error := model.GetLinksesViaLinkUpdated(offset, limit, iLinkUpdated, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesViaLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesViaLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLinkRel := self.Args("link_rel").String()
	if (offset > 0) && helper.IsHas(iLinkRel) {
		_Links, _error := model.GetLinksesViaLinkRel(offset, limit, iLinkRel, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesViaLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesViaLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLinkNotes := self.Args("link_notes").String()
	if (offset > 0) && helper.IsHas(iLinkNotes) {
		_Links, _error := model.GetLinksesViaLinkNotes(offset, limit, iLinkNotes, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesViaLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesViaLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLinkRss := self.Args("link_rss").String()
	if (offset > 0) && helper.IsHas(iLinkRss) {
		_Links, _error := model.GetLinksesViaLinkRss(offset, limit, iLinkRss, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesViaLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkUrlAndLinkNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkUrl := self.Args("link_url").String()
	iLinkName := self.Args("link_name").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkUrlAndLinkName(offset, limit, iLinkId,iLinkUrl,iLinkName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkUrlAndLinkName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkUrlAndLinkImageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkUrl := self.Args("link_url").String()
	iLinkImage := self.Args("link_image").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkUrlAndLinkImage(offset, limit, iLinkId,iLinkUrl,iLinkImage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkUrlAndLinkImage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkUrlAndLinkTargetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkUrl := self.Args("link_url").String()
	iLinkTarget := self.Args("link_target").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkUrlAndLinkTarget(offset, limit, iLinkId,iLinkUrl,iLinkTarget)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkUrlAndLinkTarget's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkUrlAndLinkDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkUrl := self.Args("link_url").String()
	iLinkDescription := self.Args("link_description").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkUrlAndLinkDescription(offset, limit, iLinkId,iLinkUrl,iLinkDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkUrlAndLinkDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkUrlAndLinkVisibleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkUrl := self.Args("link_url").String()
	iLinkVisible := self.Args("link_visible").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkUrlAndLinkVisible(offset, limit, iLinkId,iLinkUrl,iLinkVisible)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkUrlAndLinkVisible's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkUrlAndLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkUrl := self.Args("link_url").String()
	iLinkOwner := self.Args("link_owner").MustInt64()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkUrlAndLinkOwner(offset, limit, iLinkId,iLinkUrl,iLinkOwner)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkUrlAndLinkOwner's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkUrlAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkUrl := self.Args("link_url").String()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkUrlAndLinkRating(offset, limit, iLinkId,iLinkUrl,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkUrlAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkUrlAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkUrl := self.Args("link_url").String()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkUrlAndLinkUpdated(offset, limit, iLinkId,iLinkUrl,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkUrlAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkUrlAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkUrl := self.Args("link_url").String()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkUrlAndLinkRel(offset, limit, iLinkId,iLinkUrl,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkUrlAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkUrlAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkUrl := self.Args("link_url").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkUrlAndLinkNotes(offset, limit, iLinkId,iLinkUrl,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkUrlAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkUrlAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkUrl := self.Args("link_url").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkUrlAndLinkRss(offset, limit, iLinkId,iLinkUrl,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkUrlAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkNameAndLinkImageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkName := self.Args("link_name").String()
	iLinkImage := self.Args("link_image").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkNameAndLinkImage(offset, limit, iLinkId,iLinkName,iLinkImage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkNameAndLinkImage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkNameAndLinkTargetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkName := self.Args("link_name").String()
	iLinkTarget := self.Args("link_target").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkNameAndLinkTarget(offset, limit, iLinkId,iLinkName,iLinkTarget)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkNameAndLinkTarget's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkNameAndLinkDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkName := self.Args("link_name").String()
	iLinkDescription := self.Args("link_description").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkNameAndLinkDescription(offset, limit, iLinkId,iLinkName,iLinkDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkNameAndLinkDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkNameAndLinkVisibleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkName := self.Args("link_name").String()
	iLinkVisible := self.Args("link_visible").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkNameAndLinkVisible(offset, limit, iLinkId,iLinkName,iLinkVisible)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkNameAndLinkVisible's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkNameAndLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkName := self.Args("link_name").String()
	iLinkOwner := self.Args("link_owner").MustInt64()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkNameAndLinkOwner(offset, limit, iLinkId,iLinkName,iLinkOwner)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkNameAndLinkOwner's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkNameAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkName := self.Args("link_name").String()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkNameAndLinkRating(offset, limit, iLinkId,iLinkName,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkNameAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkNameAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkName := self.Args("link_name").String()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkNameAndLinkUpdated(offset, limit, iLinkId,iLinkName,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkNameAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkNameAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkName := self.Args("link_name").String()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkNameAndLinkRel(offset, limit, iLinkId,iLinkName,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkNameAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkNameAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkName := self.Args("link_name").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkNameAndLinkNotes(offset, limit, iLinkId,iLinkName,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkNameAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkNameAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkName := self.Args("link_name").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkNameAndLinkRss(offset, limit, iLinkId,iLinkName,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkNameAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkImageAndLinkTargetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkImage := self.Args("link_image").String()
	iLinkTarget := self.Args("link_target").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkImageAndLinkTarget(offset, limit, iLinkId,iLinkImage,iLinkTarget)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkImageAndLinkTarget's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkImageAndLinkDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkImage := self.Args("link_image").String()
	iLinkDescription := self.Args("link_description").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkImageAndLinkDescription(offset, limit, iLinkId,iLinkImage,iLinkDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkImageAndLinkDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkImageAndLinkVisibleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkImage := self.Args("link_image").String()
	iLinkVisible := self.Args("link_visible").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkImageAndLinkVisible(offset, limit, iLinkId,iLinkImage,iLinkVisible)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkImageAndLinkVisible's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkImageAndLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkImage := self.Args("link_image").String()
	iLinkOwner := self.Args("link_owner").MustInt64()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkImageAndLinkOwner(offset, limit, iLinkId,iLinkImage,iLinkOwner)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkImageAndLinkOwner's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkImageAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkImage := self.Args("link_image").String()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkImageAndLinkRating(offset, limit, iLinkId,iLinkImage,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkImageAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkImageAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkImage := self.Args("link_image").String()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkImageAndLinkUpdated(offset, limit, iLinkId,iLinkImage,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkImageAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkImageAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkImage := self.Args("link_image").String()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkImageAndLinkRel(offset, limit, iLinkId,iLinkImage,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkImageAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkImageAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkImage := self.Args("link_image").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkImageAndLinkNotes(offset, limit, iLinkId,iLinkImage,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkImageAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkImageAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkImage := self.Args("link_image").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkImageAndLinkRss(offset, limit, iLinkId,iLinkImage,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkImageAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkTargetAndLinkDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkTarget := self.Args("link_target").String()
	iLinkDescription := self.Args("link_description").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkTargetAndLinkDescription(offset, limit, iLinkId,iLinkTarget,iLinkDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkTargetAndLinkDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkTargetAndLinkVisibleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkTarget := self.Args("link_target").String()
	iLinkVisible := self.Args("link_visible").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkTargetAndLinkVisible(offset, limit, iLinkId,iLinkTarget,iLinkVisible)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkTargetAndLinkVisible's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkTargetAndLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkTarget := self.Args("link_target").String()
	iLinkOwner := self.Args("link_owner").MustInt64()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkTargetAndLinkOwner(offset, limit, iLinkId,iLinkTarget,iLinkOwner)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkTargetAndLinkOwner's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkTargetAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkTarget := self.Args("link_target").String()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkTargetAndLinkRating(offset, limit, iLinkId,iLinkTarget,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkTargetAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkTargetAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkTarget := self.Args("link_target").String()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkTargetAndLinkUpdated(offset, limit, iLinkId,iLinkTarget,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkTargetAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkTargetAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkTarget := self.Args("link_target").String()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkTargetAndLinkRel(offset, limit, iLinkId,iLinkTarget,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkTargetAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkTargetAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkTarget := self.Args("link_target").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkTargetAndLinkNotes(offset, limit, iLinkId,iLinkTarget,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkTargetAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkTargetAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkTarget := self.Args("link_target").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkTargetAndLinkRss(offset, limit, iLinkId,iLinkTarget,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkTargetAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkDescriptionAndLinkVisibleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkDescription := self.Args("link_description").String()
	iLinkVisible := self.Args("link_visible").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkDescriptionAndLinkVisible(offset, limit, iLinkId,iLinkDescription,iLinkVisible)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkDescriptionAndLinkVisible's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkDescriptionAndLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkDescription := self.Args("link_description").String()
	iLinkOwner := self.Args("link_owner").MustInt64()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkDescriptionAndLinkOwner(offset, limit, iLinkId,iLinkDescription,iLinkOwner)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkDescriptionAndLinkOwner's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkDescriptionAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkDescription := self.Args("link_description").String()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkDescriptionAndLinkRating(offset, limit, iLinkId,iLinkDescription,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkDescriptionAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkDescriptionAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkDescription := self.Args("link_description").String()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkDescriptionAndLinkUpdated(offset, limit, iLinkId,iLinkDescription,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkDescriptionAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkDescriptionAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkDescription := self.Args("link_description").String()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkDescriptionAndLinkRel(offset, limit, iLinkId,iLinkDescription,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkDescriptionAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkDescriptionAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkDescription := self.Args("link_description").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkDescriptionAndLinkNotes(offset, limit, iLinkId,iLinkDescription,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkDescriptionAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkDescriptionAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkDescription := self.Args("link_description").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkDescriptionAndLinkRss(offset, limit, iLinkId,iLinkDescription,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkDescriptionAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkVisibleAndLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkVisible := self.Args("link_visible").String()
	iLinkOwner := self.Args("link_owner").MustInt64()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkVisibleAndLinkOwner(offset, limit, iLinkId,iLinkVisible,iLinkOwner)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkVisibleAndLinkOwner's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkVisibleAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkVisible := self.Args("link_visible").String()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkVisibleAndLinkRating(offset, limit, iLinkId,iLinkVisible,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkVisibleAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkVisibleAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkVisible := self.Args("link_visible").String()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkVisibleAndLinkUpdated(offset, limit, iLinkId,iLinkVisible,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkVisibleAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkVisibleAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkVisible := self.Args("link_visible").String()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkVisibleAndLinkRel(offset, limit, iLinkId,iLinkVisible,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkVisibleAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkVisibleAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkVisible := self.Args("link_visible").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkVisibleAndLinkNotes(offset, limit, iLinkId,iLinkVisible,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkVisibleAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkVisibleAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkVisible := self.Args("link_visible").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkVisibleAndLinkRss(offset, limit, iLinkId,iLinkVisible,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkVisibleAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkOwnerAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkOwnerAndLinkRating(offset, limit, iLinkId,iLinkOwner,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkOwnerAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkOwnerAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkOwnerAndLinkUpdated(offset, limit, iLinkId,iLinkOwner,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkOwnerAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkOwnerAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkOwnerAndLinkRel(offset, limit, iLinkId,iLinkOwner,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkOwnerAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkOwnerAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkOwnerAndLinkNotes(offset, limit, iLinkId,iLinkOwner,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkOwnerAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkOwnerAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkOwnerAndLinkRss(offset, limit, iLinkId,iLinkOwner,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkOwnerAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkRatingAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkRatingAndLinkUpdated(offset, limit, iLinkId,iLinkRating,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkRatingAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkRatingAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkRatingAndLinkRel(offset, limit, iLinkId,iLinkRating,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkRatingAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkRatingAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkRatingAndLinkNotes(offset, limit, iLinkId,iLinkRating,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkRatingAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkRatingAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkRatingAndLinkRss(offset, limit, iLinkId,iLinkRating,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkRatingAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkUpdatedAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkUpdatedAndLinkRel(offset, limit, iLinkId,iLinkUpdated,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkUpdatedAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkUpdatedAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkUpdatedAndLinkNotes(offset, limit, iLinkId,iLinkUpdated,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkUpdatedAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkUpdatedAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkUpdatedAndLinkRss(offset, limit, iLinkId,iLinkUpdated,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkUpdatedAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkRelAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkRel := self.Args("link_rel").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkRelAndLinkNotes(offset, limit, iLinkId,iLinkRel,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkRelAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkRelAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkRel := self.Args("link_rel").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkRelAndLinkRss(offset, limit, iLinkId,iLinkRel,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkRelAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkNotesAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkNotes := self.Args("link_notes").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkNotesAndLinkRss(offset, limit, iLinkId,iLinkNotes,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkNotesAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkNameAndLinkImageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkName := self.Args("link_name").String()
	iLinkImage := self.Args("link_image").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkNameAndLinkImage(offset, limit, iLinkUrl,iLinkName,iLinkImage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkNameAndLinkImage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkNameAndLinkTargetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkName := self.Args("link_name").String()
	iLinkTarget := self.Args("link_target").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkNameAndLinkTarget(offset, limit, iLinkUrl,iLinkName,iLinkTarget)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkNameAndLinkTarget's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkNameAndLinkDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkName := self.Args("link_name").String()
	iLinkDescription := self.Args("link_description").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkNameAndLinkDescription(offset, limit, iLinkUrl,iLinkName,iLinkDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkNameAndLinkDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkNameAndLinkVisibleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkName := self.Args("link_name").String()
	iLinkVisible := self.Args("link_visible").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkNameAndLinkVisible(offset, limit, iLinkUrl,iLinkName,iLinkVisible)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkNameAndLinkVisible's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkNameAndLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkName := self.Args("link_name").String()
	iLinkOwner := self.Args("link_owner").MustInt64()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkNameAndLinkOwner(offset, limit, iLinkUrl,iLinkName,iLinkOwner)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkNameAndLinkOwner's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkNameAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkName := self.Args("link_name").String()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkNameAndLinkRating(offset, limit, iLinkUrl,iLinkName,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkNameAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkNameAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkName := self.Args("link_name").String()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkNameAndLinkUpdated(offset, limit, iLinkUrl,iLinkName,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkNameAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkNameAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkName := self.Args("link_name").String()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkNameAndLinkRel(offset, limit, iLinkUrl,iLinkName,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkNameAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkNameAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkName := self.Args("link_name").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkNameAndLinkNotes(offset, limit, iLinkUrl,iLinkName,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkNameAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkNameAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkName := self.Args("link_name").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkNameAndLinkRss(offset, limit, iLinkUrl,iLinkName,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkNameAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkImageAndLinkTargetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkImage := self.Args("link_image").String()
	iLinkTarget := self.Args("link_target").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkImageAndLinkTarget(offset, limit, iLinkUrl,iLinkImage,iLinkTarget)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkImageAndLinkTarget's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkImageAndLinkDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkImage := self.Args("link_image").String()
	iLinkDescription := self.Args("link_description").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkImageAndLinkDescription(offset, limit, iLinkUrl,iLinkImage,iLinkDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkImageAndLinkDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkImageAndLinkVisibleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkImage := self.Args("link_image").String()
	iLinkVisible := self.Args("link_visible").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkImageAndLinkVisible(offset, limit, iLinkUrl,iLinkImage,iLinkVisible)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkImageAndLinkVisible's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkImageAndLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkImage := self.Args("link_image").String()
	iLinkOwner := self.Args("link_owner").MustInt64()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkImageAndLinkOwner(offset, limit, iLinkUrl,iLinkImage,iLinkOwner)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkImageAndLinkOwner's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkImageAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkImage := self.Args("link_image").String()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkImageAndLinkRating(offset, limit, iLinkUrl,iLinkImage,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkImageAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkImageAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkImage := self.Args("link_image").String()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkImageAndLinkUpdated(offset, limit, iLinkUrl,iLinkImage,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkImageAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkImageAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkImage := self.Args("link_image").String()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkImageAndLinkRel(offset, limit, iLinkUrl,iLinkImage,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkImageAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkImageAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkImage := self.Args("link_image").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkImageAndLinkNotes(offset, limit, iLinkUrl,iLinkImage,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkImageAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkImageAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkImage := self.Args("link_image").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkImageAndLinkRss(offset, limit, iLinkUrl,iLinkImage,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkImageAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkTargetAndLinkDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkTarget := self.Args("link_target").String()
	iLinkDescription := self.Args("link_description").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkTargetAndLinkDescription(offset, limit, iLinkUrl,iLinkTarget,iLinkDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkTargetAndLinkDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkTargetAndLinkVisibleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkTarget := self.Args("link_target").String()
	iLinkVisible := self.Args("link_visible").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkTargetAndLinkVisible(offset, limit, iLinkUrl,iLinkTarget,iLinkVisible)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkTargetAndLinkVisible's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkTargetAndLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkTarget := self.Args("link_target").String()
	iLinkOwner := self.Args("link_owner").MustInt64()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkTargetAndLinkOwner(offset, limit, iLinkUrl,iLinkTarget,iLinkOwner)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkTargetAndLinkOwner's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkTargetAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkTarget := self.Args("link_target").String()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkTargetAndLinkRating(offset, limit, iLinkUrl,iLinkTarget,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkTargetAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkTargetAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkTarget := self.Args("link_target").String()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkTargetAndLinkUpdated(offset, limit, iLinkUrl,iLinkTarget,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkTargetAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkTargetAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkTarget := self.Args("link_target").String()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkTargetAndLinkRel(offset, limit, iLinkUrl,iLinkTarget,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkTargetAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkTargetAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkTarget := self.Args("link_target").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkTargetAndLinkNotes(offset, limit, iLinkUrl,iLinkTarget,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkTargetAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkTargetAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkTarget := self.Args("link_target").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkTargetAndLinkRss(offset, limit, iLinkUrl,iLinkTarget,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkTargetAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkDescriptionAndLinkVisibleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkDescription := self.Args("link_description").String()
	iLinkVisible := self.Args("link_visible").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkDescriptionAndLinkVisible(offset, limit, iLinkUrl,iLinkDescription,iLinkVisible)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkDescriptionAndLinkVisible's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkDescriptionAndLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkDescription := self.Args("link_description").String()
	iLinkOwner := self.Args("link_owner").MustInt64()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkDescriptionAndLinkOwner(offset, limit, iLinkUrl,iLinkDescription,iLinkOwner)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkDescriptionAndLinkOwner's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkDescriptionAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkDescription := self.Args("link_description").String()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkDescriptionAndLinkRating(offset, limit, iLinkUrl,iLinkDescription,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkDescriptionAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkDescriptionAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkDescription := self.Args("link_description").String()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkDescriptionAndLinkUpdated(offset, limit, iLinkUrl,iLinkDescription,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkDescriptionAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkDescriptionAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkDescription := self.Args("link_description").String()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkDescriptionAndLinkRel(offset, limit, iLinkUrl,iLinkDescription,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkDescriptionAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkDescriptionAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkDescription := self.Args("link_description").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkDescriptionAndLinkNotes(offset, limit, iLinkUrl,iLinkDescription,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkDescriptionAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkDescriptionAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkDescription := self.Args("link_description").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkDescriptionAndLinkRss(offset, limit, iLinkUrl,iLinkDescription,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkDescriptionAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkVisibleAndLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkVisible := self.Args("link_visible").String()
	iLinkOwner := self.Args("link_owner").MustInt64()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkVisibleAndLinkOwner(offset, limit, iLinkUrl,iLinkVisible,iLinkOwner)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkVisibleAndLinkOwner's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkVisibleAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkVisible := self.Args("link_visible").String()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkVisibleAndLinkRating(offset, limit, iLinkUrl,iLinkVisible,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkVisibleAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkVisibleAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkVisible := self.Args("link_visible").String()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkVisibleAndLinkUpdated(offset, limit, iLinkUrl,iLinkVisible,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkVisibleAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkVisibleAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkVisible := self.Args("link_visible").String()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkVisibleAndLinkRel(offset, limit, iLinkUrl,iLinkVisible,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkVisibleAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkVisibleAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkVisible := self.Args("link_visible").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkVisibleAndLinkNotes(offset, limit, iLinkUrl,iLinkVisible,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkVisibleAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkVisibleAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkVisible := self.Args("link_visible").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkVisibleAndLinkRss(offset, limit, iLinkUrl,iLinkVisible,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkVisibleAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkOwnerAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkOwnerAndLinkRating(offset, limit, iLinkUrl,iLinkOwner,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkOwnerAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkOwnerAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkOwnerAndLinkUpdated(offset, limit, iLinkUrl,iLinkOwner,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkOwnerAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkOwnerAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkOwnerAndLinkRel(offset, limit, iLinkUrl,iLinkOwner,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkOwnerAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkOwnerAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkOwnerAndLinkNotes(offset, limit, iLinkUrl,iLinkOwner,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkOwnerAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkOwnerAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkOwnerAndLinkRss(offset, limit, iLinkUrl,iLinkOwner,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkOwnerAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkRatingAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkRatingAndLinkUpdated(offset, limit, iLinkUrl,iLinkRating,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkRatingAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkRatingAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkRatingAndLinkRel(offset, limit, iLinkUrl,iLinkRating,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkRatingAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkRatingAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkRatingAndLinkNotes(offset, limit, iLinkUrl,iLinkRating,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkRatingAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkRatingAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkRatingAndLinkRss(offset, limit, iLinkUrl,iLinkRating,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkRatingAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkUpdatedAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkUpdatedAndLinkRel(offset, limit, iLinkUrl,iLinkUpdated,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkUpdatedAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkUpdatedAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkUpdatedAndLinkNotes(offset, limit, iLinkUrl,iLinkUpdated,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkUpdatedAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkUpdatedAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkUpdatedAndLinkRss(offset, limit, iLinkUrl,iLinkUpdated,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkUpdatedAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkRelAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkRel := self.Args("link_rel").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkRelAndLinkNotes(offset, limit, iLinkUrl,iLinkRel,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkRelAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkRelAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkRel := self.Args("link_rel").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkRelAndLinkRss(offset, limit, iLinkUrl,iLinkRel,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkRelAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkNotesAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkNotes := self.Args("link_notes").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkNotesAndLinkRss(offset, limit, iLinkUrl,iLinkNotes,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkNotesAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkImageAndLinkTargetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkImage := self.Args("link_image").String()
	iLinkTarget := self.Args("link_target").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkImageAndLinkTarget(offset, limit, iLinkName,iLinkImage,iLinkTarget)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkImageAndLinkTarget's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkImageAndLinkDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkImage := self.Args("link_image").String()
	iLinkDescription := self.Args("link_description").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkImageAndLinkDescription(offset, limit, iLinkName,iLinkImage,iLinkDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkImageAndLinkDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkImageAndLinkVisibleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkImage := self.Args("link_image").String()
	iLinkVisible := self.Args("link_visible").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkImageAndLinkVisible(offset, limit, iLinkName,iLinkImage,iLinkVisible)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkImageAndLinkVisible's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkImageAndLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkImage := self.Args("link_image").String()
	iLinkOwner := self.Args("link_owner").MustInt64()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkImageAndLinkOwner(offset, limit, iLinkName,iLinkImage,iLinkOwner)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkImageAndLinkOwner's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkImageAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkImage := self.Args("link_image").String()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkImageAndLinkRating(offset, limit, iLinkName,iLinkImage,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkImageAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkImageAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkImage := self.Args("link_image").String()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkImageAndLinkUpdated(offset, limit, iLinkName,iLinkImage,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkImageAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkImageAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkImage := self.Args("link_image").String()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkImageAndLinkRel(offset, limit, iLinkName,iLinkImage,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkImageAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkImageAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkImage := self.Args("link_image").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkImageAndLinkNotes(offset, limit, iLinkName,iLinkImage,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkImageAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkImageAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkImage := self.Args("link_image").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkImageAndLinkRss(offset, limit, iLinkName,iLinkImage,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkImageAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkTargetAndLinkDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkTarget := self.Args("link_target").String()
	iLinkDescription := self.Args("link_description").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkTargetAndLinkDescription(offset, limit, iLinkName,iLinkTarget,iLinkDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkTargetAndLinkDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkTargetAndLinkVisibleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkTarget := self.Args("link_target").String()
	iLinkVisible := self.Args("link_visible").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkTargetAndLinkVisible(offset, limit, iLinkName,iLinkTarget,iLinkVisible)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkTargetAndLinkVisible's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkTargetAndLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkTarget := self.Args("link_target").String()
	iLinkOwner := self.Args("link_owner").MustInt64()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkTargetAndLinkOwner(offset, limit, iLinkName,iLinkTarget,iLinkOwner)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkTargetAndLinkOwner's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkTargetAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkTarget := self.Args("link_target").String()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkTargetAndLinkRating(offset, limit, iLinkName,iLinkTarget,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkTargetAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkTargetAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkTarget := self.Args("link_target").String()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkTargetAndLinkUpdated(offset, limit, iLinkName,iLinkTarget,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkTargetAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkTargetAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkTarget := self.Args("link_target").String()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkTargetAndLinkRel(offset, limit, iLinkName,iLinkTarget,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkTargetAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkTargetAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkTarget := self.Args("link_target").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkTargetAndLinkNotes(offset, limit, iLinkName,iLinkTarget,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkTargetAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkTargetAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkTarget := self.Args("link_target").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkTargetAndLinkRss(offset, limit, iLinkName,iLinkTarget,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkTargetAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkDescriptionAndLinkVisibleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkDescription := self.Args("link_description").String()
	iLinkVisible := self.Args("link_visible").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkDescriptionAndLinkVisible(offset, limit, iLinkName,iLinkDescription,iLinkVisible)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkDescriptionAndLinkVisible's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkDescriptionAndLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkDescription := self.Args("link_description").String()
	iLinkOwner := self.Args("link_owner").MustInt64()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkDescriptionAndLinkOwner(offset, limit, iLinkName,iLinkDescription,iLinkOwner)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkDescriptionAndLinkOwner's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkDescriptionAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkDescription := self.Args("link_description").String()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkDescriptionAndLinkRating(offset, limit, iLinkName,iLinkDescription,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkDescriptionAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkDescriptionAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkDescription := self.Args("link_description").String()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkDescriptionAndLinkUpdated(offset, limit, iLinkName,iLinkDescription,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkDescriptionAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkDescriptionAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkDescription := self.Args("link_description").String()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkDescriptionAndLinkRel(offset, limit, iLinkName,iLinkDescription,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkDescriptionAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkDescriptionAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkDescription := self.Args("link_description").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkDescriptionAndLinkNotes(offset, limit, iLinkName,iLinkDescription,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkDescriptionAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkDescriptionAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkDescription := self.Args("link_description").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkDescriptionAndLinkRss(offset, limit, iLinkName,iLinkDescription,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkDescriptionAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkVisibleAndLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkVisible := self.Args("link_visible").String()
	iLinkOwner := self.Args("link_owner").MustInt64()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkVisibleAndLinkOwner(offset, limit, iLinkName,iLinkVisible,iLinkOwner)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkVisibleAndLinkOwner's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkVisibleAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkVisible := self.Args("link_visible").String()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkVisibleAndLinkRating(offset, limit, iLinkName,iLinkVisible,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkVisibleAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkVisibleAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkVisible := self.Args("link_visible").String()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkVisibleAndLinkUpdated(offset, limit, iLinkName,iLinkVisible,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkVisibleAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkVisibleAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkVisible := self.Args("link_visible").String()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkVisibleAndLinkRel(offset, limit, iLinkName,iLinkVisible,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkVisibleAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkVisibleAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkVisible := self.Args("link_visible").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkVisibleAndLinkNotes(offset, limit, iLinkName,iLinkVisible,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkVisibleAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkVisibleAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkVisible := self.Args("link_visible").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkVisibleAndLinkRss(offset, limit, iLinkName,iLinkVisible,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkVisibleAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkOwnerAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkOwnerAndLinkRating(offset, limit, iLinkName,iLinkOwner,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkOwnerAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkOwnerAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkOwnerAndLinkUpdated(offset, limit, iLinkName,iLinkOwner,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkOwnerAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkOwnerAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkOwnerAndLinkRel(offset, limit, iLinkName,iLinkOwner,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkOwnerAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkOwnerAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkOwnerAndLinkNotes(offset, limit, iLinkName,iLinkOwner,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkOwnerAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkOwnerAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkOwnerAndLinkRss(offset, limit, iLinkName,iLinkOwner,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkOwnerAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkRatingAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkRatingAndLinkUpdated(offset, limit, iLinkName,iLinkRating,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkRatingAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkRatingAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkRatingAndLinkRel(offset, limit, iLinkName,iLinkRating,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkRatingAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkRatingAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkRatingAndLinkNotes(offset, limit, iLinkName,iLinkRating,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkRatingAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkRatingAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkRatingAndLinkRss(offset, limit, iLinkName,iLinkRating,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkRatingAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkUpdatedAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkUpdatedAndLinkRel(offset, limit, iLinkName,iLinkUpdated,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkUpdatedAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkUpdatedAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkUpdatedAndLinkNotes(offset, limit, iLinkName,iLinkUpdated,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkUpdatedAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkUpdatedAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkUpdatedAndLinkRss(offset, limit, iLinkName,iLinkUpdated,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkUpdatedAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkRelAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkRel := self.Args("link_rel").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkRelAndLinkNotes(offset, limit, iLinkName,iLinkRel,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkRelAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkRelAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkRel := self.Args("link_rel").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkRelAndLinkRss(offset, limit, iLinkName,iLinkRel,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkRelAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkNotesAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkNotes := self.Args("link_notes").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkNotesAndLinkRss(offset, limit, iLinkName,iLinkNotes,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkNotesAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkTargetAndLinkDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkTarget := self.Args("link_target").String()
	iLinkDescription := self.Args("link_description").String()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkTargetAndLinkDescription(offset, limit, iLinkImage,iLinkTarget,iLinkDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkTargetAndLinkDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkTargetAndLinkVisibleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkTarget := self.Args("link_target").String()
	iLinkVisible := self.Args("link_visible").String()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkTargetAndLinkVisible(offset, limit, iLinkImage,iLinkTarget,iLinkVisible)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkTargetAndLinkVisible's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkTargetAndLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkTarget := self.Args("link_target").String()
	iLinkOwner := self.Args("link_owner").MustInt64()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkTargetAndLinkOwner(offset, limit, iLinkImage,iLinkTarget,iLinkOwner)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkTargetAndLinkOwner's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkTargetAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkTarget := self.Args("link_target").String()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkTargetAndLinkRating(offset, limit, iLinkImage,iLinkTarget,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkTargetAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkTargetAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkTarget := self.Args("link_target").String()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkTargetAndLinkUpdated(offset, limit, iLinkImage,iLinkTarget,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkTargetAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkTargetAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkTarget := self.Args("link_target").String()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkTargetAndLinkRel(offset, limit, iLinkImage,iLinkTarget,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkTargetAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkTargetAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkTarget := self.Args("link_target").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkTargetAndLinkNotes(offset, limit, iLinkImage,iLinkTarget,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkTargetAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkTargetAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkTarget := self.Args("link_target").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkTargetAndLinkRss(offset, limit, iLinkImage,iLinkTarget,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkTargetAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkDescriptionAndLinkVisibleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkDescription := self.Args("link_description").String()
	iLinkVisible := self.Args("link_visible").String()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkDescriptionAndLinkVisible(offset, limit, iLinkImage,iLinkDescription,iLinkVisible)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkDescriptionAndLinkVisible's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkDescriptionAndLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkDescription := self.Args("link_description").String()
	iLinkOwner := self.Args("link_owner").MustInt64()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkDescriptionAndLinkOwner(offset, limit, iLinkImage,iLinkDescription,iLinkOwner)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkDescriptionAndLinkOwner's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkDescriptionAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkDescription := self.Args("link_description").String()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkDescriptionAndLinkRating(offset, limit, iLinkImage,iLinkDescription,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkDescriptionAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkDescriptionAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkDescription := self.Args("link_description").String()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkDescriptionAndLinkUpdated(offset, limit, iLinkImage,iLinkDescription,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkDescriptionAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkDescriptionAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkDescription := self.Args("link_description").String()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkDescriptionAndLinkRel(offset, limit, iLinkImage,iLinkDescription,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkDescriptionAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkDescriptionAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkDescription := self.Args("link_description").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkDescriptionAndLinkNotes(offset, limit, iLinkImage,iLinkDescription,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkDescriptionAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkDescriptionAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkDescription := self.Args("link_description").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkDescriptionAndLinkRss(offset, limit, iLinkImage,iLinkDescription,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkDescriptionAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkVisibleAndLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkVisible := self.Args("link_visible").String()
	iLinkOwner := self.Args("link_owner").MustInt64()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkVisibleAndLinkOwner(offset, limit, iLinkImage,iLinkVisible,iLinkOwner)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkVisibleAndLinkOwner's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkVisibleAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkVisible := self.Args("link_visible").String()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkVisibleAndLinkRating(offset, limit, iLinkImage,iLinkVisible,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkVisibleAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkVisibleAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkVisible := self.Args("link_visible").String()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkVisibleAndLinkUpdated(offset, limit, iLinkImage,iLinkVisible,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkVisibleAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkVisibleAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkVisible := self.Args("link_visible").String()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkVisibleAndLinkRel(offset, limit, iLinkImage,iLinkVisible,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkVisibleAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkVisibleAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkVisible := self.Args("link_visible").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkVisibleAndLinkNotes(offset, limit, iLinkImage,iLinkVisible,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkVisibleAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkVisibleAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkVisible := self.Args("link_visible").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkVisibleAndLinkRss(offset, limit, iLinkImage,iLinkVisible,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkVisibleAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkOwnerAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkOwnerAndLinkRating(offset, limit, iLinkImage,iLinkOwner,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkOwnerAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkOwnerAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkOwnerAndLinkUpdated(offset, limit, iLinkImage,iLinkOwner,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkOwnerAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkOwnerAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkOwnerAndLinkRel(offset, limit, iLinkImage,iLinkOwner,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkOwnerAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkOwnerAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkOwnerAndLinkNotes(offset, limit, iLinkImage,iLinkOwner,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkOwnerAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkOwnerAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkOwnerAndLinkRss(offset, limit, iLinkImage,iLinkOwner,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkOwnerAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkRatingAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkRatingAndLinkUpdated(offset, limit, iLinkImage,iLinkRating,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkRatingAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkRatingAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkRatingAndLinkRel(offset, limit, iLinkImage,iLinkRating,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkRatingAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkRatingAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkRatingAndLinkNotes(offset, limit, iLinkImage,iLinkRating,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkRatingAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkRatingAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkRatingAndLinkRss(offset, limit, iLinkImage,iLinkRating,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkRatingAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkUpdatedAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkUpdatedAndLinkRel(offset, limit, iLinkImage,iLinkUpdated,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkUpdatedAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkUpdatedAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkUpdatedAndLinkNotes(offset, limit, iLinkImage,iLinkUpdated,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkUpdatedAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkUpdatedAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkUpdatedAndLinkRss(offset, limit, iLinkImage,iLinkUpdated,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkUpdatedAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkRelAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkRel := self.Args("link_rel").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkRelAndLinkNotes(offset, limit, iLinkImage,iLinkRel,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkRelAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkRelAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkRel := self.Args("link_rel").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkRelAndLinkRss(offset, limit, iLinkImage,iLinkRel,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkRelAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkNotesAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkNotes := self.Args("link_notes").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkNotesAndLinkRss(offset, limit, iLinkImage,iLinkNotes,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkNotesAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkDescriptionAndLinkVisibleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkDescription := self.Args("link_description").String()
	iLinkVisible := self.Args("link_visible").String()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkDescriptionAndLinkVisible(offset, limit, iLinkTarget,iLinkDescription,iLinkVisible)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkDescriptionAndLinkVisible's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkDescriptionAndLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkDescription := self.Args("link_description").String()
	iLinkOwner := self.Args("link_owner").MustInt64()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkDescriptionAndLinkOwner(offset, limit, iLinkTarget,iLinkDescription,iLinkOwner)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkDescriptionAndLinkOwner's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkDescriptionAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkDescription := self.Args("link_description").String()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkDescriptionAndLinkRating(offset, limit, iLinkTarget,iLinkDescription,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkDescriptionAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkDescriptionAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkDescription := self.Args("link_description").String()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkDescriptionAndLinkUpdated(offset, limit, iLinkTarget,iLinkDescription,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkDescriptionAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkDescriptionAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkDescription := self.Args("link_description").String()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkDescriptionAndLinkRel(offset, limit, iLinkTarget,iLinkDescription,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkDescriptionAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkDescriptionAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkDescription := self.Args("link_description").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkDescriptionAndLinkNotes(offset, limit, iLinkTarget,iLinkDescription,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkDescriptionAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkDescriptionAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkDescription := self.Args("link_description").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkDescriptionAndLinkRss(offset, limit, iLinkTarget,iLinkDescription,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkDescriptionAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkVisibleAndLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkVisible := self.Args("link_visible").String()
	iLinkOwner := self.Args("link_owner").MustInt64()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkVisibleAndLinkOwner(offset, limit, iLinkTarget,iLinkVisible,iLinkOwner)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkVisibleAndLinkOwner's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkVisibleAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkVisible := self.Args("link_visible").String()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkVisibleAndLinkRating(offset, limit, iLinkTarget,iLinkVisible,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkVisibleAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkVisibleAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkVisible := self.Args("link_visible").String()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkVisibleAndLinkUpdated(offset, limit, iLinkTarget,iLinkVisible,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkVisibleAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkVisibleAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkVisible := self.Args("link_visible").String()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkVisibleAndLinkRel(offset, limit, iLinkTarget,iLinkVisible,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkVisibleAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkVisibleAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkVisible := self.Args("link_visible").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkVisibleAndLinkNotes(offset, limit, iLinkTarget,iLinkVisible,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkVisibleAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkVisibleAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkVisible := self.Args("link_visible").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkVisibleAndLinkRss(offset, limit, iLinkTarget,iLinkVisible,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkVisibleAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkOwnerAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkOwnerAndLinkRating(offset, limit, iLinkTarget,iLinkOwner,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkOwnerAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkOwnerAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkOwnerAndLinkUpdated(offset, limit, iLinkTarget,iLinkOwner,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkOwnerAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkOwnerAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkOwnerAndLinkRel(offset, limit, iLinkTarget,iLinkOwner,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkOwnerAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkOwnerAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkOwnerAndLinkNotes(offset, limit, iLinkTarget,iLinkOwner,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkOwnerAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkOwnerAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkOwnerAndLinkRss(offset, limit, iLinkTarget,iLinkOwner,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkOwnerAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkRatingAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkRatingAndLinkUpdated(offset, limit, iLinkTarget,iLinkRating,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkRatingAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkRatingAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkRatingAndLinkRel(offset, limit, iLinkTarget,iLinkRating,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkRatingAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkRatingAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkRatingAndLinkNotes(offset, limit, iLinkTarget,iLinkRating,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkRatingAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkRatingAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkRatingAndLinkRss(offset, limit, iLinkTarget,iLinkRating,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkRatingAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkUpdatedAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkUpdatedAndLinkRel(offset, limit, iLinkTarget,iLinkUpdated,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkUpdatedAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkUpdatedAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkUpdatedAndLinkNotes(offset, limit, iLinkTarget,iLinkUpdated,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkUpdatedAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkUpdatedAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkUpdatedAndLinkRss(offset, limit, iLinkTarget,iLinkUpdated,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkUpdatedAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkRelAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkRel := self.Args("link_rel").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkRelAndLinkNotes(offset, limit, iLinkTarget,iLinkRel,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkRelAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkRelAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkRel := self.Args("link_rel").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkRelAndLinkRss(offset, limit, iLinkTarget,iLinkRel,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkRelAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkNotesAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkNotes := self.Args("link_notes").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkNotesAndLinkRss(offset, limit, iLinkTarget,iLinkNotes,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkNotesAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkDescriptionAndLinkVisibleAndLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkDescription := self.Args("link_description").String()
	iLinkVisible := self.Args("link_visible").String()
	iLinkOwner := self.Args("link_owner").MustInt64()

	if helper.IsHas(iLinkDescription) {
		_Links, _error := model.GetLinksesByLinkDescriptionAndLinkVisibleAndLinkOwner(offset, limit, iLinkDescription,iLinkVisible,iLinkOwner)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkDescriptionAndLinkVisibleAndLinkOwner's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkDescriptionAndLinkVisibleAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkDescription := self.Args("link_description").String()
	iLinkVisible := self.Args("link_visible").String()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkDescription) {
		_Links, _error := model.GetLinksesByLinkDescriptionAndLinkVisibleAndLinkRating(offset, limit, iLinkDescription,iLinkVisible,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkDescriptionAndLinkVisibleAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkDescriptionAndLinkVisibleAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkDescription := self.Args("link_description").String()
	iLinkVisible := self.Args("link_visible").String()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkDescription) {
		_Links, _error := model.GetLinksesByLinkDescriptionAndLinkVisibleAndLinkUpdated(offset, limit, iLinkDescription,iLinkVisible,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkDescriptionAndLinkVisibleAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkDescriptionAndLinkVisibleAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkDescription := self.Args("link_description").String()
	iLinkVisible := self.Args("link_visible").String()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkDescription) {
		_Links, _error := model.GetLinksesByLinkDescriptionAndLinkVisibleAndLinkRel(offset, limit, iLinkDescription,iLinkVisible,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkDescriptionAndLinkVisibleAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkDescriptionAndLinkVisibleAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkDescription := self.Args("link_description").String()
	iLinkVisible := self.Args("link_visible").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkDescription) {
		_Links, _error := model.GetLinksesByLinkDescriptionAndLinkVisibleAndLinkNotes(offset, limit, iLinkDescription,iLinkVisible,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkDescriptionAndLinkVisibleAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkDescriptionAndLinkVisibleAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkDescription := self.Args("link_description").String()
	iLinkVisible := self.Args("link_visible").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkDescription) {
		_Links, _error := model.GetLinksesByLinkDescriptionAndLinkVisibleAndLinkRss(offset, limit, iLinkDescription,iLinkVisible,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkDescriptionAndLinkVisibleAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkDescriptionAndLinkOwnerAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkDescription := self.Args("link_description").String()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkDescription) {
		_Links, _error := model.GetLinksesByLinkDescriptionAndLinkOwnerAndLinkRating(offset, limit, iLinkDescription,iLinkOwner,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkDescriptionAndLinkOwnerAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkDescriptionAndLinkOwnerAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkDescription := self.Args("link_description").String()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkDescription) {
		_Links, _error := model.GetLinksesByLinkDescriptionAndLinkOwnerAndLinkUpdated(offset, limit, iLinkDescription,iLinkOwner,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkDescriptionAndLinkOwnerAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkDescriptionAndLinkOwnerAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkDescription := self.Args("link_description").String()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkDescription) {
		_Links, _error := model.GetLinksesByLinkDescriptionAndLinkOwnerAndLinkRel(offset, limit, iLinkDescription,iLinkOwner,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkDescriptionAndLinkOwnerAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkDescriptionAndLinkOwnerAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkDescription := self.Args("link_description").String()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkDescription) {
		_Links, _error := model.GetLinksesByLinkDescriptionAndLinkOwnerAndLinkNotes(offset, limit, iLinkDescription,iLinkOwner,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkDescriptionAndLinkOwnerAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkDescriptionAndLinkOwnerAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkDescription := self.Args("link_description").String()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkDescription) {
		_Links, _error := model.GetLinksesByLinkDescriptionAndLinkOwnerAndLinkRss(offset, limit, iLinkDescription,iLinkOwner,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkDescriptionAndLinkOwnerAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkDescriptionAndLinkRatingAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkDescription := self.Args("link_description").String()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkDescription) {
		_Links, _error := model.GetLinksesByLinkDescriptionAndLinkRatingAndLinkUpdated(offset, limit, iLinkDescription,iLinkRating,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkDescriptionAndLinkRatingAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkDescriptionAndLinkRatingAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkDescription := self.Args("link_description").String()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkDescription) {
		_Links, _error := model.GetLinksesByLinkDescriptionAndLinkRatingAndLinkRel(offset, limit, iLinkDescription,iLinkRating,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkDescriptionAndLinkRatingAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkDescriptionAndLinkRatingAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkDescription := self.Args("link_description").String()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkDescription) {
		_Links, _error := model.GetLinksesByLinkDescriptionAndLinkRatingAndLinkNotes(offset, limit, iLinkDescription,iLinkRating,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkDescriptionAndLinkRatingAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkDescriptionAndLinkRatingAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkDescription := self.Args("link_description").String()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkDescription) {
		_Links, _error := model.GetLinksesByLinkDescriptionAndLinkRatingAndLinkRss(offset, limit, iLinkDescription,iLinkRating,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkDescriptionAndLinkRatingAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkDescriptionAndLinkUpdatedAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkDescription := self.Args("link_description").String()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkDescription) {
		_Links, _error := model.GetLinksesByLinkDescriptionAndLinkUpdatedAndLinkRel(offset, limit, iLinkDescription,iLinkUpdated,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkDescriptionAndLinkUpdatedAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkDescriptionAndLinkUpdatedAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkDescription := self.Args("link_description").String()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkDescription) {
		_Links, _error := model.GetLinksesByLinkDescriptionAndLinkUpdatedAndLinkNotes(offset, limit, iLinkDescription,iLinkUpdated,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkDescriptionAndLinkUpdatedAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkDescriptionAndLinkUpdatedAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkDescription := self.Args("link_description").String()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkDescription) {
		_Links, _error := model.GetLinksesByLinkDescriptionAndLinkUpdatedAndLinkRss(offset, limit, iLinkDescription,iLinkUpdated,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkDescriptionAndLinkUpdatedAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkDescriptionAndLinkRelAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkDescription := self.Args("link_description").String()
	iLinkRel := self.Args("link_rel").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkDescription) {
		_Links, _error := model.GetLinksesByLinkDescriptionAndLinkRelAndLinkNotes(offset, limit, iLinkDescription,iLinkRel,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkDescriptionAndLinkRelAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkDescriptionAndLinkRelAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkDescription := self.Args("link_description").String()
	iLinkRel := self.Args("link_rel").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkDescription) {
		_Links, _error := model.GetLinksesByLinkDescriptionAndLinkRelAndLinkRss(offset, limit, iLinkDescription,iLinkRel,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkDescriptionAndLinkRelAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkDescriptionAndLinkNotesAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkDescription := self.Args("link_description").String()
	iLinkNotes := self.Args("link_notes").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkDescription) {
		_Links, _error := model.GetLinksesByLinkDescriptionAndLinkNotesAndLinkRss(offset, limit, iLinkDescription,iLinkNotes,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkDescriptionAndLinkNotesAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkVisibleAndLinkOwnerAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkVisible := self.Args("link_visible").String()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkVisible) {
		_Links, _error := model.GetLinksesByLinkVisibleAndLinkOwnerAndLinkRating(offset, limit, iLinkVisible,iLinkOwner,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkVisibleAndLinkOwnerAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkVisibleAndLinkOwnerAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkVisible := self.Args("link_visible").String()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkVisible) {
		_Links, _error := model.GetLinksesByLinkVisibleAndLinkOwnerAndLinkUpdated(offset, limit, iLinkVisible,iLinkOwner,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkVisibleAndLinkOwnerAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkVisibleAndLinkOwnerAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkVisible := self.Args("link_visible").String()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkVisible) {
		_Links, _error := model.GetLinksesByLinkVisibleAndLinkOwnerAndLinkRel(offset, limit, iLinkVisible,iLinkOwner,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkVisibleAndLinkOwnerAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkVisibleAndLinkOwnerAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkVisible := self.Args("link_visible").String()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkVisible) {
		_Links, _error := model.GetLinksesByLinkVisibleAndLinkOwnerAndLinkNotes(offset, limit, iLinkVisible,iLinkOwner,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkVisibleAndLinkOwnerAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkVisibleAndLinkOwnerAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkVisible := self.Args("link_visible").String()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkVisible) {
		_Links, _error := model.GetLinksesByLinkVisibleAndLinkOwnerAndLinkRss(offset, limit, iLinkVisible,iLinkOwner,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkVisibleAndLinkOwnerAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkVisibleAndLinkRatingAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkVisible := self.Args("link_visible").String()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkVisible) {
		_Links, _error := model.GetLinksesByLinkVisibleAndLinkRatingAndLinkUpdated(offset, limit, iLinkVisible,iLinkRating,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkVisibleAndLinkRatingAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkVisibleAndLinkRatingAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkVisible := self.Args("link_visible").String()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkVisible) {
		_Links, _error := model.GetLinksesByLinkVisibleAndLinkRatingAndLinkRel(offset, limit, iLinkVisible,iLinkRating,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkVisibleAndLinkRatingAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkVisibleAndLinkRatingAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkVisible := self.Args("link_visible").String()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkVisible) {
		_Links, _error := model.GetLinksesByLinkVisibleAndLinkRatingAndLinkNotes(offset, limit, iLinkVisible,iLinkRating,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkVisibleAndLinkRatingAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkVisibleAndLinkRatingAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkVisible := self.Args("link_visible").String()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkVisible) {
		_Links, _error := model.GetLinksesByLinkVisibleAndLinkRatingAndLinkRss(offset, limit, iLinkVisible,iLinkRating,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkVisibleAndLinkRatingAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkVisibleAndLinkUpdatedAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkVisible := self.Args("link_visible").String()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkVisible) {
		_Links, _error := model.GetLinksesByLinkVisibleAndLinkUpdatedAndLinkRel(offset, limit, iLinkVisible,iLinkUpdated,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkVisibleAndLinkUpdatedAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkVisibleAndLinkUpdatedAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkVisible := self.Args("link_visible").String()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkVisible) {
		_Links, _error := model.GetLinksesByLinkVisibleAndLinkUpdatedAndLinkNotes(offset, limit, iLinkVisible,iLinkUpdated,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkVisibleAndLinkUpdatedAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkVisibleAndLinkUpdatedAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkVisible := self.Args("link_visible").String()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkVisible) {
		_Links, _error := model.GetLinksesByLinkVisibleAndLinkUpdatedAndLinkRss(offset, limit, iLinkVisible,iLinkUpdated,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkVisibleAndLinkUpdatedAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkVisibleAndLinkRelAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkVisible := self.Args("link_visible").String()
	iLinkRel := self.Args("link_rel").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkVisible) {
		_Links, _error := model.GetLinksesByLinkVisibleAndLinkRelAndLinkNotes(offset, limit, iLinkVisible,iLinkRel,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkVisibleAndLinkRelAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkVisibleAndLinkRelAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkVisible := self.Args("link_visible").String()
	iLinkRel := self.Args("link_rel").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkVisible) {
		_Links, _error := model.GetLinksesByLinkVisibleAndLinkRelAndLinkRss(offset, limit, iLinkVisible,iLinkRel,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkVisibleAndLinkRelAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkVisibleAndLinkNotesAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkVisible := self.Args("link_visible").String()
	iLinkNotes := self.Args("link_notes").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkVisible) {
		_Links, _error := model.GetLinksesByLinkVisibleAndLinkNotesAndLinkRss(offset, limit, iLinkVisible,iLinkNotes,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkVisibleAndLinkNotesAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkOwnerAndLinkRatingAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkOwner) {
		_Links, _error := model.GetLinksesByLinkOwnerAndLinkRatingAndLinkUpdated(offset, limit, iLinkOwner,iLinkRating,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkOwnerAndLinkRatingAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkOwnerAndLinkRatingAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkOwner) {
		_Links, _error := model.GetLinksesByLinkOwnerAndLinkRatingAndLinkRel(offset, limit, iLinkOwner,iLinkRating,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkOwnerAndLinkRatingAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkOwnerAndLinkRatingAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkOwner) {
		_Links, _error := model.GetLinksesByLinkOwnerAndLinkRatingAndLinkNotes(offset, limit, iLinkOwner,iLinkRating,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkOwnerAndLinkRatingAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkOwnerAndLinkRatingAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkOwner) {
		_Links, _error := model.GetLinksesByLinkOwnerAndLinkRatingAndLinkRss(offset, limit, iLinkOwner,iLinkRating,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkOwnerAndLinkRatingAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkOwnerAndLinkUpdatedAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkOwner) {
		_Links, _error := model.GetLinksesByLinkOwnerAndLinkUpdatedAndLinkRel(offset, limit, iLinkOwner,iLinkUpdated,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkOwnerAndLinkUpdatedAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkOwnerAndLinkUpdatedAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkOwner) {
		_Links, _error := model.GetLinksesByLinkOwnerAndLinkUpdatedAndLinkNotes(offset, limit, iLinkOwner,iLinkUpdated,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkOwnerAndLinkUpdatedAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkOwnerAndLinkUpdatedAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkOwner) {
		_Links, _error := model.GetLinksesByLinkOwnerAndLinkUpdatedAndLinkRss(offset, limit, iLinkOwner,iLinkUpdated,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkOwnerAndLinkUpdatedAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkOwnerAndLinkRelAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkRel := self.Args("link_rel").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkOwner) {
		_Links, _error := model.GetLinksesByLinkOwnerAndLinkRelAndLinkNotes(offset, limit, iLinkOwner,iLinkRel,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkOwnerAndLinkRelAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkOwnerAndLinkRelAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkRel := self.Args("link_rel").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkOwner) {
		_Links, _error := model.GetLinksesByLinkOwnerAndLinkRelAndLinkRss(offset, limit, iLinkOwner,iLinkRel,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkOwnerAndLinkRelAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkOwnerAndLinkNotesAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkNotes := self.Args("link_notes").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkOwner) {
		_Links, _error := model.GetLinksesByLinkOwnerAndLinkNotesAndLinkRss(offset, limit, iLinkOwner,iLinkNotes,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkOwnerAndLinkNotesAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkRatingAndLinkUpdatedAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkRating) {
		_Links, _error := model.GetLinksesByLinkRatingAndLinkUpdatedAndLinkRel(offset, limit, iLinkRating,iLinkUpdated,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkRatingAndLinkUpdatedAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkRatingAndLinkUpdatedAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkRating) {
		_Links, _error := model.GetLinksesByLinkRatingAndLinkUpdatedAndLinkNotes(offset, limit, iLinkRating,iLinkUpdated,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkRatingAndLinkUpdatedAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkRatingAndLinkUpdatedAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkRating) {
		_Links, _error := model.GetLinksesByLinkRatingAndLinkUpdatedAndLinkRss(offset, limit, iLinkRating,iLinkUpdated,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkRatingAndLinkUpdatedAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkRatingAndLinkRelAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkRel := self.Args("link_rel").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkRating) {
		_Links, _error := model.GetLinksesByLinkRatingAndLinkRelAndLinkNotes(offset, limit, iLinkRating,iLinkRel,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkRatingAndLinkRelAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkRatingAndLinkRelAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkRel := self.Args("link_rel").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkRating) {
		_Links, _error := model.GetLinksesByLinkRatingAndLinkRelAndLinkRss(offset, limit, iLinkRating,iLinkRel,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkRatingAndLinkRelAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkRatingAndLinkNotesAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkNotes := self.Args("link_notes").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkRating) {
		_Links, _error := model.GetLinksesByLinkRatingAndLinkNotesAndLinkRss(offset, limit, iLinkRating,iLinkNotes,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkRatingAndLinkNotesAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUpdatedAndLinkRelAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkRel := self.Args("link_rel").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkUpdated) {
		_Links, _error := model.GetLinksesByLinkUpdatedAndLinkRelAndLinkNotes(offset, limit, iLinkUpdated,iLinkRel,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUpdatedAndLinkRelAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUpdatedAndLinkRelAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkRel := self.Args("link_rel").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkUpdated) {
		_Links, _error := model.GetLinksesByLinkUpdatedAndLinkRelAndLinkRss(offset, limit, iLinkUpdated,iLinkRel,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUpdatedAndLinkRelAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUpdatedAndLinkNotesAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkNotes := self.Args("link_notes").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkUpdated) {
		_Links, _error := model.GetLinksesByLinkUpdatedAndLinkNotesAndLinkRss(offset, limit, iLinkUpdated,iLinkNotes,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUpdatedAndLinkNotesAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkRelAndLinkNotesAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkRel := self.Args("link_rel").String()
	iLinkNotes := self.Args("link_notes").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkRel) {
		_Links, _error := model.GetLinksesByLinkRelAndLinkNotesAndLinkRss(offset, limit, iLinkRel,iLinkNotes,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkRelAndLinkNotesAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkUrl := self.Args("link_url").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkUrl(offset, limit, iLinkId,iLinkUrl)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkName := self.Args("link_name").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkName(offset, limit, iLinkId,iLinkName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkImageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkImage := self.Args("link_image").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkImage(offset, limit, iLinkId,iLinkImage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkImage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkTargetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkTarget := self.Args("link_target").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkTarget(offset, limit, iLinkId,iLinkTarget)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkTarget's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkDescription := self.Args("link_description").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkDescription(offset, limit, iLinkId,iLinkDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkVisibleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkVisible := self.Args("link_visible").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkVisible(offset, limit, iLinkId,iLinkVisible)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkVisible's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkOwner := self.Args("link_owner").MustInt64()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkOwner(offset, limit, iLinkId,iLinkOwner)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkOwner's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkRating(offset, limit, iLinkId,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkUpdated(offset, limit, iLinkId,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkRel(offset, limit, iLinkId,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkNotes(offset, limit, iLinkId,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkIdAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkId := self.Args("link_id").MustInt64()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksesByLinkIdAndLinkRss(offset, limit, iLinkId,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkIdAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkName := self.Args("link_name").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkName(offset, limit, iLinkUrl,iLinkName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkImageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkImage := self.Args("link_image").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkImage(offset, limit, iLinkUrl,iLinkImage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkImage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkTargetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkTarget := self.Args("link_target").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkTarget(offset, limit, iLinkUrl,iLinkTarget)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkTarget's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkDescription := self.Args("link_description").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkDescription(offset, limit, iLinkUrl,iLinkDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkVisibleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkVisible := self.Args("link_visible").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkVisible(offset, limit, iLinkUrl,iLinkVisible)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkVisible's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkOwner := self.Args("link_owner").MustInt64()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkOwner(offset, limit, iLinkUrl,iLinkOwner)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkOwner's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkRating(offset, limit, iLinkUrl,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkUpdated(offset, limit, iLinkUrl,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkRel(offset, limit, iLinkUrl,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkNotes(offset, limit, iLinkUrl,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUrlAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUrl := self.Args("link_url").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksesByLinkUrlAndLinkRss(offset, limit, iLinkUrl,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUrlAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkImageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkImage := self.Args("link_image").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkImage(offset, limit, iLinkName,iLinkImage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkImage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkTargetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkTarget := self.Args("link_target").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkTarget(offset, limit, iLinkName,iLinkTarget)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkTarget's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkDescription := self.Args("link_description").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkDescription(offset, limit, iLinkName,iLinkDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkVisibleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkVisible := self.Args("link_visible").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkVisible(offset, limit, iLinkName,iLinkVisible)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkVisible's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkOwner := self.Args("link_owner").MustInt64()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkOwner(offset, limit, iLinkName,iLinkOwner)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkOwner's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkRating(offset, limit, iLinkName,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkUpdated(offset, limit, iLinkName,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkRel(offset, limit, iLinkName,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkNotes(offset, limit, iLinkName,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNameAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkName := self.Args("link_name").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksesByLinkNameAndLinkRss(offset, limit, iLinkName,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNameAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkTargetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkTarget := self.Args("link_target").String()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkTarget(offset, limit, iLinkImage,iLinkTarget)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkTarget's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkDescription := self.Args("link_description").String()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkDescription(offset, limit, iLinkImage,iLinkDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkVisibleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkVisible := self.Args("link_visible").String()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkVisible(offset, limit, iLinkImage,iLinkVisible)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkVisible's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkOwner := self.Args("link_owner").MustInt64()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkOwner(offset, limit, iLinkImage,iLinkOwner)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkOwner's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkRating(offset, limit, iLinkImage,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkUpdated(offset, limit, iLinkImage,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkRel(offset, limit, iLinkImage,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkNotes(offset, limit, iLinkImage,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkImageAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkImage := self.Args("link_image").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksesByLinkImageAndLinkRss(offset, limit, iLinkImage,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkImageAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkDescription := self.Args("link_description").String()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkDescription(offset, limit, iLinkTarget,iLinkDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkVisibleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkVisible := self.Args("link_visible").String()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkVisible(offset, limit, iLinkTarget,iLinkVisible)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkVisible's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkOwner := self.Args("link_owner").MustInt64()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkOwner(offset, limit, iLinkTarget,iLinkOwner)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkOwner's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkRating(offset, limit, iLinkTarget,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkUpdated(offset, limit, iLinkTarget,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkRel(offset, limit, iLinkTarget,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkNotes(offset, limit, iLinkTarget,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkTargetAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkTarget := self.Args("link_target").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksesByLinkTargetAndLinkRss(offset, limit, iLinkTarget,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkTargetAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkDescriptionAndLinkVisibleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkDescription := self.Args("link_description").String()
	iLinkVisible := self.Args("link_visible").String()

	if helper.IsHas(iLinkDescription) {
		_Links, _error := model.GetLinksesByLinkDescriptionAndLinkVisible(offset, limit, iLinkDescription,iLinkVisible)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkDescriptionAndLinkVisible's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkDescriptionAndLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkDescription := self.Args("link_description").String()
	iLinkOwner := self.Args("link_owner").MustInt64()

	if helper.IsHas(iLinkDescription) {
		_Links, _error := model.GetLinksesByLinkDescriptionAndLinkOwner(offset, limit, iLinkDescription,iLinkOwner)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkDescriptionAndLinkOwner's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkDescriptionAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkDescription := self.Args("link_description").String()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkDescription) {
		_Links, _error := model.GetLinksesByLinkDescriptionAndLinkRating(offset, limit, iLinkDescription,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkDescriptionAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkDescriptionAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkDescription := self.Args("link_description").String()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkDescription) {
		_Links, _error := model.GetLinksesByLinkDescriptionAndLinkUpdated(offset, limit, iLinkDescription,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkDescriptionAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkDescriptionAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkDescription := self.Args("link_description").String()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkDescription) {
		_Links, _error := model.GetLinksesByLinkDescriptionAndLinkRel(offset, limit, iLinkDescription,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkDescriptionAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkDescriptionAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkDescription := self.Args("link_description").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkDescription) {
		_Links, _error := model.GetLinksesByLinkDescriptionAndLinkNotes(offset, limit, iLinkDescription,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkDescriptionAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkDescriptionAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkDescription := self.Args("link_description").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkDescription) {
		_Links, _error := model.GetLinksesByLinkDescriptionAndLinkRss(offset, limit, iLinkDescription,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkDescriptionAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkVisibleAndLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkVisible := self.Args("link_visible").String()
	iLinkOwner := self.Args("link_owner").MustInt64()

	if helper.IsHas(iLinkVisible) {
		_Links, _error := model.GetLinksesByLinkVisibleAndLinkOwner(offset, limit, iLinkVisible,iLinkOwner)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkVisibleAndLinkOwner's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkVisibleAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkVisible := self.Args("link_visible").String()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkVisible) {
		_Links, _error := model.GetLinksesByLinkVisibleAndLinkRating(offset, limit, iLinkVisible,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkVisibleAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkVisibleAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkVisible := self.Args("link_visible").String()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkVisible) {
		_Links, _error := model.GetLinksesByLinkVisibleAndLinkUpdated(offset, limit, iLinkVisible,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkVisibleAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkVisibleAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkVisible := self.Args("link_visible").String()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkVisible) {
		_Links, _error := model.GetLinksesByLinkVisibleAndLinkRel(offset, limit, iLinkVisible,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkVisibleAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkVisibleAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkVisible := self.Args("link_visible").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkVisible) {
		_Links, _error := model.GetLinksesByLinkVisibleAndLinkNotes(offset, limit, iLinkVisible,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkVisibleAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkVisibleAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkVisible := self.Args("link_visible").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkVisible) {
		_Links, _error := model.GetLinksesByLinkVisibleAndLinkRss(offset, limit, iLinkVisible,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkVisibleAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkOwnerAndLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkRating := self.Args("link_rating").MustInt()

	if helper.IsHas(iLinkOwner) {
		_Links, _error := model.GetLinksesByLinkOwnerAndLinkRating(offset, limit, iLinkOwner,iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkOwnerAndLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkOwnerAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkOwner) {
		_Links, _error := model.GetLinksesByLinkOwnerAndLinkUpdated(offset, limit, iLinkOwner,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkOwnerAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkOwnerAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkOwner) {
		_Links, _error := model.GetLinksesByLinkOwnerAndLinkRel(offset, limit, iLinkOwner,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkOwnerAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkOwnerAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkOwner) {
		_Links, _error := model.GetLinksesByLinkOwnerAndLinkNotes(offset, limit, iLinkOwner,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkOwnerAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkOwnerAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkOwner := self.Args("link_owner").MustInt64()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkOwner) {
		_Links, _error := model.GetLinksesByLinkOwnerAndLinkRss(offset, limit, iLinkOwner,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkOwnerAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkRatingAndLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkUpdated := self.Args("link_updated").Time()

	if helper.IsHas(iLinkRating) {
		_Links, _error := model.GetLinksesByLinkRatingAndLinkUpdated(offset, limit, iLinkRating,iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkRatingAndLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkRatingAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkRating) {
		_Links, _error := model.GetLinksesByLinkRatingAndLinkRel(offset, limit, iLinkRating,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkRatingAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkRatingAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkRating) {
		_Links, _error := model.GetLinksesByLinkRatingAndLinkNotes(offset, limit, iLinkRating,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkRatingAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkRatingAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkRating := self.Args("link_rating").MustInt()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkRating) {
		_Links, _error := model.GetLinksesByLinkRatingAndLinkRss(offset, limit, iLinkRating,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkRatingAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUpdatedAndLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkRel := self.Args("link_rel").String()

	if helper.IsHas(iLinkUpdated) {
		_Links, _error := model.GetLinksesByLinkUpdatedAndLinkRel(offset, limit, iLinkUpdated,iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUpdatedAndLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUpdatedAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkUpdated) {
		_Links, _error := model.GetLinksesByLinkUpdatedAndLinkNotes(offset, limit, iLinkUpdated,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUpdatedAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkUpdatedAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkUpdated := self.Args("link_updated").Time()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkUpdated) {
		_Links, _error := model.GetLinksesByLinkUpdatedAndLinkRss(offset, limit, iLinkUpdated,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkUpdatedAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkRelAndLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkRel := self.Args("link_rel").String()
	iLinkNotes := self.Args("link_notes").String()

	if helper.IsHas(iLinkRel) {
		_Links, _error := model.GetLinksesByLinkRelAndLinkNotes(offset, limit, iLinkRel,iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkRelAndLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkRelAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkRel := self.Args("link_rel").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkRel) {
		_Links, _error := model.GetLinksesByLinkRelAndLinkRss(offset, limit, iLinkRel,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkRelAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesByLinkNotesAndLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLinkNotes := self.Args("link_notes").String()
	iLinkRss := self.Args("link_rss").String()

	if helper.IsHas(iLinkNotes) {
		_Links, _error := model.GetLinksesByLinkNotesAndLinkRss(offset, limit, iLinkNotes,iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksesByLinkNotesAndLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_Links, _error := model.GetLinkses(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinkses' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasLinksViaLinkIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLinkId := self.Args("link_id").MustInt64()
	if helper.IsHas(iLinkId) {
		_Links := model.HasLinksViaLinkId(iLinkId)
		var m = map[string]interface{}{}
		m["links"] = _Links
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasLinksViaLinkId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasLinksViaLinkUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLinkUrl := self.Args("link_url").String()
	if helper.IsHas(iLinkUrl) {
		_Links := model.HasLinksViaLinkUrl(iLinkUrl)
		var m = map[string]interface{}{}
		m["links"] = _Links
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasLinksViaLinkUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasLinksViaLinkNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLinkName := self.Args("link_name").String()
	if helper.IsHas(iLinkName) {
		_Links := model.HasLinksViaLinkName(iLinkName)
		var m = map[string]interface{}{}
		m["links"] = _Links
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasLinksViaLinkName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasLinksViaLinkImageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLinkImage := self.Args("link_image").String()
	if helper.IsHas(iLinkImage) {
		_Links := model.HasLinksViaLinkImage(iLinkImage)
		var m = map[string]interface{}{}
		m["links"] = _Links
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasLinksViaLinkImage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasLinksViaLinkTargetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLinkTarget := self.Args("link_target").String()
	if helper.IsHas(iLinkTarget) {
		_Links := model.HasLinksViaLinkTarget(iLinkTarget)
		var m = map[string]interface{}{}
		m["links"] = _Links
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasLinksViaLinkTarget's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasLinksViaLinkDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLinkDescription := self.Args("link_description").String()
	if helper.IsHas(iLinkDescription) {
		_Links := model.HasLinksViaLinkDescription(iLinkDescription)
		var m = map[string]interface{}{}
		m["links"] = _Links
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasLinksViaLinkDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasLinksViaLinkVisibleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLinkVisible := self.Args("link_visible").String()
	if helper.IsHas(iLinkVisible) {
		_Links := model.HasLinksViaLinkVisible(iLinkVisible)
		var m = map[string]interface{}{}
		m["links"] = _Links
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasLinksViaLinkVisible's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasLinksViaLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLinkOwner := self.Args("link_owner").MustInt64()
	if helper.IsHas(iLinkOwner) {
		_Links := model.HasLinksViaLinkOwner(iLinkOwner)
		var m = map[string]interface{}{}
		m["links"] = _Links
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasLinksViaLinkOwner's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasLinksViaLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLinkRating := self.Args("link_rating").MustInt()
	if helper.IsHas(iLinkRating) {
		_Links := model.HasLinksViaLinkRating(iLinkRating)
		var m = map[string]interface{}{}
		m["links"] = _Links
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasLinksViaLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasLinksViaLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLinkUpdated := self.Args("link_updated").Time()
	if helper.IsHas(iLinkUpdated) {
		_Links := model.HasLinksViaLinkUpdated(iLinkUpdated)
		var m = map[string]interface{}{}
		m["links"] = _Links
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasLinksViaLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasLinksViaLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLinkRel := self.Args("link_rel").String()
	if helper.IsHas(iLinkRel) {
		_Links := model.HasLinksViaLinkRel(iLinkRel)
		var m = map[string]interface{}{}
		m["links"] = _Links
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasLinksViaLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasLinksViaLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLinkNotes := self.Args("link_notes").String()
	if helper.IsHas(iLinkNotes) {
		_Links := model.HasLinksViaLinkNotes(iLinkNotes)
		var m = map[string]interface{}{}
		m["links"] = _Links
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasLinksViaLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasLinksViaLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLinkRss := self.Args("link_rss").String()
	if helper.IsHas(iLinkRss) {
		_Links := model.HasLinksViaLinkRss(iLinkRss)
		var m = map[string]interface{}{}
		m["links"] = _Links
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasLinksViaLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksViaLinkIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLinkId := self.Args("link_id").MustInt64()
	if helper.IsHas(iLinkId) {
		_Links, _error := model.GetLinksViaLinkId(iLinkId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksViaLinkId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksViaLinkUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLinkUrl := self.Args("link_url").String()
	if helper.IsHas(iLinkUrl) {
		_Links, _error := model.GetLinksViaLinkUrl(iLinkUrl)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksViaLinkUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksViaLinkNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLinkName := self.Args("link_name").String()
	if helper.IsHas(iLinkName) {
		_Links, _error := model.GetLinksViaLinkName(iLinkName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksViaLinkName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksViaLinkImageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLinkImage := self.Args("link_image").String()
	if helper.IsHas(iLinkImage) {
		_Links, _error := model.GetLinksViaLinkImage(iLinkImage)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksViaLinkImage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksViaLinkTargetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLinkTarget := self.Args("link_target").String()
	if helper.IsHas(iLinkTarget) {
		_Links, _error := model.GetLinksViaLinkTarget(iLinkTarget)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksViaLinkTarget's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksViaLinkDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLinkDescription := self.Args("link_description").String()
	if helper.IsHas(iLinkDescription) {
		_Links, _error := model.GetLinksViaLinkDescription(iLinkDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksViaLinkDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksViaLinkVisibleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLinkVisible := self.Args("link_visible").String()
	if helper.IsHas(iLinkVisible) {
		_Links, _error := model.GetLinksViaLinkVisible(iLinkVisible)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksViaLinkVisible's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksViaLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLinkOwner := self.Args("link_owner").MustInt64()
	if helper.IsHas(iLinkOwner) {
		_Links, _error := model.GetLinksViaLinkOwner(iLinkOwner)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksViaLinkOwner's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksViaLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLinkRating := self.Args("link_rating").MustInt()
	if helper.IsHas(iLinkRating) {
		_Links, _error := model.GetLinksViaLinkRating(iLinkRating)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksViaLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksViaLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLinkUpdated := self.Args("link_updated").Time()
	if helper.IsHas(iLinkUpdated) {
		_Links, _error := model.GetLinksViaLinkUpdated(iLinkUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksViaLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksViaLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLinkRel := self.Args("link_rel").String()
	if helper.IsHas(iLinkRel) {
		_Links, _error := model.GetLinksViaLinkRel(iLinkRel)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksViaLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksViaLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLinkNotes := self.Args("link_notes").String()
	if helper.IsHas(iLinkNotes) {
		_Links, _error := model.GetLinksViaLinkNotes(iLinkNotes)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksViaLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetLinksViaLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLinkRss := self.Args("link_rss").String()
	if helper.IsHas(iLinkRss) {
		_Links, _error := model.GetLinksViaLinkRss(iLinkRss)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the GetLinksViaLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetLinksViaLinkIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkId_ := self.Args("link_id").MustInt64()
	if helper.IsHas(LinkId_) {
		var iLinks model.Links
		self.Bind(&iLinks)
		_Links, _error := model.SetLinksViaLinkId(LinkId_, &iLinks)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the SetLinksViaLinkId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetLinksViaLinkUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkUrl_ := self.Args("link_url").String()
	if helper.IsHas(LinkUrl_) {
		var iLinks model.Links
		self.Bind(&iLinks)
		_Links, _error := model.SetLinksViaLinkUrl(LinkUrl_, &iLinks)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the SetLinksViaLinkUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetLinksViaLinkNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkName_ := self.Args("link_name").String()
	if helper.IsHas(LinkName_) {
		var iLinks model.Links
		self.Bind(&iLinks)
		_Links, _error := model.SetLinksViaLinkName(LinkName_, &iLinks)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the SetLinksViaLinkName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetLinksViaLinkImageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkImage_ := self.Args("link_image").String()
	if helper.IsHas(LinkImage_) {
		var iLinks model.Links
		self.Bind(&iLinks)
		_Links, _error := model.SetLinksViaLinkImage(LinkImage_, &iLinks)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the SetLinksViaLinkImage's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetLinksViaLinkTargetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkTarget_ := self.Args("link_target").String()
	if helper.IsHas(LinkTarget_) {
		var iLinks model.Links
		self.Bind(&iLinks)
		_Links, _error := model.SetLinksViaLinkTarget(LinkTarget_, &iLinks)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the SetLinksViaLinkTarget's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetLinksViaLinkDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkDescription_ := self.Args("link_description").String()
	if helper.IsHas(LinkDescription_) {
		var iLinks model.Links
		self.Bind(&iLinks)
		_Links, _error := model.SetLinksViaLinkDescription(LinkDescription_, &iLinks)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the SetLinksViaLinkDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetLinksViaLinkVisibleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkVisible_ := self.Args("link_visible").String()
	if helper.IsHas(LinkVisible_) {
		var iLinks model.Links
		self.Bind(&iLinks)
		_Links, _error := model.SetLinksViaLinkVisible(LinkVisible_, &iLinks)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the SetLinksViaLinkVisible's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetLinksViaLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkOwner_ := self.Args("link_owner").MustInt64()
	if helper.IsHas(LinkOwner_) {
		var iLinks model.Links
		self.Bind(&iLinks)
		_Links, _error := model.SetLinksViaLinkOwner(LinkOwner_, &iLinks)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the SetLinksViaLinkOwner's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetLinksViaLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkRating_ := self.Args("link_rating").MustInt()
	if helper.IsHas(LinkRating_) {
		var iLinks model.Links
		self.Bind(&iLinks)
		_Links, _error := model.SetLinksViaLinkRating(LinkRating_, &iLinks)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the SetLinksViaLinkRating's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetLinksViaLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkUpdated_ := self.Args("link_updated").Time()
	if helper.IsHas(LinkUpdated_) {
		var iLinks model.Links
		self.Bind(&iLinks)
		_Links, _error := model.SetLinksViaLinkUpdated(LinkUpdated_, &iLinks)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the SetLinksViaLinkUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetLinksViaLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkRel_ := self.Args("link_rel").String()
	if helper.IsHas(LinkRel_) {
		var iLinks model.Links
		self.Bind(&iLinks)
		_Links, _error := model.SetLinksViaLinkRel(LinkRel_, &iLinks)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the SetLinksViaLinkRel's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetLinksViaLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkNotes_ := self.Args("link_notes").String()
	if helper.IsHas(LinkNotes_) {
		var iLinks model.Links
		self.Bind(&iLinks)
		_Links, _error := model.SetLinksViaLinkNotes(LinkNotes_, &iLinks)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the SetLinksViaLinkNotes's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetLinksViaLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkRss_ := self.Args("link_rss").String()
	if helper.IsHas(LinkRss_) {
		var iLinks model.Links
		self.Bind(&iLinks)
		_Links, _error := model.SetLinksViaLinkRss(LinkRss_, &iLinks)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Links)
	}
	herr.Message = "Can't get to the SetLinksViaLinkRss's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddLinksHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkId_ := self.Args("link_id").MustInt64()
	LinkUrl_ := self.Args("link_url").String()
	LinkName_ := self.Args("link_name").String()
	LinkImage_ := self.Args("link_image").String()
	LinkTarget_ := self.Args("link_target").String()
	LinkDescription_ := self.Args("link_description").String()
	LinkVisible_ := self.Args("link_visible").String()
	LinkOwner_ := self.Args("link_owner").MustInt64()
	LinkRating_ := self.Args("link_rating").MustInt()
	LinkUpdated_ := self.Args("link_updated").Time()
	LinkRel_ := self.Args("link_rel").String()
	LinkNotes_ := self.Args("link_notes").String()
	LinkRss_ := self.Args("link_rss").String()

	if helper.IsHas(LinkId_) {
		_error := model.AddLinks(LinkId_,LinkUrl_,LinkName_,LinkImage_,LinkTarget_,LinkDescription_,LinkVisible_,LinkOwner_,LinkRating_,LinkUpdated_,LinkRel_,LinkNotes_,LinkRss_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddLinks's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostLinksHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iLinks model.Links
	self.Bind(&iLinks)
	_int64, _error := model.PostLinks(&iLinks)
	if (helper.IsHas(_int64)) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	} else {
		var m = map[string]interface{}{}
		m["created"] = _int64
		return self.JSON(m, macross.StatusCreated)
	}
	return self.JSON(herr)
}

func PutLinksHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iLinks model.Links
	self.Bind(&iLinks)
	_int64, _error := model.PutLinks(&iLinks)
	if (helper.IsHas(_int64)) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	} else {
		var m = map[string]interface{}{}
		m["updated"] = _int64
		return self.JSON(m)
	}
	return self.JSON(herr)
}

func PutLinksViaLinkIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkId_ := self.Args("link_id").MustInt64()
	var iLinks model.Links
	self.Bind(&iLinks)
	_int64, _error := model.PutLinksViaLinkId(LinkId_, &iLinks)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutLinksViaLinkUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkUrl_ := self.Args("link_url").String()
	var iLinks model.Links
	self.Bind(&iLinks)
	_int64, _error := model.PutLinksViaLinkUrl(LinkUrl_, &iLinks)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutLinksViaLinkNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkName_ := self.Args("link_name").String()
	var iLinks model.Links
	self.Bind(&iLinks)
	_int64, _error := model.PutLinksViaLinkName(LinkName_, &iLinks)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutLinksViaLinkImageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkImage_ := self.Args("link_image").String()
	var iLinks model.Links
	self.Bind(&iLinks)
	_int64, _error := model.PutLinksViaLinkImage(LinkImage_, &iLinks)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutLinksViaLinkTargetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkTarget_ := self.Args("link_target").String()
	var iLinks model.Links
	self.Bind(&iLinks)
	_int64, _error := model.PutLinksViaLinkTarget(LinkTarget_, &iLinks)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutLinksViaLinkDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkDescription_ := self.Args("link_description").String()
	var iLinks model.Links
	self.Bind(&iLinks)
	_int64, _error := model.PutLinksViaLinkDescription(LinkDescription_, &iLinks)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutLinksViaLinkVisibleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkVisible_ := self.Args("link_visible").String()
	var iLinks model.Links
	self.Bind(&iLinks)
	_int64, _error := model.PutLinksViaLinkVisible(LinkVisible_, &iLinks)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutLinksViaLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkOwner_ := self.Args("link_owner").MustInt64()
	var iLinks model.Links
	self.Bind(&iLinks)
	_int64, _error := model.PutLinksViaLinkOwner(LinkOwner_, &iLinks)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutLinksViaLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkRating_ := self.Args("link_rating").MustInt()
	var iLinks model.Links
	self.Bind(&iLinks)
	_int64, _error := model.PutLinksViaLinkRating(LinkRating_, &iLinks)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutLinksViaLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkUpdated_ := self.Args("link_updated").Time()
	var iLinks model.Links
	self.Bind(&iLinks)
	_int64, _error := model.PutLinksViaLinkUpdated(LinkUpdated_, &iLinks)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutLinksViaLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkRel_ := self.Args("link_rel").String()
	var iLinks model.Links
	self.Bind(&iLinks)
	_int64, _error := model.PutLinksViaLinkRel(LinkRel_, &iLinks)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutLinksViaLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkNotes_ := self.Args("link_notes").String()
	var iLinks model.Links
	self.Bind(&iLinks)
	_int64, _error := model.PutLinksViaLinkNotes(LinkNotes_, &iLinks)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutLinksViaLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkRss_ := self.Args("link_rss").String()
	var iLinks model.Links
	self.Bind(&iLinks)
	_int64, _error := model.PutLinksViaLinkRss(LinkRss_, &iLinks)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateLinksViaLinkIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkId_ := self.Args("link_id").MustInt64()
	var iLinks model.Links
	self.Bind(&iLinks)
	var iMap = helper.StructToMap(iLinks)
	_error := model.UpdateLinksViaLinkId(LinkId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateLinksViaLinkUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkUrl_ := self.Args("link_url").String()
	var iLinks model.Links
	self.Bind(&iLinks)
	var iMap = helper.StructToMap(iLinks)
	_error := model.UpdateLinksViaLinkUrl(LinkUrl_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateLinksViaLinkNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkName_ := self.Args("link_name").String()
	var iLinks model.Links
	self.Bind(&iLinks)
	var iMap = helper.StructToMap(iLinks)
	_error := model.UpdateLinksViaLinkName(LinkName_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateLinksViaLinkImageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkImage_ := self.Args("link_image").String()
	var iLinks model.Links
	self.Bind(&iLinks)
	var iMap = helper.StructToMap(iLinks)
	_error := model.UpdateLinksViaLinkImage(LinkImage_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateLinksViaLinkTargetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkTarget_ := self.Args("link_target").String()
	var iLinks model.Links
	self.Bind(&iLinks)
	var iMap = helper.StructToMap(iLinks)
	_error := model.UpdateLinksViaLinkTarget(LinkTarget_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateLinksViaLinkDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkDescription_ := self.Args("link_description").String()
	var iLinks model.Links
	self.Bind(&iLinks)
	var iMap = helper.StructToMap(iLinks)
	_error := model.UpdateLinksViaLinkDescription(LinkDescription_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateLinksViaLinkVisibleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkVisible_ := self.Args("link_visible").String()
	var iLinks model.Links
	self.Bind(&iLinks)
	var iMap = helper.StructToMap(iLinks)
	_error := model.UpdateLinksViaLinkVisible(LinkVisible_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateLinksViaLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkOwner_ := self.Args("link_owner").MustInt64()
	var iLinks model.Links
	self.Bind(&iLinks)
	var iMap = helper.StructToMap(iLinks)
	_error := model.UpdateLinksViaLinkOwner(LinkOwner_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateLinksViaLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkRating_ := self.Args("link_rating").MustInt()
	var iLinks model.Links
	self.Bind(&iLinks)
	var iMap = helper.StructToMap(iLinks)
	_error := model.UpdateLinksViaLinkRating(LinkRating_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateLinksViaLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkUpdated_ := self.Args("link_updated").Time()
	var iLinks model.Links
	self.Bind(&iLinks)
	var iMap = helper.StructToMap(iLinks)
	_error := model.UpdateLinksViaLinkUpdated(LinkUpdated_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateLinksViaLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkRel_ := self.Args("link_rel").String()
	var iLinks model.Links
	self.Bind(&iLinks)
	var iMap = helper.StructToMap(iLinks)
	_error := model.UpdateLinksViaLinkRel(LinkRel_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateLinksViaLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkNotes_ := self.Args("link_notes").String()
	var iLinks model.Links
	self.Bind(&iLinks)
	var iMap = helper.StructToMap(iLinks)
	_error := model.UpdateLinksViaLinkNotes(LinkNotes_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateLinksViaLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkRss_ := self.Args("link_rss").String()
	var iLinks model.Links
	self.Bind(&iLinks)
	var iMap = helper.StructToMap(iLinks)
	_error := model.UpdateLinksViaLinkRss(LinkRss_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteLinksHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkId_ := self.Args("link_id").MustInt64()
	_int64, _error := model.DeleteLinks(LinkId_)
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

func DeleteLinksViaLinkIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkId_ := self.Args("link_id").MustInt64()
	_error := model.DeleteLinksViaLinkId(LinkId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteLinksViaLinkUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkUrl_ := self.Args("link_url").String()
	_error := model.DeleteLinksViaLinkUrl(LinkUrl_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteLinksViaLinkNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkName_ := self.Args("link_name").String()
	_error := model.DeleteLinksViaLinkName(LinkName_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteLinksViaLinkImageHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkImage_ := self.Args("link_image").String()
	_error := model.DeleteLinksViaLinkImage(LinkImage_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteLinksViaLinkTargetHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkTarget_ := self.Args("link_target").String()
	_error := model.DeleteLinksViaLinkTarget(LinkTarget_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteLinksViaLinkDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkDescription_ := self.Args("link_description").String()
	_error := model.DeleteLinksViaLinkDescription(LinkDescription_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteLinksViaLinkVisibleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkVisible_ := self.Args("link_visible").String()
	_error := model.DeleteLinksViaLinkVisible(LinkVisible_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteLinksViaLinkOwnerHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkOwner_ := self.Args("link_owner").MustInt64()
	_error := model.DeleteLinksViaLinkOwner(LinkOwner_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteLinksViaLinkRatingHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkRating_ := self.Args("link_rating").MustInt()
	_error := model.DeleteLinksViaLinkRating(LinkRating_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteLinksViaLinkUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkUpdated_ := self.Args("link_updated").Time()
	_error := model.DeleteLinksViaLinkUpdated(LinkUpdated_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteLinksViaLinkRelHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkRel_ := self.Args("link_rel").String()
	_error := model.DeleteLinksViaLinkRel(LinkRel_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteLinksViaLinkNotesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkNotes_ := self.Args("link_notes").String()
	_error := model.DeleteLinksViaLinkNotes(LinkNotes_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteLinksViaLinkRssHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LinkRss_ := self.Args("link_rss").String()
	_error := model.DeleteLinksViaLinkRss(LinkRss_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
