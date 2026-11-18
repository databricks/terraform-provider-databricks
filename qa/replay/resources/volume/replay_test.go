package volume_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

// TestReplay runs the volume lifecycle journeys. Other themes are packages of their own
// under this directory, so each stays well within go test's per-package timeout.
func TestReplay(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
