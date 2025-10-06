// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
package tfschema

// Blockable is an interface that can be implemented by an AttributeBuilder to convert it to a BlockBuilder.
type Blockable interface {
	// ToBlock converts the AttributeBuilder to a BlockBuilder.
	ToBlock() BlockBuilder
}

// convertAttributesToBlocks converts all attributes implementing the Blockable interface to blocks except for "provider_config",
// returning a new NestedBlockObject with the converted attributes and the original blocks.
func convertAttributesToBlocks(attributes map[string]AttributeBuilder, blocks map[string]BlockBuilder) NestedBlockObject {
	newAttributes := make(map[string]AttributeBuilder)
	newBlocks := make(map[string]BlockBuilder)
	for name, attr := range attributes {
		if name == "provider_config" {
			newAttributes[name] = attr
			continue
		}
		if lnab, ok := attr.(Blockable); ok {
			newBlocks[name] = lnab.ToBlock()
		} else {
			newAttributes[name] = attr
		}
	}
	for name, block := range blocks {
		newBlocks[name] = block
	}
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
