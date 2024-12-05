package converters

import (
	"context"
	"fmt"
	"reflect"

	tfcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/tfreflect"
)

const tfSdkToGoSdkStructConversionFailureMessage = "tfsdk to gosdk struct conversion failure"
const tfSdkToGoSdkFieldConversionFailureMessage = "tfsdk to gosdk field conversion failure"

// TfSdkToGoSdkStruct converts a tfsdk struct into a gosdk struct, with the folowing rules.
//
//	types.String -> string
//	types.Bool -> bool
//	types.Int64 -> int64
//	types.Float64 -> float64
//
// NOTE:
//
// # Structs are represented as slice of structs in tfsdk, and pointers are removed
//
// ForceSendFields are populated for string, bool, int64, float64 on non null values
// types.list and types.map are not supported
// map keys should always be a string
// tfsdk structs use types.String for all enum values
func TfSdkToGoSdkStruct(ctx context.Context, tfsdk interface{}, gosdk interface{}) (d diag.Diagnostics) {
	srcVal := reflect.ValueOf(tfsdk)
	destVal := reflect.ValueOf(gosdk)

	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}

	if destVal.Kind() != reflect.Ptr {
		d.AddError(tfSdkToGoSdkStructConversionFailureMessage, fmt.Sprintf("please provide a pointer for the gosdk struct, got %s", destVal.Type().Name()))
		return
	}
	destVal = destVal.Elem()

	if srcVal.Kind() != reflect.Struct {
		d.AddError(tfSdkToGoSdkStructConversionFailureMessage, fmt.Sprintf("input should be structs, got %s,", srcVal.Type().Name()))
		return
	}

	forceSendFieldsField := destVal.FieldByName("ForceSendFields")

	var innerTypes map[string]reflect.Type
	if cftp, ok := tfsdk.(tfcommon.ComplexFieldTypeProvider); ok {
		innerTypes = cftp.GetComplexFieldTypes(ctx)
	}

	allFields := tfreflect.ListAllFields(srcVal)
	for _, field := range allFields {
		srcField := field.Value
		srcFieldName := field.StructField.Name

		srcFieldTag := field.StructField.Tag.Get("tfsdk")
		if srcFieldTag == "-" {
			continue
		}

		destField := destVal.FieldByName(srcFieldName)
		innerType := innerTypes[srcFieldTag]

		d.Append(tfSdkToGoSdkSingleField(ctx, srcField, destField, srcFieldName, &forceSendFieldsField, innerType)...)
		if d.HasError() {
			return
		}
	}

	return nil
}

func tfSdkToGoSdkSingleField(
	ctx context.Context,
	srcField reflect.Value,
	destField reflect.Value,
	srcFieldName string,
	forceSendFieldsField *reflect.Value,
	innerType reflect.Type) (d diag.Diagnostics) {

	if !destField.IsValid() {
		// Skip field that destination struct does not have.
		tflog.Trace(ctx, fmt.Sprintf("field skipped in tfsdk to gosdk conversion: destination struct does not have field %s", srcFieldName))
		return
	}

	if !destField.CanSet() {
		d.AddError(tfSdkToGoSdkFieldConversionFailureMessage, fmt.Sprintf("destination field can not be set: %T. %s", destField.Type(), common.TerraformBugErrorMessage))
		return
	}

	if srcField.Kind() != reflect.Struct {
		d.AddError(tfSdkToGoSdkFieldConversionFailureMessage, fmt.Sprintf("unexpected type %T in tfsdk structs, expected a plugin framework type. %s", srcField.Interface(), common.TerraformBugErrorMessage))
		return
	}

	// The field being processed must be an attr.Value (a field of the TF SDK struct).
	v, ok := srcField.Interface().(attr.Value)
	if !ok {
		d.AddError(tfSdkToGoSdkFieldConversionFailureMessage, fmt.Sprintf("unexpected type %T in tfsdk structs, expected a plugin framework type. %s", v, common.TerraformBugErrorMessage))
		return
	}

	if v.IsUnknown() {
		return
	}
	if shouldSetForceSendFields(v, destField) {
		addToForceSendFields(ctx, srcFieldName, forceSendFieldsField)
	}

	d.Append(tfsdkToGoSdkStructField(ctx, v, destField, srcFieldName, forceSendFieldsField, innerType)...)
	return
}

