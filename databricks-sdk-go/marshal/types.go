package marshal

import "reflect"

var basicTypes = map[reflect.Kind]bool{

	reflect.Bool: true,

	reflect.String: true,

	// Integers
	reflect.Int:     true,
	reflect.Int8:    true,
	reflect.Int16:   true,
	reflect.Int32:   true,
	reflect.Int64:   true,
	reflect.Uint:    true,
	reflect.Uint8:   true,
	reflect.Uint16:  true,
	reflect.Uint32:  true,
	reflect.Uint64:  true,
	reflect.Uintptr: true,

	// Floats
	reflect.Float32: true,
	reflect.Float64: true,

	// Complex
	reflect.Complex64:  true,
	reflect.Complex128: true,
}

func isBasicType(v reflect.Type) bool {
	b, ok := basicTypes[v.Kind()]
	return b && ok
}
