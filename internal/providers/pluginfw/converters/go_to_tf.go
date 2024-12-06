package converters

import (
	"context"
	"fmt"
	"reflect"

	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/tfreflect"
)

const goSdkToTfSdkStructConversionFailureMessage = "gosdk to tfsdk struct conversion failure"
const goSdkToTfSdkFieldConversionFailureMessage = "gosdk to tfsdk field conversion failure"

// GoSdkToTfSdkStruct converts a gosdk struct into a tfsdk struct, with the folowing rules.
//
//	string -> types.String
//	bool -> types.Bool
//	int64 -> types.Int64
//	float64 -> types.Float64
//	string -> types.String
//
// NOTE:
//
// # Structs in gosdk are represented as slices of structs in tfsdk, and pointers are removed
//
// If field name doesn't show up in ForceSendFields and the field is zero value, we set the null value on the tfsdk.
// types.list and types.map are not supported
// map keys should always be a string
// tfsdk structs use types.String for all enum values
// non-json fields will be omitted
func GoSdkToTfSdkStruct(ctx context.Context, gosdk interface{}, tfsdk interface{}) diag.Diagnostics {

	srcVal := reflect.ValueOf(gosdk)
	destVal := reflect.ValueOf(tfsdk)

	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}

	if destVal.Kind() != reflect.Ptr {
		return diag.Diagnostics{diag.NewErrorDiagnostic(goSdkToTfSdkStructConversionFailureMessage, fmt.Sprintf("please provide a pointer for the tfsdk struct, got %s", destVal.Type().Name()))}
	}
	destVal = destVal.Elem()

	if srcVal.Kind() != reflect.Struct || destVal.Kind() != reflect.Struct {
		return diag.Diagnostics{diag.NewErrorDiagnostic(goSdkToTfSdkStructConversionFailureMessage, fmt.Sprintf("input should be structs %s, %s", srcVal.Type().Name(), destVal.Type().Name()))}
	}

	var forceSendFieldVal []string

	forceSendField := srcVal.FieldByName("ForceSendFields")
	if !forceSendField.IsValid() {
		// If no forceSendField, just use an empty list.
		forceSendFieldVal = []string{}
	} else {
		forceSendFieldVal = forceSendField.Interface().([]string)
	}

	for _, field := range tfreflect.ListAllFields(srcVal) {
		srcField := field.Value
		srcFieldName := field.StructField.Name

		srcFieldTag := field.StructField.Tag.Get("json")
		if srcFieldTag == "-" {
			continue
		}
		destField := destVal.FieldByName(srcFieldName)

		if !destField.IsValid() {
			logger.Tracef(ctx, fmt.Sprintf("field skipped in gosdk to tfsdk conversion: destination struct does not have field %s", srcFieldName))
			continue
		}

		err := goSdkToTfSdkSingleField(ctx, srcField, destField, fieldInForceSendFields(srcFieldName, forceSendFieldVal))
		if err != nil {
			return diag.Diagnostics{diag.NewErrorDiagnostic(goSdkToTfSdkFieldConversionFailureMessage, err.Error())}
		}
	}
	return nil
}

