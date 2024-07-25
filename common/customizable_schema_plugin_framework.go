package common

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type CustomizableSchemaPluginFramework struct {
	attr schema.Attribute
}

func ConstructCustomizableSchema(attributes map[string]schema.Attribute) *CustomizableSchemaPluginFramework {
	attr := schema.Attribute(schema.SingleNestedAttribute{Attributes: attributes})
	return &CustomizableSchemaPluginFramework{attr: attr}
}

// Converts CustomizableSchema into a map from string to Attribute.
func (s *CustomizableSchemaPluginFramework) ToAttributeMap() map[string]schema.Attribute {
	return attributeToMap(&s.attr)
}

func attributeToMap(attr *schema.Attribute) map[string]schema.Attribute {
	var m map[string]schema.Attribute
	switch attr := (*attr).(type) {
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

func (s *CustomizableSchemaPluginFramework) AddNewField(key string, newField schema.Attribute, path ...string) *CustomizableSchemaPluginFramework {
	cb := func(a schema.Attribute) schema.Attribute {
		switch attr := a.(type) {
		case schema.SingleNestedAttribute:
			_, exists := attr.Attributes[key]
			if exists {
				panic("Cannot add new field, " + key + " already exists in the schema")
			}
			attr.Attributes[key] = newField
			return attr
		case schema.ListNestedAttribute:
			_, exists := attr.NestedObject.Attributes[key]
			if exists {
				panic("Cannot add new field, " + key + " already exists in the schema")
			}
			attr.NestedObject.Attributes[key] = newField
			return attr
		case schema.MapNestedAttribute:
			_, exists := attr.NestedObject.Attributes[key]
			if exists {
				panic("Cannot add new field, " + key + " already exists in the schema")
			}
			attr.NestedObject.Attributes[key] = newField
			return attr
		default:
			panic("attribute is not nested, cannot add field")
		}
	}

	if len(path) == 0 {
		s.attr = cb(s.attr)
	} else {
		navigateSchemaWithCallback(&s.attr, cb, path...)
	}
	return s
}

func (s *CustomizableSchemaPluginFramework) RemoveField(key string, path ...string) *CustomizableSchemaPluginFramework {
	cb := func(a schema.Attribute) schema.Attribute {
		switch attr := a.(type) {
		case schema.SingleNestedAttribute:
			_, exists := attr.Attributes[key]
			if !exists {
				panic("Cannot remove field, " + key + " does not exist in the schema")
			}
			delete(attr.Attributes, key)
			return attr
		case schema.ListNestedAttribute:
			_, exists := attr.NestedObject.Attributes[key]
			if !exists {
				panic("Cannot remove field, " + key + " does not exist in the schema")
			}
			delete(attr.NestedObject.Attributes, key)
			return attr
		case schema.MapNestedAttribute:
			_, exists := attr.NestedObject.Attributes[key]
			if !exists {
				panic("Cannot remove field, " + key + " does not exist in the schema")
			}
			delete(attr.NestedObject.Attributes, key)
			return attr
		default:
			panic("attribute is not nested, cannot add field")
		}
	}

	if len(path) == 0 {
		s.attr = cb(s.attr)
	} else {
		navigateSchemaWithCallback(&s.attr, cb, path...)
	}
	return s
}

func (s *CustomizableSchemaPluginFramework) AddValidator(v any, path ...string) *CustomizableSchemaPluginFramework {
	cb := func(a schema.Attribute) schema.Attribute {
		switch attr := a.(type) {
		case schema.SingleNestedAttribute:
			attr.Validators = append(attr.Validators, v.(validator.Object))
			return attr
		case schema.ListNestedAttribute:
			attr.Validators = append(attr.Validators, v.(validator.List))
			return attr
		case schema.MapNestedAttribute:
			attr.Validators = append(attr.Validators, v.(validator.Map))
			return attr
		case schema.BoolAttribute:
			attr.Validators = append(attr.Validators, v.(validator.Bool))
			return attr
		case schema.Float64Attribute:
			attr.Validators = append(attr.Validators, v.(validator.Float64))
			return attr
		case schema.StringAttribute:
			attr.Validators = append(attr.Validators, v.(validator.String))
			return attr
		case schema.Int64Attribute:
			attr.Validators = append(attr.Validators, v.(validator.Int64))
			return attr
		case schema.ListAttribute:
			attr.Validators = append(attr.Validators, v.(validator.List))
			return attr
		case schema.MapAttribute:
			attr.Validators = append(attr.Validators, v.(validator.Map))
			return attr
		default:
			panic(fmt.Sprintf("Unsupported type %T", s.attr))
		}
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)

	return s
}

func (s *CustomizableSchemaPluginFramework) SetOptional(path ...string) *CustomizableSchemaPluginFramework {
	cb := func(a schema.Attribute) schema.Attribute {
		switch attr := a.(type) {
		case schema.SingleNestedAttribute:
			attr.Optional = true
			attr.Required = false
			return attr
		case schema.ListNestedAttribute:
			attr.Optional = true
			attr.Required = false
			return attr
		case schema.MapNestedAttribute:
			attr.Optional = true
			attr.Required = false
			return attr
		case schema.BoolAttribute:
			attr.Optional = true
			attr.Required = false
			return attr
		case schema.Float64Attribute:
			attr.Optional = true
			attr.Required = false
			return attr
		case schema.StringAttribute:
			attr.Optional = true
			attr.Required = false
			return attr
		case schema.Int64Attribute:
			attr.Optional = true
			attr.Required = false
			return attr
		case schema.ListAttribute:
			attr.Optional = true
			attr.Required = false
			return attr
		case schema.MapAttribute:
			attr.Optional = true
			attr.Required = false
			return attr
		default:
			panic(fmt.Sprintf("Unsupported type %T", s.attr))
		}
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)

	return s
}

func (s *CustomizableSchemaPluginFramework) SetRequired(path ...string) *CustomizableSchemaPluginFramework {
	cb := func(a schema.Attribute) schema.Attribute {
		switch attr := a.(type) {
		case schema.SingleNestedAttribute:
			attr.Optional = false
			attr.Required = true
			return attr
		case schema.ListNestedAttribute:
			attr.Optional = false
			attr.Required = true
			return attr
		case schema.MapNestedAttribute:
			attr.Optional = false
			attr.Required = true
			return attr
		case schema.BoolAttribute:
			attr.Optional = false
			attr.Required = true
			return attr
		case schema.Float64Attribute:
			attr.Optional = false
			attr.Required = true
			return attr
		case schema.StringAttribute:
			attr.Optional = false
			attr.Required = true
			return attr
		case schema.Int64Attribute:
			attr.Optional = false
			attr.Required = true
			return attr
		case schema.ListAttribute:
			attr.Optional = false
			attr.Required = true
			return attr
		case schema.MapAttribute:
			attr.Optional = false
			attr.Required = true
			return attr
		default:
			panic(fmt.Sprintf("Unsupported type %T", s.attr))
		}
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)

	return s
}

func (s *CustomizableSchemaPluginFramework) SetSensitive(path ...string) *CustomizableSchemaPluginFramework {
	cb := func(a schema.Attribute) schema.Attribute {
		switch attr := a.(type) {
		case schema.SingleNestedAttribute:
			attr.Sensitive = true
			return attr
		case schema.ListNestedAttribute:
			attr.Sensitive = true
			return attr
		case schema.MapNestedAttribute:
			attr.Sensitive = true
			return attr
		case schema.BoolAttribute:
			attr.Sensitive = true
			return attr
		case schema.Float64Attribute:
			attr.Sensitive = true
			return attr
		case schema.StringAttribute:
			attr.Sensitive = true
			return attr
		case schema.Int64Attribute:
			attr.Sensitive = true
			return attr
		case schema.ListAttribute:
			attr.Sensitive = true
			return attr
		case schema.MapAttribute:
			attr.Sensitive = true
			return attr
		default:
			panic(fmt.Sprintf("Unsupported type %T", s.attr))
		}
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)
	return s
}

// Given a attribute map, navigate through the given path, panics if the path is not valid.
func MustSchemaAttributePath(attrs map[string]schema.Attribute, path ...string) schema.Attribute {
	attr := ConstructCustomizableSchema(attrs).attr

	res, err := navigateSchema(&attr, path...)
	if err != nil {
		panic(err)
	}

	return res
}

// Helper function for navigating through schema attributes, panics if path does not exist or invalid.
func navigateSchema(s *schema.Attribute, path ...string) (schema.Attribute, error) {
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
		cs = &v
	}
	return nil, fmt.Errorf("path %v is incomplete", path)
}

// Helper function for navigating through schema attributes, panics if path does not exist or invalid.
func navigateSchemaWithCallback(s *schema.Attribute, cb func(schema.Attribute) schema.Attribute, path ...string) (schema.Attribute, error) {
	cs := s
	for i, p := range path {
		m := attributeToMap(cs)

		v, ok := m[p]
		if !ok {
			return nil, fmt.Errorf("missing key %s", p)
		}

		if i == len(path)-1 {
			m[p] = cb(v)
			return m[p], nil
		}
		cs = &v
	}
	return nil, fmt.Errorf("path %v is incomplete", path)
}
