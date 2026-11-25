package marshal

import (
	"encoding/json"
	"fmt"
	"reflect"

	"golang.org/x/exp/maps"
)

const forceSendFieldName = "ForceSendFields"

// Marshal returns a JSON encoding of the given object. Included fields:
// - non-empty value
// - a basic type whose field's name is present in forceSendFields
// Nil interfaces, nil pointers, and nil or empty maps and slices are not serialized even if their field name appears in ForceSendFields
// Embedded structs are still considered a separate struct. ForceSendFields
// in an embedded struct only impact the fields of the embedded struct.
// Conversely, an embedded struct is not impacted by the ForceSendFields
// of the struct containing it.
func Marshal(object any) ([]byte, error) {
	dataMap, err := structAsMap(object)
	if err != nil {
		return nil, err
	}

	return json.Marshal(dataMap)
}

// Converts the object to a map
func structAsMap(object any) (map[string]any, error) {
	result := make(map[string]any)

	// If the object is nil or a pointer to nil, don't do anything
	if object == nil || (reflect.ValueOf(object).Kind() == reflect.Ptr && reflect.ValueOf(object).IsNil()) {
		return result, nil
	}

	value := reflect.ValueOf(object)
	value = reflect.Indirect(value)
	objectType := value.Type()

	includeFields, err := getForceSendFields(object, getTypeName(objectType))

	if err != nil {
		return nil, err
	}

	for _, field := range getTypeFields(objectType) {
		tag := field.JsonTag
		if tag.ignore {
			continue
		}

		fieldValue := value.Field(field.IndexInStruct)
		// Anonymous fields should be marshalled using the same JSON, and then merged into the same map
		if field.Anonymous && fieldValue.IsValid() {
			anonymousFieldResult, err := structAsMap(fieldValue.Interface())
			if err != nil {
				return nil, err
			}
			result = mergeMaps(anonymousFieldResult, result)
			continue
		}

		// Skip fields which should not be included
		if !includeField(tag, fieldValue, includeFields[field.Name]) {
			continue
		}

		if tag.asString {
			result[tag.name] = formatAsString(fieldValue, field.Type.Kind())
		} else {
			result[tag.name] = fieldValue.Interface()
		}
	}
	return result, nil
}

// returns the element as string
func formatAsString(v reflect.Value, kind reflect.Kind) string {
	if kind == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}

	return fmt.Sprintf("%v", v.Interface())
}

type jsonTag struct {
	name      string
	asString  bool
	ignore    bool
	omitempty bool
}

// Determines whether a field should be included or not
func includeField(tag jsonTag, value reflect.Value, forceSend bool) bool {
	if tag.ignore {
		return false
	}
	if !tag.omitempty {
		return true
	}
	return (isBasicType(value.Type()) && forceSend) || !isEmptyValue(value)
}

// isEmptyValue returns whether v is the empty value for its type.
// This implementation is based on on the encoding/json package for consistency on the results
// https://github.com/golang/go/blob/a278550c40ef3f01a5fcbef43414dc49009201f8/src/encoding/json/encode.go#L306
func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Pointer:
		return v.IsNil()
	}
	return false
}

// Merges two maps. Overrides duplicated elements, which is fine because json.Unmarshal also
// does it, and a JSON should not have duplicated entries.
func mergeMaps(m1 map[string]any, m2 map[string]any) map[string]any {
	merged := make(map[string]any)
	maps.Copy(merged, m1)
	maps.Copy(merged, m2)
	return merged
}

func getForceSendFields(v any, structName string) (map[string]bool, error) {
	// reflect.GetFieldByName panics if the field is inside a null anonymous field
	field := getFieldByName(v, forceSendFieldName)
	if !field.IsValid() {
		return nil, nil
	}
	forceSendFields, ok := field.Interface().([]string)
	if !ok {
		return nil, fmt.Errorf("invalid type for %s field", forceSendFieldName)
	}
	existingFields := getFieldNames(v)
	includeFields := make(map[string]bool)
	for _, field := range forceSendFields {
		if _, ok := existingFields[field]; ok {
			includeFields[field] = true
		} else {
			return nil, fmt.Errorf("field %s cannot be found in struct %s", field, structName)
		}
	}

	return includeFields, nil

}

func getFieldByName(v any, fieldName string) reflect.Value {
	value := reflect.ValueOf(v)
	value = reflect.Indirect(value)
	objectType := value.Type()

	for _, field := range getTypeFields(objectType) {
		name := field.Name
		if name == fieldName {
			return value.Field(field.IndexInStruct)
		}
	}
	return reflect.Value{}
}

func getFieldNames(v any) map[string]bool {
	result := map[string]bool{}
	value := reflect.ValueOf(v)
	value = reflect.Indirect(value)
	objectType := value.Type()

	for _, field := range getTypeFields(objectType) {
		name := field.Name
		result[name] = true
	}
	return result
}
