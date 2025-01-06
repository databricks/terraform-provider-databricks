package storage

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// TODO: add support for encryption parameters in S3
type GenericMount struct {
	URI     string                     `json:"uri,omitempty" tf:"force_new"`
	Options map[string]string          `json:"extra_configs,omitempty" tf:"force_new"`
	Abfs    *AzureADLSGen2MountGeneric `json:"abfs,omitempty" tf:"force_new,suppress_diff"`
	S3      *S3IamMount                `json:"s3,omitempty" tf:"force_new,suppress_diff"`
	Adl     *AzureADLSGen1MountGeneric `json:"adl,omitempty" tf:"force_new,suppress_diff"`
	Wasb    *AzureBlobMountGeneric     `json:"wasb,omitempty" tf:"force_new,suppress_diff"`
	Gs      *GSMount                   `json:"gs,omitempty" tf:"force_new,suppress_diff"`

	ClusterID      string `json:"cluster_id,omitempty" tf:"computed,force_new"`
	MountName      string `json:"name,omitempty" tf:"computed,force_new"`
	ResourceID     string `json:"resource_id,omitempty" tf:"force_new"`
	EncryptionType string `json:"encryption_type,omitempty" tf:"force_new"`
}

func (m GenericMount) getBlock() Mount {
	if m.Abfs != nil {
		return m.Abfs
	} else if m.Adl != nil {
		return m.Adl
	} else if m.Wasb != nil {
		return m.Wasb
	} else if m.S3 != nil {
		return m.S3
	} else if m.Gs != nil {
		return m.Gs
	}
	return nil
}

// Source returns URI backing the mount
func (m GenericMount) Source(client *common.DatabricksClient) string {
	if block := m.getBlock(); block != nil {
		return block.Source(client)
	}
	return m.URI
}

// Name
func (m GenericMount) Name() string {
	if m.MountName != "" {
		return m.MountName
	}
	if block := m.getBlock(); block != nil {
		return block.Name()
	}
	return ""
}

// Config returns mount configurations
func (m GenericMount) Config(client *common.DatabricksClient) map[string]string {
	if block := m.getBlock(); block != nil {
		return block.Config(client)
	}
	return m.Options
}

// ApplyDefaults tries to apply defaults to a given resource
func (m GenericMount) ValidateAndApplyDefaults(d *schema.ResourceData, client *common.DatabricksClient) error {
	if block := m.getBlock(); block != nil {
		return block.ValidateAndApplyDefaults(d, client)
	}
	if _, ok := d.GetOk("name"); !ok {
		return fmt.Errorf("value of name is not specified or empty")
	}
	if _, ok := d.GetOk("uri"); !ok {
		return fmt.Errorf("value of uri is not specified or empty")
	}
	return nil
}

// --------------- Generic ADLSgen2

func parseStorageContainerId(rid string) (string, string, error) {
	const containerRegex = `(?i)subscriptions/([^/]+)/resourceGroups/([^/]+)/providers/Microsoft.Storage/storageAccounts/([^/]+)/blobServices/default/containers/(.+)`
	containerPattern := regexp.MustCompile(containerRegex)
	match := containerPattern.FindStringSubmatch(rid)

	if len(match) == 0 {
		return "", "", fmt.Errorf("parsing failed for %s. Invalid container resource Id format", rid)
	}

	return match[3], match[4], nil
}

func getContainerDefaults(d *schema.ResourceData) (string, string, error) {
	rid := d.Get("resource_id").(string)
	if rid != "" {
		acc, cont, err := parseStorageContainerId(rid)
		return acc, cont, err
	}
	return "", "", fmt.Errorf("container_name or storage_account_name are empty, and resource_id or uri aren't specified")
}

func getTenantID(client *common.DatabricksClient) (string, error) {
	if client.Config.AzureTenantID != "" {
		return client.Config.AzureTenantID, nil
	}
	v, err := client.GetAzureJwtProperty("tid")
	if err != nil {
		return "", err
	}
	tid := strings.TrimSpace(v.(string))
	if tid != "" {
		return tid, nil
	}

	return "", fmt.Errorf("tenant_id isn't provided & we can't detect it")
}

