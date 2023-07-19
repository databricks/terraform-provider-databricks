package acceptance

import (
	"testing"
)

func TestUcAccDataSourceMetastore(t *testing.T) {
	accountLevel(t, step{
		Template: `

		resource "databricks_metastore" "test" {
			name          = "primary-{var.RANDOM}"
			storage_root  = "s3://{env.TEST_BUCKET}/test{var.RANDOM}"
			force_destroy = true
		}

		data "databricks_metastore" "this" {
			id = databricks_metastore.test.id
		}

		output "some_metastore" {
			value = data.databricks_metastore.this.metastore_info
		}`,
	})
}
