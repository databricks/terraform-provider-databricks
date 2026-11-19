package quality_monitor_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestQualityMonitor(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
