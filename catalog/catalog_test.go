package catalog_test

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stretchr/testify/require"
)

func TestUcAccCatalog(t *testing.T) {
	acceptance.LoadUcwsEnv(t)
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: fmt.Sprintf(`
		resource "databricks_catalog" "sandbox" {
			name         = "sandbox{var.RANDOM}"
			comment      = "this catalog is managed by terraform"
			properties = {
				purpose = "testing"
			}
			%s
		}`, getPredictiveOptimizationSetting(t, true)),
	})
}

func TestUcAccCatalogIsolated(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: `
		resource "databricks_catalog" "sandbox" {
			name           = "sandbox{var.STICKY_RANDOM}"
			comment        = "this catalog is managed by terraform"
			properties     = {
				purpose = "testing"
			}
		}`,
	}, acceptance.Step{
		Template: `
		resource "databricks_catalog" "sandbox" {
			name           = "sandbox{var.STICKY_RANDOM}"
			isolation_mode = "ISOLATED"
			comment        = "this catalog is managed by terraform"
			properties     = {
				purpose = "testing"
			}
		}`,
	}, acceptance.Step{
		Template: `
		resource "databricks_catalog" "sandbox" {
			name           = "sandbox{var.STICKY_RANDOM}"
			isolation_mode = "OPEN"
			comment        = "this catalog is managed by terraform"
			properties     = {
				purpose = "testing"
			}
		}`,
	})
}

func TestUcAccCatalogUpdate(t *testing.T) {
	acceptance.LoadUcwsEnv(t)
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: fmt.Sprintf(`
		resource "databricks_catalog" "sandbox" {
			name           = "sandbox{var.STICKY_RANDOM}"
			comment        = "this catalog is managed by terraform"
			properties     = {
				purpose = "testing"
			}
			%s
		}`, getPredictiveOptimizationSetting(t, true)),
	}, acceptance.Step{
		Template: fmt.Sprintf(`
		resource "databricks_catalog" "sandbox" {
			name           = "sandbox{var.STICKY_RANDOM}"
			comment        = "this catalog is managed by terraform"
			properties     = {
				purpose = "testing"
			}
			%s
			owner = "account users"
		}`, getPredictiveOptimizationSetting(t, true)),
	}, acceptance.Step{
		Template: fmt.Sprintf(`
		resource "databricks_catalog" "sandbox" {
			name           = "sandbox{var.STICKY_RANDOM}"
			comment        = "this catalog is managed by terraform"
			properties     = {
				purpose = "testing"
			}
			%s
			owner = "{env.TEST_DATA_ENG_GROUP}"
		}`, getPredictiveOptimizationSetting(t, true)),
	}, acceptance.Step{
		Template: fmt.Sprintf(`
		resource "databricks_catalog" "sandbox" {
			name           = "sandbox{var.STICKY_RANDOM}"
			comment        = "this catalog is managed by terraform - updated comment"
			properties     = {
				purpose = "testing"
			}
			%s
			owner = "{env.TEST_METASTORE_ADMIN_GROUP_NAME}"
		}`, getPredictiveOptimizationSetting(t, false)),
	}, acceptance.Step{
		// Adding options should cause the catalog to be recreated.
		Template: fmt.Sprintf(`
		resource "databricks_catalog" "sandbox" {
			name           = "sandbox{var.STICKY_RANDOM}"
			comment        = "this catalog is managed by terraform - updated comment"
			properties     = {
				purpose = "testing"
			}
			options = {
				user = "miles"
			}
			%s
			owner = "{env.TEST_METASTORE_ADMIN_GROUP_NAME}"
		}`, getPredictiveOptimizationSetting(t, false)),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				plancheck.ExpectResourceAction("databricks_catalog.sandbox", plancheck.ResourceActionDestroyBeforeCreate),
			},
		},
	})
}

