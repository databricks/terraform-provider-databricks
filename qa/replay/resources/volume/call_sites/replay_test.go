package call_sites_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

// TestReplay runs the volume journeys that pin errors and retries at each write call
// site and at the read after an update.
func TestReplay(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
