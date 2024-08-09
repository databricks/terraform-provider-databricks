package common

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

type CustomizableSchemaPluginFrameworkDataSource struct {
	attr schema.Attribute
}

func ConstructCustomizableSchemaDataSource(attributes map[string]schema.Attribute) *CustomizableSchemaPluginFrameworkDataSource {
	attr := schema.Attribute(schema.SingleNestedAttribute{Attributes: attributes})
	return &CustomizableSchemaPluginFrameworkDataSource{attr: attr}
}

// Converts CustomizableSchema into a map from string to Attribute.
func (s *CustomizableSchemaPluginFrameworkDataSource) ToAttributeMap() map[string]schema.Attribute {
	return attributeToMapDataSource(&s.attr)
}

func attributeToMapDataSource(attr *schema.Attribute) map[string]schema.Attribute {
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

func (s *CustomizableSchemaPluginFrameworkDataSource) AddNewField(key string, newField schema.Attribute, path ...string) *CustomizableSchemaPluginFrameworkDataSource {
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
		navigateSchemaWithCallbackDataSource(&s.attr, cb, path...)
	}
	return s
}

func (s *CustomizableSchemaPluginFrameworkDataSource) RemoveField(key string, path ...string) *CustomizableSchemaPluginFrameworkDataSource {
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
		navigateSchemaWithCallbackDataSource(&s.attr, cb, path...)
	}
	return s
}

func (s *CustomizableSchemaPluginFrameworkDataSource) AddValidator(v any, path ...string) *CustomizableSchemaPluginFrameworkDataSource {
	cb := func(attr schema.Attribute) schema.Attribute {
		val := reflect.ValueOf(attr)

		// Make a copy of the existing attr.
		newAttr := reflect.New(val.Type()).Elem()
		newAttr.Set(val)
		val = newAttr

		field := val.FieldByName("Validators")
		if !field.IsValid() {
			panic(fmt.Sprintf("Validators field not found in %T", attr))
		}
		if field.Kind() != reflect.Slice {
			panic(fmt.Sprintf("Validators field is not a slice in %T", attr))
		}
		if !field.CanSet() {
			panic(fmt.Sprintf("Validators field cannot be set in %T", attr))
		}

		elemType := field.Type().Elem()
		value := reflect.ValueOf(v)

		if !value.Type().AssignableTo(elemType) {
			panic(fmt.Sprintf("Value of type %T is not assignable to slice of %s", v, elemType))
		}

		// Append the value
		newSlice := reflect.Append(field, value)
		field.Set(newSlice)

		return val.Interface().(schema.Attribute)
	}

	navigateSchemaWithCallbackDataSource(&s.attr, cb, path...)

	return s
}

func (s *CustomizableSchemaPluginFrameworkDataSource) SetOptional(path ...string) *CustomizableSchemaPluginFrameworkDataSource {
	cb := func(attr schema.Attribute) schema.Attribute {
		// Get the concrete value stored in the interface
		v := reflect.ValueOf(attr)

		// Make a new addressable value and copy the original value into it
		newAttr := reflect.New(v.Type()).Elem()
		newAttr.Set(v)
		v = newAttr

		field := v.FieldByName("Required")
		if field.IsValid() && field.CanSet() {
			if field.Kind() == reflect.Bool {
				field.SetBool(false)
			} else {
				panic(fmt.Sprintf("Required is not a bool field in %T", attr))
			}
		} else {
			panic(fmt.Sprintf("Required field not found or cannot be set in %T", attr))
		}

		field = v.FieldByName("Optional")
		if field.IsValid() && field.CanSet() {
			if field.Kind() == reflect.Bool {
				field.SetBool(true)
			} else {
				panic(fmt.Sprintf("Optional is not a bool field in %T", attr))
			}
		} else {
			panic(fmt.Sprintf("Optional field not found or cannot be set in %T", attr))
		}

		return v.Interface().(schema.Attribute)
	}

	navigateSchemaWithCallbackDataSource(&s.attr, cb, path...)

	return s
}

func (s *CustomizableSchemaPluginFrameworkDataSource) SetRequired(path ...string) *CustomizableSchemaPluginFrameworkDataSource {
	cb := func(attr schema.Attribute) schema.Attribute {
		// Get the concrete value stored in the interface
		v := reflect.ValueOf(attr)

		// Make a new addressable value and copy the original value into it
		newAttr := reflect.New(v.Type()).Elem()
		newAttr.Set(v)
		v = newAttr

		field := v.FieldByName("Required")
		if field.IsValid() && field.CanSet() {
			if field.Kind() == reflect.Bool {
				field.SetBool(true)
			} else {
				panic(fmt.Sprintf("Required is not a bool field in %T", attr))
			}
		} else {
			panic(fmt.Sprintf("Required field not found or cannot be set in %T", attr))
		}

		field = v.FieldByName("Optional")
		if field.IsValid() && field.CanSet() {
			if field.Kind() == reflect.Bool {
				field.SetBool(false)
			} else {
				panic(fmt.Sprintf("Optional is not a bool field in %T", attr))
			}
		} else {
			panic(fmt.Sprintf("Optional field not found or cannot be set in %T", attr))
		}

		return v.Interface().(schema.Attribute)
	}

	navigateSchemaWithCallbackDataSource(&s.attr, cb, path...)

	return s
}

