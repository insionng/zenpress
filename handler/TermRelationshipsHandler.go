package handler

import (
	"github.com/insionng/zenpress/helper"
	"github.com/insionng/zenpress/model"
	"github.com/insionng/macross"
)

func GetTermRelationshipsesCountHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()

	if offset > 0 {
		var m = map[string]interface{}{}
		_int64, _error := model.GetTermRelationshipsesCount(offset, limit)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		m["term_relationshipssCount"] = _int64
		return self.JSON(m)
	}
	herr.Message = "Can't get to the GetTermRelationshipsesCount's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermRelationshipsCountByObjectIdHandler(self *macross.Context) error {
	ObjectId_ := self.Args("object_id").MustInt64()
	_int64 := model.GetTermRelationshipsCountByObjectId(ObjectId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["term_relationshipsCount"] = 0
	}
	m["term_relationshipsCount"] = _int64
	return self.JSON(m)
}

func GetTermRelationshipsCountByTermTaxonomyIdHandler(self *macross.Context) error {
	TermTaxonomyId_ := self.Args("term_taxonomy_id").MustInt64()
	_int64 := model.GetTermRelationshipsCountByTermTaxonomyId(TermTaxonomyId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["term_relationshipsCount"] = 0
	}
	m["term_relationshipsCount"] = _int64
	return self.JSON(m)
}

func GetTermRelationshipsCountByTermOrderHandler(self *macross.Context) error {
	TermOrder_ := self.Args("term_order").MustInt()
	_int64 := model.GetTermRelationshipsCountByTermOrder(TermOrder_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["term_relationshipsCount"] = 0
	}
	m["term_relationshipsCount"] = _int64
	return self.JSON(m)
}

func GetTermRelationshipsesByObjectIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iObjectId := self.Args("object_id").MustInt64()
	if (offset > 0) && helper.IsHas(iObjectId) {
		_TermRelationships, _error := model.GetTermRelationshipsesByObjectId(offset, limit, iObjectId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermRelationships)
	}
	herr.Message = "Can't get to the GetTermRelationshipsesByObjectId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermRelationshipsesByTermTaxonomyIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTermTaxonomyId := self.Args("term_taxonomy_id").MustInt64()
	if (offset > 0) && helper.IsHas(iTermTaxonomyId) {
		_TermRelationships, _error := model.GetTermRelationshipsesByTermTaxonomyId(offset, limit, iTermTaxonomyId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermRelationships)
	}
	herr.Message = "Can't get to the GetTermRelationshipsesByTermTaxonomyId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermRelationshipsesByTermOrderHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTermOrder := self.Args("term_order").MustInt()
	if (offset > 0) && helper.IsHas(iTermOrder) {
		_TermRelationships, _error := model.GetTermRelationshipsesByTermOrder(offset, limit, iTermOrder, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermRelationships)
	}
	herr.Message = "Can't get to the GetTermRelationshipsesByTermOrder's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermRelationshipsesByObjectIdAndTermTaxonomyIdAndTermOrderHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iObjectId := self.Args("object_id").MustInt64()
	iTermTaxonomyId := self.Args("term_taxonomy_id").MustInt64()
	iTermOrder := self.Args("term_order").MustInt()

	if helper.IsHas(iObjectId) {
		_TermRelationships, _error := model.GetTermRelationshipsesByObjectIdAndTermTaxonomyIdAndTermOrder(offset, limit, iObjectId,iTermTaxonomyId,iTermOrder)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermRelationships)
	}
	herr.Message = "Can't get to the GetTermRelationshipsesByObjectIdAndTermTaxonomyIdAndTermOrder's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermRelationshipsesByObjectIdAndTermTaxonomyIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iObjectId := self.Args("object_id").MustInt64()
	iTermTaxonomyId := self.Args("term_taxonomy_id").MustInt64()

	if helper.IsHas(iObjectId) {
		_TermRelationships, _error := model.GetTermRelationshipsesByObjectIdAndTermTaxonomyId(offset, limit, iObjectId,iTermTaxonomyId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermRelationships)
	}
	herr.Message = "Can't get to the GetTermRelationshipsesByObjectIdAndTermTaxonomyId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermRelationshipsesByObjectIdAndTermOrderHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iObjectId := self.Args("object_id").MustInt64()
	iTermOrder := self.Args("term_order").MustInt()

	if helper.IsHas(iObjectId) {
		_TermRelationships, _error := model.GetTermRelationshipsesByObjectIdAndTermOrder(offset, limit, iObjectId,iTermOrder)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermRelationships)
	}
	herr.Message = "Can't get to the GetTermRelationshipsesByObjectIdAndTermOrder's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermRelationshipsesByTermTaxonomyIdAndTermOrderHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	iTermTaxonomyId := self.Args("term_taxonomy_id").MustInt64()
	iTermOrder := self.Args("term_order").MustInt()

	if helper.IsHas(iTermTaxonomyId) {
		_TermRelationships, _error := model.GetTermRelationshipsesByTermTaxonomyIdAndTermOrder(offset, limit, iTermTaxonomyId,iTermOrder)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermRelationships)
	}
	herr.Message = "Can't get to the GetTermRelationshipsesByTermTaxonomyIdAndTermOrder's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermRelationshipsesHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	if (offset > 0) && (len(field) > 0) {
		_TermRelationships, _error := model.GetTermRelationshipses(offset, limit, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermRelationships)
	}
	herr.Message = "Can't get to the GetTermRelationshipses' args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTermRelationshipsByObjectIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iObjectId := self.Args("object_id").MustInt64()
	if helper.IsHas(iObjectId) {
		_TermRelationships := model.HasTermRelationshipsByObjectId(iObjectId)
		var m = map[string]interface{}{}
		m["term_relationships"] = _TermRelationships
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermRelationshipsByObjectId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTermRelationshipsByTermTaxonomyIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTermTaxonomyId := self.Args("term_taxonomy_id").MustInt64()
	if helper.IsHas(iTermTaxonomyId) {
		_TermRelationships := model.HasTermRelationshipsByTermTaxonomyId(iTermTaxonomyId)
		var m = map[string]interface{}{}
		m["term_relationships"] = _TermRelationships
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermRelationshipsByTermTaxonomyId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTermRelationshipsByTermOrderHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTermOrder := self.Args("term_order").MustInt()
	if helper.IsHas(iTermOrder) {
		_TermRelationships := model.HasTermRelationshipsByTermOrder(iTermOrder)
		var m = map[string]interface{}{}
		m["term_relationships"] = _TermRelationships
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermRelationshipsByTermOrder's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermRelationshipsByObjectIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iObjectId := self.Args("object_id").MustInt64()
	if helper.IsHas(iObjectId) {
		_TermRelationships, _error := model.GetTermRelationshipsByObjectId(iObjectId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermRelationships)
	}
	herr.Message = "Can't get to the GetTermRelationshipsByObjectId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermRelationshipsByTermTaxonomyIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTermTaxonomyId := self.Args("term_taxonomy_id").MustInt64()
	if helper.IsHas(iTermTaxonomyId) {
		_TermRelationships, _error := model.GetTermRelationshipsByTermTaxonomyId(iTermTaxonomyId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermRelationships)
	}
	herr.Message = "Can't get to the GetTermRelationshipsByTermTaxonomyId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermRelationshipsByTermOrderHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTermOrder := self.Args("term_order").MustInt()
	if helper.IsHas(iTermOrder) {
		_TermRelationships, _error := model.GetTermRelationshipsByTermOrder(iTermOrder)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermRelationships)
	}
	herr.Message = "Can't get to the GetTermRelationshipsByTermOrder's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermRelationshipsByObjectIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	ObjectId_ := self.Args("object_id").MustInt64()
	if helper.IsHas(ObjectId_) {
		var iTermRelationships model.TermRelationships
		self.Bind(&iTermRelationships)
		_TermRelationships, _error := model.SetTermRelationshipsByObjectId(ObjectId_, &iTermRelationships)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermRelationships)
	}
	herr.Message = "Can't get to the SetTermRelationshipsByObjectId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermRelationshipsByTermTaxonomyIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermTaxonomyId_ := self.Args("term_taxonomy_id").MustInt64()
	if helper.IsHas(TermTaxonomyId_) {
		var iTermRelationships model.TermRelationships
		self.Bind(&iTermRelationships)
		_TermRelationships, _error := model.SetTermRelationshipsByTermTaxonomyId(TermTaxonomyId_, &iTermRelationships)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermRelationships)
	}
	herr.Message = "Can't get to the SetTermRelationshipsByTermTaxonomyId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermRelationshipsByTermOrderHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermOrder_ := self.Args("term_order").MustInt()
	if helper.IsHas(TermOrder_) {
		var iTermRelationships model.TermRelationships
		self.Bind(&iTermRelationships)
		_TermRelationships, _error := model.SetTermRelationshipsByTermOrder(TermOrder_, &iTermRelationships)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermRelationships)
	}
	herr.Message = "Can't get to the SetTermRelationshipsByTermOrder's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostAddTermRelationshipsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	ObjectId_ := self.Args("object_id").MustInt64()
	TermTaxonomyId_ := self.Args("term_taxonomy_id").MustInt64()
	TermOrder_ := self.Args("term_order").MustInt()

	if helper.IsHas(ObjectId_) {
		_error := model.AddTermRelationships(ObjectId_,TermTaxonomyId_,TermOrder_)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		herr.Message = "StatusOK"
		herr.Status = macross.StatusOK
		return self.JSON(herr)
	}
	herr.Message = "Can't get to the AddTermRelationships's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostTermRelationshipsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iTermRelationships model.TermRelationships
	self.Bind(&iTermRelationships)
	_int64, _error := model.PostTermRelationships(&iTermRelationships)
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

func PutTermRelationshipsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iTermRelationships model.TermRelationships
	self.Bind(&iTermRelationships)
	_int64, _error := model.PutTermRelationships(&iTermRelationships)
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

func PutTermRelationshipsByObjectIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	ObjectId_ := self.Args("object_id").MustInt64()
	var iTermRelationships model.TermRelationships
	self.Bind(&iTermRelationships)
	_int64, _error := model.PutTermRelationshipsByObjectId(ObjectId_, &iTermRelationships)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTermRelationshipsByTermTaxonomyIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermTaxonomyId_ := self.Args("term_taxonomy_id").MustInt64()
	var iTermRelationships model.TermRelationships
	self.Bind(&iTermRelationships)
	_int64, _error := model.PutTermRelationshipsByTermTaxonomyId(TermTaxonomyId_, &iTermRelationships)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTermRelationshipsByTermOrderHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermOrder_ := self.Args("term_order").MustInt()
	var iTermRelationships model.TermRelationships
	self.Bind(&iTermRelationships)
	_int64, _error := model.PutTermRelationshipsByTermOrder(TermOrder_, &iTermRelationships)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermRelationshipsByObjectIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	ObjectId_ := self.Args("object_id").MustInt64()
	var iTermRelationships model.TermRelationships
	self.Bind(&iTermRelationships)
	var iMap = helper.StructToMap(iTermRelationships)
	_error := model.UpdateTermRelationshipsByObjectId(ObjectId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermRelationshipsByTermTaxonomyIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermTaxonomyId_ := self.Args("term_taxonomy_id").MustInt64()
	var iTermRelationships model.TermRelationships
	self.Bind(&iTermRelationships)
	var iMap = helper.StructToMap(iTermRelationships)
	_error := model.UpdateTermRelationshipsByTermTaxonomyId(TermTaxonomyId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermRelationshipsByTermOrderHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermOrder_ := self.Args("term_order").MustInt()
	var iTermRelationships model.TermRelationships
	self.Bind(&iTermRelationships)
	var iMap = helper.StructToMap(iTermRelationships)
	_error := model.UpdateTermRelationshipsByTermOrder(TermOrder_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTermRelationshipsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	ObjectId_ := self.Args("object_id").MustInt64()
	_int64, _error := model.DeleteTermRelationships(ObjectId_)
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

func DeleteTermRelationshipsByObjectIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	ObjectId_ := self.Args("object_id").MustInt64()
	_error := model.DeleteTermRelationshipsByObjectId(ObjectId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTermRelationshipsByTermTaxonomyIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermTaxonomyId_ := self.Args("term_taxonomy_id").MustInt64()
	_error := model.DeleteTermRelationshipsByTermTaxonomyId(TermTaxonomyId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTermRelationshipsByTermOrderHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermOrder_ := self.Args("term_order").MustInt()
	_error := model.DeleteTermRelationshipsByTermOrder(TermOrder_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
