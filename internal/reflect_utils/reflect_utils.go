package reflect_utils

import "reflect"

type Field struct {
	Sf reflect.StructField
	V  reflect.Value
}

func ListAllFields(v reflect.Value) []Field {
	t := v.Type()
	fields := make([]Field, 0, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		f := t.Field(i)
		if f.Anonymous {
			fields = append(fields, ListAllFields(v.Field(i))...)
		} else {
			fields = append(fields, Field{
				Sf: f,
				V:  v.Field(i),
			})
		}
	}
	return fields
}
