package tokens

import (
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/settings"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestResourceOboTokenRead(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockTokenManagementAPI().EXPECT().Get(mock.Anything, settings.GetTokenManagementRequest{
				TokenId: "abc",
			}).Return(&settings.GetTokenResponse{
				TokenInfo: &settings.TokenInfo{
					Comment:    "Hello, world!",
					ExpiryTime: time.Now().UnixMilli() + 1000,
				},
			}, nil)
		},
		Resource: ResourceOboToken(),
		Read:     true,
		New:      true,
		ID:       "abc",
	}.ApplyAndExpectData(t, map[string]any{
		"comment": "Hello, world!",
		"id":      "abc",
	})
}

func TestResourceOboTokenRead_NoExpire(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockTokenManagementAPI().EXPECT().Get(mock.Anything, settings.GetTokenManagementRequest{
				TokenId: "abc",
			}).Return(&settings.GetTokenResponse{
				TokenInfo: &settings.TokenInfo{
					Comment:    "Hello, world!",
					ExpiryTime: -1,
				},
			}, nil)
		},
		Resource: ResourceOboToken(),
		Read:     true,
		New:      true,
		ID:       "abc",
	}.ApplyAndExpectData(t, map[string]any{
		"comment": "Hello, world!",
		"id":      "abc",
	})
}

func TestResourceOboTokenRead_Expired(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockTokenManagementAPI().EXPECT().Get(mock.Anything, settings.GetTokenManagementRequest{
				TokenId: "abc",
			}).Return(&settings.GetTokenResponse{
				TokenInfo: &settings.TokenInfo{
					Comment:    "Hello, world!",
					ExpiryTime: time.Now().UnixMilli() - 1000,
				},
			}, nil)
		},
		Resource: ResourceOboToken(),
		Read:     true,
		Removed:  true,
		ID:       "abc",
	}.ApplyAndExpectData(t, map[string]any{
		"id": "",
	})
}

func TestResourceOboTokenRead_Error(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockTokenManagementAPI().EXPECT().Get(mock.Anything, settings.GetTokenManagementRequest{
				TokenId: "abc",
			}).Return(nil, &apierr.APIError{
				StatusCode: 500,
				Message:    "nope",
			})
		},
		Resource: ResourceOboToken(),
		Read:     true,
		New:      true,
		ID:       "abc",
	}.ExpectError(t, "nope")
}

func TestResourceOboTokenRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockTokenManagementAPI().EXPECT().Get(mock.Anything, settings.GetTokenManagementRequest{
				TokenId: "abc",
			}).Return(nil, &apierr.APIError{
				ErrorCode:  "NOT_FOUND",
				StatusCode: 404,
				Message:    "Token does not exist",
			})
		},
		Resource: ResourceOboToken(),
		Read:     true,
		Removed:  true,
		ID:       "abc",
	}.ApplyAndExpectData(t, map[string]any{
		"id": "",
	})
}

func TestResourceOboTokenCreate_Error(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockTokenManagementAPI().EXPECT().CreateOboToken(mock.Anything, mock.Anything).Return(nil, &apierr.APIError{
				StatusCode: 500,
				Message:    "nope",
			})
		},
		Resource: ResourceOboToken(),
		Create:   true,
		New:      true,
	}.ExpectError(t, "nope")
}

func TestResourceOboTokenCreate(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockTokenManagementAPI().EXPECT()
			e.CreateOboToken(mock.Anything, settings.CreateOboTokenRequest{
				ApplicationId:   "abc",
				LifetimeSeconds: 60,
				Comment:         "e",
			}).Return(&settings.CreateOboTokenResponse{
				TokenValue: "s#Cr3t!11",
				TokenInfo: &settings.TokenInfo{
					TokenId:    "bcd",
					ExpiryTime: time.Now().UnixMilli() + 1000,
				},
			}, nil)
			e.Get(mock.Anything, settings.GetTokenManagementRequest{
				TokenId: "bcd",
			}).Return(&settings.GetTokenResponse{
				TokenInfo: &settings.TokenInfo{
					Comment:    "Hello, world!",
					ExpiryTime: time.Now().UnixMilli() + 1000,
				},
			}, nil)
		},
		Resource: ResourceOboToken(),
		Create:   true,
		HCL: `
		application_id = "abc"
		comment = "e"
		lifetime_seconds = 60
		`,
		New: true,
	}.ApplyAndExpectData(t, map[string]any{
		"comment": "Hello, world!",
		"id":      "bcd",
	})
}

func TestResourceOboTokenDelete(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockTokenManagementAPI().EXPECT().Delete(mock.Anything, settings.DeleteTokenManagementRequest{
				TokenId: "abc",
			}).Return(nil)
		},
		Resource: ResourceOboToken(),
		Delete:   true,
		New:      true,
		ID:       "abc",
	}.ApplyNoError(t)
}

func TestResourceOboTokenCreateNoLifetimeOrComment(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockTokenManagementAPI().EXPECT()
			e.CreateOboToken(mock.Anything, settings.CreateOboTokenRequest{
				ApplicationId: "abc",
			}).Return(&settings.CreateOboTokenResponse{
				TokenValue: "s#Cr3t!11",
				TokenInfo: &settings.TokenInfo{
					TokenId:    "bcd",
					ExpiryTime: time.Now().UnixMilli() + 1000,
				},
			}, nil)
			e.Get(mock.Anything, settings.GetTokenManagementRequest{
				TokenId: "bcd",
			}).Return(&settings.GetTokenResponse{
				TokenInfo: &settings.TokenInfo{
					TokenId:    "bcd",
					ExpiryTime: time.Now().UnixMilli() + 1000,
				},
			}, nil)
		},
		Resource: ResourceOboToken(),
		Create:   true,
		HCL: `
		application_id = "abc"
		`,
		New: true,
	}.ApplyAndExpectData(t, map[string]any{
		"id": "bcd",
	})
}
