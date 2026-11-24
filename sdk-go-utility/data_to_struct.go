package sdkgoutility

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataToStruct fills result, a pointer to a modular SDK struct, from the
// resource data d, whose schema is s. It does for the structs of
// github.com/databricks/sdk-go what common.DataToStructPointer does for the
// structs of the classic SDK.
//
// For each exported field of the struct whose FieldName has an entry in s,
// DataToStruct reads the value with d.GetOk. A field without a schema entry is
// skipped, and so is a schema entry without a field. When GetOk reports the
// value unset or zero, the field keeps the value it came with, which is nil or
// empty in a new struct, so the SDK leaves it out of the request, as the
// classic client leaves out an empty omitempty field. Otherwise the field is
// replaced with:
//
//   - a string, a bool, an integer or a float, or a pointer to one, from a
//     value of that schema type. A named string type, which is how the SDK
//     declares an enum, is a string.
//   - a message, by pointer or by value, from a list of one block.
//   - a slice of messages from a list of blocks, or a slice of strings, bools,
//     integers or floats from a list of them.
//   - a map of strings to strings from a map of strings.
//
// The fields of a message are read the same way at every depth, each under
// the key that the resource data has for it, such as tags.0.key.
//
// DataToStruct returns an error instead of panicking. The error names the
// field by its key in the resource data, its Go name and its type, and never
// quotes a value of the resource data. DataToStruct returns one when:
//
//   - result is not a non-nil pointer to a struct.
//   - a field with a schema entry has another type than those above: an
//     interface, json.RawMessage, []byte, an unsigned integer, a slice of
//     slices or of pointers, a map with a non-string key or value, or a
//     pointer to anything but a scalar or a struct. The type is checked before
//     the value is read, so the error doesn't wait for a configuration that
//     sets the field.
//   - the value doesn't fit the field: its schema type is not the one that the
//     field takes, a number overflows the field, or a list has more than one
//     block for a single message. An empty string in a list of strings doesn't
//     fit either: the Terraform SDK passes it as nil, on which
//     common.DataToStructPointer panics.
//
// DataToStruct deliberately doesn't:
//
//   - send a zero value. A caller that must send an explicit false, 0 or empty
//     string sets that pointer itself after the call, where it adds a
//     ForceSendFields entry today. That includes a field that the classic
//     struct declares without omitempty: the classic client sends such a field
//     even when it is zero, and DataToStruct leaves it nil.
//   - read a set. A field whose schema type is schema.TypeSet gets the error
//     of a value that doesn't fit.
//   - convert a string to a timestamp, a duration or a field mask. Those are
//     structs in the SDK, so a string doesn't fit them, and the caller sets
//     them itself.
//   - look into embedded structs, or read json tags, tf tags or aliases. The
//     structs that need those stay with the helpers in package common.
//   - validate the resource data. It relies on the validation that the
//     Terraform SDK did against the schema.
func DataToStruct(d *schema.ResourceData, s map[string]*schema.Schema, result any) error {
	message := reflect.ValueOf(result)
	if message.Kind() != reflect.Pointer || message.Type().Elem().Kind() != reflect.Struct {
		return fmt.Errorf("DataToStruct: result is %T, not a pointer to a struct", result)
	}
	if message.IsNil() {
		return fmt.Errorf("DataToStruct: result is a nil %T", result)
	}
	return dataToStructMessage(d, s, "", message.Elem())
}

// dataToStructField is a struct field that dataToStructMessage fills from the
// resource data.
type dataToStructField struct {
	d      *schema.ResourceData
	schema *schema.Schema
	// path is the key of the field in d, such as tags.0.key.
	path string
	// owner is the struct that declares field.
	owner reflect.Type
	field reflect.StructField
}

// errorf returns an error that names the field by its key in the resource
// data, its Go name and its type, followed by the formatted problem.
func (f dataToStructField) errorf(format string, args ...any) error {
	problem := fmt.Sprintf(format, args...)
	return fmt.Errorf("%s: %s.%s of type %s %s", f.path, f.owner, f.field.Name, f.field.Type, problem)
}

// dataToStructMessage fills message, a struct, from the values that d holds
// under prefix: the empty string for the resource itself, and the key of a
// block with a trailing dot, such as "tags.0.", for a nested message. s is the
// schema at that level.
func dataToStructMessage(d *schema.ResourceData, s map[string]*schema.Schema, prefix string, message reflect.Value) error {
	owner := message.Type()
	for i := range owner.NumField() {
		field := owner.Field(i)
		name := FieldName(owner, field)
		// A nil entry is no entry: a missing key gives nil as well.
		fieldSchema := s[name]
		if fieldSchema == nil || !field.IsExported() {
			continue
		}
		f := dataToStructField{d: d, schema: fieldSchema, path: prefix + name, owner: owner, field: field}
		if !dataToStructSupports(field.Type) {
			return f.errorf("is not a type that DataToStruct supports")
		}
		raw, ok := d.GetOk(f.path)
		if !ok {
			continue
		}
		// Filling a new value and setting the field afterwards replaces the
		// field as a whole, and leaves it untouched when the value doesn't fit.
		value := reflect.New(field.Type).Elem()
		if err := f.set(value, raw); err != nil {
			return err
		}
		message.Field(i).Set(value)
	}
	return nil
}

