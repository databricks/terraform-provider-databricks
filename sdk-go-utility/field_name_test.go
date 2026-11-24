package sdkgoutility

import (
	"encoding/json"
	"errors"
	"io"
	"reflect"
	"sync"
	"testing"

	modelregistry "github.com/databricks/sdk-go/modelregistry/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// modelRegistryMessages are all messages of modelregistry v1 that have fields.
var modelRegistryMessages = []any{
	modelregistry.Activity{},
	modelregistry.ApproveTransitionRequest{},
	modelregistry.ApproveTransitionResponse{},
	modelregistry.CommentObject{},
	modelregistry.CreateCommentRequest{},
	modelregistry.CreateCommentResponse{},
	modelregistry.CreateModelVersionRequest{},
	modelregistry.CreateModelVersionResponse{},
	modelregistry.CreateRegisteredModelRequest{},
	modelregistry.CreateRegisteredModelResponse{},
	modelregistry.CreateRegistryWebhookRequest{},
	modelregistry.CreateRegistryWebhookResponse{},
	modelregistry.CreateTransitionRequest{},
	modelregistry.CreateTransitionResponse{},
	modelregistry.DeleteCommentRequest{},
	modelregistry.DeleteModelVersionRequest{},
	modelregistry.DeleteModelVersionTagRequest{},
	modelregistry.DeleteRegisteredModelRequest{},
	modelregistry.DeleteRegisteredModelTagRequest{},
	modelregistry.DeleteRegistryWebhookRequest{},
	modelregistry.DeleteTransitionRequest{},
	modelregistry.DeleteTransitionResponse{},
	modelregistry.FeatureList{},
	modelregistry.GetLatestVersionsResponse{},
	modelregistry.GetModelVersionDownloadUriRequest{},
	modelregistry.GetModelVersionDownloadUriResponse{},
	modelregistry.GetModelVersionRequest{},
	modelregistry.GetModelVersionResponse{},
	modelregistry.GetRegisteredModelDatabricksRequest{},
	modelregistry.GetRegisteredModelDatabricksResponse{},
	modelregistry.HttpUrlSpec{},
	modelregistry.JobSpec{},
	modelregistry.LinkedFeature{},
	modelregistry.ListLatestVersionsRequest{},
	modelregistry.ListRegisteredModelsRequest{},
	modelregistry.ListRegisteredModelsResponse{},
	modelregistry.ListRegistryWebhooksRequest{},
	modelregistry.ListRegistryWebhooksResponse{},
	modelregistry.ListTransitionRequest{},
	modelregistry.ListTransitionResponse{},
	modelregistry.ModelVersion{},
	modelregistry.ModelVersionDatabricks{},
	modelregistry.ModelVersionTag{},
	modelregistry.RegisteredModel{},
	modelregistry.RegisteredModelDatabricks{},
	modelregistry.RegisteredModelTag{},
	modelregistry.RegistryWebhook{},
	modelregistry.RejectTransitionRequest{},
	modelregistry.RejectTransitionResponse{},
	modelregistry.RenameRegisteredModelRequest{},
	modelregistry.RenameRegisteredModelResponse{},
	modelregistry.SearchModelVersionsRequest{},
	modelregistry.SearchModelVersionsResponse{},
	modelregistry.SearchRegisteredModelsRequest{},
	modelregistry.SearchRegisteredModelsResponse{},
	modelregistry.SetModelVersionTagRequest{},
	modelregistry.SetRegisteredModelTagRequest{},
	modelregistry.TestRegistryWebhookRequest{},
	modelregistry.TestRegistryWebhookResponse{},
	modelregistry.TransitionModelVersionStageDatabricksRequest{},
	modelregistry.TransitionModelVersionStageDatabricksResponse{},
	modelregistry.TransitionRequest{},
	modelregistry.UpdateCommentRequest{},
	modelregistry.UpdateCommentResponse{},
	modelregistry.UpdateModelVersionRequest{},
	modelregistry.UpdateModelVersionResponse{},
	modelregistry.UpdateRegisteredModelRequest{},
	modelregistry.UpdateRegisteredModelResponse{},
	modelregistry.UpdateRegistryWebhookRequest{},
	modelregistry.UpdateRegistryWebhookResponse{},
}

// setOnlyField gives a field of a modelregistry message a non-zero value and
// returns the shape of the field. It is written apart from setNonEmpty, so
// that the test doesn't prove the probe with the probe's own code.
func setOnlyField(t *testing.T, field reflect.Value) string {
	switch field.Kind() {
	case reflect.Pointer:
		field.Set(reflect.New(field.Type().Elem()))
		switch target := field.Elem(); target.Kind() {
		case reflect.Struct:
			return "nested message"
		case reflect.String:
			target.SetString("text")
			return "pointer scalar"
		case reflect.Int, reflect.Int64:
			target.SetInt(7)
			return "pointer scalar"
		case reflect.Bool:
			target.SetBool(true)
			return "pointer scalar"
		}
	case reflect.String:
		field.SetString("VALUE")
		return "enum"
	case reflect.Slice:
		field.Set(reflect.MakeSlice(field.Type(), 1, 1))
		switch element := field.Index(0); {
		case element.Kind() == reflect.Struct:
			return "repeated message"
		case element.Kind() == reflect.String && element.Type().PkgPath() != "":
			element.SetString("VALUE")
			return "repeated enum"
		case element.Kind() == reflect.String:
			element.SetString("text")
			return "repeated scalar"
		}
	}
	t.Fatalf("setOnlyField can't set a %s", field.Type())
	return ""
}

func TestFieldName_ModelRegistryMatchesMarshalJSON(t *testing.T) {
	require.GreaterOrEqual(t, len(modelRegistryMessages), 15)
	shapes := map[string]int{}
	for _, message := range modelRegistryMessages {
		structType := reflect.TypeOf(message)
		require.NotZero(t, structType.NumField(), structType.Name())
		for i := range structType.NumField() {
			field := structType.Field(i)
			t.Run(structType.Name()+"."+field.Name, func(t *testing.T) {
				value := reflect.New(structType)
				shapes[setOnlyField(t, value.Elem().Field(i))]++
				data, err := json.Marshal(value.Interface())
				require.NoError(t, err)
				var object map[string]json.RawMessage
				require.NoError(t, json.Unmarshal(data, &object))
				require.Len(t, object, 1, "JSON is %s", data)
				assert.Contains(t, object, FieldName(structType, field))
			})
		}
	}
	for _, shape := range []string{
		"pointer scalar", "enum", "nested message", "repeated message", "repeated scalar", "repeated enum",
	} {
		assert.NotZero(t, shapes[shape], "no field under test is a %s", shape)
	}
}

// fieldNames returns FieldName of every field of structType, which it passes
// to FieldName as passedType: the struct type itself or a pointer to it.
func fieldNames(passedType, structType reflect.Type) []string {
	var names []string
	for i := range structType.NumField() {
		names = append(names, FieldName(passedType, structType.Field(i)))
	}
	return names
}

func TestFieldName_ModelRegistryNames(t *testing.T) {
	tests := []struct {
		message any
		want    []string
	}{
		{modelregistry.CreateRegisteredModelRequest{}, []string{"name", "tags", "description"}},
		{modelregistry.UpdateRegisteredModelRequest{}, []string{"name", "description"}},
		{modelregistry.GetRegisteredModelDatabricksRequest{}, []string{"name"}},
		{modelregistry.DeleteRegisteredModelRequest{}, []string{"name"}},
		{modelregistry.RegisteredModelDatabricks{}, []string{
			"name", "creation_timestamp", "last_updated_timestamp", "user_id", "description",
			"latest_versions", "id", "permission_level", "tags",
		}},
		{modelregistry.RegisteredModel{}, []string{
			"name", "creation_timestamp", "last_updated_timestamp", "user_id", "description",
			"latest_versions", "tags",
		}},
		{modelregistry.RegisteredModelTag{}, []string{"key", "value"}},
		{modelregistry.ModelVersion{}, []string{
			"name", "version", "creation_timestamp", "last_updated_timestamp", "user_id",
			"current_stage", "description", "source", "run_id", "status", "status_message",
			"tags", "run_link",
		}},
	}
	for _, tt := range tests {
		structType := reflect.TypeOf(tt.message)
		t.Run(structType.Name(), func(t *testing.T) {
			assert.Equal(t, tt.want, fieldNames(structType, structType))
			assert.Equal(t, tt.want, fieldNames(reflect.PointerTo(structType), structType), "pointer to the struct")
		})
	}
}

// The types below have the shapes of modular SDK messages: public fields
// without a json tag and a MarshalJSON method that writes the API names.

type testEnum string

type testNested struct {
	Value *string
}

type isTestChoice interface{ isTestChoice() }

// testChoiceText is the wrapper struct of a oneof variant, which has no
// MarshalJSON method in the SDK either.
type testChoiceText struct {
	Text      string `fieldmask:"text"`
	PlainText string
}

func (*testChoiceText) isTestChoice() {}

// Embedded is exported so that the embedded field is.
type Embedded struct {
	Value *string
}

// testMessage has a field of every kind that the probe can set, each with an
// API name that snakeCase gets wrong, and fields that the probe can't set.
type testMessage struct {
	DisplayName            *string
	Ref                    *string
	Resources              []testNested
	PipelineTaskParameters map[string]string
	VisualizationColor01   *string
	ClusterId              *int64 `fieldmask:"cluster_id"`
	State                  testEnum
	Enabled                bool
	Count                  int
	Size                   uint32
	Ratio                  float64
	Payload                json.RawMessage
	Checksum               []byte
	Spec                   *testNested
	Choice                 isTestChoice
	Inline                 testNested
	Contents               io.ReadCloser
	Embedded
	hidden *string `fieldmask:"hidden"`
	_      [0]int  `fieldmask_oneof:"Choice"`
}

func (m testMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		DisplayName            *string           `json:"displayName,omitempty"`
		Ref                    *string           `json:"$ref,omitempty"`
		Resources              []testNested      `json:"Resources,omitempty"`
		PipelineTaskParameters map[string]string `json:"parameters,omitempty"`
		VisualizationColor01   *string           `json:"visualization_color_01,omitempty"`
		ClusterId              *int64            `json:"clusterId,omitempty"`
		State                  testEnum          `json:"stateEnum,omitempty"`
		Enabled                bool              `json:"enabledFlag,omitempty"`
		Count                  int               `json:"countValue,omitempty"`
		Size                   uint32            `json:"sizeValue,omitempty"`
		Ratio                  float64           `json:"ratioValue,omitempty"`
		Payload                json.RawMessage   `json:"rawPayload,omitempty"`
		Checksum               []byte            `json:"checksumBytes,omitempty"`
		Spec                   *testNested       `json:"specMessage,omitempty"`
		Inline                 testNested        `json:"inlined"`
		Hidden                 *string           `json:"hidden,omitempty"`
	}{
		m.DisplayName, m.Ref, m.Resources, m.PipelineTaskParameters, m.VisualizationColor01, m.ClusterId,
		m.State, m.Enabled, m.Count, m.Size, m.Ratio, m.Payload, m.Checksum, m.Spec, m.Inline, m.hidden,
	})
}

