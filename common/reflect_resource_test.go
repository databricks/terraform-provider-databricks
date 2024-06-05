package common

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

func TestReflectKind(t *testing.T) {
	assert.Equal(t, "Bool", reflectKind(1))
	assert.Equal(t, "other", reflectKind(29))
}

func TestSchemaPath(t *testing.T) {
	_, err := SchemaPath(map[string]*schema.Schema{}, "x", "y", "z")
	assert.EqualError(t, err, "missing key x")

	_, err = SchemaPath(map[string]*schema.Schema{})
	assert.EqualError(t, err, "[] does not compute")

	_, err = SchemaPath(map[string]*schema.Schema{
		"foo": {},
	}, "foo", "x")
	assert.EqualError(t, err, "foo is not nested resource")
}

func TestMustSchemaPath(t *testing.T) {
	x := MustSchemaPath(map[string]*schema.Schema{
		"foo": {},
	}, "foo")
	assert.NotNil(t, x)
}

func TestChooseFieldName(t *testing.T) {
	assert.Equal(t, "foo", chooseFieldName(reflect.StructField{
		Tag: `tf:"alias:foo"`,
	}))
}

func TestChooseFieldNameWithAliasesMap(t *testing.T) {
	type Bar struct {
		Foo string `json:"foo,omitempty"`
	}
	assert.Equal(t, "foo", chooseFieldNameWithAliases(reflect.StructField{
		Tag: `json:"bar"`,
	}, reflect.ValueOf(Bar{}).Type(), map[string]map[string]string{"common.Bar": {"bar": "foo"}}))
}

type testSliceItem struct {
	SliceItem string   `json:"slice_item,omitempty"`
	Nested    *testPtr `json:"nested,omitempty"`
}

type testPtr struct {
	PtrItem string `json:"slice_item,omitempty"`
}

type testStruct struct {
	Integer        int               `json:"integer,omitempty" tf:"default:10,min_items:invalid,max_items:invalid"`
	Float          float64           `json:"float,omitempty"`
	Bool           bool              `json:"bool,omitempty"`
	NonOptional    string            `json:"non_optional"`
	String         string            `json:"string,omitempty"`
	ComputedField  string            `json:"computed_field,omitempty" tf:"some_random_tag,computed,sensitive"`
	ForceNewField  string            `json:"force_new_field,omitempty" tf:"force_new"`
	MapField       map[string]string `json:"map_field,omitempty"`
	SliceSetStruct []testSliceItem   `json:"slice_set_struct,omitempty" tf:"slice_set"`
	SliceSetString []string          `json:"slice_set_string,omitempty" tf:"slice_set"`
	PtrItem        *testPtr          `json:"ptr_item,omitempty"`
	StringSlice    []string          `json:"string_slice,omitempty"`
	IntSlice       []int             `json:"int_slice,omitempty"`
	FloatSlice     []float64         `json:"float_slice,omitempty"`
	BoolSlice      []bool            `json:"bool_slice,omitempty"`
	TfOptional     string            `json:"tf_optional" tf:"optional"`
	Hidden         string            `json:"-"`
	Hidden2        string
	Indirect       []IndirectString `json:"indirect"`
}

type IndirectString string

type testRecursiveStruct struct {
	Task  *testJobTask `json:"task,omitempty"`
	Extra string       `json:"extra,omitempty"`
}

type testJobTask struct {
	ForEachTask *testForEachTask `json:"for_each_task,omitempty"`
	Extra       string           `json:"extra,omitempty"`
}

type testForEachTask struct {
	Task  *testJobTask `json:"task,omitempty"`
	Extra string       `json:"extra,omitempty"`
}

func (testRecursiveStruct) CustomizeSchema(s *CustomizableSchema) *CustomizableSchema {
	return s
}

func (testRecursiveStruct) MaxDepthForTypes() map[string]int {
	return map[string]int{"common.testForEachTask": 2}
}

var scm = StructToSchema(testStruct{}, nil)

var testStructFields = []string{"integer", "float", "non_optional", "string", "computed_field", "force_new_field", "map_field",
	"slice_set_struct", "slice_set_string", "ptr_item", "string_slice", "bool", "int_slice", "float_slice",
	"bool_slice", "tf_optional"}

var testStructOptionalFields = []string{"integer", "float", "string", "computed_field", "force_new_field", "map_field", "slice_set_struct",
	"ptr_item", "slice_set_string", "bool", "int_slice", "float_slice", "bool_slice", "tf_optional"}

