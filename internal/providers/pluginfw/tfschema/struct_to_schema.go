package tfschema

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/terraform-provider-databricks/common"
	tfcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/tfreflect"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
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

func typeToSchema(ctx context.Context, v reflect.Value) NestedBlockObject {
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

	for _, field := range tfreflect.ListAllFields(v) {
		typeField := field.StructField
		fieldName := typeField.Tag.Get("tfsdk")
		if fieldName == "-" {
			continue
		}
		structTag := getStructTag(typeField)
		value := field.Value.Interface()
		if _, ok := value.(attr.Value); !ok {
			panic(fmt.Errorf("unexpected type %T in tfsdk structs, expected a plugin framework value type. %s", value, common.TerraformBugErrorMessage))
		}
		switch value.(type) {
		case types.Bool:
			scmAttr[fieldName] = BoolAttributeBuilder{}
		case types.Int64:
			scmAttr[fieldName] = Int64AttributeBuilder{}
		case types.Float64:
			scmAttr[fieldName] = Float64AttributeBuilder{}
		case types.String:
			scmAttr[fieldName] = StringAttributeBuilder{}
		case types.List, types.Map, types.Object:
			// Additional metadata is required to determine the type of the list elements.
			// This is available via the ComplexFieldTypeProvider interface, implemented on the parent type.
			provider, ok := v.Interface().(tfcommon.ComplexFieldTypeProvider)
			if !ok {
				panic(fmt.Errorf("complex field types not provided for type: %T. %s", v.Interface(), common.TerraformBugErrorMessage))
			}
			complexFieldTypes := provider.GetComplexFieldTypes(ctx)
			fieldType, ok := complexFieldTypes[fieldName]
			if !ok {
				panic(fmt.Errorf("complex field type not found for field %s on type %T. %s", typeField.Name, v.Interface(), common.TerraformBugErrorMessage))
			}
			// If the field type is a "primitive", use the appropriate AttributeBuilder. This includes enums, which are treated as strings.
			// Otherwise, use ListNestedBlockBuilder.
			switch fieldType {
			// Note: The list of primitive types must match all of the possible types generated by the `attr-type` template in .codegen/model.go.tmpl.
			// If new types are added there, they must also be added here to work properly.
			case reflect.TypeOf(types.Bool{}), reflect.TypeOf(types.Int64{}), reflect.TypeOf(types.Float64{}), reflect.TypeOf(types.String{}):
				// Look up the element type from the Type() methods for TF SDK structs.
				// This is always a attr.TypeWithElementType because the field is a list or map.
				objectType := tfcommon.NewObjectTyper(v.Interface()).Type(ctx).(types.ObjectType)
				attrType, ok := objectType.AttrTypes[fieldName]
				if !ok {
					panic(fmt.Errorf("attr type not found for field %s on type %T. %s", typeField.Name, v.Interface(), common.TerraformBugErrorMessage))
				}
				containerType := attrType.(attr.TypeWithElementType)
				switch value.(type) {
				case types.List:
					scmAttr[fieldName] = ListAttributeBuilder{ElementType: containerType.ElementType()}
				case types.Map:
					scmAttr[fieldName] = MapAttributeBuilder{ElementType: containerType.ElementType()}
				}
			default:
				// The element type is a TFSDK type. Map fields are treated as MapNestedAttributes. For compatibility,
				// list fields are treated as ListNestedBlocks.
				// TODO: Change the default for lists to ListNestedAttribute.
				fieldValue := reflect.New(fieldType).Elem()

				// Generate the nested block schema
				nestedSchema := typeToSchema(ctx, fieldValue)
				switch value.(type) {
				case types.List:
					validators := []validator.List{}
					if structTag.singleObject {
						validators = append(validators, listvalidator.SizeAtMost(1))
					}
					// Note that this is being added to the block map, not the attribute map.
					scmBlock[fieldName] = ListNestedBlockBuilder{
						NestedObject: nestedSchema,
						Validators:   validators,
					}
				case types.Map:
					scmAttr[fieldName] = MapNestedAttributeBuilder{
						NestedObject: nestedSchema.ToNestedAttributeObject(),
					}
				case types.Object:
					scmAttr[fieldName] = SingleNestedAttributeBuilder{
						Attributes: nestedSchema.ToNestedAttributeObject().Attributes,
					}
				}
			}
		default:
			panic(fmt.Errorf("unexpected type %T in tfsdk structs, expected a plugin framework value type. %s", value, common.TerraformBugErrorMessage))
		}
		// types.List fields of complex types correspond to ListNestedBlock, which don't have optional/required/computed flags.
		// When these fields are later changed to use ListNestedAttribute, we can inline the if statement below, as all fields
		// will be attributes.
		if attr, ok := scmAttr[fieldName]; ok {
			if structTag.optional {
				attr = attr.SetOptional()
			} else {
				attr = attr.SetRequired()
			}
			if structTag.computed {
				attr = attr.SetComputed()
			}
			scmAttr[fieldName] = attr
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
func ResourceStructToSchema(ctx context.Context, v any, customizeSchema func(CustomizableSchema) CustomizableSchema) schema.Schema {
	attributes, blocks := ResourceStructToSchemaMap(ctx, v, customizeSchema)
	return schema.Schema{Attributes: attributes, Blocks: blocks}
}

// DataSourceStructToSchema builds a data source schema from a tfsdk struct, with custoimzations applied.
func DataSourceStructToSchema(ctx context.Context, v any, customizeSchema func(CustomizableSchema) CustomizableSchema) dataschema.Schema {
	attributes, blocks := DataSourceStructToSchemaMap(ctx, v, customizeSchema)
	return dataschema.Schema{Attributes: attributes, Blocks: blocks}
}

// ResourceStructToSchemaMap returns two maps from string to resource schema attributes and blocks using a tfsdk struct, with custoimzations applied.
func ResourceStructToSchemaMap(ctx context.Context, v any, customizeSchema func(CustomizableSchema) CustomizableSchema) (map[string]schema.Attribute, map[string]schema.Block) {
	nestedBlockObj := typeToSchema(ctx, reflect.ValueOf(v))

	if customizeSchema != nil {
		cs := customizeSchema(*ConstructCustomizableSchema(nestedBlockObj))
		return BuildResourceAttributeMap(cs.ToNestedBlockObject().Attributes), BuildResourceBlockMap(cs.ToNestedBlockObject().Blocks)
	} else {
		return BuildResourceAttributeMap(nestedBlockObj.Attributes), BuildResourceBlockMap(nestedBlockObj.Blocks)
	}
}

// DataSourceStructToSchemaMap returns twp maps from string to data source schema attributes and blocks using a tfsdk struct, with custoimzations applied.
func DataSourceStructToSchemaMap(ctx context.Context, v any, customizeSchema func(CustomizableSchema) CustomizableSchema) (map[string]dataschema.Attribute, map[string]dataschema.Block) {
	nestedBlockObj := typeToSchema(ctx, reflect.ValueOf(v))

	if customizeSchema != nil {
		cs := customizeSchema(*ConstructCustomizableSchema(nestedBlockObj))
		return BuildDataSourceAttributeMap(cs.ToNestedBlockObject().Attributes), BuildDataSourceBlockMap(cs.ToNestedBlockObject().Blocks)
	} else {
		return BuildDataSourceAttributeMap(nestedBlockObj.Attributes), BuildDataSourceBlockMap(nestedBlockObj.Blocks)
	}
}
