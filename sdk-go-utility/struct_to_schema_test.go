package sdkgoutility

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"slices"
	"testing"
	"time"
	"unsafe"

	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/sdk-go/core/types"
	modelregistry "github.com/databricks/sdk-go/modelregistry/v1"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

// The types below have the shapes of modular SDK messages: public fields
// without a json tag, pointers for optional scalars and named types for enums.

type structToSchemaTestEnum string

type structToSchemaTestCount int32

type structToSchemaTestRatio float32

type structToSchemaTestFlag bool

type structToSchemaTestByte uint8

type structToSchemaTestTag struct {
	Key   *string
	Value *string
}

// structToSchemaTestShapes holds T in every shape that a scalar has.
type structToSchemaTestShapes[T any] struct {
	Value   T
	Pointer *T
	List    []T
}

// structToSchemaTestField is a message with one field, of type T.
type structToSchemaTestField[T any] struct {
	Field T
}

type structToSchemaTestInner struct {
	Count *int64
	Tag   *structToSchemaTestTag
	Tags  []structToSchemaTestTag
}

// structToSchemaTestOuter holds the same messages at several places and
// depths, none of which is a message that contains itself.
type structToSchemaTestOuter struct {
	Name   *string
	Tag    structToSchemaTestTag
	Inner  *structToSchemaTestInner
	Inners []structToSchemaTestInner
}

// structToSchemaTestRenamed has an API name that the Go name doesn't give, as
// PipelineTask of the jobs package does.
type structToSchemaTestRenamed struct {
	PipelineTaskParameters *string
}

func (m structToSchemaTestRenamed) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Parameters *string `json:"parameters,omitempty"`
	}{m.PipelineTaskParameters})
}

// structToSchemaTestScim has an API name that is not in snake_case, as the
// SCIM messages of the SDK do.
type structToSchemaTestScim struct {
	DisplayName *string
}

func (m structToSchemaTestScim) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		DisplayName *string `json:"displayName,omitempty"`
	}{m.DisplayName})
}

type structToSchemaTestOneof interface{ isStructToSchemaTestOneof() }

type structToSchemaTestChoice struct {
	Name   *string
	Choice structToSchemaTestOneof
}

type structToSchemaTestNode struct {
	Name *string
	Next *structToSchemaTestNode
}

type structToSchemaTestTree struct {
	Children []structToSchemaTestTree
}

type structToSchemaTestCycle struct {
	There structToSchemaTestCycleBack
}

type structToSchemaTestCycleBack struct {
	Back []structToSchemaTestCycle
}

