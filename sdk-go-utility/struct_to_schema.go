package sdkgoutility

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// sdkTypesPath is the import path of the package that holds the well-known
// types of the modular SDK: Time, Duration and FieldMask.
const sdkTypesPath = "github.com/databricks/sdk-go/core/types"

// StructToSchema makes an SDKv2 schema from v, a modular SDK struct or a
// pointer to one, and returns what customize makes of it. A nil customize
// leaves the schema as built. It does for modular SDK structs what
// common.StructToSchema does for classic SDK structs.
//
// Each exported field becomes one attribute, named by FieldName:
//
//   - string, a named string type (an enum), the integer, float and bool
//     kinds, and a pointer to any of them give TypeString, TypeInt, TypeFloat
//     and TypeBool;
//   - types.Time and types.Duration of the SDK's core/types package, and
//     pointers to them, give TypeString, since they are strings in JSON;
//   - a message struct, or a pointer to one, gives a TypeList with MaxItems 1
//     of a nested resource, built by the same rules;
//   - a slice gives a TypeList of nested resources for message structs, and
//     of the matching type for the scalars above;
//   - a map of strings to strings gives a TypeMap of strings;
//   - a types.FieldMask is left out: it is request plumbing, not an attribute.
//
// Every attribute is Optional, at every depth, because the types of the
// modular SDK don't say which fields are required or computed. The caller
// marks those in customize, with common.SetRequired, common.SetReadOnly and
// common.CustomizeSchemaPath.
//
// StructToSchema panics with an error that names the struct, the field and
// its type for a field that it can't turn into an attribute: an interface (a
// oneof group or a body stream), bytes ([]byte and json.RawMessage), a slice of anything but
// message structs and scalars, a map of anything but strings to strings, a
// pointer to anything but a message struct or a scalar, an embedded field, a
// message that contains itself, and any other kind, such as an array. It
// also panics when v is not a struct or a pointer to one.
//
// StructToSchema deliberately doesn't:
//
//   - read json or tf tags, so it never sets Required, Computed, ForceNew,
//     Sensitive, Default, a diff suppression or a set. Classic SDK structs and
//     the provider's own structs stay with common.StructToSchema.
//   - flatten an embedded struct, look up the resource providers that package
//     common registers, or apply aliases. Modular SDK messages need none.
//   - cut a message that contains itself at some depth. A depth that fits one
//     resource is wrong for the next, so it panics instead of guessing.
func StructToSchema(v any, customize func(map[string]*schema.Schema) map[string]*schema.Schema) map[string]*schema.Schema {
	structType := reflect.TypeOf(v)
	if structType != nil && structType.Kind() == reflect.Pointer {
		structType = structType.Elem()
	}
	if structType == nil || structType.Kind() != reflect.Struct {
		panic(fmt.Errorf("sdkgoutility.StructToSchema: a struct or a pointer to a struct is expected, but got %T", v))
	}
	scm := messageSchema(structType, map[reflect.Type]bool{})
	if customize != nil {
		scm = customize(scm)
	}
	return scm
}

// messageSchema returns the attributes of the message structType. open holds
// the messages that are being built further up, which is how a message that
// contains itself is found.
func messageSchema(structType reflect.Type, open map[reflect.Type]bool) map[string]*schema.Schema {
	open[structType] = true
	defer delete(open, structType)
	scm := map[string]*schema.Schema{}
	for i := range structType.NumField() {
		field := structType.Field(i)
		if field.Anonymous {
			panic(unsupportedField(structType, field, "embedded fields are not supported"))
		}
		name := FieldName(structType, field)
		if name == "" || isFieldMask(field.Type) {
			continue
		}
		scm[name] = fieldSchema(structType, field, open)
	}
	return scm
}

