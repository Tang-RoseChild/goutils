package validutils

import (
	"reflect"
	"strings"
)

func FieldRequiredValid(src interface{}) bool {
	v := reflect.ValueOf(src)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return false
		}
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		if strings.Contains(string(v.Type().Field(i).Tag), "required") {
			if ZeroValue(v.Field(i)) {
				return false
			}
		}
	}

	return true
}

func ZeroValue(val reflect.Value) bool {
	switch val.Type().Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return val.Uint() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() == 0
	case reflect.String:
		return val.String() == ""
	case reflect.Chan, reflect.Ptr, reflect.Map, reflect.Slice:
		return val.IsNil()
	case reflect.Struct:
		for i := 0; i < reflect.ValueOf(val).NumField(); i++ {
			if !ZeroValue(val.Field(i)) {
				return false
			}
		}
		return true
	default:
		panic("unsupport type " + val.Type().Kind().String())
	}

	return false
}
