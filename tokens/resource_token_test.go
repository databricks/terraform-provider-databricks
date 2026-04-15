package tokens

import (
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/settings"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestResourceTokenRead(t *testing.T) {
	creationTime := time.Now().UnixMilli()
	expiryTime := time.Now().UnixMilli() + 10000
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockTokensAPI().EXPECT().ListAll(mock.Anything).Return([]settings.PublicTokenInfo{
				{
					Comment:      "Hello, world!",
					CreationTime: creationTime,
					ExpiryTime:   expiryTime,
					TokenId:      "abc",
				},
			}, nil)
		},
		Resource: ResourceToken(),
		Read:     true,
		New:      true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
	assert.Equal(t, "Hello, world!", d.Get("comment"))
	assert.Equal(t, creationTime, int64(d.Get("creation_time").(int)))
	assert.Equal(t, expiryTime, int64(d.Get("expiry_time").(int)))
	assert.Equal(t, "", d.Get("token_value"))
}

func TestResourceTokenRead_NoExpire(t *testing.T) {
	creationTime := time.Now().UnixMilli()
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockTokensAPI().EXPECT().ListAll(mock.Anything).Return([]settings.PublicTokenInfo{
				{
					Comment:      "Hello, world!",
					CreationTime: creationTime,
					ExpiryTime:   -1,
					TokenId:      "abc",
				},
			}, nil)
		},
		Resource: ResourceToken(),
		Read:     true,
		New:      true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
	assert.Equal(t, "Hello, world!", d.Get("comment"))
	assert.Equal(t, creationTime, int64(d.Get("creation_time").(int)))
	assert.Equal(t, int64(-1), int64(d.Get("expiry_time").(int)))
	assert.Equal(t, "", d.Get("token_value"))
}

func TestResourceTokenRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockTokensAPI().EXPECT().ListAll(mock.Anything).Return([]settings.PublicTokenInfo{
				{
					Comment:      "Hello, world!",
					CreationTime: 10,
					ExpiryTime:   20,
					TokenId:      "bcd",
				},
			}, nil)
		},
		Resource: ResourceToken(),
		Read:     true,
		Removed:  true,
		New:      true,
		ID:       "abc",
	}.ApplyNoError(t)
}

func TestResourceTokenRead_Expired(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockTokensAPI().EXPECT().ListAll(mock.Anything).Return([]settings.PublicTokenInfo{
				{
					Comment:      "Hello, world!",
					CreationTime: 10,
					ExpiryTime:   20,
					TokenId:      "abc",
				},
			}, nil)
		},
		Resource: ResourceToken(),
		Read:     true,
		Removed:  true,
		New:      true,
		ID:       "abc",
	}.ApplyNoError(t)
}

func TestResourceTokenRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockTokensAPI().EXPECT().ListAll(mock.Anything).Return(nil, &apierr.APIError{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			})
		},
		Resource: ResourceToken(),
		Read:     true,
		ID:       "abc",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id(), "Id should not be empty for error reads")
}

func TestResourceTokenCreate(t *testing.T) {
	creationTime := time.Now().UnixMilli()
	expiryTime := time.Now().UnixMilli() + 10000
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockTokensAPI().EXPECT()
			e.Create(mock.Anything, settings.CreateTokenRequest{
				LifetimeSeconds: 300,
				Comment:         "Hello world!",
				Scopes:          []string{"scope1", "scope2"},
			}).Return(&settings.CreateTokenResponse{
				TokenValue: "dapi...",
				TokenInfo: &settings.PublicTokenInfo{
					TokenId: "abc",
				},
			}, nil)
			e.ListAll(mock.Anything).Return([]settings.PublicTokenInfo{
				{
					Comment:      "Hello, world!",
					CreationTime: creationTime,
					ExpiryTime:   expiryTime,
					TokenId:      "abc",
				},
			}, nil)
		},
		Resource: ResourceToken(),
		State: map[string]any{
			"comment":          "Hello world!",
			"lifetime_seconds": 300,
			"scopes":           []any{"scope1", "scope2"},
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
	assert.Equal(t, "dapi...", d.Get("token_value"))
	scopes := d.Get("scopes").([]any)
	assert.Equal(t, []any{"scope1", "scope2"}, scopes)
}

func TestResourceTokenCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockTokensAPI().EXPECT().Create(mock.Anything, mock.Anything).Return(nil, &apierr.APIError{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			})
		},
		Resource: ResourceToken(),
		State: map[string]any{
			"comment":          "Hello world!",
			"lifetime_seconds": 300,
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceTokenCreate_NoExpiration(t *testing.T) {
	creationTime := time.Now().UnixMilli()
	expiryTime := time.Now().UnixMilli() + 10000
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockTokensAPI().EXPECT()
			e.Create(mock.Anything, settings.CreateTokenRequest{}).Return(&settings.CreateTokenResponse{
				TokenValue: "dapi...",
				TokenInfo: &settings.PublicTokenInfo{
					TokenId:    "abc",
					ExpiryTime: -1,
				},
			}, nil)
			e.ListAll(mock.Anything).Return([]settings.PublicTokenInfo{
				{
					CreationTime: creationTime,
					ExpiryTime:   expiryTime,
					TokenId:      "abc",
				},
			}, nil)
		},
		Resource: ResourceToken(),
		State:    map[string]any{},
		Create:   true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
	assert.Equal(t, expiryTime, int64(d.Get("expiry_time").(int)))
	assert.Equal(t, "dapi...", d.Get("token_value"))
}

func TestResourceTokenDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockTokensAPI().EXPECT().Delete(mock.Anything, settings.RevokeTokenRequest{
				TokenId: "abc",
			}).Return(nil)
		},
		Resource: ResourceToken(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceTokenDelete_NotFoundNoError(t *testing.T) {
	_, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockTokensAPI().EXPECT().Delete(mock.Anything, settings.RevokeTokenRequest{
				TokenId: "abc",
			}).Return(&apierr.APIError{
				ErrorCode:  "NOT_FOUND",
				StatusCode: 404,
				Message:    "RESOURCE_DOES_NOT_EXIST secret not found",
			})
		},
		Resource: ResourceToken(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err)
}

func TestResourceTokenDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockTokensAPI().EXPECT().Delete(mock.Anything, settings.RevokeTokenRequest{
				TokenId: "abc",
			}).Return(&apierr.APIError{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			})
		},
		Resource: ResourceToken(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id())
}
