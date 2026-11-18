package volume_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

// TestVolumeSdkGo runs volume journeys with reads through sdk-go uc/volumes/v1, selected
// by sdk_implementations.
func TestVolumeSdkGo(t *testing.T) {
	terraformtest.Run(t, "sdk_go_testdata")
}
