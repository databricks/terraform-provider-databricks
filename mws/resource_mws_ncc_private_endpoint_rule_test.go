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

func getTestNccRule() *settings.NccPrivateEndpointRule {
	return &settings.NccPrivateEndpointRule{
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
				PrivateEndpointRule: settings.CreatePrivateEndpointRule{
					ResourceId: "resource_id",
					GroupId:    "blob",
				},
			}).Return(getTestNccRule(), nil)
			e.GetPrivateEndpointRuleByNetworkConnectivityConfigIdAndPrivateEndpointRuleId(mock.Anything, "ncc_id", "rule_id").Return(getTestNccRule(), nil)
		},
		Resource:  ResourceMwsNccPrivateEndpointRule(),
		AccountID: "abc",
		HCL: `
		network_connectivity_config_id = "ncc_id"
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
				PrivateEndpointRule: settings.CreatePrivateEndpointRule{
					ResourceId: "resource_id",
					GroupId:    "blob",
				},
			}).Return(nil, &apierr.APIError{Message: "error"})
		},
		Resource:  ResourceMwsNccPrivateEndpointRule(),
		AccountID: "abc",
		HCL: `
		network_connectivity_config_id = "ncc_id"
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
		"network_connectivity_config_id": "ncc_id",
		"resource_id":                    "resource_id",
		"rule_id":                        "rule_id",
		"endpoint_name":                  "endpoint_name",
		"connection_state":               "PENDING",
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

func TestResourceNccPrivateEndpointRulePrivateEndpointRuleUpdateDomainName(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			e := a.GetMockNetworkConnectivityAPI().EXPECT()
			e.UpdatePrivateEndpointRule(mock.Anything, settings.UpdateNccPrivateEndpointRuleRequest{
				NetworkConnectivityConfigId: "ncc_id",
				PrivateEndpointRuleId:       "rule_id",
				PrivateEndpointRule: settings.UpdatePrivateEndpointRule{
					DomainNames: []string{"my-new-example.exampledomain.com", "my-new-example2.exampledomain.com"},
				},
				UpdateMask: "domain_names",
			}).Return(getTestNccRule(), nil)
			e.GetPrivateEndpointRuleByNetworkConnectivityConfigIdAndPrivateEndpointRuleId(mock.Anything, "ncc_id", "rule_id").Return(
				&settings.NccPrivateEndpointRule{
					GroupId:                     "group_id",
					ResourceId:                  "resource_id",
					RuleId:                      "rule_id",
					NetworkConnectivityConfigId: "ncc_id",
					EndpointName:                "endpoint_name",
					ConnectionState:             "PENDING",
					DomainNames:                 []string{"my-new-example.exampledomain.com", "my-new-example2.exampledomain.com"},
				}, nil)
		},
		Resource:  ResourceMwsNccPrivateEndpointRule(),
		AccountID: "abc",
		ID:        "ncc_id/rule_id",
		InstanceState: map[string]string{
			"network_connectivity_config_id": "ncc_id",
			"resource_id":                    "resource_id",
			"rule_id":                        "rule_id",
			"domain_names.#":                 "1",
			"domain_names.0":                 "my-example.exampledomain.com",
		},
		HCL: `
		network_connectivity_config_id = "ncc_id"
		resource_id = "resource_id"
		domain_names = ["my-new-example.exampledomain.com", "my-new-example2.exampledomain.com"]
		`,
		Update: true,
	}.ApplyAndExpectData(t, map[string]any{
		"id":             "ncc_id/rule_id",
		"domain_names.#": 2,
		"domain_names.0": "my-new-example.exampledomain.com",
		"domain_names.1": "my-new-example2.exampledomain.com",
	})
}

func TestResourceNccPrivateEndpointRulePrivateEndpointRuleUpdateResourceName(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			e := a.GetMockNetworkConnectivityAPI().EXPECT()
			e.UpdatePrivateEndpointRule(mock.Anything, settings.UpdateNccPrivateEndpointRuleRequest{
				NetworkConnectivityConfigId: "ncc_id",
				PrivateEndpointRuleId:       "rule_id",
				PrivateEndpointRule: settings.UpdatePrivateEndpointRule{
					ResourceNames: []string{"bucket1", "bucket2"},
				},
				UpdateMask: "resource_names",
			}).Return(getTestNccRule(), nil)
			e.GetPrivateEndpointRuleByNetworkConnectivityConfigIdAndPrivateEndpointRuleId(mock.Anything, "ncc_id", "rule_id").Return(
				&settings.NccPrivateEndpointRule{
					GroupId:                     "group_id",
					ResourceId:                  "resource_id",
					RuleId:                      "rule_id",
					NetworkConnectivityConfigId: "ncc_id",
					EndpointName:                "endpoint_name",
					ConnectionState:             "PENDING",
					ResourceNames:               []string{"bucket1", "bucket2"},
				}, nil)
		},
		Resource:  ResourceMwsNccPrivateEndpointRule(),
		AccountID: "abc",
		ID:        "ncc_id/rule_id",
		InstanceState: map[string]string{
			"network_connectivity_config_id": "ncc_id",
			"resource_id":                    "resource_id",
			"rule_id":                        "rule_id",
			"resource_names.#":               "1",
			"resource_names.0":               "bucket1",
		},
		HCL: `
		network_connectivity_config_id = "ncc_id"
		resource_id = "resource_id"
		resource_names = ["bucket1", "bucket2"]
		`,
		Update: true,
	}.ApplyAndExpectData(t, map[string]any{
		"id":               "ncc_id/rule_id",
		"resource_names.#": 2,
		"resource_names.0": "bucket1",
		"resource_names.1": "bucket2",
	})
}

func TestResourceNccPrivateEndpointRuleDelete(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockNetworkConnectivityAPI().EXPECT().
				DeletePrivateEndpointRuleByNetworkConnectivityConfigIdAndPrivateEndpointRuleId(mock.Anything, "ncc_id", "rule_id").
				Return(&settings.NccPrivateEndpointRule{}, nil)
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
		network_connectivity_config_id = "ncc_id"
		resource_id = "resource_id"
		group_id = "blob"
		`,
		Delete: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "error")
	assert.Equal(t, "ncc_id/rule_id", d.Id())
}
