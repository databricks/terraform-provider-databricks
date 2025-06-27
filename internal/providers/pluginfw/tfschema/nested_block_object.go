// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
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

func (a NestedBlockObject) ToNestedAttributeObject() NestedAttributeObject {
	attributes := make(map[string]AttributeBuilder)
	for k, v := range a.Attributes {
		attributes[k] = v
	}
	for k, v := range a.Blocks {
		attributes[k] = v.ToAttribute()
	}
	return NestedAttributeObject{
		Attributes: attributes,
	}
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
