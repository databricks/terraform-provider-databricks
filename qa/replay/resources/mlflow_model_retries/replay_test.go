// Package mlflow_model_retries_test runs the databricks_mlflow_model journeys that pin
// which responses each of its four calls retries, which it sends once, and what apply
// reports when the retry timeout ends. They are a separate package from ../mlflow_model
// so that each package stays within the go test timeout and CI balances them across
// replay shards.
package mlflow_model_retries_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestReplay(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
