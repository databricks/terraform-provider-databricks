package acceptance

import (
	"testing"
)

func TestAccSQLGlobalConfig(t *testing.T) {
	workspaceLevel(t, step{
		Template: `resource "databricks_sql_global_config" "this" {
			data_access_config = {
				"spark.sql.session.timeZone": "UTC"
			}  
		}`,
	})
}
