package common

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type CustomizableSchema struct {
	Schema         *schema.Schema
	path           []string
	isSuppressDiff bool
	context        schemaPathContext
}

func (s *CustomizableSchema) pathContainsMultipleItemsList() bool {
	schemaPath := s.context.schemaPath
	for _, scm := range schemaPath {
		if scm.Type == schema.TypeList && scm.MaxItems != 1 {
			return true
		}
	}
	return false
}

// Used to get the prefix path for functions like ConflictsWith, by joining `path` in SchemaPathContext.
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

func (s *CustomizableSchema) SchemaPath(path ...string) *CustomizableSchema {
	sch := MustSchemaPath(s.GetSchemaMap(), path...)
	return &CustomizableSchema{Schema: sch, path: path, context: s.context}
}

func (s *CustomizableSchema) GetSchemaMap() map[string]*schema.Schema {
	if s.Schema.Elem == nil {
		panic("Elem of Schema field for CustomizableSchema is nil.")
	}
	schemaResource, ok := s.Schema.Elem.(*schema.Resource)
	if !ok {
		panic("Elem of Schema field for CustomizableSchema is not a *schema.Resource.")
	}
	return schemaResource.Schema
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

func (s *CustomizableSchema) SetSliceSet() *CustomizableSchema {
	s.Schema.Type = schema.TypeSet
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

// SetSuppressDiffWithDefault suppresses the diff if the
// new value (ie value from HCL config) is not set and
// the old value (ie value from state / platform) is equal to the default value.
//
// Often Databricks HTTP APIs will return values for fields that were not set by
// the author in their terraform configuration. This function allows us to suppress
// the diff in these cases.
func (s *CustomizableSchema) SetSuppressDiffWithDefault(dv any) *CustomizableSchema {
	primitiveTypes := []schema.ValueType{schema.TypeBool, schema.TypeString, schema.TypeInt, schema.TypeFloat}
	if !slices.Contains(primitiveTypes, s.Schema.Type) {
		panic(fmt.Errorf("expected primitive type, got: %s", s.Schema.Type))
	}

	// Get zero value for the schema type
	zero := fmt.Sprintf("%v", s.Schema.Type.Zero())

	// Get string representation of the default value
	sv := fmt.Sprintf("%v", dv)

	// Suppress diff if the new value (ie value from HCL config) is not set and
	// the old value (ie value from state / platform) is equal to the default value.
	s.Schema.DiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool {
		if new == zero && old == sv {
			return true
		}
		return false
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

func (s *CustomizableSchema) SetConflictsWith(value []string) *CustomizableSchema {
	if len(value) == 0 {
		panic("SetConflictsWith cannot take in an empty list")
	}
	if s.pathContainsMultipleItemsList() {
		log.Printf("[DEBUG] ConflictsWith skipped for %v, path contains TypeList block with MaxItems not equal to 1", getPrefixedValue(s.context.path, value))
		return s
	}
	s.Schema.ConflictsWith = getPrefixedValue(s.context.path, value)
	return s
}

func (s *CustomizableSchema) SetExactlyOneOf(value []string) *CustomizableSchema {
	if len(value) == 0 {
		panic("SetExactlyOneOf cannot take in an empty list")
	}
	if s.pathContainsMultipleItemsList() {
		log.Printf("[DEBUG] ExactlyOneOf skipped for %v, path contains TypeList block with MaxItems not equal to 1", getPrefixedValue(s.context.path, value))
		return s
	}
	s.Schema.ExactlyOneOf = getPrefixedValue(s.context.path, value)
	return s
}

func (s *CustomizableSchema) SetAtLeastOneOf(value []string) *CustomizableSchema {
	if len(value) == 0 {
		panic("SetAtLeastOneOf cannot take in an empty list")
	}
	if s.pathContainsMultipleItemsList() {
		log.Printf("[DEBUG] AtLeastOneOf skipped for %v, path contains TypeList block with MaxItems not equal to 1", getPrefixedValue(s.context.path, value))
		return s
	}
	s.Schema.AtLeastOneOf = getPrefixedValue(s.context.path, value)
	return s
}

func (s *CustomizableSchema) SetRequiredWith(value []string) *CustomizableSchema {
	if len(value) == 0 {
		panic("SetRequiredWith cannot take in an empty list")
	}
	if s.pathContainsMultipleItemsList() {
		log.Printf("[DEBUG] SetRequiredWith skipped for %v, path contains TypeList block with MaxItems not equal to 1", getPrefixedValue(s.context.path, value))
		return s
	}
	s.Schema.RequiredWith = getPrefixedValue(s.context.path, value)
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
	scm := s.GetSchemaMap()
	_, exists := scm[key]
	if exists {
		panic("Cannot add new field, " + key + " already exists in the schema")
	}
	scm[key] = newField
	if s.isSuppressDiff {
		newField.DiffSuppressFunc = diffSuppressor(key, newField)
	}
	return s
}

func (s *CustomizableSchema) RemoveField(key string) *CustomizableSchema {
	scm := s.GetSchemaMap()
	_, exists := scm[key]
	if !exists {
		panic("Cannot remove new field, " + key + " does not exist in the schema")
	}
	delete(scm, key)
	return s
}
