package sdkgoutility

import (
	"encoding/json"
	"fmt"
	"io"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/sdk-go/core/types"
	modelregistry "github.com/databricks/sdk-go/modelregistry/v1"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func structToDataPtr[T any](v T) *T { return &v }

// structToDataTestData returns resource data with the plan for config, or
// without a plan for a nil config, as on an import. A new resource is one that
// the provider reads after a create or an import.
func structToDataTestData(t *testing.T, s map[string]*schema.Schema, config map[string]any, isNew bool) *schema.ResourceData {
	d := (&schema.Resource{Schema: s}).TestResourceData()
	if config != nil {
		d = schema.TestResourceDataRaw(t, s, config)
	}
	d.SetId("id")
	if isNew {
		d.MarkNewResource()
	}
	return d
}

// structToDataGet returns d.Get(name), with a set as its list, because two
// sets with the same items are not deeply equal.
func structToDataGet(d *schema.ResourceData, name string) any {
	if set, ok := d.Get(name).(*schema.Set); ok {
		return set.List()
	}
	return d.Get(name)
}

func structToDataDecode[T any](t *testing.T, body string) *T {
	var v T
	require.NoError(t, json.Unmarshal([]byte(body), &v))
	return &v
}

type structToDataTestEnum string

// structToDataClassic and structToDataModular are twins like a classic SDK struct
// and a modular one: the same fields, as tagged values and as untagged pointers.
type structToDataClassicItem struct {
	Key   string   `json:"key,omitempty"`
	Count int64    `json:"count,omitempty"`
	Flag  bool     `json:"flag,omitempty"`
	Names []string `json:"names,omitempty"`
}

// A part has no number and no bool, so that an empty part is an empty block
// unless its list is written.
type structToDataClassicPart struct {
	Key   string   `json:"key,omitempty"`
	Names []string `json:"names,omitempty"`
}

type structToDataClassic struct {
	Title   string                    `json:"title"`
	Total   int64                     `json:"total"`
	Enabled bool                      `json:"enabled"`
	Created string                    `json:"created,omitempty"`
	Timeout string                    `json:"timeout,omitempty"`
	Text    string                    `json:"text,omitempty"`
	Mode    string                    `json:"mode,omitempty" tf:"default:auto"`
	Number  int64                     `json:"number,omitempty"`
	Size    int                       `json:"size,omitempty"`
	Ratio   float64                   `json:"ratio,omitempty"`
	Flag    bool                      `json:"flag,omitempty"`
	Level   structToDataTestEnum      `json:"level,omitempty"`
	Names   []string                  `json:"names,omitempty"`
	Unique  []string                  `json:"unique,omitempty" tf:"slice_set"`
	Labels  map[string]string         `json:"labels,omitempty"`
	Nested  *structToDataClassicItem  `json:"nested,omitempty"`
	Items   []structToDataClassicItem `json:"items,omitempty"`
	Group   []structToDataClassicItem `json:"group,omitempty" tf:"slice_set"`
	Parts   []structToDataClassicPart `json:"parts,omitempty"`
}

type structToDataModularItem struct {
	Key   *string
	Count *int64
	Flag  *bool
	Names []string
}

type structToDataModularPart struct {
	Key   *string
	Names []string
}

type structToDataModular struct {
	Title   *string
	Total   *int64
	Enabled *bool
	Created *types.Time
	Timeout *types.Duration
	Text    *string
	Mode    *string
	Number  *int64
	Size    *int
	Ratio   *float64
	Flag    *bool
	Level   structToDataTestEnum
	Names   []string
	Unique  []string
	Labels  map[string]string
	Nested  *structToDataModularItem
	Items   []structToDataModularItem
	Group   []structToDataModularItem
	Parts   []structToDataModularPart
}

// structToDataTwin is a classic SDK struct and its modular twin, with what to compare them on.
type structToDataTwin struct {
	// decode reads one response into the classic struct and the modular one.
	decode    func(t *testing.T, body string) (classic, modular any)
	schemas   map[string]map[string]*schema.Schema
	configs   map[string]map[string]any
	responses map[string]string
}

// structToDataComputed returns the schema of classic with the named fields computed.
func structToDataComputed(classic any, names ...string) map[string]*schema.Schema {
	return common.StructToSchema(classic, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		for _, name := range names {
			s[name].Computed = true
		}
		return s
	})
}

