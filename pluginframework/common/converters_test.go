package pluginframework

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

type DummyTfSdk struct {
	Enabled           types.Bool                  `tfsdk:"enabled" tf:"optional"`
	Workers           types.Int64                 `tfsdk:"workers" tf:""` // Test required field
	Floats            types.Float64               `tfsdk:"floats" tf:""`  // Test required field
	Description       types.String                `tfsdk:"description" tf:""`
	Tasks             types.String                `tfsdk:"task" tf:"optional"`
	Nested            *DummyNestedTfSdk           `tfsdk:"nested" tf:"optional"`
	NoPointerNested   DummyNestedTfSdk            `tfsdk:"no_pointer_nested" tf:"optional"`
	NestedList        []DummyNestedTfSdk          `tfsdk:"nested_list" tf:"optional"`
	NestedPointerList []*DummyNestedTfSdk         `tfsdk:"nested_pointer_list" tf:"optional"`
	Map               map[string]types.String     `tfsdk:"map" tf:"optional"`
	NestedMap         map[string]DummyNestedTfSdk `tfsdk:"nested_map" tf:"optional"`
	Repeated          []types.Int64               `tfsdk:"repeated" tf:"optional"`
	Attributes        map[string]types.String     `tfsdk:"attributes" tf:"optional"`
	EnumField         types.String                `tfsdk:"enum_field" tf:"optional"`
	AdditionalField   types.String                `tfsdk:"additional_field" tf:"optional"`
	DistinctField     types.String                `tfsdk:"distinct_field" tf:"optional"` // distinct field that the gosdk struct doesn't have
	Irrelevant        types.String                `tfsdk:"-"`
}

type TestEnum string

const TestEnumA TestEnum = `TEST_ENUM_A`

const TestEnumB TestEnum = `TEST_ENUM_B`

const TestEnumC TestEnum = `TEST_ENUM_C`

