package privateendpointrule

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/mws"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-mux/tf5to6server"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	sdkschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfresource "github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stretchr/testify/mock"
)

// A local Terraform CLI exercises real plan/apply/state handling; only the
// account API is mocked so these tests need neither preview access nor cloud resources.
func TestGcpTerraformProviderHandoff(t *testing.T) {
	if os.Getenv("TF_ACC_TERRAFORM_PATH") == "" {
		t.Skip("set TF_ACC_TERRAFORM_PATH to an installed Terraform CLI to run offline lifecycle tests")
	}
	const pscURI = "projects/p/regions/r/forwardingRules/endpoint"
	for _, first := range []string{"SDKv2", "Plugin Framework"} {
		for _, tt := range []struct {
			name    string
			target  string
			gcp     settings.GcpEndpoint
			domains []string
		}{
			{
				name:   "Google APIs",
				target: `google_api_endpoints { endpoints = ["storage.googleapis.com"] }`,
				gcp: settings.GcpEndpoint{
					GoogleApiEndpoints: &settings.GoogleApiEndpoints{Endpoints: []string{"storage.googleapis.com"}},
				},
			},
			{
				name:   "all VPC-SC services",
				target: `all_vpc_sc_services = true`,
				gcp:    settings.GcpEndpoint{AllVpcScServices: true},
			},
			{
				name:    "service attachment",
				target:  `service_attachment = "projects/p/regions/r/serviceAttachments/a"`,
				gcp:     settings.GcpEndpoint{ServiceAttachment: "projects/p/regions/r/serviceAttachments/a"},
				domains: []string{"private.example.com"},
			},
		} {
			t.Run(first+"/"+tt.name, func(t *testing.T) {
				apiTarget := tt.gcp
				apiTarget.PscEndpointUri = pscURI
				rule := &settings.NccPrivateEndpointRule{
					NetworkConnectivityConfigId: "ncc_id",
					RuleId:                      "rule_id",
					ConnectionState:             "ESTABLISHED",
					Enabled:                     true,
					GcpEndpoint:                 &apiTarget,
					DomainNames:                 tt.domains,
				}
				var account *mocks.MockAccountClient
				qa.MockAccountsApply(t, func(a *mocks.MockAccountClient) {
					account = a
					a.AccountClient.Config = &config.Config{Host: "https://accounts.gcp.databricks.com", AccountID: "abc"}
					e := a.GetMockNetworkConnectivityAPI().EXPECT()
					e.CreatePrivateEndpointRule(mock.Anything, settings.CreatePrivateEndpointRuleRequest{
						NetworkConnectivityConfigId: "ncc_id",
						PrivateEndpointRule: settings.CreatePrivateEndpointRule{
							GcpEndpoint: &tt.gcp,
							DomainNames: tt.domains,
						},
					}).Return(rule, nil).Once()
					e.GetPrivateEndpointRuleByNetworkConnectivityConfigIdAndPrivateEndpointRuleId(mock.Anything, "ncc_id", "rule_id").Return(rule, nil)
					e.GetPrivateEndpointRule(mock.Anything, settings.GetPrivateEndpointRuleRequest{
						NetworkConnectivityConfigId: "ncc_id", PrivateEndpointRuleId: "rule_id",
					}).Return(rule, nil)
					if first == "SDKv2" {
						e.DeletePrivateEndpointRuleByNetworkConnectivityConfigIdAndPrivateEndpointRuleId(mock.Anything, "ncc_id", "rule_id").Return(nil, nil).Once()
					} else {
						e.DeletePrivateEndpointRule(mock.Anything, settings.DeletePrivateEndpointRuleRequest{
							NetworkConnectivityConfigId: "ncc_id", PrivateEndpointRuleId: "rule_id",
						}).Return(nil, nil).Once()
					}
					if tt.gcp.ServiceAttachment == "" {
						for _, enabled := range []bool{false, true} {
							e.UpdatePrivateEndpointRule(mock.Anything, settings.UpdateNccPrivateEndpointRuleRequest{
								NetworkConnectivityConfigId: "ncc_id", PrivateEndpointRuleId: "rule_id",
								UpdateMask: "enabled", PrivateEndpointRule: settings.UpdatePrivateEndpointRule{Enabled: enabled},
							}).Run(func(_ context.Context, req settings.UpdateNccPrivateEndpointRuleRequest) {
								rule.Enabled = req.PrivateEndpointRule.Enabled
							}).Return(rule, nil).Once()
						}
					}
				}, func(ctx context.Context, client *common.DatabricksClient) {
					sdkFactories := privateEndpointRuleSDKv2Factories(client)
					pfFactories := map[string]func() (tfprotov6.ProviderServer, error){
						"databricks": providerserver.NewProtocol6WithError(&privateEndpointRuleTestProvider{api: account.AccountClient.NetworkConnectivity}),
					}
					firstFactories, secondFactories := sdkFactories, pfFactories
					if first == "Plugin Framework" {
						firstFactories, secondFactories = pfFactories, sdkFactories
					}
					config := `resource "databricks_mws_ncc_private_endpoint_rule" "this" {
					network_connectivity_config_id = "ncc_id"
					gcp_endpoint {
						` + tt.target + `
					}
					}`
					if len(tt.domains) > 0 {
						config = strings.Replace(config, "network_connectivity_config_id", "domain_names = [\""+tt.domains[0]+"\"]\nnetwork_connectivity_config_id", 1)
					}
					const address = "databricks_mws_ncc_private_endpoint_rule.this"
					noChange := tfresource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{plancheck.ExpectResourceAction(address, plancheck.ResourceActionNoop)},
					}
					steps := []tfresource.TestStep{
						{
							ProtoV6ProviderFactories: firstFactories,
							Config:                   config,
							Check: tfresource.ComposeTestCheckFunc(
								tfresource.TestCheckResourceAttr(address, "id", "ncc_id/rule_id"),
								tfresource.TestCheckResourceAttr(address, "gcp_endpoint.0.psc_endpoint_uri", pscURI),
							),
						},
						{
							ProtoV6ProviderFactories: secondFactories,
							Config:                   config,
							ConfigPlanChecks:         noChange,
						},
						{
							ProtoV6ProviderFactories: secondFactories,
							ResourceName:             address,
							ImportState:              true,
							ImportStateId:            "ncc_id/rule_id",
							ImportStateVerify:        true,
						},
						{
							ProtoV6ProviderFactories: firstFactories,
							Config:                   config,
							ConfigPlanChecks:         noChange,
						},
					}
					if tt.gcp.ServiceAttachment == "" {
						for i, enabled := range []string{"false", "true"} {
							factories := firstFactories
							if i == 1 {
								factories = secondFactories
							}
							steps = append(steps, tfresource.TestStep{
								ProtoV6ProviderFactories: factories,
								Config:                   strings.Replace(config, "network_connectivity_config_id", "enabled = "+enabled+"\nnetwork_connectivity_config_id", 1),
								Check: tfresource.ComposeTestCheckFunc(
									tfresource.TestCheckResourceAttr(address, "id", "ncc_id/rule_id"),
									tfresource.TestCheckResourceAttr(address, "enabled", enabled),
									tfresource.TestCheckResourceAttr(address, "gcp_endpoint.0.psc_endpoint_uri", pscURI),
								),
							})
						}
						steps = append(steps, tfresource.TestStep{ProtoV6ProviderFactories: firstFactories, Config: config, ConfigPlanChecks: noChange})
					} else {
						steps = append(steps, tfresource.TestStep{
							ProtoV6ProviderFactories: firstFactories,
							Config:                   strings.Replace(config, "serviceAttachments/a", "serviceAttachments/b", 1),
							PlanOnly:                 true, ExpectNonEmptyPlan: true,
							ConfigPlanChecks: tfresource.ConfigPlanChecks{PostApplyPreRefresh: []plancheck.PlanCheck{
								plancheck.ExpectResourceAction(address, plancheck.ResourceActionDestroyBeforeCreate),
							}},
						})
					}
					tfresource.UnitTest(t, tfresource.TestCase{Steps: steps})
				})
			})
		}
	}
}

func privateEndpointRuleSDKv2Factories(client *common.DatabricksClient) map[string]func() (tfprotov6.ProviderServer, error) {
	return map[string]func() (tfprotov6.ProviderServer, error){
		"databricks": func() (tfprotov6.ProviderServer, error) {
			p := &sdkschema.Provider{
				ResourcesMap: map[string]*sdkschema.Resource{
					"databricks_mws_ncc_private_endpoint_rule": mws.ResourceMwsNccPrivateEndpointRule().ToResource(),
				},
				ConfigureContextFunc: func(context.Context, *sdkschema.ResourceData) (any, diag.Diagnostics) {
					return client, nil
				},
			}
			return tf5to6server.UpgradeServer(context.Background(), func() tfprotov5.ProviderServer { return p.GRPCProvider() })
		},
	}
}
