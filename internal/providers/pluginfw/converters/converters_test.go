package converters

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

type DummyTfSdk struct {
	Enabled           types.Bool                  `tfsdk:"enabled" tf:"optional"`
	Workers           types.Int64                 `tfsdk:"workers" tf:""`
	Floats            types.Float64               `tfsdk:"floats" tf:""`
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
	DistinctField     types.String                `tfsdk:"distinct_field" tf:"optional"`
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

// Function to construct individual test case with a pair of matching tfSdkStruct and gosdkStruct.
// Verifies that the conversion both ways are working as expected.
func RunConverterTest(t *testing.T, description string, tfSdkStruct DummyTfSdk, goSdkStruct DummyGoSdk) {
	convertedGoSdkStruct := DummyGoSdk{}
	assert.True(t, !TfSdkToGoSdkStruct(context.Background(), tfSdkStruct, &convertedGoSdkStruct).HasError())
	assert.True(t, reflect.DeepEqual(convertedGoSdkStruct, goSdkStruct), fmt.Sprintf("tfsdk to gosdk conversion - %s", description))

	convertedTfSdkStruct := DummyTfSdk{}
	assert.True(t, !GoSdkToTfSdkStruct(context.Background(), goSdkStruct, &convertedTfSdkStruct).HasError())
	assert.True(t, reflect.DeepEqual(convertedTfSdkStruct, tfSdkStruct), fmt.Sprintf("gosdk to tfsdk conversion - %s", description))
}

func TestTfSdkToGoSdkStructConversionFailure(t *testing.T) {
	tfSdkStruct := DummyTfSdk{}
	goSdkStruct := DummyGoSdk{}
	actualDiagnostics := TfSdkToGoSdkStruct(context.Background(), tfSdkStruct, goSdkStruct)
	expectedDiagnostics := diag.Diagnostics{diag.NewErrorDiagnostic(tfSdkToGoSdkStructConversionFailureMessage, "please provide a pointer for the gosdk struct, got DummyGoSdk")}
	assert.True(t, actualDiagnostics.HasError())
	assert.True(t, actualDiagnostics.Equal(expectedDiagnostics))
}

func TestGoSdkToTfSdkStructConversionFailure(t *testing.T) {
	tfSdkStruct := DummyTfSdk{}
	goSdkStruct := DummyGoSdk{}
	actualDiagnostics := GoSdkToTfSdkStruct(context.Background(), goSdkStruct, tfSdkStruct)
	expectedDiagnostics := diag.Diagnostics{diag.NewErrorDiagnostic(goSdkToTfSdkStructConversionFailureMessage, "please provide a pointer for the tfsdk struct, got DummyTfSdk")}
	assert.True(t, actualDiagnostics.HasError())
	assert.True(t, actualDiagnostics.Equal(expectedDiagnostics))
}

var tests = []struct {
	name        string
	tfSdkStruct DummyTfSdk
	goSdkStruct DummyGoSdk
}{
	{
		"string conversion",
		DummyTfSdk{Description: types.StringValue("abc")},
		DummyGoSdk{Description: "abc", ForceSendFields: []string{"Description"}},
	},
	{
		"bool conversion",
		DummyTfSdk{Enabled: types.BoolValue(true)},
		DummyGoSdk{Enabled: true, ForceSendFields: []string{"Enabled"}},
	},
	{
		"int64 conversion",
		DummyTfSdk{Workers: types.Int64Value(123)},
		DummyGoSdk{Workers: 123, ForceSendFields: []string{"Workers"}},
	},
	{
		"float64 conversion",
		DummyTfSdk{Floats: types.Float64Value(1.1)},
		DummyGoSdk{Floats: 1.1, ForceSendFields: []string{"Floats"}},
	},
	{
		"string zero value conversion",
		DummyTfSdk{Description: types.StringValue("")},
		DummyGoSdk{Description: "", ForceSendFields: []string{"Description"}},
	},
	{
		"bool zero value conversion",
		DummyTfSdk{Enabled: types.BoolValue(false)},
		DummyGoSdk{Enabled: false, ForceSendFields: []string{"Enabled"}},
	},
	{
		"int64 zero value conversion",
		DummyTfSdk{Workers: types.Int64Value(0)},
		DummyGoSdk{Workers: 0, ForceSendFields: []string{"Workers"}},
	},
	{
		"float64 zero value conversion",
		DummyTfSdk{Floats: types.Float64Value(0)},
		DummyGoSdk{Floats: 0, ForceSendFields: []string{"Floats"}},
	},
	{
		"tf null value conversion",
		DummyTfSdk{Workers: types.Int64Null()},
		DummyGoSdk{},
	},
	{
		"enum conversion",
		DummyTfSdk{EnumField: types.StringValue("TEST_ENUM_A")},
		DummyGoSdk{EnumField: TestEnumA},
	},
	{
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
	},
	{
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
	},
	{
		"list conversion",
		DummyTfSdk{Repeated: []types.Int64{types.Int64Value(12), types.Int64Value(34)}},
		DummyGoSdk{Repeated: []int64{12, 34}},
	},
	{
		"map conversion",
		DummyTfSdk{Attributes: map[string]types.String{"key": types.StringValue("value")}},
		DummyGoSdk{Attributes: map[string]string{"key": "value"}},
	},
	{
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
	},
	{
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
	},
}

func TestConverter(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) { RunConverterTest(t, test.name, test.tfSdkStruct, test.goSdkStruct) })
	}
}
