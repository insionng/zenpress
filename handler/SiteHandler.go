package handler

import (
	"github.com/insionng/zenpress/helper"
	"github.com/insionng/zenpress/model"
	"github.com/insionng/macross"
)

func GetSitesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetSitesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["sitesCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetSitesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSiteCountByIdHandler(self *macross.Context) error {
	Id_ := self.Args("id").MustInt64()
	_int64 := model.GetSiteCountById(Id_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["siteCount"] = 0
	}
	m["siteCount"] = _int64
	return self.JSON(m)
}

func GetSiteCountByDomainHandler(self *macross.Context) error {
	Domain_ := self.Args("domain").String()
	_int64 := model.GetSiteCountByDomain(Domain_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["siteCount"] = 0
	}
	m["siteCount"] = _int64
	return self.JSON(m)
}

func GetSiteCountByPathHandler(self *macross.Context) error {
	Path_ := self.Args("path").String()
	_int64 := model.GetSiteCountByPath(Path_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["siteCount"] = 0
	}
	m["siteCount"] = _int64
	return self.JSON(m)
}

func GetSitesByIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iId := self.Args("id").MustInt64()
	if (offset > 0) && helper.IsHas(iId) {
		_Site, _error := model.GetSitesById(offset, limit, iId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Site)
	}
	herr.Message = "Can't get to the GetSitesById's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitesByDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDomain := self.Args("domain").String()
	if (offset > 0) && helper.IsHas(iDomain) {
		_Site, _error := model.GetSitesByDomain(offset, limit, iDomain, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Site)
	}
	herr.Message = "Can't get to the GetSitesByDomain's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitesByPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iPath := self.Args("path").String()
	if (offset > 0) && helper.IsHas(iPath) {
		_Site, _error := model.GetSitesByPath(offset, limit, iPath, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Site)
	}
	herr.Message = "Can't get to the GetSitesByPath's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitesByIdAndDomainAndPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iDomain := self.Args("domain").String()
	iPath := self.Args("path").String()

	if helper.IsHas(iId) {
		_Site, _error := model.GetSitesByIdAndDomainAndPath(offset, limit, iId,iDomain,iPath)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Site)
	}
	herr.Message = "Can't get to the GetSitesByIdAndDomainAndPath's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitesByIdAndDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iDomain := self.Args("domain").String()

	if helper.IsHas(iId) {
		_Site, _error := model.GetSitesByIdAndDomain(offset, limit, iId,iDomain)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Site)
	}
	herr.Message = "Can't get to the GetSitesByIdAndDomain's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitesByIdAndPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iPath := self.Args("path").String()

	if helper.IsHas(iId) {
		_Site, _error := model.GetSitesByIdAndPath(offset, limit, iId,iPath)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Site)
	}
	herr.Message = "Can't get to the GetSitesByIdAndPath's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitesByDomainAndPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iPath := self.Args("path").String()

	if helper.IsHas(iDomain) {
		_Site, _error := model.GetSitesByDomainAndPath(offset, limit, iDomain,iPath)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Site)
	}
	herr.Message = "Can't get to the GetSitesByDomainAndPath's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_Site, _error := model.GetSites(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Site)
	}
	herr.Message = "Can't get to the GetSites' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSiteByIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iId := self.Args("id").MustInt64()
	if helper.IsHas(iId) {
		_Site := model.HasSiteById(iId)
		var m = map[string]interface{}{}
		m["site"] = _Site
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSiteById's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSiteByDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDomain := self.Args("domain").String()
	if helper.IsHas(iDomain) {
		_Site := model.HasSiteByDomain(iDomain)
		var m = map[string]interface{}{}
		m["site"] = _Site
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSiteByDomain's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSiteByPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iPath := self.Args("path").String()
	if helper.IsHas(iPath) {
		_Site := model.HasSiteByPath(iPath)
		var m = map[string]interface{}{}
		m["site"] = _Site
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSiteByPath's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSiteByIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iId := self.Args("id").MustInt64()
	if helper.IsHas(iId) {
		_Site, _error := model.GetSiteById(iId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Site)
	}
	herr.Message = "Can't get to the GetSiteById's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSiteByDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDomain := self.Args("domain").String()
	if helper.IsHas(iDomain) {
		_Site, _error := model.GetSiteByDomain(iDomain)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Site)
	}
	herr.Message = "Can't get to the GetSiteByDomain's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSiteByPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iPath := self.Args("path").String()
	if helper.IsHas(iPath) {
		_Site, _error := model.GetSiteByPath(iPath)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Site)
	}
	herr.Message = "Can't get to the GetSiteByPath's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSiteByIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	if helper.IsHas(Id_) {
		var iSite model.Site
		self.Bind(&iSite)
		_Site, _error := model.SetSiteById(Id_, &iSite)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Site)
	}
	herr.Message = "Can't get to the SetSiteById's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSiteByDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Domain_ := self.Args("domain").String()
	if helper.IsHas(Domain_) {
		var iSite model.Site
		self.Bind(&iSite)
		_Site, _error := model.SetSiteByDomain(Domain_, &iSite)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Site)
	}
	herr.Message = "Can't get to the SetSiteByDomain's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSiteByPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Path_ := self.Args("path").String()
	if helper.IsHas(Path_) {
		var iSite model.Site
		self.Bind(&iSite)
		_Site, _error := model.SetSiteByPath(Path_, &iSite)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Site)
	}
	herr.Message = "Can't get to the SetSiteByPath's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddSiteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	Domain_ := self.Args("domain").String()
	Path_ := self.Args("path").String()

	if helper.IsHas(Id_) {
		_error := model.AddSite(Id_,Domain_,Path_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddSite's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSiteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iSite model.Site
	self.Bind(&iSite)
	_int64, _error := model.PostSite(&iSite)
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

func PutSiteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iSite model.Site
	self.Bind(&iSite)
	_int64, _error := model.PutSite(&iSite)
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

func PutSiteByIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	var iSite model.Site
	self.Bind(&iSite)
	_int64, _error := model.PutSiteById(Id_, &iSite)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSiteByDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Domain_ := self.Args("domain").String()
	var iSite model.Site
	self.Bind(&iSite)
	_int64, _error := model.PutSiteByDomain(Domain_, &iSite)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSiteByPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Path_ := self.Args("path").String()
	var iSite model.Site
	self.Bind(&iSite)
	_int64, _error := model.PutSiteByPath(Path_, &iSite)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSiteByIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	var iSite model.Site
	self.Bind(&iSite)
	var iMap = helper.StructToMap(iSite)
	_error := model.UpdateSiteById(Id_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSiteByDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Domain_ := self.Args("domain").String()
	var iSite model.Site
	self.Bind(&iSite)
	var iMap = helper.StructToMap(iSite)
	_error := model.UpdateSiteByDomain(Domain_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSiteByPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Path_ := self.Args("path").String()
	var iSite model.Site
	self.Bind(&iSite)
	var iMap = helper.StructToMap(iSite)
	_error := model.UpdateSiteByPath(Path_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSiteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	_int64, _error := model.DeleteSite(Id_)
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

func DeleteSiteByIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	_error := model.DeleteSiteById(Id_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSiteByDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Domain_ := self.Args("domain").String()
	_error := model.DeleteSiteByDomain(Domain_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSiteByPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Path_ := self.Args("path").String()
	_error := model.DeleteSiteByPath(Path_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
