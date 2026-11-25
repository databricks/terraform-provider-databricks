package sdkgoutility

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/sdk-go/core/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// StructToData writes result, a modular SDK struct or a pointer to one, into
// the resource data d. It is the counterpart of common.StructToData for the
// structs of github.com/databricks/sdk-go, and it leaves the state that
// common.StructToData leaves for the classic twin of result.
//
// For each exported field whose FieldName has an entry in s, it calls d.Set
// with:
//
//   - the value of a scalar, or the one that a pointer to a scalar points to;
//   - the string of an enum, which is a named string type;
//   - the string of a types.Time or a types.Duration, or of a pointer to one:
//     the text of its JSON without the quotes, as 2026-01-02T03:04:05Z and
//     3.500s;
//   - a list with one map for a nested message, and with one map per message
//     for a slice of messages, where a nil message adds no map;
//   - a list for a slice of scalars, enums, times or durations;
//   - the map of a map[string]string.
//
// A nil pointer to a scalar is read as the zero value of that scalar: a nil
// *string as "", a nil *int64 as 0, a nil *float64 as 0 and a nil *bool as
// false. That is what the classic struct holds for a field that the response
// leaves out, so the rules below see the same value on both paths. A nil
// *types.Time or *types.Duration is read as "", like a nil *string, because
// the classic struct holds such a field as a string.
//
// The rules that decide whether a field is written are those of
// common.StructToData:
//
//   - An optional field is skipped when it is empty: an empty string, a nil
//     message, or a slice or map without entries. So a nil *string,
//     *types.Time or *types.Duration is skipped, while 0 and false aren't
//     empty, and a nil *int64, *float64 or *bool is written as 0 or false like
//     any other value.
//   - Unless d is a new resource, a field that the configuration doesn't set
//     is skipped, except when its schema is computed or a bool.
//   - A list of nested blocks is never written empty.
//   - In a nested block, only the first rule applies, and only to fields whose
//     schema is neither a list nor a set.
//   - At the end, every optional, non-computed field of s that has a default
//     and isn't set gets the value that d reads for it.
//
// StructToData returns an error that names what was wrong for a nil result,
// for a result that isn't a struct or a pointer to one, for a field of a type
// that it doesn't support, for a message whose schema isn't a nested block,
// for a time or a duration outside the range that its JSON form allows, and
// when d.Set fails.
//
// StructToData deliberately doesn't:
//
//   - use the difference between a nil pointer and a pointer to a zero value,
//     because the classic path can't see it and the state would change;
//   - keep the text of a time or a duration as the API sent it. The SDK
//     parses it, and StructToData writes the form of the SDK: a time in UTC
//     with a Z, and both with 0, 3, 6 or 9 fraction digits. The classic struct
//     holds the text of the API, so the state differs where the API sends
//     another form, as 2026-01-02T03:04:05.000Z or an offset;
//   - write a field mask. A types.FieldMask says which fields a request
//     updates and is no attribute of the resource. A field of that type, by
//     value or by pointer, is skipped even when s has an entry of its name;
//   - read json or tf tags or aliases, or flatten an embedded struct. Classic
//     SDK structs and the provider's own structs stay with
//     common.StructToData;
//   - support the interface of a oneof group, a body stream, raw JSON or a map
//     of anything but strings. Such a field is an error once s has its name,
//     whatever it holds;
//   - check that s fits result. A field without an entry in s is skipped, and
//     an entry without a field is left to its default.
func StructToData(result any, s map[string]*schema.Schema, d *schema.ResourceData) error {
	message, err := structToDataMessage(result)
	if err != nil {
		return err
	}
	err = structToDataEachField(message, s, func(name string, field reflect.Value, fieldSchema *schema.Schema) error {
		value, err := structToDataConvert(field, fieldSchema)
		if err != nil {
			return err
		}
		if fieldSchema.Optional && value.empty {
			return nil
		}
		// A bool is written even when it isn't configured, because GetOk can't
		// tell false from unset and drift would go unnoticed otherwise.
		if !d.IsNewResource() && !fieldSchema.Computed && fieldSchema.Type != schema.TypeBool {
			if _, configured := d.GetOk(name); !configured {
				return nil
			}
		}
		if value.blocks && value.empty {
			return nil
		}
		return d.Set(name, value.value)
	})
	if err != nil {
		return err
	}
	return structToDataWriteDefaults(s, d)
}

