package common

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dataschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type Attribute interface {
	ToDataSourceAttribute() dataschema.Attribute
	ToResourceAttribute() schema.Attribute
}

type StringAttribute struct {
	Optional bool
	Required bool
}

func (a StringAttribute) ToDataSourceAttribute() dataschema.Attribute {
	return dataschema.StringAttribute{Optional: a.Optional, Required: a.Required}
}

func (a StringAttribute) ToResourceAttribute() dataschema.Attribute {
	return schema.StringAttribute{Optional: a.Optional, Required: a.Required}
}

type Float64Attribute struct {
	Optional bool
	Required bool
}

func (a Float64Attribute) ToDataSourceAttribute() dataschema.Attribute {
	return dataschema.Float64Attribute{Optional: a.Optional, Required: a.Required}
}

func (a Float64Attribute) ToResourceAttribute() dataschema.Attribute {
	return schema.Float64Attribute{Optional: a.Optional, Required: a.Required}
}

type Int64Attribute struct {
	Optional bool
	Required bool
}

type BoolAttribute struct {
	Optional bool
	Required bool
}

type MapAttribute struct {
	ElementType attr.Type
	Optional    bool
	Required    bool
}

type ListAttribute struct {
	ElementType attr.Type
	Optional    bool
	Required    bool
}

type SingleNestedAttribute struct {
	Optional bool
	Required bool
}

type ListNestedAttribute struct {
	NestedObject NestedAttributeObject
	Optional     bool
	Required     bool
}

type NestedAttributeObject struct {
	Attributes map[string]Attribute
}

type MapNestedAttribute struct {
	NestedObject NestedAttributeObject
	Optional     bool
	Required     bool
}
