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
	datasource_schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resource_schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

type TestStringTfSdk struct {
	Description types.String `tfsdk:"description"`
}

func (TestStringTfSdk) ApplySchemaCustomizations(attrs map[string]AttributeBuilder) map[string]AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	return attrs
}

type TestBoolTfSdk struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

func (TestBoolTfSdk) ApplySchemaCustomizations(attrs map[string]AttributeBuilder) map[string]AttributeBuilder {
	attrs["enabled"] = attrs["enabled"].SetRequired()
	return attrs
}

type TestIntTfSdk struct {
	Workers types.Int64 `tfsdk:"workers"`
}

func (TestIntTfSdk) ApplySchemaCustomizations(attrs map[string]AttributeBuilder) map[string]AttributeBuilder {
	attrs["workers"] = attrs["workers"].SetOptional()
	return attrs
}

type TestNamespaceResourceTfSdk struct {
	Namespace
}

func (a TestNamespaceResourceTfSdk) ApplySchemaCustomizations(s map[string]AttributeBuilder) map[string]AttributeBuilder {
	s["provider_config"] = s["provider_config"].SetOptional()
	return s
}

func (a TestNamespaceResourceTfSdk) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"provider_config": reflect.TypeOf(ProviderConfig{}),
	}
}

type TestNamespaceDataSourceTfSdk struct {
	Namespace
}

func (a TestNamespaceDataSourceTfSdk) ApplySchemaCustomizations(s map[string]AttributeBuilder) map[string]AttributeBuilder {
	s["provider_config"] = s["provider_config"].SetOptional()
	return s
}

func (a TestNamespaceDataSourceTfSdk) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

type TestComputedTfSdk struct {
	ComputedTag  types.String `tfsdk:"computedtag"`
	MultipleTags types.String `tfsdk:"multipletags"`
	NonComputed  types.String `tfsdk:"noncomputed"`
}

func (TestComputedTfSdk) ApplySchemaCustomizations(attrs map[string]AttributeBuilder) map[string]AttributeBuilder {
	attrs["computedtag"] = attrs["computedtag"].SetComputed()
	attrs["multipletags"] = attrs["multipletags"].SetComputed().SetOptional()
	attrs["noncomputed"] = attrs["noncomputed"].SetOptional()
	return attrs
}

type TestFloatTfSdk struct {
	Float types.Float64 `tfsdk:"float"`
}

func (TestFloatTfSdk) ApplySchemaCustomizations(attrs map[string]AttributeBuilder) map[string]AttributeBuilder {
	attrs["float"] = attrs["float"].SetOptional()
	return attrs
}

type TestListTfSdk struct {
	Repeated types.List `tfsdk:"repeated"`
}

func (TestListTfSdk) ApplySchemaCustomizations(attrs map[string]AttributeBuilder) map[string]AttributeBuilder {
	attrs["repeated"] = attrs["repeated"].SetOptional()
	return attrs
}

func (TestListTfSdk) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"repeated": reflect.TypeOf(types.Int64{}),
	}
}

type TestMapTfSdk struct {
	Attributes types.Map `tfsdk:"attributes"`
}

func (TestMapTfSdk) ApplySchemaCustomizations(attrs map[string]AttributeBuilder) map[string]AttributeBuilder {
	attrs["attributes"] = attrs["attributes"].SetOptional()
	return attrs
}

func (TestMapTfSdk) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"attributes": reflect.TypeOf(types.String{}),
	}
}

type TestObjectTfSdk struct {
	Object types.Object `tfsdk:"object"`
}

func (TestObjectTfSdk) ApplySchemaCustomizations(attrs map[string]AttributeBuilder) map[string]AttributeBuilder {
	attrs["object"] = attrs["object"].SetOptional()
	return attrs
}

func (TestObjectTfSdk) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"object": reflect.TypeOf(DummyNested{}),
	}
}

type TestNestedListTfSdk struct {
	NestedList types.List `tfsdk:"nested_list"`
}

func (TestNestedListTfSdk) ApplySchemaCustomizations(attrs map[string]AttributeBuilder) map[string]AttributeBuilder {
	attrs["nested_list"] = attrs["nested_list"].SetOptional()
	return attrs
}

func (TestNestedListTfSdk) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"nested_list": reflect.TypeOf(DummyNested{}),
	}
}

type DummyNested struct {
	Name    types.String `tfsdk:"name"`
	Enabled types.Bool   `tfsdk:"enabled"`
}

