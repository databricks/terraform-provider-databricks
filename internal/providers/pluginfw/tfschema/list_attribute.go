package tfschema

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dataschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// ListAttributteBuilder represents a list of primitive types.
type ListAttributeBuilder struct {
	ElementType        attr.Type
	Optional           bool
	Required           bool
	Sensitive          bool
	Computed           bool
	DeprecationMessage string
	Validators         []validator.List
	PlanModifiers      []planmodifier.List
}

func (a ListAttributeBuilder) BuildDataSourceAttribute() dataschema.Attribute {
	return dataschema.ListAttribute{
		ElementType:        a.ElementType,
		Optional:           a.Optional,
		Required:           a.Required,
		Sensitive:          a.Sensitive,
		DeprecationMessage: a.DeprecationMessage,
		Computed:           a.Computed,
		Validators:         a.Validators,
	}
}

func (a ListAttributeBuilder) BuildResourceAttribute() schema.Attribute {
	return schema.ListAttribute{
		ElementType:        a.ElementType,
		Optional:           a.Optional,
		Required:           a.Required,
		Sensitive:          a.Sensitive,
		DeprecationMessage: a.DeprecationMessage,
		Computed:           a.Computed,
		Validators:         a.Validators,
		PlanModifiers:      a.PlanModifiers,
	}
}

func (a ListAttributeBuilder) SetOptional() BaseSchemaBuilder {
	if a.Optional && !a.Required {
		panic("attribute is already optional")
	}
	a.Optional = true
	a.Required = false
	return a
}

func (a ListAttributeBuilder) SetRequired() BaseSchemaBuilder {
	if !a.Optional && a.Required {
		panic("attribute is already required")
	}
	a.Optional = false
	a.Required = true
	return a
}

func (a ListAttributeBuilder) SetSensitive() BaseSchemaBuilder {
	if a.Sensitive {
		panic("attribute is already sensitive")
	}
	a.Sensitive = true
	return a
}

func (a ListAttributeBuilder) SetComputed() BaseSchemaBuilder {
	if a.Computed {
		panic("attribute is already computed")
	}
	a.Computed = true
	return a
}

func (a ListAttributeBuilder) SetReadOnly() BaseSchemaBuilder {
	if a.Computed && !a.Optional && !a.Required {
		panic("attribute is already read only")
	}
	a.Computed = true
	a.Optional = false
	a.Required = false
	return a
}

func (a ListAttributeBuilder) SetDeprecated(msg string) BaseSchemaBuilder {
	a.DeprecationMessage = msg
	return a
}

func (a ListAttributeBuilder) AddValidator(v validator.List) BaseSchemaBuilder {
	a.Validators = append(a.Validators, v)
	return a
}

func (a ListAttributeBuilder) AddPlanModifier(v planmodifier.List) BaseSchemaBuilder {
	a.PlanModifiers = append(a.PlanModifiers, v)
	return a
}
