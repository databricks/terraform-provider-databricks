package acceptance

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	databricks "github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw"
	"github.com/databricks/terraform-provider-databricks/internal/providers/sdkv2"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

func unifiedHostProviderFactories(unifiedHost, accountID string) map[string]func() (tfprotov6.ProviderServer, error) {
	customizer := func(cfg *config.Config) error {
		cfg.Host = unifiedHost
		cfg.AccountID = accountID
		return nil
	}
	return map[string]func() (tfprotov6.ProviderServer, error){
		"databricks": func() (tfprotov6.ProviderServer, error) {
			ctx := context.Background()
			sdkPluginProvider := sdkv2.DatabricksProvider(
				sdkv2.WithConfigCustomizer(DefaultConfigCustomizer),
				sdkv2.WithConfigCustomizer(customizer),
			)
			pluginFrameworkProvider := pluginfw.GetDatabricksProviderPluginFramework(
				pluginfw.WithConfigCustomizer(DefaultConfigCustomizer),
				pluginfw.WithConfigCustomizer(customizer),
			)
			return providers.GetProviderServer(ctx, providers.WithSdkV2Provider(sdkPluginProvider), providers.WithPluginFrameworkProvider(pluginFrameworkProvider))
		},
	}
}

func initUnifiedHostAccountEnv(t *testing.T) {
	LoadAccountEnv(t)
	unifiedHost := os.Getenv("UNIFIED_HOST")
	if unifiedHost == "" {
		Skipf(t)("UNIFIED_HOST environment variable is missing")
	}
}

func initUnifiedHostWorkspaceEnv(t *testing.T) {
	LoadWorkspaceEnv(t)
	unifiedHost := os.Getenv("UNIFIED_HOST")
	if unifiedHost == "" {
		Skipf(t)("UNIFIED_HOST environment variable is missing")
	}
}

// createJobWithProviderConfig is a shared test helper that creates a job with
// provider_config { workspace_id } and verifies it. If providerFactories is nil,
// it uses the default provider factories.
func createJobWithProviderConfig(t *testing.T, workspaceID string, providerFactories map[string]func() (tfprotov6.ProviderServer, error)) {
	jobName := "tf-" + RandomName() + "-job-1"

	step := Step{
		Template: `
		resource "databricks_job" "j1" {
			name = "` + jobName + `"
			provider_config {
				workspace_id = ` + workspaceID + `
			}
			task {
				task_key = "check"
				condition_task {
					left  = "true"
					op    = "EQUAL_TO"
					right = "true"
				}
			}
		}
		`,
		Check: ResourceCheck("databricks_job.j1", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.GetWorkspaceClientForUnifiedProvider(ctx, workspaceID)
			if err != nil {
				return err
			}
			jobID, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				return err
			}
			res, err := w.Jobs.Get(ctx, jobs.GetJobRequest{JobId: jobID})
			if err != nil {
				return err
			}
			if res.Settings.Name != jobName {
				return fmt.Errorf("expected job name %q, got %q", jobName, res.Settings.Name)
			}
			return nil
		}),
	}
	if providerFactories != nil {
		step.ProtoV6ProviderFactories = providerFactories
	}
	run(t, []Step{step})
}

func TestMwsAccUnifiedHostCreateJobs(t *testing.T) {
	initUnifiedHostAccountEnv(t)
	unifiedHost := os.Getenv("UNIFIED_HOST")
	workspaceID := GetEnvOrSkipTest(t, "TEST_WORKSPACE_ID")
	accountID := GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID")
	createJobWithProviderConfig(t, workspaceID, unifiedHostProviderFactories(unifiedHost, accountID))
}

func TestAccUnifiedHostWorkspaceCreateJobs(t *testing.T) {
	initUnifiedHostWorkspaceEnv(t)
	unifiedHost := os.Getenv("UNIFIED_HOST")
	accountID := GetEnvOrSkipTest(t, "TEST_ACCOUNT_ID")
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	workspaceID, err := w.CurrentWorkspaceID(ctx)
	if err != nil {
		t.Fatalf("failed to get current workspace ID: %s", err)
	}
	createJobWithProviderConfig(t, strconv.FormatInt(workspaceID, 10), unifiedHostProviderFactories(unifiedHost, accountID))
}

func TestMwsAccAccountHostCreateJobs(t *testing.T) {
	LoadAccountEnv(t)
	workspaceID := GetEnvOrSkipTest(t, "TEST_WORKSPACE_ID")
	createJobWithProviderConfig(t, workspaceID, nil)
}
