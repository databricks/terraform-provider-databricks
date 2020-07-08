package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// TokensAPI exposes the Secrets API
type TokensAPI struct {
	client *DatabricksClient
}

// Create creates a api token given a expiration duration and a comment
func (a TokensAPI) Create(tokenLifetime time.Duration, comment string) (model.TokenResponse, error) {
	var tokenData model.TokenResponse

	tokenCreateRequest := model.TokenRequest{
		LifetimeSeconds: int32(tokenLifetime.Seconds()),
		Comment:         comment,
	}

	tokenCreateResponse, err := a.client.performQuery(http.MethodPost, "/token/create", "2.0", nil, tokenCreateRequest)
	if err != nil {
		return tokenData, err
	}

	err = json.Unmarshal(tokenCreateResponse, &tokenData)
	return tokenData, err
}

// List will list all the token metadata and not the content of the tokens in the workspace
func (a TokensAPI) List() ([]model.TokenInfo, error) {
	var tokenListResult struct {
		TokenInfos []model.TokenInfo `json:"token_infos,omitempty"`
	}
	tokenListResponse, err := a.client.performQuery(http.MethodGet, "/token/list", "2.0", nil, tokenListResult)
	if err != nil {
		return tokenListResult.TokenInfos, err
	}
	err = json.Unmarshal(tokenListResponse, &tokenListResult)
	return tokenListResult.TokenInfos, err
}

// Read will return the token metadata and not the content of the token
func (a TokensAPI) Read(tokenID string) (model.TokenInfo, error) {
	var tokenInfo model.TokenInfo
	tokenList, err := a.List()
	if err != nil {
		return tokenInfo, err
	}
	for _, tokenInfoRecord := range tokenList {
		if tokenInfoRecord.TokenID == tokenID {
			return tokenInfoRecord, nil
		}
	}
	return tokenInfo, APIError{
		ErrorCode:  "NOT_FOUND",
		Message:    fmt.Sprintf("Unable to locate token: %s", tokenID),
		Resource:   "/api/2.0/token/list",
		StatusCode: http.StatusNotFound,
	}
}

// Delete will delete the token given a token id
func (a TokensAPI) Delete(tokenID string) error {
	tokenDeleteRequest := struct {
		TokenID string `json:"token_id,omitempty"`
	}{
		tokenID,
	}
	_, err := a.client.performQuery(http.MethodPost, "/token/delete", "2.0", nil, tokenDeleteRequest)
	return err
}
