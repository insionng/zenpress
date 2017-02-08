package handler

import (
	"github.com/insionng/zenpress/helper"
	"github.com/insionng/zenpress/model"
	"github.com/insionng/macross"
)

func GetRegistrationLogsCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetRegistrationLogsCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["registration_logsCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetRegistrationLogsCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRegistrationLogCountViaIdHandler(self *macross.Context) error {
	Id_ := self.Args("ID").MustInt64()
	_int64 := model.GetRegistrationLogCountViaId(Id_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["registration_logCount"] = 0
	}
	m["registration_logCount"] = _int64
	return self.JSON(m)
}

func GetRegistrationLogCountViaEmailHandler(self *macross.Context) error {
	Email_ := self.Args("email").String()
	_int64 := model.GetRegistrationLogCountViaEmail(Email_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["registration_logCount"] = 0
	}
	m["registration_logCount"] = _int64
	return self.JSON(m)
}

func GetRegistrationLogCountViaIpHandler(self *macross.Context) error {
	Ip_ := self.Args("IP").String()
	_int64 := model.GetRegistrationLogCountViaIp(Ip_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["registration_logCount"] = 0
	}
	m["registration_logCount"] = _int64
	return self.JSON(m)
}

func GetRegistrationLogCountViaBlogIdHandler(self *macross.Context) error {
	BlogId_ := self.Args("blog_id").MustInt64()
	_int64 := model.GetRegistrationLogCountViaBlogId(BlogId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["registration_logCount"] = 0
	}
	m["registration_logCount"] = _int64
	return self.JSON(m)
}

func GetRegistrationLogCountViaDateRegisteredHandler(self *macross.Context) error {
	DateRegistered_ := self.Args("date_registered").Time()
	_int64 := model.GetRegistrationLogCountViaDateRegistered(DateRegistered_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["registration_logCount"] = 0
	}
	m["registration_logCount"] = _int64
	return self.JSON(m)
}

func GetRegistrationLogsViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iId := self.Args("ID").MustInt64()
	if (offset > 0) && helper.IsHas(iId) {
		_RegistrationLog, _error := model.GetRegistrationLogsViaId(offset, limit, iId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the GetRegistrationLogsViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRegistrationLogsViaEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iEmail := self.Args("email").String()
	if (offset > 0) && helper.IsHas(iEmail) {
		_RegistrationLog, _error := model.GetRegistrationLogsViaEmail(offset, limit, iEmail, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the GetRegistrationLogsViaEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRegistrationLogsViaIpHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iIp := self.Args("IP").String()
	if (offset > 0) && helper.IsHas(iIp) {
		_RegistrationLog, _error := model.GetRegistrationLogsViaIp(offset, limit, iIp, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the GetRegistrationLogsViaIp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRegistrationLogsViaBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iBlogId := self.Args("blog_id").MustInt64()
	if (offset > 0) && helper.IsHas(iBlogId) {
		_RegistrationLog, _error := model.GetRegistrationLogsViaBlogId(offset, limit, iBlogId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the GetRegistrationLogsViaBlogId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRegistrationLogsViaDateRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDateRegistered := self.Args("date_registered").Time()
	if (offset > 0) && helper.IsHas(iDateRegistered) {
		_RegistrationLog, _error := model.GetRegistrationLogsViaDateRegistered(offset, limit, iDateRegistered, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the GetRegistrationLogsViaDateRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRegistrationLogsByIdAndEmailAndIpHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iEmail := self.Args("email").String()
	iIp := self.Args("ip").String()

	if helper.IsHas(iId) {
		_RegistrationLog, _error := model.GetRegistrationLogsByIdAndEmailAndIp(offset, limit, iId,iEmail,iIp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the GetRegistrationLogsByIdAndEmailAndIp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRegistrationLogsByIdAndEmailAndBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iEmail := self.Args("email").String()
	iBlogId := self.Args("blog_id").MustInt64()

	if helper.IsHas(iId) {
		_RegistrationLog, _error := model.GetRegistrationLogsByIdAndEmailAndBlogId(offset, limit, iId,iEmail,iBlogId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the GetRegistrationLogsByIdAndEmailAndBlogId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRegistrationLogsByIdAndEmailAndDateRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iEmail := self.Args("email").String()
	iDateRegistered := self.Args("date_registered").Time()

	if helper.IsHas(iId) {
		_RegistrationLog, _error := model.GetRegistrationLogsByIdAndEmailAndDateRegistered(offset, limit, iId,iEmail,iDateRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the GetRegistrationLogsByIdAndEmailAndDateRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRegistrationLogsByIdAndIpAndBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iIp := self.Args("ip").String()
	iBlogId := self.Args("blog_id").MustInt64()

	if helper.IsHas(iId) {
		_RegistrationLog, _error := model.GetRegistrationLogsByIdAndIpAndBlogId(offset, limit, iId,iIp,iBlogId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the GetRegistrationLogsByIdAndIpAndBlogId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRegistrationLogsByIdAndIpAndDateRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iIp := self.Args("ip").String()
	iDateRegistered := self.Args("date_registered").Time()

	if helper.IsHas(iId) {
		_RegistrationLog, _error := model.GetRegistrationLogsByIdAndIpAndDateRegistered(offset, limit, iId,iIp,iDateRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the GetRegistrationLogsByIdAndIpAndDateRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRegistrationLogsByIdAndBlogIdAndDateRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iBlogId := self.Args("blog_id").MustInt64()
	iDateRegistered := self.Args("date_registered").Time()

	if helper.IsHas(iId) {
		_RegistrationLog, _error := model.GetRegistrationLogsByIdAndBlogIdAndDateRegistered(offset, limit, iId,iBlogId,iDateRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the GetRegistrationLogsByIdAndBlogIdAndDateRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRegistrationLogsByEmailAndIpAndBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEmail := self.Args("email").String()
	iIp := self.Args("ip").String()
	iBlogId := self.Args("blog_id").MustInt64()

	if helper.IsHas(iEmail) {
		_RegistrationLog, _error := model.GetRegistrationLogsByEmailAndIpAndBlogId(offset, limit, iEmail,iIp,iBlogId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the GetRegistrationLogsByEmailAndIpAndBlogId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRegistrationLogsByEmailAndIpAndDateRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEmail := self.Args("email").String()
	iIp := self.Args("ip").String()
	iDateRegistered := self.Args("date_registered").Time()

	if helper.IsHas(iEmail) {
		_RegistrationLog, _error := model.GetRegistrationLogsByEmailAndIpAndDateRegistered(offset, limit, iEmail,iIp,iDateRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the GetRegistrationLogsByEmailAndIpAndDateRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRegistrationLogsByEmailAndBlogIdAndDateRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEmail := self.Args("email").String()
	iBlogId := self.Args("blog_id").MustInt64()
	iDateRegistered := self.Args("date_registered").Time()

	if helper.IsHas(iEmail) {
		_RegistrationLog, _error := model.GetRegistrationLogsByEmailAndBlogIdAndDateRegistered(offset, limit, iEmail,iBlogId,iDateRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the GetRegistrationLogsByEmailAndBlogIdAndDateRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRegistrationLogsByIpAndBlogIdAndDateRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iIp := self.Args("ip").String()
	iBlogId := self.Args("blog_id").MustInt64()
	iDateRegistered := self.Args("date_registered").Time()

	if helper.IsHas(iIp) {
		_RegistrationLog, _error := model.GetRegistrationLogsByIpAndBlogIdAndDateRegistered(offset, limit, iIp,iBlogId,iDateRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the GetRegistrationLogsByIpAndBlogIdAndDateRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRegistrationLogsByIdAndEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iEmail := self.Args("email").String()

	if helper.IsHas(iId) {
		_RegistrationLog, _error := model.GetRegistrationLogsByIdAndEmail(offset, limit, iId,iEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the GetRegistrationLogsByIdAndEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRegistrationLogsByIdAndIpHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iIp := self.Args("ip").String()

	if helper.IsHas(iId) {
		_RegistrationLog, _error := model.GetRegistrationLogsByIdAndIp(offset, limit, iId,iIp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the GetRegistrationLogsByIdAndIp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRegistrationLogsByIdAndBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iBlogId := self.Args("blog_id").MustInt64()

	if helper.IsHas(iId) {
		_RegistrationLog, _error := model.GetRegistrationLogsByIdAndBlogId(offset, limit, iId,iBlogId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the GetRegistrationLogsByIdAndBlogId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRegistrationLogsByIdAndDateRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iDateRegistered := self.Args("date_registered").Time()

	if helper.IsHas(iId) {
		_RegistrationLog, _error := model.GetRegistrationLogsByIdAndDateRegistered(offset, limit, iId,iDateRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the GetRegistrationLogsByIdAndDateRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRegistrationLogsByEmailAndIpHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEmail := self.Args("email").String()
	iIp := self.Args("ip").String()

	if helper.IsHas(iEmail) {
		_RegistrationLog, _error := model.GetRegistrationLogsByEmailAndIp(offset, limit, iEmail,iIp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the GetRegistrationLogsByEmailAndIp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRegistrationLogsByEmailAndBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEmail := self.Args("email").String()
	iBlogId := self.Args("blog_id").MustInt64()

	if helper.IsHas(iEmail) {
		_RegistrationLog, _error := model.GetRegistrationLogsByEmailAndBlogId(offset, limit, iEmail,iBlogId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the GetRegistrationLogsByEmailAndBlogId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRegistrationLogsByEmailAndDateRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iEmail := self.Args("email").String()
	iDateRegistered := self.Args("date_registered").Time()

	if helper.IsHas(iEmail) {
		_RegistrationLog, _error := model.GetRegistrationLogsByEmailAndDateRegistered(offset, limit, iEmail,iDateRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the GetRegistrationLogsByEmailAndDateRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRegistrationLogsByIpAndBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iIp := self.Args("ip").String()
	iBlogId := self.Args("blog_id").MustInt64()

	if helper.IsHas(iIp) {
		_RegistrationLog, _error := model.GetRegistrationLogsByIpAndBlogId(offset, limit, iIp,iBlogId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the GetRegistrationLogsByIpAndBlogId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRegistrationLogsByIpAndDateRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iIp := self.Args("ip").String()
	iDateRegistered := self.Args("date_registered").Time()

	if helper.IsHas(iIp) {
		_RegistrationLog, _error := model.GetRegistrationLogsByIpAndDateRegistered(offset, limit, iIp,iDateRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the GetRegistrationLogsByIpAndDateRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRegistrationLogsByBlogIdAndDateRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iBlogId := self.Args("blog_id").MustInt64()
	iDateRegistered := self.Args("date_registered").Time()

	if helper.IsHas(iBlogId) {
		_RegistrationLog, _error := model.GetRegistrationLogsByBlogIdAndDateRegistered(offset, limit, iBlogId,iDateRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the GetRegistrationLogsByBlogIdAndDateRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRegistrationLogsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_RegistrationLog, _error := model.GetRegistrationLogs(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the GetRegistrationLogs' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasRegistrationLogViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iId := self.Args("ID").MustInt64()
	if helper.IsHas(iId) {
		_RegistrationLog := model.HasRegistrationLogViaId(iId)
		var m = map[string]interface{}{}
		m["registration_log"] = _RegistrationLog
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasRegistrationLogViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasRegistrationLogViaEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEmail := self.Args("email").String()
	if helper.IsHas(iEmail) {
		_RegistrationLog := model.HasRegistrationLogViaEmail(iEmail)
		var m = map[string]interface{}{}
		m["registration_log"] = _RegistrationLog
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasRegistrationLogViaEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasRegistrationLogViaIpHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iIp := self.Args("IP").String()
	if helper.IsHas(iIp) {
		_RegistrationLog := model.HasRegistrationLogViaIp(iIp)
		var m = map[string]interface{}{}
		m["registration_log"] = _RegistrationLog
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasRegistrationLogViaIp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasRegistrationLogViaBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBlogId := self.Args("blog_id").MustInt64()
	if helper.IsHas(iBlogId) {
		_RegistrationLog := model.HasRegistrationLogViaBlogId(iBlogId)
		var m = map[string]interface{}{}
		m["registration_log"] = _RegistrationLog
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasRegistrationLogViaBlogId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasRegistrationLogViaDateRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDateRegistered := self.Args("date_registered").Time()
	if helper.IsHas(iDateRegistered) {
		_RegistrationLog := model.HasRegistrationLogViaDateRegistered(iDateRegistered)
		var m = map[string]interface{}{}
		m["registration_log"] = _RegistrationLog
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasRegistrationLogViaDateRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRegistrationLogViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iId := self.Args("ID").MustInt64()
	if helper.IsHas(iId) {
		_RegistrationLog, _error := model.GetRegistrationLogViaId(iId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the GetRegistrationLogViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRegistrationLogViaEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iEmail := self.Args("email").String()
	if helper.IsHas(iEmail) {
		_RegistrationLog, _error := model.GetRegistrationLogViaEmail(iEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the GetRegistrationLogViaEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRegistrationLogViaIpHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iIp := self.Args("IP").String()
	if helper.IsHas(iIp) {
		_RegistrationLog, _error := model.GetRegistrationLogViaIp(iIp)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the GetRegistrationLogViaIp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRegistrationLogViaBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iBlogId := self.Args("blog_id").MustInt64()
	if helper.IsHas(iBlogId) {
		_RegistrationLog, _error := model.GetRegistrationLogViaBlogId(iBlogId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the GetRegistrationLogViaBlogId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetRegistrationLogViaDateRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDateRegistered := self.Args("date_registered").Time()
	if helper.IsHas(iDateRegistered) {
		_RegistrationLog, _error := model.GetRegistrationLogViaDateRegistered(iDateRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the GetRegistrationLogViaDateRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetRegistrationLogViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("ID").MustInt64()
	if helper.IsHas(Id_) {
		var iRegistrationLog model.RegistrationLog
		self.Bind(&iRegistrationLog)
		_RegistrationLog, _error := model.SetRegistrationLogViaId(Id_, &iRegistrationLog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the SetRegistrationLogViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetRegistrationLogViaEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Email_ := self.Args("email").String()
	if helper.IsHas(Email_) {
		var iRegistrationLog model.RegistrationLog
		self.Bind(&iRegistrationLog)
		_RegistrationLog, _error := model.SetRegistrationLogViaEmail(Email_, &iRegistrationLog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the SetRegistrationLogViaEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetRegistrationLogViaIpHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Ip_ := self.Args("IP").String()
	if helper.IsHas(Ip_) {
		var iRegistrationLog model.RegistrationLog
		self.Bind(&iRegistrationLog)
		_RegistrationLog, _error := model.SetRegistrationLogViaIp(Ip_, &iRegistrationLog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the SetRegistrationLogViaIp's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetRegistrationLogViaBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BlogId_ := self.Args("blog_id").MustInt64()
	if helper.IsHas(BlogId_) {
		var iRegistrationLog model.RegistrationLog
		self.Bind(&iRegistrationLog)
		_RegistrationLog, _error := model.SetRegistrationLogViaBlogId(BlogId_, &iRegistrationLog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the SetRegistrationLogViaBlogId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetRegistrationLogViaDateRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DateRegistered_ := self.Args("date_registered").Time()
	if helper.IsHas(DateRegistered_) {
		var iRegistrationLog model.RegistrationLog
		self.Bind(&iRegistrationLog)
		_RegistrationLog, _error := model.SetRegistrationLogViaDateRegistered(DateRegistered_, &iRegistrationLog)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_RegistrationLog)
	}
	herr.Message = "Can't get to the SetRegistrationLogViaDateRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddRegistrationLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("ID").MustInt64()
	Email_ := self.Args("email").String()
	Ip_ := self.Args("IP").String()
	BlogId_ := self.Args("blog_id").MustInt64()
	DateRegistered_ := self.Args("date_registered").Time()

	if helper.IsHas(Id_) {
		_error := model.AddRegistrationLog(Id_,Email_,Ip_,BlogId_,DateRegistered_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddRegistrationLog's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostRegistrationLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iRegistrationLog model.RegistrationLog
	self.Bind(&iRegistrationLog)
	_int64, _error := model.PostRegistrationLog(&iRegistrationLog)
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

func PutRegistrationLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iRegistrationLog model.RegistrationLog
	self.Bind(&iRegistrationLog)
	_int64, _error := model.PutRegistrationLog(&iRegistrationLog)
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

func PutRegistrationLogViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("ID").MustInt64()
	var iRegistrationLog model.RegistrationLog
	self.Bind(&iRegistrationLog)
	_int64, _error := model.PutRegistrationLogViaId(Id_, &iRegistrationLog)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutRegistrationLogViaEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Email_ := self.Args("email").String()
	var iRegistrationLog model.RegistrationLog
	self.Bind(&iRegistrationLog)
	_int64, _error := model.PutRegistrationLogViaEmail(Email_, &iRegistrationLog)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutRegistrationLogViaIpHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Ip_ := self.Args("IP").String()
	var iRegistrationLog model.RegistrationLog
	self.Bind(&iRegistrationLog)
	_int64, _error := model.PutRegistrationLogViaIp(Ip_, &iRegistrationLog)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutRegistrationLogViaBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BlogId_ := self.Args("blog_id").MustInt64()
	var iRegistrationLog model.RegistrationLog
	self.Bind(&iRegistrationLog)
	_int64, _error := model.PutRegistrationLogViaBlogId(BlogId_, &iRegistrationLog)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutRegistrationLogViaDateRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DateRegistered_ := self.Args("date_registered").Time()
	var iRegistrationLog model.RegistrationLog
	self.Bind(&iRegistrationLog)
	_int64, _error := model.PutRegistrationLogViaDateRegistered(DateRegistered_, &iRegistrationLog)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateRegistrationLogViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("ID").MustInt64()
	var iRegistrationLog model.RegistrationLog
	self.Bind(&iRegistrationLog)
	var iMap = helper.StructToMap(iRegistrationLog)
	_error := model.UpdateRegistrationLogViaId(Id_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateRegistrationLogViaEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Email_ := self.Args("email").String()
	var iRegistrationLog model.RegistrationLog
	self.Bind(&iRegistrationLog)
	var iMap = helper.StructToMap(iRegistrationLog)
	_error := model.UpdateRegistrationLogViaEmail(Email_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateRegistrationLogViaIpHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Ip_ := self.Args("IP").String()
	var iRegistrationLog model.RegistrationLog
	self.Bind(&iRegistrationLog)
	var iMap = helper.StructToMap(iRegistrationLog)
	_error := model.UpdateRegistrationLogViaIp(Ip_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateRegistrationLogViaBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BlogId_ := self.Args("blog_id").MustInt64()
	var iRegistrationLog model.RegistrationLog
	self.Bind(&iRegistrationLog)
	var iMap = helper.StructToMap(iRegistrationLog)
	_error := model.UpdateRegistrationLogViaBlogId(BlogId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateRegistrationLogViaDateRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DateRegistered_ := self.Args("date_registered").Time()
	var iRegistrationLog model.RegistrationLog
	self.Bind(&iRegistrationLog)
	var iMap = helper.StructToMap(iRegistrationLog)
	_error := model.UpdateRegistrationLogViaDateRegistered(DateRegistered_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteRegistrationLogHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("ID").MustInt64()
	_int64, _error := model.DeleteRegistrationLog(Id_)
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

func DeleteRegistrationLogViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("ID").MustInt64()
	_error := model.DeleteRegistrationLogViaId(Id_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteRegistrationLogViaEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Email_ := self.Args("email").String()
	_error := model.DeleteRegistrationLogViaEmail(Email_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteRegistrationLogViaIpHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Ip_ := self.Args("IP").String()
	_error := model.DeleteRegistrationLogViaIp(Ip_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteRegistrationLogViaBlogIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	BlogId_ := self.Args("blog_id").MustInt64()
	_error := model.DeleteRegistrationLogViaBlogId(BlogId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteRegistrationLogViaDateRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DateRegistered_ := self.Args("date_registered").Time()
	_error := model.DeleteRegistrationLogViaDateRegistered(DateRegistered_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
