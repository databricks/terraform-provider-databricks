package mws

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/assert"
)

func TestResourceNetworkCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/networks",
				ExpectedRequest: Network{
					AccountID:        "abc",
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
		Resource: ResourceMwsNetworks(),
		HCL: `
		account_id = "abc"
		network_name = "Open Workers"
		security_group_ids = ["one", "two"]
		subnet_ids = ["three", "four"]
		vpc_id = "five"
		`,
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc/nid", d.Id())
}

func TestResourceNetworkCreate_GCP(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/networks",
				ExpectedRequest: Network{
					AccountID:   "abc",
					NetworkName: "Open Workers",
					GcpNetworkInfo: &GcpNetworkInfo{
						NetworkProjectId:   "project_a",
						VpcId:              "vpc_a",
						SubnetId:           "subnet_a",
						SubnetRegion:       "region_a",
					},
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
		Resource: ResourceMwsNetworks(),
		HCL: `
		account_id = "abc"
		network_name = "Open Workers"
		gcp_network_info {
			network_project_id = "project_a"
			vpc_id = "vpc_a"
			subnet_id = "subnet_a"
			subnet_region = "region_a"
        }
		`,
		Create: true,
	}.ApplyNoError(t)
}

func TestResourceNetworkCreate_GCPPsc(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/networks",
				ExpectedRequest: Network{
					AccountID:   "abc",
					NetworkName: "Open Workers",
					GcpNetworkInfo: &GcpNetworkInfo{
						NetworkProjectId:   "project_a",
						VpcId:              "vpc_a",
						SubnetId:           "subnet_a",
						SubnetRegion:       "region_a",
					},
					VPCEndpoints: &NetworkVPCEndpoints{
						RestAPI:           []string{"rest_api_endpoint"},
						DataplaneRelayAPI: []string{"dataplane_relay_endpoint"},
					},
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
		Resource: ResourceMwsNetworks(),
		HCL: `
		account_id = "abc"
		network_name = "Open Workers"
		gcp_network_info {
			network_project_id = "project_a"
			vpc_id = "vpc_a"
			subnet_id = "subnet_a"
			subnet_region = "region_a"
		}
        vpc_endpoints {
        	rest_api = ["rest_api_endpoint"]
        	dataplane_relay = ["dataplane_relay_endpoint"]
        }
		`,
		Create: true,
	}.ApplyNoError(t)
}

func TestResourceNetworkCreate_ConflictErrors(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{},
		Resource: ResourceMwsNetworks(),
		HCL: `
		account_id = "abc"
		network_name = "Open Workers"
		security_group_ids = ["one", "two"]
		subnet_ids = ["three", "four"]
		vpc_id = "five"
		vpc_endpoints {
			rest_api = ["a","b"]
			dataplane_relay = ["b", "c"]
		}
		gcp_network_info {
			network_project_id = "project_a"
			vpc_id = "vpc_a"
			subnet_id = "subnet_a"
			subnet_region = "region_a"
        }
		`,
		Create: true,
	}.Apply(t)
	assert.ErrorContains(t, err, "[gcp_network_info] Conflicting configuration arguments")
	assert.ErrorContains(t, err, "[security_group_ids] Invalid combination of arguments")
	assert.ErrorContains(t, err, "[subnet_ids] Invalid combination of arguments")
	assert.ErrorContains(t, err, "[vpc_id] Invalid combination of arguments")
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
		Resource: ResourceMwsNetworks(),
		State: map[string]any{
			"account_id":         "abc",
			"network_name":       "Open Workers",
			"security_group_ids": []any{"one", "two"},
			"subnet_ids":         []any{"three", "four"},
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
		Resource: ResourceMwsNetworks(),
		Read:     true,
		New:      true,
		ID:       "abc/nid",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc/nid", d.Id(), "Id should not be empty")
	assert.Equal(t, 0, d.Get("creation_time"))
	assert.Equal(t, "nid", d.Get("network_id"))
	assert.Equal(t, "Open Workers", d.Get("network_name"))
	assert.Equal(t, "five", d.Get("vpc_id"))
	assert.Equal(t, 789, d.Get("workspace_id"))
}

func TestResourceNetworkRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
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
		Resource: ResourceMwsNetworks(),
		Read:     true,
		Removed:  true,
		ID:       "abc/nid",
	}.ApplyNoError(t)
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
		Resource: ResourceMwsNetworks(),
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
		Resource: ResourceMwsNetworks(),
		Delete:   true,
		ID:       "abc/nid",
	}.Apply(t)
	assert.NoError(t, err)
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
		Resource: ResourceMwsNetworks(),
		Delete:   true,
		ID:       "abc/nid",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc/nid", d.Id())
}