// structToDataReadOnlySchema returns the schema of classic with every field computed and not optional.
func structToDataReadOnlySchema(classic any) map[string]*schema.Schema {
	return common.StructToSchema(classic, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		for _, fieldSchema := range s {
			structToDataReadOnly(fieldSchema)
		}
		return s
	})
}

func structToDataModelBody(fields string) string {
	return `{"registered_model_databricks": {` + fields + `}}`
}

const structToDataModelVersions = `"latest_versions": [{"name": "model", "version": "1", "current_stage": "Production",
	"status": "READY", "creation_timestamp": 1700000000002, "tags": [{"key": "a", "value": "b"}]}, {"name": "model", "version": "2"}, {}]`

var structToDataTwins = map[string]structToDataTwin{
	"ml.ModelDatabricks and modelregistry.RegisteredModelDatabricks": {
		decode: func(t *testing.T, body string) (any, any) {
			return structToDataDecode[ml.GetModelResponse](t, body).RegisteredModelDatabricks,
				structToDataDecode[modelregistry.GetRegisteredModelDatabricksResponse](t, body).RegisteredModelDatabricks
		},
		schemas: map[string]map[string]*schema.Schema{
			"schema of the twin": common.StructToSchema(ml.ModelDatabricks{}, nil),
			"read-only fields":   structToDataReadOnlySchema(ml.ModelDatabricks{}),
			"computed fields": structToDataComputed(ml.ModelDatabricks{}, "id", "user_id", "permission_level",
				"creation_timestamp", "last_updated_timestamp", "latest_versions"),
			"schema of databricks_mlflow_model": common.StructToSchema(ml.CreateModelRequest{},
				func(s map[string]*schema.Schema) map[string]*schema.Schema {
					s["name"].ForceNew = true
					s["registered_model_id"] = &schema.Schema{Type: schema.TypeString, Computed: true}
					return s
				}),
		},
		configs: map[string]map[string]any{
			"no plan":            nil,
			"nothing configured": {},
			"description and tags configured": {"name": "model", "description": "configured",
				"tags": []any{map[string]any{"key": "configured", "value": "tag"}}},
			"description configured": {"name": "model", "description": "configured"},
			"tags configured":        {"name": "model", "tags": []any{map[string]any{"key": "configured", "value": "tag"}}},
		},
		responses: map[string]string{
			"all fields": structToDataModelBody(`"name": "model", "id": "abc", "user_id": "user@example.com",
				"description": "from the API", "permission_level": "CAN_MANAGE", "creation_timestamp": 1700000000000,
				"last_updated_timestamp": 1700000000001, "tags": [{"key": "team", "value": "ml"}, {"key": "no value"}], ` +
				structToDataModelVersions),
			"without description": structToDataModelBody(`"name": "model", "tags": [{"key": "team", "value": "ml"}], ` +
				structToDataModelVersions),
			"without tags":       structToDataModelBody(`"name": "model", "description": "from the API", "user_id": "user@example.com"`),
			"only the name":      structToDataModelBody(`"name": "model"`),
			"explicit zeros":     structToDataModelBody(`"name": "", "description": "", "creation_timestamp": 0, "tags": [], "latest_versions": []`),
			"empty message":      structToDataModelBody(``),
			"only empty objects": structToDataModelBody(`"tags": [{}], "latest_versions": [{"tags": [{}]}]`),
		},
	},
	"ml.RegistryWebhook and modelregistry.RegistryWebhook": {
		decode: func(t *testing.T, body string) (any, any) {
			return structToDataDecode[ml.RegistryWebhook](t, body), structToDataDecode[modelregistry.RegistryWebhook](t, body)
		},
		schemas: map[string]map[string]*schema.Schema{
			"schema of the twin": common.StructToSchema(ml.RegistryWebhook{}, nil),
			"read-only fields":   structToDataReadOnlySchema(ml.RegistryWebhook{}),
			"computed fields": structToDataComputed(ml.RegistryWebhook{}, "id", "status", "creation_timestamp",
				"last_updated_timestamp", "http_url_spec", "events"),
		},
		configs: map[string]map[string]any{
			"no plan":            nil,
			"nothing configured": {},
			"everything configured": {"description": "configured", "status": "ACTIVE", "model_name": "configured",
				"events":        []any{"REGISTERED_MODEL_CREATED"},
				"http_url_spec": []any{map[string]any{"url": "https://configured.example.com", "enable_ssl_verification": true}},
				"job_spec":      []any{map[string]any{"job_id": "1", "workspace_url": "https://configured.example.com"}}},
		},
		responses: map[string]string{
			"all fields": `{"id": "w1", "events": ["MODEL_VERSION_CREATED", "COMMENT_CREATED"], "creation_timestamp": 1,
				"last_updated_timestamp": 2, "description": "from the API", "status": "DISABLED", "model_name": "model",
				"http_url_spec": {"url": "https://example.com", "enable_ssl_verification": true, "secret": "s"},
				"job_spec": {"job_id": "7", "workspace_url": "https://example.com", "access_token": "t"}}`,
			"false and empty": `{"id": "w1", "events": [], "http_url_spec": {"enable_ssl_verification": false}, "job_spec": {}}`,
			"empty message":   `{}`,
		},
	},
	"test structs with every kind of scalar": {
		decode: func(t *testing.T, body string) (any, any) {
			return structToDataDecode[structToDataClassic](t, body), structToDataDecode[structToDataModular](t, body)
		},
		schemas: map[string]map[string]*schema.Schema{
			"schema of the twin": common.StructToSchema(structToDataClassic{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
				s["force_destroy"] = &schema.Schema{Type: schema.TypeBool, Optional: true, Default: false}
				s["comment"] = &schema.Schema{Type: schema.TypeString, Optional: true, Default: "none"}
				return s
			}),
			"computed fields": structToDataComputed(structToDataClassic{}, "text", "number", "size", "ratio", "flag", "level",
				"names", "unique", "labels", "nested", "items", "group", "parts", "created", "timeout"),
		},
		configs: map[string]map[string]any{
			"no plan":            nil,
			"nothing configured": {},
			"everything configured": {"title": "configured", "total": 9, "enabled": true, "text": "configured", "mode": "manual",
				"number": 9, "size": 9, "ratio": 9.5, "flag": true, "level": "LOW", "names": []any{"configured"},
				"unique": []any{"configured"}, "labels": map[string]any{"configured": "yes"},
				"nested": []any{map[string]any{"key": "configured"}}, "items": []any{map[string]any{"key": "configured"}},
				"group": []any{map[string]any{"key": "configured"}}, "created": "2020-01-01T00:00:00Z", "timeout": "1s"},
			"zero values configured": {"title": "", "total": 0, "enabled": false, "text": "", "number": 0, "ratio": 0.0, "flag": false},
		},
		responses: map[string]string{
			"all fields": `{"title": "t", "total": 3, "enabled": true, "text": "from the API", "mode": "fast", "number": 7, "size": 2,
				"ratio": 1.5, "flag": true, "level": "HIGH", "names": ["a", "b"], "unique": ["b", "a", "b"], "labels": {"k": "v"},
				"nested": {"key": "n", "count": 1, "flag": true, "names": ["x"]},
				"items": [{"key": "first"}, {"count": 2, "flag": true, "names": ["y", "z"]}, {}], "group": [{"key": "g", "count": 1}],
				"parts": [{}, {"key": "p"}, {"names": []}], "created": "2023-11-14T22:13:20.120Z", "timeout": "3.500s"}`,
			"explicit zeros": `{"title": "", "total": 0, "enabled": false, "text": "", "mode": "", "number": 0, "size": 0, "ratio": 0,
				"flag": false, "level": "", "names": [], "unique": [], "labels": {}, "nested": {}, "items": [], "group": []}`,
			"only false":    `{"enabled": false, "flag": false}`,
			"empty message": `{}`,
		},
	},
}

