package handler

import (
	"github.com/insionng/zenpress/helper"
	"github.com/insionng/zenpress/model"
	"github.com/insionng/macross"
)

func GetPostmetasCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetPostmetasCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["postmetasCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetPostmetasCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetPostmetaCountViaMetaIdHandler(self *macross.Context) error {
	MetaId_ := self.Args("meta_id").MustInt64()
	_int64 := model.GetPostmetaCountViaMetaId(MetaId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["postmetaCount"] = 0
	}
	m["postmetaCount"] = _int64
	return self.JSON(m)
}

func GetPostmetaCountViaPostIdHandler(self *macross.Context) error {
	PostId_ := self.Args("post_id").MustInt64()
	_int64 := model.GetPostmetaCountViaPostId(PostId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["postmetaCount"] = 0
	}
	m["postmetaCount"] = _int64
	return self.JSON(m)
}

func GetPostmetaCountViaMetaKeyHandler(self *macross.Context) error {
	MetaKey_ := self.Args("meta_key").String()
	_int64 := model.GetPostmetaCountViaMetaKey(MetaKey_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["postmetaCount"] = 0
	}
	m["postmetaCount"] = _int64
	return self.JSON(m)
}

func GetPostmetaCountViaMetaValueHandler(self *macross.Context) error {
	MetaValue_ := self.Args("meta_value").String()
	_int64 := model.GetPostmetaCountViaMetaValue(MetaValue_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["postmetaCount"] = 0
	}
	m["postmetaCount"] = _int64
	return self.JSON(m)
}

func GetPostmetasViaMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMetaId := self.Args("meta_id").MustInt64()
	if (offset > 0) && helper.IsHas(iMetaId) {
		_Postmeta, _error := model.GetPostmetasViaMetaId(offset, limit, iMetaId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the GetPostmetasViaMetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetPostmetasViaPostIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iPostId := self.Args("post_id").MustInt64()
	if (offset > 0) && helper.IsHas(iPostId) {
		_Postmeta, _error := model.GetPostmetasViaPostId(offset, limit, iPostId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the GetPostmetasViaPostId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetPostmetasViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMetaKey := self.Args("meta_key").String()
	if (offset > 0) && helper.IsHas(iMetaKey) {
		_Postmeta, _error := model.GetPostmetasViaMetaKey(offset, limit, iMetaKey, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the GetPostmetasViaMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetPostmetasViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMetaValue := self.Args("meta_value").String()
	if (offset > 0) && helper.IsHas(iMetaValue) {
		_Postmeta, _error := model.GetPostmetasViaMetaValue(offset, limit, iMetaValue, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the GetPostmetasViaMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetPostmetasByMetaIdAndPostIdAndMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMetaId := self.Args("meta_id").MustInt64()
	iPostId := self.Args("post_id").MustInt64()
	iMetaKey := self.Args("meta_key").String()

	if helper.IsHas(iMetaId) {
		_Postmeta, _error := model.GetPostmetasByMetaIdAndPostIdAndMetaKey(offset, limit, iMetaId,iPostId,iMetaKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the GetPostmetasByMetaIdAndPostIdAndMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetPostmetasByMetaIdAndPostIdAndMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMetaId := self.Args("meta_id").MustInt64()
	iPostId := self.Args("post_id").MustInt64()
	iMetaValue := self.Args("meta_value").String()

	if helper.IsHas(iMetaId) {
		_Postmeta, _error := model.GetPostmetasByMetaIdAndPostIdAndMetaValue(offset, limit, iMetaId,iPostId,iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the GetPostmetasByMetaIdAndPostIdAndMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetPostmetasByMetaIdAndMetaKeyAndMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMetaId := self.Args("meta_id").MustInt64()
	iMetaKey := self.Args("meta_key").String()
	iMetaValue := self.Args("meta_value").String()

	if helper.IsHas(iMetaId) {
		_Postmeta, _error := model.GetPostmetasByMetaIdAndMetaKeyAndMetaValue(offset, limit, iMetaId,iMetaKey,iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the GetPostmetasByMetaIdAndMetaKeyAndMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetPostmetasByPostIdAndMetaKeyAndMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPostId := self.Args("post_id").MustInt64()
	iMetaKey := self.Args("meta_key").String()
	iMetaValue := self.Args("meta_value").String()

	if helper.IsHas(iPostId) {
		_Postmeta, _error := model.GetPostmetasByPostIdAndMetaKeyAndMetaValue(offset, limit, iPostId,iMetaKey,iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the GetPostmetasByPostIdAndMetaKeyAndMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetPostmetasByMetaIdAndPostIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMetaId := self.Args("meta_id").MustInt64()
	iPostId := self.Args("post_id").MustInt64()

	if helper.IsHas(iMetaId) {
		_Postmeta, _error := model.GetPostmetasByMetaIdAndPostId(offset, limit, iMetaId,iPostId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the GetPostmetasByMetaIdAndPostId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetPostmetasByMetaIdAndMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMetaId := self.Args("meta_id").MustInt64()
	iMetaKey := self.Args("meta_key").String()

	if helper.IsHas(iMetaId) {
		_Postmeta, _error := model.GetPostmetasByMetaIdAndMetaKey(offset, limit, iMetaId,iMetaKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the GetPostmetasByMetaIdAndMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetPostmetasByMetaIdAndMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMetaId := self.Args("meta_id").MustInt64()
	iMetaValue := self.Args("meta_value").String()

	if helper.IsHas(iMetaId) {
		_Postmeta, _error := model.GetPostmetasByMetaIdAndMetaValue(offset, limit, iMetaId,iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the GetPostmetasByMetaIdAndMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetPostmetasByPostIdAndMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPostId := self.Args("post_id").MustInt64()
	iMetaKey := self.Args("meta_key").String()

	if helper.IsHas(iPostId) {
		_Postmeta, _error := model.GetPostmetasByPostIdAndMetaKey(offset, limit, iPostId,iMetaKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the GetPostmetasByPostIdAndMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetPostmetasByPostIdAndMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPostId := self.Args("post_id").MustInt64()
	iMetaValue := self.Args("meta_value").String()

	if helper.IsHas(iPostId) {
		_Postmeta, _error := model.GetPostmetasByPostIdAndMetaValue(offset, limit, iPostId,iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the GetPostmetasByPostIdAndMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetPostmetasByMetaKeyAndMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMetaKey := self.Args("meta_key").String()
	iMetaValue := self.Args("meta_value").String()

	if helper.IsHas(iMetaKey) {
		_Postmeta, _error := model.GetPostmetasByMetaKeyAndMetaValue(offset, limit, iMetaKey,iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the GetPostmetasByMetaKeyAndMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetPostmetasHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_Postmeta, _error := model.GetPostmetas(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the GetPostmetas' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasPostmetaViaMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaId := self.Args("meta_id").MustInt64()
	if helper.IsHas(iMetaId) {
		_Postmeta := model.HasPostmetaViaMetaId(iMetaId)
		var m = map[string]interface{}{}
		m["postmeta"] = _Postmeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasPostmetaViaMetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasPostmetaViaPostIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iPostId := self.Args("post_id").MustInt64()
	if helper.IsHas(iPostId) {
		_Postmeta := model.HasPostmetaViaPostId(iPostId)
		var m = map[string]interface{}{}
		m["postmeta"] = _Postmeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasPostmetaViaPostId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasPostmetaViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaKey := self.Args("meta_key").String()
	if helper.IsHas(iMetaKey) {
		_Postmeta := model.HasPostmetaViaMetaKey(iMetaKey)
		var m = map[string]interface{}{}
		m["postmeta"] = _Postmeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasPostmetaViaMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasPostmetaViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaValue := self.Args("meta_value").String()
	if helper.IsHas(iMetaValue) {
		_Postmeta := model.HasPostmetaViaMetaValue(iMetaValue)
		var m = map[string]interface{}{}
		m["postmeta"] = _Postmeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasPostmetaViaMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetPostmetaViaMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaId := self.Args("meta_id").MustInt64()
	if helper.IsHas(iMetaId) {
		_Postmeta, _error := model.GetPostmetaViaMetaId(iMetaId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the GetPostmetaViaMetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetPostmetaViaPostIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iPostId := self.Args("post_id").MustInt64()
	if helper.IsHas(iPostId) {
		_Postmeta, _error := model.GetPostmetaViaPostId(iPostId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the GetPostmetaViaPostId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetPostmetaViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaKey := self.Args("meta_key").String()
	if helper.IsHas(iMetaKey) {
		_Postmeta, _error := model.GetPostmetaViaMetaKey(iMetaKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the GetPostmetaViaMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetPostmetaViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaValue := self.Args("meta_value").String()
	if helper.IsHas(iMetaValue) {
		_Postmeta, _error := model.GetPostmetaViaMetaValue(iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the GetPostmetaViaMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetPostmetaViaMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	if helper.IsHas(MetaId_) {
		var iPostmeta model.Postmeta
		self.Bind(&iPostmeta)
		_Postmeta, _error := model.SetPostmetaViaMetaId(MetaId_, &iPostmeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the SetPostmetaViaMetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetPostmetaViaPostIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	PostId_ := self.Args("post_id").MustInt64()
	if helper.IsHas(PostId_) {
		var iPostmeta model.Postmeta
		self.Bind(&iPostmeta)
		_Postmeta, _error := model.SetPostmetaViaPostId(PostId_, &iPostmeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the SetPostmetaViaPostId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetPostmetaViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	if helper.IsHas(MetaKey_) {
		var iPostmeta model.Postmeta
		self.Bind(&iPostmeta)
		_Postmeta, _error := model.SetPostmetaViaMetaKey(MetaKey_, &iPostmeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the SetPostmetaViaMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetPostmetaViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	if helper.IsHas(MetaValue_) {
		var iPostmeta model.Postmeta
		self.Bind(&iPostmeta)
		_Postmeta, _error := model.SetPostmetaViaMetaValue(MetaValue_, &iPostmeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the SetPostmetaViaMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddPostmetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	PostId_ := self.Args("post_id").MustInt64()
	MetaKey_ := self.Args("meta_key").String()
	MetaValue_ := self.Args("meta_value").String()

	if helper.IsHas(MetaId_) {
		_error := model.AddPostmeta(MetaId_,PostId_,MetaKey_,MetaValue_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddPostmeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostPostmetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iPostmeta model.Postmeta
	self.Bind(&iPostmeta)
	_int64, _error := model.PostPostmeta(&iPostmeta)
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

func PutPostmetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iPostmeta model.Postmeta
	self.Bind(&iPostmeta)
	_int64, _error := model.PutPostmeta(&iPostmeta)
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

func PutPostmetaViaMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	var iPostmeta model.Postmeta
	self.Bind(&iPostmeta)
	_int64, _error := model.PutPostmetaViaMetaId(MetaId_, &iPostmeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutPostmetaViaPostIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	PostId_ := self.Args("post_id").MustInt64()
	var iPostmeta model.Postmeta
	self.Bind(&iPostmeta)
	_int64, _error := model.PutPostmetaViaPostId(PostId_, &iPostmeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutPostmetaViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	var iPostmeta model.Postmeta
	self.Bind(&iPostmeta)
	_int64, _error := model.PutPostmetaViaMetaKey(MetaKey_, &iPostmeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutPostmetaViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	var iPostmeta model.Postmeta
	self.Bind(&iPostmeta)
	_int64, _error := model.PutPostmetaViaMetaValue(MetaValue_, &iPostmeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdatePostmetaViaMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	var iPostmeta model.Postmeta
	self.Bind(&iPostmeta)
	var iMap = helper.StructToMap(iPostmeta)
	_error := model.UpdatePostmetaViaMetaId(MetaId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdatePostmetaViaPostIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	PostId_ := self.Args("post_id").MustInt64()
	var iPostmeta model.Postmeta
	self.Bind(&iPostmeta)
	var iMap = helper.StructToMap(iPostmeta)
	_error := model.UpdatePostmetaViaPostId(PostId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdatePostmetaViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	var iPostmeta model.Postmeta
	self.Bind(&iPostmeta)
	var iMap = helper.StructToMap(iPostmeta)
	_error := model.UpdatePostmetaViaMetaKey(MetaKey_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdatePostmetaViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	var iPostmeta model.Postmeta
	self.Bind(&iPostmeta)
	var iMap = helper.StructToMap(iPostmeta)
	_error := model.UpdatePostmetaViaMetaValue(MetaValue_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeletePostmetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	_int64, _error := model.DeletePostmeta(MetaId_)
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

func DeletePostmetaViaMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	_error := model.DeletePostmetaViaMetaId(MetaId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeletePostmetaViaPostIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	PostId_ := self.Args("post_id").MustInt64()
	_error := model.DeletePostmetaViaPostId(PostId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeletePostmetaViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	_error := model.DeletePostmetaViaMetaKey(MetaKey_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeletePostmetaViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	_error := model.DeletePostmetaViaMetaValue(MetaValue_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
