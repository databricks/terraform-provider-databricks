package tfschema

type Blockable interface {
	ToBlock() BlockBuilder
}

func convertAttributesToBlocks(attributes map[string]AttributeBuilder, blocks map[string]BlockBuilder) NestedBlockObject {
	newAttributes := make(map[string]AttributeBuilder)
	newBlocks := make(map[string]BlockBuilder)
	for name, attr := range attributes {
		if lnab, ok := attr.(Blockable); ok {
			newBlocks[name] = lnab.ToBlock()
		} else {
			newAttributes[name] = attr
		}
	}
	for name, block := range blocks {
		newBlocks[name] = block
	}
	return NestedBlockObject{
		Attributes: newAttributes,
		Blocks:     newBlocks,
	}
}