// structToDataMessage returns the struct that result is or points to.
func structToDataMessage(result any) (reflect.Value, error) {
	v := reflect.ValueOf(result)
	if !v.IsValid() {
		return v, fmt.Errorf("result is nil, want a struct or a pointer to one")
	}
	if v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return v, fmt.Errorf("result is a nil %T, want a struct or a pointer to one", result)
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return v, fmt.Errorf("result is a %T, want a struct or a pointer to one", result)
	}
	return v, nil
}

// structToDataEachField calls visit for every exported field of message whose
// FieldName has an entry in s, except a field mask, and names the field in the
// error that visit returns.
func structToDataEachField(message reflect.Value, s map[string]*schema.Schema,
	visit func(name string, field reflect.Value, fieldSchema *schema.Schema) error,
) error {
	messageType := message.Type()
	for i := range messageType.NumField() {
		field := messageType.Field(i)
		if !field.IsExported() || structToDataIsFieldMask(field.Type) {
			continue
		}
		name := FieldName(messageType, field)
		fieldSchema, ok := s[name]
		if !ok {
			continue
		}
		if err := visit(name, message.Field(i), fieldSchema); err != nil {
			return fmt.Errorf("%s: %w", name, err)
		}
	}
	return nil
}

// structToDataIsFieldMask reports whether t is a types.FieldMask of any
// message, or a pointer to one. The type is generic, so it is told by its
// package and by its name up to the type argument.
func structToDataIsFieldMask(t reflect.Type) bool {
	t = structToDataIndirect(t)
	return t.PkgPath() == reflect.TypeFor[types.Time]().PkgPath() && strings.HasPrefix(t.Name(), "FieldMask[")
}

// structToDataWriteDefaults sets every optional, non-computed field of s that
// has a default and no value to the value that d reads for it. Such a field
// has no counterpart in the struct, as force_destroy, and would stay null in
// the state after an import, while the plan resolves null to the default.
func structToDataWriteDefaults(s map[string]*schema.Schema, d *schema.ResourceData) error {
	for name, fieldSchema := range s {
		if !fieldSchema.Optional || fieldSchema.Computed || fieldSchema.Default == nil {
			continue
		}
		if _, ok := d.GetOk(name); ok {
			continue
		}
		if err := d.Set(name, d.Get(name)); err != nil {
			return fmt.Errorf("%s: %w", name, err)
		}
	}
	return nil
}

// structToDataValue is the value of a struct field in the form that d.Set
// takes.
type structToDataValue struct {
	value any
	// empty is true for an empty string, a nil message and a slice or map
	// without entries.
	empty bool
	// blocks is true for a list of nested blocks.
	blocks bool
}

// structToDataConvert returns the value of field in the form that d.Set takes
// for fieldSchema. It decides by the Go type of the field alone, so that a
// type that it doesn't support is an error whatever the field holds.
func structToDataConvert(field reflect.Value, fieldSchema *schema.Schema) (structToDataValue, error) {
	t := field.Type()
	switch {
	case structToDataIsScalar(structToDataIndirect(t)):
		value, err := structToDataScalar(field)
		return structToDataValue{value: value, empty: value == ""}, err
	case structToDataIsMessage(structToDataIndirect(t)):
		return structToDataBlocks([]reflect.Value{field}, fieldSchema)
	case t.Kind() == reflect.Slice && structToDataIsScalar(t.Elem()):
		list := make([]any, field.Len())
		var err error
		for i := 0; i < len(list) && err == nil; i++ {
			list[i], err = structToDataScalar(field.Index(i))
		}
		return structToDataValue{value: list, empty: len(list) == 0}, err
	case t.Kind() == reflect.Slice && structToDataIsMessage(structToDataIndirect(t.Elem())):
		messages := make([]reflect.Value, field.Len())
		for i := range messages {
			messages[i] = field.Index(i)
		}
		return structToDataBlocks(messages, fieldSchema)
	case t.Kind() == reflect.Map && t.Key().Kind() == reflect.String && t.Elem().Kind() == reflect.String:
		return structToDataValue{value: field.Interface(), empty: field.Len() == 0}, nil
	}
	return structToDataValue{}, fmt.Errorf("type %s is not supported", t)
}

