package mlflow

import (
	"encoding/json"
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

// A failed PATCH still hands Terraform the planned description: Update returns the error
// without marking the change partial, so the state records the new value although the
// service kept the old one. A replay journey can't show this, because its plan checks
// run after a refresh that reads the old description back. The bodies are synthetic:
// inferred from the MLflow Model Registry API.
func TestModelUpdatePatchErrorKeepsPlannedDescriptionInState(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/mlflow/registered-models/update",
				ExpectedRequest: json.RawMessage(`{"name": "failed_update_model",
					"description": "After the update"}`),
				Response: `{"error_code": "INVALID_PARAMETER_VALUE",
					"message": "Invalid description for registered model 'failed_update_model'."}`,
				Status: 400,
			},
		},
		Resource: ResourceMlflowModel(),
		Update:   true,
		ID:       "failed_update_model",
		InstanceState: map[string]string{
			"name":                "failed_update_model",
			"description":         "Before the update",
			"registered_model_id": "0a1b2c3d4e5f60718293a4b5c6d7e891",
		},
		HCL: `
		name        = "failed_update_model"
		description = "After the update"
		`,
	}.Apply(t)

	assert.EqualError(t, err, "Invalid description for registered model 'failed_update_model'.")
	assert.Equal(t, "After the update", d.State().Attributes["description"])
	assert.Equal(t, "failed_update_model", d.State().ID)
}
