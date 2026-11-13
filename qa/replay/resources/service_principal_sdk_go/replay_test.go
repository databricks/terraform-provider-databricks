// Package service_principal_sdk_go_test runs the databricks_service_principal journeys
// with the resource on sdk-go scim/v1. They are a separate package from
// ../service_principal so that each package stays within the go test timeout and CI
// balances them across replay shards.
package service_principal_sdk_go_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

// TestServicePrincipalSdkGo runs every service principal journey with reads and writes
// through sdk-go scim/v1, selected by sdk_implementations, and journeys that switch
// between it and the existing client.
func TestServicePrincipalSdkGo(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
