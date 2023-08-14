package acceptance

import (
	"testing"
)

func TestUcAccConnectionsResourceFullLifecycle(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_connection" "this" {
			name = "name-{var.STICKY_RANDOM}"
			connection_type = "MYSQL"
			comment         = "this is a connection to mysql db"
			options         = {
				host     = "test.mysql.database.azure.com"
				port     = "3306"
				user     = "user"
				password = "password"
			}
		}`,
	})
}
