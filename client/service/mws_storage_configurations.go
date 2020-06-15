package service

import (
	"encoding/json"
	"fmt"

	"github.com/databrickslabs/databricks-terraform/client/model"

	"net/http"
)

// MWSStorageConfigurationsAPI exposes the mws storageConfiguration API
type MWSStorageConfigurationsAPI struct {
	Client *DBApiClient
}

// Create creates a configuration for the root s3 bucket
func (a MWSStorageConfigurationsAPI) Create(mwsAcctId, storageConfigurationName string, bucketName string) (model.MWSStorageConfigurations, error) {
	var mwsStorageConfigurations model.MWSStorageConfigurations

	storageConfigurationAPIPath := fmt.Sprintf("/accounts/%s/storage-configurations", mwsAcctId)

	mwsStorageConfigurationsRequest := model.MWSStorageConfigurations{
		StorageConfigurationName: storageConfigurationName,
		RootBucketInfo: &model.RootBucketInfo{
			BucketName: bucketName,
		},
	}

	resp, err := a.Client.performQuery(http.MethodPost, storageConfigurationAPIPath, "2.0", nil, mwsStorageConfigurationsRequest, nil)
	if err != nil {
		return mwsStorageConfigurations, err
	}

	err = json.Unmarshal(resp, &mwsStorageConfigurations)
	return mwsStorageConfigurations, err
}

// Read returns the configuration for the root s3 bucket and metadata for the storage configuration
func (a MWSStorageConfigurationsAPI) Read(mwsAcctId, storageConfigurationID string) (model.MWSStorageConfigurations, error) {
	var mwsStorageConfigurations model.MWSStorageConfigurations

	storageConfigurationAPIPath := fmt.Sprintf("/accounts/%s/storage-configurations/%s", mwsAcctId, storageConfigurationID)

	resp, err := a.Client.performQuery(http.MethodGet, storageConfigurationAPIPath, "2.0", nil, nil, nil)
	if err != nil {
		return mwsStorageConfigurations, err
	}

	err = json.Unmarshal(resp, &mwsStorageConfigurations)
	return mwsStorageConfigurations, err
}

// Delete deletes the configuration for the root s3 bucket
func (a MWSStorageConfigurationsAPI) Delete(mwsAcctId, storageConfigurationID string) error {
	storageConfigurationAPIPath := fmt.Sprintf("/accounts/%s/storage-configurations/%s", mwsAcctId, storageConfigurationID)

	_, err := a.Client.performQuery(http.MethodDelete, storageConfigurationAPIPath, "2.0", nil, nil, nil)

	return err
}

// List lists all the storage configurations for the root s3 buckets in the E2 account ID provided to the client config
func (a MWSStorageConfigurationsAPI) List(mwsAcctId string) ([]model.MWSStorageConfigurations, error) {
	var mwsStorageConfigurationsList []model.MWSStorageConfigurations

	storageConfigurationAPIPath := fmt.Sprintf("/accounts/%s/storage-configurations", mwsAcctId)

	resp, err := a.Client.performQuery(http.MethodGet, storageConfigurationAPIPath, "2.0", nil, nil, nil)
	if err != nil {
		return mwsStorageConfigurationsList, err
	}

	err = json.Unmarshal(resp, &mwsStorageConfigurationsList)
	return mwsStorageConfigurationsList, err
}
