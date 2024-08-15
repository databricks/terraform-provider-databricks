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
	unityWorkspaceLevel(t, LegacyStep{
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
			enable_predictive_optimization = "DISABLE"
		}

		resource "databricks_schema" "stuff" {
			catalog_name = databricks_catalog.sandbox.id
			name         = "stuff{var.RANDOM}"
			comment      = "this database is managed by terraform"
			properties = {
				kind = "various"
			}
		}

		data "databricks_schemas" "sandbox" {
			catalog_name = databricks_catalog.sandbox.id
			depends_on = [databricks_schema.things]
		}

		# This overwrites all grants on schema "things"
		resource "databricks_grants" "things" {
			schema = databricks_schema.things.id
			grant {
				principal  = "{env.TEST_DATA_ENG_GROUP}"
				privileges = ["USE_SCHEMA"]
			}
		}

		resource "databricks_grant" "stuff_test_eng_group" {
			schema = databricks_schema.stuff.id
			principal  = "{env.TEST_DATA_ENG_GROUP}"
			privileges = ["USE_SCHEMA"]
		}`,
	})
}

func schemaTemplateWithOwner(t *testing.T, comment string, owner string) string {
	return fmt.Sprintf(`
		resource "databricks_schema" "things" {
			catalog_name = databricks_catalog.sandbox.id
			name         = "things{var.STICKY_RANDOM}"
			comment      = "%s"
			properties = {
				kind = "various"
			}
			%s
			owner = "%s"
		}`, comment, getPredictiveOptimizationSetting(t, true), owner)
}

func getPredictiveOptimizationSetting(t *testing.T, enabled bool) string {
	if isGcp(t) {
		return ""
	}
	value := "ENABLE"
	if !enabled {
		value = "DISABLE"
	}
	return fmt.Sprintf(`enable_predictive_optimization = "%s"`, value)
}

func TestUcAccSchemaUpdate(t *testing.T) {
	loadUcwsEnv(t)
	unityWorkspaceLevel(t, LegacyStep{
		Template: catalogTemplate + schemaTemplateWithOwner(t, "this database is managed by terraform", "account users"),
	}, LegacyStep{
		Template: catalogTemplate + schemaTemplateWithOwner(t, "this database is managed by terraform -- updated comment", "account users"),
	}, LegacyStep{
		Template: catalogTemplate + schemaTemplateWithOwner(t, "this database is managed by terraform -- updated comment", "{env.TEST_DATA_ENG_GROUP}"),
	}, LegacyStep{
		Template: catalogTemplate + schemaTemplateWithOwner(t, "this database is managed by terraform -- updated comment 2", "{env.TEST_METASTORE_ADMIN_GROUP_NAME}"),
	})
}
