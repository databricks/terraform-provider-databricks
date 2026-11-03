package dashboard_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestDashboard(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
