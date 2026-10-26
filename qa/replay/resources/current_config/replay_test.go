package current_config_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestCurrentConfig(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
