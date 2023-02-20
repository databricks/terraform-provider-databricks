package acceptance

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestUcAccMetastoreAssignment(t *testing.T) {
	qa.RequireCloudEnv(t, "ucws")
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `resource "databricks_metastore_assignment" "this" {
				metastore_id = "{env.TEST_METASTORE_ID}"
				workspace_id = {env.TEST_WORKSPACE_ID}
			}`,
		},
	})
}
