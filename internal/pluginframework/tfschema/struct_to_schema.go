package tfschema

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/tfreflect"
	dataschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func typeToSchema(v reflect.Value) map[string]AttributeBuilder {
	scm := map[string]AttributeBuilder{}
	rk := v.Kind()
	if rk == reflect.Ptr {
		v = v.Elem()
		rk = v.Kind()
	}
	if rk != reflect.Struct {
		panic(fmt.Errorf("schema value of Struct is expected, but got %s: %#v. %s", rk.String(), v, common.TerraformBugErrorMessage))
	}
	fields := tfreflect.ListAllFields(v)
	for _, field := range fields {
		typeField := field.StructField
		fieldName := typeField.Tag.Get("tfsdk")
		if fieldName == "-" {
			continue
		}
		isOptional := fieldIsOptional(typeField)
		kind := typeField.Type.Kind()
		value := field.Value
		typeFieldType := typeField.Type
		if kind == reflect.Ptr {
			typeFieldType = typeFieldType.Elem()
			kind = typeFieldType.Kind()
			value = reflect.New(typeFieldType).Elem()
		}
		if kind == reflect.Slice {
			elemType := typeFieldType.Elem()
			if elemType.Kind() == reflect.Ptr {
				elemType = elemType.Elem()
			}
			if elemType.Kind() != reflect.Struct {
				panic(fmt.Errorf("unsupported slice value for %s: %s. %s", fieldName, elemType.Kind().String(), common.TerraformBugErrorMessage))
			}
			switch elemType {
			case reflect.TypeOf(types.Bool{}):
				scm[fieldName] = ListAttributeBuilder{ElementType: types.BoolType, Optional: isOptional, Required: !isOptional}
			case reflect.TypeOf(types.Int64{}):
				scm[fieldName] = ListAttributeBuilder{ElementType: types.Int64Type, Optional: isOptional, Required: !isOptional}
			case reflect.TypeOf(types.Float64{}):
				scm[fieldName] = ListAttributeBuilder{ElementType: types.Float64Type, Optional: isOptional, Required: !isOptional}
			case reflect.TypeOf(types.String{}):
				scm[fieldName] = ListAttributeBuilder{ElementType: types.StringType, Optional: isOptional, Required: !isOptional}
			default:
				// Nested struct
				nestedScm := typeToSchema(reflect.New(elemType).Elem())
				scm[fieldName] = ListNestedAttributeBuilder{NestedObject: NestedAttributeObject{Attributes: nestedScm}, Optional: isOptional, Required: !isOptional}
			}
		} else if kind == reflect.Map {
			elemType := typeFieldType.Elem()
			if elemType.Kind() == reflect.Ptr {
				elemType = elemType.Elem()
			}
			if elemType.Kind() != reflect.Struct {
				panic(fmt.Errorf("unsupported map value for %s: %s. %s", fieldName, elemType.Kind().String(), common.TerraformBugErrorMessage))
			}
			switch elemType {
			case reflect.TypeOf(types.Bool{}):
				scm[fieldName] = MapAttributeBuilder{ElementType: types.BoolType, Optional: isOptional, Required: !isOptional}
			case reflect.TypeOf(types.Int64{}):
				scm[fieldName] = MapAttributeBuilder{ElementType: types.Int64Type, Optional: isOptional, Required: !isOptional}
			case reflect.TypeOf(types.Float64{}):
				scm[fieldName] = MapAttributeBuilder{ElementType: types.Float64Type, Optional: isOptional, Required: !isOptional}
			case reflect.TypeOf(types.String{}):
				scm[fieldName] = MapAttributeBuilder{ElementType: types.StringType, Optional: isOptional, Required: !isOptional}
			default:
				// Nested struct
				nestedScm := typeToSchema(reflect.New(elemType).Elem())
				scm[fieldName] = MapNestedAttributeBuilder{NestedObject: NestedAttributeObject{Attributes: nestedScm}, Optional: isOptional, Required: !isOptional}
			}
		} else if kind == reflect.Struct {
			switch value.Interface().(type) {
			case types.Bool:
				scm[fieldName] = BoolAttributeBuilder{Optional: isOptional, Required: !isOptional}
			case types.Int64:
				scm[fieldName] = Int64AttributeBuilder{Optional: isOptional, Required: !isOptional}
			case types.Float64:
				scm[fieldName] = Float64AttributeBuilder{Optional: isOptional, Required: !isOptional}
			case types.String:
				scm[fieldName] = StringAttributeBuilder{Optional: isOptional, Required: !isOptional}
			case types.List:
				panic(fmt.Errorf("types.List should never be used in tfsdk structs. %s", common.TerraformBugErrorMessage))
			case types.Map:
				panic(fmt.Errorf("types.Map should never be used in tfsdk structs. %s", common.TerraformBugErrorMessage))
			default:
				// If it is a real stuct instead of a tfsdk type, recursively resolve it.
				elem := typeFieldType
				sv := reflect.New(elem)
				nestedScm := typeToSchema(sv)
				scm[fieldName] = SingleNestedAttributeBuilder{Attributes: nestedScm, Optional: isOptional, Required: !isOptional}
			}
		} else {
			panic(fmt.Errorf("unknown type for field: %s. %s", typeField.Name, common.TerraformBugErrorMessage))
		}
	}
	return scm
}

func fieldIsOptional(field reflect.StructField) bool {
	tagValue := field.Tag.Get("tf")
	return strings.Contains(tagValue, "optional")
}

// ResourceStructToSchema builds a resource schema from a tfsdk struct, with custoimzations applied.
func ResourceStructToSchema(v any, customizeSchema func(CustomizableSchema) CustomizableSchema) schema.Schema {
	attributes := ResourceStructToSchemaMap(v, customizeSchema)
	return schema.Schema{Attributes: attributes}
}

// DataSourceStructToSchema builds a data source schema from a tfsdk struct, with custoimzations applied.
func DataSourceStructToSchema(v any, customizeSchema func(CustomizableSchema) CustomizableSchema) dataschema.Schema {
	attributes := DataSourceStructToSchemaMap(v, customizeSchema)
	return dataschema.Schema{Attributes: attributes}
}

// ResourceStructToSchemaMap returns a map from string to resource schema attributes using a tfsdk struct, with custoimzations applied.
func ResourceStructToSchemaMap(v any, customizeSchema func(CustomizableSchema) CustomizableSchema) map[string]schema.Attribute {
	attributes := typeToSchema(reflect.ValueOf(v))

	if customizeSchema != nil {
		cs := customizeSchema(*ConstructCustomizableSchema(attributes))
		return BuildResourceAttributeMap(cs.ToAttributeMap())
	} else {
		return BuildResourceAttributeMap(attributes)
	}
}

// DataSourceStructToSchemaMap returns a map from string to data source schema attributes using a tfsdk struct, with custoimzations applied.
func DataSourceStructToSchemaMap(v any, customizeSchema func(CustomizableSchema) CustomizableSchema) map[string]dataschema.Attribute {
	attributes := typeToSchema(reflect.ValueOf(v))

	if customizeSchema != nil {
		cs := customizeSchema(*ConstructCustomizableSchema(attributes))
		return BuildDataSourceAttributeMap(cs.ToAttributeMap())
	} else {
		return BuildDataSourceAttributeMap(attributes)
	}
}
