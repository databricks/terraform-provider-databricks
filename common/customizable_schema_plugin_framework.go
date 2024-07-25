package common

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
)

type CustomizableSchemaPluginFramework struct {
	attr schema.Attribute
}

func ConstructCustomizableSchema(attributes map[string]schema.Attribute) *CustomizableSchemaPluginFramework {
	attr := schema.Attribute(schema.SingleNestedAttribute{Attributes: attributes})
	return &CustomizableSchemaPluginFramework{attr: attr}
}

func (s *CustomizableSchemaPluginFramework) SchemaPath(path ...string) *CustomizableSchemaPluginFramework {
	attr, err := navigateSchema(s.attr, path...)
	if err != nil {
		panic(err)
	}

	return &CustomizableSchemaPluginFramework{attr}
}

// Converts CustomizableSchema into a map from string to Attribute.
func (s *CustomizableSchemaPluginFramework) ToAttributeMap() map[string]schema.Attribute {
	return attributeToMap(s.attr)
}

func attributeToMap(attr schema.Attribute) map[string]schema.Attribute {
	var m map[string]schema.Attribute

	switch attr := attr.(type) {
	case schema.SingleNestedAttribute:
		m = attr.Attributes
	case schema.ListNestedAttribute:
		m = attr.NestedObject.Attributes
	case schema.MapNestedAttribute:
		m = attr.NestedObject.Attributes
	default:
		panic(fmt.Errorf("cannot convert to map, attribute is not nested"))
	}

	return m
}

func (s *CustomizableSchemaPluginFramework) AddNewField(key string, newField schema.Attribute) *CustomizableSchemaPluginFramework {
	switch attr := s.attr.(type) {
	case schema.SingleNestedAttribute:
		_, exists := attr.Attributes[key]
		if exists {
			panic("Cannot add new field, " + key + " already exists in the schema")
		}
		attr.Attributes[key] = newField
		attr.Required = true
		s.attr = attr
	case schema.ListNestedAttribute:
		_, exists := attr.NestedObject.Attributes[key]
		if exists {
			panic("Cannot add new field, " + key + " already exists in the schema")
		}
		attr.NestedObject.Attributes[key] = newField
		s.attr = attr
	case schema.MapNestedAttribute:
		_, exists := attr.NestedObject.Attributes[key]
		if exists {
			panic("Cannot add new field, " + key + " already exists in the schema")
		}
		attr.NestedObject.Attributes[key] = newField
		s.attr = attr
	default:
		panic("attribute is not nested, cannot add field")
	}

	return s
}

func (s *CustomizableSchemaPluginFramework) RemoveField(key string) *CustomizableSchemaPluginFramework {
	switch attr := s.attr.(type) {
	case schema.SingleNestedAttribute:
		_, exists := attr.Attributes[key]
		if !exists {
			panic("Cannot remove field, " + key + " does not exist in the schema")
		}
		delete(attr.Attributes, key)
		s.attr = attr
	case schema.ListNestedAttribute:
		_, exists := attr.NestedObject.Attributes[key]
		if !exists {
			panic("Cannot remove field, " + key + " does not exist in the schema")
		}
		delete(attr.NestedObject.Attributes, key)
		s.attr = attr
	case schema.MapNestedAttribute:
		_, exists := attr.NestedObject.Attributes[key]
		if !exists {
			panic("Cannot remove field, " + key + " does not exist in the schema")
		}
		delete(attr.NestedObject.Attributes, key)
		s.attr = attr
	default:
		panic("attribute is not nested, cannot add field")
	}

	return s
}

// Given a attribute map, navigate through the given path, panics if the path is not valid.
func MustSchemaAttributePath(attrs map[string]schema.Attribute, path ...string) schema.Attribute {
	return ConstructCustomizableSchema(attrs).SchemaPath(path...).attr
}

// Helper function for navigating through schema attributes, panics if path does not exist or invalid.
func navigateSchema(s schema.Attribute, path ...string) (schema.Attribute, error) {
	cs := s
	for i, p := range path {
		m := attributeToMap(cs)

		v, ok := m[p]
		if !ok {
			return nil, fmt.Errorf("missing key %s", p)
		}
		if i == len(path)-1 {
			return v, nil
		}
		cs = v
	}
	return nil, fmt.Errorf("path %v is incomplete", path)
}
