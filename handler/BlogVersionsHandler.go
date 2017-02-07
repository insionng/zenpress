package handler

import (
	"github.com/insionng/zenpress/helper"
	"github.com/insionng/zenpress/model"
	"github.com/insionng/macross"
)

func GetBlogVersionsesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetBlogVersionsesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["blog_versionssCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetBlogVersionsesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogVersionsCountByBlogIdHandler(self *macross.Context) error {
	BlogId_ := self.Args("blog_id").MustInt64()
	_int64 := model.GetBlogVersionsCountByBlogId(BlogId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["blog_versionsCount"] = 0
	}
	m["blog_versionsCount"] = _int64
	return self.JSON(m)
}

func GetBlogVersionsCountByDbVersionHandler(self *macross.Context) error {
	DbVersion_ := self.Args("db_version").String()
	_int64 := model.GetBlogVersionsCountByDbVersion(DbVersion_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["blog_versionsCount"] = 0
	}
	m["blog_versionsCount"] = _int64
	return self.JSON(m)
}

func GetBlogVersionsCountByLastUpdatedHandler(self *macross.Context) error {
	LastUpdated_ := self.Args("last_updated").Time()
	_int64 := model.GetBlogVersionsCountByLastUpdated(LastUpdated_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["blog_versionsCount"] = 0
	}
	m["blog_versionsCount"] = _int64
	return self.JSON(m)
}

func GetBlogVersionsesByBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iBlogId := self.Args("blog_id").MustInt64()
	if (offset > 0) && helper.IsHas(iBlogId) {
		_BlogVersions, _error := model.GetBlogVersionsesByBlogId(offset, limit, iBlogId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlogVersions)
	}
	herr.Message = "Can't get to the GetBlogVersionsesByBlogId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogVersionsesByDbVersionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDbVersion := self.Args("db_version").String()
	if (offset > 0) && helper.IsHas(iDbVersion) {
		_BlogVersions, _error := model.GetBlogVersionsesByDbVersion(offset, limit, iDbVersion, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlogVersions)
	}
	herr.Message = "Can't get to the GetBlogVersionsesByDbVersion's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogVersionsesByLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLastUpdated := self.Args("last_updated").Time()
	if (offset > 0) && helper.IsHas(iLastUpdated) {
		_BlogVersions, _error := model.GetBlogVersionsesByLastUpdated(offset, limit, iLastUpdated, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlogVersions)
	}
	herr.Message = "Can't get to the GetBlogVersionsesByLastUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogVersionsesByBlogIdAndDbVersionAndLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iDbVersion := self.Args("db_version").String()
	iLastUpdated := self.Args("last_updated").Time()

	if helper.IsHas(iBlogId) {
		_BlogVersions, _error := model.GetBlogVersionsesByBlogIdAndDbVersionAndLastUpdated(offset, limit, iBlogId,iDbVersion,iLastUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlogVersions)
	}
	herr.Message = "Can't get to the GetBlogVersionsesByBlogIdAndDbVersionAndLastUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogVersionsesByBlogIdAndDbVersionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iDbVersion := self.Args("db_version").String()

	if helper.IsHas(iBlogId) {
		_BlogVersions, _error := model.GetBlogVersionsesByBlogIdAndDbVersion(offset, limit, iBlogId,iDbVersion)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlogVersions)
	}
	herr.Message = "Can't get to the GetBlogVersionsesByBlogIdAndDbVersion's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogVersionsesByBlogIdAndLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iLastUpdated := self.Args("last_updated").Time()

	if helper.IsHas(iBlogId) {
		_BlogVersions, _error := model.GetBlogVersionsesByBlogIdAndLastUpdated(offset, limit, iBlogId,iLastUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlogVersions)
	}
	herr.Message = "Can't get to the GetBlogVersionsesByBlogIdAndLastUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogVersionsesByDbVersionAndLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDbVersion := self.Args("db_version").String()
	iLastUpdated := self.Args("last_updated").Time()

	if helper.IsHas(iDbVersion) {
		_BlogVersions, _error := model.GetBlogVersionsesByDbVersionAndLastUpdated(offset, limit, iDbVersion,iLastUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlogVersions)
	}
	herr.Message = "Can't get to the GetBlogVersionsesByDbVersionAndLastUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogVersionsesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_BlogVersions, _error := model.GetBlogVersionses(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlogVersions)
	}
	herr.Message = "Can't get to the GetBlogVersionses' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlogVersionsByBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBlogId := self.Args("blog_id").MustInt64()
	if helper.IsHas(iBlogId) {
		_BlogVersions := model.HasBlogVersionsByBlogId(iBlogId)
		var m = map[string]interface{}{}
		m["blog_versions"] = _BlogVersions
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlogVersionsByBlogId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlogVersionsByDbVersionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDbVersion := self.Args("db_version").String()
	if helper.IsHas(iDbVersion) {
		_BlogVersions := model.HasBlogVersionsByDbVersion(iDbVersion)
		var m = map[string]interface{}{}
		m["blog_versions"] = _BlogVersions
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlogVersionsByDbVersion's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlogVersionsByLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLastUpdated := self.Args("last_updated").Time()
	if helper.IsHas(iLastUpdated) {
		_BlogVersions := model.HasBlogVersionsByLastUpdated(iLastUpdated)
		var m = map[string]interface{}{}
		m["blog_versions"] = _BlogVersions
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlogVersionsByLastUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogVersionsByBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBlogId := self.Args("blog_id").MustInt64()
	if helper.IsHas(iBlogId) {
		_BlogVersions, _error := model.GetBlogVersionsByBlogId(iBlogId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlogVersions)
	}
	herr.Message = "Can't get to the GetBlogVersionsByBlogId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogVersionsByDbVersionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDbVersion := self.Args("db_version").String()
	if helper.IsHas(iDbVersion) {
		_BlogVersions, _error := model.GetBlogVersionsByDbVersion(iDbVersion)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlogVersions)
	}
	herr.Message = "Can't get to the GetBlogVersionsByDbVersion's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogVersionsByLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLastUpdated := self.Args("last_updated").Time()
	if helper.IsHas(iLastUpdated) {
		_BlogVersions, _error := model.GetBlogVersionsByLastUpdated(iLastUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlogVersions)
	}
	herr.Message = "Can't get to the GetBlogVersionsByLastUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlogVersionsByBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BlogId_ := self.Args("blog_id").MustInt64()
	if helper.IsHas(BlogId_) {
		var iBlogVersions model.BlogVersions
		self.Bind(&iBlogVersions)
		_BlogVersions, _error := model.SetBlogVersionsByBlogId(BlogId_, &iBlogVersions)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlogVersions)
	}
	herr.Message = "Can't get to the SetBlogVersionsByBlogId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlogVersionsByDbVersionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DbVersion_ := self.Args("db_version").String()
	if helper.IsHas(DbVersion_) {
		var iBlogVersions model.BlogVersions
		self.Bind(&iBlogVersions)
		_BlogVersions, _error := model.SetBlogVersionsByDbVersion(DbVersion_, &iBlogVersions)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlogVersions)
	}
	herr.Message = "Can't get to the SetBlogVersionsByDbVersion's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlogVersionsByLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LastUpdated_ := self.Args("last_updated").Time()
	if helper.IsHas(LastUpdated_) {
		var iBlogVersions model.BlogVersions
		self.Bind(&iBlogVersions)
		_BlogVersions, _error := model.SetBlogVersionsByLastUpdated(LastUpdated_, &iBlogVersions)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlogVersions)
	}
	herr.Message = "Can't get to the SetBlogVersionsByLastUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddBlogVersionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BlogId_ := self.Args("blog_id").MustInt64()
	DbVersion_ := self.Args("db_version").String()
	LastUpdated_ := self.Args("last_updated").Time()

	if helper.IsHas(BlogId_) {
		_error := model.AddBlogVersions(BlogId_,DbVersion_,LastUpdated_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddBlogVersions's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostBlogVersionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iBlogVersions model.BlogVersions
	self.Bind(&iBlogVersions)
	_int64, _error := model.PostBlogVersions(&iBlogVersions)
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

func PutBlogVersionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iBlogVersions model.BlogVersions
	self.Bind(&iBlogVersions)
	_int64, _error := model.PutBlogVersions(&iBlogVersions)
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

func PutBlogVersionsByBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BlogId_ := self.Args("blog_id").MustInt64()
	var iBlogVersions model.BlogVersions
	self.Bind(&iBlogVersions)
	_int64, _error := model.PutBlogVersionsByBlogId(BlogId_, &iBlogVersions)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlogVersionsByDbVersionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DbVersion_ := self.Args("db_version").String()
	var iBlogVersions model.BlogVersions
	self.Bind(&iBlogVersions)
	_int64, _error := model.PutBlogVersionsByDbVersion(DbVersion_, &iBlogVersions)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlogVersionsByLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LastUpdated_ := self.Args("last_updated").Time()
	var iBlogVersions model.BlogVersions
	self.Bind(&iBlogVersions)
	_int64, _error := model.PutBlogVersionsByLastUpdated(LastUpdated_, &iBlogVersions)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlogVersionsByBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BlogId_ := self.Args("blog_id").MustInt64()
	var iBlogVersions model.BlogVersions
	self.Bind(&iBlogVersions)
	var iMap = helper.StructToMap(iBlogVersions)
	_error := model.UpdateBlogVersionsByBlogId(BlogId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlogVersionsByDbVersionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DbVersion_ := self.Args("db_version").String()
	var iBlogVersions model.BlogVersions
	self.Bind(&iBlogVersions)
	var iMap = helper.StructToMap(iBlogVersions)
	_error := model.UpdateBlogVersionsByDbVersion(DbVersion_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlogVersionsByLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LastUpdated_ := self.Args("last_updated").Time()
	var iBlogVersions model.BlogVersions
	self.Bind(&iBlogVersions)
	var iMap = helper.StructToMap(iBlogVersions)
	_error := model.UpdateBlogVersionsByLastUpdated(LastUpdated_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlogVersionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BlogId_ := self.Args("blog_id").MustInt64()
	_int64, _error := model.DeleteBlogVersions(BlogId_)
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

func DeleteBlogVersionsByBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BlogId_ := self.Args("blog_id").MustInt64()
	_error := model.DeleteBlogVersionsByBlogId(BlogId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlogVersionsByDbVersionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DbVersion_ := self.Args("db_version").String()
	_error := model.DeleteBlogVersionsByDbVersion(DbVersion_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlogVersionsByLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LastUpdated_ := self.Args("last_updated").Time()
	_error := model.DeleteBlogVersionsByLastUpdated(LastUpdated_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
