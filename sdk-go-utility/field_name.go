package sdkgoutility

import (
	"encoding/json"
	"reflect"
	"strings"
	"sync"
	"unicode"
)

// FieldName returns the name that the Databricks API uses for field, a field
// of the modular SDK struct structType. It is the key of the field in the JSON
// that the MarshalJSON method of the struct writes.
//
// Modular SDK structs carry no json tag, so FieldName asks the SDK: it sets
// only that field on an otherwise empty message, marshals the message and
// reads the key that appears. It does so once for all fields of a struct and
// keeps the result. When that doesn't yield a name, the name is the fieldmask
// tag of the field, which the SDK generator sets to the API name on the
// messages that field masks address, and otherwise the snake_case form of the
// Go name, as fieldNameSnakeCase derives it. That is the case when:
//
//   - structType has no MarshalJSON method, as the wrapper struct of a oneof
//     variant, or is nil or not a struct;
//   - the field can't be given a value, as the interface field of a oneof
//     group, a body stream or a nested struct that is not a pointer;
//   - the message doesn't marshal with the field set, or the field doesn't add
//     exactly one key.
//
// structType is the struct that declares field, or a pointer to it. FieldName
// needs it because the Go name alone doesn't decide the API name: DisplayName
// is displayName in SCIM and display_name elsewhere, and a reflect.StructField
// doesn't say which struct it comes from. FieldName returns the empty string
// for an unexported field.
//
// FieldName deliberately doesn't:
//
//   - read json or tf tags. Classic SDK structs and the provider's own structs
//     stay with the helpers in package common.
//   - tell whether the API has the field in its JSON at all. A oneof group, a
//     body stream and a field that travels in an HTTP header get their
//     snake_case name like any other field, although no JSON has that key. A
//     helper that walks a struct recognizes those shapes itself.
//   - look into the type of the field. It names one field and never builds
//     the path of a nested one.
func FieldName(structType reflect.Type, field reflect.StructField) string {
	if !field.IsExported() {
		return ""
	}
	if name, ok := fieldNameProbeCached(structType)[field.Name]; ok {
		return name
	}
	if name := field.Tag.Get("fieldmask"); name != "" {
		return name
	}
	return fieldNameSnakeCase(field.Name)
}

var fieldNameMarshalerType = reflect.TypeFor[json.Marshaler]()

// fieldNameProbeCache holds the result of fieldNameProbe for every struct type
// that FieldName has seen, as a map[string]string by reflect.Type.
var fieldNameProbeCache sync.Map

// fieldNameProbeCached returns the JSON keys of the fields of structType by Go
// name, as far as fieldNameProbe finds them. It returns nil for a type that
// fieldNameProbe can't work on.
func fieldNameProbeCached(structType reflect.Type) map[string]string {
	for structType != nil && structType.Kind() == reflect.Pointer {
		structType = structType.Elem()
	}
	if structType == nil || structType.Kind() != reflect.Struct ||
		!reflect.PointerTo(structType).Implements(fieldNameMarshalerType) {
		return nil
	}
	names, ok := fieldNameProbeCache.Load(structType)
	if !ok {
		names, _ = fieldNameProbeCache.LoadOrStore(structType, fieldNameProbe(structType))
	}
	return names.(map[string]string)
}

// fieldNameProbe finds the JSON key of every field of structType that it can.
// It sets one field at a time on an otherwise empty message, marshals the
// message and takes the one key that the empty message doesn't have. Comparing
// with the empty message keeps a key that the message always writes from
// hiding the keys of the other fields.
func fieldNameProbe(structType reflect.Type) map[string]string {
	names := map[string]string{}
	empty := fieldNameJSONKeys(reflect.New(structType))
	for i := range structType.NumField() {
		message := reflect.New(structType)
		if !structType.Field(i).IsExported() || !fieldNameSetNonEmpty(message.Elem().Field(i)) {
			continue
		}
		var added []string
		for key := range fieldNameJSONKeys(message) {
			if _, always := empty[key]; !always {
				added = append(added, key)
			}
		}
		if len(added) == 1 {
			names[structType.Field(i).Name] = added[0]
		}
	}
	return names
}

// fieldNameJSONKeys returns the keys of the JSON object that message marshals
// to. It returns nil when message fails to marshal or marshals to something
// else.
func fieldNameJSONKeys(message reflect.Value) map[string]json.RawMessage {
	data, err := json.Marshal(message.Interface())
	if err != nil {
		return nil
	}
	var keys map[string]json.RawMessage
	if err := json.Unmarshal(data, &keys); err != nil {
		return nil
	}
	return keys
}

// fieldNameSetNonEmpty gives field a value that the omitempty option of a json
// tag doesn't drop, and reports whether it could. It can't for an interface,
// which needs a concrete type that only the declaring package knows, and for a
// struct, which omitempty never drops.
func fieldNameSetNonEmpty(field reflect.Value) bool {
	switch field.Kind() {
	case reflect.Pointer:
		field.Set(reflect.New(field.Type().Elem()))
	case reflect.String:
		field.SetString("x")
	case reflect.Bool:
		field.SetBool(true)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		field.SetInt(1)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		field.SetUint(1)
	case reflect.Float32, reflect.Float64:
		field.SetFloat(1)
	case reflect.Slice:
		if field.Type().Elem().Kind() == reflect.Uint8 {
			// Bytes that are JSON, because the field can be a json.RawMessage.
			field.SetBytes([]byte("0"))
		} else {
			field.Set(reflect.MakeSlice(field.Type(), 1, 1))
		}
	case reflect.Map:
		entry := reflect.MakeMapWithSize(field.Type(), 1)
		entry.SetMapIndex(reflect.Zero(field.Type().Key()), reflect.Zero(field.Type().Elem()))
		field.Set(entry)
	default:
		return false
	}
	return true
}

// fieldNameSnakeCase inverts the naming of the SDK generator, which builds the
// Go name of a field by capitalizing the first letter of every
// underscore-separated word of its API name: registered_model_id becomes
// RegisteredModelId. So a capital letter starts a new word, and nothing else
// does:
//
//   - Every capital in a run of capitals is a word of its own. The generator
//     capitalizes only the first letter of a word (http_status is HttpStatus
//     and t_shirt_size is TShirtSize), so the rule doesn't guess acronyms:
//     HTTPStatus gives h_t_t_p_status, not http_status.
//   - A digit never starts a word and stays with the letters before it:
//     Ipv6Cidr gives ipv6_cidr, S3Uri gives s3_uri and Oauth2Token gives
//     oauth2_token. An API word that starts with a digit can't be told apart
//     from that: visualization_color_01 is VisualizationColor01 in Go, which
//     gives visualization_color01.
//   - An underscore that the Go name already has is kept, and the capital
//     after it adds no second one.
//
// It is only the fallback of FieldName, because not every API name is in
// snake_case: SCIM has displayName and $ref, for instance.
func fieldNameSnakeCase(goName string) string {
	var b strings.Builder
	var previous rune
	for i, r := range goName {
		if unicode.IsUpper(r) && i > 0 && previous != '_' {
			b.WriteByte('_')
		}
		b.WriteRune(unicode.ToLower(r))
		previous = r
	}
	return b.String()
}
