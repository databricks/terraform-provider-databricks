// Package volume_errors_test runs the databricks_volume journeys that pin failed calls:
// error bodies, statuses that are not retried, undecodable responses and private-link
// login redirects. They are a separate package from ../volume so that each package stays
// within the go test timeout and CI balances them across replay shards.
package volume_errors_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestReplay(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
