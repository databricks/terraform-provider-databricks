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
	"github.com/golang-jwt/jwt/v4"
	"github.com/hashicorp/go-retryablehttp"

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
	auth, err := aa.getClientSecretAuthorizer(armDatabricksResourceID)
	require.Nil(t, auth)
	require.EqualError(t, err, "parameter 'clientID' cannot be empty")

	aa.AzureTenantID = "a"
	aa.AzureClientID = "b"
	aa.AzureClientSecret = "c"
	auth, err = aa.getClientSecretAuthorizer(armDatabricksResourceID)
	require.NotNil(t, auth)
	require.NoError(t, err)

	auth, err = aa.getClientSecretAuthorizer(env.ServiceManagementEndpoint)
	require.NotNil(t, auth)
	require.NoError(t, err)
}

func TestEnsureWorkspaceURL_CornerCases(t *testing.T) {
	aa := DatabricksClient{}
	env, err := aa.getAzureEnvironment()
	require.NoError(t, err)
	aa.AzureEnvironment = &env

	err = aa.ensureWorkspaceURL(context.Background(), nil)
	assert.EqualError(t, err, "somehow resource id is not set")
}

func TestDatabricksClient_ensureWorkspaceURL(t *testing.T) {
	aa := DatabricksClient{InsecureSkipVerify: true}
	aa.configureHTTPCLient()
	env, err := aa.getAzureEnvironment()
	require.NoError(t, err)
	aa.AzureEnvironment = &env

	cnt := []int{0}
	var serverURL string
	server := httptest.NewUnstartedServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			if req.RequestURI ==
				"/subscriptions/a/resourceGroups/b/providers/Microsoft.Databricks/workspaces/c?api-version=2018-04-01" {
				_, err := rw.Write([]byte(fmt.Sprintf(`{"properties": {"workspaceUrl": "%s"}}`,
					strings.ReplaceAll(serverURL, "https://", ""))))
				assert.NoError(t, err)
				cnt[0]++
				return
			}
			assert.Fail(t, fmt.Sprintf("Received unexpected call: %s %s",
				req.Method, req.RequestURI))
		}))
	server.StartTLS()
	serverURL = server.URL
	defer server.Close()

	aa.AzureResourceID = "/subscriptions/a/resourceGroups/b/providers/Microsoft.Databricks/workspaces/c"
	// resource management endpoints end with a trailing slash in url
	aa.AzureEnvironment = &azure.Environment{
		ResourceManagerEndpoint: fmt.Sprintf("%s/", server.URL),
	}

	token := &adal.Token{
		AccessToken: "TestToken",
		Resource:    "https://azure.microsoft.com/",
		Type:        "Bearer",
	}
	authorizer := autorest.NewBearerAuthorizer(token)
	err = aa.ensureWorkspaceURL(context.Background(), authorizer)
	assert.NoError(t, err)

	err = aa.ensureWorkspaceURL(context.Background(), authorizer)
	assert.NoError(t, err)
	assert.Equal(t, 1, cnt[0],
		"Calls to Azure Management API must be done only once")
}

