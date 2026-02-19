package sql_test

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/qa/lock"
	"github.com/databricks/databricks-sdk-go/qa/lock/core"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type SqlGlobalConfigLock struct {
	WorkspaceHost string
}

func (s SqlGlobalConfigLock) GetLockId() string {
	return fmt.Sprintf("WorkspaceHost:%s;type=SqlGlobalConfig", strings.ReplaceAll(s.WorkspaceHost, "/", "_"))
}

func getSqlGlobalConfigLockable(t *testing.T) core.Lockable {
	return lock.NewLockable(SqlGlobalConfigLock{WorkspaceHost: acceptance.GetEnvOrSkipTest(t, "DATABRICKS_HOST")})
}

func makeSqlGlobalConfig(extraConfig string) string {
	return fmt.Sprintf(`
resource "databricks_sql_global_config" "this" {
	data_access_config = {
		"spark.sql.session.timeZone": "UTC"
	}
	%s
}`, extraConfig)
}

func TestAccSQLGlobalConfig(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	acceptance.WorkspaceLevel(t, acceptance.Step{
		PreConfig: func() {
			ctx := context.Background()
			_, err := lock.Acquire(ctx, getSqlGlobalConfigLockable(t), lock.InTest(t))
			require.NoError(t, err)
		},
		Template: makeSqlGlobalConfig(""),
	})
}

func TestAccSQLGlobalConfigServerless(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	if acceptance.IsGcp(t) {
		acceptance.Skipf(t)("GCP does not support serverless compute")
	}

	checkServerlessEnabled := func(enabled bool) func(state *terraform.State) error {
		return func(state *terraform.State) error {
			enableServerlessComputeStr := state.Modules[0].Resources["databricks_sql_global_config.this"].Primary.Attributes["enable_serverless_compute"]
			enableServerlessCompute, err := strconv.ParseBool(enableServerlessComputeStr)
			require.NoError(t, err)
			assert.Equal(t, enabled, enableServerlessCompute)
			return nil
		}
	}

	acceptance.WorkspaceLevel(t, acceptance.Step{
		PreConfig: func() {
			ctx := context.Background()
			_, err := lock.Acquire(ctx, getSqlGlobalConfigLockable(t), lock.InTest(t))
			require.NoError(t, err)
		},
		Template: makeSqlGlobalConfig("enable_serverless_compute = true"),
		Check:    checkServerlessEnabled(true),
	}, acceptance.Step{
		Template: makeSqlGlobalConfig(""),
		Check:    checkServerlessEnabled(true),
	}, acceptance.Step{
		Template: makeSqlGlobalConfig("enable_serverless_compute = false"),
		Check:    checkServerlessEnabled(false),
	})
}
