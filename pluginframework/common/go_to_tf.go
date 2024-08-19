package pluginframework

import (
	"context"
	"fmt"
	"reflect"

	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/databricks/terraform-provider-databricks/internal/reflect_utils"
)

// Converts a gosdk struct into a tfsdk struct, with the folowing rules.
//
//	string -> types.String
//	bool -> types.Bool
//	int64 -> types.Int64
//	float64 -> types.Float64
//	string -> types.String
//
// NOTE:
//
//	If field name doesn't show up in ForceSendFields and the field is zero value, we set the null value on the tfsdk.
//	types.list and types.map are not supported
//	map keys should always be a string
//	tfsdk structs use types.String for all enum values
//	non-json fields will be omitted
func GoSdkToTfSdkStruct(gosdk interface{}, tfsdk interface{}, ctx context.Context) diag.Diagnostics {

	srcVal := reflect.ValueOf(gosdk)
	destVal := reflect.ValueOf(tfsdk)

	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}

	if destVal.Kind() != reflect.Ptr {
		return diag.Diagnostics{diag.NewErrorDiagnostic("please provide a pointer for the tfsdk struct", "tfsdk to gosdk struct conversion failure")}
	}
	destVal = destVal.Elem()

	if srcVal.Kind() != reflect.Struct || destVal.Kind() != reflect.Struct {
		return diag.Diagnostics{diag.NewErrorDiagnostic(fmt.Sprintf("input should be structs %s, %s", srcVal.Type().Name(), destVal.Type().Name()), "tfsdk to gosdk struct conversion failure")}
	}

	var forceSendFieldVal []string

	forceSendField := srcVal.FieldByName("ForceSendFields")
	if !forceSendField.IsValid() {
		// If no forceSendField, just use an empty list.
		forceSendFieldVal = []string{}
	} else {
		forceSendFieldVal = forceSendField.Interface().([]string)
	}

	for _, field := range reflect_utils.ListAllFields(srcVal) {
		srcField := field.V
		srcFieldName := field.Sf.Name

		srcFieldTag := field.Sf.Tag.Get("json")
		if srcFieldTag == "-" {
			continue
		}
		destField := destVal.FieldByName(srcFieldName)

		if !destField.IsValid() {
			logger.Tracef(ctx, fmt.Sprintf("field skipped in gosdk to tfsdk conversion: destination struct does not have field %s", srcFieldName))
			continue
		}

		err := goSdkToTfSdkSingleField(srcField, destField, fieldInForceSendFields(srcFieldName, forceSendFieldVal), ctx)
		if err != nil {
			return diag.Diagnostics{diag.NewErrorDiagnostic(err.Error(), "gosdk to tfsdk field conversion failure")}
		}
	}
	return nil
}

func goSdkToTfSdkSingleField(srcField reflect.Value, destField reflect.Value, forceSendField bool, ctx context.Context) error {

	if !destField.CanSet() {
		panic(fmt.Errorf("destination field can not be set: %s", destField.Type().Name()))
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
		destField.Set(reflect.New(destField.Type().Elem()))

		// Recursively populate the nested struct.
		if diags := GoSdkToTfSdkStruct(srcFieldValue, destField.Interface(), ctx); diags.ErrorsCount() > 0 {
			panic("Error converting gosdk to tfsdk struct")
		}
	case reflect.Bool:
		boolVal := srcFieldValue.(bool)
		// check if the value is non-zero or if the field is in the forceSendFields list
		if boolVal || forceSendField {
			destField.Set(reflect.ValueOf(types.BoolValue(boolVal)))
		} else {
			destField.Set(reflect.ValueOf(types.BoolNull()))
		}
	case reflect.Int64:
		// convert any kind of integer to int64
		intVal := srcField.Convert(reflect.TypeOf(int64(0))).Interface().(int64)
		// check if the value is non-zero or if the field is in the forceSendFields list
		if intVal != 0 || forceSendField {
			destField.Set(reflect.ValueOf(types.Int64Value(int64(intVal))))
		} else {
			destField.Set(reflect.ValueOf(types.Int64Null()))
		}
	case reflect.Float64:
		// convert any kind of float to float64
		float64Val := srcField.Convert(reflect.TypeOf(float64(0))).Interface().(float64)
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
		// resolve the nested struct by recursively calling the function
		if GoSdkToTfSdkStruct(srcFieldValue, destField.Addr().Interface(), ctx).ErrorsCount() > 0 {
			panic("Error converting gosdk to tfsdk struct")
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
			if err := goSdkToTfSdkSingleField(srcElem, destElem, true, ctx); err != nil {
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
			if err := goSdkToTfSdkSingleField(srcMapValue, destMapValue, true, ctx); err != nil {
				return err
			}
			destMap.SetMapIndex(destMapKey, destMapValue)
		}
		destField.Set(destMap)
	default:
		panic(fmt.Errorf("unknown type for field: %s", srcField.Type().Name()))
	}
	return nil
}

func getStringFromEnum(srcField reflect.Value) string {
	var stringMethod reflect.Value
	if srcField.CanAddr() {
		stringMethod = srcField.Addr().MethodByName("String")
	} else {
		// Create a new addressable variable to call the String method
		addr := reflect.New(srcField.Type()).Elem()
		addr.Set(srcField)
		stringMethod = addr.Addr().MethodByName("String")
	}
	if stringMethod.IsValid() {
		stringResult := stringMethod.Call(nil)
		if len(stringResult) == 1 {
			return stringResult[0].Interface().(string)
		} else {
			panic("num get string has more than one result")
		}
	} else {
		panic("enum does not have valid .String() method")
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
