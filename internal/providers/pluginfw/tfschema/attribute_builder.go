package tfschema

import (
	dataschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// AttributeBuilder is the common interface for all attributes, it can be used to build data source attribute and resource attribute.
// We need this because in terraform plugin framework, the datasource schema and resource schema are in two separate packages.
// This common interface prevents us from keeping two copies of StructToSchema and CustomizableSchema.
type AttributeBuilder interface {
	BaseSchemaBuilder
	BuildDataSourceAttribute() dataschema.Attribute
	BuildResourceAttribute() schema.Attribute
}

// BuildDataSourceAttributeMap takes a map from string to AttributeBuilder and returns a map from string to datasource.schema.Attribute.
func BuildDataSourceAttributeMap(attributes map[string]AttributeBuilder) map[string]dataschema.Attribute {
	dataSourceAttributes := make(map[string]dataschema.Attribute)

	for key, attribute := range attributes {
		dataSourceAttributes[key] = attribute.BuildDataSourceAttribute()
	}

	return dataSourceAttributes
}

// BuildResourceAttributeMap takes a map from string to AttributeBuilder and returns a map from string to resource.schema.Attribute.
func BuildResourceAttributeMap(attributes map[string]AttributeBuilder) map[string]schema.Attribute {
	resourceAttributes := make(map[string]schema.Attribute)

	for key, attribute := range attributes {
		resourceAttributes[key] = attribute.BuildResourceAttribute()
	}

	return resourceAttributes
}
