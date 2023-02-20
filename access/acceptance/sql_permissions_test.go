package acceptance

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccTableACL(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	if cloudEnv == "" {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	t.Parallel()

	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	info, err := w.Clusters.GetOrCreateRunningCluster(ctx, "tf-dummy")
	require.NoError(t, err)

	talbeName := qa.RandomName("table_acl_")
	cr := w.CommandExecutor.Execute(ctx, info.ClusterId, "python",
		fmt.Sprintf("spark.range(10).write.saveAsTable('%s')",
			talbeName))
	require.False(t, cr.Failed(), cr.Error())
	t.Setenv("TABLE_ACL_TEST_TABLE", talbeName)
	defer func() {
		cr := w.CommandExecutor.Execute(ctx, info.ClusterId, "sql",
			fmt.Sprintf("DROP TABLE %s", talbeName))
		assert.False(t, cr.Failed(), cr.Error())
	}()

	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			resource "databricks_sql_permissions" "this" {
				table = "{env.TABLE_ACL_TEST_TABLE}"

				privilege_assignments {
					principal = "users"
					privileges = ["SELECT"]
				}
			}`,
		},
	})
}
