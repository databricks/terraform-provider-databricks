package service

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/stretchr/testify/assert"
)

func TestAzureAuth_resourceID(t *testing.T) {
	aa := AzureAuth{}
	assert.Equal(t, "", aa.resourceID())

	aa.ResourceID = "/foo/bar"
	assert.Equal(t, "/foo/bar", aa.resourceID())
	aa.ResourceID = ""

	aa.SubscriptionID = "a"
	assert.Equal(t, "", aa.resourceID())
	aa.ResourceGroup = "b"
	assert.Equal(t, "", aa.resourceID())
	aa.WorkspaceName = "c"
	assert.Equal(t, "/subscriptions/a/resourceGroups/b/providers/Microsoft.Databricks/workspaces/c", aa.resourceID())
}

func TestAzureAuth_isClientSecretSet(t *testing.T) {
	aa := AzureAuth{}
	assert.False(t, aa.isClientSecretSet())

	aa.ClientID = "a"
	assert.False(t, aa.isClientSecretSet())
	aa.ClientSecret = "b"
	assert.False(t, aa.isClientSecretSet())
	aa.TenantID = "c"
	assert.True(t, aa.isClientSecretSet())
}

func TestAzureAuth_configureWithClientSecret(t *testing.T) {
	aa := AzureAuth{}
	auth, err := aa.configureWithClientSecret2()
	assert.Nil(t, auth)
	assert.NoError(t, err)

	aa.ResourceID = "/a/b/c"
	auth, err = aa.configureWithClientSecret2()
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
			if req.RequestURI == "/a/b/c?api-version=2018-04-01" {
				rw.Write([]byte(fmt.Sprintf(`{"properties": {"workspaceUrl": "%s"}}`,
					strings.ReplaceAll(serverURL, "https://", ""))))
				return
			}
			if req.RequestURI == "/api/2.0/token/create" {
				assert.Equal(t, token.AccessToken, req.Header.Get("X-Databricks-Azure-SP-Management-Token"))
				assert.Equal(t, "Bearer "+token.AccessToken, req.Header.Get("Authorization"))
				assert.Equal(t, aa.ResourceID, req.Header.Get("X-Databricks-Azure-Workspace-Resource-Id"))
				rw.Write([]byte(fmt.Sprintf(`{
					"token_value": "%s", 
					"token_info": {
						"token_id": "qwertyu",
						"creation_time": 1234567,
						"expiry_time": 1234568
					}
				}`, dummyPAT)))
				return
			}
			if req.RequestURI == "/api/2.0/clusters/list-zones?" {
				assert.Equal(t, "Bearer "+dummyPAT, req.Header.Get("Authorization"))
				rw.Write([]byte(`{"zones": ["a", "b", "c"]}`))
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
	auth, err = aa.configureWithClientSecret2()
	assert.NotNil(t, auth)
	assert.NoError(t, err)

	client := DatabricksClient{InsecureSkipVerify: true}
	client.configureHTTPCLient()
	aa.databricksClient = &client
	client.AzureAuth = aa
	err = client.findAndApplyAuthorizer()
	assert.NoError(t, err)

	// testing here happens within http server block
	zi, err := client.Clusters().ListZones()
	assert.NotNil(t, zi)
	assert.NoError(t, err)
	assert.Len(t, zi.Zones, 3)
}

// getAndAssertEnv fetches the env for testing and also asserts that the env value is not Zero i.e ""
func getAndAssertEnv(t *testing.T, key string) string {
	value, present := os.LookupEnv(key)
	assert.True(t, present, fmt.Sprintf("Env variable %s is not set", key))
	return value
}

// func TestAzureAccAuth_TestPatTokenDuration(t *testing.T) {
// 	if _, ok := os.LookupEnv("TF_ACC"); !ok {
// 		t.Skip("Acceptance tests skipped unless env 'TF_ACC' set")
// 	}

// 	client := DatabricksClient{
// 		AzureAuth: AzureAuth{
// 			ManagedResourceGroup: getAndAssertEnv(t, "DATABRICKS_AZURE_MANAGED_RESOURCE_GROUP"),
// 			AzureRegion:          getAndAssertEnv(t, "AZURE_REGION"),
// 			WorkspaceName:        getAndAssertEnv(t, "DATABRICKS_AZURE_WORKSPACE_NAME"),
// 			ResourceGroup:        getAndAssertEnv(t, "DATABRICKS_AZURE_RESOURCE_GROUP"),
// 			SubscriptionID:       getAndAssertEnv(t, "DATABRICKS_AZURE_SUBSCRIPTION_ID"),
// 			TenantID:             getAndAssertEnv(t, "DATABRICKS_AZURE_TENANT_ID"),
// 			ClientID:             getAndAssertEnv(t, "DATABRICKS_AZURE_CLIENT_ID"),
// 			ClientSecret:         getAndAssertEnv(t, "DATABRICKS_AZURE_CLIENT_SECRET"),
// 		},
// 	}
// 	err := client.Configure("dev-integration")
// 	assert.NoError(t, err, err)

// 	// Time in milliseconds
// 	tokenActualDuration := client.tokenExpiryTime - client.tokenCreateTime
// 	assert.Equal(t, patTokenSeconds, (time.Duration(tokenActualDuration) * time.Millisecond).Seconds(),
// 		"duration should be the same")
// }

func TestAzureAccAuthCreateApiToken(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	client := DatabricksClient{
		AzureAuth: AzureAuth{
			ManagedResourceGroup: getAndAssertEnv(t, "DATABRICKS_AZURE_MANAGED_RESOURCE_GROUP"),
			AzureRegion:          getAndAssertEnv(t, "AZURE_REGION"),
			WorkspaceName:        getAndAssertEnv(t, "DATABRICKS_AZURE_WORKSPACE_NAME"),
			ResourceGroup:        getAndAssertEnv(t, "DATABRICKS_AZURE_RESOURCE_GROUP"),
			SubscriptionID:       getAndAssertEnv(t, "DATABRICKS_AZURE_SUBSCRIPTION_ID"),
			TenantID:             getAndAssertEnv(t, "DATABRICKS_AZURE_TENANT_ID"),
			ClientID:             getAndAssertEnv(t, "DATABRICKS_AZURE_CLIENT_ID"),
			ClientSecret:         getAndAssertEnv(t, "DATABRICKS_AZURE_CLIENT_SECRET"),
		},
	}
	err := client.Configure("dev-integration")
	assert.NoError(t, err, err)
	instancePoolInfo, instancePoolErr := client.InstancePools().Create(model.InstancePool{
		InstancePoolName:                   "my_instance_pool",
		MinIdleInstances:                   0,
		MaxCapacity:                        10,
		NodeTypeID:                         "Standard_DS3_v2",
		IdleInstanceAutoTerminationMinutes: 20,
		PreloadedSparkVersions: []string{
			"6.3.x-scala2.11",
		},
	})
	defer func() {
		err := client.InstancePools().Delete(instancePoolInfo.InstancePoolID)
		assert.NoError(t, err, err)
	}()

	assert.NoError(t, instancePoolErr, instancePoolErr)
}
