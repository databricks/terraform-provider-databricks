package databricks

import (
	"fmt"
	"log"
	"strings"

	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAzureAdlsGen2Mount() *schema.Resource {
	return &schema.Resource{
		Create: resourceAzureAdlsGen2Create,
		Read:   resourceAzureAdlsGen2Read,
		Delete: resourceAzureAdlsGen2Delete,

		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"container_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"storage_account_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"directory": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: ValidateMountDirectory,
			},
			"mount_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tenant_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"client_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"client_secret_scope": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"client_secret_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"initialize_file_system": {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAzureAdlsGen2Create(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
	clusterID := d.Get("cluster_id").(string)
	err := changeClusterIntoRunningState(clusterID, client)
	if err != nil {
		return err
	}
	containerName := d.Get("container_name").(string)
	storageAccountName := d.Get("storage_account_name").(string)
	directory := d.Get("directory").(string)
	mountName := d.Get("mount_name").(string)
	tenantID := d.Get("tenant_id").(string)
	clientID := d.Get("client_id").(string)
	clientSecretScope := d.Get("client_secret_scope").(string)
	clientSecretKey := d.Get("client_secret_key").(string)
	initializeFileSystem := d.Get("initialize_file_system").(bool)

	adlsGen2Mount := NewAzureADLSGen2Mount(containerName, storageAccountName, directory, mountName, clientID, tenantID,
		clientSecretScope, clientSecretKey, initializeFileSystem)

	err = adlsGen2Mount.Create(client.Commands(), clusterID)
	if err != nil {
		return err
	}
	d.SetId(mountName)

	err = d.Set("cluster_id", clusterID)
	if err != nil {
		return err
	}
	err = d.Set("mount_name", mountName)
	if err != nil {
		return err
	}
	err = d.Set("tenant_id", tenantID)
	if err != nil {
		return err
	}
	err = d.Set("client_id", clientID)
	if err != nil {
		return err
	}
	err = d.Set("client_secret_scope", clientSecretScope)
	if err != nil {
		return err
	}
	err = d.Set("client_secret_key", clientSecretKey)
	if err != nil {
		return err
	}
	err = d.Set("initialize_file_system", initializeFileSystem)
	if err != nil {
		return err
	}

	return resourceAzureAdlsGen2Read(d, m)
}
func resourceAzureAdlsGen2Read(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
	clusterID := d.Get("cluster_id").(string)
	containerName := d.Get("container_name").(string)
	storageAccountName := d.Get("storage_account_name").(string)
	directory := d.Get("directory").(string)
	mountName := d.Get("mount_name").(string)
	tenantID := d.Get("tenant_id").(string)
	clientID := d.Get("client_id").(string)
	clientSecretScope := d.Get("client_secret_scope").(string)
	clientSecretKey := d.Get("client_secret_key").(string)
	initializeFileSystem := d.Get("initialize_file_system").(bool)

	err := changeClusterIntoRunningState(clusterID, client)
	if err != nil {
		if isClusterMissing(err.Error(), clusterID) {
			log.Printf("Unable to refresh mount '%s' as cluster '%s' is missing", mountName, clusterID)
			d.SetId("")
			return nil
		}
		return err
	}

	adlsGen2Mount := NewAzureADLSGen2Mount(containerName, storageAccountName, directory, mountName, clientID, tenantID,
		clientSecretScope, clientSecretKey, initializeFileSystem)

	url, err := adlsGen2Mount.Read(client.Commands(), clusterID)
	if err != nil {
		//Reset id in case of inability to find mount
		if strings.Contains(err.Error(), "Unable to find mount point!") ||
			strings.Contains(err.Error(), fmt.Sprintf("/mnt/%s does not exist.", mountName)) {
			d.SetId("")
			return nil
		}
		return err
	}
	container, storageAcc, dir, err := ProcessAzureWasbAbfssUris(url)
	if err != nil {
		return err
	}
	err = d.Set("container_name", container)
	if err != nil {
		return err
	}
	err = d.Set("storage_account_name", storageAcc)
	if err != nil {
		return err
	}
	err = d.Set("directory", dir)
	return err
}

func resourceAzureAdlsGen2Delete(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
	clusterID := d.Get("cluster_id").(string)
	err := changeClusterIntoRunningState(clusterID, client)
	if err != nil {
		return err
	}
	containerName := d.Get("container_name").(string)
	storageAccountName := d.Get("storage_account_name").(string)
	directory := d.Get("directory").(string)
	mountName := d.Get("mount_name").(string)
	tenantID := d.Get("tenant_id").(string)
	clientID := d.Get("client_id").(string)
	clientSecretScope := d.Get("client_secret_scope").(string)
	clientSecretKey := d.Get("client_secret_key").(string)
	initializeFileSystem := d.Get("initialize_file_system").(bool)

	adlsGen2Mount := NewAzureADLSGen2Mount(containerName, storageAccountName, directory, mountName, clientID, tenantID,
		clientSecretScope, clientSecretKey, initializeFileSystem)
	return adlsGen2Mount.Delete(client.Commands(), clusterID)
}