// structToSchemaTestMessages are all messages of modelregistry v1.
var structToSchemaTestMessages = []any{
	modelregistry.Activity{}, modelregistry.ApproveTransitionRequest{}, modelregistry.ApproveTransitionResponse{},
	modelregistry.CommentObject{}, modelregistry.CreateCommentRequest{}, modelregistry.CreateCommentResponse{},
	modelregistry.CreateModelVersionRequest{}, modelregistry.CreateModelVersionResponse{},
	modelregistry.CreateRegisteredModelRequest{}, modelregistry.CreateRegisteredModelResponse{},
	modelregistry.CreateRegistryWebhookRequest{}, modelregistry.CreateRegistryWebhookResponse{},
	modelregistry.CreateTransitionRequest{}, modelregistry.CreateTransitionResponse{},
	modelregistry.DeleteCommentRequest{}, modelregistry.DeleteCommentResponse{},
	modelregistry.DeleteModelVersionRequest{}, modelregistry.DeleteModelVersionResponse{},
	modelregistry.DeleteModelVersionTagRequest{}, modelregistry.DeleteModelVersionTagResponse{},
	modelregistry.DeleteRegisteredModelRequest{}, modelregistry.DeleteRegisteredModelResponse{},
	modelregistry.DeleteRegisteredModelTagRequest{}, modelregistry.DeleteRegisteredModelTagResponse{},
	modelregistry.DeleteRegistryWebhookRequest{}, modelregistry.DeleteRegistryWebhookResponse{},
	modelregistry.DeleteTransitionRequest{}, modelregistry.DeleteTransitionResponse{},
	modelregistry.FeatureList{}, modelregistry.GetLatestVersionsResponse{},
	modelregistry.GetModelVersionDownloadUriRequest{}, modelregistry.GetModelVersionDownloadUriResponse{},
	modelregistry.GetModelVersionRequest{}, modelregistry.GetModelVersionResponse{},
	modelregistry.GetRegisteredModelDatabricksRequest{}, modelregistry.GetRegisteredModelDatabricksResponse{},
	modelregistry.HttpUrlSpec{}, modelregistry.JobSpec{}, modelregistry.LinkedFeature{},
	modelregistry.ListLatestVersionsRequest{}, modelregistry.ListRegisteredModelsRequest{},
	modelregistry.ListRegisteredModelsResponse{}, modelregistry.ListRegistryWebhooksRequest{},
	modelregistry.ListRegistryWebhooksResponse{}, modelregistry.ListTransitionRequest{},
	modelregistry.ListTransitionResponse{}, modelregistry.ModelVersion{}, modelregistry.ModelVersionDatabricks{},
	modelregistry.ModelVersionTag{}, modelregistry.RegisteredModel{}, modelregistry.RegisteredModelDatabricks{},
	modelregistry.RegisteredModelTag{}, modelregistry.RegistryWebhook{}, modelregistry.RejectTransitionRequest{},
	modelregistry.RejectTransitionResponse{}, modelregistry.RenameRegisteredModelRequest{},
	modelregistry.RenameRegisteredModelResponse{}, modelregistry.SearchModelVersionsRequest{},
	modelregistry.SearchModelVersionsResponse{}, modelregistry.SearchRegisteredModelsRequest{},
	modelregistry.SearchRegisteredModelsResponse{}, modelregistry.SetModelVersionTagRequest{},
	modelregistry.SetModelVersionTagResponse{}, modelregistry.SetRegisteredModelTagRequest{},
	modelregistry.SetRegisteredModelTagResponse{}, modelregistry.TestRegistryWebhookRequest{},
	modelregistry.TestRegistryWebhookResponse{}, modelregistry.TransitionModelVersionStageDatabricksRequest{},
	modelregistry.TransitionModelVersionStageDatabricksResponse{}, modelregistry.TransitionRequest{},
	modelregistry.UpdateCommentRequest{}, modelregistry.UpdateCommentResponse{},
	modelregistry.UpdateModelVersionRequest{}, modelregistry.UpdateModelVersionResponse{},
	modelregistry.UpdateRegisteredModelRequest{}, modelregistry.UpdateRegisteredModelResponse{},
	modelregistry.UpdateRegistryWebhookRequest{}, modelregistry.UpdateRegistryWebhookResponse{},
}

// structToSchemaTestScalar is the attribute expected for a scalar field.
func structToSchemaTestScalar(valueType schema.ValueType) *schema.Schema {
	return &schema.Schema{Type: valueType, Optional: true}
}

// structToSchemaTestList is the attribute expected for a slice field: elem is
// a *schema.Schema for scalars and a *schema.Resource for messages.
func structToSchemaTestList(elem any) *schema.Schema {
	return &schema.Schema{Type: schema.TypeList, Optional: true, Elem: elem}
}

// structToSchemaTestBlock is the attribute expected for a message field.
func structToSchemaTestBlock(nested map[string]*schema.Schema) *schema.Schema {
	return &schema.Schema{Type: schema.TypeList, Optional: true, MaxItems: 1, Elem: &schema.Resource{Schema: nested}}
}

// structToSchemaTestStrings is the schema expected for a message whose fields
// are all strings.
func structToSchemaTestStrings(names ...string) map[string]*schema.Schema {
	scm := map[string]*schema.Schema{}
	for _, name := range names {
		scm[name] = structToSchemaTestScalar(schema.TypeString)
	}
	return scm
}

// structToSchemaTestOuterSchema is the schema expected for
// structToSchemaTestOuter.
func structToSchemaTestOuterSchema() map[string]*schema.Schema {
	inner := func() map[string]*schema.Schema {
		return map[string]*schema.Schema{
			"count": structToSchemaTestScalar(schema.TypeInt),
			"tag":   structToSchemaTestBlock(structToSchemaTestStrings("key", "value")),
			"tags":  structToSchemaTestList(&schema.Resource{Schema: structToSchemaTestStrings("key", "value")}),
		}
	}
	return map[string]*schema.Schema{
		"name":   structToSchemaTestScalar(schema.TypeString),
		"tag":    structToSchemaTestBlock(structToSchemaTestStrings("key", "value")),
		"inner":  structToSchemaTestBlock(inner()),
		"inners": structToSchemaTestList(&schema.Resource{Schema: inner()}),
	}
}

