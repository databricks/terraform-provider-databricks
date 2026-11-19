package connection_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

// TestReplay runs the volume journeys that pin connection resets, missing responses,
// broken bodies and request visitor failures. They are a package of their own so the
// volume journeys stay within go test's per-package timeout.
func TestReplay(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
