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

func GetTermmetaCountByMetaIdHandler(self *macross.Context) error {
	MetaId_ := self.Args("meta_id").MustInt64()
	_int64 := model.GetTermmetaCountByMetaId(MetaId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["termmetaCount"] = 0
	}
	m["termmetaCount"] = _int64
	return self.JSON(m)
}

func GetTermmetaCountByTermIdHandler(self *macross.Context) error {
	TermId_ := self.Args("term_id").MustInt64()
	_int64 := model.GetTermmetaCountByTermId(TermId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["termmetaCount"] = 0
	}
	m["termmetaCount"] = _int64
	return self.JSON(m)
}

func GetTermmetaCountByMetaKeyHandler(self *macross.Context) error {
	MetaKey_ := self.Args("meta_key").String()
	_int64 := model.GetTermmetaCountByMetaKey(MetaKey_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["termmetaCount"] = 0
	}
	m["termmetaCount"] = _int64
	return self.JSON(m)
}

func GetTermmetaCountByMetaValueHandler(self *macross.Context) error {
	MetaValue_ := self.Args("meta_value").String()
	_int64 := model.GetTermmetaCountByMetaValue(MetaValue_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["termmetaCount"] = 0
	}
	m["termmetaCount"] = _int64
	return self.JSON(m)
}

func GetTermmetasByMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMetaId := self.Args("meta_id").MustInt64()
	if (offset > 0) && helper.IsHas(iMetaId) {
		_Termmeta, _error := model.GetTermmetasByMetaId(offset, limit, iMetaId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the GetTermmetasByMetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermmetasByTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTermId := self.Args("term_id").MustInt64()
	if (offset > 0) && helper.IsHas(iTermId) {
		_Termmeta, _error := model.GetTermmetasByTermId(offset, limit, iTermId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the GetTermmetasByTermId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermmetasByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMetaKey := self.Args("meta_key").String()
	if (offset > 0) && helper.IsHas(iMetaKey) {
		_Termmeta, _error := model.GetTermmetasByMetaKey(offset, limit, iMetaKey, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the GetTermmetasByMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermmetasByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMetaValue := self.Args("meta_value").String()
	if (offset > 0) && helper.IsHas(iMetaValue) {
		_Termmeta, _error := model.GetTermmetasByMetaValue(offset, limit, iMetaValue, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the GetTermmetasByMetaValue's args."
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

func GetHasTermmetaByMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaId := self.Args("meta_id").MustInt64()
	if helper.IsHas(iMetaId) {
		_Termmeta := model.HasTermmetaByMetaId(iMetaId)
		var m = map[string]interface{}{}
		m["termmeta"] = _Termmeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermmetaByMetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTermmetaByTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTermId := self.Args("term_id").MustInt64()
	if helper.IsHas(iTermId) {
		_Termmeta := model.HasTermmetaByTermId(iTermId)
		var m = map[string]interface{}{}
		m["termmeta"] = _Termmeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermmetaByTermId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTermmetaByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaKey := self.Args("meta_key").String()
	if helper.IsHas(iMetaKey) {
		_Termmeta := model.HasTermmetaByMetaKey(iMetaKey)
		var m = map[string]interface{}{}
		m["termmeta"] = _Termmeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermmetaByMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTermmetaByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaValue := self.Args("meta_value").String()
	if helper.IsHas(iMetaValue) {
		_Termmeta := model.HasTermmetaByMetaValue(iMetaValue)
		var m = map[string]interface{}{}
		m["termmeta"] = _Termmeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermmetaByMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermmetaByMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaId := self.Args("meta_id").MustInt64()
	if helper.IsHas(iMetaId) {
		_Termmeta, _error := model.GetTermmetaByMetaId(iMetaId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the GetTermmetaByMetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermmetaByTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTermId := self.Args("term_id").MustInt64()
	if helper.IsHas(iTermId) {
		_Termmeta, _error := model.GetTermmetaByTermId(iTermId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the GetTermmetaByTermId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermmetaByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaKey := self.Args("meta_key").String()
	if helper.IsHas(iMetaKey) {
		_Termmeta, _error := model.GetTermmetaByMetaKey(iMetaKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the GetTermmetaByMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermmetaByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaValue := self.Args("meta_value").String()
	if helper.IsHas(iMetaValue) {
		_Termmeta, _error := model.GetTermmetaByMetaValue(iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the GetTermmetaByMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermmetaByMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	if helper.IsHas(MetaId_) {
		var iTermmeta model.Termmeta
		self.Bind(&iTermmeta)
		_Termmeta, _error := model.SetTermmetaByMetaId(MetaId_, &iTermmeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the SetTermmetaByMetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermmetaByTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermId_ := self.Args("term_id").MustInt64()
	if helper.IsHas(TermId_) {
		var iTermmeta model.Termmeta
		self.Bind(&iTermmeta)
		_Termmeta, _error := model.SetTermmetaByTermId(TermId_, &iTermmeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the SetTermmetaByTermId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermmetaByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	if helper.IsHas(MetaKey_) {
		var iTermmeta model.Termmeta
		self.Bind(&iTermmeta)
		_Termmeta, _error := model.SetTermmetaByMetaKey(MetaKey_, &iTermmeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the SetTermmetaByMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermmetaByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	if helper.IsHas(MetaValue_) {
		var iTermmeta model.Termmeta
		self.Bind(&iTermmeta)
		_Termmeta, _error := model.SetTermmetaByMetaValue(MetaValue_, &iTermmeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Termmeta)
	}
	herr.Message = "Can't get to the SetTermmetaByMetaValue's args."
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

func PutTermmetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iTermmeta model.Termmeta
	self.Bind(&iTermmeta)
	_int64, _error := model.PutTermmeta(&iTermmeta)
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

func PutTermmetaByMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	var iTermmeta model.Termmeta
	self.Bind(&iTermmeta)
	_int64, _error := model.PutTermmetaByMetaId(MetaId_, &iTermmeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTermmetaByTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermId_ := self.Args("term_id").MustInt64()
	var iTermmeta model.Termmeta
	self.Bind(&iTermmeta)
	_int64, _error := model.PutTermmetaByTermId(TermId_, &iTermmeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTermmetaByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	var iTermmeta model.Termmeta
	self.Bind(&iTermmeta)
	_int64, _error := model.PutTermmetaByMetaKey(MetaKey_, &iTermmeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTermmetaByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	var iTermmeta model.Termmeta
	self.Bind(&iTermmeta)
	_int64, _error := model.PutTermmetaByMetaValue(MetaValue_, &iTermmeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermmetaByMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	var iTermmeta model.Termmeta
	self.Bind(&iTermmeta)
	var iMap = helper.StructToMap(iTermmeta)
	_error := model.UpdateTermmetaByMetaId(MetaId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermmetaByTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermId_ := self.Args("term_id").MustInt64()
	var iTermmeta model.Termmeta
	self.Bind(&iTermmeta)
	var iMap = helper.StructToMap(iTermmeta)
	_error := model.UpdateTermmetaByTermId(TermId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermmetaByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	var iTermmeta model.Termmeta
	self.Bind(&iTermmeta)
	var iMap = helper.StructToMap(iTermmeta)
	_error := model.UpdateTermmetaByMetaKey(MetaKey_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermmetaByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	var iTermmeta model.Termmeta
	self.Bind(&iTermmeta)
	var iMap = helper.StructToMap(iTermmeta)
	_error := model.UpdateTermmetaByMetaValue(MetaValue_, &iMap)
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

func DeleteTermmetaByMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	_error := model.DeleteTermmetaByMetaId(MetaId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTermmetaByTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermId_ := self.Args("term_id").MustInt64()
	_error := model.DeleteTermmetaByTermId(TermId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTermmetaByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	_error := model.DeleteTermmetaByMetaKey(MetaKey_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTermmetaByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	_error := model.DeleteTermmetaByMetaValue(MetaValue_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
