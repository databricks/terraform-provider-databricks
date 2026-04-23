package acceptance

import (
	"context"
	"regexp"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/terraform-provider-databricks/internal/providers"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw"
	"github.com/databricks/terraform-provider-databricks/internal/providers/sdkv2"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// unifiedHostMockProviderFactories returns provider factories whose resolved
// host type is UnifiedHost. The customizer resets the config so EnsureResolved
// runs again with an in-memory HostMetadataResolver that always returns
// UnifiedHost — no network call, no env dependency.
func unifiedHostMockProviderFactories() map[string]func() (tfprotov6.ProviderServer, error) {
	customizer := func(cfg *config.Config) error {
		*cfg = config.Config{
			Host:  "https://unifiedhost.databricks.com",
			Token: "test-token",
			HostMetadataResolver: func(ctx context.Context, _ string) (*config.HostMetadata, error) {
				return &config.HostMetadata{HostType: config.UnifiedHost}, nil
			},
		}
		return cfg.EnsureResolved()
	}
	return map[string]func() (tfprotov6.ProviderServer, error){
		"databricks": func() (tfprotov6.ProviderServer, error) {
			ctx := context.Background()
			sdkPluginProvider := sdkv2.DatabricksProvider(
				sdkv2.WithConfigCustomizer(customizer),
			)
			pluginFrameworkProvider := pluginfw.GetDatabricksProviderPluginFramework(
				pluginfw.WithConfigCustomizer(customizer),
			)
			return providers.GetProviderServer(ctx,
				providers.WithSdkV2Provider(sdkPluginProvider),
				providers.WithPluginFrameworkProvider(pluginFrameworkProvider),
			)
		},
	}
}

// dualResourceUnifiedHostPlanTest plans the given HCL against a UnifiedHost-
// mocked provider and asserts it fails with the missing-api error.
func dualResourceUnifiedHostPlanTest(t *testing.T, hcl string) {
	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: unifiedHostMockProviderFactories(),
		Steps: []resource.TestStep{{
			Config:      hcl,
			PlanOnly:    true,
			ExpectError: regexp.MustCompile(`please set api to account or workspace`),
		}},
	})
}

func TestAccDualResource_UnifiedHost_User_MissingApi(t *testing.T) {
	dualResourceUnifiedHostPlanTest(t, `
		resource "databricks_user" "this" {
			user_name = "test@example.com"
		}
	`)
}

func TestAccDualResource_UnifiedHost_ServicePrincipal_MissingApi(t *testing.T) {
	dualResourceUnifiedHostPlanTest(t, `
		resource "databricks_service_principal" "this" {
			display_name = "test-sp"
		}
	`)
}

func TestAccDualResource_UnifiedHost_Group_MissingApi(t *testing.T) {
	dualResourceUnifiedHostPlanTest(t, `
		resource "databricks_group" "this" {
			display_name = "test-group"
		}
	`)
}

func TestAccDualResource_UnifiedHost_GroupRole_MissingApi(t *testing.T) {
	dualResourceUnifiedHostPlanTest(t, `
		resource "databricks_group_role" "this" {
			group_id = "123"
			role     = "test-role"
		}
	`)
}

func TestAccDualResource_UnifiedHost_GroupMember_MissingApi(t *testing.T) {
	dualResourceUnifiedHostPlanTest(t, `
		resource "databricks_group_member" "this" {
			group_id  = "123"
			member_id = "456"
		}
	`)
}

func TestAccDualResource_UnifiedHost_UserRole_MissingApi(t *testing.T) {
	dualResourceUnifiedHostPlanTest(t, `
		resource "databricks_user_role" "this" {
			user_id = "123"
			role    = "test-role"
		}
	`)
}

func TestAccDualResource_UnifiedHost_ServicePrincipalRole_MissingApi(t *testing.T) {
	dualResourceUnifiedHostPlanTest(t, `
		resource "databricks_service_principal_role" "this" {
			service_principal_id = "123"
			role                 = "test-role"
		}
	`)
}

func TestAccDualResource_UnifiedHost_UserInstanceProfile_MissingApi(t *testing.T) {
	dualResourceUnifiedHostPlanTest(t, `
		resource "databricks_user_instance_profile" "this" {
			user_id             = "123"
			instance_profile_id = "arn:aws:iam::123456789012:instance-profile/test"
		}
	`)
}

func TestAccDualResource_UnifiedHost_GroupInstanceProfile_MissingApi(t *testing.T) {
	dualResourceUnifiedHostPlanTest(t, `
		resource "databricks_group_instance_profile" "this" {
			group_id            = "123"
			instance_profile_id = "arn:aws:iam::123456789012:instance-profile/test"
		}
	`)
}

func TestAccDualResource_UnifiedHost_ServicePrincipalSecret_MissingApi(t *testing.T) {
	dualResourceUnifiedHostPlanTest(t, `
		resource "databricks_service_principal_secret" "this" {
			service_principal_id = "123"
		}
	`)
}

func TestAccDualResource_UnifiedHost_AccessControlRuleSet_MissingApi(t *testing.T) {
	dualResourceUnifiedHostPlanTest(t, `
		resource "databricks_access_control_rule_set" "this" {
			name = "accounts/abc/servicePrincipals/123/ruleSets/default"
			grant_rules {
				principals = ["groups/admins"]
				role       = "roles/servicePrincipal.user"
			}
		}
	`)
}

func TestAccDualResource_UnifiedHost_Metastore_MissingApi(t *testing.T) {
	dualResourceUnifiedHostPlanTest(t, `
		resource "databricks_metastore" "this" {
			name         = "test-metastore"
			storage_root = "s3://test-bucket/"
			region       = "us-east-1"
		}
	`)
}

func TestAccDualResource_UnifiedHost_StorageCredential_MissingApi(t *testing.T) {
	dualResourceUnifiedHostPlanTest(t, `
		resource "databricks_storage_credential" "this" {
			name = "test-credential"
			aws_iam_role {
				role_arn = "arn:aws:iam::123456789012:role/test"
			}
		}
	`)
}

func TestAccDualResource_UnifiedHost_MetastoreDataAccess_MissingApi(t *testing.T) {
	dualResourceUnifiedHostPlanTest(t, `
		resource "databricks_metastore_data_access" "this" {
			metastore_id = "11111111-1111-1111-1111-111111111111"
			name         = "test-data-access"
			aws_iam_role {
				role_arn = "arn:aws:iam::123456789012:role/test"
			}
		}
	`)
}

func TestAccDualResource_UnifiedHost_MetastoreAssignment_MissingApi(t *testing.T) {
	dualResourceUnifiedHostPlanTest(t, `
		resource "databricks_metastore_assignment" "this" {
			metastore_id = "11111111-1111-1111-1111-111111111111"
			workspace_id = 123456
		}
	`)
}

func TestAccDualDataSource_UnifiedHost_User_MissingApi(t *testing.T) {
	dualResourceUnifiedHostPlanTest(t, `
		data "databricks_user" "this" {
			user_name = "test@example.com"
		}
	`)
}

func TestAccDualDataSource_UnifiedHost_Group_MissingApi(t *testing.T) {
	dualResourceUnifiedHostPlanTest(t, `
		data "databricks_group" "this" {
			display_name = "test-group"
		}
	`)
}

func TestAccDualDataSource_UnifiedHost_CurrentConfig_MissingApi(t *testing.T) {
	dualResourceUnifiedHostPlanTest(t, `
		data "databricks_current_config" "this" {
			cloud = "aws"
		}
	`)
}
