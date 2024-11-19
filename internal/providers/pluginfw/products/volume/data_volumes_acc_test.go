package volume_test

import (
	"strconv"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func checkDataSourceVolumesPopulated(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		_, ok := s.Modules[0].Resources["data.databricks_volumes.this"]
		require.True(t, ok, "data.databricks_volumes.this has to be there")
		num_volumes, _ := strconv.Atoi(s.Modules[0].Outputs["volumes"].Value.(string))
		assert.GreaterOrEqual(t, num_volumes, 1)
		return nil
	}
}

func TestUcAccDataSourceVolumes(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
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
		data "databricks_volumes" "this" {
			catalog_name = databricks_catalog.sandbox.name
			schema_name = databricks_schema.things.name
			depends_on = [ databricks_volume.this ] 
		}
		output "volumes" {
			value = length(data.databricks_volumes.this.ids)
		}
		`,
		Check: checkDataSourceVolumesPopulated(t),
	})
}
