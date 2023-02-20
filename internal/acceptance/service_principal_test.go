package acceptance

import (
	"testing"
)

const azureSpn = `resource "databricks_service_principal" "this" {
	application_id = "00000000-1234-5678-0000-000000000001"
	display_name = "SPN {var.RANDOM}"
	force = true
}`

const awsSpn = `resource "databricks_service_principal" "this" {
	display_name = "SPN {var.RANDOM}"
}`

func TestMwsAccServicePrincipalResourceOnAzure(t *testing.T) {
	GetEnvOrSkipTest(t, "ARM_CLIENT_ID")
	accountLevel(t, step{
		Template: azureSpn,
	})
}

func TestAccServicePrincipalResourceOnAzure(t *testing.T) {
	GetEnvOrSkipTest(t, "ARM_CLIENT_ID")
	workspaceLevel(t, step{
		Template: azureSpn,
	})
}

func TestMwsAccServicePrincipalResourceOnAws(t *testing.T) {
	GetEnvOrSkipTest(t, "TEST_ROOT_BUCKET")
	accountLevel(t, step{
		Template: awsSpn,
	})
}

func TestAccServicePrincipalResourceOnAws(t *testing.T) {
	GetEnvOrSkipTest(t, "TEST_EC2_INSTANCE_PROFILE")
	workspaceLevel(t, step{
		Template: awsSpn,
	})
}
