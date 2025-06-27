package storage

import (
	"fmt"
	"strings"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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

func getAzureDomain(client *common.DatabricksClient) string {
	domains := map[string]string{
		"PUBLIC":       "core.windows.net",
		"USGOVERNMENT": "core.usgovcloudapi.net",
		"CHINA":        "core.chinacloudapi.cn",
	}
	azureEnvironment := client.Config.Environment().AzureEnvironment.Name
	domain, ok := domains[strings.ToUpper(azureEnvironment)]
	if !ok {
		panic(fmt.Sprintf("Unknown Azure environment: '%s'", azureEnvironment))
	}
	return domain
}

// Source returns ABFSS URI backing the mount
func (m AzureADLSGen2Mount) Source(client *common.DatabricksClient) string {
	return fmt.Sprintf("abfss://%s@%s.dfs.%s%s", m.ContainerName, m.StorageAccountName, getAzureDomain(client), m.Directory)
}

func (m AzureADLSGen2Mount) Name() string {
	return m.ContainerName
}

func (m AzureADLSGen2Mount) ValidateAndApplyDefaults(d *schema.ResourceData, client *common.DatabricksClient) error {
	return nil
}

// Config returns mount configurations
func (m AzureADLSGen2Mount) Config(client *common.DatabricksClient) map[string]string {
	aadEndpoint := client.Config.Environment().AzureActiveDirectoryEndpoint()
	return map[string]string{
		"fs.azure.account.auth.type":                          "OAuth",
		"fs.azure.account.oauth.provider.type":                "org.apache.hadoop.fs.azurebfs.oauth2.ClientCredsTokenProvider",
		"fs.azure.account.oauth2.client.id":                   m.ClientID,
		"fs.azure.account.oauth2.client.secret":               fmt.Sprintf("{{secrets/%s/%s}}", m.SecretScope, m.SecretKey),
		"fs.azure.account.oauth2.client.endpoint":             fmt.Sprintf("%s/%s/oauth2/token", aadEndpoint, m.TenantID),
		"fs.azure.createRemoteFileSystemDuringInitialization": fmt.Sprintf("%t", m.InitializeFileSystem),
	}
}

// ResourceAzureAdlsGen2Mount creates the resource
func ResourceAzureAdlsGen2Mount() common.Resource {
	return deprecatedMountResource(commonMountResource(AzureADLSGen2Mount{}, map[string]*schema.Schema{
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
		"environment": {
			Type:         schema.TypeString,
			Optional:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringInSlice([]string{"PUBLIC", "USGOVERNMENT", "CHINA"}, false),
			Default:      "PUBLIC",
		},
	}))
}
