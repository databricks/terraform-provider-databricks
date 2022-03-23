package common

import (
	"context"
	"io/ioutil"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
)

// Fake credentials for testing only. Private key was generated using:
// $ openssl genrsa -out name_of_private_key.pem 512
const fakeCredentialsJson = `{
	"type": "service_account",
	"project_id": "fake-project-id",
	"private_key_id": "123456789",
	"private_key": "-----BEGIN RSA PRIVATE KEY-----\nMIIBOwIBAAJBAO3U0tUlSLIX06qGqalP+MUScgnmB9scOZ/fZNOykU+QLufdgqDe\nA59CfpytNg/zts8BsRgSTRiLs+6jLhjK6LkCAwEAAQJBALDuGibNROaQyTvcUI2P\n2/8oOMRaZ8++kLP56jV/a5DmwIYt5t5c35/LUWR2GA/7nvQvOJ1XZ6U+uyciOKGg\nJ7ECIQD5k+N8jIMJiobULRLAJgEWQat158sWQ3G23NakdJEWlQIhAPPzjhV+iuJ9\nu4SMOP0BLGgbjWQtna75/cOC916EmLSVAiBrBY7MTti2E7ADdhyPRvy6VYi386Cz\nuFIf7w0f0liRDQIgWD/XOndYjq6lU0HWq8/s3Ix7Da5iyJWu8zdBfXPCOjECIQC1\nusZas9Gcfu4oc3g29c0aQ5IozUTnJhAPjljxj3PmZg==\n-----END RSA PRIVATE KEY-----",
	"client_email": "fake-sa2@example.com",
	"client_id": "987654321",
	"auth_uri": "https://accounts.example.com/o/oauth2/auth",
	"token_uri": "http://127.0.0.1/token",
	"auth_provider_x509_cert_url": "http://127.0.0.1/oauth2/v1/certs",
	"client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/my-sa2%40example.com"
  }`

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

	_, err := client.getGoogleOIDCSource(context.Background())
	require.NoError(t, err)
}

func TestConfigureWithGoogleCredentialsUsingInvalidJson(t *testing.T) {
	defer CleanupEnvironment()()
	client := &DatabricksClient{
		Host:                 "https://accounts.gcp.databricks.com/",
		GoogleServiceAccount: "my-sa2@fake-project-id.iam.gserviceaccount.com",
		GoogleCredentials:    fakeCredentialsJson,
	}
	client.configureHTTPCLient()

	_, err := client.configureWithGoogleCrendentials(context.Background())
	assert.ErrorContains(t, err,
		"could not obtain OIDC token from JSON: oauth2: cannot fetch token")
	// TODO: To get better test coverage, start a local OAuth2 server to support
	// exercising a full token fetching flow.
}

func TestConfigureWithGoogleCredentialsUsingInvalidCredsFile(t *testing.T) {
	defer CleanupEnvironment()()

	// Write fakeCredentialsJson to a temp file to pass to GoogleCredentials.
	credsPath := filepath.Join(t.TempDir(), "creds.json")
	err := ioutil.WriteFile(credsPath, []byte(fakeCredentialsJson), 0600)
	assert.Nil(t, err, "Failed writing input JSON file.")

	client := &DatabricksClient{
		Host:                 "https://accounts.gcp.databricks.com/",
		GoogleServiceAccount: "my-sa2@fake-project-id.iam.gserviceaccount.com",
		GoogleCredentials:    credsPath,
	}
	client.configureHTTPCLient()

	_, err = client.configureWithGoogleCrendentials(context.Background())
	assert.ErrorContains(t, err,
		"could not obtain OIDC token from JSON: oauth2: cannot fetch token")
}

func TestConfigureWithGoogleCredentialsNoCredsFile(t *testing.T) {
	defer CleanupEnvironment()()
	client := &DatabricksClient{
		Host:                 "https://accounts.gcp.databricks.com/",
		GoogleServiceAccount: "my-sa2@fake-project-id.iam.gserviceaccount.com",
		GoogleCredentials:    "/tmp/this/path/does/not/exist/creds.json",
	}
	client.configureHTTPCLient()

	_, err := client.configureWithGoogleCrendentials(context.Background())
	assert.ErrorContains(t, err,
		"could not obtain OIDC token from JSON: invalid character")
}

func TestConfigureWithGoogleCredentialsEmptyHost(t *testing.T) {
	defer CleanupEnvironment()()
	client := &DatabricksClient{
		Host:                 "",
		GoogleServiceAccount: "my-sa2@fake-project-id.iam.gserviceaccount.com",
		GoogleCredentials:    fakeCredentialsJson,
	}
	client.configureHTTPCLient()

	ret, err := client.configureWithGoogleCrendentials(context.Background())
	assert.Nil(t, ret)
	assert.Nil(t, err)
}

func TestConfigureWithGoogleForAccountsAPI(t *testing.T) {
	defer CleanupEnvironment()()
	client := &DatabricksClient{
		Host:                 "https://accounts.gcp.databricks.com/",
		GoogleServiceAccount: "a",
	}
	client.configureHTTPCLient()

	_, err := client.configureWithGoogleForAccountsAPI(context.Background())
	assert.Error(t, err)

	client.googleAuthOptions = []option.ClientOption{option.WithoutAuthentication()}
	a, err := client.configureWithGoogleForAccountsAPI(context.Background())
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

	_, err := client.configureWithGoogleForWorkspace(context.Background())
	assert.Error(t, err)

	client.googleAuthOptions = []option.ClientOption{option.WithoutAuthentication()}
	a, err := client.configureWithGoogleForWorkspace(context.Background())
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
