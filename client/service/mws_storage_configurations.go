package service

import (
	"fmt"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// MWSStorageConfigurationsAPI exposes the mws storageConfiguration API
type MWSStorageConfigurationsAPI struct {
	client *DatabricksClient
}

// Create creates a configuration for the root s3 bucket
func (a MWSStorageConfigurationsAPI) Create(mwsAcctId, storageConfigurationName string, bucketName string) (model.MWSStorageConfigurations, error) {
	var mwsStorageConfigurations model.MWSStorageConfigurations
	storageConfigurationAPIPath := fmt.Sprintf("/accounts/%s/storage-configurations", mwsAcctId)
	err := a.client.post(storageConfigurationAPIPath, model.MWSStorageConfigurations{
		StorageConfigurationName: storageConfigurationName,
		RootBucketInfo: &model.RootBucketInfo{
			BucketName: bucketName,
		},
	}, &mwsStorageConfigurations)
	return mwsStorageConfigurations, err
}

// Read returns the configuration for the root s3 bucket and metadata for the storage configuration
func (a MWSStorageConfigurationsAPI) Read(mwsAcctId, storageConfigurationID string) (model.MWSStorageConfigurations, error) {
	var mwsStorageConfigurations model.MWSStorageConfigurations
	storageConfigurationAPIPath := fmt.Sprintf("/accounts/%s/storage-configurations/%s", mwsAcctId, storageConfigurationID)
	err := a.client.get(storageConfigurationAPIPath, nil, &mwsStorageConfigurations)
	return mwsStorageConfigurations, err
}

// Delete deletes the configuration for the root s3 bucket
func (a MWSStorageConfigurationsAPI) Delete(mwsAcctId, storageConfigurationID string) error {
	storageConfigurationAPIPath := fmt.Sprintf("/accounts/%s/storage-configurations/%s", mwsAcctId, storageConfigurationID)
	return a.client.delete(storageConfigurationAPIPath, nil)
}

// List lists all the storage configurations for the root s3 buckets in the E2 account ID provided to the client config
func (a MWSStorageConfigurationsAPI) List(mwsAcctId string) ([]model.MWSStorageConfigurations, error) {
	var mwsStorageConfigurationsList []model.MWSStorageConfigurations
	storageConfigurationAPIPath := fmt.Sprintf("/accounts/%s/storage-configurations", mwsAcctId)
	err := a.client.get(storageConfigurationAPIPath, nil, &mwsStorageConfigurationsList)
	return mwsStorageConfigurationsList, err
}
