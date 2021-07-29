package common

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
)

func TestGoogleOIDC(t *testing.T) {
	defer CleanupEnvironment()()
	client := &DatabricksClient{
		Host:                 "http://localhost",
		GoogleServiceAccount: "a",
		googleAuthOptions: []option.ClientOption{
			option.WithoutAuthentication(),
		},
	}
	client.configureHTTPCLient()

	_, err := client.getGoogleOIDCSource()
	require.NoError(t, err)
}

func TestConfigureWithGoogleForAccountsAPI(t *testing.T) {
	defer CleanupEnvironment()()
	client := &DatabricksClient{
		Host:                 "https://accounts.gcp.databricks.com/",
		GoogleServiceAccount: "a",
	}
	client.configureHTTPCLient()

	_, err := client.configureWithGoogleForAccountsAPI()
	assert.Error(t, err)

	client.googleAuthOptions = []option.ClientOption{option.WithoutAuthentication()}
	a, err := client.configureWithGoogleForAccountsAPI()
	require.NoError(t, err)
	assert.NotNil(t, a)
}

func TestConfigureWithGoogleForWorkspace(t *testing.T) {
	defer CleanupEnvironment()()
	client := &DatabricksClient{
		Host:                 "https://123.4.gcp.databricks.com/",
		GoogleServiceAccount: "a",
	}
	client.configureHTTPCLient()

	_, err := client.configureWithGoogleForWorkspace()
	assert.Error(t, err)

	client.googleAuthOptions = []option.ClientOption{option.WithoutAuthentication()}
	a, err := client.configureWithGoogleForWorkspace()
	require.NoError(t, err)
	assert.NotNil(t, a)
}

func TestNewOidcAuthorizerForAccountsAPI(t *testing.T) {
	token := oauth2.Token{
		AccessToken: "abc",
		TokenType:   "Bearer",
	}
	auth := newOidcAuthorizerForAccountsAPI(
		oauth2.StaticTokenSource(&token),
		oauth2.StaticTokenSource(&token))
	request := httptest.NewRequest("GET", "http://localhost", nil)
	err := auth(request)
	require.NoError(t, err)

	assert.Equal(t, "Bearer abc", request.Header.Get("Authorization"))
	assert.Equal(t, "abc", request.Header.Get("X-Databricks-GCP-SA-Access-Token"))
}

func TestNewOidcAuthorizerForWorkspace(t *testing.T) {
	token := oauth2.Token{
		AccessToken: "abc",
		TokenType:   "Bearer",
	}
	auth := newOidcAuthorizerForWorkspace(
		oauth2.StaticTokenSource(&token))
	request := httptest.NewRequest("GET", "http://localhost", nil)
	err := auth(request)
	require.NoError(t, err)

	assert.Equal(t, "Bearer abc", request.Header.Get("Authorization"))
	assert.Equal(t, "", request.Header.Get("X-Databricks-GCP-SA-Access-Token"))
}
