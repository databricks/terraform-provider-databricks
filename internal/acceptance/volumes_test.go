package acceptance

import (
	"testing"
)

func TestAccVolumesResourceFullLifecycle(t *testing.T) {
	workspaceLevel(t, step{
		Template: `
		resource "databricks_volumes" "this" {
			name = "tf-{var.RANDOM}"
			owner = "abc"
			comment = "comment abc"
			catalog_name = "catalog abc"
			volume_type = "volume abc"
			schema_name = "schema abc" 
		}`,
	}, step{
		Template: `
		resource "databricks_volumes" "this" {
			name = "tf-{var.RANDOM}"
			owner = "def"
			comment = "comment def"
			catalog_name = "catalog def"
			volume_type = "volume def"
			schema_name = "schema def" 
		}`,
	})
}
