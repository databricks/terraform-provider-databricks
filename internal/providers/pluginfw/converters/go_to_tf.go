package converters

import (
	"context"
	"fmt"
	"reflect"

	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/databricks/terraform-provider-databricks/common"
	tfcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/tfreflect"
)

const goSdkToTfSdkStructConversionFailureMessage = "gosdk to tfsdk struct conversion failure"
const goSdkToTfSdkFieldConversionFailureMessage = "gosdk to tfsdk field conversion failure"

// GoSdkToTfSdkStruct converts a gosdk struct into a tfsdk struct, with the folowing rules.
//
// string -> types.String
// bool -> types.Bool
// int64 -> types.Int64
// float64 -> types.Float64
// Struct and pointer to struct -> types.List
// Slice -> types.List
// Map -> types.Map
//
// `gosdk` parameter must be a struct or pointer to a struct. `tfsdk` must be a pointer to the corresponding
// TF SDK structure.
//
// Structs in Go SDK are represented as types.Lists.
// If field name doesn't show up in ForceSendFields and the field is zero value, we set the null value on the tfsdk.
// Map keys must always be strings.
// TF SDK structs use types.String for all enum values.
// Non-JSON fields will be omitted.
func GoSdkToTfSdkStruct(ctx context.Context, gosdk interface{}, tfsdk interface{}) (d diag.Diagnostics) {
	srcVal := reflect.ValueOf(gosdk)
	destVal := reflect.ValueOf(tfsdk)

	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}

	if destVal.Kind() != reflect.Ptr {
		d.AddError(goSdkToTfSdkStructConversionFailureMessage, fmt.Sprintf("please provide a pointer for the tfsdk struct, got %s", destVal.Type().Name()))
		return
	}
	destVal = destVal.Elem()

	if srcVal.Kind() != reflect.Struct || destVal.Kind() != reflect.Struct {
		d.AddError(goSdkToTfSdkStructConversionFailureMessage, fmt.Sprintf("input should be structs, got %s, %s", srcVal.Type().Name(), destVal.Type().Name()))
		return
	}

	// complexFieldTypes captures the elements within a types.List, types.Object, or types.Map.
	var complexFieldTypes map[string]reflect.Type
	if cftp, ok := destVal.Interface().(tfcommon.ComplexFieldTypeProvider); ok {
		complexFieldTypes = cftp.GetComplexFieldTypes(ctx)
	}

	// objectType is the type of the destination struct. Entries from this are used when constructing
	// plugin framework attr.Values for fields in the object.
	objectType := tfcommon.NewObjectValuable(tfsdk).Type(ctx).(types.ObjectType)

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
		destField := destVal.FieldByName(toTfSdkName(srcFieldName))
		destFieldType, ok := destVal.Type().FieldByName(field.StructField.Name)
		if !ok {
			d.AddError(goSdkToTfSdkStructConversionFailureMessage, fmt.Sprintf("destination struct does not have field %s. %s", srcFieldName, common.TerraformBugErrorMessage))
			return
		}
		destFieldName := destFieldType.Tag.Get("tfsdk")
		if destFieldName == "-" {
			continue
		}

		if !destField.IsValid() {
			logger.Tracef(ctx, fmt.Sprintf("field skipped in gosdk to tfsdk conversion: destination struct does not have field %s", srcFieldName))
			continue
		}
		innerType := objectType.AttrTypes[destFieldName]
		complexFieldType := complexFieldTypes[destFieldName]

		d.Append(goSdkToTfSdkSingleField(ctx, srcField, destField, fieldInForceSendFields(srcFieldName, forceSendFieldVal), innerType, complexFieldType)...)
		if d.HasError() {
			return
		}
	}
	return
}