// TestStructToData_SameStateAsCommonStructToData decodes one response into a
// classic struct and into its modular twin, writes each into its own resource
// data, the classic one with common.StructToData, and compares the states.
func TestStructToData_SameStateAsCommonStructToData(t *testing.T) {
	for twinName, twin := range structToDataTwins {
		for schemaName, s := range twin.schemas {
			for configName, config := range twin.configs {
				for responseName, body := range twin.responses {
					for _, isNew := range []bool{true, false} {
						name := fmt.Sprintf("%s/%s/%s/%s/new=%t", twinName, schemaName, configName, responseName, isNew)
						t.Run(name, func(t *testing.T) {
							classic, modular := twin.decode(t, body)
							classicData := structToDataTestData(t, s, config, isNew)
							modularData := structToDataTestData(t, s, config, isNew)
							require.NoError(t, common.StructToData(classic, s, classicData))
							require.NoError(t, StructToData(modular, s, modularData))
							assert.Equal(t, classicData.State().Attributes, modularData.State().Attributes)
							for attribute := range s {
								assert.Equal(t, structToDataGet(classicData, attribute), structToDataGet(modularData, attribute), attribute)
							}
						})
					}
				}
			}
		}
	}
}

type structToDataTestItem struct {
	Key   *string
	Count *int64
	Names []string
	Oneof any
	Mask  *types.FieldMask[structToDataTestItem]
}

