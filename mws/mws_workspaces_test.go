package mws_test

import (
	"context"
	"errors"
	"fmt"
	"os"
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
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/option"
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

func TestMwsAccWorkspaces_Serverless(t *testing.T) {
	acceptance.LoadAccountEnv(t)
	if !acceptance.IsAws(t) {
		acceptance.Skipf(t)("TestMwsAccWorkspaces_Serverless is currently only supported on AWS")
	}
	acceptance.AccountLevel(t, acceptance.Step{
		Template: `
		resource "databricks_mws_workspaces" "this" {
			account_id      = "{env.DATABRICKS_ACCOUNT_ID}"
			workspace_name  = "terra-{var.RANDOM}"
			aws_region      = "{env.AWS_REGION}"
			compute_mode    = "SERVERLESS"
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

func TestMwsAccGcpWorkspacesProvisioningToRunning(t *testing.T) {
	acceptance.LoadAccountEnv(t)
	if !acceptance.IsGcp(t) {
		acceptance.Skipf(t)("TestMwsAccGcpWorkspacesProvisioningToRunning is currently only supported on GCP")
	}

	var serviceAccountEmail string

	acceptance.AccountLevel(t, acceptance.Step{
		Template: `
		resource "databricks_mws_workspaces" "this" {
			account_id                = "{env.DATABRICKS_ACCOUNT_ID}"
			workspace_name            = "{env.TEST_PREFIX}-{var.STICKY_RANDOM}"
			location                  = "{env.GOOGLE_REGION}"
			expected_workspace_status = "PROVISIONING"

			cloud_resource_container {
				gcp {
					project_id = "{env.GOOGLE_PROJECT}"
				}
			}
		}`,
		Check: func(s *terraform.State) error {
			rs, ok := s.RootModule().Resources["databricks_mws_workspaces.this"]
			if !ok {
				return fmt.Errorf("databricks_mws_workspaces.this not found")
			}
			if rs.Primary.Attributes["workspace_id"] == "" {
				return fmt.Errorf("workspace_id is empty")
			}

			// Capture service account email for use in PreConfig
			serviceAccountEmail = rs.Primary.Attributes["gcp_workspace_sa"]
			if serviceAccountEmail == "" {
				return fmt.Errorf("gcp_workspace_sa is empty")
			}

			expectedStatus := "PROVISIONING"
			if status := rs.Primary.Attributes["workspace_status"]; status != expectedStatus {
				return fmt.Errorf("expected workspace_status to be %s, got %s", expectedStatus, status)
			}
			return nil
		},
	}, acceptance.Step{
		PreConfig: func() {
			grantGcpWorkspacePermissions(t, serviceAccountEmail)
		},
		Template: `
		resource "databricks_mws_networks" "this" {
			account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
			network_name = "network-{var.STICKY_RANDOM}"
			gcp_network_info {
				network_project_id = "{env.GOOGLE_PROJECT}"
				vpc_id = "{env.TEST_VPC_ID}"
				subnet_id = "{env.TEST_SUBNET_ID}"
				subnet_region = "{env.GOOGLE_REGION}"
			}
		}
		resource "databricks_mws_workspaces" "this" {
			account_id                = "{env.DATABRICKS_ACCOUNT_ID}"
			workspace_name            = "{env.TEST_PREFIX}-{var.STICKY_RANDOM}"
			location                  = "{env.GOOGLE_REGION}"
			expected_workspace_status = "RUNNING"

			cloud_resource_container {
				gcp {
					project_id = "{env.GOOGLE_PROJECT}"
				}
			}

			network_id = databricks_mws_networks.this.network_id
		}`,
		Check: func(s *terraform.State) error {
			rs, ok := s.RootModule().Resources["databricks_mws_workspaces.this"]
			if !ok {
				return fmt.Errorf("databricks_mws_workspaces.this not found")
			}
			if rs.Primary.Attributes["workspace_id"] == "" {
				return fmt.Errorf("workspace_id is empty")
			}

			expectedStatus := "RUNNING"
			if status := rs.Primary.Attributes["workspace_status"]; status != expectedStatus {
				return fmt.Errorf("expected workspace_status to be %s, got %s", expectedStatus, status)
			}
			return nil
		},
	})
}

func TestMwsAccGcpWorkspacesUnsetExpectedState(t *testing.T) {
	acceptance.LoadAccountEnv(t)
	if !acceptance.IsGcp(t) {
		acceptance.Skipf(t)("TestMwsAccGcpWorkspacesUnsetExpectedState is currently only supported on GCP")
	}

	var serviceAccountEmail string

	acceptance.AccountLevel(t, acceptance.Step{
		Template: `
		resource "databricks_mws_workspaces" "this" {
			account_id                = "{env.DATABRICKS_ACCOUNT_ID}"
			workspace_name            = "{env.TEST_PREFIX}-{var.STICKY_RANDOM}"
			location                  = "{env.GOOGLE_REGION}"
			expected_workspace_status = "PROVISIONING"

			cloud_resource_container {
				gcp {
					project_id = "{env.GOOGLE_PROJECT}"
				}
			}
		}`,
		Check: func(s *terraform.State) error {
			rs, ok := s.RootModule().Resources["databricks_mws_workspaces.this"]
			if !ok {
				return fmt.Errorf("databricks_mws_workspaces.this not found")
			}
			if rs.Primary.Attributes["workspace_id"] == "" {
				return fmt.Errorf("workspace_id is empty")
			}

			// Capture service account email for use in PreConfig
			serviceAccountEmail = rs.Primary.Attributes["gcp_workspace_sa"]
			if serviceAccountEmail == "" {
				return fmt.Errorf("gcp_workspace_sa is empty")
			}

			expectedStatus := "PROVISIONING"
			if status := rs.Primary.Attributes["workspace_status"]; status != expectedStatus {
				return fmt.Errorf("expected workspace_status to be %s, got %s", expectedStatus, status)
			}
			return nil
		},
	}, acceptance.Step{
		PreConfig: func() {
			grantGcpWorkspacePermissions(t, serviceAccountEmail)
		},
		Template: `
		resource "databricks_mws_networks" "this" {
			account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
			network_name = "network-{var.STICKY_RANDOM}"
			gcp_network_info {
				network_project_id = "{env.GOOGLE_PROJECT}"
				vpc_id = "{env.TEST_VPC_ID}"
				subnet_id = "{env.TEST_SUBNET_ID}"
				subnet_region = "{env.GOOGLE_REGION}"
			}
		}
		resource "databricks_mws_workspaces" "this" {
			account_id      = "{env.DATABRICKS_ACCOUNT_ID}"
			workspace_name  = "{env.TEST_PREFIX}-{var.STICKY_RANDOM}"
			location        = "{env.GOOGLE_REGION}"

			cloud_resource_container {
				gcp {
					project_id = "{env.GOOGLE_PROJECT}"
				}
			}

			network_id = databricks_mws_networks.this.network_id
		}`,
		Check: func(s *terraform.State) error {
			rs, ok := s.RootModule().Resources["databricks_mws_workspaces.this"]
			if !ok {
				return fmt.Errorf("databricks_mws_workspaces.this not found")
			}
			if rs.Primary.Attributes["workspace_id"] == "" {
				return fmt.Errorf("workspace_id is empty")
			}

			expectedStatus := "RUNNING"
			if status := rs.Primary.Attributes["workspace_status"]; status != expectedStatus {
				return fmt.Errorf("expected workspace_status to be %s, got %s", expectedStatus, status)
			}
			return nil
		},
	})
}

func TestMwsAccGcpByovpcWorkspaces(t *testing.T) {
	acceptance.AccountLevel(t, acceptance.Step{
		Template: `
		resource "databricks_mws_networks" "this" {
			account_id   = "{env.DATABRICKS_ACCOUNT_ID}"
			network_name = "psc-network-{var.STICKY_RANDOM}"
			gcp_network_info {
				network_project_id = "{env.GOOGLE_PROJECT}"
				vpc_id = "{env.TEST_VPC_ID}"
				subnet_id = "{env.TEST_SUBNET_ID}"
				subnet_region = "{env.GOOGLE_REGION}"
			}
		}
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
				return sdkv2.ConfigureDatabricksClient(ctx, rd, acceptance.DefaultConfigCustomizer)
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

// grantGcpWorkspacePermissions is a helper function for PreConfig that grants GCP permissions
// to a workspace service account using a pre-created custom role
func grantGcpWorkspacePermissions(t *testing.T, serviceAccountEmail string) {
	ctx := context.Background()

	// Grant GCP permissions to the workspace service account
	projectID := os.Getenv("GOOGLE_PROJECT")
	if projectID == "" {
		t.Fatal("GOOGLE_PROJECT environment variable is not set")
	}

	// Get the pre-created custom role ID from environment variable
	roleID := os.Getenv("TEST_LEAST_PRIVILEGED_WORKSPACE_ROLE_ID")
	if roleID == "" {
		// TODO: fail here if roleID is not set
		roleID = fmt.Sprintf("projects/%s/roles/databricksLeastPrivilegedWorkspaceRole", projectID)
	}

	t.Logf("Assigning role %s to service account: %s on project: %s", roleID, serviceAccountEmail, projectID)
	if err := assignGcpRoleToServiceAccount(ctx, projectID, serviceAccountEmail, roleID); err != nil {
		t.Fatalf("Failed to assign GCP role: %v", err)
	}

	// Clean up permissions after the test completes
	t.Cleanup(func() {
		cleanupCtx := context.Background()
		t.Logf("Removing role %s from service account: %s on project: %s", roleID, serviceAccountEmail, projectID)
		if err := removeGcpRoleFromServiceAccount(cleanupCtx, projectID, serviceAccountEmail, roleID); err != nil {
			t.Logf("Warning: Failed to remove GCP role: %v", err)
		}
	})
}

func getCloudResourceManagerClient(ctx context.Context) (*cloudresourcemanager.Service, error) {
	options := []option.ClientOption{
		option.WithScopes(cloudresourcemanager.CloudPlatformScope),
	}
	// In integration tests, the GOOGLE_CREDENTIALS environment variable specifies the credentials
	credentialsJSON := os.Getenv("GOOGLE_CREDENTIALS")
	if credentialsJSON != "" {
		options = append(options, option.WithCredentialsJSON([]byte(credentialsJSON)))
	}
	return cloudresourcemanager.NewService(ctx, options...)
}

// assignGcpRoleToServiceAccount assigns a custom role to a service account on a GCP project
func assignGcpRoleToServiceAccount(ctx context.Context, projectID, serviceAccountEmail, roleID string) error {
	crmService, err := getCloudResourceManagerClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to create Cloud Resource Manager client: %w", err)
	}

	// Get the current IAM policy
	policy, err := crmService.Projects.GetIamPolicy(projectID, &cloudresourcemanager.GetIamPolicyRequest{}).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("failed to get IAM policy: %w", err)
	}

	member := fmt.Sprintf("serviceAccount:%s", serviceAccountEmail)

	// Check if binding already exists for this custom role
	memberExists := false
	bindingIndex := -1

	for i, binding := range policy.Bindings {
		if binding.Role == roleID {
			bindingIndex = i
			// Check if member already exists
			for _, m := range binding.Members {
				if m == member {
					memberExists = true
					break
				}
			}
			break
		}
	}

	if !memberExists {
		if bindingIndex >= 0 {
			// Binding exists, add member to it
			policy.Bindings[bindingIndex].Members = append(policy.Bindings[bindingIndex].Members, member)
		} else {
			// Binding doesn't exist, create it
			newBinding := &cloudresourcemanager.Binding{
				Role:    roleID,
				Members: []string{member},
			}
			policy.Bindings = append(policy.Bindings, newBinding)
		}

		// Set the updated IAM policy
		_, err = crmService.Projects.SetIamPolicy(projectID, &cloudresourcemanager.SetIamPolicyRequest{
			Policy: policy,
		}).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("failed to set IAM policy: %w", err)
		}
	}

	return nil
}

// removeGcpRoleFromServiceAccount removes a custom role from a service account on a GCP project
func removeGcpRoleFromServiceAccount(ctx context.Context, projectID, serviceAccountEmail, roleName string) error {
	crmService, err := getCloudResourceManagerClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to create Cloud Resource Manager client: %w", err)
	}

	// Get the current IAM policy
	policy, err := crmService.Projects.GetIamPolicy(projectID, &cloudresourcemanager.GetIamPolicyRequest{}).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("failed to get IAM policy: %w", err)
	}

	member := fmt.Sprintf("serviceAccount:%s", serviceAccountEmail)

	// Remove the member from the custom role binding
	for i, binding := range policy.Bindings {
		if binding.Role == roleName {
			newMembers := []string{}
			for _, m := range binding.Members {
				if m != member {
					newMembers = append(newMembers, m)
				}
			}
			policy.Bindings[i].Members = newMembers
		}
	}

	// Remove empty bindings
	newBindings := []*cloudresourcemanager.Binding{}
	for _, binding := range policy.Bindings {
		if len(binding.Members) > 0 {
			newBindings = append(newBindings, binding)
		}
	}
	policy.Bindings = newBindings

	// Set the updated IAM policy
	_, err = crmService.Projects.SetIamPolicy(projectID, &cloudresourcemanager.SetIamPolicyRequest{
		Policy: policy,
	}).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("failed to set IAM policy: %w", err)
	}

	return nil
}
