package names_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

// TestVolumeSdkGo runs the databricks_volume journeys that pin the request paths for
// names with punctuation, spaces, non-ASCII or upper-case characters, with the volume on
// sdk-go uc/volumes/v1.
func TestVolumeSdkGo(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
