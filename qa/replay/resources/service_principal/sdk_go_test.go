package service_principal_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

// TestServicePrincipalSdkGo runs service principal journeys with reads through sdk-go
// scim/v1, selected by sdk_implementations.
func TestServicePrincipalSdkGo(t *testing.T) {
	terraformtest.Run(t, "sdk_go_testdata")
}