// structToDataTestMask is a field mask that is not empty.
var structToDataTestMask, _ = types.NewFieldMask[structToDataTestItem]("*")

// structToDataTestMessage has a field of every shape that StructToData
// supports, and those that it skips: hidden is unexported, Extra has no entry
// in the schema, and Mask, MaskRef and NoMask are field masks.
type structToDataTestMessage struct {
	Text    *string
	Number  *int64
	Size    *int
	Short   *int32
	Ratio   *float64
	Half    *float32
	Flag    *bool
	Level   structToDataTestEnum
	Grade   *structToDataTestEnum
	Plain   string
	Whole   int
	Sure    bool
	Names   []string
	Unique  []string
	Levels  []structToDataTestEnum
	Numbers []int64
	Ratios  []float64
	Flags   []bool
	Labels  map[string]string
	Nested  *structToDataTestItem
	Inline  structToDataTestItem
	Items   []structToDataTestItem
	Refs    []*structToDataTestItem
	Created *types.Time
	Timeout *types.Duration
	Started types.Time
	Wait    types.Duration
	Times   []types.Time
	Mask    types.FieldMask[structToDataTestItem]
	MaskRef *types.FieldMask[structToDataTestItem]
	NoMask  *types.FieldMask[structToDataTestItem]
	hidden  *string
	Extra   *string
}

// structToDataTestSchema returns a schema for structToDataTestMessage without
// an entry for Extra, where kind makes every field optional, required or computed.
func structToDataTestSchema(kind func(*schema.Schema)) map[string]*schema.Schema {
	block := &schema.Resource{Schema: map[string]*schema.Schema{
		"key": {Type: schema.TypeString, Optional: true}, "count": {Type: schema.TypeInt, Optional: true},
		"names": {Type: schema.TypeList, Optional: true, Elem: &schema.Schema{Type: schema.TypeString}},
	}}
	list := func(elem schema.ValueType) *schema.Schema {
		return &schema.Schema{Type: schema.TypeList, Elem: &schema.Schema{Type: elem}}
	}
	s := map[string]*schema.Schema{
		"text": {Type: schema.TypeString}, "number": {Type: schema.TypeInt}, "size": {Type: schema.TypeInt},
		"short": {Type: schema.TypeInt}, "ratio": {Type: schema.TypeFloat}, "half": {Type: schema.TypeFloat},
		"flag": {Type: schema.TypeBool}, "level": {Type: schema.TypeString}, "grade": {Type: schema.TypeString},
		"plain": {Type: schema.TypeString}, "whole": {Type: schema.TypeInt}, "sure": {Type: schema.TypeBool},
		"names": list(schema.TypeString), "levels": list(schema.TypeString), "numbers": list(schema.TypeInt),
		"ratios": list(schema.TypeFloat), "flags": list(schema.TypeBool),
		"unique": {Type: schema.TypeSet, Elem: &schema.Schema{Type: schema.TypeString}},
		"labels": {Type: schema.TypeMap, Elem: &schema.Schema{Type: schema.TypeString}},
		"nested": {Type: schema.TypeList, MaxItems: 1, Elem: block}, "inline": {Type: schema.TypeList, MaxItems: 1, Elem: block},
		"items": {Type: schema.TypeList, Elem: block}, "refs": {Type: schema.TypeList, Elem: block},
		"created": {Type: schema.TypeString}, "timeout": {Type: schema.TypeString}, "started": {Type: schema.TypeString},
		"wait": {Type: schema.TypeString}, "times": list(schema.TypeString), "hidden": {Type: schema.TypeString},
		"mask": {Type: schema.TypeString}, "mask_ref": {Type: schema.TypeString}, "no_mask": {Type: schema.TypeString},
	}
	for _, fieldSchema := range s {
		kind(fieldSchema)
	}
	return s
}

func structToDataOptional(s *schema.Schema) { s.Optional = true }
func structToDataRequired(s *schema.Schema) { s.Required = true }
func structToDataReadOnly(s *schema.Schema) { s.Optional, s.Computed = false, true }

