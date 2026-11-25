// Package mlflow_model_updates_test runs the databricks_mlflow_model journeys that pin
// the PATCH requests of updates, drift, and how a refresh turns the GET response into
// state. They are a separate package from ../mlflow_model so that each package stays
// within the go test timeout and CI balances them across replay shards.
package mlflow_model_updates_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestReplay(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
