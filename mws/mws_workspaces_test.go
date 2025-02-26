package mws_test

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/internal/providers"
	"github.com/databricks/terraform-provider-databricks/internal/providers/sdkv2"
	"github.com/databricks/terraform-provider-databricks/tokens"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
)

func TestMwsAccWorkspaces(t *testing.T) {
	acceptance.AccountLevel(t, acceptance.Step{
		Template: `
		resource "databricks_mws_credentials" "this" {
			account_id       = "{env.DATABRICKS_ACCOUNT_ID}"
			credentials_name = "credentials-ws-{var.RANDOM}"
			role_arn         = "{env.TEST_CROSSACCOUNT_ARN}"
		}
		resource "databricks_mws_customer_managed_keys" "this" {
			account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
			aws_key_info {
				key_arn   = "{env.TEST_MANAGED_KMS_KEY_ARN}"
				key_alias = "{env.TEST_MANAGED_KMS_KEY_ALIAS}"
			}
			use_cases = ["MANAGED_SERVICES"]
		}
		resource "databricks_mws_storage_configurations" "this" {
			account_id                 = "{env.DATABRICKS_ACCOUNT_ID}"
			storage_configuration_name = "storage-ws-{var.RANDOM}"
			bucket_name                = "{env.TEST_ROOT_BUCKET}"
		}
		resource "databricks_mws_networks" "this" {
			account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
			network_name = "network-ws-{var.RANDOM}"
			vpc_id       = "{env.TEST_VPC_ID}"
			subnet_ids   = [
				"{env.TEST_SUBNET_PRIVATE}",
				"{env.TEST_SUBNET_PRIVATE2}",
			]
			security_group_ids = [
				"{env.TEST_SECURITY_GROUP}",
			]
		}
		resource "databricks_mws_workspaces" "this" {
			account_id      = "{env.DATABRICKS_ACCOUNT_ID}"
			workspace_name  = "terra-{var.RANDOM}"
			aws_region      = "{env.AWS_REGION}"
	
			network_id = databricks_mws_networks.this.network_id
			credentials_id = databricks_mws_credentials.this.credentials_id
			storage_configuration_id = databricks_mws_storage_configurations.this.storage_configuration_id
			managed_services_customer_managed_key_id = databricks_mws_customer_managed_keys.this.customer_managed_key_id

			custom_tags = {
				"randomkey" = "randomvalue"
			}

			token {
				comment = "Test {var.RANDOM}"
			}
		}`,
	})
}

