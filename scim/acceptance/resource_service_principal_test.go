package acceptance

import (
	"os"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/internal/acceptance"
)

func TestAccServicePrincipalResourceOnAzure(t *testing.T) {
	if cloud, ok := os.LookupEnv("CLOUD_ENV"); !ok || cloud != "azure" {
		t.Skip("Test is only for CLOUD_ENV=azure")
	}
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `resource "databricks_service_principal" "this" {
				application_id = "00000000-1234-5678-0000-000000000001"
				display_name = "SPN {var.RANDOM}"
			}`,
		},
	})
}

func TestAccServicePrincipalResourceOnAws(t *testing.T) {
	if cloud, ok := os.LookupEnv("CLOUD_ENV"); !ok || cloud != "AWS" {
		t.Skip("Test is only for CLOUD_ENV=AWS")
	}
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `resource "databricks_service_principal" "this" {
				display_name = "SPN {var.RANDOM}"
			}`,
		},
	})
}
