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

func GetUsermetaCountByUmetaIdHandler(self *macross.Context) error {
	UmetaId_ := self.Args("umeta_id").MustInt64()
	_int64 := model.GetUsermetaCountByUmetaId(UmetaId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["usermetaCount"] = 0
	}
	m["usermetaCount"] = _int64
	return self.JSON(m)
}

func GetUsermetaCountByUserIdHandler(self *macross.Context) error {
	UserId_ := self.Args("user_id").MustInt64()
	_int64 := model.GetUsermetaCountByUserId(UserId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["usermetaCount"] = 0
	}
	m["usermetaCount"] = _int64
	return self.JSON(m)
}

func GetUsermetaCountByMetaKeyHandler(self *macross.Context) error {
	MetaKey_ := self.Args("meta_key").String()
	_int64 := model.GetUsermetaCountByMetaKey(MetaKey_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["usermetaCount"] = 0
	}
	m["usermetaCount"] = _int64
	return self.JSON(m)
}

func GetUsermetaCountByMetaValueHandler(self *macross.Context) error {
	MetaValue_ := self.Args("meta_value").String()
	_int64 := model.GetUsermetaCountByMetaValue(MetaValue_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["usermetaCount"] = 0
	}
	m["usermetaCount"] = _int64
	return self.JSON(m)
}

func GetUsermetasByUmetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUmetaId := self.Args("umeta_id").MustInt64()
	if (offset > 0) && helper.IsHas(iUmetaId) {
		_Usermeta, _error := model.GetUsermetasByUmetaId(offset, limit, iUmetaId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the GetUsermetasByUmetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsermetasByUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUserId := self.Args("user_id").MustInt64()
	if (offset > 0) && helper.IsHas(iUserId) {
		_Usermeta, _error := model.GetUsermetasByUserId(offset, limit, iUserId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the GetUsermetasByUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsermetasByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMetaKey := self.Args("meta_key").String()
	if (offset > 0) && helper.IsHas(iMetaKey) {
		_Usermeta, _error := model.GetUsermetasByMetaKey(offset, limit, iMetaKey, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the GetUsermetasByMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsermetasByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMetaValue := self.Args("meta_value").String()
	if (offset > 0) && helper.IsHas(iMetaValue) {
		_Usermeta, _error := model.GetUsermetasByMetaValue(offset, limit, iMetaValue, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the GetUsermetasByMetaValue's args."
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

func GetHasUsermetaByUmetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUmetaId := self.Args("umeta_id").MustInt64()
	if helper.IsHas(iUmetaId) {
		_Usermeta := model.HasUsermetaByUmetaId(iUmetaId)
		var m = map[string]interface{}{}
		m["usermeta"] = _Usermeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsermetaByUmetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsermetaByUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserId := self.Args("user_id").MustInt64()
	if helper.IsHas(iUserId) {
		_Usermeta := model.HasUsermetaByUserId(iUserId)
		var m = map[string]interface{}{}
		m["usermeta"] = _Usermeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsermetaByUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsermetaByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaKey := self.Args("meta_key").String()
	if helper.IsHas(iMetaKey) {
		_Usermeta := model.HasUsermetaByMetaKey(iMetaKey)
		var m = map[string]interface{}{}
		m["usermeta"] = _Usermeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsermetaByMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsermetaByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaValue := self.Args("meta_value").String()
	if helper.IsHas(iMetaValue) {
		_Usermeta := model.HasUsermetaByMetaValue(iMetaValue)
		var m = map[string]interface{}{}
		m["usermeta"] = _Usermeta
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsermetaByMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsermetaByUmetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUmetaId := self.Args("umeta_id").MustInt64()
	if helper.IsHas(iUmetaId) {
		_Usermeta, _error := model.GetUsermetaByUmetaId(iUmetaId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the GetUsermetaByUmetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsermetaByUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserId := self.Args("user_id").MustInt64()
	if helper.IsHas(iUserId) {
		_Usermeta, _error := model.GetUsermetaByUserId(iUserId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the GetUsermetaByUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsermetaByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaKey := self.Args("meta_key").String()
	if helper.IsHas(iMetaKey) {
		_Usermeta, _error := model.GetUsermetaByMetaKey(iMetaKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the GetUsermetaByMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsermetaByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMetaValue := self.Args("meta_value").String()
	if helper.IsHas(iMetaValue) {
		_Usermeta, _error := model.GetUsermetaByMetaValue(iMetaValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the GetUsermetaByMetaValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsermetaByUmetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UmetaId_ := self.Args("umeta_id").MustInt64()
	if helper.IsHas(UmetaId_) {
		var iUsermeta model.Usermeta
		self.Bind(&iUsermeta)
		_Usermeta, _error := model.SetUsermetaByUmetaId(UmetaId_, &iUsermeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the SetUsermetaByUmetaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsermetaByUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserId_ := self.Args("user_id").MustInt64()
	if helper.IsHas(UserId_) {
		var iUsermeta model.Usermeta
		self.Bind(&iUsermeta)
		_Usermeta, _error := model.SetUsermetaByUserId(UserId_, &iUsermeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the SetUsermetaByUserId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsermetaByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	if helper.IsHas(MetaKey_) {
		var iUsermeta model.Usermeta
		self.Bind(&iUsermeta)
		_Usermeta, _error := model.SetUsermetaByMetaKey(MetaKey_, &iUsermeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the SetUsermetaByMetaKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsermetaByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	if helper.IsHas(MetaValue_) {
		var iUsermeta model.Usermeta
		self.Bind(&iUsermeta)
		_Usermeta, _error := model.SetUsermetaByMetaValue(MetaValue_, &iUsermeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Usermeta)
	}
	herr.Message = "Can't get to the SetUsermetaByMetaValue's args."
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

func PutUsermetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iUsermeta model.Usermeta
	self.Bind(&iUsermeta)
	_int64, _error := model.PutUsermeta(&iUsermeta)
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

func PutUsermetaByUmetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UmetaId_ := self.Args("umeta_id").MustInt64()
	var iUsermeta model.Usermeta
	self.Bind(&iUsermeta)
	_int64, _error := model.PutUsermetaByUmetaId(UmetaId_, &iUsermeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsermetaByUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserId_ := self.Args("user_id").MustInt64()
	var iUsermeta model.Usermeta
	self.Bind(&iUsermeta)
	_int64, _error := model.PutUsermetaByUserId(UserId_, &iUsermeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsermetaByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	var iUsermeta model.Usermeta
	self.Bind(&iUsermeta)
	_int64, _error := model.PutUsermetaByMetaKey(MetaKey_, &iUsermeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsermetaByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	var iUsermeta model.Usermeta
	self.Bind(&iUsermeta)
	_int64, _error := model.PutUsermetaByMetaValue(MetaValue_, &iUsermeta)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsermetaByUmetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UmetaId_ := self.Args("umeta_id").MustInt64()
	var iUsermeta model.Usermeta
	self.Bind(&iUsermeta)
	var iMap = helper.StructToMap(iUsermeta)
	_error := model.UpdateUsermetaByUmetaId(UmetaId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsermetaByUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserId_ := self.Args("user_id").MustInt64()
	var iUsermeta model.Usermeta
	self.Bind(&iUsermeta)
	var iMap = helper.StructToMap(iUsermeta)
	_error := model.UpdateUsermetaByUserId(UserId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsermetaByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	var iUsermeta model.Usermeta
	self.Bind(&iUsermeta)
	var iMap = helper.StructToMap(iUsermeta)
	_error := model.UpdateUsermetaByMetaKey(MetaKey_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsermetaByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	var iUsermeta model.Usermeta
	self.Bind(&iUsermeta)
	var iMap = helper.StructToMap(iUsermeta)
	_error := model.UpdateUsermetaByMetaValue(MetaValue_, &iMap)
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

func DeleteUsermetaByUmetaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UmetaId_ := self.Args("umeta_id").MustInt64()
	_error := model.DeleteUsermetaByUmetaId(UmetaId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsermetaByUserIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserId_ := self.Args("user_id").MustInt64()
	_error := model.DeleteUsermetaByUserId(UserId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsermetaByMetaKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaKey_ := self.Args("meta_key").String()
	_error := model.DeleteUsermetaByMetaKey(MetaKey_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsermetaByMetaValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	MetaValue_ := self.Args("meta_value").String()
	_error := model.DeleteUsermetaByMetaValue(MetaValue_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
