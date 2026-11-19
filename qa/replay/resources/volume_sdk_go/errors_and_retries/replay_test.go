package errors_and_retries_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

// TestVolumeSdkGo runs the databricks_volume journeys that pin error bodies, missing
// volumes, response decoding errors, the private link login redirect and retries, with
// the volume on sdk-go uc/volumes/v1.
func TestVolumeSdkGo(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