var testStructRequiredFields = []string{"non_optional"}

var testStructComputedFields = []string{"computed_field"}

var testStructForceNewFields = []string{"force_new_field"}

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
		"force_new_field":  schema.TypeString,
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

func TestStructToSchema_forced_new_values_set(t *testing.T) {
	for _, field := range testStructForceNewFields {
		requiredField, ok := scm[field]
		assert.Truef(t, ok, "%s key not found", field)
		assert.Truef(t, requiredField.ForceNew, "force_new should be set to true in field: %s", field)
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
	Workers     int               `json:"workers,omitempty" tf:"suppress_diff"`
	Description string            `json:"description,omitempty"`
	Addresses   []Address         `json:"addresses,omitempty" tf:"min_items:1,max_items:10"`
	Unique      []Address         `json:"unique,omitempty" tf:"slice_set"`
	Things      []string          `json:"things,omitempty" tf:"slice_set"`
	Tags        map[string]string `json:"tags,omitempty" tf:"max_items:5"`
	Home        *Address          `json:"home,omitempty" tf:"group:v,suppress_diff"`
	House       *Address          `json:"house,omitempty" tf:"group:v"`
	Other       *Address          `json:"other,omitempty"`
}

type AddressNoTfTag struct {
	Line      string `json:"line"`
	Lijn      string `json:"lijn"`
	IsPrimary bool   `json:"primary"`

	OptionalString string `json:"optional_string,omitempty"`
	RequiredString string `json:"required_string"`
}

type DummyNoTfTag struct {
	Enabled     bool              `json:"enabled"`
	Workers     int               `json:"workers,omitempty"`
	Description string            `json:"description,omitempty"`
	Addresses   []AddressNoTfTag  `json:"addresses,omitempty"`
	Things      []string          `json:"things,omitempty"`
	Tags        map[string]string `json:"tags,omitempty"`
	Home        *AddressNoTfTag   `json:"home,omitempty"`
	House       *AddressNoTfTag   `json:"house,omitempty"`
	Other       *AddressNoTfTag   `json:"other,omitempty"`
}

type DummyResourceProvider struct {
	DummyNoTfTag
}

func (DummyResourceProvider) Aliases() map[string]map[string]string {
	return map[string]map[string]string{"common.DummyResourceProvider": {"enabled": "enabled_alias"},
		"common.AddressNoTfTag": {"primary": "primary_alias"}}
}

func (DummyResourceProvider) CustomizeSchema(s *CustomizableSchema) *CustomizableSchema {
	s.SchemaPath("addresses").SetMinItems(1)
	s.SchemaPath("addresses").SetMaxItems(10)
	s.SchemaPath("tags").SetMaxItems(5)
	s.SchemaPath("home").SetSuppressDiff()
	s.SchemaPath("things").Schema.Type = schema.TypeSet
	return s
}

var dummy = DummyNoTfTag{
	Enabled:     true,
	Workers:     1004,
	Description: "something",
	Addresses: []AddressNoTfTag{
		{
			Line:      "abc",
			IsPrimary: false,
		},
		{
			Line:      "def",
			IsPrimary: true,
		},
	},
	Things: []string{"one", "two", "two"},
	Tags: map[string]string{
		"Foo": "Bar",
	},
	Home: &AddressNoTfTag{
		Line:      "bcd",
		IsPrimary: true,
	},
}

func TestStructToDataAndBack(t *testing.T) {
	d := schema.TestResourceDataRaw(t, scm, map[string]any{})
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
	DataToStructPointer(d, scm, &r)
}

func TestStructToDataAndBackPanic(t *testing.T) {
	defer func() {
		p := recover()
		err := p.(error)
		assert.EqualError(t, err, "pointer is expected, but got Int: 1")
	}()
	DataToStructPointer(nil, nil, 1)
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
	err := iterFields(v, []string{"x"}, scm, nil, nil)
	assert.EqualError(t, err, "value of Struct is expected, but got String: \"x\"")

	v = reflect.ValueOf(testStruct{})
	err = iterFields(v, []string{}, map[string]*schema.Schema{
		"integer": {
			Type: schema.TypeInt,
		},
	}, nil, nil)
	assert.EqualError(t, err, "inconsistency: integer has omitempty, but is not optional")

	err = iterFields(v, []string{}, map[string]*schema.Schema{
		"non_optional": {
			Type:     schema.TypeString,
			Default:  nil,
			Optional: true,
		},
	}, nil, nil)
	assert.EqualError(t, err, "inconsistency: non_optional is optional, default is empty, but has no omitempty")

	err = iterFields(v, []string{}, map[string]*schema.Schema{
		"non_optional": {
			Type:     schema.TypeString,
			Default:  "_",
			Optional: true,
		},
	}, nil, func(fieldSchema *schema.Schema, path []string, valueField *reflect.Value) error {
		return fmt.Errorf("test error")
	})
	assert.EqualError(t, err, "non_optional: test error")
}

func TestCollectionToMaps(t *testing.T) {
	v, err := collectionToMaps([]string{"a", "b"}, nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, []any{"a", "b"}, v)

	_, err = collectionToMaps([]int{1, 2}, &schema.Schema{
		Elem: schema.TypeBool,
	}, nil)
	assert.EqualError(t, err, "not resource")
}

func TestStructToSchemaWithResourceProviderCustomization(t *testing.T) {
	s := StructToSchema(DummyResourceProvider{}, nil)
	assert.NotNil(t, s)
	assert.Equal(t, 5, s["tags"].MaxItems)
	assert.Equal(t, 10, s["addresses"].MaxItems)
}

func TestStructToSchemaWithResourceProviderAliases(t *testing.T) {
	s := StructToSchema(DummyResourceProvider{}, nil)
	sp, err := SchemaPath(s, "enabled_alias")
	assert.NoError(t, err)
	assert.Equal(t, schema.TypeBool, sp.Type)
}

func TestStructToDataWithResourceProviderStruct(t *testing.T) {
	s := StructToSchema(DummyResourceProvider{}, nil)

	dummyResourceProvider := DummyResourceProvider{DummyNoTfTag: dummy}
	d := schema.TestResourceDataRaw(t, s, map[string]any{})
	d.MarkNewResource()
	err := StructToData(dummyResourceProvider, s, d)
	assert.NoError(t, err)

	assert.Equal(t, "something", d.Get("description"))
	assert.Equal(t, true, d.Get("enabled_alias")) // Testing aliases.
	assert.Equal(t, 2, d.Get("addresses.#"))

	assert.NotNil(t, s["home"].DiffSuppressFunc)
	assert.True(t, s["home"].DiffSuppressFunc("home.#", "1", "0", d))
	assert.False(t, s["home"].DiffSuppressFunc("home.#", "1", "1", d))

	{
		//lint:ignore SA1019 Empty optional string should not be set.
		_, ok := d.GetOkExists("addresses.0.optional_string")
		assert.Falsef(t, ok, "Empty optional string should not be set in ResourceData")
	}

	{
		//lint:ignore SA1019 Empty required string should be set.
		_, ok := d.GetOkExists("addresses.0.required_string")
		assert.Truef(t, ok, "Empty required string should be set in ResourceData")
	}
}

func TestDataToStructPointerWithResourceProviderStruct(t *testing.T) {
	s := StructToSchema(DummyResourceProvider{}, nil)
	d := schema.TestResourceDataRaw(t, s, map[string]any{})
	d.MarkNewResource()
	dummyResourceProvider := DummyResourceProvider{DummyNoTfTag: dummy}
	err := StructToData(dummyResourceProvider, s, d)
	assert.NoError(t, err)
	var dummyCopy DummyResourceProvider
	DataToStructPointer(d, s, &dummyCopy)

	assert.Equal(t, len(dummyCopy.Addresses), len(dummy.Addresses))
	assert.Equal(t, dummyCopy.Enabled, dummy.Enabled)
	assert.Len(t, dummyCopy.Things, 2)

	err = d.Set("addresses", []any{
		map[string]string{
			"line": "ABC",
			"lijn": "CBA",
		},
	})
	assert.NoError(t, err)

	DataToStructPointer(d, s, &dummyCopy)
}

func TestStructToData_EmptyField(t *testing.T) {
	type EmptyField struct{}
	type Container struct {
		EmptyField *EmptyField `json:"empty_field,omitempty"`
	}
	s := StructToSchema(Container{}, nil)
	assert.NotNil(t, s)

	dummy := Container{
		EmptyField: &EmptyField{},
	}

	d := schema.TestResourceDataRaw(t, s, map[string]any{})
	d.MarkNewResource()
	err := StructToData(dummy, s, d)
	assert.NoError(t, err)
	assert.Equal(t, 1, d.Get("empty_field.#"))
}

func TestStructToData_EmptyFieldNil(t *testing.T) {
	type EmptyField struct{}
	type Container struct {
		EmptyField *EmptyField `json:"empty_field,omitempty"`
	}
	s := StructToSchema(Container{}, nil)
	assert.NotNil(t, s)

	dummy := Container{
		EmptyField: nil,
	}

	d := schema.TestResourceDataRaw(t, s, map[string]any{})
	d.MarkNewResource()
	err := StructToData(dummy, s, d)
	assert.NoError(t, err)
	assert.Equal(t, 0, d.Get("empty_field.#"))
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

	d := schema.TestResourceDataRaw(t, s, map[string]any{})
	d.MarkNewResource()
	err = StructToData(dummy, s, d)
	assert.NoError(t, err)

	assert.Equal(t, "something", d.Get("description"))
	assert.Equal(t, false, d.Get("enabled"))
	assert.Equal(t, 2, d.Get("addresses.#"))

	assert.NotNil(t, s["home"].DiffSuppressFunc)
	assert.True(t, s["home"].DiffSuppressFunc("home.#", "1", "0", d))
	assert.False(t, s["home"].DiffSuppressFunc("home.#", "1", "1", d))

	{
		//lint:ignore SA1019 Empty optional string should not be set.
		_, ok := d.GetOkExists("addresses.0.optional_string")
		assert.Falsef(t, ok, "Empty optional string should not be set in ResourceData")
	}

	{
		//lint:ignore SA1019 Empty required string should be set.
		_, ok := d.GetOkExists("addresses.0.required_string")
		assert.Truef(t, ok, "Empty required string should be set in ResourceData")
	}

	var dummyCopy Dummy
	DataToStructPointer(d, s, &dummyCopy)

	assert.Equal(t, len(dummyCopy.Addresses), len(dummy.Addresses))
	assert.Len(t, dummyCopy.Things, 2)
	assert.Len(t, dummy.Things, 3)

	err = d.Set("addresses", []any{
		map[string]string{
			"line": "ABC",
			"lijn": "CBA",
		},
	})
	assert.NoError(t, err)

	DataToStructPointer(d, s, &dummyCopy)
}

func TestDiffSuppressor(t *testing.T) {
	stringSchema := &schema.Schema{
		Type: schema.TypeString,
	}
	dsf := diffSuppressor("foo", stringSchema)
	d := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
		"foo": {
			Type:     schema.TypeBool,
			Optional: true,
		},
	}, map[string]any{})
	// no suppress
	assert.False(t, dsf("", "old", "new", d))
	// suppress
	assert.True(t, dsf("", "old", "", d))
}