func TestStructToSchema_Scalars(t *testing.T) {
	tests := []struct {
		v    any
		want schema.ValueType
	}{
		{structToSchemaTestShapes[string]{}, schema.TypeString},
		{structToSchemaTestShapes[structToSchemaTestEnum]{}, schema.TypeString},
		{structToSchemaTestShapes[types.Time]{}, schema.TypeString},
		{structToSchemaTestShapes[types.Duration]{}, schema.TypeString},
		{structToSchemaTestShapes[int]{}, schema.TypeInt},
		{structToSchemaTestShapes[int8]{}, schema.TypeInt},
		{structToSchemaTestShapes[int16]{}, schema.TypeInt},
		{structToSchemaTestShapes[int32]{}, schema.TypeInt},
		{structToSchemaTestShapes[int64]{}, schema.TypeInt},
		{structToSchemaTestShapes[uint]{}, schema.TypeInt},
		{structToSchemaTestShapes[uint16]{}, schema.TypeInt},
		{structToSchemaTestShapes[uint32]{}, schema.TypeInt},
		{structToSchemaTestShapes[uint64]{}, schema.TypeInt},
		{structToSchemaTestShapes[uintptr]{}, schema.TypeInt},
		{structToSchemaTestShapes[structToSchemaTestCount]{}, schema.TypeInt},
		{structToSchemaTestShapes[float32]{}, schema.TypeFloat},
		{structToSchemaTestShapes[float64]{}, schema.TypeFloat},
		{structToSchemaTestShapes[structToSchemaTestRatio]{}, schema.TypeFloat},
		{structToSchemaTestShapes[bool]{}, schema.TypeBool},
		{structToSchemaTestShapes[structToSchemaTestFlag]{}, schema.TypeBool},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%T", tt.v), func(t *testing.T) {
			got := StructToSchema(tt.v, nil)
			assert.Equal(t, map[string]*schema.Schema{
				"value":   structToSchemaTestScalar(tt.want),
				"pointer": structToSchemaTestScalar(tt.want),
				"list":    structToSchemaTestList(&schema.Schema{Type: tt.want}),
			}, got)
			assert.NoError(t, schema.InternalMap(got).InternalValidate(nil))
		})
	}
}

func TestStructToSchema_Fields(t *testing.T) {
	tag := func() map[string]*schema.Schema { return structToSchemaTestStrings("key", "value") }
	field := func(attribute *schema.Schema) map[string]*schema.Schema {
		return map[string]*schema.Schema{"field": attribute}
	}
	stringMap := &schema.Schema{Type: schema.TypeMap, Optional: true, Elem: schema.TypeString}
	tests := []struct {
		name string
		v    any
		want map[string]*schema.Schema
	}{
		// A slice of uint8 is bytes, so uint8 is not in TestStructToSchema_Scalars.
		{"uint8", structToSchemaTestField[uint8]{}, field(structToSchemaTestScalar(schema.TypeInt))},
		{"pointer to uint8", structToSchemaTestField[*uint8]{}, field(structToSchemaTestScalar(schema.TypeInt))},
		{"named uint8", structToSchemaTestField[structToSchemaTestByte]{}, field(structToSchemaTestScalar(schema.TypeInt))},
		{"message", structToSchemaTestField[structToSchemaTestTag]{}, field(structToSchemaTestBlock(tag()))},
		{"pointer to message", structToSchemaTestField[*structToSchemaTestTag]{}, field(structToSchemaTestBlock(tag()))},
		{"slice of messages", structToSchemaTestField[[]structToSchemaTestTag]{},
			field(structToSchemaTestList(&schema.Resource{Schema: tag()}))},
		{"message without fields", structToSchemaTestField[*struct{}]{},
			field(structToSchemaTestBlock(map[string]*schema.Schema{}))},
		// Only Time of the SDK is a string: time.Time is a message without public fields.
		{"struct named Time of another package", structToSchemaTestField[time.Time]{},
			field(structToSchemaTestBlock(map[string]*schema.Schema{}))},
		{"map of strings", structToSchemaTestField[map[string]string]{}, field(stringMap)},
		{"map of enums", structToSchemaTestField[map[structToSchemaTestEnum]structToSchemaTestEnum]{}, field(stringMap)},
		{"field masks are left out", struct {
			Name       *string
			UpdateMask *types.FieldMask[structToSchemaTestTag]
			Mask       types.FieldMask[structToSchemaTestTag]
		}{}, structToSchemaTestStrings("name")},
		{"unexported and blank fields are left out", struct {
			Name   *string
			hidden *string
			_      [0]int
		}{}, structToSchemaTestStrings("name")},
		{"no fields", struct{}{}, map[string]*schema.Schema{}},
		{"pointer to a struct", &structToSchemaTestTag{}, tag()},
		{"nil pointer to a struct", (*structToSchemaTestTag)(nil), tag()},
		{"name from the fieldmask tag", struct {
			Name *string `fieldmask:"display_name"`
		}{}, structToSchemaTestStrings("display_name")},
		{"name from MarshalJSON", structToSchemaTestRenamed{}, structToSchemaTestStrings("parameters")},
		{"name from MarshalJSON of a nested message", structToSchemaTestField[*structToSchemaTestRenamed]{},
			field(structToSchemaTestBlock(structToSchemaTestStrings("parameters")))},
		{"name from MarshalJSON of a message in a slice", structToSchemaTestField[[]structToSchemaTestRenamed]{},
			field(structToSchemaTestList(&schema.Resource{Schema: structToSchemaTestStrings("parameters")}))},
		{"same message at several places and depths", structToSchemaTestOuter{}, structToSchemaTestOuterSchema()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StructToSchema(tt.v, nil)
			assert.Equal(t, tt.want, got)
			assert.NoError(t, schema.InternalMap(got).InternalValidate(nil))
		})
	}
}

