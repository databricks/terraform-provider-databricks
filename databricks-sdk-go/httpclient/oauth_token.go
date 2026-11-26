package httpclient

import (
	"context"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

const JWTGrantType = "urn:ietf:params:oauth:grant-type:jwt-bearer"

// getOAuthTokenRequest is the request to get an OAuth token. It follows the OAuth 2.0 Rich Authorization Requests specification.
// https://datatracker.ietf.org/doc/html/rfc9396
type getOAuthTokenRequest struct {
	// Defines the method used to get the token.
	GrantType string `url:"grant_type"`
	// An array of authorization details that the token should be scoped to. This needs to be passed in string format.
	AuthorizationDetails string `url:"authorization_details"`
	// The token that will be exchanged for an OAuth token.
	Assertion string `url:"assertion"`
}

// oAuthToken represents an OAuth token as defined by the OAuth 2.0 Authorization Framework.
// https://datatracker.ietf.org/doc/html/rfc6749.
//
// The Go SDK maintains its own implementation of OAuth because Go's oauth2
// library lacks two features that we depend on:
// 1. The ability to use an arbitrary assertion with the JWT grant type.
// 2. The ability to set authorization_details when getting an OAuth token.
type oAuthToken struct {
	// The access token issued by the authorization server. This is the token that will be used to authenticate requests.
	AccessToken string `json:"access_token"`
	// Time in seconds until the token expires.
	ExpiresIn int `json:"expires_in"`
	// The scope of the token. This is a space-separated list of strings that represent the permissions granted by the token.
	Scope string `json:"scope"`
	// The type of token that was issued.
	TokenType string `json:"token_type"`
}

// Returns a new OAuth token using the provided token. The token must be a JWT token.
// The resulting token is scoped to the authorization details provided.
//
// **NOTE:** Experimental: This API may change or be removed in a future release
// without warning.
func (c *ApiClient) GetOAuthToken(ctx context.Context, authDetails string, token *oauth2.Token) (*oauth2.Token, error) {
	path := "/oidc/v1/token"
	data := getOAuthTokenRequest{
		GrantType:            JWTGrantType,
		AuthorizationDetails: authDetails,
		Assertion:            token.AccessToken,
	}
	var response oAuthToken
	opts := []DoOption{
		WithUrlEncodedData(data),
		WithResponseUnmarshal(&response),
	}
	err := c.Do(ctx, http.MethodPost, path, opts...)
	if err != nil {
		return nil, err
	}
	return &oauth2.Token{
		AccessToken: response.AccessToken,
		TokenType:   response.TokenType,
		Expiry:      time.Now().Add(time.Duration(response.ExpiresIn) * time.Second),
	}, nil
}
