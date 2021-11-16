package common

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddSpManagementTokenVisitor(t *testing.T) {
	aa := DatabricksClient{}
	r := httptest.NewRequest("GET", "/a/b/c", http.NoBody)
	err := aa.addSpManagementTokenVisitor(r, &autorest.BearerAuthorizer{})
	assert.EqualError(t, err, "token provider is nil")
}

func TestAddSpManagementTokenVisitor_Refreshed(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")

	aa := DatabricksClient{}
	r := httptest.NewRequest("GET", "/a/b/c", http.NoBody)
	rct := refreshableCliToken{
		lock:           &sync.RWMutex{},
		resource:       "x",
		refreshMinutes: 6,
	}
	err := aa.addSpManagementTokenVisitor(r, autorest.NewBearerAuthorizer(&rct))
	require.NoError(t, err)
	assert.Equal(t, "...", r.Header[http.CanonicalHeaderKey("X-Databricks-Azure-SP-Management-Token")][0])
}

func TestAddSpManagementTokenVisitor_RefreshedError(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")
	os.Setenv("FAIL", "yes")

	aa := DatabricksClient{}
	r := httptest.NewRequest("GET", "/a/b/c", http.NoBody)
	rct := refreshableCliToken{
		lock:           &sync.RWMutex{},
		resource:       "x",
		refreshMinutes: 6,
	}
	err := aa.addSpManagementTokenVisitor(r, autorest.NewBearerAuthorizer(&rct))
	require.EqualError(t, err, "cannot refresh AAD token: cannot get access token: This is just a failing script.\n")

	err = aa.addSpManagementTokenVisitor(r, autorest.NewBasicAuthorizer("a", "b"))
	require.Error(t, err)
}

func TestGetClientSecretAuthorizer(t *testing.T) {
	aa := DatabricksClient{}
	env, err := aa.getAzureEnvironment()
	require.NoError(t, err)
	aa.AzureEnvironment = &env
	auth, err := aa.getClientSecretAuthorizer(AzureDatabricksResourceID)
	require.Nil(t, auth)
	require.EqualError(t, err, "parameter 'clientID' cannot be empty")

	aa.AzureTenantID = "a"
	aa.AzureClientID = "b"
	aa.AzureClientSecret = "c"
	auth, err = aa.getClientSecretAuthorizer(AzureDatabricksResourceID)
	require.NotNil(t, auth)
	require.NoError(t, err)
}

func TestDatabricksClient_configureWithClientSecretAAD(t *testing.T) {
	client := DatabricksClient{InsecureSkipVerify: true}
	token := &adal.Token{
		AccessToken: "TestToken",
		Resource:    "https://azure.microsoft.com/",
		Type:        "Bearer",
	}
	client.azureAuthorizer = autorest.NewBearerAuthorizer(token)

	var serverURL string
	server := httptest.NewUnstartedServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			if req.RequestURI ==
				"/subscriptions/a/resourceGroups/b/providers/Microsoft.Databricks/workspaces/c?api-version=2018-04-01" {
				_, err := rw.Write([]byte(fmt.Sprintf(`{"properties": {"workspaceUrl": "%s"}}`,
					strings.ReplaceAll(serverURL, "https://", ""))))
				assert.NoError(t, err)
				return
			}
			if req.RequestURI == "/api/2.0/clusters/list-zones" {
				assert.Equal(t, token.AccessToken, req.Header.Get("X-Databricks-Azure-SP-Management-Token"))
				assert.Equal(t, "Bearer "+token.AccessToken, req.Header.Get("Authorization"))
				_, err := rw.Write([]byte(`{"zones": ["a", "b", "c"]}`))
				assert.NoError(t, err)
				return
			}
			assert.Fail(t, fmt.Sprintf("Received unexpected call: %s %s",
				req.Method, req.RequestURI))
		}))
	server.StartTLS()
	serverURL = server.URL
	defer server.Close()

	client.AzureClientID = "a"
	client.AzureClientSecret = "b"
	client.AzureTenantID = "c"
	// resource management endpoints end with a trailing slash in url
	client.AzureEnvironment = &azure.Environment{
		ResourceManagerEndpoint: fmt.Sprintf("%s/", server.URL),
	}
	client.configureHTTPCLient()
	ctx := context.Background()
	auth, err := client.configureWithAzureClientSecret(ctx)
	assert.NoError(t, err)

	client.authVisitor = auth
	client.configureHTTPCLient()

	type ZonesInfo struct {
		Zones       []string `json:"zones,omitempty"`
		DefaultZone string   `json:"default_zone,omitempty"`
	}
	var zi ZonesInfo
	err = client.Get(context.Background(), "/clusters/list-zones", nil, &zi)
	assert.NotNil(t, zi)
	assert.NoError(t, err)
	assert.Len(t, zi.Zones, 3)
}

func TestAzureEnvironment_WithAzureManagementEndpoint(t *testing.T) {
	fakeEndpoint := "http://google.com"
	aa := DatabricksClient{
		AzureEnvironment: &azure.Environment{
			ResourceManagerEndpoint: fakeEndpoint,
		},
	}
	env, err := aa.getAzureEnvironment()
	assert.Nil(t, err)
	// This value should be populated with azureManagementEndpoint for testing
	assert.Equal(t, env.ResourceManagerEndpoint, fakeEndpoint)
	// The rest should be nill
	assert.Equal(t, env.ActiveDirectoryEndpoint, "")

	// Making the azureManagementEndpoint empty should yield PublicCloud
	aa.AzureEnvironment = nil
	env, err = aa.getAzureEnvironment()
	assert.Nil(t, err)
	assert.Equal(t, azure.PublicCloud, env)
}

