package tfschema

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

type TestStringTfSdk struct {
	Description types.String `tfsdk:"description" tf:"optional"`
}

type TestBoolTfSdk struct {
	Enabled types.Bool `tfsdk:"enabled" tf:""`
}

type TestIntTfSdk struct {
	Workers types.Int64 `tfsdk:"workers" tf:"optional"`
}

type TestComputedTfSdk struct {
	ComputedTag  types.String `tfsdk:"computedtag" tf:"computed"`
	MultipleTags types.String `tfsdk:"multipletags" tf:"computed,optional"`
	NonComputed  types.String `tfsdk:"noncomputed" tf:"optional"`
}

type TestFloatTfSdk struct {
	Float types.Float64 `tfsdk:"float" tf:"optional"`
}

type TestListTfSdk struct {
	Repeated []types.Int64 `tfsdk:"repeated" tf:"optional"`
}

type TestMapTfSdk struct {
	Attributes map[string]types.String `tfsdk:"attributes" tf:"optional"`
}

type TestNestedListTfSdk struct {
	NestedList []DummyNested `tfsdk:"nested_list" tf:"optional"`
}

type DummyNested struct {
	Name    types.String `tfsdk:"name" tf:"optional"`
	Enabled types.Bool   `tfsdk:"enabled" tf:"optional"`
}

type DummyDoubleNested struct {
	Nested *DummyNested `tfsdk:"nested" tf:"optional"`
}

type TestNestedMapTfSdk struct {
	NestedMap map[string]DummyNested `tfsdk:"nested_map" tf:"optional"`
}

type TestPointerTfSdk struct {
	Nested *DummyNested `tfsdk:"nested" tf:"optional"`
}

type TestNestedPointerTfSdk struct {
	Nested DummyDoubleNested `tfsdk:"nested" tf:"optional"`
}

var tests = []struct {
	name       string
	testStruct any
}{
	{
		"string conversion",
		TestStringTfSdk{Description: types.StringValue("abc")},
	},
	{
		"bool conversion",
		TestBoolTfSdk{Enabled: types.BoolValue(true)},
	},
	{
		"int conversion",
		TestIntTfSdk{Workers: types.Int64Value(1)},
	},
	{
		"float conversion",
		TestFloatTfSdk{Float: types.Float64Value(1.1)},
	},
	{
		"list conversion",
		TestListTfSdk{Repeated: []types.Int64{types.Int64Value(12), types.Int64Value(34)}},
	},
	{
		"map conversion",
		TestMapTfSdk{Attributes: map[string]types.String{"key": types.StringValue("value")}},
	},
	{
		"nested list conversion",
		TestNestedListTfSdk{NestedList: []DummyNested{
			{
				Name:    types.StringValue("abc"),
				Enabled: types.BoolValue(true),
			},
			{
				Name:    types.StringValue("def"),
				Enabled: types.BoolValue(false),
			},
		}},
	},
	{
		"nested map conversion",
		TestNestedMapTfSdk{NestedMap: map[string]DummyNested{
			"key1": {
				Name:    types.StringValue("abc"),
				Enabled: types.BoolValue(true),
			},
			"key2": {
				Name:    types.StringValue("def"),
				Enabled: types.BoolValue(false),
			},
		}},
	},
	{
		"pointer to a struct conversion",
		TestPointerTfSdk{
			&DummyNested{
				Name:    types.StringValue("def"),
				Enabled: types.BoolValue(true),
			},
		},
	},
	{
		"nested pointer to a struct conversion",
		TestNestedPointerTfSdk{
			DummyDoubleNested{
				Nested: &DummyNested{
					Name:    types.StringValue("def"),
					Enabled: types.BoolValue(true),
				},
			},
		},
	},
}

// StructToSchemaConversionTestCase runs a single test case to verify StructToSchema works for both data source and resource.
func StructToSchemaConversionTestCase(t *testing.T, description string, testStruct any) {
	scm := ResourceStructToSchema(testStruct, nil)
	state := tfsdk.State{
		Schema: scm,
	}
	// Assert we can properly set the state, this means the schema and the struct are consistent.
	assert.True(t, !state.Set(context.Background(), testStruct).HasError(), fmt.Sprintf("ResourceStructToSchema - %s", description))

	data_scm := DataSourceStructToSchema(testStruct, nil)
	data_state := tfsdk.State{
		Schema: data_scm,
	}
	// Assert we can properly set the state, this means the schema and the struct are consistent.
	assert.True(t, !data_state.Set(context.Background(), testStruct).HasError(), fmt.Sprintf("DataSourceStructToSchema - %s", description))
}