func TestFieldName_ProbesMarshalJSON(t *testing.T) {
	want := map[string]string{
		// The JSON key, where snakeCase gives another name.
		"DisplayName":            "displayName",
		"Ref":                    "$ref",
		"Resources":              "Resources",
		"PipelineTaskParameters": "parameters",
		"VisualizationColor01":   "visualization_color_01",
		// The JSON key wins over the fieldmask tag.
		"ClusterId": "clusterId",
		// The other kinds of field that the probe sets.
		"State":    "stateEnum",
		"Enabled":  "enabledFlag",
		"Count":    "countValue",
		"Size":     "sizeValue",
		"Ratio":    "ratioValue",
		"Payload":  "rawPayload",
		"Checksum": "checksumBytes",
		"Spec":     "specMessage",
		// Fields that the probe can't set get their snake_case name: a oneof
		// group, a struct that is not a pointer, whose key the message always
		// writes, a body stream and an embedded struct.
		"Choice":   "choice",
		"Inline":   "inline",
		"Contents": "contents",
		"Embedded": "embedded",
		// Unexported fields have no name, with or without a tag.
		"hidden": "",
		"_":      "",
	}
	structType := reflect.TypeOf(testMessage{})
	require.Equal(t, len(want), structType.NumField())
	for i := range structType.NumField() {
		field := structType.Field(i)
		t.Run(field.Name, func(t *testing.T) {
			assert.Equal(t, want[field.Name], FieldName(structType, field))
			assert.Equal(t, want[field.Name], FieldName(reflect.PointerTo(structType), field), "pointer to the struct")
			pointerToPointer := reflect.PointerTo(reflect.PointerTo(structType))
			assert.Equal(t, want[field.Name], FieldName(pointerToPointer, field), "pointer to a pointer")
		})
	}
}

