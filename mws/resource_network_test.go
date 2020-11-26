package mws

import (
	"context"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"

	"github.com/stretchr/testify/assert"
)

func TestMWSNetworks(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	acctID := qa.GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID")
	client := common.CommonEnvironmentClient()
	ctx := context.Background()
	networksAPI := NewNetworksAPI(ctx, client)
	networksList, err := networksAPI.List(acctID)
	assert.NoError(t, err, err)
	t.Log(networksList)

	network := Network{
		AccountID:        acctID,
		NetworkName:      qa.RandomName(),
		VPCID:            "vpc-0abcdef1234567890",
		SubnetIds:        []string{"subnet-0123456789abcdef0", "subnet-0fedcba9876543210"},
		SecurityGroupIds: []string{"sg-0a1b2c3d4e5f6a7b8"},
	}
	err = networksAPI.Create(&network)
	assert.NoError(t, err, err)
	defer func() {
		err = networksAPI.Delete(acctID, network.NetworkID)
		assert.NoError(t, err, err)
	}()

	myNetworkFull, err := networksAPI.Read(acctID, network.NetworkID)
	assert.NoError(t, err, err)
	t.Log(myNetworkFull)
}

func TestResourceNetworkCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/networks",
				ExpectedRequest: Network{
					SecurityGroupIds: []string{"one", "two"},
					NetworkName:      "Open Workers",
					VPCID:            "five",
					SubnetIds:        []string{"four", "three"},
				},
				Response: Network{
					AccountID: "abc",
					NetworkID: "nid",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/networks/nid",
				Response: Network{
					NetworkID:        "nid",
					SecurityGroupIds: []string{"one", "two"},
					NetworkName:      "Open Workers",
					VPCID:            "five",
					SubnetIds:        []string{"four", "three"},
				},
			},
		},
		Resource: ResourceNetwork(),
		HCL: `
		account_id = "abc"
		network_name = "Open Workers"
		security_group_ids = ["one", "two"]
		subnet_ids = ["three", "four"]
		vpc_id = "five"
		`,
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/nid", d.Id())
}

func TestResourceNetworkCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/networks",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceNetwork(),
		State: map[string]interface{}{
			"account_id":         "abc",
			"network_name":       "Open Workers",
			"security_group_ids": []interface{}{"one", "two"},
			"subnet_ids":         []interface{}{"three", "four"},
			"vpc_id":             "five",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceNetworkRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/networks/nid",
				Response: Network{
					NetworkID:        "nid",
					SecurityGroupIds: []string{"one", "two"},
					NetworkName:      "Open Workers",
					VPCID:            "five",
					SubnetIds:        []string{"four", "three"},
					WorkspaceID:      789,
				},
			},
		},
		Resource: ResourceNetwork(),
		Read:     true,
		ID:       "abc/nid",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/nid", d.Id(), "Id should not be empty")
	assert.Equal(t, 0, d.Get("creation_time"))
	assert.Equal(t, "nid", d.Get("network_id"))
	assert.Equal(t, "Open Workers", d.Get("network_name"))
	assert.Equal(t, "five", d.Get("vpc_id"))
	assert.Equal(t, 789, d.Get("workspace_id"))
}

func TestResourceNetworkRead_NotFound(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/networks/nid",
				Response: common.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Item not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceNetwork(),
		Read:     true,
		ID:       "abc/nid",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id(), "Id should be empty for missing resources")
}

func TestResourceNetworkRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/networks/nid",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceNetwork(),
		Read:     true,
		ID:       "abc/nid",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc/nid", d.Id(), "Id should not be empty for error reads")
}

func TestResourceNetworkDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/accounts/abc/networks/nid",
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/networks/nid",
				Response: Network{
					AccountID:   "abc",
					NetworkID:   "nid",
					NetworkName: "Open Workers",
					VPCID:       "five",
					VPCStatus:   "SOMETHING",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/networks/nid",
				Response: common.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Yes, it's not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceNetwork(),
		Delete:   true,
		ID:       "abc/nid",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/nid", d.Id())
}

func TestResourceNetworkDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/accounts/abc/networks/nid",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceNetwork(),
		Delete:   true,
		ID:       "abc/nid",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc/nid", d.Id())
}