func TestDiffSuppressorWhenNumberExplicitlyChangedToZero(t *testing.T) {
	intSchema := &schema.Schema{
		Type: schema.TypeInt,
	}
	dsf := diffSuppressor("foo", intSchema)
	noChange := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
		"foo": {
			Type:     schema.TypeInt,
			Optional: true,
		},
	}, map[string]any{})
	// no suppress
	assert.False(t, dsf("foo", "1", "2", noChange))
	// suppress
	assert.True(t, dsf("foo", "1", "0", noChange))

	change := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
		"foo": {
			Type:     schema.TypeInt,
			Optional: true,
		},
	}, map[string]any{
		"foo": 1,
	})

	// no suppress
	assert.False(t, dsf("foo", "1", "2", change))
	assert.False(t, dsf("foo", "1", "0", change))
}

func TestTypeToSchemaNoStruct(t *testing.T) {
	defer func() {
		p := recover()
		assert.Equal(t,
			"Schema value of Struct is expected, but got Int: 1",
			fmt.Sprintf("%s", p))
	}()
	v := reflect.ValueOf(1)
	typeToSchema(v, nil, getEmptyTrackingContext())
}

func TestTypeToSchemaUnsupported(t *testing.T) {
	defer func() {
		p := recover()
		assert.Equal(t, "unknown type for new: Chan",
			fmt.Sprintf("%s", p))
	}()
	type nonsense struct {
		New chan int `json:"new"`
	}
	v := reflect.ValueOf(nonsense{})
	typeToSchema(v, nil, getEmptyTrackingContext())
}