func TestMwsAccWorkspacesTokenUpdate(t *testing.T) {
	acceptance.AccountLevel(t, acceptance.Step{
		Template: `
		resource "databricks_mws_credentials" "this" {
			account_id       = "{env.DATABRICKS_ACCOUNT_ID}"
			credentials_name = "credentials-ws-{var.RANDOM}"
			role_arn         = "{env.TEST_CROSSACCOUNT_ARN}"
		}
		resource "databricks_mws_customer_managed_keys" "this" {
			account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
			aws_key_info {
				key_arn   = "{env.TEST_MANAGED_KMS_KEY_ARN}"
				key_alias = "{env.TEST_MANAGED_KMS_KEY_ALIAS}"
			}
			use_cases = ["MANAGED_SERVICES"]
		}
		resource "databricks_mws_storage_configurations" "this" {
			account_id                 = "{env.DATABRICKS_ACCOUNT_ID}"
			storage_configuration_name = "storage-ws-{var.RANDOM}"
			bucket_name                = "{env.TEST_ROOT_BUCKET}"
		}
		resource "databricks_mws_networks" "this" {
			account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
			network_name = "network-ws-{var.RANDOM}"
			vpc_id       = "{env.TEST_VPC_ID}"
			subnet_ids   = [
				"{env.TEST_SUBNET_PRIVATE}",
				"{env.TEST_SUBNET_PRIVATE2}",
			]
			security_group_ids = [
				"{env.TEST_SECURITY_GROUP}",
			]
		}			
		resource "databricks_mws_workspaces" "this" {
			account_id      = "{env.DATABRICKS_ACCOUNT_ID}"
			workspace_name  = "terra-{var.RANDOM}"
			aws_region      = "{env.AWS_REGION}"
	
			network_id = databricks_mws_networks.this.network_id
			credentials_id = databricks_mws_credentials.this.credentials_id
			storage_configuration_id = databricks_mws_storage_configurations.this.storage_configuration_id
			managed_services_customer_managed_key_id = databricks_mws_customer_managed_keys.this.customer_managed_key_id

			token {
				comment = "test foo"
			}
		}`,
		Check: acceptance.ResourceCheckWithState("databricks_mws_workspaces.this",
			func(ctx context.Context, client *common.DatabricksClient, state *terraform.InstanceState) error {
				workspaceUrl, ok := state.Attributes["workspace_url"]
				assert.True(t, ok, "workspace_url is absent from databricks_mws_workspaces instance state")

				workspaceClient, err := client.ClientForHost(ctx, workspaceUrl)
				assert.NoError(t, err)

				tokensAPI := tokens.NewTokensAPI(ctx, workspaceClient)
				tokens, err := tokensAPI.List()
				assert.NoError(t, err)

				foundFoo := false
				foundBar := false
				for _, token := range tokens {
					if token.Comment == "test foo" {
						foundFoo = true
					}
					if token.Comment == "test bar" {
						foundBar = true
					}
				}
				assert.True(t, foundFoo)
				assert.False(t, foundBar)
				return nil
			}),
	},
		acceptance.Step{
			Template: `
		resource "databricks_mws_credentials" "this" {
			account_id       = "{env.DATABRICKS_ACCOUNT_ID}"
			credentials_name = "credentials-ws-{var.RANDOM}"
			role_arn         = "{env.TEST_CROSSACCOUNT_ARN}"
		}
		resource "databricks_mws_customer_managed_keys" "this" {
			account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
			aws_key_info {
				key_arn   = "{env.TEST_MANAGED_KMS_KEY_ARN}"
				key_alias = "{env.TEST_MANAGED_KMS_KEY_ALIAS}"
			}
			use_cases = ["MANAGED_SERVICES"]
		}
		resource "databricks_mws_storage_configurations" "this" {
			account_id                 = "{env.DATABRICKS_ACCOUNT_ID}"
			storage_configuration_name = "storage-ws-{var.RANDOM}"
			bucket_name                = "{env.TEST_ROOT_BUCKET}"
		}
		resource "databricks_mws_networks" "this" {
			account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
			network_name = "network-ws-{var.RANDOM}"
			vpc_id       = "{env.TEST_VPC_ID}"
			subnet_ids   = [
				"{env.TEST_SUBNET_PRIVATE}",
				"{env.TEST_SUBNET_PRIVATE2}",
			]
			security_group_ids = [
				"{env.TEST_SECURITY_GROUP}",
			]
		}			
		resource "databricks_mws_workspaces" "this" {
			account_id      = "{env.DATABRICKS_ACCOUNT_ID}"
			workspace_name  = "terra-{var.RANDOM}"
			aws_region      = "{env.AWS_REGION}"
	
			network_id = databricks_mws_networks.this.network_id
			credentials_id = databricks_mws_credentials.this.credentials_id
			storage_configuration_id = databricks_mws_storage_configurations.this.storage_configuration_id
			managed_services_customer_managed_key_id = databricks_mws_customer_managed_keys.this.customer_managed_key_id

			token {
				comment = "test bar"
			}
		}`,
			Check: acceptance.ResourceCheckWithState("databricks_mws_workspaces.this",
				func(ctx context.Context, client *common.DatabricksClient, state *terraform.InstanceState) error {
					workspaceUrl, ok := state.Attributes["workspace_url"]
					assert.True(t, ok, "workspace_url is absent from databricks_mws_workspaces instance state")

					workspaceClient, err := client.ClientForHost(ctx, workspaceUrl)
					assert.NoError(t, err)

					tokensAPI := tokens.NewTokensAPI(ctx, workspaceClient)
					tokens, err := tokensAPI.List()
					assert.NoError(t, err)

					foundFoo := false
					foundBar := false
					for _, token := range tokens {
						if token.Comment == "test foo" {
							foundFoo = true
						}
						if token.Comment == "test bar" {
							foundBar = true
						}
					}
					assert.False(t, foundFoo)
					assert.True(t, foundBar)
					return nil
				}),
		})
}

