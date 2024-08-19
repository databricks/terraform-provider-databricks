package pluginframework

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/databricks/terraform-provider-databricks/internal/reflect_utils"
)

// Converts a tfsdk struct into a gosdk struct, with the folowing rules.
//
//	types.String -> string
//	types.Bool -> bool
//	types.Int64 -> int64
//	types.Float64 -> float64
//	types.String -> string
//
// NOTE:
//
//	ForceSendFields are populated for string, bool, int64, float64 on non null values
//	types.list and types.map are not supported
//	map keys should always be a string
//	tfsdk structs use types.String for all enum values
func TfSdkToGoSdkStruct(tfsdk interface{}, gosdk interface{}, ctx context.Context) diag.Diagnostics {
	srcVal := reflect.ValueOf(tfsdk)
	destVal := reflect.ValueOf(gosdk)

	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}

	if destVal.Kind() != reflect.Ptr {
		return diag.Diagnostics{diag.NewErrorDiagnostic(fmt.Sprintf("please provide a pointer for the gosdk struct, got %s", destVal.Type().Name()), "tfsdk to gosdk struct conversion failure")}
	}
	destVal = destVal.Elem()

	if srcVal.Kind() != reflect.Struct {
		return diag.Diagnostics{diag.NewErrorDiagnostic(fmt.Sprintf("input should be structs, got %s,", srcVal.Type().Name()), "tfsdk to gosdk struct conversion failure")}
	}

	forceSendFieldsField := destVal.FieldByName("ForceSendFields")

	allFields := reflect_utils.ListAllFields(srcVal)
	for _, field := range allFields {
		srcField := field.V
		srcFieldName := field.Sf.Name

		srcFieldTag := field.Sf.Tag.Get("tfsdk")
		if srcFieldTag == "-" {
			continue
		}

		destField := destVal.FieldByName(srcFieldName)

		err := tfSdkToGoSdkSingleField(srcField, destField, srcFieldName, &forceSendFieldsField, ctx)
		if err != nil {
			return diag.Diagnostics{diag.NewErrorDiagnostic(err.Error(), "tfsdk to gosdk field conversion failure")}
		}
	}

	return diag.Diagnostics{}
}

func tfSdkToGoSdkSingleField(srcField reflect.Value, destField reflect.Value, srcFieldName string, forceSendFieldsField *reflect.Value, ctx context.Context) error {

	if !destField.IsValid() {
		// Skip field that destination struct does not have.
		logger.Tracef(ctx, fmt.Sprintf("field skipped in tfsdk to gosdk conversion: destination struct does not have field %s", srcFieldName))
		return nil
	}

	if !destField.CanSet() {
		panic(fmt.Errorf("destination field can not be set: %s", destField.Type().Name()))
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
		if TfSdkToGoSdkStruct(srcFieldValue, destField.Interface(), ctx).ErrorsCount() > 0 {
			panic("Error converting tfsdk to gosdk struct")
		}
	} else if srcField.Kind() == reflect.Struct {
		tfsdkToGoSdkStructField(srcField, destField, srcFieldName, forceSendFieldsField, ctx)
	} else if srcField.Kind() == reflect.Slice {
		if srcField.IsNil() {
			// Skip nils
			return nil
		}
		destSlice := reflect.MakeSlice(destField.Type(), srcField.Len(), srcField.Cap())
		for j := 0; j < srcField.Len(); j++ {

			srcElem := srcField.Index(j)

			destElem := destSlice.Index(j)
			if err := tfSdkToGoSdkSingleField(srcElem, destElem, "", nil, ctx); err != nil {
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
			if err := tfSdkToGoSdkSingleField(srcMapValue, destMapValue, "", nil, ctx); err != nil {
				return err
			}
			destMap.SetMapIndex(destMapKey, destMapValue)
		}
		destField.Set(destMap)
	} else {
		panic(fmt.Errorf("unknown type for field: %s", srcField.Type().Name()))
	}
	return nil

}

func tfsdkToGoSdkStructField(srcField reflect.Value, destField reflect.Value, srcFieldName string, forceSendFieldsField *reflect.Value, ctx context.Context) {
	srcFieldValue := srcField.Interface()
	switch v := srcFieldValue.(type) {
	case types.Bool:
		destField.SetBool(v.ValueBool())
		if !v.IsNull() {
			addToForceSendFields(srcFieldName, forceSendFieldsField)
		}
	case types.Int64:
		destField.SetInt(v.ValueInt64())
		if !v.IsNull() {
			addToForceSendFields(srcFieldName, forceSendFieldsField)
		}
	case types.Float64:
		destField.SetFloat(v.ValueFloat64())
		if !v.IsNull() {
			addToForceSendFields(srcFieldName, forceSendFieldsField)
		}
	case types.String:
		if destField.Type().Name() != "string" {
			// This is the case for enum.

			// Skip unset value.
			if srcField.IsZero() {
				return
			}

			// Find the Set method
			destVal := reflect.New(destField.Type())
			setMethod := destVal.MethodByName("Set")
			if !setMethod.IsValid() {
				panic(fmt.Sprintf("set method not found on enum type: %s", destField.Type().Name()))
			}

			// Prepare the argument for the Set method
			arg := reflect.ValueOf(v.ValueString())

			// Call the Set method
			result := setMethod.Call([]reflect.Value{arg})
			if len(result) != 0 {
				if err, ok := result[0].Interface().(error); ok && err != nil {
					panic(err)
				}
			}
			// We don't need to set ForceSendFields for enums because the value is never going to be a zero value (empty string).
			destField.Set(destVal.Elem())
		} else {
			destField.SetString(v.ValueString())
			if !v.IsNull() {
				addToForceSendFields(srcFieldName, forceSendFieldsField)
			}
		}
	case types.List:
		panic("types.List should never be used, use go native slices instead")
	case types.Map:
		panic("types.Map should never be used, use go native maps instead")
	default:
		if srcField.IsZero() {
			// Skip zeros
			return
		}
		// If it is a real stuct instead of a tfsdk type, recursively resolve it.
		if TfSdkToGoSdkStruct(srcFieldValue, destField.Addr().Interface(), ctx).ErrorsCount() > 0 {
			panic("Error converting tfsdk to gosdk struct")
		}
	}
}

func addToForceSendFields(fieldName string, forceSendFieldsField *reflect.Value) {
	if forceSendFieldsField == nil || !forceSendFieldsField.IsValid() || !forceSendFieldsField.CanSet() {
		log.Printf("[Debug] forceSendFieldsField is nil, invalid or not settable. %s", fieldName)
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
