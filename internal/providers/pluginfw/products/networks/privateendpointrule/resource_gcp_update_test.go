package privateendpointrule

import (
	"context"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	tfresource "github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/stretchr/testify/mock"
)

func TestGcpTerraformTargetUpdatesWithDefaults(t *testing.T) {
	for _, defaults := range []string{"omitted", "explicit", "null"} {
		t.Run(defaults, func(t *testing.T) {
			const address = "databricks_mws_ncc_private_endpoint_rule.this"
			const pscURI = "projects/p/regions/r/forwardingRules/endpoint"
			googleAPIs := settings.GcpEndpoint{GoogleApiEndpoints: &settings.GoogleApiEndpoints{Endpoints: []string{"storage.googleapis.com"}}}
			allServices := settings.GcpEndpoint{AllVpcScServices: true}
			rule := &settings.NccPrivateEndpointRule{
				RuleId: "rule_id", ConnectionState: "ESTABLISHED", Enabled: true,
				GcpEndpoint: &settings.GcpEndpoint{GoogleApiEndpoints: googleAPIs.GoogleApiEndpoints, PscEndpointUri: pscURI},
			}
			var account *mocks.MockAccountClient
			qa.MockAccountsApply(t, func(a *mocks.MockAccountClient) {
				account = a
				a.AccountClient.Config = &config.Config{Host: "https://accounts.gcp.databricks.com", AccountID: "abc"}
				e := a.GetMockNetworkConnectivityAPI().EXPECT()
				e.CreatePrivateEndpointRule(mock.Anything, settings.CreatePrivateEndpointRuleRequest{
					NetworkConnectivityConfigId: "ncc_id",
					PrivateEndpointRule:         settings.CreatePrivateEndpointRule{GcpEndpoint: &googleAPIs},
				}).Return(rule, nil).Once()
				e.GetPrivateEndpointRule(mock.Anything, settings.GetPrivateEndpointRuleRequest{
					NetworkConnectivityConfigId: "ncc_id", PrivateEndpointRuleId: "rule_id",
				}).Return(rule, nil)
				for _, target := range []settings.GcpEndpoint{allServices, googleAPIs} {
					e.UpdatePrivateEndpointRule(mock.Anything, settings.UpdateNccPrivateEndpointRuleRequest{
						NetworkConnectivityConfigId: "ncc_id", PrivateEndpointRuleId: "rule_id",
						UpdateMask: "gcp_endpoint", PrivateEndpointRule: settings.UpdatePrivateEndpointRule{GcpEndpoint: &target},
					}).Run(func(_ context.Context, req settings.UpdateNccPrivateEndpointRuleRequest) {
						next := *req.PrivateEndpointRule.GcpEndpoint
						next.PscEndpointUri = pscURI
						rule.GcpEndpoint = &next
					}).Return(rule, nil).Once()
				}
				e.DeletePrivateEndpointRule(mock.Anything, settings.DeletePrivateEndpointRuleRequest{
					NetworkConnectivityConfigId: "ncc_id", PrivateEndpointRuleId: "rule_id",
				}).Return(nil, nil).Once()
			}, func(ctx context.Context, client *common.DatabricksClient) {
				googleConfig := `resource "databricks_mws_ncc_private_endpoint_rule" "this" {
					network_connectivity_config_id = "ncc_id"
					gcp_endpoint {
						google_api_endpoints { endpoints = ["storage.googleapis.com"] }
					}
				}`
				allConfig := strings.Replace(googleConfig, `google_api_endpoints { endpoints = ["storage.googleapis.com"] }`, `all_vpc_sc_services = true`, 1)
				if defaults == "explicit" {
					googleConfig = strings.Replace(googleConfig, "gcp_endpoint {", "gcp_endpoint {\nall_vpc_sc_services = false\nservice_attachment = \"\"", 1)
					allConfig = strings.Replace(allConfig, "gcp_endpoint {", "gcp_endpoint {\nservice_attachment = \"\"", 1)
				} else if defaults == "null" {
					googleConfig = strings.Replace(googleConfig, "gcp_endpoint {", "gcp_endpoint {\nall_vpc_sc_services = null\nservice_attachment = null", 1)
					allConfig = strings.Replace(allConfig, "gcp_endpoint {", "gcp_endpoint {\nservice_attachment = null", 1)
				}
				factories := map[string]func() (tfprotov6.ProviderServer, error){
					"databricks": providerserver.NewProtocol6WithError(&privateEndpointRuleTestProvider{api: account.AccountClient.NetworkConnectivity}),
				}
				noChange := tfresource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(address, plancheck.ResourceActionNoop),
				}}
				targetPlan := func(allServices bool) tfresource.ConfigPlanChecks {
					return tfresource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(address, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(address, tfjsonpath.New("gcp_endpoint").AtSliceIndex(0).AtMapKey("all_vpc_sc_services"), knownvalue.Bool(allServices)),
					}}
				}
				check := func(allServices string) tfresource.TestCheckFunc {
					return tfresource.ComposeTestCheckFunc(
						tfresource.TestCheckResourceAttr(address, "id", "ncc_id/rule_id"),
						tfresource.TestCheckResourceAttr(address, "gcp_endpoint.0.all_vpc_sc_services", allServices),
						tfresource.TestCheckResourceAttr(address, "gcp_endpoint.0.service_attachment", ""),
						tfresource.TestCheckResourceAttr(address, "gcp_endpoint.0.psc_endpoint_uri", pscURI),
					)
				}
				tfresource.UnitTest(t, tfresource.TestCase{Steps: []tfresource.TestStep{
					{ProtoV6ProviderFactories: factories, Config: googleConfig, Check: check("false")},
					{ProtoV6ProviderFactories: factories, Config: allConfig, ConfigPlanChecks: targetPlan(true), Check: check("true")},
					{ProtoV6ProviderFactories: factories, Config: allConfig, ConfigPlanChecks: noChange},
					{ProtoV6ProviderFactories: factories, Config: googleConfig, ConfigPlanChecks: targetPlan(false), Check: check("false")},
					{ProtoV6ProviderFactories: factories, Config: googleConfig, ConfigPlanChecks: noChange},
				}})
			})
		})
	}
}
