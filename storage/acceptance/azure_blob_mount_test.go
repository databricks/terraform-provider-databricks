package acceptance

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/storage"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
)

func TestAzureAccBlobMount(t *testing.T) {
	client, mp := mountPointThroughReusedCluster(t)
	storageAccountName := qa.GetEnvOrSkipTest(t, "TEST_STORAGE_V2_ACCOUNT")
	accountKey := qa.GetEnvOrSkipTest(t, "TEST_STORAGE_V2_KEY")
	container := qa.GetEnvOrSkipTest(t, "TEST_STORAGE_V2_WASBS")
	testWithNewSecretScope(t, func(scope, key string) {
		testMounting(t, mp, storage.AzureBlobMount{
			StorageAccountName: storageAccountName,
			ContainerName:      container,
			SecretScope:        scope,
			SecretKey:          key,
			Directory:          "/",
		})
	}, client, mp.Name, accountKey)
}
