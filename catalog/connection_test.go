package catalog_test

import (
	"fmt"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func connectionTemplateWithOwner(host string, owner string) string {
	return fmt.Sprintf(`
	resource "databricks_connection" "this" {
		name = "tf-test-connection-{var.STICKY_RANDOM}"
		connection_type = "MYSQL"
		comment         = "this is a connection to mysql db"
		options         = {
			host     = "%s"
			port     = "3306"
			user     = "user"
			password = "password"
		}
		owner = "%s"
	}
	`, host, owner)
}

func connectionTemplateWithoutOwner() string {
	return `
	resource "databricks_connection" "this" {
		name = "tf-test-connection-{var.STICKY_RANDOM}"
		connection_type = "BIGQUERY"
		comment         = "test"
		options = {
			GoogleServiceAccountKeyJson = <<-EOT
				{
					"type": "service_account",
					"project_id": "PROJECT_ID",
					"private_key_id": "KEY_ID",
					"private_key": "-----BEGIN PRIVATE KEY-----\nPRIVATE_KEY\n-----END PRIVATE KEY-----\n",
					"client_email": "SERVICE_ACCOUNT_EMAIL",
					"client_id": "CLIENT_ID",
					"auth_uri": "https://accounts.google.com/o/oauth2/auth",
					"token_uri": "https://accounts.google.com/o/oauth2/token",
					"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
					"client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/SERVICE_ACCOUNT_EMAIL",
					"universe_domain": "googleapis.com"
				}
			EOT
		}
	}
	`
}
func TestUcAccConnectionsResourceFullLifecycle(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: connectionTemplateWithOwner("test.mysql.database.azure.com", "account users"),
	}, acceptance.Step{
		Template: connectionTemplateWithOwner("test.mysql.database.aws.com", "account users"),
	}, acceptance.Step{
		Template: connectionTemplateWithOwner("test.mysql.database.azure.com", "{env.TEST_METASTORE_ADMIN_GROUP_NAME}"),
	})
}

func TestUcAccConnectionsWithoutOwnerResourceFullLifecycle(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: connectionTemplateWithoutOwner(),
	}, acceptance.Step{
		Template: connectionTemplateWithoutOwner(),
	})
}

// schemaLevelConnectionTemplate creates a connection inside a schema via parent, with a
// configurable host option.
func schemaLevelConnectionTemplate(host string) string {
	return fmt.Sprintf(`
	resource "databricks_catalog" "this" {
		name = "tf_test_sc_cat_{var.STICKY_RANDOM}"
	}
	resource "databricks_schema" "this" {
		catalog_name = databricks_catalog.this.name
		name         = "tf_test_sc_schema"
	}
	resource "databricks_connection" "this" {
		name            = "tf-test-schema-conn-{var.STICKY_RANDOM}"
		connection_type = "HTTP"
		parent          = "schemas/${databricks_catalog.this.name}.${databricks_schema.this.name}"
		comment         = "schema-level connection acceptance test"
		options = {
			host         = "%s"
			port         = "8433"
			base_path    = "/api/"
			bearer_token = "bearer_token"
		}
	}
	`, host)
}

func TestUcAccConnectionsSchemaLevelResourceFullLifecycle(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: schemaLevelConnectionTemplate("https://example.com"),
		// The HTTP bearer backend returns a server-managed auth_scheme option that is not in
		// config, so the post-apply plan is non-empty.
		ExpectNonEmptyPlan: true,
	}, acceptance.Step{
		// Re-applying the identical config must be an in-place update, never a destroy/recreate:
		// this proves the parent/full_name round-trip is stable (no ForceNew drift).
		Template: schemaLevelConnectionTemplate("https://example.com"),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				plancheck.ExpectResourceAction("databricks_connection.this", plancheck.ResourceActionUpdate),
			},
		},
		ExpectNonEmptyPlan: true,
	}, acceptance.Step{
		// In-place update of a mutable option (host). The Check confirms the new value round-trips
		// on read; the plan check confirms an in-place update rather than a replace.
		Template: schemaLevelConnectionTemplate("https://example2.com"),
		Check:    resource.TestCheckResourceAttr("databricks_connection.this", "options.host", "https://example2.com"),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				plancheck.ExpectResourceAction("databricks_connection.this", plancheck.ResourceActionUpdate),
			},
		},
		ExpectNonEmptyPlan: true,
	}, acceptance.Step{
		// Import by metastore_id|full_name. Attribute verification is skipped because options
		// (write-only bearer_token, server-managed auth_scheme) cannot match config on import.
		ResourceName: "databricks_connection.this",
		ImportState:  true,
	})
}
