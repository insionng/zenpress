package helper

import (
	"reflect"
)

func IsHas(i interface{}) bool {
	if i != nil {
		return true
	}
	return false
}

func StructToMap(i interface{}) (m map[string]interface{}) {
	m = make(map[string]interface{})
	vt := reflect.TypeOf(i)
	vv := reflect.ValueOf(i)
	for i := 0; i < vt.NumField(); i++ {
		f := vt.Field(i)
		m[f.Name] = vv.FieldByName(f.Name).String()
	}
	return
}
