package handler

import (
	"github.com/insionng/zenpress/helper"
	"github.com/insionng/zenpress/model"
	"github.com/insionng/macross"
)

func GetTermsesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetTermsesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["termssCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetTermsesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermsCountViaTermIdHandler(self *macross.Context) error {
	TermId_ := self.Args("term_id").MustInt64()
	_int64 := model.GetTermsCountViaTermId(TermId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["termsCount"] = 0
	}
	m["termsCount"] = _int64
	return self.JSON(m)
}

func GetTermsCountViaNameHandler(self *macross.Context) error {
	Name_ := self.Args("name").String()
	_int64 := model.GetTermsCountViaName(Name_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["termsCount"] = 0
	}
	m["termsCount"] = _int64
	return self.JSON(m)
}

func GetTermsCountViaSlugHandler(self *macross.Context) error {
	Slug_ := self.Args("slug").String()
	_int64 := model.GetTermsCountViaSlug(Slug_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["termsCount"] = 0
	}
	m["termsCount"] = _int64
	return self.JSON(m)
}

func GetTermsCountViaTermGroupHandler(self *macross.Context) error {
	TermGroup_ := self.Args("term_group").MustInt64()
	_int64 := model.GetTermsCountViaTermGroup(TermGroup_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["termsCount"] = 0
	}
	m["termsCount"] = _int64
	return self.JSON(m)
}

func GetTermsesViaTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTermId := self.Args("term_id").MustInt64()
	if (offset > 0) && helper.IsHas(iTermId) {
		_Terms, _error := model.GetTermsesViaTermId(offset, limit, iTermId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the GetTermsesViaTermId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermsesViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iName := self.Args("name").String()
	if (offset > 0) && helper.IsHas(iName) {
		_Terms, _error := model.GetTermsesViaName(offset, limit, iName, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the GetTermsesViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermsesViaSlugHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iSlug := self.Args("slug").String()
	if (offset > 0) && helper.IsHas(iSlug) {
		_Terms, _error := model.GetTermsesViaSlug(offset, limit, iSlug, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the GetTermsesViaSlug's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermsesViaTermGroupHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTermGroup := self.Args("term_group").MustInt64()
	if (offset > 0) && helper.IsHas(iTermGroup) {
		_Terms, _error := model.GetTermsesViaTermGroup(offset, limit, iTermGroup, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the GetTermsesViaTermGroup's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermsesByTermIdAndNameAndSlugHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermId := self.Args("term_id").MustInt64()
	iName := self.Args("name").String()
	iSlug := self.Args("slug").String()

	if helper.IsHas(iTermId) {
		_Terms, _error := model.GetTermsesByTermIdAndNameAndSlug(offset, limit, iTermId,iName,iSlug)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the GetTermsesByTermIdAndNameAndSlug's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermsesByTermIdAndNameAndTermGroupHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermId := self.Args("term_id").MustInt64()
	iName := self.Args("name").String()
	iTermGroup := self.Args("term_group").MustInt64()

	if helper.IsHas(iTermId) {
		_Terms, _error := model.GetTermsesByTermIdAndNameAndTermGroup(offset, limit, iTermId,iName,iTermGroup)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the GetTermsesByTermIdAndNameAndTermGroup's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermsesByTermIdAndSlugAndTermGroupHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermId := self.Args("term_id").MustInt64()
	iSlug := self.Args("slug").String()
	iTermGroup := self.Args("term_group").MustInt64()

	if helper.IsHas(iTermId) {
		_Terms, _error := model.GetTermsesByTermIdAndSlugAndTermGroup(offset, limit, iTermId,iSlug,iTermGroup)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the GetTermsesByTermIdAndSlugAndTermGroup's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermsesByNameAndSlugAndTermGroupHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iSlug := self.Args("slug").String()
	iTermGroup := self.Args("term_group").MustInt64()

	if helper.IsHas(iName) {
		_Terms, _error := model.GetTermsesByNameAndSlugAndTermGroup(offset, limit, iName,iSlug,iTermGroup)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the GetTermsesByNameAndSlugAndTermGroup's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermsesByTermIdAndNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermId := self.Args("term_id").MustInt64()
	iName := self.Args("name").String()

	if helper.IsHas(iTermId) {
		_Terms, _error := model.GetTermsesByTermIdAndName(offset, limit, iTermId,iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the GetTermsesByTermIdAndName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermsesByTermIdAndSlugHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermId := self.Args("term_id").MustInt64()
	iSlug := self.Args("slug").String()

	if helper.IsHas(iTermId) {
		_Terms, _error := model.GetTermsesByTermIdAndSlug(offset, limit, iTermId,iSlug)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the GetTermsesByTermIdAndSlug's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermsesByTermIdAndTermGroupHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermId := self.Args("term_id").MustInt64()
	iTermGroup := self.Args("term_group").MustInt64()

	if helper.IsHas(iTermId) {
		_Terms, _error := model.GetTermsesByTermIdAndTermGroup(offset, limit, iTermId,iTermGroup)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the GetTermsesByTermIdAndTermGroup's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermsesByNameAndSlugHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iSlug := self.Args("slug").String()

	if helper.IsHas(iName) {
		_Terms, _error := model.GetTermsesByNameAndSlug(offset, limit, iName,iSlug)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the GetTermsesByNameAndSlug's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermsesByNameAndTermGroupHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iName := self.Args("name").String()
	iTermGroup := self.Args("term_group").MustInt64()

	if helper.IsHas(iName) {
		_Terms, _error := model.GetTermsesByNameAndTermGroup(offset, limit, iName,iTermGroup)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the GetTermsesByNameAndTermGroup's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermsesBySlugAndTermGroupHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iSlug := self.Args("slug").String()
	iTermGroup := self.Args("term_group").MustInt64()

	if helper.IsHas(iSlug) {
		_Terms, _error := model.GetTermsesBySlugAndTermGroup(offset, limit, iSlug,iTermGroup)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the GetTermsesBySlugAndTermGroup's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermsesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_Terms, _error := model.GetTermses(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the GetTermses' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTermsViaTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTermId := self.Args("term_id").MustInt64()
	if helper.IsHas(iTermId) {
		_Terms := model.HasTermsViaTermId(iTermId)
		var m = map[string]interface{}{}
		m["terms"] = _Terms
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermsViaTermId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTermsViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iName := self.Args("name").String()
	if helper.IsHas(iName) {
		_Terms := model.HasTermsViaName(iName)
		var m = map[string]interface{}{}
		m["terms"] = _Terms
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermsViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTermsViaSlugHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSlug := self.Args("slug").String()
	if helper.IsHas(iSlug) {
		_Terms := model.HasTermsViaSlug(iSlug)
		var m = map[string]interface{}{}
		m["terms"] = _Terms
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermsViaSlug's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTermsViaTermGroupHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTermGroup := self.Args("term_group").MustInt64()
	if helper.IsHas(iTermGroup) {
		_Terms := model.HasTermsViaTermGroup(iTermGroup)
		var m = map[string]interface{}{}
		m["terms"] = _Terms
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermsViaTermGroup's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermsViaTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTermId := self.Args("term_id").MustInt64()
	if helper.IsHas(iTermId) {
		_Terms, _error := model.GetTermsViaTermId(iTermId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the GetTermsViaTermId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermsViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iName := self.Args("name").String()
	if helper.IsHas(iName) {
		_Terms, _error := model.GetTermsViaName(iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the GetTermsViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermsViaSlugHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSlug := self.Args("slug").String()
	if helper.IsHas(iSlug) {
		_Terms, _error := model.GetTermsViaSlug(iSlug)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the GetTermsViaSlug's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermsViaTermGroupHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTermGroup := self.Args("term_group").MustInt64()
	if helper.IsHas(iTermGroup) {
		_Terms, _error := model.GetTermsViaTermGroup(iTermGroup)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the GetTermsViaTermGroup's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermsViaTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermId_ := self.Args("term_id").MustInt64()
	if helper.IsHas(TermId_) {
		var iTerms model.Terms
		self.Bind(&iTerms)
		_Terms, _error := model.SetTermsViaTermId(TermId_, &iTerms)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the SetTermsViaTermId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermsViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	if helper.IsHas(Name_) {
		var iTerms model.Terms
		self.Bind(&iTerms)
		_Terms, _error := model.SetTermsViaName(Name_, &iTerms)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the SetTermsViaName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermsViaSlugHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Slug_ := self.Args("slug").String()
	if helper.IsHas(Slug_) {
		var iTerms model.Terms
		self.Bind(&iTerms)
		_Terms, _error := model.SetTermsViaSlug(Slug_, &iTerms)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the SetTermsViaSlug's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermsViaTermGroupHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermGroup_ := self.Args("term_group").MustInt64()
	if helper.IsHas(TermGroup_) {
		var iTerms model.Terms
		self.Bind(&iTerms)
		_Terms, _error := model.SetTermsViaTermGroup(TermGroup_, &iTerms)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the SetTermsViaTermGroup's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddTermsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermId_ := self.Args("term_id").MustInt64()
	Name_ := self.Args("name").String()
	Slug_ := self.Args("slug").String()
	TermGroup_ := self.Args("term_group").MustInt64()

	if helper.IsHas(TermId_) {
		_error := model.AddTerms(TermId_,Name_,Slug_,TermGroup_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddTerms's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostTermsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iTerms model.Terms
	self.Bind(&iTerms)
	_int64, _error := model.PostTerms(&iTerms)
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

func PutTermsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iTerms model.Terms
	self.Bind(&iTerms)
	_int64, _error := model.PutTerms(&iTerms)
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

func PutTermsViaTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermId_ := self.Args("term_id").MustInt64()
	var iTerms model.Terms
	self.Bind(&iTerms)
	_int64, _error := model.PutTermsViaTermId(TermId_, &iTerms)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTermsViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	var iTerms model.Terms
	self.Bind(&iTerms)
	_int64, _error := model.PutTermsViaName(Name_, &iTerms)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTermsViaSlugHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Slug_ := self.Args("slug").String()
	var iTerms model.Terms
	self.Bind(&iTerms)
	_int64, _error := model.PutTermsViaSlug(Slug_, &iTerms)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTermsViaTermGroupHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermGroup_ := self.Args("term_group").MustInt64()
	var iTerms model.Terms
	self.Bind(&iTerms)
	_int64, _error := model.PutTermsViaTermGroup(TermGroup_, &iTerms)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermsViaTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermId_ := self.Args("term_id").MustInt64()
	var iTerms model.Terms
	self.Bind(&iTerms)
	var iMap = helper.StructToMap(iTerms)
	_error := model.UpdateTermsViaTermId(TermId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermsViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	var iTerms model.Terms
	self.Bind(&iTerms)
	var iMap = helper.StructToMap(iTerms)
	_error := model.UpdateTermsViaName(Name_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermsViaSlugHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Slug_ := self.Args("slug").String()
	var iTerms model.Terms
	self.Bind(&iTerms)
	var iMap = helper.StructToMap(iTerms)
	_error := model.UpdateTermsViaSlug(Slug_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermsViaTermGroupHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermGroup_ := self.Args("term_group").MustInt64()
	var iTerms model.Terms
	self.Bind(&iTerms)
	var iMap = helper.StructToMap(iTerms)
	_error := model.UpdateTermsViaTermGroup(TermGroup_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTermsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermId_ := self.Args("term_id").MustInt64()
	_int64, _error := model.DeleteTerms(TermId_)
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

func DeleteTermsViaTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermId_ := self.Args("term_id").MustInt64()
	_error := model.DeleteTermsViaTermId(TermId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTermsViaNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	_error := model.DeleteTermsViaName(Name_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTermsViaSlugHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Slug_ := self.Args("slug").String()
	_error := model.DeleteTermsViaSlug(Slug_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTermsViaTermGroupHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermGroup_ := self.Args("term_group").MustInt64()
	_error := model.DeleteTermsViaTermGroup(TermGroup_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
