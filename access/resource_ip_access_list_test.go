package access

// REST API: https://docs.databricks.com/dev-tools/api/latest/ip-access-list.html#operation/create-list

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/assert"
)

var (
	TestingId               = "234567"
	TestingLabel            = "Naughty"
	TestingListTypeString   = "BLOCK"
	TestingListType         = settings.ListType("BLOCK")
	TestingEnabled          = true
	TestingIpAddresses      = []string{"1.2.3.4", "1.2.4.0/24"}
	TestingIpAddressesState = []any{"1.2.3.4", "1.2.4.0/24"}
)

func TestIPACLCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/ip-access-lists",
				ExpectedRequest: settings.CreateIpAccessList{
					Label:       TestingLabel,
					ListType:    TestingListType,
					IpAddresses: TestingIpAddresses,
				},
				Response: settings.CreateIpAccessListResponse{
					IpAccessList: &settings.IpAccessListInfo{
						ListId:       TestingId,
						Label:        TestingLabel,
						ListType:     TestingListType,
						IpAddresses:  TestingIpAddresses,
						AddressCount: 2,
						CreatedAt:    87939234,
						CreatedBy:    1234556,
						UpdatedAt:    87939234,
						UpdatedBy:    1234556,
						Enabled:      TestingEnabled,
					},
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/ip-access-lists/" + TestingId + "?",
				Response: settings.CreateIpAccessListResponse{
					IpAccessList: &settings.IpAccessListInfo{
						ListId:       TestingId,
						Label:        TestingLabel,
						ListType:     TestingListType,
						IpAddresses:  TestingIpAddresses,
						AddressCount: 2,
						CreatedAt:    87939234,
						CreatedBy:    1234556,
						UpdatedAt:    87939234,
						UpdatedBy:    1234556,
						Enabled:      TestingEnabled,
					},
				},
			},
		},
		Resource: ResourceIPAccessList(),
		State: map[string]any{
			"label":        TestingLabel,
			"list_type":    TestingListTypeString,
			"ip_addresses": TestingIpAddressesState,
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, TestingId, d.Id())
	assert.Equal(t, TestingLabel, d.Get("label"))
	assert.Equal(t, TestingListTypeString, d.Get("list_type"))
	assert.Equal(t, TestingEnabled, d.Get("enabled"))
	assert.Equal(t, 2, d.Get("ip_addresses.#"))
}

func TestAPIACLCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/ip-access-lists",
				Response: common.APIErrorBody{
					ErrorCode: "RESOURCE_ALREADY_EXISTS",
					Message:   "IP access list with type (" + TestingListTypeString + ") and label (" + TestingLabel + ") already exists",
				},
				Status: 400,
			},
		},
		Resource: ResourceIPAccessList(),
		State: map[string]any{
			"label":        TestingLabel,
			"list_type":    TestingListTypeString,
			"ip_addresses": TestingIpAddressesState,
		},
		Create: true,
	}.Apply(t)
	assert.Error(t, err)
	qa.AssertErrorStartsWith(t, err, "IP access list with type")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestIPACLUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/ip-access-lists/" + TestingId + "?",
				Response: settings.CreateIpAccessListResponse{
					IpAccessList: &settings.IpAccessListInfo{
						ListId:       TestingId,
						Label:        TestingLabel,
						ListType:     TestingListType,
						IpAddresses:  TestingIpAddresses,
						AddressCount: 2,
						CreatedAt:    87939234,
						CreatedBy:    1234556,
						UpdatedAt:    87939234,
						UpdatedBy:    1234556,
						Enabled:      TestingEnabled,
					},
				},
			},
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.0/ip-access-lists/" + TestingId,
				ExpectedRequest: ipAccessListUpdateRequest{
					Label:       TestingLabel,
					ListType:    TestingListType,
					IpAddresses: TestingIpAddresses,
					Enabled:     TestingEnabled,
				},
				Response: settings.CreateIpAccessListResponse{
					IpAccessList: &settings.IpAccessListInfo{
						ListId:       TestingId,
						Label:        TestingLabel,
						ListType:     TestingListType,
						IpAddresses:  TestingIpAddresses,
						AddressCount: 2,
						CreatedAt:    87939234,
						CreatedBy:    1234556,
						UpdatedAt:    87939234,
						UpdatedBy:    1234556,
						Enabled:      TestingEnabled,
					},
				},
			},
		},
		Resource: ResourceIPAccessList(),
		State: map[string]any{
			"label":        TestingLabel,
			"list_type":    TestingListTypeString,
			"ip_addresses": TestingIpAddressesState,
		},
		Update: true,
		ID:     TestingId,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, TestingId, d.Id())
	assert.Equal(t, TestingLabel, d.Get("label"))
	assert.Equal(t, TestingListTypeString, d.Get("list_type"))
	assert.Equal(t, TestingEnabled, d.Get("enabled"))
	assert.Equal(t, 2, d.Get("ip_addresses.#"))
}

