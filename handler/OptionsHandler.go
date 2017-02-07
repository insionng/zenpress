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

func GetOptionsCountByOptionIdHandler(self *macross.Context) error {
	OptionId_ := self.Args("option_id").MustInt64()
	_int64 := model.GetOptionsCountByOptionId(OptionId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["optionsCount"] = 0
	}
	m["optionsCount"] = _int64
	return self.JSON(m)
}

func GetOptionsCountByOptionNameHandler(self *macross.Context) error {
	OptionName_ := self.Args("option_name").String()
	_int64 := model.GetOptionsCountByOptionName(OptionName_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["optionsCount"] = 0
	}
	m["optionsCount"] = _int64
	return self.JSON(m)
}

func GetOptionsCountByOptionValueHandler(self *macross.Context) error {
	OptionValue_ := self.Args("option_value").String()
	_int64 := model.GetOptionsCountByOptionValue(OptionValue_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["optionsCount"] = 0
	}
	m["optionsCount"] = _int64
	return self.JSON(m)
}

func GetOptionsCountByAutoloadHandler(self *macross.Context) error {
	Autoload_ := self.Args("autoload").String()
	_int64 := model.GetOptionsCountByAutoload(Autoload_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["optionsCount"] = 0
	}
	m["optionsCount"] = _int64
	return self.JSON(m)
}

func GetOptionsesByOptionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iOptionId := self.Args("option_id").MustInt64()
	if (offset > 0) && helper.IsHas(iOptionId) {
		_Options, _error := model.GetOptionsesByOptionId(offset, limit, iOptionId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the GetOptionsesByOptionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetOptionsesByOptionNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iOptionName := self.Args("option_name").String()
	if (offset > 0) && helper.IsHas(iOptionName) {
		_Options, _error := model.GetOptionsesByOptionName(offset, limit, iOptionName, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the GetOptionsesByOptionName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetOptionsesByOptionValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iOptionValue := self.Args("option_value").String()
	if (offset > 0) && helper.IsHas(iOptionValue) {
		_Options, _error := model.GetOptionsesByOptionValue(offset, limit, iOptionValue, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the GetOptionsesByOptionValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetOptionsesByAutoloadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iAutoload := self.Args("autoload").String()
	if (offset > 0) && helper.IsHas(iAutoload) {
		_Options, _error := model.GetOptionsesByAutoload(offset, limit, iAutoload, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the GetOptionsesByAutoload's args."
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

func GetHasOptionsByOptionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iOptionId := self.Args("option_id").MustInt64()
	if helper.IsHas(iOptionId) {
		_Options := model.HasOptionsByOptionId(iOptionId)
		var m = map[string]interface{}{}
		m["options"] = _Options
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasOptionsByOptionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasOptionsByOptionNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iOptionName := self.Args("option_name").String()
	if helper.IsHas(iOptionName) {
		_Options := model.HasOptionsByOptionName(iOptionName)
		var m = map[string]interface{}{}
		m["options"] = _Options
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasOptionsByOptionName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasOptionsByOptionValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iOptionValue := self.Args("option_value").String()
	if helper.IsHas(iOptionValue) {
		_Options := model.HasOptionsByOptionValue(iOptionValue)
		var m = map[string]interface{}{}
		m["options"] = _Options
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasOptionsByOptionValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasOptionsByAutoloadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iAutoload := self.Args("autoload").String()
	if helper.IsHas(iAutoload) {
		_Options := model.HasOptionsByAutoload(iAutoload)
		var m = map[string]interface{}{}
		m["options"] = _Options
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasOptionsByAutoload's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetOptionsByOptionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iOptionId := self.Args("option_id").MustInt64()
	if helper.IsHas(iOptionId) {
		_Options, _error := model.GetOptionsByOptionId(iOptionId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the GetOptionsByOptionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetOptionsByOptionNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iOptionName := self.Args("option_name").String()
	if helper.IsHas(iOptionName) {
		_Options, _error := model.GetOptionsByOptionName(iOptionName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the GetOptionsByOptionName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetOptionsByOptionValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iOptionValue := self.Args("option_value").String()
	if helper.IsHas(iOptionValue) {
		_Options, _error := model.GetOptionsByOptionValue(iOptionValue)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the GetOptionsByOptionValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetOptionsByAutoloadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iAutoload := self.Args("autoload").String()
	if helper.IsHas(iAutoload) {
		_Options, _error := model.GetOptionsByAutoload(iAutoload)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the GetOptionsByAutoload's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetOptionsByOptionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	OptionId_ := self.Args("option_id").MustInt64()
	if helper.IsHas(OptionId_) {
		var iOptions model.Options
		self.Bind(&iOptions)
		_Options, _error := model.SetOptionsByOptionId(OptionId_, &iOptions)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the SetOptionsByOptionId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetOptionsByOptionNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	OptionName_ := self.Args("option_name").String()
	if helper.IsHas(OptionName_) {
		var iOptions model.Options
		self.Bind(&iOptions)
		_Options, _error := model.SetOptionsByOptionName(OptionName_, &iOptions)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the SetOptionsByOptionName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetOptionsByOptionValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	OptionValue_ := self.Args("option_value").String()
	if helper.IsHas(OptionValue_) {
		var iOptions model.Options
		self.Bind(&iOptions)
		_Options, _error := model.SetOptionsByOptionValue(OptionValue_, &iOptions)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the SetOptionsByOptionValue's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetOptionsByAutoloadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Autoload_ := self.Args("autoload").String()
	if helper.IsHas(Autoload_) {
		var iOptions model.Options
		self.Bind(&iOptions)
		_Options, _error := model.SetOptionsByAutoload(Autoload_, &iOptions)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Options)
	}
	herr.Message = "Can't get to the SetOptionsByAutoload's args."
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

func PutOptionsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iOptions model.Options
	self.Bind(&iOptions)
	_int64, _error := model.PutOptions(&iOptions)
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

func PutOptionsByOptionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	OptionId_ := self.Args("option_id").MustInt64()
	var iOptions model.Options
	self.Bind(&iOptions)
	_int64, _error := model.PutOptionsByOptionId(OptionId_, &iOptions)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutOptionsByOptionNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	OptionName_ := self.Args("option_name").String()
	var iOptions model.Options
	self.Bind(&iOptions)
	_int64, _error := model.PutOptionsByOptionName(OptionName_, &iOptions)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutOptionsByOptionValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	OptionValue_ := self.Args("option_value").String()
	var iOptions model.Options
	self.Bind(&iOptions)
	_int64, _error := model.PutOptionsByOptionValue(OptionValue_, &iOptions)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutOptionsByAutoloadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Autoload_ := self.Args("autoload").String()
	var iOptions model.Options
	self.Bind(&iOptions)
	_int64, _error := model.PutOptionsByAutoload(Autoload_, &iOptions)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateOptionsByOptionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	OptionId_ := self.Args("option_id").MustInt64()
	var iOptions model.Options
	self.Bind(&iOptions)
	var iMap = helper.StructToMap(iOptions)
	_error := model.UpdateOptionsByOptionId(OptionId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateOptionsByOptionNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	OptionName_ := self.Args("option_name").String()
	var iOptions model.Options
	self.Bind(&iOptions)
	var iMap = helper.StructToMap(iOptions)
	_error := model.UpdateOptionsByOptionName(OptionName_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateOptionsByOptionValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	OptionValue_ := self.Args("option_value").String()
	var iOptions model.Options
	self.Bind(&iOptions)
	var iMap = helper.StructToMap(iOptions)
	_error := model.UpdateOptionsByOptionValue(OptionValue_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateOptionsByAutoloadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Autoload_ := self.Args("autoload").String()
	var iOptions model.Options
	self.Bind(&iOptions)
	var iMap = helper.StructToMap(iOptions)
	_error := model.UpdateOptionsByAutoload(Autoload_, &iMap)
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

func DeleteOptionsByOptionIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	OptionId_ := self.Args("option_id").MustInt64()
	_error := model.DeleteOptionsByOptionId(OptionId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteOptionsByOptionNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	OptionName_ := self.Args("option_name").String()
	_error := model.DeleteOptionsByOptionName(OptionName_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteOptionsByOptionValueHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	OptionValue_ := self.Args("option_value").String()
	_error := model.DeleteOptionsByOptionValue(OptionValue_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteOptionsByAutoloadHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Autoload_ := self.Args("autoload").String()
	_error := model.DeleteOptionsByAutoload(Autoload_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
