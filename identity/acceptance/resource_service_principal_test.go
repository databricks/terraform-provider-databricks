package acceptance

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/internal/acceptance"
)

func TestAzureAccServicePrincipalResource(t *testing.T) {
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `resource "databricks_service_principal" "this" {
				application_id = "00000000-1234-5678-0000-000000000001"
				display_name = "SPN {var.RANDOM}"
			}`,
		},
	})
}

func TestAwsAccServicePrincipalResource(t *testing.T) {
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `resource "databricks_service_principal" "this" {
				display_name = "SPN {var.RANDOM}"
			}`,
		},
	})
}
