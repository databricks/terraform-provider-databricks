// Package mlflow_model_errors_test runs the databricks_mlflow_model journeys that pin
// failed API calls: the diagnostic, the state and the next plan after an error on each
// call, and which errors count as a missing model. They are a separate package from
// ../mlflow_model so that each package stays within the go test timeout.
package mlflow_model_errors_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestReplay(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
