package handler

import (
	"github.com/insionng/zenpress/helper"
	"github.com/insionng/zenpress/model"
	"github.com/insionng/macross"
)

func GetCommentmetasCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetCommentmetasCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["commentmetasCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetCommentmetasCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentmetaCountViaMetaIdHandler(self *macross.Context) error {
	MetaId_ := self.Args("meta_id").MustInt64()
	_int64 := model.GetCommentmetaCountViaMetaId(MetaId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["commentmetaCount"] = 0
	}
	m["commentmetaCount"] = _int64
	return self.JSON(m)
}

func GetCommentmetaCountViaCommentIdHandler(self *macross.Context) error {
	CommentId_ := self.Args("comment_id").MustInt64()
	_int64 := model.GetCommentmetaCountViaCommentId(CommentId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["commentmetaCount"] = 0
	}
	m["commentmetaCount"] = _int64
	return self.JSON(m)
}

func GetCommentmetaCountViaMetaKeyHandler(self *macross.Context) error {
	MetaKey_ := self.Args("meta_key").String()
	_int64 := model.GetCommentmetaCountViaMetaKey(MetaKey_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["commentmetaCount"] = 0
	}
	m["commentmetaCount"] = _int64
	return self.JSON(m)
}

func GetCommentmetaCountViaMetaValueHandler(self *macross.Context) error {
	MetaValue_ := self.Args("meta_value").String()
	_int64 := model.GetCommentmetaCountViaMetaValue(MetaValue_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["commentmetaCount"] = 0
	}
	m["commentmetaCount"] = _int64
	return self.JSON(m)
}

func GetCommentmetasViaMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMetaId := self.Args("meta_id").MustInt64()
	if (offset > 0) && helper.IsHas(iMetaId) {
		_Commentmeta, _error := model.GetCommentmetasViaMetaId(offset, limit, iMetaId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the GetCommentmetasViaMetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentmetasViaCommentIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCommentId := self.Args("comment_id").MustInt64()
	if (offset > 0) && helper.IsHas(iCommentId) {
		_Commentmeta, _error := model.GetCommentmetasViaCommentId(offset, limit, iCommentId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the GetCommentmetasViaCommentId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentmetasViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMetaKey := self.Args("meta_key").String()
	if (offset > 0) && helper.IsHas(iMetaKey) {
		_Commentmeta, _error := model.GetCommentmetasViaMetaKey(offset, limit, iMetaKey, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the GetCommentmetasViaMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentmetasViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMetaValue := self.Args("meta_value").String()
	if (offset > 0) && helper.IsHas(iMetaValue) {
		_Commentmeta, _error := model.GetCommentmetasViaMetaValue(offset, limit, iMetaValue, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the GetCommentmetasViaMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentmetasByMetaIdAndCommentIdAndMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMetaId := self.Args("meta_id").MustInt64()
	iCommentId := self.Args("comment_id").MustInt64()
	iMetaKey := self.Args("meta_key").String()

	if helper.IsHas(iMetaId) {
		_Commentmeta, _error := model.GetCommentmetasByMetaIdAndCommentIdAndMetaKey(offset, limit, iMetaId,iCommentId,iMetaKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the GetCommentmetasByMetaIdAndCommentIdAndMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentmetasByMetaIdAndCommentIdAndMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMetaId := self.Args("meta_id").MustInt64()
	iCommentId := self.Args("comment_id").MustInt64()
	iMetaValue := self.Args("meta_value").String()

	if helper.IsHas(iMetaId) {
		_Commentmeta, _error := model.GetCommentmetasByMetaIdAndCommentIdAndMetaValue(offset, limit, iMetaId,iCommentId,iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the GetCommentmetasByMetaIdAndCommentIdAndMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentmetasByMetaIdAndMetaKeyAndMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMetaId := self.Args("meta_id").MustInt64()
	iMetaKey := self.Args("meta_key").String()
	iMetaValue := self.Args("meta_value").String()

	if helper.IsHas(iMetaId) {
		_Commentmeta, _error := model.GetCommentmetasByMetaIdAndMetaKeyAndMetaValue(offset, limit, iMetaId,iMetaKey,iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the GetCommentmetasByMetaIdAndMetaKeyAndMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentmetasByCommentIdAndMetaKeyAndMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iMetaKey := self.Args("meta_key").String()
	iMetaValue := self.Args("meta_value").String()

	if helper.IsHas(iCommentId) {
		_Commentmeta, _error := model.GetCommentmetasByCommentIdAndMetaKeyAndMetaValue(offset, limit, iCommentId,iMetaKey,iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the GetCommentmetasByCommentIdAndMetaKeyAndMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentmetasByMetaIdAndCommentIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMetaId := self.Args("meta_id").MustInt64()
	iCommentId := self.Args("comment_id").MustInt64()

	if helper.IsHas(iMetaId) {
		_Commentmeta, _error := model.GetCommentmetasByMetaIdAndCommentId(offset, limit, iMetaId,iCommentId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the GetCommentmetasByMetaIdAndCommentId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentmetasByMetaIdAndMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMetaId := self.Args("meta_id").MustInt64()
	iMetaKey := self.Args("meta_key").String()

	if helper.IsHas(iMetaId) {
		_Commentmeta, _error := model.GetCommentmetasByMetaIdAndMetaKey(offset, limit, iMetaId,iMetaKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the GetCommentmetasByMetaIdAndMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentmetasByMetaIdAndMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMetaId := self.Args("meta_id").MustInt64()
	iMetaValue := self.Args("meta_value").String()

	if helper.IsHas(iMetaId) {
		_Commentmeta, _error := model.GetCommentmetasByMetaIdAndMetaValue(offset, limit, iMetaId,iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the GetCommentmetasByMetaIdAndMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentmetasByCommentIdAndMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iMetaKey := self.Args("meta_key").String()

	if helper.IsHas(iCommentId) {
		_Commentmeta, _error := model.GetCommentmetasByCommentIdAndMetaKey(offset, limit, iCommentId,iMetaKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the GetCommentmetasByCommentIdAndMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentmetasByCommentIdAndMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iCommentId := self.Args("comment_id").MustInt64()
	iMetaValue := self.Args("meta_value").String()

	if helper.IsHas(iCommentId) {
		_Commentmeta, _error := model.GetCommentmetasByCommentIdAndMetaValue(offset, limit, iCommentId,iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the GetCommentmetasByCommentIdAndMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentmetasByMetaKeyAndMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMetaKey := self.Args("meta_key").String()
	iMetaValue := self.Args("meta_value").String()

	if helper.IsHas(iMetaKey) {
		_Commentmeta, _error := model.GetCommentmetasByMetaKeyAndMetaValue(offset, limit, iMetaKey,iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the GetCommentmetasByMetaKeyAndMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentmetasHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_Commentmeta, _error := model.GetCommentmetas(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the GetCommentmetas' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentmetaViaMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaId := self.Args("meta_id").MustInt64()
	if helper.IsHas(iMetaId) {
		_Commentmeta := model.HasCommentmetaViaMetaId(iMetaId)
		var m = map[string]interface{}{}
		m["commentmeta"] = _Commentmeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentmetaViaMetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentmetaViaCommentIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentId := self.Args("comment_id").MustInt64()
	if helper.IsHas(iCommentId) {
		_Commentmeta := model.HasCommentmetaViaCommentId(iCommentId)
		var m = map[string]interface{}{}
		m["commentmeta"] = _Commentmeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentmetaViaCommentId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentmetaViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaKey := self.Args("meta_key").String()
	if helper.IsHas(iMetaKey) {
		_Commentmeta := model.HasCommentmetaViaMetaKey(iMetaKey)
		var m = map[string]interface{}{}
		m["commentmeta"] = _Commentmeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentmetaViaMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentmetaViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaValue := self.Args("meta_value").String()
	if helper.IsHas(iMetaValue) {
		_Commentmeta := model.HasCommentmetaViaMetaValue(iMetaValue)
		var m = map[string]interface{}{}
		m["commentmeta"] = _Commentmeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentmetaViaMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentmetaViaMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaId := self.Args("meta_id").MustInt64()
	if helper.IsHas(iMetaId) {
		_Commentmeta, _error := model.GetCommentmetaViaMetaId(iMetaId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the GetCommentmetaViaMetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentmetaViaCommentIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentId := self.Args("comment_id").MustInt64()
	if helper.IsHas(iCommentId) {
		_Commentmeta, _error := model.GetCommentmetaViaCommentId(iCommentId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the GetCommentmetaViaCommentId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentmetaViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaKey := self.Args("meta_key").String()
	if helper.IsHas(iMetaKey) {
		_Commentmeta, _error := model.GetCommentmetaViaMetaKey(iMetaKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the GetCommentmetaViaMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentmetaViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaValue := self.Args("meta_value").String()
	if helper.IsHas(iMetaValue) {
		_Commentmeta, _error := model.GetCommentmetaViaMetaValue(iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the GetCommentmetaViaMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentmetaViaMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	if helper.IsHas(MetaId_) {
		var iCommentmeta model.Commentmeta
		self.Bind(&iCommentmeta)
		_Commentmeta, _error := model.SetCommentmetaViaMetaId(MetaId_, &iCommentmeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the SetCommentmetaViaMetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentmetaViaCommentIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentId_ := self.Args("comment_id").MustInt64()
	if helper.IsHas(CommentId_) {
		var iCommentmeta model.Commentmeta
		self.Bind(&iCommentmeta)
		_Commentmeta, _error := model.SetCommentmetaViaCommentId(CommentId_, &iCommentmeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the SetCommentmetaViaCommentId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentmetaViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	if helper.IsHas(MetaKey_) {
		var iCommentmeta model.Commentmeta
		self.Bind(&iCommentmeta)
		_Commentmeta, _error := model.SetCommentmetaViaMetaKey(MetaKey_, &iCommentmeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the SetCommentmetaViaMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentmetaViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	if helper.IsHas(MetaValue_) {
		var iCommentmeta model.Commentmeta
		self.Bind(&iCommentmeta)
		_Commentmeta, _error := model.SetCommentmetaViaMetaValue(MetaValue_, &iCommentmeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the SetCommentmetaViaMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddCommentmetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	CommentId_ := self.Args("comment_id").MustInt64()
	MetaKey_ := self.Args("meta_key").String()
	MetaValue_ := self.Args("meta_value").String()

	if helper.IsHas(MetaId_) {
		_error := model.AddCommentmeta(MetaId_,CommentId_,MetaKey_,MetaValue_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddCommentmeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostCommentmetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iCommentmeta model.Commentmeta
	self.Bind(&iCommentmeta)
	_int64, _error := model.PostCommentmeta(&iCommentmeta)
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

func PutCommentmetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iCommentmeta model.Commentmeta
	self.Bind(&iCommentmeta)
	_int64, _error := model.PutCommentmeta(&iCommentmeta)
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

func PutCommentmetaViaMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	var iCommentmeta model.Commentmeta
	self.Bind(&iCommentmeta)
	_int64, _error := model.PutCommentmetaViaMetaId(MetaId_, &iCommentmeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentmetaViaCommentIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentId_ := self.Args("comment_id").MustInt64()
	var iCommentmeta model.Commentmeta
	self.Bind(&iCommentmeta)
	_int64, _error := model.PutCommentmetaViaCommentId(CommentId_, &iCommentmeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentmetaViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	var iCommentmeta model.Commentmeta
	self.Bind(&iCommentmeta)
	_int64, _error := model.PutCommentmetaViaMetaKey(MetaKey_, &iCommentmeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentmetaViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	var iCommentmeta model.Commentmeta
	self.Bind(&iCommentmeta)
	_int64, _error := model.PutCommentmetaViaMetaValue(MetaValue_, &iCommentmeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentmetaViaMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	var iCommentmeta model.Commentmeta
	self.Bind(&iCommentmeta)
	var iMap = helper.StructToMap(iCommentmeta)
	_error := model.UpdateCommentmetaViaMetaId(MetaId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentmetaViaCommentIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentId_ := self.Args("comment_id").MustInt64()
	var iCommentmeta model.Commentmeta
	self.Bind(&iCommentmeta)
	var iMap = helper.StructToMap(iCommentmeta)
	_error := model.UpdateCommentmetaViaCommentId(CommentId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentmetaViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	var iCommentmeta model.Commentmeta
	self.Bind(&iCommentmeta)
	var iMap = helper.StructToMap(iCommentmeta)
	_error := model.UpdateCommentmetaViaMetaKey(MetaKey_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentmetaViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	var iCommentmeta model.Commentmeta
	self.Bind(&iCommentmeta)
	var iMap = helper.StructToMap(iCommentmeta)
	_error := model.UpdateCommentmetaViaMetaValue(MetaValue_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentmetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	_int64, _error := model.DeleteCommentmeta(MetaId_)
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

func DeleteCommentmetaViaMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	_error := model.DeleteCommentmetaViaMetaId(MetaId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentmetaViaCommentIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentId_ := self.Args("comment_id").MustInt64()
	_error := model.DeleteCommentmetaViaCommentId(CommentId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentmetaViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	_error := model.DeleteCommentmetaViaMetaKey(MetaKey_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentmetaViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	_error := model.DeleteCommentmetaViaMetaValue(MetaValue_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
