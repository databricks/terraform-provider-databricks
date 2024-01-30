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

func connectionTemplateWithoutOwner(host string) string {
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
	}
	`, host)
}
func TestUcAccConnectionsResourceFullLifecycle(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: connectionTemplateWithOwner("test.mysql.database.azure.com", "account users"),
	}, step{
		Template: connectionTemplateWithOwner("test.mysql.database.aws.com", "account users"),
	}, step{
		Template: connectionTemplateWithOwner("test.mysql.database.azure.com", "{env.TEST_METASTORE_ADMIN_GROUP_NAME}"),
	})
}

func TestUcAccConnectionsWithoutOwnerResourceFullLifecycle(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: connectionTemplateWithoutOwner("test.mysql.database.azure.com"),
	}, step{
		Template: connectionTemplateWithoutOwner("test.mysql.database.aws.com"),
	})
}