type data map[string]any

func (a data) GetOk(key string) (any, bool) {
	v, ok := a[key]
	return v, ok
}

func (a data) GetOkExists(key string) (any, bool) {
	return a.GetOk(key)
}

func TestDiffToStructPointerPanic(t *testing.T) {
	type Nonsense struct {
		New int `json:"new,omitempty"`
	}
	s := schema.InternalMap(map[string]*schema.Schema{
		"new": {
			Type:     schema.TypeInt,
			Optional: true,
		},
	})
	defer func() {
		p := recover()
		err := p.(error)
		assert.EqualError(t, err, "pointer is expected, but got Struct: common.Nonsense{New:0}")
	}()
	DiffToStructPointer(data{"new": "3"}, s, Nonsense{})
}

func TestDiffToStructPointer(t *testing.T) {
	type Nonsense struct {
		New int `json:"new,omitempty"`
	}
	s := schema.InternalMap(map[string]*schema.Schema{
		"new": {
			Type:     schema.TypeInt,
			Optional: true,
		},
	})
	var n Nonsense
	DiffToStructPointer(data{"new": 3}, s, &n)
	assert.Equal(t, 3, n.New)
}

func TestReadListFromData(t *testing.T) {
	err := readListFromData([]string{}, data{}, []any{}, nil, nil, nil, nil)
	assert.NoError(t, err)

	x := reflect.ValueOf(0)
	err = readListFromData([]string{}, data{}, []any{1}, &x, nil, nil, nil)
	assert.EqualError(t, err, "[[1]] unknown collection field")
}

