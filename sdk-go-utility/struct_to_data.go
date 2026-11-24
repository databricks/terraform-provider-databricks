package sdkgoutility

import (
	"fmt"
	"reflect"

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
//   - a list with one map for a nested message, and with one map per message
//     for a slice of messages, where a nil message adds no map;
//   - a list for a slice of scalars or enums;
//   - the map of a map[string]string.
//
// A nil pointer to a scalar is read as the zero value of that scalar. That is
// what the classic struct holds for a field that the response leaves out, so
// the rules below see the same value on both paths.
//
// The rules that decide whether a field is written are those of
// common.StructToData:
//
//   - An optional field is skipped when it is empty: an empty string, a nil
//     message, or a slice or map without entries. A zero number and false
//     aren't empty.
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
// and when d.Set fails.
//
// StructToData deliberately doesn't:
//
//   - use the difference between a nil pointer and a pointer to a zero value,
//     because the classic path can't see it and the state would change;
//   - read json or tf tags, aliases or embedded structs. Classic SDK structs
//     and the provider's own structs stay with common.StructToData;
//   - support the interface of a oneof group, a body stream, raw JSON or a map
//     of anything but strings. Such a field is an error once s has its name;
//   - check that s fits result. A field without an entry in s is skipped, and
//     an entry without a field is left to its default.
func StructToData(result any, s map[string]*schema.Schema, d *schema.ResourceData) error {
	message, err := messageOf(result)
	if err != nil {
		return err
	}
	err = eachField(message, s, func(name string, field reflect.Value, fieldSchema *schema.Schema) error {
		value, err := convert(field, fieldSchema)
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
	return writeDefaults(s, d)
}

// messageOf returns the struct that result is or points to.
func messageOf(result any) (reflect.Value, error) {
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

// eachField calls visit for every exported field of message whose FieldName
// has an entry in s, and names the field in the error that visit returns.
func eachField(message reflect.Value, s map[string]*schema.Schema,
	visit func(name string, field reflect.Value, fieldSchema *schema.Schema) error,
) error {
	messageType := message.Type()
	for i := range messageType.NumField() {
		if !messageType.Field(i).IsExported() {
			continue
		}
		name := FieldName(messageType, messageType.Field(i))
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

// writeDefaults sets every optional, non-computed field of s that has a
// default and no value to the value that d reads for it. Such a field has no
// counterpart in the struct, as force_destroy, and would stay null in the
// state after an import, while the plan resolves null to the default.
func writeDefaults(s map[string]*schema.Schema, d *schema.ResourceData) error {
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

// converted is the value of a struct field in the form that d.Set takes.
type converted struct {
	value any
	// empty is true for an empty string, a nil message and a slice or map
	// without entries.
	empty bool
	// blocks is true for a list of nested blocks.
	blocks bool
}

// convert returns the value of field in the form that d.Set takes for
// fieldSchema. It decides by the Go type of the field alone, so that a type
// that it doesn't support is an error whatever the field holds.
func convert(field reflect.Value, fieldSchema *schema.Schema) (converted, error) {
	t := field.Type()
	switch {
	case isScalar(indirect(t)):
		value := scalarOf(field)
		return converted{value: value, empty: value == ""}, nil
	case indirect(t).Kind() == reflect.Struct:
		return blocksOf([]reflect.Value{field}, fieldSchema)
	case t.Kind() == reflect.Slice && isScalar(t.Elem()):
		list := make([]any, field.Len())
		for i := range list {
			list[i] = scalarOf(field.Index(i))
		}
		return converted{value: list, empty: len(list) == 0}, nil
	case t.Kind() == reflect.Slice && indirect(t.Elem()).Kind() == reflect.Struct:
		messages := make([]reflect.Value, field.Len())
		for i := range messages {
			messages[i] = field.Index(i)
		}
		return blocksOf(messages, fieldSchema)
	case t.Kind() == reflect.Map && t.Key().Kind() == reflect.String && t.Elem().Kind() == reflect.String:
		return converted{value: field.Interface(), empty: field.Len() == 0}, nil
	}
	return converted{}, fmt.Errorf("type %s is not supported", t)
}

// indirect returns the type that t points to, and t itself when it is not a
// pointer.
func indirect(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Pointer {
		return t.Elem()
	}
	return t
}

// isScalar reports whether t is a string, a bool, a signed integer or a float,
// or a type named after one, as an enum. Unsigned integers are left out, so
// that a slice of bytes, which holds raw JSON, isn't taken for a list.
func isScalar(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.String, reflect.Bool, reflect.Float32, reflect.Float64,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	}
	return false
}

// scalarOf returns the scalar that v is or points to as a string, bool, int64
// or float64, whatever the name of its type. A nil pointer gives the zero
// value.
func scalarOf(v reflect.Value) any {
	if v.Kind() == reflect.Pointer {
		if v.IsNil() {
			v = reflect.Zero(v.Type().Elem())
		} else {
			v = v.Elem()
		}
	}
	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Bool:
		return v.Bool()
	case reflect.Float32, reflect.Float64:
		return v.Float()
	}
	return v.Int()
}

// blocksOf returns messages, structs or pointers to them, as the list of maps
// that d.Set takes for the nested blocks of fieldSchema. A nil pointer adds no
// block.
func blocksOf(messages []reflect.Value, fieldSchema *schema.Schema) (converted, error) {
	nested, ok := fieldSchema.Elem.(*schema.Resource)
	if !ok {
		return converted{}, fmt.Errorf("the field is a message, but its schema is not a nested block")
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
		err := eachField(message, nested.Schema, func(name string, field reflect.Value, fieldSchema *schema.Schema) error {
			value, err := convert(field, fieldSchema)
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
			return converted{}, err
		}
		blocks = append(blocks, block)
	}
	return converted{value: blocks, empty: len(blocks) == 0, blocks: true}, nil
}
