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

func GetSiteCountViaIdHandler(self *macross.Context) error {
	Id_ := self.Args("id").MustInt64()
	_int64 := model.GetSiteCountViaId(Id_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["siteCount"] = 0
	}
	m["siteCount"] = _int64
	return self.JSON(m)
}

func GetSiteCountViaDomainHandler(self *macross.Context) error {
	Domain_ := self.Args("domain").String()
	_int64 := model.GetSiteCountViaDomain(Domain_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["siteCount"] = 0
	}
	m["siteCount"] = _int64
	return self.JSON(m)
}

func GetSiteCountViaPathHandler(self *macross.Context) error {
	Path_ := self.Args("path").String()
	_int64 := model.GetSiteCountViaPath(Path_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["siteCount"] = 0
	}
	m["siteCount"] = _int64
	return self.JSON(m)
}

func GetSitesViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iId := self.Args("id").MustInt64()
	if (offset > 0) && helper.IsHas(iId) {
		_Site, _error := model.GetSitesViaId(offset, limit, iId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Site)
	}
	herr.Message = "Can't get to the GetSitesViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitesViaDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDomain := self.Args("domain").String()
	if (offset > 0) && helper.IsHas(iDomain) {
		_Site, _error := model.GetSitesViaDomain(offset, limit, iDomain, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Site)
	}
	herr.Message = "Can't get to the GetSitesViaDomain's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitesViaPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iPath := self.Args("path").String()
	if (offset > 0) && helper.IsHas(iPath) {
		_Site, _error := model.GetSitesViaPath(offset, limit, iPath, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Site)
	}
	herr.Message = "Can't get to the GetSitesViaPath's args."
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

func GetHasSiteViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iId := self.Args("id").MustInt64()
	if helper.IsHas(iId) {
		_Site := model.HasSiteViaId(iId)
		var m = map[string]interface{}{}
		m["site"] = _Site
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSiteViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSiteViaDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDomain := self.Args("domain").String()
	if helper.IsHas(iDomain) {
		_Site := model.HasSiteViaDomain(iDomain)
		var m = map[string]interface{}{}
		m["site"] = _Site
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSiteViaDomain's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSiteViaPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iPath := self.Args("path").String()
	if helper.IsHas(iPath) {
		_Site := model.HasSiteViaPath(iPath)
		var m = map[string]interface{}{}
		m["site"] = _Site
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSiteViaPath's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSiteViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iId := self.Args("id").MustInt64()
	if helper.IsHas(iId) {
		_Site, _error := model.GetSiteViaId(iId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Site)
	}
	herr.Message = "Can't get to the GetSiteViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSiteViaDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDomain := self.Args("domain").String()
	if helper.IsHas(iDomain) {
		_Site, _error := model.GetSiteViaDomain(iDomain)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Site)
	}
	herr.Message = "Can't get to the GetSiteViaDomain's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSiteViaPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iPath := self.Args("path").String()
	if helper.IsHas(iPath) {
		_Site, _error := model.GetSiteViaPath(iPath)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Site)
	}
	herr.Message = "Can't get to the GetSiteViaPath's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSiteViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	if helper.IsHas(Id_) {
		var iSite model.Site
		self.Bind(&iSite)
		_Site, _error := model.SetSiteViaId(Id_, &iSite)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Site)
	}
	herr.Message = "Can't get to the SetSiteViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSiteViaDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Domain_ := self.Args("domain").String()
	if helper.IsHas(Domain_) {
		var iSite model.Site
		self.Bind(&iSite)
		_Site, _error := model.SetSiteViaDomain(Domain_, &iSite)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Site)
	}
	herr.Message = "Can't get to the SetSiteViaDomain's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSiteViaPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Path_ := self.Args("path").String()
	if helper.IsHas(Path_) {
		var iSite model.Site
		self.Bind(&iSite)
		_Site, _error := model.SetSiteViaPath(Path_, &iSite)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Site)
	}
	herr.Message = "Can't get to the SetSiteViaPath's args."
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

func PutSiteHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iSite model.Site
	self.Bind(&iSite)
	_int64, _error := model.PutSite(&iSite)
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

func PutSiteViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	var iSite model.Site
	self.Bind(&iSite)
	_int64, _error := model.PutSiteViaId(Id_, &iSite)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSiteViaDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Domain_ := self.Args("domain").String()
	var iSite model.Site
	self.Bind(&iSite)
	_int64, _error := model.PutSiteViaDomain(Domain_, &iSite)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSiteViaPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Path_ := self.Args("path").String()
	var iSite model.Site
	self.Bind(&iSite)
	_int64, _error := model.PutSiteViaPath(Path_, &iSite)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSiteViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	var iSite model.Site
	self.Bind(&iSite)
	var iMap = helper.StructToMap(iSite)
	_error := model.UpdateSiteViaId(Id_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSiteViaDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Domain_ := self.Args("domain").String()
	var iSite model.Site
	self.Bind(&iSite)
	var iMap = helper.StructToMap(iSite)
	_error := model.UpdateSiteViaDomain(Domain_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSiteViaPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Path_ := self.Args("path").String()
	var iSite model.Site
	self.Bind(&iSite)
	var iMap = helper.StructToMap(iSite)
	_error := model.UpdateSiteViaPath(Path_, &iMap)
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

func DeleteSiteViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("id").MustInt64()
	_error := model.DeleteSiteViaId(Id_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSiteViaDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Domain_ := self.Args("domain").String()
	_error := model.DeleteSiteViaDomain(Domain_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSiteViaPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Path_ := self.Args("path").String()
	_error := model.DeleteSiteViaPath(Path_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