// testPicky marshals to an object, but not for every field.
type testPicky struct {
	Good       *string
	Broken     *string
	Twice      *string
	Silent     *string
	MaskedName *string `fieldmask:"masked"`
}

func (m testPicky) MarshalJSON() ([]byte, error) {
	if m.Broken != nil {
		return nil, errors.New("broken is set")
	}
	object := map[string]any{}
	if m.Good != nil {
		object["goodKey"] = m.Good
	}
	if m.Twice != nil {
		object["twiceA"], object["twiceB"] = m.Twice, m.Twice
	}
	return json.Marshal(object)
}

type testNeverMarshals struct {
	UserName *string
}

func (testNeverMarshals) MarshalJSON() ([]byte, error) {
	return nil, errors.New("never")
}

type testMarshalsToList struct {
	UserName *string
}

func (testMarshalsToList) MarshalJSON() ([]byte, error) {
	return []byte(`["userName"]`), nil
}

type testMarshalsToNull struct {
	UserName *string
}

func (testMarshalsToNull) MarshalJSON() ([]byte, error) {
	return []byte(`null`), nil
}

type testPointerReceiver struct {
	UserName *string
}

func (m *testPointerReceiver) MarshalJSON() ([]byte, error) {
	if m.UserName == nil {
		return []byte(`{}`), nil
	}
	return []byte(`{"userName":"x"}`), nil
}

