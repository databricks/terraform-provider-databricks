package marshal

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"reflect"
	"strconv"
	"strings"
)

var (
	int64Type      = reflect.TypeFor[int64]()
	rawMessageType = reflect.TypeFor[json.RawMessage]()
)

type jsonReplacement struct {
	start int
	end   int
	value []byte
}

type int64Normalizer struct {
	data         []byte
	decoder      *json.Decoder
	replacements []jsonReplacement
}

// normalizeInt64Strings rewrites quoted integers as JSON numbers wherever the
// destination type expects an int64. It uses only the destination's type and
// does not modify the destination value. This preprocessing is necessary
// because encoding/json cannot accept both quoted and unquoted integers for a
// plain int64 field. JSON values that do not require normalization preserve
// their original bytes.
func normalizeInt64Strings(data []byte, destination any) ([]byte, error) {
	normalizer := int64Normalizer{
		data:    data,
		decoder: json.NewDecoder(bytes.NewReader(data)),
	}
	if err := normalizer.normalizeValue(reflect.TypeOf(destination)); err != nil {
		return nil, err
	}
	if err := normalizer.decoder.Decode(&struct{}{}); err != io.EOF {
		if err == nil {
			return nil, errors.New("multiple JSON values")
		}
		return nil, err
	}
	if len(normalizer.replacements) == 0 {
		return data, nil
	}

	normalized := make([]byte, 0, len(data))
	last := 0
	for _, replacement := range normalizer.replacements {
		normalized = append(normalized, data[last:replacement.start]...)
		normalized = append(normalized, replacement.value...)
		last = replacement.end
	}
	normalized = append(normalized, data[last:]...)
	return normalized, nil
}

func (n *int64Normalizer) normalizeValue(target reflect.Type) error {
	for target != nil && target.Kind() == reflect.Pointer {
		target = target.Elem()
	}
	if target == rawMessageType {
		return n.skipValue()
	}
	if target == int64Type {
		return n.normalizeInt64()
	}
	if target == nil {
		return n.skipValue()
	}

	switch target.Kind() {
	case reflect.Struct:
		return n.normalizeStruct(target)
	case reflect.Array, reflect.Slice:
		return n.normalizeArray(target.Elem())
	case reflect.Map:
		return n.normalizeMap(target.Elem())
	default:
		return n.skipValue()
	}
}

func (n *int64Normalizer) normalizeInt64() error {
	raw, start, end, err := n.decodeRawMessage()
	if err != nil || len(raw) == 0 || raw[0] != '"' {
		return err
	}

	var text string
	if err := json.Unmarshal(raw, &text); err != nil {
		return err
	}
	value, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		return err
	}
	n.replacements = append(n.replacements, jsonReplacement{
		start: start,
		end:   end,
		value: strconv.AppendInt(nil, value, 10),
	})
	return nil
}

func (n *int64Normalizer) normalizeStruct(target reflect.Type) error {
	if n.nextValueByte() != '{' {
		return n.skipValue()
	}
	if _, err := n.decoder.Token(); err != nil {
		return err
	}
	for n.decoder.More() {
		keyStart := n.nextValuePosition()
		token, err := n.decoder.Token()
		if err != nil {
			return err
		}
		name, ok := token.(string)
		if !ok {
			return errors.New("invalid JSON object name")
		}

		field, ok := findJSONField(target, name)
		if !ok || field.asString {
			if err := n.skipValue(); err != nil {
				return err
			}
			continue
		}
		if field.name != name {
			encodedName, err := json.Marshal(field.name)
			if err != nil {
				return err
			}
			n.replacements = append(n.replacements, jsonReplacement{
				start: keyStart,
				end:   int(n.decoder.InputOffset()),
				value: encodedName,
			})
		}
		if err := n.normalizeValue(field.fieldType); err != nil {
			return err
		}
	}
	_, err := n.decoder.Token()
	return err
}

func (n *int64Normalizer) normalizeArray(elementType reflect.Type) error {
	if n.nextValueByte() != '[' {
		return n.skipValue()
	}
	if _, err := n.decoder.Token(); err != nil {
		return err
	}
	for n.decoder.More() {
		if err := n.normalizeValue(elementType); err != nil {
			return err
		}
	}
	_, err := n.decoder.Token()
	return err
}

func (n *int64Normalizer) normalizeMap(elementType reflect.Type) error {
	if n.nextValueByte() != '{' {
		return n.skipValue()
	}
	if _, err := n.decoder.Token(); err != nil {
		return err
	}
	for n.decoder.More() {
		if _, err := n.decoder.Token(); err != nil {
			return err
		}
		if err := n.normalizeValue(elementType); err != nil {
			return err
		}
	}
	_, err := n.decoder.Token()
	return err
}

func (n *int64Normalizer) skipValue() error {
	_, _, _, err := n.decodeRawMessage()
	return err
}

func (n *int64Normalizer) decodeRawMessage() (json.RawMessage, int, int, error) {
	var raw json.RawMessage
	if err := n.decoder.Decode(&raw); err != nil {
		return nil, 0, 0, err
	}
	end := int(n.decoder.InputOffset())
	return raw, end - len(raw), end, nil
}

func (n *int64Normalizer) nextValueByte() byte {
	position := n.nextValuePosition()
	if position >= len(n.data) {
		return 0
	}
	return n.data[position]
}

func (n *int64Normalizer) nextValuePosition() int {
	position := int(n.decoder.InputOffset())
	for position < len(n.data) {
		switch n.data[position] {
		case ' ', '\t', '\r', '\n', ',', ':':
			position++
		default:
			return position
		}
	}
	return position
}

type jsonField struct {
	name      string
	fieldType reflect.Type
	asString  bool
}

func findJSONField(target reflect.Type, name string) (jsonField, bool) {
	var folded jsonField
	for _, field := range reflect.VisibleFields(target) {
		tag := parseJSONTag(field)
		if tag.ignore || field.Anonymous || field.PkgPath != "" {
			continue
		}
		if tag.name == name {
			return jsonField{tag.name, field.Type, tag.asString}, true
		}
		if folded.fieldType == nil && strings.EqualFold(tag.name, name) {
			folded = jsonField{tag.name, field.Type, tag.asString}
		}
	}
	return folded, folded.fieldType != nil
}
