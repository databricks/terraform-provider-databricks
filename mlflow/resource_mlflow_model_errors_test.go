package mlflow

import (
	"encoding/json"
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

// A failed update still hands Terraform the planned description: Update returns the error
// without marking the change partial, so the state records the new value whether or not
// the service has it. A replay journey shows this only through its plan without refresh,
// because its state checks run after a refresh that reads the description back. The
// bodies are synthetic: inferred from the MLflow Model Registry API.
func TestModelUpdateErrorsKeepPlannedDescriptionInState(t *testing.T) {
	patchRequest := json.RawMessage(`{"name": "failed_update_model",
		"description": "After the update"}`)
	for _, tc := range []struct {
		name      string
		fixtures  []qa.HTTPFixture
		wantError string
	}{
		{
			// The service kept the old description, so the state is ahead of it.
			name: "patch_answered_400",
			fixtures: []qa.HTTPFixture{
				{
					Method:          "PATCH",
					Resource:        "/api/2.0/mlflow/registered-models/update",
					ExpectedRequest: patchRequest,
					Response: `{"error_code": "INVALID_PARAMETER_VALUE",
						"message": "Invalid description for registered model 'failed_update_model'."}`,
					Status: 400,
				},
			},
			wantError: "Invalid description for registered model 'failed_update_model'.",
		},
		{
			// The PATCH succeeded, so the state matches the service although the apply fails.
			name: "read_after_patch_answered_500",
			fixtures: []qa.HTTPFixture{
				{
					Method:          "PATCH",
					Resource:        "/api/2.0/mlflow/registered-models/update",
					ExpectedRequest: patchRequest,
					Response: `{"registered_model": {"name": "failed_update_model",
						"description": "After the update",
						"creation_timestamp": 1700000000000,
						"last_updated_timestamp": 1700000100000,
						"user_id": "tester@example.com"}}`,
				},
				{
					Method:   "GET",
					Resource: "/api/2.0/mlflow/databricks/registered-models/get?name=failed_update_model",
					Response: `{"error_code": "INTERNAL_ERROR",
						"message": "An internal error occurred while reading the registered model."}`,
					Status: 500,
				},
			},
			wantError: "An internal error occurred while reading the registered model.",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			d, err := qa.ResourceFixture{
				Fixtures: tc.fixtures,
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

			assert.EqualError(t, err, tc.wantError)
			assert.Equal(t, "After the update", d.State().Attributes["description"])
			assert.Equal(t, "failed_update_model", d.State().ID)
		})
	}
}
