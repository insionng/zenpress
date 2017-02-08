package handler

import (
	"github.com/insionng/zenpress/helper"
	"github.com/insionng/zenpress/model"
	"github.com/insionng/macross"
)

func GetSignupsesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetSignupsesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["signupssCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetSignupsesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsCountViaSignupIdHandler(self *macross.Context) error {
	SignupId_ := self.Args("signup_id").MustInt64()
	_int64 := model.GetSignupsCountViaSignupId(SignupId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["signupsCount"] = 0
	}
	m["signupsCount"] = _int64
	return self.JSON(m)
}

func GetSignupsCountViaDomainHandler(self *macross.Context) error {
	Domain_ := self.Args("domain").String()
	_int64 := model.GetSignupsCountViaDomain(Domain_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["signupsCount"] = 0
	}
	m["signupsCount"] = _int64
	return self.JSON(m)
}

func GetSignupsCountViaPathHandler(self *macross.Context) error {
	Path_ := self.Args("path").String()
	_int64 := model.GetSignupsCountViaPath(Path_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["signupsCount"] = 0
	}
	m["signupsCount"] = _int64
	return self.JSON(m)
}

func GetSignupsCountViaTitleHandler(self *macross.Context) error {
	Title_ := self.Args("title").String()
	_int64 := model.GetSignupsCountViaTitle(Title_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["signupsCount"] = 0
	}
	m["signupsCount"] = _int64
	return self.JSON(m)
}

func GetSignupsCountViaUserLoginHandler(self *macross.Context) error {
	UserLogin_ := self.Args("user_login").String()
	_int64 := model.GetSignupsCountViaUserLogin(UserLogin_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["signupsCount"] = 0
	}
	m["signupsCount"] = _int64
	return self.JSON(m)
}

func GetSignupsCountViaUserEmailHandler(self *macross.Context) error {
	UserEmail_ := self.Args("user_email").String()
	_int64 := model.GetSignupsCountViaUserEmail(UserEmail_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["signupsCount"] = 0
	}
	m["signupsCount"] = _int64
	return self.JSON(m)
}

func GetSignupsCountViaRegisteredHandler(self *macross.Context) error {
	Registered_ := self.Args("registered").Time()
	_int64 := model.GetSignupsCountViaRegistered(Registered_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["signupsCount"] = 0
	}
	m["signupsCount"] = _int64
	return self.JSON(m)
}

func GetSignupsCountViaActivatedHandler(self *macross.Context) error {
	Activated_ := self.Args("activated").Time()
	_int64 := model.GetSignupsCountViaActivated(Activated_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["signupsCount"] = 0
	}
	m["signupsCount"] = _int64
	return self.JSON(m)
}

func GetSignupsCountViaActiveHandler(self *macross.Context) error {
	Active_ := self.Args("active").MustInt()
	_int64 := model.GetSignupsCountViaActive(Active_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["signupsCount"] = 0
	}
	m["signupsCount"] = _int64
	return self.JSON(m)
}

func GetSignupsCountViaActivationKeyHandler(self *macross.Context) error {
	ActivationKey_ := self.Args("activation_key").String()
	_int64 := model.GetSignupsCountViaActivationKey(ActivationKey_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["signupsCount"] = 0
	}
	m["signupsCount"] = _int64
	return self.JSON(m)
}

func GetSignupsCountViaMetaHandler(self *macross.Context) error {
	Meta_ := self.Args("meta").String()
	_int64 := model.GetSignupsCountViaMeta(Meta_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["signupsCount"] = 0
	}
	m["signupsCount"] = _int64
	return self.JSON(m)
}

func GetSignupsesViaSignupIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iSignupId := self.Args("signup_id").MustInt64()
	if (offset > 0) && helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesViaSignupId(offset, limit, iSignupId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesViaSignupId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesViaDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDomain := self.Args("domain").String()
	if (offset > 0) && helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesViaDomain(offset, limit, iDomain, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesViaDomain's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesViaPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iPath := self.Args("path").String()
	if (offset > 0) && helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesViaPath(offset, limit, iPath, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesViaPath's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTitle := self.Args("title").String()
	if (offset > 0) && helper.IsHas(iTitle) {
		_Signups, _error := model.GetSignupsesViaTitle(offset, limit, iTitle, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesViaTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesViaUserLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUserLogin := self.Args("user_login").String()
	if (offset > 0) && helper.IsHas(iUserLogin) {
		_Signups, _error := model.GetSignupsesViaUserLogin(offset, limit, iUserLogin, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesViaUserLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesViaUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUserEmail := self.Args("user_email").String()
	if (offset > 0) && helper.IsHas(iUserEmail) {
		_Signups, _error := model.GetSignupsesViaUserEmail(offset, limit, iUserEmail, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesViaUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesViaRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iRegistered := self.Args("registered").Time()
	if (offset > 0) && helper.IsHas(iRegistered) {
		_Signups, _error := model.GetSignupsesViaRegistered(offset, limit, iRegistered, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesViaRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesViaActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iActivated := self.Args("activated").Time()
	if (offset > 0) && helper.IsHas(iActivated) {
		_Signups, _error := model.GetSignupsesViaActivated(offset, limit, iActivated, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesViaActivated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesViaActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iActive := self.Args("active").MustInt()
	if (offset > 0) && helper.IsHas(iActive) {
		_Signups, _error := model.GetSignupsesViaActive(offset, limit, iActive, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesViaActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesViaActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iActivationKey := self.Args("activation_key").String()
	if (offset > 0) && helper.IsHas(iActivationKey) {
		_Signups, _error := model.GetSignupsesViaActivationKey(offset, limit, iActivationKey, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesViaActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesViaMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iMeta := self.Args("meta").String()
	if (offset > 0) && helper.IsHas(iMeta) {
		_Signups, _error := model.GetSignupsesViaMeta(offset, limit, iMeta, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesViaMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndDomainAndPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iDomain := self.Args("domain").String()
	iPath := self.Args("path").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndDomainAndPath(offset, limit, iSignupId,iDomain,iPath)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndDomainAndPath's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndDomainAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iDomain := self.Args("domain").String()
	iTitle := self.Args("title").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndDomainAndTitle(offset, limit, iSignupId,iDomain,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndDomainAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndDomainAndUserLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iDomain := self.Args("domain").String()
	iUserLogin := self.Args("user_login").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndDomainAndUserLogin(offset, limit, iSignupId,iDomain,iUserLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndDomainAndUserLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndDomainAndUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iDomain := self.Args("domain").String()
	iUserEmail := self.Args("user_email").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndDomainAndUserEmail(offset, limit, iSignupId,iDomain,iUserEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndDomainAndUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndDomainAndRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iDomain := self.Args("domain").String()
	iRegistered := self.Args("registered").Time()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndDomainAndRegistered(offset, limit, iSignupId,iDomain,iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndDomainAndRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndDomainAndActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iDomain := self.Args("domain").String()
	iActivated := self.Args("activated").Time()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndDomainAndActivated(offset, limit, iSignupId,iDomain,iActivated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndDomainAndActivated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndDomainAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iDomain := self.Args("domain").String()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndDomainAndActive(offset, limit, iSignupId,iDomain,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndDomainAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndDomainAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iDomain := self.Args("domain").String()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndDomainAndActivationKey(offset, limit, iSignupId,iDomain,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndDomainAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndDomainAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iDomain := self.Args("domain").String()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndDomainAndMeta(offset, limit, iSignupId,iDomain,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndDomainAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndPathAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iPath := self.Args("path").String()
	iTitle := self.Args("title").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndPathAndTitle(offset, limit, iSignupId,iPath,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndPathAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndPathAndUserLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iPath := self.Args("path").String()
	iUserLogin := self.Args("user_login").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndPathAndUserLogin(offset, limit, iSignupId,iPath,iUserLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndPathAndUserLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndPathAndUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iPath := self.Args("path").String()
	iUserEmail := self.Args("user_email").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndPathAndUserEmail(offset, limit, iSignupId,iPath,iUserEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndPathAndUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndPathAndRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iPath := self.Args("path").String()
	iRegistered := self.Args("registered").Time()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndPathAndRegistered(offset, limit, iSignupId,iPath,iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndPathAndRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndPathAndActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iPath := self.Args("path").String()
	iActivated := self.Args("activated").Time()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndPathAndActivated(offset, limit, iSignupId,iPath,iActivated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndPathAndActivated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndPathAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iPath := self.Args("path").String()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndPathAndActive(offset, limit, iSignupId,iPath,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndPathAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndPathAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iPath := self.Args("path").String()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndPathAndActivationKey(offset, limit, iSignupId,iPath,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndPathAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndPathAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iPath := self.Args("path").String()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndPathAndMeta(offset, limit, iSignupId,iPath,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndPathAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndTitleAndUserLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iTitle := self.Args("title").String()
	iUserLogin := self.Args("user_login").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndTitleAndUserLogin(offset, limit, iSignupId,iTitle,iUserLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndTitleAndUserLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndTitleAndUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iTitle := self.Args("title").String()
	iUserEmail := self.Args("user_email").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndTitleAndUserEmail(offset, limit, iSignupId,iTitle,iUserEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndTitleAndUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndTitleAndRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iTitle := self.Args("title").String()
	iRegistered := self.Args("registered").Time()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndTitleAndRegistered(offset, limit, iSignupId,iTitle,iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndTitleAndRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndTitleAndActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iTitle := self.Args("title").String()
	iActivated := self.Args("activated").Time()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndTitleAndActivated(offset, limit, iSignupId,iTitle,iActivated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndTitleAndActivated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndTitleAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iTitle := self.Args("title").String()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndTitleAndActive(offset, limit, iSignupId,iTitle,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndTitleAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndTitleAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iTitle := self.Args("title").String()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndTitleAndActivationKey(offset, limit, iSignupId,iTitle,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndTitleAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndTitleAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iTitle := self.Args("title").String()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndTitleAndMeta(offset, limit, iSignupId,iTitle,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndTitleAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndUserLoginAndUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iUserLogin := self.Args("user_login").String()
	iUserEmail := self.Args("user_email").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndUserLoginAndUserEmail(offset, limit, iSignupId,iUserLogin,iUserEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndUserLoginAndUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndUserLoginAndRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iUserLogin := self.Args("user_login").String()
	iRegistered := self.Args("registered").Time()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndUserLoginAndRegistered(offset, limit, iSignupId,iUserLogin,iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndUserLoginAndRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndUserLoginAndActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iUserLogin := self.Args("user_login").String()
	iActivated := self.Args("activated").Time()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndUserLoginAndActivated(offset, limit, iSignupId,iUserLogin,iActivated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndUserLoginAndActivated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndUserLoginAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iUserLogin := self.Args("user_login").String()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndUserLoginAndActive(offset, limit, iSignupId,iUserLogin,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndUserLoginAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndUserLoginAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iUserLogin := self.Args("user_login").String()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndUserLoginAndActivationKey(offset, limit, iSignupId,iUserLogin,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndUserLoginAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndUserLoginAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iUserLogin := self.Args("user_login").String()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndUserLoginAndMeta(offset, limit, iSignupId,iUserLogin,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndUserLoginAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndUserEmailAndRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iUserEmail := self.Args("user_email").String()
	iRegistered := self.Args("registered").Time()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndUserEmailAndRegistered(offset, limit, iSignupId,iUserEmail,iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndUserEmailAndRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndUserEmailAndActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iUserEmail := self.Args("user_email").String()
	iActivated := self.Args("activated").Time()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndUserEmailAndActivated(offset, limit, iSignupId,iUserEmail,iActivated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndUserEmailAndActivated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndUserEmailAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iUserEmail := self.Args("user_email").String()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndUserEmailAndActive(offset, limit, iSignupId,iUserEmail,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndUserEmailAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndUserEmailAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iUserEmail := self.Args("user_email").String()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndUserEmailAndActivationKey(offset, limit, iSignupId,iUserEmail,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndUserEmailAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndUserEmailAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iUserEmail := self.Args("user_email").String()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndUserEmailAndMeta(offset, limit, iSignupId,iUserEmail,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndUserEmailAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndRegisteredAndActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iRegistered := self.Args("registered").Time()
	iActivated := self.Args("activated").Time()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndRegisteredAndActivated(offset, limit, iSignupId,iRegistered,iActivated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndRegisteredAndActivated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndRegisteredAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iRegistered := self.Args("registered").Time()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndRegisteredAndActive(offset, limit, iSignupId,iRegistered,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndRegisteredAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndRegisteredAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iRegistered := self.Args("registered").Time()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndRegisteredAndActivationKey(offset, limit, iSignupId,iRegistered,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndRegisteredAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndRegisteredAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iRegistered := self.Args("registered").Time()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndRegisteredAndMeta(offset, limit, iSignupId,iRegistered,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndRegisteredAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndActivatedAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iActivated := self.Args("activated").Time()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndActivatedAndActive(offset, limit, iSignupId,iActivated,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndActivatedAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndActivatedAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iActivated := self.Args("activated").Time()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndActivatedAndActivationKey(offset, limit, iSignupId,iActivated,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndActivatedAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndActivatedAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iActivated := self.Args("activated").Time()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndActivatedAndMeta(offset, limit, iSignupId,iActivated,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndActivatedAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndActiveAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iActive := self.Args("active").MustInt()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndActiveAndActivationKey(offset, limit, iSignupId,iActive,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndActiveAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndActiveAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iActive := self.Args("active").MustInt()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndActiveAndMeta(offset, limit, iSignupId,iActive,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndActiveAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndActivationKeyAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iActivationKey := self.Args("activation_key").String()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndActivationKeyAndMeta(offset, limit, iSignupId,iActivationKey,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndActivationKeyAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndPathAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iPath := self.Args("path").String()
	iTitle := self.Args("title").String()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndPathAndTitle(offset, limit, iDomain,iPath,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndPathAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndPathAndUserLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iPath := self.Args("path").String()
	iUserLogin := self.Args("user_login").String()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndPathAndUserLogin(offset, limit, iDomain,iPath,iUserLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndPathAndUserLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndPathAndUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iPath := self.Args("path").String()
	iUserEmail := self.Args("user_email").String()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndPathAndUserEmail(offset, limit, iDomain,iPath,iUserEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndPathAndUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndPathAndRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iPath := self.Args("path").String()
	iRegistered := self.Args("registered").Time()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndPathAndRegistered(offset, limit, iDomain,iPath,iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndPathAndRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndPathAndActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iPath := self.Args("path").String()
	iActivated := self.Args("activated").Time()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndPathAndActivated(offset, limit, iDomain,iPath,iActivated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndPathAndActivated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndPathAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iPath := self.Args("path").String()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndPathAndActive(offset, limit, iDomain,iPath,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndPathAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndPathAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iPath := self.Args("path").String()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndPathAndActivationKey(offset, limit, iDomain,iPath,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndPathAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndPathAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iPath := self.Args("path").String()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndPathAndMeta(offset, limit, iDomain,iPath,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndPathAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndTitleAndUserLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iTitle := self.Args("title").String()
	iUserLogin := self.Args("user_login").String()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndTitleAndUserLogin(offset, limit, iDomain,iTitle,iUserLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndTitleAndUserLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndTitleAndUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iTitle := self.Args("title").String()
	iUserEmail := self.Args("user_email").String()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndTitleAndUserEmail(offset, limit, iDomain,iTitle,iUserEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndTitleAndUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndTitleAndRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iTitle := self.Args("title").String()
	iRegistered := self.Args("registered").Time()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndTitleAndRegistered(offset, limit, iDomain,iTitle,iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndTitleAndRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndTitleAndActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iTitle := self.Args("title").String()
	iActivated := self.Args("activated").Time()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndTitleAndActivated(offset, limit, iDomain,iTitle,iActivated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndTitleAndActivated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndTitleAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iTitle := self.Args("title").String()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndTitleAndActive(offset, limit, iDomain,iTitle,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndTitleAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndTitleAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iTitle := self.Args("title").String()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndTitleAndActivationKey(offset, limit, iDomain,iTitle,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndTitleAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndTitleAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iTitle := self.Args("title").String()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndTitleAndMeta(offset, limit, iDomain,iTitle,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndTitleAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndUserLoginAndUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iUserLogin := self.Args("user_login").String()
	iUserEmail := self.Args("user_email").String()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndUserLoginAndUserEmail(offset, limit, iDomain,iUserLogin,iUserEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndUserLoginAndUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndUserLoginAndRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iUserLogin := self.Args("user_login").String()
	iRegistered := self.Args("registered").Time()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndUserLoginAndRegistered(offset, limit, iDomain,iUserLogin,iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndUserLoginAndRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndUserLoginAndActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iUserLogin := self.Args("user_login").String()
	iActivated := self.Args("activated").Time()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndUserLoginAndActivated(offset, limit, iDomain,iUserLogin,iActivated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndUserLoginAndActivated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndUserLoginAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iUserLogin := self.Args("user_login").String()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndUserLoginAndActive(offset, limit, iDomain,iUserLogin,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndUserLoginAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndUserLoginAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iUserLogin := self.Args("user_login").String()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndUserLoginAndActivationKey(offset, limit, iDomain,iUserLogin,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndUserLoginAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndUserLoginAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iUserLogin := self.Args("user_login").String()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndUserLoginAndMeta(offset, limit, iDomain,iUserLogin,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndUserLoginAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndUserEmailAndRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iUserEmail := self.Args("user_email").String()
	iRegistered := self.Args("registered").Time()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndUserEmailAndRegistered(offset, limit, iDomain,iUserEmail,iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndUserEmailAndRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndUserEmailAndActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iUserEmail := self.Args("user_email").String()
	iActivated := self.Args("activated").Time()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndUserEmailAndActivated(offset, limit, iDomain,iUserEmail,iActivated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndUserEmailAndActivated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndUserEmailAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iUserEmail := self.Args("user_email").String()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndUserEmailAndActive(offset, limit, iDomain,iUserEmail,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndUserEmailAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndUserEmailAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iUserEmail := self.Args("user_email").String()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndUserEmailAndActivationKey(offset, limit, iDomain,iUserEmail,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndUserEmailAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndUserEmailAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iUserEmail := self.Args("user_email").String()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndUserEmailAndMeta(offset, limit, iDomain,iUserEmail,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndUserEmailAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndRegisteredAndActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iRegistered := self.Args("registered").Time()
	iActivated := self.Args("activated").Time()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndRegisteredAndActivated(offset, limit, iDomain,iRegistered,iActivated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndRegisteredAndActivated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndRegisteredAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iRegistered := self.Args("registered").Time()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndRegisteredAndActive(offset, limit, iDomain,iRegistered,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndRegisteredAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndRegisteredAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iRegistered := self.Args("registered").Time()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndRegisteredAndActivationKey(offset, limit, iDomain,iRegistered,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndRegisteredAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndRegisteredAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iRegistered := self.Args("registered").Time()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndRegisteredAndMeta(offset, limit, iDomain,iRegistered,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndRegisteredAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndActivatedAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iActivated := self.Args("activated").Time()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndActivatedAndActive(offset, limit, iDomain,iActivated,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndActivatedAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndActivatedAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iActivated := self.Args("activated").Time()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndActivatedAndActivationKey(offset, limit, iDomain,iActivated,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndActivatedAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndActivatedAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iActivated := self.Args("activated").Time()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndActivatedAndMeta(offset, limit, iDomain,iActivated,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndActivatedAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndActiveAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iActive := self.Args("active").MustInt()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndActiveAndActivationKey(offset, limit, iDomain,iActive,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndActiveAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndActiveAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iActive := self.Args("active").MustInt()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndActiveAndMeta(offset, limit, iDomain,iActive,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndActiveAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndActivationKeyAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iActivationKey := self.Args("activation_key").String()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndActivationKeyAndMeta(offset, limit, iDomain,iActivationKey,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndActivationKeyAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndTitleAndUserLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iTitle := self.Args("title").String()
	iUserLogin := self.Args("user_login").String()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndTitleAndUserLogin(offset, limit, iPath,iTitle,iUserLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndTitleAndUserLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndTitleAndUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iTitle := self.Args("title").String()
	iUserEmail := self.Args("user_email").String()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndTitleAndUserEmail(offset, limit, iPath,iTitle,iUserEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndTitleAndUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndTitleAndRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iTitle := self.Args("title").String()
	iRegistered := self.Args("registered").Time()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndTitleAndRegistered(offset, limit, iPath,iTitle,iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndTitleAndRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndTitleAndActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iTitle := self.Args("title").String()
	iActivated := self.Args("activated").Time()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndTitleAndActivated(offset, limit, iPath,iTitle,iActivated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndTitleAndActivated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndTitleAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iTitle := self.Args("title").String()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndTitleAndActive(offset, limit, iPath,iTitle,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndTitleAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndTitleAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iTitle := self.Args("title").String()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndTitleAndActivationKey(offset, limit, iPath,iTitle,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndTitleAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndTitleAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iTitle := self.Args("title").String()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndTitleAndMeta(offset, limit, iPath,iTitle,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndTitleAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndUserLoginAndUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iUserLogin := self.Args("user_login").String()
	iUserEmail := self.Args("user_email").String()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndUserLoginAndUserEmail(offset, limit, iPath,iUserLogin,iUserEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndUserLoginAndUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndUserLoginAndRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iUserLogin := self.Args("user_login").String()
	iRegistered := self.Args("registered").Time()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndUserLoginAndRegistered(offset, limit, iPath,iUserLogin,iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndUserLoginAndRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndUserLoginAndActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iUserLogin := self.Args("user_login").String()
	iActivated := self.Args("activated").Time()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndUserLoginAndActivated(offset, limit, iPath,iUserLogin,iActivated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndUserLoginAndActivated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndUserLoginAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iUserLogin := self.Args("user_login").String()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndUserLoginAndActive(offset, limit, iPath,iUserLogin,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndUserLoginAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndUserLoginAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iUserLogin := self.Args("user_login").String()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndUserLoginAndActivationKey(offset, limit, iPath,iUserLogin,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndUserLoginAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndUserLoginAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iUserLogin := self.Args("user_login").String()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndUserLoginAndMeta(offset, limit, iPath,iUserLogin,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndUserLoginAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndUserEmailAndRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iUserEmail := self.Args("user_email").String()
	iRegistered := self.Args("registered").Time()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndUserEmailAndRegistered(offset, limit, iPath,iUserEmail,iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndUserEmailAndRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndUserEmailAndActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iUserEmail := self.Args("user_email").String()
	iActivated := self.Args("activated").Time()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndUserEmailAndActivated(offset, limit, iPath,iUserEmail,iActivated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndUserEmailAndActivated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndUserEmailAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iUserEmail := self.Args("user_email").String()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndUserEmailAndActive(offset, limit, iPath,iUserEmail,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndUserEmailAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndUserEmailAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iUserEmail := self.Args("user_email").String()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndUserEmailAndActivationKey(offset, limit, iPath,iUserEmail,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndUserEmailAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndUserEmailAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iUserEmail := self.Args("user_email").String()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndUserEmailAndMeta(offset, limit, iPath,iUserEmail,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndUserEmailAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndRegisteredAndActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iRegistered := self.Args("registered").Time()
	iActivated := self.Args("activated").Time()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndRegisteredAndActivated(offset, limit, iPath,iRegistered,iActivated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndRegisteredAndActivated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndRegisteredAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iRegistered := self.Args("registered").Time()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndRegisteredAndActive(offset, limit, iPath,iRegistered,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndRegisteredAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndRegisteredAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iRegistered := self.Args("registered").Time()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndRegisteredAndActivationKey(offset, limit, iPath,iRegistered,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndRegisteredAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndRegisteredAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iRegistered := self.Args("registered").Time()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndRegisteredAndMeta(offset, limit, iPath,iRegistered,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndRegisteredAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndActivatedAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iActivated := self.Args("activated").Time()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndActivatedAndActive(offset, limit, iPath,iActivated,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndActivatedAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndActivatedAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iActivated := self.Args("activated").Time()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndActivatedAndActivationKey(offset, limit, iPath,iActivated,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndActivatedAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndActivatedAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iActivated := self.Args("activated").Time()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndActivatedAndMeta(offset, limit, iPath,iActivated,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndActivatedAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndActiveAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iActive := self.Args("active").MustInt()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndActiveAndActivationKey(offset, limit, iPath,iActive,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndActiveAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndActiveAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iActive := self.Args("active").MustInt()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndActiveAndMeta(offset, limit, iPath,iActive,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndActiveAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndActivationKeyAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iActivationKey := self.Args("activation_key").String()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndActivationKeyAndMeta(offset, limit, iPath,iActivationKey,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndActivationKeyAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByTitleAndUserLoginAndUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iUserLogin := self.Args("user_login").String()
	iUserEmail := self.Args("user_email").String()

	if helper.IsHas(iTitle) {
		_Signups, _error := model.GetSignupsesByTitleAndUserLoginAndUserEmail(offset, limit, iTitle,iUserLogin,iUserEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByTitleAndUserLoginAndUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByTitleAndUserLoginAndRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iUserLogin := self.Args("user_login").String()
	iRegistered := self.Args("registered").Time()

	if helper.IsHas(iTitle) {
		_Signups, _error := model.GetSignupsesByTitleAndUserLoginAndRegistered(offset, limit, iTitle,iUserLogin,iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByTitleAndUserLoginAndRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByTitleAndUserLoginAndActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iUserLogin := self.Args("user_login").String()
	iActivated := self.Args("activated").Time()

	if helper.IsHas(iTitle) {
		_Signups, _error := model.GetSignupsesByTitleAndUserLoginAndActivated(offset, limit, iTitle,iUserLogin,iActivated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByTitleAndUserLoginAndActivated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByTitleAndUserLoginAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iUserLogin := self.Args("user_login").String()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iTitle) {
		_Signups, _error := model.GetSignupsesByTitleAndUserLoginAndActive(offset, limit, iTitle,iUserLogin,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByTitleAndUserLoginAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByTitleAndUserLoginAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iUserLogin := self.Args("user_login").String()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iTitle) {
		_Signups, _error := model.GetSignupsesByTitleAndUserLoginAndActivationKey(offset, limit, iTitle,iUserLogin,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByTitleAndUserLoginAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByTitleAndUserLoginAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iUserLogin := self.Args("user_login").String()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iTitle) {
		_Signups, _error := model.GetSignupsesByTitleAndUserLoginAndMeta(offset, limit, iTitle,iUserLogin,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByTitleAndUserLoginAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByTitleAndUserEmailAndRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iUserEmail := self.Args("user_email").String()
	iRegistered := self.Args("registered").Time()

	if helper.IsHas(iTitle) {
		_Signups, _error := model.GetSignupsesByTitleAndUserEmailAndRegistered(offset, limit, iTitle,iUserEmail,iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByTitleAndUserEmailAndRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByTitleAndUserEmailAndActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iUserEmail := self.Args("user_email").String()
	iActivated := self.Args("activated").Time()

	if helper.IsHas(iTitle) {
		_Signups, _error := model.GetSignupsesByTitleAndUserEmailAndActivated(offset, limit, iTitle,iUserEmail,iActivated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByTitleAndUserEmailAndActivated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByTitleAndUserEmailAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iUserEmail := self.Args("user_email").String()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iTitle) {
		_Signups, _error := model.GetSignupsesByTitleAndUserEmailAndActive(offset, limit, iTitle,iUserEmail,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByTitleAndUserEmailAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByTitleAndUserEmailAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iUserEmail := self.Args("user_email").String()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iTitle) {
		_Signups, _error := model.GetSignupsesByTitleAndUserEmailAndActivationKey(offset, limit, iTitle,iUserEmail,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByTitleAndUserEmailAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByTitleAndUserEmailAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iUserEmail := self.Args("user_email").String()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iTitle) {
		_Signups, _error := model.GetSignupsesByTitleAndUserEmailAndMeta(offset, limit, iTitle,iUserEmail,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByTitleAndUserEmailAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByTitleAndRegisteredAndActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iRegistered := self.Args("registered").Time()
	iActivated := self.Args("activated").Time()

	if helper.IsHas(iTitle) {
		_Signups, _error := model.GetSignupsesByTitleAndRegisteredAndActivated(offset, limit, iTitle,iRegistered,iActivated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByTitleAndRegisteredAndActivated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByTitleAndRegisteredAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iRegistered := self.Args("registered").Time()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iTitle) {
		_Signups, _error := model.GetSignupsesByTitleAndRegisteredAndActive(offset, limit, iTitle,iRegistered,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByTitleAndRegisteredAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByTitleAndRegisteredAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iRegistered := self.Args("registered").Time()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iTitle) {
		_Signups, _error := model.GetSignupsesByTitleAndRegisteredAndActivationKey(offset, limit, iTitle,iRegistered,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByTitleAndRegisteredAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByTitleAndRegisteredAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iRegistered := self.Args("registered").Time()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iTitle) {
		_Signups, _error := model.GetSignupsesByTitleAndRegisteredAndMeta(offset, limit, iTitle,iRegistered,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByTitleAndRegisteredAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByTitleAndActivatedAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iActivated := self.Args("activated").Time()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iTitle) {
		_Signups, _error := model.GetSignupsesByTitleAndActivatedAndActive(offset, limit, iTitle,iActivated,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByTitleAndActivatedAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByTitleAndActivatedAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iActivated := self.Args("activated").Time()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iTitle) {
		_Signups, _error := model.GetSignupsesByTitleAndActivatedAndActivationKey(offset, limit, iTitle,iActivated,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByTitleAndActivatedAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByTitleAndActivatedAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iActivated := self.Args("activated").Time()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iTitle) {
		_Signups, _error := model.GetSignupsesByTitleAndActivatedAndMeta(offset, limit, iTitle,iActivated,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByTitleAndActivatedAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByTitleAndActiveAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iActive := self.Args("active").MustInt()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iTitle) {
		_Signups, _error := model.GetSignupsesByTitleAndActiveAndActivationKey(offset, limit, iTitle,iActive,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByTitleAndActiveAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByTitleAndActiveAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iActive := self.Args("active").MustInt()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iTitle) {
		_Signups, _error := model.GetSignupsesByTitleAndActiveAndMeta(offset, limit, iTitle,iActive,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByTitleAndActiveAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByTitleAndActivationKeyAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iActivationKey := self.Args("activation_key").String()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iTitle) {
		_Signups, _error := model.GetSignupsesByTitleAndActivationKeyAndMeta(offset, limit, iTitle,iActivationKey,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByTitleAndActivationKeyAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserLoginAndUserEmailAndRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserEmail := self.Args("user_email").String()
	iRegistered := self.Args("registered").Time()

	if helper.IsHas(iUserLogin) {
		_Signups, _error := model.GetSignupsesByUserLoginAndUserEmailAndRegistered(offset, limit, iUserLogin,iUserEmail,iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserLoginAndUserEmailAndRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserLoginAndUserEmailAndActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserEmail := self.Args("user_email").String()
	iActivated := self.Args("activated").Time()

	if helper.IsHas(iUserLogin) {
		_Signups, _error := model.GetSignupsesByUserLoginAndUserEmailAndActivated(offset, limit, iUserLogin,iUserEmail,iActivated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserLoginAndUserEmailAndActivated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserLoginAndUserEmailAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserEmail := self.Args("user_email").String()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iUserLogin) {
		_Signups, _error := model.GetSignupsesByUserLoginAndUserEmailAndActive(offset, limit, iUserLogin,iUserEmail,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserLoginAndUserEmailAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserLoginAndUserEmailAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserEmail := self.Args("user_email").String()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iUserLogin) {
		_Signups, _error := model.GetSignupsesByUserLoginAndUserEmailAndActivationKey(offset, limit, iUserLogin,iUserEmail,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserLoginAndUserEmailAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserLoginAndUserEmailAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserEmail := self.Args("user_email").String()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iUserLogin) {
		_Signups, _error := model.GetSignupsesByUserLoginAndUserEmailAndMeta(offset, limit, iUserLogin,iUserEmail,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserLoginAndUserEmailAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserLoginAndRegisteredAndActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iRegistered := self.Args("registered").Time()
	iActivated := self.Args("activated").Time()

	if helper.IsHas(iUserLogin) {
		_Signups, _error := model.GetSignupsesByUserLoginAndRegisteredAndActivated(offset, limit, iUserLogin,iRegistered,iActivated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserLoginAndRegisteredAndActivated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserLoginAndRegisteredAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iRegistered := self.Args("registered").Time()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iUserLogin) {
		_Signups, _error := model.GetSignupsesByUserLoginAndRegisteredAndActive(offset, limit, iUserLogin,iRegistered,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserLoginAndRegisteredAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserLoginAndRegisteredAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iRegistered := self.Args("registered").Time()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iUserLogin) {
		_Signups, _error := model.GetSignupsesByUserLoginAndRegisteredAndActivationKey(offset, limit, iUserLogin,iRegistered,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserLoginAndRegisteredAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserLoginAndRegisteredAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iRegistered := self.Args("registered").Time()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iUserLogin) {
		_Signups, _error := model.GetSignupsesByUserLoginAndRegisteredAndMeta(offset, limit, iUserLogin,iRegistered,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserLoginAndRegisteredAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserLoginAndActivatedAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iActivated := self.Args("activated").Time()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iUserLogin) {
		_Signups, _error := model.GetSignupsesByUserLoginAndActivatedAndActive(offset, limit, iUserLogin,iActivated,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserLoginAndActivatedAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserLoginAndActivatedAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iActivated := self.Args("activated").Time()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iUserLogin) {
		_Signups, _error := model.GetSignupsesByUserLoginAndActivatedAndActivationKey(offset, limit, iUserLogin,iActivated,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserLoginAndActivatedAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserLoginAndActivatedAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iActivated := self.Args("activated").Time()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iUserLogin) {
		_Signups, _error := model.GetSignupsesByUserLoginAndActivatedAndMeta(offset, limit, iUserLogin,iActivated,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserLoginAndActivatedAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserLoginAndActiveAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iActive := self.Args("active").MustInt()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iUserLogin) {
		_Signups, _error := model.GetSignupsesByUserLoginAndActiveAndActivationKey(offset, limit, iUserLogin,iActive,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserLoginAndActiveAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserLoginAndActiveAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iActive := self.Args("active").MustInt()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iUserLogin) {
		_Signups, _error := model.GetSignupsesByUserLoginAndActiveAndMeta(offset, limit, iUserLogin,iActive,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserLoginAndActiveAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserLoginAndActivationKeyAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iActivationKey := self.Args("activation_key").String()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iUserLogin) {
		_Signups, _error := model.GetSignupsesByUserLoginAndActivationKeyAndMeta(offset, limit, iUserLogin,iActivationKey,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserLoginAndActivationKeyAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserEmailAndRegisteredAndActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iRegistered := self.Args("registered").Time()
	iActivated := self.Args("activated").Time()

	if helper.IsHas(iUserEmail) {
		_Signups, _error := model.GetSignupsesByUserEmailAndRegisteredAndActivated(offset, limit, iUserEmail,iRegistered,iActivated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserEmailAndRegisteredAndActivated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserEmailAndRegisteredAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iRegistered := self.Args("registered").Time()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iUserEmail) {
		_Signups, _error := model.GetSignupsesByUserEmailAndRegisteredAndActive(offset, limit, iUserEmail,iRegistered,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserEmailAndRegisteredAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserEmailAndRegisteredAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iRegistered := self.Args("registered").Time()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iUserEmail) {
		_Signups, _error := model.GetSignupsesByUserEmailAndRegisteredAndActivationKey(offset, limit, iUserEmail,iRegistered,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserEmailAndRegisteredAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserEmailAndRegisteredAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iRegistered := self.Args("registered").Time()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iUserEmail) {
		_Signups, _error := model.GetSignupsesByUserEmailAndRegisteredAndMeta(offset, limit, iUserEmail,iRegistered,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserEmailAndRegisteredAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserEmailAndActivatedAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iActivated := self.Args("activated").Time()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iUserEmail) {
		_Signups, _error := model.GetSignupsesByUserEmailAndActivatedAndActive(offset, limit, iUserEmail,iActivated,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserEmailAndActivatedAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserEmailAndActivatedAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iActivated := self.Args("activated").Time()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iUserEmail) {
		_Signups, _error := model.GetSignupsesByUserEmailAndActivatedAndActivationKey(offset, limit, iUserEmail,iActivated,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserEmailAndActivatedAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserEmailAndActivatedAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iActivated := self.Args("activated").Time()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iUserEmail) {
		_Signups, _error := model.GetSignupsesByUserEmailAndActivatedAndMeta(offset, limit, iUserEmail,iActivated,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserEmailAndActivatedAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserEmailAndActiveAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iActive := self.Args("active").MustInt()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iUserEmail) {
		_Signups, _error := model.GetSignupsesByUserEmailAndActiveAndActivationKey(offset, limit, iUserEmail,iActive,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserEmailAndActiveAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserEmailAndActiveAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iActive := self.Args("active").MustInt()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iUserEmail) {
		_Signups, _error := model.GetSignupsesByUserEmailAndActiveAndMeta(offset, limit, iUserEmail,iActive,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserEmailAndActiveAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserEmailAndActivationKeyAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iActivationKey := self.Args("activation_key").String()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iUserEmail) {
		_Signups, _error := model.GetSignupsesByUserEmailAndActivationKeyAndMeta(offset, limit, iUserEmail,iActivationKey,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserEmailAndActivationKeyAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByRegisteredAndActivatedAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iActivated := self.Args("activated").Time()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iRegistered) {
		_Signups, _error := model.GetSignupsesByRegisteredAndActivatedAndActive(offset, limit, iRegistered,iActivated,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByRegisteredAndActivatedAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByRegisteredAndActivatedAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iActivated := self.Args("activated").Time()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iRegistered) {
		_Signups, _error := model.GetSignupsesByRegisteredAndActivatedAndActivationKey(offset, limit, iRegistered,iActivated,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByRegisteredAndActivatedAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByRegisteredAndActivatedAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iActivated := self.Args("activated").Time()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iRegistered) {
		_Signups, _error := model.GetSignupsesByRegisteredAndActivatedAndMeta(offset, limit, iRegistered,iActivated,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByRegisteredAndActivatedAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByRegisteredAndActiveAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iActive := self.Args("active").MustInt()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iRegistered) {
		_Signups, _error := model.GetSignupsesByRegisteredAndActiveAndActivationKey(offset, limit, iRegistered,iActive,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByRegisteredAndActiveAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByRegisteredAndActiveAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iActive := self.Args("active").MustInt()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iRegistered) {
		_Signups, _error := model.GetSignupsesByRegisteredAndActiveAndMeta(offset, limit, iRegistered,iActive,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByRegisteredAndActiveAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByRegisteredAndActivationKeyAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iActivationKey := self.Args("activation_key").String()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iRegistered) {
		_Signups, _error := model.GetSignupsesByRegisteredAndActivationKeyAndMeta(offset, limit, iRegistered,iActivationKey,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByRegisteredAndActivationKeyAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByActivatedAndActiveAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iActivated := self.Args("activated").Time()
	iActive := self.Args("active").MustInt()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iActivated) {
		_Signups, _error := model.GetSignupsesByActivatedAndActiveAndActivationKey(offset, limit, iActivated,iActive,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByActivatedAndActiveAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByActivatedAndActiveAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iActivated := self.Args("activated").Time()
	iActive := self.Args("active").MustInt()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iActivated) {
		_Signups, _error := model.GetSignupsesByActivatedAndActiveAndMeta(offset, limit, iActivated,iActive,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByActivatedAndActiveAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByActivatedAndActivationKeyAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iActivated := self.Args("activated").Time()
	iActivationKey := self.Args("activation_key").String()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iActivated) {
		_Signups, _error := model.GetSignupsesByActivatedAndActivationKeyAndMeta(offset, limit, iActivated,iActivationKey,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByActivatedAndActivationKeyAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByActiveAndActivationKeyAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iActive := self.Args("active").MustInt()
	iActivationKey := self.Args("activation_key").String()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iActive) {
		_Signups, _error := model.GetSignupsesByActiveAndActivationKeyAndMeta(offset, limit, iActive,iActivationKey,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByActiveAndActivationKeyAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iDomain := self.Args("domain").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndDomain(offset, limit, iSignupId,iDomain)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndDomain's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iPath := self.Args("path").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndPath(offset, limit, iSignupId,iPath)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndPath's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iTitle := self.Args("title").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndTitle(offset, limit, iSignupId,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndUserLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iUserLogin := self.Args("user_login").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndUserLogin(offset, limit, iSignupId,iUserLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndUserLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iUserEmail := self.Args("user_email").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndUserEmail(offset, limit, iSignupId,iUserEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iRegistered := self.Args("registered").Time()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndRegistered(offset, limit, iSignupId,iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iActivated := self.Args("activated").Time()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndActivated(offset, limit, iSignupId,iActivated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndActivated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndActive(offset, limit, iSignupId,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndActivationKey(offset, limit, iSignupId,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesBySignupIdAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSignupId := self.Args("signup_id").MustInt64()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsesBySignupIdAndMeta(offset, limit, iSignupId,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesBySignupIdAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iPath := self.Args("path").String()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndPath(offset, limit, iDomain,iPath)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndPath's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iTitle := self.Args("title").String()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndTitle(offset, limit, iDomain,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndUserLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iUserLogin := self.Args("user_login").String()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndUserLogin(offset, limit, iDomain,iUserLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndUserLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iUserEmail := self.Args("user_email").String()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndUserEmail(offset, limit, iDomain,iUserEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iRegistered := self.Args("registered").Time()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndRegistered(offset, limit, iDomain,iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iActivated := self.Args("activated").Time()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndActivated(offset, limit, iDomain,iActivated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndActivated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndActive(offset, limit, iDomain,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndActivationKey(offset, limit, iDomain,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByDomainAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDomain := self.Args("domain").String()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsesByDomainAndMeta(offset, limit, iDomain,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByDomainAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iTitle := self.Args("title").String()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndTitle(offset, limit, iPath,iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndUserLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iUserLogin := self.Args("user_login").String()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndUserLogin(offset, limit, iPath,iUserLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndUserLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iUserEmail := self.Args("user_email").String()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndUserEmail(offset, limit, iPath,iUserEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iRegistered := self.Args("registered").Time()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndRegistered(offset, limit, iPath,iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iActivated := self.Args("activated").Time()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndActivated(offset, limit, iPath,iActivated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndActivated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndActive(offset, limit, iPath,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndActivationKey(offset, limit, iPath,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByPathAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iPath := self.Args("path").String()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsesByPathAndMeta(offset, limit, iPath,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByPathAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByTitleAndUserLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iUserLogin := self.Args("user_login").String()

	if helper.IsHas(iTitle) {
		_Signups, _error := model.GetSignupsesByTitleAndUserLogin(offset, limit, iTitle,iUserLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByTitleAndUserLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByTitleAndUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iUserEmail := self.Args("user_email").String()

	if helper.IsHas(iTitle) {
		_Signups, _error := model.GetSignupsesByTitleAndUserEmail(offset, limit, iTitle,iUserEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByTitleAndUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByTitleAndRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iRegistered := self.Args("registered").Time()

	if helper.IsHas(iTitle) {
		_Signups, _error := model.GetSignupsesByTitleAndRegistered(offset, limit, iTitle,iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByTitleAndRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByTitleAndActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iActivated := self.Args("activated").Time()

	if helper.IsHas(iTitle) {
		_Signups, _error := model.GetSignupsesByTitleAndActivated(offset, limit, iTitle,iActivated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByTitleAndActivated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByTitleAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iTitle) {
		_Signups, _error := model.GetSignupsesByTitleAndActive(offset, limit, iTitle,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByTitleAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByTitleAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iTitle) {
		_Signups, _error := model.GetSignupsesByTitleAndActivationKey(offset, limit, iTitle,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByTitleAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByTitleAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTitle := self.Args("title").String()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iTitle) {
		_Signups, _error := model.GetSignupsesByTitleAndMeta(offset, limit, iTitle,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByTitleAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserLoginAndUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserEmail := self.Args("user_email").String()

	if helper.IsHas(iUserLogin) {
		_Signups, _error := model.GetSignupsesByUserLoginAndUserEmail(offset, limit, iUserLogin,iUserEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserLoginAndUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserLoginAndRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iRegistered := self.Args("registered").Time()

	if helper.IsHas(iUserLogin) {
		_Signups, _error := model.GetSignupsesByUserLoginAndRegistered(offset, limit, iUserLogin,iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserLoginAndRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserLoginAndActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iActivated := self.Args("activated").Time()

	if helper.IsHas(iUserLogin) {
		_Signups, _error := model.GetSignupsesByUserLoginAndActivated(offset, limit, iUserLogin,iActivated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserLoginAndActivated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserLoginAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iUserLogin) {
		_Signups, _error := model.GetSignupsesByUserLoginAndActive(offset, limit, iUserLogin,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserLoginAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserLoginAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iUserLogin) {
		_Signups, _error := model.GetSignupsesByUserLoginAndActivationKey(offset, limit, iUserLogin,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserLoginAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserLoginAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iUserLogin) {
		_Signups, _error := model.GetSignupsesByUserLoginAndMeta(offset, limit, iUserLogin,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserLoginAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserEmailAndRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iRegistered := self.Args("registered").Time()

	if helper.IsHas(iUserEmail) {
		_Signups, _error := model.GetSignupsesByUserEmailAndRegistered(offset, limit, iUserEmail,iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserEmailAndRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserEmailAndActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iActivated := self.Args("activated").Time()

	if helper.IsHas(iUserEmail) {
		_Signups, _error := model.GetSignupsesByUserEmailAndActivated(offset, limit, iUserEmail,iActivated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserEmailAndActivated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserEmailAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iUserEmail) {
		_Signups, _error := model.GetSignupsesByUserEmailAndActive(offset, limit, iUserEmail,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserEmailAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserEmailAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iUserEmail) {
		_Signups, _error := model.GetSignupsesByUserEmailAndActivationKey(offset, limit, iUserEmail,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserEmailAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByUserEmailAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iUserEmail) {
		_Signups, _error := model.GetSignupsesByUserEmailAndMeta(offset, limit, iUserEmail,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByUserEmailAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByRegisteredAndActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iActivated := self.Args("activated").Time()

	if helper.IsHas(iRegistered) {
		_Signups, _error := model.GetSignupsesByRegisteredAndActivated(offset, limit, iRegistered,iActivated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByRegisteredAndActivated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByRegisteredAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iRegistered) {
		_Signups, _error := model.GetSignupsesByRegisteredAndActive(offset, limit, iRegistered,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByRegisteredAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByRegisteredAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iRegistered) {
		_Signups, _error := model.GetSignupsesByRegisteredAndActivationKey(offset, limit, iRegistered,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByRegisteredAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByRegisteredAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iRegistered := self.Args("registered").Time()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iRegistered) {
		_Signups, _error := model.GetSignupsesByRegisteredAndMeta(offset, limit, iRegistered,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByRegisteredAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByActivatedAndActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iActivated := self.Args("activated").Time()
	iActive := self.Args("active").MustInt()

	if helper.IsHas(iActivated) {
		_Signups, _error := model.GetSignupsesByActivatedAndActive(offset, limit, iActivated,iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByActivatedAndActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByActivatedAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iActivated := self.Args("activated").Time()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iActivated) {
		_Signups, _error := model.GetSignupsesByActivatedAndActivationKey(offset, limit, iActivated,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByActivatedAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByActivatedAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iActivated := self.Args("activated").Time()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iActivated) {
		_Signups, _error := model.GetSignupsesByActivatedAndMeta(offset, limit, iActivated,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByActivatedAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByActiveAndActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iActive := self.Args("active").MustInt()
	iActivationKey := self.Args("activation_key").String()

	if helper.IsHas(iActive) {
		_Signups, _error := model.GetSignupsesByActiveAndActivationKey(offset, limit, iActive,iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByActiveAndActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByActiveAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iActive := self.Args("active").MustInt()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iActive) {
		_Signups, _error := model.GetSignupsesByActiveAndMeta(offset, limit, iActive,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByActiveAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesByActivationKeyAndMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iActivationKey := self.Args("activation_key").String()
	iMeta := self.Args("meta").String()

	if helper.IsHas(iActivationKey) {
		_Signups, _error := model.GetSignupsesByActivationKeyAndMeta(offset, limit, iActivationKey,iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsesByActivationKeyAndMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_Signups, _error := model.GetSignupses(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupses' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSignupsViaSignupIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSignupId := self.Args("signup_id").MustInt64()
	if helper.IsHas(iSignupId) {
		_Signups := model.HasSignupsViaSignupId(iSignupId)
		var m = map[string]interface{}{}
		m["signups"] = _Signups
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSignupsViaSignupId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSignupsViaDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDomain := self.Args("domain").String()
	if helper.IsHas(iDomain) {
		_Signups := model.HasSignupsViaDomain(iDomain)
		var m = map[string]interface{}{}
		m["signups"] = _Signups
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSignupsViaDomain's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSignupsViaPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iPath := self.Args("path").String()
	if helper.IsHas(iPath) {
		_Signups := model.HasSignupsViaPath(iPath)
		var m = map[string]interface{}{}
		m["signups"] = _Signups
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSignupsViaPath's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSignupsViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTitle := self.Args("title").String()
	if helper.IsHas(iTitle) {
		_Signups := model.HasSignupsViaTitle(iTitle)
		var m = map[string]interface{}{}
		m["signups"] = _Signups
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSignupsViaTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSignupsViaUserLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserLogin := self.Args("user_login").String()
	if helper.IsHas(iUserLogin) {
		_Signups := model.HasSignupsViaUserLogin(iUserLogin)
		var m = map[string]interface{}{}
		m["signups"] = _Signups
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSignupsViaUserLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSignupsViaUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserEmail := self.Args("user_email").String()
	if helper.IsHas(iUserEmail) {
		_Signups := model.HasSignupsViaUserEmail(iUserEmail)
		var m = map[string]interface{}{}
		m["signups"] = _Signups
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSignupsViaUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSignupsViaRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRegistered := self.Args("registered").Time()
	if helper.IsHas(iRegistered) {
		_Signups := model.HasSignupsViaRegistered(iRegistered)
		var m = map[string]interface{}{}
		m["signups"] = _Signups
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSignupsViaRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSignupsViaActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iActivated := self.Args("activated").Time()
	if helper.IsHas(iActivated) {
		_Signups := model.HasSignupsViaActivated(iActivated)
		var m = map[string]interface{}{}
		m["signups"] = _Signups
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSignupsViaActivated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSignupsViaActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iActive := self.Args("active").MustInt()
	if helper.IsHas(iActive) {
		_Signups := model.HasSignupsViaActive(iActive)
		var m = map[string]interface{}{}
		m["signups"] = _Signups
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSignupsViaActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSignupsViaActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iActivationKey := self.Args("activation_key").String()
	if helper.IsHas(iActivationKey) {
		_Signups := model.HasSignupsViaActivationKey(iActivationKey)
		var m = map[string]interface{}{}
		m["signups"] = _Signups
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSignupsViaActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasSignupsViaMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMeta := self.Args("meta").String()
	if helper.IsHas(iMeta) {
		_Signups := model.HasSignupsViaMeta(iMeta)
		var m = map[string]interface{}{}
		m["signups"] = _Signups
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasSignupsViaMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsViaSignupIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSignupId := self.Args("signup_id").MustInt64()
	if helper.IsHas(iSignupId) {
		_Signups, _error := model.GetSignupsViaSignupId(iSignupId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsViaSignupId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsViaDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDomain := self.Args("domain").String()
	if helper.IsHas(iDomain) {
		_Signups, _error := model.GetSignupsViaDomain(iDomain)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsViaDomain's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsViaPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iPath := self.Args("path").String()
	if helper.IsHas(iPath) {
		_Signups, _error := model.GetSignupsViaPath(iPath)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsViaPath's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTitle := self.Args("title").String()
	if helper.IsHas(iTitle) {
		_Signups, _error := model.GetSignupsViaTitle(iTitle)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsViaTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsViaUserLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserLogin := self.Args("user_login").String()
	if helper.IsHas(iUserLogin) {
		_Signups, _error := model.GetSignupsViaUserLogin(iUserLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsViaUserLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsViaUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserEmail := self.Args("user_email").String()
	if helper.IsHas(iUserEmail) {
		_Signups, _error := model.GetSignupsViaUserEmail(iUserEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsViaUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsViaRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iRegistered := self.Args("registered").Time()
	if helper.IsHas(iRegistered) {
		_Signups, _error := model.GetSignupsViaRegistered(iRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsViaRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsViaActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iActivated := self.Args("activated").Time()
	if helper.IsHas(iActivated) {
		_Signups, _error := model.GetSignupsViaActivated(iActivated)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsViaActivated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsViaActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iActive := self.Args("active").MustInt()
	if helper.IsHas(iActive) {
		_Signups, _error := model.GetSignupsViaActive(iActive)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsViaActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsViaActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iActivationKey := self.Args("activation_key").String()
	if helper.IsHas(iActivationKey) {
		_Signups, _error := model.GetSignupsViaActivationKey(iActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsViaActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetSignupsViaMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iMeta := self.Args("meta").String()
	if helper.IsHas(iMeta) {
		_Signups, _error := model.GetSignupsViaMeta(iMeta)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the GetSignupsViaMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSignupsViaSignupIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	SignupId_ := self.Args("signup_id").MustInt64()
	if helper.IsHas(SignupId_) {
		var iSignups model.Signups
		self.Bind(&iSignups)
		_Signups, _error := model.SetSignupsViaSignupId(SignupId_, &iSignups)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the SetSignupsViaSignupId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSignupsViaDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Domain_ := self.Args("domain").String()
	if helper.IsHas(Domain_) {
		var iSignups model.Signups
		self.Bind(&iSignups)
		_Signups, _error := model.SetSignupsViaDomain(Domain_, &iSignups)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the SetSignupsViaDomain's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSignupsViaPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Path_ := self.Args("path").String()
	if helper.IsHas(Path_) {
		var iSignups model.Signups
		self.Bind(&iSignups)
		_Signups, _error := model.SetSignupsViaPath(Path_, &iSignups)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the SetSignupsViaPath's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSignupsViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Title_ := self.Args("title").String()
	if helper.IsHas(Title_) {
		var iSignups model.Signups
		self.Bind(&iSignups)
		_Signups, _error := model.SetSignupsViaTitle(Title_, &iSignups)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the SetSignupsViaTitle's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSignupsViaUserLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserLogin_ := self.Args("user_login").String()
	if helper.IsHas(UserLogin_) {
		var iSignups model.Signups
		self.Bind(&iSignups)
		_Signups, _error := model.SetSignupsViaUserLogin(UserLogin_, &iSignups)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the SetSignupsViaUserLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSignupsViaUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserEmail_ := self.Args("user_email").String()
	if helper.IsHas(UserEmail_) {
		var iSignups model.Signups
		self.Bind(&iSignups)
		_Signups, _error := model.SetSignupsViaUserEmail(UserEmail_, &iSignups)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the SetSignupsViaUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSignupsViaRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Registered_ := self.Args("registered").Time()
	if helper.IsHas(Registered_) {
		var iSignups model.Signups
		self.Bind(&iSignups)
		_Signups, _error := model.SetSignupsViaRegistered(Registered_, &iSignups)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the SetSignupsViaRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSignupsViaActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Activated_ := self.Args("activated").Time()
	if helper.IsHas(Activated_) {
		var iSignups model.Signups
		self.Bind(&iSignups)
		_Signups, _error := model.SetSignupsViaActivated(Activated_, &iSignups)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the SetSignupsViaActivated's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSignupsViaActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Active_ := self.Args("active").MustInt()
	if helper.IsHas(Active_) {
		var iSignups model.Signups
		self.Bind(&iSignups)
		_Signups, _error := model.SetSignupsViaActive(Active_, &iSignups)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the SetSignupsViaActive's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSignupsViaActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	ActivationKey_ := self.Args("activation_key").String()
	if helper.IsHas(ActivationKey_) {
		var iSignups model.Signups
		self.Bind(&iSignups)
		_Signups, _error := model.SetSignupsViaActivationKey(ActivationKey_, &iSignups)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the SetSignupsViaActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetSignupsViaMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Meta_ := self.Args("meta").String()
	if helper.IsHas(Meta_) {
		var iSignups model.Signups
		self.Bind(&iSignups)
		_Signups, _error := model.SetSignupsViaMeta(Meta_, &iSignups)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Signups)
	}
	herr.Message = "Can't get to the SetSignupsViaMeta's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddSignupsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	SignupId_ := self.Args("signup_id").MustInt64()
	Domain_ := self.Args("domain").String()
	Path_ := self.Args("path").String()
	Title_ := self.Args("title").String()
	UserLogin_ := self.Args("user_login").String()
	UserEmail_ := self.Args("user_email").String()
	Registered_ := self.Args("registered").Time()
	Activated_ := self.Args("activated").Time()
	Active_ := self.Args("active").MustInt()
	ActivationKey_ := self.Args("activation_key").String()
	Meta_ := self.Args("meta").String()

	if helper.IsHas(SignupId_) {
		_error := model.AddSignups(SignupId_,Domain_,Path_,Title_,UserLogin_,UserEmail_,Registered_,Activated_,Active_,ActivationKey_,Meta_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddSignups's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSignupsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iSignups model.Signups
	self.Bind(&iSignups)
	_int64, _error := model.PostSignups(&iSignups)
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

func PutSignupsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iSignups model.Signups
	self.Bind(&iSignups)
	_int64, _error := model.PutSignups(&iSignups)
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

func PutSignupsViaSignupIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	SignupId_ := self.Args("signup_id").MustInt64()
	var iSignups model.Signups
	self.Bind(&iSignups)
	_int64, _error := model.PutSignupsViaSignupId(SignupId_, &iSignups)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSignupsViaDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Domain_ := self.Args("domain").String()
	var iSignups model.Signups
	self.Bind(&iSignups)
	_int64, _error := model.PutSignupsViaDomain(Domain_, &iSignups)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSignupsViaPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Path_ := self.Args("path").String()
	var iSignups model.Signups
	self.Bind(&iSignups)
	_int64, _error := model.PutSignupsViaPath(Path_, &iSignups)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSignupsViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Title_ := self.Args("title").String()
	var iSignups model.Signups
	self.Bind(&iSignups)
	_int64, _error := model.PutSignupsViaTitle(Title_, &iSignups)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSignupsViaUserLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserLogin_ := self.Args("user_login").String()
	var iSignups model.Signups
	self.Bind(&iSignups)
	_int64, _error := model.PutSignupsViaUserLogin(UserLogin_, &iSignups)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSignupsViaUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserEmail_ := self.Args("user_email").String()
	var iSignups model.Signups
	self.Bind(&iSignups)
	_int64, _error := model.PutSignupsViaUserEmail(UserEmail_, &iSignups)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSignupsViaRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Registered_ := self.Args("registered").Time()
	var iSignups model.Signups
	self.Bind(&iSignups)
	_int64, _error := model.PutSignupsViaRegistered(Registered_, &iSignups)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSignupsViaActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Activated_ := self.Args("activated").Time()
	var iSignups model.Signups
	self.Bind(&iSignups)
	_int64, _error := model.PutSignupsViaActivated(Activated_, &iSignups)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSignupsViaActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Active_ := self.Args("active").MustInt()
	var iSignups model.Signups
	self.Bind(&iSignups)
	_int64, _error := model.PutSignupsViaActive(Active_, &iSignups)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSignupsViaActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	ActivationKey_ := self.Args("activation_key").String()
	var iSignups model.Signups
	self.Bind(&iSignups)
	_int64, _error := model.PutSignupsViaActivationKey(ActivationKey_, &iSignups)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutSignupsViaMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Meta_ := self.Args("meta").String()
	var iSignups model.Signups
	self.Bind(&iSignups)
	_int64, _error := model.PutSignupsViaMeta(Meta_, &iSignups)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSignupsViaSignupIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	SignupId_ := self.Args("signup_id").MustInt64()
	var iSignups model.Signups
	self.Bind(&iSignups)
	var iMap = helper.StructToMap(iSignups)
	_error := model.UpdateSignupsViaSignupId(SignupId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSignupsViaDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Domain_ := self.Args("domain").String()
	var iSignups model.Signups
	self.Bind(&iSignups)
	var iMap = helper.StructToMap(iSignups)
	_error := model.UpdateSignupsViaDomain(Domain_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSignupsViaPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Path_ := self.Args("path").String()
	var iSignups model.Signups
	self.Bind(&iSignups)
	var iMap = helper.StructToMap(iSignups)
	_error := model.UpdateSignupsViaPath(Path_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSignupsViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Title_ := self.Args("title").String()
	var iSignups model.Signups
	self.Bind(&iSignups)
	var iMap = helper.StructToMap(iSignups)
	_error := model.UpdateSignupsViaTitle(Title_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSignupsViaUserLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserLogin_ := self.Args("user_login").String()
	var iSignups model.Signups
	self.Bind(&iSignups)
	var iMap = helper.StructToMap(iSignups)
	_error := model.UpdateSignupsViaUserLogin(UserLogin_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSignupsViaUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserEmail_ := self.Args("user_email").String()
	var iSignups model.Signups
	self.Bind(&iSignups)
	var iMap = helper.StructToMap(iSignups)
	_error := model.UpdateSignupsViaUserEmail(UserEmail_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSignupsViaRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Registered_ := self.Args("registered").Time()
	var iSignups model.Signups
	self.Bind(&iSignups)
	var iMap = helper.StructToMap(iSignups)
	_error := model.UpdateSignupsViaRegistered(Registered_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSignupsViaActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Activated_ := self.Args("activated").Time()
	var iSignups model.Signups
	self.Bind(&iSignups)
	var iMap = helper.StructToMap(iSignups)
	_error := model.UpdateSignupsViaActivated(Activated_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSignupsViaActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Active_ := self.Args("active").MustInt()
	var iSignups model.Signups
	self.Bind(&iSignups)
	var iMap = helper.StructToMap(iSignups)
	_error := model.UpdateSignupsViaActive(Active_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSignupsViaActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	ActivationKey_ := self.Args("activation_key").String()
	var iSignups model.Signups
	self.Bind(&iSignups)
	var iMap = helper.StructToMap(iSignups)
	_error := model.UpdateSignupsViaActivationKey(ActivationKey_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateSignupsViaMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Meta_ := self.Args("meta").String()
	var iSignups model.Signups
	self.Bind(&iSignups)
	var iMap = helper.StructToMap(iSignups)
	_error := model.UpdateSignupsViaMeta(Meta_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSignupsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	SignupId_ := self.Args("signup_id").MustInt64()
	_int64, _error := model.DeleteSignups(SignupId_)
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

func DeleteSignupsViaSignupIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	SignupId_ := self.Args("signup_id").MustInt64()
	_error := model.DeleteSignupsViaSignupId(SignupId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSignupsViaDomainHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Domain_ := self.Args("domain").String()
	_error := model.DeleteSignupsViaDomain(Domain_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSignupsViaPathHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Path_ := self.Args("path").String()
	_error := model.DeleteSignupsViaPath(Path_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSignupsViaTitleHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Title_ := self.Args("title").String()
	_error := model.DeleteSignupsViaTitle(Title_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSignupsViaUserLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserLogin_ := self.Args("user_login").String()
	_error := model.DeleteSignupsViaUserLogin(UserLogin_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSignupsViaUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserEmail_ := self.Args("user_email").String()
	_error := model.DeleteSignupsViaUserEmail(UserEmail_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSignupsViaRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Registered_ := self.Args("registered").Time()
	_error := model.DeleteSignupsViaRegistered(Registered_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSignupsViaActivatedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Activated_ := self.Args("activated").Time()
	_error := model.DeleteSignupsViaActivated(Activated_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSignupsViaActiveHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Active_ := self.Args("active").MustInt()
	_error := model.DeleteSignupsViaActive(Active_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSignupsViaActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	ActivationKey_ := self.Args("activation_key").String()
	_error := model.DeleteSignupsViaActivationKey(ActivationKey_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteSignupsViaMetaHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Meta_ := self.Args("meta").String()
	_error := model.DeleteSignupsViaMeta(Meta_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