// structToDataIndirect returns the type that t points to, and t itself when it
// is not a pointer.
func structToDataIndirect(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Pointer {
		return t.Elem()
	}
	return t
}

// structToDataIsScalar reports whether t is a string, a bool, a signed integer
// or a float, or a type named after one, as an enum, or a types.Time or a
// types.Duration, which are strings in JSON. Unsigned integers are left out,
// so that a slice of bytes, which holds raw JSON, isn't taken for a list.
func structToDataIsScalar(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.String, reflect.Bool, reflect.Float32, reflect.Float64,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	}
	return t == reflect.TypeFor[types.Time]() || t == reflect.TypeFor[types.Duration]()
}

// structToDataIsMessage reports whether t is a struct that is written as a
// nested block, which is every struct but a time and a duration.
func structToDataIsMessage(t reflect.Type) bool {
	return t.Kind() == reflect.Struct && !structToDataIsScalar(t)
}

// structToDataScalar returns the scalar that v is or points to as a string,
// bool, int64 or float64, whatever the name of its type. A nil pointer gives
// the zero value, which is what the classic struct holds for a field that the
// response leaves out. For a time and a duration that is the empty string,
// because the classic struct holds them as strings.
func structToDataScalar(v reflect.Value) (any, error) {
	if v.Kind() == reflect.Pointer {
		switch {
		case !v.IsNil():
			v = v.Elem()
		case v.Type().Elem().Kind() == reflect.Struct:
			return "", nil
		default:
			v = reflect.Zero(v.Type().Elem())
		}
	}
	switch v.Kind() {
	case reflect.String:
		return v.String(), nil
	case reflect.Bool:
		return v.Bool(), nil
	case reflect.Float32, reflect.Float64:
		return v.Float(), nil
	case reflect.Struct:
		return structToDataText(v.Interface().(json.Marshaler))
	}
	return v.Int(), nil
}

// structToDataText returns the text of the JSON string that value, a time or
// a duration, marshals to. It calls MarshalJSON itself, so that the error for
// a value outside the range of the JSON form is the one of the SDK.
func structToDataText(value json.Marshaler) (string, error) {
	data, err := value.MarshalJSON()
	if err != nil {
		return "", err
	}
	var text string
	err = json.Unmarshal(data, &text)
	return text, err
}

// structToDataBlocks returns messages, structs or pointers to them, as the
// list of maps that d.Set takes for the nested blocks of fieldSchema. A nil
// pointer adds no block.
func structToDataBlocks(messages []reflect.Value, fieldSchema *schema.Schema) (structToDataValue, error) {
	nested, ok := fieldSchema.Elem.(*schema.Resource)
	if !ok {
		return structToDataValue{}, fmt.Errorf("a message needs a nested block, but the schema is a %s without one", fieldSchema.Type)
	}
	blocks := []any{}
	for _, message := range messages {
		if message.Kind() == reflect.Pointer {
			if message.IsNil() {
				continue
			}
			message = message.Elem()
		}
		block := map[string]any{}
		err := structToDataEachField(message, nested.Schema, func(name string, field reflect.Value, fieldSchema *schema.Schema) error {
			value, err := structToDataConvert(field, fieldSchema)
			if err != nil {
				return err
			}
			isCollection := fieldSchema.Type == schema.TypeList || fieldSchema.Type == schema.TypeSet
			if !isCollection && fieldSchema.Optional && value.empty {
				return nil
			}
			block[name] = value.value
			return nil
		})
		if err != nil {
			return structToDataValue{}, err
		}
		blocks = append(blocks, block)
	}
	return structToDataValue{value: blocks, empty: len(blocks) == 0, blocks: true}, nil
}