// AzureADLSGen2Mount describes the object for a azure datalake gen 2 storage mount
type AzureADLSGen2MountGeneric struct {
	ContainerName        string `json:"container_name,omitempty" tf:"computed,force_new"`
	StorageAccountName   string `json:"storage_account_name,omitempty" tf:"computed,force_new"`
	Directory            string `json:"directory,omitempty" tf:"force_new"`
	ClientID             string `json:"client_id" tf:"force_new"`
	TenantID             string `json:"tenant_id,omitempty" tf:"computed,force_new"`
	SecretScope          string `json:"client_secret_scope" tf:"force_new"`
	SecretKey            string `json:"client_secret_key" tf:"force_new"`
	InitializeFileSystem bool   `json:"initialize_file_system" tf:"force_new"`
}

// Source returns ABFSS URI backing the mount
func (m *AzureADLSGen2MountGeneric) Source(client *common.DatabricksClient) string {
	return fmt.Sprintf("abfss://%s@%s.dfs.%s%s", m.ContainerName, m.StorageAccountName, getAzureDomain(client), m.Directory)
}

func (m *AzureADLSGen2MountGeneric) Name() string {
	return m.ContainerName
}

func (m *AzureADLSGen2MountGeneric) ValidateAndApplyDefaults(d *schema.ResourceData, client *common.DatabricksClient) error {
	if m.ContainerName == "" || m.StorageAccountName == "" {
		acc, cont, err := getContainerDefaults(d)
		if err != nil {
			return err
		}
		m.ContainerName = cont
		m.StorageAccountName = acc
	}
	nm := d.Get("name").(string)
	if nm == "" {
		d.Set("name", m.Name())
	}
	if m.TenantID == "" {
		tenant_id, err := getTenantID(client)
		if err != nil {
			return fmt.Errorf("tenant_id is not defined, and we can't extract it: %w", err)
		}
		log.Printf("[DEBUG] Got tenant_id: %s", tenant_id)
		m.TenantID = tenant_id
	}
	return nil
}

// Config returns mount configurations
func (m *AzureADLSGen2MountGeneric) Config(client *common.DatabricksClient) map[string]string {
	aadEndpoint := client.Config.Environment().AzureActiveDirectoryEndpoint()
	return map[string]string{
		"fs.azure.account.auth.type":                          "OAuth",
		"fs.azure.account.oauth.provider.type":                "org.apache.hadoop.fs.azurebfs.oauth2.ClientCredsTokenProvider",
		"fs.azure.account.oauth2.client.id":                   m.ClientID,
		"fs.azure.account.oauth2.client.secret":               fmt.Sprintf("{{secrets/%s/%s}}", m.SecretScope, m.SecretKey),
		"fs.azure.account.oauth2.client.endpoint":             fmt.Sprintf("%s%s/oauth2/token", aadEndpoint, m.TenantID),
		"fs.azure.createRemoteFileSystemDuringInitialization": fmt.Sprintf("%t", m.InitializeFileSystem),
	}
}

// --------------- Generic ADLSgen1

// AzureADLSGen1Mount describes the object for a azure datalake gen 1 storage mount
type AzureADLSGen1MountGeneric struct {
	StorageResource string `json:"storage_resource_name,omitempty" tf:"computed,force_new"`
	Directory       string `json:"directory,omitempty" tf:"force_new"`
	PrefixType      string `json:"spark_conf_prefix,omitempty" tf:"default:fs.adl,force_new"`
	ClientID        string `json:"client_id" tf:"force_new"`
	TenantID        string `json:"tenant_id,omitempty" tf:"computed,force_new"`
	SecretScope     string `json:"client_secret_scope" tf:"force_new"`
	SecretKey       string `json:"client_secret_key" tf:"force_new"`
}

// Source ...
func (m *AzureADLSGen1MountGeneric) Source(_ *common.DatabricksClient) string {
	return fmt.Sprintf("adl://%s.azuredatalakestore.net%s", m.StorageResource, m.Directory)
}

func (m *AzureADLSGen1MountGeneric) Name() string {
	return m.StorageResource
}

