// Package mlflow_model_lifecycle_test pins the create, replacement and delete requests
// of databricks_mlflow_model. It is a separate package from ../mlflow_model so that each
// package stays within the go test timeout and CI balances them across replay shards.
package mlflow_model_lifecycle_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestReplay(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
