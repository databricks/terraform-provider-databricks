package acceptance

import (
	"fmt"
	"testing"
)

const catalogTemplate = `
	resource "databricks_catalog" "sandbox" {
		name         = "sandbox{var.STICKY_RANDOM}"
		comment      = "this catalog is managed by terraform"
		properties = {
			purpose = "testing"
		}
	}
`

func TestUcAccSchema(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: catalogTemplate + `
		data "databricks_catalogs" "all" {
			depends_on = [databricks_catalog.sandbox]
		}

		resource "databricks_grants" "sandbox" {
			catalog = databricks_catalog.sandbox.name
			grant {
				principal  = "{env.TEST_DATA_SCI_GROUP}"
				privileges = ["USE_CATALOG", "CREATE_SCHEMA"]
			}
			grant {
				principal  = "{env.TEST_DATA_ENG_GROUP}"
				privileges = ["USE_CATALOG"]
			}
		}
		
		resource "databricks_schema" "things" {
			catalog_name = databricks_catalog.sandbox.id
			name         = "things{var.RANDOM}"
			comment      = "this database is managed by terraform"
			properties = {
				kind = "various"
			}
		}

		data "databricks_schemas" "sandbox" {
			catalog_name = databricks_catalog.sandbox.id
			depends_on = [databricks_schema.things]
		}			  

		resource "databricks_grants" "things" {
			schema = databricks_schema.things.id
			grant {
				principal  = "{env.TEST_DATA_ENG_GROUP}"
				privileges = ["USE_SCHEMA"]
			}
		}`,
	})
}

func schemaTemplateWithOwner(comment string, owner string) string {
	return fmt.Sprintf(`
		resource "databricks_schema" "things" {
			catalog_name = databricks_catalog.sandbox.id
			name         = "things{var.STICKY_RANDOM}"
			comment      = "%s"
			properties = {
				kind = "various"
			}
			owner = "%s"
		}`, comment, owner)
}

func TestUcAccSchemaUpdate(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: catalogTemplate + schemaTemplateWithOwner("this database is managed by terraform", "account users"),
	}, step{
		Template: catalogTemplate + schemaTemplateWithOwner("this database is managed by terraform -- updated comment", "account users"),
	}, step{
		Template: catalogTemplate + schemaTemplateWithOwner("this database is managed by terraform -- updated comment 2", "{env.TEST_METASTORE_ADMIN_GROUP_NAME}"),
	})
}
