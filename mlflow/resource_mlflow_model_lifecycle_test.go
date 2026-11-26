package mlflow

import (
	"encoding/json"
	"sync/atomic"
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The tests in this file pin databricks_mlflow_model behavior that the replay journeys
// in qa/replay/resources/mlflow_model_lifecycle cannot express. Requests and responses
// are literal JSON rather than classic SDK structs, so the tests do not depend on the
// client the resource uses. Responses are synthetic, inferred from the MLflow Model
// Registry API, not recorded.

// countedResponse is a fixture response that counts how often the fixture server sent
// it. qa.ResourceFixture does not fail on a fixture that no request used, so a call
// whose response the resource ignores needs the count to prove it was sent.
type countedResponse struct {
	body string
	sent *atomic.Int32
}

func (r countedResponse) MarshalJSON() ([]byte, error) {
	r.sent.Add(1)
	return []byte(r.body), nil
}

// A replay journey matches the decoded query, so it cannot tell "+" from "%20" or "/"
// from "%2F". qa.HTTPFixture matches the request URI byte for byte. The classic client
// form-encodes the name in the query of the GET and the DELETE, with the space as "+",
// and the POST and PATCH bodies carry the name unchanged.
func TestModelLifecycleNameWithReservedCharacters(t *testing.T) {
	const (
		name      = "lifecycle a&b?c#d%e/f+gé"
		query     = "?name=lifecycle+a%26b%3Fc%23d%25e%2Ff%2Bg%C3%A9"
		getURI    = "/api/2.0/mlflow/databricks/registered-models/get" + query
		deleteURI = "/api/2.0/mlflow/registered-models/delete" + query
		modelID   = "293a4b5c6d7e8f010203040506070809"
	)

	t.Run("create", func(t *testing.T) {
		qa.ResourceFixture{
			Fixtures: []qa.HTTPFixture{
				{
					Method:          "POST",
					Resource:        "/api/2.0/mlflow/registered-models/create",
					ExpectedRequest: json.RawMessage(`{"name": "lifecycle a&b?c#d%e/f+gé"}`),
					Response: `{"registered_model": {
						"name": "lifecycle a&b?c#d%e/f+gé",
						"creation_timestamp": 1700000000701,
						"last_updated_timestamp": 1700000000701,
						"user_id": "replay-admin@example.com"}}`,
				},
				{
					Method:   "GET",
					Resource: getURI,
					Response: `{"registered_model_databricks": {
						"name": "lifecycle a&b?c#d%e/f+gé",
						"id": "293a4b5c6d7e8f010203040506070809",
						"creation_timestamp": 1700000000701,
						"last_updated_timestamp": 1700000000701,
						"user_id": "replay-admin@example.com",
						"permission_level": "CAN_MANAGE"}}`,
				},
			},
			Resource: ResourceMlflowModel(),
			Create:   true,
			State:    map[string]any{"name": name},
		}.ApplyAndExpectData(t, map[string]any{
			"id":                  name,
			"name":                name,
			"registered_model_id": modelID,
		})
	})

	t.Run("update", func(t *testing.T) {
		var patches atomic.Int32
		qa.ResourceFixture{
			Fixtures: []qa.HTTPFixture{
				{
					Method:   "PATCH",
					Resource: "/api/2.0/mlflow/registered-models/update",
					ExpectedRequest: json.RawMessage(`{"name": "lifecycle a&b?c#d%e/f+gé",
						"description": "Described"}`),
					Response: countedResponse{`{"registered_model": {
						"name": "lifecycle a&b?c#d%e/f+gé",
						"description": "Described",
						"creation_timestamp": 1700000000701,
						"last_updated_timestamp": 1700000000702,
						"user_id": "replay-admin@example.com"}}`, &patches},
				},
				{
					Method:   "GET",
					Resource: getURI,
					Response: `{"registered_model_databricks": {
						"name": "lifecycle a&b?c#d%e/f+gé",
						"id": "293a4b5c6d7e8f010203040506070809",
						"description": "Described",
						"creation_timestamp": 1700000000701,
						"last_updated_timestamp": 1700000000702,
						"user_id": "replay-admin@example.com",
						"permission_level": "CAN_MANAGE"}}`,
				},
			},
			Resource: ResourceMlflowModel(),
			Update:   true,
			ID:       name,
			InstanceState: map[string]string{
				"name":                name,
				"registered_model_id": modelID,
			},
			State: map[string]any{"name": name, "description": "Described"},
		}.ApplyAndExpectData(t, map[string]any{
			"id":          name,
			"name":        name,
			"description": "Described",
		})
		assert.Equal(t, int32(1), patches.Load(), "PATCH calls")
	})

	t.Run("delete", func(t *testing.T) {
		var deletes atomic.Int32
		qa.ResourceFixture{
			Fixtures: []qa.HTTPFixture{
				{
					Method:   "DELETE",
					Resource: deleteURI,
					Response: countedResponse{`{}`, &deletes},
				},
			},
			Resource: ResourceMlflowModel(),
			Delete:   true,
			ID:       name,
			State:    map[string]any{"name": name},
		}.ApplyNoError(t)
		assert.Equal(t, int32(1), deletes.Load(), "DELETE calls")
	})
}

// The id is the name in the create response, not the configured name. When the two
// differ in letter case, the GET after the create asks for the response's name, state
// takes that name, and the next plan replaces the model to restore the configured name.
// A replay journey cannot express this, because its harness requires an empty plan after
// every apply.
func TestModelLifecycleCreateResponseNameInAnotherLetterCase(t *testing.T) {
	const config = `name = "Lifecycle-Case"`

	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "POST",
				Resource:        "/api/2.0/mlflow/registered-models/create",
				ExpectedRequest: json.RawMessage(`{"name": "Lifecycle-Case"}`),
				Response: `{"registered_model": {
					"name": "lifecycle-case",
					"creation_timestamp": 1700000000801,
					"last_updated_timestamp": 1700000000801,
					"user_id": "replay-admin@example.com"}}`,
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/databricks/registered-models/get?name=lifecycle-case",
				Response: `{"registered_model_databricks": {
					"name": "lifecycle-case",
					"id": "3a4b5c6d7e8f01020304050607080910",
					"creation_timestamp": 1700000000801,
					"last_updated_timestamp": 1700000000801,
					"user_id": "replay-admin@example.com",
					"permission_level": "CAN_MANAGE"}}`,
			},
		},
		Resource: ResourceMlflowModel(),
		Create:   true,
		HCL:      config,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "lifecycle-case", d.Id())
	assert.Equal(t, "lifecycle-case", d.Get("name"))
	assert.Equal(t, "3a4b5c6d7e8f01020304050607080910", d.Get("registered_model_id"))

	// The plan for the same config, from the state the create left.
	qa.ResourceFixture{
		Resource:      ResourceMlflowModel(),
		ID:            "lifecycle-case",
		InstanceState: d.State().Attributes,
		HCL:           config,
		ExpectedDiff: map[string]*terraform.ResourceAttrDiff{
			"name": {Old: "lifecycle-case", New: "Lifecycle-Case", RequiresNew: true},
			"registered_model_id": {
				Old: "3a4b5c6d7e8f01020304050607080910", NewComputed: true,
			},
			"provider_config.#": {Old: "1", NewComputed: true},
		},
	}.ApplyNoError(t)
}
