package tfschema

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/tfreflect"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	dataschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type structTag struct {
	optional     bool
	computed     bool
	singleObject bool
}

func typeToSchema(v reflect.Value) NestedBlockObject {
	scmAttr := map[string]AttributeBuilder{}
	scmBlock := map[string]BlockBuilder{}
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
		structTag := getStructTag(typeField)
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
				scmAttr[fieldName] = ListAttributeBuilder{
					ElementType: types.BoolType,
					Optional:    structTag.optional,
					Required:    !structTag.optional,
					Computed:    structTag.computed,
				}
			case reflect.TypeOf(types.Int64{}):
				scmAttr[fieldName] = ListAttributeBuilder{
					ElementType: types.Int64Type,
					Optional:    structTag.optional,
					Required:    !structTag.optional,
					Computed:    structTag.computed,
				}
			case reflect.TypeOf(types.Float64{}):
				scmAttr[fieldName] = ListAttributeBuilder{
					ElementType: types.Float64Type,
					Optional:    structTag.optional,
					Required:    !structTag.optional,
					Computed:    structTag.computed,
				}
			case reflect.TypeOf(types.String{}):
				scmAttr[fieldName] = ListAttributeBuilder{
					ElementType: types.StringType,
					Optional:    structTag.optional,
					Required:    !structTag.optional,
					Computed:    structTag.computed,
				}
			default:
				// Nested struct
				nestedScm := typeToSchema(reflect.New(elemType).Elem())
				var validators []validator.List
				if structTag.singleObject {
					validators = append(validators, listvalidator.SizeAtMost(1))
				}
				scmBlock[fieldName] = ListNestedBlockBuilder{
					NestedObject: &NestedBlockObject{
						Attributes: nestedScm.Attributes,
						Blocks:     nestedScm.Blocks,
					},
					Optional:   structTag.optional,
					Required:   !structTag.optional,
					Computed:   structTag.computed,
					Validators: validators,
				}
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
				scmAttr[fieldName] = MapAttributeBuilder{
					ElementType: types.BoolType,
					Optional:    structTag.optional,
					Required:    !structTag.optional,
					Computed:    structTag.computed,
				}
			case reflect.TypeOf(types.Int64{}):
				scmAttr[fieldName] = MapAttributeBuilder{
					ElementType: types.Int64Type,
					Optional:    structTag.optional,
					Required:    !structTag.optional,
					Computed:    structTag.computed,
				}
			case reflect.TypeOf(types.Float64{}):
				scmAttr[fieldName] = MapAttributeBuilder{
					ElementType: types.Float64Type,
					Optional:    structTag.optional,
					Required:    !structTag.optional,
					Computed:    structTag.computed,
				}
			case reflect.TypeOf(types.String{}):
				scmAttr[fieldName] = MapAttributeBuilder{
					ElementType: types.StringType,
					Optional:    structTag.optional,
					Required:    !structTag.optional,
					Computed:    structTag.computed,
				}
			default:
				// Nested struct
				nestedScm := typeToSchema(reflect.New(elemType).Elem())
				scmAttr[fieldName] = MapNestedAttributeBuilder{
					NestedObject: NestedAttributeObject{
						Attributes: nestedScm.Attributes,
					},
					Optional: structTag.optional,
					Required: !structTag.optional,
					Computed: structTag.computed,
				}
			}
		} else if kind == reflect.Struct {
			switch value.Interface().(type) {
			case types.Bool:
				scmAttr[fieldName] = BoolAttributeBuilder{
					Optional: structTag.optional,
					Required: !structTag.optional,
					Computed: structTag.computed,
				}
			case types.Int64:
				scmAttr[fieldName] = Int64AttributeBuilder{
					Optional: structTag.optional,
					Required: !structTag.optional,
					Computed: structTag.computed,
				}
			case types.Float64:
				scmAttr[fieldName] = Float64AttributeBuilder{
					Optional: structTag.optional,
					Required: !structTag.optional,
					Computed: structTag.computed,
				}
			case types.String:
				scmAttr[fieldName] = StringAttributeBuilder{
					Optional: structTag.optional,
					Required: !structTag.optional,
					Computed: structTag.computed,
				}
			case types.List:
				panic(fmt.Errorf("types.List should never be used in tfsdk structs. %s", common.TerraformBugErrorMessage))
			case types.Map:
				panic(fmt.Errorf("types.Map should never be used in tfsdk structs. %s", common.TerraformBugErrorMessage))
			default:
				// If it is a real stuct instead of a tfsdk type, recursively resolve it.
				elem := typeFieldType
				sv := reflect.New(elem)
				nestedScm := typeToSchema(sv)
				scmBlock[fieldName] = ListNestedBlockBuilder{
					NestedObject: &nestedScm,
					Optional:     structTag.optional,
					Required:     !structTag.optional,
					Computed:     structTag.computed,
				}
			}
		} else {
			panic(fmt.Errorf("unknown type for field: %s. %s", typeField.Name, common.TerraformBugErrorMessage))
		}
	}
	return NestedBlockObject{Attributes: scmAttr, Blocks: scmBlock}
}