func TestDatabricksClient_configureWithClientSecretAAD(t *testing.T) {
	client := DatabricksClient{InsecureSkipVerify: true}
	client.AzureResourceID = "/subscriptions/a/resourceGroups/b/providers/Microsoft.Databricks/workspaces/c"

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
				assert.Equal(t, client.AzureResourceID, req.Header.Get("X-Databricks-Azure-Workspace-Resource-Id"))
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

func newTestJwt(t *testing.T, claims jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	result, err := token.SignedString([]byte("_"))
	assert.NoError(t, err)
	return result
}

func TestGetJWTProperty_AzureCli(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")
	os.Setenv("TF_AAD_TOKEN", newTestJwt(t, jwt.MapClaims{
		"tid": "some-tenant",
	}))

	srv, client := setupJwtTestClient()
	assert.NotNil(t, srv)
	defer srv.Close()
	assert.NotNil(t, client)

	tid, err := client.GetAzureJwtProperty("tid")
	require.NoError(t, err)
	assert.Equal(t, "some-tenant", tid.(string))
}

func TestGetJWTProperty_Authenticate_Fail(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")
	os.Setenv("FAIL", "yes")

	client := &DatabricksClient{
		Host: "https://adb-1232.azuredatabricks.net",
	}
	_, err := client.GetAzureJwtProperty("tid")
	require.EqualError(t, err, "cannot configure azure-cli auth: "+
		"Invoking Azure CLI failed with the following error: "+
		"This is just a failing script.\n. "+
		"Please check https://registry.terraform.io/providers/"+
		"databrickslabs/databricks/latest/docs#authentication for details")
}

func TestGetJWTProperty_makeGetRequest_Fail(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")
	os.Setenv("TF_AAD_TOKEN", newTestJwt(t, jwt.MapClaims{
		"tid": "some-tenant",
	}))

	srv, client := setupJwtTestClient()
	assert.NotNil(t, srv)
	defer srv.Close()
	assert.NotNil(t, client)
	client.Host = "%🙀.azuredatabricks.net"

	_, err := client.GetAzureJwtProperty("tid")
	require.EqualError(t, err, `parse "%🙀.azuredatabricks.net": invalid URL escape "%\xf0\x9f"`)
}

func TestGetJWTProperty_authVisitor_Fail(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")

	client := &DatabricksClient{
		Host: "https://adb-1232.azuredatabricks.net",
		authVisitor: func(r *http.Request) error {
			return fmt.Errorf("fails for the test")
		},
	}
	_, err := client.GetAzureJwtProperty("tid")
	require.EqualError(t, err, "fails for the test")
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
	os.Setenv("TF_AAD_TOKEN", newTestJwt(t, jwt.MapClaims{
		"something": "else",
	}))

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

func TestNoMsiAvailable(t *testing.T) {
	msiAvailabilityChecker = func(ctx context.Context, s adal.Sender) bool {
		// without this shim test will fail on GitHub, as it runs on Azure
		return false
	}
	_, err := (&DatabricksClient{
		AzureResourceID: "/a/b/c",
		AzureUseMSI:     true,
		httpClient: &retryablehttp.Client{
			HTTPClient: http.DefaultClient,
			RetryMax:   0,
			CheckRetry: func(ctx context.Context, resp *http.Response, err error) (bool, error) {
				return false, err
			},
		},
	}).configureWithAzureManagedIdentity(context.Background())
	assert.EqualError(t, err, "managed identity is not available")
}

func TestMsiWorks(t *testing.T) {
	msiAvailabilityChecker = func(ctx context.Context, s adal.Sender) bool {
		return true
	}
	auth, err := (&DatabricksClient{
		AzureResourceID: "/a/b/c",
		AzureUseMSI:     true,
		Host:            "abc.azuredatabricks.net",
		AzureEnvironment: &azure.Environment{
			ResourceManagerEndpoint:   fmt.Sprintf("%s/", "http://localhost"),
			ServiceManagementEndpoint: "sm",
		},
		httpClient: &retryablehttp.Client{
			HTTPClient: http.DefaultClient,
			RetryMax:   0,
		},
	}).configureWithAzureManagedIdentity(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, auth)
}

func TestSimpleAADRequestVisitor_FailManagement(t *testing.T) {
	_, err := (&DatabricksClient{
		AzureEnvironment: &azure.Environment{
			ServiceManagementEndpoint: "x",
		},
	}).simpleAADRequestVisitor(context.Background(),
		func(resource string) (autorest.Authorizer, error) {
			return nil, fmt.Errorf("🤨")
		})
	assert.EqualError(t, err, "cannot authorize management: 🤨")
}

func TestSimpleAADRequestVisitor_FailWorkspaceUrl(t *testing.T) {
	_, err := (&DatabricksClient{
		AzureEnvironment: &azure.Environment{
			ServiceManagementEndpoint: "x",
		},
	}).simpleAADRequestVisitor(context.Background(),
		func(resource string) (autorest.Authorizer, error) {
			return autorest.NullAuthorizer{}, nil
		})
	assert.EqualError(t, err, "cannot get workspace: somehow resource id is not set")
}

func TestSimpleAADRequestVisitor_FailPlatformAuth(t *testing.T) {
	_, err := (&DatabricksClient{
		Host: "abc.azuredatabricks.net",
		AzureEnvironment: &azure.Environment{
			ServiceManagementEndpoint: "x",
		},
	}).simpleAADRequestVisitor(context.Background(),
		func(resource string) (autorest.Authorizer, error) {
			if resource == armDatabricksResourceID {
				return nil, fmt.Errorf("🤨")
			}
			return autorest.NullAuthorizer{}, nil
		})
	assert.EqualError(t, err, "cannot authorize databricks: 🤨")
}
