package names_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

// TestReplay runs the volume journeys that pin the request paths for names with
// punctuation, spaces, non-ASCII or upper-case characters.
func TestReplay(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
