package access

// REST API: https://docs.databricks.com/dev-tools/api/latest/ip-access-list.html#operation/create-list

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	qa.ResourceFixture{
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
	}.ApplyAndExpectData(t,
		map[string]any{
			"id":             TestingId,
			"label":          TestingLabel,
			"list_type":      TestingListTypeString,
			"enabled":        TestingEnabled,
			"ip_addresses.#": 2,
		})
}

func TestAPIACLCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/ip-access-lists",
				Response: apierr.APIErrorBody{
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
	qa.ResourceFixture{
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
	}.ApplyAndExpectData(t,
		map[string]any{
			"id":             TestingId,
			"label":          TestingLabel,
			"list_type":      TestingListTypeString,
			"enabled":        TestingEnabled,
			"ip_addresses.#": 2,
		})
}

func TestIPACLUpdate_Error(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.0/ip-access-lists/" + TestingId,
				ExpectedRequest: ipAccessListUpdateRequest{
					Enabled: TestingEnabled,
				},
				Response: apierr.APIErrorBody{
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
	qa.ResourceFixture{
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
	}.ApplyAndExpectData(t,
		map[string]any{
			"id":             TestingId,
			"label":          TestingLabel,
			"list_type":      TestingListTypeString,
			"enabled":        TestingEnabled,
			"ip_addresses.#": 2,
		})
}

func TestIPACLRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/ip-access-lists/" + TestingId + "?",
				Response: apierr.APIErrorBody{
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
				Response: apierr.APIErrorBody{
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
				Response: apierr.APIErrorBody{
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

func TestListIpAccessLists(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/ip-access-lists",
			Response: map[string]any{},
		},
	})
	require.NoError(t, err)

	w, err := client.WorkspaceClient()
	require.NoError(t, err)

	defer server.Close()
	require.NoError(t, err)

	ctx := context.Background()
	ipLists, err := w.IpAccessLists.Impl().List(ctx)

	require.NoError(t, err)
	assert.Equal(t, 0, len(ipLists.IpAccessLists))
}

func TestAccIPACLCreate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/accounts/100/ip-access-lists",
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
						Enabled:      false,
					},
				},
			},
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.0/accounts/100/ip-access-lists/" + TestingId,
				ExpectedRequest: ipAccessListUpdateRequest{
					Label:       TestingLabel,
					ListType:    TestingListType,
					IpAddresses: TestingIpAddresses,
					Enabled:     TestingEnabled,
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/accounts/100/ip-access-lists/" + TestingId + "?",
				Response: settings.GetIpAccessListResponse{
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
		Resource:  ResourceIPAccessList(),
		AccountID: "100",
		State: map[string]any{
			"label":        TestingLabel,
			"list_type":    TestingListTypeString,
			"ip_addresses": TestingIpAddressesState,
		},
		Create: true,
	}.ApplyAndExpectData(t,
		map[string]any{
			"id":             TestingId,
			"label":          TestingLabel,
			"list_type":      TestingListTypeString,
			"enabled":        TestingEnabled,
			"ip_addresses.#": 2,
		})
}

func TestAccIPACLCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/accounts/100/ip-access-lists",
				Response: apierr.APIErrorBody{
					ErrorCode: "RESOURCE_ALREADY_EXISTS",
					Message:   "IP access list with type (" + TestingListTypeString + ") and label (" + TestingLabel + ") already exists",
				},
				Status: 400,
			},
		},
		Resource:  ResourceIPAccessList(),
		AccountID: "100",
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

func TestAccIPACLUpdate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPatch,
				Resource: fmt.Sprintf("/api/2.0/accounts/100/ip-access-lists/%s", TestingId),
				ExpectedRequest: ipAccessListUpdateRequest{
					Label:       TestingLabel,
					ListType:    TestingListType,
					IpAddresses: TestingIpAddresses,
					Enabled:     TestingEnabled,
				},
			},
			{
				Method:   http.MethodGet,
				Resource: fmt.Sprintf("/api/2.0/accounts/100/ip-access-lists/%s?", TestingId),
				Response: settings.GetIpAccessListResponse{
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
		Update:    true,
		ID:        TestingId,
		AccountID: "100",
	}.ApplyAndExpectData(t,
		map[string]any{
			"id":             TestingId,
			"label":          TestingLabel,
			"list_type":      TestingListTypeString,
			"enabled":        TestingEnabled,
			"ip_addresses.#": 2,
		})
}

func TestAccIPACLUpdate_Error(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.0/accounts/100/ip-access-lists/" + TestingId,
				ExpectedRequest: ipAccessListUpdateRequest{
					Enabled: TestingEnabled,
				},
				Response: apierr.APIErrorBody{
					ErrorCode: "SERVER_ERROR",
					Message:   "Something unexpected happened",
				},
				Status: 500,
			},
		},
		Resource:  ResourceIPAccessList(),
		Update:    true,
		ID:        TestingId,
		AccountID: "100",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Something unexpected")
}

func TestAccIPACLRead(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/accounts/100/ip-access-lists/" + TestingId + "?",
				Response: settings.GetIpAccessListResponse{
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
		Resource:  ResourceIPAccessList(),
		Read:      true,
		New:       true,
		ID:        TestingId,
		AccountID: "100",
	}.ApplyAndExpectData(t,
		map[string]any{
			"id":             TestingId,
			"label":          TestingLabel,
			"list_type":      TestingListTypeString,
			"enabled":        TestingEnabled,
			"ip_addresses.#": 2,
		})
}

func TestAccIPACLRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/accounts/100/ip-access-lists/" + TestingId + "?",
				Response: apierr.APIErrorBody{
					ErrorCode: "RESOURCE_DOES_NOT_EXIST",
					Message:   "Can't find an IP access list with id: " + TestingId + ".",
				},
				Status: 404,
			},
		},
		Resource:  ResourceIPAccessList(),
		Read:      true,
		Removed:   true,
		ID:        TestingId,
		AccountID: "100",
	}.ApplyNoError(t)
}

func TestAccIPACLRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/accounts/100/ip-access-lists/" + TestingId + "?",
				Response: apierr.APIErrorBody{
					ErrorCode: "SERVER_ERROR",
					Message:   "Something unexpected happened",
				},
				Status: 500,
			},
		},
		Resource:  ResourceIPAccessList(),
		Read:      true,
		ID:        TestingId,
		AccountID: "100",
	}.Apply(t)
	assert.Error(t, err)
	qa.AssertErrorStartsWith(t, err, "Something unexpected happened")
	assert.Equal(t, TestingId, d.Id(), "Id should not be empty for error reads")
}

func TestAccIPACLDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodDelete,
				Resource: fmt.Sprintf("/api/2.0/accounts/100/ip-access-lists/%s?", TestingId),
			},
		},
		Resource:  ResourceIPAccessList(),
		Delete:    true,
		AccountID: "100",
		ID:        TestingId,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, TestingId, d.Id())
}

func TestAccIPACLDelete_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodDelete,
				Resource: fmt.Sprintf("/api/2.0/accounts/100/ip-access-lists/%s?", TestingId),
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_STATE",
					Message:   "Something went wrong",
				},
				Status: 400,
			},
		},
		Resource:  ResourceIPAccessList(),
		Delete:    true,
		Removed:   true,
		AccountID: "100",
		ID:        TestingId,
	}.ExpectError(t, "Something went wrong")
}