func (m *AzureADLSGen1MountGeneric) ValidateAndApplyDefaults(d *schema.ResourceData, client *common.DatabricksClient) error {
	rid := d.Get("resource_id").(string)
	if m.StorageResource == "" {
		if rid != "" {
			res, err := parseAzureResourceID(rid)
			if err != nil {
				return err
			}
			if res.resourceType != "accounts" || res.provider != "Microsoft.DataLakeStore" {
				return fmt.Errorf("incorrect resource type or provider in resource_id: %s", rid)
			}
			m.StorageResource = res.resourceName
		} else {
			return fmt.Errorf("storage_resource_name is empty, and resource_id or uri aren't specified")
		}
	}
	nm := d.Get("name").(string)
	if nm == "" {
		d.Set("name", m.Name())
	}
	if m.TenantID == "" {
		tenant_id, err := getTenantID(client)
		if err != nil {
			return fmt.Errorf("tenant_id is not defined, and we can't extract it: %w", err)
		}
		m.TenantID = tenant_id
	}
	return nil
}

// Config ...
func (m *AzureADLSGen1MountGeneric) Config(client *common.DatabricksClient) map[string]string {
	aadEndpoint := client.Config.Environment().AzureActiveDirectoryEndpoint()
	return map[string]string{
		m.PrefixType + ".oauth2.access.token.provider.type": "ClientCredential",
		m.PrefixType + ".oauth2.client.id":                  m.ClientID,
		m.PrefixType + ".oauth2.credential":                 fmt.Sprintf("{{secrets/%s/%s}}", m.SecretScope, m.SecretKey),
		m.PrefixType + ".oauth2.refresh.url":                fmt.Sprintf("%s%s/oauth2/token", aadEndpoint, m.TenantID),
	}
}

// --------------- Generic Azure Blob Storage

// AzureBlobMount describes the object for a azure blob storage mount - a.k.a. NativeAzureFileSystem
type AzureBlobMountGeneric struct {
	ContainerName      string `json:"container_name,omitempty" tf:"computed,force_new"`
	StorageAccountName string `json:"storage_account_name,omitempty" tf:"computed,force_new"`
	Directory          string `json:"directory,omitempty" tf:"force_new"`
	AuthType           string `json:"auth_type" tf:"force_new"`
	SecretScope        string `json:"token_secret_scope" tf:"force_new"`
	SecretKey          string `json:"token_secret_key" tf:"force_new"`
}

// Source ...
func (m *AzureBlobMountGeneric) Source(client *common.DatabricksClient) string {
	return fmt.Sprintf("wasbs://%[1]s@%[2]s.blob.%[3]s%[4]s",
		m.ContainerName, m.StorageAccountName, getAzureDomain(client), m.Directory)
}

func (m *AzureBlobMountGeneric) Name() string {
	return m.ContainerName
}

func (m *AzureBlobMountGeneric) ValidateAndApplyDefaults(d *schema.ResourceData, client *common.DatabricksClient) error {
	if m.ContainerName == "" || m.StorageAccountName == "" {
		acc, cont, err := getContainerDefaults(d)
		if err != nil {
			return err
		}
		m.ContainerName = cont
		m.StorageAccountName = acc
	}
	nm := d.Get("name").(string)
	if nm == "" {
		d.Set("name", m.Name())
	}

	return nil
}

// Config ...
func (m *AzureBlobMountGeneric) Config(client *common.DatabricksClient) map[string]string {
	var confKey string
	if m.AuthType == "SAS" {
		confKey = fmt.Sprintf("fs.azure.sas.%s.%s.blob.core.windows.net", m.ContainerName, m.StorageAccountName)
	} else {
		confKey = fmt.Sprintf("fs.azure.account.key.%s.blob.core.windows.net", m.StorageAccountName)
	}
	return map[string]string{
		confKey: fmt.Sprintf("{{secrets/%s/%s}}", m.SecretScope, m.SecretKey),
	}
}

type resourceDetails struct {
	provider     string
	resourceType string
	resourceName string
}

var azResRE = regexp.MustCompile(`(?i)subscriptions/[^/]+/resourceGroups/[^/]+/providers/([^/]+?)/([^/]+?)/(.+)`)

// simplified parsing of azure resource IDs
// See https://learn.microsoft.com/en-us/azure/azure-resource-manager/templates/template-functions-resource?tabs=json#resourceid
func parseAzureResourceID(resourceID string) (*resourceDetails, error) {
	m := azResRE.FindStringSubmatch(resourceID)
	if len(m) == 0 {
		return nil, fmt.Errorf("invalid azure resource id: %s", resourceID)
	}
	v := strings.Split(m[3], "/")
	return &resourceDetails{
		provider:     m[1],
		resourceType: m[2],
		resourceName: v[len(v)-1],
	}, nil
}
