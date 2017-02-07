package handler

import (
	"github.com/insionng/zenpress/helper"
	"github.com/insionng/zenpress/model"
	"github.com/insionng/macross"
)

func GetBlogsesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetBlogsesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["blogssCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetBlogsesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsCountByBlogIdHandler(self *macross.Context) error {
	BlogId_ := self.Args("blog_id").MustInt64()
	_int64 := model.GetBlogsCountByBlogId(BlogId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["blogsCount"] = 0
	}
	m["blogsCount"] = _int64
	return self.JSON(m)
}

func GetBlogsCountBySiteIdHandler(self *macross.Context) error {
	SiteId_ := self.Args("site_id").MustInt64()
	_int64 := model.GetBlogsCountBySiteId(SiteId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["blogsCount"] = 0
	}
	m["blogsCount"] = _int64
	return self.JSON(m)
}

func GetBlogsCountByDomainHandler(self *macross.Context) error {
	Domain_ := self.Args("domain").String()
	_int64 := model.GetBlogsCountByDomain(Domain_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["blogsCount"] = 0
	}
	m["blogsCount"] = _int64
	return self.JSON(m)
}

func GetBlogsCountByPathHandler(self *macross.Context) error {
	Path_ := self.Args("path").String()
	_int64 := model.GetBlogsCountByPath(Path_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["blogsCount"] = 0
	}
	m["blogsCount"] = _int64
	return self.JSON(m)
}

func GetBlogsCountByRegisteredHandler(self *macross.Context) error {
	Registered_ := self.Args("registered").Time()
	_int64 := model.GetBlogsCountByRegistered(Registered_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["blogsCount"] = 0
	}
	m["blogsCount"] = _int64
	return self.JSON(m)
}

func GetBlogsCountByLastUpdatedHandler(self *macross.Context) error {
	LastUpdated_ := self.Args("last_updated").Time()
	_int64 := model.GetBlogsCountByLastUpdated(LastUpdated_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["blogsCount"] = 0
	}
	m["blogsCount"] = _int64
	return self.JSON(m)
}

func GetBlogsCountByPublicHandler(self *macross.Context) error {
	Public_ := self.Args("public").MustInt()
	_int64 := model.GetBlogsCountByPublic(Public_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["blogsCount"] = 0
	}
	m["blogsCount"] = _int64
	return self.JSON(m)
}

func GetBlogsCountByArchivedHandler(self *macross.Context) error {
	Archived_ := self.Args("archived").MustInt()
	_int64 := model.GetBlogsCountByArchived(Archived_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["blogsCount"] = 0
	}
	m["blogsCount"] = _int64
	return self.JSON(m)
}

func GetBlogsCountByMatureHandler(self *macross.Context) error {
	Mature_ := self.Args("mature").MustInt()
	_int64 := model.GetBlogsCountByMature(Mature_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["blogsCount"] = 0
	}
	m["blogsCount"] = _int64
	return self.JSON(m)
}

func GetBlogsCountBySpamHandler(self *macross.Context) error {
	Spam_ := self.Args("spam").MustInt()
	_int64 := model.GetBlogsCountBySpam(Spam_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["blogsCount"] = 0
	}
	m["blogsCount"] = _int64
	return self.JSON(m)
}

func GetBlogsCountByDeletedHandler(self *macross.Context) error {
	Deleted_ := self.Args("deleted").MustInt()
	_int64 := model.GetBlogsCountByDeleted(Deleted_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["blogsCount"] = 0
	}
	m["blogsCount"] = _int64
	return self.JSON(m)
}

func GetBlogsCountByLangIdHandler(self *macross.Context) error {
	LangId_ := self.Args("lang_id").MustInt()
	_int64 := model.GetBlogsCountByLangId(LangId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["blogsCount"] = 0
	}
	m["blogsCount"] = _int64
	return self.JSON(m)
}

func GetBlogsesByBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iBlogId := self.Args("blog_id").MustInt64()
	if (offset > 0) && helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogId(offset, limit, iBlogId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iSiteId := self.Args("site_id").MustInt64()
	if (offset > 0) && helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteId(offset, limit, iSiteId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDomain := self.Args("domain").String()
	if (offset > 0) && helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomain(offset, limit, iDomain, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomain's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iPath := self.Args("path").String()
	if (offset > 0) && helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPath(offset, limit, iPath, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPath's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iRegistered := self.Args("registered").Time()
	if (offset > 0) && helper.IsHas(iRegistered) {
		_Blogs, _error := model.GetBlogsesByRegistered(offset, limit, iRegistered, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLastUpdated := self.Args("last_updated").Time()
	if (offset > 0) && helper.IsHas(iLastUpdated) {
		_Blogs, _error := model.GetBlogsesByLastUpdated(offset, limit, iLastUpdated, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByLastUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPublicHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iPublic := self.Args("public").MustInt()
	if (offset > 0) && helper.IsHas(iPublic) {
		_Blogs, _error := model.GetBlogsesByPublic(offset, limit, iPublic, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPublic's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iArchived := self.Args("archived").MustInt()
	if (offset > 0) && helper.IsHas(iArchived) {
		_Blogs, _error := model.GetBlogsesByArchived(offset, limit, iArchived, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByArchived's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMature := self.Args("mature").MustInt()
	if (offset > 0) && helper.IsHas(iMature) {
		_Blogs, _error := model.GetBlogsesByMature(offset, limit, iMature, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iSpam := self.Args("spam").MustInt()
	if (offset > 0) && helper.IsHas(iSpam) {
		_Blogs, _error := model.GetBlogsesBySpam(offset, limit, iSpam, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDeleted := self.Args("deleted").MustInt()
	if (offset > 0) && helper.IsHas(iDeleted) {
		_Blogs, _error := model.GetBlogsesByDeleted(offset, limit, iDeleted, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLangId := self.Args("lang_id").MustInt()
	if (offset > 0) && helper.IsHas(iLangId) {
		_Blogs, _error := model.GetBlogsesByLangId(offset, limit, iLangId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndSiteIdAndDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iSiteId := self.Args("site_id").MustInt64()
	iDomain := self.Args("domain").String()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndSiteIdAndDomain(offset, limit, iBlogId,iSiteId,iDomain)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndSiteIdAndDomain's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndSiteIdAndPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iSiteId := self.Args("site_id").MustInt64()
	iPath := self.Args("path").String()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndSiteIdAndPath(offset, limit, iBlogId,iSiteId,iPath)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndSiteIdAndPath's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndSiteIdAndRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iSiteId := self.Args("site_id").MustInt64()
	iRegistered := self.Args("registered").Time()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndSiteIdAndRegistered(offset, limit, iBlogId,iSiteId,iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndSiteIdAndRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndSiteIdAndLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iSiteId := self.Args("site_id").MustInt64()
	iLastUpdated := self.Args("last_updated").Time()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndSiteIdAndLastUpdated(offset, limit, iBlogId,iSiteId,iLastUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndSiteIdAndLastUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndSiteIdAndPublicHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iSiteId := self.Args("site_id").MustInt64()
	iPublic := self.Args("public").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndSiteIdAndPublic(offset, limit, iBlogId,iSiteId,iPublic)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndSiteIdAndPublic's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndSiteIdAndArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iSiteId := self.Args("site_id").MustInt64()
	iArchived := self.Args("archived").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndSiteIdAndArchived(offset, limit, iBlogId,iSiteId,iArchived)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndSiteIdAndArchived's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndSiteIdAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iSiteId := self.Args("site_id").MustInt64()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndSiteIdAndMature(offset, limit, iBlogId,iSiteId,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndSiteIdAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndSiteIdAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iSiteId := self.Args("site_id").MustInt64()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndSiteIdAndSpam(offset, limit, iBlogId,iSiteId,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndSiteIdAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndSiteIdAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iSiteId := self.Args("site_id").MustInt64()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndSiteIdAndDeleted(offset, limit, iBlogId,iSiteId,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndSiteIdAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndSiteIdAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iSiteId := self.Args("site_id").MustInt64()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndSiteIdAndLangId(offset, limit, iBlogId,iSiteId,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndSiteIdAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndDomainAndPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iDomain := self.Args("domain").String()
	iPath := self.Args("path").String()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndDomainAndPath(offset, limit, iBlogId,iDomain,iPath)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndDomainAndPath's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndDomainAndRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iDomain := self.Args("domain").String()
	iRegistered := self.Args("registered").Time()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndDomainAndRegistered(offset, limit, iBlogId,iDomain,iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndDomainAndRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndDomainAndLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iDomain := self.Args("domain").String()
	iLastUpdated := self.Args("last_updated").Time()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndDomainAndLastUpdated(offset, limit, iBlogId,iDomain,iLastUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndDomainAndLastUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndDomainAndPublicHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iDomain := self.Args("domain").String()
	iPublic := self.Args("public").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndDomainAndPublic(offset, limit, iBlogId,iDomain,iPublic)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndDomainAndPublic's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndDomainAndArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iDomain := self.Args("domain").String()
	iArchived := self.Args("archived").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndDomainAndArchived(offset, limit, iBlogId,iDomain,iArchived)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndDomainAndArchived's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndDomainAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iDomain := self.Args("domain").String()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndDomainAndMature(offset, limit, iBlogId,iDomain,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndDomainAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndDomainAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iDomain := self.Args("domain").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndDomainAndSpam(offset, limit, iBlogId,iDomain,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndDomainAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndDomainAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iDomain := self.Args("domain").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndDomainAndDeleted(offset, limit, iBlogId,iDomain,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndDomainAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndDomainAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iDomain := self.Args("domain").String()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndDomainAndLangId(offset, limit, iBlogId,iDomain,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndDomainAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndPathAndRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iPath := self.Args("path").String()
	iRegistered := self.Args("registered").Time()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndPathAndRegistered(offset, limit, iBlogId,iPath,iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndPathAndRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndPathAndLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iPath := self.Args("path").String()
	iLastUpdated := self.Args("last_updated").Time()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndPathAndLastUpdated(offset, limit, iBlogId,iPath,iLastUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndPathAndLastUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndPathAndPublicHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iPath := self.Args("path").String()
	iPublic := self.Args("public").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndPathAndPublic(offset, limit, iBlogId,iPath,iPublic)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndPathAndPublic's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndPathAndArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iPath := self.Args("path").String()
	iArchived := self.Args("archived").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndPathAndArchived(offset, limit, iBlogId,iPath,iArchived)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndPathAndArchived's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndPathAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iPath := self.Args("path").String()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndPathAndMature(offset, limit, iBlogId,iPath,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndPathAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndPathAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iPath := self.Args("path").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndPathAndSpam(offset, limit, iBlogId,iPath,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndPathAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndPathAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iPath := self.Args("path").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndPathAndDeleted(offset, limit, iBlogId,iPath,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndPathAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndPathAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iPath := self.Args("path").String()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndPathAndLangId(offset, limit, iBlogId,iPath,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndPathAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndRegisteredAndLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iRegistered := self.Args("registered").Time()
	iLastUpdated := self.Args("last_updated").Time()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndRegisteredAndLastUpdated(offset, limit, iBlogId,iRegistered,iLastUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndRegisteredAndLastUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndRegisteredAndPublicHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iRegistered := self.Args("registered").Time()
	iPublic := self.Args("public").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndRegisteredAndPublic(offset, limit, iBlogId,iRegistered,iPublic)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndRegisteredAndPublic's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndRegisteredAndArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iRegistered := self.Args("registered").Time()
	iArchived := self.Args("archived").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndRegisteredAndArchived(offset, limit, iBlogId,iRegistered,iArchived)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndRegisteredAndArchived's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndRegisteredAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iRegistered := self.Args("registered").Time()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndRegisteredAndMature(offset, limit, iBlogId,iRegistered,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndRegisteredAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndRegisteredAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iRegistered := self.Args("registered").Time()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndRegisteredAndSpam(offset, limit, iBlogId,iRegistered,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndRegisteredAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndRegisteredAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iRegistered := self.Args("registered").Time()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndRegisteredAndDeleted(offset, limit, iBlogId,iRegistered,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndRegisteredAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndRegisteredAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iRegistered := self.Args("registered").Time()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndRegisteredAndLangId(offset, limit, iBlogId,iRegistered,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndRegisteredAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndLastUpdatedAndPublicHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iLastUpdated := self.Args("last_updated").Time()
	iPublic := self.Args("public").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndLastUpdatedAndPublic(offset, limit, iBlogId,iLastUpdated,iPublic)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndLastUpdatedAndPublic's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndLastUpdatedAndArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iLastUpdated := self.Args("last_updated").Time()
	iArchived := self.Args("archived").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndLastUpdatedAndArchived(offset, limit, iBlogId,iLastUpdated,iArchived)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndLastUpdatedAndArchived's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndLastUpdatedAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iLastUpdated := self.Args("last_updated").Time()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndLastUpdatedAndMature(offset, limit, iBlogId,iLastUpdated,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndLastUpdatedAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndLastUpdatedAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iLastUpdated := self.Args("last_updated").Time()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndLastUpdatedAndSpam(offset, limit, iBlogId,iLastUpdated,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndLastUpdatedAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndLastUpdatedAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iLastUpdated := self.Args("last_updated").Time()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndLastUpdatedAndDeleted(offset, limit, iBlogId,iLastUpdated,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndLastUpdatedAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndLastUpdatedAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iLastUpdated := self.Args("last_updated").Time()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndLastUpdatedAndLangId(offset, limit, iBlogId,iLastUpdated,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndLastUpdatedAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndPublicAndArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iPublic := self.Args("public").MustInt()
	iArchived := self.Args("archived").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndPublicAndArchived(offset, limit, iBlogId,iPublic,iArchived)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndPublicAndArchived's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndPublicAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iPublic := self.Args("public").MustInt()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndPublicAndMature(offset, limit, iBlogId,iPublic,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndPublicAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndPublicAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iPublic := self.Args("public").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndPublicAndSpam(offset, limit, iBlogId,iPublic,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndPublicAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndPublicAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iPublic := self.Args("public").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndPublicAndDeleted(offset, limit, iBlogId,iPublic,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndPublicAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndPublicAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iPublic := self.Args("public").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndPublicAndLangId(offset, limit, iBlogId,iPublic,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndPublicAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndArchivedAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iArchived := self.Args("archived").MustInt()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndArchivedAndMature(offset, limit, iBlogId,iArchived,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndArchivedAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndArchivedAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iArchived := self.Args("archived").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndArchivedAndSpam(offset, limit, iBlogId,iArchived,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndArchivedAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndArchivedAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iArchived := self.Args("archived").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndArchivedAndDeleted(offset, limit, iBlogId,iArchived,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndArchivedAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndArchivedAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iArchived := self.Args("archived").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndArchivedAndLangId(offset, limit, iBlogId,iArchived,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndArchivedAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndMatureAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iMature := self.Args("mature").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndMatureAndSpam(offset, limit, iBlogId,iMature,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndMatureAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndMatureAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iMature := self.Args("mature").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndMatureAndDeleted(offset, limit, iBlogId,iMature,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndMatureAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndMatureAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iMature := self.Args("mature").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndMatureAndLangId(offset, limit, iBlogId,iMature,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndMatureAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndSpamAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iSpam := self.Args("spam").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndSpamAndDeleted(offset, limit, iBlogId,iSpam,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndSpamAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndSpamAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iSpam := self.Args("spam").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndSpamAndLangId(offset, limit, iBlogId,iSpam,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndSpamAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndDeletedAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iDeleted := self.Args("deleted").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndDeletedAndLangId(offset, limit, iBlogId,iDeleted,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndDeletedAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndDomainAndPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iDomain := self.Args("domain").String()
	iPath := self.Args("path").String()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndDomainAndPath(offset, limit, iSiteId,iDomain,iPath)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndDomainAndPath's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndDomainAndRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iDomain := self.Args("domain").String()
	iRegistered := self.Args("registered").Time()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndDomainAndRegistered(offset, limit, iSiteId,iDomain,iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndDomainAndRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndDomainAndLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iDomain := self.Args("domain").String()
	iLastUpdated := self.Args("last_updated").Time()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndDomainAndLastUpdated(offset, limit, iSiteId,iDomain,iLastUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndDomainAndLastUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndDomainAndPublicHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iDomain := self.Args("domain").String()
	iPublic := self.Args("public").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndDomainAndPublic(offset, limit, iSiteId,iDomain,iPublic)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndDomainAndPublic's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndDomainAndArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iDomain := self.Args("domain").String()
	iArchived := self.Args("archived").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndDomainAndArchived(offset, limit, iSiteId,iDomain,iArchived)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndDomainAndArchived's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndDomainAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iDomain := self.Args("domain").String()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndDomainAndMature(offset, limit, iSiteId,iDomain,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndDomainAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndDomainAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iDomain := self.Args("domain").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndDomainAndSpam(offset, limit, iSiteId,iDomain,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndDomainAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndDomainAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iDomain := self.Args("domain").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndDomainAndDeleted(offset, limit, iSiteId,iDomain,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndDomainAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndDomainAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iDomain := self.Args("domain").String()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndDomainAndLangId(offset, limit, iSiteId,iDomain,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndDomainAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndPathAndRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iPath := self.Args("path").String()
	iRegistered := self.Args("registered").Time()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndPathAndRegistered(offset, limit, iSiteId,iPath,iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndPathAndRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndPathAndLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iPath := self.Args("path").String()
	iLastUpdated := self.Args("last_updated").Time()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndPathAndLastUpdated(offset, limit, iSiteId,iPath,iLastUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndPathAndLastUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndPathAndPublicHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iPath := self.Args("path").String()
	iPublic := self.Args("public").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndPathAndPublic(offset, limit, iSiteId,iPath,iPublic)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndPathAndPublic's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndPathAndArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iPath := self.Args("path").String()
	iArchived := self.Args("archived").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndPathAndArchived(offset, limit, iSiteId,iPath,iArchived)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndPathAndArchived's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndPathAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iPath := self.Args("path").String()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndPathAndMature(offset, limit, iSiteId,iPath,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndPathAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndPathAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iPath := self.Args("path").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndPathAndSpam(offset, limit, iSiteId,iPath,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndPathAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndPathAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iPath := self.Args("path").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndPathAndDeleted(offset, limit, iSiteId,iPath,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndPathAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndPathAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iPath := self.Args("path").String()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndPathAndLangId(offset, limit, iSiteId,iPath,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndPathAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndRegisteredAndLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iRegistered := self.Args("registered").Time()
	iLastUpdated := self.Args("last_updated").Time()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndRegisteredAndLastUpdated(offset, limit, iSiteId,iRegistered,iLastUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndRegisteredAndLastUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndRegisteredAndPublicHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iRegistered := self.Args("registered").Time()
	iPublic := self.Args("public").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndRegisteredAndPublic(offset, limit, iSiteId,iRegistered,iPublic)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndRegisteredAndPublic's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndRegisteredAndArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iRegistered := self.Args("registered").Time()
	iArchived := self.Args("archived").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndRegisteredAndArchived(offset, limit, iSiteId,iRegistered,iArchived)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndRegisteredAndArchived's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndRegisteredAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iRegistered := self.Args("registered").Time()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndRegisteredAndMature(offset, limit, iSiteId,iRegistered,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndRegisteredAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndRegisteredAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iRegistered := self.Args("registered").Time()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndRegisteredAndSpam(offset, limit, iSiteId,iRegistered,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndRegisteredAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndRegisteredAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iRegistered := self.Args("registered").Time()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndRegisteredAndDeleted(offset, limit, iSiteId,iRegistered,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndRegisteredAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndRegisteredAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iRegistered := self.Args("registered").Time()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndRegisteredAndLangId(offset, limit, iSiteId,iRegistered,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndRegisteredAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndLastUpdatedAndPublicHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iLastUpdated := self.Args("last_updated").Time()
	iPublic := self.Args("public").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndLastUpdatedAndPublic(offset, limit, iSiteId,iLastUpdated,iPublic)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndLastUpdatedAndPublic's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndLastUpdatedAndArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iLastUpdated := self.Args("last_updated").Time()
	iArchived := self.Args("archived").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndLastUpdatedAndArchived(offset, limit, iSiteId,iLastUpdated,iArchived)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndLastUpdatedAndArchived's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndLastUpdatedAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iLastUpdated := self.Args("last_updated").Time()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndLastUpdatedAndMature(offset, limit, iSiteId,iLastUpdated,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndLastUpdatedAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndLastUpdatedAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iLastUpdated := self.Args("last_updated").Time()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndLastUpdatedAndSpam(offset, limit, iSiteId,iLastUpdated,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndLastUpdatedAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndLastUpdatedAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iLastUpdated := self.Args("last_updated").Time()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndLastUpdatedAndDeleted(offset, limit, iSiteId,iLastUpdated,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndLastUpdatedAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndLastUpdatedAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iLastUpdated := self.Args("last_updated").Time()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndLastUpdatedAndLangId(offset, limit, iSiteId,iLastUpdated,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndLastUpdatedAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndPublicAndArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iPublic := self.Args("public").MustInt()
	iArchived := self.Args("archived").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndPublicAndArchived(offset, limit, iSiteId,iPublic,iArchived)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndPublicAndArchived's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndPublicAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iPublic := self.Args("public").MustInt()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndPublicAndMature(offset, limit, iSiteId,iPublic,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndPublicAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndPublicAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iPublic := self.Args("public").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndPublicAndSpam(offset, limit, iSiteId,iPublic,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndPublicAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndPublicAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iPublic := self.Args("public").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndPublicAndDeleted(offset, limit, iSiteId,iPublic,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndPublicAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndPublicAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iPublic := self.Args("public").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndPublicAndLangId(offset, limit, iSiteId,iPublic,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndPublicAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndArchivedAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iArchived := self.Args("archived").MustInt()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndArchivedAndMature(offset, limit, iSiteId,iArchived,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndArchivedAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndArchivedAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iArchived := self.Args("archived").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndArchivedAndSpam(offset, limit, iSiteId,iArchived,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndArchivedAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndArchivedAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iArchived := self.Args("archived").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndArchivedAndDeleted(offset, limit, iSiteId,iArchived,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndArchivedAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndArchivedAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iArchived := self.Args("archived").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndArchivedAndLangId(offset, limit, iSiteId,iArchived,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndArchivedAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndMatureAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iMature := self.Args("mature").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndMatureAndSpam(offset, limit, iSiteId,iMature,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndMatureAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndMatureAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iMature := self.Args("mature").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndMatureAndDeleted(offset, limit, iSiteId,iMature,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndMatureAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndMatureAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iMature := self.Args("mature").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndMatureAndLangId(offset, limit, iSiteId,iMature,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndMatureAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndSpamAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iSpam := self.Args("spam").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndSpamAndDeleted(offset, limit, iSiteId,iSpam,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndSpamAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndSpamAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iSpam := self.Args("spam").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndSpamAndLangId(offset, limit, iSiteId,iSpam,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndSpamAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndDeletedAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iDeleted := self.Args("deleted").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndDeletedAndLangId(offset, limit, iSiteId,iDeleted,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndDeletedAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndPathAndRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iPath := self.Args("path").String()
	iRegistered := self.Args("registered").Time()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndPathAndRegistered(offset, limit, iDomain,iPath,iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndPathAndRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndPathAndLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iPath := self.Args("path").String()
	iLastUpdated := self.Args("last_updated").Time()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndPathAndLastUpdated(offset, limit, iDomain,iPath,iLastUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndPathAndLastUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndPathAndPublicHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iPath := self.Args("path").String()
	iPublic := self.Args("public").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndPathAndPublic(offset, limit, iDomain,iPath,iPublic)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndPathAndPublic's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndPathAndArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iPath := self.Args("path").String()
	iArchived := self.Args("archived").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndPathAndArchived(offset, limit, iDomain,iPath,iArchived)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndPathAndArchived's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndPathAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iPath := self.Args("path").String()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndPathAndMature(offset, limit, iDomain,iPath,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndPathAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndPathAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iPath := self.Args("path").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndPathAndSpam(offset, limit, iDomain,iPath,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndPathAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndPathAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iPath := self.Args("path").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndPathAndDeleted(offset, limit, iDomain,iPath,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndPathAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndPathAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iPath := self.Args("path").String()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndPathAndLangId(offset, limit, iDomain,iPath,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndPathAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndRegisteredAndLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iRegistered := self.Args("registered").Time()
	iLastUpdated := self.Args("last_updated").Time()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndRegisteredAndLastUpdated(offset, limit, iDomain,iRegistered,iLastUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndRegisteredAndLastUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndRegisteredAndPublicHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iRegistered := self.Args("registered").Time()
	iPublic := self.Args("public").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndRegisteredAndPublic(offset, limit, iDomain,iRegistered,iPublic)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndRegisteredAndPublic's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndRegisteredAndArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iRegistered := self.Args("registered").Time()
	iArchived := self.Args("archived").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndRegisteredAndArchived(offset, limit, iDomain,iRegistered,iArchived)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndRegisteredAndArchived's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndRegisteredAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iRegistered := self.Args("registered").Time()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndRegisteredAndMature(offset, limit, iDomain,iRegistered,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndRegisteredAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndRegisteredAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iRegistered := self.Args("registered").Time()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndRegisteredAndSpam(offset, limit, iDomain,iRegistered,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndRegisteredAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndRegisteredAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iRegistered := self.Args("registered").Time()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndRegisteredAndDeleted(offset, limit, iDomain,iRegistered,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndRegisteredAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndRegisteredAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iRegistered := self.Args("registered").Time()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndRegisteredAndLangId(offset, limit, iDomain,iRegistered,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndRegisteredAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndLastUpdatedAndPublicHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iLastUpdated := self.Args("last_updated").Time()
	iPublic := self.Args("public").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndLastUpdatedAndPublic(offset, limit, iDomain,iLastUpdated,iPublic)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndLastUpdatedAndPublic's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndLastUpdatedAndArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iLastUpdated := self.Args("last_updated").Time()
	iArchived := self.Args("archived").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndLastUpdatedAndArchived(offset, limit, iDomain,iLastUpdated,iArchived)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndLastUpdatedAndArchived's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndLastUpdatedAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iLastUpdated := self.Args("last_updated").Time()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndLastUpdatedAndMature(offset, limit, iDomain,iLastUpdated,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndLastUpdatedAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndLastUpdatedAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iLastUpdated := self.Args("last_updated").Time()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndLastUpdatedAndSpam(offset, limit, iDomain,iLastUpdated,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndLastUpdatedAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndLastUpdatedAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iLastUpdated := self.Args("last_updated").Time()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndLastUpdatedAndDeleted(offset, limit, iDomain,iLastUpdated,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndLastUpdatedAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndLastUpdatedAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iLastUpdated := self.Args("last_updated").Time()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndLastUpdatedAndLangId(offset, limit, iDomain,iLastUpdated,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndLastUpdatedAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndPublicAndArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iPublic := self.Args("public").MustInt()
	iArchived := self.Args("archived").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndPublicAndArchived(offset, limit, iDomain,iPublic,iArchived)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndPublicAndArchived's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndPublicAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iPublic := self.Args("public").MustInt()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndPublicAndMature(offset, limit, iDomain,iPublic,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndPublicAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndPublicAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iPublic := self.Args("public").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndPublicAndSpam(offset, limit, iDomain,iPublic,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndPublicAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndPublicAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iPublic := self.Args("public").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndPublicAndDeleted(offset, limit, iDomain,iPublic,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndPublicAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndPublicAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iPublic := self.Args("public").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndPublicAndLangId(offset, limit, iDomain,iPublic,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndPublicAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndArchivedAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iArchived := self.Args("archived").MustInt()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndArchivedAndMature(offset, limit, iDomain,iArchived,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndArchivedAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndArchivedAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iArchived := self.Args("archived").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndArchivedAndSpam(offset, limit, iDomain,iArchived,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndArchivedAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndArchivedAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iArchived := self.Args("archived").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndArchivedAndDeleted(offset, limit, iDomain,iArchived,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndArchivedAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndArchivedAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iArchived := self.Args("archived").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndArchivedAndLangId(offset, limit, iDomain,iArchived,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndArchivedAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndMatureAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iMature := self.Args("mature").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndMatureAndSpam(offset, limit, iDomain,iMature,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndMatureAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndMatureAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iMature := self.Args("mature").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndMatureAndDeleted(offset, limit, iDomain,iMature,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndMatureAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndMatureAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iMature := self.Args("mature").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndMatureAndLangId(offset, limit, iDomain,iMature,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndMatureAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndSpamAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iSpam := self.Args("spam").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndSpamAndDeleted(offset, limit, iDomain,iSpam,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndSpamAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndSpamAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iSpam := self.Args("spam").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndSpamAndLangId(offset, limit, iDomain,iSpam,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndSpamAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndDeletedAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iDeleted := self.Args("deleted").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndDeletedAndLangId(offset, limit, iDomain,iDeleted,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndDeletedAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndRegisteredAndLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iRegistered := self.Args("registered").Time()
	iLastUpdated := self.Args("last_updated").Time()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndRegisteredAndLastUpdated(offset, limit, iPath,iRegistered,iLastUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndRegisteredAndLastUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndRegisteredAndPublicHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iRegistered := self.Args("registered").Time()
	iPublic := self.Args("public").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndRegisteredAndPublic(offset, limit, iPath,iRegistered,iPublic)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndRegisteredAndPublic's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndRegisteredAndArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iRegistered := self.Args("registered").Time()
	iArchived := self.Args("archived").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndRegisteredAndArchived(offset, limit, iPath,iRegistered,iArchived)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndRegisteredAndArchived's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndRegisteredAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iRegistered := self.Args("registered").Time()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndRegisteredAndMature(offset, limit, iPath,iRegistered,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndRegisteredAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndRegisteredAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iRegistered := self.Args("registered").Time()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndRegisteredAndSpam(offset, limit, iPath,iRegistered,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndRegisteredAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndRegisteredAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iRegistered := self.Args("registered").Time()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndRegisteredAndDeleted(offset, limit, iPath,iRegistered,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndRegisteredAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndRegisteredAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iRegistered := self.Args("registered").Time()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndRegisteredAndLangId(offset, limit, iPath,iRegistered,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndRegisteredAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndLastUpdatedAndPublicHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iLastUpdated := self.Args("last_updated").Time()
	iPublic := self.Args("public").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndLastUpdatedAndPublic(offset, limit, iPath,iLastUpdated,iPublic)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndLastUpdatedAndPublic's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndLastUpdatedAndArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iLastUpdated := self.Args("last_updated").Time()
	iArchived := self.Args("archived").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndLastUpdatedAndArchived(offset, limit, iPath,iLastUpdated,iArchived)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndLastUpdatedAndArchived's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndLastUpdatedAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iLastUpdated := self.Args("last_updated").Time()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndLastUpdatedAndMature(offset, limit, iPath,iLastUpdated,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndLastUpdatedAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndLastUpdatedAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iLastUpdated := self.Args("last_updated").Time()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndLastUpdatedAndSpam(offset, limit, iPath,iLastUpdated,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndLastUpdatedAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndLastUpdatedAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iLastUpdated := self.Args("last_updated").Time()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndLastUpdatedAndDeleted(offset, limit, iPath,iLastUpdated,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndLastUpdatedAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndLastUpdatedAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iLastUpdated := self.Args("last_updated").Time()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndLastUpdatedAndLangId(offset, limit, iPath,iLastUpdated,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndLastUpdatedAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndPublicAndArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iPublic := self.Args("public").MustInt()
	iArchived := self.Args("archived").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndPublicAndArchived(offset, limit, iPath,iPublic,iArchived)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndPublicAndArchived's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndPublicAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iPublic := self.Args("public").MustInt()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndPublicAndMature(offset, limit, iPath,iPublic,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndPublicAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndPublicAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iPublic := self.Args("public").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndPublicAndSpam(offset, limit, iPath,iPublic,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndPublicAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndPublicAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iPublic := self.Args("public").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndPublicAndDeleted(offset, limit, iPath,iPublic,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndPublicAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndPublicAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iPublic := self.Args("public").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndPublicAndLangId(offset, limit, iPath,iPublic,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndPublicAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndArchivedAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iArchived := self.Args("archived").MustInt()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndArchivedAndMature(offset, limit, iPath,iArchived,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndArchivedAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndArchivedAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iArchived := self.Args("archived").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndArchivedAndSpam(offset, limit, iPath,iArchived,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndArchivedAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndArchivedAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iArchived := self.Args("archived").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndArchivedAndDeleted(offset, limit, iPath,iArchived,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndArchivedAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndArchivedAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iArchived := self.Args("archived").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndArchivedAndLangId(offset, limit, iPath,iArchived,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndArchivedAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndMatureAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iMature := self.Args("mature").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndMatureAndSpam(offset, limit, iPath,iMature,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndMatureAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndMatureAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iMature := self.Args("mature").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndMatureAndDeleted(offset, limit, iPath,iMature,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndMatureAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndMatureAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iMature := self.Args("mature").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndMatureAndLangId(offset, limit, iPath,iMature,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndMatureAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndSpamAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iSpam := self.Args("spam").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndSpamAndDeleted(offset, limit, iPath,iSpam,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndSpamAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndSpamAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iSpam := self.Args("spam").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndSpamAndLangId(offset, limit, iPath,iSpam,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndSpamAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndDeletedAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iDeleted := self.Args("deleted").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndDeletedAndLangId(offset, limit, iPath,iDeleted,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndDeletedAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByRegisteredAndLastUpdatedAndPublicHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iLastUpdated := self.Args("last_updated").Time()
	iPublic := self.Args("public").MustInt()

	if helper.IsHas(iRegistered) {
		_Blogs, _error := model.GetBlogsesByRegisteredAndLastUpdatedAndPublic(offset, limit, iRegistered,iLastUpdated,iPublic)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByRegisteredAndLastUpdatedAndPublic's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByRegisteredAndLastUpdatedAndArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iLastUpdated := self.Args("last_updated").Time()
	iArchived := self.Args("archived").MustInt()

	if helper.IsHas(iRegistered) {
		_Blogs, _error := model.GetBlogsesByRegisteredAndLastUpdatedAndArchived(offset, limit, iRegistered,iLastUpdated,iArchived)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByRegisteredAndLastUpdatedAndArchived's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByRegisteredAndLastUpdatedAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iLastUpdated := self.Args("last_updated").Time()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iRegistered) {
		_Blogs, _error := model.GetBlogsesByRegisteredAndLastUpdatedAndMature(offset, limit, iRegistered,iLastUpdated,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByRegisteredAndLastUpdatedAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByRegisteredAndLastUpdatedAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iLastUpdated := self.Args("last_updated").Time()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iRegistered) {
		_Blogs, _error := model.GetBlogsesByRegisteredAndLastUpdatedAndSpam(offset, limit, iRegistered,iLastUpdated,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByRegisteredAndLastUpdatedAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByRegisteredAndLastUpdatedAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iLastUpdated := self.Args("last_updated").Time()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iRegistered) {
		_Blogs, _error := model.GetBlogsesByRegisteredAndLastUpdatedAndDeleted(offset, limit, iRegistered,iLastUpdated,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByRegisteredAndLastUpdatedAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByRegisteredAndLastUpdatedAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iLastUpdated := self.Args("last_updated").Time()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iRegistered) {
		_Blogs, _error := model.GetBlogsesByRegisteredAndLastUpdatedAndLangId(offset, limit, iRegistered,iLastUpdated,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByRegisteredAndLastUpdatedAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByRegisteredAndPublicAndArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iPublic := self.Args("public").MustInt()
	iArchived := self.Args("archived").MustInt()

	if helper.IsHas(iRegistered) {
		_Blogs, _error := model.GetBlogsesByRegisteredAndPublicAndArchived(offset, limit, iRegistered,iPublic,iArchived)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByRegisteredAndPublicAndArchived's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByRegisteredAndPublicAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iPublic := self.Args("public").MustInt()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iRegistered) {
		_Blogs, _error := model.GetBlogsesByRegisteredAndPublicAndMature(offset, limit, iRegistered,iPublic,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByRegisteredAndPublicAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByRegisteredAndPublicAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iPublic := self.Args("public").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iRegistered) {
		_Blogs, _error := model.GetBlogsesByRegisteredAndPublicAndSpam(offset, limit, iRegistered,iPublic,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByRegisteredAndPublicAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByRegisteredAndPublicAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iPublic := self.Args("public").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iRegistered) {
		_Blogs, _error := model.GetBlogsesByRegisteredAndPublicAndDeleted(offset, limit, iRegistered,iPublic,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByRegisteredAndPublicAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByRegisteredAndPublicAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iPublic := self.Args("public").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iRegistered) {
		_Blogs, _error := model.GetBlogsesByRegisteredAndPublicAndLangId(offset, limit, iRegistered,iPublic,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByRegisteredAndPublicAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByRegisteredAndArchivedAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iArchived := self.Args("archived").MustInt()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iRegistered) {
		_Blogs, _error := model.GetBlogsesByRegisteredAndArchivedAndMature(offset, limit, iRegistered,iArchived,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByRegisteredAndArchivedAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByRegisteredAndArchivedAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iArchived := self.Args("archived").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iRegistered) {
		_Blogs, _error := model.GetBlogsesByRegisteredAndArchivedAndSpam(offset, limit, iRegistered,iArchived,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByRegisteredAndArchivedAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByRegisteredAndArchivedAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iArchived := self.Args("archived").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iRegistered) {
		_Blogs, _error := model.GetBlogsesByRegisteredAndArchivedAndDeleted(offset, limit, iRegistered,iArchived,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByRegisteredAndArchivedAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByRegisteredAndArchivedAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iArchived := self.Args("archived").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iRegistered) {
		_Blogs, _error := model.GetBlogsesByRegisteredAndArchivedAndLangId(offset, limit, iRegistered,iArchived,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByRegisteredAndArchivedAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByRegisteredAndMatureAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iMature := self.Args("mature").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iRegistered) {
		_Blogs, _error := model.GetBlogsesByRegisteredAndMatureAndSpam(offset, limit, iRegistered,iMature,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByRegisteredAndMatureAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByRegisteredAndMatureAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iMature := self.Args("mature").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iRegistered) {
		_Blogs, _error := model.GetBlogsesByRegisteredAndMatureAndDeleted(offset, limit, iRegistered,iMature,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByRegisteredAndMatureAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByRegisteredAndMatureAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iMature := self.Args("mature").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iRegistered) {
		_Blogs, _error := model.GetBlogsesByRegisteredAndMatureAndLangId(offset, limit, iRegistered,iMature,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByRegisteredAndMatureAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByRegisteredAndSpamAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iSpam := self.Args("spam").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iRegistered) {
		_Blogs, _error := model.GetBlogsesByRegisteredAndSpamAndDeleted(offset, limit, iRegistered,iSpam,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByRegisteredAndSpamAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByRegisteredAndSpamAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iSpam := self.Args("spam").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iRegistered) {
		_Blogs, _error := model.GetBlogsesByRegisteredAndSpamAndLangId(offset, limit, iRegistered,iSpam,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByRegisteredAndSpamAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByRegisteredAndDeletedAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iDeleted := self.Args("deleted").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iRegistered) {
		_Blogs, _error := model.GetBlogsesByRegisteredAndDeletedAndLangId(offset, limit, iRegistered,iDeleted,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByRegisteredAndDeletedAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByLastUpdatedAndPublicAndArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLastUpdated := self.Args("last_updated").Time()
	iPublic := self.Args("public").MustInt()
	iArchived := self.Args("archived").MustInt()

	if helper.IsHas(iLastUpdated) {
		_Blogs, _error := model.GetBlogsesByLastUpdatedAndPublicAndArchived(offset, limit, iLastUpdated,iPublic,iArchived)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByLastUpdatedAndPublicAndArchived's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByLastUpdatedAndPublicAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLastUpdated := self.Args("last_updated").Time()
	iPublic := self.Args("public").MustInt()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iLastUpdated) {
		_Blogs, _error := model.GetBlogsesByLastUpdatedAndPublicAndMature(offset, limit, iLastUpdated,iPublic,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByLastUpdatedAndPublicAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByLastUpdatedAndPublicAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLastUpdated := self.Args("last_updated").Time()
	iPublic := self.Args("public").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iLastUpdated) {
		_Blogs, _error := model.GetBlogsesByLastUpdatedAndPublicAndSpam(offset, limit, iLastUpdated,iPublic,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByLastUpdatedAndPublicAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByLastUpdatedAndPublicAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLastUpdated := self.Args("last_updated").Time()
	iPublic := self.Args("public").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iLastUpdated) {
		_Blogs, _error := model.GetBlogsesByLastUpdatedAndPublicAndDeleted(offset, limit, iLastUpdated,iPublic,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByLastUpdatedAndPublicAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByLastUpdatedAndPublicAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLastUpdated := self.Args("last_updated").Time()
	iPublic := self.Args("public").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iLastUpdated) {
		_Blogs, _error := model.GetBlogsesByLastUpdatedAndPublicAndLangId(offset, limit, iLastUpdated,iPublic,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByLastUpdatedAndPublicAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByLastUpdatedAndArchivedAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLastUpdated := self.Args("last_updated").Time()
	iArchived := self.Args("archived").MustInt()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iLastUpdated) {
		_Blogs, _error := model.GetBlogsesByLastUpdatedAndArchivedAndMature(offset, limit, iLastUpdated,iArchived,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByLastUpdatedAndArchivedAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByLastUpdatedAndArchivedAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLastUpdated := self.Args("last_updated").Time()
	iArchived := self.Args("archived").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iLastUpdated) {
		_Blogs, _error := model.GetBlogsesByLastUpdatedAndArchivedAndSpam(offset, limit, iLastUpdated,iArchived,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByLastUpdatedAndArchivedAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByLastUpdatedAndArchivedAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLastUpdated := self.Args("last_updated").Time()
	iArchived := self.Args("archived").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iLastUpdated) {
		_Blogs, _error := model.GetBlogsesByLastUpdatedAndArchivedAndDeleted(offset, limit, iLastUpdated,iArchived,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByLastUpdatedAndArchivedAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByLastUpdatedAndArchivedAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLastUpdated := self.Args("last_updated").Time()
	iArchived := self.Args("archived").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iLastUpdated) {
		_Blogs, _error := model.GetBlogsesByLastUpdatedAndArchivedAndLangId(offset, limit, iLastUpdated,iArchived,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByLastUpdatedAndArchivedAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByLastUpdatedAndMatureAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLastUpdated := self.Args("last_updated").Time()
	iMature := self.Args("mature").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iLastUpdated) {
		_Blogs, _error := model.GetBlogsesByLastUpdatedAndMatureAndSpam(offset, limit, iLastUpdated,iMature,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByLastUpdatedAndMatureAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByLastUpdatedAndMatureAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLastUpdated := self.Args("last_updated").Time()
	iMature := self.Args("mature").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iLastUpdated) {
		_Blogs, _error := model.GetBlogsesByLastUpdatedAndMatureAndDeleted(offset, limit, iLastUpdated,iMature,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByLastUpdatedAndMatureAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByLastUpdatedAndMatureAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLastUpdated := self.Args("last_updated").Time()
	iMature := self.Args("mature").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iLastUpdated) {
		_Blogs, _error := model.GetBlogsesByLastUpdatedAndMatureAndLangId(offset, limit, iLastUpdated,iMature,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByLastUpdatedAndMatureAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByLastUpdatedAndSpamAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLastUpdated := self.Args("last_updated").Time()
	iSpam := self.Args("spam").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iLastUpdated) {
		_Blogs, _error := model.GetBlogsesByLastUpdatedAndSpamAndDeleted(offset, limit, iLastUpdated,iSpam,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByLastUpdatedAndSpamAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByLastUpdatedAndSpamAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLastUpdated := self.Args("last_updated").Time()
	iSpam := self.Args("spam").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iLastUpdated) {
		_Blogs, _error := model.GetBlogsesByLastUpdatedAndSpamAndLangId(offset, limit, iLastUpdated,iSpam,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByLastUpdatedAndSpamAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByLastUpdatedAndDeletedAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLastUpdated := self.Args("last_updated").Time()
	iDeleted := self.Args("deleted").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iLastUpdated) {
		_Blogs, _error := model.GetBlogsesByLastUpdatedAndDeletedAndLangId(offset, limit, iLastUpdated,iDeleted,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByLastUpdatedAndDeletedAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPublicAndArchivedAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPublic := self.Args("public").MustInt()
	iArchived := self.Args("archived").MustInt()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iPublic) {
		_Blogs, _error := model.GetBlogsesByPublicAndArchivedAndMature(offset, limit, iPublic,iArchived,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPublicAndArchivedAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPublicAndArchivedAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPublic := self.Args("public").MustInt()
	iArchived := self.Args("archived").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iPublic) {
		_Blogs, _error := model.GetBlogsesByPublicAndArchivedAndSpam(offset, limit, iPublic,iArchived,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPublicAndArchivedAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPublicAndArchivedAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPublic := self.Args("public").MustInt()
	iArchived := self.Args("archived").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iPublic) {
		_Blogs, _error := model.GetBlogsesByPublicAndArchivedAndDeleted(offset, limit, iPublic,iArchived,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPublicAndArchivedAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPublicAndArchivedAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPublic := self.Args("public").MustInt()
	iArchived := self.Args("archived").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iPublic) {
		_Blogs, _error := model.GetBlogsesByPublicAndArchivedAndLangId(offset, limit, iPublic,iArchived,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPublicAndArchivedAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPublicAndMatureAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPublic := self.Args("public").MustInt()
	iMature := self.Args("mature").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iPublic) {
		_Blogs, _error := model.GetBlogsesByPublicAndMatureAndSpam(offset, limit, iPublic,iMature,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPublicAndMatureAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPublicAndMatureAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPublic := self.Args("public").MustInt()
	iMature := self.Args("mature").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iPublic) {
		_Blogs, _error := model.GetBlogsesByPublicAndMatureAndDeleted(offset, limit, iPublic,iMature,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPublicAndMatureAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPublicAndMatureAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPublic := self.Args("public").MustInt()
	iMature := self.Args("mature").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iPublic) {
		_Blogs, _error := model.GetBlogsesByPublicAndMatureAndLangId(offset, limit, iPublic,iMature,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPublicAndMatureAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPublicAndSpamAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPublic := self.Args("public").MustInt()
	iSpam := self.Args("spam").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iPublic) {
		_Blogs, _error := model.GetBlogsesByPublicAndSpamAndDeleted(offset, limit, iPublic,iSpam,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPublicAndSpamAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPublicAndSpamAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPublic := self.Args("public").MustInt()
	iSpam := self.Args("spam").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iPublic) {
		_Blogs, _error := model.GetBlogsesByPublicAndSpamAndLangId(offset, limit, iPublic,iSpam,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPublicAndSpamAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPublicAndDeletedAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPublic := self.Args("public").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iPublic) {
		_Blogs, _error := model.GetBlogsesByPublicAndDeletedAndLangId(offset, limit, iPublic,iDeleted,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPublicAndDeletedAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByArchivedAndMatureAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iArchived := self.Args("archived").MustInt()
	iMature := self.Args("mature").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iArchived) {
		_Blogs, _error := model.GetBlogsesByArchivedAndMatureAndSpam(offset, limit, iArchived,iMature,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByArchivedAndMatureAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByArchivedAndMatureAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iArchived := self.Args("archived").MustInt()
	iMature := self.Args("mature").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iArchived) {
		_Blogs, _error := model.GetBlogsesByArchivedAndMatureAndDeleted(offset, limit, iArchived,iMature,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByArchivedAndMatureAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByArchivedAndMatureAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iArchived := self.Args("archived").MustInt()
	iMature := self.Args("mature").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iArchived) {
		_Blogs, _error := model.GetBlogsesByArchivedAndMatureAndLangId(offset, limit, iArchived,iMature,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByArchivedAndMatureAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByArchivedAndSpamAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iArchived := self.Args("archived").MustInt()
	iSpam := self.Args("spam").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iArchived) {
		_Blogs, _error := model.GetBlogsesByArchivedAndSpamAndDeleted(offset, limit, iArchived,iSpam,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByArchivedAndSpamAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByArchivedAndSpamAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iArchived := self.Args("archived").MustInt()
	iSpam := self.Args("spam").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iArchived) {
		_Blogs, _error := model.GetBlogsesByArchivedAndSpamAndLangId(offset, limit, iArchived,iSpam,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByArchivedAndSpamAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByArchivedAndDeletedAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iArchived := self.Args("archived").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iArchived) {
		_Blogs, _error := model.GetBlogsesByArchivedAndDeletedAndLangId(offset, limit, iArchived,iDeleted,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByArchivedAndDeletedAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByMatureAndSpamAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMature := self.Args("mature").MustInt()
	iSpam := self.Args("spam").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iMature) {
		_Blogs, _error := model.GetBlogsesByMatureAndSpamAndDeleted(offset, limit, iMature,iSpam,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByMatureAndSpamAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByMatureAndSpamAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMature := self.Args("mature").MustInt()
	iSpam := self.Args("spam").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iMature) {
		_Blogs, _error := model.GetBlogsesByMatureAndSpamAndLangId(offset, limit, iMature,iSpam,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByMatureAndSpamAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByMatureAndDeletedAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMature := self.Args("mature").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iMature) {
		_Blogs, _error := model.GetBlogsesByMatureAndDeletedAndLangId(offset, limit, iMature,iDeleted,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByMatureAndDeletedAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySpamAndDeletedAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSpam := self.Args("spam").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iSpam) {
		_Blogs, _error := model.GetBlogsesBySpamAndDeletedAndLangId(offset, limit, iSpam,iDeleted,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySpamAndDeletedAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndSiteIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iSiteId := self.Args("site_id").MustInt64()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndSiteId(offset, limit, iBlogId,iSiteId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndSiteId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iDomain := self.Args("domain").String()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndDomain(offset, limit, iBlogId,iDomain)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndDomain's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iPath := self.Args("path").String()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndPath(offset, limit, iBlogId,iPath)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndPath's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iRegistered := self.Args("registered").Time()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndRegistered(offset, limit, iBlogId,iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iLastUpdated := self.Args("last_updated").Time()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndLastUpdated(offset, limit, iBlogId,iLastUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndLastUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndPublicHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iPublic := self.Args("public").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndPublic(offset, limit, iBlogId,iPublic)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndPublic's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iArchived := self.Args("archived").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndArchived(offset, limit, iBlogId,iArchived)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndArchived's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndMature(offset, limit, iBlogId,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndSpam(offset, limit, iBlogId,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndDeleted(offset, limit, iBlogId,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByBlogIdAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsesByBlogIdAndLangId(offset, limit, iBlogId,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByBlogIdAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iDomain := self.Args("domain").String()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndDomain(offset, limit, iSiteId,iDomain)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndDomain's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iPath := self.Args("path").String()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndPath(offset, limit, iSiteId,iPath)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndPath's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iRegistered := self.Args("registered").Time()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndRegistered(offset, limit, iSiteId,iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iLastUpdated := self.Args("last_updated").Time()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndLastUpdated(offset, limit, iSiteId,iLastUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndLastUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndPublicHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iPublic := self.Args("public").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndPublic(offset, limit, iSiteId,iPublic)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndPublic's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iArchived := self.Args("archived").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndArchived(offset, limit, iSiteId,iArchived)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndArchived's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndMature(offset, limit, iSiteId,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndSpam(offset, limit, iSiteId,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndDeleted(offset, limit, iSiteId,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySiteIdAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsesBySiteIdAndLangId(offset, limit, iSiteId,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySiteIdAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iPath := self.Args("path").String()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndPath(offset, limit, iDomain,iPath)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndPath's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iRegistered := self.Args("registered").Time()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndRegistered(offset, limit, iDomain,iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iLastUpdated := self.Args("last_updated").Time()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndLastUpdated(offset, limit, iDomain,iLastUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndLastUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndPublicHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iPublic := self.Args("public").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndPublic(offset, limit, iDomain,iPublic)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndPublic's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iArchived := self.Args("archived").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndArchived(offset, limit, iDomain,iArchived)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndArchived's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndMature(offset, limit, iDomain,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndSpam(offset, limit, iDomain,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndDeleted(offset, limit, iDomain,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDomainAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsesByDomainAndLangId(offset, limit, iDomain,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDomainAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iRegistered := self.Args("registered").Time()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndRegistered(offset, limit, iPath,iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iLastUpdated := self.Args("last_updated").Time()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndLastUpdated(offset, limit, iPath,iLastUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndLastUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndPublicHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iPublic := self.Args("public").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndPublic(offset, limit, iPath,iPublic)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndPublic's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iArchived := self.Args("archived").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndArchived(offset, limit, iPath,iArchived)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndArchived's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndMature(offset, limit, iPath,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndSpam(offset, limit, iPath,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndDeleted(offset, limit, iPath,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPathAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsesByPathAndLangId(offset, limit, iPath,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPathAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByRegisteredAndLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iLastUpdated := self.Args("last_updated").Time()

	if helper.IsHas(iRegistered) {
		_Blogs, _error := model.GetBlogsesByRegisteredAndLastUpdated(offset, limit, iRegistered,iLastUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByRegisteredAndLastUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByRegisteredAndPublicHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iPublic := self.Args("public").MustInt()

	if helper.IsHas(iRegistered) {
		_Blogs, _error := model.GetBlogsesByRegisteredAndPublic(offset, limit, iRegistered,iPublic)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByRegisteredAndPublic's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByRegisteredAndArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iArchived := self.Args("archived").MustInt()

	if helper.IsHas(iRegistered) {
		_Blogs, _error := model.GetBlogsesByRegisteredAndArchived(offset, limit, iRegistered,iArchived)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByRegisteredAndArchived's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByRegisteredAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iRegistered) {
		_Blogs, _error := model.GetBlogsesByRegisteredAndMature(offset, limit, iRegistered,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByRegisteredAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByRegisteredAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iRegistered) {
		_Blogs, _error := model.GetBlogsesByRegisteredAndSpam(offset, limit, iRegistered,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByRegisteredAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByRegisteredAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iRegistered) {
		_Blogs, _error := model.GetBlogsesByRegisteredAndDeleted(offset, limit, iRegistered,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByRegisteredAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByRegisteredAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iRegistered) {
		_Blogs, _error := model.GetBlogsesByRegisteredAndLangId(offset, limit, iRegistered,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByRegisteredAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByLastUpdatedAndPublicHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLastUpdated := self.Args("last_updated").Time()
	iPublic := self.Args("public").MustInt()

	if helper.IsHas(iLastUpdated) {
		_Blogs, _error := model.GetBlogsesByLastUpdatedAndPublic(offset, limit, iLastUpdated,iPublic)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByLastUpdatedAndPublic's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByLastUpdatedAndArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLastUpdated := self.Args("last_updated").Time()
	iArchived := self.Args("archived").MustInt()

	if helper.IsHas(iLastUpdated) {
		_Blogs, _error := model.GetBlogsesByLastUpdatedAndArchived(offset, limit, iLastUpdated,iArchived)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByLastUpdatedAndArchived's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByLastUpdatedAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLastUpdated := self.Args("last_updated").Time()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iLastUpdated) {
		_Blogs, _error := model.GetBlogsesByLastUpdatedAndMature(offset, limit, iLastUpdated,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByLastUpdatedAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByLastUpdatedAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLastUpdated := self.Args("last_updated").Time()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iLastUpdated) {
		_Blogs, _error := model.GetBlogsesByLastUpdatedAndSpam(offset, limit, iLastUpdated,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByLastUpdatedAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByLastUpdatedAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLastUpdated := self.Args("last_updated").Time()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iLastUpdated) {
		_Blogs, _error := model.GetBlogsesByLastUpdatedAndDeleted(offset, limit, iLastUpdated,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByLastUpdatedAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByLastUpdatedAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iLastUpdated := self.Args("last_updated").Time()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iLastUpdated) {
		_Blogs, _error := model.GetBlogsesByLastUpdatedAndLangId(offset, limit, iLastUpdated,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByLastUpdatedAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPublicAndArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPublic := self.Args("public").MustInt()
	iArchived := self.Args("archived").MustInt()

	if helper.IsHas(iPublic) {
		_Blogs, _error := model.GetBlogsesByPublicAndArchived(offset, limit, iPublic,iArchived)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPublicAndArchived's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPublicAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPublic := self.Args("public").MustInt()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iPublic) {
		_Blogs, _error := model.GetBlogsesByPublicAndMature(offset, limit, iPublic,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPublicAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPublicAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPublic := self.Args("public").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iPublic) {
		_Blogs, _error := model.GetBlogsesByPublicAndSpam(offset, limit, iPublic,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPublicAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPublicAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPublic := self.Args("public").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iPublic) {
		_Blogs, _error := model.GetBlogsesByPublicAndDeleted(offset, limit, iPublic,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPublicAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByPublicAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPublic := self.Args("public").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iPublic) {
		_Blogs, _error := model.GetBlogsesByPublicAndLangId(offset, limit, iPublic,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByPublicAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByArchivedAndMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iArchived := self.Args("archived").MustInt()
	iMature := self.Args("mature").MustInt()

	if helper.IsHas(iArchived) {
		_Blogs, _error := model.GetBlogsesByArchivedAndMature(offset, limit, iArchived,iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByArchivedAndMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByArchivedAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iArchived := self.Args("archived").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iArchived) {
		_Blogs, _error := model.GetBlogsesByArchivedAndSpam(offset, limit, iArchived,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByArchivedAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByArchivedAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iArchived := self.Args("archived").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iArchived) {
		_Blogs, _error := model.GetBlogsesByArchivedAndDeleted(offset, limit, iArchived,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByArchivedAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByArchivedAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iArchived := self.Args("archived").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iArchived) {
		_Blogs, _error := model.GetBlogsesByArchivedAndLangId(offset, limit, iArchived,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByArchivedAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByMatureAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMature := self.Args("mature").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iMature) {
		_Blogs, _error := model.GetBlogsesByMatureAndSpam(offset, limit, iMature,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByMatureAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByMatureAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMature := self.Args("mature").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iMature) {
		_Blogs, _error := model.GetBlogsesByMatureAndDeleted(offset, limit, iMature,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByMatureAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByMatureAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMature := self.Args("mature").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iMature) {
		_Blogs, _error := model.GetBlogsesByMatureAndLangId(offset, limit, iMature,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByMatureAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySpamAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSpam := self.Args("spam").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iSpam) {
		_Blogs, _error := model.GetBlogsesBySpamAndDeleted(offset, limit, iSpam,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySpamAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesBySpamAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSpam := self.Args("spam").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iSpam) {
		_Blogs, _error := model.GetBlogsesBySpamAndLangId(offset, limit, iSpam,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesBySpamAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesByDeletedAndLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDeleted := self.Args("deleted").MustInt()
	iLangId := self.Args("lang_id").MustInt()

	if helper.IsHas(iDeleted) {
		_Blogs, _error := model.GetBlogsesByDeletedAndLangId(offset, limit, iDeleted,iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsesByDeletedAndLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_Blogs, _error := model.GetBlogses(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogses' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlogsByBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBlogId := self.Args("blog_id").MustInt64()
	if helper.IsHas(iBlogId) {
		_Blogs := model.HasBlogsByBlogId(iBlogId)
		var m = map[string]interface{}{}
		m["blogs"] = _Blogs
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlogsByBlogId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlogsBySiteIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSiteId := self.Args("site_id").MustInt64()
	if helper.IsHas(iSiteId) {
		_Blogs := model.HasBlogsBySiteId(iSiteId)
		var m = map[string]interface{}{}
		m["blogs"] = _Blogs
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlogsBySiteId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlogsByDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDomain := self.Args("domain").String()
	if helper.IsHas(iDomain) {
		_Blogs := model.HasBlogsByDomain(iDomain)
		var m = map[string]interface{}{}
		m["blogs"] = _Blogs
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlogsByDomain's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlogsByPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iPath := self.Args("path").String()
	if helper.IsHas(iPath) {
		_Blogs := model.HasBlogsByPath(iPath)
		var m = map[string]interface{}{}
		m["blogs"] = _Blogs
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlogsByPath's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlogsByRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRegistered := self.Args("registered").Time()
	if helper.IsHas(iRegistered) {
		_Blogs := model.HasBlogsByRegistered(iRegistered)
		var m = map[string]interface{}{}
		m["blogs"] = _Blogs
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlogsByRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlogsByLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLastUpdated := self.Args("last_updated").Time()
	if helper.IsHas(iLastUpdated) {
		_Blogs := model.HasBlogsByLastUpdated(iLastUpdated)
		var m = map[string]interface{}{}
		m["blogs"] = _Blogs
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlogsByLastUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlogsByPublicHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iPublic := self.Args("public").MustInt()
	if helper.IsHas(iPublic) {
		_Blogs := model.HasBlogsByPublic(iPublic)
		var m = map[string]interface{}{}
		m["blogs"] = _Blogs
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlogsByPublic's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlogsByArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iArchived := self.Args("archived").MustInt()
	if helper.IsHas(iArchived) {
		_Blogs := model.HasBlogsByArchived(iArchived)
		var m = map[string]interface{}{}
		m["blogs"] = _Blogs
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlogsByArchived's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlogsByMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMature := self.Args("mature").MustInt()
	if helper.IsHas(iMature) {
		_Blogs := model.HasBlogsByMature(iMature)
		var m = map[string]interface{}{}
		m["blogs"] = _Blogs
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlogsByMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlogsBySpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSpam := self.Args("spam").MustInt()
	if helper.IsHas(iSpam) {
		_Blogs := model.HasBlogsBySpam(iSpam)
		var m = map[string]interface{}{}
		m["blogs"] = _Blogs
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlogsBySpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlogsByDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDeleted := self.Args("deleted").MustInt()
	if helper.IsHas(iDeleted) {
		_Blogs := model.HasBlogsByDeleted(iDeleted)
		var m = map[string]interface{}{}
		m["blogs"] = _Blogs
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlogsByDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlogsByLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangId := self.Args("lang_id").MustInt()
	if helper.IsHas(iLangId) {
		_Blogs := model.HasBlogsByLangId(iLangId)
		var m = map[string]interface{}{}
		m["blogs"] = _Blogs
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlogsByLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsByBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBlogId := self.Args("blog_id").MustInt64()
	if helper.IsHas(iBlogId) {
		_Blogs, _error := model.GetBlogsByBlogId(iBlogId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsByBlogId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsBySiteIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSiteId := self.Args("site_id").MustInt64()
	if helper.IsHas(iSiteId) {
		_Blogs, _error := model.GetBlogsBySiteId(iSiteId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsBySiteId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsByDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDomain := self.Args("domain").String()
	if helper.IsHas(iDomain) {
		_Blogs, _error := model.GetBlogsByDomain(iDomain)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsByDomain's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsByPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iPath := self.Args("path").String()
	if helper.IsHas(iPath) {
		_Blogs, _error := model.GetBlogsByPath(iPath)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsByPath's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsByRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRegistered := self.Args("registered").Time()
	if helper.IsHas(iRegistered) {
		_Blogs, _error := model.GetBlogsByRegistered(iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsByRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsByLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLastUpdated := self.Args("last_updated").Time()
	if helper.IsHas(iLastUpdated) {
		_Blogs, _error := model.GetBlogsByLastUpdated(iLastUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsByLastUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsByPublicHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iPublic := self.Args("public").MustInt()
	if helper.IsHas(iPublic) {
		_Blogs, _error := model.GetBlogsByPublic(iPublic)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsByPublic's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsByArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iArchived := self.Args("archived").MustInt()
	if helper.IsHas(iArchived) {
		_Blogs, _error := model.GetBlogsByArchived(iArchived)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsByArchived's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsByMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMature := self.Args("mature").MustInt()
	if helper.IsHas(iMature) {
		_Blogs, _error := model.GetBlogsByMature(iMature)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsByMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsBySpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSpam := self.Args("spam").MustInt()
	if helper.IsHas(iSpam) {
		_Blogs, _error := model.GetBlogsBySpam(iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsBySpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsByDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDeleted := self.Args("deleted").MustInt()
	if helper.IsHas(iDeleted) {
		_Blogs, _error := model.GetBlogsByDeleted(iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsByDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogsByLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLangId := self.Args("lang_id").MustInt()
	if helper.IsHas(iLangId) {
		_Blogs, _error := model.GetBlogsByLangId(iLangId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the GetBlogsByLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlogsByBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BlogId_ := self.Args("blog_id").MustInt64()
	if helper.IsHas(BlogId_) {
		var iBlogs model.Blogs
		self.Bind(&iBlogs)
		_Blogs, _error := model.SetBlogsByBlogId(BlogId_, &iBlogs)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the SetBlogsByBlogId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlogsBySiteIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	SiteId_ := self.Args("site_id").MustInt64()
	if helper.IsHas(SiteId_) {
		var iBlogs model.Blogs
		self.Bind(&iBlogs)
		_Blogs, _error := model.SetBlogsBySiteId(SiteId_, &iBlogs)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the SetBlogsBySiteId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlogsByDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Domain_ := self.Args("domain").String()
	if helper.IsHas(Domain_) {
		var iBlogs model.Blogs
		self.Bind(&iBlogs)
		_Blogs, _error := model.SetBlogsByDomain(Domain_, &iBlogs)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the SetBlogsByDomain's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlogsByPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Path_ := self.Args("path").String()
	if helper.IsHas(Path_) {
		var iBlogs model.Blogs
		self.Bind(&iBlogs)
		_Blogs, _error := model.SetBlogsByPath(Path_, &iBlogs)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the SetBlogsByPath's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlogsByRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Registered_ := self.Args("registered").Time()
	if helper.IsHas(Registered_) {
		var iBlogs model.Blogs
		self.Bind(&iBlogs)
		_Blogs, _error := model.SetBlogsByRegistered(Registered_, &iBlogs)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the SetBlogsByRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlogsByLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LastUpdated_ := self.Args("last_updated").Time()
	if helper.IsHas(LastUpdated_) {
		var iBlogs model.Blogs
		self.Bind(&iBlogs)
		_Blogs, _error := model.SetBlogsByLastUpdated(LastUpdated_, &iBlogs)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the SetBlogsByLastUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlogsByPublicHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Public_ := self.Args("public").MustInt()
	if helper.IsHas(Public_) {
		var iBlogs model.Blogs
		self.Bind(&iBlogs)
		_Blogs, _error := model.SetBlogsByPublic(Public_, &iBlogs)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the SetBlogsByPublic's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlogsByArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Archived_ := self.Args("archived").MustInt()
	if helper.IsHas(Archived_) {
		var iBlogs model.Blogs
		self.Bind(&iBlogs)
		_Blogs, _error := model.SetBlogsByArchived(Archived_, &iBlogs)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the SetBlogsByArchived's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlogsByMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Mature_ := self.Args("mature").MustInt()
	if helper.IsHas(Mature_) {
		var iBlogs model.Blogs
		self.Bind(&iBlogs)
		_Blogs, _error := model.SetBlogsByMature(Mature_, &iBlogs)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the SetBlogsByMature's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlogsBySpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Spam_ := self.Args("spam").MustInt()
	if helper.IsHas(Spam_) {
		var iBlogs model.Blogs
		self.Bind(&iBlogs)
		_Blogs, _error := model.SetBlogsBySpam(Spam_, &iBlogs)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the SetBlogsBySpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlogsByDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	if helper.IsHas(Deleted_) {
		var iBlogs model.Blogs
		self.Bind(&iBlogs)
		_Blogs, _error := model.SetBlogsByDeleted(Deleted_, &iBlogs)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the SetBlogsByDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlogsByLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LangId_ := self.Args("lang_id").MustInt()
	if helper.IsHas(LangId_) {
		var iBlogs model.Blogs
		self.Bind(&iBlogs)
		_Blogs, _error := model.SetBlogsByLangId(LangId_, &iBlogs)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Blogs)
	}
	herr.Message = "Can't get to the SetBlogsByLangId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddBlogsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BlogId_ := self.Args("blog_id").MustInt64()
	SiteId_ := self.Args("site_id").MustInt64()
	Domain_ := self.Args("domain").String()
	Path_ := self.Args("path").String()
	Registered_ := self.Args("registered").Time()
	LastUpdated_ := self.Args("last_updated").Time()
	Public_ := self.Args("public").MustInt()
	Archived_ := self.Args("archived").MustInt()
	Mature_ := self.Args("mature").MustInt()
	Spam_ := self.Args("spam").MustInt()
	Deleted_ := self.Args("deleted").MustInt()
	LangId_ := self.Args("lang_id").MustInt()

	if helper.IsHas(BlogId_) {
		_error := model.AddBlogs(BlogId_,SiteId_,Domain_,Path_,Registered_,LastUpdated_,Public_,Archived_,Mature_,Spam_,Deleted_,LangId_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddBlogs's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostBlogsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iBlogs model.Blogs
	self.Bind(&iBlogs)
	_int64, _error := model.PostBlogs(&iBlogs)
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

func PutBlogsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iBlogs model.Blogs
	self.Bind(&iBlogs)
	_int64, _error := model.PutBlogs(&iBlogs)
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

func PutBlogsByBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BlogId_ := self.Args("blog_id").MustInt64()
	var iBlogs model.Blogs
	self.Bind(&iBlogs)
	_int64, _error := model.PutBlogsByBlogId(BlogId_, &iBlogs)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlogsBySiteIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	SiteId_ := self.Args("site_id").MustInt64()
	var iBlogs model.Blogs
	self.Bind(&iBlogs)
	_int64, _error := model.PutBlogsBySiteId(SiteId_, &iBlogs)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlogsByDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Domain_ := self.Args("domain").String()
	var iBlogs model.Blogs
	self.Bind(&iBlogs)
	_int64, _error := model.PutBlogsByDomain(Domain_, &iBlogs)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlogsByPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Path_ := self.Args("path").String()
	var iBlogs model.Blogs
	self.Bind(&iBlogs)
	_int64, _error := model.PutBlogsByPath(Path_, &iBlogs)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlogsByRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Registered_ := self.Args("registered").Time()
	var iBlogs model.Blogs
	self.Bind(&iBlogs)
	_int64, _error := model.PutBlogsByRegistered(Registered_, &iBlogs)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlogsByLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LastUpdated_ := self.Args("last_updated").Time()
	var iBlogs model.Blogs
	self.Bind(&iBlogs)
	_int64, _error := model.PutBlogsByLastUpdated(LastUpdated_, &iBlogs)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlogsByPublicHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Public_ := self.Args("public").MustInt()
	var iBlogs model.Blogs
	self.Bind(&iBlogs)
	_int64, _error := model.PutBlogsByPublic(Public_, &iBlogs)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlogsByArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Archived_ := self.Args("archived").MustInt()
	var iBlogs model.Blogs
	self.Bind(&iBlogs)
	_int64, _error := model.PutBlogsByArchived(Archived_, &iBlogs)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlogsByMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Mature_ := self.Args("mature").MustInt()
	var iBlogs model.Blogs
	self.Bind(&iBlogs)
	_int64, _error := model.PutBlogsByMature(Mature_, &iBlogs)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlogsBySpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Spam_ := self.Args("spam").MustInt()
	var iBlogs model.Blogs
	self.Bind(&iBlogs)
	_int64, _error := model.PutBlogsBySpam(Spam_, &iBlogs)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlogsByDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	var iBlogs model.Blogs
	self.Bind(&iBlogs)
	_int64, _error := model.PutBlogsByDeleted(Deleted_, &iBlogs)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlogsByLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LangId_ := self.Args("lang_id").MustInt()
	var iBlogs model.Blogs
	self.Bind(&iBlogs)
	_int64, _error := model.PutBlogsByLangId(LangId_, &iBlogs)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlogsByBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BlogId_ := self.Args("blog_id").MustInt64()
	var iBlogs model.Blogs
	self.Bind(&iBlogs)
	var iMap = helper.StructToMap(iBlogs)
	_error := model.UpdateBlogsByBlogId(BlogId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlogsBySiteIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	SiteId_ := self.Args("site_id").MustInt64()
	var iBlogs model.Blogs
	self.Bind(&iBlogs)
	var iMap = helper.StructToMap(iBlogs)
	_error := model.UpdateBlogsBySiteId(SiteId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlogsByDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Domain_ := self.Args("domain").String()
	var iBlogs model.Blogs
	self.Bind(&iBlogs)
	var iMap = helper.StructToMap(iBlogs)
	_error := model.UpdateBlogsByDomain(Domain_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlogsByPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Path_ := self.Args("path").String()
	var iBlogs model.Blogs
	self.Bind(&iBlogs)
	var iMap = helper.StructToMap(iBlogs)
	_error := model.UpdateBlogsByPath(Path_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlogsByRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Registered_ := self.Args("registered").Time()
	var iBlogs model.Blogs
	self.Bind(&iBlogs)
	var iMap = helper.StructToMap(iBlogs)
	_error := model.UpdateBlogsByRegistered(Registered_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlogsByLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LastUpdated_ := self.Args("last_updated").Time()
	var iBlogs model.Blogs
	self.Bind(&iBlogs)
	var iMap = helper.StructToMap(iBlogs)
	_error := model.UpdateBlogsByLastUpdated(LastUpdated_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlogsByPublicHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Public_ := self.Args("public").MustInt()
	var iBlogs model.Blogs
	self.Bind(&iBlogs)
	var iMap = helper.StructToMap(iBlogs)
	_error := model.UpdateBlogsByPublic(Public_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlogsByArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Archived_ := self.Args("archived").MustInt()
	var iBlogs model.Blogs
	self.Bind(&iBlogs)
	var iMap = helper.StructToMap(iBlogs)
	_error := model.UpdateBlogsByArchived(Archived_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlogsByMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Mature_ := self.Args("mature").MustInt()
	var iBlogs model.Blogs
	self.Bind(&iBlogs)
	var iMap = helper.StructToMap(iBlogs)
	_error := model.UpdateBlogsByMature(Mature_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlogsBySpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Spam_ := self.Args("spam").MustInt()
	var iBlogs model.Blogs
	self.Bind(&iBlogs)
	var iMap = helper.StructToMap(iBlogs)
	_error := model.UpdateBlogsBySpam(Spam_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlogsByDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	var iBlogs model.Blogs
	self.Bind(&iBlogs)
	var iMap = helper.StructToMap(iBlogs)
	_error := model.UpdateBlogsByDeleted(Deleted_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlogsByLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LangId_ := self.Args("lang_id").MustInt()
	var iBlogs model.Blogs
	self.Bind(&iBlogs)
	var iMap = helper.StructToMap(iBlogs)
	_error := model.UpdateBlogsByLangId(LangId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlogsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BlogId_ := self.Args("blog_id").MustInt64()
	_int64, _error := model.DeleteBlogs(BlogId_)
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

func DeleteBlogsByBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BlogId_ := self.Args("blog_id").MustInt64()
	_error := model.DeleteBlogsByBlogId(BlogId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlogsBySiteIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	SiteId_ := self.Args("site_id").MustInt64()
	_error := model.DeleteBlogsBySiteId(SiteId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlogsByDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Domain_ := self.Args("domain").String()
	_error := model.DeleteBlogsByDomain(Domain_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlogsByPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Path_ := self.Args("path").String()
	_error := model.DeleteBlogsByPath(Path_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlogsByRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Registered_ := self.Args("registered").Time()
	_error := model.DeleteBlogsByRegistered(Registered_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlogsByLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LastUpdated_ := self.Args("last_updated").Time()
	_error := model.DeleteBlogsByLastUpdated(LastUpdated_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlogsByPublicHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Public_ := self.Args("public").MustInt()
	_error := model.DeleteBlogsByPublic(Public_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlogsByArchivedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Archived_ := self.Args("archived").MustInt()
	_error := model.DeleteBlogsByArchived(Archived_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlogsByMatureHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Mature_ := self.Args("mature").MustInt()
	_error := model.DeleteBlogsByMature(Mature_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlogsBySpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Spam_ := self.Args("spam").MustInt()
	_error := model.DeleteBlogsBySpam(Spam_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlogsByDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	_error := model.DeleteBlogsByDeleted(Deleted_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlogsByLangIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LangId_ := self.Args("lang_id").MustInt()
	_error := model.DeleteBlogsByLangId(LangId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
