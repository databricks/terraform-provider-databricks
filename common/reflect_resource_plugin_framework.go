package common

import (
	"context"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Converts a tfsdk struct into a go-sdk struct.
func TfSdkToGoSdkStruct(tfsdk interface{}, gosdk interface{}, ctx context.Context) error {
	srcVal := reflect.ValueOf(tfsdk)
	destVal := reflect.ValueOf(gosdk)

	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}

	if destVal.Kind() != reflect.Ptr {
		panic("please provide a pointer for the gosdk struct")
	}
	destVal = destVal.Elem()

	if srcVal.Kind() != reflect.Struct || destVal.Kind() != reflect.Struct {
		panic("input should be structs")
	}

	forceSendFieldsField := destVal.FieldByName("ForceSendFields")

	srcType := srcVal.Type()
	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Field(i)
		srcFieldName := srcType.Field(i).Name

		destField := destVal.FieldByName(srcFieldName)

		srcFieldTag := srcType.Field(i).Tag.Get("tfsdk")
		if srcFieldTag == "-" {
			continue
		}

		tfSdkToGoSdkSingleField(srcField, destField, srcFieldName, &forceSendFieldsField, ctx)
	}

	return nil
}

func tfSdkToGoSdkSingleField(srcField reflect.Value, destField reflect.Value, srcFieldName string, forceSendFieldsField *reflect.Value, ctx context.Context) error {

	if !destField.IsValid() {
		panic(fmt.Errorf("destination field is not valid: %s", destField.Type().Name()))
	}

	if !destField.CanSet() {
		panic(fmt.Errorf("destination field can not be set: %s", destField.Type().Name()))
	}
	srcFieldValue := srcField.Interface()

	if srcFieldValue == nil {
		return nil
	} else if srcField.Kind() == reflect.Ptr {
		// Allocate new memory for the destination field
		destField.Set(reflect.New(destField.Type().Elem()))

		// Recursively populate the nested struct.
		if err := TfSdkToGoSdkStruct(srcFieldValue, destField.Interface(), ctx); err != nil {
			return err
		}
	} else if srcField.Kind() == reflect.Struct {
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
			destField.SetString(v.ValueString())
			if !v.IsNull() {
				addToForceSendFields(srcFieldName, forceSendFieldsField)
			}
		case types.List:
			diag := v.ElementsAs(ctx, destField.Addr().Interface(), false)
			if len(diag) != 0 {
				panic("Error")
			}
		case types.Map:
			v.ElementsAs(ctx, destField.Addr().Interface(), false)
		default:
			// If it is a real stuct instead of a tfsdk type, recursively resolve it.
			if err := TfSdkToGoSdkStruct(srcFieldValue, destField.Addr().Interface(), ctx); err != nil {
				return err
			}
		}
	} else if srcField.Kind() == reflect.Slice {
		destSlice := reflect.MakeSlice(destField.Type(), srcField.Len(), srcField.Cap())
		for j := 0; j < srcField.Len(); j++ {
			nestedSrcField := srcField.Index(j)
			nestedSrcField.Kind()

			srcElem := srcField.Index(j)

			destElem := destSlice.Index(j)
			if err := tfSdkToGoSdkSingleField(srcElem, destElem, "", nil, ctx); err != nil {
				return err
			}
		}
		destField.Set(destSlice)
	} else if srcField.Kind() == reflect.Map {
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

// Converts a go-sdk struct into a tfsdk struct.
func GoSdkToTfSdkStruct(gosdk interface{}, tfsdk interface{}, ctx context.Context) error {

	srcVal := reflect.ValueOf(gosdk)
	destVal := reflect.ValueOf(tfsdk)

	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}

	if destVal.Kind() != reflect.Ptr {
		panic("please provide a pointer for the tfsdk struct")
	}
	destVal = destVal.Elem()

	if srcVal.Kind() != reflect.Struct || destVal.Kind() != reflect.Struct {
		panic(fmt.Errorf("input should be structs %s, %s", srcVal.Type().Name(), destVal.Type().Name()))
	}

	forceSendField := srcVal.FieldByName("ForceSendFields")
	if !forceSendField.IsValid() {
		panic(fmt.Errorf("go sdk struct does not have valid ForceSendField %s", srcVal.Type().Name()))
	}
	switch forceSendField.Interface().(type) {
	case []string:
	default:
		panic(fmt.Errorf("ForceSendField is not of type []string"))
	}
	forceSendFieldVal := forceSendField.Interface().([]string)
	srcType := srcVal.Type()
	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Field(i)
		srcFieldName := srcVal.Type().Field(i).Name

		destField := destVal.FieldByName(srcFieldName)
		srcFieldTag := srcType.Field(i).Tag.Get("json")
		if srcFieldTag == "-" {
			continue
		}
		goSdkToTfSdkSingleField(srcField, destField, srcFieldName, forceSendFieldVal, false, ctx)
	}
	return nil
}