var structToDataFullMessage = structToDataTestMessage{
	Text: structToDataPtr("a"), Number: structToDataPtr(int64(7)), Size: structToDataPtr(3), Short: structToDataPtr(int32(4)),
	Ratio: structToDataPtr(1.5), Half: structToDataPtr(float32(0.5)), Flag: structToDataPtr(true), Level: "HIGH",
	Grade: structToDataPtr(structToDataTestEnum("B")), Plain: "p", Whole: 5, Sure: true, Names: []string{"x", "y"},
	Unique: []string{"u", "u"}, Levels: []structToDataTestEnum{"LOW", "HIGH"}, Numbers: []int64{1, 2}, Ratios: []float64{0.25},
	Flags: []bool{true, false}, Labels: map[string]string{"k": "v"},
	Nested: &structToDataTestItem{Key: structToDataPtr("n"), Count: structToDataPtr(int64(1)), Names: []string{"m"}},
	Inline: structToDataTestItem{Key: structToDataPtr("i")}, Items: []structToDataTestItem{{Key: structToDataPtr("first")}, {}},
	Refs:    []*structToDataTestItem{nil, {Count: structToDataPtr(int64(2))}, nil},
	Created: &types.Time{Seconds: 1700000000, Nanos: 120000000}, Timeout: &types.Duration{Seconds: 3, Nanos: 500000000},
	Started: types.Time{Seconds: 1}, Wait: types.Duration{Seconds: -90}, Times: []types.Time{{}, {Seconds: 60, Nanos: 1}},
	Mask: *structToDataTestMask, MaskRef: structToDataTestMask, hidden: structToDataPtr("h"), Extra: structToDataPtr("e"),
}

// The states are flatmaps, as terraform.InstanceState holds them.
var (
	structToDataFullState = map[string]string{
		"text": "a", "number": "7", "size": "3", "short": "4", "ratio": "1.5", "half": "0.5", "flag": "true", "level": "HIGH",
		"grade": "B", "plain": "p", "whole": "5", "sure": "true", "names.#": "2", "names.0": "x", "names.1": "y",
		"unique.#": "1", "unique.2728031848": "u", "levels.#": "2", "levels.0": "LOW", "levels.1": "HIGH",
		"numbers.#": "2", "numbers.0": "1", "numbers.1": "2", "ratios.#": "1", "ratios.0": "0.25",
		"flags.#": "2", "flags.0": "true", "flags.1": "false", "labels.%": "1", "labels.k": "v",
		"nested.#": "1", "nested.0.key": "n", "nested.0.count": "1", "nested.0.names.#": "1", "nested.0.names.0": "m",
		"inline.#": "1", "inline.0.key": "i", "inline.0.count": "0", "inline.0.names.#": "0",
		"items.#": "2", "items.0.key": "first", "items.0.count": "0", "items.0.names.#": "0",
		"items.1.key": "", "items.1.count": "0", "items.1.names.#": "0",
		"refs.#": "1", "refs.0.key": "", "refs.0.count": "2", "refs.0.names.#": "0",
		"created": "2023-11-14T22:13:20.120Z", "timeout": "3.500s", "started": "1970-01-01T00:00:01Z", "wait": "-90s",
		"times.#": "2", "times.0": "1970-01-01T00:00:00Z", "times.1": "1970-01-01T00:01:00.000000001Z",
	}
	// A nil *string, *types.Time and *types.Duration is empty and skipped. A
	// nil *int64, *int, *int32, *float64, *float32 and *bool is 0 or false and
	// written, and so is a time or a duration that is not a pointer.
	structToDataZeroState = map[string]string{
		"number": "0", "size": "0", "short": "0", "ratio": "0", "half": "0", "flag": "false", "whole": "0", "sure": "false",
		"started": "1970-01-01T00:00:00Z", "wait": "0s",
		"inline.#": "1", "inline.0.key": "", "inline.0.count": "0", "inline.0.names.#": "0",
	}
	// A field that isn't optional is written even when it is empty, except
	// for a list of nested blocks.
	structToDataRequiredZeroState = map[string]string{
		"text": "", "level": "", "grade": "", "plain": "", "names.#": "0", "unique.#": "0", "levels.#": "0",
		"numbers.#": "0", "ratios.#": "0", "flags.#": "0", "labels.%": "0",
		"created": "", "timeout": "", "started": "1970-01-01T00:00:00Z", "wait": "0s", "times.#": "0",
		"number": "0", "size": "0", "short": "0", "ratio": "0", "half": "0", "flag": "false", "whole": "0", "sure": "false",
		"inline.#": "1", "inline.0.key": "", "inline.0.count": "0", "inline.0.names.#": "0",
	}
)

