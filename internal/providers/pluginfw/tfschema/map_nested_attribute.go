package tfschema

import (
	dataschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// MapNestedAttributteBuilder represents a map of complex (non-primitive) types.
type MapNestedAttributeBuilder struct {
	NestedObject       NestedAttributeObject
	Optional           bool
	Required           bool
	Sensitive          bool
	Computed           bool
	DeprecationMessage string
	Validators         []validator.Map
	PlanModifiers      []planmodifier.Map
}

func (a MapNestedAttributeBuilder) BuildDataSourceAttribute() dataschema.Attribute {
	return dataschema.MapNestedAttribute{
		NestedObject:       a.NestedObject.BuildDataSourceAttribute(),
		Optional:           a.Optional,
		Required:           a.Required,
		Sensitive:          a.Sensitive,
		DeprecationMessage: a.DeprecationMessage,
		Computed:           a.Computed,
		Validators:         a.Validators,
	}
}

func (a MapNestedAttributeBuilder) BuildResourceAttribute() schema.Attribute {
	return schema.MapNestedAttribute{
		NestedObject:       a.NestedObject.BuildResourceAttribute(),
		Optional:           a.Optional,
		Required:           a.Required,
		Sensitive:          a.Sensitive,
		DeprecationMessage: a.DeprecationMessage,
		Computed:           a.Computed,
		Validators:         a.Validators,
		PlanModifiers:      a.PlanModifiers,
	}
}

func (a MapNestedAttributeBuilder) SetOptional() AttributeBuilder {
	if a.Optional && !a.Required {
		panic("attribute is already optional")
	}
	a.Optional = true
	a.Required = false
	return a
}

func (a MapNestedAttributeBuilder) SetRequired() AttributeBuilder {
	if !a.Optional && a.Required {
		panic("attribute is already required")
	}
	a.Optional = false
	a.Required = true
	return a
}

func (a MapNestedAttributeBuilder) SetSensitive() AttributeBuilder {
	if a.Sensitive {
		panic("attribute is already sensitive")
	}
	a.Sensitive = true
	return a
}

func (a MapNestedAttributeBuilder) SetComputed() AttributeBuilder {
	if a.Computed {
		panic("attribute is already computed")
	}
	a.Computed = true
	return a
}

func (a MapNestedAttributeBuilder) SetReadOnly() AttributeBuilder {
	if a.Computed && !a.Optional && !a.Required {
		panic("attribute is already read only")
	}
	a.Computed = true
	a.Optional = false
	a.Required = false
	return a
}

func (a MapNestedAttributeBuilder) SetDeprecated(msg string) BaseSchemaBuilder {
	a.DeprecationMessage = msg
	return a
}

func (a MapNestedAttributeBuilder) AddValidator(v validator.Map) AttributeBuilder {
	a.Validators = append(a.Validators, v)
	return a
}

func (a MapNestedAttributeBuilder) AddPlanModifier(v planmodifier.Map) AttributeBuilder {
	a.PlanModifiers = append(a.PlanModifiers, v)
	return a
}
