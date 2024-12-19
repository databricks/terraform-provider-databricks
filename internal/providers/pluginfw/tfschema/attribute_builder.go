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

	// SetOptional sets the attribute as optional in the schema. This does not affect whether the attribute is computed.
	// It fails if the attribute is already optional.
	SetOptional() AttributeBuilder

	// SetRequired sets the attribute as required in the schema. This does not affect whether the attribute is computed.
	// It fails if the attribute is already required.
	SetRequired() AttributeBuilder

	// SetSensitive sets the attribute as sensitive in the schema. It fails if the attribute is already sensitive.
	SetSensitive() AttributeBuilder

	// SetComputed sets the attribute as computed in the schema. It fails if the attribute is already computed.
	SetComputed() AttributeBuilder

	// Sets the attribute as read-only in the schema, i.e. computed and neither optional or required. It fails if the
	// attribute is already read-only.
	SetReadOnly() AttributeBuilder

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