func tfsdkToGoSdkStructField(
	ctx context.Context,
	srcFieldValue attr.Value,
	destField reflect.Value,
	srcFieldName string,
	forceSendFieldsField *reflect.Value,
	innerType reflect.Type) (d diag.Diagnostics) {
	switch v := srcFieldValue.(type) {
	case types.Bool:
		destField.SetBool(v.ValueBool())
	case types.Int64:
		destField.SetInt(v.ValueInt64())
	case types.Float64:
		destField.SetFloat(v.ValueFloat64())
	case types.String:
		if destField.Type().Name() != "string" {
			// This is the case for enum.

			// Skip unset value.
			if v.ValueString() == "" {
				return
			}

			destVal := convertToEnumValue(v, destField.Type())
			destField.Set(destVal)
		} else {
			destField.SetString(v.ValueString())
		}
	case types.List:
		// Empty lists correspond to nil slices or the struct zero value.
		if v.IsNull() {
			return
		}

		// Read the nested elements into the TFSDK struct slice
		// This is a slice of either TFSDK structs or bools, ints, strings, and floats from the TF plugin framework types.
		innerValue := reflect.New(reflect.SliceOf(innerType))
		d.Append(v.ElementsAs(ctx, innerValue.Interface(), true)...)
		if d.HasError() {
			return
		}

		// Recursively call TFSDK to GOSDK conversion for each element in the list. If this corresponds to a slice,
		// the target type is the slice element type. If it corresponds to a struct, the target type is the struct type.
		// If it corresponds to a pointer, the target type is the type pointed to by the pointer.
		var destInnerType reflect.Type
		if destField.Type().Kind() == reflect.Slice {
			destInnerType = destField.Type().Elem()
		} else {
			if innerValue.Elem().Len() > 1 {
				d.AddError(tfSdkToGoSdkFieldConversionFailureMessage, fmt.Sprintf("The length of a slice can not be greater than 1 if it is representing a struct, %s", common.TerraformBugErrorMessage))
				return
			}
			// Case of types.List <-> struct or ptr
			if destField.Type().Kind() == reflect.Ptr {
				destInnerType = destField.Type().Elem()
			} else {
				destInnerType = destField.Type()
			}
		}

		// Recursively call TFSDK to GOSDK conversion for each element in the list
		converted := reflect.MakeSlice(reflect.SliceOf(destInnerType), 0, innerValue.Elem().Len())
		for i := 0; i < innerValue.Elem().Len(); i++ {
			vv := innerValue.Elem().Index(i).Interface()
			nextDest := reflect.New(destInnerType)
			// If the element is a primitive type, we can convert it by recursively calling this function.
			// Otherwise, it is a TF SDK struct, and we need to call TfSdkToGoSdkStruct to convert it.
			switch typedVv := vv.(type) {
			case types.Bool, types.String, types.Int64, types.Float64:
				d.Append(tfsdkToGoSdkStructField(ctx, typedVv.(attr.Value), nextDest.Elem(), srcFieldName, forceSendFieldsField, innerType)...)
			default:
				d.Append(TfSdkToGoSdkStruct(ctx, vv, nextDest.Interface())...)
			}
			if d.HasError() {
				return
			}
			converted = reflect.Append(converted, reflect.Indirect(nextDest))
		}

		if destField.Type().Kind() == reflect.Slice {
			destField.Set(converted)
		} else if destField.Type().Kind() == reflect.Ptr {
			destField.Set(converted.Index(0).Addr())
		} else {
			destField.Set(converted.Index(0))
		}
	case types.Map:
		// Empty maps correspond to nil maps or the struct zero value.
		if v.IsNull() {
			return
		}

		// Read the nested elements into the TFSDK struct map
		// This is a map from string to either TFSDK structs or bools, ints, strings, and floats from the TF plugin framework types.
		innerValue := reflect.New(reflect.MapOf(reflect.TypeOf(""), innerType))
		d.Append(v.ElementsAs(ctx, innerValue.Interface(), true)...)
		if d.HasError() {
			return
		}

		// Recursively call TFSDK to GOSDK conversion for each element in the map
		destType := destField.Type().Elem()
		converted := reflect.MakeMap(reflect.MapOf(reflect.TypeOf(""), destType))
		for _, key := range innerValue.Elem().MapKeys() {
			vv := innerValue.Elem().MapIndex(key).Interface()
			nextDest := reflect.New(destType)
			// If the element is a primitive type, we can convert it by recursively calling this function.
			// Otherwise, it is a TF SDK struct, and we need to call TfSdkToGoSdkStruct to convert it.
			switch typedVv := vv.(type) {
			case types.Bool, types.String, types.Int64, types.Float64:
				d.Append(tfsdkToGoSdkStructField(ctx, typedVv.(attr.Value), nextDest.Elem(), srcFieldName, forceSendFieldsField, innerType)...)
			default:
				d.Append(TfSdkToGoSdkStruct(ctx, vv, nextDest.Interface())...)
			}
			if d.HasError() {
				return
			}
			converted.SetMapIndex(key, nextDest.Elem())
		}

		destField.Set(converted)
	case types.Object:
		if v.IsNull() {
			return
		}

		innerValue := reflect.New(innerType)
		d.Append(v.As(ctx, innerValue.Interface(), basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
		if d.HasError() {
			return
		}

		d.Append(TfSdkToGoSdkStruct(ctx, innerValue.Interface(), destField.Addr().Interface())...)
	case types.Set, types.Tuple:
		d.AddError(tfSdkToGoSdkFieldConversionFailureMessage, fmt.Sprintf("%T is not currently supported as a source field. %s", v, common.TerraformBugErrorMessage))
		return
	default:
		d.AddError(tfSdkToGoSdkFieldConversionFailureMessage, fmt.Sprintf("unexpected type %T in tfsdk structs, expected a plugin framework type. %s", v, common.TerraformBugErrorMessage))
		return
	}
	return
}

func shouldSetForceSendFields(srcFieldValue attr.Value, destField reflect.Value) bool {
	if srcFieldValue.IsNull() {
		return false
	}
	// Don't set forceSendFields for enums
	// We don't need to set ForceSendFields for enums because the value is never going to be a zero value (empty string).
	if _, ok := srcFieldValue.(types.String); ok && destField.Type().Name() != "string" {
		return false
	}
	// Don't set forceSendFields for lists
	if _, ok := srcFieldValue.(types.List); ok && destField.Kind() == reflect.Slice {
		return false
	}

	return true
}

func addToForceSendFields(ctx context.Context, fieldName string, forceSendFieldsField *reflect.Value) {
	if forceSendFieldsField == nil || !forceSendFieldsField.IsValid() || !forceSendFieldsField.CanSet() {
		tflog.Debug(ctx, fmt.Sprintf("[Debug] forceSendFieldsField is nil, invalid or not settable. %s", fieldName))
		return
	}
	// Initialize forceSendFields if it is a zero Value
	if forceSendFieldsField.IsZero() {
		// Create a new slice of strings
		newSlice := []string{}
		forceSendFieldsField.Set(reflect.ValueOf(&newSlice).Elem())
	}
	forceSendFields := forceSendFieldsField.Interface().([]string)
	forceSendFields = append(forceSendFields, fieldName)
	forceSendFieldsField.Set(reflect.ValueOf(forceSendFields))
}

// Returns a reflect.Value of the enum type with the value set to the given string.
func convertToEnumValue(v types.String, destType reflect.Type) reflect.Value {
	// Find the Set method
	destVal := reflect.New(destType)
	setMethod := destVal.MethodByName("Set")
	if !setMethod.IsValid() {
		panic(fmt.Sprintf("set method not found on enum type: %s. %s", destType.Name(), common.TerraformBugErrorMessage))
	}

	// Prepare the argument for the Set method
	arg := reflect.ValueOf(v.ValueString())

	// Call the Set method
	result := setMethod.Call([]reflect.Value{arg})
	if len(result) != 0 {
		if err, ok := result[0].Interface().(error); ok && err != nil {
			panic(fmt.Sprintf("%s. %s", err, common.TerraformBugErrorMessage))
		}
	}
	return destVal.Elem()
}
