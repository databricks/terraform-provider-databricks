package sdkgoutility

import (
	"encoding/json"
	"io"
	"reflect"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/sdk-go/core/types"
	modelregistry "github.com/databricks/sdk-go/modelregistry/v1"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Every identifier of this file starts with the name of the helper, because the other helpers of the
// package declare their own. The types have no MarshalJSON method, so FieldName gives snake_case names.

type dataToStructEnum string

type dataToStructLeaf struct {
	Name  *string
	Count *int64
}

type dataToStructBranch struct {
	Label  *string
	Leaf   *dataToStructLeaf
	Inline dataToStructLeaf
	Leaves []dataToStructLeaf
}

// dataToStructShapes has a field of every shape that DataToStruct sets.
type dataToStructShapes struct {
	StringPointer  *string
	BoolPointer    *bool
	IntPointer     *int
	Int32Pointer   *int32
	Int64Pointer   *int64
	Float32Pointer *float32
	Float64Pointer *float64
	EnumPointer    *dataToStructEnum
	StringValue    string
	BoolValue      bool
	IntValue       int
	Int8Value      int8
	Int64Value     int64
	Float32Value   float32
	Float64Value   float64
	EnumValue      dataToStructEnum
	MessagePointer *dataToStructBranch
	MessageValue   dataToStructBranch
	Messages       []dataToStructBranch
	Strings        []string
	Bools          []bool
	Ints           []int
	Int64s         []int64
	Floats         []float64
	Enums          []dataToStructEnum
	Labels         map[string]string
	EnumLabels     map[dataToStructEnum]dataToStructEnum
}

// dataToStructOne has one field, of the type T and with the key "field".
type dataToStructOne[T any] struct{ Field T }

// dataToStructChoice has the shape of the interface of a oneof group.
type dataToStructChoice interface{ isDataToStructChoice() }

func dataToStructPtr[T any](v T) *T { return &v }

// dataToStructScalars returns the schema of an optional string, bool, integer and float.
func dataToStructScalars() (str, boolean, integer, float *schema.Schema) {
	optional := func(t schema.ValueType) *schema.Schema { return &schema.Schema{Type: t, Optional: true} }
	return optional(schema.TypeString), optional(schema.TypeBool), optional(schema.TypeInt), optional(schema.TypeFloat)
}

// dataToStructOf is the schema of a list, a set or a map of elem: a
// *schema.Schema for scalars and a *schema.Resource for blocks.
func dataToStructOf(t schema.ValueType, elem any) *schema.Schema {
	return &schema.Schema{Type: t, Optional: true, Elem: elem}
}
func dataToStructList(elem any) *schema.Schema { return dataToStructOf(schema.TypeList, elem) }
func dataToStructSet(elem any) *schema.Schema  { return dataToStructOf(schema.TypeSet, elem) }
func dataToStructMap(elem any) *schema.Schema  { return dataToStructOf(schema.TypeMap, elem) }
func dataToStructBlocks(block map[string]*schema.Schema) *schema.Schema {
	return dataToStructList(&schema.Resource{Schema: block})
}

func dataToStructLeafSchema() map[string]*schema.Schema {
	str, _, integer, _ := dataToStructScalars()
	return map[string]*schema.Schema{"name": str, "count": integer}
}

func dataToStructBranchSchema() map[string]*schema.Schema {
	str, _, _, _ := dataToStructScalars()
	leaves := dataToStructBlocks(dataToStructLeafSchema())
	return map[string]*schema.Schema{"label": str, "leaf": leaves, "inline": leaves, "leaves": leaves}
}

func dataToStructShapesSchema() map[string]*schema.Schema {
	str, boolean, integer, float := dataToStructScalars()
	return map[string]*schema.Schema{
		"string_pointer": str, "bool_pointer": boolean, "int_pointer": integer, "int32_pointer": integer,
		"int64_pointer": integer, "float32_pointer": float, "float64_pointer": float, "enum_pointer": str,
		"string_value": str, "bool_value": boolean, "int_value": integer, "int8_value": integer,
		"int64_value": integer, "float32_value": float, "float64_value": float, "enum_value": str,
		"message_pointer": dataToStructBlocks(dataToStructBranchSchema()), "message_value": dataToStructBlocks(dataToStructBranchSchema()),
		"messages": dataToStructBlocks(dataToStructBranchSchema()),
		"strings":  dataToStructList(str), "bools": dataToStructList(boolean), "ints": dataToStructList(integer), "int64s": dataToStructList(integer),
		"floats": dataToStructList(float), "enums": dataToStructList(str),
		"labels": dataToStructMap(str), "enum_labels": dataToStructMap(str),
	}
}

func TestDataToStruct_SetsEveryShape(t *testing.T) {
	s := dataToStructShapesSchema()
	leaf := map[string]any{"name": "leaf", "count": 7}
	branch := map[string]any{
		"label": "branch", "leaf": []any{leaf}, "inline": []any{leaf},
		"leaves": []any{leaf, map[string]any{"name": "second"}},
	}
	d := schema.TestResourceDataRaw(t, s, map[string]any{
		"string_pointer": "a", "bool_pointer": true, "int_pointer": 1, "int32_pointer": 2, "int64_pointer": 3,
		"float32_pointer": 1.5, "float64_pointer": 2.5, "enum_pointer": "ON",
		"string_value": "b", "bool_value": true, "int_value": 4, "int8_value": 5, "int64_value": 6,
		"float32_value": 3.5, "float64_value": 4.5, "enum_value": "OFF",
		"message_pointer": []any{branch}, "message_value": []any{branch},
		"messages": []any{branch, map[string]any{"label": "other"}},
		"strings":  []any{"x", "y"}, "bools": []any{true, false}, "ints": []any{1, 0}, "int64s": []any{2, 3},
		"floats": []any{0.5, 0.0}, "enums": []any{"ON", "OFF"},
		"labels": map[string]any{"k": "v", "empty": ""}, "enum_labels": map[string]any{"ON": "OFF"},
	})
	wantLeaf := dataToStructLeaf{Name: dataToStructPtr("leaf"), Count: dataToStructPtr(int64(7))}
	wantBranch := dataToStructBranch{
		Label: dataToStructPtr("branch"), Leaf: &wantLeaf, Inline: wantLeaf,
		Leaves: []dataToStructLeaf{wantLeaf, {Name: dataToStructPtr("second")}},
	}
	var got dataToStructShapes
	require.NoError(t, DataToStruct(d, s, &got))
	assert.Equal(t, dataToStructShapes{
		StringPointer: dataToStructPtr("a"), BoolPointer: dataToStructPtr(true), IntPointer: dataToStructPtr(1), Int32Pointer: dataToStructPtr(int32(2)),
		Int64Pointer: dataToStructPtr(int64(3)), Float32Pointer: dataToStructPtr(float32(1.5)), Float64Pointer: dataToStructPtr(2.5),
		EnumPointer: dataToStructPtr(dataToStructEnum("ON")),
		StringValue: "b", BoolValue: true, IntValue: 4, Int8Value: 5, Int64Value: 6,
		Float32Value: 3.5, Float64Value: 4.5, EnumValue: "OFF",
		MessagePointer: &wantBranch, MessageValue: wantBranch,
		Messages: []dataToStructBranch{wantBranch, {Label: dataToStructPtr("other")}},
		Strings:  []string{"x", "y"}, Bools: []bool{true, false}, Ints: []int{1, 0}, Int64s: []int64{2, 3},
		Floats: []float64{0.5, 0}, Enums: []dataToStructEnum{"ON", "OFF"},
		Labels: map[string]string{"k": "v", "empty": ""}, EnumLabels: map[dataToStructEnum]dataToStructEnum{"ON": "OFF"},
	}, got)
}

func TestDataToStruct_UnsetAndZeroStayEmpty(t *testing.T) {
	s := dataToStructShapesSchema()
	for name, config := range map[string]map[string]any{
		"nothing is set": {},
		"every value is zero": {
			"string_pointer": "", "bool_pointer": false, "int_pointer": 0, "int32_pointer": 0, "int64_pointer": 0,
			"float32_pointer": 0.0, "float64_pointer": 0.0, "enum_pointer": "",
			"string_value": "", "bool_value": false, "int_value": 0, "int8_value": 0, "int64_value": 0,
			"float32_value": 0.0, "float64_value": 0.0, "enum_value": "",
			"message_pointer": []any{}, "message_value": []any{}, "messages": []any{},
			"strings": []any{}, "bools": []any{}, "ints": []any{}, "int64s": []any{}, "floats": []any{},
			"enums": []any{}, "labels": map[string]any{}, "enum_labels": map[string]any{},
		},
	} {
		t.Run(name, func(t *testing.T) {
			var got dataToStructShapes
			require.NoError(t, DataToStruct(schema.TestResourceDataRaw(t, s, config), s, &got))
			// assert.Equal tells a nil pointer, slice or map from an empty one.
			assert.Equal(t, dataToStructShapes{}, got)
		})
	}
}

func TestDataToStruct_NestedBlocks(t *testing.T) {
	s := dataToStructShapesSchema()
	empty := map[string]any{}
	for _, tc := range []struct {
		name   string
		config map[string]any
		want   dataToStructShapes
	}{
		{
			name:   "an empty block is an empty message",
			config: map[string]any{"message_pointer": []any{empty}, "message_value": []any{empty}, "messages": []any{empty}},
			want:   dataToStructShapes{MessagePointer: &dataToStructBranch{}, Messages: []dataToStructBranch{{}}},
		},
		{
			name: "a zero value in a block stays nil",
			config: map[string]any{"message_pointer": []any{map[string]any{
				"label": "", "leaf": []any{map[string]any{"name": "n", "count": 0}},
				"leaves": []any{map[string]any{"name": "", "count": 0}},
			}}},
			want: dataToStructShapes{MessagePointer: &dataToStructBranch{Leaf: &dataToStructLeaf{Name: dataToStructPtr("n")}, Leaves: []dataToStructLeaf{{}}}},
		},
		{
			name: "the third level is read under its own key",
			config: map[string]any{"messages": []any{empty, map[string]any{
				"inline": []any{map[string]any{"count": 1}},
				"leaves": []any{empty, map[string]any{"count": 2}, map[string]any{"name": "third"}},
			}}},
			want: dataToStructShapes{Messages: []dataToStructBranch{{}, {
				Inline: dataToStructLeaf{Count: dataToStructPtr(int64(1))},
				Leaves: []dataToStructLeaf{{}, {Count: dataToStructPtr(int64(2))}, {Name: dataToStructPtr("third")}},
			}}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var got dataToStructShapes
			require.NoError(t, DataToStruct(schema.TestResourceDataRaw(t, s, tc.config), s, &got))
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestDataToStruct_ModelRegistryMessages(t *testing.T) {
	str, boolean, integer, _ := dataToStructScalars()
	for _, tc := range []struct {
		s      map[string]*schema.Schema
		config map[string]any
		result any
		want   any
	}{
		{
			s: map[string]*schema.Schema{"name": str, "description": str, "tags": dataToStructBlocks(map[string]*schema.Schema{"key": str, "value": str})},
			config: map[string]any{"name": "model", "description": "about", "tags": []any{
				map[string]any{"key": "a", "value": "1"}, map[string]any{"key": "b"},
			}},
			result: &modelregistry.CreateRegisteredModelRequest{},
			want: &modelregistry.CreateRegisteredModelRequest{
				Name: dataToStructPtr("model"), Description: dataToStructPtr("about"),
				Tags: []modelregistry.RegisteredModelTag{{Key: dataToStructPtr("a"), Value: dataToStructPtr("1")}, {Key: dataToStructPtr("b")}},
			},
		},
		{
			// Enums by value and in a slice, and messages by pointer with a bool.
			s: map[string]*schema.Schema{
				"model_name": str, "events": dataToStructList(str), "description": str, "status": str,
				"http_url_spec": dataToStructBlocks(map[string]*schema.Schema{"url": str, "enable_ssl_verification": boolean, "secret": str, "authorization": str}),
				"job_spec":      dataToStructBlocks(map[string]*schema.Schema{"job_id": str, "workspace_url": str, "access_token": str}),
			},
			config: map[string]any{
				"model_name": "model", "events": []any{"MODEL_VERSION_CREATED", "COMMENT_CREATED"}, "status": "ACTIVE",
				"http_url_spec": []any{map[string]any{"url": "https://example.com", "enable_ssl_verification": true}},
				"job_spec":      []any{map[string]any{"job_id": "1", "access_token": "token"}},
			},
			result: &modelregistry.CreateRegistryWebhookRequest{},
			want: &modelregistry.CreateRegistryWebhookRequest{
				ModelName:   dataToStructPtr("model"),
				Events:      []modelregistry.RegistryWebhookEvent{modelregistry.RegistryWebhookEvent_ModelVersionCreated, modelregistry.RegistryWebhookEvent_CommentCreated},
				Status:      modelregistry.RegistryWebhookStatus_Active,
				HttpUrlSpec: &modelregistry.HttpUrlSpec{Url: dataToStructPtr("https://example.com"), EnableSslVerification: dataToStructPtr(true)},
				JobSpec:     &modelregistry.JobSpec{JobId: dataToStructPtr("1"), AccessToken: dataToStructPtr("token")},
			},
		},
		{
			s:      map[string]*schema.Schema{"filter": str, "max_results": integer, "order_by": dataToStructList(str), "page_token": str},
			config: map[string]any{"filter": "name LIKE 'm%'", "max_results": 50, "order_by": []any{"name ASC", "timestamp DESC"}},
			result: &modelregistry.SearchRegisteredModelsRequest{},
			want: &modelregistry.SearchRegisteredModelsRequest{
				Filter: dataToStructPtr("name LIKE 'm%'"), MaxResults: dataToStructPtr(int64(50)), OrderBy: []string{"name ASC", "timestamp DESC"},
			},
		},
		{
			s:      map[string]*schema.Schema{"status_code": integer, "body": str},
			config: map[string]any{"status_code": 200, "body": "ok"},
			result: &modelregistry.TestRegistryWebhookResponse{},
			want:   &modelregistry.TestRegistryWebhookResponse{StatusCode: dataToStructPtr(200), Body: dataToStructPtr("ok")},
		},
		{
			// A false bool is left out, like the other zero values.
			s:      map[string]*schema.Schema{"name": str, "version": str, "stage": str, "archive_existing_versions": boolean},
			config: map[string]any{"name": "model", "version": "1", "stage": "Staging", "archive_existing_versions": false},
			result: &modelregistry.TransitionModelVersionStageDatabricksRequest{},
			want:   &modelregistry.TransitionModelVersionStageDatabricksRequest{Name: dataToStructPtr("model"), Version: dataToStructPtr("1"), Stage: dataToStructPtr("Staging")},
		},
	} {
		t.Run(reflect.TypeOf(tc.result).Elem().Name(), func(t *testing.T) {
			require.NoError(t, DataToStruct(schema.TestResourceDataRaw(t, tc.s, tc.config), tc.s, tc.result))
			assert.Equal(t, tc.want, tc.result)
		})
	}
}

// dataToStructSkips has the fields that DataToStruct must leave alone.
type dataToStructSkips struct {
	Known    *string
	NoEntry  *string
	NilEntry *string
	// Unsupported has no schema entry, so that its type doesn't matter.
	Unsupported any
	hidden      *string
	// The fields of an embedded struct, name and count, are not looked into.
	dataToStructLeaf
}

func TestDataToStruct_SkipsFieldsAndEntries(t *testing.T) {
	str, _, _, _ := dataToStructScalars()
	inData := map[string]*schema.Schema{"known": str, "no_entry": str, "nil_entry": str, "unsupported": str, "hidden": str, "no_field": str, "name": str}
	d := schema.TestResourceDataRaw(t, inData, map[string]any{"known": "new", "no_entry": "x", "nil_entry": "x", "unsupported": "x", "hidden": "x", "no_field": "x", "name": "x"})
	// The empty string is the FieldName of the unexported fields.
	s := map[string]*schema.Schema{"known": str, "nil_entry": nil, "hidden": str, "": str, "no_field": str, "name": str}
	got := dataToStructSkips{Known: dataToStructPtr("old"), NoEntry: dataToStructPtr("kept"), hidden: dataToStructPtr("kept")}
	want := dataToStructSkips{Known: dataToStructPtr("new"), NoEntry: dataToStructPtr("kept"), hidden: dataToStructPtr("kept")}
	require.NoError(t, DataToStruct(d, s, &got))
	assert.Equal(t, want, got)

	// A field that the data leaves unset keeps the value it came with.
	require.NoError(t, DataToStruct(schema.TestResourceDataRaw(t, inData, map[string]any{}), s, &got))
	assert.Equal(t, want, got)

	// A nil schema skips every field.
	got = dataToStructSkips{}
	require.NoError(t, DataToStruct(d, nil, &got))
	assert.Equal(t, dataToStructSkips{}, got)
}

func TestDataToStruct_ResultMustBePointerToStruct(t *testing.T) {
	str, _, _, _ := dataToStructScalars()
	s := map[string]*schema.Schema{"name": str}
	d := schema.TestResourceDataRaw(t, s, map[string]any{"name": "model"})
	request := &modelregistry.CreateRegisteredModelRequest{}
	var nilRequest *modelregistry.CreateRegisteredModelRequest
	for want, result := range map[string]any{
		"DataToStruct: result is <nil>, not a pointer to a struct":                                        nil,
		"DataToStruct: result is modelregistry.CreateRegisteredModelRequest, not a pointer to a struct":   *request,
		"DataToStruct: result is *string, not a pointer to a struct":                                      dataToStructPtr("x"),
		"DataToStruct: result is **modelregistry.CreateRegisteredModelRequest, not a pointer to a struct": &request,
		"DataToStruct: result is map[string]interface {}, not a pointer to a struct":                      map[string]any{},
		"DataToStruct: result is a nil *modelregistry.CreateRegisteredModelRequest":                       nilRequest,
	} {
		assert.EqualError(t, DataToStruct(d, s, result), want)
	}
}

func TestDataToStruct_UnsupportedFieldType(t *testing.T) {
	str, _, _, _ := dataToStructScalars()
	s := map[string]*schema.Schema{"field": str}
	// The type is rejected whether or not the data has a value for the field.
	for _, config := range []map[string]any{{}, {"field": "secret"}} {
		d := schema.TestResourceDataRaw(t, s, config)
		for _, result := range []any{
			// The interface of a oneof group and a body stream are interfaces.
			&dataToStructOne[any]{}, &dataToStructOne[dataToStructChoice]{}, &dataToStructOne[io.ReadCloser]{},
			&dataToStructOne[json.RawMessage]{}, &dataToStructOne[[]byte]{}, &dataToStructOne[[][]string]{}, &dataToStructOne[[]json.RawMessage]{},
			&dataToStructOne[[]*dataToStructLeaf]{}, &dataToStructOne[[]any]{}, &dataToStructOne[[]*types.Time]{}, &dataToStructOne[**types.Time]{},
			&dataToStructOne[map[string]int]{}, &dataToStructOne[map[string]any]{}, &dataToStructOne[map[string]json.RawMessage]{},
			&dataToStructOne[map[int]string]{}, &dataToStructOne[map[string]types.Duration]{}, &dataToStructOne[**string]{}, &dataToStructOne[*[]string]{},
			&dataToStructOne[*map[string]string]{}, &dataToStructOne[*any]{}, &dataToStructOne[uint]{}, &dataToStructOne[*uint64]{},
			&dataToStructOne[[2]string]{}, &dataToStructOne[complex128]{}, &dataToStructOne[chan int]{}, &dataToStructOne[func()]{},
		} {
			// The names come from reflect, because json.RawMessage has another
			// one in a Go version where it is an alias.
			owner := reflect.TypeOf(result).Elem()
			fieldType := owner.Field(0).Type.String()
			t.Run(fieldType, func(t *testing.T) {
				assert.EqualError(t, DataToStruct(d, s, result),
					"field: "+owner.String()+".Field of type "+fieldType+" is not a type that DataToStruct supports")
			})
		}
	}
	// The whole error once, without reflect.
	assert.EqualError(t, DataToStruct(schema.TestResourceDataRaw(t, s, nil), s, &dataToStructOne[[]byte]{}),
		"field: sdkgoutility.dataToStructOne[[]uint8].Field of type []uint8 is not a type that DataToStruct supports")
}

func TestDataToStruct_UnsupportedFieldTypeInBlock(t *testing.T) {
	str, _, _, _ := dataToStructScalars()
	s := map[string]*schema.Schema{"field": dataToStructBlocks(map[string]*schema.Schema{"field": str})}
	// The fields of a block are walked, and their types checked, when the
	// block is set.
	var unset dataToStructOne[[]dataToStructOne[any]]
	require.NoError(t, DataToStruct(schema.TestResourceDataRaw(t, s, map[string]any{}), s, &unset))
	assert.Nil(t, unset.Field)

	d := schema.TestResourceDataRaw(t, s, map[string]any{"field": []any{map[string]any{}}})
	const problem = "sdkgoutility.dataToStructOne[interface {}].Field of type interface {} is not a type that DataToStruct supports"
	assert.EqualError(t, DataToStruct(d, s, &dataToStructOne[[]dataToStructOne[any]]{}), "field.0.field: "+problem)
	assert.EqualError(t, DataToStruct(d, s, &dataToStructOne[*dataToStructOne[any]]{}), "field.0.field: "+problem)
	assert.EqualError(t, DataToStruct(d, s, &dataToStructOne[dataToStructOne[any]]{}), "field.0.field: "+problem)
}

func TestDataToStruct_ValueDoesNotFit(t *testing.T) {
	str, boolean, integer, float := dataToStructScalars()
	leaves := dataToStructBlocks(dataToStructLeafSchema())
	oneLeaf := []any{map[string]any{"name": "secret"}}
	const notBlocks = " takes a list of blocks, but the value in the resource data is of type "
	const noBlock = " takes a list of blocks, but the schema of the field has no block"
	const notHeld = " can't hold the value in the resource data, which is "
	const notScalar = notHeld + "of type "
	const notTime = "not a timestamp that types.Time decodes, such as 2006-01-02T15:04:05Z"
	const notDuration = "not a duration that types.Duration decodes, such as 3.5s"
	for _, tc := range []struct {
		name   string
		schema *schema.Schema
		value  any
		result any
		want   string // what follows "of type " in the error
	}{
		{"a string in an integer", str, "secret", &dataToStructOne[*int64]{}, "*int64" + notScalar + "string"},
		{"an integer in a string", integer, 1, &dataToStructOne[*string]{}, "*string" + notScalar + "int"},
		{"an integer in a float", integer, 1, &dataToStructOne[*float64]{}, "*float64" + notScalar + "int"},
		{"a float in an integer", float, 1.5, &dataToStructOne[int64]{}, "int64" + notScalar + "float64"},
		{"a bool in an enum", boolean, true, &dataToStructOne[dataToStructEnum]{}, "sdkgoutility.dataToStructEnum" + notScalar + "bool"},
		{"a string in a bool", str, "secret", &dataToStructOne[*bool]{}, "*bool" + notScalar + "string"},
		{"an integer overflows int8", integer, 300, &dataToStructOne[int8]{}, "int8" + notScalar + "int"},
		{"an integer overflows *int32", integer, 1 << 40, &dataToStructOne[*int32]{}, "*int32" + notScalar + "int"},
		{"a float overflows float32", float, 1e300, &dataToStructOne[float32]{}, "float32" + notScalar + "float64"},
		{"a word in a timestamp", str, "secret", &dataToStructOne[*types.Time]{}, "*types.Time" + notHeld + notTime},
		{"a date in a timestamp", str, "2025-11-25", &dataToStructOne[types.Time]{}, "types.Time" + notHeld + notTime},
		{"a word in a duration", str, "secret", &dataToStructOne[types.Duration]{}, "types.Duration" + notHeld + notDuration},
		{"minutes in a duration", str, "5m", &dataToStructOne[*types.Duration]{}, "*types.Duration" + notHeld + notDuration},
		{"an integer in a timestamp", integer, 1764028800, &dataToStructOne[*types.Time]{}, "*types.Time" + notScalar + "int"},
		{"a string in a message", str, "secret", &dataToStructOne[*dataToStructLeaf]{}, "*sdkgoutility.dataToStructLeaf" + notBlocks + "string"},
		{"a map in a message", dataToStructMap(str), map[string]any{"name": "secret"}, &dataToStructOne[dataToStructLeaf]{},
			"sdkgoutility.dataToStructLeaf" + notBlocks + "map[string]interface {}"},
		{"a list of strings in a message", dataToStructList(str), []any{"secret"}, &dataToStructOne[dataToStructLeaf]{}, "sdkgoutility.dataToStructLeaf" + noBlock},
		{"two blocks in a message", leaves, []any{map[string]any{"name": "secret"}, map[string]any{}}, &dataToStructOne[*dataToStructLeaf]{},
			"*sdkgoutility.dataToStructLeaf takes one block, but the resource data has 2"},
		{"a list of strings in messages", dataToStructList(str), []any{"secret"}, &dataToStructOne[[]dataToStructLeaf]{}, "[]sdkgoutility.dataToStructLeaf" + noBlock},
		{"a string in messages", str, "secret", &dataToStructOne[[]dataToStructLeaf]{}, "[]sdkgoutility.dataToStructLeaf" + notBlocks + "string"},
		{"a set of blocks in messages", dataToStructSet(&schema.Resource{Schema: dataToStructLeafSchema()}), oneLeaf, &dataToStructOne[[]dataToStructLeaf]{},
			"[]sdkgoutility.dataToStructLeaf" + notBlocks + "*schema.Set"},
		{"a string in a slice", str, "secret", &dataToStructOne[[]string]{},
			"[]string takes a list, but the value in the resource data is of type string"},
		{"a set in a slice", dataToStructSet(str), []any{"secret"}, &dataToStructOne[[]string]{},
			"[]string takes a list, but the value in the resource data is of type *schema.Set"},
		{"a map in a slice", dataToStructMap(str), map[string]any{"a": "secret"}, &dataToStructOne[[]string]{},
			"[]string takes a list, but the value in the resource data is of type map[string]interface {}"},
		{"blocks in a slice of strings", leaves, oneLeaf, &dataToStructOne[[]string]{},
			"[]string can't hold item 0 of the list, which is of type map[string]interface {}"},
		{"strings in a slice of integers", dataToStructList(str), []any{"secret"}, &dataToStructOne[[]int64]{},
			"[]int64 can't hold item 0 of the list, which is of type string"},
		{"an empty string in a slice of strings", dataToStructList(str), []any{"secret", ""}, &dataToStructOne[[]string]{},
			"[]string can't hold item 1 of the list, which is of type <nil>"},
		{"the second timestamp isn't one", dataToStructList(str), []any{"2025-11-25T00:00:00Z", "secret"}, &dataToStructOne[[]types.Time]{},
			"[]types.Time can't hold item 1 of the list, which is " + notTime},
		{"the second integer overflows", dataToStructList(integer), []any{1, 300}, &dataToStructOne[[]int8]{},
			"[]int8 can't hold item 1 of the list, which is of type int"},
		{"a string in a map", str, "secret", &dataToStructOne[map[string]string]{},
			"map[string]string takes a map, but the value in the resource data is of type string"},
		{"a list in a map", dataToStructList(str), []any{"secret"}, &dataToStructOne[map[string]string]{},
			"map[string]string takes a map, but the value in the resource data is of type []interface {}"},
		{"a map of integers in a map of strings", dataToStructMap(integer), map[string]any{"a": 1}, &dataToStructOne[map[string]string]{},
			"map[string]string can't hold a value of the map, which is of type int"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			s := map[string]*schema.Schema{"field": tc.schema}
			d := schema.TestResourceDataRaw(t, s, map[string]any{"field": tc.value})
			result := reflect.ValueOf(tc.result).Elem()
			err := DataToStruct(d, s, tc.result)
			require.EqualError(t, err, "field: "+result.Type().String()+".Field of type "+tc.want)
			assert.NotContains(t, err.Error(), "secret", "the error quotes a value of the resource data")
			assert.True(t, result.Field(0).IsZero(), "the field is set although the value doesn't fit")
		})
	}
}

func TestDataToStruct_ErrorNamesNestedField(t *testing.T) {
	str, _, _, _ := dataToStructScalars()
	// count is a string here, which doesn't fit the *int64 of dataToStructLeaf.
	leaves := dataToStructBlocks(map[string]*schema.Schema{"name": str, "count": str})
	branches := dataToStructBlocks(map[string]*schema.Schema{"leaf": leaves, "inline": leaves, "leaves": leaves})
	s := map[string]*schema.Schema{"message_pointer": branches, "message_value": branches, "messages": branches}
	many := map[string]any{"count": "many"}
	const problem = ": sdkgoutility.dataToStructLeaf.Count of type *int64 can't hold the value in the resource data, which is of type string"
	for key, config := range map[string]map[string]any{
		"message_pointer.0.leaf.0.count": {"message_pointer": []any{map[string]any{"leaf": []any{many}}}},
		"message_value.0.inline.0.count": {"message_value": []any{map[string]any{"inline": []any{many}}}},
		"messages.1.leaves.2.count": {"messages": []any{map[string]any{}, map[string]any{
			"leaves": []any{map[string]any{}, map[string]any{"name": "fits"}, many},
		}}},
	} {
		t.Run(key, func(t *testing.T) {
			var got dataToStructShapes
			assert.EqualError(t, DataToStruct(schema.TestResourceDataRaw(t, s, config), s, &got), key+problem)
			assert.Equal(t, dataToStructShapes{}, got, "a field is set although a value below it doesn't fit")
		})
	}
}

// dataToStructReadsJSON checks that DataToStruct reads the string that the JSON encoding of want writes into a
// field of its type by value, by pointer and in a slice, and that an unset attribute leaves those fields as they came.
func dataToStructReadsJSON[T any](t *testing.T, want T) {
	encoded, err := json.Marshal([]T{want, want})
	require.NoError(t, err)
	var texts []any
	require.NoError(t, json.Unmarshal(encoded, &texts))
	str, _, _, _ := dataToStructScalars()
	s := map[string]*schema.Schema{"value": str, "pointer": str, "slice": dataToStructList(str)}
	var got, filled struct {
		Value   T
		Pointer *T
		Slice   []T
	}
	filled.Value, filled.Pointer, filled.Slice = want, &want, []T{want, want}
	d := schema.TestResourceDataRaw(t, s, map[string]any{"value": texts[0], "pointer": texts[0], "slice": texts})
	require.NoError(t, DataToStruct(d, s, &got), texts[0])
	assert.Equal(t, filled, got, texts[0])
	require.NoError(t, DataToStruct(schema.TestResourceDataRaw(t, s, map[string]any{}), s, &got))
	assert.Equal(t, filled, got, "an unset attribute changed a field")
}

func TestDataToStruct_ReadsTimeAndDuration(t *testing.T) {
	for _, want := range []types.Time{{}, {Seconds: 1764028800, Nanos: 120000000}, {Seconds: -1, Nanos: 999999999}, {Seconds: types.MaxTimestampSeconds}} {
		dataToStructReadsJSON(t, want)
	}
	for _, want := range []types.Duration{{}, {Seconds: 3, Nanos: 500000000}, {Seconds: -90}, {Nanos: -1}, {Seconds: 315576000000}} {
		dataToStructReadsJSON(t, want)
	}
	// The decoding of types.Time also takes a numeric offset, which its encoding never writes.
	str, _, _, _ := dataToStructScalars()
	s := map[string]*schema.Schema{"field": str}
	var got dataToStructOne[types.Time]
	require.NoError(t, DataToStruct(schema.TestResourceDataRaw(t, s, map[string]any{"field": "2025-11-25T05:30:00.5+05:30"}), s, &got))
	assert.Equal(t, types.Time{Seconds: 1764028800, Nanos: 500000000}, got.Field)
}

func dataToStructJSON(t *testing.T, message any) string {
	data, err := json.Marshal(message)
	require.NoError(t, err)
	return string(data)
}

// TestDataToStruct_SameRequestsAsClassic fills a classic request and the modular request that replaces it from the
// same resource data, and compares the JSON that the two send.
func TestDataToStruct_SameRequestsAsClassic(t *testing.T) {
	createSchema := common.StructToSchema(ml.CreateModelRequest{}, nil)
	newCreate := func() (any, any) { return &ml.CreateModelRequest{}, &modelregistry.CreateRegisteredModelRequest{} }
	newUpdate := func() (any, any) { return &ml.UpdateModelRequest{}, &modelregistry.UpdateRegisteredModelRequest{} }
	twins := []struct {
		name     string
		s        map[string]*schema.Schema
		requests func() (classic, modular any)
		hasTags  bool
	}{
		{"create", createSchema, newCreate, true},
		{"update", common.StructToSchema(ml.UpdateModelRequest{}, nil), newUpdate, false},
		// databricks_mlflow_model passes the schema of the create request to every request.
		{"update with the schema of create", createSchema, newUpdate, false},
	}
	type tag = map[string]any
	for _, tc := range []struct {
		name   string
		config map[string]any
		want   string // the JSON of the request, without its tags
		tags   string // the JSON of the tags of a request that has them
	}{
		{"only name", map[string]any{"name": "model"}, `{"name":"model"}`, ""},
		{"name, description and two tags",
			map[string]any{"name": "model", "description": "about", "tags": []any{tag{"key": "a", "value": "1"}, tag{"key": "b", "value": "2"}}},
			`{"name":"model","description":"about"}`, `[{"key":"a","value":"1"},{"key":"b","value":"2"}]`},
		{"a tag with only a key", map[string]any{"name": "model", "tags": []any{tag{"key": "a"}}}, `{"name":"model"}`, `[{"key":"a"}]`},
		{"a tag with only a value", map[string]any{"name": "model", "tags": []any{tag{"value": "1"}}}, `{"name":"model"}`, `[{"value":"1"}]`},
		{"a tag with an empty value", map[string]any{"name": "model", "tags": []any{tag{"key": "a", "value": ""}}}, `{"name":"model"}`, `[{"key":"a"}]`},
		{"an empty description", map[string]any{"name": "model", "description": ""}, `{"name":"model"}`, ""},
		{"no tags", map[string]any{"name": "model", "description": "about", "tags": []any{}}, `{"name":"model","description":"about"}`, ""},
		{"an empty tags block", map[string]any{"name": "model", "tags": []any{tag{}}}, `{"name":"model"}`, `[{}]`},
		{"an empty tags block before a tag", map[string]any{"name": "model", "tags": []any{tag{}, tag{"key": "a", "value": "1"}}},
			`{"name":"model"}`, `[{},{"key":"a","value":"1"}]`},
	} {
		for _, twin := range twins {
			t.Run(twin.name+"/"+tc.name, func(t *testing.T) {
				d := schema.TestResourceDataRaw(t, twin.s, tc.config)
				classic, modular := twin.requests()
				common.DataToStructPointer(d, twin.s, classic)
				require.NoError(t, DataToStruct(d, twin.s, modular))
				want := tc.want
				if twin.hasTags && tc.tags != "" {
					want = want[:len(want)-1] + `,"tags":` + tc.tags + "}"
				}
				assert.JSONEq(t, dataToStructJSON(t, classic), dataToStructJSON(t, modular))
				assert.JSONEq(t, want, dataToStructJSON(t, modular))
			})
		}
	}
}

// TestDataToStruct_LeavesOutZeroThatClassicSends pins the request that differs from the classic path: name has no
// omitempty in the classic struct, so the classic client sends an empty name, and DataToStruct leaves it out until
// the caller sets the pointer.
func TestDataToStruct_LeavesOutZeroThatClassicSends(t *testing.T) {
	s := common.StructToSchema(ml.CreateModelRequest{}, nil)
	d := schema.TestResourceDataRaw(t, s, map[string]any{"name": ""})
	var classic ml.CreateModelRequest
	var modular modelregistry.CreateRegisteredModelRequest
	common.DataToStructPointer(d, s, &classic)
	require.NoError(t, DataToStruct(d, s, &modular))
	assert.JSONEq(t, `{"name":""}`, dataToStructJSON(t, classic))
	assert.JSONEq(t, `{}`, dataToStructJSON(t, modular))

	modular.Name = dataToStructPtr(d.Get("name").(string))
	assert.JSONEq(t, dataToStructJSON(t, classic), dataToStructJSON(t, modular))
}
