package acceptance

import (
	"testing"
)

func TestAccVolumesResourceFullLifecycle(t *testing.T) {
	workspaceLevel(t, step{
		Template: `
		resource "databricks_schema" "this" {
			name 		 = "testschema"
			catalog_name = "main"
		}

		resource "databricks_volumes" "this" {
			name = "name-abc"
			owner = "owner-abc"
			comment = "comment-abc"
			catalog_name = "main"
			schema_name = databricks_schema.this.name 
			volume_type = "EXTERNAL"
			storage_location   = "s3://{env.TEST_BUCKET}/sometestingpath"
		}`,
	}, step{
		Template: `
		resource "databricks_schema" "this" {
			name 		 = "testschema"
			catalog_name = "main"
		}

		resource "databricks_volumes" "this" {
			name = "name-def"
			owner = "owner-def"
			comment = "comment-def"
			catalog_name = "main"
			schema_name = databricks_schema.this.name 
			volume_type = "EXTERNAL"
			storage_location   = "s3://{env.TEST_BUCKET}/sometestingpath"
		}`,
	})
}
