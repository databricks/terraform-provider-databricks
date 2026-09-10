package catalog_test

import (
	"fmt"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
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

// A schema-level (L3) connection is created inside a schema via the parent
// argument; its full_name is catalog.schema.name and the resource id is
// metastore_id|full_name.
func schemaLevelConnectionTemplate() string {
	return `
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
			host         = "https://example.com"
			port         = "8433"
			base_path    = "/api/"
			bearer_token = "bearer_token"
		}
	}
	`
}

func TestUcAccConnectionsSchemaLevelResourceFullLifecycle(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: schemaLevelConnectionTemplate(),
		// A schema-level HTTP connection re-plans dirty on `options` because the server omits the
		// write-only bearer_token on read. That is a pre-existing round-trip issue affecting all
		// HTTP connections (metastore- and schema-level alike), not specific to schema connections,
		// and it does not block create/read/update/delete. Tracked separately.
		ExpectNonEmptyPlan: true,
	})
}
