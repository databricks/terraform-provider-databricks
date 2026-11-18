package schema_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

// TestReplay runs the volume journeys that pin plan-time validation and the handling of
// configured arguments.
func TestReplay(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