func (f *TestEnum) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TestEnum) Set(v string) error {
	switch v {
	case `TEST_ENUM_A`, `TEST_ENUM_B`, `TEST_ENUM_C`:
		*f = TestEnum(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "TEST_ENUM_A", "TEST_ENUM_B", "TEST_ENUM_C"`, v)
	}
}

func (f *TestEnum) Type() string {
	return "MonitorInfoStatus"
}

type DummyNestedTfSdk struct {
	Name    types.String `tfsdk:"name" tf:"optional"`
	Enabled types.Bool   `tfsdk:"enabled" tf:"optional"`
}

type DummyGoSdk struct {
	Enabled           bool                        `json:"enabled"`
	Workers           int64                       `json:"workers"`
	Floats            float64                     `json:"floats"`
	Description       string                      `json:"description"`
	Tasks             string                      `json:"tasks"`
	Nested            *DummyNestedGoSdk           `json:"nested"`
	NoPointerNested   DummyNestedGoSdk            `json:"no_pointer_nested"`
	NestedList        []DummyNestedGoSdk          `json:"nested_list"`
	NestedPointerList []*DummyNestedGoSdk         `json:"nested_pointer_list"`
	Map               map[string]string           `json:"map"`
	NestedMap         map[string]DummyNestedGoSdk `json:"nested_map"`
	Repeated          []int64                     `json:"repeated"`
	Attributes        map[string]string           `json:"attributes"`
	EnumField         TestEnum                    `json:"enum_field"`
	AdditionalField   string                      `json:"additional_field"`
	DistinctField     string                      `json:"distinct_field"` // distinct field that the tfsdk struct doesn't have
	ForceSendFields   []string                    `json:"-"`
}

type DummyNestedGoSdk struct {
	Name            string   `json:"name"`
	Enabled         bool     `json:"enabled"`
	ForceSendFields []string `json:"-"`
}

type emptyCtx struct{}

func (emptyCtx) Deadline() (deadline time.Time, ok bool) {
	return
}

func (emptyCtx) Done() <-chan struct{} {
	return nil
}

func (emptyCtx) Err() error {
	return nil
}

func (emptyCtx) Value(key any) any {
	return nil
}

var ctx = emptyCtx{}

// Constructing a dummy tfsdk struct.
var tfSdkStructComplex = DummyTfSdk{
	Enabled:     types.BoolValue(false),
	Workers:     types.Int64Value(12),
	Description: types.StringValue("abc"),
	Tasks:       types.StringNull(),
	Nested: &DummyNestedTfSdk{
		Name:    types.StringValue("def"),
		Enabled: types.BoolValue(true),
	},
	NoPointerNested: DummyNestedTfSdk{
		Name:    types.StringValue("def"),
		Enabled: types.BoolValue(true),
	},
	NestedList: []DummyNestedTfSdk{
		{
			Name:    types.StringValue("def"),
			Enabled: types.BoolValue(true),
		},
		{
			Name:    types.StringValue("def"),
			Enabled: types.BoolValue(true),
		},
	},
	Map: map[string]types.String{
		"key1": types.StringValue("value1"),
		"key2": types.StringValue("value2"),
	},
	NestedMap: map[string]DummyNestedTfSdk{
		"key1": {
			Name:    types.StringValue("abc"),
			Enabled: types.BoolValue(false),
		},
		"key2": {
			Name:    types.StringValue("def"),
			Enabled: types.BoolValue(true),
		},
	},
	NestedPointerList: []*DummyNestedTfSdk{
		{
			Name:    types.StringValue("def"),
			Enabled: types.BoolValue(true),
		},
		{
			Name:    types.StringValue("def"),
			Enabled: types.BoolValue(true),
		},
	},
	Attributes: map[string]types.String{"key": types.StringValue("value")},
	EnumField:  types.StringValue("TEST_ENUM_A"),
	Repeated:   []types.Int64{types.Int64Value(12), types.Int64Value(34)},
}

func TestGetAndSetPluginFramework(t *testing.T) {
	// Also test StructToSchema.
	scm := PluginFrameworkResourceStructToSchema(DummyTfSdk{}, nil)
	state := tfsdk.State{
		Schema: scm,
	}

	// Test optional is being handled properly.
	assert.True(t, scm.Attributes["enabled"].IsOptional())
	assert.True(t, scm.Attributes["workers"].IsRequired())
	assert.True(t, scm.Attributes["floats"].IsRequired())

	// Assert that we can set state from the tfsdk struct.
	diags := state.Set(ctx, tfSdkStructComplex)
	assert.Len(t, diags, 0)

	// Assert that we can get a struct from the state.
	getterStruct := DummyTfSdk{}
	diags = state.Get(ctx, &getterStruct)
	assert.Len(t, diags, 0)

	// Assert the struct populated from .Get is exactly the same as the original tfsdk struct.
	assert.True(t, reflect.DeepEqual(getterStruct, tfSdkStructComplex))
}

func TestStructConversion(t *testing.T) {
	// Convert from tfsdk to gosdk struct using the converter function
	convertedGoSdkStruct := DummyGoSdk{}
	diags := TfSdkToGoSdkStruct(tfSdkStructComplex, &convertedGoSdkStruct, ctx)
	assert.True(t, !diags.HasError())

	// Convert the gosdk struct back to tfsdk struct
	convertedTfSdkStruct := DummyTfSdk{}
	diags = GoSdkToTfSdkStruct(convertedGoSdkStruct, &convertedTfSdkStruct, ctx)
	assert.True(t, !diags.HasError())

	// Assert that the struct is exactly the same after tfsdk --> gosdk --> tfsdk
	assert.True(t, reflect.DeepEqual(tfSdkStructComplex, convertedTfSdkStruct))
}

// Function to construct individual test case with a pair of matching tfSdkStruct and gosdkStruct.
// Verifies that the conversion both ways are working as expected.
func ConverterTestCase(t *testing.T, description string, tfSdkStruct DummyTfSdk, goSdkStruct DummyGoSdk) {
	convertedGoSdkStruct := DummyGoSdk{}
	assert.True(t, !TfSdkToGoSdkStruct(tfSdkStruct, &convertedGoSdkStruct, ctx).HasError())
	assert.True(t, reflect.DeepEqual(convertedGoSdkStruct, goSdkStruct), fmt.Sprintf("tfsdk to gosdk conversion - %s", description))

	convertedTfSdkStruct := DummyTfSdk{}
	assert.True(t, !GoSdkToTfSdkStruct(goSdkStruct, &convertedTfSdkStruct, ctx).HasError())
	assert.True(t, reflect.DeepEqual(convertedTfSdkStruct, tfSdkStruct), fmt.Sprintf("gosdk to tfsdk conversion - %s", description))
}

func TestConverter(t *testing.T) {
	ConverterTestCase(
		t,
		"string conversion",
		DummyTfSdk{Description: types.StringValue("abc")},
		DummyGoSdk{Description: "abc", ForceSendFields: []string{"Description"}},
	)
	ConverterTestCase(
		t,
		"bool conversion",
		DummyTfSdk{Enabled: types.BoolValue(true)},
		DummyGoSdk{Enabled: true, ForceSendFields: []string{"Enabled"}},
	)
	ConverterTestCase(
		t,
		"int64 conversion",
		DummyTfSdk{Workers: types.Int64Value(123)},
		DummyGoSdk{Workers: 123, ForceSendFields: []string{"Workers"}},
	)
	ConverterTestCase(
		t,
		"tf null value conversion",
		DummyTfSdk{Workers: types.Int64Null()},
		DummyGoSdk{},
	)
	ConverterTestCase(
		t,
		"float64 conversion",
		DummyTfSdk{Floats: types.Float64Value(1.1)},
		DummyGoSdk{Floats: 1.1, ForceSendFields: []string{"Floats"}},
	)
	ConverterTestCase(
		t,
		"enum conversion",
		DummyTfSdk{EnumField: types.StringValue("TEST_ENUM_A")},
		DummyGoSdk{EnumField: TestEnumA},
	)
	ConverterTestCase(
		t,
		"struct conversion",
		DummyTfSdk{NoPointerNested: DummyNestedTfSdk{
			Name:    types.StringValue("def"),
			Enabled: types.BoolValue(true),
		}},
		DummyGoSdk{NoPointerNested: DummyNestedGoSdk{
			Name:            "def",
			Enabled:         true,
			ForceSendFields: []string{"Name", "Enabled"},
		}},
	)
	ConverterTestCase(
		t,
		"pointer conversion",
		DummyTfSdk{Nested: &DummyNestedTfSdk{
			Name:    types.StringValue("def"),
			Enabled: types.BoolValue(true),
		}},
		DummyGoSdk{Nested: &DummyNestedGoSdk{
			Name:            "def",
			Enabled:         true,
			ForceSendFields: []string{"Name", "Enabled"},
		}},
	)
	ConverterTestCase(
		t,
		"list conversion",
		DummyTfSdk{Repeated: []types.Int64{types.Int64Value(12), types.Int64Value(34)}},
		DummyGoSdk{Repeated: []int64{12, 34}},
	)
	ConverterTestCase(
		t,
		"map conversion",
		DummyTfSdk{Attributes: map[string]types.String{"key": types.StringValue("value")}},
		DummyGoSdk{Attributes: map[string]string{"key": "value"}},
	)
	ConverterTestCase(
		t,
		"nested list conversion",
		DummyTfSdk{NestedList: []DummyNestedTfSdk{
			{
				Name:    types.StringValue("def"),
				Enabled: types.BoolValue(true),
			},
			{
				Name:    types.StringValue("def"),
				Enabled: types.BoolValue(true),
			},
		}},
		DummyGoSdk{NestedList: []DummyNestedGoSdk{
			{
				Name:            "def",
				Enabled:         true,
				ForceSendFields: []string{"Name", "Enabled"},
			},
			{
				Name:            "def",
				Enabled:         true,
				ForceSendFields: []string{"Name", "Enabled"},
			},
		}},
	)
	ConverterTestCase(
		t,
		"nested list conversion",
		DummyTfSdk{NestedList: []DummyNestedTfSdk{
			{
				Name:    types.StringValue("abc"),
				Enabled: types.BoolValue(true),
			},
			{
				Name:    types.StringValue("def"),
				Enabled: types.BoolValue(false),
			},
		}},
		DummyGoSdk{NestedList: []DummyNestedGoSdk{
			{
				Name:            "abc",
				Enabled:         true,
				ForceSendFields: []string{"Name", "Enabled"},
			},
			{
				Name:            "def",
				Enabled:         false,
				ForceSendFields: []string{"Name", "Enabled"},
			},
		}},
	)
	ConverterTestCase(
		t,
		"nested map conversion",
		DummyTfSdk{NestedMap: map[string]DummyNestedTfSdk{
			"key1": {
				Name:    types.StringValue("abc"),
				Enabled: types.BoolValue(true),
			},
			"key2": {
				Name:    types.StringValue("def"),
				Enabled: types.BoolValue(false),
			},
		}},
		DummyGoSdk{NestedMap: map[string]DummyNestedGoSdk{
			"key1": {
				Name:            "abc",
				Enabled:         true,
				ForceSendFields: []string{"Name", "Enabled"},
			},
			"key2": {
				Name:            "def",
				Enabled:         false,
				ForceSendFields: []string{"Name", "Enabled"},
			},
		}},
	)
}
