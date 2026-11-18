package routing_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

// TestReplay runs the volume journeys that pin which workspace each call goes to and how
// provider_config is populated.
func TestReplay(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