// StructToSchema takes the name from FieldName as it is, also when SDKv2
// doesn't accept it as the name of an attribute.
func TestStructToSchema_KeepsNameThatIsNotInSnakeCase(t *testing.T) {
	got := StructToSchema(structToSchemaTestScim{}, nil)
	assert.Equal(t, structToSchemaTestStrings("displayName"), got)
	assert.ErrorContains(t, schema.InternalMap(got).InternalValidate(nil),
		"displayName: Field name may only contain lowercase alphanumeric characters & underscores")
}

func TestStructToSchema_ModelRegistrySchemas(t *testing.T) {
	tests := []struct {
		v    any
		want map[string]*schema.Schema
	}{
		{modelregistry.CreateRegisteredModelRequest{}, map[string]*schema.Schema{
			"name":        structToSchemaTestScalar(schema.TypeString),
			"tags":        structToSchemaTestList(&schema.Resource{Schema: structToSchemaTestStrings("key", "value")}),
			"description": structToSchemaTestScalar(schema.TypeString),
		}},
		{modelregistry.CreateRegistryWebhookRequest{}, map[string]*schema.Schema{
			"model_name":  structToSchemaTestScalar(schema.TypeString),
			"events":      structToSchemaTestList(&schema.Schema{Type: schema.TypeString}),
			"description": structToSchemaTestScalar(schema.TypeString),
			"status":      structToSchemaTestScalar(schema.TypeString),
			"http_url_spec": structToSchemaTestBlock(map[string]*schema.Schema{
				"url":                     structToSchemaTestScalar(schema.TypeString),
				"enable_ssl_verification": structToSchemaTestScalar(schema.TypeBool),
				"secret":                  structToSchemaTestScalar(schema.TypeString),
				"authorization":           structToSchemaTestScalar(schema.TypeString),
			}),
			"job_spec": structToSchemaTestBlock(structToSchemaTestStrings("job_id", "workspace_url", "access_token")),
		}},
		{modelregistry.SearchModelVersionsRequest{}, map[string]*schema.Schema{
			"filter":      structToSchemaTestScalar(schema.TypeString),
			"max_results": structToSchemaTestScalar(schema.TypeInt),
			"order_by":    structToSchemaTestList(&schema.Schema{Type: schema.TypeString}),
			"page_token":  structToSchemaTestScalar(schema.TypeString),
		}},
		{&modelregistry.TestRegistryWebhookResponse{}, map[string]*schema.Schema{
			"status_code": structToSchemaTestScalar(schema.TypeInt),
			"body":        structToSchemaTestScalar(schema.TypeString),
		}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%T", tt.v), func(t *testing.T) {
			assert.Equal(t, tt.want, StructToSchema(tt.v, nil))
		})
	}
}

