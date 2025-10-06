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

const volumesTemplate = `
	resource "databricks_catalog" "sandbox" {
		name         = "sandbox{var.STICKY_RANDOM}"
		comment      = "this catalog is managed by terraform"
		properties = {
			purpose = "testing"
		}
	}

	resource "databricks_schema" "things" {
		catalog_name = databricks_catalog.sandbox.id
		name         = "things{var.STICKY_RANDOM}"
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
	output "volumes" {
		value = length(data.databricks_volumes.this.ids)
	}
`

func TestUcAccDataSourceVolumes(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: volumesTemplate + `
		data "databricks_volumes" "this" {
			catalog_name = databricks_catalog.sandbox.name
			schema_name = databricks_schema.things.name
			depends_on = [ databricks_volume.this ]
		}
		`,
		Check: checkDataSourceVolumesPopulated(t),
	}, acceptance.Step{
		Template: volumesTemplate + `
		data "databricks_volumes" "this" {
			catalog_name = databricks_catalog.sandbox.name
			schema_name = databricks_schema.things.name
			depends_on = [ databricks_volume.this ]
			provider_config = {
				workspace_id = "{env.THIS_WORKSPACE_ID}"
			}
		}
		`,
		Check: checkDataSourceVolumesPopulated(t),
	})
}