// dataToStructSupports reports whether DataToStruct fills a field of type t.
func dataToStructSupports(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Pointer, reflect.Slice:
		return t.Elem().Kind() == reflect.Struct || dataToStructIsScalar(t.Elem())
	case reflect.Map:
		return t.Key().Kind() == reflect.String && t.Elem().Kind() == reflect.String
	case reflect.Struct:
		return true
	default:
		return dataToStructIsScalar(t)
	}
}

// dataToStructIsScalar reports whether t has one of the kinds that a
// schema.TypeString, TypeBool, TypeInt or TypeFloat value can be stored in.
// Unsigned integers are not among them: the modular SDK has no such field, and
// []byte is a slice of them.
func dataToStructIsScalar(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.String, reflect.Bool, reflect.Float32, reflect.Float64,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	default:
		return false
	}
}

// set fills target, a zero value of a supported type, from raw, the value that
// GetOk returned for the field.
func (f dataToStructField) set(target reflect.Value, raw any) error {
	switch target.Kind() {
	case reflect.Pointer:
		pointee := reflect.New(target.Type().Elem())
		target.Set(pointee)
		return f.set(pointee.Elem(), raw)
	case reflect.Struct:
		blocks, blockSchema, err := f.blocks(raw)
		if err != nil {
			return err
		}
		if len(blocks) != 1 {
			return f.errorf("takes one block, but the resource data has %d", len(blocks))
		}
		return dataToStructMessage(f.d, blockSchema, f.path+".0.", target)
	case reflect.Slice:
		return f.setSlice(target, raw)
	case reflect.Map:
		return f.setMap(target, raw)
	default:
		if !dataToStructSetScalar(target, raw) {
			return f.errorf("can't hold the value in the resource data, which is of type %T", raw)
		}
		return nil
	}
}

// blocks returns raw as the list of blocks that a message or a slice of
// messages is read from, with the schema of one block. The items of the list
// are not used: the fields of a block are read from the resource data by key.
func (f dataToStructField) blocks(raw any) ([]any, map[string]*schema.Schema, error) {
	blocks, ok := raw.([]any)
	if !ok {
		return nil, nil, f.errorf("takes a list of blocks, but the value in the resource data is of type %T", raw)
	}
	block, ok := f.schema.Elem.(*schema.Resource)
	if !ok {
		return nil, nil, f.errorf("takes a list of blocks, but the schema of the field has no block")
	}
	return blocks, block.Schema, nil
}

// setSlice fills target, an empty slice of messages or of scalars, from raw.
func (f dataToStructField) setSlice(target reflect.Value, raw any) error {
	if target.Type().Elem().Kind() == reflect.Struct {
		blocks, blockSchema, err := f.blocks(raw)
		if err != nil {
			return err
		}
		target.Set(reflect.MakeSlice(target.Type(), len(blocks), len(blocks)))
		for i := range blocks {
			prefix := f.path + "." + strconv.Itoa(i) + "."
			if err := dataToStructMessage(f.d, blockSchema, prefix, target.Index(i)); err != nil {
				return err
			}
		}
		return nil
	}
	items, ok := raw.([]any)
	if !ok {
		return f.errorf("takes a list, but the value in the resource data is of type %T", raw)
	}
	target.Set(reflect.MakeSlice(target.Type(), len(items), len(items)))
	for i, item := range items {
		if !dataToStructSetScalar(target.Index(i), item) {
			return f.errorf("can't hold item %d of the list, which is of type %T", i, item)
		}
	}
	return nil
}

// setMap fills target, a nil map of strings to strings, from raw.
func (f dataToStructField) setMap(target reflect.Value, raw any) error {
	entries, ok := raw.(map[string]any)
	if !ok {
		return f.errorf("takes a map, but the value in the resource data is of type %T", raw)
	}
	mapType := target.Type()
	target.Set(reflect.MakeMapWithSize(mapType, len(entries)))
	for key, entry := range entries {
		value := reflect.New(mapType.Elem()).Elem()
		if !dataToStructSetScalar(value, entry) {
			return f.errorf("can't hold a value of the map, which is of type %T", entry)
		}
		target.SetMapIndex(reflect.ValueOf(key).Convert(mapType.Key()), value)
	}
	return nil
}

// dataToStructSetScalar sets target from raw and reports whether raw fits: raw
// has the Go type that the Terraform SDK gives a value of the kind of target,
// and a number doesn't overflow target. It leaves target untouched otherwise,
// and for a target of any other kind.
func dataToStructSetScalar(target reflect.Value, raw any) bool {
	switch value := raw.(type) {
	case string:
		if target.Kind() == reflect.String {
			target.SetString(value)
			return true
		}
	case bool:
		if target.Kind() == reflect.Bool {
			target.SetBool(value)
			return true
		}
	case int:
		if target.CanInt() && !target.OverflowInt(int64(value)) {
			target.SetInt(int64(value))
			return true
		}
	case float64:
		if target.CanFloat() && !target.OverflowFloat(value) {
			target.SetFloat(value)
			return true
		}
	}
	return false
}
