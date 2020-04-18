package db

import (
	clientLib "github.com/databrickslabs/databricks-terraform/client"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"path/filepath"
	"strings"
)

func resourceDBFSFileSync() *schema.Resource {
	return &schema.Resource{
		Create: resourceDBFSFileSyncCreate,
		Read:   resourceDBFSFileSyncRead,
		Delete: resourceDBFSFileSyncDelete,
		Update: resourceDBFSFileSyncUpdate,

		Schema: map[string]*schema.Schema{
			"src_path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tgt_path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"file_size": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"mkdirs": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"token": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func resourceDBFSFileSyncCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(service.DBApiClient)
	srcPath := d.Get("src_path").(string)
	tgtPath := d.Get("tgt_path").(string)
	mkdirs := d.Get("mkdirs").(bool)

	if mkdirs == true {
		parentDir := filepath.Dir(tgtPath)
		err := client.DBFS().Mkdirs(parentDir)
		if err != nil {
			return err
		}
	}

	apiClient := parseSchemaToDBAPIClient(d)
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
	client := m.(service.DBApiClient)
	srcPath := d.Get("src_path").(string)
	tgtPath := d.Get("tgt_path").(string)

	var srcAPIDBFSClient service.DBFSAPI
	srcAPICLient := parseSchemaToDBAPIClient(d)
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
	client := m.(service.DBApiClient)
	id := strings.Split(d.Id(), "|||")[1]

	err := client.DBFS().Delete(id, false)
	return err
}

func parseSchemaToDBAPIClient(d *schema.ResourceData) *service.DBApiClient {
	host, hostOk := d.GetOk("host")
	token, tokenOk := d.GetOk("token")
	var option clientLib.DBApiClientConfig
	if hostOk && tokenOk {
		option.Host = host.(string)
		option.Token = token.(string)
		var dbClient service.DBApiClient
		dbClient.SetConfig(&option)
		return &dbClient
	}
	return nil
}
