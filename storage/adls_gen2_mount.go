package storage

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// AzureADLSGen2Mount describes the object for a azure datalake gen 2 storage mount
type AzureADLSGen2Mount struct {
	ContainerName        string `json:"container_name"`
	StorageAccountName   string `json:"storage_account_name"`
	Directory            string `json:"directory,omitempty"`
	ClientID             string `json:"client_id"`
	TenantID             string `json:"tenant_id"`
	SecretScope          string `json:"client_secret_scope"`
	SecretKey            string `json:"client_secret_key"`
	InitializeFileSystem bool   `json:"initialize_file_system"`
}

// Source returns ABFSS URI backing the mount
func (m AzureADLSGen2Mount) Source() string {
	return fmt.Sprintf("abfss://%s@%s.dfs.core.windows.net%s",
		m.ContainerName, m.StorageAccountName, m.Directory)
}

// Config returns mount configurations
func (m AzureADLSGen2Mount) Config() map[string]string {
	return map[string]string{
		"fs.azure.account.auth.type":                          "OAuth",
		"fs.azure.account.oauth.provider.type":                "org.apache.hadoop.fs.azurebfs.oauth2.ClientCredsTokenProvider",
		"fs.azure.account.oauth2.client.id":                   m.ClientID,
		"fs.azure.account.oauth2.client.secret":               fmt.Sprintf("{secrets/%s/%s}", m.SecretScope, m.SecretKey),
		"fs.azure.account.oauth2.client.endpoint":             fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/token", m.TenantID),
		"fs.azure.createRemoteFileSystemDuringInitialization": fmt.Sprintf("%t", m.InitializeFileSystem),
	}
}

func ResourceAzureAdlsGen2Mount() *schema.Resource {
	return commonMountResource(AzureADLSGen2Mount{}, map[string]*schema.Schema{
		"cluster_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		"source": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"mount_name": {
			// TODO: have it by default as storage_resource_name
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
	})
}
