package global_init_script_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestGlobalInitScript(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
