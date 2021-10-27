package storage

import (
	"context"
	"crypto/md5"
	"fmt"
	"net/url"
	"strings"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/compute"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// TODO: add support for encryption parameters in S3
type GenericMount struct {
	URI     string                     `json:"uri,omitempty" tf:"force_new"`
	Options map[string]string          `json:"extra_configs,omitempty" tf:"force_new"`
	Abfs    *AzureADLSGen2MountGeneric `json:"abfs,omitempty" tf:"force_new"`
	S3      *S3IamMount                `json:"s3,omitempty" tf:"force_new"`
	Adl     *AzureADLSGen1MountGeneric `json:"adl,omitempty" tf:"force_new"`
	Wasb    *AzureBlobMountGeneric     `json:"wasb,omitempty" tf:"force_new"`
	Gs      *GSMount                   `json:"gs,omitempty" tf:"force_new"`

	ClusterID      string `json:"cluster_id,omitempty" tf:"computed,force_new"`
	MountName      string `json:"mount_name,omitempty" tf:"computed,force_new"`
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
func (m GenericMount) Source() string {
	if block := m.getBlock(); block != nil {
		return block.Source()
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
	if _, ok := d.GetOk("mount_name"); !ok {
		return fmt.Errorf("value of mount_name is not specified or empty")
	}
	if _, ok := d.GetOk("uri"); !ok {
		return fmt.Errorf("value of uri is not specified or empty")
	}
	return nil
}

// --------------- Google Cloud Storage

// GSMount describes the object for a GS mount using google service account
type GSMount struct {
	BucketName     string `json:"bucket_name" tf:"force_new"`
	ServiceAccount string `json:"service_account,omitempty" tf:"force_new"`
}

// Source ...
func (m GSMount) Source() string {
	return fmt.Sprintf("gs://%s", m.BucketName)
}

func (m GSMount) Name() string {
	return m.BucketName
}

func (m GSMount) ValidateAndApplyDefaults(d *schema.ResourceData, client *common.DatabricksClient) error {
	nm := d.Get("mount_name").(string)
	if nm != "" {
		return nil
	}
	nm = m.Name()
	if nm != "" {
		d.Set("mount_name", nm)
		return nil
	}
	nm = d.Get("resource_id").(string)
	if nm != "" {
		d.Set("mount_name", nm)
		return nil
	}
	uri := d.Get("uri").(string)
	if uri == "" {
		return fmt.Errorf("mount_name isn't specified, and no other data to infer it (bucket_name, uri or resource_id)")
	}
	u, err := url.Parse(uri)
	if err != nil {
		return fmt.Errorf("can't parse GCS URI")
	}
	if u.Scheme != "gs" {
		return fmt.Errorf("unsupported schema for GCS URI: %s", u.Scheme)
	}
	if u.Host == "" {
		return fmt.Errorf("no bucket name specified in URI")
	}
	d.Set("mount_name", u.Host)

	return nil
}

// Config ...
func (m GSMount) Config(client *common.DatabricksClient) map[string]string {
	return make(map[string]string) // return empty map so nil map does not marshal to null
}

func preprocessGsMount(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m interface{}) error {
	var gm GenericMount
	if err := common.DataToStructPointer(d, s, &gm); err != nil {
		return err
	}
	// TODO: move into Validate function
	if !(strings.HasPrefix(gm.URI, "gs://") || gm.Gs != nil) {
		return nil
	}
	clusterID := gm.ClusterID
	serviceAccount := ""
	if gm.Gs != nil {
		serviceAccount = gm.Gs.ServiceAccount
	}
	if clusterID == "" && serviceAccount == "" {
		return fmt.Errorf("either cluster_id or service_account must be specified to mount GCS bucket")
	}
	clustersAPI := compute.NewClustersAPI(ctx, m)
	if clusterID != "" {
		clusterInfo, err := clustersAPI.Get(clusterID)
		if err != nil {
			return err
		}
		if clusterInfo.GcpAttributes == nil {
			return fmt.Errorf("cluster %s must have GCP attributes", clusterID)
		}
		if len(clusterInfo.GcpAttributes.GoogleServiceAccount) == 0 {
			return fmt.Errorf("cluster %s must have GCP service account attached", clusterID)
		}
	}
	if serviceAccount != "" {
		cluster, err := GetOrCreateMountingClusterWithGcpServiceAccount(clustersAPI, serviceAccount)
		if err != nil {
			return err
		}
		return d.Set("cluster_id", cluster.ClusterID)
	}
	return nil
}

// GetOrCreateMountingClusterWithGcpServiceAccount ...
func GetOrCreateMountingClusterWithGcpServiceAccount(
	clustersAPI compute.ClustersAPI, serviceAccount string) (i compute.ClusterInfo, err error) {
	clusterName := fmt.Sprintf("terraform-mount-gcs-%x", md5.Sum([]byte(serviceAccount)))
	cluster := getCommonClusterObject(clustersAPI, clusterName)
	cluster.GcpAttributes = &compute.GcpAttributes{GoogleServiceAccount: serviceAccount}
	return clustersAPI.GetOrCreateRunningCluster(clusterName, cluster)
}

// --------------- Generic AWS S3

// S3IamMount describes the object for a aws mount using iam role
type S3IamMount struct {
	BucketName      string `json:"bucket_name" tf:"force_new"`
	InstanceProfile string `json:"instance_profile,omitempty" tf:"force_new"`
}

// Source ...
func (m S3IamMount) Source() string {
	return fmt.Sprintf("s3a://%s", m.BucketName)
}

// Name ...
func (m S3IamMount) Name() string {
	return m.BucketName
}

// Config ...
func (m S3IamMount) Config(client *common.DatabricksClient) map[string]string {
	return make(map[string]string) // return empty map so nil map does not marshal to null
}

func (m S3IamMount) ValidateAndApplyDefaults(d *schema.ResourceData, client *common.DatabricksClient) error {
	nm := d.Get("mount_name").(string)
	if nm != "" {
		return nil
	}
	nm = m.Name()
	if nm != "" {
		d.Set("mount_name", nm)
		return nil
	}
	nm = d.Get("resource_id").(string)
	if nm != "" {
		d.Set("mount_name", nm)
		return nil
	}
	uri := d.Get("uri").(string)
	if uri == "" {
		return fmt.Errorf("mount_name isn't specified, and no other data to infer it (bucket_name, uri or resource_id)")
	}
	u, err := url.Parse(uri)
	if err != nil {
		return fmt.Errorf("can't parse S3 URI")
	}
	if u.Scheme != "s3" && u.Scheme != "s3a" {
		return fmt.Errorf("unsupported schema for S3 URI: %s", u.Scheme)
	}
	if u.Host == "" {
		return fmt.Errorf("no bucket name specified in URI")
	}
	d.Set("mount_name", u.Host)

	return nil
}

func preprocessS3MountGeneric(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m interface{}) error {
	var gm GenericMount
	if err := common.DataToStructPointer(d, s, &gm); err != nil {
		return err
	}
	// TODO: move into Validate function
	if !(strings.HasPrefix(gm.URI, "s3://") || strings.HasPrefix(gm.URI, "s3a://") || gm.S3 != nil) {
		return nil
	}
	clusterID := gm.ClusterID
	instanceProfile := ""
	if gm.S3 != nil {
		instanceProfile = gm.S3.InstanceProfile
	}
	if clusterID == "" && instanceProfile == "" {
		return fmt.Errorf("either cluster_id or instance_profile must be specified to mount S3 bucket")
	}
	clustersAPI := compute.NewClustersAPI(ctx, m)
	if clusterID != "" {
		clusterInfo, err := clustersAPI.Get(clusterID)
		if err != nil {
			return err
		}
		if clusterInfo.AwsAttributes == nil {
			return fmt.Errorf("cluster %s must have AWS attributes", clusterID)
		}
		if len(clusterInfo.AwsAttributes.InstanceProfileArn) == 0 {
			return fmt.Errorf("cluster %s must have EC2 instance profile attached", clusterID)
		}
	}
	if instanceProfile != "" {
		cluster, err := GetOrCreateMountingClusterWithInstanceProfile(clustersAPI, instanceProfile)
		if err != nil {
			return err
		}
		return d.Set("cluster_id", cluster.ClusterID)
	}
	return nil
}

// --------------- Generic ADLSgen2

// AzureADLSGen2Mount describes the object for a azure datalake gen 2 storage mount
type AzureADLSGen2MountGeneric struct {
	ContainerName        string `json:"container_name,omitempty" tf:"computed,force_new"`
	StorageAccountName   string `json:"storage_account_name,omitempty" tf:"computed,force_new"`
	Directory            string `json:"directory,omitempty" tf:"force_new"`
	ClientID             string `json:"client_id,omitempty" tf:"computed,force_new"`
	TenantID             string `json:"tenant_id,omitempty" tf:"computed,force_new"`
	SecretScope          string `json:"client_secret_scope" tf:"force_new"`
	SecretKey            string `json:"client_secret_key" tf:"force_new"`
	InitializeFileSystem bool   `json:"initialize_file_system" tf:"force_new"`
}

// Source returns ABFSS URI backing the mount
func (m AzureADLSGen2MountGeneric) Source() string {
	return fmt.Sprintf("abfss://%s@%s.dfs.core.windows.net%s",
		m.ContainerName, m.StorageAccountName, m.Directory)
}

func (m AzureADLSGen2MountGeneric) Name() string {
	return fmt.Sprintf("%s-%s", m.StorageAccountName, m.ContainerName)
}

func (m AzureADLSGen2MountGeneric) ValidateAndApplyDefaults(d *schema.ResourceData, client *common.DatabricksClient) error {
	return nil
}

// Config returns mount configurations
func (m AzureADLSGen2MountGeneric) Config(client *common.DatabricksClient) map[string]string {
	aadEndpoint := client.AzureEnvironment.ActiveDirectoryEndpoint
	return map[string]string{
		"fs.azure.account.auth.type":                          "OAuth",
		"fs.azure.account.oauth.provider.type":                "org.apache.hadoop.fs.azurebfs.oauth2.ClientCredsTokenProvider",
		"fs.azure.account.oauth2.client.id":                   m.ClientID,
		"fs.azure.account.oauth2.client.secret":               fmt.Sprintf("{{secrets/%s/%s}}", m.SecretScope, m.SecretKey),
		"fs.azure.account.oauth2.client.endpoint":             fmt.Sprintf("%s/%s/oauth2/token", aadEndpoint, m.TenantID),
		"fs.azure.createRemoteFileSystemDuringInitialization": fmt.Sprintf("%t", m.InitializeFileSystem),
	}
}

// --------------- Generic ADLSgen1

// AzureADLSGen1Mount describes the object for a azure datalake gen 1 storage mount
type AzureADLSGen1MountGeneric struct {
	StorageResource string `json:"storage_resource_name,omitempty" tf:"computed,force_new"`
	Directory       string `json:"directory,omitempty" tf:"force_new"`
	PrefixType      string `json:"spark_conf_prefix,omitempty" tf:"default:fs.adl,force_new"`
	ClientID        string `json:"client_id,omitempty" tf:"computed,force_new"`
	TenantID        string `json:"tenant_id,omitempty" tf:"computed,force_new"`
	SecretScope     string `json:"client_secret_scope" tf:"force_new"`
	SecretKey       string `json:"client_secret_key" tf:"force_new"`
}

// Source ...
func (m AzureADLSGen1MountGeneric) Source() string {
	return fmt.Sprintf("adl://%s.azuredatalakestore.net%s", m.StorageResource, m.Directory)
}

func (m AzureADLSGen1MountGeneric) Name() string {
	return m.StorageResource
}

func (m AzureADLSGen1MountGeneric) ValidateAndApplyDefaults(d *schema.ResourceData, client *common.DatabricksClient) error {
	return nil
}

// Config ...
func (m AzureADLSGen1MountGeneric) Config(client *common.DatabricksClient) map[string]string {
	aadEndpoint := client.AzureEnvironment.ActiveDirectoryEndpoint
	return map[string]string{
		m.PrefixType + ".oauth2.access.token.provider.type": "ClientCredential",

		m.PrefixType + ".oauth2.client.id":   m.ClientID,
		m.PrefixType + ".oauth2.credential":  fmt.Sprintf("{{secrets/%s/%s}}", m.SecretScope, m.SecretKey),
		m.PrefixType + ".oauth2.refresh.url": fmt.Sprintf("%s/%s/oauth2/token", aadEndpoint, m.TenantID),
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
func (m AzureBlobMountGeneric) Source() string {
	return fmt.Sprintf("wasbs://%[1]s@%[2]s.blob.core.windows.net%[3]s",
		m.ContainerName, m.StorageAccountName, m.Directory)
}

func (m AzureBlobMountGeneric) Name() string {
	return fmt.Sprintf("%s-%s", m.StorageAccountName, m.ContainerName)
}

func (m AzureBlobMountGeneric) ValidateAndApplyDefaults(d *schema.ResourceData, client *common.DatabricksClient) error {
	return nil
}

// Config ...
func (m AzureBlobMountGeneric) Config(client *common.DatabricksClient) map[string]string {
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
