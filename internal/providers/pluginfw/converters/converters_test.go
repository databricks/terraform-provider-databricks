package converters

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	tfcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

type DummyTfSdk struct {
	Enabled           types.Bool    `tfsdk:"enabled" tf:"optional"`
	Workers           types.Int64   `tfsdk:"workers" tf:""`
	Floats            types.Float64 `tfsdk:"floats" tf:""`
	Description       types.String  `tfsdk:"description" tf:""`
	Tasks             types.String  `tfsdk:"task" tf:"optional"`
	NoPointerNested   types.List    `tfsdk:"no_pointer_nested" tf:"optional"`
	NestedList        types.List    `tfsdk:"nested_list" tf:"optional"`
	NestedPointerList types.List    `tfsdk:"nested_pointer_list" tf:"optional"`
	Map               types.Map     `tfsdk:"map" tf:"optional"`
	NestedMap         types.Map     `tfsdk:"nested_map" tf:"optional"`
	Repeated          types.List    `tfsdk:"repeated" tf:"optional"`
	Attributes        types.Map     `tfsdk:"attributes" tf:"optional"`
	EnumField         types.String  `tfsdk:"enum_field" tf:"optional"`
	AdditionalField   types.String  `tfsdk:"additional_field" tf:"optional"`
	DistinctField     types.String  `tfsdk:"distinct_field" tf:"optional"`
	SliceStructPtr    types.List    `tfsdk:"slice_struct_ptr" tf:"optional"`
	Irrelevant        types.String  `tfsdk:"-"`
	Object            types.Object  `tfsdk:"object" tf:"optional"`
}

func (DummyTfSdk) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"no_pointer_nested":   reflect.TypeOf(DummyNestedTfSdk{}),
		"nested_list":         reflect.TypeOf(DummyNestedTfSdk{}),
		"nested_pointer_list": reflect.TypeOf(DummyNestedTfSdk{}),
		"map":                 reflect.TypeOf(types.String{}),
		"nested_map":          reflect.TypeOf(DummyNestedTfSdk{}),
		"repeated":            reflect.TypeOf(types.Int64{}),
		"attributes":          reflect.TypeOf(types.String{}),
		"slice_struct":        reflect.TypeOf(DummyNestedTfSdk{}),
		"slice_struct_ptr":    reflect.TypeOf(DummyNestedTfSdk{}),
		"object":              reflect.TypeOf(DummyNestedTfSdk{}),
	}
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
	SliceStructPtr    *DummyNestedGoSdk           `json:"slice_struct_ptr"`
	ForceSendFields   []string                    `json:"-"`
	Object            DummyNestedGoSdk            `json:"object"`
}

type DummyNestedGoSdk struct {
	Name            string   `json:"name"`
	Enabled         bool     `json:"enabled"`
	ForceSendFields []string `json:"-"`
}

// This function is used to populate empty fields in the tfsdk struct with null values.
// This is required because the Go->TF conversion function instantiates list, map, and
// object fields with empty values, which are not equal to null values in the tfsdk struct.
func populateEmptyFields(c DummyTfSdk) DummyTfSdk {
	complexFields := c.GetComplexFieldTypes(context.Background())
	v := reflect.ValueOf(&c).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		// If the field is a simple type, the zero value is OK.
		switch field.Type() {
		case reflect.TypeOf(types.Bool{}), reflect.TypeOf(types.Int64{}), reflect.TypeOf(types.Float64{}), reflect.TypeOf(types.String{}):
			continue
		}
		if !field.IsZero() {
			continue
		}

		tfsdkName := v.Type().Field(i).Tag.Get("tfsdk")
		complexType, ok := complexFields[tfsdkName]
		if !ok {
			continue
		}

		var typ attr.Type
		switch complexType {
		case reflect.TypeOf(types.Bool{}):
			typ = types.BoolType
		case reflect.TypeOf(types.Int64{}):
			typ = types.Int64Type
		case reflect.TypeOf(types.Float64{}):
			typ = types.Float64Type
		case reflect.TypeOf(types.String{}):
			typ = types.StringType
		default:
			innerVal := reflect.New(complexType).Elem().Interface()
			typ = tfcommon.NewObjectValuable(innerVal).Type(context.Background())
		}
		switch field.Type() {
		case reflect.TypeOf(types.List{}):
			value := types.ListNull(typ)
			field.Set(reflect.ValueOf(value))
		case reflect.TypeOf(types.Map{}):
			value := types.MapNull(typ)
			field.Set(reflect.ValueOf(value))
		case reflect.TypeOf(types.Object{}):
			objectType := typ.(types.ObjectType)
			value := types.ObjectNull(objectType.AttrTypes)
			field.Set(reflect.ValueOf(value))
		}
	}
	return v.Interface().(DummyTfSdk)
}

