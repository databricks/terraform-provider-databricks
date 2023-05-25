package acceptance

import (
	"testing"
)

func TestAccVolumesResourceFullLifecycle(t *testing.T) {
	t.Skip("Not running until we have UC Volumes enabled in our aws test workspace")
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
			storage_location   = "s3://{env.TEST_BUCKET}/sometestpath-for-uc-volumes"
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
			storage_location   = "s3://{env.TEST_BUCKET}/sometestpath-for-uc-volumes"
		}`,
	})
}
