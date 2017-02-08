package handler

import (
	"github.com/insionng/zenpress/helper"
	"github.com/insionng/zenpress/model"
	"github.com/insionng/macross"
)

func GetOptionsesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetOptionsesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["optionssCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetOptionsesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetOptionsCountViaOptionIdHandler(self *macross.Context) error {
	OptionId_ := self.Args("option_id").MustInt64()
	_int64 := model.GetOptionsCountViaOptionId(OptionId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["optionsCount"] = 0
	}
	m["optionsCount"] = _int64
	return self.JSON(m)
}

func GetOptionsCountViaOptionNameHandler(self *macross.Context) error {
	OptionName_ := self.Args("option_name").String()
	_int64 := model.GetOptionsCountViaOptionName(OptionName_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["optionsCount"] = 0
	}
	m["optionsCount"] = _int64
	return self.JSON(m)
}

func GetOptionsCountViaOptionValueHandler(self *macross.Context) error {
	OptionValue_ := self.Args("option_value").String()
	_int64 := model.GetOptionsCountViaOptionValue(OptionValue_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["optionsCount"] = 0
	}
	m["optionsCount"] = _int64
	return self.JSON(m)
}

func GetOptionsCountViaAutoloadHandler(self *macross.Context) error {
	Autoload_ := self.Args("autoload").String()
	_int64 := model.GetOptionsCountViaAutoload(Autoload_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["optionsCount"] = 0
	}
	m["optionsCount"] = _int64
	return self.JSON(m)
}

func GetOptionsesViaOptionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iOptionId := self.Args("option_id").MustInt64()
	if (offset > 0) && helper.IsHas(iOptionId) {
		_Options, _error := model.GetOptionsesViaOptionId(offset, limit, iOptionId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the GetOptionsesViaOptionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetOptionsesViaOptionNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iOptionName := self.Args("option_name").String()
	if (offset > 0) && helper.IsHas(iOptionName) {
		_Options, _error := model.GetOptionsesViaOptionName(offset, limit, iOptionName, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the GetOptionsesViaOptionName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetOptionsesViaOptionValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iOptionValue := self.Args("option_value").String()
	if (offset > 0) && helper.IsHas(iOptionValue) {
		_Options, _error := model.GetOptionsesViaOptionValue(offset, limit, iOptionValue, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the GetOptionsesViaOptionValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetOptionsesViaAutoloadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iAutoload := self.Args("autoload").String()
	if (offset > 0) && helper.IsHas(iAutoload) {
		_Options, _error := model.GetOptionsesViaAutoload(offset, limit, iAutoload, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the GetOptionsesViaAutoload's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetOptionsesByOptionIdAndOptionNameAndOptionValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iOptionId := self.Args("option_id").MustInt64()
	iOptionName := self.Args("option_name").String()
	iOptionValue := self.Args("option_value").String()

	if helper.IsHas(iOptionId) {
		_Options, _error := model.GetOptionsesByOptionIdAndOptionNameAndOptionValue(offset, limit, iOptionId,iOptionName,iOptionValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the GetOptionsesByOptionIdAndOptionNameAndOptionValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetOptionsesByOptionIdAndOptionNameAndAutoloadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iOptionId := self.Args("option_id").MustInt64()
	iOptionName := self.Args("option_name").String()
	iAutoload := self.Args("autoload").String()

	if helper.IsHas(iOptionId) {
		_Options, _error := model.GetOptionsesByOptionIdAndOptionNameAndAutoload(offset, limit, iOptionId,iOptionName,iAutoload)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the GetOptionsesByOptionIdAndOptionNameAndAutoload's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetOptionsesByOptionIdAndOptionValueAndAutoloadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iOptionId := self.Args("option_id").MustInt64()
	iOptionValue := self.Args("option_value").String()
	iAutoload := self.Args("autoload").String()

	if helper.IsHas(iOptionId) {
		_Options, _error := model.GetOptionsesByOptionIdAndOptionValueAndAutoload(offset, limit, iOptionId,iOptionValue,iAutoload)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the GetOptionsesByOptionIdAndOptionValueAndAutoload's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetOptionsesByOptionNameAndOptionValueAndAutoloadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iOptionName := self.Args("option_name").String()
	iOptionValue := self.Args("option_value").String()
	iAutoload := self.Args("autoload").String()

	if helper.IsHas(iOptionName) {
		_Options, _error := model.GetOptionsesByOptionNameAndOptionValueAndAutoload(offset, limit, iOptionName,iOptionValue,iAutoload)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the GetOptionsesByOptionNameAndOptionValueAndAutoload's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetOptionsesByOptionIdAndOptionNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iOptionId := self.Args("option_id").MustInt64()
	iOptionName := self.Args("option_name").String()

	if helper.IsHas(iOptionId) {
		_Options, _error := model.GetOptionsesByOptionIdAndOptionName(offset, limit, iOptionId,iOptionName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the GetOptionsesByOptionIdAndOptionName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetOptionsesByOptionIdAndOptionValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iOptionId := self.Args("option_id").MustInt64()
	iOptionValue := self.Args("option_value").String()

	if helper.IsHas(iOptionId) {
		_Options, _error := model.GetOptionsesByOptionIdAndOptionValue(offset, limit, iOptionId,iOptionValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the GetOptionsesByOptionIdAndOptionValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetOptionsesByOptionIdAndAutoloadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iOptionId := self.Args("option_id").MustInt64()
	iAutoload := self.Args("autoload").String()

	if helper.IsHas(iOptionId) {
		_Options, _error := model.GetOptionsesByOptionIdAndAutoload(offset, limit, iOptionId,iAutoload)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the GetOptionsesByOptionIdAndAutoload's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetOptionsesByOptionNameAndOptionValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iOptionName := self.Args("option_name").String()
	iOptionValue := self.Args("option_value").String()

	if helper.IsHas(iOptionName) {
		_Options, _error := model.GetOptionsesByOptionNameAndOptionValue(offset, limit, iOptionName,iOptionValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the GetOptionsesByOptionNameAndOptionValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetOptionsesByOptionNameAndAutoloadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iOptionName := self.Args("option_name").String()
	iAutoload := self.Args("autoload").String()

	if helper.IsHas(iOptionName) {
		_Options, _error := model.GetOptionsesByOptionNameAndAutoload(offset, limit, iOptionName,iAutoload)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the GetOptionsesByOptionNameAndAutoload's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetOptionsesByOptionValueAndAutoloadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iOptionValue := self.Args("option_value").String()
	iAutoload := self.Args("autoload").String()

	if helper.IsHas(iOptionValue) {
		_Options, _error := model.GetOptionsesByOptionValueAndAutoload(offset, limit, iOptionValue,iAutoload)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the GetOptionsesByOptionValueAndAutoload's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetOptionsesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_Options, _error := model.GetOptionses(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the GetOptionses' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasOptionsViaOptionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iOptionId := self.Args("option_id").MustInt64()
	if helper.IsHas(iOptionId) {
		_Options := model.HasOptionsViaOptionId(iOptionId)
		var m = map[string]interface{}{}
		m["options"] = _Options
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasOptionsViaOptionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasOptionsViaOptionNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iOptionName := self.Args("option_name").String()
	if helper.IsHas(iOptionName) {
		_Options := model.HasOptionsViaOptionName(iOptionName)
		var m = map[string]interface{}{}
		m["options"] = _Options
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasOptionsViaOptionName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasOptionsViaOptionValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iOptionValue := self.Args("option_value").String()
	if helper.IsHas(iOptionValue) {
		_Options := model.HasOptionsViaOptionValue(iOptionValue)
		var m = map[string]interface{}{}
		m["options"] = _Options
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasOptionsViaOptionValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasOptionsViaAutoloadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iAutoload := self.Args("autoload").String()
	if helper.IsHas(iAutoload) {
		_Options := model.HasOptionsViaAutoload(iAutoload)
		var m = map[string]interface{}{}
		m["options"] = _Options
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasOptionsViaAutoload's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetOptionsViaOptionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iOptionId := self.Args("option_id").MustInt64()
	if helper.IsHas(iOptionId) {
		_Options, _error := model.GetOptionsViaOptionId(iOptionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the GetOptionsViaOptionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetOptionsViaOptionNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iOptionName := self.Args("option_name").String()
	if helper.IsHas(iOptionName) {
		_Options, _error := model.GetOptionsViaOptionName(iOptionName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the GetOptionsViaOptionName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetOptionsViaOptionValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iOptionValue := self.Args("option_value").String()
	if helper.IsHas(iOptionValue) {
		_Options, _error := model.GetOptionsViaOptionValue(iOptionValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the GetOptionsViaOptionValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetOptionsViaAutoloadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iAutoload := self.Args("autoload").String()
	if helper.IsHas(iAutoload) {
		_Options, _error := model.GetOptionsViaAutoload(iAutoload)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the GetOptionsViaAutoload's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetOptionsViaOptionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	OptionId_ := self.Args("option_id").MustInt64()
	if helper.IsHas(OptionId_) {
		var iOptions model.Options
		self.Bind(&iOptions)
		_Options, _error := model.SetOptionsViaOptionId(OptionId_, &iOptions)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the SetOptionsViaOptionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetOptionsViaOptionNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	OptionName_ := self.Args("option_name").String()
	if helper.IsHas(OptionName_) {
		var iOptions model.Options
		self.Bind(&iOptions)
		_Options, _error := model.SetOptionsViaOptionName(OptionName_, &iOptions)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the SetOptionsViaOptionName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetOptionsViaOptionValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	OptionValue_ := self.Args("option_value").String()
	if helper.IsHas(OptionValue_) {
		var iOptions model.Options
		self.Bind(&iOptions)
		_Options, _error := model.SetOptionsViaOptionValue(OptionValue_, &iOptions)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the SetOptionsViaOptionValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetOptionsViaAutoloadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Autoload_ := self.Args("autoload").String()
	if helper.IsHas(Autoload_) {
		var iOptions model.Options
		self.Bind(&iOptions)
		_Options, _error := model.SetOptionsViaAutoload(Autoload_, &iOptions)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the SetOptionsViaAutoload's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddOptionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	OptionId_ := self.Args("option_id").MustInt64()
	OptionName_ := self.Args("option_name").String()
	OptionValue_ := self.Args("option_value").String()
	Autoload_ := self.Args("autoload").String()

	if helper.IsHas(OptionId_) {
		_error := model.AddOptions(OptionId_,OptionName_,OptionValue_,Autoload_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddOptions's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostOptionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iOptions model.Options
	self.Bind(&iOptions)
	_int64, _error := model.PostOptions(&iOptions)
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

func PutOptionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iOptions model.Options
	self.Bind(&iOptions)
	_int64, _error := model.PutOptions(&iOptions)
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

func PutOptionsViaOptionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	OptionId_ := self.Args("option_id").MustInt64()
	var iOptions model.Options
	self.Bind(&iOptions)
	_int64, _error := model.PutOptionsViaOptionId(OptionId_, &iOptions)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutOptionsViaOptionNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	OptionName_ := self.Args("option_name").String()
	var iOptions model.Options
	self.Bind(&iOptions)
	_int64, _error := model.PutOptionsViaOptionName(OptionName_, &iOptions)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutOptionsViaOptionValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	OptionValue_ := self.Args("option_value").String()
	var iOptions model.Options
	self.Bind(&iOptions)
	_int64, _error := model.PutOptionsViaOptionValue(OptionValue_, &iOptions)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutOptionsViaAutoloadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Autoload_ := self.Args("autoload").String()
	var iOptions model.Options
	self.Bind(&iOptions)
	_int64, _error := model.PutOptionsViaAutoload(Autoload_, &iOptions)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateOptionsViaOptionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	OptionId_ := self.Args("option_id").MustInt64()
	var iOptions model.Options
	self.Bind(&iOptions)
	var iMap = helper.StructToMap(iOptions)
	_error := model.UpdateOptionsViaOptionId(OptionId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateOptionsViaOptionNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	OptionName_ := self.Args("option_name").String()
	var iOptions model.Options
	self.Bind(&iOptions)
	var iMap = helper.StructToMap(iOptions)
	_error := model.UpdateOptionsViaOptionName(OptionName_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateOptionsViaOptionValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	OptionValue_ := self.Args("option_value").String()
	var iOptions model.Options
	self.Bind(&iOptions)
	var iMap = helper.StructToMap(iOptions)
	_error := model.UpdateOptionsViaOptionValue(OptionValue_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateOptionsViaAutoloadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Autoload_ := self.Args("autoload").String()
	var iOptions model.Options
	self.Bind(&iOptions)
	var iMap = helper.StructToMap(iOptions)
	_error := model.UpdateOptionsViaAutoload(Autoload_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteOptionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	OptionId_ := self.Args("option_id").MustInt64()
	_int64, _error := model.DeleteOptions(OptionId_)
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

func DeleteOptionsViaOptionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	OptionId_ := self.Args("option_id").MustInt64()
	_error := model.DeleteOptionsViaOptionId(OptionId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteOptionsViaOptionNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	OptionName_ := self.Args("option_name").String()
	_error := model.DeleteOptionsViaOptionName(OptionName_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteOptionsViaOptionValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	OptionValue_ := self.Args("option_value").String()
	_error := model.DeleteOptionsViaOptionValue(OptionValue_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteOptionsViaAutoloadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Autoload_ := self.Args("autoload").String()
	_error := model.DeleteOptionsViaAutoload(Autoload_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
