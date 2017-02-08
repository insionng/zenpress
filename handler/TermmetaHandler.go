package handler

import (
	"github.com/insionng/zenpress/helper"
	"github.com/insionng/zenpress/model"
	"github.com/insionng/macross"
)

func GetTermmetasCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetTermmetasCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["termmetasCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetTermmetasCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermmetaCountViaMetaIdHandler(self *macross.Context) error {
	MetaId_ := self.Args("meta_id").MustInt64()
	_int64 := model.GetTermmetaCountViaMetaId(MetaId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["termmetaCount"] = 0
	}
	m["termmetaCount"] = _int64
	return self.JSON(m)
}

func GetTermmetaCountViaTermIdHandler(self *macross.Context) error {
	TermId_ := self.Args("term_id").MustInt64()
	_int64 := model.GetTermmetaCountViaTermId(TermId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["termmetaCount"] = 0
	}
	m["termmetaCount"] = _int64
	return self.JSON(m)
}

func GetTermmetaCountViaMetaKeyHandler(self *macross.Context) error {
	MetaKey_ := self.Args("meta_key").String()
	_int64 := model.GetTermmetaCountViaMetaKey(MetaKey_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["termmetaCount"] = 0
	}
	m["termmetaCount"] = _int64
	return self.JSON(m)
}

func GetTermmetaCountViaMetaValueHandler(self *macross.Context) error {
	MetaValue_ := self.Args("meta_value").String()
	_int64 := model.GetTermmetaCountViaMetaValue(MetaValue_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["termmetaCount"] = 0
	}
	m["termmetaCount"] = _int64
	return self.JSON(m)
}

func GetTermmetasViaMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMetaId := self.Args("meta_id").MustInt64()
	if (offset > 0) && helper.IsHas(iMetaId) {
		_Termmeta, _error := model.GetTermmetasViaMetaId(offset, limit, iMetaId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the GetTermmetasViaMetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermmetasViaTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTermId := self.Args("term_id").MustInt64()
	if (offset > 0) && helper.IsHas(iTermId) {
		_Termmeta, _error := model.GetTermmetasViaTermId(offset, limit, iTermId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the GetTermmetasViaTermId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermmetasViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMetaKey := self.Args("meta_key").String()
	if (offset > 0) && helper.IsHas(iMetaKey) {
		_Termmeta, _error := model.GetTermmetasViaMetaKey(offset, limit, iMetaKey, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the GetTermmetasViaMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermmetasViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMetaValue := self.Args("meta_value").String()
	if (offset > 0) && helper.IsHas(iMetaValue) {
		_Termmeta, _error := model.GetTermmetasViaMetaValue(offset, limit, iMetaValue, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the GetTermmetasViaMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermmetasByMetaIdAndTermIdAndMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMetaId := self.Args("meta_id").MustInt64()
	iTermId := self.Args("term_id").MustInt64()
	iMetaKey := self.Args("meta_key").String()

	if helper.IsHas(iMetaId) {
		_Termmeta, _error := model.GetTermmetasByMetaIdAndTermIdAndMetaKey(offset, limit, iMetaId,iTermId,iMetaKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the GetTermmetasByMetaIdAndTermIdAndMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermmetasByMetaIdAndTermIdAndMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMetaId := self.Args("meta_id").MustInt64()
	iTermId := self.Args("term_id").MustInt64()
	iMetaValue := self.Args("meta_value").String()

	if helper.IsHas(iMetaId) {
		_Termmeta, _error := model.GetTermmetasByMetaIdAndTermIdAndMetaValue(offset, limit, iMetaId,iTermId,iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the GetTermmetasByMetaIdAndTermIdAndMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermmetasByMetaIdAndMetaKeyAndMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMetaId := self.Args("meta_id").MustInt64()
	iMetaKey := self.Args("meta_key").String()
	iMetaValue := self.Args("meta_value").String()

	if helper.IsHas(iMetaId) {
		_Termmeta, _error := model.GetTermmetasByMetaIdAndMetaKeyAndMetaValue(offset, limit, iMetaId,iMetaKey,iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the GetTermmetasByMetaIdAndMetaKeyAndMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermmetasByTermIdAndMetaKeyAndMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermId := self.Args("term_id").MustInt64()
	iMetaKey := self.Args("meta_key").String()
	iMetaValue := self.Args("meta_value").String()

	if helper.IsHas(iTermId) {
		_Termmeta, _error := model.GetTermmetasByTermIdAndMetaKeyAndMetaValue(offset, limit, iTermId,iMetaKey,iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the GetTermmetasByTermIdAndMetaKeyAndMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermmetasByMetaIdAndTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMetaId := self.Args("meta_id").MustInt64()
	iTermId := self.Args("term_id").MustInt64()

	if helper.IsHas(iMetaId) {
		_Termmeta, _error := model.GetTermmetasByMetaIdAndTermId(offset, limit, iMetaId,iTermId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the GetTermmetasByMetaIdAndTermId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermmetasByMetaIdAndMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMetaId := self.Args("meta_id").MustInt64()
	iMetaKey := self.Args("meta_key").String()

	if helper.IsHas(iMetaId) {
		_Termmeta, _error := model.GetTermmetasByMetaIdAndMetaKey(offset, limit, iMetaId,iMetaKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the GetTermmetasByMetaIdAndMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermmetasByMetaIdAndMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMetaId := self.Args("meta_id").MustInt64()
	iMetaValue := self.Args("meta_value").String()

	if helper.IsHas(iMetaId) {
		_Termmeta, _error := model.GetTermmetasByMetaIdAndMetaValue(offset, limit, iMetaId,iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the GetTermmetasByMetaIdAndMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermmetasByTermIdAndMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermId := self.Args("term_id").MustInt64()
	iMetaKey := self.Args("meta_key").String()

	if helper.IsHas(iTermId) {
		_Termmeta, _error := model.GetTermmetasByTermIdAndMetaKey(offset, limit, iTermId,iMetaKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the GetTermmetasByTermIdAndMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermmetasByTermIdAndMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermId := self.Args("term_id").MustInt64()
	iMetaValue := self.Args("meta_value").String()

	if helper.IsHas(iTermId) {
		_Termmeta, _error := model.GetTermmetasByTermIdAndMetaValue(offset, limit, iTermId,iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the GetTermmetasByTermIdAndMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermmetasByMetaKeyAndMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMetaKey := self.Args("meta_key").String()
	iMetaValue := self.Args("meta_value").String()

	if helper.IsHas(iMetaKey) {
		_Termmeta, _error := model.GetTermmetasByMetaKeyAndMetaValue(offset, limit, iMetaKey,iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the GetTermmetasByMetaKeyAndMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermmetasHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_Termmeta, _error := model.GetTermmetas(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the GetTermmetas' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTermmetaViaMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaId := self.Args("meta_id").MustInt64()
	if helper.IsHas(iMetaId) {
		_Termmeta := model.HasTermmetaViaMetaId(iMetaId)
		var m = map[string]interface{}{}
		m["termmeta"] = _Termmeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermmetaViaMetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTermmetaViaTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTermId := self.Args("term_id").MustInt64()
	if helper.IsHas(iTermId) {
		_Termmeta := model.HasTermmetaViaTermId(iTermId)
		var m = map[string]interface{}{}
		m["termmeta"] = _Termmeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermmetaViaTermId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTermmetaViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaKey := self.Args("meta_key").String()
	if helper.IsHas(iMetaKey) {
		_Termmeta := model.HasTermmetaViaMetaKey(iMetaKey)
		var m = map[string]interface{}{}
		m["termmeta"] = _Termmeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermmetaViaMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTermmetaViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaValue := self.Args("meta_value").String()
	if helper.IsHas(iMetaValue) {
		_Termmeta := model.HasTermmetaViaMetaValue(iMetaValue)
		var m = map[string]interface{}{}
		m["termmeta"] = _Termmeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermmetaViaMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermmetaViaMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaId := self.Args("meta_id").MustInt64()
	if helper.IsHas(iMetaId) {
		_Termmeta, _error := model.GetTermmetaViaMetaId(iMetaId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the GetTermmetaViaMetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermmetaViaTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTermId := self.Args("term_id").MustInt64()
	if helper.IsHas(iTermId) {
		_Termmeta, _error := model.GetTermmetaViaTermId(iTermId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the GetTermmetaViaTermId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermmetaViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaKey := self.Args("meta_key").String()
	if helper.IsHas(iMetaKey) {
		_Termmeta, _error := model.GetTermmetaViaMetaKey(iMetaKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the GetTermmetaViaMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermmetaViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaValue := self.Args("meta_value").String()
	if helper.IsHas(iMetaValue) {
		_Termmeta, _error := model.GetTermmetaViaMetaValue(iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the GetTermmetaViaMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermmetaViaMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	if helper.IsHas(MetaId_) {
		var iTermmeta model.Termmeta
		self.Bind(&iTermmeta)
		_Termmeta, _error := model.SetTermmetaViaMetaId(MetaId_, &iTermmeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the SetTermmetaViaMetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermmetaViaTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermId_ := self.Args("term_id").MustInt64()
	if helper.IsHas(TermId_) {
		var iTermmeta model.Termmeta
		self.Bind(&iTermmeta)
		_Termmeta, _error := model.SetTermmetaViaTermId(TermId_, &iTermmeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the SetTermmetaViaTermId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermmetaViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	if helper.IsHas(MetaKey_) {
		var iTermmeta model.Termmeta
		self.Bind(&iTermmeta)
		_Termmeta, _error := model.SetTermmetaViaMetaKey(MetaKey_, &iTermmeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the SetTermmetaViaMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermmetaViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	if helper.IsHas(MetaValue_) {
		var iTermmeta model.Termmeta
		self.Bind(&iTermmeta)
		_Termmeta, _error := model.SetTermmetaViaMetaValue(MetaValue_, &iTermmeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the SetTermmetaViaMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddTermmetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	TermId_ := self.Args("term_id").MustInt64()
	MetaKey_ := self.Args("meta_key").String()
	MetaValue_ := self.Args("meta_value").String()

	if helper.IsHas(MetaId_) {
		_error := model.AddTermmeta(MetaId_,TermId_,MetaKey_,MetaValue_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddTermmeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostTermmetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iTermmeta model.Termmeta
	self.Bind(&iTermmeta)
	_int64, _error := model.PostTermmeta(&iTermmeta)
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

func PutTermmetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iTermmeta model.Termmeta
	self.Bind(&iTermmeta)
	_int64, _error := model.PutTermmeta(&iTermmeta)
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

func PutTermmetaViaMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	var iTermmeta model.Termmeta
	self.Bind(&iTermmeta)
	_int64, _error := model.PutTermmetaViaMetaId(MetaId_, &iTermmeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTermmetaViaTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermId_ := self.Args("term_id").MustInt64()
	var iTermmeta model.Termmeta
	self.Bind(&iTermmeta)
	_int64, _error := model.PutTermmetaViaTermId(TermId_, &iTermmeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTermmetaViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	var iTermmeta model.Termmeta
	self.Bind(&iTermmeta)
	_int64, _error := model.PutTermmetaViaMetaKey(MetaKey_, &iTermmeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTermmetaViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	var iTermmeta model.Termmeta
	self.Bind(&iTermmeta)
	_int64, _error := model.PutTermmetaViaMetaValue(MetaValue_, &iTermmeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermmetaViaMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	var iTermmeta model.Termmeta
	self.Bind(&iTermmeta)
	var iMap = helper.StructToMap(iTermmeta)
	_error := model.UpdateTermmetaViaMetaId(MetaId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermmetaViaTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermId_ := self.Args("term_id").MustInt64()
	var iTermmeta model.Termmeta
	self.Bind(&iTermmeta)
	var iMap = helper.StructToMap(iTermmeta)
	_error := model.UpdateTermmetaViaTermId(TermId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermmetaViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	var iTermmeta model.Termmeta
	self.Bind(&iTermmeta)
	var iMap = helper.StructToMap(iTermmeta)
	_error := model.UpdateTermmetaViaMetaKey(MetaKey_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermmetaViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	var iTermmeta model.Termmeta
	self.Bind(&iTermmeta)
	var iMap = helper.StructToMap(iTermmeta)
	_error := model.UpdateTermmetaViaMetaValue(MetaValue_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTermmetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	_int64, _error := model.DeleteTermmeta(MetaId_)
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

func DeleteTermmetaViaMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	_error := model.DeleteTermmetaViaMetaId(MetaId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTermmetaViaTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermId_ := self.Args("term_id").MustInt64()
	_error := model.DeleteTermmetaViaTermId(TermId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTermmetaViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	_error := model.DeleteTermmetaViaMetaKey(MetaKey_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTermmetaViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	_error := model.DeleteTermmetaViaMetaValue(MetaValue_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
