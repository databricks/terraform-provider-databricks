// Package routing_test runs the sdk-go twins of the ../../volume_routing journeys. They
// are a separate package so that each package stays within the go test timeout.
package routing_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestReplay(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
