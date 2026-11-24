// Request and response bodies in this file are literal wire JSON, so that the tests pin
// the bytes on the wire whichever client the resource uses. The responses are synthetic:
// inferred from the MLflow Model Registry API, not recorded. Where a fixture departs from
// that API, a comment on it says how.

package mlflow

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestModelCreateMVP(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "POST",
				Resource:        "/api/2.0/mlflow/registered-models/create",
				ExpectedRequest: json.RawMessage(`{"name": "xyz"}`),
				Response:        `{"registered_model": {"name": "xyz"}}`,
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/mlflow/databricks/registered-models/get?name=xyz",
				Response:     `{"registered_model_databricks": {"name": "xyz", "id": "123"}}`,
				ReuseRequest: true,
			},
		},
		Resource: ResourceMlflowModel(),
		Create:   true,
		HCL: `
		name = "xyz"
		`,
	}.ApplyAndExpectData(t, map[string]any{"id": "xyz", "registered_model_id": "123"})
}

func TestModelCreateWithTags(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/mlflow/registered-models/create",
				ExpectedRequest: json.RawMessage(`{
					"name": "xyz",
					"tags": [
						{"key": "key1", "value": "value1"},
						{"key": "key2", "value": "value2"}
					]
				}`),
				Response: `{"registered_model": {
					"name": "xyz",
					"tags": [
						{"key": "key1", "value": "value1"},
						{"key": "key2", "value": "value2"}
					]
				}}`,
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/databricks/registered-models/get?name=xyz",
				Response: `{"registered_model_databricks": {
					"name": "xyz",
					"tags": [
						{"key": "key1", "value": "value1"},
						{"key": "key2", "value": "value2"}
					]
				}}`,
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/databricks/registered-models/get?name=xyz",
				Response: `{"registered_model_databricks": {
					"name": "xyz",
					"tags": [
						{"key": "key1", "value": "value1"},
						{"key": "key2", "value": "value2"}
					]
				}}`,
			},
		},
		Resource: ResourceMlflowModel(),
		Create:   true,
		HCL: `
		name = "xyz"
		tags {
				key = "key1"
				value = "value1"
		    }
		tags {
				key = "key2"
				value = "value2"
			  }
		`,
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "xyz", d.Id(), "Resource ID should not be empty")
	assert.Equal(t, "xyz", d.Get("name"), "Name should be set")
	assert.Equal(t, d.Get("name"), d.Id(), "Name and Id should match")

}

func TestModelCreatePostError(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/mlflow/registered-models/create",
				ExpectedRequest: json.RawMessage(`{
					"name": "xyz",
					"tags": [
						{"key": "key1", "value": "value1"},
						{"key": "key2", "value": "value2"}
					]
				}`),
				// A 400 with a success body, which the API doesn't send.
				Response: `{"registered_model": {
					"name": "xyz",
					"tags": [
						{"key": "key1", "value": "value1"},
						{"key": "key2", "value": "value2"}
					]
				}}`,
				Status: 400,
			},
		},
		Resource: ResourceMlflowModel(),
		Create:   true,
		HCL: `
		name = "xyz"
		tags {
				key = "key1"
				value = "value1"
		    }
		tags {
				key = "key2"
				value = "value2"
			  }
		`,
	}.Apply(t)

	assert.Error(t, err)
	// The body has neither error_code nor message, so the error text is empty.
	assert.EqualError(t, err, "")
}

func TestModelRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/databricks/registered-models/get?name=xyz",
				Response: `{"registered_model_databricks": {"name": "xyz"}}`,
			},
		},
		Resource: ResourceMlflowModel(),
		Read:     true,
		ID:       "xyz",
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "xyz", d.Id(), "Resource ID should not be empty")
}

func TestModelReadGetError(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/databricks/registered-models/get?name=xyz",
				// A 400 with a success body, which the API doesn't send.
				Response: `{"registered_model_databricks": {"name": "xyz"}}`,
				Status:   400,
			},
		},
		Resource: ResourceMlflowModel(),
		Read:     true,
		ID:       "xyz",
	}.Apply(t)

	assert.Error(t, err)
	// The body has neither error_code nor message, so the error text is empty.
	assert.EqualError(t, err, "")
}

func TestModelUpdate(t *testing.T) {
	newDescription := "updatedddescription"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/mlflow/registered-models/update",
				ExpectedRequest: json.RawMessage(
					`{"name": "xyz", "description": "updatedddescription"}`),
				// A bare model, where the API wraps it in registered_model. The provider
				// doesn't read the response.
				Response: `{
					"name": "xyz",
					"description": "thedescription",
					"tags": [
						{"key": "key1", "value": "value1"},
						{"key": "key2", "value": "value2"}
					]
				}`,
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/databricks/registered-models/get?name=xyz",
				Response: `{"registered_model_databricks": {
					"name": "xyz", "description": "updatedddescription"
				}}`,
			},
		},
		Resource:    ResourceMlflowModel(),
		Update:      true,
		RequiresNew: true,
		ID:          "xyz",
		State: map[string]any{
			"name": "xyz",
		},
		HCL: fmt.Sprintf(`
		name = "xyz"
		description = "%s"
		`, newDescription),
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "xyz", d.Id(), "Resource ID should not be empty")
	assert.Equal(t, newDescription, d.Get("description"), "Description should be updated")
}

func TestModelUpdatePatchError(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/mlflow/registered-models/update",
				ExpectedRequest: json.RawMessage(
					`{"name": "xyz", "description": "updateddescription"}`),
				// A 400 with a bare model as its body, which the API doesn't send.
				Response: `{
					"name": "xyz",
					"description": "updateddescription",
					"tags": [
						{"key": "key1", "value": "value1"},
						{"key": "key2", "value": "value2"}
					]
				}`,
				Status: 400,
			},
		},
		Resource:    ResourceMlflowModel(),
		Update:      true,
		RequiresNew: true,
		ID:          "xyz",
		State: map[string]any{
			"name": "xyz",
		},
		HCL: `
		name = "xyz"
		description = "updateddescription"
		`,
	}.Apply(t)

	assert.Error(t, err)
	// The body has neither error_code nor message, so the error text is empty.
	assert.EqualError(t, err, "")
}
func TestModelDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/mlflow/registered-models/delete?name=xyz",
				Response: nil,
				Status:   200,
			},
		},
		Resource: ResourceMlflowModel(),
		Delete:   true,
		ID:       "xyz",
		HCL: `
		name = "xyz"
		`,
	}.Apply(t)

	assert.NoError(t, err)
	assert.Equal(t, "xyz", d.Id(), "Resource ID should not be empty")
}

func TestModelDeleteError(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/mlflow/registered-models/delete?name=xyz",
				Status:   400,
			},
		},
		Resource: ResourceMlflowModel(),
		Delete:   true,
		ID:       "xyz",
		HCL: `
		name = "xyz"
		`,
	}.Apply(t)

	assert.Error(t, err)
}

func TestMlflowModelCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceMlflowModel())
}
