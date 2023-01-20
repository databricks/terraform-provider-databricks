package acceptance

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestUcAccGroupDataWorkspace(t *testing.T) {
	qa.RequireCloudEnv(t, "ucws")
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `data "databricks_account_group" "account_users" {
				display_name = account users"
			}`,
			Check: func(s *terraform.State) error {
				_, ok := s.Modules[0].Resources["data.databricks_group.account_users"]
				require.True(t, ok, "data.databricks_group.account_users has to be there")
				return nil
			},
		},
	})
}