func getStructTag(field reflect.StructField) structTag {
	tagValue := field.Tag.Get("tf")
	return structTag{
		optional:     strings.Contains(tagValue, "optional"),
		computed:     strings.Contains(tagValue, "computed"),
		singleObject: strings.Contains(tagValue, "object"),
	}
}

// ResourceStructToSchema builds a resource schema from a tfsdk struct, with custoimzations applied.
func ResourceStructToSchema(v any, customizeSchema func(CustomizableSchema) CustomizableSchema) schema.Schema {
	attributes, blocks := ResourceStructToSchemaMap(v, customizeSchema)
	return schema.Schema{Attributes: attributes, Blocks: blocks}
}

// DataSourceStructToSchema builds a data source schema from a tfsdk struct, with custoimzations applied.
func DataSourceStructToSchema(v any, customizeSchema func(CustomizableSchema) CustomizableSchema) dataschema.Schema {
	attributes, blocks := DataSourceStructToSchemaMap(v, customizeSchema)
	return dataschema.Schema{Attributes: attributes, Blocks: blocks}
}

// ResourceStructToSchemaMap returns two maps from string to resource schema attributes and blocks using a tfsdk struct, with custoimzations applied.
func ResourceStructToSchemaMap(v any, customizeSchema func(CustomizableSchema) CustomizableSchema) (map[string]schema.Attribute, map[string]schema.Block) {
	nestedBlockObj := typeToSchema(reflect.ValueOf(v))

	if customizeSchema != nil {
		cs := customizeSchema(*ConstructCustomizableSchema(nestedBlockObj))
		return BuildResourceAttributeMap(cs.ToNestedBlockObject().Attributes), BuildResourceBlockMap(cs.ToNestedBlockObject().Blocks)
	} else {
		return BuildResourceAttributeMap(nestedBlockObj.Attributes), BuildResourceBlockMap(nestedBlockObj.Blocks)
	}
}

// DataSourceStructToSchemaMap returns twp maps from string to data source schema attributes and blocks using a tfsdk struct, with custoimzations applied.
func DataSourceStructToSchemaMap(v any, customizeSchema func(CustomizableSchema) CustomizableSchema) (map[string]dataschema.Attribute, map[string]dataschema.Block) {
	nestedBlockObj := typeToSchema(reflect.ValueOf(v))

	if customizeSchema != nil {
		cs := customizeSchema(*ConstructCustomizableSchema(nestedBlockObj))
		return BuildDataSourceAttributeMap(cs.ToNestedBlockObject().Attributes), BuildDataSourceBlockMap(cs.ToNestedBlockObject().Blocks)
	} else {
		return BuildDataSourceAttributeMap(nestedBlockObj.Attributes), BuildDataSourceBlockMap(nestedBlockObj.Blocks)
	}
}