// goSdkToTfSdkSingleField converts a single field from a Go SDK struct to a TF SDK struct.
// The `srcField` is the field in the Go SDK struct, and `destField` is the field on which
// the value will be set in the TF SDK struct. Note that unlike GoSdkToTfSdkStruct, the
// `destField` parameter is not a pointer to the field, but the field itself. The `tfType`
// parameter is the Terraform type of the field, and `complexFieldType` is the runtime
// type of the field. These parameters are only needed when the field is a list, object, or
// map.
func goSdkToTfSdkSingleField(
	ctx context.Context,
	srcField reflect.Value,
	destField reflect.Value,
	forceSendField bool,
	tfType attr.Type,
	innerType reflect.Type) (d diag.Diagnostics) {
	if !destField.CanSet() {
		d.AddError(goSdkToTfSdkStructConversionFailureMessage, fmt.Sprintf("destination field can not be set: %s. %s", destField.Type().Name(), common.TerraformBugErrorMessage))
		return d
	}

	switch srcField.Kind() {
	case reflect.Ptr:
		// This corresponds to either a types.List or types.Object.
		// If nil, set the destination field to the null value of the appropriate type.
		if srcField.IsNil() {
			setFieldToNull(destField, tfType)
			return nil
		}

		// Otherwise, dereference the pointer and continue.
		srcField = srcField.Elem()
		d.Append(goSdkToTfSdkSingleField(ctx, srcField, destField, forceSendField, tfType, innerType)...)
	case reflect.Bool:
		boolVal := srcField.Interface().(bool)
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
			var ds diag.Diagnostics
			strVal, ds = getStringFromEnum(srcField)
			d.Append(ds...)
			if d.HasError() {
				return
			}
		} else {
			strVal = srcField.Interface().(string)
		}
		// check if the value is non-zero or if the field is in the forceSendFields list
		if strVal != "" || forceSendField {
			destField.Set(reflect.ValueOf(types.StringValue(strVal)))
		} else {
			destField.Set(reflect.ValueOf(types.StringNull()))
		}
	case reflect.Struct:
		// This corresponds to either a types.List or types.Object.
		// If the struct is zero value, set the destination field to the null value of the appropriate type.
		if srcField.IsZero() {
			setFieldToNull(destField, tfType)
			return
		}

		// If the destination field is a types.List, treat the source field as a slice with length 1
		// containing only this struct.
		if destField.Type() == reflect.TypeOf(types.List{}) {
			listSrc := reflect.MakeSlice(reflect.SliceOf(srcField.Type()), 1, 1)
			listSrc.Index(0).Set(srcField)
			return goSdkToTfSdkSingleField(ctx, listSrc, destField, forceSendField, tfType, innerType)
		}

		// Otherwise, the destination field is a types.Object. Convert the nested struct to the corresponding
		// TFSDK struct, then set the destination field to the object
		dest := reflect.New(innerType).Interface()
		d.Append(GoSdkToTfSdkStruct(ctx, srcField.Interface(), dest)...)
		if d.HasError() {
			return
		}
		objectType, ok := tfType.(types.ObjectType)
		if !ok {
			d.AddError(goSdkToTfSdkFieldConversionFailureMessage, fmt.Sprintf("inner type is not an object type: %s. %s", tfType, common.TerraformBugErrorMessage))
			return
		}
		objectVal, ds := types.ObjectValueFrom(ctx, objectType.AttrTypes, dest)
		d.Append(ds...)
		if d.HasError() {
			return
		}
		destField.Set(reflect.ValueOf(objectVal))
	case reflect.Slice:
		// This always corresponds to a types.List.
		listType, ok := tfType.(types.ListType)
		if !ok {
			d.AddError(goSdkToTfSdkFieldConversionFailureMessage, fmt.Sprintf("inner type is not a list type: %s. %s", tfType, common.TerraformBugErrorMessage))
			return
		}
		if srcField.Len() == 0 {
			// If the destination field is a types.List, treat the source field as an empty slice.
			emptyList := types.ListNull(listType.ElemType)
			destField.Set(reflect.ValueOf(emptyList))
			return
		}

		// Convert each element of the slice to the corresponding inner type.
		elements := make([]any, 0, srcField.Len())
		for i := 0; i < srcField.Len(); i++ {
			element := reflect.New(innerType)
			// If the element is a primitive type, we can convert it by recursively calling this function.
			// Otherwise, it is a struct, and we need to convert it by calling GoSdkToTfSdkStruct.
			switch innerType {
			case reflect.TypeOf(types.String{}), reflect.TypeOf(types.Bool{}), reflect.TypeOf(types.Int64{}), reflect.TypeOf(types.Float64{}):
				d.Append(goSdkToTfSdkSingleField(ctx, srcField.Index(i), element.Elem(), true, listType.ElemType, innerType)...)
			default:
				d.Append(GoSdkToTfSdkStruct(ctx, srcField.Index(i).Interface(), element.Interface())...)
			}
			if d.HasError() {
				return
			}
			elements = append(elements, element.Interface())
		}

		// Construct the Terraform value and set it.
		destVal, ds := types.ListValueFrom(ctx, listType.ElemType, elements)
		d.Append(ds...)
		if d.HasError() {
			return
		}
		destField.Set(reflect.ValueOf(destVal))
	case reflect.Map:
		// This always corresponds to a types.Map.
		mapType, ok := tfType.(types.MapType)
		if !ok {
			d.AddError(goSdkToTfSdkFieldConversionFailureMessage, fmt.Sprintf("inner type is not a map type: %s. %s", tfType, common.TerraformBugErrorMessage))
			return
		}
		if srcField.Len() == 0 {
			// If the destination field is a types.Map, treat the source field as an empty map.
			emptyMap := types.MapNull(mapType.ElemType)
			destField.Set(reflect.ValueOf(emptyMap))
			return
		}

		// Convert each key-value pair of the map to the corresponding inner type.
		destMap := map[string]any{}
		for _, key := range srcField.MapKeys() {
			srcMapValue := srcField.MapIndex(key)
			destMapValue := reflect.New(innerType)
			// If the element is a primitive type, we can convert it by recursively calling this function.
			// Otherwise, it is a struct, and we need to convert it by calling GoSdkToTfSdkStruct.
			switch innerType {
			case reflect.TypeOf(types.String{}), reflect.TypeOf(types.Bool{}), reflect.TypeOf(types.Int64{}), reflect.TypeOf(types.Float64{}):
				d.Append(goSdkToTfSdkSingleField(ctx, srcMapValue, destMapValue.Elem(), true, mapType.ElemType, innerType)...)
			default:
				d.Append(GoSdkToTfSdkStruct(ctx, srcMapValue.Interface(), destMapValue.Interface())...)
			}
			if d.HasError() {
				return
			}
			destMap[key.String()] = destMapValue.Interface()
		}

		// Construct the Terraform value and set it.
		destVal, ds := types.MapValueFrom(ctx, mapType.ElemType, destMap)
		d.Append(ds...)
		if d.HasError() {
			return
		}
		destField.Set(reflect.ValueOf(destVal))
	default:
		d.AddError(goSdkToTfSdkFieldConversionFailureMessage, fmt.Sprintf("%s is not currently supported as a source field. %s", srcField.Type().Name(), common.TerraformBugErrorMessage))
	}
	return
}