func (s *CustomizableSchemaPluginFrameworkDataSource) SetSensitive(path ...string) *CustomizableSchemaPluginFrameworkDataSource {
	cb := func(attr schema.Attribute) schema.Attribute {
		// Get the concrete value stored in the interface
		v := reflect.ValueOf(attr)

		// Make a new addressable value and copy the original value into it
		newAttr := reflect.New(v.Type()).Elem()
		newAttr.Set(v)
		v = newAttr

		field := v.FieldByName("Sensitive")
		if field.IsValid() && field.CanSet() {
			if field.Kind() == reflect.Bool {
				field.SetBool(true)
			} else {
				panic(fmt.Sprintf("Sensitive is not a bool field in %T", attr))
			}
		} else {
			panic(fmt.Sprintf("Sensitive field not found or cannot be set in %T", attr))
		}

		return v.Interface().(schema.Attribute)
	}

	navigateSchemaWithCallbackDataSource(&s.attr, cb, path...)
	return s
}

func (s *CustomizableSchemaPluginFrameworkDataSource) SetDeprecated(msg string, path ...string) *CustomizableSchemaPluginFrameworkDataSource {
	cb := func(attr schema.Attribute) schema.Attribute {
		// Get the concrete value stored in the interface
		v := reflect.ValueOf(attr)

		// Make a new addressable value and copy the original value into it
		newAttr := reflect.New(v.Type()).Elem()
		newAttr.Set(v)
		v = newAttr

		field := v.FieldByName("DeprecationMessage")
		if field.IsValid() && field.CanSet() {
			if field.Kind() == reflect.String {
				field.SetString(msg)
			} else {
				panic(fmt.Sprintf("DeprecationMessage is not a string field in %T", attr))
			}
		} else {
			panic(fmt.Sprintf("DeprecationMessage field not found or cannot be set in %T", attr))
		}

		return v.Interface().(schema.Attribute)
	}

	navigateSchemaWithCallbackDataSource(&s.attr, cb, path...)

	return s
}

func (s *CustomizableSchemaPluginFrameworkDataSource) SetComputed(path ...string) *CustomizableSchemaPluginFrameworkDataSource {
	cb := func(attr schema.Attribute) schema.Attribute {
		// Get the concrete value stored in the interface
		v := reflect.ValueOf(attr)

		// Make a new addressable value and copy the original value into it
		newAttr := reflect.New(v.Type()).Elem()
		newAttr.Set(v)
		v = newAttr

		field := v.FieldByName("Computed")
		if field.IsValid() && field.CanSet() {
			if field.Kind() == reflect.Bool {
				field.SetBool(true)
			} else {
				panic(fmt.Sprintf("Computed is not a bool field in %T", attr))
			}
		} else {
			panic(fmt.Sprintf("Computed field not found or cannot be set in %T", attr))
		}

		return v.Interface().(schema.Attribute)
	}

	navigateSchemaWithCallbackDataSource(&s.attr, cb, path...)
	return s
}

// SetReadOnly sets the schema to be read-only (i.e. computed, non-optional).
// This should be used for fields that are not user-configurable but are returned
// by the platform.
func (s *CustomizableSchemaPluginFrameworkDataSource) SetReadOnly(path ...string) *CustomizableSchemaPluginFrameworkDataSource {
	cb := func(attr schema.Attribute) schema.Attribute {
		// Get the concrete value stored in the interface
		v := reflect.ValueOf(attr)

		// Make a new addressable value and copy the original value into it
		newAttr := reflect.New(v.Type()).Elem()
		newAttr.Set(v)
		v = newAttr

		field := v.FieldByName("Computed")
		if field.IsValid() && field.CanSet() {
			if field.Kind() == reflect.Bool {
				field.SetBool(true)
			} else {
				panic(fmt.Sprintf("Computed is not a bool field in %T", attr))
			}
		} else {
			panic(fmt.Sprintf("Computed field not found or cannot be set in %T", attr))
		}

		field = v.FieldByName("Optional")
		if field.IsValid() && field.CanSet() {
			if field.Kind() == reflect.Bool {
				field.SetBool(false)
			} else {
				panic(fmt.Sprintf("Optional is not a bool field in %T", attr))
			}
		} else {
			panic(fmt.Sprintf("Optional field not found or cannot be set in %T", attr))
		}

		field = v.FieldByName("Required")
		if field.IsValid() && field.CanSet() {
			if field.Kind() == reflect.Bool {
				field.SetBool(false)
			} else {
				panic(fmt.Sprintf("Required is not a bool field in %T", attr))
			}
		} else {
			panic(fmt.Sprintf("Required field not found or cannot be set in %T", attr))
		}

		return v.Interface().(schema.Attribute)
	}

	navigateSchemaWithCallbackDataSource(&s.attr, cb, path...)

	return s
}

// Given a attribute map, navigate through the given path, panics if the path is not valid.
func MustSchemaAttributePathDataSource(attrs map[string]schema.Attribute, path ...string) schema.Attribute {
	attr := ConstructCustomizableSchemaDataSource(attrs).attr

	res, err := navigateSchemaDataSource(&attr, path...)
	if err != nil {
		panic(err)
	}

	return res
}

// Helper function for navigating through schema attributes, panics if path does not exist or invalid.
func navigateSchemaDataSource(s *schema.Attribute, path ...string) (schema.Attribute, error) {
	current_scm := s
	for i, p := range path {
		m := attributeToMapDataSource(current_scm)

		v, ok := m[p]
		if !ok {
			return nil, fmt.Errorf("missing key %s", p)
		}

		if i == len(path)-1 {
			return v, nil
		}
		current_scm = &v
	}
	return nil, fmt.Errorf("path %v is incomplete", path)
}

// Helper function for navigating through schema attributes, panics if path does not exist or invalid.
func navigateSchemaWithCallbackDataSource(s *schema.Attribute, cb func(schema.Attribute) schema.Attribute, path ...string) (schema.Attribute, error) {
	current_scm := s
	for i, p := range path {
		m := attributeToMapDataSource(current_scm)

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
