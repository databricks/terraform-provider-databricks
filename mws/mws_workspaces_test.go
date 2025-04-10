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
	"github.com/databricks/databricks-sdk-go/service/provisioning"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/internal/providers"
	"github.com/databricks/terraform-provider-databricks/internal/providers/sdkv2"
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

type expectNotDestroyed struct {
	addr string
}

func ExpectNotDestroyed(addr string) expectNotDestroyed {
	return expectNotDestroyed{addr: addr}
}

func (e expectNotDestroyed) CheckPlan(ctx context.Context, req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
	for _, resource := range req.Plan.ResourceChanges {
		if resource.Address != e.addr {
			continue
		}
		actions := resource.Change.Actions
		if actions.DestroyBeforeCreate() || actions.CreateBeforeDestroy() || actions.Delete() {
			resp.Error = fmt.Errorf("resource %s is marked for destruction", e.addr)
			return
		}
	}
}

func TestMwsAccWorkspaces_TokenUpdate(t *testing.T) {
	tokenUpdateTemplate := func(token, customTags string) string {
		tokenBlock := ``
		if token != "" {
			tokenBlock = fmt.Sprintf(`
			token {
				%s
			}`, token)
		}
		customTagsBlock := ``
		if customTags != "" {
			customTagsBlock = fmt.Sprintf(`
			custom_tags = {
				%s
			}`, customTags)
		}
		return fmt.Sprintf(`
		resource "databricks_mws_credentials" "this" {
			account_id       = "{env.DATABRICKS_ACCOUNT_ID}"
			credentials_name = "credentials-ws-{var.STICKY_RANDOM}"
			role_arn         = "{env.TEST_CROSSACCOUNT_ARN}"
		}
		resource "databricks_mws_storage_configurations" "this" {
			account_id                 = "{env.DATABRICKS_ACCOUNT_ID}"
			storage_configuration_name = "storage-ws-{var.STICKY_RANDOM}"
			bucket_name                = "{env.TEST_ROOT_BUCKET}"
		}
		resource "databricks_mws_workspaces" "this" {
			account_id      = "{env.DATABRICKS_ACCOUNT_ID}"
			workspace_name  = "terra-{var.STICKY_RANDOM}"
			aws_region      = "{env.AWS_REGION}"
	
			credentials_id = databricks_mws_credentials.this.credentials_id
			storage_configuration_id = databricks_mws_storage_configurations.this.storage_configuration_id

			%s
			%s
		}`, tokenBlock, customTagsBlock)
	}

	checkWorkspace := func(f func(instanceState map[string]string, w *databricks.WorkspaceClient) error) func(*terraform.State) error {
		return func(s *terraform.State) error {
			state, ok := s.RootModule().Resources["databricks_mws_workspaces.this"]
			if !ok {
				return fmt.Errorf("resource not found in state")
			}
			a := databricks.Must(databricks.NewAccountClient())
			ctx := context.Background()
			workspace, err := a.Workspaces.Get(ctx, provisioning.GetWorkspaceRequest{WorkspaceId: common.MustInt64(state.Primary.Attributes["workspace_id"])})
			assert.NoError(t, err)

			w, err := a.GetWorkspaceClient(*workspace)
			assert.NoError(t, err)

			return f(state.Primary.Attributes, w)
		}
	}

	checkTokenExists := checkWorkspace(func(instanceState map[string]string, w *databricks.WorkspaceClient) error {
		tokenId := instanceState["token.0.token_id"]
		assert.NotEmpty(t, tokenId)
		tokens := w.Tokens.List(context.Background())
		ctx := context.Background()
		for tokens.HasNext(ctx) {
			token, err := tokens.Next(ctx)
			if err != nil {
				return fmt.Errorf("error fetching tokens: %w", err)
			}
			if token.TokenId == tokenId {
				return nil
			}
		}
		return fmt.Errorf("token %s not found", tokenId)
	})

	var oldTokenId string
	acceptance.AccountLevel(t, acceptance.Step{
		Template: tokenUpdateTemplate(`comment = "test foo"`, ""),
		Check:    checkTokenExists,
	}, acceptance.Step{
		// Updating the comment causes the old token to be deleted and a new one to be created
		Template: tokenUpdateTemplate(`comment = "test bar"`, ""),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{ExpectNotDestroyed("databricks_mws_workspaces.this")},
		},
		Check: resource.ComposeAggregateTestCheckFunc(
			checkTokenExists,
			// Capture the token ID at the end of this step to verify it is not changed in future steps.
			func(s *terraform.State) error {
				state, ok := s.RootModule().Resources["databricks_mws_workspaces.this"]
				if !ok {
					return fmt.Errorf("resource not found in state")
				}
				instanceState := state.Primary.Attributes
				oldTokenId = instanceState["token.0.token_id"]
				return nil
			},
		),
	}, acceptance.Step{
		// Modifying the tags doesn't change the token but does modify the workspace.
		Template: tokenUpdateTemplate(`comment = "test bar"`, `"Key" = "Value"`),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{ExpectNotDestroyed("databricks_mws_workspaces.this")},
		},
		Check: func(s *terraform.State) error {
			state, ok := s.RootModule().Resources["databricks_mws_workspaces.this"]
			if !ok {
				return fmt.Errorf("resource not found in state")
			}
			instanceState := state.Primary.Attributes
			assert.Equal(t, instanceState["custom_tags.Key"], "Value")
			assert.Equal(t, instanceState["token.0.token_id"], oldTokenId)
			return nil
		},
	}, acceptance.Step{
		// It is also possible to modify the token comment and tags at the same time.
		Template: tokenUpdateTemplate(`comment = "test quux"`, `"Key" = "Value2"`),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{ExpectNotDestroyed("databricks_mws_workspaces.this")},
		},
		Check: func(s *terraform.State) error {
			state, ok := s.RootModule().Resources["databricks_mws_workspaces.this"]
			if !ok {
				return fmt.Errorf("resource not found in state")
			}
			instanceState := state.Primary.Attributes
			assert.Equal(t, instanceState["custom_tags.Key"], "Value2")
			assert.NotEqual(t, instanceState["token.0.token_id"], oldTokenId)
			return nil
		},
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
