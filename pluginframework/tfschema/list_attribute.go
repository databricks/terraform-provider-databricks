package tfschema

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dataschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// ListAttributte represents a list of primitive types.
type ListAttribute struct {
	ElementType        attr.Type
	Optional           bool
	Required           bool
	Sensitive          bool
	Computed           bool
	DeprecationMessage string
	Validators         []validator.List
}

func (a ListAttribute) BuildDataSourceAttribute() dataschema.Attribute {
	return dataschema.ListAttribute{ElementType: a.ElementType, Optional: a.Optional, Required: a.Required, Sensitive: a.Sensitive, DeprecationMessage: a.DeprecationMessage, Computed: a.Computed, Validators: a.Validators}
}

func (a ListAttribute) BuildResourceAttribute() schema.Attribute {
	return schema.ListAttribute{ElementType: a.ElementType, Optional: a.Optional, Required: a.Required, Sensitive: a.Sensitive, DeprecationMessage: a.DeprecationMessage, Computed: a.Computed, Validators: a.Validators}
}

func (a ListAttribute) SetOptional() AttributeBuilder {
	if a.Optional && !a.Required {
		panic("attribute is already optional")
	}
	a.Optional = true
	a.Required = false
	return a
}

func (a ListAttribute) SetRequired() AttributeBuilder {
	if !a.Optional && a.Required {
		panic("attribute is already required")
	}
	a.Optional = false
	a.Required = true
	return a
}

func (a ListAttribute) SetSensitive() AttributeBuilder {
	if a.Sensitive {
		panic("attribute is already sensitive")
	}
	a.Sensitive = true
	return a
}

func (a ListAttribute) SetComputed() AttributeBuilder {
	if a.Computed {
		panic("attribute is already computed")
	}
	a.Computed = true
	return a
}

func (a ListAttribute) SetReadOnly() AttributeBuilder {
	if a.Computed && !a.Optional && !a.Required {
		panic("attribute is already read only")
	}
	a.Computed = true
	a.Optional = false
	a.Required = false
	return a
}

func (a ListAttribute) SetDeprecated(msg string) AttributeBuilder {
	a.DeprecationMessage = msg
	return a
}

func (a ListAttribute) AddValidator(v validator.List) AttributeBuilder {
	a.Validators = append(a.Validators, v)
	return a
}
