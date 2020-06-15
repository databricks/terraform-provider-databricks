package databricks

import (
	"log"
	"strings"

	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceMWSStorageConfigurations() *schema.Resource {
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
	client := m.(*service.DBApiClient)
	storageConfigurationName := d.Get("storage_configuration_name").(string)
	bucketName := d.Get("bucket_name").(string)
	mwsAcctID := d.Get("account_id").(string)
	storageConfiguration, err := client.MWSStorageConfigurations().Create(mwsAcctID, storageConfigurationName, bucketName)
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
	client := m.(*service.DBApiClient)
	packagedMwsID, err := unpackMWSAccountID(id)
	if err != nil {
		return err
	}
	storageConifiguration, err := client.MWSStorageConfigurations().Read(packagedMwsID.MwsAcctID, packagedMwsID.ResourceID)
	if err != nil {
		if isE2StorageConfigurationsMissing(err.Error()) {
			log.Printf("Missing mws storage configurations with id: %s.", id)
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
	client := m.(*service.DBApiClient)
	packagedMwsID, err := unpackMWSAccountID(id)
	if err != nil {
		return err
	}
	err = client.MWSStorageConfigurations().Delete(packagedMwsID.MwsAcctID, packagedMwsID.ResourceID)
	return err
}

func isE2StorageConfigurationsMissing(errorMsg string) bool {
	return strings.Contains(errorMsg, "RESOURCE_DOES_NOT_EXIST")
}
