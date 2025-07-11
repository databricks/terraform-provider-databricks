package scim_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

const spnBySCIMID = `
resource "databricks_service_principal" "this" {
	display_name = "SPN {var.RANDOM}"
}

data "databricks_service_principal" "this" {
	scim_id = databricks_service_principal.this.id
	depends_on = [databricks_service_principal.this]
}`

const azureSpnBySCIMID = `
resource "databricks_service_principal" "this" {
	application_id = "{var.RANDOM_UUID}"
	display_name = "SPN {var.RANDOM}"
	force = true
}

data "databricks_service_principal" "this" {
	scim_id = databricks_service_principal.this.id 
	depends_on = [databricks_service_principal.this]
}`

func TestAccDataSourceSPNOnAWSBySCIMID(t *testing.T) {
	acceptance.GetEnvOrSkipTest(t, "TEST_EC2_INSTANCE_PROFILE")
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: spnBySCIMID,
	})
}

func TestAccDataSourceSPNOnGCPBySCIMID(t *testing.T) {
	acceptance.GetEnvOrSkipTest(t, "GOOGLE_CREDENTIALS")
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: spnBySCIMID,
	})
}

func TestAccDataSourceSPNOnAzureBySCIMID(t *testing.T) {
	acceptance.GetEnvOrSkipTest(t, "ARM_CLIENT_ID")
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: azureSpnBySCIMID,
	})
}
