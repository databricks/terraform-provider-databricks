package service

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMWSStorageConfigurations(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	acctId := os.Getenv("DATABRICKS_MWS_ACCT_ID")
	client := GetIntegrationMWSAPIClient()
	storageConfigsList, err := client.MWSStorageConfigurations().List(acctId)
	assert.NoError(t, err, err)
	t.Log(storageConfigsList)

	storageConfig, err := client.MWSStorageConfigurations().Create(acctId, "sri-mws-terraform-storage-root-bucket", "sri-root-s3-bucket")
	assert.NoError(t, err, err)

	myStorageConfig, err := client.MWSStorageConfigurations().Read(acctId, storageConfig.StorageConfigurationID)
	assert.NoError(t, err, err)
	t.Log(myStorageConfig.RootBucketInfo.BucketName)

	defer func() {
		err = client.MWSStorageConfigurations().Delete(acctId, storageConfig.StorageConfigurationID)
		assert.NoError(t, err, err)
	}()
}
