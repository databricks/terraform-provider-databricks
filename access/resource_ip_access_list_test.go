package access

// REST API: https://docs.databricks.com/dev-tools/api/latest/ip-access-list.html#operation/create-list

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			api := w.GetMockIpAccessListsAPI().EXPECT()
			api.Create(mock.Anything, settings.CreateIpAccessList{
				Label:       TestingLabel,
				ListType:    TestingListType,
				IpAddresses: TestingIpAddresses,
			}).Return(&settings.CreateIpAccessListResponse{
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
			}, nil)
			api.GetByIpAccessListId(mock.Anything, TestingId).Return(&settings.FetchIpAccessListResponse{
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
			}, nil)
		},
		Resource: ResourceIPAccessList(),
		State: map[string]any{
			"label":        TestingLabel,
			"list_type":    TestingListTypeString,
			"ip_addresses": TestingIpAddressesState,
		},
		Create: true,
	}.ApplyAndExpectData(t, map[string]any{
		"id":           TestingId,
		"label":        TestingLabel,
		"list_type":    TestingListTypeString,
		"enabled":      TestingEnabled,
		"ip_addresses": TestingIpAddressesState,
	})
}

func TestAPIACLCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockIpAccessListsAPI().EXPECT().
				Create(mock.Anything, settings.CreateIpAccessList{
					Label:       TestingLabel,
					ListType:    TestingListType,
					IpAddresses: TestingIpAddresses,
				}).Return(nil, &apierr.APIError{
				ErrorCode:  "RESOURCE_ALREADY_EXISTS",
				StatusCode: 400,
				Message:    "IP access list with type (" + TestingListTypeString + ") and label (" + TestingLabel + ") already exists",
			})
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
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			api := w.GetMockIpAccessListsAPI().EXPECT()
			api.GetByIpAccessListId(mock.Anything, TestingId).Return(&settings.FetchIpAccessListResponse{
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
			}, nil)
			api.Update(mock.Anything, settings.UpdateIpAccessList{
				IpAccessListId: TestingId,
				Label:          TestingLabel,
				ListType:       TestingListType,
				IpAddresses:    TestingIpAddresses,
				Enabled:        TestingEnabled,
			}).Return(nil)
		},
		Resource: ResourceIPAccessList(),
		State: map[string]any{
			"label":        TestingLabel,
			"list_type":    TestingListTypeString,
			"ip_addresses": TestingIpAddressesState,
		},
		Update: true,
		ID:     TestingId,
	}.ApplyAndExpectData(t, map[string]any{
		"id":           TestingId,
		"label":        TestingLabel,
		"list_type":    TestingListTypeString,
		"enabled":      TestingEnabled,
		"ip_addresses": TestingIpAddressesState,
	})
}

func TestIPACLUpdate_Error(t *testing.T) {
	_, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockIpAccessListsAPI().EXPECT().
				Update(mock.Anything, mock.MatchedBy(func(req settings.UpdateIpAccessList) bool {
					return req.IpAccessListId == TestingId && req.Enabled == TestingEnabled
				})).Return(&apierr.APIError{
				ErrorCode:  "SERVER_ERROR",
				StatusCode: 500,
				Message:    "Something unexpected happened",
			})
		},
		Resource: ResourceIPAccessList(),
		Update:   true,
		ID:       TestingId,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Something unexpected")
}

func TestIPACLRead(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockIpAccessListsAPI().EXPECT().
				GetByIpAccessListId(mock.Anything, TestingId).
				Return(&settings.FetchIpAccessListResponse{
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
				}, nil)
		},
		Resource: ResourceIPAccessList(),
		Read:     true,
		New:      true,
		ID:       TestingId,
	}.ApplyAndExpectData(t, map[string]any{
		"id":           TestingId,
		"label":        TestingLabel,
		"list_type":    TestingListTypeString,
		"enabled":      TestingEnabled,
		"ip_addresses": TestingIpAddressesState,
	})
}

func TestIPACLRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockIpAccessListsAPI().EXPECT().
				GetByIpAccessListId(mock.Anything, TestingId).
				Return(nil, &apierr.APIError{
					ErrorCode:  "RESOURCE_DOES_NOT_EXIST",
					StatusCode: 404,
					Message:    "Can't find an IP access list with id: " + TestingId + ".",
				})
		},
		Resource: ResourceIPAccessList(),
		Read:     true,
		Removed:  true,
		ID:       TestingId,
	}.ApplyNoError(t)
}

func TestIPACLRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockIpAccessListsAPI().EXPECT().
				GetByIpAccessListId(mock.Anything, TestingId).
				Return(nil, &apierr.APIError{
					ErrorCode:  "SERVER_ERROR",
					StatusCode: 500,
					Message:    "Something unexpected happened",
				})
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
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockIpAccessListsAPI().EXPECT().
				DeleteByIpAccessListId(mock.Anything, TestingId).
				Return(nil)
		},
		Resource: ResourceIPAccessList(),
		Delete:   true,
		ID:       TestingId,
	}.ApplyAndExpectData(t, map[string]any{
		"id": TestingId,
	})
}

func TestIPACLDelete_Error(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockIpAccessListsAPI().EXPECT().
				DeleteByIpAccessListId(mock.Anything, TestingId).
				Return(&apierr.APIError{
					ErrorCode:  "INVALID_STATE",
					StatusCode: 400,
					Message:    "Something went wrong",
				})
		},
		Resource: ResourceIPAccessList(),
		Delete:   true,
		Removed:  true,
		ID:       TestingId,
	}.ExpectError(t, "Something went wrong")
}
