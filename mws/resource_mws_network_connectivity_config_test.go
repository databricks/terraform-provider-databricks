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

func getTestNcc() *settings.NetworkConnectivityConfiguration {
	return &settings.NetworkConnectivityConfiguration{
		AccountId:                   "abc",
		Name:                        "ncc_name",
		Region:                      "ar",
		NetworkConnectivityConfigId: "ncc_id",
	}
}

func TestResourceNccCreate(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			e := a.GetMockNetworkConnectivityAPI().EXPECT()
			e.CreateNetworkConnectivityConfiguration(mock.Anything, settings.CreateNetworkConnectivityConfigRequest{
				NetworkConnectivityConfig: settings.CreateNetworkConnectivityConfiguration{
					Name:   "ncc_name",
					Region: "ar",
				},
			}).Return(getTestNcc(), nil)
			e.GetNetworkConnectivityConfigurationByNetworkConnectivityConfigId(mock.Anything, "ncc_id").Return(getTestNcc(), nil)
		},
		Resource:  ResourceMwsNetworkConnectivityConfig(),
		AccountID: "abc",
		Host:      "https://accounts.cloud.databricks.com",
		HCL: `
		name = "ncc_name"
		region = "ar"
		`,
		Create: true,
	}.ApplyAndExpectData(t, map[string]any{"id": "abc/ncc_id"})
}

func TestResourceNccCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			e := a.GetMockNetworkConnectivityAPI().EXPECT()
			e.CreateNetworkConnectivityConfiguration(mock.Anything, settings.CreateNetworkConnectivityConfigRequest{
				NetworkConnectivityConfig: settings.CreateNetworkConnectivityConfiguration{
					Name:   "ncc_name",
					Region: "ar",
				},
			}).Return(nil, &apierr.APIError{Message: "error"})
		},
		Resource:  ResourceMwsNetworkConnectivityConfig(),
		AccountID: "abc",
		Host:      "https://accounts.cloud.databricks.com",
		HCL: `
		name = "ncc_name"
		region = "ar"
		`,
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "error")
	assert.Equal(t, "", d.Id())
}

func TestResourceNccRead(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockNetworkConnectivityAPI().EXPECT().GetNetworkConnectivityConfigurationByNetworkConnectivityConfigId(mock.Anything, "ncc_id").Return(getTestNcc(), nil)
		},
		Resource:  ResourceMwsNetworkConnectivityConfig(),
		AccountID: "abc",
		Host:      "https://accounts.cloud.databricks.com",
		Read:      true,
		New:       true,
		ID:        "abc/ncc_id",
	}.ApplyAndExpectData(t, map[string]any{
		"id":                             "abc/ncc_id",
		"account_id":                     "abc",
		"name":                           "ncc_name",
		"region":                         "ar",
		"network_connectivity_config_id": "ncc_id",
	})
}

func TestResourceNccRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockNetworkConnectivityAPI().EXPECT().GetNetworkConnectivityConfigurationByNetworkConnectivityConfigId(mock.Anything, "ncc_id").Return(nil, &apierr.APIError{Message: "error"})
		},
		Resource:  ResourceMwsNetworkConnectivityConfig(),
		AccountID: "abc",
		Host:      "https://accounts.cloud.databricks.com",
		Read:      true,
		ID:        "abc/ncc_id",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "error")
	assert.Equal(t, "abc/ncc_id", d.Id())
}

func TestResourceNccDelete(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockNetworkConnectivityAPI().EXPECT().DeleteNetworkConnectivityConfigurationByNetworkConnectivityConfigId(mock.Anything, "ncc_id").Return(nil)
		},
		Resource:  ResourceMwsNetworkConnectivityConfig(),
		AccountID: "abc",
		Host:      "https://accounts.cloud.databricks.com",
		Delete:    true,
		ID:        "abc/ncc_id",
	}.ApplyAndExpectData(t, map[string]any{"id": "abc/ncc_id"})
}

func TestResourceNccDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockNetworkConnectivityAPI().EXPECT().DeleteNetworkConnectivityConfigurationByNetworkConnectivityConfigId(mock.Anything, "ncc_id").Return(&apierr.APIError{Message: "error"})
		},
		Resource:  ResourceMwsNetworkConnectivityConfig(),
		AccountID: "abc",
		Host:      "https://accounts.cloud.databricks.com",
		ID:        "abc/ncc_id",
		HCL: `
		name = "ncc_name"
		region = "ar"
		`,
		Delete: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "error")
	assert.Equal(t, "abc/ncc_id", d.Id())
}

func TestResourceNccDelete_RetriesOnStillInUseRules(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			e := a.GetMockNetworkConnectivityAPI().EXPECT()
			call1 := e.DeleteNetworkConnectivityConfigurationByNetworkConnectivityConfigId(
				mock.Anything, "ncc_id",
			).Return(&apierr.APIError{
				Message: "Network Connectivity Config ncc_id is unable to be deleted because it has one or more private endpoint rules.",
			})
			call1.Repeatability = 1
			e.DeleteNetworkConnectivityConfigurationByNetworkConnectivityConfigId(
				mock.Anything, "ncc_id",
			).Return(nil)
		},
		Resource:  ResourceMwsNetworkConnectivityConfig(),
		AccountID: "abc",
		Host:      "https://accounts.cloud.databricks.com",
		InstanceState: map[string]string{
			"account_id":                     "abc",
			"network_connectivity_config_id": "ncc_id",
			"name":                           "ncc_name",
			"region":                         "ar",
		},
		ID:     "abc/ncc_id",
		Delete: true,
	}.ApplyNoError(t)
}

func TestResourceNccDelete_RetriesOnStillInUseWorkspaces(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			e := a.GetMockNetworkConnectivityAPI().EXPECT()
			call1 := e.DeleteNetworkConnectivityConfigurationByNetworkConnectivityConfigId(
				mock.Anything, "ncc_id",
			).Return(&apierr.APIError{
				Message: "Network Connectivity Config ncc_id is unable to be deleted because it is attached to one or more workspaces: 12345",
			})
			call1.Repeatability = 1
			e.DeleteNetworkConnectivityConfigurationByNetworkConnectivityConfigId(
				mock.Anything, "ncc_id",
			).Return(nil)
		},
		Resource:  ResourceMwsNetworkConnectivityConfig(),
		AccountID: "abc",
		Host:      "https://accounts.cloud.databricks.com",
		InstanceState: map[string]string{
			"account_id":                     "abc",
			"network_connectivity_config_id": "ncc_id",
			"name":                           "ncc_name",
			"region":                         "ar",
		},
		ID:     "abc/ncc_id",
		Delete: true,
	}.ApplyNoError(t)
}

func TestResourceNccDelete_NonRetryableErrorFailsFast(t *testing.T) {
	_, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			e := a.GetMockNetworkConnectivityAPI().EXPECT()
			e.DeleteNetworkConnectivityConfigurationByNetworkConnectivityConfigId(
				mock.Anything, "ncc_id",
			).Return(&apierr.APIError{Message: "permission denied"}).Once()
		},
		Resource:  ResourceMwsNetworkConnectivityConfig(),
		AccountID: "abc",
		Host:      "https://accounts.cloud.databricks.com",
		InstanceState: map[string]string{
			"account_id":                     "abc",
			"network_connectivity_config_id": "ncc_id",
			"name":                           "ncc_name",
			"region":                         "ar",
		},
		ID:     "abc/ncc_id",
		Delete: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "permission denied")
}
