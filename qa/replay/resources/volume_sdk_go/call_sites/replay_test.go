package call_sites_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

// TestVolumeSdkGo runs the databricks_volume journeys that pin errors and retries at each
// write call site and at the read after an update, with the volume on sdk-go
// uc/volumes/v1.
func TestVolumeSdkGo(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