// structToSchemaTestCheckMessage fails unless scm has exactly one attribute
// for every field of structType, named by FieldName, with Optional set and
// nothing else beside its type, at every depth.
func structToSchemaTestCheckMessage(t *testing.T, structType reflect.Type, scm map[string]*schema.Schema) {
	assert.Len(t, scm, structType.NumField(), "%s", structType)
	for i := range structType.NumField() {
		field := structType.Field(i)
		attribute, ok := scm[FieldName(structType, field)]
		if !assert.True(t, ok, "%s.%s has no attribute", structType, field.Name) {
			continue
		}
		typeOnly := schema.Schema{Type: attribute.Type, MaxItems: attribute.MaxItems, Elem: attribute.Elem, Optional: true}
		assert.Equal(t, typeOnly, *attribute, "%s.%s", structType, field.Name)
		fieldType := field.Type
		if fieldType.Kind() == reflect.Pointer || fieldType.Kind() == reflect.Slice {
			fieldType = fieldType.Elem()
		}
		if resource, ok := attribute.Elem.(*schema.Resource); ok {
			structToSchemaTestCheckMessage(t, fieldType, resource.Schema)
		}
	}
}

func TestStructToSchema_ModelRegistryMessages(t *testing.T) {
	for _, message := range structToSchemaTestMessages {
		t.Run(fmt.Sprintf("%T", message), func(t *testing.T) {
			got := StructToSchema(message, nil)
			structToSchemaTestCheckMessage(t, reflect.TypeOf(message), got)
			assert.NoError(t, schema.InternalMap(got).InternalValidate(nil))
		})
	}
}

// structToSchemaTestClearOptionality clears Required and Optional of every
// attribute of scm, at every depth, and returns the sorted paths of the
// attributes that were Required.
func structToSchemaTestClearOptionality(prefix string, scm map[string]*schema.Schema) []string {
	var required []string
	for name, attribute := range scm {
		if attribute.Required {
			required = append(required, prefix+name)
		}
		attribute.Required, attribute.Optional = false, false
		if resource, ok := attribute.Elem.(*schema.Resource); ok {
			required = append(required, structToSchemaTestClearOptionality(prefix+name+".", resource.Schema)...)
		}
	}
	slices.Sort(required)
	return required
}

func TestStructToSchema_SameAsClassicTwin(t *testing.T) {
	tests := []struct {
		classic any
		modular any
		// required are the attributes that the json tags of the classic struct
		// make Required, which is the only difference between the two schemas.
		required []string
	}{
		{ml.CreateModelRequest{}, modelregistry.CreateRegisteredModelRequest{}, []string{"name"}},
		{ml.UpdateModelRequest{}, modelregistry.UpdateRegisteredModelRequest{}, []string{"name"}},
		{ml.ModelDatabricks{}, modelregistry.RegisteredModelDatabricks{}, nil},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%T", tt.modular), func(t *testing.T) {
			classic := common.StructToSchema(tt.classic, nil)
			modular := StructToSchema(tt.modular, nil)
			assert.Equal(t, tt.required, structToSchemaTestClearOptionality("", classic))
			assert.Empty(t, structToSchemaTestClearOptionality("", modular))
			assert.Equal(t, classic, modular)
		})
	}
}

func TestStructToSchema_Customize(t *testing.T) {
	got := StructToSchema(modelregistry.CreateRegisteredModelRequest{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		common.SetRequired(s["name"])
		common.SetReadOnly(s["description"])
		common.CustomizeSchemaPath(s, "tags", "key").SetRequired()
		s["registered_model_id"] = &schema.Schema{Type: schema.TypeString, Computed: true}
		return s
	})
	assert.Equal(t, map[string]*schema.Schema{
		"name":        {Type: schema.TypeString, Required: true},
		"description": {Type: schema.TypeString, Computed: true},
		"tags": structToSchemaTestList(&schema.Resource{Schema: map[string]*schema.Schema{
			"key":   {Type: schema.TypeString, Required: true},
			"value": structToSchemaTestScalar(schema.TypeString),
		}}),
		"registered_model_id": {Type: schema.TypeString, Computed: true},
	}, got)
	assert.NoError(t, schema.InternalMap(got).InternalValidate(nil))
}

func TestStructToSchema_ReturnsWhatCustomizeReturns(t *testing.T) {
	replacement := map[string]*schema.Schema{"other": structToSchemaTestScalar(schema.TypeBool)}
	var built map[string]*schema.Schema
	got := StructToSchema(structToSchemaTestTag{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		built = s
		return replacement
	})
	assert.Equal(t, structToSchemaTestStrings("key", "value"), built)
	assert.Equal(t, replacement, got)
}

