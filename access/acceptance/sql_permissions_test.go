package acceptance

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/commands"
	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/internal/acceptance"
	"github.com/databrickslabs/terraform-provider-databricks/internal/compute"
	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccTableACL(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	if cloudEnv == "" {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}

	client := common.CommonEnvironmentClient()
	client.WithCommandExecutor(func(ctx context.Context,
		dc *common.DatabricksClient) common.CommandExecutor {
		return commands.NewCommandsAPI(ctx, dc)
	})

	shell := client.CommandExecutor(context.Background())
	clusterInfo := compute.NewTinyClusterInCommonPoolPossiblyReused()
	talbeName := qa.RandomName("table_acl_")

	cr := shell.Execute(clusterInfo.ClusterID, "python",
		fmt.Sprintf("spark.range(10).write.saveAsTable('%s')",
			talbeName))
	require.False(t, cr.Failed(), cr.Error())
	os.Setenv("TABLE_ACL_TEST_TABLE", talbeName)
	defer func() {
		cr := shell.Execute(clusterInfo.ClusterID, "sql",
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
