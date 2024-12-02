package tfschema

import (
	dataschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// ListNestedBlockBuilder represents a list of complex (non-primitive) types.
// To be compatible with our sdkv2 schema, all struct types in the gosdk are represented with this type.
type ListNestedBlockBuilder struct {
	NestedObject       NestedBlockObject
	DeprecationMessage string
	Validators         []validator.List
	PlanModifiers      []planmodifier.List
}

func (a ListNestedBlockBuilder) ToAttribute() AttributeBuilder {
	return ListNestedAttributeBuilder{
		NestedObject:       a.NestedObject.ToNestedAttributeObject(),
		DeprecationMessage: a.DeprecationMessage,
		Validators:         a.Validators,
		PlanModifiers:      a.PlanModifiers,
	}
}

func (a ListNestedBlockBuilder) BuildDataSourceBlock() dataschema.Block {
	return dataschema.ListNestedBlock{
		NestedObject:       a.NestedObject.BuildDataSourceAttribute(),
		DeprecationMessage: a.DeprecationMessage,
		Validators:         a.Validators,
	}
}

func (a ListNestedBlockBuilder) BuildResourceBlock() schema.Block {
	return schema.ListNestedBlock{
		NestedObject:       a.NestedObject.BuildResourceAttribute(),
		DeprecationMessage: a.DeprecationMessage,
		Validators:         a.Validators,
		PlanModifiers:      a.PlanModifiers,
	}
}

func (a ListNestedBlockBuilder) SetDeprecated(msg string) BaseSchemaBuilder {
	a.DeprecationMessage = msg
	return a
}

func (a ListNestedBlockBuilder) AddValidator(v validator.List) BaseSchemaBuilder {
	a.Validators = append(a.Validators, v)
	return a
}

func (a ListNestedBlockBuilder) AddPlanModifier(v planmodifier.List) BaseSchemaBuilder {
	a.PlanModifiers = append(a.PlanModifiers, v)
	return a
}
