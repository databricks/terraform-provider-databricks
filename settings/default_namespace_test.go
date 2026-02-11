package settings_test

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func defaultNamespaceSettingTemplate(providerConfig string) string {
	return fmt.Sprintf(`
	resource "databricks_default_namespace_setting" "this" {
		namespace {
			value = "namespace_value"
		}
		%s
	}
	`, providerConfig)
}

func TestAccDefaultNamespaceSetting_ProviderConfig_Invalid(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: defaultNamespaceSettingTemplate(`
			provider_config {
				workspace_id = "invalid"
			}
		`),
		ExpectError: regexp.MustCompile(`workspace_id must be a positive integer without leading zeros`),
		PlanOnly:    true,
	})
}

func TestAccDefaultNamespaceSetting_ProviderConfig_Mismatched(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: defaultNamespaceSettingTemplate(`
			provider_config {
				workspace_id = "123"
			}
		`),
		ExpectError: regexp.MustCompile(`workspace_id mismatch.*please check the workspace_id provided in provider_config`),
	})
}

func TestAccDefaultNamespaceSetting_ProviderConfig_Required(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: defaultNamespaceSettingTemplate(`
			provider_config {
			}
		`),
		ExpectError: regexp.MustCompile(`The argument "workspace_id" is required, but no definition was found.`),
		PlanOnly:    true,
	})
}

func TestAccDefaultNamespaceSetting_ProviderConfig_EmptyID(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: defaultNamespaceSettingTemplate(`
			provider_config {
				workspace_id = ""
			}
		`),
		ExpectError: regexp.MustCompile(`expected "provider_config.0.workspace_id" to not be an empty string`),
		PlanOnly:    true,
	})
}

func TestAccDefaultNamespaceSetting_ProviderConfig_Match(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	workspaceID, err := w.CurrentWorkspaceID(ctx)
	require.NoError(t, err)
	workspaceIDStr := strconv.FormatInt(workspaceID, 10)
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: defaultNamespaceSettingTemplate(""),
	}, acceptance.Step{
		Template: defaultNamespaceSettingTemplate(fmt.Sprintf(`
			provider_config {
				workspace_id = "%s"
			}
		`, workspaceIDStr)),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				plancheck.ExpectResourceAction("databricks_default_namespace_setting.this", plancheck.ResourceActionUpdate),
			},
		},
	})
}

func TestAccDefaultNamespaceSetting_ProviderConfig_Recreate(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	workspaceID, err := w.CurrentWorkspaceID(ctx)
	require.NoError(t, err)
	workspaceIDStr := strconv.FormatInt(workspaceID, 10)
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: defaultNamespaceSettingTemplate(""),
	}, acceptance.Step{
		Template: defaultNamespaceSettingTemplate(fmt.Sprintf(`
			provider_config {
				workspace_id = "%s"
			}
		`, workspaceIDStr)),
	}, acceptance.Step{
		Template: defaultNamespaceSettingTemplate(`
			provider_config {
				workspace_id = "123"
			}
		`),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PostApplyPreRefresh: []plancheck.PlanCheck{
				plancheck.ExpectResourceAction("databricks_default_namespace_setting.this", plancheck.ResourceActionDestroyBeforeCreate),
			},
		},
		PlanOnly:           true,
		ExpectNonEmptyPlan: true,
	})
}

func TestAccDefaultNamespaceSetting_ProviderConfig_Remove(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	workspaceID, err := w.CurrentWorkspaceID(ctx)
	require.NoError(t, err)
	workspaceIDStr := strconv.FormatInt(workspaceID, 10)
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: defaultNamespaceSettingTemplate(""),
	}, acceptance.Step{
		Template: defaultNamespaceSettingTemplate(fmt.Sprintf(`
			provider_config {
				workspace_id = "%s"
			}
		`, workspaceIDStr)),
	}, acceptance.Step{
		Template: defaultNamespaceSettingTemplate(""),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				plancheck.ExpectResourceAction("databricks_default_namespace_setting.this", plancheck.ResourceActionUpdate),
			},
		},
	})
}

func TestAccDefaultNamespaceSetting(t *testing.T) {
	template := `
	resource "databricks_default_namespace_setting" "this" {
		namespace {
			value = "namespace_value"
		}
	}
	`
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: template,
		Check: acceptance.ResourceCheckWithState("databricks_default_namespace_setting.this",
			func(ctx context.Context, client *common.DatabricksClient, state *terraform.InstanceState) error {
				ctx = context.WithValue(ctx, common.Api, common.API_2_1)
				w, err := client.WorkspaceClient()
				require.NoError(t, err)
				etag := state.Attributes["etag"]
				require.NotEmpty(t, etag)
				res, err := w.Settings.DefaultNamespace().Get(ctx, settings.GetDefaultNamespaceSettingRequest{
					Etag: etag,
				})
				require.NoError(t, err)
				// Check that the resource has been created and that it has the correct value.
				assert.Equal(t, res.Namespace.Value, "namespace_value")
				return nil
			}),
	},
		acceptance.Step{
			Template: template,
			Destroy:  true,
			Check: acceptance.ResourceCheck("databricks_default_namespace_setting.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
				ctx = context.WithValue(ctx, common.Api, common.API_2_1)
				w, err := client.WorkspaceClient()
				require.NoError(t, err)
				// Terraform Check returns the latest resource status before it is destroyed, which has an outdated eTag.
				// We are making an update call to get the correct eTag in the response error.
				_, err = w.Settings.DefaultNamespace().Update(ctx, settings.UpdateDefaultNamespaceSettingRequest{
					AllowMissing: true,
					Setting: settings.DefaultNamespaceSetting{
						Namespace: settings.StringMessage{
							Value: "this_call_should_fail",
						},
					},
					FieldMask: "namespace.value",
				})
				assert.Error(t, err)
				var aerr *apierr.APIError
				if !errors.As(err, &aerr) {
					assert.FailNow(t, "cannot parse error message %v", err)
				}
				etag := aerr.Details[0].Metadata["etag"]
				_, err = w.Settings.DefaultNamespace().Get(ctx, settings.GetDefaultNamespaceSettingRequest{
					Etag: etag,
				})
				if !errors.As(err, &aerr) {
					assert.FailNow(t, "cannot parse error message %v", err)
				}

				assert.Equal(t, aerr.ErrorCode, "RESOURCE_DOES_NOT_EXIST")
				return nil
			}),
		},
	)
}
