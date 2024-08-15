package acceptance

import (
	"fmt"
	"testing"
)

func connectionTemplateWithOwner(host string, owner string) string {
	return fmt.Sprintf(`
	resource "databricks_connection" "this" {
		name = "name-{var.STICKY_RANDOM}"
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
		name = "name-{var.STICKY_RANDOM}"
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
	unityWorkspaceLevel(t, LegacyStep{
		Template: connectionTemplateWithOwner("test.mysql.database.azure.com", "account users"),
	}, LegacyStep{
		Template: connectionTemplateWithOwner("test.mysql.database.aws.com", "account users"),
	}, LegacyStep{
		Template: connectionTemplateWithOwner("test.mysql.database.azure.com", "{env.TEST_METASTORE_ADMIN_GROUP_NAME}"),
	})
}

func TestUcAccConnectionsWithoutOwnerResourceFullLifecycle(t *testing.T) {
	unityWorkspaceLevel(t, LegacyStep{
		Template: connectionTemplateWithoutOwner(),
	}, LegacyStep{
		Template: connectionTemplateWithoutOwner(),
	})
}
