package converters

import (
	"context"
	"fmt"
	"reflect"

	"github.com/databricks/databricks-sdk-go/logger"
	tfcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
//	types.String -> string
//
// NOTE:
//
// # Structs are represented as slice of structs in tfsdk, and pointers are removed
//
// ForceSendFields are populated for string, bool, int64, float64 on non null values
// types.list and types.map are not supported
// map keys should always be a string
// tfsdk structs use types.String for all enum values
func TfSdkToGoSdkStruct(ctx context.Context, tfsdk interface{}, gosdk interface{}) diag.Diagnostics {
	srcVal := reflect.ValueOf(tfsdk)
	destVal := reflect.ValueOf(gosdk)

	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}

	if destVal.Kind() != reflect.Ptr {
		return diag.Diagnostics{diag.NewErrorDiagnostic(tfSdkToGoSdkStructConversionFailureMessage, fmt.Sprintf("please provide a pointer for the gosdk struct, got %s", destVal.Type().Name()))}
	}
	destVal = destVal.Elem()

	if srcVal.Kind() != reflect.Struct {
		return diag.Diagnostics{diag.NewErrorDiagnostic(tfSdkToGoSdkStructConversionFailureMessage, fmt.Sprintf("input should be structs, got %s,", srcVal.Type().Name()))}
	}

	forceSendFieldsField := destVal.FieldByName("ForceSendFields")

	var innerTypes map[string]reflect.Type
	if cftp, ok := tfsdk.(tfcommon.ComplexFieldTypeProvider); ok {
		innerTypes = cftp.GetComplexFieldTypes()
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

		err := tfSdkToGoSdkSingleField(ctx, srcField, destField, srcFieldName, &forceSendFieldsField, innerType)
		if err != nil {
			return diag.Diagnostics{diag.NewErrorDiagnostic(tfSdkToGoSdkFieldConversionFailureMessage, err.Error())}
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
	innerType reflect.Type) error {

	if !destField.IsValid() {
		// Skip field that destination struct does not have.
		logger.Tracef(ctx, fmt.Sprintf("field skipped in tfsdk to gosdk conversion: destination struct does not have field %s", srcFieldName))
		return nil
	}

	if !destField.CanSet() {
		panic(fmt.Errorf("destination field can not be set: %T. %s", destField.Type(), common.TerraformBugErrorMessage))
	}
	srcFieldValue := srcField.Interface()

	if srcFieldValue == nil {
		return nil
	} else if srcField.Kind() == reflect.Ptr {
		if srcField.IsNil() {
			// Skip nils
			return nil
		}
		// Allocate new memory for the destination field
		destField.Set(reflect.New(destField.Type().Elem()))

		// Recursively populate the nested struct.
		if TfSdkToGoSdkStruct(ctx, srcFieldValue, destField.Interface()).HasError() {
			panic(fmt.Sprintf("%s. %s", tfSdkToGoSdkStructConversionFailureMessage, common.TerraformBugErrorMessage))
		}
	} else if srcField.Kind() == reflect.Slice && destField.Kind() == reflect.Struct {
		if srcField.IsNil() {
			// Skip nils
			return nil
		}
		assertStructSliceLengthIsOne(srcField)
		tfsdkToGoSdkStructField(srcField.Index(0), destField, srcFieldName, forceSendFieldsField, ctx, innerType)
	} else if srcField.Kind() == reflect.Slice && destField.Kind() == reflect.Ptr {
		if srcField.IsNil() {
			// Skip nils
			return nil
		}
		destField.Set(reflect.New(destField.Type().Elem()))

		assertStructSliceLengthIsOne(srcField)

		// Recursively populate the nested struct.
		if TfSdkToGoSdkStruct(ctx, srcField.Index(0).Interface(), destField.Interface()).HasError() {
			panic(fmt.Sprintf("%s. %s", tfSdkToGoSdkStructConversionFailureMessage, common.TerraformBugErrorMessage))
		}
	} else if srcField.Kind() == reflect.Struct {
		tfsdkToGoSdkStructField(srcField, destField, srcFieldName, forceSendFieldsField, ctx, innerType)
	} else if srcField.Kind() == reflect.Slice {
		if srcField.IsNil() {
			// Skip nils
			return nil
		}
		destSlice := reflect.MakeSlice(destField.Type(), srcField.Len(), srcField.Cap())
		for j := 0; j < srcField.Len(); j++ {

			srcElem := srcField.Index(j)

			destElem := destSlice.Index(j)
			if err := tfSdkToGoSdkSingleField(ctx, srcElem, destElem, "", nil, innerType); err != nil {
				return err
			}
		}
		destField.Set(destSlice)
	} else if srcField.Kind() == reflect.Map {
		if srcField.IsNil() {
			// Skip nils
			return nil
		}
		destMap := reflect.MakeMap(destField.Type())
		for _, key := range srcField.MapKeys() {
			srcMapValue := srcField.MapIndex(key)
			destMapValue := reflect.New(destField.Type().Elem()).Elem()
			destMapKey := reflect.ValueOf(key.Interface())
			if err := tfSdkToGoSdkSingleField(ctx, srcMapValue, destMapValue, "", nil, innerType); err != nil {
				return err
			}
			destMap.SetMapIndex(destMapKey, destMapValue)
		}
		destField.Set(destMap)
	} else {
		panic(fmt.Errorf("unknown type for field: %s. %s", srcField.Type().Name(), common.TerraformBugErrorMessage))
	}
	return nil

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

func tfsdkToGoSdkStructField(
	srcField reflect.Value,
	destField reflect.Value,
	srcFieldName string,
	forceSendFieldsField *reflect.Value,
	ctx context.Context,
	innerType reflect.Type) {
	srcFieldValue := srcField.Interface()
	switch v := srcFieldValue.(type) {
	case types.Bool:
		destField.SetBool(v.ValueBool())
		if !v.IsNull() {
			addToForceSendFields(ctx, srcFieldName, forceSendFieldsField)
		}
	case types.Int64:
		destField.SetInt(v.ValueInt64())
		if !v.IsNull() {
			addToForceSendFields(ctx, srcFieldName, forceSendFieldsField)
		}
	case types.Float64:
		destField.SetFloat(v.ValueFloat64())
		if !v.IsNull() {
			addToForceSendFields(ctx, srcFieldName, forceSendFieldsField)
		}
	case types.String:
		if destField.Type().Name() != "string" {
			// This is the case for enum.

			// Skip unset value.
			if srcField.IsZero() || v.ValueString() == "" {
				return
			}

			destVal := convertToEnumValue(v, destField.Type())
			// We don't need to set ForceSendFields for enums because the value is never going to be a zero value (empty string).
			destField.Set(destVal)
		} else {
			destField.SetString(v.ValueString())
			if !v.IsNull() {
				addToForceSendFields(ctx, srcFieldName, forceSendFieldsField)
			}
		}
	case types.List:
		if v.IsNull() || v.IsUnknown() {
			return
		}

		var d diag.Diagnostics
		innerValue := reflect.New(reflect.SliceOf(innerType))

		// Read the nested elements into the TFSDK struct slice
		d = v.ElementsAs(ctx, innerValue.Interface(), true)
		if d.HasError() {
			for _, diag := range d {
				tflog.Error(ctx, fmt.Sprintf("[ERROR] %s: %s", diag.Summary(), diag.Detail()))
			}
			panic("failed to convert list")
		}

		// Recursively call TFSDK to GOSDK conversion for each element in the list
		var destInnerType reflect.Type
		if destField.Type().Kind() == reflect.Slice {
			destInnerType = destField.Type().Elem()
		} else {
			assertStructSliceLengthIsOne(innerValue.Elem())
			// Case of types.List <-> struct or ptr
			if destField.Type().Kind() == reflect.Ptr {
				destInnerType = destField.Type().Elem()
			} else {
				destInnerType = destField.Type()
			}
		}

		converted := reflect.MakeSlice(reflect.SliceOf(destInnerType), 0, innerValue.Elem().Len())
		for i := 0; i < innerValue.Elem().Len(); i++ {
			next := reflect.New(destInnerType).Elem()
			err := tfSdkToGoSdkSingleField(ctx, innerValue.Elem().Index(i), next, srcFieldName, forceSendFieldsField, innerType)
			if err != nil {
				panic(err)
			}
			converted = reflect.Append(converted, next)
		}

		if destField.Type().Kind() == reflect.Slice {
			destField.Set(converted)
		} else if destField.Type().Kind() == reflect.Ptr {
			destField.Set(converted.Index(0).Addr())
		} else {
			destField.Set(converted.Index(0))
		}
	case types.Object, types.Map, types.Set, types.Tuple:
		panic(fmt.Sprintf("%T should never be used, use go native maps instead. %s", v, common.TerraformBugErrorMessage))
	default:
		if srcField.IsZero() {
			// Skip zeros
			return
		}
		// If it is a real stuct instead of a tfsdk type, recursively resolve it.
		if TfSdkToGoSdkStruct(ctx, srcFieldValue, destField.Addr().Interface()).HasError() {
			panic(fmt.Sprintf("%s. %s", tfSdkToGoSdkStructConversionFailureMessage, common.TerraformBugErrorMessage))
		}
	}
}

func assertStructSliceLengthIsOne(srcSlice reflect.Value) {
	if srcSlice.Len() > 1 {
		panic(fmt.Sprintf("The length of a slice can not be greater than 1 if it is representing a struct, %s", common.TerraformBugErrorMessage))
	}
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