func TestStructToData_State(t *testing.T) {
	zeroPointers := structToDataTestMessage{
		Text: structToDataPtr(""), Number: structToDataPtr(int64(0)), Size: structToDataPtr(0), Short: structToDataPtr(int32(0)),
		Ratio: structToDataPtr(0.0), Half: structToDataPtr(float32(0)), Flag: structToDataPtr(false),
		Grade: structToDataPtr(structToDataTestEnum("")), Names: []string{}, Unique: []string{}, Levels: []structToDataTestEnum{},
		Numbers: []int64{}, Ratios: []float64{}, Flags: []bool{}, Labels: map[string]string{},
		Items: []structToDataTestItem{}, Refs: []*structToDataTestItem{nil, nil},
	}
	defaults := map[string]*schema.Schema{
		"text": {Type: schema.TypeString, Optional: true}, "force_destroy": {Type: schema.TypeBool, Optional: true, Default: false},
		"comment": {Type: schema.TypeString, Optional: true, Default: "none"}, "retries": {Type: schema.TypeInt, Optional: true, Default: 0},
		"computed": {Type: schema.TypeInt, Optional: true, Computed: true, Default: 0}, "required": {Type: schema.TypeInt, Required: true},
		"no_default": {Type: schema.TypeInt, Optional: true},
	}
	optional, readOnly := structToDataTestSchema(structToDataOptional), structToDataTestSchema(structToDataReadOnly)
	for name, c := range map[string]struct {
		schema   map[string]*schema.Schema
		config   map[string]any
		existing bool
		result   any
		want     map[string]string
	}{
		"every shape":                  {schema: optional, result: structToDataFullMessage, want: structToDataFullState},
		"pointer to a struct":          {schema: optional, result: &structToDataFullMessage, want: structToDataFullState},
		"nil pointers are zero values": {schema: optional, result: structToDataTestMessage{}, want: structToDataZeroState},
		"pointers to zero values give the state of nil pointers": {schema: optional, result: zeroPointers, want: structToDataZeroState},
		"a nil time or duration is empty like a nil string": {result: structToDataTestMessage{}, want: map[string]string{"timeout": ""},
			schema: map[string]*schema.Schema{"created": {Type: schema.TypeString, Optional: true}, "timeout": {Type: schema.TypeString, Computed: true}}},
		"field masks are skipped also when their schema is a block": {result: structToDataFullMessage, want: map[string]string{},
			schema: map[string]*schema.Schema{"mask": optional["nested"], "mask_ref": readOnly["items"], "no_mask": readOnly["nested"]}},
		"a nested field mask is skipped": {result: structToDataTestMessage{Nested: &structToDataTestItem{Mask: structToDataTestMask}}, want: map[string]string{"nested.#": "1"},
			schema: map[string]*schema.Schema{"nested": {Type: schema.TypeList, Computed: true, Elem: &schema.Resource{Schema: map[string]*schema.Schema{"mask": readOnly["items"]}}}}},
		"required and empty":                     {schema: structToDataTestSchema(structToDataRequired), result: structToDataTestMessage{}, want: structToDataRequiredZeroState},
		"computed and empty":                     {schema: readOnly, result: &zeroPointers, want: structToDataRequiredZeroState},
		"existing resource takes only the bools": {schema: optional, existing: true, result: structToDataFullMessage, want: map[string]string{"flag": "true", "sure": "true"}},
		"existing resource takes what is configured": {schema: optional, existing: true,
			config: map[string]any{"text": "old", "number": 1, "names": []any{"old"}, "labels": map[string]any{"old": "label"},
				"nested": []any{map[string]any{"key": "old"}}, "flag": true},
			result: structToDataTestMessage{Text: structToDataPtr("new"), Number: structToDataPtr(int64(2)), Names: []string{"new"},
				Labels: map[string]string{"new": "label"}, Nested: &structToDataTestItem{Key: structToDataPtr("new")}, Whole: 5},
			want: map[string]string{"text": "new", "number": "2", "names.#": "1", "names.0": "new", "labels.%": "1",
				"labels.new": "label", "nested.#": "1", "nested.0.key": "new", "nested.0.count": "0", "nested.0.names.#": "0",
				"flag": "false", "sure": "false"}},
		"existing resource keeps what the response leaves empty": {schema: optional, existing: true,
			config: map[string]any{"text": "old", "number": 1, "names": []any{"old"}, "nested": []any{map[string]any{"key": "old"}}},
			result: structToDataTestMessage{},
			want: map[string]string{"text": "old", "number": "0", "names.#": "1", "names.0": "old", "nested.#": "1",
				"nested.0.key": "old", "nested.0.count": "0", "nested.0.names.#": "0", "flag": "false", "sure": "false"}},
		"existing resource takes every computed field": {schema: readOnly, existing: true, result: structToDataFullMessage, want: structToDataFullState},
		"defaults without a plan are the zero values": {schema: defaults, result: structToDataTestMessage{Text: structToDataPtr("a")},
			want: map[string]string{"text": "a", "force_destroy": "false", "comment": "", "retries": "0"}},
		"defaults of the plan stay": {schema: defaults, config: map[string]any{}, result: structToDataTestMessage{},
			want: map[string]string{"force_destroy": "false", "comment": "none", "retries": "0", "computed": "0"}},
		"configured values win over defaults": {schema: defaults, existing: true,
			config: map[string]any{"force_destroy": true, "comment": "configured", "retries": 3, "required": 1},
			result: structToDataTestMessage{},
			want:   map[string]string{"force_destroy": "true", "comment": "configured", "retries": "3", "required": "1", "computed": "0"}},
		"modelregistry.RegisteredModelDatabricks": {schema: common.StructToSchema(ml.CreateModelRequest{}, nil),
			result: &modelregistry.RegisteredModelDatabricks{Name: structToDataPtr("model"), Id: structToDataPtr("abc"),
				PermissionLevel: modelregistry.PermissionLevel_CanEdit, CreationTimestamp: structToDataPtr(int64(1)),
				Tags: []modelregistry.RegisteredModelTag{{Key: structToDataPtr("team"), Value: structToDataPtr("ml")}, {}}},
			want: map[string]string{"name": "model", "tags.#": "2", "tags.0.key": "team", "tags.0.value": "ml"}},
		"modelregistry.RegistryWebhook": {schema: common.StructToSchema(ml.RegistryWebhook{}, nil),
			result: modelregistry.RegistryWebhook{Id: structToDataPtr("w1"), Status: modelregistry.RegistryWebhookStatus_Active,
				Events:      []modelregistry.RegistryWebhookEvent{modelregistry.RegistryWebhookEvent_CommentCreated},
				HttpUrlSpec: &modelregistry.HttpUrlSpec{Url: structToDataPtr("https://example.com"), Secret: structToDataPtr("s")}},
			want: map[string]string{"status": "ACTIVE", "events.#": "1", "events.0": "COMMENT_CREATED",
				"creation_timestamp": "0", "last_updated_timestamp": "0", "http_url_spec.#": "1",
				"http_url_spec.0.url": "https://example.com", "http_url_spec.0.enable_ssl_verification": "false"}},
	} {
		t.Run(name, func(t *testing.T) {
			d := structToDataTestData(t, c.schema, c.config, !c.existing)
			require.NoError(t, StructToData(c.result, c.schema, d))
			got := d.State().Attributes
			delete(got, "id")
			assert.Equal(t, c.want, got)
		})
	}
}

