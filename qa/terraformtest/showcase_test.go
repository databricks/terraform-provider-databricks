package terraformtest_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestTerraform(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
