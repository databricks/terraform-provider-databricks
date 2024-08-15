package acceptance

import (
	"testing"
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
	GetEnvOrSkipTest(t, "TEST_EC2_INSTANCE_PROFILE")
	workspaceLevel(t, LegacyStep{
		Template: spns,
	})
}

func TestAccDataSourceSPNsOnGCP(t *testing.T) {
	GetEnvOrSkipTest(t, "GOOGLE_CREDENTIALS")
	workspaceLevel(t, LegacyStep{
		Template: spns,
	})
}

func TestAccDataSourceSPNsOnAzure(t *testing.T) {
	GetEnvOrSkipTest(t, "ARM_CLIENT_ID")
	workspaceLevel(t, LegacyStep{
		Template: azureSpns,
	})
}
