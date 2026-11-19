package schema_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

// TestVolumeSdkGo runs the databricks_volume journeys that pin plan-time validation and
// configured-argument handling, with the volume on sdk-go uc/volumes/v1.
func TestVolumeSdkGo(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
