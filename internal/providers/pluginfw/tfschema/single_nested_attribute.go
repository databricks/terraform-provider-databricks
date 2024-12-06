package tfschema

import (
	dataschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// SingleNestedAttributteBuilder represents single complex (struct) types.
type SingleNestedAttributeBuilder struct {
	Attributes         map[string]AttributeBuilder
	Optional           bool
	Required           bool
	Sensitive          bool
	Computed           bool
	DeprecationMessage string
	Validators         []validator.Object
	PlanModifiers      []planmodifier.Object
}

func (a SingleNestedAttributeBuilder) BuildDataSourceAttribute() dataschema.Attribute {
	return dataschema.SingleNestedAttribute{
		Attributes:         BuildDataSourceAttributeMap(a.Attributes),
		Optional:           a.Optional,
		Required:           a.Required,
		Sensitive:          a.Sensitive,
		DeprecationMessage: a.DeprecationMessage,
		Computed:           a.Computed,
		Validators:         a.Validators,
	}
}

func (a SingleNestedAttributeBuilder) BuildResourceAttribute() schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes:         BuildResourceAttributeMap(a.Attributes),
		Optional:           a.Optional,
		Required:           a.Required,
		Sensitive:          a.Sensitive,
		DeprecationMessage: a.DeprecationMessage,
		Computed:           a.Computed,
		Validators:         a.Validators,
		PlanModifiers:      a.PlanModifiers,
	}
}

func (a SingleNestedAttributeBuilder) SetOptional() AttributeBuilder {
	if a.Optional && !a.Required {
		panic("attribute is already optional")
	}
	a.Optional = true
	a.Required = false
	return a
}

func (a SingleNestedAttributeBuilder) SetRequired() AttributeBuilder {
	if !a.Optional && a.Required {
		panic("attribute is already required")
	}
	a.Optional = false
	a.Required = true
	return a
}

func (a SingleNestedAttributeBuilder) SetSensitive() AttributeBuilder {
	if a.Sensitive {
		panic("attribute is already sensitive")
	}
	a.Sensitive = true
	return a
}

func (a SingleNestedAttributeBuilder) SetComputed() AttributeBuilder {
	if a.Computed {
		panic("attribute is already computed")
	}
	a.Computed = true
	return a
}

func (a SingleNestedAttributeBuilder) SetReadOnly() AttributeBuilder {
	if a.Computed && !a.Optional && !a.Required {
		panic("attribute is already read only")
	}
	a.Computed = true
	a.Optional = false
	a.Required = false
	return a
}

func (a SingleNestedAttributeBuilder) SetDeprecated(msg string) BaseSchemaBuilder {
	a.DeprecationMessage = msg
	return a
}

func (a SingleNestedAttributeBuilder) AddValidator(v validator.Object) AttributeBuilder {
	a.Validators = append(a.Validators, v)
	return a
}

func (a SingleNestedAttributeBuilder) AddPlanModifier(v planmodifier.Object) AttributeBuilder {
	a.PlanModifiers = append(a.PlanModifiers, v)
	return a
}
