package service

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMwsAccStorageConfigurations(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	acctID := os.Getenv("DATABRICKS_ACCOUNT_ID")
	client := CommonEnvironmentClient()
	storageConfigsList, err := client.MWSStorageConfigurations().List(acctID)
	assert.NoError(t, err, err)
	t.Log(storageConfigsList)

	storageConfig, err := client.MWSStorageConfigurations().Create(acctID, "sri-mws-terraform-storage-root-bucket", "sri-root-s3-bucket")
	assert.NoError(t, err, err)

	myStorageConfig, err := client.MWSStorageConfigurations().Read(acctID, storageConfig.StorageConfigurationID)
	assert.NoError(t, err, err)
	t.Log(myStorageConfig.RootBucketInfo.BucketName)

	defer func() {
		err = client.MWSStorageConfigurations().Delete(acctID, storageConfig.StorageConfigurationID)
		assert.NoError(t, err, err)
	}()
}
