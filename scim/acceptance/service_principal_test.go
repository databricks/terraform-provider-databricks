package acceptance

import (
	"os"
	"strings"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestMwsAccServicePrincipalResourceOnAzure(t *testing.T) {
	TestAccServicePrincipalResourceOnAzure(t)
}
func TestAccServicePrincipalResourceOnAzure(t *testing.T) {
	if cloud, ok := os.LookupEnv("CLOUD_ENV"); !ok || !strings.Contains(cloud, "azure") {
		t.Skip("Test is only for CLOUD_ENV=azure")
	}
	t.Parallel()
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `resource "databricks_service_principal" "this" {
				application_id = "00000000-1234-5678-0000-000000000001"
				display_name = "SPN {var.RANDOM}"
				force = true
			}`,
		},
	})
}
func TestMwsAccServicePrincipalResourceOnAws(t *testing.T) {
	TestAccServicePrincipalResourceOnAws(t)
}
func TestAccServicePrincipalResourceOnAws(t *testing.T) {
	if cloud, ok := os.LookupEnv("CLOUD_ENV"); !ok || cloud != "aws" {
		t.Skip("Test is only for CLOUD_ENV=aws")
	}
	t.Parallel()
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `resource "databricks_service_principal" "this" {
				display_name = "SPN {var.RANDOM}"
			}`,
		},
	})
}
