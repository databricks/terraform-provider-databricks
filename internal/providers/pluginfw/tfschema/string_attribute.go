// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
package tfschema

import (
	dataschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type StringAttributeBuilder struct {
	Computed           bool
	CustomType         basetypes.StringTypable
	DeprecationMessage string
	Optional           bool
	PlanModifiers      []planmodifier.String
	Required           bool
	Sensitive          bool
	Validators         []validator.String
}

func (a StringAttributeBuilder) BuildDataSourceAttribute() dataschema.Attribute {
	return dataschema.StringAttribute{
		Computed:           a.Computed,
		CustomType:         a.CustomType,
		DeprecationMessage: a.DeprecationMessage,
		Optional:           a.Optional,
		Required:           a.Required,
		Sensitive:          a.Sensitive,
		Validators:         a.Validators,
	}
}

func (a StringAttributeBuilder) BuildResourceAttribute() schema.Attribute {
	return schema.StringAttribute{
		Computed:           a.Computed,
		CustomType:         a.CustomType,
		DeprecationMessage: a.DeprecationMessage,
		Optional:           a.Optional,
		PlanModifiers:      a.PlanModifiers,
		Required:           a.Required,
		Sensitive:          a.Sensitive,
		Validators:         a.Validators,
	}
}

func (a StringAttributeBuilder) SetOptional() AttributeBuilder {
	if a.Optional && !a.Required {
		panic("attribute is already optional")
	}
	a.Optional = true
	a.Required = false
	return a
}

func (a StringAttributeBuilder) SetRequired() AttributeBuilder {
	if !a.Optional && a.Required {
		panic("attribute is already required")
	}
	a.Optional = false
	a.Required = true
	return a
}

func (a StringAttributeBuilder) SetSensitive() AttributeBuilder {
	if a.Sensitive {
		panic("attribute is already sensitive")
	}
	a.Sensitive = true
	return a
}

func (a StringAttributeBuilder) SetComputed() AttributeBuilder {
	if a.Computed {
		panic("attribute is already computed")
	}
	a.Computed = true
	return a
}

func (a StringAttributeBuilder) SetReadOnly() AttributeBuilder {
	if a.Computed && !a.Optional && !a.Required {
		panic("attribute is already read only")
	}
	a.Computed = true
	a.Optional = false
	a.Required = false
	return a
}

func (a StringAttributeBuilder) SetDeprecated(msg string) BaseSchemaBuilder {
	a.DeprecationMessage = msg
	return a
}

func (a StringAttributeBuilder) AddValidator(v validator.String) AttributeBuilder {
	a.Validators = append(a.Validators, v)
	return a
}

func (a StringAttributeBuilder) AddPlanModifier(v planmodifier.String) AttributeBuilder {
	a.PlanModifiers = append(a.PlanModifiers, v)
	return a
}
