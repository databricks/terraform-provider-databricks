package marshal

import (
	"encoding/json"
	"errors"
	"reflect"
)

// Unmarshals a JSON element and fills in the ForceSendFields field if
// the struct contains it. Only anotates basic types in the ForceSendFields.
func Unmarshal(data []byte, v any) error {
	if len(data) == 0 {
		return nil
	}
	data, err := normalizeInt64Strings(data, v)
	if err != nil {
		return err
	}

	var jsonFields map[string]json.RawMessage
	if err = json.Unmarshal([]byte(data), &jsonFields); err != nil {
		return err
	}

	value := reflect.ValueOf(v)
	value = reflect.Indirect(value)

	objectType := value.Type()
	foundFields := []string{}

	for _, field := range getTypeFields(objectType) {
		if field.JsonTag.ignore {
			continue
		}
		index := field.IndexInStruct
		fieldValue := value.Field(index)
		fieldType := fieldValue.Type()
		tag := field.JsonTag
		if field.Anonymous {
			setField(fieldValue, data)
			continue
		}

		if err != nil {
			return err
		}

		value, ok := jsonFields[tag.name]
		if !ok {
			continue
		}

		if isBasicType(fieldType) {
			foundFields = append(foundFields, field.Name)
		}

		if !fieldValue.CanAddr() {
			return errors.New("cannot address field")
		}

		err = setField(fieldValue, value)
		if err != nil {
			return err
		}
	}

	return setForceSendFields(v, foundFields)

}

func setField(field reflect.Value, value []byte) error {
	pointer := field.Addr().Interface()
	err := json.Unmarshal(value, pointer)
	if err == nil {
		return nil
	}
	// We use godss/yaml to convert YAML into JSON.
	// This library stops converting when it finds a custom Marshaller,
	// Since strings in YAML may not have quotes, they won't be added.
	// So we add them manually
	quoted, err := json.Marshal(string(value))
	if err != nil {
		return err
	}
	return json.Unmarshal(quoted, pointer)
}

func setForceSendFields(v any, presentFields []string) error {
	if len(presentFields) == 0 {
		return nil
	}

	value := reflect.ValueOf(v)
	if value.Kind() != reflect.Ptr {
		return nil
	}

	// reflex.GetFieldByName may return the field from an anonymous field
	// if the extending struct doesn't have the field.
	field := getFieldByName(v, forceSendFieldName)
	if !field.IsValid() {
		return nil
	}

	if !field.CanSet() || field.Kind() != reflect.Slice {
		return errors.New("cannot set field")
	}

	presentFieldsValue := reflect.ValueOf(presentFields)
	field.Set(presentFieldsValue)

	return nil
}
