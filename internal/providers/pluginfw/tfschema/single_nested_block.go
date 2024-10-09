package tfschema

import (
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
	dataschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// SingleNestedBlockBuilder represents a list of complex (non-primitive) types.
type SingleNestedBlockBuilder struct {
	NestedObject       NestedBlockObject
	Optional           bool
	Required           bool
	Sensitive          bool
	Computed           bool
	DeprecationMessage string
	Validators         []validator.Object
	PlanModifiers      []planmodifier.Object
}

func (a SingleNestedBlockBuilder) BuildDataSourceAttribute() dataschema.Attribute {
	panic(fmt.Errorf("BuildDataSourceBlock should never be called for SingleNestedBlockBuilder. %s", common.TerraformBugErrorMessage))
}

func (a SingleNestedBlockBuilder) BuildResourceAttribute() schema.Attribute {
	panic(fmt.Errorf("BuildResourceBlock should never be called for SingleNestedBlockBuilder. %s", common.TerraformBugErrorMessage))
}

func (a SingleNestedBlockBuilder) BuildDataSourceBlock() dataschema.Block {
	return dataschema.SingleNestedBlock{
		Attributes:         a.NestedObject.BuildDataSourceAttribute().Attributes,
		Blocks:             a.NestedObject.BuildDataSourceAttribute().Blocks,
		DeprecationMessage: a.DeprecationMessage,
		Validators:         a.Validators,
	}
}

func (a SingleNestedBlockBuilder) BuildResourceBlock() schema.Block {
	return schema.SingleNestedBlock{
		Attributes:         a.NestedObject.BuildResourceAttribute().Attributes,
		Blocks:             a.NestedObject.BuildResourceAttribute().Blocks,
		DeprecationMessage: a.DeprecationMessage,
		Validators:         a.Validators,
		PlanModifiers:      a.PlanModifiers,
	}
}

func (a SingleNestedBlockBuilder) SetOptional() BaseSchemaBuilder {
	if a.Optional && !a.Required {
		panic("attribute is already optional")
	}
	a.Optional = true
	a.Required = false
	return a
}

func (a SingleNestedBlockBuilder) SetRequired() BaseSchemaBuilder {
	if !a.Optional && a.Required {
		panic("attribute is already required")
	}
	a.Optional = false
	a.Required = true
	return a
}

func (a SingleNestedBlockBuilder) SetSensitive() BaseSchemaBuilder {
	if a.Sensitive {
		panic("attribute is already sensitive")
	}
	a.Sensitive = true
	return a
}

func (a SingleNestedBlockBuilder) SetComputed() BaseSchemaBuilder {
	if a.Computed {
		panic("attribute is already computed")
	}
	a.Computed = true
	return a
}

func (a SingleNestedBlockBuilder) SetReadOnly() BaseSchemaBuilder {
	if a.Computed && !a.Optional && !a.Required {
		panic("attribute is already read only")
	}
	a.Computed = true
	a.Optional = false
	a.Required = false
	return a
}

func (a SingleNestedBlockBuilder) SetDeprecated(msg string) BaseSchemaBuilder {
	a.DeprecationMessage = msg
	return a
}

func (a SingleNestedBlockBuilder) AddValidator(v validator.Object) BaseSchemaBuilder {
	a.Validators = append(a.Validators, v)
	return a
}

func (a SingleNestedBlockBuilder) AddPlanModifier(v planmodifier.Object) BaseSchemaBuilder {
	a.PlanModifiers = append(a.PlanModifiers, v)
	return a
}
