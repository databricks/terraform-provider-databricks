package mlflow

import (
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/qa"
)

// These tests pin what an update does with tags. Update builds an ml.UpdateModelRequest,
// which has only name and description, so the PATCH never carries tags and no other
// call sets them. The read after the update then decides the state (StructToData in
// common/reflect_resource.go). A changed tag leaves a diff after apply, which the
// replay harness rejects, so the tag updates are Go tests and not journeys in
// qa/replay/resources/mlflow_model_updates.
//
// Fixtures are synthetic: request and response bodies are literal wire JSON inferred
// from the MLflow Model Registry API, not recorded. They are not classic SDK structs or
// mocks, so the tests still apply once the resource uses another client.

// tagsUpdate is one update of a model: the state before it, the new configuration, the
// PATCH body it must send, and the responses of the PATCH and of the GET after it.
type tagsUpdate struct {
	state         map[string]string
	hcl           string
	patchBody     string
	patchResponse string
	getResponse   string
}

// apply runs the update and returns the state after it. The fixture list starts with
// the host-metadata stub so that the fixture server works on this slice, where it
// blanks each fixture it serves: a fixture that still has a method was never requested.
func (u tagsUpdate) apply(t *testing.T) map[string]string {
	fixtures := []qa.HTTPFixture{
		{
			Method:       "GET",
			Resource:     "/.well-known/databricks-config",
			Status:       404,
			Response:     `{"message": "Not Found"}`,
			ReuseRequest: true,
		},
		{
			Method:          "PATCH",
			Resource:        "/api/2.0/mlflow/registered-models/update",
			ExpectedRequest: json.RawMessage(u.patchBody),
			Response:        u.patchResponse,
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/mlflow/databricks/registered-models/get?name=pinned-model",
			Response: u.getResponse,
		},
	}
	d, err := qa.ResourceFixture{
		Fixtures:      fixtures,
		Resource:      ResourceMlflowModel(),
		Update:        true,
		ID:            "pinned-model",
		InstanceState: u.state,
		HCL:           u.hcl,
	}.Apply(t)
	require.NoError(t, err)
	assert.Empty(t, fixtures[1].Method, "the PATCH was not sent")
	assert.Empty(t, fixtures[2].Method, "the GET after the update was not sent")
	return d.State().Attributes
}

// expectNextPlan asserts the diff that planning the same configuration against state
// gives, where an empty map means no diff. The empty fixture list fails any API call.
func (u tagsUpdate) expectNextPlan(t *testing.T, state map[string]string,
	diff map[string]*terraform.ResourceAttrDiff) {
	_, err := qa.ResourceFixture{
		Fixtures:      []qa.HTTPFixture{},
		Resource:      ResourceMlflowModel(),
		InstanceState: state,
		HCL:           u.hcl,
		ExpectedDiff:  diff,
	}.Apply(t)
	require.NoError(t, err)
}

// TestModelUpdateTagsChangeLeavesDiff pins a changed tag value. The PATCH carries name
// and the unchanged description, without tags. The GET after it returns the old tag,
// which state adopts, so the next plan shows the same change again.
func TestModelUpdateTagsChangeLeavesDiff(t *testing.T) {
	update := tagsUpdate{
		state: map[string]string{
			"name":                           "pinned-model",
			"description":                    "first description",
			"registered_model_id":            "0123456789abcdef0123456789abcdef",
			"provider_config.#":              "1",
			"provider_config.0.workspace_id": "12345",
			"tags.#":                         "1",
			"tags.0.key":                     "team",
			"tags.0.value":                   "ml",
		},
		hcl: `
		name        = "pinned-model"
		description = "first description"
		tags {
			key   = "team"
			value = "platform"
		}`,
		patchBody: `{"name": "pinned-model", "description": "first description"}`,
		patchResponse: `{"registered_model": {
			"name": "pinned-model",
			"creation_timestamp": 1700000000000,
			"last_updated_timestamp": 1700000001000,
			"user_id": "replay-admin@example.com",
			"description": "first description",
			"tags": [{"key": "team", "value": "ml"}]}}`,
		getResponse: `{"registered_model_databricks": {
			"name": "pinned-model",
			"creation_timestamp": 1700000000000,
			"last_updated_timestamp": 1700000001000,
			"user_id": "replay-admin@example.com",
			"description": "first description",
			"tags": [{"key": "team", "value": "ml"}],
			"id": "0123456789abcdef0123456789abcdef",
			"permission_level": "CAN_MANAGE"}}`,
	}
	state := update.apply(t)
	assert.Equal(t, map[string]string{
		"id":                             "pinned-model",
		"name":                           "pinned-model",
		"description":                    "first description",
		"registered_model_id":            "0123456789abcdef0123456789abcdef",
		"provider_config.#":              "1",
		"provider_config.0.workspace_id": "12345",
		"tags.#":                         "1",
		"tags.0.key":                     "team",
		"tags.0.value":                   "ml",
	}, state)
	update.expectNextPlan(t, state, map[string]*terraform.ResourceAttrDiff{
		"tags.0.value": {Old: "ml", New: "platform"},
	})
}

