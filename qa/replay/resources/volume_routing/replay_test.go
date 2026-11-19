// Package volume_routing_test runs the databricks_volume journeys that pin which
// workspace each call goes to and how provider_config is resolved. They are a separate
// package from ../volume so that each package stays within the go test timeout and CI
// balances them across replay shards.
package volume_routing_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestReplay(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
