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

func TestReplay(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
