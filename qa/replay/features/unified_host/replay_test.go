package unified_host_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestUnifiedHost(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
