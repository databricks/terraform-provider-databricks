package util

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/stretchr/testify/assert"
)

type testSliceItem struct {
	SliceItem string `json:"slice_item,omitempty"`
}

type testPtr struct {
	PtrItem string `json:"slice_item,omitempty"`
}

type testStruct struct {
	Integer        int               `json:"integer,omitempty"`
	Float          float64           `json:"float,omitempty"`
	Bool           bool              `json:"bool,omitempty"`
	NonOptional    string            `json:"non_optional"`
	String         string            `json:"string,omitempty"`
	ComputedField  string            `json:"computed_field,omitempty" tf:"some_random_tag,computed"`
	MapField       map[string]string `json:"map_field,omitempty"`
	SliceSetStruct []testSliceItem   `json:"slice_set_struct,omitempty" tf:"slice_set"`
	SliceSetString []string          `json:"slice_set_string,omitempty" tf:"slice_set"`
	PtrItem        *testPtr          `json:"ptr_item,omitempty"`
	StringSlice    []string          `json:"string_slice,omitempty"`
	IntSlice       []int             `json:"int_slice,omitempty"`
	FloatSlice     []float64         `json:"float_slice,omitempty"`
	BoolSlice      []bool            `json:"bool_slice,omitempty"`
}

var scm = StructToSchema(testStruct{}, nil)

var testStructFields = []string{"integer", "float", "non_optional", "string", "computed_field", "map_field",
	"slice_set_struct", "slice_set_string", "ptr_item", "string_slice", "bool", "int_slice", "float_slice",
	"bool_slice"}

var testStructOptionalFields = []string{"integer", "float", "string", "computed_field", "map_field", "slice_set_struct",
	"ptr_item", "slice_set_string", "bool", "int_slice", "float_slice", "bool_slice"}

var testStructRequiredFields = []string{"non_optional"}

var testStructComputedFields = []string{"computed_field"}

var testStructPtrFields = []string{"ptr_item"}

var testStructSliceStructFields = []string{"slice_set_struct"}

var testStructSliceNonStructFields = []string{"slice_set_string", "string_slice", "int_slice", "float_slice",
	"bool_slice"}

func TestStructToSchema_type(t *testing.T) {
	expectedMap := map[string]schema.ValueType{
		"integer":          schema.TypeInt,
		"float":            schema.TypeFloat,
		"bool":             schema.TypeBool,
		"non_optional":     schema.TypeString,
		"string":           schema.TypeString,
		"computed_field":   schema.TypeString,
		"slice_set_struct": schema.TypeSet,
		"map_field":        schema.TypeMap,
		"slice_set_string": schema.TypeSet,
		"ptr_item":         schema.TypeList,
		"string_slice":     schema.TypeList,
		"int_slice":        schema.TypeList,
		"float_slice":      schema.TypeList,
		"bool_slice":       schema.TypeList,
	}
	for field, value := range expectedMap {
		requiredField, ok := scm[field]
		assert.Truef(t, ok, "%s key not found", field)
		assert.Equalf(t, value, requiredField.Type, "%s key should be of type %v but got %v", field,
			value, requiredField.Type)
	}
}

func TestStructToSchema_slice_set_non_struct(t *testing.T) {
	for _, field := range testStructSliceNonStructFields {
		requiredField, ok := scm[field]
		assert.Truef(t, ok, "%s key not found", field)
		assert.Equalf(t, reflect.TypeOf(&schema.Schema{}), reflect.TypeOf(requiredField.Elem),
			"elem in field: %s should be a type of resource", field)
	}
}

func TestStructToSchema_slice_set_struct(t *testing.T) {
	for _, field := range testStructSliceStructFields {
		requiredField, ok := scm[field]
		assert.Truef(t, ok, "%s key not found", field)
		assert.Equalf(t, reflect.TypeOf(&schema.Resource{}), reflect.TypeOf(requiredField.Elem),
			"elem in field: %s should be a type of resource", field)
	}
}

func TestStructToSchema_ptr_set(t *testing.T) {
	for _, field := range testStructPtrFields {
		requiredField, ok := scm[field]
		assert.Truef(t, ok, "%s key not found", field)
		assert.Equalf(t, 1, requiredField.MaxItems, "max items for field: %s should be 1", field)
		assert.Equalf(t, reflect.TypeOf(&schema.Resource{}), reflect.TypeOf(requiredField.Elem),
			"elem in field: %s should be a type of resource", field)
	}
}

func TestStructToSchema_optional_values_set(t *testing.T) {
	for _, field := range testStructOptionalFields {
		requiredField, ok := scm[field]
		assert.Truef(t, ok, "%s key not found", field)
		assert.Truef(t, requiredField.Optional, "optional should be set to true in field: %s", field)
		assert.Falsef(t, requiredField.Required, "required should be set to false in field: %s", field)
	}
}

func TestStructToSchema_computed_values_set(t *testing.T) {
	for _, field := range testStructComputedFields {
		requiredField, ok := scm[field]
		assert.Truef(t, ok, "%s key not found", field)
		assert.Truef(t, requiredField.Computed, "computed should be set to true in field: %s", field)
	}
}

func TestStructToSchema_required_values_set(t *testing.T) {
	for _, field := range testStructRequiredFields {
		requiredField, ok := scm[field]
		assert.Truef(t, ok, "%s key not found", field)
		assert.Truef(t, requiredField.Required, "required should be set to true in field: %s", field)
		assert.Falsef(t, requiredField.Optional, "optional should be set to false in field: %s", field)
	}
}

func TestStructToSchema_base_values_are_set(t *testing.T) {
	for _, field := range testStructFields {
		_, ok := scm[field]
		assert.Truef(t, ok, "%s key not found", field)
	}
}
