package acceptance

import (
	"context"
	"os"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/providers"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw"
	"github.com/databricks/terraform-provider-databricks/internal/providers/sdkv2"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

func initUnifiedHostEnv(t *testing.T) {
	initTest(t, "account")
	unifiedHost := os.Getenv("UNIFIED_HOST")
	if unifiedHost == "" {
		Skipf(t)("UNIFIED_HOST environment variable is missing")
	}
	// Override provider env vars so the Terraform provider uses unified host.
	os.Setenv("DATABRICKS_HOST", unifiedHost)
	os.Setenv("DATABRICKS_EXPERIMENTAL_IS_UNIFIED_HOST", "true")
}

// noOidcProviderFactories creates provider factories that skip the OidcConfigCustomizer.
// This is needed for unified host tests where the auth is handled via client credentials,
// not github-oidc.
func noOidcProviderFactories() map[string]func() (tfprotov6.ProviderServer, error) {
	return map[string]func() (tfprotov6.ProviderServer, error){
		"databricks": func() (tfprotov6.ProviderServer, error) {
			ctx := context.Background()
			sdkPluginProvider := sdkv2.DatabricksProvider(sdkv2.WithConfigCustomizer(DefaultConfigCustomizer))
			pluginFrameworkProvider := pluginfw.GetDatabricksProviderPluginFramework(pluginfw.WithConfigCustomizer(DefaultConfigCustomizer))
			return providers.GetProviderServer(ctx, providers.WithSdkV2Provider(sdkPluginProvider), providers.WithPluginFrameworkProvider(pluginFrameworkProvider))
		},
	}
}

func TestAccUnifiedHostCreateJobsAWS(t *testing.T) {
	initUnifiedHostEnv(t)
	if !IsAws(t) {
		Skipf(t)("This test is only running on AWS")
	}
	run(t, []Step{
		{
			Template: `
			resource "databricks_mws_workspaces" "ws1" {
				account_id     = "{env.DATABRICKS_ACCOUNT_ID}"
				workspace_name = "tf-unified-{var.RANDOM}-1"
				aws_region     = "{env.AWS_REGION}"
				compute_mode   = "SERVERLESS"
			}

			resource "databricks_mws_workspaces" "ws2" {
				account_id     = "{env.DATABRICKS_ACCOUNT_ID}"
				workspace_name = "tf-unified-{var.RANDOM}-2"
				aws_region     = "{env.AWS_REGION}"
				compute_mode   = "SERVERLESS"
			}

			resource "databricks_job" "ws1" {
				name = "tf-unified-{var.RANDOM}-job-1"
				provider_config {
					workspace_id = databricks_mws_workspaces.ws1.workspace_id
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

			resource "databricks_job" "ws2" {
				name = "tf-unified-{var.RANDOM}-job-2"
				provider_config {
					workspace_id = databricks_mws_workspaces.ws2.workspace_id
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
			ProtoV6ProviderFactories: noOidcProviderFactories(),
		},
	})
}
