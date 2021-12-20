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
	Hidden         string            `json:"-"`
	Hidden2        string
}

var scm = StructToSchema(testStruct{}, nil)

var testStructFields = []string{"integer", "float", "non_optional", "string", "computed_field", "force_new_field", "map_field",
	"slice_set_struct", "slice_set_string", "ptr_item", "string_slice", "bool", "int_slice", "float_slice",
	"bool_slice"}

var testStructOptionalFields = []string{"integer", "float", "string", "computed_field", "force_new_field", "map_field", "slice_set_struct",
	"ptr_item", "slice_set_string", "bool", "int_slice", "float_slice", "bool_slice"}

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
	Workers     int               `json:"workers,omitempty"`
	Description string            `json:"description,omitempty"`
	Addresses   []Address         `json:"addresses,omitempty" tf:"max_items:10"`
	Unique      []Address         `json:"unique,omitempty" tf:"slice_set"`
	Things      []string          `json:"things,omitempty" tf:"slice_set"`
	Tags        map[string]string `json:"tags,omitempty" tf:"max_items:5"`
	Home        *Address          `json:"home,omitempty" tf:"group:v,suppress_diff"`
	House       *Address          `json:"house,omitempty" tf:"group:v"`
	Other       *Address          `json:"other,omitempty"`
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

	err = DataToStructPointer(d, scm, 1)
	assert.EqualError(t, err, "pointer is expected, but got Int: 1")
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
	assert.EqualError(t, err, "value of Struct is expected, but got String: \"x\"")

	v = reflect.ValueOf(testStruct{})
	err = iterFields(v, []string{}, map[string]*schema.Schema{
		"integer": {
			Type: schema.TypeInt,
		},
	}, nil)
	assert.EqualError(t, err, "inconsistency: integer has omitempty, but is not optional")

	err = iterFields(v, []string{}, map[string]*schema.Schema{
		"non_optional": {
			Type:     schema.TypeString,
			Default:  nil,
			Optional: true,
		},
	}, nil)
	assert.EqualError(t, err, "inconsistency: non_optional is optional, default is empty, but has no omitempty")

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

func TestDiffSuppressor(t *testing.T) {
	dsf := diffSuppressor("")
	d := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
		"foo": {
			Type:     schema.TypeBool,
			Optional: true,
		},
	}, map[string]interface{}{})
	// no suppress
	assert.False(t, dsf("", "old", "new", d))
	// suppress
	assert.True(t, dsf("", "old", "", d))
}

func TestTypeToSchemaNoStruct(t *testing.T) {
	defer func() {
		p := recover()
		assert.Equal(t,
			"Schema value of Struct is expected, but got Int: 1",
			fmt.Sprintf("%s", p))
	}()
	v := reflect.ValueOf(1)
	typeToSchema(v, v.Type(), []string{})
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
	typeToSchema(v, v.Type(), []string{})
}

type data map[string]interface{}

func (a data) GetOk(key string) (interface{}, bool) {
	v, ok := a[key]
	return v, ok
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
	err := DiffToStructPointer(data{"new": "3"}, s, Nonsense{})
	assert.EqualError(t, err, "pointer is expected, but got Struct: common.Nonsense{New:0}")

	var n Nonsense
	err = DiffToStructPointer(data{"new": 3}, s, &n)
	assert.NoError(t, err)
	assert.Equal(t, 3, n.New)
}

func TestReadListFromData(t *testing.T) {
	err := readListFromData([]string{}, data{}, []interface{}{}, nil, nil, nil)
	assert.NoError(t, err)

	x := reflect.ValueOf(0)
	err = readListFromData([]string{}, data{}, []interface{}{1}, &x, nil, nil)
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
	err := readReflectValueFromData([]string{}, data{"new": 0.123, "invalid": 1}, rv, s)
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
	d := schema.TestResourceDataRaw(t, s, map[string]interface{}{})
	d.Set("ints", []int{1})
	d.Set("addrs", []interface{}{
		map[string]interface{}{
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
	err := DataToReflectValue(nil, &schema.Resource{Schema: map[string]*schema.Schema{}}, reflect.ValueOf(0))
	assert.EqualError(t, err, "value of Struct is expected, but got Int: 0")
}