// Function to construct individual test case with a pair of matching tfSdkStruct and gosdkStruct.
// Verifies that the conversion both ways are working as expected.
func RunConverterTest(t *testing.T, description string, tfSdkStruct DummyTfSdk, goSdkStruct DummyGoSdk) {
	convertedGoSdkStruct := DummyGoSdk{}
	d := TfSdkToGoSdkStruct(context.Background(), tfSdkStruct, &convertedGoSdkStruct)
	if d.HasError() {
		t.Errorf("tfsdk to gosdk conversion: %s", pluginfwcommon.DiagToString(d))
	}
	assert.Equal(t, goSdkStruct, convertedGoSdkStruct, fmt.Sprintf("tfsdk to gosdk conversion - %s", description))

	convertedTfSdkStruct := DummyTfSdk{}
	d = GoSdkToTfSdkStruct(context.Background(), goSdkStruct, &convertedTfSdkStruct)
	if d.HasError() {
		t.Errorf("gosdk to tfsdk conversion: %s", pluginfwcommon.DiagToString(d))
	}
	assert.Equal(t, populateEmptyFields(tfSdkStruct), convertedTfSdkStruct, fmt.Sprintf("gosdk to tfsdk conversion - %s", description))
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

var dummyType = tfcommon.NewObjectValuable(DummyNestedTfSdk{}).Type(context.Background()).(types.ObjectType)

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
		DummyTfSdk{NoPointerNested: types.ListValueMust(
			dummyType, []attr.Value{
				types.ObjectValueMust(dummyType.AttrTypes, map[string]attr.Value{
					"name":    types.StringValue("def"),
					"enabled": types.BoolValue(true),
				}),
			}),
		},
		DummyGoSdk{NoPointerNested: DummyNestedGoSdk{
			Name:            "def",
			Enabled:         true,
			ForceSendFields: []string{"Name", "Enabled"},
		}, ForceSendFields: []string{"NoPointerNested"}},
	},
	{
		"list conversion",
		DummyTfSdk{Repeated: types.ListValueMust(types.Int64Type, []attr.Value{types.Int64Value(12), types.Int64Value(34)})},
		DummyGoSdk{Repeated: []int64{12, 34}},
	},
	{
		"map conversion",
		DummyTfSdk{Attributes: types.MapValueMust(types.StringType, map[string]attr.Value{"key": types.StringValue("value")})},
		DummyGoSdk{Attributes: map[string]string{"key": "value"}, ForceSendFields: []string{"Attributes"}},
	},
	{
		"nested list conversion",
		DummyTfSdk{NestedList: types.ListValueMust(dummyType,
			[]attr.Value{
				types.ObjectValueMust(dummyType.AttrTypes, map[string]attr.Value{
					"name":    types.StringValue("def"),
					"enabled": types.BoolValue(true),
				}),
				types.ObjectValueMust(dummyType.AttrTypes, map[string]attr.Value{
					"name":    types.StringValue("def"),
					"enabled": types.BoolValue(true),
				}),
			}),
		},
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
		DummyTfSdk{NestedMap: types.MapValueMust(dummyType, map[string]attr.Value{
			"key1": types.ObjectValueMust(dummyType.AttrTypes, map[string]attr.Value{
				"name":    types.StringValue("abc"),
				"enabled": types.BoolValue(true),
			}),
			"key2": types.ObjectValueMust(dummyType.AttrTypes, map[string]attr.Value{
				"name":    types.StringValue("def"),
				"enabled": types.BoolValue(false),
			}),
		})},
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
		}, ForceSendFields: []string{"NestedMap"}},
	},
	{
		"list representation of struct pointer conversion", // we use list with one element in the tfsdk to represent struct in gosdk
		DummyTfSdk{SliceStructPtr: types.ListValueMust(dummyType,
			[]attr.Value{
				types.ObjectValueMust(dummyType.AttrTypes, map[string]attr.Value{
					"name":    types.StringValue("def"),
					"enabled": types.BoolValue(true),
				}),
			}),
		},
		DummyGoSdk{SliceStructPtr: &DummyNestedGoSdk{
			Name:            "def",
			Enabled:         true,
			ForceSendFields: []string{"Name", "Enabled"},
		}, ForceSendFields: []string{"SliceStructPtr"}},
	},
	{
		"object conversion",
		DummyTfSdk{Object: types.ObjectValueMust(dummyType.AttrTypes, map[string]attr.Value{
			"name":    types.StringValue("def"),
			"enabled": types.BoolValue(true),
		})},
		DummyGoSdk{Object: DummyNestedGoSdk{
			Name:            "def",
			Enabled:         true,
			ForceSendFields: []string{"Name", "Enabled"},
		}, ForceSendFields: []string{"Object"}},
	},
}

func TestConverter(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) { RunConverterTest(t, test.name, test.tfSdkStruct, test.goSdkStruct) })
	}
}
