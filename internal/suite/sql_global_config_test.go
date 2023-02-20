package acceptance

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccSQLGlobalConfig(t *testing.T) {
	t.Parallel()
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `resource "databricks_sql_global_config" "this" {
				data_access_config = {
					"spark.sql.session.timeZone": "UTC"
				}  
			}`,
		},
	})
}
