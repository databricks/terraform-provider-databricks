package tfschema

import (
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
	dataschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// SingleNestedBlockBuilder represents a single nested complex (non-primitive) type.
type SingleNestedBlockBuilder struct {
	NestedObject       NestedBlockObject
	DeprecationMessage string
	Validators         []validator.Object
	PlanModifiers      []planmodifier.Object
}

func (a SingleNestedBlockBuilder) ToAttribute() AttributeBuilder {
	return SingleNestedAttributeBuilder{
		Attributes:         a.NestedObject.ToNestedAttributeObject().Attributes,
		DeprecationMessage: a.DeprecationMessage,
		Validators:         a.Validators,
		PlanModifiers:      a.PlanModifiers,
	}
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
