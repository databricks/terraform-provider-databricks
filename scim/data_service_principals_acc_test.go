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

// Regression coverage for the same class of issue as #5664: the data source
// is non-dual yet has provider_config in its schema, and at account level the
// post-Read hook tries to resolve a workspace_id that doesn't exist.
const accountSpnsByDisplayName = `
resource "databricks_service_principal" "this" {
	display_name = "SPN {var.RANDOM}"
	api          = "account"
}

data "databricks_service_principals" "this" {
	display_name_contains = databricks_service_principal.this.display_name
	api                   = "account"
	depends_on            = [databricks_service_principal.this]
}`

func TestMwsAccDataSourceSPNsByDisplayNameOnAccount(t *testing.T) {
	acceptance.AccountLevel(t, acceptance.Step{
		Template: accountSpnsByDisplayName,
	})
}

const accountSpnsByDisplayNameNoApi = `
resource "databricks_service_principal" "this" {
	display_name = "SPN {var.RANDOM}"
	api          = "account"
}

data "databricks_service_principals" "this" {
	display_name_contains = databricks_service_principal.this.display_name
	depends_on            = [databricks_service_principal.this]
}`

func TestMwsAccDataSourceSPNsByDisplayNameOnAccountNoApi(t *testing.T) {
	acceptance.AccountLevel(t, acceptance.Step{
		Template: accountSpnsByDisplayNameNoApi,
	})
}