func TestAzureEnvironment(t *testing.T) {
	aa := DatabricksClient{}
	env, err := aa.getAzureEnvironment()

	assert.Nil(t, err)
	assert.Equal(t, azure.PublicCloud, env)

	aa.AzurermEnvironment = "public"
	env, err = aa.getAzureEnvironment()
	assert.Nil(t, err)
	assert.Equal(t, azure.PublicCloud, env)

	aa.AzurermEnvironment = "china"
	env, err = aa.getAzureEnvironment()
	assert.Nil(t, err)
	assert.Equal(t, azure.ChinaCloud, env)

	aa.AzurermEnvironment = "german"
	env, err = aa.getAzureEnvironment()
	assert.Nil(t, err)
	assert.Equal(t, azure.GermanCloud, env)

	aa.AzurermEnvironment = "usgovernment"
	env, err = aa.getAzureEnvironment()
	assert.Nil(t, err)
	assert.Equal(t, azure.USGovernmentCloud, env)

	aa.AzurermEnvironment = "xyzdummy"
	_, err = aa.getAzureEnvironment()
	assert.NotNil(t, err)
}

func TestMaybeExtendError(t *testing.T) {
	err := fmt.Errorf("Some test")
	err2 := maybeExtendAuthzError(err)

	assert.EqualError(t, err2, err.Error())

	msg := "Azure authorization error. Does your SPN"

	err = fmt.Errorf("something does not have authorization to perform action abc")
	err2 = maybeExtendAuthzError(err)
	assert.True(t, strings.HasPrefix(err2.Error(), msg), err2.Error())

	err = APIError{StatusCode: 403}
	err2 = maybeExtendAuthzError(err)
	assert.True(t, strings.HasPrefix(err2.Error(), msg), err2.Error())
}

func TestGetJWTProperty_AzureCLI_SP(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")

	aa := DatabricksClient{
		AzureClientID:     "a",
		AzureClientSecret: "b",
		AzureTenantID:     "c",
		Host:              "https://adb-1232.azuredatabricks.net",
	}
	tid, err := aa.GetAzureJwtProperty("tid")
	assert.NoError(t, err)
	assert.Equal(t, "c", tid)
}

func TestGetJWTProperty_NonAzure(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")

	aa := DatabricksClient{
		Host:  "https://abc.cloud.databricks.com",
		Token: "abc",
	}
	_, err := aa.GetAzureJwtProperty("tid")
	require.EqualError(t, err, "can't get Azure JWT token in non-Azure environment")
}

func TestGetJWTProperty_AzureCli_Error(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")

	// token without expiry in this case
	client, server := singleRequestServer(t, "POST", "/api/2.0/token/create", `{
		"token_value": "abc"
	}`)
	defer server.Close()

	client.Host = "https://adb-1232.azuredatabricks.net"

	_, err := client.GetAzureJwtProperty("tid")
	require.EqualError(t, err, "unexpected end of JSON input")
}

func setupJwtTestClient() (*httptest.Server, *DatabricksClient) {
	client := DatabricksClient{InsecureSkipVerify: true,
		Host: "https://adb-1232.azuredatabricks.net",
	}
	ctx := context.Background()

	server := httptest.NewUnstartedServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
		}))
	server.StartTLS()

	// resource management endpoints end with a trailing slash in url
	client.AzureEnvironment = &azure.Environment{
		ResourceManagerEndpoint: fmt.Sprintf("%s/", server.URL),
	}
	auth, err := client.configureWithAzureCLI(ctx)
	if err != nil || auth == nil {
		return nil, nil
	}

	client.authVisitor = auth
	client.configureHTTPCLient()

	return server, &client
}

func TestGetJWTProperty_AzureCli(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")
	os.Setenv("TF_AAD_TOKEN", "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE2MzU2OTU4MzksImV4cCI6MTY2NzIzMTgzOSwiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoianJvY2tldEBleGFtcGxlLmNvbSIsInRpZCI6ImFiYyJ9._G1DrR4DspidISpsra8UnecV_FV4zMlJDtSNzaS0UxI")

	srv, client := setupJwtTestClient()
	assert.NotNil(t, srv)
	defer srv.Close()
	assert.NotNil(t, client)

	tid, err := client.GetAzureJwtProperty("tid")
	require.NoError(t, err)
	assert.Equal(t, "abc", tid.(string))
}

func TestGetJWTProperty_AzureCli_Error_DB_PAT(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")
	os.Setenv("TF_AAD_TOKEN", "dapi123")

	srv, client := setupJwtTestClient()
	assert.NotNil(t, srv)
	defer srv.Close()
	assert.NotNil(t, client)

	_, err := client.GetAzureJwtProperty("tid")
	require.EqualError(t, err, "can't use Databricks PAT")
}

func TestGetJWTProperty_AzureCli_Error_No_TenantID(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")
	os.Setenv("TF_AAD_TOKEN", "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE2MzU3OTMwOTksImV4cCI6MTY2NzMyOTA5OSwiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoianJvY2tldEBleGFtcGxlLmNvbSJ9.bWKxWSNburEQPGx71pxO1xTa8cuGue5PeXn407voUMI")

	srv, client := setupJwtTestClient()
	assert.NotNil(t, srv)
	defer srv.Close()
	assert.NotNil(t, client)

	_, err := client.GetAzureJwtProperty("tid")
	require.EqualError(t, err, "can't find field 'tid' in parsed JWT")
}

func TestGetJWTProperty_AzureCli_Error_EmptyToken(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")
	os.Setenv("TF_AAD_TOKEN", "   ")

	srv, client := setupJwtTestClient()
	assert.NotNil(t, srv)
	defer srv.Close()
	assert.NotNil(t, client)

	_, err := client.GetAzureJwtProperty("tid")
	require.EqualError(t, err, "can't obtain Azure JWT token")
}
