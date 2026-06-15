package tokens

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"
)

// TokenRequest asks for a token
type TokenRequest struct {
	LifetimeSeconds int32    `json:"lifetime_seconds,omitempty"`
	Comment         string   `json:"comment,omitempty"`
	Scopes          []string `json:"scopes,omitempty"`
}

// TokenResponse is a struct that contains information about token that is created from the create tokens api
type TokenResponse struct {
	TokenValue string     `json:"token_value,omitempty"`
	TokenInfo  *TokenInfo `json:"token_info,omitempty"`
}

// TokenInfo is a struct that contains metadata about a given token
type TokenInfo struct {
	TokenID      string `json:"token_id,omitempty"`
	CreationTime int64  `json:"creation_time,omitempty"`
	ExpiryTime   int64  `json:"expiry_time,omitempty"`
	Comment      string `json:"comment,omitempty"`
}

// TokenList ...
type TokenList struct {
	TokenInfos []TokenInfo `json:"token_infos,omitempty"`
}

// NewTokensAPI creates TokensAPI instance from provider meta
func NewTokensAPI(ctx context.Context, m any) TokensAPI {
	return TokensAPI{m.(*common.DatabricksClient), ctx}
}

// TokensAPI exposes the Secrets API
type TokensAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Create creates a api token given a token request
func (a TokensAPI) Create(request TokenRequest) (r TokenResponse, err error) {
	err = a.client.Post(a.context, "/token/create", request, &r)
	return
}

// List will list all the token metadata and not the content of the tokens in the workspace
func (a TokensAPI) List() ([]TokenInfo, error) {
	var tokenListResult TokenList
	err := a.client.Get(a.context, "/token/list", nil, &tokenListResult)
	return tokenListResult.TokenInfos, err
}

// Read will return the token metadata and not the content of the token
func (a TokensAPI) Read(tokenID string) (TokenInfo, error) {
	var tokenInfo TokenInfo
	tokenList, err := a.List()
	if err != nil {
		return tokenInfo, err
	}
	for _, tokenInfoRecord := range tokenList {
		if tokenInfoRecord.TokenID == tokenID {
			return tokenInfoRecord, nil
		}
	}
	return tokenInfo, &apierr.APIError{
		ErrorCode:  "NOT_FOUND",
		StatusCode: 404,
		Message:    fmt.Sprintf("Unable to locate token: %s", tokenID),
	}
}

// Delete will delete the token given a token id
func (a TokensAPI) Delete(tokenID string) error {
	err := a.client.Post(a.context, "/token/delete", map[string]string{
		"token_id": tokenID,
	}, nil)
	return common.IgnoreNotFoundError(err) // ignore not found error on delete, as it is idempotent
}
