// Package volume_test runs the databricks_volume lifecycle, import, name and refresh
// journeys and the User-Agent and traceparent checks. Call-site, error, retry, routing,
// schema and update journeys are in the sibling volume_* packages, which keep each
// package within the go test timeout.
package volume_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestReplay(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