func goSdkToTfSdkSingleField(ctx context.Context, srcField reflect.Value, destField reflect.Value, forceSendField bool) error {

	if !destField.CanSet() {
		panic(fmt.Errorf("destination field can not be set: %s. %s", destField.Type().Name(), common.TerraformBugErrorMessage))
	}

	srcFieldValue := srcField.Interface()

	if srcFieldValue == nil {
		return nil
	}

	switch srcField.Kind() {
	case reflect.Ptr:
		if srcField.IsNil() {
			// Skip nils
			return nil
		}

		var fieldToSetInterface any

		if destField.Kind() == reflect.Slice {
			sliceType := destField.Type()
			newSlice := reflect.MakeSlice(sliceType, 1, 1)
			newSlice.Index(0).Set(reflect.New(sliceType.Elem()).Elem())

			destField.Set(newSlice)
			fieldToSetInterface = newSlice.Index(0).Addr().Interface()
		} else {
			destField.Set(reflect.New(destField.Type().Elem()))
			fieldToSetInterface = destField.Interface()
		}

		// Recursively populate the nested struct.
		if GoSdkToTfSdkStruct(ctx, srcFieldValue, fieldToSetInterface).HasError() {
			panic(fmt.Sprintf("%s. %s", goSdkToTfSdkStructConversionFailureMessage, common.TerraformBugErrorMessage))
		}
	case reflect.Bool:
		boolVal := srcFieldValue.(bool)
		// check if the value is non-zero or if the field is in the forceSendFields list
		if boolVal || forceSendField {
			destField.Set(reflect.ValueOf(types.BoolValue(boolVal)))
		} else {
			destField.Set(reflect.ValueOf(types.BoolNull()))
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		// convert any kind of integer to int64
		intVal := srcField.Convert(reflect.TypeOf(int64(0))).Int()
		// check if the value is non-zero or if the field is in the forceSendFields list
		if intVal != 0 || forceSendField {
			destField.Set(reflect.ValueOf(types.Int64Value(intVal)))
		} else {
			destField.Set(reflect.ValueOf(types.Int64Null()))
		}
	case reflect.Float32, reflect.Float64:
		// convert any kind of float to float64
		float64Val := srcField.Convert(reflect.TypeOf(float64(0))).Float()
		// check if the value is non-zero or if the field is in the forceSendFields list
		if float64Val != 0 || forceSendField {
			destField.Set(reflect.ValueOf(types.Float64Value(float64Val)))
		} else {
			destField.Set(reflect.ValueOf(types.Float64Null()))
		}
	case reflect.String:
		var strVal string
		if srcField.Type().Name() != "string" {
			// This case is for Enum Types.
			strVal = getStringFromEnum(srcField)
		} else {
			strVal = srcFieldValue.(string)
		}
		// check if the value is non-zero or if the field is in the forceSendFields list
		if strVal != "" || forceSendField {
			destField.Set(reflect.ValueOf(types.StringValue(strVal)))
		} else {
			destField.Set(reflect.ValueOf(types.StringNull()))
		}
	case reflect.Struct:
		if srcField.IsZero() {
			// Skip zeros
			return nil
		}
		var dest any
		if destField.Kind() == reflect.Slice {
			// allocate a slice first
			destSlice := reflect.MakeSlice(destField.Type(), 1, 1)
			destField.Set(destSlice)
			dest = destSlice.Index(0).Addr().Interface()
		} else {
			dest = destField.Addr().Interface()
		}
		// resolve the nested struct by recursively calling the function
		if GoSdkToTfSdkStruct(ctx, srcFieldValue, dest).HasError() {
			panic(fmt.Sprintf("%s. %s", goSdkToTfSdkStructConversionFailureMessage, common.TerraformBugErrorMessage))
		}
	case reflect.Slice:
		if srcField.IsNil() {
			// Skip nils
			return nil
		}
		destSlice := reflect.MakeSlice(destField.Type(), srcField.Len(), srcField.Cap())
		for j := 0; j < srcField.Len(); j++ {

			srcElem := srcField.Index(j)

			destElem := destSlice.Index(j)
			if err := goSdkToTfSdkSingleField(ctx, srcElem, destElem, true); err != nil {
				return err
			}
		}
		destField.Set(destSlice)
	case reflect.Map:
		if srcField.IsNil() {
			// Skip nils
			return nil
		}
		destMap := reflect.MakeMap(destField.Type())
		for _, key := range srcField.MapKeys() {
			srcMapValue := srcField.MapIndex(key)
			destMapValue := reflect.New(destField.Type().Elem()).Elem()
			destMapKey := reflect.ValueOf(key.Interface())
			if err := goSdkToTfSdkSingleField(ctx, srcMapValue, destMapValue, true); err != nil {
				return err
			}
			destMap.SetMapIndex(destMapKey, destMapValue)
		}
		destField.Set(destMap)
	default:
		panic(fmt.Errorf("unknown type for field: %s. %s", srcField.Type().Name(), common.TerraformBugErrorMessage))
	}
	return nil
}

// Get the string value of an enum by calling the .String() method on the enum object.
func getStringFromEnum(srcField reflect.Value) string {
	var stringMethod reflect.Value
	if srcField.CanAddr() {
		stringMethod = srcField.Addr().MethodByName("String")
	} else {
		// This case is for the unit tests because the enum values will be const and we cannot get the address.
		// If cannot get addr, create a new addressable variable to call the String method
		addr := reflect.New(srcField.Type()).Elem()
		addr.Set(srcField)
		stringMethod = addr.Addr().MethodByName("String")
	}
	if stringMethod.IsValid() {
		stringResult := stringMethod.Call(nil)
		if len(stringResult) == 1 {
			return stringResult[0].Interface().(string)
		} else {
			panic(fmt.Sprintf("num get string has more than one result. %s", common.TerraformBugErrorMessage))
		}
	} else {
		panic(fmt.Sprintf("enum does not have valid .String() method. %s", common.TerraformBugErrorMessage))
	}
}

func fieldInForceSendFields(fieldName string, forceSendFields []string) bool {
	for _, field := range forceSendFields {
		if field == fieldName {
			return true
		}
	}
	return false
}
