package common

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/stretchr/testify/assert"
)

func TestAzureAuth_resourceID(t *testing.T) {
	aa := AzureAuth{}
	assert.Equal(t, "", aa.resourceID())

	aa.ResourceID = "/subscriptions/a/resourceGroups/b/providers/Microsoft.Databricks/workspaces/c"
	assert.Equal(t, "/subscriptions/a/resourceGroups/b/providers/Microsoft.Databricks/workspaces/c", aa.resourceID())
	assert.Equal(t, "a", aa.SubscriptionID)
	assert.Equal(t, "b", aa.ResourceGroup)
	assert.Equal(t, "c", aa.WorkspaceName)

	aa = AzureAuth{}
	aa.SubscriptionID = "a"
	assert.Equal(t, "", aa.resourceID())
	aa.ResourceGroup = "b"
	assert.Equal(t, "", aa.resourceID())
	aa.WorkspaceName = "c"
	assert.Equal(t, "/subscriptions/a/resourceGroups/b/providers/Microsoft.Databricks/workspaces/c", aa.resourceID())
}

func TestAzureAuth_isClientSecretSet(t *testing.T) {
	aa := AzureAuth{}
	assert.False(t, aa.IsClientSecretSet())

	aa.ClientID = "a"
	assert.False(t, aa.IsClientSecretSet())
	aa.ClientSecret = "b"
	assert.False(t, aa.IsClientSecretSet())
	aa.TenantID = "c"
	assert.True(t, aa.IsClientSecretSet())
}

func TestAzureAuth_ensureWorkspaceURL(t *testing.T) {
	aa := AzureAuth{}

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

	aa.ResourceID = "/subscriptions/a/resourceGroups/b/providers/Microsoft.Databricks/workspaces/c"
	aa.azureManagementEndpoint = server.URL

	client := DatabricksClient{InsecureSkipVerify: true}
	client.configureHTTPCLient()
	aa.databricksClient = &client
	client.AzureAuth = aa

	token := &adal.Token{
		AccessToken: "TestToken",
		Resource:    "https://azure.microsoft.com/",
		Type:        "Bearer",
	}
	authorizer := autorest.NewBearerAuthorizer(token)
	err := aa.ensureWorkspaceURL(authorizer)
	assert.NoError(t, err)

	err = aa.ensureWorkspaceURL(authorizer)
	assert.NoError(t, err)
	assert.Equal(t, 1, cnt[0],
		"Calls to Azure Management API must be done only once")
}

func TestAzureAuth_configureWithClientSecret(t *testing.T) {
	aa := AzureAuth{}
	auth, err := aa.configureWithClientSecret()
	assert.Nil(t, auth)
	assert.NoError(t, err)

	aa.ResourceID = "/subscriptions/a/resourceGroups/b/providers/Microsoft.Databricks/workspaces/c"
	auth, err = aa.configureWithClientSecret()
	assert.Nil(t, auth)
	assert.NoError(t, err)

	token := &adal.Token{
		AccessToken: "TestToken",
		Resource:    "https://azure.microsoft.com/",
		Type:        "Bearer",
	}
	aa.authorizer = autorest.NewBearerAuthorizer(token)

	var serverURL string
	dummyPAT := "dapi234567"
	server := httptest.NewUnstartedServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			if req.RequestURI ==
				"/subscriptions/a/resourceGroups/b/providers/Microsoft.Databricks/workspaces/c?api-version=2018-04-01" {
				_, err := rw.Write([]byte(fmt.Sprintf(`{"properties": {"workspaceUrl": "%s"}}`,
					strings.ReplaceAll(serverURL, "https://", ""))))
				assert.NoError(t, err)
				return
			}
			if req.RequestURI == "/api/2.0/token/create" {
				assert.Equal(t, token.AccessToken, req.Header.Get("X-Databricks-Azure-SP-Management-Token"))
				assert.Equal(t, "Bearer "+token.AccessToken, req.Header.Get("Authorization"))
				assert.Equal(t, aa.ResourceID, req.Header.Get("X-Databricks-Azure-Workspace-Resource-Id"))
				_, err := rw.Write([]byte(fmt.Sprintf(`{
					"token_value": "%s", 
					"token_info": {
						"token_id": "qwertyu",
						"creation_time": 1234567,
						"expiry_time": 1234568
					}
				}`, dummyPAT)))
				assert.NoError(t, err)
				return
			}
			if req.RequestURI == "/api/2.0/clusters/list-zones" {
				assert.Equal(t, "Bearer "+dummyPAT, req.Header.Get("Authorization"))
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

	aa.ClientID = "a"
	aa.ClientSecret = "b"
	aa.TenantID = "c"
	aa.azureManagementEndpoint = server.URL
	auth, err = aa.configureWithClientSecret()
	assert.NotNil(t, auth)
	assert.NoError(t, err)

	client := DatabricksClient{InsecureSkipVerify: true}
	client.authVisitor = auth
	err = client.Configure()
	assert.NoError(t, err)
	aa.databricksClient = &client
	client.AzureAuth = aa

	type ZonesInfo struct {
		Zones       []string `json:"zones,omitempty"`
		DefaultZone string   `json:"default_zone,omitempty"`
	}
	var zi ZonesInfo
	err = client.Get("/clusters/list-zones", nil, &zi)
	assert.NotNil(t, zi)
	assert.NoError(t, err)
	assert.Len(t, zi.Zones, 3)
}
