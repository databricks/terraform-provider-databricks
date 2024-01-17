package common

import (
	"fmt"
	"reflect"
	"slices"
	"strings"
)

// SetForceSendFields adds any fields specified in the `fields` parameter to the ForceSendFields field of the
// request structure if they are present in the resource state. The provided fields must match the JSON tag
// for some field in the request structure. This ensures that fields explicitly set to the zero value of its
// type (e.g. `0` for an `int`) will be serialized and sent to the platform.
//
// This function requires that the request structure has a `ForceSendFields` field of type `[]string`. If not,
// it panics with an appropriate error message.
func SetForceSendFields(req any, d attributeGetter, fields []string) {
	rv := reflect.ValueOf(req)
	if rv.Kind() != reflect.Ptr {
		panic("request argument to setForceSendFields must be a pointer")
	}
	rv = rv.Elem()
	forceSendFieldsField := rv.FieldByName("ForceSendFields")
	if !forceSendFieldsField.IsValid() {
		panic("request argument to setForceSendFields must have ForceSendFields field")
	}
	forceSendFields, ok := forceSendFieldsField.Interface().([]string)
	if !ok {
		panic(fmt.Errorf("request argument to setForceSendFields must have ForceSendFields field of type []string (got %s)", forceSendFieldsField.Type()))
	}
	fs := listAllFields(rv)
	for _, fieldName := range fields {
		found := false
		var structField reflect.StructField
		for _, f := range fs {
			fn := chooseFieldName(f.sf)
			if fn != "-" && fn == fieldName {
				found = true
				structField = f.sf
				break
			}
		}
		if !found {
			allFieldNames := make([]string, 0)
			for _, f := range fs {
				fn := chooseFieldName(f.sf)
				if fn == "-" || fn == "force_send_fields" {
					continue
				}
				allFieldNames = append(allFieldNames, fn)
			}
			panic(fmt.Errorf("unexpected field %s not found in request structure, expected one of: %s", fieldName, strings.Join(allFieldNames, ", ")))
		}
		// Check if the field was ever set, even to the zero value of the type.
		if _, ok := d.GetOkExists(fieldName); !ok {
			continue
		}
		if !slices.Contains(forceSendFields, structField.Name) {
			forceSendFields = append(forceSendFields, structField.Name)
		}
	}
	forceSendFieldsField.Set(reflect.ValueOf(forceSendFields))
}