// testScalar has a MarshalJSON method and is not a struct.
type testScalar string

func (testScalar) MarshalJSON() ([]byte, error) {
	return []byte(`{"userName":"x"}`), nil
}

// testPlain has no MarshalJSON method.
type testPlain struct {
	UserName *string
	Masked   *string `fieldmask:"maskedName"`
	EmptyTag *string `fieldmask:""`
	JsonOnly *string `json:"jsonName" tf:"alias:tf_name"`
	Both     *string `json:"bothJson" fieldmask:"both_name"`
}

func TestFieldName_FallsBackToTagAndSnakeCase(t *testing.T) {
	field := func(structType reflect.Type, name string) reflect.StructField {
		found, ok := structType.FieldByName(name)
		require.True(t, ok, name)
		return found
	}
	picky := reflect.TypeOf(testPicky{})
	plain := reflect.TypeOf(testPlain{})
	variant := reflect.TypeOf(testChoiceText{})
	userName := field(plain, "UserName")
	tests := []struct {
		name       string
		structType reflect.Type
		field      reflect.StructField
		want       string
	}{
		{"probed next to fields that fail", picky, field(picky, "Good"), "goodKey"},
		{"marshal fails with the field set", picky, field(picky, "Broken"), "broken"},
		{"field adds two keys", picky, field(picky, "Twice"), "twice"},
		{"field adds no key", picky, field(picky, "Silent"), "silent"},
		{"field adds no key and has a tag", picky, field(picky, "MaskedName"), "masked"},
		{"marshal always fails", reflect.TypeOf(testNeverMarshals{}), userName, "user_name"},
		{"marshals to a list", reflect.TypeOf(testMarshalsToList{}), userName, "user_name"},
		{"marshals to null", reflect.TypeOf(testMarshalsToNull{}), userName, "user_name"},
		{"MarshalJSON on the pointer", reflect.TypeOf(testPointerReceiver{}), userName, "userName"},
		{"MarshalJSON on a type that is not a struct", reflect.TypeOf(testScalar("")), userName, "user_name"},
		{"no MarshalJSON", plain, userName, "user_name"},
		{"no MarshalJSON, fieldmask tag", plain, field(plain, "Masked"), "maskedName"},
		{"no MarshalJSON, empty fieldmask tag", plain, field(plain, "EmptyTag"), "empty_tag"},
		{"no MarshalJSON, json and tf tags are not read", plain, field(plain, "JsonOnly"), "json_only"},
		{"no MarshalJSON, fieldmask tag next to a json tag", plain, field(plain, "Both"), "both_name"},
		{"oneof variant with a tag", variant, field(variant, "Text"), "text"},
		{"oneof variant without a tag", variant, field(variant, "PlainText"), "plain_text"},
		{"nil struct type", nil, userName, "user_name"},
		{"nil struct type, fieldmask tag", nil, field(plain, "Masked"), "maskedName"},
		{"struct type that is not a struct", reflect.TypeOf(0), userName, "user_name"},
		{"field of another struct", reflect.TypeOf(testMessage{}), userName, "user_name"},
		{"zero field", plain, reflect.StructField{}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, FieldName(tt.structType, tt.field))
		})
	}
}

var testCountedMarshals int

type testCounted struct {
	First  *string
	Second *string
}

func (m testCounted) MarshalJSON() ([]byte, error) {
	testCountedMarshals++
	object := map[string]any{}
	if m.First != nil {
		object["firstKey"] = m.First
	}
	if m.Second != nil {
		object["secondKey"] = m.Second
	}
	return json.Marshal(object)
}

