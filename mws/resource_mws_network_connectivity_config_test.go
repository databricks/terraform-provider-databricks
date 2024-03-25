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

func TestResourceNCCCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			e := a.GetMockNetworkConnectivityAPI().EXPECT()
			e.CreateNetworkConnectivityConfiguration(mock.Anything, settings.CreateNetworkConnectivityConfigRequest{
				Name:   "ncc_name",
				Region: "ar",
			}).Return(getTestNcc(), nil)
			e.GetNetworkConnectivityConfigurationByNetworkConnectivityConfigId(mock.Anything, "ncc_id").Return(getTestNcc(), nil)
		},
		Resource:  ResourceMwsNetworkConnectivityConfig(),
		AccountID: "abc",
		HCL: `
		name = "ncc_name"
		region = "ar"
		`,
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc/ncc_id", d.Id())
}

func TestResourceNCCCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			e := a.GetMockNetworkConnectivityAPI().EXPECT()
			e.CreateNetworkConnectivityConfiguration(mock.Anything, settings.CreateNetworkConnectivityConfigRequest{
				Name:   "ncc_name",
				Region: "ar",
			}).Return(nil, &apierr.APIError{Message: "error"})
		},
		Resource:  ResourceMwsNetworkConnectivityConfig(),
		AccountID: "abc",
		HCL: `
		name = "ncc_name"
		region = "ar"
		`,
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "error")
	assert.Equal(t, "", d.Id())
}

func TestResourceNCCRead(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			e := a.GetMockNetworkConnectivityAPI().EXPECT()
			e.GetNetworkConnectivityConfigurationByNetworkConnectivityConfigId(mock.Anything, "ncc_id").Return(getTestNcc(), nil)
		},
		Resource:  ResourceMwsNetworkConnectivityConfig(),
		AccountID: "abc",
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

func TestResourceNCCRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			e := a.GetMockNetworkConnectivityAPI().EXPECT()
			e.GetNetworkConnectivityConfigurationByNetworkConnectivityConfigId(mock.Anything, "ncc_id").Return(nil, &apierr.APIError{Message: "error"})
		},
		Resource:  ResourceMwsNetworkConnectivityConfig(),
		AccountID: "abc",
		Read:      true,
		ID:        "abc/ncc_id",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "error")
	assert.Equal(t, "abc/ncc_id", d.Id())
}

func TestResourceNCCDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			e := a.GetMockNetworkConnectivityAPI().EXPECT()
			e.DeleteNetworkConnectivityConfigurationByNetworkConnectivityConfigId(mock.Anything, "ncc_id").Return(nil)
		},
		Resource:  ResourceMwsNetworkConnectivityConfig(),
		AccountID: "abc",
		Delete:    true,
		ID:        "abc/ncc_id",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc/ncc_id", d.Id())
}

func TestResourceNCCDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			e := a.GetMockNetworkConnectivityAPI().EXPECT()
			e.DeleteNetworkConnectivityConfigurationByNetworkConnectivityConfigId(mock.Anything, "ncc_id").Return(&apierr.APIError{Message: "error"})
		},
		Resource:  ResourceMwsNetworkConnectivityConfig(),
		AccountID: "abc",
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