func goSdkToTfSdkSingleField(srcField reflect.Value, destField reflect.Value, srcFieldName string, forceSendField []string, alwaysAdd bool, ctx context.Context) error {

	if !destField.IsValid() {
		panic(fmt.Errorf("destination field is not valid: %s", destField.Type().Name()))
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
		destField.Set(reflect.New(destField.Type().Elem()))

		// Recursively populate the nested struct.
		if err := GoSdkToTfSdkStruct(srcFieldValue, destField.Interface(), ctx); err != nil {
			return err
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
		strVal := srcFieldValue.(string)
		// check if alwaysAdd is false or the value is zero or if the field is in the forceSendFields list
		if alwaysAdd || !(strVal == "" && !checkTheStringInForceSendFields(srcFieldName, forceSendField)) {
			destField.Set(reflect.ValueOf(types.StringValue(strVal)))
		} else {
			destField.Set(reflect.ValueOf(types.StringNull()))
		}
	case reflect.Struct:
		// resolve the nested struct by recursively calling the function
		if err := GoSdkToTfSdkStruct(srcFieldValue, destField.Addr().Interface(), ctx); err != nil {
			return err
		}
	case reflect.Slice:
		destSlice := reflect.MakeSlice(destField.Type(), srcField.Len(), srcField.Cap())
		for j := 0; j < srcField.Len(); j++ {
			nestedSrcField := srcField.Index(j)
			nestedSrcField.Kind()

			srcElem := srcField.Index(j)

			destElem := destSlice.Index(j)
			if err := goSdkToTfSdkSingleField(srcElem, destElem, "", nil, true, ctx); err != nil {
				return err
			}
		}
		destField.Set(destSlice)
	case reflect.Map:
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

func addToForceSendFields(fieldName string, forceSendFieldsField *reflect.Value) {
	if forceSendFieldsField == nil {
		return
	}
	forceSendFields := forceSendFieldsField.Interface().([]string)
	forceSendFields = append(forceSendFields, fieldName)
	forceSendFieldsField.Set(reflect.ValueOf(forceSendFields))
}

func checkTheStringInForceSendFields(fieldName string, forceSendFields []string) bool {
	for _, field := range forceSendFields {
		if field == fieldName {
			return true
		}
	}
	return false
}

func pluginFrameworkTypeToSchema(v reflect.Value) map[string]schema.Attribute {
	rk := v.Kind()
	if rk == reflect.Ptr {
		v = v.Elem()
		rk = v.Kind()
	}
	if rk != reflect.Struct {
		panic(fmt.Errorf("Schema value of Struct is expected, but got %s: %#v", reflectKind(rk), v))
	}
	fields := listAllFields(v)
	for _, field := range fields {
		typeField := field.sf
		fieldName := typeField.Tag.Get("tfsdk")
		if fieldName == "-" {
			continue
		}
		// TODO: handle optional and all kinds of stuff
		kind := typeField.Type.Kind()
		if kind == reflect.Ptr {
			// In the schema ptr and non ptr are treated the same way
			// Seems like in the original typeToSchema implementation we assumed that the ptr can only be a ptr to a struct
			// If so, we just need to recursively call this function with its elem
			// Otherwise we need more helper functions

		} else if kind == reflect.Slice {
			// If it's a slice, we need to check what type it is
			//   If it's a slice of primitive types (tfsdk primiritive types), it's a schema.ListAttribute
			//   If it's a slice of nested types (struct), it's a schema.ListNestedAttribute
			//     If it's a list of struct, we can create a schema.NestedAttributeObject and call the function recursively
			//     It seems like the two scenarios below are not covered in the existing structToSchema
			//       If it's a list of list, what to do?
			//       If it's a list of map, what to do?

		} else if kind == reflect.Map {
			// If it's a slice, we need to check what type it is
			//   If it's a map with values of primitive types (tfsdk primiritive types), it's a schema.MapAttribute
			//   If it's a map with values of nested types (struct), it's a schema.MapNestedAttribute
			// Complicated nested maps are not covered for structToSchema's existing implementation

		} else if kind == reflect.Struct {
			switch field.v.Interface().(type) {
			case types.Bool:
				println("bool!")
				println(field.sf.Name)
			case types.Int64:
				println("int64!")
				println(field.sf.Name)
			case types.Float64:
				println("float64!")
				println(field.sf.Name)
			case types.String:
				println("string!")
				println(field.sf.Name)
			case types.List:
				println("list!")
				println(field.sf.Name)
			case types.Map:
				println("map!")
				println(field.sf.Name)
			default:
				// If it is a real stuct instead of a tfsdk type, recursively resolve it.
				println("real struct")
				println(field.sf.Name)
				return nil
			}

		} else {
			panic(fmt.Errorf("unknown type for field: %s", typeField.Name))
		}
	}
	return nil
}
