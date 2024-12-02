package tfschema

// BaseSchemaBuilder is the common interface for all blocks and attributes, it can be used to build data source and resource.
// Both AttributeBuilder and BlockBuilder extend this interface.
type BaseSchemaBuilder interface {
	SetOptional() BaseSchemaBuilder
	SetRequired() BaseSchemaBuilder
	SetSensitive() BaseSchemaBuilder
	SetComputed() BaseSchemaBuilder
	SetReadOnly() BaseSchemaBuilder
	SetDeprecated(string) BaseSchemaBuilder
}