func TestStructToSchemaConversion(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) { StructToSchemaConversionTestCase(t, test.name, test.testStruct) })
	}
}

func TestStructToSchemaOptionalVsRequiredField(t *testing.T) {
	// Test that description is an optional field.
	scm := ResourceStructToSchema(TestStringTfSdk{}, nil)
	assert.True(t, scm.Attributes["description"].IsOptional())
	assert.True(t, !scm.Attributes["description"].IsRequired())

	// Test that enabled is a required field.
	data_scm := DataSourceStructToSchema(TestBoolTfSdk{}, nil)
	assert.True(t, !data_scm.Attributes["enabled"].IsOptional())
	assert.True(t, data_scm.Attributes["enabled"].IsRequired())
}

func testStructToSchemaPanics(t *testing.T, testStruct any, expectedError string) {
	defer func() {
		err := recover()
		if err == nil {
			t.Fatalf("did not fail with error")
		}
		var errMsg string
		switch e := err.(type) {
		case error:
			errMsg = e.Error()
		case string:
			errMsg = e
		default:
			t.Fatalf("recovered panic is of unknown type: %T, value: %v", e, e)
		}

		if !strings.Contains(errMsg, expectedError) {
			t.Fatalf("error %s did not include expected error message %q", errMsg, expectedError)
		}
	}()
	ResourceStructToSchema(testStruct, nil)
}

type TestTfSdkList struct {
	Description types.List `tfsdk:"description" tf:"optional"`
}

type TestTfSdkMap struct {
	Description types.Map `tfsdk:"description" tf:"optional"`
}

type TestSliceOfSlice struct {
	NestedList [][]string `tfsdk:"nested_list" tf:"optional"`
}

type TestMapOfMap struct {
	NestedMap map[string]map[string]string `tfsdk:"nested_map" tf:"optional"`
}

var error_tests = []struct {
	name          string
	testStruct    any
	expectedError string
}{
	{
		"tf list conversion",
		TestTfSdkList{},
		fmt.Sprintf("types.List should never be used in tfsdk structs. %s", common.TerraformBugErrorMessage),
	},
	{
		"tf map conversion",
		TestTfSdkMap{},
		fmt.Sprintf("types.Map should never be used in tfsdk structs. %s", common.TerraformBugErrorMessage),
	},
	{
		"non-struct conversion",
		"Abc",
		fmt.Sprintf("schema value of Struct is expected, but got string: \"Abc\". %s", common.TerraformBugErrorMessage),
	},
	{
		"slice of slice conversion",
		TestSliceOfSlice{},
		fmt.Sprintf("unsupported slice value for nested_list: slice. %s", common.TerraformBugErrorMessage),
	},
	{
		"map of map conversion",
		TestMapOfMap{},
		fmt.Sprintf("unsupported map value for nested_map: map. %s", common.TerraformBugErrorMessage),
	},
}

func TestStructToSchemaExpectedError(t *testing.T) {
	for _, test := range error_tests {
		t.Run(test.name, func(t *testing.T) { testStructToSchemaPanics(t, test.testStruct, test.expectedError) })
	}
}

func TestComputedField(t *testing.T) {
	// Test that ComputedTag field is computed and required
	scm := ResourceStructToSchema(TestComputedTfSdk{}, nil)
	assert.True(t, scm.Attributes["computedtag"].IsComputed())
	assert.True(t, scm.Attributes["computedtag"].IsRequired())

	// Test that MultipleTags field is computed and optional
	assert.True(t, scm.Attributes["multipletags"].IsComputed())
	assert.True(t, scm.Attributes["multipletags"].IsOptional())

	// Test that NonComputed field is not computed
	assert.True(t, !scm.Attributes["noncomputed"].IsComputed())
}
