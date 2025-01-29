package acceptance

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/require"
)

func checkDataSourceSchema(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		_, ok := s.Modules[0].Resources["data.databricks_schema.this"]
		require.True(t, ok, "data.databricks_schema.this has to be there")
		return nil
	}
}

func TestUcAccDataSourceSchema(t *testing.T) {
	UnityWorkspaceLevel(t, Step{
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

		data "databricks_schema" "this" {
			name = databricks_schema.things.id
		}
		`,
		Check: checkDataSourceSchema(t),
	})
}
