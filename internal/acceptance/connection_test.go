package acceptance

import (
	"testing"
)

func TestUcAccConnectionsResourceFullLifecycle(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_volume" "this" {
			name = "name-{var.STICKY_RANDOM}"
			comment = "comment-{var.STICKY_RANDOM}"
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
