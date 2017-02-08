package handler

import (
	"github.com/insionng/zenpress/helper"
	"github.com/insionng/zenpress/model"
	"github.com/insionng/macross"
)

func GetTermTaxonomiesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetTermTaxonomiesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["term_taxonomysCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomyCountViaTermTaxonomyIdHandler(self *macross.Context) error {
	TermTaxonomyId_ := self.Args("term_taxonomy_id").MustInt64()
	_int64 := model.GetTermTaxonomyCountViaTermTaxonomyId(TermTaxonomyId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["term_taxonomyCount"] = 0
	}
	m["term_taxonomyCount"] = _int64
	return self.JSON(m)
}

func GetTermTaxonomyCountViaTermIdHandler(self *macross.Context) error {
	TermId_ := self.Args("term_id").MustInt64()
	_int64 := model.GetTermTaxonomyCountViaTermId(TermId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["term_taxonomyCount"] = 0
	}
	m["term_taxonomyCount"] = _int64
	return self.JSON(m)
}

func GetTermTaxonomyCountViaTaxonomyHandler(self *macross.Context) error {
	Taxonomy_ := self.Args("taxonomy").String()
	_int64 := model.GetTermTaxonomyCountViaTaxonomy(Taxonomy_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["term_taxonomyCount"] = 0
	}
	m["term_taxonomyCount"] = _int64
	return self.JSON(m)
}

func GetTermTaxonomyCountViaDescriptionHandler(self *macross.Context) error {
	Description_ := self.Args("description").String()
	_int64 := model.GetTermTaxonomyCountViaDescription(Description_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["term_taxonomyCount"] = 0
	}
	m["term_taxonomyCount"] = _int64
	return self.JSON(m)
}

func GetTermTaxonomyCountViaParentHandler(self *macross.Context) error {
	Parent_ := self.Args("parent").MustInt64()
	_int64 := model.GetTermTaxonomyCountViaParent(Parent_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["term_taxonomyCount"] = 0
	}
	m["term_taxonomyCount"] = _int64
	return self.JSON(m)
}

func GetTermTaxonomyCountViaCountHandler(self *macross.Context) error {
	Count_ := self.Args("count").MustInt64()
	_int64 := model.GetTermTaxonomyCountViaCount(Count_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["term_taxonomyCount"] = 0
	}
	m["term_taxonomyCount"] = _int64
	return self.JSON(m)
}

