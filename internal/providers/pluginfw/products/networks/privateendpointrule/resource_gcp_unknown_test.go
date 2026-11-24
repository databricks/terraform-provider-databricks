package privateendpointrule

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	tfresource "github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/stretchr/testify/mock"
)

func TestGcpTerraformUnknownScalarSelectors(t *testing.T) {
	const address = "databricks_mws_ncc_private_endpoint_rule.this"
	const pscURI = "projects/p/regions/r/forwardingRules/endpoint"
	for _, tt := range []struct {
		name       string
		input      string
		target     string
		attributes []string
		gcp        settings.GcpEndpoint
	}{
		{
			name: "all services", input: "true",
			target:     "all_vpc_sc_services = terraform_data.target.output",
			attributes: []string{"all_vpc_sc_services"},
			gcp:        settings.GcpEndpoint{AllVpcScServices: true},
		},
		{
			name: "service attachment", input: `"projects/p/regions/r/serviceAttachments/a"`,
			target:     "service_attachment = terraform_data.target.output",
			attributes: []string{"service_attachment"},
			gcp:        settings.GcpEndpoint{ServiceAttachment: "projects/p/regions/r/serviceAttachments/a"},
		},
		{
			name: "explicit defaults", input: `{ all_vpc_sc_services = false, service_attachment = "" }`,
			target: `all_vpc_sc_services = terraform_data.target.output.all_vpc_sc_services
				service_attachment = terraform_data.target.output.service_attachment
				google_api_endpoints { endpoints = ["storage.googleapis.com"] }`,
			attributes: []string{"all_vpc_sc_services", "service_attachment"},
			gcp:        settings.GcpEndpoint{GoogleApiEndpoints: &settings.GoogleApiEndpoints{Endpoints: []string{"storage.googleapis.com"}}},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			apiTarget := tt.gcp
			apiTarget.PscEndpointUri = pscURI
			rule := &settings.NccPrivateEndpointRule{
				RuleId: "rule_id", ConnectionState: "ESTABLISHED", Enabled: true, GcpEndpoint: &apiTarget,
			}
			var account *mocks.MockAccountClient
			qa.MockAccountsApply(t, func(a *mocks.MockAccountClient) {
				account = a
				a.AccountClient.Config = &config.Config{Host: "https://accounts.gcp.databricks.com", AccountID: "abc"}
				e := a.GetMockNetworkConnectivityAPI().EXPECT()
				e.CreatePrivateEndpointRule(mock.Anything, settings.CreatePrivateEndpointRuleRequest{
					NetworkConnectivityConfigId: "ncc_id",
					PrivateEndpointRule:         settings.CreatePrivateEndpointRule{GcpEndpoint: &tt.gcp},
				}).Return(rule, nil).Once()
				e.GetPrivateEndpointRule(mock.Anything, settings.GetPrivateEndpointRuleRequest{
					NetworkConnectivityConfigId: "ncc_id", PrivateEndpointRuleId: "rule_id",
				}).Return(rule, nil)
				e.DeletePrivateEndpointRule(mock.Anything, settings.DeletePrivateEndpointRuleRequest{
					NetworkConnectivityConfigId: "ncc_id", PrivateEndpointRuleId: "rule_id",
				}).Return(nil, nil).Once()
			}, func(ctx context.Context, client *common.DatabricksClient) {
				factories := map[string]func() (tfprotov6.ProviderServer, error){
					"databricks": providerserver.NewProtocol6WithError(&privateEndpointRuleTestProvider{api: account.AccountClient.NetworkConnectivity}),
				}
				// The built-in terraform_data resource supplies values only at apply;
				// schema defaults must not replace these unknown inputs during planning.
				config := `resource "terraform_data" "target" { input = ` + tt.input + ` }
				resource "databricks_mws_ncc_private_endpoint_rule" "this" {
					network_connectivity_config_id = "ncc_id"
					gcp_endpoint {
						` + tt.target + `
					}
				}`
				checks := []plancheck.PlanCheck{plancheck.ExpectResourceAction(address, plancheck.ResourceActionCreate)}
				for _, attribute := range tt.attributes {
					checks = append(checks, plancheck.ExpectUnknownValue(address, tfjsonpath.New("gcp_endpoint").AtSliceIndex(0).AtMapKey(attribute)))
				}
				tfresource.UnitTest(t, tfresource.TestCase{Steps: []tfresource.TestStep{
					{
						ProtoV6ProviderFactories: factories, Config: config,
						ConfigPlanChecks: tfresource.ConfigPlanChecks{PreApply: checks},
						Check:            tfresource.TestCheckResourceAttr(address, "gcp_endpoint.0.psc_endpoint_uri", pscURI),
					},
					{
						ProtoV6ProviderFactories: factories, Config: config,
						ConfigPlanChecks: tfresource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(address, plancheck.ResourceActionNoop),
						}},
					},
				}})
			})
		})
	}
}
