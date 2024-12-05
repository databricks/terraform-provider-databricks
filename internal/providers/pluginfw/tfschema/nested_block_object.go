package tfschema

import (
	dataschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// NestedAttributteObject is the intermediate type for nested complex (non-primitive) types.
type NestedBlockObject struct {
	Attributes    map[string]AttributeBuilder
	Blocks        map[string]BlockBuilder
	Validators    []validator.Object
	PlanModifiers []planmodifier.Object
}

func (a NestedBlockObject) ToNestedAttributeObject() NestedAttributeObject {
	attributes := make(map[string]AttributeBuilder)
	for k, v := range a.Attributes {
		attributes[k] = v
	}
	for k, v := range a.Blocks {
		attributes[k] = v.ToAttribute()
	}
	return NestedAttributeObject{
		Attributes: attributes,
	}
}

func (a NestedBlockObject) ToNestedAttributeObject() NestedAttributeObject {
	attributes := make(map[string]AttributeBuilder)
	for k, v := range a.Attributes {
		attributes[k] = v
	}
	for k, v := range a.Blocks {
		attributes[k] = v.ToAttribute()
	}
	return NestedAttributeObject{
		Attributes: attributes,
	}
}

func (a NestedBlockObject) BuildDataSourceAttribute() dataschema.NestedBlockObject {
	dataSourceAttributes := BuildDataSourceAttributeMap(a.Attributes)
	dataSourceBlocks := BuildDataSourceBlockMap(a.Blocks)

	return dataschema.NestedBlockObject{
		Attributes: dataSourceAttributes,
		Blocks:     dataSourceBlocks,
		Validators: a.Validators,
	}
}

func (a NestedBlockObject) BuildResourceAttribute() schema.NestedBlockObject {
	resourceAttributes := BuildResourceAttributeMap(a.Attributes)
	resourceBlocks := BuildResourceBlockMap(a.Blocks)

	return schema.NestedBlockObject{
		Attributes:    resourceAttributes,
		Blocks:        resourceBlocks,
		Validators:    a.Validators,
		PlanModifiers: a.PlanModifiers,
	}
}

func (a *NestedBlockObject) AddValidator(v validator.Object) {
	a.Validators = append(a.Validators, v)
}

func (a *NestedBlockObject) AddPlanModifier(p planmodifier.Object) {
	a.PlanModifiers = append(a.PlanModifiers, p)
}
