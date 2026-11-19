// Package volume_updates_test runs the databricks_volume journeys that pin the PATCH
// requests of updates, including renames, owner changes and rollbacks. They are a
// separate package from ../volume so that each package stays within the go test timeout
// and CI balances them across replay shards.
package volume_updates_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestReplay(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
