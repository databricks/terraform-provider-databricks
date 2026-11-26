// Package mlflow_model_schema_test runs the databricks_mlflow_model journeys that pin
// the importer and how configured arguments are validated and sent. They are a separate
// package from ../mlflow_model so that each package stays within the go test timeout and
// CI balances them across replay shards.
package mlflow_model_schema_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestReplay(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
