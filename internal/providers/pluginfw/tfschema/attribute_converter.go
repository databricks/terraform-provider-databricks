// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
package tfschema

import "maps"

// Blockable is an interface that can be implemented by an AttributeBuilder to convert it to a BlockBuilder.
type Blockable interface {
	// ToBlock converts the AttributeBuilder to a BlockBuilder.
	ToBlock() BlockBuilder
}

// convertAttributesToBlocks converts all attributes implementing the Blockable interface to blocks, returning
// a new NestedBlockObject with the converted attributes and the original blocks.
// SingleNestedAttributeBuilder (types.Object) is excluded from conversion because it is used by
// fields like provider_config that are new to the Plugin Framework and have no SDKv2 block equivalent.
func convertAttributesToBlocks(attributes map[string]AttributeBuilder, blocks map[string]BlockBuilder) NestedBlockObject {
	newAttributes := make(map[string]AttributeBuilder)
	newBlocks := make(map[string]BlockBuilder)
	for name, attr := range attributes {
		if _, isSingle := attr.(SingleNestedAttributeBuilder); isSingle {
			// SingleNestedAttributeBuilder corresponds to types.Object, which has no
			// SDKv2 block equivalent. Keep it as an attribute.
			newAttributes[name] = attr
		} else if lnab, ok := attr.(Blockable); ok {
			newBlocks[name] = lnab.ToBlock()
		} else {
			newAttributes[name] = attr
		}
	}
	maps.Copy(newBlocks, blocks)
	if len(newAttributes) == 0 {
		newAttributes = nil
	}
	if len(newBlocks) == 0 {
		newBlocks = nil
	}
	return NestedBlockObject{
		Attributes: newAttributes,
		Blocks:     newBlocks,
	}
}
