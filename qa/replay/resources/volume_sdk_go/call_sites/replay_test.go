// Package call_sites_test runs the sdk-go twins of the ../../volume_call_sites journeys,
// in their own package so it stays within the go test timeout.
package call_sites_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestReplay(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
