package service

import (
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
func (a TokensAPI) Create(tokenLifetime time.Duration, comment string) (r model.TokenResponse, err error) {
	err = a.client.post("/token/create", model.TokenRequest{
		LifetimeSeconds: int32(tokenLifetime.Seconds()),
		Comment:         comment,
	}, &r)
	return
}

// List will list all the token metadata and not the content of the tokens in the workspace
func (a TokensAPI) List() ([]model.TokenInfo, error) {
	var tokenListResult model.TokenList
	err := a.client.get("/token/list", nil, &tokenListResult)
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
	return a.client.post("/token/delete", map[string]string{
		"token_id": tokenID,
	}, nil)
}
