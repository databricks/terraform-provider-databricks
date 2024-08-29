package acceptance

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func checkSchemasDataSourcePopulated(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		r, ok := s.Modules[0].Resources["data.databricks_schemas.things"]
		require.True(t, ok, "data.databricks_schemas.things has to be there")
		attr := r.Primary.Attributes

		assert.Equal(t, s.Modules[0].Resources["databricks_schema.things"].Primary.ID, attr["ids.1"])
		assert.Equal(t, s.Modules[0].Resources["databricks_schema.things_2"].Primary.ID, attr["ids.2"])
		return nil
	}
}
func TestUcAccDataSourceSchemas(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_catalog" "sandbox" {
			name         = "sandbox{var.RANDOM}"
			comment      = "this catalog is managed by terraform"
			properties = {
				purpose = "testing"
			}
		}
		
		resource "databricks_schema" "things" {
			catalog_name = databricks_catalog.sandbox.id
			name         = "things"
			comment      = "this database is managed by terraform"
			properties = {
				kind = "various"
			}
		}

		resource "databricks_schema" "things_2" {
			catalog_name = databricks_catalog.sandbox.id
			name         = "things2"
			comment      = "this database is managed by terraform"
			properties = {
				kind = "various"
			}
		}

		data "databricks_schemas" "things" {
			catalog_name = databricks_catalog.sandbox.id
			depends_on = [
				databricks_schema.things,
				databricks_schema.things_2
			]
		}`,
		Check: checkSchemasDataSourcePopulated(t),
	})
}