// fieldSchema returns the attribute for field, a field of structType.
func fieldSchema(structType reflect.Type, field reflect.StructField, open map[reflect.Type]bool) *schema.Schema {
	fieldType := field.Type
	isPointer := fieldType.Kind() == reflect.Pointer
	if isPointer {
		fieldType = fieldType.Elem()
	}
	attribute := &schema.Schema{Optional: true}
	if scalar, ok := scalarType(fieldType); ok {
		attribute.Type = scalar
		return attribute
	}
	switch {
	case fieldType.Kind() == reflect.Struct:
		attribute.Type = schema.TypeList
		attribute.MaxItems = 1
		attribute.Elem = nestedResource(structType, field, fieldType, open)
	case isPointer:
		panic(unsupportedField(structType, field, "a pointer is supported only to a message struct or a scalar"))
	case fieldType.Kind() == reflect.Slice:
		attribute.Type = schema.TypeList
		attribute.Elem = sliceElem(structType, field, open)
	case fieldType.Kind() == reflect.Map:
		if fieldType.Key().Kind() != reflect.String || fieldType.Elem().Kind() != reflect.String {
			panic(unsupportedField(structType, field, "a map is supported only of strings to strings"))
		}
		attribute.Type = schema.TypeMap
		attribute.Elem = schema.TypeString
	case fieldType.Kind() == reflect.Interface:
		panic(unsupportedField(structType, field, "an interface, such as a oneof group or a body stream, has no schema type"))
	default:
		panic(unsupportedField(structType, field, "this kind of type has no schema type"))
	}
	return attribute
}

// sliceElem returns the Elem of the list attribute for field, a slice field
// of structType: a resource for message structs and a schema for scalars.
func sliceElem(structType reflect.Type, field reflect.StructField, open map[reflect.Type]bool) any {
	elemType := field.Type.Elem()
	// encoding/json writes a slice of any uint8 type as one base64 string.
	if elemType.Kind() == reflect.Uint8 {
		panic(unsupportedField(structType, field, "bytes have no schema type"))
	}
	if scalar, ok := scalarType(elemType); ok {
		return &schema.Schema{Type: scalar}
	}
	if elemType.Kind() == reflect.Struct {
		return nestedResource(structType, field, elemType, open)
	}
	panic(unsupportedField(structType, field, "a slice is supported only of message structs and scalars"))
}

// nestedResource returns the resource for the message messageType, which
// field of structType holds directly, through a pointer or in a slice.
func nestedResource(structType reflect.Type, field reflect.StructField, messageType reflect.Type, open map[reflect.Type]bool) *schema.Resource {
	if open[messageType] {
		panic(unsupportedField(structType, field, fmt.Sprintf("the message %s contains itself", messageType)))
	}
	return &schema.Resource{Schema: messageSchema(messageType, open)}
}

// scalarType returns the schema type of an attribute that holds one value of
// type t, and false when t is not such a type.
func scalarType(t reflect.Type) (schema.ValueType, bool) {
	switch t.Kind() {
	case reflect.String:
		return schema.TypeString, true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return schema.TypeInt, true
	case reflect.Float32, reflect.Float64:
		return schema.TypeFloat, true
	case reflect.Bool:
		return schema.TypeBool, true
	case reflect.Struct:
		if name := wellKnownTypeName(t); name == "Time" || name == "Duration" {
			return schema.TypeString, true
		}
	}
	return schema.TypeInvalid, false
}

// isFieldMask tells whether t is a types.FieldMask or a pointer to one.
func isFieldMask(t reflect.Type) bool {
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	return wellKnownTypeName(t) == "FieldMask"
}

// wellKnownTypeName returns the name of t without its type arguments when the
// SDK's core/types package declares t, and the empty string otherwise.
func wellKnownTypeName(t reflect.Type) string {
	if t.PkgPath() != sdkTypesPath {
		return ""
	}
	name, _, _ := strings.Cut(t.Name(), "[")
	return name
}

// unsupportedField returns the error that StructToSchema panics with for
// field, a field of structType that it can't turn into an attribute.
func unsupportedField(structType reflect.Type, field reflect.StructField, reason string) error {
	return fmt.Errorf("sdkgoutility.StructToSchema: field %s of %s has unsupported type %s: %s",
		field.Name, structType, field.Type, reason)
}
