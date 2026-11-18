package updates_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

// TestReplay runs the volume journeys that pin the requests updates send, including owner
// changes, renames and rollbacks.
func TestReplay(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
