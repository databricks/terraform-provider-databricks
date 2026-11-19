// Package volume_lifecycle_test runs the databricks_volume journeys that pin the recorded
// acceptance lifecycles and the create, import, refresh, drift, replacement and delete
// journeys. They are a separate package from ../volume so that each package stays within
// the go test timeout and CI balances them across replay shards.
package volume_lifecycle_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestReplay(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
