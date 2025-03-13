package catalog_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/require"
)

func TestUcAccDataSourceCatalog(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: `
		resource "databricks_catalog" "sandbox" {
			name         = "sandbox{var.RANDOM}"
			comment      = "this catalog is managed by terraform"
			properties = {
				purpose = "testing"
			}
		}
		
		data "databricks_catalog" "this" {
			name = databricks_catalog.sandbox.name
			depends_on = [
				databricks_catalog.sandbox,
			]
		}
		`,
		Check: func(s *terraform.State) error {
			_, ok := s.Modules[0].Resources["data.databricks_catalog.this"]
			require.True(t, ok, "data.databricks_catalog.this has to be there")
			return nil
		},
	})
}
