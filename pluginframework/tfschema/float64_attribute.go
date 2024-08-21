package tfschema

import (
	dataschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type Float64Attribute struct {
	Optional           bool
	Required           bool
	Sensitive          bool
	Computed           bool
	DeprecationMessage string
	Validators         []validator.Float64
}

func (a Float64Attribute) BuildDataSourceAttribute() dataschema.Attribute {
	return dataschema.Float64Attribute{Optional: a.Optional, Required: a.Required, Sensitive: a.Sensitive, DeprecationMessage: a.DeprecationMessage, Computed: a.Computed, Validators: a.Validators}
}

func (a Float64Attribute) BuildResourceAttribute() schema.Attribute {
	return schema.Float64Attribute{Optional: a.Optional, Required: a.Required, Sensitive: a.Sensitive, DeprecationMessage: a.DeprecationMessage, Computed: a.Computed, Validators: a.Validators}
}

func (a Float64Attribute) SetOptional() AttributeBuilder {
	if a.Optional && !a.Required {
		panic("attribute is already optional")
	}
	a.Optional = true
	a.Required = false
	return a
}

func (a Float64Attribute) SetRequired() AttributeBuilder {
	if !a.Optional && a.Required {
		panic("attribute is already required")
	}
	a.Optional = false
	a.Required = true
	return a
}

func (a Float64Attribute) SetSensitive() AttributeBuilder {
	if a.Sensitive {
		panic("attribute is already sensitive")
	}
	a.Sensitive = true
	return a
}

func (a Float64Attribute) SetComputed() AttributeBuilder {
	if a.Computed {
		panic("attribute is already computed")
	}
	a.Computed = true
	return a
}

func (a Float64Attribute) SetReadOnly() AttributeBuilder {
	if a.Computed && !a.Optional && !a.Required {
		panic("attribute is already read only")
	}
	a.Computed = true
	a.Optional = false
	a.Required = false
	return a
}

func (a Float64Attribute) SetDeprecated(msg string) AttributeBuilder {
	a.DeprecationMessage = msg
	return a
}

func (a Float64Attribute) AddValidator(v validator.Float64) AttributeBuilder {
	a.Validators = append(a.Validators, v)
	return a
}
