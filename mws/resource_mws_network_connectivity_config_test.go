package mws

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func getTestNccRule() *settings.NccAzurePrivateEndpointRule {
	return &settings.NccAzurePrivateEndpointRule{
		GroupId:                     "group_id",
		ResourceId:                  "resource_id",
		RuleId:                      "rule_id",
		NetworkConnectivityConfigId: "ncc_id",
		EndpointName:                "endpoint_name",
		ConnectionState:             "PENDING",
	}
}

func TestResourceNccPrivateEndpointRulePrivateEndpointRuleCreate(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			e := a.GetMockNetworkConnectivityAPI().EXPECT()
			e.CreatePrivateEndpointRule(mock.Anything, settings.CreatePrivateEndpointRuleRequest{
				NetworkConnectivityConfigId: "ncc_id",
				ResourceId:                  "resource_id",
				GroupId:                     "blob",
			}).Return(getTestNccRule(), nil)
			e.GetPrivateEndpointRuleByNetworkConnectivityConfigIdAndPrivateEndpointRuleId(mock.Anything, "ncc_id", "rule_id").Return(getTestNccRule(), nil)
		},
		Resource:  ResourceMwsNccPrivateEndpointRule(),
		AccountID: "abc",
		HCL: `
		ncc_id = "ncc_id"
		resource_id = "resource_id"
		group_id = "blob"
		`,
		Create: true,
	}.ApplyAndExpectData(t, map[string]any{"id": "ncc_id/rule_id"})
}

func TestResourceNccPrivateEndpointRulePrivateEndpointRuleCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			e := a.GetMockNetworkConnectivityAPI().EXPECT()
			e.CreatePrivateEndpointRule(mock.Anything, settings.CreatePrivateEndpointRuleRequest{
				NetworkConnectivityConfigId: "ncc_id",
				ResourceId:                  "resource_id",
				GroupId:                     "blob",
			}).Return(nil, &apierr.APIError{Message: "error"})
		},
		Resource:  ResourceMwsNccPrivateEndpointRule(),
		AccountID: "abc",
		HCL: `
		ncc_id = "ncc_id"
		resource_id = "resource_id"
		group_id = "blob"
		`,
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "error")
	assert.Equal(t, "", d.Id())
}

func TestResourceNccPrivateEndpointRuleRead(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockNetworkConnectivityAPI().EXPECT().
				GetPrivateEndpointRuleByNetworkConnectivityConfigIdAndPrivateEndpointRuleId(mock.Anything, "ncc_id", "rule_id").
				Return(getTestNccRule(), nil)
		},
		Resource:  ResourceMwsNccPrivateEndpointRule(),
		AccountID: "abc",
		Read:      true,
		New:       true,
		ID:        "ncc_id/rule_id",
	}.ApplyAndExpectData(t, map[string]any{
		"id":                             "ncc_id/rule_id",
		"account_id":                     "abc",
		"name":                           "ncc_name",
		"region":                         "ar",
		"network_connectivity_config_id": "ncc_id",
	})
}

func TestResourceNccPrivateEndpointRuleRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockNetworkConnectivityAPI().EXPECT().
				GetPrivateEndpointRuleByNetworkConnectivityConfigIdAndPrivateEndpointRuleId(mock.Anything, "ncc_id", "rule_id").
				Return(nil, &apierr.APIError{Message: "error"})
		},
		Resource:  ResourceMwsNccPrivateEndpointRule(),
		AccountID: "abc",
		Read:      true,
		ID:        "ncc_id/rule_id",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "error")
	assert.Equal(t, "ncc_id/rule_id", d.Id())
}

func TestResourceNccPrivateEndpointRuleDelete(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockNetworkConnectivityAPI().EXPECT().
				DeletePrivateEndpointRuleByNetworkConnectivityConfigIdAndPrivateEndpointRuleId(mock.Anything, "ncc_id", "rule_id").
				Return(&settings.NccAzurePrivateEndpointRule{}, nil)
		},
		Resource:  ResourceMwsNccPrivateEndpointRule(),
		AccountID: "abc",
		Delete:    true,
		ID:        "ncc_id/rule_id",
	}.ApplyAndExpectData(t, map[string]any{"id": "ncc_id/rule_id"})
}

func TestResourceNccPrivateEndpointRuleDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockNetworkConnectivityAPI().EXPECT().
				DeletePrivateEndpointRuleByNetworkConnectivityConfigIdAndPrivateEndpointRuleId(mock.Anything, "ncc_id", "rule_id").
				Return(nil, &apierr.APIError{Message: "error"})
		},
		Resource:  ResourceMwsNccPrivateEndpointRule(),
		AccountID: "abc",
		ID:        "ncc_id/rule_id",
		HCL: `
		ncc_id = "ncc_id"
		resource_id = "resource_id"
		group_id = "blob"
		`,
		Delete: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "error")
	assert.Equal(t, "ncc_id/rule_id", d.Id())
}