// A customization of one attribute must not reach the attribute of the same
// message at another place, or the schema that a later call returns.
func TestStructToSchema_SharesNoAttributes(t *testing.T) {
	customized := StructToSchema(structToSchemaTestOuter{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		common.CustomizeSchemaPath(s, "inner", "tag", "key").SetRequired()
		return s
	})
	assert.True(t, common.MustSchemaPath(customized, "inner", "tag", "key").Required)
	for _, path := range [][]string{{"tag", "key"}, {"inner", "tags", "key"}, {"inners", "tag", "key"}} {
		assert.Equal(t, structToSchemaTestScalar(schema.TypeString), common.MustSchemaPath(customized, path...), "%v", path)
	}
	assert.Equal(t, structToSchemaTestOuterSchema(), StructToSchema(structToSchemaTestOuter{}, nil))
}

func TestStructToSchema_ReadsResourceData(t *testing.T) {
	scm := StructToSchema(modelregistry.CreateRegistryWebhookRequest{}, nil)
	d := schema.TestResourceDataRaw(t, scm, map[string]any{
		"model_name": "model",
		"events":     []any{"MODEL_VERSION_CREATED", "COMMENT_CREATED"},
		"http_url_spec": []any{map[string]any{
			"url":                     "https://example.com/hook",
			"enable_ssl_verification": true,
		}},
	})
	assert.Equal(t, "model", d.Get("model_name"))
	assert.Equal(t, []any{"MODEL_VERSION_CREATED", "COMMENT_CREATED"}, d.Get("events"))
	assert.Equal(t, "https://example.com/hook", d.Get("http_url_spec.0.url"))
	assert.Equal(t, true, d.Get("http_url_spec.0.enable_ssl_verification"))
}

func TestStructToSchema_UnsupportedFieldTypes(t *testing.T) {
	const (
		anInterface = "an interface, such as a oneof group or a body stream, has no schema type"
		bytes       = "bytes have no schema type"
		aSlice      = "a slice is supported only of message structs and scalars"
		aMap        = "a map is supported only of strings to strings"
		aPointer    = "a pointer is supported only to a message struct or a scalar"
		otherKind   = "this kind of type has no schema type"
	)
	tests := []struct {
		v         any
		fieldType string
		reason    string
	}{
		{structToSchemaTestField[structToSchemaTestOneof]{}, "sdkgoutility.structToSchemaTestOneof", anInterface},
		{structToSchemaTestField[any]{}, "interface {}", anInterface},
		{structToSchemaTestField[io.ReadCloser]{}, "io.ReadCloser", anInterface},
		// json.RawMessage is an alias of jsontext.Value in newer versions of Go.
		{structToSchemaTestField[json.RawMessage]{}, reflect.TypeFor[json.RawMessage]().String(), bytes},
		{structToSchemaTestField[[]byte]{}, "[]uint8", bytes},
		{structToSchemaTestField[[]structToSchemaTestByte]{}, "[]sdkgoutility.structToSchemaTestByte", bytes},
		{structToSchemaTestField[[][]string]{}, "[][]string", aSlice},
		{structToSchemaTestField[[][]byte]{}, "[][]uint8", aSlice},
		{structToSchemaTestField[[][]structToSchemaTestTag]{}, "[][]sdkgoutility.structToSchemaTestTag", aSlice},
		{structToSchemaTestField[[]*string]{}, "[]*string", aSlice},
		{structToSchemaTestField[[]*structToSchemaTestTag]{}, "[]*sdkgoutility.structToSchemaTestTag", aSlice},
		{structToSchemaTestField[[]map[string]string]{}, "[]map[string]string", aSlice},
		{structToSchemaTestField[[]any]{}, "[]interface {}", aSlice},
		{structToSchemaTestField[map[string]int64]{}, "map[string]int64", aMap},
		{structToSchemaTestField[map[string]bool]{}, "map[string]bool", aMap},
		{structToSchemaTestField[map[string]structToSchemaTestTag]{}, "map[string]sdkgoutility.structToSchemaTestTag", aMap},
		{structToSchemaTestField[map[string]*string]{}, "map[string]*string", aMap},
		{structToSchemaTestField[map[string][]string]{}, "map[string][]string", aMap},
		{structToSchemaTestField[map[string]any]{}, "map[string]interface {}", aMap},
		{structToSchemaTestField[map[int]string]{}, "map[int]string", aMap},
		{structToSchemaTestField[*[]string]{}, "*[]string", aPointer},
		{structToSchemaTestField[*map[string]string]{}, "*map[string]string", aPointer},
		{structToSchemaTestField[**string]{}, "**string", aPointer},
		{structToSchemaTestField[**structToSchemaTestTag]{}, "**sdkgoutility.structToSchemaTestTag", aPointer},
		{structToSchemaTestField[*any]{}, "*interface {}", aPointer},
		{structToSchemaTestField[[2]string]{}, "[2]string", otherKind},
		{structToSchemaTestField[chan string]{}, "chan string", otherKind},
		{structToSchemaTestField[func()]{}, "func()", otherKind},
		{structToSchemaTestField[complex128]{}, "complex128", otherKind},
		{structToSchemaTestField[unsafe.Pointer]{}, "unsafe.Pointer", otherKind},
	}
	for _, tt := range tests {
		t.Run(tt.fieldType, func(t *testing.T) {
			want := fmt.Sprintf("sdkgoutility.StructToSchema: field Field of %T has unsupported type %s: %s",
				tt.v, tt.fieldType, tt.reason)
			assert.PanicsWithError(t, want, func() { StructToSchema(tt.v, nil) })
		})
	}
}

