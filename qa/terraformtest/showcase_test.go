package terraformtest_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestTerraformGroupLifecycle(t *testing.T) {
	terraformtest.Run(t)
}
