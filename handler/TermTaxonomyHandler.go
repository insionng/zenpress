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

func GetTermTaxonomyCountByTermTaxonomyIdHandler(self *macross.Context) error {
	TermTaxonomyId_ := self.Args("term_taxonomy_id").MustInt64()
	_int64 := model.GetTermTaxonomyCountByTermTaxonomyId(TermTaxonomyId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["term_taxonomyCount"] = 0
	}
	m["term_taxonomyCount"] = _int64
	return self.JSON(m)
}

func GetTermTaxonomyCountByTermIdHandler(self *macross.Context) error {
	TermId_ := self.Args("term_id").MustInt64()
	_int64 := model.GetTermTaxonomyCountByTermId(TermId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["term_taxonomyCount"] = 0
	}
	m["term_taxonomyCount"] = _int64
	return self.JSON(m)
}

func GetTermTaxonomyCountByTaxonomyHandler(self *macross.Context) error {
	Taxonomy_ := self.Args("taxonomy").String()
	_int64 := model.GetTermTaxonomyCountByTaxonomy(Taxonomy_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["term_taxonomyCount"] = 0
	}
	m["term_taxonomyCount"] = _int64
	return self.JSON(m)
}

func GetTermTaxonomyCountByDescriptionHandler(self *macross.Context) error {
	Description_ := self.Args("description").String()
	_int64 := model.GetTermTaxonomyCountByDescription(Description_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["term_taxonomyCount"] = 0
	}
	m["term_taxonomyCount"] = _int64
	return self.JSON(m)
}

func GetTermTaxonomyCountByParentHandler(self *macross.Context) error {
	Parent_ := self.Args("parent").MustInt64()
	_int64 := model.GetTermTaxonomyCountByParent(Parent_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["term_taxonomyCount"] = 0
	}
	m["term_taxonomyCount"] = _int64
	return self.JSON(m)
}

func GetTermTaxonomyCountByCountHandler(self *macross.Context) error {
	Count_ := self.Args("count").MustInt64()
	_int64 := model.GetTermTaxonomyCountByCount(Count_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["term_taxonomyCount"] = 0
	}
	m["term_taxonomyCount"] = _int64
	return self.JSON(m)
}

