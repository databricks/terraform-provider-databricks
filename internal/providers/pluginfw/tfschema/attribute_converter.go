package tfschema

type BlockToAttributeConverter interface {
	// ConvertBlockToAttribute converts a contained block to its corresponding attribute type.
	ConvertBlockToAttribute(string) BaseSchemaBuilder
}
