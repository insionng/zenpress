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

func GetPostmetaCountByMetaIdHandler(self *macross.Context) error {
	MetaId_ := self.Args("meta_id").MustInt64()
	_int64 := model.GetPostmetaCountByMetaId(MetaId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["postmetaCount"] = 0
	}
	m["postmetaCount"] = _int64
	return self.JSON(m)
}

func GetPostmetaCountByPostIdHandler(self *macross.Context) error {
	PostId_ := self.Args("post_id").MustInt64()
	_int64 := model.GetPostmetaCountByPostId(PostId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["postmetaCount"] = 0
	}
	m["postmetaCount"] = _int64
	return self.JSON(m)
}

func GetPostmetaCountByMetaKeyHandler(self *macross.Context) error {
	MetaKey_ := self.Args("meta_key").String()
	_int64 := model.GetPostmetaCountByMetaKey(MetaKey_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["postmetaCount"] = 0
	}
	m["postmetaCount"] = _int64
	return self.JSON(m)
}

func GetPostmetaCountByMetaValueHandler(self *macross.Context) error {
	MetaValue_ := self.Args("meta_value").String()
	_int64 := model.GetPostmetaCountByMetaValue(MetaValue_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["postmetaCount"] = 0
	}
	m["postmetaCount"] = _int64
	return self.JSON(m)
}

func GetPostmetasByMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMetaId := self.Args("meta_id").MustInt64()
	if (offset > 0) && helper.IsHas(iMetaId) {
		_Postmeta, _error := model.GetPostmetasByMetaId(offset, limit, iMetaId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the GetPostmetasByMetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetPostmetasByPostIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iPostId := self.Args("post_id").MustInt64()
	if (offset > 0) && helper.IsHas(iPostId) {
		_Postmeta, _error := model.GetPostmetasByPostId(offset, limit, iPostId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the GetPostmetasByPostId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetPostmetasByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMetaKey := self.Args("meta_key").String()
	if (offset > 0) && helper.IsHas(iMetaKey) {
		_Postmeta, _error := model.GetPostmetasByMetaKey(offset, limit, iMetaKey, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the GetPostmetasByMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetPostmetasByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMetaValue := self.Args("meta_value").String()
	if (offset > 0) && helper.IsHas(iMetaValue) {
		_Postmeta, _error := model.GetPostmetasByMetaValue(offset, limit, iMetaValue, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the GetPostmetasByMetaValue's args."
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

func GetHasPostmetaByMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaId := self.Args("meta_id").MustInt64()
	if helper.IsHas(iMetaId) {
		_Postmeta := model.HasPostmetaByMetaId(iMetaId)
		var m = map[string]interface{}{}
		m["postmeta"] = _Postmeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasPostmetaByMetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasPostmetaByPostIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iPostId := self.Args("post_id").MustInt64()
	if helper.IsHas(iPostId) {
		_Postmeta := model.HasPostmetaByPostId(iPostId)
		var m = map[string]interface{}{}
		m["postmeta"] = _Postmeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasPostmetaByPostId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasPostmetaByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaKey := self.Args("meta_key").String()
	if helper.IsHas(iMetaKey) {
		_Postmeta := model.HasPostmetaByMetaKey(iMetaKey)
		var m = map[string]interface{}{}
		m["postmeta"] = _Postmeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasPostmetaByMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasPostmetaByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaValue := self.Args("meta_value").String()
	if helper.IsHas(iMetaValue) {
		_Postmeta := model.HasPostmetaByMetaValue(iMetaValue)
		var m = map[string]interface{}{}
		m["postmeta"] = _Postmeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasPostmetaByMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetPostmetaByMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaId := self.Args("meta_id").MustInt64()
	if helper.IsHas(iMetaId) {
		_Postmeta, _error := model.GetPostmetaByMetaId(iMetaId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the GetPostmetaByMetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetPostmetaByPostIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iPostId := self.Args("post_id").MustInt64()
	if helper.IsHas(iPostId) {
		_Postmeta, _error := model.GetPostmetaByPostId(iPostId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the GetPostmetaByPostId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetPostmetaByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaKey := self.Args("meta_key").String()
	if helper.IsHas(iMetaKey) {
		_Postmeta, _error := model.GetPostmetaByMetaKey(iMetaKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the GetPostmetaByMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetPostmetaByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaValue := self.Args("meta_value").String()
	if helper.IsHas(iMetaValue) {
		_Postmeta, _error := model.GetPostmetaByMetaValue(iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the GetPostmetaByMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetPostmetaByMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	if helper.IsHas(MetaId_) {
		var iPostmeta model.Postmeta
		self.Bind(&iPostmeta)
		_Postmeta, _error := model.SetPostmetaByMetaId(MetaId_, &iPostmeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the SetPostmetaByMetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetPostmetaByPostIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	PostId_ := self.Args("post_id").MustInt64()
	if helper.IsHas(PostId_) {
		var iPostmeta model.Postmeta
		self.Bind(&iPostmeta)
		_Postmeta, _error := model.SetPostmetaByPostId(PostId_, &iPostmeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the SetPostmetaByPostId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetPostmetaByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	if helper.IsHas(MetaKey_) {
		var iPostmeta model.Postmeta
		self.Bind(&iPostmeta)
		_Postmeta, _error := model.SetPostmetaByMetaKey(MetaKey_, &iPostmeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the SetPostmetaByMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetPostmetaByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	if helper.IsHas(MetaValue_) {
		var iPostmeta model.Postmeta
		self.Bind(&iPostmeta)
		_Postmeta, _error := model.SetPostmetaByMetaValue(MetaValue_, &iPostmeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Postmeta)
	}
	herr.Message = "Can't get to the SetPostmetaByMetaValue's args."
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

func PutPostmetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iPostmeta model.Postmeta
	self.Bind(&iPostmeta)
	_int64, _error := model.PutPostmeta(&iPostmeta)
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

func PutPostmetaByMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	var iPostmeta model.Postmeta
	self.Bind(&iPostmeta)
	_int64, _error := model.PutPostmetaByMetaId(MetaId_, &iPostmeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutPostmetaByPostIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	PostId_ := self.Args("post_id").MustInt64()
	var iPostmeta model.Postmeta
	self.Bind(&iPostmeta)
	_int64, _error := model.PutPostmetaByPostId(PostId_, &iPostmeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutPostmetaByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	var iPostmeta model.Postmeta
	self.Bind(&iPostmeta)
	_int64, _error := model.PutPostmetaByMetaKey(MetaKey_, &iPostmeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutPostmetaByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	var iPostmeta model.Postmeta
	self.Bind(&iPostmeta)
	_int64, _error := model.PutPostmetaByMetaValue(MetaValue_, &iPostmeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdatePostmetaByMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	var iPostmeta model.Postmeta
	self.Bind(&iPostmeta)
	var iMap = helper.StructToMap(iPostmeta)
	_error := model.UpdatePostmetaByMetaId(MetaId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdatePostmetaByPostIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	PostId_ := self.Args("post_id").MustInt64()
	var iPostmeta model.Postmeta
	self.Bind(&iPostmeta)
	var iMap = helper.StructToMap(iPostmeta)
	_error := model.UpdatePostmetaByPostId(PostId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdatePostmetaByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	var iPostmeta model.Postmeta
	self.Bind(&iPostmeta)
	var iMap = helper.StructToMap(iPostmeta)
	_error := model.UpdatePostmetaByMetaKey(MetaKey_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdatePostmetaByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	var iPostmeta model.Postmeta
	self.Bind(&iPostmeta)
	var iMap = helper.StructToMap(iPostmeta)
	_error := model.UpdatePostmetaByMetaValue(MetaValue_, &iMap)
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

func DeletePostmetaByMetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaId_ := self.Args("meta_id").MustInt64()
	_error := model.DeletePostmetaByMetaId(MetaId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeletePostmetaByPostIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	PostId_ := self.Args("post_id").MustInt64()
	_error := model.DeletePostmetaByPostId(PostId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeletePostmetaByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	_error := model.DeletePostmetaByMetaKey(MetaKey_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeletePostmetaByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	_error := model.DeletePostmetaByMetaValue(MetaValue_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
