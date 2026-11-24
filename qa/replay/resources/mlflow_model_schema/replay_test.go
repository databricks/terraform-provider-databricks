// Package mlflow_model_schema_test runs the databricks_mlflow_model journeys that pin
// the importer and how configured arguments are validated and sent. They are a separate
// package from ../mlflow_model so that each package stays within the go test timeout and
// CI balances them across replay shards.
package mlflow_model_schema_test

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
