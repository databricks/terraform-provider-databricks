package tfschema

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	tfcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/hashicorp/terraform-plugin-framework/attr"
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
	Repeated types.List `tfsdk:"repeated" tf:"optional"`
}

func (TestListTfSdk) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"repeated": reflect.TypeOf(types.Int64{}),
	}
}

type TestMapTfSdk struct {
	Attributes types.Map `tfsdk:"attributes" tf:"optional"`
}

func (TestMapTfSdk) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"attributes": reflect.TypeOf(types.String{}),
	}
}

type TestObjectTfSdk struct {
	Object types.Object `tfsdk:"object" tf:"optional"`
}

func (TestObjectTfSdk) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"object": reflect.TypeOf(DummyNested{}),
	}
}

type TestNestedListTfSdk struct {
	NestedList types.List `tfsdk:"nested_list" tf:"optional"`
}

func (TestNestedListTfSdk) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"nested_list": reflect.TypeOf(DummyNested{}),
	}
}

type DummyNested struct {
	Name    types.String `tfsdk:"name" tf:"optional"`
	Enabled types.Bool   `tfsdk:"enabled" tf:"optional"`
}

type TestNestedMapTfSdk struct {
	NestedMap types.Map `tfsdk:"nested_map" tf:"optional"`
}

func (TestNestedMapTfSdk) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"nested_map": reflect.TypeOf(DummyNested{}),
	}
}

var dummyType = tfcommon.NewObjectTyper(DummyNested{}).Type(context.Background()).(types.ObjectType)

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
		TestListTfSdk{Repeated: types.ListValueMust(types.Int64Type, []attr.Value{types.Int64Value(12), types.Int64Value(34)})},
	},
	{
		"map conversion",
		TestMapTfSdk{Attributes: types.MapValueMust(types.StringType, map[string]attr.Value{"key": types.StringValue("value")})},
	},
	{
		"object conversion",
		TestObjectTfSdk{Object: types.ObjectValueMust(dummyType.AttrTypes, map[string]attr.Value{
			"name":    types.StringValue("abc"),
			"enabled": types.BoolValue(true),
		})},
	},
	{
		"nested list conversion",
		TestNestedListTfSdk{NestedList: types.ListValueMust(dummyType,
			[]attr.Value{
				types.ObjectValueMust(dummyType.AttrTypes, map[string]attr.Value{
					"name":    types.StringValue("abc"),
					"enabled": types.BoolValue(true),
				}),
				types.ObjectValueMust(dummyType.AttrTypes, map[string]attr.Value{
					"name":    types.StringValue("def"),
					"enabled": types.BoolValue(false),
				}),
			}),
		},
	},
	{
		"nested map conversion",
		TestNestedMapTfSdk{NestedMap: types.MapValueMust(dummyType, map[string]attr.Value{
			"key1": types.ObjectValueMust(dummyType.AttrTypes, map[string]attr.Value{
				"name":    types.StringValue("abc"),
				"enabled": types.BoolValue(true),
			}),
			"key2": types.ObjectValueMust(dummyType.AttrTypes, map[string]attr.Value{
				"name":    types.StringValue("def"),
				"enabled": types.BoolValue(false),
			}),
		}),
		},
	},
}

// StructToSchemaConversionTestCase runs a single test case to verify StructToSchema works for both data source and resource.
func StructToSchemaConversionTestCase(t *testing.T, description string, testStruct any) {
	scm := ResourceStructToSchema(context.Background(), testStruct, nil)
	state := tfsdk.State{
		Schema: scm,
	}
	// Assert we can properly set the state, this means the schema and the struct are consistent.
	d := state.Set(context.Background(), testStruct)
	if d.HasError() {
		t.Errorf("ResourceStructToSchema - %s: %s", description, tfcommon.DiagToString(d))
	}

	data_scm := DataSourceStructToSchema(context.Background(), testStruct, nil)
	data_state := tfsdk.State{
		Schema: data_scm,
	}
	// Assert we can properly set the state, this means the schema and the struct are consistent.
	d = data_state.Set(context.Background(), testStruct)
	if d.HasError() {
		t.Errorf("DataSourceStructToSchema - %s: %s", description, tfcommon.DiagToString(d))
	}
}

func TestStructToSchemaConversion(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) { StructToSchemaConversionTestCase(t, test.name, test.testStruct) })
	}
}

func TestStructToSchemaOptionalVsRequiredField(t *testing.T) {
	// Test that description is an optional field.
	scm := ResourceStructToSchema(context.Background(), TestStringTfSdk{}, nil)
	assert.True(t, scm.Attributes["description"].IsOptional())
	assert.True(t, !scm.Attributes["description"].IsRequired())

	// Test that enabled is a required field.
	data_scm := DataSourceStructToSchema(context.Background(), TestBoolTfSdk{}, nil)
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
	ResourceStructToSchema(context.Background(), testStruct, nil)
}

type TestTfSdkList struct {
	Description types.List `tfsdk:"description" tf:"optional"`
}

type TestTfSdkMapWithoutMetadata struct {
	Description types.Map `tfsdk:"description" tf:"optional"`
}

type TestSliceOfSlice struct {
	NestedList [][]string `tfsdk:"nested_list" tf:"optional"`
}

var error_tests = []struct {
	name          string
	testStruct    any
	expectedError string
}{
	{
		"tfsdk struct without complex field types conversion",
		TestTfSdkMapWithoutMetadata{},
		fmt.Sprintf("complex field types not provided for type: tfschema.TestTfSdkMapWithoutMetadata. %s", common.TerraformBugErrorMessage),
	},
	{
		"non-struct conversion",
		"Abc",
		fmt.Sprintf("schema value of Struct is expected, but got string: \"Abc\". %s", common.TerraformBugErrorMessage),
	},
	{
		"slice of slice conversion",
		TestSliceOfSlice{},
		fmt.Sprintf("unexpected type [][]string in tfsdk structs, expected a plugin framework value type. %s", common.TerraformBugErrorMessage),
	},
}

func TestStructToSchemaExpectedError(t *testing.T) {
	for _, test := range error_tests {
		t.Run(test.name, func(t *testing.T) { testStructToSchemaPanics(t, test.testStruct, test.expectedError) })
	}
}

func TestComputedField(t *testing.T) {
	// Test that ComputedTag field is computed and required
	scm := ResourceStructToSchema(context.Background(), TestComputedTfSdk{}, nil)
	assert.True(t, scm.Attributes["computedtag"].IsComputed())
	assert.True(t, scm.Attributes["computedtag"].IsRequired())

	// Test that MultipleTags field is computed and optional
	assert.True(t, scm.Attributes["multipletags"].IsComputed())
	assert.True(t, scm.Attributes["multipletags"].IsOptional())

	// Test that NonComputed field is not computed
	assert.True(t, !scm.Attributes["noncomputed"].IsComputed())
}