// Create a connection to an HMS catalog and update authorized_paths.
func TestUcAccCatalogHmsConnectionUpdate(t *testing.T) {
	authorizedPath := fmt.Sprintf("s3://%s/path/to/authorized", qa.RandomName("hms-bucket-"))
	otherAuthorizedPath := fmt.Sprintf("s3://%s/path/to/authorized", qa.RandomName("hms-other-bucket-"))
	otherInfra := fmt.Sprintf(`
		resource "databricks_connection" "sandbox" {
			name = "hms_connection{var.STICKY_RANDOM}"
			connection_type = "HIVE_METASTORE"
			comment         = "created in TestUcAccCatalogHmsConnectionUpdate"
			options = {
				host     = "test.mysql.database.azure.com"
				port     = "3306"
				user     = "user"
				password = "password"
				database = "metastore{var.STICKY_RANDOM}"
				db_type  = "MYSQL"
				version  = "2.3"
			}
			properties = {
				purpose = "testing"
			}
		}
		resource "databricks_storage_credential" "external" {
			name = "cred-{var.STICKY_RANDOM}"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			comment = "created in TestUcAccCatalogHmsConnectionUpdate"
		}
		resource "databricks_external_location" "sandbox" {
		    name = "sandbox{var.STICKY_RANDOM}"
			comment = "created in TestUcAccCatalogHmsConnectionUpdate"
			url = "%s"
			credential_name = "${databricks_storage_credential.external.name}"
			skip_validation = true
		}
		resource "databricks_external_location" "sandbox-other" {
		    name = "sandbox-other{var.STICKY_RANDOM}"
			comment = "created in TestUcAccCatalogHmsConnectionUpdate"
			url = "%s"
			credential_name = "${databricks_storage_credential.external.name}"
			skip_validation = true
		}
	`, authorizedPath, otherAuthorizedPath)
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: otherInfra + fmt.Sprintf(`
		resource "databricks_catalog" "sandbox" {
			name           = "sandbox{var.STICKY_RANDOM}"
			comment        = "created in TestUcAccCatalogHmsConnectionUpdate"
			connection_name = "${databricks_connection.sandbox.name}"
			options = {
				authorized_paths = "%s"
			}
			lifecycle {
			    prevent_destroy = true
			}
			depends_on = [databricks_external_location.sandbox, databricks_external_location.sandbox-other]
		}`, authorizedPath),
	}, acceptance.Step{
		Template: otherInfra + fmt.Sprintf(`
		resource "databricks_catalog" "sandbox" {
			name           = "sandbox{var.STICKY_RANDOM}"
			comment        = "created in TestUcAccCatalogHmsConnectionUpdate"
			connection_name = "${databricks_connection.sandbox.name}"
			options = {
				authorized_paths = "%s,%s"
			}
			lifecycle {
			    prevent_destroy = true
			}
		}`, authorizedPath, otherAuthorizedPath),
	}, acceptance.Step{
		Template: otherInfra + fmt.Sprintf(`
		resource "databricks_catalog" "sandbox" {
			name           = "sandbox{var.STICKY_RANDOM}"
			comment        = "created in TestUcAccCatalogHmsConnectionUpdate"
			connection_name = "${databricks_connection.sandbox.name}"
			options = {
				authorized_paths = "%s,%s"
				other_option = "value"
			}
			lifecycle {
			    prevent_destroy = true
			}
		}`, authorizedPath, otherAuthorizedPath),
		ExpectError: regexp.MustCompile("Instance cannot be destroyed"),
	}, acceptance.Step{
		Template: otherInfra + fmt.Sprintf(`
		resource "databricks_catalog" "sandbox" {
			name           = "sandbox{var.STICKY_RANDOM}"
			comment        = "created in TestUcAccCatalogHmsConnectionUpdate"
			connection_name = "${databricks_connection.sandbox.name}"
			options = {
				authorized_paths = "%s,%s"
			}
		}`, authorizedPath, otherAuthorizedPath),
	})
}

func catalogProviderConfigTemplate(catalogName string, providerConfig string) string {
	return fmt.Sprintf(`
	resource "databricks_catalog" "this" {
		name = "%s"
		comment = "test catalog"
		force_destroy = true
		%s
	}
`, catalogName, providerConfig)
}

