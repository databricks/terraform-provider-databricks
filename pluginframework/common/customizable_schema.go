package pluginframework

import (
	"fmt"
)

type CustomizableSchemaPluginFramework struct {
	attr Attribute
}

func ConstructCustomizableSchema(attributes map[string]Attribute) *CustomizableSchemaPluginFramework {
	attr := Attribute(SingleNestedAttribute{Attributes: attributes})
	return &CustomizableSchemaPluginFramework{attr: attr}
}

// Converts CustomizableSchema into a map from string to Attribute.
func (s *CustomizableSchemaPluginFramework) ToAttributeMap() map[string]Attribute {
	return attributeToMap(&s.attr)
}

func attributeToMap(attr *Attribute) map[string]Attribute {
	var m map[string]Attribute
	switch attr := (*attr).(type) {
	case SingleNestedAttribute:
		m = attr.Attributes
	case ListNestedAttribute:
		m = attr.NestedObject.Attributes
	case MapNestedAttribute:
		m = attr.NestedObject.Attributes
	default:
		panic(fmt.Errorf("cannot convert to map, attribute is not nested"))
	}

	return m
}

func (s *CustomizableSchemaPluginFramework) AddValidator(v any, path ...string) *CustomizableSchemaPluginFramework {
	cb := func(attr Attribute) Attribute {
		return attr.AddValidators(v)
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)

	return s
}

func (s *CustomizableSchemaPluginFramework) SetOptional(path ...string) *CustomizableSchemaPluginFramework {
	cb := func(attr Attribute) Attribute {
		return attr.SetOptional()
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)

	return s
}

func (s *CustomizableSchemaPluginFramework) SetRequired(path ...string) *CustomizableSchemaPluginFramework {
	cb := func(attr Attribute) Attribute {
		return attr.SetRequired()
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)

	return s
}

func (s *CustomizableSchemaPluginFramework) SetSensitive(path ...string) *CustomizableSchemaPluginFramework {
	cb := func(attr Attribute) Attribute {
		return attr.SetSensitive()
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)
	return s
}

func (s *CustomizableSchemaPluginFramework) SetDeprecated(msg string, path ...string) *CustomizableSchemaPluginFramework {
	cb := func(attr Attribute) Attribute {
		return attr.SetDeprecated(msg)
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)

	return s
}

func (s *CustomizableSchemaPluginFramework) SetComputed(path ...string) *CustomizableSchemaPluginFramework {
	cb := func(attr Attribute) Attribute {
		return attr.SetComputed()
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)
	return s
}

// SetReadOnly sets the schema to be read-only (i.e. computed, non-optional).
// This should be used for fields that are not user-configurable but are returned
// by the platform.
func (s *CustomizableSchemaPluginFramework) SetReadOnly(path ...string) *CustomizableSchemaPluginFramework {
	cb := func(attr Attribute) Attribute {
		return attr.SetReadOnly()
	}

	navigateSchemaWithCallback(&s.attr, cb, path...)

	return s
}

// Helper function for navigating through schema attributes, panics if path does not exist or invalid.
func navigateSchemaWithCallback(s *Attribute, cb func(Attribute) Attribute, path ...string) (Attribute, error) {
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
