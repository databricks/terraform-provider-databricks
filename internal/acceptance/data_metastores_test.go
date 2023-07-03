package acceptance

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func checkMetastoreIdExists(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		r, ok := s.Modules[0].Resources["data.databricks_metastores.test"]
		require.True(t, ok, "data.databricks_metastores.test has to be there")
		attr := r.Primary.Attributes
		idToCheck := s.Modules[0].Resources["databricks_metastores.test"].Primary.ID
		s.Modules[0].Resources["databricks_metastores.test"].Primary.ID = attr["all_metastores.0.id"]
		assert.Contains(t, )
		return nil
	}
}

func TestUcAccDataSourceMetastore(t *testing.T) {
	accountLevel(t, step{
		Template: `

		resource "databricks_metastore" "temp" {
			name          = "primary-{var.RANDOM}"
			storage_root  = "s3://{env.TEST_METASTORE_BUCKET}/test{var.RANDOM}"
			force_destroy = true
		}

		data "databricks_metastores" "test" {}
		
		output "all_metastores" {
			value = data.databricks_metastores.test
		}`,
		Check: checkMetastoreIdExists(t),
	})
}
