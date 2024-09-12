package tfschema

import (
	dataschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type Float64AttributeBuilder struct {
	Optional           bool
	Required           bool
	Sensitive          bool
	Computed           bool
	DeprecationMessage string
	Validators         []validator.Float64
	PlanModifiers      []planmodifier.Float64
}

func (a Float64AttributeBuilder) BuildDataSourceAttribute() dataschema.Attribute {
	return dataschema.Float64Attribute{
		Optional:           a.Optional,
		Required:           a.Required,
		Sensitive:          a.Sensitive,
		DeprecationMessage: a.DeprecationMessage,
		Computed:           a.Computed,
		Validators:         a.Validators,
	}
}

func (a Float64AttributeBuilder) BuildResourceAttribute() schema.Attribute {
	return schema.Float64Attribute{
		Optional:           a.Optional,
		Required:           a.Required,
		Sensitive:          a.Sensitive,
		DeprecationMessage: a.DeprecationMessage,
		Computed:           a.Computed,
		Validators:         a.Validators,
		PlanModifiers:      a.PlanModifiers,
	}
}

func (a Float64AttributeBuilder) SetOptional() AttributeBuilder {
	if a.Optional && !a.Required {
		panic("attribute is already optional")
	}
	a.Optional = true
	a.Required = false
	return a
}

func (a Float64AttributeBuilder) SetRequired() AttributeBuilder {
	if !a.Optional && a.Required {
		panic("attribute is already required")
	}
	a.Optional = false
	a.Required = true
	return a
}

func (a Float64AttributeBuilder) SetSensitive() AttributeBuilder {
	if a.Sensitive {
		panic("attribute is already sensitive")
	}
	a.Sensitive = true
	return a
}

func (a Float64AttributeBuilder) SetComputed() AttributeBuilder {
	if a.Computed {
		panic("attribute is already computed")
	}
	a.Computed = true
	return a
}

func (a Float64AttributeBuilder) SetReadOnly() AttributeBuilder {
	if a.Computed && !a.Optional && !a.Required {
		panic("attribute is already read only")
	}
	a.Computed = true
	a.Optional = false
	a.Required = false
	return a
}

func (a Float64AttributeBuilder) SetDeprecated(msg string) AttributeBuilder {
	a.DeprecationMessage = msg
	return a
}

func (a Float64AttributeBuilder) AddValidator(v validator.Float64) AttributeBuilder {
	a.Validators = append(a.Validators, v)
	return a
}

func (a Float64AttributeBuilder) AddPlanModifier(v planmodifier.Float64) AttributeBuilder {
	a.PlanModifiers = append(a.PlanModifiers, v)
	return a
}