// setFieldToNull sets the destination field to the null value of the appropriate type.
func setFieldToNull(destField reflect.Value, innerType attr.Type) {
	switch destField.Type() {
	case reflect.TypeOf(types.List{}):
		// If the destination field is a types.List, treat the source field as an empty slice.
		listType := innerType.(types.ListType)
		emptyList := types.ListNull(listType.ElemType)
		destField.Set(reflect.ValueOf(emptyList))
	case reflect.TypeOf(types.Object{}):
		// If the destination field is a types.Object, treat the source field as an empty object.
		innerType := innerType.(types.ObjectType)
		destField.Set(reflect.ValueOf(types.ObjectNull(innerType.AttrTypes)))
	}
}

// Get the string value of an enum by calling the .String() method on the enum object.
func getStringFromEnum(srcField reflect.Value) (s string, d diag.Diagnostics) {
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
			s = stringResult[0].Interface().(string)
			return
		}
		d.AddError(goSdkToTfSdkFieldConversionFailureMessage, fmt.Sprintf("num get string has more than one result. %s", common.TerraformBugErrorMessage))
		return
	}
	d.AddError(goSdkToTfSdkFieldConversionFailureMessage, fmt.Sprintf("enum does not have valid .String() method. %s", common.TerraformBugErrorMessage))
	return
}

func fieldInForceSendFields(fieldName string, forceSendFields []string) bool {
	for _, field := range forceSendFields {
		if field == fieldName {
			return true
		}
	}
	return false
}
