// Package retries_test runs the sdk-go twins of the ../../volume_retries journeys, in
// their own package so it stays within the go test timeout.
package retries_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestReplay(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
