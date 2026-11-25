// Response bodies in this file are literal wire JSON, so that the tests pin the bytes on
// the wire whichever client the resource uses. They are synthetic: inferred from the
// MLflow Model Registry API, not recorded.

package mlflow

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestMlflowModelCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceMlflowModel())
}

// TestMlflowModelErrorWithoutErrorFields pins the error of a 400 whose JSON body has none
// of error_code, message, error or detail, such as a body from a gateway or proxy in
// front of the API. The API itself sends error_code and message.
func TestMlflowModelErrorWithoutErrorFields(t *testing.T) {
	testCases := []struct {
		name    string
		fixture qa.ResourceFixture
	}{
		{
			name: "create",
			fixture: qa.ResourceFixture{
				Fixtures: []qa.HTTPFixture{{
					Method:   "POST",
					Resource: "/api/2.0/mlflow/registered-models/create",
					Status:   400,
					Response: `{"registered_model": {"name": "xyz"}}`,
				}},
				Create: true,
				HCL:    `name = "xyz"`,
			},
		},
		{
			name: "read",
			fixture: qa.ResourceFixture{
				Fixtures: []qa.HTTPFixture{{
					Method:   "GET",
					Resource: "/api/2.0/mlflow/databricks/registered-models/get?name=xyz",
					Status:   400,
					Response: `{"registered_model_databricks": {"name": "xyz"}}`,
				}},
				Read: true,
				ID:   "xyz",
			},
		},
		{
			name: "update",
			fixture: qa.ResourceFixture{
				Fixtures: []qa.HTTPFixture{{
					Method:   "PATCH",
					Resource: "/api/2.0/mlflow/registered-models/update",
					Status:   400,
					Response: `{"registered_model": {"name": "xyz", "description": "updated"}}`,
				}},
				Update:      true,
				RequiresNew: true,
				ID:          "xyz",
				State:       map[string]any{"name": "xyz"},
				HCL: `
				name        = "xyz"
				description = "updated"
				`,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.fixture.Resource = ResourceMlflowModel()
			_, err := tc.fixture.Apply(t)
			// The SDK takes the empty message of any JSON object as the error text.
			assert.EqualError(t, err, "")
		})
	}
}
