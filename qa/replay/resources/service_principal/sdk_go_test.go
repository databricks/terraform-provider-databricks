package service_principal_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

// TestServicePrincipalSdkGo runs every service principal journey with reads and writes
// through sdk-go scim/v1, selected by sdk_implementations, and journeys that switch
// between it and the existing client.
func TestServicePrincipalSdkGo(t *testing.T) {
	terraformtest.Run(t, "sdk_go_testdata")
}