// TestModelUpdateTagsAddedOrRemoved pins adding tags to a model without tags and
// removing all tags. Each sends a PATCH with only name. Unlike a changed tag, neither
// leaves a diff: state follows the configuration and the next plan is empty, while the
// API still has no tags after the add and still has the tag after the removal.
func TestModelUpdateTagsAddedOrRemoved(t *testing.T) {
	cases := []struct {
		name   string
		update tagsUpdate
		state  map[string]string
	}{
		{
			name: "added_to_a_model_without_tags",
			update: tagsUpdate{
				state: map[string]string{
					"name":                           "pinned-model",
					"registered_model_id":            "0123456789abcdef0123456789abcdef",
					"provider_config.#":              "1",
					"provider_config.0.workspace_id": "12345",
				},
				hcl: `
				name = "pinned-model"
				tags {
					key   = "team"
					value = "ml"
				}`,
				patchBody: `{"name": "pinned-model"}`,
				patchResponse: `{"registered_model": {
					"name": "pinned-model",
					"creation_timestamp": 1700000000000,
					"last_updated_timestamp": 1700000001000,
					"user_id": "replay-admin@example.com"}}`,
				getResponse: `{"registered_model_databricks": {
					"name": "pinned-model",
					"creation_timestamp": 1700000000000,
					"last_updated_timestamp": 1700000001000,
					"user_id": "replay-admin@example.com",
					"id": "0123456789abcdef0123456789abcdef",
					"permission_level": "CAN_MANAGE"}}`,
			},
			state: map[string]string{
				"id":                             "pinned-model",
				"name":                           "pinned-model",
				"registered_model_id":            "0123456789abcdef0123456789abcdef",
				"provider_config.#":              "1",
				"provider_config.0.workspace_id": "12345",
				"tags.#":                         "1",
				"tags.0.key":                     "team",
				"tags.0.value":                   "ml",
			},
		},
		{
			name: "all_removed",
			update: tagsUpdate{
				state: map[string]string{
					"name":                           "pinned-model",
					"registered_model_id":            "0123456789abcdef0123456789abcdef",
					"provider_config.#":              "1",
					"provider_config.0.workspace_id": "12345",
					"tags.#":                         "1",
					"tags.0.key":                     "team",
					"tags.0.value":                   "ml",
				},
				hcl:       `name = "pinned-model"`,
				patchBody: `{"name": "pinned-model"}`,
				patchResponse: `{"registered_model": {
					"name": "pinned-model",
					"creation_timestamp": 1700000000000,
					"last_updated_timestamp": 1700000001000,
					"user_id": "replay-admin@example.com",
					"tags": [{"key": "team", "value": "ml"}]}}`,
				getResponse: `{"registered_model_databricks": {
					"name": "pinned-model",
					"creation_timestamp": 1700000000000,
					"last_updated_timestamp": 1700000001000,
					"user_id": "replay-admin@example.com",
					"tags": [{"key": "team", "value": "ml"}],
					"id": "0123456789abcdef0123456789abcdef",
					"permission_level": "CAN_MANAGE"}}`,
			},
			state: map[string]string{
				"id":                             "pinned-model",
				"name":                           "pinned-model",
				"registered_model_id":            "0123456789abcdef0123456789abcdef",
				"provider_config.#":              "1",
				"provider_config.0.workspace_id": "12345",
				"tags.#":                         "0",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			state := tc.update.apply(t)
			assert.Equal(t, tc.state, state)
			tc.update.expectNextPlan(t, state, map[string]*terraform.ResourceAttrDiff{})
		})
	}
}
