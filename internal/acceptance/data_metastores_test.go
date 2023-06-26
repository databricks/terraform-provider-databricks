package acceptance

import (
	"testing"
)

func TestUcAccDataSourceMetastore(t *testing.T) {
	accountLevel(t, step{
		Template: `
		data "databricks_metastores" "this" {}
		
		output "all_metastores" {
			value = data.databricks_metastores.this
		}`,
	})
}
