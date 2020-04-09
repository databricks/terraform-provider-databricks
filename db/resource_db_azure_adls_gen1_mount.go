package db

import (
	"fmt"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"strings"
)

func resourceAzureAdlsGen1Mount() *schema.Resource {
	return &schema.Resource{
		Create: resourceAzureAdlsGen1Create,
		Read:   resourceAzureAdlsGen1Read,
		Delete: resourceAzureAdlsGen1Delete,

		Schema: map[string]*schema.Schema{
			"cluster_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"storage_resource_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"spark_conf_prefix": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "fs.adl",
				ValidateFunc: validation.StringInSlice([]string{"fs.adl", "dfs.adls"}, false),
				ForceNew:     true,
			},
			"directory": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				//Default:  "/",
				ForceNew: true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errors []error) {
					directory := val.(string)
					if strings.HasPrefix(directory, "/") {
						return
					} else {
						errors = append(errors, fmt.Errorf("%s must start with /, got: %s", key, val))
					}
					return
				},
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
		},
	}
}

func resourceAzureAdlsGen1Create(d *schema.ResourceData, m interface{}) error {
	client := m.(service.DBApiClient)
	clusterId := d.Get("cluster_id").(string)
	err := changeClusterIntoRunningState(clusterId, client)
	if err != nil {
		return err
	}
	storageResourceName := d.Get("storage_resource_name").(string)
	sparkConfPrefix := d.Get("spark_conf_prefix").(string)
	directory := d.Get("directory").(string)
	mountName := d.Get("mount_name").(string)
	tenantId := d.Get("tenant_id").(string)
	clientId := d.Get("client_id").(string)
	clientSecretScope := d.Get("client_secret_scope").(string)
	clientSecretKey := d.Get("client_secret_key").(string)

	adlsGen1Mount := service.NewAzureADLSGen1Mount(storageResourceName, directory, mountName,
		sparkConfPrefix, clientId, tenantId, clientSecretScope, clientSecretKey)

	err = adlsGen1Mount.Create(client, clusterId)
	if err != nil {
		return err
	}
	d.SetId(mountName)

	err = d.Set("cluster_id", clusterId)
	if err != nil {
		return err
	}
	err = d.Set("mount_name", mountName)
	if err != nil {
		return err
	}
	err = d.Set("tenant_id", tenantId)
	if err != nil {
		return err
	}
	err = d.Set("client_id", clientId)
	if err != nil {
		return err
	}
	err = d.Set("spark_conf_prefix", sparkConfPrefix)
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

	return resourceAzureAdlsGen1Read(d, m)
}
func resourceAzureAdlsGen1Read(d *schema.ResourceData, m interface{}) error {
	client := m.(service.DBApiClient)
	clusterId := d.Get("cluster_id").(string)
	err := changeClusterIntoRunningState(clusterId, client)
	if err != nil {
		return err
	}
	storageResourceName := d.Get("storage_resource_name").(string)
	sparkConfPrefix := d.Get("spark_conf_prefix").(string)
	directory := d.Get("directory").(string)
	mountName := d.Get("mount_name").(string)
	tenantId := d.Get("tenant_id").(string)
	clientId := d.Get("client_id").(string)
	clientSecretScope := d.Get("client_secret_scope").(string)
	clientSecretKey := d.Get("client_secret_key").(string)

	adlsGen1Mount := service.NewAzureADLSGen1Mount(storageResourceName, directory, mountName,
		sparkConfPrefix, clientId, tenantId, clientSecretScope, clientSecretKey)

	url, err := adlsGen1Mount.Read(client, clusterId)
	if err != nil {
		//Reset id in case of inability to find mount
		if strings.Contains(err.Error(), "Unable to find mount point!") ||
			strings.Contains(err.Error(), fmt.Sprintf("/mnt/%s does not exist.", mountName)) {
			d.SetId("")
			return nil
		}
		return err
	}
	storageResourceName, dir, err := service.ProcessAzureAdlsGen1Uri(url)
	err = d.Set("storage_resource_name", storageResourceName)
	if err != nil {
		return err
	}
	err = d.Set("directory", dir)
	return err
}

func resourceAzureAdlsGen1Delete(d *schema.ResourceData, m interface{}) error {
	client := m.(service.DBApiClient)
	clusterId := d.Get("cluster_id").(string)
	err := changeClusterIntoRunningState(clusterId, client)
	if err != nil {
		return err
	}
	storageResourceName := d.Get("storage_resource_name").(string)
	sparkConfPrefix := d.Get("spark_conf_prefix").(string)
	directory := d.Get("directory").(string)
	mountName := d.Get("mount_name").(string)
	tenantId := d.Get("tenant_id").(string)
	clientId := d.Get("client_id").(string)
	clientSecretScope := d.Get("client_secret_scope").(string)
	clientSecretKey := d.Get("client_secret_key").(string)

	adlsGen1Mount := service.NewAzureADLSGen1Mount(storageResourceName, directory, mountName,
		sparkConfPrefix, clientId, tenantId, clientSecretScope, clientSecretKey)
	return adlsGen1Mount.Delete(client, clusterId)
}
