package flatten

import (
	"reflect"
)

func Flatten(in interface{}) []interface{} {
	var retur []interface{}
	if in == nil {
		return retur
	}
	_, ok := in.(int)
	if ok {
		retur = append(retur, in)
		return retur
	}
	switch reflect.TypeOf(in).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(in)
		if s.Len() >= 0 {
			for i := 0; i < s.Len(); i++ {
				if s.Index(i).IsNil() {
					break
				}
				back := Flatten(s.Index(i).Elem().Interface())
				for _, v := range back {
					retur = append(retur, v)
				}
			}
		}
	}
	return retur
}
