package common

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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

	if srcVal.Kind() != reflect.Struct {
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
		// Skip field that destination struct does not have.
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
			panic("types.List should never be used, use go native slices instead")
		case types.Map:
			panic("types.Map should never be used, use go native maps instead")
		default:
			if srcField.IsNil() {
				// Skip nils
				return nil
			}
			// If it is a real stuct instead of a tfsdk type, recursively resolve it.
			if err := TfSdkToGoSdkStruct(srcFieldValue, destField.Addr().Interface(), ctx); err != nil {
				return err
			}
		}
	} else if srcField.Kind() == reflect.Slice {
		if srcField.IsNil() {
			// Skip nils
			return nil
		}
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
		// Skip field that destination struct does not have.
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
		var strVal string
		if srcField.Type().Name() != "string" {
			// This case is for Enum Types.
			stringMethod := srcField.Addr().MethodByName("String")
			if stringMethod.IsValid() {
				stringResult := stringMethod.Call(nil)
				if len(stringResult) == 1 {
					strVal = stringResult[0].Interface().(string)
				} else {
					log.Printf("[DEBUG] Enum get string has more than one result")
					strVal = ""
				}
			} else {
				log.Printf("[DEBUG] Enum does not have valid .String() method")
				strVal = ""
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
		if srcField.IsNil() {
			// Skip nils
			return nil
		}
		// resolve the nested struct by recursively calling the function
		if err := GoSdkToTfSdkStruct(srcFieldValue, destField.Addr().Interface(), ctx); err != nil {
			return err
		}
	case reflect.Slice:
		if srcField.IsNil() {
			// Skip nils
			return nil
		}
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

func addToForceSendFields(fieldName string, forceSendFieldsField *reflect.Value) {
	if forceSendFieldsField == nil || !forceSendFieldsField.IsValid() || !forceSendFieldsField.CanSet() {
		log.Printf("[Debug] forceSendFieldsField is nil, invalid or not settable for %s", fieldName)
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

func checkTheStringInForceSendFields(fieldName string, forceSendFields []string) bool {
	for _, field := range forceSendFields {
		if field == fieldName {
			return true
		}
	}
	return false
}

func pluginFrameworkResourceTypeToSchema(v reflect.Value) map[string]schema.Attribute {
	scm := map[string]schema.Attribute{}
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
		isOptional := fieldIsOptional(typeField)
		// For now everything is marked as optional. TODO: add tf annotations for computed, optional, etc.
		kind := typeField.Type.Kind()
		if kind == reflect.Ptr {
			elem := typeField.Type.Elem()
			sv := reflect.New(elem).Elem()
			nestedScm := pluginFrameworkResourceTypeToSchema(sv)
			scm[fieldName] = schema.SingleNestedAttribute{Attributes: nestedScm, Optional: isOptional, Required: !isOptional}
		} else if kind == reflect.Slice {
			elem := typeField.Type.Elem()
			if elem.Kind() == reflect.Ptr {
				elem = elem.Elem()
			}
			if elem.Kind() != reflect.Struct {
				panic(fmt.Errorf("unsupported slice value for %s: %s", fieldName, reflectKind(elem.Kind())))
			}
			switch elem {
			case reflect.TypeOf(types.Bool{}):
				scm[fieldName] = schema.ListAttribute{ElementType: types.BoolType, Optional: isOptional, Required: !isOptional}
			case reflect.TypeOf(types.Int64{}):
				scm[fieldName] = schema.ListAttribute{ElementType: types.Int64Type, Optional: isOptional, Required: !isOptional}
			case reflect.TypeOf(types.Float64{}):
				scm[fieldName] = schema.ListAttribute{ElementType: types.Float64Type, Optional: isOptional, Required: !isOptional}
			case reflect.TypeOf(types.String{}):
				scm[fieldName] = schema.ListAttribute{ElementType: types.StringType, Optional: isOptional, Required: !isOptional}
			default:
				// Nested struct
				nestedScm := pluginFrameworkResourceTypeToSchema(reflect.New(elem).Elem())
				scm[fieldName] = schema.ListNestedAttribute{NestedObject: schema.NestedAttributeObject{Attributes: nestedScm}, Optional: isOptional, Required: !isOptional}
			}
		} else if kind == reflect.Map {
			elem := typeField.Type.Elem()
			if elem.Kind() == reflect.Ptr {
				elem = elem.Elem()
			}
			if elem.Kind() != reflect.Struct {
				panic(fmt.Errorf("unsupported map value for %s: %s", fieldName, reflectKind(elem.Kind())))
			}
			switch elem {
			case reflect.TypeOf(types.Bool{}):
				scm[fieldName] = schema.MapAttribute{ElementType: types.BoolType, Optional: isOptional, Required: !isOptional}
			case reflect.TypeOf(types.Int64{}):
				scm[fieldName] = schema.MapAttribute{ElementType: types.Int64Type, Optional: isOptional, Required: !isOptional}
			case reflect.TypeOf(types.Float64{}):
				scm[fieldName] = schema.MapAttribute{ElementType: types.Float64Type, Optional: isOptional, Required: !isOptional}
			case reflect.TypeOf(types.String{}):
				scm[fieldName] = schema.MapAttribute{ElementType: types.StringType, Optional: isOptional, Required: !isOptional}
			default:
				// Nested struct
				nestedScm := pluginFrameworkResourceTypeToSchema(reflect.New(elem).Elem())
				scm[fieldName] = schema.MapNestedAttribute{NestedObject: schema.NestedAttributeObject{Attributes: nestedScm}, Optional: isOptional, Required: !isOptional}
			}
		} else if kind == reflect.Struct {
			switch field.v.Interface().(type) {
			case types.Bool:
				scm[fieldName] = schema.BoolAttribute{Optional: isOptional, Required: !isOptional}
			case types.Int64:
				scm[fieldName] = schema.Int64Attribute{Optional: isOptional, Required: !isOptional}
			case types.Float64:
				scm[fieldName] = schema.Float64Attribute{Optional: isOptional, Required: !isOptional}
			case types.String:
				scm[fieldName] = schema.StringAttribute{Optional: isOptional, Required: !isOptional}
			case types.List:
				panic("types.List should never be used in tfsdk structs")
			case types.Map:
				panic("types.Map should never be used in tfsdk structs")
			default:
				// If it is a real stuct instead of a tfsdk type, recursively resolve it.
				elem := typeField.Type
				sv := reflect.New(elem)
				nestedScm := pluginFrameworkResourceTypeToSchema(sv)
				scm[fieldName] = schema.SingleNestedAttribute{Attributes: nestedScm, Optional: isOptional, Required: !isOptional}
			}
		} else if kind == reflect.String {
			// This case is for Enum types.
			scm[fieldName] = schema.StringAttribute{Optional: isOptional, Required: !isOptional}
		} else {
			panic(fmt.Errorf("unknown type for field: %s", typeField.Name))
		}
	}
	return scm
}

func fieldIsOptional(field reflect.StructField) bool {
	tagValue := field.Tag.Get("tf")
	return strings.Contains(tagValue, "optional")
}

func PluginFrameworkResourceStructToSchema(v any, customizeSchema func(CustomizableSchemaPluginFramework) CustomizableSchemaPluginFramework) schema.Schema {
	attributes := PluginFrameworkResourceStructToSchemaMap(v, customizeSchema)
	return schema.Schema{Attributes: attributes}
}

func PluginFrameworkResourceStructToSchemaMap(v any, customizeSchema func(CustomizableSchemaPluginFramework) CustomizableSchemaPluginFramework) map[string]schema.Attribute {
	attributes := pluginFrameworkResourceTypeToSchema(reflect.ValueOf(v))

	if customizeSchema != nil {
		cs := customizeSchema(*ConstructCustomizableSchema(attributes))
		return cs.ToAttributeMap()
	} else {
		return attributes
	}
}
