package handler

import (
	"github.com/insionng/zenpress/helper"
	"github.com/insionng/zenpress/model"
	"github.com/insionng/macross"
)

func GetUsersesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetUsersesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["userssCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetUsersesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersCountViaIdHandler(self *macross.Context) error {
	Id_ := self.Args("ID").MustInt64()
	_int64 := model.GetUsersCountViaId(Id_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["usersCount"] = 0
	}
	m["usersCount"] = _int64
	return self.JSON(m)
}

func GetUsersCountViaUserLoginHandler(self *macross.Context) error {
	UserLogin_ := self.Args("user_login").String()
	_int64 := model.GetUsersCountViaUserLogin(UserLogin_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["usersCount"] = 0
	}
	m["usersCount"] = _int64
	return self.JSON(m)
}

func GetUsersCountViaUserPassHandler(self *macross.Context) error {
	UserPass_ := self.Args("user_pass").String()
	_int64 := model.GetUsersCountViaUserPass(UserPass_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["usersCount"] = 0
	}
	m["usersCount"] = _int64
	return self.JSON(m)
}

func GetUsersCountViaUserNicenameHandler(self *macross.Context) error {
	UserNicename_ := self.Args("user_nicename").String()
	_int64 := model.GetUsersCountViaUserNicename(UserNicename_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["usersCount"] = 0
	}
	m["usersCount"] = _int64
	return self.JSON(m)
}

func GetUsersCountViaUserEmailHandler(self *macross.Context) error {
	UserEmail_ := self.Args("user_email").String()
	_int64 := model.GetUsersCountViaUserEmail(UserEmail_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["usersCount"] = 0
	}
	m["usersCount"] = _int64
	return self.JSON(m)
}

func GetUsersCountViaUserUrlHandler(self *macross.Context) error {
	UserUrl_ := self.Args("user_url").String()
	_int64 := model.GetUsersCountViaUserUrl(UserUrl_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["usersCount"] = 0
	}
	m["usersCount"] = _int64
	return self.JSON(m)
}

func GetUsersCountViaUserRegisteredHandler(self *macross.Context) error {
	UserRegistered_ := self.Args("user_registered").Time()
	_int64 := model.GetUsersCountViaUserRegistered(UserRegistered_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["usersCount"] = 0
	}
	m["usersCount"] = _int64
	return self.JSON(m)
}

func GetUsersCountViaUserActivationKeyHandler(self *macross.Context) error {
	UserActivationKey_ := self.Args("user_activation_key").String()
	_int64 := model.GetUsersCountViaUserActivationKey(UserActivationKey_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["usersCount"] = 0
	}
	m["usersCount"] = _int64
	return self.JSON(m)
}

func GetUsersCountViaUserStatusHandler(self *macross.Context) error {
	UserStatus_ := self.Args("user_status").MustInt()
	_int64 := model.GetUsersCountViaUserStatus(UserStatus_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["usersCount"] = 0
	}
	m["usersCount"] = _int64
	return self.JSON(m)
}

func GetUsersCountViaDisplayNameHandler(self *macross.Context) error {
	DisplayName_ := self.Args("display_name").String()
	_int64 := model.GetUsersCountViaDisplayName(DisplayName_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["usersCount"] = 0
	}
	m["usersCount"] = _int64
	return self.JSON(m)
}

func GetUsersCountViaSpamHandler(self *macross.Context) error {
	Spam_ := self.Args("spam").MustInt()
	_int64 := model.GetUsersCountViaSpam(Spam_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["usersCount"] = 0
	}
	m["usersCount"] = _int64
	return self.JSON(m)
}

func GetUsersCountViaDeletedHandler(self *macross.Context) error {
	Deleted_ := self.Args("deleted").MustInt()
	_int64 := model.GetUsersCountViaDeleted(Deleted_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["usersCount"] = 0
	}
	m["usersCount"] = _int64
	return self.JSON(m)
}

func GetUsersesViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iId := self.Args("ID").MustInt64()
	if (offset > 0) && helper.IsHas(iId) {
		_Users, _error := model.GetUsersesViaId(offset, limit, iId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesViaUserLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUserLogin := self.Args("user_login").String()
	if (offset > 0) && helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesViaUserLogin(offset, limit, iUserLogin, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesViaUserLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesViaUserPassHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUserPass := self.Args("user_pass").String()
	if (offset > 0) && helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesViaUserPass(offset, limit, iUserPass, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesViaUserPass's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesViaUserNicenameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUserNicename := self.Args("user_nicename").String()
	if (offset > 0) && helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesViaUserNicename(offset, limit, iUserNicename, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesViaUserNicename's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesViaUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUserEmail := self.Args("user_email").String()
	if (offset > 0) && helper.IsHas(iUserEmail) {
		_Users, _error := model.GetUsersesViaUserEmail(offset, limit, iUserEmail, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesViaUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesViaUserUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUserUrl := self.Args("user_url").String()
	if (offset > 0) && helper.IsHas(iUserUrl) {
		_Users, _error := model.GetUsersesViaUserUrl(offset, limit, iUserUrl, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesViaUserUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesViaUserRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUserRegistered := self.Args("user_registered").Time()
	if (offset > 0) && helper.IsHas(iUserRegistered) {
		_Users, _error := model.GetUsersesViaUserRegistered(offset, limit, iUserRegistered, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesViaUserRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesViaUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUserActivationKey := self.Args("user_activation_key").String()
	if (offset > 0) && helper.IsHas(iUserActivationKey) {
		_Users, _error := model.GetUsersesViaUserActivationKey(offset, limit, iUserActivationKey, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesViaUserActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesViaUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iUserStatus := self.Args("user_status").MustInt()
	if (offset > 0) && helper.IsHas(iUserStatus) {
		_Users, _error := model.GetUsersesViaUserStatus(offset, limit, iUserStatus, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesViaUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesViaDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDisplayName := self.Args("display_name").String()
	if (offset > 0) && helper.IsHas(iDisplayName) {
		_Users, _error := model.GetUsersesViaDisplayName(offset, limit, iDisplayName, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesViaDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesViaSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iSpam := self.Args("spam").MustInt()
	if (offset > 0) && helper.IsHas(iSpam) {
		_Users, _error := model.GetUsersesViaSpam(offset, limit, iSpam, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesViaSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDeleted := self.Args("deleted").MustInt()
	if (offset > 0) && helper.IsHas(iDeleted) {
		_Users, _error := model.GetUsersesViaDeleted(offset, limit, iDeleted, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserLoginAndUserPassHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserLogin := self.Args("user_login").String()
	iUserPass := self.Args("user_pass").String()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserLoginAndUserPass(offset, limit, iId,iUserLogin,iUserPass)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserLoginAndUserPass's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserLoginAndUserNicenameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserLogin := self.Args("user_login").String()
	iUserNicename := self.Args("user_nicename").String()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserLoginAndUserNicename(offset, limit, iId,iUserLogin,iUserNicename)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserLoginAndUserNicename's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserLoginAndUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserLogin := self.Args("user_login").String()
	iUserEmail := self.Args("user_email").String()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserLoginAndUserEmail(offset, limit, iId,iUserLogin,iUserEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserLoginAndUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserLoginAndUserUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserLogin := self.Args("user_login").String()
	iUserUrl := self.Args("user_url").String()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserLoginAndUserUrl(offset, limit, iId,iUserLogin,iUserUrl)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserLoginAndUserUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserLoginAndUserRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserLogin := self.Args("user_login").String()
	iUserRegistered := self.Args("user_registered").Time()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserLoginAndUserRegistered(offset, limit, iId,iUserLogin,iUserRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserLoginAndUserRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserLoginAndUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserLogin := self.Args("user_login").String()
	iUserActivationKey := self.Args("user_activation_key").String()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserLoginAndUserActivationKey(offset, limit, iId,iUserLogin,iUserActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserLoginAndUserActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserLoginAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserLogin := self.Args("user_login").String()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserLoginAndUserStatus(offset, limit, iId,iUserLogin,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserLoginAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserLoginAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserLogin := self.Args("user_login").String()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserLoginAndDisplayName(offset, limit, iId,iUserLogin,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserLoginAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserLoginAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserLogin := self.Args("user_login").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserLoginAndSpam(offset, limit, iId,iUserLogin,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserLoginAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserLoginAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserLogin := self.Args("user_login").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserLoginAndDeleted(offset, limit, iId,iUserLogin,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserLoginAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserPassAndUserNicenameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserPass := self.Args("user_pass").String()
	iUserNicename := self.Args("user_nicename").String()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserPassAndUserNicename(offset, limit, iId,iUserPass,iUserNicename)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserPassAndUserNicename's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserPassAndUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserPass := self.Args("user_pass").String()
	iUserEmail := self.Args("user_email").String()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserPassAndUserEmail(offset, limit, iId,iUserPass,iUserEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserPassAndUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserPassAndUserUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserPass := self.Args("user_pass").String()
	iUserUrl := self.Args("user_url").String()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserPassAndUserUrl(offset, limit, iId,iUserPass,iUserUrl)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserPassAndUserUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserPassAndUserRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserPass := self.Args("user_pass").String()
	iUserRegistered := self.Args("user_registered").Time()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserPassAndUserRegistered(offset, limit, iId,iUserPass,iUserRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserPassAndUserRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserPassAndUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserPass := self.Args("user_pass").String()
	iUserActivationKey := self.Args("user_activation_key").String()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserPassAndUserActivationKey(offset, limit, iId,iUserPass,iUserActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserPassAndUserActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserPassAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserPass := self.Args("user_pass").String()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserPassAndUserStatus(offset, limit, iId,iUserPass,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserPassAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserPassAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserPass := self.Args("user_pass").String()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserPassAndDisplayName(offset, limit, iId,iUserPass,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserPassAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserPassAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserPass := self.Args("user_pass").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserPassAndSpam(offset, limit, iId,iUserPass,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserPassAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserPassAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserPass := self.Args("user_pass").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserPassAndDeleted(offset, limit, iId,iUserPass,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserPassAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserNicenameAndUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserNicename := self.Args("user_nicename").String()
	iUserEmail := self.Args("user_email").String()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserNicenameAndUserEmail(offset, limit, iId,iUserNicename,iUserEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserNicenameAndUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserNicenameAndUserUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserNicename := self.Args("user_nicename").String()
	iUserUrl := self.Args("user_url").String()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserNicenameAndUserUrl(offset, limit, iId,iUserNicename,iUserUrl)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserNicenameAndUserUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserNicenameAndUserRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserNicename := self.Args("user_nicename").String()
	iUserRegistered := self.Args("user_registered").Time()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserNicenameAndUserRegistered(offset, limit, iId,iUserNicename,iUserRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserNicenameAndUserRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserNicenameAndUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserNicename := self.Args("user_nicename").String()
	iUserActivationKey := self.Args("user_activation_key").String()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserNicenameAndUserActivationKey(offset, limit, iId,iUserNicename,iUserActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserNicenameAndUserActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserNicenameAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserNicename := self.Args("user_nicename").String()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserNicenameAndUserStatus(offset, limit, iId,iUserNicename,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserNicenameAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserNicenameAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserNicename := self.Args("user_nicename").String()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserNicenameAndDisplayName(offset, limit, iId,iUserNicename,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserNicenameAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserNicenameAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserNicename := self.Args("user_nicename").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserNicenameAndSpam(offset, limit, iId,iUserNicename,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserNicenameAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserNicenameAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserNicename := self.Args("user_nicename").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserNicenameAndDeleted(offset, limit, iId,iUserNicename,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserNicenameAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserEmailAndUserUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserEmail := self.Args("user_email").String()
	iUserUrl := self.Args("user_url").String()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserEmailAndUserUrl(offset, limit, iId,iUserEmail,iUserUrl)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserEmailAndUserUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserEmailAndUserRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserEmail := self.Args("user_email").String()
	iUserRegistered := self.Args("user_registered").Time()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserEmailAndUserRegistered(offset, limit, iId,iUserEmail,iUserRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserEmailAndUserRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserEmailAndUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserEmail := self.Args("user_email").String()
	iUserActivationKey := self.Args("user_activation_key").String()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserEmailAndUserActivationKey(offset, limit, iId,iUserEmail,iUserActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserEmailAndUserActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserEmailAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserEmail := self.Args("user_email").String()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserEmailAndUserStatus(offset, limit, iId,iUserEmail,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserEmailAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserEmailAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserEmail := self.Args("user_email").String()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserEmailAndDisplayName(offset, limit, iId,iUserEmail,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserEmailAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserEmailAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserEmail := self.Args("user_email").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserEmailAndSpam(offset, limit, iId,iUserEmail,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserEmailAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserEmailAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserEmail := self.Args("user_email").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserEmailAndDeleted(offset, limit, iId,iUserEmail,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserEmailAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserUrlAndUserRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserUrl := self.Args("user_url").String()
	iUserRegistered := self.Args("user_registered").Time()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserUrlAndUserRegistered(offset, limit, iId,iUserUrl,iUserRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserUrlAndUserRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserUrlAndUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserUrl := self.Args("user_url").String()
	iUserActivationKey := self.Args("user_activation_key").String()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserUrlAndUserActivationKey(offset, limit, iId,iUserUrl,iUserActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserUrlAndUserActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserUrlAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserUrl := self.Args("user_url").String()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserUrlAndUserStatus(offset, limit, iId,iUserUrl,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserUrlAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserUrlAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserUrl := self.Args("user_url").String()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserUrlAndDisplayName(offset, limit, iId,iUserUrl,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserUrlAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserUrlAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserUrl := self.Args("user_url").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserUrlAndSpam(offset, limit, iId,iUserUrl,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserUrlAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserUrlAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserUrl := self.Args("user_url").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserUrlAndDeleted(offset, limit, iId,iUserUrl,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserUrlAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserRegisteredAndUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserRegistered := self.Args("user_registered").Time()
	iUserActivationKey := self.Args("user_activation_key").String()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserRegisteredAndUserActivationKey(offset, limit, iId,iUserRegistered,iUserActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserRegisteredAndUserActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserRegisteredAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserRegistered := self.Args("user_registered").Time()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserRegisteredAndUserStatus(offset, limit, iId,iUserRegistered,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserRegisteredAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserRegisteredAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserRegistered := self.Args("user_registered").Time()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserRegisteredAndDisplayName(offset, limit, iId,iUserRegistered,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserRegisteredAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserRegisteredAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserRegistered := self.Args("user_registered").Time()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserRegisteredAndSpam(offset, limit, iId,iUserRegistered,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserRegisteredAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserRegisteredAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserRegistered := self.Args("user_registered").Time()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserRegisteredAndDeleted(offset, limit, iId,iUserRegistered,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserRegisteredAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserActivationKeyAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserActivationKey := self.Args("user_activation_key").String()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserActivationKeyAndUserStatus(offset, limit, iId,iUserActivationKey,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserActivationKeyAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserActivationKeyAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserActivationKey := self.Args("user_activation_key").String()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserActivationKeyAndDisplayName(offset, limit, iId,iUserActivationKey,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserActivationKeyAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserActivationKeyAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserActivationKey := self.Args("user_activation_key").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserActivationKeyAndSpam(offset, limit, iId,iUserActivationKey,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserActivationKeyAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserActivationKeyAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserActivationKey := self.Args("user_activation_key").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserActivationKeyAndDeleted(offset, limit, iId,iUserActivationKey,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserActivationKeyAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserStatusAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserStatus := self.Args("user_status").MustInt()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserStatusAndDisplayName(offset, limit, iId,iUserStatus,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserStatusAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserStatusAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserStatus := self.Args("user_status").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserStatusAndSpam(offset, limit, iId,iUserStatus,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserStatusAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserStatusAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserStatus := self.Args("user_status").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserStatusAndDeleted(offset, limit, iId,iUserStatus,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserStatusAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndDisplayNameAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iDisplayName := self.Args("display_name").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndDisplayNameAndSpam(offset, limit, iId,iDisplayName,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndDisplayNameAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndDisplayNameAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iDisplayName := self.Args("display_name").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndDisplayNameAndDeleted(offset, limit, iId,iDisplayName,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndDisplayNameAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndSpamAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iSpam := self.Args("spam").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndSpamAndDeleted(offset, limit, iId,iSpam,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndSpamAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserPassAndUserNicenameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserPass := self.Args("user_pass").String()
	iUserNicename := self.Args("user_nicename").String()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserPassAndUserNicename(offset, limit, iUserLogin,iUserPass,iUserNicename)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserPassAndUserNicename's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserPassAndUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserPass := self.Args("user_pass").String()
	iUserEmail := self.Args("user_email").String()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserPassAndUserEmail(offset, limit, iUserLogin,iUserPass,iUserEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserPassAndUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserPassAndUserUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserPass := self.Args("user_pass").String()
	iUserUrl := self.Args("user_url").String()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserPassAndUserUrl(offset, limit, iUserLogin,iUserPass,iUserUrl)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserPassAndUserUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserPassAndUserRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserPass := self.Args("user_pass").String()
	iUserRegistered := self.Args("user_registered").Time()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserPassAndUserRegistered(offset, limit, iUserLogin,iUserPass,iUserRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserPassAndUserRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserPassAndUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserPass := self.Args("user_pass").String()
	iUserActivationKey := self.Args("user_activation_key").String()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserPassAndUserActivationKey(offset, limit, iUserLogin,iUserPass,iUserActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserPassAndUserActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserPassAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserPass := self.Args("user_pass").String()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserPassAndUserStatus(offset, limit, iUserLogin,iUserPass,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserPassAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserPassAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserPass := self.Args("user_pass").String()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserPassAndDisplayName(offset, limit, iUserLogin,iUserPass,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserPassAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserPassAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserPass := self.Args("user_pass").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserPassAndSpam(offset, limit, iUserLogin,iUserPass,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserPassAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserPassAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserPass := self.Args("user_pass").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserPassAndDeleted(offset, limit, iUserLogin,iUserPass,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserPassAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserNicenameAndUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserNicename := self.Args("user_nicename").String()
	iUserEmail := self.Args("user_email").String()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserNicenameAndUserEmail(offset, limit, iUserLogin,iUserNicename,iUserEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserNicenameAndUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserNicenameAndUserUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserNicename := self.Args("user_nicename").String()
	iUserUrl := self.Args("user_url").String()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserNicenameAndUserUrl(offset, limit, iUserLogin,iUserNicename,iUserUrl)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserNicenameAndUserUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserNicenameAndUserRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserNicename := self.Args("user_nicename").String()
	iUserRegistered := self.Args("user_registered").Time()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserNicenameAndUserRegistered(offset, limit, iUserLogin,iUserNicename,iUserRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserNicenameAndUserRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserNicenameAndUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserNicename := self.Args("user_nicename").String()
	iUserActivationKey := self.Args("user_activation_key").String()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserNicenameAndUserActivationKey(offset, limit, iUserLogin,iUserNicename,iUserActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserNicenameAndUserActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserNicenameAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserNicename := self.Args("user_nicename").String()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserNicenameAndUserStatus(offset, limit, iUserLogin,iUserNicename,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserNicenameAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserNicenameAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserNicename := self.Args("user_nicename").String()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserNicenameAndDisplayName(offset, limit, iUserLogin,iUserNicename,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserNicenameAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserNicenameAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserNicename := self.Args("user_nicename").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserNicenameAndSpam(offset, limit, iUserLogin,iUserNicename,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserNicenameAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserNicenameAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserNicename := self.Args("user_nicename").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserNicenameAndDeleted(offset, limit, iUserLogin,iUserNicename,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserNicenameAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserEmailAndUserUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserEmail := self.Args("user_email").String()
	iUserUrl := self.Args("user_url").String()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserEmailAndUserUrl(offset, limit, iUserLogin,iUserEmail,iUserUrl)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserEmailAndUserUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserEmailAndUserRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserEmail := self.Args("user_email").String()
	iUserRegistered := self.Args("user_registered").Time()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserEmailAndUserRegistered(offset, limit, iUserLogin,iUserEmail,iUserRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserEmailAndUserRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserEmailAndUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserEmail := self.Args("user_email").String()
	iUserActivationKey := self.Args("user_activation_key").String()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserEmailAndUserActivationKey(offset, limit, iUserLogin,iUserEmail,iUserActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserEmailAndUserActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserEmailAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserEmail := self.Args("user_email").String()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserEmailAndUserStatus(offset, limit, iUserLogin,iUserEmail,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserEmailAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserEmailAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserEmail := self.Args("user_email").String()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserEmailAndDisplayName(offset, limit, iUserLogin,iUserEmail,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserEmailAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserEmailAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserEmail := self.Args("user_email").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserEmailAndSpam(offset, limit, iUserLogin,iUserEmail,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserEmailAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserEmailAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserEmail := self.Args("user_email").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserEmailAndDeleted(offset, limit, iUserLogin,iUserEmail,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserEmailAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserUrlAndUserRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserUrl := self.Args("user_url").String()
	iUserRegistered := self.Args("user_registered").Time()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserUrlAndUserRegistered(offset, limit, iUserLogin,iUserUrl,iUserRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserUrlAndUserRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserUrlAndUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserUrl := self.Args("user_url").String()
	iUserActivationKey := self.Args("user_activation_key").String()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserUrlAndUserActivationKey(offset, limit, iUserLogin,iUserUrl,iUserActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserUrlAndUserActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserUrlAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserUrl := self.Args("user_url").String()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserUrlAndUserStatus(offset, limit, iUserLogin,iUserUrl,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserUrlAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserUrlAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserUrl := self.Args("user_url").String()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserUrlAndDisplayName(offset, limit, iUserLogin,iUserUrl,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserUrlAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserUrlAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserUrl := self.Args("user_url").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserUrlAndSpam(offset, limit, iUserLogin,iUserUrl,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserUrlAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserUrlAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserUrl := self.Args("user_url").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserUrlAndDeleted(offset, limit, iUserLogin,iUserUrl,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserUrlAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserRegisteredAndUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserRegistered := self.Args("user_registered").Time()
	iUserActivationKey := self.Args("user_activation_key").String()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserRegisteredAndUserActivationKey(offset, limit, iUserLogin,iUserRegistered,iUserActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserRegisteredAndUserActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserRegisteredAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserRegistered := self.Args("user_registered").Time()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserRegisteredAndUserStatus(offset, limit, iUserLogin,iUserRegistered,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserRegisteredAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserRegisteredAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserRegistered := self.Args("user_registered").Time()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserRegisteredAndDisplayName(offset, limit, iUserLogin,iUserRegistered,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserRegisteredAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserRegisteredAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserRegistered := self.Args("user_registered").Time()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserRegisteredAndSpam(offset, limit, iUserLogin,iUserRegistered,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserRegisteredAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserRegisteredAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserRegistered := self.Args("user_registered").Time()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserRegisteredAndDeleted(offset, limit, iUserLogin,iUserRegistered,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserRegisteredAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserActivationKeyAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserActivationKey := self.Args("user_activation_key").String()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserActivationKeyAndUserStatus(offset, limit, iUserLogin,iUserActivationKey,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserActivationKeyAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserActivationKeyAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserActivationKey := self.Args("user_activation_key").String()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserActivationKeyAndDisplayName(offset, limit, iUserLogin,iUserActivationKey,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserActivationKeyAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserActivationKeyAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserActivationKey := self.Args("user_activation_key").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserActivationKeyAndSpam(offset, limit, iUserLogin,iUserActivationKey,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserActivationKeyAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserActivationKeyAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserActivationKey := self.Args("user_activation_key").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserActivationKeyAndDeleted(offset, limit, iUserLogin,iUserActivationKey,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserActivationKeyAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserStatusAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserStatus := self.Args("user_status").MustInt()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserStatusAndDisplayName(offset, limit, iUserLogin,iUserStatus,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserStatusAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserStatusAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserStatus := self.Args("user_status").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserStatusAndSpam(offset, limit, iUserLogin,iUserStatus,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserStatusAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserStatusAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserStatus := self.Args("user_status").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserStatusAndDeleted(offset, limit, iUserLogin,iUserStatus,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserStatusAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndDisplayNameAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iDisplayName := self.Args("display_name").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndDisplayNameAndSpam(offset, limit, iUserLogin,iDisplayName,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndDisplayNameAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndDisplayNameAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iDisplayName := self.Args("display_name").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndDisplayNameAndDeleted(offset, limit, iUserLogin,iDisplayName,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndDisplayNameAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndSpamAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iSpam := self.Args("spam").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndSpamAndDeleted(offset, limit, iUserLogin,iSpam,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndSpamAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserNicenameAndUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserNicename := self.Args("user_nicename").String()
	iUserEmail := self.Args("user_email").String()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserNicenameAndUserEmail(offset, limit, iUserPass,iUserNicename,iUserEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserNicenameAndUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserNicenameAndUserUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserNicename := self.Args("user_nicename").String()
	iUserUrl := self.Args("user_url").String()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserNicenameAndUserUrl(offset, limit, iUserPass,iUserNicename,iUserUrl)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserNicenameAndUserUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserNicenameAndUserRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserNicename := self.Args("user_nicename").String()
	iUserRegistered := self.Args("user_registered").Time()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserNicenameAndUserRegistered(offset, limit, iUserPass,iUserNicename,iUserRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserNicenameAndUserRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserNicenameAndUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserNicename := self.Args("user_nicename").String()
	iUserActivationKey := self.Args("user_activation_key").String()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserNicenameAndUserActivationKey(offset, limit, iUserPass,iUserNicename,iUserActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserNicenameAndUserActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserNicenameAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserNicename := self.Args("user_nicename").String()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserNicenameAndUserStatus(offset, limit, iUserPass,iUserNicename,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserNicenameAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserNicenameAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserNicename := self.Args("user_nicename").String()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserNicenameAndDisplayName(offset, limit, iUserPass,iUserNicename,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserNicenameAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserNicenameAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserNicename := self.Args("user_nicename").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserNicenameAndSpam(offset, limit, iUserPass,iUserNicename,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserNicenameAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserNicenameAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserNicename := self.Args("user_nicename").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserNicenameAndDeleted(offset, limit, iUserPass,iUserNicename,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserNicenameAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserEmailAndUserUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserEmail := self.Args("user_email").String()
	iUserUrl := self.Args("user_url").String()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserEmailAndUserUrl(offset, limit, iUserPass,iUserEmail,iUserUrl)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserEmailAndUserUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserEmailAndUserRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserEmail := self.Args("user_email").String()
	iUserRegistered := self.Args("user_registered").Time()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserEmailAndUserRegistered(offset, limit, iUserPass,iUserEmail,iUserRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserEmailAndUserRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserEmailAndUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserEmail := self.Args("user_email").String()
	iUserActivationKey := self.Args("user_activation_key").String()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserEmailAndUserActivationKey(offset, limit, iUserPass,iUserEmail,iUserActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserEmailAndUserActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserEmailAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserEmail := self.Args("user_email").String()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserEmailAndUserStatus(offset, limit, iUserPass,iUserEmail,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserEmailAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserEmailAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserEmail := self.Args("user_email").String()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserEmailAndDisplayName(offset, limit, iUserPass,iUserEmail,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserEmailAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserEmailAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserEmail := self.Args("user_email").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserEmailAndSpam(offset, limit, iUserPass,iUserEmail,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserEmailAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserEmailAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserEmail := self.Args("user_email").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserEmailAndDeleted(offset, limit, iUserPass,iUserEmail,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserEmailAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserUrlAndUserRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserUrl := self.Args("user_url").String()
	iUserRegistered := self.Args("user_registered").Time()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserUrlAndUserRegistered(offset, limit, iUserPass,iUserUrl,iUserRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserUrlAndUserRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserUrlAndUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserUrl := self.Args("user_url").String()
	iUserActivationKey := self.Args("user_activation_key").String()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserUrlAndUserActivationKey(offset, limit, iUserPass,iUserUrl,iUserActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserUrlAndUserActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserUrlAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserUrl := self.Args("user_url").String()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserUrlAndUserStatus(offset, limit, iUserPass,iUserUrl,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserUrlAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserUrlAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserUrl := self.Args("user_url").String()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserUrlAndDisplayName(offset, limit, iUserPass,iUserUrl,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserUrlAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserUrlAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserUrl := self.Args("user_url").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserUrlAndSpam(offset, limit, iUserPass,iUserUrl,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserUrlAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserUrlAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserUrl := self.Args("user_url").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserUrlAndDeleted(offset, limit, iUserPass,iUserUrl,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserUrlAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserRegisteredAndUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserRegistered := self.Args("user_registered").Time()
	iUserActivationKey := self.Args("user_activation_key").String()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserRegisteredAndUserActivationKey(offset, limit, iUserPass,iUserRegistered,iUserActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserRegisteredAndUserActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserRegisteredAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserRegistered := self.Args("user_registered").Time()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserRegisteredAndUserStatus(offset, limit, iUserPass,iUserRegistered,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserRegisteredAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserRegisteredAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserRegistered := self.Args("user_registered").Time()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserRegisteredAndDisplayName(offset, limit, iUserPass,iUserRegistered,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserRegisteredAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserRegisteredAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserRegistered := self.Args("user_registered").Time()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserRegisteredAndSpam(offset, limit, iUserPass,iUserRegistered,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserRegisteredAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserRegisteredAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserRegistered := self.Args("user_registered").Time()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserRegisteredAndDeleted(offset, limit, iUserPass,iUserRegistered,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserRegisteredAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserActivationKeyAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserActivationKey := self.Args("user_activation_key").String()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserActivationKeyAndUserStatus(offset, limit, iUserPass,iUserActivationKey,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserActivationKeyAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserActivationKeyAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserActivationKey := self.Args("user_activation_key").String()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserActivationKeyAndDisplayName(offset, limit, iUserPass,iUserActivationKey,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserActivationKeyAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserActivationKeyAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserActivationKey := self.Args("user_activation_key").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserActivationKeyAndSpam(offset, limit, iUserPass,iUserActivationKey,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserActivationKeyAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserActivationKeyAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserActivationKey := self.Args("user_activation_key").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserActivationKeyAndDeleted(offset, limit, iUserPass,iUserActivationKey,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserActivationKeyAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserStatusAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserStatus := self.Args("user_status").MustInt()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserStatusAndDisplayName(offset, limit, iUserPass,iUserStatus,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserStatusAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserStatusAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserStatus := self.Args("user_status").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserStatusAndSpam(offset, limit, iUserPass,iUserStatus,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserStatusAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserStatusAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserStatus := self.Args("user_status").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserStatusAndDeleted(offset, limit, iUserPass,iUserStatus,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserStatusAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndDisplayNameAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iDisplayName := self.Args("display_name").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndDisplayNameAndSpam(offset, limit, iUserPass,iDisplayName,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndDisplayNameAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndDisplayNameAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iDisplayName := self.Args("display_name").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndDisplayNameAndDeleted(offset, limit, iUserPass,iDisplayName,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndDisplayNameAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndSpamAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iSpam := self.Args("spam").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndSpamAndDeleted(offset, limit, iUserPass,iSpam,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndSpamAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndUserEmailAndUserUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iUserEmail := self.Args("user_email").String()
	iUserUrl := self.Args("user_url").String()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndUserEmailAndUserUrl(offset, limit, iUserNicename,iUserEmail,iUserUrl)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndUserEmailAndUserUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndUserEmailAndUserRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iUserEmail := self.Args("user_email").String()
	iUserRegistered := self.Args("user_registered").Time()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndUserEmailAndUserRegistered(offset, limit, iUserNicename,iUserEmail,iUserRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndUserEmailAndUserRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndUserEmailAndUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iUserEmail := self.Args("user_email").String()
	iUserActivationKey := self.Args("user_activation_key").String()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndUserEmailAndUserActivationKey(offset, limit, iUserNicename,iUserEmail,iUserActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndUserEmailAndUserActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndUserEmailAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iUserEmail := self.Args("user_email").String()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndUserEmailAndUserStatus(offset, limit, iUserNicename,iUserEmail,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndUserEmailAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndUserEmailAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iUserEmail := self.Args("user_email").String()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndUserEmailAndDisplayName(offset, limit, iUserNicename,iUserEmail,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndUserEmailAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndUserEmailAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iUserEmail := self.Args("user_email").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndUserEmailAndSpam(offset, limit, iUserNicename,iUserEmail,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndUserEmailAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndUserEmailAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iUserEmail := self.Args("user_email").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndUserEmailAndDeleted(offset, limit, iUserNicename,iUserEmail,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndUserEmailAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndUserUrlAndUserRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iUserUrl := self.Args("user_url").String()
	iUserRegistered := self.Args("user_registered").Time()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndUserUrlAndUserRegistered(offset, limit, iUserNicename,iUserUrl,iUserRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndUserUrlAndUserRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndUserUrlAndUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iUserUrl := self.Args("user_url").String()
	iUserActivationKey := self.Args("user_activation_key").String()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndUserUrlAndUserActivationKey(offset, limit, iUserNicename,iUserUrl,iUserActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndUserUrlAndUserActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndUserUrlAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iUserUrl := self.Args("user_url").String()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndUserUrlAndUserStatus(offset, limit, iUserNicename,iUserUrl,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndUserUrlAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndUserUrlAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iUserUrl := self.Args("user_url").String()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndUserUrlAndDisplayName(offset, limit, iUserNicename,iUserUrl,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndUserUrlAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndUserUrlAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iUserUrl := self.Args("user_url").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndUserUrlAndSpam(offset, limit, iUserNicename,iUserUrl,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndUserUrlAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndUserUrlAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iUserUrl := self.Args("user_url").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndUserUrlAndDeleted(offset, limit, iUserNicename,iUserUrl,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndUserUrlAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndUserRegisteredAndUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iUserRegistered := self.Args("user_registered").Time()
	iUserActivationKey := self.Args("user_activation_key").String()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndUserRegisteredAndUserActivationKey(offset, limit, iUserNicename,iUserRegistered,iUserActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndUserRegisteredAndUserActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndUserRegisteredAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iUserRegistered := self.Args("user_registered").Time()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndUserRegisteredAndUserStatus(offset, limit, iUserNicename,iUserRegistered,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndUserRegisteredAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndUserRegisteredAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iUserRegistered := self.Args("user_registered").Time()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndUserRegisteredAndDisplayName(offset, limit, iUserNicename,iUserRegistered,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndUserRegisteredAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndUserRegisteredAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iUserRegistered := self.Args("user_registered").Time()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndUserRegisteredAndSpam(offset, limit, iUserNicename,iUserRegistered,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndUserRegisteredAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndUserRegisteredAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iUserRegistered := self.Args("user_registered").Time()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndUserRegisteredAndDeleted(offset, limit, iUserNicename,iUserRegistered,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndUserRegisteredAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndUserActivationKeyAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iUserActivationKey := self.Args("user_activation_key").String()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndUserActivationKeyAndUserStatus(offset, limit, iUserNicename,iUserActivationKey,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndUserActivationKeyAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndUserActivationKeyAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iUserActivationKey := self.Args("user_activation_key").String()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndUserActivationKeyAndDisplayName(offset, limit, iUserNicename,iUserActivationKey,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndUserActivationKeyAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndUserActivationKeyAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iUserActivationKey := self.Args("user_activation_key").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndUserActivationKeyAndSpam(offset, limit, iUserNicename,iUserActivationKey,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndUserActivationKeyAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndUserActivationKeyAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iUserActivationKey := self.Args("user_activation_key").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndUserActivationKeyAndDeleted(offset, limit, iUserNicename,iUserActivationKey,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndUserActivationKeyAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndUserStatusAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iUserStatus := self.Args("user_status").MustInt()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndUserStatusAndDisplayName(offset, limit, iUserNicename,iUserStatus,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndUserStatusAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndUserStatusAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iUserStatus := self.Args("user_status").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndUserStatusAndSpam(offset, limit, iUserNicename,iUserStatus,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndUserStatusAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndUserStatusAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iUserStatus := self.Args("user_status").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndUserStatusAndDeleted(offset, limit, iUserNicename,iUserStatus,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndUserStatusAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndDisplayNameAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iDisplayName := self.Args("display_name").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndDisplayNameAndSpam(offset, limit, iUserNicename,iDisplayName,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndDisplayNameAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndDisplayNameAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iDisplayName := self.Args("display_name").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndDisplayNameAndDeleted(offset, limit, iUserNicename,iDisplayName,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndDisplayNameAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndSpamAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iSpam := self.Args("spam").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndSpamAndDeleted(offset, limit, iUserNicename,iSpam,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndSpamAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserEmailAndUserUrlAndUserRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iUserUrl := self.Args("user_url").String()
	iUserRegistered := self.Args("user_registered").Time()

	if helper.IsHas(iUserEmail) {
		_Users, _error := model.GetUsersesByUserEmailAndUserUrlAndUserRegistered(offset, limit, iUserEmail,iUserUrl,iUserRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserEmailAndUserUrlAndUserRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserEmailAndUserUrlAndUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iUserUrl := self.Args("user_url").String()
	iUserActivationKey := self.Args("user_activation_key").String()

	if helper.IsHas(iUserEmail) {
		_Users, _error := model.GetUsersesByUserEmailAndUserUrlAndUserActivationKey(offset, limit, iUserEmail,iUserUrl,iUserActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserEmailAndUserUrlAndUserActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserEmailAndUserUrlAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iUserUrl := self.Args("user_url").String()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iUserEmail) {
		_Users, _error := model.GetUsersesByUserEmailAndUserUrlAndUserStatus(offset, limit, iUserEmail,iUserUrl,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserEmailAndUserUrlAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserEmailAndUserUrlAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iUserUrl := self.Args("user_url").String()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserEmail) {
		_Users, _error := model.GetUsersesByUserEmailAndUserUrlAndDisplayName(offset, limit, iUserEmail,iUserUrl,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserEmailAndUserUrlAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserEmailAndUserUrlAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iUserUrl := self.Args("user_url").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserEmail) {
		_Users, _error := model.GetUsersesByUserEmailAndUserUrlAndSpam(offset, limit, iUserEmail,iUserUrl,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserEmailAndUserUrlAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserEmailAndUserUrlAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iUserUrl := self.Args("user_url").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserEmail) {
		_Users, _error := model.GetUsersesByUserEmailAndUserUrlAndDeleted(offset, limit, iUserEmail,iUserUrl,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserEmailAndUserUrlAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserEmailAndUserRegisteredAndUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iUserRegistered := self.Args("user_registered").Time()
	iUserActivationKey := self.Args("user_activation_key").String()

	if helper.IsHas(iUserEmail) {
		_Users, _error := model.GetUsersesByUserEmailAndUserRegisteredAndUserActivationKey(offset, limit, iUserEmail,iUserRegistered,iUserActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserEmailAndUserRegisteredAndUserActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserEmailAndUserRegisteredAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iUserRegistered := self.Args("user_registered").Time()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iUserEmail) {
		_Users, _error := model.GetUsersesByUserEmailAndUserRegisteredAndUserStatus(offset, limit, iUserEmail,iUserRegistered,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserEmailAndUserRegisteredAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserEmailAndUserRegisteredAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iUserRegistered := self.Args("user_registered").Time()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserEmail) {
		_Users, _error := model.GetUsersesByUserEmailAndUserRegisteredAndDisplayName(offset, limit, iUserEmail,iUserRegistered,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserEmailAndUserRegisteredAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserEmailAndUserRegisteredAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iUserRegistered := self.Args("user_registered").Time()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserEmail) {
		_Users, _error := model.GetUsersesByUserEmailAndUserRegisteredAndSpam(offset, limit, iUserEmail,iUserRegistered,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserEmailAndUserRegisteredAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserEmailAndUserRegisteredAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iUserRegistered := self.Args("user_registered").Time()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserEmail) {
		_Users, _error := model.GetUsersesByUserEmailAndUserRegisteredAndDeleted(offset, limit, iUserEmail,iUserRegistered,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserEmailAndUserRegisteredAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserEmailAndUserActivationKeyAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iUserActivationKey := self.Args("user_activation_key").String()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iUserEmail) {
		_Users, _error := model.GetUsersesByUserEmailAndUserActivationKeyAndUserStatus(offset, limit, iUserEmail,iUserActivationKey,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserEmailAndUserActivationKeyAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserEmailAndUserActivationKeyAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iUserActivationKey := self.Args("user_activation_key").String()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserEmail) {
		_Users, _error := model.GetUsersesByUserEmailAndUserActivationKeyAndDisplayName(offset, limit, iUserEmail,iUserActivationKey,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserEmailAndUserActivationKeyAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserEmailAndUserActivationKeyAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iUserActivationKey := self.Args("user_activation_key").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserEmail) {
		_Users, _error := model.GetUsersesByUserEmailAndUserActivationKeyAndSpam(offset, limit, iUserEmail,iUserActivationKey,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserEmailAndUserActivationKeyAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserEmailAndUserActivationKeyAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iUserActivationKey := self.Args("user_activation_key").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserEmail) {
		_Users, _error := model.GetUsersesByUserEmailAndUserActivationKeyAndDeleted(offset, limit, iUserEmail,iUserActivationKey,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserEmailAndUserActivationKeyAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserEmailAndUserStatusAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iUserStatus := self.Args("user_status").MustInt()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserEmail) {
		_Users, _error := model.GetUsersesByUserEmailAndUserStatusAndDisplayName(offset, limit, iUserEmail,iUserStatus,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserEmailAndUserStatusAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserEmailAndUserStatusAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iUserStatus := self.Args("user_status").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserEmail) {
		_Users, _error := model.GetUsersesByUserEmailAndUserStatusAndSpam(offset, limit, iUserEmail,iUserStatus,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserEmailAndUserStatusAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserEmailAndUserStatusAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iUserStatus := self.Args("user_status").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserEmail) {
		_Users, _error := model.GetUsersesByUserEmailAndUserStatusAndDeleted(offset, limit, iUserEmail,iUserStatus,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserEmailAndUserStatusAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserEmailAndDisplayNameAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iDisplayName := self.Args("display_name").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserEmail) {
		_Users, _error := model.GetUsersesByUserEmailAndDisplayNameAndSpam(offset, limit, iUserEmail,iDisplayName,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserEmailAndDisplayNameAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserEmailAndDisplayNameAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iDisplayName := self.Args("display_name").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserEmail) {
		_Users, _error := model.GetUsersesByUserEmailAndDisplayNameAndDeleted(offset, limit, iUserEmail,iDisplayName,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserEmailAndDisplayNameAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserEmailAndSpamAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iSpam := self.Args("spam").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserEmail) {
		_Users, _error := model.GetUsersesByUserEmailAndSpamAndDeleted(offset, limit, iUserEmail,iSpam,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserEmailAndSpamAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserUrlAndUserRegisteredAndUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserUrl := self.Args("user_url").String()
	iUserRegistered := self.Args("user_registered").Time()
	iUserActivationKey := self.Args("user_activation_key").String()

	if helper.IsHas(iUserUrl) {
		_Users, _error := model.GetUsersesByUserUrlAndUserRegisteredAndUserActivationKey(offset, limit, iUserUrl,iUserRegistered,iUserActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserUrlAndUserRegisteredAndUserActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserUrlAndUserRegisteredAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserUrl := self.Args("user_url").String()
	iUserRegistered := self.Args("user_registered").Time()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iUserUrl) {
		_Users, _error := model.GetUsersesByUserUrlAndUserRegisteredAndUserStatus(offset, limit, iUserUrl,iUserRegistered,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserUrlAndUserRegisteredAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserUrlAndUserRegisteredAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserUrl := self.Args("user_url").String()
	iUserRegistered := self.Args("user_registered").Time()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserUrl) {
		_Users, _error := model.GetUsersesByUserUrlAndUserRegisteredAndDisplayName(offset, limit, iUserUrl,iUserRegistered,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserUrlAndUserRegisteredAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserUrlAndUserRegisteredAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserUrl := self.Args("user_url").String()
	iUserRegistered := self.Args("user_registered").Time()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserUrl) {
		_Users, _error := model.GetUsersesByUserUrlAndUserRegisteredAndSpam(offset, limit, iUserUrl,iUserRegistered,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserUrlAndUserRegisteredAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserUrlAndUserRegisteredAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserUrl := self.Args("user_url").String()
	iUserRegistered := self.Args("user_registered").Time()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserUrl) {
		_Users, _error := model.GetUsersesByUserUrlAndUserRegisteredAndDeleted(offset, limit, iUserUrl,iUserRegistered,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserUrlAndUserRegisteredAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserUrlAndUserActivationKeyAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserUrl := self.Args("user_url").String()
	iUserActivationKey := self.Args("user_activation_key").String()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iUserUrl) {
		_Users, _error := model.GetUsersesByUserUrlAndUserActivationKeyAndUserStatus(offset, limit, iUserUrl,iUserActivationKey,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserUrlAndUserActivationKeyAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserUrlAndUserActivationKeyAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserUrl := self.Args("user_url").String()
	iUserActivationKey := self.Args("user_activation_key").String()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserUrl) {
		_Users, _error := model.GetUsersesByUserUrlAndUserActivationKeyAndDisplayName(offset, limit, iUserUrl,iUserActivationKey,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserUrlAndUserActivationKeyAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserUrlAndUserActivationKeyAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserUrl := self.Args("user_url").String()
	iUserActivationKey := self.Args("user_activation_key").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserUrl) {
		_Users, _error := model.GetUsersesByUserUrlAndUserActivationKeyAndSpam(offset, limit, iUserUrl,iUserActivationKey,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserUrlAndUserActivationKeyAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserUrlAndUserActivationKeyAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserUrl := self.Args("user_url").String()
	iUserActivationKey := self.Args("user_activation_key").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserUrl) {
		_Users, _error := model.GetUsersesByUserUrlAndUserActivationKeyAndDeleted(offset, limit, iUserUrl,iUserActivationKey,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserUrlAndUserActivationKeyAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserUrlAndUserStatusAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserUrl := self.Args("user_url").String()
	iUserStatus := self.Args("user_status").MustInt()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserUrl) {
		_Users, _error := model.GetUsersesByUserUrlAndUserStatusAndDisplayName(offset, limit, iUserUrl,iUserStatus,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserUrlAndUserStatusAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserUrlAndUserStatusAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserUrl := self.Args("user_url").String()
	iUserStatus := self.Args("user_status").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserUrl) {
		_Users, _error := model.GetUsersesByUserUrlAndUserStatusAndSpam(offset, limit, iUserUrl,iUserStatus,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserUrlAndUserStatusAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserUrlAndUserStatusAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserUrl := self.Args("user_url").String()
	iUserStatus := self.Args("user_status").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserUrl) {
		_Users, _error := model.GetUsersesByUserUrlAndUserStatusAndDeleted(offset, limit, iUserUrl,iUserStatus,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserUrlAndUserStatusAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserUrlAndDisplayNameAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserUrl := self.Args("user_url").String()
	iDisplayName := self.Args("display_name").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserUrl) {
		_Users, _error := model.GetUsersesByUserUrlAndDisplayNameAndSpam(offset, limit, iUserUrl,iDisplayName,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserUrlAndDisplayNameAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserUrlAndDisplayNameAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserUrl := self.Args("user_url").String()
	iDisplayName := self.Args("display_name").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserUrl) {
		_Users, _error := model.GetUsersesByUserUrlAndDisplayNameAndDeleted(offset, limit, iUserUrl,iDisplayName,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserUrlAndDisplayNameAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserUrlAndSpamAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserUrl := self.Args("user_url").String()
	iSpam := self.Args("spam").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserUrl) {
		_Users, _error := model.GetUsersesByUserUrlAndSpamAndDeleted(offset, limit, iUserUrl,iSpam,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserUrlAndSpamAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserRegisteredAndUserActivationKeyAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserRegistered := self.Args("user_registered").Time()
	iUserActivationKey := self.Args("user_activation_key").String()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iUserRegistered) {
		_Users, _error := model.GetUsersesByUserRegisteredAndUserActivationKeyAndUserStatus(offset, limit, iUserRegistered,iUserActivationKey,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserRegisteredAndUserActivationKeyAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserRegisteredAndUserActivationKeyAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserRegistered := self.Args("user_registered").Time()
	iUserActivationKey := self.Args("user_activation_key").String()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserRegistered) {
		_Users, _error := model.GetUsersesByUserRegisteredAndUserActivationKeyAndDisplayName(offset, limit, iUserRegistered,iUserActivationKey,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserRegisteredAndUserActivationKeyAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserRegisteredAndUserActivationKeyAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserRegistered := self.Args("user_registered").Time()
	iUserActivationKey := self.Args("user_activation_key").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserRegistered) {
		_Users, _error := model.GetUsersesByUserRegisteredAndUserActivationKeyAndSpam(offset, limit, iUserRegistered,iUserActivationKey,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserRegisteredAndUserActivationKeyAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserRegisteredAndUserActivationKeyAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserRegistered := self.Args("user_registered").Time()
	iUserActivationKey := self.Args("user_activation_key").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserRegistered) {
		_Users, _error := model.GetUsersesByUserRegisteredAndUserActivationKeyAndDeleted(offset, limit, iUserRegistered,iUserActivationKey,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserRegisteredAndUserActivationKeyAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserRegisteredAndUserStatusAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserRegistered := self.Args("user_registered").Time()
	iUserStatus := self.Args("user_status").MustInt()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserRegistered) {
		_Users, _error := model.GetUsersesByUserRegisteredAndUserStatusAndDisplayName(offset, limit, iUserRegistered,iUserStatus,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserRegisteredAndUserStatusAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserRegisteredAndUserStatusAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserRegistered := self.Args("user_registered").Time()
	iUserStatus := self.Args("user_status").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserRegistered) {
		_Users, _error := model.GetUsersesByUserRegisteredAndUserStatusAndSpam(offset, limit, iUserRegistered,iUserStatus,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserRegisteredAndUserStatusAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserRegisteredAndUserStatusAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserRegistered := self.Args("user_registered").Time()
	iUserStatus := self.Args("user_status").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserRegistered) {
		_Users, _error := model.GetUsersesByUserRegisteredAndUserStatusAndDeleted(offset, limit, iUserRegistered,iUserStatus,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserRegisteredAndUserStatusAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserRegisteredAndDisplayNameAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserRegistered := self.Args("user_registered").Time()
	iDisplayName := self.Args("display_name").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserRegistered) {
		_Users, _error := model.GetUsersesByUserRegisteredAndDisplayNameAndSpam(offset, limit, iUserRegistered,iDisplayName,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserRegisteredAndDisplayNameAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserRegisteredAndDisplayNameAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserRegistered := self.Args("user_registered").Time()
	iDisplayName := self.Args("display_name").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserRegistered) {
		_Users, _error := model.GetUsersesByUserRegisteredAndDisplayNameAndDeleted(offset, limit, iUserRegistered,iDisplayName,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserRegisteredAndDisplayNameAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserRegisteredAndSpamAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserRegistered := self.Args("user_registered").Time()
	iSpam := self.Args("spam").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserRegistered) {
		_Users, _error := model.GetUsersesByUserRegisteredAndSpamAndDeleted(offset, limit, iUserRegistered,iSpam,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserRegisteredAndSpamAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserActivationKeyAndUserStatusAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserActivationKey := self.Args("user_activation_key").String()
	iUserStatus := self.Args("user_status").MustInt()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserActivationKey) {
		_Users, _error := model.GetUsersesByUserActivationKeyAndUserStatusAndDisplayName(offset, limit, iUserActivationKey,iUserStatus,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserActivationKeyAndUserStatusAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserActivationKeyAndUserStatusAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserActivationKey := self.Args("user_activation_key").String()
	iUserStatus := self.Args("user_status").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserActivationKey) {
		_Users, _error := model.GetUsersesByUserActivationKeyAndUserStatusAndSpam(offset, limit, iUserActivationKey,iUserStatus,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserActivationKeyAndUserStatusAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserActivationKeyAndUserStatusAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserActivationKey := self.Args("user_activation_key").String()
	iUserStatus := self.Args("user_status").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserActivationKey) {
		_Users, _error := model.GetUsersesByUserActivationKeyAndUserStatusAndDeleted(offset, limit, iUserActivationKey,iUserStatus,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserActivationKeyAndUserStatusAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserActivationKeyAndDisplayNameAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserActivationKey := self.Args("user_activation_key").String()
	iDisplayName := self.Args("display_name").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserActivationKey) {
		_Users, _error := model.GetUsersesByUserActivationKeyAndDisplayNameAndSpam(offset, limit, iUserActivationKey,iDisplayName,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserActivationKeyAndDisplayNameAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserActivationKeyAndDisplayNameAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserActivationKey := self.Args("user_activation_key").String()
	iDisplayName := self.Args("display_name").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserActivationKey) {
		_Users, _error := model.GetUsersesByUserActivationKeyAndDisplayNameAndDeleted(offset, limit, iUserActivationKey,iDisplayName,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserActivationKeyAndDisplayNameAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserActivationKeyAndSpamAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserActivationKey := self.Args("user_activation_key").String()
	iSpam := self.Args("spam").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserActivationKey) {
		_Users, _error := model.GetUsersesByUserActivationKeyAndSpamAndDeleted(offset, limit, iUserActivationKey,iSpam,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserActivationKeyAndSpamAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserStatusAndDisplayNameAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserStatus := self.Args("user_status").MustInt()
	iDisplayName := self.Args("display_name").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserStatus) {
		_Users, _error := model.GetUsersesByUserStatusAndDisplayNameAndSpam(offset, limit, iUserStatus,iDisplayName,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserStatusAndDisplayNameAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserStatusAndDisplayNameAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserStatus := self.Args("user_status").MustInt()
	iDisplayName := self.Args("display_name").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserStatus) {
		_Users, _error := model.GetUsersesByUserStatusAndDisplayNameAndDeleted(offset, limit, iUserStatus,iDisplayName,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserStatusAndDisplayNameAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserStatusAndSpamAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserStatus := self.Args("user_status").MustInt()
	iSpam := self.Args("spam").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserStatus) {
		_Users, _error := model.GetUsersesByUserStatusAndSpamAndDeleted(offset, limit, iUserStatus,iSpam,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserStatusAndSpamAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByDisplayNameAndSpamAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDisplayName := self.Args("display_name").String()
	iSpam := self.Args("spam").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iDisplayName) {
		_Users, _error := model.GetUsersesByDisplayNameAndSpamAndDeleted(offset, limit, iDisplayName,iSpam,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByDisplayNameAndSpamAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserLogin := self.Args("user_login").String()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserLogin(offset, limit, iId,iUserLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserPassHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserPass := self.Args("user_pass").String()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserPass(offset, limit, iId,iUserPass)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserPass's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserNicenameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserNicename := self.Args("user_nicename").String()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserNicename(offset, limit, iId,iUserNicename)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserNicename's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserEmail := self.Args("user_email").String()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserEmail(offset, limit, iId,iUserEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserUrl := self.Args("user_url").String()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserUrl(offset, limit, iId,iUserUrl)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserRegistered := self.Args("user_registered").Time()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserRegistered(offset, limit, iId,iUserRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserActivationKey := self.Args("user_activation_key").String()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserActivationKey(offset, limit, iId,iUserActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndUserStatus(offset, limit, iId,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndDisplayName(offset, limit, iId,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndSpam(offset, limit, iId,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByIdAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iId := self.Args("id").MustInt64()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersesByIdAndDeleted(offset, limit, iId,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByIdAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserPassHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserPass := self.Args("user_pass").String()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserPass(offset, limit, iUserLogin,iUserPass)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserPass's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserNicenameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserNicename := self.Args("user_nicename").String()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserNicename(offset, limit, iUserLogin,iUserNicename)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserNicename's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserEmail := self.Args("user_email").String()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserEmail(offset, limit, iUserLogin,iUserEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserUrl := self.Args("user_url").String()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserUrl(offset, limit, iUserLogin,iUserUrl)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserRegistered := self.Args("user_registered").Time()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserRegistered(offset, limit, iUserLogin,iUserRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserActivationKey := self.Args("user_activation_key").String()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserActivationKey(offset, limit, iUserLogin,iUserActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndUserStatus(offset, limit, iUserLogin,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndDisplayName(offset, limit, iUserLogin,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndSpam(offset, limit, iUserLogin,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserLoginAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserLogin := self.Args("user_login").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersesByUserLoginAndDeleted(offset, limit, iUserLogin,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserLoginAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserNicenameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserNicename := self.Args("user_nicename").String()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserNicename(offset, limit, iUserPass,iUserNicename)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserNicename's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserEmail := self.Args("user_email").String()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserEmail(offset, limit, iUserPass,iUserEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserUrl := self.Args("user_url").String()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserUrl(offset, limit, iUserPass,iUserUrl)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserRegistered := self.Args("user_registered").Time()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserRegistered(offset, limit, iUserPass,iUserRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserActivationKey := self.Args("user_activation_key").String()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserActivationKey(offset, limit, iUserPass,iUserActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndUserStatus(offset, limit, iUserPass,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndDisplayName(offset, limit, iUserPass,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndSpam(offset, limit, iUserPass,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserPassAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserPass := self.Args("user_pass").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersesByUserPassAndDeleted(offset, limit, iUserPass,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserPassAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iUserEmail := self.Args("user_email").String()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndUserEmail(offset, limit, iUserNicename,iUserEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndUserUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iUserUrl := self.Args("user_url").String()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndUserUrl(offset, limit, iUserNicename,iUserUrl)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndUserUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndUserRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iUserRegistered := self.Args("user_registered").Time()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndUserRegistered(offset, limit, iUserNicename,iUserRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndUserRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iUserActivationKey := self.Args("user_activation_key").String()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndUserActivationKey(offset, limit, iUserNicename,iUserActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndUserActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndUserStatus(offset, limit, iUserNicename,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndDisplayName(offset, limit, iUserNicename,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndSpam(offset, limit, iUserNicename,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserNicenameAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserNicename := self.Args("user_nicename").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersesByUserNicenameAndDeleted(offset, limit, iUserNicename,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserNicenameAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserEmailAndUserUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iUserUrl := self.Args("user_url").String()

	if helper.IsHas(iUserEmail) {
		_Users, _error := model.GetUsersesByUserEmailAndUserUrl(offset, limit, iUserEmail,iUserUrl)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserEmailAndUserUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserEmailAndUserRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iUserRegistered := self.Args("user_registered").Time()

	if helper.IsHas(iUserEmail) {
		_Users, _error := model.GetUsersesByUserEmailAndUserRegistered(offset, limit, iUserEmail,iUserRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserEmailAndUserRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserEmailAndUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iUserActivationKey := self.Args("user_activation_key").String()

	if helper.IsHas(iUserEmail) {
		_Users, _error := model.GetUsersesByUserEmailAndUserActivationKey(offset, limit, iUserEmail,iUserActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserEmailAndUserActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserEmailAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iUserEmail) {
		_Users, _error := model.GetUsersesByUserEmailAndUserStatus(offset, limit, iUserEmail,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserEmailAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserEmailAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserEmail) {
		_Users, _error := model.GetUsersesByUserEmailAndDisplayName(offset, limit, iUserEmail,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserEmailAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserEmailAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserEmail) {
		_Users, _error := model.GetUsersesByUserEmailAndSpam(offset, limit, iUserEmail,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserEmailAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserEmailAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserEmail := self.Args("user_email").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserEmail) {
		_Users, _error := model.GetUsersesByUserEmailAndDeleted(offset, limit, iUserEmail,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserEmailAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserUrlAndUserRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserUrl := self.Args("user_url").String()
	iUserRegistered := self.Args("user_registered").Time()

	if helper.IsHas(iUserUrl) {
		_Users, _error := model.GetUsersesByUserUrlAndUserRegistered(offset, limit, iUserUrl,iUserRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserUrlAndUserRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserUrlAndUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserUrl := self.Args("user_url").String()
	iUserActivationKey := self.Args("user_activation_key").String()

	if helper.IsHas(iUserUrl) {
		_Users, _error := model.GetUsersesByUserUrlAndUserActivationKey(offset, limit, iUserUrl,iUserActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserUrlAndUserActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserUrlAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserUrl := self.Args("user_url").String()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iUserUrl) {
		_Users, _error := model.GetUsersesByUserUrlAndUserStatus(offset, limit, iUserUrl,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserUrlAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserUrlAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserUrl := self.Args("user_url").String()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserUrl) {
		_Users, _error := model.GetUsersesByUserUrlAndDisplayName(offset, limit, iUserUrl,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserUrlAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserUrlAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserUrl := self.Args("user_url").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserUrl) {
		_Users, _error := model.GetUsersesByUserUrlAndSpam(offset, limit, iUserUrl,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserUrlAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserUrlAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserUrl := self.Args("user_url").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserUrl) {
		_Users, _error := model.GetUsersesByUserUrlAndDeleted(offset, limit, iUserUrl,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserUrlAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserRegisteredAndUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserRegistered := self.Args("user_registered").Time()
	iUserActivationKey := self.Args("user_activation_key").String()

	if helper.IsHas(iUserRegistered) {
		_Users, _error := model.GetUsersesByUserRegisteredAndUserActivationKey(offset, limit, iUserRegistered,iUserActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserRegisteredAndUserActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserRegisteredAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserRegistered := self.Args("user_registered").Time()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iUserRegistered) {
		_Users, _error := model.GetUsersesByUserRegisteredAndUserStatus(offset, limit, iUserRegistered,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserRegisteredAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserRegisteredAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserRegistered := self.Args("user_registered").Time()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserRegistered) {
		_Users, _error := model.GetUsersesByUserRegisteredAndDisplayName(offset, limit, iUserRegistered,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserRegisteredAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserRegisteredAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserRegistered := self.Args("user_registered").Time()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserRegistered) {
		_Users, _error := model.GetUsersesByUserRegisteredAndSpam(offset, limit, iUserRegistered,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserRegisteredAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserRegisteredAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserRegistered := self.Args("user_registered").Time()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserRegistered) {
		_Users, _error := model.GetUsersesByUserRegisteredAndDeleted(offset, limit, iUserRegistered,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserRegisteredAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserActivationKeyAndUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserActivationKey := self.Args("user_activation_key").String()
	iUserStatus := self.Args("user_status").MustInt()

	if helper.IsHas(iUserActivationKey) {
		_Users, _error := model.GetUsersesByUserActivationKeyAndUserStatus(offset, limit, iUserActivationKey,iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserActivationKeyAndUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserActivationKeyAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserActivationKey := self.Args("user_activation_key").String()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserActivationKey) {
		_Users, _error := model.GetUsersesByUserActivationKeyAndDisplayName(offset, limit, iUserActivationKey,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserActivationKeyAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserActivationKeyAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserActivationKey := self.Args("user_activation_key").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserActivationKey) {
		_Users, _error := model.GetUsersesByUserActivationKeyAndSpam(offset, limit, iUserActivationKey,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserActivationKeyAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserActivationKeyAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserActivationKey := self.Args("user_activation_key").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserActivationKey) {
		_Users, _error := model.GetUsersesByUserActivationKeyAndDeleted(offset, limit, iUserActivationKey,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserActivationKeyAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserStatusAndDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserStatus := self.Args("user_status").MustInt()
	iDisplayName := self.Args("display_name").String()

	if helper.IsHas(iUserStatus) {
		_Users, _error := model.GetUsersesByUserStatusAndDisplayName(offset, limit, iUserStatus,iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserStatusAndDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserStatusAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserStatus := self.Args("user_status").MustInt()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iUserStatus) {
		_Users, _error := model.GetUsersesByUserStatusAndSpam(offset, limit, iUserStatus,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserStatusAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByUserStatusAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iUserStatus := self.Args("user_status").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iUserStatus) {
		_Users, _error := model.GetUsersesByUserStatusAndDeleted(offset, limit, iUserStatus,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByUserStatusAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByDisplayNameAndSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDisplayName := self.Args("display_name").String()
	iSpam := self.Args("spam").MustInt()

	if helper.IsHas(iDisplayName) {
		_Users, _error := model.GetUsersesByDisplayNameAndSpam(offset, limit, iDisplayName,iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByDisplayNameAndSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesByDisplayNameAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDisplayName := self.Args("display_name").String()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iDisplayName) {
		_Users, _error := model.GetUsersesByDisplayNameAndDeleted(offset, limit, iDisplayName,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesByDisplayNameAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesBySpamAndDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSpam := self.Args("spam").MustInt()
	iDeleted := self.Args("deleted").MustInt()

	if helper.IsHas(iSpam) {
		_Users, _error := model.GetUsersesBySpamAndDeleted(offset, limit, iSpam,iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersesBySpamAndDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_Users, _error := model.GetUserses(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUserses' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iId := self.Args("ID").MustInt64()
	if helper.IsHas(iId) {
		_Users := model.HasUsersViaId(iId)
		var m = map[string]interface{}{}
		m["users"] = _Users
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersViaUserLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserLogin := self.Args("user_login").String()
	if helper.IsHas(iUserLogin) {
		_Users := model.HasUsersViaUserLogin(iUserLogin)
		var m = map[string]interface{}{}
		m["users"] = _Users
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersViaUserLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersViaUserPassHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserPass := self.Args("user_pass").String()
	if helper.IsHas(iUserPass) {
		_Users := model.HasUsersViaUserPass(iUserPass)
		var m = map[string]interface{}{}
		m["users"] = _Users
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersViaUserPass's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersViaUserNicenameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserNicename := self.Args("user_nicename").String()
	if helper.IsHas(iUserNicename) {
		_Users := model.HasUsersViaUserNicename(iUserNicename)
		var m = map[string]interface{}{}
		m["users"] = _Users
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersViaUserNicename's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersViaUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserEmail := self.Args("user_email").String()
	if helper.IsHas(iUserEmail) {
		_Users := model.HasUsersViaUserEmail(iUserEmail)
		var m = map[string]interface{}{}
		m["users"] = _Users
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersViaUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersViaUserUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserUrl := self.Args("user_url").String()
	if helper.IsHas(iUserUrl) {
		_Users := model.HasUsersViaUserUrl(iUserUrl)
		var m = map[string]interface{}{}
		m["users"] = _Users
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersViaUserUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersViaUserRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserRegistered := self.Args("user_registered").Time()
	if helper.IsHas(iUserRegistered) {
		_Users := model.HasUsersViaUserRegistered(iUserRegistered)
		var m = map[string]interface{}{}
		m["users"] = _Users
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersViaUserRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersViaUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserActivationKey := self.Args("user_activation_key").String()
	if helper.IsHas(iUserActivationKey) {
		_Users := model.HasUsersViaUserActivationKey(iUserActivationKey)
		var m = map[string]interface{}{}
		m["users"] = _Users
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersViaUserActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersViaUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserStatus := self.Args("user_status").MustInt()
	if helper.IsHas(iUserStatus) {
		_Users := model.HasUsersViaUserStatus(iUserStatus)
		var m = map[string]interface{}{}
		m["users"] = _Users
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersViaUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersViaDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDisplayName := self.Args("display_name").String()
	if helper.IsHas(iDisplayName) {
		_Users := model.HasUsersViaDisplayName(iDisplayName)
		var m = map[string]interface{}{}
		m["users"] = _Users
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersViaDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersViaSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSpam := self.Args("spam").MustInt()
	if helper.IsHas(iSpam) {
		_Users := model.HasUsersViaSpam(iSpam)
		var m = map[string]interface{}{}
		m["users"] = _Users
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersViaSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasUsersViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDeleted := self.Args("deleted").MustInt()
	if helper.IsHas(iDeleted) {
		_Users := model.HasUsersViaDeleted(iDeleted)
		var m = map[string]interface{}{}
		m["users"] = _Users
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasUsersViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iId := self.Args("ID").MustInt64()
	if helper.IsHas(iId) {
		_Users, _error := model.GetUsersViaId(iId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersViaUserLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserLogin := self.Args("user_login").String()
	if helper.IsHas(iUserLogin) {
		_Users, _error := model.GetUsersViaUserLogin(iUserLogin)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersViaUserLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersViaUserPassHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserPass := self.Args("user_pass").String()
	if helper.IsHas(iUserPass) {
		_Users, _error := model.GetUsersViaUserPass(iUserPass)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersViaUserPass's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersViaUserNicenameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserNicename := self.Args("user_nicename").String()
	if helper.IsHas(iUserNicename) {
		_Users, _error := model.GetUsersViaUserNicename(iUserNicename)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersViaUserNicename's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersViaUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserEmail := self.Args("user_email").String()
	if helper.IsHas(iUserEmail) {
		_Users, _error := model.GetUsersViaUserEmail(iUserEmail)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersViaUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersViaUserUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserUrl := self.Args("user_url").String()
	if helper.IsHas(iUserUrl) {
		_Users, _error := model.GetUsersViaUserUrl(iUserUrl)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersViaUserUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersViaUserRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserRegistered := self.Args("user_registered").Time()
	if helper.IsHas(iUserRegistered) {
		_Users, _error := model.GetUsersViaUserRegistered(iUserRegistered)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersViaUserRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersViaUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserActivationKey := self.Args("user_activation_key").String()
	if helper.IsHas(iUserActivationKey) {
		_Users, _error := model.GetUsersViaUserActivationKey(iUserActivationKey)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersViaUserActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersViaUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iUserStatus := self.Args("user_status").MustInt()
	if helper.IsHas(iUserStatus) {
		_Users, _error := model.GetUsersViaUserStatus(iUserStatus)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersViaUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersViaDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDisplayName := self.Args("display_name").String()
	if helper.IsHas(iDisplayName) {
		_Users, _error := model.GetUsersViaDisplayName(iDisplayName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersViaDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersViaSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSpam := self.Args("spam").MustInt()
	if helper.IsHas(iSpam) {
		_Users, _error := model.GetUsersViaSpam(iSpam)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersViaSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetUsersViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDeleted := self.Args("deleted").MustInt()
	if helper.IsHas(iDeleted) {
		_Users, _error := model.GetUsersViaDeleted(iDeleted)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the GetUsersViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("ID").MustInt64()
	if helper.IsHas(Id_) {
		var iUsers model.Users
		self.Bind(&iUsers)
		_Users, _error := model.SetUsersViaId(Id_, &iUsers)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the SetUsersViaId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersViaUserLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserLogin_ := self.Args("user_login").String()
	if helper.IsHas(UserLogin_) {
		var iUsers model.Users
		self.Bind(&iUsers)
		_Users, _error := model.SetUsersViaUserLogin(UserLogin_, &iUsers)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the SetUsersViaUserLogin's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersViaUserPassHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserPass_ := self.Args("user_pass").String()
	if helper.IsHas(UserPass_) {
		var iUsers model.Users
		self.Bind(&iUsers)
		_Users, _error := model.SetUsersViaUserPass(UserPass_, &iUsers)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the SetUsersViaUserPass's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersViaUserNicenameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserNicename_ := self.Args("user_nicename").String()
	if helper.IsHas(UserNicename_) {
		var iUsers model.Users
		self.Bind(&iUsers)
		_Users, _error := model.SetUsersViaUserNicename(UserNicename_, &iUsers)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the SetUsersViaUserNicename's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersViaUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserEmail_ := self.Args("user_email").String()
	if helper.IsHas(UserEmail_) {
		var iUsers model.Users
		self.Bind(&iUsers)
		_Users, _error := model.SetUsersViaUserEmail(UserEmail_, &iUsers)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the SetUsersViaUserEmail's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersViaUserUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserUrl_ := self.Args("user_url").String()
	if helper.IsHas(UserUrl_) {
		var iUsers model.Users
		self.Bind(&iUsers)
		_Users, _error := model.SetUsersViaUserUrl(UserUrl_, &iUsers)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the SetUsersViaUserUrl's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersViaUserRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserRegistered_ := self.Args("user_registered").Time()
	if helper.IsHas(UserRegistered_) {
		var iUsers model.Users
		self.Bind(&iUsers)
		_Users, _error := model.SetUsersViaUserRegistered(UserRegistered_, &iUsers)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the SetUsersViaUserRegistered's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersViaUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserActivationKey_ := self.Args("user_activation_key").String()
	if helper.IsHas(UserActivationKey_) {
		var iUsers model.Users
		self.Bind(&iUsers)
		_Users, _error := model.SetUsersViaUserActivationKey(UserActivationKey_, &iUsers)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the SetUsersViaUserActivationKey's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersViaUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserStatus_ := self.Args("user_status").MustInt()
	if helper.IsHas(UserStatus_) {
		var iUsers model.Users
		self.Bind(&iUsers)
		_Users, _error := model.SetUsersViaUserStatus(UserStatus_, &iUsers)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the SetUsersViaUserStatus's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersViaDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DisplayName_ := self.Args("display_name").String()
	if helper.IsHas(DisplayName_) {
		var iUsers model.Users
		self.Bind(&iUsers)
		_Users, _error := model.SetUsersViaDisplayName(DisplayName_, &iUsers)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the SetUsersViaDisplayName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersViaSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Spam_ := self.Args("spam").MustInt()
	if helper.IsHas(Spam_) {
		var iUsers model.Users
		self.Bind(&iUsers)
		_Users, _error := model.SetUsersViaSpam(Spam_, &iUsers)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the SetUsersViaSpam's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetUsersViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	if helper.IsHas(Deleted_) {
		var iUsers model.Users
		self.Bind(&iUsers)
		_Users, _error := model.SetUsersViaDeleted(Deleted_, &iUsers)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Users)
	}
	herr.Message = "Can't get to the SetUsersViaDeleted's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddUsersHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("ID").MustInt64()
	UserLogin_ := self.Args("user_login").String()
	UserPass_ := self.Args("user_pass").String()
	UserNicename_ := self.Args("user_nicename").String()
	UserEmail_ := self.Args("user_email").String()
	UserUrl_ := self.Args("user_url").String()
	UserRegistered_ := self.Args("user_registered").Time()
	UserActivationKey_ := self.Args("user_activation_key").String()
	UserStatus_ := self.Args("user_status").MustInt()
	DisplayName_ := self.Args("display_name").String()
	Spam_ := self.Args("spam").MustInt()
	Deleted_ := self.Args("deleted").MustInt()

	if helper.IsHas(Id_) {
		_error := model.AddUsers(Id_,UserLogin_,UserPass_,UserNicename_,UserEmail_,UserUrl_,UserRegistered_,UserActivationKey_,UserStatus_,DisplayName_,Spam_,Deleted_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddUsers's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostUsersHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iUsers model.Users
	self.Bind(&iUsers)
	_int64, _error := model.PostUsers(&iUsers)
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

func PutUsersHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iUsers model.Users
	self.Bind(&iUsers)
	_int64, _error := model.PutUsers(&iUsers)
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

func PutUsersViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("ID").MustInt64()
	var iUsers model.Users
	self.Bind(&iUsers)
	_int64, _error := model.PutUsersViaId(Id_, &iUsers)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsersViaUserLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserLogin_ := self.Args("user_login").String()
	var iUsers model.Users
	self.Bind(&iUsers)
	_int64, _error := model.PutUsersViaUserLogin(UserLogin_, &iUsers)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsersViaUserPassHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserPass_ := self.Args("user_pass").String()
	var iUsers model.Users
	self.Bind(&iUsers)
	_int64, _error := model.PutUsersViaUserPass(UserPass_, &iUsers)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsersViaUserNicenameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserNicename_ := self.Args("user_nicename").String()
	var iUsers model.Users
	self.Bind(&iUsers)
	_int64, _error := model.PutUsersViaUserNicename(UserNicename_, &iUsers)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsersViaUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserEmail_ := self.Args("user_email").String()
	var iUsers model.Users
	self.Bind(&iUsers)
	_int64, _error := model.PutUsersViaUserEmail(UserEmail_, &iUsers)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsersViaUserUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserUrl_ := self.Args("user_url").String()
	var iUsers model.Users
	self.Bind(&iUsers)
	_int64, _error := model.PutUsersViaUserUrl(UserUrl_, &iUsers)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsersViaUserRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserRegistered_ := self.Args("user_registered").Time()
	var iUsers model.Users
	self.Bind(&iUsers)
	_int64, _error := model.PutUsersViaUserRegistered(UserRegistered_, &iUsers)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsersViaUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserActivationKey_ := self.Args("user_activation_key").String()
	var iUsers model.Users
	self.Bind(&iUsers)
	_int64, _error := model.PutUsersViaUserActivationKey(UserActivationKey_, &iUsers)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsersViaUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserStatus_ := self.Args("user_status").MustInt()
	var iUsers model.Users
	self.Bind(&iUsers)
	_int64, _error := model.PutUsersViaUserStatus(UserStatus_, &iUsers)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsersViaDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DisplayName_ := self.Args("display_name").String()
	var iUsers model.Users
	self.Bind(&iUsers)
	_int64, _error := model.PutUsersViaDisplayName(DisplayName_, &iUsers)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsersViaSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Spam_ := self.Args("spam").MustInt()
	var iUsers model.Users
	self.Bind(&iUsers)
	_int64, _error := model.PutUsersViaSpam(Spam_, &iUsers)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUsersViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	var iUsers model.Users
	self.Bind(&iUsers)
	_int64, _error := model.PutUsersViaDeleted(Deleted_, &iUsers)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("ID").MustInt64()
	var iUsers model.Users
	self.Bind(&iUsers)
	var iMap = helper.StructToMap(iUsers)
	_error := model.UpdateUsersViaId(Id_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersViaUserLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserLogin_ := self.Args("user_login").String()
	var iUsers model.Users
	self.Bind(&iUsers)
	var iMap = helper.StructToMap(iUsers)
	_error := model.UpdateUsersViaUserLogin(UserLogin_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersViaUserPassHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserPass_ := self.Args("user_pass").String()
	var iUsers model.Users
	self.Bind(&iUsers)
	var iMap = helper.StructToMap(iUsers)
	_error := model.UpdateUsersViaUserPass(UserPass_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersViaUserNicenameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserNicename_ := self.Args("user_nicename").String()
	var iUsers model.Users
	self.Bind(&iUsers)
	var iMap = helper.StructToMap(iUsers)
	_error := model.UpdateUsersViaUserNicename(UserNicename_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersViaUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserEmail_ := self.Args("user_email").String()
	var iUsers model.Users
	self.Bind(&iUsers)
	var iMap = helper.StructToMap(iUsers)
	_error := model.UpdateUsersViaUserEmail(UserEmail_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersViaUserUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserUrl_ := self.Args("user_url").String()
	var iUsers model.Users
	self.Bind(&iUsers)
	var iMap = helper.StructToMap(iUsers)
	_error := model.UpdateUsersViaUserUrl(UserUrl_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersViaUserRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserRegistered_ := self.Args("user_registered").Time()
	var iUsers model.Users
	self.Bind(&iUsers)
	var iMap = helper.StructToMap(iUsers)
	_error := model.UpdateUsersViaUserRegistered(UserRegistered_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersViaUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserActivationKey_ := self.Args("user_activation_key").String()
	var iUsers model.Users
	self.Bind(&iUsers)
	var iMap = helper.StructToMap(iUsers)
	_error := model.UpdateUsersViaUserActivationKey(UserActivationKey_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersViaUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserStatus_ := self.Args("user_status").MustInt()
	var iUsers model.Users
	self.Bind(&iUsers)
	var iMap = helper.StructToMap(iUsers)
	_error := model.UpdateUsersViaUserStatus(UserStatus_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersViaDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DisplayName_ := self.Args("display_name").String()
	var iUsers model.Users
	self.Bind(&iUsers)
	var iMap = helper.StructToMap(iUsers)
	_error := model.UpdateUsersViaDisplayName(DisplayName_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersViaSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Spam_ := self.Args("spam").MustInt()
	var iUsers model.Users
	self.Bind(&iUsers)
	var iMap = helper.StructToMap(iUsers)
	_error := model.UpdateUsersViaSpam(Spam_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateUsersViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	var iUsers model.Users
	self.Bind(&iUsers)
	var iMap = helper.StructToMap(iUsers)
	_error := model.UpdateUsersViaDeleted(Deleted_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("ID").MustInt64()
	_int64, _error := model.DeleteUsers(Id_)
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

func DeleteUsersViaIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Id_ := self.Args("ID").MustInt64()
	_error := model.DeleteUsersViaId(Id_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersViaUserLoginHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserLogin_ := self.Args("user_login").String()
	_error := model.DeleteUsersViaUserLogin(UserLogin_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersViaUserPassHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserPass_ := self.Args("user_pass").String()
	_error := model.DeleteUsersViaUserPass(UserPass_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersViaUserNicenameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserNicename_ := self.Args("user_nicename").String()
	_error := model.DeleteUsersViaUserNicename(UserNicename_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersViaUserEmailHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserEmail_ := self.Args("user_email").String()
	_error := model.DeleteUsersViaUserEmail(UserEmail_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersViaUserUrlHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserUrl_ := self.Args("user_url").String()
	_error := model.DeleteUsersViaUserUrl(UserUrl_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersViaUserRegisteredHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserRegistered_ := self.Args("user_registered").Time()
	_error := model.DeleteUsersViaUserRegistered(UserRegistered_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersViaUserActivationKeyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserActivationKey_ := self.Args("user_activation_key").String()
	_error := model.DeleteUsersViaUserActivationKey(UserActivationKey_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersViaUserStatusHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	UserStatus_ := self.Args("user_status").MustInt()
	_error := model.DeleteUsersViaUserStatus(UserStatus_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersViaDisplayNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	DisplayName_ := self.Args("display_name").String()
	_error := model.DeleteUsersViaDisplayName(DisplayName_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersViaSpamHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Spam_ := self.Args("spam").MustInt()
	_error := model.DeleteUsersViaSpam(Spam_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteUsersViaDeletedHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Deleted_ := self.Args("deleted").MustInt()
	_error := model.DeleteUsersViaDeleted(Deleted_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
