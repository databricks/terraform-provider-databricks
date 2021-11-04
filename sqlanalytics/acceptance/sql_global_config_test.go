package acceptance

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/internal/acceptance"
)

func TestPreviewAccSQLGlobalConfig(t *testing.T) {
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
