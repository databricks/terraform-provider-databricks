package access

// REST API: https://docs.databricks.com/dev-tools/api/latest/ip-access-list.html#operation/create-list

import (
	"errors"
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
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
	TestingTimestamp        = int64(87939234)
	TestingUserId           = int64(1234556)
	TestDataExpected        = map[string]any{
		"id":             TestingId,
		"label":          TestingLabel,
		"list_type":      TestingListTypeString,
		"enabled":        TestingEnabled,
		"ip_addresses.#": 2,
	}
	TestIpAccessList = &settings.IpAccessListInfo{
		ListId:       TestingId,
		Label:        TestingLabel,
		ListType:     TestingListType,
		IpAddresses:  TestingIpAddresses,
		AddressCount: 2,
		CreatedAt:    TestingTimestamp,
		CreatedBy:    TestingUserId,
		UpdatedAt:    TestingTimestamp,
		UpdatedBy:    TestingUserId,
		Enabled:      TestingEnabled,
	}
)

func TestIPACLCreate(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockIpAccessListsAPI().EXPECT()
			e.Create(mock.Anything, settings.CreateIpAccessList{
				Label:       TestingLabel,
				ListType:    TestingListType,
				IpAddresses: TestingIpAddresses,
			}).Return(&settings.CreateIpAccessListResponse{
				IpAccessList: TestIpAccessList,
			}, nil)
		},
		Resource: ResourceIPAccessList(),
		State: map[string]any{
			"label":        TestingLabel,
			"list_type":    TestingListTypeString,
			"ip_addresses": TestingIpAddressesState,
		},
		Create: true,
	}.ApplyAndExpectData(t, TestDataExpected)
	qa.ResourceFixture{
		MockAccountClientFunc: func(w *mocks.MockAccountClient) {
			e := w.GetMockAccountIpAccessListsAPI().EXPECT()
			e.Create(mock.Anything, settings.CreateIpAccessList{
				Label:       TestingLabel,
				ListType:    TestingListType,
				IpAddresses: TestingIpAddresses,
			}).Return(&settings.CreateIpAccessListResponse{
				IpAccessList: TestIpAccessList,
			}, nil)
			e.GetByIpAccessListId(mock.Anything, TestingId).Return(&settings.GetIpAccessListResponse{
				IpAccessList: TestIpAccessList,
			}, nil)
			e.Update(mock.Anything, settings.UpdateIpAccessList{
				IpAccessListId: TestingId,
				Enabled:        TestingEnabled,
				Label:          TestingLabel,
				ListType:       TestingListType,
				IpAddresses:    TestingIpAddresses,
			}).Return(nil)
		},
		Resource: ResourceIPAccessList(),
		State: map[string]any{
			"label":        TestingLabel,
			"list_type":    TestingListTypeString,
			"ip_addresses": TestingIpAddressesState,
		},
		AccountID: "100",
		Create:    true,
	}.ApplyAndExpectData(t, TestDataExpected)
}

func TestAPIACLCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockIpAccessListsAPI().EXPECT()
			e.Create(mock.Anything, settings.CreateIpAccessList{
				Label:       TestingLabel,
				ListType:    TestingListType,
				IpAddresses: TestingIpAddresses,
			}).Return(nil,
				errors.New("IP access list with type ("+TestingListTypeString+") and label ("+TestingLabel+") already exists"))
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
	d, err = qa.ResourceFixture{
		MockAccountClientFunc: func(w *mocks.MockAccountClient) {
			e := w.GetMockAccountIpAccessListsAPI().EXPECT()
			e.Create(mock.Anything, settings.CreateIpAccessList{
				Label:       TestingLabel,
				ListType:    TestingListType,
				IpAddresses: TestingIpAddresses,
			}).Return(nil,
				errors.New("IP access list with type ("+TestingListTypeString+") and label ("+TestingLabel+") already exists"))
		},
		Resource: ResourceIPAccessList(),
		State: map[string]any{
			"label":        TestingLabel,
			"list_type":    TestingListTypeString,
			"ip_addresses": TestingIpAddressesState,
		},
		AccountID: "100",
		Create:    true,
	}.Apply(t)
	assert.Error(t, err)
	qa.AssertErrorStartsWith(t, err, "IP access list with type")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestIPACLUpdate(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockIpAccessListsAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateIpAccessList{
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
	}.ApplyAndExpectData(t, TestDataExpected)
	qa.ResourceFixture{
		MockAccountClientFunc: func(w *mocks.MockAccountClient) {
			e := w.GetMockAccountIpAccessListsAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateIpAccessList{
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
		Update:    true,
		AccountID: "100",
		ID:        TestingId,
	}.ApplyAndExpectData(t, TestDataExpected)
}

func TestIPACLUpdate_Error(t *testing.T) {
	_, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockIpAccessListsAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateIpAccessList{
				IpAccessListId: TestingId,
				Enabled:        TestingEnabled,
			}).Return(errors.New("Something unexpected happened"))
		},
		Resource: ResourceIPAccessList(),
		Update:   true,
		ID:       TestingId,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Something unexpected")
	_, err = qa.ResourceFixture{
		MockAccountClientFunc: func(w *mocks.MockAccountClient) {
			e := w.GetMockAccountIpAccessListsAPI().EXPECT()
			e.Update(mock.Anything, settings.UpdateIpAccessList{
				IpAccessListId: TestingId,
				Enabled:        TestingEnabled,
			}).Return(errors.New("Something unexpected happened"))
		},
		Resource:  ResourceIPAccessList(),
		Update:    true,
		AccountID: "100",
		ID:        TestingId,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Something unexpected")
}

func TestIPACLRead(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockIpAccessListsAPI().EXPECT()
			e.GetByIpAccessListId(mock.Anything, TestingId).Return(&settings.FetchIpAccessListResponse{
				IpAccessList: TestIpAccessList,
			}, nil)
		},
		Resource: ResourceIPAccessList(),
		Read:     true,
		New:      true,
		ID:       TestingId,
	}.ApplyAndExpectData(t, TestDataExpected)
	qa.ResourceFixture{
		MockAccountClientFunc: func(w *mocks.MockAccountClient) {
			e := w.GetMockAccountIpAccessListsAPI().EXPECT()
			e.GetByIpAccessListId(mock.Anything, TestingId).Return(&settings.GetIpAccessListResponse{
				IpAccessList: TestIpAccessList,
			}, nil)
		},
		Resource:  ResourceIPAccessList(),
		Read:      true,
		New:       true,
		AccountID: "100",
		ID:        TestingId,
	}.ApplyAndExpectData(t, TestDataExpected)
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
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockIpAccessListsAPI().EXPECT()
			e.GetByIpAccessListId(mock.Anything, TestingId).Return(nil,
				errors.New("Something unexpected happened"))
		},
		Resource: ResourceIPAccessList(),
		Read:     true,
		ID:       TestingId,
	}.Apply(t)
	assert.Error(t, err)
	qa.AssertErrorStartsWith(t, err, "Something unexpected happened")
	assert.Equal(t, TestingId, d.Id(), "Id should not be empty for error reads")
	d, err = qa.ResourceFixture{
		MockAccountClientFunc: func(w *mocks.MockAccountClient) {
			e := w.GetMockAccountIpAccessListsAPI().EXPECT()
			e.GetByIpAccessListId(mock.Anything, TestingId).Return(nil,
				errors.New("Something unexpected happened"))
		},
		Resource:  ResourceIPAccessList(),
		Read:      true,
		AccountID: "100",
		ID:        TestingId,
	}.Apply(t)
	assert.Error(t, err)
	qa.AssertErrorStartsWith(t, err, "Something unexpected happened")
	assert.Equal(t, TestingId, d.Id(), "Id should not be empty for error reads")
}

func TestIPACLDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockIpAccessListsAPI().EXPECT()
			e.DeleteByIpAccessListId(mock.Anything, TestingId).Return(nil)
		},
		Resource: ResourceIPAccessList(),
		Delete:   true,
		ID:       TestingId,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, TestingId, d.Id())
	d, err = qa.ResourceFixture{
		MockAccountClientFunc: func(w *mocks.MockAccountClient) {
			e := w.GetMockAccountIpAccessListsAPI().EXPECT()
			e.DeleteByIpAccessListId(mock.Anything, TestingId).Return(nil)
		},
		Resource:  ResourceIPAccessList(),
		Delete:    true,
		AccountID: "100",
		ID:        TestingId,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, TestingId, d.Id())
}

func TestIPACLDelete_Error(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockIpAccessListsAPI().EXPECT()
			e.DeleteByIpAccessListId(mock.Anything, TestingId).Return(errors.New("Something went wrong"))
		},
		Resource: ResourceIPAccessList(),
		Delete:   true,
		Removed:  true,
		ID:       TestingId,
	}.ExpectError(t, "Something went wrong")
	qa.ResourceFixture{
		MockAccountClientFunc: func(w *mocks.MockAccountClient) {
			e := w.GetMockAccountIpAccessListsAPI().EXPECT()
			e.DeleteByIpAccessListId(mock.Anything, TestingId).Return(errors.New("Something went wrong"))
		},
		Resource:  ResourceIPAccessList(),
		Delete:    true,
		Removed:   true,
		AccountID: "100",
		ID:        TestingId,
	}.ExpectError(t, "Something went wrong")
}
