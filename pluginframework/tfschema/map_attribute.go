package tfschema

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dataschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// MapAttributte represents a map of primitive types.
type MapAttribute struct {
	ElementType        attr.Type
	Optional           bool
	Required           bool
	Sensitive          bool
	Computed           bool
	DeprecationMessage string
	Validators         []validator.Map
}

func (a MapAttribute) BuildDataSourceAttribute() dataschema.Attribute {
	return dataschema.MapAttribute{ElementType: a.ElementType, Optional: a.Optional, Required: a.Required, Sensitive: a.Sensitive, DeprecationMessage: a.DeprecationMessage, Computed: a.Computed, Validators: a.Validators}
}

func (a MapAttribute) BuildResourceAttribute() schema.Attribute {
	return schema.MapAttribute{ElementType: a.ElementType, Optional: a.Optional, Required: a.Required, Sensitive: a.Sensitive, DeprecationMessage: a.DeprecationMessage, Computed: a.Computed, Validators: a.Validators}
}

func (a MapAttribute) SetOptional() AttributeBuilder {
	if a.Optional && !a.Required {
		panic("attribute is already optional")
	}
	a.Optional = true
	a.Required = false
	return a
}

func (a MapAttribute) SetRequired() AttributeBuilder {
	if !a.Optional && a.Required {
		panic("attribute is already required")
	}
	a.Optional = false
	a.Required = true
	return a
}

func (a MapAttribute) SetSensitive() AttributeBuilder {
	if a.Sensitive {
		panic("attribute is already sensitive")
	}
	a.Sensitive = true
	return a
}

func (a MapAttribute) SetComputed() AttributeBuilder {
	if a.Computed {
		panic("attribute is already computed")
	}
	a.Computed = true
	return a
}

func (a MapAttribute) SetReadOnly() AttributeBuilder {
	if a.Computed && !a.Optional && !a.Required {
		panic("attribute is already read only")
	}
	a.Computed = true
	a.Optional = false
	a.Required = false
	return a
}

func (a MapAttribute) SetDeprecated(msg string) AttributeBuilder {
	a.DeprecationMessage = msg
	return a
}

func (a MapAttribute) AddValidator(v validator.Map) AttributeBuilder {
	a.Validators = append(a.Validators, v)
	return a
}
