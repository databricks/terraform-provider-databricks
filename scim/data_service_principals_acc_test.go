package scim_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

const spns = `
resource "databricks_service_principal" "this" {
	display_name = "SPN {var.RANDOM}"
}

data databricks_service_principals "this" {
	display_name_contains = ""
	depends_on = [databricks_service_principal.this]
}`

const azureSpns = `
resource "databricks_service_principal" "this" {
	application_id = "{var.RANDOM_UUID}"
	display_name = "SPN {var.RANDOM}"
	force = true
}

data databricks_service_principals "this" {
	display_name_contains = ""
	depends_on = [databricks_service_principal.this]
}`

func TestAccDataSourceSPNsOnAWS(t *testing.T) {
	acceptance.GetEnvOrSkipTest(t, "TEST_EC2_INSTANCE_PROFILE")
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: spns,
	})
}

func TestAccDataSourceSPNsOnGCP(t *testing.T) {
	acceptance.GetEnvOrSkipTest(t, "GOOGLE_CREDENTIALS")
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: spns,
	})
}

func TestAccDataSourceSPNsOnAzure(t *testing.T) {
	acceptance.GetEnvOrSkipTest(t, "ARM_CLIENT_ID")
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: azureSpns,
	})
}