func GetTermTaxonomiesByTermTaxonomyIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTermTaxonomyId := self.Args("term_taxonomy_id").MustInt64()
	if (offset > 0) && helper.IsHas(iTermTaxonomyId) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTermTaxonomyId(offset, limit, iTermTaxonomyId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTermTaxonomyId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTermId := self.Args("term_id").MustInt64()
	if (offset > 0) && helper.IsHas(iTermId) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTermId(offset, limit, iTermId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTermId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByTaxonomyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTaxonomy := self.Args("taxonomy").String()
	if (offset > 0) && helper.IsHas(iTaxonomy) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByTaxonomy(offset, limit, iTaxonomy, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByTaxonomy's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iDescription := self.Args("description").String()
	if (offset > 0) && helper.IsHas(iDescription) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByDescription(offset, limit, iDescription, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iParent := self.Args("parent").MustInt64()
	if (offset > 0) && helper.IsHas(iParent) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByParent(offset, limit, iParent, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomiesByCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iCount := self.Args("count").MustInt64()
	if (offset > 0) && helper.IsHas(iCount) {
		_TermTaxonomy, _error := model.GetTermTaxonomiesByCount(offset, limit, iCount, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomiesByCount's args."
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

func GetHasTermTaxonomyByTermTaxonomyIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTermTaxonomyId := self.Args("term_taxonomy_id").MustInt64()
	if helper.IsHas(iTermTaxonomyId) {
		_TermTaxonomy := model.HasTermTaxonomyByTermTaxonomyId(iTermTaxonomyId)
		var m = map[string]interface{}{}
		m["term_taxonomy"] = _TermTaxonomy
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermTaxonomyByTermTaxonomyId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTermTaxonomyByTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTermId := self.Args("term_id").MustInt64()
	if helper.IsHas(iTermId) {
		_TermTaxonomy := model.HasTermTaxonomyByTermId(iTermId)
		var m = map[string]interface{}{}
		m["term_taxonomy"] = _TermTaxonomy
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermTaxonomyByTermId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTermTaxonomyByTaxonomyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTaxonomy := self.Args("taxonomy").String()
	if helper.IsHas(iTaxonomy) {
		_TermTaxonomy := model.HasTermTaxonomyByTaxonomy(iTaxonomy)
		var m = map[string]interface{}{}
		m["term_taxonomy"] = _TermTaxonomy
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermTaxonomyByTaxonomy's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTermTaxonomyByDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDescription := self.Args("description").String()
	if helper.IsHas(iDescription) {
		_TermTaxonomy := model.HasTermTaxonomyByDescription(iDescription)
		var m = map[string]interface{}{}
		m["term_taxonomy"] = _TermTaxonomy
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermTaxonomyByDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTermTaxonomyByParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iParent := self.Args("parent").MustInt64()
	if helper.IsHas(iParent) {
		_TermTaxonomy := model.HasTermTaxonomyByParent(iParent)
		var m = map[string]interface{}{}
		m["term_taxonomy"] = _TermTaxonomy
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermTaxonomyByParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTermTaxonomyByCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCount := self.Args("count").MustInt64()
	if helper.IsHas(iCount) {
		_TermTaxonomy := model.HasTermTaxonomyByCount(iCount)
		var m = map[string]interface{}{}
		m["term_taxonomy"] = _TermTaxonomy
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermTaxonomyByCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomyByTermTaxonomyIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTermTaxonomyId := self.Args("term_taxonomy_id").MustInt64()
	if helper.IsHas(iTermTaxonomyId) {
		_TermTaxonomy, _error := model.GetTermTaxonomyByTermTaxonomyId(iTermTaxonomyId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomyByTermTaxonomyId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomyByTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTermId := self.Args("term_id").MustInt64()
	if helper.IsHas(iTermId) {
		_TermTaxonomy, _error := model.GetTermTaxonomyByTermId(iTermId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomyByTermId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomyByTaxonomyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTaxonomy := self.Args("taxonomy").String()
	if helper.IsHas(iTaxonomy) {
		_TermTaxonomy, _error := model.GetTermTaxonomyByTaxonomy(iTaxonomy)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomyByTaxonomy's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomyByDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iDescription := self.Args("description").String()
	if helper.IsHas(iDescription) {
		_TermTaxonomy, _error := model.GetTermTaxonomyByDescription(iDescription)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomyByDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomyByParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iParent := self.Args("parent").MustInt64()
	if helper.IsHas(iParent) {
		_TermTaxonomy, _error := model.GetTermTaxonomyByParent(iParent)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomyByParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermTaxonomyByCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iCount := self.Args("count").MustInt64()
	if helper.IsHas(iCount) {
		_TermTaxonomy, _error := model.GetTermTaxonomyByCount(iCount)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the GetTermTaxonomyByCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermTaxonomyByTermTaxonomyIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermTaxonomyId_ := self.Args("term_taxonomy_id").MustInt64()
	if helper.IsHas(TermTaxonomyId_) {
		var iTermTaxonomy model.TermTaxonomy
		self.Bind(&iTermTaxonomy)
		_TermTaxonomy, _error := model.SetTermTaxonomyByTermTaxonomyId(TermTaxonomyId_, &iTermTaxonomy)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the SetTermTaxonomyByTermTaxonomyId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermTaxonomyByTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermId_ := self.Args("term_id").MustInt64()
	if helper.IsHas(TermId_) {
		var iTermTaxonomy model.TermTaxonomy
		self.Bind(&iTermTaxonomy)
		_TermTaxonomy, _error := model.SetTermTaxonomyByTermId(TermId_, &iTermTaxonomy)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the SetTermTaxonomyByTermId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermTaxonomyByTaxonomyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Taxonomy_ := self.Args("taxonomy").String()
	if helper.IsHas(Taxonomy_) {
		var iTermTaxonomy model.TermTaxonomy
		self.Bind(&iTermTaxonomy)
		_TermTaxonomy, _error := model.SetTermTaxonomyByTaxonomy(Taxonomy_, &iTermTaxonomy)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the SetTermTaxonomyByTaxonomy's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermTaxonomyByDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Description_ := self.Args("description").String()
	if helper.IsHas(Description_) {
		var iTermTaxonomy model.TermTaxonomy
		self.Bind(&iTermTaxonomy)
		_TermTaxonomy, _error := model.SetTermTaxonomyByDescription(Description_, &iTermTaxonomy)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the SetTermTaxonomyByDescription's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermTaxonomyByParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Parent_ := self.Args("parent").MustInt64()
	if helper.IsHas(Parent_) {
		var iTermTaxonomy model.TermTaxonomy
		self.Bind(&iTermTaxonomy)
		_TermTaxonomy, _error := model.SetTermTaxonomyByParent(Parent_, &iTermTaxonomy)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the SetTermTaxonomyByParent's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermTaxonomyByCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Count_ := self.Args("count").MustInt64()
	if helper.IsHas(Count_) {
		var iTermTaxonomy model.TermTaxonomy
		self.Bind(&iTermTaxonomy)
		_TermTaxonomy, _error := model.SetTermTaxonomyByCount(Count_, &iTermTaxonomy)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermTaxonomy)
	}
	herr.Message = "Can't get to the SetTermTaxonomyByCount's args."
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

func PutTermTaxonomyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iTermTaxonomy model.TermTaxonomy
	self.Bind(&iTermTaxonomy)
	_int64, _error := model.PutTermTaxonomy(&iTermTaxonomy)
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

func PutTermTaxonomyByTermTaxonomyIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermTaxonomyId_ := self.Args("term_taxonomy_id").MustInt64()
	var iTermTaxonomy model.TermTaxonomy
	self.Bind(&iTermTaxonomy)
	_int64, _error := model.PutTermTaxonomyByTermTaxonomyId(TermTaxonomyId_, &iTermTaxonomy)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTermTaxonomyByTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermId_ := self.Args("term_id").MustInt64()
	var iTermTaxonomy model.TermTaxonomy
	self.Bind(&iTermTaxonomy)
	_int64, _error := model.PutTermTaxonomyByTermId(TermId_, &iTermTaxonomy)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTermTaxonomyByTaxonomyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Taxonomy_ := self.Args("taxonomy").String()
	var iTermTaxonomy model.TermTaxonomy
	self.Bind(&iTermTaxonomy)
	_int64, _error := model.PutTermTaxonomyByTaxonomy(Taxonomy_, &iTermTaxonomy)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTermTaxonomyByDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Description_ := self.Args("description").String()
	var iTermTaxonomy model.TermTaxonomy
	self.Bind(&iTermTaxonomy)
	_int64, _error := model.PutTermTaxonomyByDescription(Description_, &iTermTaxonomy)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTermTaxonomyByParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Parent_ := self.Args("parent").MustInt64()
	var iTermTaxonomy model.TermTaxonomy
	self.Bind(&iTermTaxonomy)
	_int64, _error := model.PutTermTaxonomyByParent(Parent_, &iTermTaxonomy)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTermTaxonomyByCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Count_ := self.Args("count").MustInt64()
	var iTermTaxonomy model.TermTaxonomy
	self.Bind(&iTermTaxonomy)
	_int64, _error := model.PutTermTaxonomyByCount(Count_, &iTermTaxonomy)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermTaxonomyByTermTaxonomyIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermTaxonomyId_ := self.Args("term_taxonomy_id").MustInt64()
	var iTermTaxonomy model.TermTaxonomy
	self.Bind(&iTermTaxonomy)
	var iMap = helper.StructToMap(iTermTaxonomy)
	_error := model.UpdateTermTaxonomyByTermTaxonomyId(TermTaxonomyId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermTaxonomyByTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermId_ := self.Args("term_id").MustInt64()
	var iTermTaxonomy model.TermTaxonomy
	self.Bind(&iTermTaxonomy)
	var iMap = helper.StructToMap(iTermTaxonomy)
	_error := model.UpdateTermTaxonomyByTermId(TermId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermTaxonomyByTaxonomyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Taxonomy_ := self.Args("taxonomy").String()
	var iTermTaxonomy model.TermTaxonomy
	self.Bind(&iTermTaxonomy)
	var iMap = helper.StructToMap(iTermTaxonomy)
	_error := model.UpdateTermTaxonomyByTaxonomy(Taxonomy_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermTaxonomyByDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Description_ := self.Args("description").String()
	var iTermTaxonomy model.TermTaxonomy
	self.Bind(&iTermTaxonomy)
	var iMap = helper.StructToMap(iTermTaxonomy)
	_error := model.UpdateTermTaxonomyByDescription(Description_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermTaxonomyByParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Parent_ := self.Args("parent").MustInt64()
	var iTermTaxonomy model.TermTaxonomy
	self.Bind(&iTermTaxonomy)
	var iMap = helper.StructToMap(iTermTaxonomy)
	_error := model.UpdateTermTaxonomyByParent(Parent_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermTaxonomyByCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Count_ := self.Args("count").MustInt64()
	var iTermTaxonomy model.TermTaxonomy
	self.Bind(&iTermTaxonomy)
	var iMap = helper.StructToMap(iTermTaxonomy)
	_error := model.UpdateTermTaxonomyByCount(Count_, &iMap)
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

func DeleteTermTaxonomyByTermTaxonomyIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermTaxonomyId_ := self.Args("term_taxonomy_id").MustInt64()
	_error := model.DeleteTermTaxonomyByTermTaxonomyId(TermTaxonomyId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTermTaxonomyByTermIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermId_ := self.Args("term_id").MustInt64()
	_error := model.DeleteTermTaxonomyByTermId(TermId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTermTaxonomyByTaxonomyHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Taxonomy_ := self.Args("taxonomy").String()
	_error := model.DeleteTermTaxonomyByTaxonomy(Taxonomy_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTermTaxonomyByDescriptionHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Description_ := self.Args("description").String()
	_error := model.DeleteTermTaxonomyByDescription(Description_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTermTaxonomyByParentHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Parent_ := self.Args("parent").MustInt64()
	_error := model.DeleteTermTaxonomyByParent(Parent_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTermTaxonomyByCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	Count_ := self.Args("count").MustInt64()
	_error := model.DeleteTermTaxonomyByCount(Count_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
