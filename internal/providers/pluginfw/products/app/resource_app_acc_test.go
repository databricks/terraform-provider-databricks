package app_test

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const baseResources = `
	resource "databricks_secret_scope" "this" {
		name = "tf-{var.STICKY_RANDOM}"
	}

	resource "databricks_secret" "this" {
	    scope = databricks_secret_scope.this.name
		key = "tf-{var.STICKY_RANDOM}"
		string_value = "secret"
	}

	resource "databricks_sql_endpoint" "this" {
		name = "tf-{var.STICKY_RANDOM}"
		cluster_size = "2X-Small"
		max_num_clusters = 1

		tags {
			custom_tags {
				key   = "Owner"
				value = "eng-dev-ecosystem-team_at_databricks.com"
			}
		}
	}

	resource "databricks_job" "this" {
		name = "tf-{var.STICKY_RANDOM}"
	}

	resource "databricks_model_serving" "this" {
		name = "tf-{var.STICKY_RANDOM}"
		config {
			served_models {
				name = "prod_model"
				model_name = "experiment-fixture-model"
				model_version = "1"
				workload_size = "Small"
				scale_to_zero_enabled = true
			}
		}
	}
`

func makeTemplate(description string) string {
	appTemplate := baseResources + `
	resource "databricks_app" "this" {
		name = "tf-{var.STICKY_RANDOM}"
		description = "%s"
		resources = [{
			name = "secret"
			description = "secret for app"
			secret = {
				scope = databricks_secret_scope.this.name
				key = databricks_secret.this.key
				permission = "MANAGE"
			}
		}, {
			name = "warehouse"
			description = "warehouse for app"
			job = {
				id = databricks_job.this.id
				permission = "CAN_MANAGE"
			}
		}, {
			name = "serving endpoint"
			description = "serving endpoint for app"
			serving_endpoint = {
				name = databricks_model_serving.this.name
				permission = "CAN_MANAGE"
			}
		}, {
			name = "sql warehouse"
			description = "sql warehouse for app"
			sql_warehouse = {
				id = databricks_sql_endpoint.this.id
				permission = "CAN_MANAGE"
			}
		}]
	}`
	return fmt.Sprintf(appTemplate, description)
}

var templateWithInvalidResource = `
	resource "databricks_app" "this" {
		name = "tf-{var.STICKY_RANDOM}"
		description = "My app"
		resources = [{
			name = "invalid resource"
			description = "invalid resource for app"
			secret = {
				permission = "CAN_MANAGE"
				key = "test"
				scope = "test"
			}
			sql_warehouse = {
				id = "123"
				permission = "CAN_MANAGE"
			}
		}]
	}`

func TestAccApp_InvalidResource(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: templateWithInvalidResource,
		ExpectError: regexp.MustCompile(`2 attributes specified when one \(and only one\) of
\[.*resources\[0\].*<\.job,.*\.<\.serving_endpoint,.*\.<\.sql_warehouse,.*\.<\.uc_securable.*\]
is required`),
	})
}

func TestAccAppResource(t *testing.T) {
	var updateTime string
	acceptance.LoadWorkspaceEnv(t)
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: makeTemplate("My app"),
		Check: func(s *terraform.State) error {
			updateTime = s.RootModule().Resources["databricks_app.this"].Primary.Attributes["update_time"]
			return nil
		},
	}, acceptance.Step{
		Template: makeTemplate("My new app"),
		Check: func(s *terraform.State) error {
			var newUpdateTime = s.RootModule().Resources["databricks_app.this"].Primary.Attributes["update_time"]
			assert.NotEqual(t, updateTime, newUpdateTime)
			return nil
		},
	}, acceptance.Step{
		ImportState:       true,
		ResourceName:      "databricks_app.this",
		ImportStateIdFunc: acceptance.BuildImportStateIdFunc("databricks_app.this", "name"),
		// I cannot enable ImportStateVerify because computed fields don't appear to be filled in during import.
		// ImportStateVerify: true,
		ImportStateVerifyIdentifierAttribute: "name",
	})
}

