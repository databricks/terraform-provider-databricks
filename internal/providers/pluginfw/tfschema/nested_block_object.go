package tfschema

import (
	dataschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// NestedAttributteObject is the intermediate type for nested complex (non-primitive) types.
type NestedBlockObject struct {
	Attributes map[string]AttributeBuilder
	Blocks     map[string]BlockBuilder
}

func (a NestedBlockObject) BuildDataSourceAttribute() dataschema.NestedBlockObject {
	dataSourceAttributes := BuildDataSourceAttributeMap(a.Attributes)
	dataSourceBlocks := BuildDataSourceBlockMap(a.Blocks)

	return dataschema.NestedBlockObject{
		Attributes: dataSourceAttributes,
		Blocks:     dataSourceBlocks,
	}
}

func (a NestedBlockObject) BuildResourceAttribute() schema.NestedBlockObject {
	resourceAttributes := BuildResourceAttributeMap(a.Attributes)
	resourceBlocks := BuildResourceBlockMap(a.Blocks)

	return schema.NestedBlockObject{
		Attributes: resourceAttributes,
		Blocks:     resourceBlocks,
	}
}
