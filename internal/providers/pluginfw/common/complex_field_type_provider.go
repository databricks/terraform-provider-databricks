package common

import (
	"context"
	"reflect"
)

// ComplexFieldTypeProvider must be implemented by any TFSDK structure that contains
// a complex field (list, map, object). Such fields do not include sufficient type
// information to understand the type of the contained elements in the case of a list
// or map, or the fields in the case of an object. This interface enables callers
// to recover that information.
type ComplexFieldTypeProvider interface {
	// GetComplexFieldTypes returns a map from field name to the type of the value in
	// the list, map or object. The keys of the map must match the value of the
	// `tfsdk` tag on the field. There must be one entry in the map for each field
	// that has type types.List, types.Map or types.Object.
	//
	// If the field has type types.List or types.Map, the reflect.Type instance may
	// correspond to either a primitive value (e.g. types.String) or a TFSDK structure.
	// It is not allowed to return a reflect.Type that corresponds to a type value
	// (e.g. types.StringType).
	//
	// If the field has type types.Object, the reflect.Type instance must correspond
	// to a TFSDK structure.
	GetComplexFieldTypes(context.Context) map[string]reflect.Type
}