// structToDataTestUnsupported has a field of every kind of unsupported type.
type structToDataTestUnsupported struct {
	Oneof    any
	Stream   io.Reader
	Raw      []byte // as json.RawMessage
	Counts   map[string]int64
	ByNumber map[int]string
	Deep     **string
	Matrix   [][]string
	Texts    []*string
	Unsigned *uint32
	Channel  chan string
	Instants []*types.Time
}

func TestStructToData_UnsupportedFieldType(t *testing.T) {
	for name, fieldType := range map[string]string{
		"oneof": "interface {}", "stream": "io.Reader", "raw": "[]uint8", "counts": "map[string]int64",
		"by_number": "map[int]string", "deep": "**string", "matrix": "[][]string", "texts": "[]*string",
		"unsigned": "*uint32", "channel": "chan string", "instants": "[]*types.Time",
	} {
		t.Run(name, func(t *testing.T) {
			// The error doesn't depend on what the field holds or on its schema.
			for _, fieldSchema := range []*schema.Schema{{Type: schema.TypeString, Optional: true}, {Type: schema.TypeInt, Computed: true}} {
				s := map[string]*schema.Schema{name: fieldSchema}
				for _, isNew := range []bool{true, false} {
					err := StructToData(structToDataTestUnsupported{}, s, structToDataTestData(t, s, nil, isNew))
					assert.EqualError(t, err, name+": type "+fieldType+" is not supported")
				}
			}
		})
	}
}

