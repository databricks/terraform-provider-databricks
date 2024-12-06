package tfschema

import (
	dataschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type BoolAttributeBuilder struct {
	Optional           bool
	Required           bool
	Sensitive          bool
	Computed           bool
	DeprecationMessage string
	Validators         []validator.Bool
	PlanModifiers      []planmodifier.Bool
}

func (a BoolAttributeBuilder) BuildDataSourceAttribute() dataschema.Attribute {
	return dataschema.BoolAttribute{
		Optional:           a.Optional,
		Required:           a.Required,
		Sensitive:          a.Sensitive,
		DeprecationMessage: a.DeprecationMessage,
		Computed:           a.Computed,
		Validators:         a.Validators,
	}
}

func (a BoolAttributeBuilder) BuildResourceAttribute() schema.Attribute {
	return schema.BoolAttribute{
		Optional:           a.Optional,
		Required:           a.Required,
		Sensitive:          a.Sensitive,
		DeprecationMessage: a.DeprecationMessage,
		Computed:           a.Computed,
		Validators:         a.Validators,
		PlanModifiers:      a.PlanModifiers,
	}
}

func (a BoolAttributeBuilder) SetOptional() AttributeBuilder {
	if a.Optional && !a.Required {
		panic("attribute is already optional")
	}
	a.Optional = true
	a.Required = false
	return a
}

func (a BoolAttributeBuilder) SetRequired() AttributeBuilder {
	if !a.Optional && a.Required {
		panic("attribute is already required")
	}
	a.Optional = false
	a.Required = true
	return a
}

func (a BoolAttributeBuilder) SetSensitive() AttributeBuilder {
	if a.Sensitive {
		panic("attribute is already sensitive")
	}
	a.Sensitive = true
	return a
}

func (a BoolAttributeBuilder) SetComputed() AttributeBuilder {
	if a.Computed {
		panic("attribute is already computed")
	}
	a.Computed = true
	return a
}

func (a BoolAttributeBuilder) SetReadOnly() AttributeBuilder {
	if a.Computed && !a.Optional && !a.Required {
		panic("attribute is already read only")
	}
	a.Computed = true
	a.Optional = false
	a.Required = false
	return a
}

func (a BoolAttributeBuilder) SetDeprecated(msg string) BaseSchemaBuilder {
	a.DeprecationMessage = msg
	return a
}

func (a BoolAttributeBuilder) AddValidator(v validator.Bool) AttributeBuilder {
	a.Validators = append(a.Validators, v)
	return a
}

func (a BoolAttributeBuilder) AddPlanModifier(v planmodifier.Bool) AttributeBuilder {
	a.PlanModifiers = append(a.PlanModifiers, v)
	return a
}
