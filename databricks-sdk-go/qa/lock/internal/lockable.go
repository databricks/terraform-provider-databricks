package internal

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

func GenerateBlobName[R any](r R) string {
	rv := reflect.ValueOf(r)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	switch rv.Kind() {
	case reflect.String:
		return rv.String()
	case reflect.Struct:
		id := bytes.Buffer{}
		t := rv.Type()
		if rv.NumField() == 0 {
			return t.Name()
		}
		id.WriteString(fmt.Sprintf("%s:", t.Name()))
		for i := 0; i < rv.NumField(); i++ {
			fieldName := t.Field(i).Name
			id.WriteString(fmt.Sprintf("%s=", fieldName))
			switch v := rv.Field(i).Interface().(type) {
			case string:
				noSlashes := strings.ReplaceAll(v, "/", "_")
				id.WriteString(noSlashes)
			case int, int8, int16, int32, int64:
				id.WriteString(fmt.Sprintf("%d", v))
			default:
				panic(fmt.Errorf("unsupported type %T for field %s, must be string or int", v, fieldName))
			}
			if i < rv.NumField()-1 {
				id.WriteString(";")
			}
		}
		return id.String()
	default:
		panic(fmt.Errorf("unsupported type %T, must be string or struct (or pointer to string or pointer to struct)", r))
	}
}
