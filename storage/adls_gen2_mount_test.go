package storage

import (
	"testing"

	"github.com/databrickslabs/databricks-terraform/internal/qa"
)

func TestAzureAccADLSv2Mount(t *testing.T) {
	client, mp := mountPointThroughReusedCluster(t)
	if !client.AzureAuth.IsClientSecretSet() {
		t.Skip("Test is meant only for client-secret conf Azure")
	}
	storageAccountName := qa.GetEnvOrSkipTest(t, "TEST_STORAGE_V2_ACCOUNT")
	container := qa.GetEnvOrSkipTest(t, "TEST_STORAGE_V2_ABFSS")
	testWithNewSecretScope(t, func(scope, key string) {
		testMounting(t, mp, AzureADLSGen2Mount{
			ClientID:             client.AzureAuth.ClientID,
			TenantID:             client.AzureAuth.TenantID,
			StorageAccountName:   storageAccountName,
			ContainerName:        container,
			SecretScope:          scope,
			SecretKey:            key,
			InitializeFileSystem: true,
			Directory:            "/",
		})
	}, client, mp.name, client.AzureAuth.ClientSecret)
}
