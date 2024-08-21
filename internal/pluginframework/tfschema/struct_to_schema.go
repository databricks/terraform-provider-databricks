package tfschema

import (
	"fmt"
	"reflect"
	"strings"

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
		panic(fmt.Errorf("schema value of Struct is expected, but got %s: %#v", rk.String(), v))
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
		if kind == reflect.Ptr {
			elem := typeField.Type.Elem()
			sv := reflect.New(elem).Elem()
			nestedScm := typeToSchema(sv)
			scm[fieldName] = SingleNestedAttributeBuilder{Attributes: nestedScm, Optional: isOptional, Required: !isOptional}
		} else if kind == reflect.Slice {
			elemType := typeField.Type.Elem()
			if elemType.Kind() == reflect.Ptr {
				elemType = elemType.Elem()
			}
			if elemType.Kind() != reflect.Struct {
				panic(fmt.Errorf("unsupported slice value for %s: %s", fieldName, elemType.Kind().String()))
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
			elemType := typeField.Type.Elem()
			if elemType.Kind() == reflect.Ptr {
				elemType = elemType.Elem()
			}
			if elemType.Kind() != reflect.Struct {
				panic(fmt.Errorf("unsupported map value for %s: %s", fieldName, elemType.Kind().String()))
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
			switch field.Value.Interface().(type) {
			case types.Bool:
				scm[fieldName] = BoolAttributeBuilder{Optional: isOptional, Required: !isOptional}
			case types.Int64:
				scm[fieldName] = Int64AttributeBuilder{Optional: isOptional, Required: !isOptional}
			case types.Float64:
				scm[fieldName] = Float64AttributeBuilder{Optional: isOptional, Required: !isOptional}
			case types.String:
				scm[fieldName] = StringAttributeBuilder{Optional: isOptional, Required: !isOptional}
			case types.List:
				panic("types.List should never be used in tfsdk structs")
			case types.Map:
				panic("types.Map should never be used in tfsdk structs")
			default:
				// If it is a real stuct instead of a tfsdk type, recursively resolve it.
				elem := typeField.Type
				sv := reflect.New(elem)
				nestedScm := typeToSchema(sv)
				scm[fieldName] = SingleNestedAttributeBuilder{Attributes: nestedScm, Optional: isOptional, Required: !isOptional}
			}
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

func ResourceStructToSchema(v any, customizeSchema func(CustomizableSchema) CustomizableSchema) schema.Schema {
	attributes := ResourceStructToSchemaMap(v, customizeSchema)
	return schema.Schema{Attributes: attributes}
}

func DataSourceStructToSchema(v any, customizeSchema func(CustomizableSchema) CustomizableSchema) dataschema.Schema {
	attributes := DataSourceStructToSchemaMap(v, customizeSchema)
	return dataschema.Schema{Attributes: attributes}
}

func ResourceStructToSchemaMap(v any, customizeSchema func(CustomizableSchema) CustomizableSchema) map[string]schema.Attribute {
	attributes := typeToSchema(reflect.ValueOf(v))

	if customizeSchema != nil {
		cs := customizeSchema(*ConstructCustomizableSchema(attributes))
		return BuildResourceAttributeMap(cs.ToAttributeMap())
	} else {
		return BuildResourceAttributeMap(attributes)
	}
}

func DataSourceStructToSchemaMap(v any, customizeSchema func(CustomizableSchema) CustomizableSchema) map[string]dataschema.Attribute {
	attributes := typeToSchema(reflect.ValueOf(v))

	if customizeSchema != nil {
		cs := customizeSchema(*ConstructCustomizableSchema(attributes))
		return BuildDataSourceAttributeMap(cs.ToAttributeMap())
	} else {
		return BuildDataSourceAttributeMap(attributes)
	}
}
