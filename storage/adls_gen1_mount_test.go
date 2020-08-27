package storage

import (
	"testing"

	"github.com/databrickslabs/databricks-terraform/internal/qa"
)

func TestAzureAccADLSv1Mount(t *testing.T) {
	client, mp := mountPointThroughReusedCluster(t)
	if !client.AzureAuth.IsClientSecretSet() {
		t.Skip("Test is meant only for client-secret conf Azure")
	}
	storageResource := qa.GetEnvOrSkipTest(t, "TEST_DATA_LAKE_STORE_NAME")
	testWithNewSecretScope(t, func(scope, key string) {
		testMounting(t, mp, AzureADLSGen1Mount{
			ClientID:        client.AzureAuth.ClientID,
			TenantID:        client.AzureAuth.TenantID,
			PrefixType:      "dfs.adls",
			StorageResource: storageResource,
			Directory:       "/",
			SecretScope:     scope,
			SecretKey:       key,
		})
	}, client, mp.name, client.AzureAuth.ClientSecret)
}
