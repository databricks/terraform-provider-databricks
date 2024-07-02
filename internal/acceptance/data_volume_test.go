package acceptance

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func checkDataSourceVolume(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		r, ok := s.Modules[0].Resources["data.databricks_volume.this"]
		require.True(t, ok, "data.databricks_volume.this has to be there")
		assert.Equal(t, s.Modules[0].Resources["databricks_volume.this"].Primary.ID, r.Primary.Attributes["full_name"])
		return nil
	}
}
func TestUcAccDataSourceVolume(t *testing.T) {
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
			name         = "things{var.RANDOM}"
			comment      = "this database is managed by terraform"
			properties = {
				kind = "various"
			}
		}

		resource "databricks_volume" "this" {
			name         = "volume_data_source_test"
			catalog_name = databricks_catalog.sandbox.name
			schema_name  = databricks_schema.things.name
			volume_type  = "MANAGED"      
		}

		data "databricks_volume" "this" {
			name = databricks_volume.this.id
			depends_on = [ databricks_volume.this ] 
		}
		`,
		Check: checkDataSourceVolume(t),
	})
}
