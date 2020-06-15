package databricks

import (
	"path/filepath"
	"strings"

	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceDBFSFileSync() *schema.Resource {
	return &schema.Resource{
		Create: resourceDBFSFileSyncCreate,
		Read:   resourceDBFSFileSyncRead,
		Delete: resourceDBFSFileSyncDelete,
		Update: resourceDBFSFileSyncUpdate,

		Schema: map[string]*schema.Schema{
			"src_path": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tgt_path": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"file_size": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"mkdirs": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"host": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"token": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func resourceDBFSFileSyncCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
	srcPath := d.Get("src_path").(string)
	tgtPath := d.Get("tgt_path").(string)
	mkdirs := d.Get("mkdirs").(bool)

	if mkdirs {
		parentDir := filepath.Dir(tgtPath)
		err := client.DBFS().Mkdirs(parentDir)
		if err != nil {
			return err
		}
	}

	apiClient := parseSchemaToDBAPIClient(d, client)
	err := client.DBFS().Copy(srcPath, tgtPath, apiClient, true)
	if err != nil {
		return err
	}

	d.SetId(srcPath + "|||" + tgtPath)
	err = d.Set("mkdirs", mkdirs)
	if err != nil {
		return err
	}
	err = d.Set("tgt_path", tgtPath)
	if err != nil {
		return err
	}
	err = d.Set("src_path", srcPath)
	if err != nil {
		return err
	}

	return resourceDBFSFileSyncRead(d, m)
}

func resourceDBFSFileSyncUpdate(d *schema.ResourceData, m interface{}) error {
	mkdirs := d.Get("mkdirs").(bool)

	err := d.Set("mkdirs", mkdirs)
	if err != nil {
		return err
	}
	return resourceDBFSFileSyncRead(d, m)
}

func resourceDBFSFileSyncRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
	srcPath := d.Get("src_path").(string)
	tgtPath := d.Get("tgt_path").(string)

	var srcAPIDBFSClient service.DBFSAPI
	srcAPICLient := parseSchemaToDBAPIClient(d, client)
	if srcAPICLient != nil {
		srcAPIDBFSClient = srcAPICLient.DBFS()
	} else {
		srcAPIDBFSClient = client.DBFS()
	}
	srcInfo, err := srcAPIDBFSClient.Status(srcPath)
	if err != nil {
		return err
	}

	tgtInfo, err := client.DBFS().Status(tgtPath)
	if err != nil {
		return err
	}
	err = d.Set("file_size", srcInfo.FileSize)
	if err != nil {
		return err
	}
	err = d.Set("file_size", tgtInfo.FileSize)
	return err
}

func resourceDBFSFileSyncDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
	id := strings.Split(d.Id(), "|||")[1]

	err := client.DBFS().Delete(id, false)
	return err
}

func parseSchemaToDBAPIClient(d *schema.ResourceData, currentClient *service.DBApiClient) *service.DBApiClient {
	host, hostOk := d.GetOk("host")
	token, tokenOk := d.GetOk("token")
	var config service.DBApiClientConfig
	if hostOk && tokenOk {
		config.Host = host.(string)
		config.Token = token.(string)
		config.UserAgent = currentClient.Config.UserAgent
		var dbClient service.DBApiClient
		dbClient.SetConfig(&config)
		return &dbClient
	}
	return nil
}
