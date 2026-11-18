package errors_and_retries_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

// TestReplay runs the volume journeys that pin error diagnostics, response decoding and
// which failures the client retries.
func TestReplay(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
