package validutils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFieldRequiredValid(t *testing.T) {
	type D struct {
		Name   string            `yaml:"name, required"`
		Skills map[string]string `required`
	}

	var d1 D
	valid := FieldRequiredValid(d1)
	if valid {
		t.Error("should not valid ")
	}

	d1.Name = "123"
	valid = FieldRequiredValid(d1)
	if valid {
		t.Error("should not valid ")
	}

	d1.Skills = map[string]string{"test": "good"}
	valid = FieldRequiredValid(d1)
	if !valid {
		t.Error("should valid ")
	}
}

func TestZeroValue(t *testing.T) {
	ints := []interface{}{"abc", int(1), int8(2), int64(3), uint(4), uint8(5), "", 0}
	for idx := 0; idx < len(ints); idx++ {
		if idx >= len(ints)-2 {
			if !ZeroValue(reflect.ValueOf(ints[idx])) {
				t.Error(fmt.Sprintf("idx(%d) should be true ", idx))
			}
		} else {
			if ZeroValue(reflect.ValueOf(ints[idx])) {
				t.Error(fmt.Sprintf("idx(%d) should not be true ", idx))
			}
		}
	}
	var (
		v        int
		nilIntP  *int
		nilChan  chan int
		nilMap   map[string]string
		nilSlice []string
	)

	ptrs := []interface{}{make(chan int), make(map[string]string), make([]int, 10), &v, nilIntP, nilChan, nilMap, nilSlice}

	for idx := 0; idx < len(ptrs); idx++ {
		if idx >= len(ptrs)-4 {
			if !ZeroValue(reflect.ValueOf(ptrs[idx])) {
				t.Error(fmt.Sprintf("idx(%d) should be true ", idx))
			}
		} else {
			if ZeroValue(reflect.ValueOf(ptrs[idx])) {
				t.Error(fmt.Sprintf("idx(%d) should not be true ", idx))
			}
		}
	}

	type _D struct {
		Name  string
		Ptr   *int
		NodeP *_D
	}
	structs := []interface{}{_D{Name: "123"}, _D{NodeP: &_D{Ptr: &v}}, _D{}}
	for idx := 0; idx < len(structs); idx++ {
		if idx >= len(structs)-1 {
			if !ZeroValue(reflect.ValueOf(structs[idx])) {
				t.Error(fmt.Sprintf("idx(%d) should be true ", idx))
			}
		} else {
			if ZeroValue(reflect.ValueOf(structs[idx])) {
				t.Error(fmt.Sprintf("idx(%d) should not be true ", idx))
			}
		}
	}
}