func GetTermTaxonomiesViaTermTaxonomyIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTermTaxonomyId := self.Args("term_taxonomy_id").MustInt64()
	if (offset > 0) && helper.IsHas(iTermTaxonomyId) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesViaTermTaxonomyId(offset, limit, iTermTaxonomyId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesViaTermTaxonomyId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesViaTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTermId := self.Args("term_id").MustInt64()
	if (offset > 0) && helper.IsHas(iTermId) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesViaTermId(offset, limit, iTermId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesViaTermId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesViaTaxonomyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTaxonomy := self.Args("taxonomy").String()
	if (offset > 0) && helper.IsHas(iTaxonomy) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesViaTaxonomy(offset, limit, iTaxonomy, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesViaTaxonomy's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesViaDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDescription := self.Args("description").String()
	if (offset > 0) && helper.IsHas(iDescription) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesViaDescription(offset, limit, iDescription, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesViaDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesViaParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iParent := self.Args("parent").MustInt64()
	if (offset > 0) && helper.IsHas(iParent) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesViaParent(offset, limit, iParent, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesViaParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesViaCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCount := self.Args("count").MustInt64()
	if (offset > 0) && helper.IsHas(iCount) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesViaCount(offset, limit, iCount, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesViaCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTermTaxonomyIdAndTermIdAndTaxonomyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermTaxonomyId := self.Args("term_taxonomy_id").MustInt64()
	iTermId := self.Args("term_id").MustInt64()
	iTaxonomy := self.Args("taxonomy").String()

	if helper.IsHas(iTermTaxonomyId) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTermTaxonomyIdAndTermIdAndTaxonomy(offset, limit, iTermTaxonomyId,iTermId,iTaxonomy)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTermTaxonomyIdAndTermIdAndTaxonomy's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTermTaxonomyIdAndTermIdAndDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermTaxonomyId := self.Args("term_taxonomy_id").MustInt64()
	iTermId := self.Args("term_id").MustInt64()
	iDescription := self.Args("description").String()

	if helper.IsHas(iTermTaxonomyId) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTermTaxonomyIdAndTermIdAndDescription(offset, limit, iTermTaxonomyId,iTermId,iDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTermTaxonomyIdAndTermIdAndDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTermTaxonomyIdAndTermIdAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermTaxonomyId := self.Args("term_taxonomy_id").MustInt64()
	iTermId := self.Args("term_id").MustInt64()
	iParent := self.Args("parent").MustInt64()

	if helper.IsHas(iTermTaxonomyId) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTermTaxonomyIdAndTermIdAndParent(offset, limit, iTermTaxonomyId,iTermId,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTermTaxonomyIdAndTermIdAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTermTaxonomyIdAndTermIdAndCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermTaxonomyId := self.Args("term_taxonomy_id").MustInt64()
	iTermId := self.Args("term_id").MustInt64()
	iCount := self.Args("count").MustInt64()

	if helper.IsHas(iTermTaxonomyId) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTermTaxonomyIdAndTermIdAndCount(offset, limit, iTermTaxonomyId,iTermId,iCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTermTaxonomyIdAndTermIdAndCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTermTaxonomyIdAndTaxonomyAndDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermTaxonomyId := self.Args("term_taxonomy_id").MustInt64()
	iTaxonomy := self.Args("taxonomy").String()
	iDescription := self.Args("description").String()

	if helper.IsHas(iTermTaxonomyId) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTermTaxonomyIdAndTaxonomyAndDescription(offset, limit, iTermTaxonomyId,iTaxonomy,iDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTermTaxonomyIdAndTaxonomyAndDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTermTaxonomyIdAndTaxonomyAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermTaxonomyId := self.Args("term_taxonomy_id").MustInt64()
	iTaxonomy := self.Args("taxonomy").String()
	iParent := self.Args("parent").MustInt64()

	if helper.IsHas(iTermTaxonomyId) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTermTaxonomyIdAndTaxonomyAndParent(offset, limit, iTermTaxonomyId,iTaxonomy,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTermTaxonomyIdAndTaxonomyAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTermTaxonomyIdAndTaxonomyAndCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermTaxonomyId := self.Args("term_taxonomy_id").MustInt64()
	iTaxonomy := self.Args("taxonomy").String()
	iCount := self.Args("count").MustInt64()

	if helper.IsHas(iTermTaxonomyId) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTermTaxonomyIdAndTaxonomyAndCount(offset, limit, iTermTaxonomyId,iTaxonomy,iCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTermTaxonomyIdAndTaxonomyAndCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTermTaxonomyIdAndDescriptionAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermTaxonomyId := self.Args("term_taxonomy_id").MustInt64()
	iDescription := self.Args("description").String()
	iParent := self.Args("parent").MustInt64()

	if helper.IsHas(iTermTaxonomyId) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTermTaxonomyIdAndDescriptionAndParent(offset, limit, iTermTaxonomyId,iDescription,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTermTaxonomyIdAndDescriptionAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTermTaxonomyIdAndDescriptionAndCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermTaxonomyId := self.Args("term_taxonomy_id").MustInt64()
	iDescription := self.Args("description").String()
	iCount := self.Args("count").MustInt64()

	if helper.IsHas(iTermTaxonomyId) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTermTaxonomyIdAndDescriptionAndCount(offset, limit, iTermTaxonomyId,iDescription,iCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTermTaxonomyIdAndDescriptionAndCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTermTaxonomyIdAndParentAndCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermTaxonomyId := self.Args("term_taxonomy_id").MustInt64()
	iParent := self.Args("parent").MustInt64()
	iCount := self.Args("count").MustInt64()

	if helper.IsHas(iTermTaxonomyId) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTermTaxonomyIdAndParentAndCount(offset, limit, iTermTaxonomyId,iParent,iCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTermTaxonomyIdAndParentAndCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTermIdAndTaxonomyAndDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermId := self.Args("term_id").MustInt64()
	iTaxonomy := self.Args("taxonomy").String()
	iDescription := self.Args("description").String()

	if helper.IsHas(iTermId) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTermIdAndTaxonomyAndDescription(offset, limit, iTermId,iTaxonomy,iDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTermIdAndTaxonomyAndDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTermIdAndTaxonomyAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermId := self.Args("term_id").MustInt64()
	iTaxonomy := self.Args("taxonomy").String()
	iParent := self.Args("parent").MustInt64()

	if helper.IsHas(iTermId) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTermIdAndTaxonomyAndParent(offset, limit, iTermId,iTaxonomy,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTermIdAndTaxonomyAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTermIdAndTaxonomyAndCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermId := self.Args("term_id").MustInt64()
	iTaxonomy := self.Args("taxonomy").String()
	iCount := self.Args("count").MustInt64()

	if helper.IsHas(iTermId) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTermIdAndTaxonomyAndCount(offset, limit, iTermId,iTaxonomy,iCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTermIdAndTaxonomyAndCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTermIdAndDescriptionAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermId := self.Args("term_id").MustInt64()
	iDescription := self.Args("description").String()
	iParent := self.Args("parent").MustInt64()

	if helper.IsHas(iTermId) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTermIdAndDescriptionAndParent(offset, limit, iTermId,iDescription,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTermIdAndDescriptionAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTermIdAndDescriptionAndCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermId := self.Args("term_id").MustInt64()
	iDescription := self.Args("description").String()
	iCount := self.Args("count").MustInt64()

	if helper.IsHas(iTermId) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTermIdAndDescriptionAndCount(offset, limit, iTermId,iDescription,iCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTermIdAndDescriptionAndCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTermIdAndParentAndCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermId := self.Args("term_id").MustInt64()
	iParent := self.Args("parent").MustInt64()
	iCount := self.Args("count").MustInt64()

	if helper.IsHas(iTermId) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTermIdAndParentAndCount(offset, limit, iTermId,iParent,iCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTermIdAndParentAndCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTaxonomyAndDescriptionAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTaxonomy := self.Args("taxonomy").String()
	iDescription := self.Args("description").String()
	iParent := self.Args("parent").MustInt64()

	if helper.IsHas(iTaxonomy) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTaxonomyAndDescriptionAndParent(offset, limit, iTaxonomy,iDescription,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTaxonomyAndDescriptionAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTaxonomyAndDescriptionAndCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTaxonomy := self.Args("taxonomy").String()
	iDescription := self.Args("description").String()
	iCount := self.Args("count").MustInt64()

	if helper.IsHas(iTaxonomy) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTaxonomyAndDescriptionAndCount(offset, limit, iTaxonomy,iDescription,iCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTaxonomyAndDescriptionAndCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTaxonomyAndParentAndCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTaxonomy := self.Args("taxonomy").String()
	iParent := self.Args("parent").MustInt64()
	iCount := self.Args("count").MustInt64()

	if helper.IsHas(iTaxonomy) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTaxonomyAndParentAndCount(offset, limit, iTaxonomy,iParent,iCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTaxonomyAndParentAndCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByDescriptionAndParentAndCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iParent := self.Args("parent").MustInt64()
	iCount := self.Args("count").MustInt64()

	if helper.IsHas(iDescription) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByDescriptionAndParentAndCount(offset, limit, iDescription,iParent,iCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByDescriptionAndParentAndCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTermTaxonomyIdAndTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermTaxonomyId := self.Args("term_taxonomy_id").MustInt64()
	iTermId := self.Args("term_id").MustInt64()

	if helper.IsHas(iTermTaxonomyId) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTermTaxonomyIdAndTermId(offset, limit, iTermTaxonomyId,iTermId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTermTaxonomyIdAndTermId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTermTaxonomyIdAndTaxonomyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermTaxonomyId := self.Args("term_taxonomy_id").MustInt64()
	iTaxonomy := self.Args("taxonomy").String()

	if helper.IsHas(iTermTaxonomyId) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTermTaxonomyIdAndTaxonomy(offset, limit, iTermTaxonomyId,iTaxonomy)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTermTaxonomyIdAndTaxonomy's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTermTaxonomyIdAndDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermTaxonomyId := self.Args("term_taxonomy_id").MustInt64()
	iDescription := self.Args("description").String()

	if helper.IsHas(iTermTaxonomyId) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTermTaxonomyIdAndDescription(offset, limit, iTermTaxonomyId,iDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTermTaxonomyIdAndDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTermTaxonomyIdAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermTaxonomyId := self.Args("term_taxonomy_id").MustInt64()
	iParent := self.Args("parent").MustInt64()

	if helper.IsHas(iTermTaxonomyId) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTermTaxonomyIdAndParent(offset, limit, iTermTaxonomyId,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTermTaxonomyIdAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTermTaxonomyIdAndCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermTaxonomyId := self.Args("term_taxonomy_id").MustInt64()
	iCount := self.Args("count").MustInt64()

	if helper.IsHas(iTermTaxonomyId) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTermTaxonomyIdAndCount(offset, limit, iTermTaxonomyId,iCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTermTaxonomyIdAndCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTermIdAndTaxonomyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermId := self.Args("term_id").MustInt64()
	iTaxonomy := self.Args("taxonomy").String()

	if helper.IsHas(iTermId) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTermIdAndTaxonomy(offset, limit, iTermId,iTaxonomy)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTermIdAndTaxonomy's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTermIdAndDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermId := self.Args("term_id").MustInt64()
	iDescription := self.Args("description").String()

	if helper.IsHas(iTermId) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTermIdAndDescription(offset, limit, iTermId,iDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTermIdAndDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTermIdAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermId := self.Args("term_id").MustInt64()
	iParent := self.Args("parent").MustInt64()

	if helper.IsHas(iTermId) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTermIdAndParent(offset, limit, iTermId,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTermIdAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTermIdAndCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermId := self.Args("term_id").MustInt64()
	iCount := self.Args("count").MustInt64()

	if helper.IsHas(iTermId) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTermIdAndCount(offset, limit, iTermId,iCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTermIdAndCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTaxonomyAndDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTaxonomy := self.Args("taxonomy").String()
	iDescription := self.Args("description").String()

	if helper.IsHas(iTaxonomy) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTaxonomyAndDescription(offset, limit, iTaxonomy,iDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTaxonomyAndDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTaxonomyAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTaxonomy := self.Args("taxonomy").String()
	iParent := self.Args("parent").MustInt64()

	if helper.IsHas(iTaxonomy) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTaxonomyAndParent(offset, limit, iTaxonomy,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTaxonomyAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTaxonomyAndCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTaxonomy := self.Args("taxonomy").String()
	iCount := self.Args("count").MustInt64()

	if helper.IsHas(iTaxonomy) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTaxonomyAndCount(offset, limit, iTaxonomy,iCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTaxonomyAndCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByDescriptionAndParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iParent := self.Args("parent").MustInt64()

	if helper.IsHas(iDescription) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByDescriptionAndParent(offset, limit, iDescription,iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByDescriptionAndParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByDescriptionAndCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iDescription := self.Args("description").String()
	iCount := self.Args("count").MustInt64()

	if helper.IsHas(iDescription) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByDescriptionAndCount(offset, limit, iDescription,iCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByDescriptionAndCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByParentAndCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iParent := self.Args("parent").MustInt64()
	iCount := self.Args("count").MustInt64()

	if helper.IsHas(iParent) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByParentAndCount(offset, limit, iParent,iCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByParentAndCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_TermTaxonomy, _error := model.GetTermTaxonomies(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomies' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTermTaxonomyViaTermTaxonomyIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTermTaxonomyId := self.Args("term_taxonomy_id").MustInt64()
	if helper.IsHas(iTermTaxonomyId) {
		_TermTaxonomy := model.HasTermTaxonomyViaTermTaxonomyId(iTermTaxonomyId)
		var m = map[string]interface{}{}
		m["term_taxonomy"] = _TermTaxonomy
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermTaxonomyViaTermTaxonomyId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTermTaxonomyViaTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTermId := self.Args("term_id").MustInt64()
	if helper.IsHas(iTermId) {
		_TermTaxonomy := model.HasTermTaxonomyViaTermId(iTermId)
		var m = map[string]interface{}{}
		m["term_taxonomy"] = _TermTaxonomy
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermTaxonomyViaTermId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTermTaxonomyViaTaxonomyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTaxonomy := self.Args("taxonomy").String()
	if helper.IsHas(iTaxonomy) {
		_TermTaxonomy := model.HasTermTaxonomyViaTaxonomy(iTaxonomy)
		var m = map[string]interface{}{}
		m["term_taxonomy"] = _TermTaxonomy
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermTaxonomyViaTaxonomy's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTermTaxonomyViaDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDescription := self.Args("description").String()
	if helper.IsHas(iDescription) {
		_TermTaxonomy := model.HasTermTaxonomyViaDescription(iDescription)
		var m = map[string]interface{}{}
		m["term_taxonomy"] = _TermTaxonomy
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermTaxonomyViaDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTermTaxonomyViaParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iParent := self.Args("parent").MustInt64()
	if helper.IsHas(iParent) {
		_TermTaxonomy := model.HasTermTaxonomyViaParent(iParent)
		var m = map[string]interface{}{}
		m["term_taxonomy"] = _TermTaxonomy
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermTaxonomyViaParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTermTaxonomyViaCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCount := self.Args("count").MustInt64()
	if helper.IsHas(iCount) {
		_TermTaxonomy := model.HasTermTaxonomyViaCount(iCount)
		var m = map[string]interface{}{}
		m["term_taxonomy"] = _TermTaxonomy
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermTaxonomyViaCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomyViaTermTaxonomyIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTermTaxonomyId := self.Args("term_taxonomy_id").MustInt64()
	if helper.IsHas(iTermTaxonomyId) {
		_TermTaxonomy, _error := model.GetTermTaxonomyViaTermTaxonomyId(iTermTaxonomyId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomyViaTermTaxonomyId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomyViaTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTermId := self.Args("term_id").MustInt64()
	if helper.IsHas(iTermId) {
		_TermTaxonomy, _error := model.GetTermTaxonomyViaTermId(iTermId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomyViaTermId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomyViaTaxonomyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTaxonomy := self.Args("taxonomy").String()
	if helper.IsHas(iTaxonomy) {
		_TermTaxonomy, _error := model.GetTermTaxonomyViaTaxonomy(iTaxonomy)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomyViaTaxonomy's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomyViaDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDescription := self.Args("description").String()
	if helper.IsHas(iDescription) {
		_TermTaxonomy, _error := model.GetTermTaxonomyViaDescription(iDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomyViaDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomyViaParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iParent := self.Args("parent").MustInt64()
	if helper.IsHas(iParent) {
		_TermTaxonomy, _error := model.GetTermTaxonomyViaParent(iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomyViaParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomyViaCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCount := self.Args("count").MustInt64()
	if helper.IsHas(iCount) {
		_TermTaxonomy, _error := model.GetTermTaxonomyViaCount(iCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomyViaCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermTaxonomyViaTermTaxonomyIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermTaxonomyId_ := self.Args("term_taxonomy_id").MustInt64()
	if helper.IsHas(TermTaxonomyId_) {
		var iTermTaxonomy model.TermTaxonomy
		self.Bind(&iTermTaxonomy)
		_TermTaxonomy, _error := model.SetTermTaxonomyViaTermTaxonomyId(TermTaxonomyId_, &iTermTaxonomy)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the SetTermTaxonomyViaTermTaxonomyId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermTaxonomyViaTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermId_ := self.Args("term_id").MustInt64()
	if helper.IsHas(TermId_) {
		var iTermTaxonomy model.TermTaxonomy
		self.Bind(&iTermTaxonomy)
		_TermTaxonomy, _error := model.SetTermTaxonomyViaTermId(TermId_, &iTermTaxonomy)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the SetTermTaxonomyViaTermId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermTaxonomyViaTaxonomyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Taxonomy_ := self.Args("taxonomy").String()
	if helper.IsHas(Taxonomy_) {
		var iTermTaxonomy model.TermTaxonomy
		self.Bind(&iTermTaxonomy)
		_TermTaxonomy, _error := model.SetTermTaxonomyViaTaxonomy(Taxonomy_, &iTermTaxonomy)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the SetTermTaxonomyViaTaxonomy's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermTaxonomyViaDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Description_ := self.Args("description").String()
	if helper.IsHas(Description_) {
		var iTermTaxonomy model.TermTaxonomy
		self.Bind(&iTermTaxonomy)
		_TermTaxonomy, _error := model.SetTermTaxonomyViaDescription(Description_, &iTermTaxonomy)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the SetTermTaxonomyViaDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermTaxonomyViaParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Parent_ := self.Args("parent").MustInt64()
	if helper.IsHas(Parent_) {
		var iTermTaxonomy model.TermTaxonomy
		self.Bind(&iTermTaxonomy)
		_TermTaxonomy, _error := model.SetTermTaxonomyViaParent(Parent_, &iTermTaxonomy)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the SetTermTaxonomyViaParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermTaxonomyViaCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Count_ := self.Args("count").MustInt64()
	if helper.IsHas(Count_) {
		var iTermTaxonomy model.TermTaxonomy
		self.Bind(&iTermTaxonomy)
		_TermTaxonomy, _error := model.SetTermTaxonomyViaCount(Count_, &iTermTaxonomy)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the SetTermTaxonomyViaCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddTermTaxonomyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermTaxonomyId_ := self.Args("term_taxonomy_id").MustInt64()
	TermId_ := self.Args("term_id").MustInt64()
	Taxonomy_ := self.Args("taxonomy").String()
	Description_ := self.Args("description").String()
	Parent_ := self.Args("parent").MustInt64()
	Count_ := self.Args("count").MustInt64()

	if helper.IsHas(TermTaxonomyId_) {
		_error := model.AddTermTaxonomy(TermTaxonomyId_,TermId_,Taxonomy_,Description_,Parent_,Count_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddTermTaxonomy's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostTermTaxonomyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iTermTaxonomy model.TermTaxonomy
	self.Bind(&iTermTaxonomy)
	_int64, _error := model.PostTermTaxonomy(&iTermTaxonomy)
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

func PutTermTaxonomyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iTermTaxonomy model.TermTaxonomy
	self.Bind(&iTermTaxonomy)
	_int64, _error := model.PutTermTaxonomy(&iTermTaxonomy)
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

func PutTermTaxonomyViaTermTaxonomyIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermTaxonomyId_ := self.Args("term_taxonomy_id").MustInt64()
	var iTermTaxonomy model.TermTaxonomy
	self.Bind(&iTermTaxonomy)
	_int64, _error := model.PutTermTaxonomyViaTermTaxonomyId(TermTaxonomyId_, &iTermTaxonomy)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTermTaxonomyViaTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermId_ := self.Args("term_id").MustInt64()
	var iTermTaxonomy model.TermTaxonomy
	self.Bind(&iTermTaxonomy)
	_int64, _error := model.PutTermTaxonomyViaTermId(TermId_, &iTermTaxonomy)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTermTaxonomyViaTaxonomyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Taxonomy_ := self.Args("taxonomy").String()
	var iTermTaxonomy model.TermTaxonomy
	self.Bind(&iTermTaxonomy)
	_int64, _error := model.PutTermTaxonomyViaTaxonomy(Taxonomy_, &iTermTaxonomy)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTermTaxonomyViaDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Description_ := self.Args("description").String()
	var iTermTaxonomy model.TermTaxonomy
	self.Bind(&iTermTaxonomy)
	_int64, _error := model.PutTermTaxonomyViaDescription(Description_, &iTermTaxonomy)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTermTaxonomyViaParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Parent_ := self.Args("parent").MustInt64()
	var iTermTaxonomy model.TermTaxonomy
	self.Bind(&iTermTaxonomy)
	_int64, _error := model.PutTermTaxonomyViaParent(Parent_, &iTermTaxonomy)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTermTaxonomyViaCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Count_ := self.Args("count").MustInt64()
	var iTermTaxonomy model.TermTaxonomy
	self.Bind(&iTermTaxonomy)
	_int64, _error := model.PutTermTaxonomyViaCount(Count_, &iTermTaxonomy)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermTaxonomyViaTermTaxonomyIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermTaxonomyId_ := self.Args("term_taxonomy_id").MustInt64()
	var iTermTaxonomy model.TermTaxonomy
	self.Bind(&iTermTaxonomy)
	var iMap = helper.StructToMap(iTermTaxonomy)
	_error := model.UpdateTermTaxonomyViaTermTaxonomyId(TermTaxonomyId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermTaxonomyViaTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermId_ := self.Args("term_id").MustInt64()
	var iTermTaxonomy model.TermTaxonomy
	self.Bind(&iTermTaxonomy)
	var iMap = helper.StructToMap(iTermTaxonomy)
	_error := model.UpdateTermTaxonomyViaTermId(TermId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermTaxonomyViaTaxonomyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Taxonomy_ := self.Args("taxonomy").String()
	var iTermTaxonomy model.TermTaxonomy
	self.Bind(&iTermTaxonomy)
	var iMap = helper.StructToMap(iTermTaxonomy)
	_error := model.UpdateTermTaxonomyViaTaxonomy(Taxonomy_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermTaxonomyViaDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Description_ := self.Args("description").String()
	var iTermTaxonomy model.TermTaxonomy
	self.Bind(&iTermTaxonomy)
	var iMap = helper.StructToMap(iTermTaxonomy)
	_error := model.UpdateTermTaxonomyViaDescription(Description_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermTaxonomyViaParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Parent_ := self.Args("parent").MustInt64()
	var iTermTaxonomy model.TermTaxonomy
	self.Bind(&iTermTaxonomy)
	var iMap = helper.StructToMap(iTermTaxonomy)
	_error := model.UpdateTermTaxonomyViaParent(Parent_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermTaxonomyViaCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Count_ := self.Args("count").MustInt64()
	var iTermTaxonomy model.TermTaxonomy
	self.Bind(&iTermTaxonomy)
	var iMap = helper.StructToMap(iTermTaxonomy)
	_error := model.UpdateTermTaxonomyViaCount(Count_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTermTaxonomyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermTaxonomyId_ := self.Args("term_taxonomy_id").MustInt64()
	_int64, _error := model.DeleteTermTaxonomy(TermTaxonomyId_)
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

func DeleteTermTaxonomyViaTermTaxonomyIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermTaxonomyId_ := self.Args("term_taxonomy_id").MustInt64()
	_error := model.DeleteTermTaxonomyViaTermTaxonomyId(TermTaxonomyId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTermTaxonomyViaTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermId_ := self.Args("term_id").MustInt64()
	_error := model.DeleteTermTaxonomyViaTermId(TermId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTermTaxonomyViaTaxonomyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Taxonomy_ := self.Args("taxonomy").String()
	_error := model.DeleteTermTaxonomyViaTaxonomy(Taxonomy_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTermTaxonomyViaDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Description_ := self.Args("description").String()
	_error := model.DeleteTermTaxonomyViaDescription(Description_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTermTaxonomyViaParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Parent_ := self.Args("parent").MustInt64()
	_error := model.DeleteTermTaxonomyViaParent(Parent_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTermTaxonomyViaCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Count_ := self.Args("count").MustInt64()
	_error := model.DeleteTermTaxonomyViaCount(Count_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
