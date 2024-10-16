package tfschema

import (
	dataschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// BlockBuilder is the common interface for all blocks, it can be used to build data source blocks and resource blocks.
// We need this because in terraform plugin framework, the datasource schema and resource schema are in two separate packages.
// This common interface prevents us from keeping two copies of StructToSchema and CustomizableSchema.
type BlockBuilder interface {
	BaseSchemaBuilder
	BuildDataSourceBlock() dataschema.Block
	BuildResourceBlock() schema.Block
}

func BuildDataSourceBlockMap(attributes map[string]BlockBuilder) map[string]dataschema.Block {
	dataSourceAttributes := make(map[string]dataschema.Block)

	for key, attribute := range attributes {
		dataSourceAttributes[key] = attribute.BuildDataSourceBlock()
	}

	return dataSourceAttributes
}

func BuildResourceBlockMap(attributes map[string]BlockBuilder) map[string]schema.Block {
	resourceAttributes := make(map[string]schema.Block)

	for key, attribute := range attributes {
		resourceAttributes[key] = attribute.BuildResourceBlock()
	}

	return resourceAttributes
}
