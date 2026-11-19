// Package volume_test runs the databricks_volume name journeys and the User-Agent and
// traceparent checks. The other journeys are in the sibling volume_* packages, one per
// theme, which keep each package within the go test timeout.
package volume_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestReplay(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
