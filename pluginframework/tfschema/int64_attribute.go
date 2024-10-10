package tfschema

import (
	dataschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type Int64Attribute struct {
	Optional           bool
	Required           bool
	Sensitive          bool
	Computed           bool
	DeprecationMessage string
	Validators         []validator.Int64
}

func (a Int64Attribute) BuildDataSourceAttribute() dataschema.Attribute {
	return dataschema.Int64Attribute{Optional: a.Optional, Required: a.Required, Sensitive: a.Sensitive, DeprecationMessage: a.DeprecationMessage, Computed: a.Computed, Validators: a.Validators}
}

func (a Int64Attribute) BuildResourceAttribute() schema.Attribute {
	return schema.Int64Attribute{Optional: a.Optional, Required: a.Required, Sensitive: a.Sensitive, DeprecationMessage: a.DeprecationMessage, Computed: a.Computed, Validators: a.Validators}
}

func (a Int64Attribute) SetOptional() AttributeBuilder {
	if a.Optional && !a.Required {
		panic("attribute is already optional")
	}
	a.Optional = true
	a.Required = false
	return a
}

func (a Int64Attribute) SetRequired() AttributeBuilder {
	if !a.Optional && a.Required {
		panic("attribute is already required")
	}
	a.Optional = false
	a.Required = true
	return a
}

func (a Int64Attribute) SetSensitive() AttributeBuilder {
	if a.Sensitive {
		panic("attribute is already sensitive")
	}
	a.Sensitive = true
	return a
}

func (a Int64Attribute) SetComputed() AttributeBuilder {
	if a.Computed {
		panic("attribute is already computed")
	}
	a.Computed = true
	return a
}

func (a Int64Attribute) SetReadOnly() AttributeBuilder {
	if a.Computed && !a.Optional && !a.Required {
		panic("attribute is already read only")
	}
	a.Computed = true
	a.Optional = false
	a.Required = false
	return a
}

func (a Int64Attribute) SetDeprecated(msg string) AttributeBuilder {
	a.DeprecationMessage = msg
	return a
}

func (a Int64Attribute) AddValidator(v validator.Int64) AttributeBuilder {
	a.Validators = append(a.Validators, v)
	return a
}
