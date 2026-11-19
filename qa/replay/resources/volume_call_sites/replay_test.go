// Package volume_call_sites_test runs the databricks_volume journeys that pin the errors
// and retries of each call in create, update and delete. They are a separate package from
// ../volume so that each package stays within the go test timeout and CI balances them
// across replay shards.
package volume_call_sites_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestReplay(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
