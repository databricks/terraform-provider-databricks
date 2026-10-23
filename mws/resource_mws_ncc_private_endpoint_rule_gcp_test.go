package mws

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestResourceNccPrivateEndpointRuleGcpUpdatesIgnoreReadOnlyChanges(t *testing.T) {
	for _, targetChanged := range []bool{false, true} {
		name := "URI and enablement"
		if targetChanged {
			name = "URI and target"
		}
		t.Run(name, func(t *testing.T) {
			resource := ResourceMwsNccPrivateEndpointRule()
			state := &terraform.InstanceState{
				ID: "ncc_id/rule_id",
				Attributes: map[string]string{
					"network_connectivity_config_id":     "ncc_id",
					"rule_id":                            "rule_id",
					"enabled":                            "true",
					"gcp_endpoint.#":                     "1",
					"gcp_endpoint.0.all_vpc_sc_services": "true",
					"gcp_endpoint.0.psc_endpoint_uri":    "old-uri",
				},
			}
			diff := &terraform.InstanceDiff{Attributes: map[string]*terraform.ResourceAttrDiff{
				"gcp_endpoint.0.psc_endpoint_uri": {Old: "old-uri", New: "new-uri"},
			}}
			request := settings.UpdateNccPrivateEndpointRuleRequest{
				NetworkConnectivityConfigId: "ncc_id", PrivateEndpointRuleId: "rule_id",
				UpdateMask: "enabled", PrivateEndpointRule: settings.UpdatePrivateEndpointRule{Enabled: false},
			}
			if targetChanged {
				diff.Attributes["gcp_endpoint.0.all_vpc_sc_services"] = &terraform.ResourceAttrDiff{Old: "true", New: "false"}
				diff.Attributes["gcp_endpoint.0.google_api_endpoints.#"] = &terraform.ResourceAttrDiff{New: "1"}
				diff.Attributes["gcp_endpoint.0.google_api_endpoints.0.endpoints.#"] = &terraform.ResourceAttrDiff{New: "1"}
				diff.Attributes["gcp_endpoint.0.google_api_endpoints.0.endpoints.0"] = &terraform.ResourceAttrDiff{New: "storage.googleapis.com"}
				request.UpdateMask = "gcp_endpoint"
				request.PrivateEndpointRule.GcpEndpoint = &settings.GcpEndpoint{
					GoogleApiEndpoints: &settings.GoogleApiEndpoints{Endpoints: []string{"storage.googleapis.com"}},
				}
			} else {
				diff.Attributes["enabled"] = &terraform.ResourceAttrDiff{Old: "true", New: "false"}
			}
			data, err := schema.InternalMap(resource.Schema).Data(state, diff)
			require.NoError(t, err)
			require.True(t, data.HasChange("gcp_endpoint"))
			qa.MockAccountsApply(t, func(a *mocks.MockAccountClient) {
				a.AccountClient.Config = &config.Config{Host: "https://accounts.gcp.databricks.com"}
				a.GetMockNetworkConnectivityAPI().EXPECT().UpdatePrivateEndpointRule(mock.Anything, request).
					Return(&settings.NccPrivateEndpointRule{}, nil).Once()
			}, func(ctx context.Context, client *common.DatabricksClient) {
				require.NoError(t, resource.Update(ctx, data, client))
			})
		})
	}
}