func (DummyNested) ApplySchemaCustomizations(attrs map[string]AttributeBuilder) map[string]AttributeBuilder {
	attrs["name"] = attrs["name"].SetOptional()
	attrs["enabled"] = attrs["enabled"].SetOptional()
	return attrs
}

type TestNestedMapTfSdk struct {
	NestedMap types.Map `tfsdk:"nested_map"`
}

func (TestNestedMapTfSdk) ApplySchemaCustomizations(attrs map[string]AttributeBuilder) map[string]AttributeBuilder {
	attrs["nested_map"] = attrs["nested_map"].SetOptional()
	return attrs
}

func (TestNestedMapTfSdk) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"nested_map": reflect.TypeOf(DummyNested{}),
	}
}

var dummyType = tfcommon.NewObjectTyper(DummyNested{}).Type(context.Background()).(types.ObjectType)

type TestDeprecatedTfTags struct {
	Foo types.String `tfsdk:"foo" tf:"computed"`
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
	{
		"namespace resource conversion",
		TestNamespaceResourceTfSdk{Namespace: Namespace{ProviderConfig: types.ObjectValueMust(ProviderConfig{}.Type(context.Background()).(types.ObjectType).AttrTypes, map[string]attr.Value{
			"workspace_id": types.StringValue("1234567890"),
		})}},
	},
	{
		"namespace data source conversion",
		TestNamespaceDataSourceTfSdk{Namespace: Namespace{ProviderConfig: types.ObjectValueMust(ProviderConfigData{}.Type(context.Background()).(types.ObjectType).AttrTypes, map[string]attr.Value{
			"workspace_id": types.StringValue("1234567890"),
		})}},
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

func TestStructToSchemaNamespace(t *testing.T) {
	// Test that provider_config is an optional field.
	scm := ResourceStructToSchema(context.Background(), TestNamespaceResourceTfSdk{}, nil)
	assert.True(t, scm.Attributes["provider_config"].IsOptional())

	data_scm := DataSourceStructToSchema(context.Background(), TestNamespaceDataSourceTfSdk{}, nil)
	assert.True(t, data_scm.Attributes["provider_config"].IsOptional())

	// Test that workspace_id is a required field.
	scm = ResourceStructToSchema(context.Background(), TestNamespaceResourceTfSdk{}, nil)
	assert.True(t, scm.Attributes["provider_config"].(resource_schema.SingleNestedAttribute).Attributes["workspace_id"].IsRequired())

	data_scm = DataSourceStructToSchema(context.Background(), TestNamespaceDataSourceTfSdk{}, nil)
	assert.True(t, data_scm.Attributes["provider_config"].(datasource_schema.SingleNestedAttribute).Attributes["workspace_id"].IsRequired())
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
	Description types.List `tfsdk:"description"`
}

func (TestTfSdkList) ApplySchemaCustomizations(attrs map[string]AttributeBuilder) map[string]AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	return attrs
}

type TestTfSdkMapWithoutMetadata struct {
	Description types.Map `tfsdk:"description"`
}

func (TestTfSdkMapWithoutMetadata) ApplySchemaCustomizations(attrs map[string]AttributeBuilder) map[string]AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	return attrs
}

type TestSliceOfSlice struct {
	NestedList [][]string `tfsdk:"nested_list"`
}

func (TestSliceOfSlice) ApplySchemaCustomizations(attrs map[string]AttributeBuilder) map[string]AttributeBuilder {
	attrs["nested_list"] = attrs["nested_list"].SetOptional()
	return attrs
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
	// Test that ComputedTag field is computed
	scm := ResourceStructToSchema(context.Background(), TestComputedTfSdk{}, nil)
	assert.True(t, scm.Attributes["computedtag"].IsComputed())
	// Computed fields can never be required
	assert.False(t, scm.Attributes["computedtag"].IsRequired())

	// Test that MultipleTags field is computed and optional
	assert.True(t, scm.Attributes["multipletags"].IsComputed())
	assert.True(t, scm.Attributes["multipletags"].IsOptional())

	// Test that NonComputed field is not computed
	assert.True(t, !scm.Attributes["noncomputed"].IsComputed())
}

func TestDeprecatedTagsPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for TestDeprecatedTfTags")
		}
	}()
	ResourceStructToSchema(context.Background(), TestDeprecatedTfTags{}, nil)
}