func TestStructToSchema_Panics(t *testing.T) {
	const (
		notAStruct = "sdkgoutility.StructToSchema: a struct or a pointer to a struct is expected, but got "
		oneof      = "sdkgoutility.StructToSchema: field Choice of sdkgoutility.structToSchemaTestChoice has unsupported type " +
			"sdkgoutility.structToSchemaTestOneof: an interface, such as a oneof group or a body stream, has no schema type"
	)
	tests := []struct {
		name string
		v    any
		want string
	}{
		{"nil", nil, notAStruct + "<nil>"},
		{"string", "name", notAStruct + "string"},
		{"pointer to a string", new(string), notAStruct + "*string"},
		{"slice of structs", []structToSchemaTestTag{}, notAStruct + "[]sdkgoutility.structToSchemaTestTag"},
		{"pointer to a pointer to a struct", new(*structToSchemaTestTag), notAStruct + "**sdkgoutility.structToSchemaTestTag"},
		{"oneof group", structToSchemaTestChoice{}, oneof},
		{"oneof group of a nested message", structToSchemaTestField[*structToSchemaTestChoice]{}, oneof},
		{"oneof group of a message in a slice", &structToSchemaTestField[[]structToSchemaTestChoice]{}, oneof},
		{"embedded struct", struct{ structToSchemaTestTag }{}, "sdkgoutility.StructToSchema: field structToSchemaTestTag of " +
			"struct { sdkgoutility.structToSchemaTestTag } has unsupported type sdkgoutility.structToSchemaTestTag: " +
			"embedded fields are not supported"},
		{"embedded pointer to an exported struct", struct{ *types.Time }{}, "sdkgoutility.StructToSchema: field Time of " +
			"struct { *types.Time } has unsupported type *types.Time: embedded fields are not supported"},
		{"message that contains itself through a pointer", structToSchemaTestNode{}, "sdkgoutility.StructToSchema: field Next of " +
			"sdkgoutility.structToSchemaTestNode has unsupported type *sdkgoutility.structToSchemaTestNode: " +
			"the message sdkgoutility.structToSchemaTestNode contains itself"},
		{"message that contains itself in a slice", structToSchemaTestTree{}, "sdkgoutility.StructToSchema: field Children of " +
			"sdkgoutility.structToSchemaTestTree has unsupported type []sdkgoutility.structToSchemaTestTree: " +
			"the message sdkgoutility.structToSchemaTestTree contains itself"},
		{"message that contains itself through another message", structToSchemaTestCycle{}, "sdkgoutility.StructToSchema: field Back of " +
			"sdkgoutility.structToSchemaTestCycleBack has unsupported type []sdkgoutility.structToSchemaTestCycle: " +
			"the message sdkgoutility.structToSchemaTestCycle contains itself"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.PanicsWithError(t, tt.want, func() {
				StructToSchema(tt.v, func(map[string]*schema.Schema) map[string]*schema.Schema {
					t.Error("customize ran although the schema can't be built")
					return nil
				})
			})
		})
	}
}
