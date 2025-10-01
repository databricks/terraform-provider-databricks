package common

// ProviderConfig represents provider configuration fields that may be embedded in resource and data source types
type ProviderConfig struct {
	WorkspaceId string `json:"workspace_id,omitempty" tf:"optional"`
}

// CustomizeSchemaFunc is a function type that can be used to customize schema
type CustomizeSchemaFunc func(s *CustomizableSchema) *CustomizableSchema
