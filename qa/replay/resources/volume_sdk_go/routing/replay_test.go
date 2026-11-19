package routing_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

// TestVolumeSdkGo runs the databricks_volume journeys that pin which workspace each call
// goes to and how provider_config is populated, with the volume on sdk-go uc/volumes/v1.
func TestVolumeSdkGo(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