func TestIPACLUpdate_Error(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.0/ip-access-lists/" + TestingId,
				ExpectedRequest: settings.UpdateIpAccessList{
					Enabled: TestingEnabled,
				},
				Response: common.APIErrorBody{
					ErrorCode: "SERVER_ERROR",
					Message:   "Something unexpected happened",
				},
				Status: 500,
			},
		},
		Resource: ResourceIPAccessList(),
		Update:   true,
		ID:       TestingId,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Something unexpected")
}

func TestIPACLRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/ip-access-lists/" + TestingId + "?",
				Response: settings.FetchIpAccessListResponse{
					IpAccessList: &settings.IpAccessListInfo{
						ListId:       TestingId,
						Label:        TestingLabel,
						ListType:     TestingListType,
						IpAddresses:  TestingIpAddresses,
						AddressCount: 2,
						CreatedAt:    87939234,
						CreatedBy:    1234556,
						UpdatedAt:    87939234,
						UpdatedBy:    1234556,
						Enabled:      TestingEnabled,
					},
				},
			},
		},
		Resource: ResourceIPAccessList(),
		Read:     true,
		New:      true,
		ID:       TestingId,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, TestingId, d.Id())
	assert.Equal(t, TestingLabel, d.Get("label"))
	assert.Equal(t, TestingListTypeString, d.Get("list_type"))
	assert.Equal(t, TestingEnabled, d.Get("enabled"))
	assert.Equal(t, 2, d.Get("ip_addresses.#"))
}

func TestIPACLRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/ip-access-lists/" + TestingId + "?",
				Response: common.APIErrorBody{
					ErrorCode: "RESOURCE_DOES_NOT_EXIST",
					Message:   "Can't find an IP access list with id: " + TestingId + ".",
				},
				Status: 404,
			},
		},
		Resource: ResourceIPAccessList(),
		Read:     true,
		Removed:  true,
		ID:       TestingId,
	}.ApplyNoError(t)
}

func TestIPACLRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/ip-access-lists/" + TestingId + "?",
				Response: common.APIErrorBody{
					ErrorCode: "SERVER_ERROR",
					Message:   "Something unexpected happened",
				},
				Status: 500,
			},
		},
		Resource: ResourceIPAccessList(),
		Read:     true,
		ID:       TestingId,
	}.Apply(t)
	assert.Error(t, err)
	qa.AssertErrorStartsWith(t, err, "Something unexpected happened")
	assert.Equal(t, TestingId, d.Id(), "Id should not be empty for error reads")
}

func TestIPACLDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodDelete,
				Resource: fmt.Sprintf("/api/2.0/ip-access-lists/%s?", TestingId),
			},
		},
		Resource: ResourceIPAccessList(),
		Delete:   true,
		ID:       TestingId,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, TestingId, d.Id())
}

func TestIPACLDelete_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodDelete,
				Resource: fmt.Sprintf("/api/2.0/ip-access-lists/%s?", TestingId),
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_STATE",
					Message:   "Something went wrong",
				},
				Status: 400,
			},
		},
		Resource: ResourceIPAccessList(),
		Delete:   true,
		Removed:  true,
		ID:       TestingId,
	}.ExpectError(t, "Something went wrong")
}