func TestAccCatalog_ProviderConfig_Invalid(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: catalogProviderConfigTemplate("test_catalog_{var.STICKY_RANDOM}", `
			provider_config {
				workspace_id = "invalid"
			}
		`),
		ExpectError: regexp.MustCompile(`workspace_id must be a positive integer without leading zeros`),
		PlanOnly:    true,
	})
}

func TestAccCatalog_ProviderConfig_Required(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: catalogProviderConfigTemplate("test_catalog_{var.STICKY_RANDOM}", `
			provider_config {
			}
		`),
		ExpectError: regexp.MustCompile(`The argument "workspace_id" is required, but no definition was found.`),
		PlanOnly:    true,
	})
}

func TestAccCatalog_ProviderConfig_EmptyID(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: catalogProviderConfigTemplate("test_catalog_{var.STICKY_RANDOM}", `
			provider_config {
				workspace_id = ""
			}
		`),
		ExpectError: regexp.MustCompile(`expected "provider_config.0.workspace_id" to not be an empty string`),
		PlanOnly:    true,
	})
}

func TestAccCatalog_ProviderConfig_Mismatched(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: catalogProviderConfigTemplate("test_catalog_{var.STICKY_RANDOM}", `
			provider_config {
				workspace_id = "123"
			}
		`),
		ExpectError: regexp.MustCompile(`workspace_id mismatch.*please check the workspace_id provided in provider_config`),
	})
}

func TestAccCatalog_ProviderConfig_Match(t *testing.T) {
	acceptance.LoadUcwsEnv(t)
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	workspaceID, err := w.CurrentWorkspaceID(ctx)
	require.NoError(t, err)
	workspaceIDStr := strconv.FormatInt(workspaceID, 10)
	catalogName := "test_catalog_{var.STICKY_RANDOM}"
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: catalogProviderConfigTemplate(catalogName, ""),
	}, acceptance.Step{
		Template: catalogProviderConfigTemplate(catalogName, fmt.Sprintf(`
			provider_config {
				workspace_id = "%s"
			}
		`, workspaceIDStr)),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				plancheck.ExpectResourceAction("databricks_catalog.this", plancheck.ResourceActionUpdate),
			},
		},
	})
}

func TestAccCatalog_ProviderConfig_Recreate(t *testing.T) {
	acceptance.LoadUcwsEnv(t)
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	workspaceID, err := w.CurrentWorkspaceID(ctx)
	require.NoError(t, err)
	workspaceIDStr := strconv.FormatInt(workspaceID, 10)
	catalogName := "test_catalog_{var.STICKY_RANDOM}"
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: catalogProviderConfigTemplate(catalogName, ""),
	}, acceptance.Step{
		Template: catalogProviderConfigTemplate(catalogName, fmt.Sprintf(`
			provider_config {
				workspace_id = "%s"
			}
		`, workspaceIDStr)),
	}, acceptance.Step{
		Template: catalogProviderConfigTemplate(catalogName, `
			provider_config {
				workspace_id = "123"
			}
		`),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PostApplyPreRefresh: []plancheck.PlanCheck{
				plancheck.ExpectResourceAction("databricks_catalog.this", plancheck.ResourceActionUpdate),
			},
		},
		PlanOnly:           true,
		ExpectNonEmptyPlan: true,
	})
}

func TestAccCatalog_ProviderConfig_Remove(t *testing.T) {
	acceptance.LoadUcwsEnv(t)
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	workspaceID, err := w.CurrentWorkspaceID(ctx)
	require.NoError(t, err)
	workspaceIDStr := strconv.FormatInt(workspaceID, 10)
	catalogName := "test_catalog_{var.STICKY_RANDOM}"
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: catalogProviderConfigTemplate(catalogName, ""),
	}, acceptance.Step{
		Template: catalogProviderConfigTemplate(catalogName, fmt.Sprintf(`
			provider_config {
				workspace_id = "%s"
			}
		`, workspaceIDStr)),
	}, acceptance.Step{
		Template: catalogProviderConfigTemplate(catalogName, ""),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				plancheck.ExpectResourceAction("databricks_catalog.this", plancheck.ResourceActionUpdate),
			},
		},
	})
}
