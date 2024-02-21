package acceptance

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/qa/lock"
	"github.com/databricks/databricks-sdk-go/qa/lock/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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
	return lock.NewLockable(SqlGlobalConfigLock{WorkspaceHost: GetEnvOrSkipTest(t, "DATABRICKS_HOST")})
}

func checkServerlessEnabled(t *testing.T, state *terraform.State, enabled bool) {
	enableServerlessComputeStr := state.Modules[0].Resources["databricks_sql_global_config.this"].Primary.Attributes["enable_serverless_compute"]
	enableServerlessCompute, err := strconv.ParseBool(enableServerlessComputeStr)
	require.NoError(t, err)
	assert.Equal(t, enabled, enableServerlessCompute)
}

func TestAccSQLGlobalConfig(t *testing.T) {
	loadDebugEnvIfRunsFromIDE(t, "workspace")
	ctx := context.Background()
	_, err := lock.Acquire(ctx, getSqlGlobalConfigLockable(t), lock.InTest(t))
	require.NoError(t, err)
	workspaceLevel(t, step{
		Template: `resource "databricks_sql_global_config" "this" {
			data_access_config = {
				"spark.sql.session.timeZone": "UTC"
			}  
		}`,
	})
}

func TestAccSQLGlobalConfigServerless(t *testing.T) {
	loadDebugEnvIfRunsFromIDE(t, "workspace")
	if strings.Contains(os.Getenv("CLOUD_ENV"), "gcp") {
		skipf(t)("GCP does not support serverless compute")
	}
	ctx := context.Background()
	l, err := lock.Acquire(ctx, getSqlGlobalConfigLockable(t), lock.InTest(t))
	require.NoError(t, err)
	defer l.Unlock()
	workspaceLevel(t, step{
		Template: `resource "databricks_sql_global_config" "this" {
			enable_serverless_compute = true
			data_access_config = {
				"spark.sql.session.timeZone": "UTC"
			}
		}`,
		Check: func(s *terraform.State) error {
			checkServerlessEnabled(t, s, true)
			return nil
		},
	}, step{
		Template: `resource "databricks_sql_global_config" "this" {
			data_access_config = {
				"spark.sql.session.timeZone": "UTC"
			}
		}`,
		Check: func(s *terraform.State) error {
			checkServerlessEnabled(t, s, true)
			return nil
		},
	})
}
