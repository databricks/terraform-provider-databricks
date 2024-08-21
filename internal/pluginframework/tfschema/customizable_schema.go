package tfschema

import (
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// CustomizableSchema is a wrapper struct on top of AttributeBuilder that can be used to navigate through nested schema add customizations.
type CustomizableSchema struct {
	attr AttributeBuilder
}

// ConstructCustomizableSchema constructs a CustomizableSchema given a map from string to AttributeBuilder.
func ConstructCustomizableSchema(attributes map[string]AttributeBuilder) *CustomizableSchema {
	attr := AttributeBuilder(SingleNestedAttributeBuilder{Attributes: attributes})
	return &CustomizableSchema{attr: attr}
}

// ToAttributeMap converts CustomizableSchema into a map from string to Attribute.
func (s *CustomizableSchema) ToAttributeMap() map[string]AttributeBuilder {
	return attributeToMap(&s.attr)
}

// attributeToMap converts AttributeBuilder into a map from string to AttributeBuilder.
func attributeToMap(attr *AttributeBuilder) map[string]AttributeBuilder {
	var m map[string]AttributeBuilder
	switch attr := (*attr).(type) {
	case SingleNestedAttributeBuilder:
		m = attr.Attributes
	case ListNestedAttributeBuilder:
		m = attr.NestedObject.Attributes
	case MapNestedAttributeBuilder:
		m = attr.NestedObject.Attributes
	default:
		panic(fmt.Errorf("cannot convert to map, attribute is not nested"))
	}

	return m
}

func (s *CustomizableSchema) AddValidator(v any, path ...string) *CustomizableSchema {
	cb := func(attr AttributeBuilder) AttributeBuilder {
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
			panic(fmt.Errorf("cannot add validator, attribute builder type is invalid. %s", common.TerraformBugErrorMessage))
		}
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)

	return s
}

func (s *CustomizableSchema) SetOptional(path ...string) *CustomizableSchema {
	cb := func(attr AttributeBuilder) AttributeBuilder {
		return attr.SetOptional()
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)

	return s
}

func (s *CustomizableSchema) SetRequired(path ...string) *CustomizableSchema {
	cb := func(attr AttributeBuilder) AttributeBuilder {
		return attr.SetRequired()
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)

	return s
}

func (s *CustomizableSchema) SetSensitive(path ...string) *CustomizableSchema {
	cb := func(attr AttributeBuilder) AttributeBuilder {
		return attr.SetSensitive()
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)
	return s
}

func (s *CustomizableSchema) SetDeprecated(msg string, path ...string) *CustomizableSchema {
	cb := func(attr AttributeBuilder) AttributeBuilder {
		return attr.SetDeprecated(msg)
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)

	return s
}

func (s *CustomizableSchema) SetComputed(path ...string) *CustomizableSchema {
	cb := func(attr AttributeBuilder) AttributeBuilder {
		return attr.SetComputed()
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)
	return s
}

// SetReadOnly sets the schema to be read-only (i.e. computed, non-optional).
// This should be used for fields that are not user-configurable but are returned
// by the platform.
func (s *CustomizableSchema) SetReadOnly(path ...string) *CustomizableSchema {
	cb := func(attr AttributeBuilder) AttributeBuilder {
		return attr.SetReadOnly()
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)

	return s
}

// navigateSchemaWithCallback navigates through schema attributes and executes callback on the target, panics if path does not exist or invalid.
func navigateSchemaWithCallback(s *AttributeBuilder, cb func(AttributeBuilder) AttributeBuilder, path ...string) (AttributeBuilder, error) {
	current_scm := s
	for i, p := range path {
		m := attributeToMap(current_scm)

		v, ok := m[p]
		if !ok {
			return nil, fmt.Errorf("missing key %s", p)
		}

		if i == len(path)-1 {
			m[p] = cb(v)
			return m[p], nil
		}
		current_scm = &v
	}
	return nil, fmt.Errorf("path %v is incomplete", path)
}