func TestReadReflectValueFromDataCornerCases(t *testing.T) {
	type Nonsense struct {
		New     float64 `json:"new,omitempty"`
		Invalid int     `json:"invalid"`
	}
	s := schema.InternalMap(map[string]*schema.Schema{
		"new": {
			Type:     schema.TypeFloat,
			Optional: true,
		},
		"invalid": {
			Type:     schema.TypeInvalid,
			Required: true,
		},
	})
	var n Nonsense
	v := reflect.ValueOf(&n)
	rv := v.Elem()
	err := readReflectValueFromData([]string{}, data{"new": 0.123, "invalid": 1}, rv, s, nil)
	assert.EqualError(t, err, "invalid: invalid[1] unsupported field type")
}

func TestStructToData_CornerCases(t *testing.T) {
	type Nonsense struct {
		WillBeIgnored int       `json:"will_be_ignored,omitempty"`
		Ints          []int     `json:"ints"`
		Addresses     []Address `json:"addrs"`
	}
	s := schema.InternalMap(map[string]*schema.Schema{
		"will_be_ignored": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"ints": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
		"addrs": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"line": {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
	})
	d := schema.TestResourceDataRaw(t, s, map[string]any{})
	d.Set("ints", []int{1})
	d.Set("addrs", []any{
		map[string]any{
			"line": "a",
		},
	})
	err := StructToData(Nonsense{
		WillBeIgnored: 1,
		Ints:          []int{},
		Addresses:     []Address{},
	}, s, d)
	assert.NoError(t, err)
}

func TestDataToReflectValueBypass(t *testing.T) {
	err := DataToReflectValue(nil, map[string]*schema.Schema{}, reflect.ValueOf(0))
	assert.EqualError(t, err, "value of Struct is expected, but got Int: 0")
}

func TestDataResource(t *testing.T) {
	r := func() *schema.Resource {
		type entry struct {
			In  string `json:"in"`
			Out string `json:"out,omitempty" tf:"computed"`
		}
		return DataResource(entry{}, func(ctx context.Context, e any, c *DatabricksClient) error {
			dto := e.(*entry)
			dto.Out = "out: " + dto.In
			if dto.In == "fail" {
				return fmt.Errorf("happens")
			}
			return nil
		}).ToResource()
	}()
	d := r.TestResourceData()
	d.Set("in", "test")

	diags := r.ReadContext(context.Background(), d, &DatabricksClient{})
	assert.Len(t, diags, 0)
	assert.Equal(t, "out: test", d.Get("out"))
	assert.Equal(t, "_", d.Id())

	d.Set("in", "fail")
	diags = r.ReadContext(context.Background(), d, &DatabricksClient{})
	assert.Len(t, diags, 1)
}

func TestDataResourceWithID(t *testing.T) {
	r := func() *schema.Resource {
		type entry struct {
			In  string `json:"in"`
			ID  string `json:"id,omitempty" tf:"computed"`
			Out string `json:"out,omitempty" tf:"computed"`
		}
		return DataResource(entry{}, func(ctx context.Context, e any, c *DatabricksClient) error {
			dto := e.(*entry)
			dto.Out = "out: " + dto.In
			dto.ID = "abc"
			return nil
		}).ToResource()
	}()
	d := r.TestResourceData()
	d.Set("in", "id")
	diags := r.ReadContext(context.Background(), d, &DatabricksClient{})
	assert.Len(t, diags, 0)
	assert.Equal(t, "out: id", d.Get("out"))
	assert.Equal(t, "abc", d.Id())
}

func TestStructToSchema_go_sdk_embedded(t *testing.T) {
	type T struct {
		sql.GetWarehouseResponse
		Extra string `json:"extra,omitempty" tf:"computed"`
	}
	s := StructToSchema(T{}, nil)
	autoStopMins, ok := s["auto_stop_mins"]
	assert.True(t, ok)
	assert.Equal(t, schema.TypeInt, autoStopMins.Type)
	assert.True(t, autoStopMins.Optional)
}

func TestStructToData_go_sdk_embedded(t *testing.T) {
	type T struct {
		sql.GetWarehouseResponse
		Extra string `json:"extra,omitempty" tf:"computed"`
	}
	s := StructToSchema(T{}, nil)
	SetRequired(s["cluster_size"])
	r := &schema.Resource{
		Schema: s,
	}
	d := r.TestResourceData()
	d.Set("cluster_size", "original")
	err := StructToData(T{
		GetWarehouseResponse: sql.GetWarehouseResponse{
			ClusterSize: "abc123",
		},
		Extra: "extra",
	}, s, d)
	assert.NoError(t, err)
	assert.Equal(t, "abc123", d.Get("cluster_size"))
	assert.Equal(t, "extra", d.Get("extra"))
}

func TestStructToSchema_go_sdk_field(t *testing.T) {
	type T struct {
		Warehouse *sql.GetWarehouseResponse `json:"warehouse,omitempty"`
		Extra     string                    `json:"extra,omitempty" tf:"computed"`
	}
	s := StructToSchema(T{}, nil)
	warehouse, ok := s["warehouse"]
	assert.True(t, ok)
	autoStopMins, ok := warehouse.Elem.(*schema.Resource).Schema["auto_stop_mins"]
	assert.True(t, ok)
	assert.Equal(t, schema.TypeInt, autoStopMins.Type)
	assert.True(t, autoStopMins.Optional)
}

func TestStructToData_go_sdk_field(t *testing.T) {
	type T struct {
		Warehouse *sql.GetWarehouseResponse `json:"warehouse,omitempty"`
		Extra     string                    `json:"extra,omitempty" tf:"computed"`
	}
	s := StructToSchema(T{}, nil)
	SetRequired(MustSchemaPath(s, "warehouse", "cluster_size"))
	r := &schema.Resource{
		Schema: s,
	}
	d := r.TestResourceData()
	d.Set("warehouse", []map[string]any{
		{
			"cluster_size": "original",
		},
	})
	err := StructToData(T{
		Warehouse: &sql.GetWarehouseResponse{
			ClusterSize: "abc123",
		},
		Extra: "extra",
	}, s, d)
	assert.NoError(t, err)
	assert.Equal(t, "abc123", d.Get("warehouse.0.cluster_size"))
	assert.Equal(t, "extra", d.Get("extra"))
}

func TestStructToSchema_recursive(t *testing.T) {
	s := StructToSchema(testRecursiveStruct{}, nil)
	// Assert that the recursion cannot go beyond 2 levels deep.
	_, err := SchemaPath(s, "task", "for_each_task")
	assert.NoError(t, err)
	_, err = SchemaPath(s, "task", "for_each_task", "task", "for_each_task")
	assert.NoError(t, err)
	_, err = SchemaPath(s, "task", "for_each_task", "task", "for_each_task", "task")
	assert.NoError(t, err)
	// Should error out on the 3rd level of for_each_task.
	_, err = SchemaPath(s, "task", "for_each_task", "task", "for_each_task", "task", "for_each_task")
	assert.Error(t, err)
}

func TestStructToData_IndirectString(t *testing.T) {
	d := schema.TestResourceDataRaw(t, scm, map[string]any{})
	d.MarkNewResource()
	err := StructToData(testStruct{
		Indirect: []IndirectString{"a"},
	}, scm, d)
	assert.NoError(t, err)
}
