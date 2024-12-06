package common

import (
	"context"
	"fmt"
	"reflect"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/tfreflect"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type ObjectValuable struct {
	// A TF SDK structure.
	// If this contains types.List, types.Map, or types.Object, it must implement the
	// ComplexFieldTypesProvider interface.
	inner any
}

// Construct a new ObjectValuable.
// TFSDK structs automatically implement ObjectValuable, so they are returned as-is.
// Hand-written structs do not necessarily implement ObjectValuable, so this is a
// convenience implementation using reflection.
func NewObjectValuable(inner any) ObjectValuable {
	if ov, ok := inner.(ObjectValuable); ok {
		return ov
	}
	return ObjectValuable{inner: inner}
}

// Equal implements basetypes.ObjectValuable.
func (o ObjectValuable) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ObjectValuable) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ObjectValuable) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ObjectValuable) String() string {
	return fmt.Sprintf("%v", o.inner)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ObjectValuable) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o.inner,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ObjectValuable) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o ObjectValuable) Type(ctx context.Context) attr.Type {
	attrs := map[string]attr.Type{}

	// Tolerate pointers.
	rv := reflect.Indirect(reflect.ValueOf(o.inner))
	for _, field := range tfreflect.ListAllFields(rv) {
		typeField := field.StructField
		fieldName := typeField.Tag.Get("tfsdk")
		if fieldName == "-" {
			continue
		}
		// If it is a simple type, we can determine the type from the reflect.Type.
		if t, ok := getAttrType(field.Value); ok {
			attrs[fieldName] = t
			continue
		}

		// Otherwise, additional metadata is required to determine the type of the list elements.
		// This is available via the ComplexFieldTypeProvider interface, implemented on the parent type.
		provider, ok := o.inner.(ComplexFieldTypeProvider)
		if !ok {
			panic(fmt.Errorf("complex field types not provided for type: %T. %s", o.inner, common.TerraformBugErrorMessage))
		}
		complexFieldTypes := provider.GetComplexFieldTypes(ctx)
		fieldType, ok := complexFieldTypes[fieldName]
		if !ok {
			panic(fmt.Errorf("complex field type not found for field %s on type %T. %s", typeField.Name, o.inner, common.TerraformBugErrorMessage))
		}

		// This is either a "simple" type or a TF SDK structure.
		var innerType attr.Type
		if t, ok := getAttrType(fieldType); ok {
			innerType = t
		} else {
			// If this is a TF SDK structure, we need to recursively determine the type.
			nested := reflect.New(fieldType).Elem().Interface()
			ov := ObjectValuable{inner: nested}
			innerType = ov.Type(ctx)
		}

		switch field.Value.Interface().(type) {
		case types.List:
			attrs[fieldName] = types.ListType{ElemType: innerType}
		case types.Map:
			attrs[fieldName] = types.MapType{ElemType: innerType}
		case types.Object:
			// Objects are only used for nested structures, not primitives, so we must go through
			// the else case above.
			innerType, ok = innerType.(basetypes.ObjectType)
			if !ok {
				panic(fmt.Errorf("expected ObjectType, got %T", innerType))
			}
			attrs[fieldName] = innerType
		}
	}

	return basetypes.ObjectType{
		AttrTypes: attrs,
	}
}

var simpleTypeMap = map[reflect.Type]attr.Type{
	reflect.TypeOf(types.Bool{}):    types.BoolType,
	reflect.TypeOf(types.Int64{}):   types.Int64Type,
	reflect.TypeOf(types.Float64{}): types.Float64Type,
	reflect.TypeOf(types.String{}):  types.StringType,
}

// getAttrType returns the attr.Type for the given value. The value can be a
// reflect.Type instance or a Terraform type instance.
func getAttrType(v any) (attr.Type, bool) {
	if r, ok := v.(reflect.Type); ok {
		t, ok := simpleTypeMap[r]
		return t, ok
	}
	if rv, ok := v.(reflect.Value); ok {
		t, ok := simpleTypeMap[rv.Type()]
		return t, ok
	}
	t, ok := simpleTypeMap[reflect.TypeOf(v)]
	return t, ok
}

var _ basetypes.ObjectValuable = ObjectValuable{}
