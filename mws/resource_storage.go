package mws

import (
	"context"
	"fmt"
	"log"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// NewStorageConfigurationsAPI creates MWSStorageConfigurationsAPI instance from provider meta
func NewStorageConfigurationsAPI(m interface{}) StorageConfigurationsAPI {
	return StorageConfigurationsAPI{m.(*common.DatabricksClient), context.TODO()}
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

func ResourceStorageConfiguration() *schema.Resource {
	return &schema.Resource{
		Create: resourceMWSStorageConfigurationsCreate,
		Read:   resourceMWSStorageConfigurationsRead,
		Delete: resourceMWSStorageConfigurationsDelete,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
				ForceNew:  true,
			},
			"storage_configuration_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"bucket_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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

func resourceMWSStorageConfigurationsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*common.DatabricksClient)
	storageConfigurationName := d.Get("storage_configuration_name").(string)
	bucketName := d.Get("bucket_name").(string)
	mwsAcctID := d.Get("account_id").(string)
	storageConfiguration, err := NewStorageConfigurationsAPI(client).Create(mwsAcctID, storageConfigurationName, bucketName)
	if err != nil {
		return err
	}
	storageConfigurationResourceID := PackagedMWSIds{
		MwsAcctID:  mwsAcctID,
		ResourceID: storageConfiguration.StorageConfigurationID,
	}
	d.SetId(packMWSAccountID(storageConfigurationResourceID))
	return resourceMWSStorageConfigurationsRead(d, m)
}

func resourceMWSStorageConfigurationsRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*common.DatabricksClient)
	packagedMwsID, err := UnpackMWSAccountID(id)
	if err != nil {
		return err
	}
	storageConifiguration, err := NewStorageConfigurationsAPI(client).Read(packagedMwsID.MwsAcctID, packagedMwsID.ResourceID)
	if err != nil {
		if e, ok := err.(common.APIError); ok && e.IsMissing() {
			log.Printf("missing resource due to error: %v\n", e)
			d.SetId("")
			return nil
		}
		return err
	}
	err = d.Set("storage_configuration_name", storageConifiguration.StorageConfigurationName)
	if err != nil {
		return err
	}
	err = d.Set("bucket_name", storageConifiguration.RootBucketInfo.BucketName)
	if err != nil {
		return err
	}
	err = d.Set("account_id", storageConifiguration.AccountID)
	if err != nil {
		return err
	}
	err = d.Set("creation_time", storageConifiguration.CreationTime)
	if err != nil {
		return err
	}
	err = d.Set("storage_configuration_id", storageConifiguration.StorageConfigurationID)
	if err != nil {
		return err
	}

	return nil
}

func resourceMWSStorageConfigurationsDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*common.DatabricksClient)
	packagedMwsID, err := UnpackMWSAccountID(id)
	if err != nil {
		return err
	}
	err = NewStorageConfigurationsAPI(client).Delete(packagedMwsID.MwsAcctID, packagedMwsID.ResourceID)
	return err
}
