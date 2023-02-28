package acceptance

import (
	"context"
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestAccTableACL(t *testing.T) {
	talbeName := qa.RandomName("table_acl_")
	workspaceLevel(t, step{
		Template: `data "databricks_current_user" "me" {}`,
		Check: func(s *terraform.State) error {
			// we need to run table creation within check code of previous step,
			// so that we know that environment variables are loaded by WorkspaceLevel
			// in the necessary way.
			ctx := context.Background()
			w := databricks.Must(databricks.NewWorkspaceClient())
			info, err := w.Clusters.GetOrCreateRunningCluster(ctx, "tf-dummy")
			require.NoError(t, err)
			cr := w.CommandExecutor.Execute(ctx, info.ClusterId, "python",
				fmt.Sprintf("spark.range(10).write.saveAsTable('%s')", talbeName))
			require.False(t, cr.Failed(), cr.Error())

			t.Cleanup(func() {
				cr := w.CommandExecutor.Execute(ctx, info.ClusterId, "sql",
					fmt.Sprintf("DROP TABLE %s", talbeName))
				require.False(t, cr.Failed(), cr.Error())
			})
			return nil
		},
	}, step{
		Template: `
		resource "databricks_sql_permissions" "this" {
			table = "` + talbeName + `"

			privilege_assignments {
				principal = "users"
				privileges = ["SELECT"]
			}
		}`,
	})
}
