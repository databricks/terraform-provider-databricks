package main

import "encoding/json"

// ProviderSchemas is the top-level shape of `terraform providers schema -json`.
type ProviderSchemas struct {
	FormatVersion   string                    `json:"format_version"`
	ProviderSchemas map[string]ProviderSchema `json:"provider_schemas"`
}

type ProviderSchema struct {
	Provider          Schema            `json:"provider"`
	ResourceSchemas   map[string]Schema `json:"resource_schemas"`
	DataSourceSchemas map[string]Schema `json:"data_source_schemas"`
}

type Schema struct {
	Version int   `json:"version"`
	Block   Block `json:"block"`
}

type Block struct {
	Attributes      map[string]Attribute `json:"attributes,omitempty"`
	BlockTypes      map[string]BlockType `json:"block_types,omitempty"`
	Description     string               `json:"description,omitempty"`
	DescriptionKind string               `json:"description_kind,omitempty"`
	Deprecated      bool                 `json:"deprecated,omitempty"`
}

// Attribute is one field on a block. The `type` value is a cty type
// descriptor — either a JSON string ("string", "number", "bool") or a
// nested JSON array describing collection / object types. We keep it as
// RawMessage and compare canonical JSON bytes; per the rule taxonomy any
// type delta (including element-type widening) is breaking.
type Attribute struct {
	Type            json.RawMessage `json:"type,omitempty"`
	NestedType      json.RawMessage `json:"nested_type,omitempty"`
	Description     string          `json:"description,omitempty"`
	DescriptionKind string          `json:"description_kind,omitempty"`
	Required        bool            `json:"required,omitempty"`
	Optional        bool            `json:"optional,omitempty"`
	Computed        bool            `json:"computed,omitempty"`
	Sensitive       bool            `json:"sensitive,omitempty"`
	Deprecated      bool            `json:"deprecated,omitempty"`
}

type BlockType struct {
	NestingMode string `json:"nesting_mode"`
	Block       Block  `json:"block"`
	MinItems    int    `json:"min_items,omitempty"`
	MaxItems    int    `json:"max_items,omitempty"`
}

// IsComputedOnly returns true when the attribute is computed-only (i.e. an
// output the user cannot set). Required+Computed and Optional+Computed
// attributes are *settable*.
func (a Attribute) IsComputedOnly() bool {
	return a.Computed && !a.Required && !a.Optional
}

// IsSettable returns true when the user can set the attribute in config.
func (a Attribute) IsSettable() bool {
	return a.Required || a.Optional
}
