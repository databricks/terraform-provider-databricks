package acceptance

import (
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"

	"testing"
)

func TestAccDataSourceQueryableJob(t *testing.T) {
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			data "databricks_queryable_job" "this" {
			}`,
		},
	})
}
