// Package lifecycle_test runs the sdk-go twins of the ../../volume_lifecycle journeys, in
// their own package so it stays within the go test timeout.
package lifecycle_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestReplay(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
