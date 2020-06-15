package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// TokensAPI exposes the Secrets API
type TokensAPI struct {
	Client *DBApiClient
}

// Create creates a api token given a expiration duration and a comment
func (a TokensAPI) Create(lifeTimeSeconds int32, comment string) (model.TokenResponse, error) {
	var tokenData model.TokenResponse

	tokenCreateRequest := struct {
		LifetimeSeconds int32  `json:"lifetime_seconds,omitempty"`
		Comment         string `json:"comment,omitempty"`
	}{
		lifeTimeSeconds,
		comment,
	}

	tokenCreateResponse, err := a.Client.performQuery(http.MethodPost, "/token/create", "2.0", nil, tokenCreateRequest, nil)
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
	tokenListResponse, err := a.Client.performQuery(http.MethodGet, "/token/list", "2.0", nil, tokenListResult, nil)
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
	return tokenInfo, fmt.Errorf("Unable to locate token: %s", tokenID)
}

// Delete will delete the token given a token id
func (a TokensAPI) Delete(tokenID string) error {
	tokenDeleteRequest := struct {
		TokenID string `json:"token_id,omitempty"`
	}{
		tokenID,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/token/delete", "2.0", nil, tokenDeleteRequest, nil)
	return err
}