func TestMwsAccGcpWorkspaces(t *testing.T) {
	acceptance.AccountLevel(t, acceptance.Step{
		Template: `
		resource "databricks_mws_workspaces" "this" {
			account_id      = "{env.DATABRICKS_ACCOUNT_ID}"
			workspace_name  = "{env.TEST_PREFIX}-{var.RANDOM}"
			location        = "{env.GOOGLE_REGION}"
	
			cloud_resource_container {
				gcp {
					project_id = "{env.GOOGLE_PROJECT}"
				}
			}
		}`,
	})
}

func TestMwsAccGcpByovpcWorkspaces(t *testing.T) {
	commonResources := `
		resource "databricks_mws_networks" "this" {
			account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
			network_name = "psc-network-{var.STICKY_RANDOM}"
			gcp_network_info {
				network_project_id = "{env.GOOGLE_PROJECT}"
				vpc_id = "{env.TEST_VPC_ID}"
				subnet_id = "{env.TEST_SUBNET_ID}"
				subnet_region = "{env.GOOGLE_REGION}"
				pod_ip_range_name = "pods"
				service_ip_range_name = "service"
			}
		}
		`
	acceptance.AccountLevel(t, acceptance.Step{
		Template: commonResources + `
		resource "databricks_mws_workspaces" "this" {
			account_id      = "{env.DATABRICKS_ACCOUNT_ID}"
			workspace_name  = "psc-test-{var.STICKY_RANDOM}"
			location        = "{env.GOOGLE_REGION}"
	
			cloud_resource_container {
				gcp {
					project_id = "{env.GOOGLE_PROJECT}"
				}
			}
						
			gke_config {
				connectivity_type = "PRIVATE_NODE_PUBLIC_MASTER"
				master_ip_range = "10.3.0.0/28"
			}
            
			network_id = databricks_mws_networks.this.network_id
		}`,
	}, acceptance.Step{
		// Changing the workspace name recreates the workspace.
		Template: commonResources + `
		resource "databricks_mws_workspaces" "this" {
			account_id      = "{env.DATABRICKS_ACCOUNT_ID}"
			workspace_name  = "psc-test-new-{var.STICKY_RANDOM}"
			location        = "{env.GOOGLE_REGION}"
	
			cloud_resource_container {
				gcp {
					project_id = "{env.GOOGLE_PROJECT}"
				}
			}
						
			gke_config {
				connectivity_type = "PRIVATE_NODE_PUBLIC_MASTER"
				master_ip_range = "10.3.0.0/28"
			}
            
			network_id = databricks_mws_networks.this.network_id
		}`,
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				checkResourceActions{"databricks_mws_workspaces.this", []tfjson.Action{tfjson.ActionDelete, tfjson.ActionCreate}},
			},
		},
	}, acceptance.Step{
		// Removing gke_config is a no-op because of suppress_diff.
		Template: commonResources + `
		resource "databricks_mws_workspaces" "this" {
			account_id      = "{env.DATABRICKS_ACCOUNT_ID}"
			workspace_name  = "psc-test-new-{var.STICKY_RANDOM}"
			location        = "{env.GOOGLE_REGION}"
	
			cloud_resource_container {
				gcp {
					project_id = "{env.GOOGLE_PROJECT}"
				}
			}
            
			network_id = databricks_mws_networks.this.network_id
		}`,
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{
				checkResourceActions{"databricks_mws_workspaces.this", []tfjson.Action{tfjson.ActionNoop}},
			},
		},
		Check: func(s *terraform.State) error {
			r := s.RootModule().Resources["databricks_mws_workspaces.this"].Primary
			assert.Empty(t, r.Attributes["gke_config"])
			return nil
		},
	})
}