func TestFieldName_ProbesStructOnce(t *testing.T) {
	structType := reflect.TypeOf(testCounted{})
	assert.Equal(t, "firstKey", FieldName(structType, structType.Field(0)))
	assert.Equal(t, 3, testCountedMarshals, "the empty message and one per field")
	for range 3 {
		assert.Equal(t, "firstKey", FieldName(structType, structType.Field(0)))
		assert.Equal(t, "secondKey", FieldName(reflect.PointerTo(structType), structType.Field(1)))
	}
	assert.Equal(t, 3, testCountedMarshals, "later calls don't marshal")
}

type testConcurrent struct {
	UserName *string
}

func (m testConcurrent) MarshalJSON() ([]byte, error) {
	if m.UserName == nil {
		return []byte(`{}`), nil
	}
	return []byte(`{"userName":"x"}`), nil
}

func TestFieldName_Concurrent(t *testing.T) {
	structType := reflect.TypeOf(testConcurrent{})
	names := make([]string, 32)
	var wg sync.WaitGroup
	for i := range names {
		wg.Go(func() {
			names[i] = FieldName(structType, structType.Field(0))
		})
	}
	wg.Wait()
	for _, name := range names {
		assert.Equal(t, "userName", name)
	}
}

func TestSetNonEmpty(t *testing.T) {
	type kinds struct {
		Pointer   *string
		Struct    *testNested
		String    string
		Enum      testEnum
		Bool      bool
		Int       int
		Int64     int64
		Uint8     uint8
		Float32   float32
		Strings   []string
		Bytes     []byte
		Raw       json.RawMessage
		Map       map[string]testNested
		Interface isTestChoice
		Any       any
		Nested    testNested
		Array     [2]string
		Chan      chan int
		Func      func()
		Complex   complex128
	}
	wantSet := map[string]bool{
		"Pointer": true, "Struct": true, "String": true, "Enum": true, "Bool": true, "Int": true,
		"Int64": true, "Uint8": true, "Float32": true, "Strings": true, "Bytes": true, "Raw": true, "Map": true,
	}
	value := reflect.ValueOf(&kinds{}).Elem()
	for i := range value.NumField() {
		name := value.Type().Field(i).Name
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, wantSet[name], setNonEmpty(value.Field(i)))
			assert.Equal(t, !wantSet[name], value.Field(i).IsZero(), "a field is zero exactly when it wasn't set")
		})
	}
	set := value.Interface().(kinds)
	assert.Len(t, set.Strings, 1)
	assert.Len(t, set.Map, 1)
	assert.True(t, json.Valid(set.Raw), "a json.RawMessage must hold JSON to marshal")
	data, err := json.Marshal(set.Raw)
	assert.NoError(t, err)
	assert.Equal(t, string(set.Raw), string(data))
}

func TestSnakeCase(t *testing.T) {
	tests := []struct{ goName, want string }{
		{"", ""},
		{"A", "a"},
		{"Name", "name"},
		{"Id", "id"},
		{"RegisteredModelId", "registered_model_id"},
		{"LatestVersions", "latest_versions"},
		// A digit stays with the letters before it.
		{"Ipv6Cidr", "ipv6_cidr"},
		{"S3Uri", "s3_uri"},
		{"S3", "s3"},
		{"Oauth2Token", "oauth2_token"},
		{"Ai21labsApiKey", "ai21labs_api_key"},
		{"R2TempCredentials", "r2_temp_credentials"},
		{"X509Certificate", "x509_certificate"},
		{"VisualizationColor01", "visualization_color01"},
		// Every capital in a run of capitals is a word.
		{"HTTPStatus", "h_t_t_p_status"},
		{"TShirtSize", "t_shirt_size"},
		{"UserID", "user_i_d"},
		// An underscore is kept and not doubled.
		{"Foo_Bar", "foo_bar"},
		{"Foo_bar", "foo_bar"},
		{"Foo__Bar", "foo__bar"},
		{"Trailing_", "trailing_"},
		{"already_snake", "already_snake"},
		{"lowerCamel", "lower_camel"},
		// Letters outside ASCII follow the same rule.
		{"ÉtatCivil", "état_civil"},
		{"NomÉtat", "nom_état"},
	}
	for _, tt := range tests {
		t.Run(tt.goName, func(t *testing.T) {
			assert.Equal(t, tt.want, snakeCase(tt.goName))
		})
	}
}
