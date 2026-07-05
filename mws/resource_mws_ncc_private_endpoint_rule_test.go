package mws

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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
		Host:      "https://accounts.cloud.databricks.com",
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
		Host:      "https://accounts.cloud.databricks.com",
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
		Host:      "https://accounts.cloud.databricks.com",
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
		Host:      "https://accounts.cloud.databricks.com",
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
		Host:      "https://accounts.cloud.databricks.com",
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
		Host:      "https://accounts.cloud.databricks.com",
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

func TestResourceNccPrivateEndpointRulePrivateEndpointRuleUpdateEnabled(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			e := a.GetMockNetworkConnectivityAPI().EXPECT()
			e.UpdatePrivateEndpointRule(mock.Anything, settings.UpdateNccPrivateEndpointRuleRequest{
				NetworkConnectivityConfigId: "ncc_id",
				PrivateEndpointRuleId:       "rule_id",
				PrivateEndpointRule:         settings.UpdatePrivateEndpointRule{Enabled: false},
				UpdateMask:                  "enabled",
			}).Return(getTestNccRule(), nil)
			e.GetPrivateEndpointRuleByNetworkConnectivityConfigIdAndPrivateEndpointRuleId(mock.Anything, "ncc_id", "rule_id").Return(
				&settings.NccPrivateEndpointRule{
					GroupId:                     "group_id",
					ResourceId:                  "resource_id",
					RuleId:                      "rule_id",
					NetworkConnectivityConfigId: "ncc_id",
					EndpointName:                "endpoint_name",
					ConnectionState:             "PENDING",
					ResourceNames:               []string{"bucket1"},
				}, nil)
		},
		Resource:  ResourceMwsNccPrivateEndpointRule(),
		AccountID: "abc",
		Host:      "https://accounts.cloud.databricks.com",
		ID:        "ncc_id/rule_id",
		InstanceState: map[string]string{
			"network_connectivity_config_id": "ncc_id",
			"resource_id":                    "resource_id",
			"rule_id":                        "rule_id",
			"resource_names.#":               "1",
			"resource_names.0":               "bucket1",
			"enabled":                        "true",
		},
		HCL: `
		network_connectivity_config_id = "ncc_id"
		resource_id = "resource_id"
		resource_names = ["bucket1"]
		enabled = false
		`,
		Update: true,
	}.ApplyAndExpectData(t, map[string]any{
		"id":               "ncc_id/rule_id",
		"resource_names.#": 1,
		"resource_names.0": "bucket1",
		"enabled":          false,
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
		Host:      "https://accounts.cloud.databricks.com",
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
		Host:      "https://accounts.cloud.databricks.com",
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

// account_id must be Computed so that a value populated by the backend on Read
// does not register as user-driven drift. The sibling
// databricks_mws_network_connectivity_config marks account_id Computed for the
// same reason; the omission here is the root cause of #5347.
func TestResourceNccPrivateEndpointRule_AccountIdIsComputed(t *testing.T) {
	s := ResourceMwsNccPrivateEndpointRule().Schema
	accountID, ok := s["account_id"]
	assert.True(t, ok, "account_id should be present in the schema")
	assert.True(t, accountID.Computed,
		"account_id must be Computed; otherwise a backend-populated value triggers spurious drift and an empty-update-mask error")
}

// Reproduces the customer-facing symptom of #5347: when a prior Read has
// populated account_id into state (the backend intermittently returns it) but
// HCL doesn't set it, no in-place update should be planned. Before the fix the
// diff plans account_id = "..." -> null, which then triggers an Update API
// call with an empty update_mask, which the backend rejects with
// "Update mask must be specified".
func TestResourceNccPrivateEndpointRule_NoDriftWhenBackendReturnsAccountId(t *testing.T) {
	qa.ResourceFixture{
		Resource:  ResourceMwsNccPrivateEndpointRule(),
		AccountID: "abc",
		Host:      "https://accounts.cloud.databricks.com",
		ID:        "ncc_id/rule_id",
		InstanceState: map[string]string{
			"network_connectivity_config_id": "ncc_id",
			"rule_id":                        "rule_id",
			"account_id":                     "abc",
			"group_id":                       "blob",
			"resource_id":                    "resource_id",
			"endpoint_name":                  "endpoint_name",
			"connection_state":               "PENDING",
			"creation_time":                  "0",
			"updated_time":                   "0",
			"vpc_endpoint_id":                "",
			"enabled":                        "false",
			// Server-populated read-only fields a prior Read writes into state.
			// Present here so the now-Computed attributes do not show as a diff.
			"deactivated":    "false",
			"deactivated_at": "0",
			"error_message":  "",
		},
		HCL: `
		network_connectivity_config_id = "ncc_id"
		resource_id = "resource_id"
		group_id = "blob"
		`,
		ExpectedDiff: map[string]*terraform.ResourceAttrDiff{},
	}.ApplyNoError(t)
}
