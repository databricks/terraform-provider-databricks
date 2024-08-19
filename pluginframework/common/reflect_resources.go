package pluginframework

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/terraform-provider-databricks/common"
	dataschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func pluginFrameworkTypeToSchema(v reflect.Value) map[string]Attribute {
	scm := map[string]Attribute{}
	rk := v.Kind()
	if rk == reflect.Ptr {
		v = v.Elem()
		rk = v.Kind()
	}
	if rk != reflect.Struct {
		panic(fmt.Errorf("Schema value of Struct is expected, but got %s: %#v", rk.String(), v))
	}
	fields := common.ListAllFields(v)
	for _, field := range fields {
		typeField := field.Sf
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
			nestedScm := pluginFrameworkTypeToSchema(sv)
			scm[fieldName] = SingleNestedAttribute{Attributes: nestedScm, Optional: isOptional, Required: !isOptional}
		} else if kind == reflect.Slice {
			elem := typeField.Type.Elem()
			if elem.Kind() == reflect.Ptr {
				elem = elem.Elem()
			}
			if elem.Kind() != reflect.Struct {
				panic(fmt.Errorf("unsupported slice value for %s: %s", fieldName, elem.Kind().String()))
			}
			switch elem {
			case reflect.TypeOf(types.Bool{}):
				scm[fieldName] = ListAttribute{ElementType: types.BoolType, Optional: isOptional, Required: !isOptional}
			case reflect.TypeOf(types.Int64{}):
				scm[fieldName] = ListAttribute{ElementType: types.Int64Type, Optional: isOptional, Required: !isOptional}
			case reflect.TypeOf(types.Float64{}):
				scm[fieldName] = ListAttribute{ElementType: types.Float64Type, Optional: isOptional, Required: !isOptional}
			case reflect.TypeOf(types.String{}):
				scm[fieldName] = ListAttribute{ElementType: types.StringType, Optional: isOptional, Required: !isOptional}
			default:
				// Nested struct
				nestedScm := pluginFrameworkTypeToSchema(reflect.New(elem).Elem())
				scm[fieldName] = ListNestedAttribute{NestedObject: NestedAttributeObject{Attributes: nestedScm}, Optional: isOptional, Required: !isOptional}
			}
		} else if kind == reflect.Map {
			elem := typeField.Type.Elem()
			if elem.Kind() == reflect.Ptr {
				elem = elem.Elem()
			}
			if elem.Kind() != reflect.Struct {
				panic(fmt.Errorf("unsupported map value for %s: %s", fieldName, elem.Kind().String()))
			}
			switch elem {
			case reflect.TypeOf(types.Bool{}):
				scm[fieldName] = MapAttribute{ElementType: types.BoolType, Optional: isOptional, Required: !isOptional}
			case reflect.TypeOf(types.Int64{}):
				scm[fieldName] = MapAttribute{ElementType: types.Int64Type, Optional: isOptional, Required: !isOptional}
			case reflect.TypeOf(types.Float64{}):
				scm[fieldName] = MapAttribute{ElementType: types.Float64Type, Optional: isOptional, Required: !isOptional}
			case reflect.TypeOf(types.String{}):
				scm[fieldName] = MapAttribute{ElementType: types.StringType, Optional: isOptional, Required: !isOptional}
			default:
				// Nested struct
				nestedScm := pluginFrameworkTypeToSchema(reflect.New(elem).Elem())
				scm[fieldName] = MapNestedAttribute{NestedObject: NestedAttributeObject{Attributes: nestedScm}, Optional: isOptional, Required: !isOptional}
			}
		} else if kind == reflect.Struct {
			switch field.V.Interface().(type) {
			case types.Bool:
				scm[fieldName] = BoolAttribute{Optional: isOptional, Required: !isOptional}
			case types.Int64:
				scm[fieldName] = Int64Attribute{Optional: isOptional, Required: !isOptional}
			case types.Float64:
				scm[fieldName] = Float64Attribute{Optional: isOptional, Required: !isOptional}
			case types.String:
				scm[fieldName] = StringAttribute{Optional: isOptional, Required: !isOptional}
			case types.List:
				panic("types.List should never be used in tfsdk structs")
			case types.Map:
				panic("types.Map should never be used in tfsdk structs")
			default:
				// If it is a real stuct instead of a tfsdk type, recursively resolve it.
				elem := typeField.Type
				sv := reflect.New(elem)
				nestedScm := pluginFrameworkTypeToSchema(sv)
				scm[fieldName] = SingleNestedAttribute{Attributes: nestedScm, Optional: isOptional, Required: !isOptional}
			}
		} else if kind == reflect.String {
			// This case is for Enum types.
			scm[fieldName] = StringAttribute{Optional: isOptional, Required: !isOptional}
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

func PluginFrameworkDataSourceStructToSchema(v any, customizeSchema func(CustomizableSchemaPluginFramework) CustomizableSchemaPluginFramework) dataschema.Schema {
	attributes := PluginFrameworkDataSourceStructToSchemaMap(v, customizeSchema)
	return dataschema.Schema{Attributes: attributes}
}

func PluginFrameworkResourceStructToSchemaMap(v any, customizeSchema func(CustomizableSchemaPluginFramework) CustomizableSchemaPluginFramework) map[string]schema.Attribute {
	attributes := pluginFrameworkTypeToSchema(reflect.ValueOf(v))

	if customizeSchema != nil {
		cs := customizeSchema(*ConstructCustomizableSchema(attributes))
		return ToResourceAttributeMap(cs.ToAttributeMap())
	} else {
		return ToResourceAttributeMap(attributes)
	}
}

func PluginFrameworkDataSourceStructToSchemaMap(v any, customizeSchema func(CustomizableSchemaPluginFramework) CustomizableSchemaPluginFramework) map[string]dataschema.Attribute {
	attributes := pluginFrameworkTypeToSchema(reflect.ValueOf(v))

	if customizeSchema != nil {
		cs := customizeSchema(*ConstructCustomizableSchema(attributes))
		return ToDataSourceAttributeMap(cs.ToAttributeMap())
	} else {
		return ToDataSourceAttributeMap(attributes)
	}
}