func TestMwsAccGcpPscWorkspaces(t *testing.T) {
	t.Skip()
	// private access settings are not enabled in our new E2 account.
	acceptance.AccountLevel(t, acceptance.Step{
		Template: `
		resource "databricks_mws_networks" "this" {
			account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
			network_name = "{env.TEST_PREFIX}-network-{var.RANDOM}"
			gcp_network_info {
				network_project_id = "{env.GOOGLE_PROJECT}"
				vpc_id = "{env.VPC_NETWORK_ID}"
				subnet_id = "{env.SUBNET_ID}"
				subnet_region = "{env.GOOGLE_REGION}"
				pod_ip_range_name = "{env.POD_IP_RANGE_NAME}"
				service_ip_range_name = "{env.SVC_IP_RANGE_NAME}"
			}
		}

		resource "databricks_mws_private_access_settings" "this" {
			account_id = "{env.DATABRICKS_ACCOUNT_ID}"
			private_access_settings_name = "tf-pas-{var.RANDOM}"
			region = "{env.GOOGLE_REGION}"
			public_access_enabled = true
			private_access_level = "ACCOUNT"
		}
		
		resource "databricks_mws_workspaces" "this" {
			account_id      = "{env.DATABRICKS_ACCOUNT_ID}"
			workspace_name  = "{env.TEST_PREFIX}-{var.RANDOM}"
			location        = "{env.GOOGLE_REGION}"
	
			cloud_resource_container {
				gcp {
					project_id = "{env.GOOGLE_PROJECT}"
				}
			}
            
            private_access_settings_id = databricks_mws_private_access_settings.this.private_access_settings_id
			network_id = databricks_mws_networks.this.network_id
			
			gke_config {
				connectivity_type = "PRIVATE_NODE_PUBLIC_MASTER"
				master_ip_range = "10.3.0.0/28"
			}
		}`,
	})
}

