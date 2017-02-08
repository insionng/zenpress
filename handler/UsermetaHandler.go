package handler

import (
	"github.com/insionng/zenpress/helper"
	"github.com/insionng/zenpress/model"
	"github.com/insionng/macross"
)

func GetUsermetasCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetUsermetasCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["usermetasCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetUsermetasCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsermetaCountViaUmetaIdHandler(self *macross.Context) error {
	UmetaId_ := self.Args("umeta_id").MustInt64()
	_int64 := model.GetUsermetaCountViaUmetaId(UmetaId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["usermetaCount"] = 0
	}
	m["usermetaCount"] = _int64
	return self.JSON(m)
}

func GetUsermetaCountViaUserIdHandler(self *macross.Context) error {
	UserId_ := self.Args("user_id").MustInt64()
	_int64 := model.GetUsermetaCountViaUserId(UserId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["usermetaCount"] = 0
	}
	m["usermetaCount"] = _int64
	return self.JSON(m)
}

func GetUsermetaCountViaMetaKeyHandler(self *macross.Context) error {
	MetaKey_ := self.Args("meta_key").String()
	_int64 := model.GetUsermetaCountViaMetaKey(MetaKey_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["usermetaCount"] = 0
	}
	m["usermetaCount"] = _int64
	return self.JSON(m)
}

func GetUsermetaCountViaMetaValueHandler(self *macross.Context) error {
	MetaValue_ := self.Args("meta_value").String()
	_int64 := model.GetUsermetaCountViaMetaValue(MetaValue_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["usermetaCount"] = 0
	}
	m["usermetaCount"] = _int64
	return self.JSON(m)
}

func GetUsermetasViaUmetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUmetaId := self.Args("umeta_id").MustInt64()
	if (offset > 0) && helper.IsHas(iUmetaId) {
		_Usermeta, _error := model.GetUsermetasViaUmetaId(offset, limit, iUmetaId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the GetUsermetasViaUmetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsermetasViaUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUserId := self.Args("user_id").MustInt64()
	if (offset > 0) && helper.IsHas(iUserId) {
		_Usermeta, _error := model.GetUsermetasViaUserId(offset, limit, iUserId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the GetUsermetasViaUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsermetasViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMetaKey := self.Args("meta_key").String()
	if (offset > 0) && helper.IsHas(iMetaKey) {
		_Usermeta, _error := model.GetUsermetasViaMetaKey(offset, limit, iMetaKey, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the GetUsermetasViaMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsermetasViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMetaValue := self.Args("meta_value").String()
	if (offset > 0) && helper.IsHas(iMetaValue) {
		_Usermeta, _error := model.GetUsermetasViaMetaValue(offset, limit, iMetaValue, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the GetUsermetasViaMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsermetasByUmetaIdAndUserIdAndMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUmetaId := self.Args("umeta_id").MustInt64()
	iUserId := self.Args("user_id").MustInt64()
	iMetaKey := self.Args("meta_key").String()

	if helper.IsHas(iUmetaId) {
		_Usermeta, _error := model.GetUsermetasByUmetaIdAndUserIdAndMetaKey(offset, limit, iUmetaId,iUserId,iMetaKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the GetUsermetasByUmetaIdAndUserIdAndMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsermetasByUmetaIdAndUserIdAndMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUmetaId := self.Args("umeta_id").MustInt64()
	iUserId := self.Args("user_id").MustInt64()
	iMetaValue := self.Args("meta_value").String()

	if helper.IsHas(iUmetaId) {
		_Usermeta, _error := model.GetUsermetasByUmetaIdAndUserIdAndMetaValue(offset, limit, iUmetaId,iUserId,iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the GetUsermetasByUmetaIdAndUserIdAndMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsermetasByUmetaIdAndMetaKeyAndMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUmetaId := self.Args("umeta_id").MustInt64()
	iMetaKey := self.Args("meta_key").String()
	iMetaValue := self.Args("meta_value").String()

	if helper.IsHas(iUmetaId) {
		_Usermeta, _error := model.GetUsermetasByUmetaIdAndMetaKeyAndMetaValue(offset, limit, iUmetaId,iMetaKey,iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the GetUsermetasByUmetaIdAndMetaKeyAndMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsermetasByUserIdAndMetaKeyAndMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserId := self.Args("user_id").MustInt64()
	iMetaKey := self.Args("meta_key").String()
	iMetaValue := self.Args("meta_value").String()

	if helper.IsHas(iUserId) {
		_Usermeta, _error := model.GetUsermetasByUserIdAndMetaKeyAndMetaValue(offset, limit, iUserId,iMetaKey,iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the GetUsermetasByUserIdAndMetaKeyAndMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsermetasByUmetaIdAndUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUmetaId := self.Args("umeta_id").MustInt64()
	iUserId := self.Args("user_id").MustInt64()

	if helper.IsHas(iUmetaId) {
		_Usermeta, _error := model.GetUsermetasByUmetaIdAndUserId(offset, limit, iUmetaId,iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the GetUsermetasByUmetaIdAndUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsermetasByUmetaIdAndMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUmetaId := self.Args("umeta_id").MustInt64()
	iMetaKey := self.Args("meta_key").String()

	if helper.IsHas(iUmetaId) {
		_Usermeta, _error := model.GetUsermetasByUmetaIdAndMetaKey(offset, limit, iUmetaId,iMetaKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the GetUsermetasByUmetaIdAndMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsermetasByUmetaIdAndMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUmetaId := self.Args("umeta_id").MustInt64()
	iMetaValue := self.Args("meta_value").String()

	if helper.IsHas(iUmetaId) {
		_Usermeta, _error := model.GetUsermetasByUmetaIdAndMetaValue(offset, limit, iUmetaId,iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the GetUsermetasByUmetaIdAndMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsermetasByUserIdAndMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserId := self.Args("user_id").MustInt64()
	iMetaKey := self.Args("meta_key").String()

	if helper.IsHas(iUserId) {
		_Usermeta, _error := model.GetUsermetasByUserIdAndMetaKey(offset, limit, iUserId,iMetaKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the GetUsermetasByUserIdAndMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsermetasByUserIdAndMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserId := self.Args("user_id").MustInt64()
	iMetaValue := self.Args("meta_value").String()

	if helper.IsHas(iUserId) {
		_Usermeta, _error := model.GetUsermetasByUserIdAndMetaValue(offset, limit, iUserId,iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the GetUsermetasByUserIdAndMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsermetasByMetaKeyAndMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iMetaKey := self.Args("meta_key").String()
	iMetaValue := self.Args("meta_value").String()

	if helper.IsHas(iMetaKey) {
		_Usermeta, _error := model.GetUsermetasByMetaKeyAndMetaValue(offset, limit, iMetaKey,iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the GetUsermetasByMetaKeyAndMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsermetasHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_Usermeta, _error := model.GetUsermetas(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the GetUsermetas' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsermetaViaUmetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUmetaId := self.Args("umeta_id").MustInt64()
	if helper.IsHas(iUmetaId) {
		_Usermeta := model.HasUsermetaViaUmetaId(iUmetaId)
		var m = map[string]interface{}{}
		m["usermeta"] = _Usermeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsermetaViaUmetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsermetaViaUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserId := self.Args("user_id").MustInt64()
	if helper.IsHas(iUserId) {
		_Usermeta := model.HasUsermetaViaUserId(iUserId)
		var m = map[string]interface{}{}
		m["usermeta"] = _Usermeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsermetaViaUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsermetaViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaKey := self.Args("meta_key").String()
	if helper.IsHas(iMetaKey) {
		_Usermeta := model.HasUsermetaViaMetaKey(iMetaKey)
		var m = map[string]interface{}{}
		m["usermeta"] = _Usermeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsermetaViaMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsermetaViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaValue := self.Args("meta_value").String()
	if helper.IsHas(iMetaValue) {
		_Usermeta := model.HasUsermetaViaMetaValue(iMetaValue)
		var m = map[string]interface{}{}
		m["usermeta"] = _Usermeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsermetaViaMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsermetaViaUmetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUmetaId := self.Args("umeta_id").MustInt64()
	if helper.IsHas(iUmetaId) {
		_Usermeta, _error := model.GetUsermetaViaUmetaId(iUmetaId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the GetUsermetaViaUmetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsermetaViaUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserId := self.Args("user_id").MustInt64()
	if helper.IsHas(iUserId) {
		_Usermeta, _error := model.GetUsermetaViaUserId(iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the GetUsermetaViaUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsermetaViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaKey := self.Args("meta_key").String()
	if helper.IsHas(iMetaKey) {
		_Usermeta, _error := model.GetUsermetaViaMetaKey(iMetaKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the GetUsermetaViaMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsermetaViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaValue := self.Args("meta_value").String()
	if helper.IsHas(iMetaValue) {
		_Usermeta, _error := model.GetUsermetaViaMetaValue(iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the GetUsermetaViaMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsermetaViaUmetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UmetaId_ := self.Args("umeta_id").MustInt64()
	if helper.IsHas(UmetaId_) {
		var iUsermeta model.Usermeta
		self.Bind(&iUsermeta)
		_Usermeta, _error := model.SetUsermetaViaUmetaId(UmetaId_, &iUsermeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the SetUsermetaViaUmetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsermetaViaUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserId_ := self.Args("user_id").MustInt64()
	if helper.IsHas(UserId_) {
		var iUsermeta model.Usermeta
		self.Bind(&iUsermeta)
		_Usermeta, _error := model.SetUsermetaViaUserId(UserId_, &iUsermeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the SetUsermetaViaUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsermetaViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	if helper.IsHas(MetaKey_) {
		var iUsermeta model.Usermeta
		self.Bind(&iUsermeta)
		_Usermeta, _error := model.SetUsermetaViaMetaKey(MetaKey_, &iUsermeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the SetUsermetaViaMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsermetaViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	if helper.IsHas(MetaValue_) {
		var iUsermeta model.Usermeta
		self.Bind(&iUsermeta)
		_Usermeta, _error := model.SetUsermetaViaMetaValue(MetaValue_, &iUsermeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the SetUsermetaViaMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddUsermetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UmetaId_ := self.Args("umeta_id").MustInt64()
	UserId_ := self.Args("user_id").MustInt64()
	MetaKey_ := self.Args("meta_key").String()
	MetaValue_ := self.Args("meta_value").String()

	if helper.IsHas(UmetaId_) {
		_error := model.AddUsermeta(UmetaId_,UserId_,MetaKey_,MetaValue_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddUsermeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostUsermetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iUsermeta model.Usermeta
	self.Bind(&iUsermeta)
	_int64, _error := model.PostUsermeta(&iUsermeta)
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

func PutUsermetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iUsermeta model.Usermeta
	self.Bind(&iUsermeta)
	_int64, _error := model.PutUsermeta(&iUsermeta)
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

func PutUsermetaViaUmetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UmetaId_ := self.Args("umeta_id").MustInt64()
	var iUsermeta model.Usermeta
	self.Bind(&iUsermeta)
	_int64, _error := model.PutUsermetaViaUmetaId(UmetaId_, &iUsermeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsermetaViaUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserId_ := self.Args("user_id").MustInt64()
	var iUsermeta model.Usermeta
	self.Bind(&iUsermeta)
	_int64, _error := model.PutUsermetaViaUserId(UserId_, &iUsermeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsermetaViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	var iUsermeta model.Usermeta
	self.Bind(&iUsermeta)
	_int64, _error := model.PutUsermetaViaMetaKey(MetaKey_, &iUsermeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsermetaViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	var iUsermeta model.Usermeta
	self.Bind(&iUsermeta)
	_int64, _error := model.PutUsermetaViaMetaValue(MetaValue_, &iUsermeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsermetaViaUmetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UmetaId_ := self.Args("umeta_id").MustInt64()
	var iUsermeta model.Usermeta
	self.Bind(&iUsermeta)
	var iMap = helper.StructToMap(iUsermeta)
	_error := model.UpdateUsermetaViaUmetaId(UmetaId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsermetaViaUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserId_ := self.Args("user_id").MustInt64()
	var iUsermeta model.Usermeta
	self.Bind(&iUsermeta)
	var iMap = helper.StructToMap(iUsermeta)
	_error := model.UpdateUsermetaViaUserId(UserId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsermetaViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	var iUsermeta model.Usermeta
	self.Bind(&iUsermeta)
	var iMap = helper.StructToMap(iUsermeta)
	_error := model.UpdateUsermetaViaMetaKey(MetaKey_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsermetaViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	var iUsermeta model.Usermeta
	self.Bind(&iUsermeta)
	var iMap = helper.StructToMap(iUsermeta)
	_error := model.UpdateUsermetaViaMetaValue(MetaValue_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsermetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UmetaId_ := self.Args("umeta_id").MustInt64()
	_int64, _error := model.DeleteUsermeta(UmetaId_)
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

func DeleteUsermetaViaUmetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UmetaId_ := self.Args("umeta_id").MustInt64()
	_error := model.DeleteUsermetaViaUmetaId(UmetaId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsermetaViaUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserId_ := self.Args("user_id").MustInt64()
	_error := model.DeleteUsermetaViaUserId(UserId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsermetaViaMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	_error := model.DeleteUsermetaViaMetaKey(MetaKey_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsermetaViaMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	_error := model.DeleteUsermetaViaMetaValue(MetaValue_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
