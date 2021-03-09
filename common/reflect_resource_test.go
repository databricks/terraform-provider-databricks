package common

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

func TestReflectKind(t *testing.T) {
	assert.Equal(t, "Bool", reflectKind(1))
}

func TestSchemaPath(t *testing.T) {
	_, err := SchemaPath(map[string]*schema.Schema{}, "x", "y", "z")
	assert.EqualError(t, err, "Missing key x")

	_, err = SchemaPath(map[string]*schema.Schema{})
	assert.EqualError(t, err, "[] does not compute")

	_, err = SchemaPath(map[string]*schema.Schema{
		"foo": {},
	}, "foo", "x")
	assert.EqualError(t, err, "foo is not nested resource")
}

func TestChooseFieldName(t *testing.T) {
	assert.Equal(t, "foo", chooseFieldName(reflect.StructField{
		Tag: `tf:"alias:foo"`,
	}))
}

type testSliceItem struct {
	SliceItem string   `json:"slice_item,omitempty"`
	Nested    *testPtr `json:"nested,omitempty"`
}

type testPtr struct {
	PtrItem string `json:"slice_item,omitempty"`
}

type testStruct struct {
	Integer        int               `json:"integer,omitempty" tf:"default:10,max_items:invalid"`
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
	Hidden         string            `json:"-"`
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

type Address struct {
	Line      string `json:"line" tf:"group:v"`
	Lijn      string `json:"lijn" tf:"group:v"`
	IsPrimary bool   `json:"primary"`

	OptionalString string `json:"optional_string,omitempty"`
	RequiredString string `json:"required_string"`
}

type Dummy struct {
	Enabled     bool              `json:"enabled" tf:"conflicts:workers"`
	Workers     int               `json:"workers,omitempty"`
	Description string            `json:"description,omitempty"`
	Addresses   []Address         `json:"addresses,omitempty" tf:"max_items:10"`
	Unique      []Address         `json:"unique,omitempty" tf:"slice_set"`
	Things      []string          `json:"things,omitempty" tf:"slice_set"`
	Tags        map[string]string `json:"tags,omitempty" tf:"max_items:5"`
	Home        *Address          `json:"home,omitempty" tf:"group:v"`
	House       *Address          `json:"house,omitempty" tf:"group:v"`
}

func TestStructToDataAndBack(t *testing.T) {
	d := schema.TestResourceDataRaw(t, scm, map[string]interface{}{})
	d.MarkNewResource()
	err := StructToData(testStruct{
		ComputedField: "x",
		BoolSlice:     []bool{true, false},
		FloatSlice:    []float64{.87, .98},
		IntSlice:      []int{435, 23, 6},
		MapField: map[string]string{
			"x": "y",
		},
		StringSlice: []string{"a", "b"},
		PtrItem: &testPtr{
			PtrItem: "x",
		},
		Hidden:         "x",
		SliceSetString: []string{"a", "b"},
		SliceSetStruct: []testSliceItem{
			{
				SliceItem: "x",
				Nested: &testPtr{
					PtrItem: "y",
				},
			},
		},
	}, scm, d)
	assert.NoError(t, err)

	var r testStruct
	err = DataToStructPointer(d, scm, &r)
	assert.NoError(t, err)
}

func TestSetPrimitiveOfKind(t *testing.T) {
	err := setPrimitiveValueOfKind("a.b.c", reflect.String, reflect.Value{}, []string{"_"})
	assert.EqualError(t, err, "a.b.c[[_]] is not a string")
	err = setPrimitiveValueOfKind("a.b.c", reflect.Int, reflect.Value{}, []string{"_"})
	assert.EqualError(t, err, "a.b.c[[_]] is not an int")
	err = setPrimitiveValueOfKind("a.b.c", reflect.Float64, reflect.Value{}, []string{"_"})
	assert.EqualError(t, err, "a.b.c[[_]] is not a float64")
	err = setPrimitiveValueOfKind("a.b.c", reflect.Bool, reflect.Value{}, []string{"_"})
	assert.EqualError(t, err, "a.b.c[[_]] is not a bool")
	err = setPrimitiveValueOfKind("a.b.c", reflect.Slice, reflect.Value{}, []string{"_"})
	assert.EqualError(t, err, "a.b.c[[_]] is not a valid primitive")
}

func TestPrimitiveReflectValueFromInterface(t *testing.T) {
	_, err := primitiveReflectValueFromInterface(reflect.String, []string{"_"}, "a", "b")
	assert.NoError(t, err)

	_, err = primitiveReflectValueFromInterface(reflect.Int, []string{"_"}, "a", "b")
	assert.EqualError(t, err, "a[b] '[_]' is not Int")
	_, err = primitiveReflectValueFromInterface(reflect.Int, 452345, "a", "b")
	assert.NoError(t, err)

	_, err = primitiveReflectValueFromInterface(reflect.Float32, []string{"_"}, "a", "b")
	assert.EqualError(t, err, "a[b] '[_]' is not Float32")
	_, err = primitiveReflectValueFromInterface(reflect.Float32, float32(1.3), "a", "b")
	assert.NoError(t, err)

	_, err = primitiveReflectValueFromInterface(reflect.Float64, []string{"_"}, "a", "b")
	assert.EqualError(t, err, "a[b] '[_]' is not Float64")
	_, err = primitiveReflectValueFromInterface(reflect.Float64, float64(1.3), "a", "b")
	assert.NoError(t, err)

	_, err = primitiveReflectValueFromInterface(reflect.Bool, []string{"_"}, "a", "b")
	assert.EqualError(t, err, "a[b] '[_]' is not Bool")
	_, err = primitiveReflectValueFromInterface(reflect.Bool, true, "a", "b")
	assert.NoError(t, err)

	_, err = primitiveReflectValueFromInterface(reflect.Slice, []string{"_"}, "a", "b")
	assert.EqualError(t, err, "a[b] '[_]' is not valid primitive")
}

func TestIterFields(t *testing.T) {
	v := reflect.ValueOf("x")
	err := iterFields(v, []string{"x"}, scm, nil)
	assert.EqualError(t, err, "Value of Struct is expected, but got String: \"x\"")

	v = reflect.ValueOf(testStruct{})
	err = iterFields(v, []string{}, map[string]*schema.Schema{
		"integer": {
			Type: schema.TypeInt,
		},
	}, nil)
	assert.EqualError(t, err, "Inconsistency: integer has omitempty, but is not optional")

	err = iterFields(v, []string{}, map[string]*schema.Schema{
		"non_optional": {
			Type:     schema.TypeString,
			Default:  nil,
			Optional: true,
		},
	}, nil)
	assert.EqualError(t, err, "Inconsistency: non_optional is optional, default is empty, but has no omitempty")

	err = iterFields(v, []string{}, map[string]*schema.Schema{
		"non_optional": {
			Type:     schema.TypeString,
			Default:  "_",
			Optional: true,
		},
	}, func(fieldSchema *schema.Schema, path []string, valueField *reflect.Value) error {
		return fmt.Errorf("test error")
	})
	assert.EqualError(t, err, "non_optional: test error")
}

func TestCollectionToMaps(t *testing.T) {
	v, err := collectionToMaps([]string{"a", "b"}, nil)
	assert.NoError(t, err)
	assert.Equal(t, []interface{}{"a", "b"}, v)

	_, err = collectionToMaps([]int{1, 2}, &schema.Schema{
		Elem: schema.TypeBool,
	})
	assert.EqualError(t, err, "not resource")
}

func TestStructToData(t *testing.T) {
	s := StructToSchema(Dummy{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		return s
	})
	assert.NotNil(t, s)
	assert.Equal(t, 5, s["tags"].MaxItems)
	assert.Equal(t, 10, s["addresses"].MaxItems)

	sp, err := SchemaPath(s, "addresses", "line")
	assert.NoError(t, err)
	assert.Equal(t, schema.TypeString, sp.Type)

	dummy := Dummy{
		Enabled:     false,
		Workers:     1004,
		Description: "something",
		Addresses: []Address{
			{
				Line:      "abc",
				IsPrimary: false,
			},
			{
				Line:      "def",
				IsPrimary: true,
			},
		},
		Unique: []Address{
			{
				Line:      "oop",
				IsPrimary: false,
			},
		},
		Things: []string{"one", "two", "two"},
		Tags: map[string]string{
			"Foo": "Bar",
		},
		Home: &Address{
			Line:      "bcd",
			IsPrimary: true,
		},
	}

	d := schema.TestResourceDataRaw(t, s, map[string]interface{}{})
	d.MarkNewResource()
	err = StructToData(dummy, s, d)
	assert.NoError(t, err)

	assert.Equal(t, "something", d.Get("description"))
	assert.Equal(t, false, d.Get("enabled"))
	assert.Equal(t, 2, d.Get("addresses.#"))

	// Empty optional string should not be set.
	{
		_, ok := d.GetOkExists("addresses.0.optional_string")
		assert.Falsef(t, ok, "Empty optional string should not be set in ResourceData")
	}

	// Empty required string should be set.
	{
		_, ok := d.GetOkExists("addresses.0.required_string")
		assert.Truef(t, ok, "Empty required string should be set in ResourceData")
	}

	var dummyCopy Dummy
	err = DataToStructPointer(d, s, &dummyCopy)
	assert.NoError(t, err)

	assert.Equal(t, len(dummyCopy.Addresses), len(dummy.Addresses))
	assert.Len(t, dummyCopy.Things, 2)
	assert.Len(t, dummy.Things, 3)

	err = d.Set("addresses", []interface{}{
		map[string]string{
			"line": "ABC",
			"lijn": "CBA",
		},
	})
	assert.NoError(t, err)

	err = DataToStructPointer(d, s, &dummyCopy)
	assert.NoError(t, err)
}
