// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
package tfreflect

import (
	"reflect"
)

type Field struct {
	StructField reflect.StructField
	Value       reflect.Value
}

// ListAllFields takes in a reflect.Value of a struct, returns all of the fields for both struct field and
// the value. This function also extracts and flattens the anonymous fields nested inside.
func ListAllFields(v reflect.Value) []Field {
	t := v.Type()
	fields := make([]Field, 0, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		f := t.Field(i)
		if f.Anonymous {
			fields = append(fields, ListAllFields(v.Field(i))...)
		} else {
			fields = append(fields, Field{
				StructField: f,
				Value:       v.Field(i),
			})
		}
	}
	return fields
}
