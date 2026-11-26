// Package mlflow_model_error_bodies_test runs the databricks_mlflow_model journeys that
// pin which error bodies drop the model on refresh and which count as deleted on DELETE.
// The classic SDK's error parsers decide that from the body, not the resource from the
// status. They are a separate package from ../mlflow_model so that each package stays
// within the go test timeout.
package mlflow_model_error_bodies_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

// TestReplay runs the journeys on the classic SDK path of databricks_mlflow_model. It
// clears the opt-in, so that the environment of the test process can't switch the path.
func TestReplay(t *testing.T) {
	t.Setenv("USE_SDK_GO_RESOURCES", "")
	terraformtest.Run(t, "testdata")
}

// TestReplaySdkGo runs the same journeys on the modular SDK path of the resource.
func TestReplaySdkGo(t *testing.T) {
	t.Setenv("USE_SDK_GO_RESOURCES", "databricks_mlflow_model")
	terraformtest.Run(t, "testdata")
}
