package decoding_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

// TestReplay runs the volume journeys that pin how 200 response bodies are decoded. They
// are a package of their own so the volume journeys stay within go test's per-package
// timeout.
func TestReplay(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
