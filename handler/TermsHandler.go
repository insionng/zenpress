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

func GetTermsCountByTermIdHandler(self *macross.Context) error {
	TermId_ := self.Args("term_id").MustInt64()
	_int64 := model.GetTermsCountByTermId(TermId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["termsCount"] = 0
	}
	m["termsCount"] = _int64
	return self.JSON(m)
}

func GetTermsCountByNameHandler(self *macross.Context) error {
	Name_ := self.Args("name").String()
	_int64 := model.GetTermsCountByName(Name_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["termsCount"] = 0
	}
	m["termsCount"] = _int64
	return self.JSON(m)
}

func GetTermsCountBySlugHandler(self *macross.Context) error {
	Slug_ := self.Args("slug").String()
	_int64 := model.GetTermsCountBySlug(Slug_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["termsCount"] = 0
	}
	m["termsCount"] = _int64
	return self.JSON(m)
}

func GetTermsCountByTermGroupHandler(self *macross.Context) error {
	TermGroup_ := self.Args("term_group").MustInt64()
	_int64 := model.GetTermsCountByTermGroup(TermGroup_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["termsCount"] = 0
	}
	m["termsCount"] = _int64
	return self.JSON(m)
}

func GetTermsesByTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTermId := self.Args("term_id").MustInt64()
	if (offset > 0) && helper.IsHas(iTermId) {
		_Terms, _error := model.GetTermsesByTermId(offset, limit, iTermId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the GetTermsesByTermId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermsesByNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iName := self.Args("name").String()
	if (offset > 0) && helper.IsHas(iName) {
		_Terms, _error := model.GetTermsesByName(offset, limit, iName, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the GetTermsesByName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermsesBySlugHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iSlug := self.Args("slug").String()
	if (offset > 0) && helper.IsHas(iSlug) {
		_Terms, _error := model.GetTermsesBySlug(offset, limit, iSlug, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the GetTermsesBySlug's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermsesByTermGroupHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTermGroup := self.Args("term_group").MustInt64()
	if (offset > 0) && helper.IsHas(iTermGroup) {
		_Terms, _error := model.GetTermsesByTermGroup(offset, limit, iTermGroup, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the GetTermsesByTermGroup's args."
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

func GetHasTermsByTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTermId := self.Args("term_id").MustInt64()
	if helper.IsHas(iTermId) {
		_Terms := model.HasTermsByTermId(iTermId)
		var m = map[string]interface{}{}
		m["terms"] = _Terms
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermsByTermId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTermsByNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iName := self.Args("name").String()
	if helper.IsHas(iName) {
		_Terms := model.HasTermsByName(iName)
		var m = map[string]interface{}{}
		m["terms"] = _Terms
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermsByName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTermsBySlugHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSlug := self.Args("slug").String()
	if helper.IsHas(iSlug) {
		_Terms := model.HasTermsBySlug(iSlug)
		var m = map[string]interface{}{}
		m["terms"] = _Terms
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermsBySlug's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTermsByTermGroupHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTermGroup := self.Args("term_group").MustInt64()
	if helper.IsHas(iTermGroup) {
		_Terms := model.HasTermsByTermGroup(iTermGroup)
		var m = map[string]interface{}{}
		m["terms"] = _Terms
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermsByTermGroup's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermsByTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTermId := self.Args("term_id").MustInt64()
	if helper.IsHas(iTermId) {
		_Terms, _error := model.GetTermsByTermId(iTermId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the GetTermsByTermId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermsByNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iName := self.Args("name").String()
	if helper.IsHas(iName) {
		_Terms, _error := model.GetTermsByName(iName)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the GetTermsByName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermsBySlugHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iSlug := self.Args("slug").String()
	if helper.IsHas(iSlug) {
		_Terms, _error := model.GetTermsBySlug(iSlug)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the GetTermsBySlug's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermsByTermGroupHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTermGroup := self.Args("term_group").MustInt64()
	if helper.IsHas(iTermGroup) {
		_Terms, _error := model.GetTermsByTermGroup(iTermGroup)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the GetTermsByTermGroup's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermsByTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermId_ := self.Args("term_id").MustInt64()
	if helper.IsHas(TermId_) {
		var iTerms model.Terms
		self.Bind(&iTerms)
		_Terms, _error := model.SetTermsByTermId(TermId_, &iTerms)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the SetTermsByTermId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermsByNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	if helper.IsHas(Name_) {
		var iTerms model.Terms
		self.Bind(&iTerms)
		_Terms, _error := model.SetTermsByName(Name_, &iTerms)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the SetTermsByName's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermsBySlugHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Slug_ := self.Args("slug").String()
	if helper.IsHas(Slug_) {
		var iTerms model.Terms
		self.Bind(&iTerms)
		_Terms, _error := model.SetTermsBySlug(Slug_, &iTerms)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the SetTermsBySlug's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermsByTermGroupHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermGroup_ := self.Args("term_group").MustInt64()
	if helper.IsHas(TermGroup_) {
		var iTerms model.Terms
		self.Bind(&iTerms)
		_Terms, _error := model.SetTermsByTermGroup(TermGroup_, &iTerms)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_Terms)
	}
	herr.Message = "Can't get to the SetTermsByTermGroup's args."
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

func PutTermsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iTerms model.Terms
	self.Bind(&iTerms)
	_int64, _error := model.PutTerms(&iTerms)
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

func PutTermsByTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermId_ := self.Args("term_id").MustInt64()
	var iTerms model.Terms
	self.Bind(&iTerms)
	_int64, _error := model.PutTermsByTermId(TermId_, &iTerms)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTermsByNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	var iTerms model.Terms
	self.Bind(&iTerms)
	_int64, _error := model.PutTermsByName(Name_, &iTerms)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTermsBySlugHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Slug_ := self.Args("slug").String()
	var iTerms model.Terms
	self.Bind(&iTerms)
	_int64, _error := model.PutTermsBySlug(Slug_, &iTerms)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTermsByTermGroupHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermGroup_ := self.Args("term_group").MustInt64()
	var iTerms model.Terms
	self.Bind(&iTerms)
	_int64, _error := model.PutTermsByTermGroup(TermGroup_, &iTerms)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermsByTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermId_ := self.Args("term_id").MustInt64()
	var iTerms model.Terms
	self.Bind(&iTerms)
	var iMap = helper.StructToMap(iTerms)
	_error := model.UpdateTermsByTermId(TermId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermsByNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	var iTerms model.Terms
	self.Bind(&iTerms)
	var iMap = helper.StructToMap(iTerms)
	_error := model.UpdateTermsByName(Name_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermsBySlugHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Slug_ := self.Args("slug").String()
	var iTerms model.Terms
	self.Bind(&iTerms)
	var iMap = helper.StructToMap(iTerms)
	_error := model.UpdateTermsBySlug(Slug_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermsByTermGroupHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermGroup_ := self.Args("term_group").MustInt64()
	var iTerms model.Terms
	self.Bind(&iTerms)
	var iMap = helper.StructToMap(iTerms)
	_error := model.UpdateTermsByTermGroup(TermGroup_, &iMap)
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

func DeleteTermsByTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermId_ := self.Args("term_id").MustInt64()
	_error := model.DeleteTermsByTermId(TermId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTermsByNameHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Name_ := self.Args("name").String()
	_error := model.DeleteTermsByName(Name_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTermsBySlugHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Slug_ := self.Args("slug").String()
	_error := model.DeleteTermsBySlug(Slug_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTermsByTermGroupHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermGroup_ := self.Args("term_group").MustInt64()
	_error := model.DeleteTermsByTermGroup(TermGroup_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
