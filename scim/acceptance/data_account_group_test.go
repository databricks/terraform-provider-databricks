package acceptance

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestUcAccGroupDataWorkspace(t *testing.T) {
	qa.RequireCloudEnv(t, "ucws")
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `data "databricks_account_group" "data_eng_group" {
				display_name = "{env.TEST_DATA_ENG_GROUP}"
			}
			output "data_eng_group_id" {
				value = data.databricks_account_group.data_eng_group.id
            }
			`,
		},
	})
}
