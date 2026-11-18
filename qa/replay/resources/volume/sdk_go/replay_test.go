package sdk_go_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

// TestVolumeSdkGo runs volume journeys with reads through sdk-go uc/volumes/v1, selected
// by sdk_implementations. They are a package of their own so that they don't add to the
// existing volume journeys under go test's per-package timeout.
func TestVolumeSdkGo(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
