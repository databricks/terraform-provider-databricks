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

// Regression test for https://github.com/databricks/terraform-provider-databricks/issues/5664.
// The data source must work against an account-level provider with no workspace_id
// configured anywhere. Before adding the `api` field, the post-Read provider_config
// hook tried to resolve a workspace_id and failed with strconv.ParseInt: parsing "".
const accountSpnByDisplayName = `
resource "databricks_service_principal" "this" {
	display_name = "SPN {var.RANDOM}"
	api          = "account"
}

data "databricks_service_principal" "this" {
	display_name = databricks_service_principal.this.display_name
	api          = "account"
	depends_on   = [databricks_service_principal.this]
}`

func TestMwsAccDataSourceSPNByDisplayNameOnAccount(t *testing.T) {
	acceptance.AccountLevel(t, acceptance.Step{
		Template: accountSpnByDisplayName,
	})
}

// Same as above but does not set `api` on the data source — exercises the
// host-based account-level inference path so the regression in 1.114.0 cannot
// reappear under the no-explicit-api configuration the user originally hit.
const accountSpnByDisplayNameNoApi = `
resource "databricks_service_principal" "this" {
	display_name = "SPN {var.RANDOM}"
	api          = "account"
}

data "databricks_service_principal" "this" {
	display_name = databricks_service_principal.this.display_name
	depends_on   = [databricks_service_principal.this]
}`

func TestMwsAccDataSourceSPNByDisplayNameOnAccountNoApi(t *testing.T) {
	acceptance.AccountLevel(t, acceptance.Step{
		Template: accountSpnByDisplayNameNoApi,
	})
}