func TestMwsAccAwsChangeToServicePrincipal(t *testing.T) {
	if !acceptance.IsAws(t) {
		acceptance.Skipf(t)("TestMwsAccAwsChangeToServicePrincipal should only run on AWS")
	}
	workspaceTemplate := func(tokenBlock string) string {
		return `
		resource "databricks_mws_credentials" "this" {
			account_id       = "{env.DATABRICKS_ACCOUNT_ID}"
			credentials_name = "credentials-ws-{var.STICKY_RANDOM}"
			role_arn         = "{env.TEST_CROSSACCOUNT_ARN}"
		}
		resource "databricks_mws_customer_managed_keys" "this" {
			account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
			aws_key_info {
				key_arn   = "{env.TEST_MANAGED_KMS_KEY_ARN}"
				key_alias = "{env.TEST_MANAGED_KMS_KEY_ALIAS}"
			}
			use_cases = ["MANAGED_SERVICES"]
		}
		resource "databricks_mws_storage_configurations" "this" {
			account_id                 = "{env.DATABRICKS_ACCOUNT_ID}"
			storage_configuration_name = "storage-ws-{var.STICKY_RANDOM}"
			bucket_name                = "{env.TEST_ROOT_BUCKET}"
		}
		resource "databricks_mws_networks" "this" {
			account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
			network_name = "network-ws-{var.STICKY_RANDOM}"
			vpc_id       = "{env.TEST_VPC_ID}"
			subnet_ids   = [
				"{env.TEST_SUBNET_PRIVATE}",
				"{env.TEST_SUBNET_PRIVATE2}",
			]
			security_group_ids = [
				"{env.TEST_SECURITY_GROUP}",
			]
		}
		resource "databricks_mws_workspaces" "this" {
			account_id      = "{env.DATABRICKS_ACCOUNT_ID}"
			workspace_name  = "terra-{var.STICKY_RANDOM}"
			aws_region      = "{env.AWS_REGION}"

			network_id = databricks_mws_networks.this.network_id
			credentials_id = databricks_mws_credentials.this.credentials_id
			storage_configuration_id = databricks_mws_storage_configurations.this.storage_configuration_id
			managed_services_customer_managed_key_id = databricks_mws_customer_managed_keys.this.customer_managed_key_id

			custom_tags = {
				"randomkey" = "randomvalue"
			}

			` + tokenBlock + `
		}
		`
	}
	servicePrincipal := `
		resource "databricks_service_principal" "this" {
		    display_name = "tf-new-sp-{var.STICKY_RANDOM}"
			disable_as_user_deletion = false
		}
		resource "databricks_service_principal_role" "this" {
		    service_principal_id = databricks_service_principal.this.id
			role = "account_admin"
		}
		resource "databricks_service_principal_secret" "this" {
		    service_principal_id = databricks_service_principal.this.id
		}
		`

	var pr *schema.Provider
	providerFactory := map[string]func() (tfprotov6.ProviderServer, error){
		"databricks": func() (tfprotov6.ProviderServer, error) {
			return providers.GetProviderServer(context.Background(), providers.WithSdkV2Provider(pr))
		},
	}
	acceptance.AccountLevel(t, acceptance.Step{
		Template: workspaceTemplate(`token { comment = "Test {var.STICKY_RANDOM}" }`) + servicePrincipal,
		Check: func(s *terraform.State) error {
			spId := s.RootModule().Resources["databricks_service_principal.this"].Primary.ID
			spAppId := s.RootModule().Resources["databricks_service_principal.this"].Primary.Attributes["application_id"]
			spSecret := s.RootModule().Resources["databricks_service_principal_secret.this"].Primary.Attributes["secret"]
			pr = sdkv2.DatabricksProvider()
			rd := schema.TestResourceDataRaw(t, pr.Schema, map[string]interface{}{
				"client_id":     spAppId,
				"client_secret": spSecret,
			})
			fmt.Printf("client_id: %s, client_secret: %s\n", spAppId, spSecret)
			pr.ConfigureContextFunc = func(ctx context.Context, c *schema.ResourceData) (interface{}, diag.Diagnostics) {
				return sdkv2.ConfigureDatabricksClient(ctx, rd)
			}
			logger.DefaultLogger = &logger.SimpleLogger{
				Level: logger.LevelDebug,
			}
			// wait until SP exists
			for i := 100; i >= 0; i-- {
				a := databricks.Must(databricks.NewAccountClient(&databricks.Config{
					ClientID:     spAppId,
					ClientSecret: spSecret,
				}))
				_, err := a.ServicePrincipals.GetById(context.Background(), spId)
				if err == nil {
					break
				}
				if i == 0 {
					return errors.New("service principal not found")
				}
				if errors.Is(err, databricks.ErrUnauthenticated) {
					fmt.Println("waiting for SP to be ready (sleeping 5 seconds, trying ", i, " more times)")
					time.Sleep(5 * time.Second)
				}
			}
			return nil
		},
	}, acceptance.Step{
		// Tolerate existing token
		Template:                 workspaceTemplate(`token { comment = "Test {var.STICKY_RANDOM}" }`) + servicePrincipal,
		ProtoV6ProviderFactories: providerFactory,
	}, acceptance.Step{
		// Allow the token to be removed
		Template:                 workspaceTemplate(``) + servicePrincipal,
		ProtoV6ProviderFactories: providerFactory,
	}, acceptance.Step{
		// Fail when adding the token back
		Template:                 workspaceTemplate(`token { comment = "Test {var.STICKY_RANDOM}" }`) + servicePrincipal,
		ProtoV6ProviderFactories: providerFactory,
		ExpectError:              regexp.MustCompile(`cannot create token: the principal used by Databricks \(client ID .*\) is not authorized to create a token in this workspace`),
	}, acceptance.Step{
		// Use the original provider for a final step to clean up the newly created service principal
		Template: workspaceTemplate(``) + servicePrincipal,
	})
}
