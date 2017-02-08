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

func GetBlogVersionsCountViaBlogIdHandler(self *macross.Context) error {
	BlogId_ := self.Args("blog_id").MustInt64()
	_int64 := model.GetBlogVersionsCountViaBlogId(BlogId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["blog_versionsCount"] = 0
	}
	m["blog_versionsCount"] = _int64
	return self.JSON(m)
}

func GetBlogVersionsCountViaDbVersionHandler(self *macross.Context) error {
	DbVersion_ := self.Args("db_version").String()
	_int64 := model.GetBlogVersionsCountViaDbVersion(DbVersion_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["blog_versionsCount"] = 0
	}
	m["blog_versionsCount"] = _int64
	return self.JSON(m)
}

func GetBlogVersionsCountViaLastUpdatedHandler(self *macross.Context) error {
	LastUpdated_ := self.Args("last_updated").Time()
	_int64 := model.GetBlogVersionsCountViaLastUpdated(LastUpdated_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["blog_versionsCount"] = 0
	}
	m["blog_versionsCount"] = _int64
	return self.JSON(m)
}

func GetBlogVersionsesViaBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iBlogId := self.Args("blog_id").MustInt64()
	if (offset > 0) && helper.IsHas(iBlogId) {
		_BlogVersions, _error := model.GetBlogVersionsesViaBlogId(offset, limit, iBlogId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlogVersions)
	}
	herr.Message = "Can't get to the GetBlogVersionsesViaBlogId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogVersionsesViaDbVersionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDbVersion := self.Args("db_version").String()
	if (offset > 0) && helper.IsHas(iDbVersion) {
		_BlogVersions, _error := model.GetBlogVersionsesViaDbVersion(offset, limit, iDbVersion, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlogVersions)
	}
	herr.Message = "Can't get to the GetBlogVersionsesViaDbVersion's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogVersionsesViaLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iLastUpdated := self.Args("last_updated").Time()
	if (offset > 0) && helper.IsHas(iLastUpdated) {
		_BlogVersions, _error := model.GetBlogVersionsesViaLastUpdated(offset, limit, iLastUpdated, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlogVersions)
	}
	herr.Message = "Can't get to the GetBlogVersionsesViaLastUpdated's args."
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

func GetHasBlogVersionsViaBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBlogId := self.Args("blog_id").MustInt64()
	if helper.IsHas(iBlogId) {
		_BlogVersions := model.HasBlogVersionsViaBlogId(iBlogId)
		var m = map[string]interface{}{}
		m["blog_versions"] = _BlogVersions
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlogVersionsViaBlogId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlogVersionsViaDbVersionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDbVersion := self.Args("db_version").String()
	if helper.IsHas(iDbVersion) {
		_BlogVersions := model.HasBlogVersionsViaDbVersion(iDbVersion)
		var m = map[string]interface{}{}
		m["blog_versions"] = _BlogVersions
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlogVersionsViaDbVersion's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasBlogVersionsViaLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLastUpdated := self.Args("last_updated").Time()
	if helper.IsHas(iLastUpdated) {
		_BlogVersions := model.HasBlogVersionsViaLastUpdated(iLastUpdated)
		var m = map[string]interface{}{}
		m["blog_versions"] = _BlogVersions
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasBlogVersionsViaLastUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogVersionsViaBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBlogId := self.Args("blog_id").MustInt64()
	if helper.IsHas(iBlogId) {
		_BlogVersions, _error := model.GetBlogVersionsViaBlogId(iBlogId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlogVersions)
	}
	herr.Message = "Can't get to the GetBlogVersionsViaBlogId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogVersionsViaDbVersionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDbVersion := self.Args("db_version").String()
	if helper.IsHas(iDbVersion) {
		_BlogVersions, _error := model.GetBlogVersionsViaDbVersion(iDbVersion)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlogVersions)
	}
	herr.Message = "Can't get to the GetBlogVersionsViaDbVersion's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetBlogVersionsViaLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iLastUpdated := self.Args("last_updated").Time()
	if helper.IsHas(iLastUpdated) {
		_BlogVersions, _error := model.GetBlogVersionsViaLastUpdated(iLastUpdated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlogVersions)
	}
	herr.Message = "Can't get to the GetBlogVersionsViaLastUpdated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlogVersionsViaBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BlogId_ := self.Args("blog_id").MustInt64()
	if helper.IsHas(BlogId_) {
		var iBlogVersions model.BlogVersions
		self.Bind(&iBlogVersions)
		_BlogVersions, _error := model.SetBlogVersionsViaBlogId(BlogId_, &iBlogVersions)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlogVersions)
	}
	herr.Message = "Can't get to the SetBlogVersionsViaBlogId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlogVersionsViaDbVersionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DbVersion_ := self.Args("db_version").String()
	if helper.IsHas(DbVersion_) {
		var iBlogVersions model.BlogVersions
		self.Bind(&iBlogVersions)
		_BlogVersions, _error := model.SetBlogVersionsViaDbVersion(DbVersion_, &iBlogVersions)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlogVersions)
	}
	herr.Message = "Can't get to the SetBlogVersionsViaDbVersion's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetBlogVersionsViaLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LastUpdated_ := self.Args("last_updated").Time()
	if helper.IsHas(LastUpdated_) {
		var iBlogVersions model.BlogVersions
		self.Bind(&iBlogVersions)
		_BlogVersions, _error := model.SetBlogVersionsViaLastUpdated(LastUpdated_, &iBlogVersions)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_BlogVersions)
	}
	herr.Message = "Can't get to the SetBlogVersionsViaLastUpdated's args."
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

func PutBlogVersionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iBlogVersions model.BlogVersions
	self.Bind(&iBlogVersions)
	_int64, _error := model.PutBlogVersions(&iBlogVersions)
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

func PutBlogVersionsViaBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BlogId_ := self.Args("blog_id").MustInt64()
	var iBlogVersions model.BlogVersions
	self.Bind(&iBlogVersions)
	_int64, _error := model.PutBlogVersionsViaBlogId(BlogId_, &iBlogVersions)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlogVersionsViaDbVersionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DbVersion_ := self.Args("db_version").String()
	var iBlogVersions model.BlogVersions
	self.Bind(&iBlogVersions)
	_int64, _error := model.PutBlogVersionsViaDbVersion(DbVersion_, &iBlogVersions)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutBlogVersionsViaLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LastUpdated_ := self.Args("last_updated").Time()
	var iBlogVersions model.BlogVersions
	self.Bind(&iBlogVersions)
	_int64, _error := model.PutBlogVersionsViaLastUpdated(LastUpdated_, &iBlogVersions)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlogVersionsViaBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BlogId_ := self.Args("blog_id").MustInt64()
	var iBlogVersions model.BlogVersions
	self.Bind(&iBlogVersions)
	var iMap = helper.StructToMap(iBlogVersions)
	_error := model.UpdateBlogVersionsViaBlogId(BlogId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlogVersionsViaDbVersionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DbVersion_ := self.Args("db_version").String()
	var iBlogVersions model.BlogVersions
	self.Bind(&iBlogVersions)
	var iMap = helper.StructToMap(iBlogVersions)
	_error := model.UpdateBlogVersionsViaDbVersion(DbVersion_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateBlogVersionsViaLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LastUpdated_ := self.Args("last_updated").Time()
	var iBlogVersions model.BlogVersions
	self.Bind(&iBlogVersions)
	var iMap = helper.StructToMap(iBlogVersions)
	_error := model.UpdateBlogVersionsViaLastUpdated(LastUpdated_, &iMap)
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

func DeleteBlogVersionsViaBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BlogId_ := self.Args("blog_id").MustInt64()
	_error := model.DeleteBlogVersionsViaBlogId(BlogId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlogVersionsViaDbVersionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DbVersion_ := self.Args("db_version").String()
	_error := model.DeleteBlogVersionsViaDbVersion(DbVersion_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteBlogVersionsViaLastUpdatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	LastUpdated_ := self.Args("last_updated").Time()
	_error := model.DeleteBlogVersionsViaLastUpdated(LastUpdated_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
