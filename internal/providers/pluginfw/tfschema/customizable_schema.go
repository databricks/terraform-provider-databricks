package tfschema

import (
	"fmt"
	"reflect"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// CustomizableSchema is a wrapper struct on top of AttributeBuilder that can be used to navigate through nested schema add customizations.
type CustomizableSchema struct {
	attr BaseSchemaBuilder
}

// ConstructCustomizableSchema constructs a CustomizableSchema given a map from string to AttributeBuilder.
func ConstructCustomizableSchema(nestedObject NestedBlockObject) *CustomizableSchema {
	attr := AttributeBuilder(SingleNestedBlockBuilder{NestedObject: nestedObject})
	return &CustomizableSchema{attr: attr}
}

// ToAttributeMap converts CustomizableSchema into a map from string to Attribute.
func (s *CustomizableSchema) ToAttributeMap() NestedBlockObject {
	return attributeToMap(&s.attr)
}

// attributeToMap converts AttributeBuilder into a map from string to AttributeBuilder.
func attributeToMap(attr *BaseSchemaBuilder) NestedBlockObject {
	var res = NestedBlockObject{}
	switch attr := (*attr).(type) {
	case SingleNestedAttributeBuilder:
		res.Attributes = attr.Attributes
	case ListNestedAttributeBuilder:
		res.Attributes = attr.NestedObject.Attributes
	case MapNestedAttributeBuilder:
		res.Attributes = attr.NestedObject.Attributes
	case SingleNestedBlockBuilder:
		res.Attributes = attr.NestedObject.Attributes
		res.Blocks = attr.NestedObject.Blocks
	case ListNestedBlockBuilder:
		res.Attributes = attr.NestedObject.Attributes
		res.Blocks = attr.NestedObject.Blocks
	default:
		panic(fmt.Errorf("cannot convert to map, attribute is not nested"))
	}

	return res
}

func (s *CustomizableSchema) AddValidator(v any, path ...string) *CustomizableSchema {
	cb := func(attr BaseSchemaBuilder) BaseSchemaBuilder {
		switch a := attr.(type) {
		case BoolAttributeBuilder:
			return a.AddValidator(v.(validator.Bool))
		case Float64AttributeBuilder:
			return a.AddValidator(v.(validator.Float64))
		case Int64AttributeBuilder:
			return a.AddValidator(v.(validator.Int64))
		case ListAttributeBuilder:
			return a.AddValidator(v.(validator.List))
		case ListNestedAttributeBuilder:
			return a.AddValidator(v.(validator.List))
		case MapAttributeBuilder:
			return a.AddValidator(v.(validator.Map))
		case MapNestedAttributeBuilder:
			return a.AddValidator(v.(validator.Map))
		case SingleNestedAttributeBuilder:
			return a.AddValidator(v.(validator.Object))
		case StringAttributeBuilder:
			return a.AddValidator(v.(validator.String))
		default:
			panic(fmt.Errorf("cannot add validator, attribute builder type is invalid: %s. %s", reflect.TypeOf(attr).String(), common.TerraformBugErrorMessage))
		}
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)

	return s
}

func (s *CustomizableSchema) AddPlanModifier(v any, path ...string) *CustomizableSchema {
	cb := func(attr BaseSchemaBuilder) BaseSchemaBuilder {
		switch a := attr.(type) {
		case BoolAttributeBuilder:
			return a.AddPlanModifier(v.(planmodifier.Bool))
		case Float64AttributeBuilder:
			return a.AddPlanModifier(v.(planmodifier.Float64))
		case Int64AttributeBuilder:
			return a.AddPlanModifier(v.(planmodifier.Int64))
		case ListAttributeBuilder:
			return a.AddPlanModifier(v.(planmodifier.List))
		case ListNestedAttributeBuilder:
			return a.AddPlanModifier(v.(planmodifier.List))
		case MapAttributeBuilder:
			return a.AddPlanModifier(v.(planmodifier.Map))
		case MapNestedAttributeBuilder:
			return a.AddPlanModifier(v.(planmodifier.Map))
		case SingleNestedAttributeBuilder:
			return a.AddPlanModifier(v.(planmodifier.Object))
		case StringAttributeBuilder:
			return a.AddPlanModifier(v.(planmodifier.String))
		case ListNestedBlockBuilder:
			return a.AddPlanModifier(v.(planmodifier.List))
		case SingleNestedBlockBuilder:
			return a.AddPlanModifier(v.(planmodifier.Object))
		default:
			panic(fmt.Errorf("cannot add planmodifier, attribute builder type is invalid: %s. %s", reflect.TypeOf(attr).String(), common.TerraformBugErrorMessage))
		}
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)

	return s
}

func (s *CustomizableSchema) SetOptional(path ...string) *CustomizableSchema {
	cb := func(attr BaseSchemaBuilder) BaseSchemaBuilder {
		return attr.SetOptional()
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)

	return s
}

func (s *CustomizableSchema) SetRequired(path ...string) *CustomizableSchema {
	cb := func(attr BaseSchemaBuilder) BaseSchemaBuilder {
		return attr.SetRequired()
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)

	return s
}

func (s *CustomizableSchema) SetSensitive(path ...string) *CustomizableSchema {
	cb := func(attr BaseSchemaBuilder) BaseSchemaBuilder {
		return attr.SetSensitive()
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)
	return s
}

func (s *CustomizableSchema) SetDeprecated(msg string, path ...string) *CustomizableSchema {
	cb := func(attr BaseSchemaBuilder) BaseSchemaBuilder {
		return attr.SetDeprecated(msg)
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)

	return s
}

func (s *CustomizableSchema) SetComputed(path ...string) *CustomizableSchema {
	cb := func(attr BaseSchemaBuilder) BaseSchemaBuilder {
		return attr.SetComputed()
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)
	return s
}

// SetReadOnly sets the schema to be read-only (i.e. computed, non-optional).
// This should be used for fields that are not user-configurable but are returned
// by the platform.
func (s *CustomizableSchema) SetReadOnly(path ...string) *CustomizableSchema {
	cb := func(attr BaseSchemaBuilder) BaseSchemaBuilder {
		return attr.SetReadOnly()
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)

	return s
}

// navigateSchemaWithCallback navigates through schema attributes and executes callback on the target, panics if path does not exist or invalid.
func navigateSchemaWithCallback(s *BaseSchemaBuilder, cb func(BaseSchemaBuilder) BaseSchemaBuilder, path ...string) (BaseSchemaBuilder, error) {
	current_scm := s
	for i, p := range path {
		m := attributeToMap(current_scm)
		m_attr := m.Attributes
		m_block := m.Blocks

		if v, ok := m_attr[p]; ok {
			if i == len(path)-1 {
				new_v := cb(v).(AttributeBuilder)
				m_attr[p] = new_v
				return m_attr[p], nil
			}
			casted_v := v.(BaseSchemaBuilder)
			current_scm = &casted_v
		} else if v, ok := m_block[p]; ok {
			if i == len(path)-1 {
				new_v := cb(v).(BlockBuilder)
				m_block[p] = new_v
				return m_block[p], nil
			}
			casted_v := v.(BaseSchemaBuilder)
			current_scm = &casted_v
		} else {
			return nil, fmt.Errorf("missing key %s", p)
		}

	}
	return nil, fmt.Errorf("path %v is incomplete", path)
}