func TestStructToData_Errors(t *testing.T) {
	text := map[string]*schema.Schema{"text": {Type: schema.TypeString, Optional: true}}
	tag := &modelregistry.RegisteredModelTag{}
	oneofBlock := &schema.Resource{Schema: map[string]*schema.Schema{"oneof": {Type: schema.TypeString, Optional: true}}}
	for name, c := range map[string]struct {
		result any
		schema map[string]*schema.Schema
		// dataSchema is the schema of the resource data, when it isn't schema.
		dataSchema map[string]*schema.Schema
		want       string
	}{
		"nil": {result: nil, schema: text, want: "result is nil, want a struct or a pointer to one"},
		"nil pointer": {result: (*modelregistry.RegisteredModelDatabricks)(nil), schema: text,
			want: "result is a nil *modelregistry.RegisteredModelDatabricks, want a struct or a pointer to one"},
		"string": {result: "model", schema: text, want: "result is a string, want a struct or a pointer to one"},
		"slice of structs": {result: []modelregistry.RegisteredModelTag{{}}, schema: text,
			want: "result is a []modelregistry.RegisteredModelTag, want a struct or a pointer to one"},
		"pointer to a pointer": {result: &tag, schema: text, want: "result is a **modelregistry.RegisteredModelTag, want a struct or a pointer to one"},
		"message without a nested block": {result: structToDataTestMessage{},
			schema: map[string]*schema.Schema{"nested": {Type: schema.TypeString, Optional: true}},
			want:   "nested: a message needs a nested block, but the schema is a TypeString without one"},
		"messages in a list of strings": {result: modelregistry.RegisteredModelDatabricks{},
			schema: map[string]*schema.Schema{"tags": {Type: schema.TypeList, Optional: true, Elem: &schema.Schema{Type: schema.TypeString}}},
			want:   "tags: a message needs a nested block, but the schema is a TypeList without one"},
		"unsupported type in a nested message": {result: structToDataTestMessage{Nested: &structToDataTestItem{}},
			schema: map[string]*schema.Schema{"nested": {Type: schema.TypeList, Optional: true, Elem: oneofBlock}},
			want:   "nested: oneof: type interface {} is not supported"},
		"unsupported type in a slice of messages": {result: structToDataTestMessage{Items: []structToDataTestItem{{}}},
			schema: map[string]*schema.Schema{"items": {Type: schema.TypeList, Computed: true, Elem: oneofBlock}},
			want:   "items: oneof: type interface {} is not supported"},
		"duration outside the range of its JSON form": {result: structToDataTestMessage{Timeout: &types.Duration{Seconds: 1, Nanos: -1}},
			schema: map[string]*schema.Schema{"timeout": {Type: schema.TypeString, Optional: true}},
			want:   "timeout: duration seconds and nanoseconds have different signs"},
		"time in a list outside the range of its JSON form": {result: structToDataTestMessage{Times: []types.Time{{}, {Nanos: -1}}},
			schema: map[string]*schema.Schema{"times": {Type: schema.TypeList, Computed: true, Elem: &schema.Schema{Type: schema.TypeString}}},
			want:   "times: timestamp nanoseconds -1 outside protobuf range"},
		"d.Set fails on a list for a string": {result: structToDataTestMessage{Names: []string{"a"}},
			schema: map[string]*schema.Schema{"names": {Type: schema.TypeString, Computed: true}},
			want:   "names: names: '' expected type 'string', got unconvertible type '[]interface {}'"},
		"d.Set fails on a field that the resource data lacks": {result: structToDataTestMessage{Text: structToDataPtr("a")},
			schema: map[string]*schema.Schema{"text": {Type: schema.TypeString, Computed: true}}, dataSchema: oneofBlock.Schema,
			want: `text: Invalid address to set: []string{"text"}`},
		"d.Set fails on the default of a field that the resource data lacks": {result: structToDataTestMessage{},
			schema:     map[string]*schema.Schema{"force_destroy": {Type: schema.TypeBool, Optional: true, Default: false}},
			dataSchema: text, want: `force_destroy: Invalid address to set: []string{"force_destroy"}`},
	} {
		t.Run(name, func(t *testing.T) {
			dataSchema := c.dataSchema
			if dataSchema == nil {
				dataSchema = c.schema
			}
			for _, isNew := range []bool{true, false} {
				d := structToDataTestData(t, dataSchema, nil, isNew)
				assert.ErrorContains(t, StructToData(c.result, c.schema, d), c.want)
				assert.Equal(t, map[string]string{"id": "id"}, d.State().Attributes)
			}
		})
	}
}
