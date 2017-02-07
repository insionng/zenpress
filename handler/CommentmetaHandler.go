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

func GetCommentmetaCountByMetaIdHandler(self *macross.Context) error {
	MetaId_ := self.Args("meta_id").MustInt64()
	_int64 := model.GetCommentmetaCountByMetaId(MetaId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["commentmetaCount"] = 0
	}
	m["commentmetaCount"] = _int64
	return self.JSON(m)
}

func GetCommentmetaCountByCommentIdHandler(self *macross.Context) error {
	CommentId_ := self.Args("comment_id").MustInt64()
	_int64 := model.GetCommentmetaCountByCommentId(CommentId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["commentmetaCount"] = 0
	}
	m["commentmetaCount"] = _int64
	return self.JSON(m)
}

func GetCommentmetaCountByMetaKeyHandler(self *macross.Context) error {
	MetaKey_ := self.Args("meta_key").String()
	_int64 := model.GetCommentmetaCountByMetaKey(MetaKey_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["commentmetaCount"] = 0
	}
	m["commentmetaCount"] = _int64
	return self.JSON(m)
}

func GetCommentmetaCountByMetaValueHandler(self *macross.Context) error {
	MetaValue_ := self.Args("meta_value").String()
	_int64 := model.GetCommentmetaCountByMetaValue(MetaValue_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["commentmetaCount"] = 0
	}
	m["commentmetaCount"] = _int64
	return self.JSON(m)
}

func GetCommentmetasByMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMetaId := self.Args("meta_id").MustInt64()
	if (offset > 0) && helper.IsHas(iMetaId) {
		_Commentmeta, _error := model.GetCommentmetasByMetaId(offset, limit, iMetaId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the GetCommentmetasByMetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentmetasByCommentIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCommentId := self.Args("comment_id").MustInt64()
	if (offset > 0) && helper.IsHas(iCommentId) {
		_Commentmeta, _error := model.GetCommentmetasByCommentId(offset, limit, iCommentId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the GetCommentmetasByCommentId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentmetasByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMetaKey := self.Args("meta_key").String()
	if (offset > 0) && helper.IsHas(iMetaKey) {
		_Commentmeta, _error := model.GetCommentmetasByMetaKey(offset, limit, iMetaKey, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the GetCommentmetasByMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentmetasByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMetaValue := self.Args("meta_value").String()
	if (offset > 0) && helper.IsHas(iMetaValue) {
		_Commentmeta, _error := model.GetCommentmetasByMetaValue(offset, limit, iMetaValue, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the GetCommentmetasByMetaValue's args."
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

func GetHasCommentmetaByMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaId := self.Args("meta_id").MustInt64()
	if helper.IsHas(iMetaId) {
		_Commentmeta := model.HasCommentmetaByMetaId(iMetaId)
		var m = map[string]interface{}{}
		m["commentmeta"] = _Commentmeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentmetaByMetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentmetaByCommentIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentId := self.Args("comment_id").MustInt64()
	if helper.IsHas(iCommentId) {
		_Commentmeta := model.HasCommentmetaByCommentId(iCommentId)
		var m = map[string]interface{}{}
		m["commentmeta"] = _Commentmeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentmetaByCommentId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentmetaByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaKey := self.Args("meta_key").String()
	if helper.IsHas(iMetaKey) {
		_Commentmeta := model.HasCommentmetaByMetaKey(iMetaKey)
		var m = map[string]interface{}{}
		m["commentmeta"] = _Commentmeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentmetaByMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasCommentmetaByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaValue := self.Args("meta_value").String()
	if helper.IsHas(iMetaValue) {
		_Commentmeta := model.HasCommentmetaByMetaValue(iMetaValue)
		var m = map[string]interface{}{}
		m["commentmeta"] = _Commentmeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasCommentmetaByMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentmetaByMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaId := self.Args("meta_id").MustInt64()
	if helper.IsHas(iMetaId) {
		_Commentmeta, _error := model.GetCommentmetaByMetaId(iMetaId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the GetCommentmetaByMetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentmetaByCommentIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCommentId := self.Args("comment_id").MustInt64()
	if helper.IsHas(iCommentId) {
		_Commentmeta, _error := model.GetCommentmetaByCommentId(iCommentId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the GetCommentmetaByCommentId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentmetaByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaKey := self.Args("meta_key").String()
	if helper.IsHas(iMetaKey) {
		_Commentmeta, _error := model.GetCommentmetaByMetaKey(iMetaKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the GetCommentmetaByMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetCommentmetaByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaValue := self.Args("meta_value").String()
	if helper.IsHas(iMetaValue) {
		_Commentmeta, _error := model.GetCommentmetaByMetaValue(iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the GetCommentmetaByMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentmetaByMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	if helper.IsHas(MetaId_) {
		var iCommentmeta model.Commentmeta
		self.Bind(&iCommentmeta)
		_Commentmeta, _error := model.SetCommentmetaByMetaId(MetaId_, &iCommentmeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the SetCommentmetaByMetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentmetaByCommentIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentId_ := self.Args("comment_id").MustInt64()
	if helper.IsHas(CommentId_) {
		var iCommentmeta model.Commentmeta
		self.Bind(&iCommentmeta)
		_Commentmeta, _error := model.SetCommentmetaByCommentId(CommentId_, &iCommentmeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the SetCommentmetaByCommentId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentmetaByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	if helper.IsHas(MetaKey_) {
		var iCommentmeta model.Commentmeta
		self.Bind(&iCommentmeta)
		_Commentmeta, _error := model.SetCommentmetaByMetaKey(MetaKey_, &iCommentmeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the SetCommentmetaByMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetCommentmetaByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	if helper.IsHas(MetaValue_) {
		var iCommentmeta model.Commentmeta
		self.Bind(&iCommentmeta)
		_Commentmeta, _error := model.SetCommentmetaByMetaValue(MetaValue_, &iCommentmeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Commentmeta)
	}
	herr.Message = "Can't get to the SetCommentmetaByMetaValue's args."
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

func PutCommentmetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iCommentmeta model.Commentmeta
	self.Bind(&iCommentmeta)
	_int64, _error := model.PutCommentmeta(&iCommentmeta)
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

func PutCommentmetaByMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	var iCommentmeta model.Commentmeta
	self.Bind(&iCommentmeta)
	_int64, _error := model.PutCommentmetaByMetaId(MetaId_, &iCommentmeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentmetaByCommentIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentId_ := self.Args("comment_id").MustInt64()
	var iCommentmeta model.Commentmeta
	self.Bind(&iCommentmeta)
	_int64, _error := model.PutCommentmetaByCommentId(CommentId_, &iCommentmeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentmetaByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	var iCommentmeta model.Commentmeta
	self.Bind(&iCommentmeta)
	_int64, _error := model.PutCommentmetaByMetaKey(MetaKey_, &iCommentmeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutCommentmetaByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	var iCommentmeta model.Commentmeta
	self.Bind(&iCommentmeta)
	_int64, _error := model.PutCommentmetaByMetaValue(MetaValue_, &iCommentmeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentmetaByMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	var iCommentmeta model.Commentmeta
	self.Bind(&iCommentmeta)
	var iMap = helper.StructToMap(iCommentmeta)
	_error := model.UpdateCommentmetaByMetaId(MetaId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentmetaByCommentIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentId_ := self.Args("comment_id").MustInt64()
	var iCommentmeta model.Commentmeta
	self.Bind(&iCommentmeta)
	var iMap = helper.StructToMap(iCommentmeta)
	_error := model.UpdateCommentmetaByCommentId(CommentId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentmetaByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	var iCommentmeta model.Commentmeta
	self.Bind(&iCommentmeta)
	var iMap = helper.StructToMap(iCommentmeta)
	_error := model.UpdateCommentmetaByMetaKey(MetaKey_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateCommentmetaByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	var iCommentmeta model.Commentmeta
	self.Bind(&iCommentmeta)
	var iMap = helper.StructToMap(iCommentmeta)
	_error := model.UpdateCommentmetaByMetaValue(MetaValue_, &iMap)
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

func DeleteCommentmetaByMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	_error := model.DeleteCommentmetaByMetaId(MetaId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentmetaByCommentIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	CommentId_ := self.Args("comment_id").MustInt64()
	_error := model.DeleteCommentmetaByCommentId(CommentId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentmetaByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	_error := model.DeleteCommentmetaByMetaKey(MetaKey_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteCommentmetaByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	_error := model.DeleteCommentmetaByMetaValue(MetaValue_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
