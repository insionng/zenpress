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

func GetSitemetaCountByMetaIdHandler(self *macross.Context) error {
	MetaId_ := self.Args("meta_id").MustInt64()
	_int64 := model.GetSitemetaCountByMetaId(MetaId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["sitemetaCount"] = 0
	}
	m["sitemetaCount"] = _int64
	return self.JSON(m)
}

func GetSitemetaCountBySiteIdHandler(self *macross.Context) error {
	SiteId_ := self.Args("site_id").MustInt64()
	_int64 := model.GetSitemetaCountBySiteId(SiteId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["sitemetaCount"] = 0
	}
	m["sitemetaCount"] = _int64
	return self.JSON(m)
}

func GetSitemetaCountByMetaKeyHandler(self *macross.Context) error {
	MetaKey_ := self.Args("meta_key").String()
	_int64 := model.GetSitemetaCountByMetaKey(MetaKey_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["sitemetaCount"] = 0
	}
	m["sitemetaCount"] = _int64
	return self.JSON(m)
}

func GetSitemetaCountByMetaValueHandler(self *macross.Context) error {
	MetaValue_ := self.Args("meta_value").String()
	_int64 := model.GetSitemetaCountByMetaValue(MetaValue_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["sitemetaCount"] = 0
	}
	m["sitemetaCount"] = _int64
	return self.JSON(m)
}

func GetSitemetasByMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMetaId := self.Args("meta_id").MustInt64()
	if (offset > 0) && helper.IsHas(iMetaId) {
		_Sitemeta, _error := model.GetSitemetasByMetaId(offset, limit, iMetaId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the GetSitemetasByMetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitemetasBySiteIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iSiteId := self.Args("site_id").MustInt64()
	if (offset > 0) && helper.IsHas(iSiteId) {
		_Sitemeta, _error := model.GetSitemetasBySiteId(offset, limit, iSiteId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the GetSitemetasBySiteId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitemetasByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMetaKey := self.Args("meta_key").String()
	if (offset > 0) && helper.IsHas(iMetaKey) {
		_Sitemeta, _error := model.GetSitemetasByMetaKey(offset, limit, iMetaKey, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the GetSitemetasByMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitemetasByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMetaValue := self.Args("meta_value").String()
	if (offset > 0) && helper.IsHas(iMetaValue) {
		_Sitemeta, _error := model.GetSitemetasByMetaValue(offset, limit, iMetaValue, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the GetSitemetasByMetaValue's args."
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

func GetHasSitemetaByMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaId := self.Args("meta_id").MustInt64()
	if helper.IsHas(iMetaId) {
		_Sitemeta := model.HasSitemetaByMetaId(iMetaId)
		var m = map[string]interface{}{}
		m["sitemeta"] = _Sitemeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSitemetaByMetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSitemetaBySiteIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSiteId := self.Args("site_id").MustInt64()
	if helper.IsHas(iSiteId) {
		_Sitemeta := model.HasSitemetaBySiteId(iSiteId)
		var m = map[string]interface{}{}
		m["sitemeta"] = _Sitemeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSitemetaBySiteId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSitemetaByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaKey := self.Args("meta_key").String()
	if helper.IsHas(iMetaKey) {
		_Sitemeta := model.HasSitemetaByMetaKey(iMetaKey)
		var m = map[string]interface{}{}
		m["sitemeta"] = _Sitemeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSitemetaByMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSitemetaByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaValue := self.Args("meta_value").String()
	if helper.IsHas(iMetaValue) {
		_Sitemeta := model.HasSitemetaByMetaValue(iMetaValue)
		var m = map[string]interface{}{}
		m["sitemeta"] = _Sitemeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSitemetaByMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitemetaByMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaId := self.Args("meta_id").MustInt64()
	if helper.IsHas(iMetaId) {
		_Sitemeta, _error := model.GetSitemetaByMetaId(iMetaId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the GetSitemetaByMetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitemetaBySiteIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSiteId := self.Args("site_id").MustInt64()
	if helper.IsHas(iSiteId) {
		_Sitemeta, _error := model.GetSitemetaBySiteId(iSiteId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the GetSitemetaBySiteId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitemetaByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaKey := self.Args("meta_key").String()
	if helper.IsHas(iMetaKey) {
		_Sitemeta, _error := model.GetSitemetaByMetaKey(iMetaKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the GetSitemetaByMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSitemetaByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaValue := self.Args("meta_value").String()
	if helper.IsHas(iMetaValue) {
		_Sitemeta, _error := model.GetSitemetaByMetaValue(iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the GetSitemetaByMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSitemetaByMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	if helper.IsHas(MetaId_) {
		var iSitemeta model.Sitemeta
		self.Bind(&iSitemeta)
		_Sitemeta, _error := model.SetSitemetaByMetaId(MetaId_, &iSitemeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the SetSitemetaByMetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSitemetaBySiteIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	SiteId_ := self.Args("site_id").MustInt64()
	if helper.IsHas(SiteId_) {
		var iSitemeta model.Sitemeta
		self.Bind(&iSitemeta)
		_Sitemeta, _error := model.SetSitemetaBySiteId(SiteId_, &iSitemeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the SetSitemetaBySiteId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSitemetaByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	if helper.IsHas(MetaKey_) {
		var iSitemeta model.Sitemeta
		self.Bind(&iSitemeta)
		_Sitemeta, _error := model.SetSitemetaByMetaKey(MetaKey_, &iSitemeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the SetSitemetaByMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSitemetaByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	if helper.IsHas(MetaValue_) {
		var iSitemeta model.Sitemeta
		self.Bind(&iSitemeta)
		_Sitemeta, _error := model.SetSitemetaByMetaValue(MetaValue_, &iSitemeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Sitemeta)
	}
	herr.Message = "Can't get to the SetSitemetaByMetaValue's args."
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

func PutSitemetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iSitemeta model.Sitemeta
	self.Bind(&iSitemeta)
	_int64, _error := model.PutSitemeta(&iSitemeta)
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

func PutSitemetaByMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	var iSitemeta model.Sitemeta
	self.Bind(&iSitemeta)
	_int64, _error := model.PutSitemetaByMetaId(MetaId_, &iSitemeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSitemetaBySiteIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	SiteId_ := self.Args("site_id").MustInt64()
	var iSitemeta model.Sitemeta
	self.Bind(&iSitemeta)
	_int64, _error := model.PutSitemetaBySiteId(SiteId_, &iSitemeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSitemetaByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	var iSitemeta model.Sitemeta
	self.Bind(&iSitemeta)
	_int64, _error := model.PutSitemetaByMetaKey(MetaKey_, &iSitemeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSitemetaByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	var iSitemeta model.Sitemeta
	self.Bind(&iSitemeta)
	_int64, _error := model.PutSitemetaByMetaValue(MetaValue_, &iSitemeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSitemetaByMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	var iSitemeta model.Sitemeta
	self.Bind(&iSitemeta)
	var iMap = helper.StructToMap(iSitemeta)
	_error := model.UpdateSitemetaByMetaId(MetaId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSitemetaBySiteIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	SiteId_ := self.Args("site_id").MustInt64()
	var iSitemeta model.Sitemeta
	self.Bind(&iSitemeta)
	var iMap = helper.StructToMap(iSitemeta)
	_error := model.UpdateSitemetaBySiteId(SiteId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSitemetaByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	var iSitemeta model.Sitemeta
	self.Bind(&iSitemeta)
	var iMap = helper.StructToMap(iSitemeta)
	_error := model.UpdateSitemetaByMetaKey(MetaKey_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSitemetaByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	var iSitemeta model.Sitemeta
	self.Bind(&iSitemeta)
	var iMap = helper.StructToMap(iSitemeta)
	_error := model.UpdateSitemetaByMetaValue(MetaValue_, &iMap)
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

func DeleteSitemetaByMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	_error := model.DeleteSitemetaByMetaId(MetaId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSitemetaBySiteIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	SiteId_ := self.Args("site_id").MustInt64()
	_error := model.DeleteSitemetaBySiteId(SiteId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSitemetaByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	_error := model.DeleteSitemetaByMetaKey(MetaKey_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSitemetaByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	_error := model.DeleteSitemetaByMetaValue(MetaValue_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
