package pluginframework

import (
	"context"
	"fmt"
	"reflect"

	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Converts a go-sdk struct into a tfsdk struct.
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
		switch forceSendField.Interface().(type) {
		case []string:
		default:
			panic(fmt.Errorf("ForceSendField is not of type []string"))
		}
		forceSendFieldVal = forceSendField.Interface().([]string)
	}

	for _, field := range common.ListAllFields(srcVal) {
		srcField := field.V
		srcFieldName := field.Sf.Name

		destField := destVal.FieldByName(srcFieldName)
		srcFieldTag := field.Sf.Tag.Get("json")
		if srcFieldTag == "-" {
			continue
		}
		err := goSdkToTfSdkSingleField(srcField, destField, srcFieldName, forceSendFieldVal, false, ctx)
		if err != nil {
			return diag.Diagnostics{diag.NewErrorDiagnostic(err.Error(), "gosdk to tfsdk field conversion failure")}
		}
	}
	return nil
}

func goSdkToTfSdkSingleField(srcField reflect.Value, destField reflect.Value, srcFieldName string, forceSendField []string, alwaysAdd bool, ctx context.Context) error {

	if !destField.IsValid() {
		// Skip field that destination struct does not have.
		logger.Tracef(ctx, fmt.Sprintf("field skipped in gosdk to tfsdk conversion: destination struct does not have field %s", srcFieldName))
		return nil
	}

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
		// check if alwaysAdd is false or the value is zero or if the field is in the forceSendFields list
		if alwaysAdd || !(!boolVal && !checkTheStringInForceSendFields(srcFieldName, forceSendField)) {
			destField.Set(reflect.ValueOf(types.BoolValue(boolVal)))
		} else {
			destField.Set(reflect.ValueOf(types.BoolNull()))
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		// convert any kind of integer to int64
		intVal := srcField.Convert(reflect.TypeOf(int64(0))).Interface().(int64)
		// check if alwaysAdd is true or the value is zero or if the field is in the forceSendFields list
		if alwaysAdd || !(intVal == 0 && !checkTheStringInForceSendFields(srcFieldName, forceSendField)) {
			destField.Set(reflect.ValueOf(types.Int64Value(int64(intVal))))
		} else {
			destField.Set(reflect.ValueOf(types.Int64Null()))
		}
	case reflect.Float32, reflect.Float64:
		// convert any kind of float to float64
		float64Val := srcField.Convert(reflect.TypeOf(float64(0))).Interface().(float64)
		// check if alwaysAdd is true or the value is zero or if the field is in the forceSendFields list
		if alwaysAdd || !(float64Val == 0 && !checkTheStringInForceSendFields(srcFieldName, forceSendField)) {
			destField.Set(reflect.ValueOf(types.Float64Value(float64Val)))
		} else {
			destField.Set(reflect.ValueOf(types.Float64Null()))
		}
	case reflect.String:
		var strVal string
		if srcField.Type().Name() != "string" {
			// This case is for Enum Types.
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
					strVal = stringResult[0].Interface().(string)
				} else {
					panic("num get string has more than one result")
				}
			} else {
				panic("enum does not have valid .String() method")
			}
		} else {
			strVal = srcFieldValue.(string)
		}
		// check if alwaysAdd is false or the value is zero or if the field is in the forceSendFields list
		if alwaysAdd || !(strVal == "" && !checkTheStringInForceSendFields(srcFieldName, forceSendField)) {
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
			if err := goSdkToTfSdkSingleField(srcElem, destElem, "", nil, true, ctx); err != nil {
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
			if err := goSdkToTfSdkSingleField(srcMapValue, destMapValue, "", nil, true, ctx); err != nil {
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

func checkTheStringInForceSendFields(fieldName string, forceSendFields []string) bool {
	for _, field := range forceSendFields {
		if field == fieldName {
			return true
		}
	}
	return false
}
