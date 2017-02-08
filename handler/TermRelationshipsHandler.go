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

func GetTermRelationshipsCountViaObjectIdHandler(self *macross.Context) error {
	ObjectId_ := self.Args("object_id").MustInt64()
	_int64 := model.GetTermRelationshipsCountViaObjectId(ObjectId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["term_relationshipsCount"] = 0
	}
	m["term_relationshipsCount"] = _int64
	return self.JSON(m)
}

func GetTermRelationshipsCountViaTermTaxonomyIdHandler(self *macross.Context) error {
	TermTaxonomyId_ := self.Args("term_taxonomy_id").MustInt64()
	_int64 := model.GetTermRelationshipsCountViaTermTaxonomyId(TermTaxonomyId_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["term_relationshipsCount"] = 0
	}
	m["term_relationshipsCount"] = _int64
	return self.JSON(m)
}

func GetTermRelationshipsCountViaTermOrderHandler(self *macross.Context) error {
	TermOrder_ := self.Args("term_order").MustInt()
	_int64 := model.GetTermRelationshipsCountViaTermOrder(TermOrder_)
	var m = map[string]interface{}{}
	if _int64 <= 0 {
		m["term_relationshipsCount"] = 0
	}
	m["term_relationshipsCount"] = _int64
	return self.JSON(m)
}

func GetTermRelationshipsesViaObjectIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iObjectId := self.Args("object_id").MustInt64()
	if (offset > 0) && helper.IsHas(iObjectId) {
		_TermRelationships, _error := model.GetTermRelationshipsesViaObjectId(offset, limit, iObjectId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermRelationships)
	}
	herr.Message = "Can't get to the GetTermRelationshipsesViaObjectId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermRelationshipsesViaTermTaxonomyIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTermTaxonomyId := self.Args("term_taxonomy_id").MustInt64()
	if (offset > 0) && helper.IsHas(iTermTaxonomyId) {
		_TermRelationships, _error := model.GetTermRelationshipsesViaTermTaxonomyId(offset, limit, iTermTaxonomyId, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermRelationships)
	}
	herr.Message = "Can't get to the GetTermRelationshipsesViaTermTaxonomyId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermRelationshipsesViaTermOrderHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	offset := self.Args("offset").MustInt()
	limit := self.Args("limit").MustInt()
	field := self.Args("field").String()
	iTermOrder := self.Args("term_order").MustInt()
	if (offset > 0) && helper.IsHas(iTermOrder) {
		_TermRelationships, _error := model.GetTermRelationshipsesViaTermOrder(offset, limit, iTermOrder, field)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermRelationships)
	}
	herr.Message = "Can't get to the GetTermRelationshipsesViaTermOrder's args."
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

func GetHasTermRelationshipsViaObjectIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iObjectId := self.Args("object_id").MustInt64()
	if helper.IsHas(iObjectId) {
		_TermRelationships := model.HasTermRelationshipsViaObjectId(iObjectId)
		var m = map[string]interface{}{}
		m["term_relationships"] = _TermRelationships
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermRelationshipsViaObjectId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTermRelationshipsViaTermTaxonomyIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTermTaxonomyId := self.Args("term_taxonomy_id").MustInt64()
	if helper.IsHas(iTermTaxonomyId) {
		_TermRelationships := model.HasTermRelationshipsViaTermTaxonomyId(iTermTaxonomyId)
		var m = map[string]interface{}{}
		m["term_relationships"] = _TermRelationships
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermRelationshipsViaTermTaxonomyId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetHasTermRelationshipsViaTermOrderHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTermOrder := self.Args("term_order").MustInt()
	if helper.IsHas(iTermOrder) {
		_TermRelationships := model.HasTermRelationshipsViaTermOrder(iTermOrder)
		var m = map[string]interface{}{}
		m["term_relationships"] = _TermRelationships
		return self.JSON(m)
	}
	herr.Message = "Can't get to the HasTermRelationshipsViaTermOrder's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermRelationshipsViaObjectIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iObjectId := self.Args("object_id").MustInt64()
	if helper.IsHas(iObjectId) {
		_TermRelationships, _error := model.GetTermRelationshipsViaObjectId(iObjectId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermRelationships)
	}
	herr.Message = "Can't get to the GetTermRelationshipsViaObjectId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermRelationshipsViaTermTaxonomyIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTermTaxonomyId := self.Args("term_taxonomy_id").MustInt64()
	if helper.IsHas(iTermTaxonomyId) {
		_TermRelationships, _error := model.GetTermRelationshipsViaTermTaxonomyId(iTermTaxonomyId)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermRelationships)
	}
	herr.Message = "Can't get to the GetTermRelationshipsViaTermTaxonomyId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func GetTermRelationshipsViaTermOrderHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	iTermOrder := self.Args("term_order").MustInt()
	if helper.IsHas(iTermOrder) {
		_TermRelationships, _error := model.GetTermRelationshipsViaTermOrder(iTermOrder)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermRelationships)
	}
	herr.Message = "Can't get to the GetTermRelationshipsViaTermOrder's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermRelationshipsViaObjectIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	ObjectId_ := self.Args("object_id").MustInt64()
	if helper.IsHas(ObjectId_) {
		var iTermRelationships model.TermRelationships
		self.Bind(&iTermRelationships)
		_TermRelationships, _error := model.SetTermRelationshipsViaObjectId(ObjectId_, &iTermRelationships)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermRelationships)
	}
	herr.Message = "Can't get to the SetTermRelationshipsViaObjectId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermRelationshipsViaTermTaxonomyIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermTaxonomyId_ := self.Args("term_taxonomy_id").MustInt64()
	if helper.IsHas(TermTaxonomyId_) {
		var iTermRelationships model.TermRelationships
		self.Bind(&iTermRelationships)
		_TermRelationships, _error := model.SetTermRelationshipsViaTermTaxonomyId(TermTaxonomyId_, &iTermRelationships)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermRelationships)
	}
	herr.Message = "Can't get to the SetTermRelationshipsViaTermTaxonomyId's args."
	return self.JSON(herr, macross.StatusServiceUnavailable)
}

func PostSetTermRelationshipsViaTermOrderHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermOrder_ := self.Args("term_order").MustInt()
	if helper.IsHas(TermOrder_) {
		var iTermRelationships model.TermRelationships
		self.Bind(&iTermRelationships)
		_TermRelationships, _error := model.SetTermRelationshipsViaTermOrder(TermOrder_, &iTermRelationships)
		if _error != nil {
			herr.Message = _error.Error()
			return self.JSON(herr, macross.StatusServiceUnavailable)
		}
		return self.JSON(_TermRelationships)
	}
	herr.Message = "Can't get to the SetTermRelationshipsViaTermOrder's args."
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

func PutTermRelationshipsHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	var iTermRelationships model.TermRelationships
	self.Bind(&iTermRelationships)
	_int64, _error := model.PutTermRelationships(&iTermRelationships)
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

func PutTermRelationshipsViaObjectIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	ObjectId_ := self.Args("object_id").MustInt64()
	var iTermRelationships model.TermRelationships
	self.Bind(&iTermRelationships)
	_int64, _error := model.PutTermRelationshipsViaObjectId(ObjectId_, &iTermRelationships)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTermRelationshipsViaTermTaxonomyIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermTaxonomyId_ := self.Args("term_taxonomy_id").MustInt64()
	var iTermRelationships model.TermRelationships
	self.Bind(&iTermRelationships)
	_int64, _error := model.PutTermRelationshipsViaTermTaxonomyId(TermTaxonomyId_, &iTermRelationships)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutTermRelationshipsViaTermOrderHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermOrder_ := self.Args("term_order").MustInt()
	var iTermRelationships model.TermRelationships
	self.Bind(&iTermRelationships)
	_int64, _error := model.PutTermRelationshipsViaTermOrder(TermOrder_, &iTermRelationships)
	if (_int64 <= 0) || (_error != nil) {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermRelationshipsViaObjectIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	ObjectId_ := self.Args("object_id").MustInt64()
	var iTermRelationships model.TermRelationships
	self.Bind(&iTermRelationships)
	var iMap = helper.StructToMap(iTermRelationships)
	_error := model.UpdateTermRelationshipsViaObjectId(ObjectId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermRelationshipsViaTermTaxonomyIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermTaxonomyId_ := self.Args("term_taxonomy_id").MustInt64()
	var iTermRelationships model.TermRelationships
	self.Bind(&iTermRelationships)
	var iMap = helper.StructToMap(iTermRelationships)
	_error := model.UpdateTermRelationshipsViaTermTaxonomyId(TermTaxonomyId_, &iMap)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func PutUpdateTermRelationshipsViaTermOrderHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermOrder_ := self.Args("term_order").MustInt()
	var iTermRelationships model.TermRelationships
	self.Bind(&iTermRelationships)
	var iMap = helper.StructToMap(iTermRelationships)
	_error := model.UpdateTermRelationshipsViaTermOrder(TermOrder_, &iMap)
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

func DeleteTermRelationshipsViaObjectIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	ObjectId_ := self.Args("object_id").MustInt64()
	_error := model.DeleteTermRelationshipsViaObjectId(ObjectId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTermRelationshipsViaTermTaxonomyIdHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermTaxonomyId_ := self.Args("term_taxonomy_id").MustInt64()
	_error := model.DeleteTermRelationshipsViaTermTaxonomyId(TermTaxonomyId_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}

func DeleteTermRelationshipsViaTermOrderHandler(self *macross.Context) error {
	var herr = new(macross.HTTPError)
	herr.Message = "ErrServiceUnavailable"
	herr.Status = macross.StatusServiceUnavailable
	TermOrder_ := self.Args("term_order").MustInt()
	_error := model.DeleteTermRelationshipsViaTermOrder(TermOrder_)
	if _error != nil {
		herr.Message = _error.Error()
		return self.JSON(herr, macross.StatusServiceUnavailable)
	}
	herr.Message = "StatusOK"
	herr.Status = macross.StatusOK
	return self.JSON(herr)
}
