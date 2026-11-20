// Package volume_retries_test runs the databricks_volume journeys that pin retries of
// transient errors and connection resets, retry exhaustion and HTTP timeouts. They are a
// separate package from ../volume so that each package stays within the go test timeout
// and CI balances them across replay shards.
package volume_retries_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestReplay(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
