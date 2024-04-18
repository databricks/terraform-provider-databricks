package mws

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// NewStorageConfigurationsAPI creates MWSStorageConfigurationsAPI instance from provider meta
func NewStorageConfigurationsAPI(ctx context.Context, m any) StorageConfigurationsAPI {
	return StorageConfigurationsAPI{m.(*common.DatabricksClient), ctx}
}

// StorageConfigurationsAPI exposes the mws storageConfiguration API
type StorageConfigurationsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Create creates a configuration for the root s3 bucket
func (a StorageConfigurationsAPI) Create(mwsAcctID, storageConfigurationName string, bucketName string) (StorageConfiguration, error) {
	var mwsStorageConfigurations StorageConfiguration
	storageConfigurationAPIPath := fmt.Sprintf("/accounts/%s/storage-configurations", mwsAcctID)
	err := a.client.Post(a.context, storageConfigurationAPIPath, StorageConfiguration{
		StorageConfigurationName: storageConfigurationName,
		RootBucketInfo: &RootBucketInfo{
			BucketName: bucketName,
		},
	}, &mwsStorageConfigurations)
	return mwsStorageConfigurations, err
}

// Read returns the configuration for the root s3 bucket and metadata for the storage configuration
func (a StorageConfigurationsAPI) Read(mwsAcctID, storageConfigurationID string) (StorageConfiguration, error) {
	var mwsStorageConfigurations StorageConfiguration
	storageConfigurationAPIPath := fmt.Sprintf("/accounts/%s/storage-configurations/%s", mwsAcctID, storageConfigurationID)
	err := a.client.Get(a.context, storageConfigurationAPIPath, nil, &mwsStorageConfigurations)
	return mwsStorageConfigurations, err
}

// Delete deletes the configuration for the root s3 bucket
func (a StorageConfigurationsAPI) Delete(mwsAcctID, storageConfigurationID string) error {
	storageConfigurationAPIPath := fmt.Sprintf("/accounts/%s/storage-configurations/%s", mwsAcctID, storageConfigurationID)
	return a.client.Delete(a.context, storageConfigurationAPIPath, nil)
}

// List lists all the storage configurations for the root s3 buckets in the account ID provided to the client config
func (a StorageConfigurationsAPI) List(mwsAcctID string) ([]StorageConfiguration, error) {
	var mwsStorageConfigurationsList []StorageConfiguration
	storageConfigurationAPIPath := fmt.Sprintf("/accounts/%s/storage-configurations", mwsAcctID)
	err := a.client.Get(a.context, storageConfigurationAPIPath, nil, &mwsStorageConfigurationsList)
	return mwsStorageConfigurationsList, err
}

func ResourceMwsStorageConfigurations() common.Resource {
	p := common.NewPairSeparatedID("account_id", "storage_configuration_id", "/")
	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			name := d.Get("storage_configuration_name").(string)
			bucketName := d.Get("bucket_name").(string)
			accountID := d.Get("account_id").(string)
			storageConfiguration, err := NewStorageConfigurationsAPI(ctx, c).Create(accountID, name, bucketName)
			if err != nil {
				return err
			}
			d.Set("storage_configuration_id", storageConfiguration.StorageConfigurationID)
			p.Pack(d)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			accountID, storageID, err := p.Unpack(d)
			if err != nil {
				return err
			}
			storageConifiguration, err := NewStorageConfigurationsAPI(ctx, c).Read(accountID, storageID)
			if err != nil {
				return err
			}
			d.Set("storage_configuration_name", storageConifiguration.StorageConfigurationName)
			d.Set("bucket_name", storageConifiguration.RootBucketInfo.BucketName)
			return d.Set("creation_time", storageConifiguration.CreationTime)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			accountID, storageID, err := p.Unpack(d)
			if err != nil {
				return err
			}
			return NewStorageConfigurationsAPI(ctx, c).Delete(accountID, storageID)
		},
		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"storage_configuration_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"bucket_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"creation_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"storage_configuration_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
