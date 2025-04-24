package access_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/require"
)

func TestAccSqlPermissions(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	tableName := qa.RandomName("table_acl_")
	clusterId := acceptance.GetEnvOrSkipTest(t, "TEST_TABLE_ACL_CLUSTER_ID")
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	err := w.Clusters.EnsureClusterIsRunning(ctx, clusterId)
	require.NoError(t, err)

	executor, err := w.CommandExecution.Start(ctx, clusterId, compute.LanguagePython)
	require.NoError(t, err)
	t.Cleanup(func() {
		err = executor.Destroy(ctx)
		require.NoError(t, err)
	})

	cr, err := executor.Execute(ctx, fmt.Sprintf("spark.range(10).write.saveAsTable('%s')", tableName))
	require.NoError(t, err)
	require.False(t, cr.Failed(), cr.Error())

	t.Cleanup(func() {
		cr, err = executor.Execute(ctx, fmt.Sprintf(`spark.sql("DROP TABLE %s")`, tableName))
		require.NoError(t, err)
		require.False(t, cr.Failed(), cr.Error())
	})
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		resource "databricks_sql_permissions" "this" {
			table = "` + tableName + `"

			privilege_assignments {
				principal = "users"
				privileges = ["SELECT"]
			}
		}`,
	})
}
