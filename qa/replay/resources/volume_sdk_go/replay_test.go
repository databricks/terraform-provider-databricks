// Package volume_sdk_go_test runs the databricks_volume journeys with reads through
// sdk-go uc/volumes/v1. They are a separate package from ../volume so that each package
// stays within the go test timeout and CI balances them across replay shards.
package volume_sdk_go_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

// TestVolumeSdkGo runs volume journeys with reads through sdk-go uc/volumes/v1, selected
// by sdk_implementations, and a journey that switches between it and the existing client.
func TestVolumeSdkGo(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
