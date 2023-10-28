package mws

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getTestNcc() NetworkConnectivityConfig {
	return NetworkConnectivityConfig{
		AccountID:                   "abc",
		Name:                        "ncc_name",
		Region:                      "ar",
		NetworkConnectivityConfigID: "ncc_id",
	}
}

func TestResourceNCCCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/network-connectivity-configs",
				ExpectedRequest: NetworkConnectivityConfig{
					AccountID: "abc",
					Name:      "ncc_name",
					Region:    "ar",
				},
				Response: NetworkConnectivityConfig{
					NetworkConnectivityConfigID: "ncc_id",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/network-connectivity-configs/ncc_id",
				Response: getTestNcc(),
			},
		},
		Resource: ResourceMwsNetworkConnectivityConfig(),
		HCL: `
		account_id = "abc"
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
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/network-connectivity-configs",
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "error",
				},
				Status: 400,
			},
		},
		Resource: ResourceMwsNetworkConnectivityConfig(),
		HCL: `
		account_id = "abc"
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
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/network-connectivity-configs/ncc_id",
				Response: getTestNcc(),
			},
		},
		Resource: ResourceMwsNetworkConnectivityConfig(),
		Read:     true,
		New:      true,
		ID:       "abc/ncc_id",
	}.ApplyAndExpectData(t, map[string]any{
		"id":                             "abc/ncc_id",
		"account_id":                     "abc",
		"name":                           "ncc_name",
		"region":                         "ar",
		"network_connectivity_config_id": "ncc_id",
	})
}

func TestResourceNCCRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/network-connectivity-configs/ncc_id",
				Response: apierr.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Item not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceMwsNetworkConnectivityConfig(),
		Read:     true,
		Removed:  true,
		ID:       "abc/ncc_id",
	}.ApplyNoError(t)
}

func TestResourceNCCRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/network-connectivity-configs/ncc_id",
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "error",
				},
				Status: 400,
			},
		},
		Resource: ResourceMwsNetworkConnectivityConfig(),
		Read:     true,
		ID:       "abc/ncc_id",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "error")
	assert.Equal(t, "abc/ncc_id", d.Id())
}

func TestResourceNCCDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/accounts/abc/network-connectivity-configs/ncc_id",
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/network-connectivity-configs/ncc_id",
				Response: getTestNcc(),
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/network-connectivity-configs/ncc_id",
				Response: apierr.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceMwsNetworkConnectivityConfig(),
		Delete:   true,
		ID:       "abc/ncc_id",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc/ncc_id", d.Id())
}

func TestResourceNCCDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/accounts/abc/network-connectivity-configs/ncc_id",
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "error",
				},
				Status: 400,
			},
		},
		Resource: ResourceMwsNetworkConnectivityConfig(),
		ID:       "abc/ncc_id",
		HCL: `
		account_id = "abc"
		name = "ncc_name"
		region = "ar"
		network_connectivity_config_id = "ncc_id"
		`,
		Delete: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "error")
	assert.Equal(t, "abc/ncc_id", d.Id())
}

func TestResourceNCCList(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/network-connectivity-configs",
			Response: []NetworkConnectivityConfig{getTestNcc()},
		},
	})
	require.NoError(t, err)
	defer server.Close()

	l, err := NewNetworkConnectivityConfigAPI(context.Background(), client).List("abc")
	require.NoError(t, err)
	assert.Len(t, l, 1)
}
