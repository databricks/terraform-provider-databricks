package tfschema

import (
	dataschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// NestedAttributteObject is the intermediate type for nested complex (non-primitive) types.
type NestedAttributeObject struct {
	Attributes map[string]AttributeBuilder
}

func (a NestedAttributeObject) BuildDataSourceAttribute() dataschema.NestedAttributeObject {
	dataSourceAttributes := BuildDataSourceAttributeMap(a.Attributes)

	return dataschema.NestedAttributeObject{
		Attributes: dataSourceAttributes,
	}
}

func (a NestedAttributeObject) BuildResourceAttribute() schema.NestedAttributeObject {
	resourceAttributes := BuildResourceAttributeMap(a.Attributes)

	return schema.NestedAttributeObject{
		Attributes: resourceAttributes,
	}
}

func (a NestedAttributeObject) SetReadOnly() {
	for attr, attrV := range a.Attributes {
		a.Attributes[attr] = attrV.SetReadOnly()
	}
}
