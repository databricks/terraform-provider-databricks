package common

import (
	"strings"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type CustomizableSchema struct {
	Schema         *schema.Schema
	path           []string
	isSuppressDiff bool
}

func CustomizeSchemaPath(s map[string]*schema.Schema, path ...string) *CustomizableSchema {
	if len(path) == 0 {
		// Wrapping the input map into a schema when the path is empty.
		// The primary use case for this situation is for adding a new field at the top level.
		wrappedSch := &schema.Schema{
			Elem: &schema.Resource{
				Schema: s,
			},
		}
		return &CustomizableSchema{Schema: wrappedSch}
	}
	sch := MustSchemaPath(s, path...)
	return &CustomizableSchema{Schema: sch, path: path}
}

func (s *CustomizableSchema) SetOptional() *CustomizableSchema {
	s.Schema.Optional = true
	s.Schema.Required = false
	return s
}

func (s *CustomizableSchema) SetComputed() *CustomizableSchema {
	s.Schema.Computed = true
	return s
}

func (s *CustomizableSchema) SetDefault(value any) *CustomizableSchema {
	s.Schema.Default = value
	s.Schema.Optional = true
	s.Schema.Required = false
	return s
}

// SetReadOnly sets the schema to be read-only (i.e. computed, non-optional).
// This should be used for fields that are not user-configurable but are returned
// by the platform.
func (s *CustomizableSchema) SetReadOnly() *CustomizableSchema {
	s.Schema.Optional = false
	s.Schema.Required = false
	s.Schema.MaxItems = 0
	s.Schema.Computed = true
	return s
}

// SetRequired sets the schema to be required.
func (s *CustomizableSchema) SetRequired() *CustomizableSchema {
	s.Schema.Optional = false
	s.Schema.Required = true
	s.Schema.Computed = false
	return s
}

func (s *CustomizableSchema) SetSuppressDiff() *CustomizableSchema {
	s.Schema.DiffSuppressFunc = diffSuppressor(s.path[len(s.path)-1], s.Schema)
	s.isSuppressDiff = true
	if s.Schema.Type == schema.TypeList && s.Schema.MaxItems == 1 {
		// If it is a list with max items = 1, it means the corresponding sdk schema type is a struct or a ptr.
		// In this case we would like to set the diff suppressor for the underlying fields as well.
		resource, ok := s.Schema.Elem.(*schema.Resource)
		if !ok {
			panic("Cannot cast Elem into Resource type.")
		}
		nestedSchema := resource.Schema
		for k, v := range nestedSchema {
			v.DiffSuppressFunc = diffSuppressor(k, v)
		}
	}
	return s
}

func (s *CustomizableSchema) SetCustomSuppressDiff(suppressor func(k, old, new string, d *schema.ResourceData) bool) *CustomizableSchema {
	s.Schema.DiffSuppressFunc = suppressor
	return s
}

func (s *CustomizableSchema) SetSensitive() *CustomizableSchema {
	s.Schema.Sensitive = true
	return s
}

func (s *CustomizableSchema) SetForceNew() *CustomizableSchema {
	s.Schema.ForceNew = true
	return s
}

func (s *CustomizableSchema) SetMaxItems(value int) *CustomizableSchema {
	s.Schema.MaxItems = value
	return s
}

func (s *CustomizableSchema) SetMinItems(value int) *CustomizableSchema {
	s.Schema.MinItems = value
	return s
}

func getPrefixedValue(path []string, value []string) []string {
	var prefix string
	if len(path) != 0 {
		prefix = strings.Join(path, ".") + "."
	} else {
		prefix = ""
	}
	prefixedPaths := make([]string, len(value))
	for i, item := range value {
		prefixedPaths[i] = prefix + item
	}
	return prefixedPaths
}

// Path is used to construct prefixes for cases where the resource is referenced from the resourceProviderRegistry.
// The path should be directly taken from the input of CustomizeSchema function.
func (s *CustomizableSchema) SetConflictsWith(path []string, value []string) *CustomizableSchema {
	if len(value) == 0 {
		panic("SetConflictsWith cannot take in an empty list")
	}
	s.Schema.ConflictsWith = getPrefixedValue(path, value)
	return s
}

func (s *CustomizableSchema) SetExactlyOneOf(path []string, value []string) *CustomizableSchema {
	if len(value) == 0 {
		panic("SetExactlyOneOf cannot take in an empty list")
	}
	s.Schema.ExactlyOneOf = getPrefixedValue(path, value)
	return s
}

func (s *CustomizableSchema) SetAtLeastOneOf(path []string, value []string) *CustomizableSchema {
	if len(value) == 0 {
		panic("SetAtLeastOneOf cannot take in an empty list")
	}
	s.Schema.AtLeastOneOf = getPrefixedValue(path, value)
	return s
}

func (s *CustomizableSchema) SetRequiredWith(path []string, value []string) *CustomizableSchema {
	if len(value) == 0 {
		panic("SetRequiredWith cannot take in an empty list")
	}
	s.Schema.RequiredWith = getPrefixedValue(path, value)
	return s
}

func (s *CustomizableSchema) SetDeprecated(reason string) *CustomizableSchema {
	s.Schema.Deprecated = reason
	return s
}

func (s *CustomizableSchema) SetValidateFunc(validate func(interface{}, string) ([]string, []error)) *CustomizableSchema {
	s.Schema.ValidateFunc = validate
	return s
}

func (s *CustomizableSchema) SetValidateDiagFunc(validate func(interface{}, cty.Path) diag.Diagnostics) *CustomizableSchema {
	s.Schema.ValidateDiagFunc = validate
	return s
}

func (s *CustomizableSchema) AddNewField(key string, newField *schema.Schema) *CustomizableSchema {
	cv, ok := s.Schema.Elem.(*schema.Resource)
	if !ok {
		panic("Cannot add new field, target is not nested resource")
	}
	_, exists := cv.Schema[key]
	if exists {
		panic("Cannot add new field, " + key + " already exists in the schema")
	}
	cv.Schema[key] = newField
	if s.isSuppressDiff {
		newField.DiffSuppressFunc = diffSuppressor(key, newField)
	}
	return s
}
