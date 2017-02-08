package handler

import (
	"github.com/insionng/zenpress/helper"
	"github.com/insionng/zenpress/model"
	"github.com/insionng/macross"
)

func GetSitemetasCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetSitemetasCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["sitemetasCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetSitemetasCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitemetaCountViaMetaIdHandler(self *macross.Context) error {
	MetaId_ := self.Args("meta_id").MustInt64()
	_int64 := model.GetSitemetaCountViaMetaId(MetaId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["sitemetaCount"] = 0
	}
	m["sitemetaCount"] = _int64
	return self.JSON(m)
}

func GetSitemetaCountViaSiteIdHandler(self *macross.Context) error {
	SiteId_ := self.Args("site_id").MustInt64()
	_int64 := model.GetSitemetaCountViaSiteId(SiteId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["sitemetaCount"] = 0
	}
	m["sitemetaCount"] = _int64
	return self.JSON(m)
}

func GetSitemetaCountViaMetaKeyHandler(self *macross.Context) error {
	MetaKey_ := self.Args("meta_key").String()
	_int64 := model.GetSitemetaCountViaMetaKey(MetaKey_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["sitemetaCount"] = 0
	}
	m["sitemetaCount"] = _int64
	return self.JSON(m)
}

func GetSitemetaCountViaMetaValueHandler(self *macross.Context) error {
	MetaValue_ := self.Args("meta_value").String()
	_int64 := model.GetSitemetaCountViaMetaValue(MetaValue_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["sitemetaCount"] = 0
	}
	m["sitemetaCount"] = _int64
	return self.JSON(m)
}

func GetSitemetasViaMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMetaId := self.Args("meta_id").MustInt64()
	if (offset > 0) && helper.IsHas(iMetaId) {
		_Sitemeta, _error := model.GetSitemetasViaMetaId(offset, limit, iMetaId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the GetSitemetasViaMetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitemetasViaSiteIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iSiteId := self.Args("site_id").MustInt64()
	if (offset > 0) && helper.IsHas(iSiteId) {
		_Sitemeta, _error := model.GetSitemetasViaSiteId(offset, limit, iSiteId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the GetSitemetasViaSiteId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitemetasViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMetaKey := self.Args("meta_key").String()
	if (offset > 0) && helper.IsHas(iMetaKey) {
		_Sitemeta, _error := model.GetSitemetasViaMetaKey(offset, limit, iMetaKey, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the GetSitemetasViaMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitemetasViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMetaValue := self.Args("meta_value").String()
	if (offset > 0) && helper.IsHas(iMetaValue) {
		_Sitemeta, _error := model.GetSitemetasViaMetaValue(offset, limit, iMetaValue, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the GetSitemetasViaMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitemetasByMetaIdAndSiteIdAndMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMetaId := self.Args("meta_id").MustInt64()
	iSiteId := self.Args("site_id").MustInt64()
	iMetaKey := self.Args("meta_key").String()

	if helper.IsHas(iMetaId) {
		_Sitemeta, _error := model.GetSitemetasByMetaIdAndSiteIdAndMetaKey(offset, limit, iMetaId,iSiteId,iMetaKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the GetSitemetasByMetaIdAndSiteIdAndMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitemetasByMetaIdAndSiteIdAndMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMetaId := self.Args("meta_id").MustInt64()
	iSiteId := self.Args("site_id").MustInt64()
	iMetaValue := self.Args("meta_value").String()

	if helper.IsHas(iMetaId) {
		_Sitemeta, _error := model.GetSitemetasByMetaIdAndSiteIdAndMetaValue(offset, limit, iMetaId,iSiteId,iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the GetSitemetasByMetaIdAndSiteIdAndMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitemetasByMetaIdAndMetaKeyAndMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMetaId := self.Args("meta_id").MustInt64()
	iMetaKey := self.Args("meta_key").String()
	iMetaValue := self.Args("meta_value").String()

	if helper.IsHas(iMetaId) {
		_Sitemeta, _error := model.GetSitemetasByMetaIdAndMetaKeyAndMetaValue(offset, limit, iMetaId,iMetaKey,iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the GetSitemetasByMetaIdAndMetaKeyAndMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitemetasBySiteIdAndMetaKeyAndMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iMetaKey := self.Args("meta_key").String()
	iMetaValue := self.Args("meta_value").String()

	if helper.IsHas(iSiteId) {
		_Sitemeta, _error := model.GetSitemetasBySiteIdAndMetaKeyAndMetaValue(offset, limit, iSiteId,iMetaKey,iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the GetSitemetasBySiteIdAndMetaKeyAndMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitemetasByMetaIdAndSiteIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMetaId := self.Args("meta_id").MustInt64()
	iSiteId := self.Args("site_id").MustInt64()

	if helper.IsHas(iMetaId) {
		_Sitemeta, _error := model.GetSitemetasByMetaIdAndSiteId(offset, limit, iMetaId,iSiteId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the GetSitemetasByMetaIdAndSiteId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitemetasByMetaIdAndMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMetaId := self.Args("meta_id").MustInt64()
	iMetaKey := self.Args("meta_key").String()

	if helper.IsHas(iMetaId) {
		_Sitemeta, _error := model.GetSitemetasByMetaIdAndMetaKey(offset, limit, iMetaId,iMetaKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the GetSitemetasByMetaIdAndMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitemetasByMetaIdAndMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMetaId := self.Args("meta_id").MustInt64()
	iMetaValue := self.Args("meta_value").String()

	if helper.IsHas(iMetaId) {
		_Sitemeta, _error := model.GetSitemetasByMetaIdAndMetaValue(offset, limit, iMetaId,iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the GetSitemetasByMetaIdAndMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitemetasBySiteIdAndMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iMetaKey := self.Args("meta_key").String()

	if helper.IsHas(iSiteId) {
		_Sitemeta, _error := model.GetSitemetasBySiteIdAndMetaKey(offset, limit, iSiteId,iMetaKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the GetSitemetasBySiteIdAndMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitemetasBySiteIdAndMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSiteId := self.Args("site_id").MustInt64()
	iMetaValue := self.Args("meta_value").String()

	if helper.IsHas(iSiteId) {
		_Sitemeta, _error := model.GetSitemetasBySiteIdAndMetaValue(offset, limit, iSiteId,iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the GetSitemetasBySiteIdAndMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitemetasByMetaKeyAndMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMetaKey := self.Args("meta_key").String()
	iMetaValue := self.Args("meta_value").String()

	if helper.IsHas(iMetaKey) {
		_Sitemeta, _error := model.GetSitemetasByMetaKeyAndMetaValue(offset, limit, iMetaKey,iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the GetSitemetasByMetaKeyAndMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitemetasHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_Sitemeta, _error := model.GetSitemetas(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the GetSitemetas' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSitemetaViaMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaId := self.Args("meta_id").MustInt64()
	if helper.IsHas(iMetaId) {
		_Sitemeta := model.HasSitemetaViaMetaId(iMetaId)
		var m = map[string]interface{}{}
		m["sitemeta"] = _Sitemeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSitemetaViaMetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSitemetaViaSiteIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSiteId := self.Args("site_id").MustInt64()
	if helper.IsHas(iSiteId) {
		_Sitemeta := model.HasSitemetaViaSiteId(iSiteId)
		var m = map[string]interface{}{}
		m["sitemeta"] = _Sitemeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSitemetaViaSiteId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSitemetaViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaKey := self.Args("meta_key").String()
	if helper.IsHas(iMetaKey) {
		_Sitemeta := model.HasSitemetaViaMetaKey(iMetaKey)
		var m = map[string]interface{}{}
		m["sitemeta"] = _Sitemeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSitemetaViaMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSitemetaViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaValue := self.Args("meta_value").String()
	if helper.IsHas(iMetaValue) {
		_Sitemeta := model.HasSitemetaViaMetaValue(iMetaValue)
		var m = map[string]interface{}{}
		m["sitemeta"] = _Sitemeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSitemetaViaMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitemetaViaMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaId := self.Args("meta_id").MustInt64()
	if helper.IsHas(iMetaId) {
		_Sitemeta, _error := model.GetSitemetaViaMetaId(iMetaId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the GetSitemetaViaMetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitemetaViaSiteIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSiteId := self.Args("site_id").MustInt64()
	if helper.IsHas(iSiteId) {
		_Sitemeta, _error := model.GetSitemetaViaSiteId(iSiteId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the GetSitemetaViaSiteId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitemetaViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaKey := self.Args("meta_key").String()
	if helper.IsHas(iMetaKey) {
		_Sitemeta, _error := model.GetSitemetaViaMetaKey(iMetaKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the GetSitemetaViaMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitemetaViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaValue := self.Args("meta_value").String()
	if helper.IsHas(iMetaValue) {
		_Sitemeta, _error := model.GetSitemetaViaMetaValue(iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the GetSitemetaViaMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSitemetaViaMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	if helper.IsHas(MetaId_) {
		var iSitemeta model.Sitemeta
		self.Bind(&iSitemeta)
		_Sitemeta, _error := model.SetSitemetaViaMetaId(MetaId_, &iSitemeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the SetSitemetaViaMetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSitemetaViaSiteIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	SiteId_ := self.Args("site_id").MustInt64()
	if helper.IsHas(SiteId_) {
		var iSitemeta model.Sitemeta
		self.Bind(&iSitemeta)
		_Sitemeta, _error := model.SetSitemetaViaSiteId(SiteId_, &iSitemeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the SetSitemetaViaSiteId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSitemetaViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	if helper.IsHas(MetaKey_) {
		var iSitemeta model.Sitemeta
		self.Bind(&iSitemeta)
		_Sitemeta, _error := model.SetSitemetaViaMetaKey(MetaKey_, &iSitemeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the SetSitemetaViaMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSitemetaViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	if helper.IsHas(MetaValue_) {
		var iSitemeta model.Sitemeta
		self.Bind(&iSitemeta)
		_Sitemeta, _error := model.SetSitemetaViaMetaValue(MetaValue_, &iSitemeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the SetSitemetaViaMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddSitemetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	SiteId_ := self.Args("site_id").MustInt64()
	MetaKey_ := self.Args("meta_key").String()
	MetaValue_ := self.Args("meta_value").String()

	if helper.IsHas(MetaId_) {
		_error := model.AddSitemeta(MetaId_,SiteId_,MetaKey_,MetaValue_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddSitemeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSitemetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iSitemeta model.Sitemeta
	self.Bind(&iSitemeta)
	_int64, _error := model.PostSitemeta(&iSitemeta)
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

func PutSitemetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iSitemeta model.Sitemeta
	self.Bind(&iSitemeta)
	_int64, _error := model.PutSitemeta(&iSitemeta)
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

func PutSitemetaViaMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	var iSitemeta model.Sitemeta
	self.Bind(&iSitemeta)
	_int64, _error := model.PutSitemetaViaMetaId(MetaId_, &iSitemeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSitemetaViaSiteIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	SiteId_ := self.Args("site_id").MustInt64()
	var iSitemeta model.Sitemeta
	self.Bind(&iSitemeta)
	_int64, _error := model.PutSitemetaViaSiteId(SiteId_, &iSitemeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSitemetaViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	var iSitemeta model.Sitemeta
	self.Bind(&iSitemeta)
	_int64, _error := model.PutSitemetaViaMetaKey(MetaKey_, &iSitemeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSitemetaViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	var iSitemeta model.Sitemeta
	self.Bind(&iSitemeta)
	_int64, _error := model.PutSitemetaViaMetaValue(MetaValue_, &iSitemeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSitemetaViaMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	var iSitemeta model.Sitemeta
	self.Bind(&iSitemeta)
	var iMap = helper.StructToMap(iSitemeta)
	_error := model.UpdateSitemetaViaMetaId(MetaId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSitemetaViaSiteIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	SiteId_ := self.Args("site_id").MustInt64()
	var iSitemeta model.Sitemeta
	self.Bind(&iSitemeta)
	var iMap = helper.StructToMap(iSitemeta)
	_error := model.UpdateSitemetaViaSiteId(SiteId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSitemetaViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	var iSitemeta model.Sitemeta
	self.Bind(&iSitemeta)
	var iMap = helper.StructToMap(iSitemeta)
	_error := model.UpdateSitemetaViaMetaKey(MetaKey_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSitemetaViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	var iSitemeta model.Sitemeta
	self.Bind(&iSitemeta)
	var iMap = helper.StructToMap(iSitemeta)
	_error := model.UpdateSitemetaViaMetaValue(MetaValue_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSitemetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	_int64, _error := model.DeleteSitemeta(MetaId_)
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

func DeleteSitemetaViaMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	_error := model.DeleteSitemetaViaMetaId(MetaId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSitemetaViaSiteIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	SiteId_ := self.Args("site_id").MustInt64()
	_error := model.DeleteSitemetaViaSiteId(SiteId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSitemetaViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	_error := model.DeleteSitemetaViaMetaKey(MetaKey_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSitemetaViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	_error := model.DeleteSitemetaViaMetaValue(MetaValue_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