func TestAccAppResource_NoCompute(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
	resource "databricks_secret_scope" "this" {
		name = "tf-{var.STICKY_RANDOM}"
	}

	resource "databricks_secret" "this" {
	    scope = databricks_secret_scope.this.name
		key = "tf-{var.STICKY_RANDOM}"
		string_value = "secret"
	}
	resource "databricks_app" "this" {
		no_compute = true
		name = "tf-{var.STICKY_RANDOM}"
		description = "no_compute app"
		resources = [{
			name = "secret"
			description = "secret for app"
			secret = {
				scope = databricks_secret_scope.this.name
				key = databricks_secret.this.key
				permission = "MANAGE"
			}
		}]
	}
		`,
		Check: func(s *terraform.State) error {
			computeStatus := s.RootModule().Resources["databricks_app.this"].Primary.Attributes["compute_status.state"]
			assert.Equal(t, "STOPPED", computeStatus)
			return nil
		},
	})
}

var deletedOutsideTemplate = `
	resource "databricks_secret_scope" "this" {
		name = "tf-{var.STICKY_RANDOM}"
	}

	resource "databricks_secret" "this" {
	    scope = databricks_secret_scope.this.name
		key = "tf-{var.STICKY_RANDOM}"
		string_value = "secret"
	}

	resource "databricks_app" "this" {
		no_compute = true
		name = "tf-{var.STICKY_RANDOM}"
		description = "deleted outside terraform test"
		resources = [{
			name = "secret"
			description = "secret for app"
			secret = {
				scope = databricks_secret_scope.this.name
				key = databricks_secret.this.key
				permission = "MANAGE"
			}
		}]
	}
`

func TestAccAppResource_DeletedOutsideTerraform(t *testing.T) {
	var appName string
	acceptance.LoadWorkspaceEnv(t)
	if acceptance.IsGcp(t) {
		acceptance.Skipf(t)("not available on GCP")
	}
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: deletedOutsideTemplate,
		Check: func(s *terraform.State) error {
			appName = s.RootModule().Resources["databricks_app.this"].Primary.Attributes["name"]
			return nil
		},
	}, acceptance.Step{
		PreConfig: func() {
			w := databricks.Must(databricks.NewWorkspaceClient())
			_, err := w.Apps.DeleteByName(context.Background(), appName)
			require.NoError(t, err)
			// Wait for the app to be fully deleted before proceeding,
			// otherwise recreating it will conflict with the DELETING state.
			for {
				_, err := w.Apps.GetByName(context.Background(), appName)
				if apierr.IsMissing(err) {
					break
				}
				require.NoError(t, err)
				time.Sleep(30 * time.Second)
			}
		},
		Template: deletedOutsideTemplate,
	})
}

func appTemplate(provider_config string) string {
	return fmt.Sprintf(`
		resource "databricks_secret_scope" "this" {
		name = "tf-{var.STICKY_RANDOM}"
	}
	resource "databricks_secret" "this" {
	    scope = databricks_secret_scope.this.name
		key = "tf-{var.STICKY_RANDOM}"
		string_value = "secret"
	}
	resource "databricks_app" "this" {
		%s
		no_compute = true
		name = "tf-{var.STICKY_RANDOM}"
		description = "no_compute app"
		resources = [{
			name = "secret"
			description = "secret for app"
			secret = {
				scope = databricks_secret_scope.this.name
				key = databricks_secret.this.key
				permission = "MANAGE"
			}
		}]
	}
	`, provider_config)
}

func TestAccApp_ProviderConfig_Invalid(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: appTemplate(`
			provider_config = {
				workspace_id = "invalid"
			}
		`),
		ExpectError: regexp.MustCompile(
			`Attribute provider_config\.workspace_id\s+workspace_id must be a valid integer`,
		),
		PlanOnly: true,
	})
}

func TestAccApp_ProviderConfig_Mismatched(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: appTemplate(`
			provider_config = {
				workspace_id = "123"
			}
		`),
		ExpectError: regexp.MustCompile(
			`(?s)failed to get workspace client.*workspace_id mismatch` +
				`.*please check the workspace_id provided in ` +
				`provider_config`,
		),
	})
}

func TestAccApp_ProviderConfig_Apply(t *testing.T) {
	acceptance.LoadUcwsEnv(t)
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	workspaceID, err := w.CurrentWorkspaceID(ctx)
	require.NoError(t, err)
	workspaceIDStr := strconv.FormatInt(workspaceID, 10)
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: appTemplate(``),
	}, acceptance.Step{
		Template: appTemplate(fmt.Sprintf(`
			provider_config = {
				workspace_id = "%s"
			}
		`, workspaceIDStr)),
	})
}
