package apiselection_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestAPISelection(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
